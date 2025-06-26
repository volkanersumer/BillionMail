package settings

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/settings/v1"
)

func (c *ControllerV1) DeleteIPWhitelist(ctx context.Context, req *v1.DeleteIPWhitelistReq) (res *v1.DeleteIPWhitelistRes, err error) {
	res = &v1.DeleteIPWhitelistRes{}

	if req.ID <= 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "ID cannot be empty")))
		return
	}

	// Check if it exists
	count, err := g.DB().Model("bm_console_ip_whitelist").Where("id", req.ID).Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to query IP whitelist: {}", err)))
		return
	}
	if count == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "The IP whitelist does not exist")))
		return
	}

	_, err = g.DB().Model("bm_console_ip_whitelist").Where("id", req.ID).Delete()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to delete IP from whitelist: {}", err)))
		return
	}

	res.SetSuccess(public.LangCtx(ctx, "IP removed from whitelist successfully"))
	return res, nil
}
