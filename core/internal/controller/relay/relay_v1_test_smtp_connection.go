package relay

import (
	"billionmail-core/api/relay/v1"
	"billionmail-core/internal/service/public"
	"billionmail-core/internal/service/relay"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"net/smtp"
)

// customPlainAuth A custom PLAIN authentication method that allows authentication over unencrypted connections
type customPlainAuth struct {
	identity, username, password string
	host                         string
}

func (a *customPlainAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	//Ignoring connection status and host checks, PLAIN authentication is enforced
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

		return nil, fmt.Errorf("unexpected server challenge")
	}
	return nil, nil
}

func newCustomPlainAuth(identity, username, password, host string) smtp.Auth {
	return &customPlainAuth{
		identity: identity,
		username: username,
		password: password,
		host:     host,
	}
}

// TestSmtpConnection
func (c *ControllerV1) TestSmtpConnection(ctx context.Context, req *v1.TestSmtpConnectionReq) (res *v1.TestSmtpConnectionRes, err error) {
	res = &v1.TestSmtpConnectionRes{}

	result := relay.TestSmtpConnection(
		req.RelayHost,
		req.RelayPort,
		req.AuthUser,
		req.AuthPassword,
	)

	if !result.Success {
		res.SetError(gerror.New(public.LangCtx(ctx, result.Message)))
		return res, nil
	}

	res.Data.ServerInfo = result.ServerInfo
	res.Data.TlsStatus = result.TlsStatus
	res.Data.AuthStatus = result.AuthStatus
	res.Data.ConnectionTime = result.ConnectionTime
	res.Data.SupportedAuth = result.SupportedAuth
	res.Data.RecommendConfig = result.RecommendConfig

	res.SetSuccess(public.LangCtx(ctx, "The SMTP connection test was successful"))

	return res, nil
}
