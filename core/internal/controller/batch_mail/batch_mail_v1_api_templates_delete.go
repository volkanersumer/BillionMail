package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ApiTemplatesDelete(ctx context.Context, req *v1.ApiTemplatesDeleteReq) (res *v1.ApiTemplatesDeleteRes, err error) {
	res = &v1.ApiTemplatesDeleteRes{}

	// 验证API是否存在
	count, err := g.DB().Model("api_templates").Where("id", req.ID).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New(public.LangCtx(ctx, "API不存在"))
	}

	// 开启事务

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除API相关的发送日志
		_, err = tx.Model("api_mail_logs").Where("api_id", req.ID).Delete()
		if err != nil {
			return err
		}

		// 删除API模板
		_, err = tx.Model("api_templates").Where("id", req.ID).Delete()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	res.SetSuccess(public.LangCtx(ctx, "删除API成功"))
	return res, nil
}
