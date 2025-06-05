package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) ApiTemplatesUpdate(ctx context.Context, req *v1.ApiTemplatesUpdateReq) (res *v1.ApiTemplatesUpdateRes, err error) {
	res = &v1.ApiTemplatesUpdateRes{}

	// 验证API是否存在
	count, err := g.DB().Model("api_templates").Where("id", req.ID).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New(public.LangCtx(ctx, "API不存在"))
	}

	// 验证模板是否存在
	count, err = g.DB().Model("email_templates").Where("id", req.TemplateId).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New(public.LangCtx(ctx, "邮件模板不存在"))
	}
	// 当前时间
	now := time.Now().Unix()

	// 更新API模板
	_, err = g.DB().Model("api_templates").
		Where("id", req.ID).
		Update(g.Map{
			"api_name":    req.ApiName,
			"template_id": req.TemplateId,
			"subject":     req.Subject,
			"addresser":   req.Addresser,
			"full_name":   req.FullName,
			"unsubscribe": req.Unsubscribe,
			"track_open":  req.TrackOpen,
			"track_click": req.TrackClick,
			"active":      req.Active,
			"update_time": now,
		})

	if err != nil {
		return nil, err
	}

	res.SetSuccess(public.LangCtx(ctx, "更新API成功"))
	return res, nil
}
