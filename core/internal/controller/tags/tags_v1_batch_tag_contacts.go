package tags

import (
	"billionmail-core/api/tags/v1"
	"billionmail-core/internal/consts"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"time"
)

func (c *ControllerV1) BatchTagContacts(ctx context.Context, req *v1.BatchTagContactsReq) (res *v1.BatchTagContactsRes, err error) {
	res = &v1.BatchTagContactsRes{}

	// 验证联系人组是否存在
	var groupCount int
	groupCount, err = g.DB().Model("bm_contact_groups").Where("id", req.GroupId).Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check group")))
		return
	}
	if groupCount == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "Group not found")))
		return
	}

	// 验证标签是否存在且属于该组
	var tag struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		GroupId int    `json:"group_id"`
	}
	err = g.DB().Model("bm_tags").Where("id", req.TagId).Scan(&tag)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check tag")))
		return
	}
	if tag.Id == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "Tag not found")))
		return
	}
	if tag.GroupId != req.GroupId {
		res.SetError(gerror.New(public.LangCtx(ctx, "Tag does not belong to this group")))
		return
	}

	// 解析邮箱列表
	emailList := parseEmailList(req.Data)
	if len(emailList) == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "No valid emails provided")))
		return
	}

	g.Log().Infof(ctx, "BatchTagContacts: Processing %d emails for tag %s (ID: %d), mode: %d",
		len(emailList), tag.Name, req.TagId, req.MarkInclude)

	var processedCount, skippedCount, errorCount int

	// 执行批量标记操作
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if req.MarkInclude == 1 {
			// 标记包含：为指定邮箱添加标签
			processedCount, skippedCount, errorCount, err = c.batchAddTags(ctx, tx, req.GroupId, req.TagId, emailList)
		} else {
			// 标记���包含：为指定邮箱以外的所有邮箱添加标签
			processedCount, skippedCount, errorCount, err = c.batchAddTagsExclude(ctx, tx, req.GroupId, req.TagId, emailList)
		}
		return err
	})

	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to process batch tag operation")))
		return
	}

	// 记录操作日志
	operation := " Include"
	if req.MarkInclude == 0 {
		operation = " Not included"
	}
	_ = public.WriteLog(ctx, public.LogParams{
		Type: consts.LOGTYPE.Tag,
		Log: "Batch add" + " tag: " + tag.Name + operation + " for " +
			g.NewVar(processedCount).String() + " contacts successfully",
		Data: map[string]interface{}{
			"tag_id":       req.TagId,
			"tag_name":     tag.Name,
			"group_id":     req.GroupId,
			"operation":    operation,
			"total_emails": len(emailList),
			"processed":    processedCount,
			"skipped":      skippedCount,
			"errors":       errorCount,
		},
	})

	// 构建成功消息
	successMsg := public.LangCtx(ctx, "Batch tag operation completed")
	if req.MarkInclude == 1 {
		successMsg = public.LangCtx(ctx, "Successfully added tag to {} contacts", processedCount)
	} else {
		successMsg = public.LangCtx(ctx, "Successfully removed tag from {} contacts", processedCount)
	}

	if skippedCount > 0 {
		successMsg += public.LangCtx(ctx, " ( {}  skipped,  {}  errors)", skippedCount, errorCount)
	}

	res.SetSuccess(successMsg)
	return
}

// batchAddTags 批量添加标签
func (c *ControllerV1) batchAddTags(ctx context.Context, tx gdb.TX, groupId, tagId int, emailList []string) (processed, skipped, errors int, err error) {
	const batchSize = 500 // PostgreSQL 安全批量大小
	now := time.Now().Unix()

	// 分批处理邮箱列表
	for i := 0; i < len(emailList); i += batchSize {
		end := i + batchSize
		if end > len(emailList) {
			end = len(emailList)
		}

		currentBatch := emailList[i:end]
		g.Log().Debug(ctx, "Processing batch %d-%d of %d emails", i+1, end, len(emailList))

		// 查询当前批次中存在于组中的联系人
		var existingContacts []struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}
		err = tx.Model("bm_contacts").
			Fields("id, email").
			Where("group_id", groupId).
			WhereIn("email", currentBatch).
			Scan(&existingContacts)
		if err != nil {
			g.Log().Error(ctx, "Failed to query existing contacts: %v", err)
			errors += len(currentBatch)
			continue
		}

		if len(existingContacts) == 0 {
			skipped += len(currentBatch)
			continue
		}

		// 提取联系人ID
		contactIds := make([]int, len(existingContacts))
		emailToContactId := make(map[string]int)
		for j, contact := range existingContacts {
			contactIds[j] = contact.Id
			emailToContactId[contact.Email] = contact.Id
		}

		// 查询已存在的标签关系，避免重复插入
		var existingTags []struct {
			ContactId int `json:"contact_id"`
		}
		err = tx.Model("bm_contact_tags").
			Fields("contact_id").
			Where("tag_id", tagId).
			WhereIn("contact_id", contactIds).
			Scan(&existingTags)
		if err != nil {
			g.Log().Error(ctx, "Failed to query existing tags: %v", err)
			errors += len(existingContacts)
			continue
		}

		// 创建已存在标签的联系人ID集合
		existingTagContactIds := make(map[int]bool)
		for _, tag := range existingTags {
			existingTagContactIds[tag.ContactId] = true
		}

		// 准备插入数据（仅插入未标记的联���人）
		var insertData []g.Map
		for _, contact := range existingContacts {
			if !existingTagContactIds[contact.Id] {
				insertData = append(insertData, g.Map{
					"contact_id":  contact.Id,
					"tag_id":      tagId,
					"create_time": now,
				})
			}
		}

		if len(insertData) > 0 {
			// 批量插入标签关系
			_, err = tx.Model("bm_contact_tags").InsertIgnore(insertData)
			if err != nil {
				g.Log().Error(ctx, "Failed to insert contact tags: %v", err)
				errors += len(insertData)
				continue
			}
			processed += len(insertData)
		}

		// 计算跳过的邮箱数量（组中不存在的邮箱 + 已有标签的邮箱）
		skippedInBatch := len(currentBatch) - len(existingContacts) + len(existingTags)
		skipped += skippedInBatch
	}

	return processed, skipped, errors, nil
}

// batchRemoveTags 批量移除标签
func (c *ControllerV1) batchRemoveTags(ctx context.Context, tx gdb.TX, groupId, tagId int, emailList []string) (processed, skipped, errors int, err error) {
	const batchSize = 500 // PostgreSQL 安全批量大小

	// 分批处理邮箱列表
	for i := 0; i < len(emailList); i += batchSize {
		end := i + batchSize
		if end > len(emailList) {
			end = len(emailList)
		}

		currentBatch := emailList[i:end]
		g.Log().Debug(ctx, "Processing batch %d-%d of %d emails", i+1, end, len(emailList))

		// 查询当前批次中存在于组中且有该标签的联系人
		var taggedContacts []struct {
			ContactId int    `json:"contact_id"`
			Email     string `json:"email"`
		}
		err = tx.Model("bm_contacts c").
			LeftJoin("bm_contact_tags ct", "c.id = ct.contact_id").
			Fields("c.id as contact_id, c.email").
			Where("c.group_id", groupId).
			Where("c.active", 1).
			Where("ct.tag_id", tagId).
			WhereIn("c.email", currentBatch).
			Scan(&taggedContacts)
		if err != nil {
			g.Log().Error(ctx, "Failed to query tagged contacts: %v", err)
			errors += len(currentBatch)
			continue
		}

		if len(taggedContacts) == 0 {
			skipped += len(currentBatch)
			continue
		}

		// 提取联系人ID
		contactIds := make([]int, len(taggedContacts))
		for j, contact := range taggedContacts {
			contactIds[j] = contact.ContactId
		}

		// 批量删除标签关系
		result, err := tx.Model("bm_contact_tags").
			Where("tag_id", tagId).
			WhereIn("contact_id", contactIds).
			Delete()
		if err != nil {
			g.Log().Error(ctx, "Failed to delete contact tags: %v", err)
			errors += len(taggedContacts)
			continue
		}

		affected, _ := result.RowsAffected()
		processed += int(affected)

		// 计算跳过的邮箱数量（组中不存在的邮箱或没有该标签的邮箱）
		skippedInBatch := len(currentBatch) - len(taggedContacts)
		skipped += skippedInBatch
	}

	return processed, skipped, errors, nil
}

// batchAddTagsExclude 批量为指定邮箱以外的联系人添加标签
func (c *ControllerV1) batchAddTagsExclude(ctx context.Context, tx gdb.TX, groupId, tagId int, excludeEmailList []string) (processed, skipped, errors int, err error) {
	const batchSize = 500 // PostgreSQL 安全批量大小
	now := time.Now().Unix()

	g.Log().Infof(ctx, "batchAddTagsExclude: Excluding %d emails from tagging", len(excludeEmailList))

	// 创建排除邮箱的映射表，用于快速查找
	excludeEmailMap := make(map[string]bool)
	for _, email := range excludeEmailList {
		excludeEmailMap[strings.ToLower(email)] = true
	}

	// 分批查询组中的所有联系人（排除指定邮箱）
	var offset = 0
	for {
		// 查询当前批次的联系人（不在排除列表中的）
		var allContacts []struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}
		err = tx.Model("bm_contacts").
			Fields("id, email").
			Where("group_id", groupId).
			Limit(batchSize).
			Offset(offset).
			Scan(&allContacts)
		if err != nil {
			g.Log().Error(ctx, "Failed to query group contacts: %v", err)
			errors += batchSize
			break
		}

		// 如果没有更多联系人，退出循环
		if len(allContacts) == 0 {
			break
		}

		g.Log().Debug(ctx, "Processing batch starting from offset %d, found %d contacts", offset, len(allContacts))

		// 过滤出不在排除列表中的联系人
		var targetContacts []struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}
		for _, contact := range allContacts {
			if !excludeEmailMap[strings.ToLower(contact.Email)] {
				targetContacts = append(targetContacts, contact)
			} else {
				g.Log().Debug(ctx, "Excluding email: %s", contact.Email)
			}
		}

		if len(targetContacts) == 0 {
			skipped += len(allContacts)
			offset += batchSize
			continue
		}

		// 提取目标联系人ID
		targetContactIds := make([]int, len(targetContacts))
		for j, contact := range targetContacts {
			targetContactIds[j] = contact.Id
		}

		// 查询已存在的标签关系，避免重复插入
		var existingTags []struct {
			ContactId int `json:"contact_id"`
		}
		err = tx.Model("bm_contact_tags").
			Fields("contact_id").
			Where("tag_id", tagId).
			WhereIn("contact_id", targetContactIds).
			Scan(&existingTags)
		if err != nil {
			g.Log().Error(ctx, "Failed to query existing tags: %v", err)
			errors += len(targetContacts)
			offset += batchSize
			continue
		}

		// 创建已存在标签的联系人ID集合
		existingTagContactIds := make(map[int]bool)
		for _, tag := range existingTags {
			existingTagContactIds[tag.ContactId] = true
		}

		// 准备插入数据（仅插入未标记的联系人）
		var insertData []g.Map
		for _, contact := range targetContacts {
			if !existingTagContactIds[contact.Id] {
				insertData = append(insertData, g.Map{
					"contact_id":  contact.Id,
					"tag_id":      tagId,
					"create_time": now,
				})
			}
		}

		if len(insertData) > 0 {
			// 批量插入标签关系
			_, err = tx.Model("bm_contact_tags").InsertIgnore(insertData)
			if err != nil {
				g.Log().Error(ctx, "Failed to insert contact tags: %v", err)
				errors += len(insertData)
			} else {
				processed += len(insertData)
				g.Log().Debug(ctx, "Successfully tagged %d contacts in this batch", len(insertData))
			}
		}

		// 计算跳过的联系人数量（已有标签的 + 被排除的）
		skippedInBatch := len(allContacts) - len(targetContacts) + len(existingTags)
		skipped += skippedInBatch

		// 如果返回的联系人数量少于批量大小，说明已经处理完所有联系人
		if len(allContacts) < batchSize {
			break
		}

		offset += batchSize
	}

	g.Log().Infof(ctx, "batchAddTagsExclude completed: processed=%d, skipped=%d, errors=%d", processed, skipped, errors)
	return processed, skipped, errors, nil
}

// parseEmailList 解析邮箱列表
func parseEmailList(data string) []string {
	if data == "" {
		return nil
	}

	// 分割邮箱，支持换行符、逗号等分隔符
	lines := strings.FieldsFunc(data, func(r rune) bool {
		return r == '\n' || r == '\r' || r == ',' || r == ';'
	})

	var emails []string
	emailMap := make(map[string]bool) // 用于去重

	for _, line := range lines {
		email := strings.TrimSpace(line)
		if email != "" && !emailMap[email] {
			// 简单的邮箱格式验证
			if strings.Contains(email, "@") && strings.Contains(email, ".") {
				emails = append(emails, strings.ToLower(email))
				emailMap[email] = true
			}
		}

	}

	return emails
}
