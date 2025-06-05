package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func generateApiKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (c *ControllerV1) ApiTemplatesCreate(ctx context.Context, req *v1.ApiTemplatesCreateReq) (res *v1.ApiTemplatesCreateRes, err error) {
	res = &v1.ApiTemplatesCreateRes{}

	// 验证模板是否存在
	count, err := g.DB().Model("email_templates").Where("id", req.TemplateId).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New(public.LangCtx(ctx, "邮件模板不存在"))
	}

	// 生成API密钥
	apiKey, err := generateApiKey()
	if err != nil {
		return nil, err
	}

	// 创建API模板
	_, err = g.DB().Model("api_templates").Insert(g.Map{
		"api_key":     apiKey,
		"api_name":    req.ApiName,
		"template_id": req.TemplateId,
		"subject":     req.Subject,
		"addresser":   req.Addresser,
		"full_name":   req.FullName,
		"unsubscribe": req.Unsubscribe,
		"track_open":  1,
		"track_click": 1,
		"active":      req.Active,
	})

	if err != nil {
		return nil, err
	}

	res.SetSuccess(public.LangCtx(ctx, "创建API成功"))
	return res, nil
}
