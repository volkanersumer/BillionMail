package subscribe_list

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"

	"billionmail-core/api/subscribe_list/v1"
)

func (c *ControllerV1) SubscribeConfirm(ctx context.Context, req *v1.SubscribeConfirmReq) (res *v1.SubscribeConfirmRes, err error) {
	res = &v1.SubscribeConfirmRes{}

	email, groupToken, err := getEmailFromToken(req.Token)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "确认链接无效")))
		return
	}
	if email == "" || groupToken == "" {
		res.SetError(gerror.New(public.LangCtx(ctx, "确认链接无效")))
		return
	}

	// 1. 查找订阅组
	group, err := getGroupByToken(groupToken)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "订阅组不存在")))
		return
	}

	// 2. 查找联系人
	contact, err := getContactByEmailAndGroup(req.Email, group.Id)
	if err != nil || contact == nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "联系人不存在")))
		return
	}
	hostUrl := public.GethostUrl()
	if contact.Status == 1 {
		// 已确认
		if group.AlreadyUrl != "" {
			res.Data = group.AlreadyUrl
		} else {
			res.Data = fmt.Sprintf("%s/already_subscribed.html", hostUrl)
		}
		res.SetSuccess(public.LangCtx(ctx, "已订阅"))
		return
	}

	// 4. 更新联系人状态为已确认
	err = updateContactStatus(req.Email, group.Id, 1)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "更新联系人状态失败")))
		return
	}

	// 5. 发送欢迎邮件
	if group.SendWelcomeEmail == 1 {
		if group.WelcomeHtml == "" {
			// 取默认欢迎邮件模版
			group.WelcomeHtml, _ = GetDefaultTemplate(1)
		}
		if group.WelcomeSubject == "" {
			group.WelcomeSubject = "默认主题"
		}
		gtimer.AddOnce(500*time.Millisecond, func() {

			err = sendMail(ctx, group.WelcomeHtml, req.Email, group.WelcomeSubject, "")
			if err != nil {
				g.Log().Error(ctx, "Failed to send welcome email: {}", err)
				return
			}

		})
	}


	// 6. 跳转到订阅成功页面
	if group.SuccessUrl != "" {
		res.Data = group.SuccessUrl
	} else {
		res.Data = fmt.Sprintf("%s/subscribe_success.html", hostUrl)
	}
	res.SetSuccess(public.LangCtx(ctx, "订阅成功"))
	return
}
