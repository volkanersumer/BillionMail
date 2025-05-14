package relay

import (
	"billionmail-core/internal/model/entity"
	"context"

	"billionmail-core/api/relay/v1"
	"billionmail-core/internal/service/public"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) UpdateRelayConfig(ctx context.Context, req *v1.UpdateRelayConfigReq) (res *v1.UpdateRelayConfigRes, err error) {
	res = &v1.UpdateRelayConfigRes{}

	var relayInfo entity.BmRelay
	err = g.DB().Model("bm_relay").Where("id", req.ID).Scan(&relayInfo)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check relay configuration: {}", err.Error())))
		return res, nil
	}
	if relayInfo.Id == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "Relay configuration does not exist")))
		return res, nil
	}
	// Domain updates are not allowed
	// If the domain is updated, check if the domain is already used by another record
	//if req.SenderDomain != "" && req.SenderDomain != relayInfo.SenderDomain {
	//	count, err := g.DB().Model("bm_relay").
	//		Where("sender_domain", req.SenderDomain).
	//		Where("id != ?", req.ID).
	//		Count()
	//	if err != nil {
	//		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check domain uniqueness: {}", err.Error())))
	//		return res, nil
	//	}
	//	if count > 0 {
	//		res.SetError(gerror.New(public.LangCtx(ctx, "Relay configuration for this sender domain already exists: {}", req.SenderDomain)))
	//		return res, nil
	//	}
	//}

	updateData := g.Map{
		"update_time": time.Now().Unix(),
	}
	// Update only the provided fields
	if req.Remark != "" {
		updateData["remark"] = req.Remark
	}
	//if req.SenderDomain != "" {
	//	updateData["sender_domain"] = req.SenderDomain
	//}
	if req.RelayHost != "" {
		updateData["relay_host"] = req.RelayHost
	}
	if req.RelayPort != "" {
		updateData["relay_port"] = req.RelayPort
	}
	if req.AuthUser != "" {
		updateData["auth_user"] = req.AuthUser
	}
	// Special handling for password - encrypt and store if a new password is provided
	if req.AuthPassword != "" && req.AuthPassword != "********" {
		encryptedPass, err := EncryptPassword(ctx, req.AuthPassword)
		if err != nil {
			res.SetError(err)
			return res, nil
		}
		updateData["auth_password"] = encryptedPass
	}
	if req.IP != "" {
		updateData["ip"] = req.IP
	}
	if req.Host != "" {
		updateData["host"] = req.Host
	}
	// Active can be updated separately to 0 as a disable operation
	if req.Active >= 0 {
		updateData["active"] = req.Active
	}

	_, err = g.DB().Model("bm_relay").Where("id", req.ID).Update(updateData)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to update relay configuration: {}", err.Error())))
		return res, nil
	}
	// Sync configuration to Postfix
	if err := SyncRelayConfigsToPostfix(ctx); err != nil {

		res.SetError(gerror.New(public.LangCtx(ctx, "Updated successfully but failed to sync configuration: {}", err.Error())))
		return res, nil
	}
	// Retrieve the latest information for generating SPF records
	var latestRelayInfo entity.BmRelay
	_ = g.DB().Model("bm_relay").Where("id", req.ID).Scan(&latestRelayInfo)

	spfRecord := GenerateSPFRecord(latestRelayInfo.IP, latestRelayInfo.Host, latestRelayInfo.SenderDomain)
	res.SPFRecord = v1.DNSRecord{
		Type:  "TXT",
		Host:  "@",
		Value: spfRecord,
	}

	if len(updateData) > 2 { // In addition to update_time, other fields have been updated
		res.SetSuccess(public.LangCtx(ctx, "Relay configuration updated successfully"))
	} else if req.Active == 1 {
		res.SetSuccess(public.LangCtx(ctx, "Relay configuration has been enabled"))
	} else if req.Active == 0 {
		res.SetSuccess(public.LangCtx(ctx, "Relay configuration has been disabled"))
	} else {
		res.SetSuccess(public.LangCtx(ctx, "Relay configuration updated successfully"))
	}

	return res, nil
}
