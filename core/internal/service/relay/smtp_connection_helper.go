package relay

import (
	"billionmail-core/api/relay/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/public"
	"context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"net"
	"net/smtp"
	"strings"
	"sync"
	"time"
)

// customPlainAuth Custom implementation of PLAIN authentication method, allowing authentication over non-encrypted connections.
type customPlainAuth struct {
	identity, username, password string
	host                         string
}

func (a *customPlainAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// Ignore connection status and host checks, force PLAIN authentication.
	resp := []byte{}
	resp = append(resp, []byte(a.identity)...)
	resp = append(resp, 0)
	resp = append(resp, []byte(a.username)...)
	resp = append(resp, 0)
	resp = append(resp, []byte(a.password)...)

	return "PLAIN", resp, nil
}

func (a *customPlainAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		// Should not be called.
		return nil, fmt.Errorf("unexpected server challenge")
	}
	return nil, nil
}

// NewCustomPlainAuth Creates a custom authentication method.
func NewCustomPlainAuth(identity, username, password, host string) smtp.Auth {
	return &customPlainAuth{
		identity: identity,
		username: username,
		password: password,
		host:     host,
	}
}

// Cache instance for storing SMTP connection states.
var smtpStatusCache = gcache.New()

// Cache expiration time (5 minutes).
const smtpStatusCacheExpire = 5 * time.Minute

// SMTP status cache key format.
const smtpStatusCacheKeyFormat = "smtp_status:%s:%s:%s"

// SmtpConnectionStatus Represents the SMTP connection status.
type SmtpConnectionStatus struct {
	Status   bool   `json:"status"`   // Whether the connection was successful.
	Msg      string `json:"msg"`      // Status description.
	LastTime int64  `json:"lastTime"` // Last test time (Unix timestamp).
}

// SmtpConnectionTestResult Complete SMTP connection test result.
type SmtpConnectionTestResult struct {
	Success         bool   // Whether the connection test was successful.
	Message         string // Connection test result message.
	ServerInfo      string // Server information (server welcome message).
	TlsStatus       string // TLS status (enabled and type).
	AuthStatus      string // Authentication status (success, failure, or unauthenticated).
	ConnectionTime  int64  // Time to establish the connection (milliseconds).
	SupportedAuth   string // Authentication methods supported by the server.
	RecommendConfig bool   // Whether this configuration is recommended.
}

// TestSmtpConnection Complete SMTP connection testing method.
func TestSmtpConnection(host, port, user, password string) *SmtpConnectionTestResult {
	result := &SmtpConnectionTestResult{
		Success:         false,
		Message:         "",
		ServerInfo:      "",
		TlsStatus:       "",
		AuthStatus:      "",
		ConnectionTime:  0,
		SupportedAuth:   "",
		RecommendConfig: false,
	}

	startTime := time.Now()

	// Set default values
	tlsProtocol := "STARTTLS"
	skipTlsVerify := 1

	// Connection timeout setting
	timeout := 10 * time.Second

	// Attempt to connect to the SMTP server
	serverAddr := fmt.Sprintf("%s:%s", host, port)

	// First, try a plain TCP connection
	conn, err := net.DialTimeout("tcp", serverAddr, timeout)
	if err != nil {
		result.Message = fmt.Sprintf("Unable to connect to SMTP server: %v", err)
		return result
	}
	defer conn.Close()

	// Create an SMTP client
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		result.Message = fmt.Sprintf("Failed to create SMTP client: %v", err)
		return result
	}
	defer client.Close()

	// Record server information
	result.ServerInfo = fmt.Sprintf("Connected to %s", serverAddr)

	// Automatically determine whether TLS is needed based on the port
	isTLS := false
	if port == "465" {
		// Port 465 typically requires direct SSL/TLS, but we already established a plain connection, so we just record the state.
		result.TlsStatus = "Port 465 typically requires direct TLS connection."
	} else if port == "587" || strings.ToUpper(tlsProtocol) == "STARTTLS" {
		// Port 587 typically uses STARTTLS
		tlsConfig := &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: skipTlsVerify == 1,
			MinVersion:         tls.VersionTLS12,
		}

		if err = client.StartTLS(tlsConfig); err != nil {
			result.TlsStatus = fmt.Sprintf("STARTTLS failed: %v", err)
		} else {
			result.TlsStatus = "STARTTLS enabled."
			isTLS = true
		}
	} else {
		// Other ports (e.g., 25) default to plaintext
		result.TlsStatus = "TLS not enabled (plaintext connection)."
	}

	// 发送EHLO命令
	//if err = client.Hello(heloName); err != nil {
	//	result.Message = fmt.Sprintf("EHLO命令失败: %v", err)
	//	return result
	//}

	// Check if AUTH is supported
	if ok, _ := client.Extension("AUTH"); ok {
		result.SupportedAuth = "Server supports authentication."
	} else {
		result.SupportedAuth = "Server does not support authentication or requires TLS."
	}

	// If authentication information is provided, attempt authentication
	if user != "" && password != "" {
		var auth smtp.Auth

		// Select authentication method based on the port and TLS status
		if port == "25" && !isTLS {
			// Port 25 with no TLS uses custom authentication method
			auth = NewCustomPlainAuth("", user, password, host)
		} else {
			// Other cases use standard authentication method
			auth = smtp.PlainAuth("", user, password, host)
		}

		if err = client.Auth(auth); err != nil {
			result.AuthStatus = fmt.Sprintf("Authentication failed: %v", err)
			result.Message = fmt.Sprintf("SMTP authentication failed: %v", err)
			return result
		}
		result.AuthStatus = "Authentication succeeded."
	} else {
		result.AuthStatus = "No authentication information provided."
	}

	// Calculate connection time
	connectionTime := time.Since(startTime).Milliseconds()
	result.ConnectionTime = connectionTime

	// Set success status
	result.Success = true
	result.Message = "SMTP connection test succeeded."

	// Provide configuration recommendations based on connection time and authentication status
	if connectionTime < 1000 && result.AuthStatus == "Authentication succeeded." {
		result.RecommendConfig = true
	}

	return result
}

// UpdateRelayStatus Periodically updates the SMTP connection status of all relay configurations.
func UpdateRelayStatus(ctx context.Context) {
	g.Log().Debug(ctx, "Starting periodic update of SMTP connection statuses...")

	// Query all active relay configurations.
	var relayConfigs []*entity.BmRelay
	err := g.DB().Model("bm_relay").Where("active", 1).Scan(&relayConfigs)
	if err != nil {
		//g.Log().Error(ctx, "Failed to query relay configurations:", err)
		return
	}

	if len(relayConfigs) == 0 {
		//g.Log().Debug(ctx, "No active relay configurations found, skipping SMTP status update.")
		return
	}

	//g.Log().Debugf(ctx, "Found %d active relay configurations, starting to check SMTP connection statuses.", len(relayConfigs))

	// Use WaitGroup to wait for all checks to complete.
	var wg sync.WaitGroup
	// Limit concurrency to avoid establishing too many connections simultaneously.
	semaphore := make(chan struct{}, 5)

	for _, config := range relayConfigs {
		wg.Add(1)

		// Use goroutines for concurrent checks.
		go func(cfg *entity.BmRelay) {
			defer wg.Done()

			// Acquire semaphore to limit concurrency.
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Decrypt password.
			password, err := DecryptPassword1(ctx, cfg.AuthPassword)
			if err != nil {
				g.Log().Warningf(ctx, "Failed to decrypt password for relay configuration ID=%d: %v", cfg.Id, err)
				return
			}

			// Perform SMTP connection check.
			status := checkSmtpConnectionImpl(cfg.RelayHost, cfg.RelayPort, cfg.AuthUser, password)

			// Generate cache key.
			cacheKey := fmt.Sprintf(smtpStatusCacheKeyFormat, cfg.RelayHost, cfg.RelayPort, cfg.AuthUser)

			// Store result in cache.
			err = smtpStatusCache.Set(ctx, cacheKey, status, smtpStatusCacheExpire)
			if err != nil {
				g.Log().Warningf(ctx, "Failed to cache SMTP status for relay configuration ID=%d: %v", cfg.Id, err)
			}

			g.Log().Debugf(ctx, "Updated SMTP status for relay configuration (ID=%d, Host=%s:%s): %v - %s",
				cfg.Id, cfg.RelayHost, cfg.RelayPort, status.Status, status.Msg)
		}(config)
	}

	// Wait for all checks to complete.
	wg.Wait()
	g.Log().Debug(ctx, "All SMTP connection statuses have been updated.")
}

// CheckSmtpConnection
func CheckSmtpConnection(ctx context.Context, host, port, user, password string) *SmtpConnectionStatus {

	cacheKey := fmt.Sprintf(smtpStatusCacheKeyFormat, host, port, user)

	if v, err := smtpStatusCache.Get(ctx, cacheKey); err == nil && v != nil {
		var status SmtpConnectionStatus
		if err := v.Scan(&status); err == nil {
			return &status
		}
	}

	return &SmtpConnectionStatus{
		Status:   false,
		Msg:      "Status unknown (under check)",
		LastTime: time.Now().Unix(),
	}
}

// ForceCheckSmtpConnection
func ForceCheckSmtpConnection(ctx context.Context, host, port, user, password string) *SmtpConnectionStatus {

	status := checkSmtpConnectionImpl(host, port, user, password)

	cacheKey := fmt.Sprintf(smtpStatusCacheKeyFormat, host, port, user)

	smtpStatusCache.Set(ctx, cacheKey, status, smtpStatusCacheExpire)

	return status
}

// checkSmtpConnectionImpl Actual implementation of SMTP connection checking.
func checkSmtpConnectionImpl(host, port, user, password string) *SmtpConnectionStatus {
	status := &SmtpConnectionStatus{
		Status:   false,
		Msg:      "Not checked",
		LastTime: time.Now().Unix(),
	}

	// Connection timeout setting (short time for batch checks)
	timeout := 5 * time.Second

	// Attempt TCP connection
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", serverAddr, timeout)
	if err != nil {
		status.Msg = public.Lang("Connection failed: {}", err)
		return status
	}
	defer conn.Close()

	// Create SMTP client
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		status.Msg = public.Lang("Failed to create SMTP client: {}", err)
		return status
	}
	defer client.Close()

	// Determine if TLS is needed based on the port
	if port == "587" {
		// Port 587 typically uses STARTTLS
		tlsConfig := &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS12,
		}

		if err = client.StartTLS(tlsConfig); err != nil {
			status.Msg = public.Lang("STARTTLS failed: {}", err)
			return status
		}
	}

	// If authentication information is provided, attempt authentication
	if user != "" && password != "" {
		var auth smtp.Auth

		// Select authentication method based on the port
		if port == "25" {
			// Port 25 uses custom authentication method
			auth = NewCustomPlainAuth("", user, password, host)
		} else {
			// Other cases use standard authentication method
			auth = smtp.PlainAuth("", user, password, host)
		}

		if err = client.Auth(auth); err != nil {
			status.Msg = public.Lang("Authentication failed: {}", err)
			return status
		}
	}

	// All checks passed
	status.Status = true
	status.Msg = public.Lang("Connection normal")
	return status
}

// ForceRefreshSmtpStatus Forces refresh of the SMTP connection status cache.
func ForceRefreshSmtpStatus(ctx context.Context, relayConfig *v1.BmRelay) *SmtpConnectionStatus {
	// Generate cache key
	cacheKey := fmt.Sprintf(smtpStatusCacheKeyFormat, relayConfig.RelayHost, relayConfig.RelayPort, relayConfig.AuthUser)

	// Perform actual check
	status := checkSmtpConnectionImpl(relayConfig.RelayHost, relayConfig.RelayPort, relayConfig.AuthUser, relayConfig.AuthPassword)

	// Update cache
	smtpStatusCache.Set(ctx, cacheKey, status, smtpStatusCacheExpire)

	return status
}

// ClearSmtpStatusCache Clears the SMTP status cache.
func ClearSmtpStatusCache(ctx context.Context) error {
	return smtpStatusCache.Clear(ctx)
}

func DecryptPassword1(ctx context.Context, encryptedHex string) (string, error) {

	password, err := DecryptPassword(ctx, encryptedHex)
	if password == "" {
		password = "********"
	}
	return password, err
}

func DecryptPassword(ctx context.Context, encryptedHex string) (string, error) {
	if encryptedHex == "" {
		return "", nil
	}

	encryptedBytes, err := hex.DecodeString(encryptedHex)
	if err != nil {
		//g.Log().Errorf(ctx, "Decryption failed, invalid hex format: %v", err)
		return "", gerror.Wrap(err, public.LangCtx(ctx, "Password format is incorrect"))
	}

	relayEncryptionKey, err := GetRelayEncryptionKey()
	if err != nil {
		return "", gerror.Wrap(err, public.LangCtx(ctx, "Failed to retrieve encryption key"))
	}

	keyBytes, err := hex.DecodeString(relayEncryptionKey)
	if err != nil {
		return "", gerror.Wrap(err, public.LangCtx(ctx, "Failed to parse encryption key"))
	}

	if len(keyBytes) < 16 {
		return "", gerror.New(public.LangCtx(ctx, "Encryption key length is insufficient"))
	}
	keyBytes = keyBytes[:16]

	decrypted, err := gaes.Decrypt(encryptedBytes, keyBytes)
	if err != nil {
		//g.Log().Errorf(ctx, "Password decryption failed: %v", err)
		return "", gerror.Wrap(err, public.LangCtx(ctx, "Password decryption failed"))
	}

	return string(decrypted), nil
}

func GetRelayEncryptionKey() (string, error) {

	val, _ := g.DB().Model("bm_options").
		Where("name", "relay_encryption_key").
		Value("value")

	if val != nil && val.String() != "" {
		return val.String(), nil
	} else {
		return "", gerror.New("Failed to insert new key and retrieve key again")
	}

}
