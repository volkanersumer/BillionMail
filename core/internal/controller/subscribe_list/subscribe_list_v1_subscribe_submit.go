package subscribe_list

import (
	"billionmail-core/api/subscribe_list/v1"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) SubscribeSubmit(ctx context.Context, req *v1.SubscribeSubmitReq) (res *v1.SubscribeSubmitRes, err error) {

	res = &v1.SubscribeSubmitRes{}

	// 1. 查找订阅组
	group, err := getGroupByToken(req.Token)
	if err != nil || group.Id == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "订阅组不存在")))
		return
	}

	hostUrl := public.GethostUrl()

	// 2. 查找联系人  存在联系人且已确认 -- 已订阅    存在联系人但未确认 -- 二次确认   不存在联系人 -- 一次确认
	contact, err := getContactByEmailAndGroup(req.Email, group.Id)

	if contact.Id != 0 {
		if contact.Active == 0 {
			// 存在联系人但退订，更新状态为1  改为已订阅
			err = updateContactStatus(req.Email, group.Id, 1)
			if err != nil {
				res.SetError(gerror.New(public.LangCtx(ctx, "更新联系人状态失败")))
				return
			}

			if group.SuccessUrl != "" {
				res.Data = group.SuccessUrl
			} else {
				res.Data = fmt.Sprintf("%s/subscribe_success.html", hostUrl)

			}
			res.SetSuccess(public.LangCtx(ctx, "已订阅1"))
			return
		}
		if contact.Status == 1 {
			// 已订阅
			if group.AlreadyUrl != "" {
				res.Data = group.AlreadyUrl
			} else {
				res.Data = fmt.Sprintf("%s/already_subscribed.html", hostUrl)
			}
			res.SetSuccess(public.LangCtx(ctx, "已订阅2"))
			return
		} else {
			// 存在联系人但未确认
		}
	}

	// 3. 判断是否二次确认
	if group.DoubleOptin == 1 {
		// 二次确认
		// 3.1 添加/更新联系人，status=0
		err = addOrUpdateContact(req.Email, group.Id, req.Attribs, 0)
		if err != nil {
			res.SetError(gerror.New(public.LangCtx(ctx, "添加联系人失败")))
			return
		}

		// 3.2 发送确认邮件
		confirmToken := generateConfirmToken(req.Email, group.Token)
		confirmUrl := buildConfirmUrl(confirmToken)

		if group.ConfirmHtml == "" {
			group.ConfirmHtml, _ = GetDefaultTemplate(2)
		}
		if group.ConfirmSubject == "" {
			group.ConfirmSubject = "默认主题"
		}
		gtimer.AddOnce(500*time.Millisecond, func() {
			err = sendMail(ctx, group.ConfirmHtml, req.Email, group.ConfirmSubject, confirmUrl)
			if err != nil {
				g.Log().Error(ctx, "Failed to send confirm email: {}", err)
				return
			}

		})

		// 3.3 跳转到订阅确认页面
		if group.ConfirmUrl != "" {
			res.Data = group.ConfirmUrl
		} else {
			res.Data = fmt.Sprintf("%s/subscribe_confirm.html", hostUrl)
		}

		res.SetError(gerror.New(public.LangCtx(ctx, "请检查您的邮箱进行确认订阅")))
		return
	} else {
		// 一次确认
		// 4.1 添加/更新联系人，status=1
		err = addOrUpdateContact(req.Email, group.Id, req.Attribs, 1)
		if err != nil {
			res.SetError(gerror.New(public.LangCtx(ctx, "添加联系人失败")))
			return
		}
		// 4.2 发送欢迎邮件 当组内允许发送时发送
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

		// 4.3 跳转到订阅成功页面
		if group.SuccessUrl != "" {
			res.Data = group.SuccessUrl
		} else {
			res.Data = fmt.Sprintf("%s/subscribe_success.html", hostUrl)
		}
		res.SetSuccess(public.LangCtx(ctx, "Subscribe successfully"))
		return

	}
}
