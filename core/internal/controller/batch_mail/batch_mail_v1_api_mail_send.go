package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/mail_service"
	"billionmail-core/internal/service/maillog_stat"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"time"
)

func (c *ControllerV1) ApiMailSend(ctx context.Context, req *v1.ApiMailSendReq) (res *v1.ApiMailSendRes, err error) {
	res = &v1.ApiMailSendRes{}

	// 1. 校验API Key
	apiTemplate, err := getApiTemplateByKey(ctx, req.ApiKey)
	if err != nil {
		res.Code = 1001
		res.SetError(gerror.New(public.LangCtx(ctx, "API密钥无效")))
		return res, nil
	}

	// 2. 校验邮件模板
	emailTemplate, err := getEmailTemplateById(ctx, apiTemplate.TemplateId)
	if err != nil {
		res.Code = 1004
		res.SetError(gerror.New(public.LangCtx(ctx, "邮件模板不存在")))
		return res, nil
	}

	// 3. 校验收件人
	if req.Recipient == "" || !strings.Contains(req.Recipient, "@") {
		res.Code = 1003
		res.SetError(gerror.New(public.LangCtx(ctx, "收件人无效")))
		return res, nil
	}

	// 4. 处理联系人和分组
	contact, err := ensureContactAndGroup(ctx, req.Recipient, apiTemplate.Id)
	if err != nil {
		res.Code = 1003
		res.SetError(gerror.New(public.LangCtx(ctx, "收件人处理失败: {}", err.Error())))
		return res, nil
	}

	// 发件人处理
	if req.Addresser == "" {
		req.Addresser = apiTemplate.Addresser
	}

	// 5. 处理邮件内容和主题
	content, subject := processMailContentAndSubject(ctx, emailTemplate.Content, apiTemplate.Subject, apiTemplate, contact, req)

	// 6. 发送邮件
	err = sendApiMail(ctx, apiTemplate, subject, content, req.Recipient, req.Addresser)
	if err != nil {
		res.Code = 1005
		res.SetError(gerror.New(public.LangCtx(ctx, "发送邮件失败: {}", err.Error())))
		return res, nil
	}

	res.SetSuccess(public.LangCtx(ctx, "发送邮件成功"))
	return res, nil
}

// 获取API模板
func getApiTemplateByKey(ctx context.Context, apiKey string) (*entity.ApiTemplates, error) {
	var apiTemplate entity.ApiTemplates
	err := g.DB().Model("api_templates").Where("api_key", apiKey).Where("active", 1).Scan(&apiTemplate)
	if err != nil || apiTemplate.Id == 0 {
		return nil, gerror.New("API密钥无效")
	}
	return &apiTemplate, nil
}

// 获取邮件模板
func getEmailTemplateById(ctx context.Context, templateId int) (*entity.EmailTemplate, error) {
	var emailTemplate entity.EmailTemplate
	err := g.DB().Model("email_templates").Where("id", templateId).Scan(&emailTemplate)
	if err != nil || emailTemplate.Id == 0 {
		return nil, gerror.New("邮件模板不存在")
	}
	return &emailTemplate, nil
}

// 确保联系人和分组存在
func ensureContactAndGroup(ctx context.Context, email string, apiId int) (entity.Contact, error) {
	var contact entity.Contact
	err := g.DB().Model("bm_contacts").Where("email", email).Scan(&contact)
	now := int(time.Now().Unix())
	if err != nil {
		return contact, err
	}
	if contact.Id == 0 {
		apiGroupName := fmt.Sprintf("api_group_%d", apiId)
		var group entity.ContactGroup
		err = g.DB().Model("bm_contact_groups").Where("name", apiGroupName).Scan(&group)
		if err != nil {
			return contact, err
		}
		if group.Id == 0 {
			groupResult, err := g.DB().Model("bm_contact_groups").Insert(g.Map{
				"name":        apiGroupName,
				"description": fmt.Sprintf("API %d 自动创建的联系人组", apiId),
				"create_time": now,
				"update_time": now,
			})
			if err != nil {
				return contact, err
			}
			groupId, _ := groupResult.LastInsertId()
			group.Id = int(groupId)
		}
		contactResult, err := g.DB().Model("bm_contacts").Insert(g.Map{
			"email":       email,
			"group_id":    group.Id,
			"active":      1,
			"status":      1,
			"create_time": now,
		})
		if err != nil {
			return contact, err
		}
		contactId, _ := contactResult.LastInsertId()
		contact.Id = int(contactId)
		contact.Email = email
		contact.GroupId = group.Id
		contact.Active = 1
		contact.Status = 1
		contact.CreateTime = now
	}
	return contact, nil
}

// 处理邮件内容和主题
func processMailContentAndSubject(ctx context.Context, content, subject string, apiTemplate *entity.ApiTemplates, contact entity.Contact, req *v1.ApiMailSendReq) (string, string) {
	// 退订链接处理
	if apiTemplate.Unsubscribe == 1 {
		if !strings.Contains(content, "__UNSUBSCRIBE_URL__") && !strings.Contains(content, "{{ UnsubscribeURL . }}") {
			content = public.AddUnsubscribeButton(content)
		}
		domain := domains.GetBaseURLBySender(req.Addresser)
		unsubscribeURL := fmt.Sprintf("%s/api/unsubscribe", domain)
		groupURL := fmt.Sprintf("%s/api/unsubscribe/user_group", domain)
		jwtToken, _ := batch_mail.GenerateUnsubscribeJWT(req.Recipient, apiTemplate.TemplateId, apiTemplate.Id)
		unsubscribeJumpURL := fmt.Sprintf("%s/unsubscribe.html?jwt=%s&email=%s&url_type=%s&url_unsubscribe=%s", domain, jwtToken, req.Recipient, groupURL, unsubscribeURL)

		if contact.Id > 0 {
			engine := batch_mail.GetTemplateEngine()
			renderedContent, err := engine.RenderEmailTemplate(ctx, content, &contact, nil, unsubscribeJumpURL)
			if err == nil {
				content = renderedContent
			}
			renderedSubject, err := engine.RenderEmailTemplate(ctx, subject, &contact, nil, unsubscribeJumpURL)
			if err == nil {
				subject = renderedSubject
			}
		} else {
			content = strings.ReplaceAll(content, "{{ UnsubscribeURL . }}", unsubscribeJumpURL)
		}
	} else if contact.Id > 0 {
		engine := batch_mail.GetTemplateEngine()
		renderedContent, err := engine.RenderEmailTemplate(ctx, content, &contact, nil, "")
		if err == nil {
			content = renderedContent
		}
		renderedSubject, err := engine.RenderEmailTemplate(ctx, subject, &contact, nil, "")
		if err == nil {
			subject = renderedSubject
		}
	}
	return content, subject
}

// 发送邮件
func sendApiMail(ctx context.Context, apiTemplate *entity.ApiTemplates, subject, content, recipient, addresser string) error {

	// 创建邮件发送器
	sender, err := mail_service.NewEmailSenderWithLocal(addresser)
	if err != nil {
		return gerror.New(public.LangCtx(ctx, "创建邮件发送器失败: {}", err))
	}
	defer sender.Close()

	// 生成消息ID
	messageId := sender.GenerateMessageID()
	// 添加邮件追踪
	baseURL := domains.GetBaseURLBySender(addresser)
	// 追踪id加上一千万 防止和营销任务的id冲突
	apiTemplate_id := apiTemplate.Id + 1000000000
	mailTracker := maillog_stat.NewMailTracker(content, apiTemplate_id, messageId, recipient, baseURL)
	mailTracker.TrackLinks()
	mailTracker.AppendTrackingPixel()
	content = mailTracker.GetHTML()

	// 创建邮件消息
	message := mail_service.NewMessage(subject, content)
	message.SetMessageID(messageId)
	// 设置发件人显示名称
	if apiTemplate.FullName != "" {
		message.SetRealName(apiTemplate.FullName)
	}
	// 发送邮件
	err = sender.Send(message, []string{recipient})
	if err != nil {
		return err
	}
	// 记录邮件日志
	messageId = strings.Trim(messageId, "<>")
	_, err = g.DB().Model("api_mail_logs").Insert(g.Map{
		"api_id":     apiTemplate.Id,
		"recipient":  recipient,
		"message_id": messageId,
		"addresser":  addresser,
	})
	if err != nil {
		return err
	}
	return nil
}
