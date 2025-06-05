package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/mail_service"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ApiMailSend(ctx context.Context, req *v1.ApiMailSendReq) (res *v1.ApiMailSendRes, err error) {
	res = &v1.ApiMailSendRes{}

	// 获取API模板信息
	var apiTemplate entity.ApiTemplates
	err = g.DB().Model("api_templates").
		Where("api_key", req.ApiKey).
		Where("active", 1).
		Scan(&apiTemplate)
	if err != nil {
		return nil, err
	}
	if apiTemplate.Id == 0 {
		return nil, gerror.New(public.LangCtx(ctx, "API不存在或未启用"))
	}

	// 获取邮件模板
	var emailTemplate entity.EmailTemplate
	err = g.DB().Model("email_templates").
		Where("id", apiTemplate.TemplateId).
		Scan(&emailTemplate)
	if err != nil {
		return nil, err
	}
	if emailTemplate.Id == 0 {
		return nil, gerror.New(public.LangCtx(ctx, "邮件模板不存在"))
	}

	// 使用指定的发件人或默认发件人
	addresser := req.Addresser
	if addresser == "" {
		addresser = apiTemplate.Addresser
	}

	// 创建邮件发送器
	sender, err := mail_service.NewEmailSenderWithLocal(addresser)
	if err != nil {
		return nil, gerror.New(public.LangCtx(ctx, "创建邮件发送器失败: {}", err))
	}
	defer sender.Close()

	// 生成消息ID
	messageId := sender.GenerateMessageID()

	// 创建邮件消息
	message := mail_service.NewMessage(apiTemplate.Subject, emailTemplate.Content)
	message.SetMessageID(messageId)

	// 设置发件人显示名称
	if apiTemplate.FullName != "" {
		message.SetRealName(apiTemplate.FullName)
	}

	// 发送邮件
	err = sender.Send(message, []string{req.Recipient})
	if err != nil {
		return nil, gerror.New(public.LangCtx(ctx, "发送邮件失败: {}", err))
	}

	// 记录发送日志
	_, err = g.DB().Model("api_mail_logs").Insert(g.Map{
		"api_id":     apiTemplate.Id,
		"recipient":  req.Recipient,
		"message_id": messageId,
		"addresser":  addresser,
	})
	if err != nil {
		g.Log().Error(ctx, "记录发送日志失败:", err)
	}

	res.SetSuccess(public.LangCtx(ctx, "发送邮件成功"))
	return res, nil
}
