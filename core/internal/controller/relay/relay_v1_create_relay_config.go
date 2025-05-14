package relay

import (
	"context"

	"billionmail-core/api/relay/v1"
	"billionmail-core/internal/service/public"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// CreateRelayConfig Creates a relay configuration
func (c *ControllerV1) CreateRelayConfig(ctx context.Context, req *v1.CreateRelayConfigReq) (res *v1.CreateRelayConfigRes, err error) {
	res = &v1.CreateRelayConfigRes{}
	// Check if the domain already exists
	count, err := g.DB().Model("bm_relay").Where("sender_domain", req.SenderDomain).Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check relay configuration: {}", err.Error())))
		return res, nil
	}
	if count > 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "Relay configuration for this sender domain already exists: {}", req.SenderDomain)))
		return res, nil
	}
	// Encrypt the password
	encryptedPass, err := EncryptPassword(ctx, req.AuthPassword)
	if err != nil {
		res.SetError(err)
		return res, nil
	}
	// Prepare data
	now := time.Now().Unix()
	_, err = g.DB().Model("bm_relay").Insert(g.Map{
		"remark":        req.Remark,
		"sender_domain": req.SenderDomain,
		"relay_host":    req.RelayHost,
		"relay_port":    req.RelayPort,
		"auth_user":     req.AuthUser,
		"auth_password": encryptedPass,
		"ip":            req.IP,
		"host":          req.Host,
		"create_time":   now,
		"update_time":   now,
		"active":        req.Active,
	})

	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create relay configuration: {}", err.Error())))
		return res, nil
	}
	// Sync configuration to Postfix
	if err := SyncRelayConfigsToPostfix(ctx); err != nil {
		// g.Log().Error(ctx, "Failed to sync relay configuration to Postfix:", err)
		res.SetError(gerror.New(public.LangCtx(ctx, "Created successfully but failed to sync configuration: {}", err.Error())))
		return res, nil
	}
	// Generate SPF record suggestion
	spfRecord := GenerateSPFRecord(req.IP, req.Host, req.SenderDomain)
	res.SPFRecord = v1.DNSRecord{
		Type:  "TXT",
		Host:  "@",
		Value: spfRecord,
	}
	res.SetSuccess(public.LangCtx(ctx, "Relay configuration created successfully"))
	return res, nil
}
