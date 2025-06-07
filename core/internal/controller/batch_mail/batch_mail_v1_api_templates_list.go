package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ApiTemplatesList(ctx context.Context, req *v1.ApiTemplatesListReq) (res *v1.ApiTemplatesListRes, err error) {
	res = &v1.ApiTemplatesListRes{}

	// 构建查询条件
	model := g.DB().Model("api_templates").Safe()

	// 添加搜索条件
	if req.Keyword != "" {
		model = model.WhereLike("api_name", "%"+req.Keyword+"%")
	}

	// 获取总数
	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []*v1.ApiTemplatesInfo
	err = model.Page(req.Page, req.PageSize).
		OrderDesc("id").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		// 构建基础查询
		query := g.DB().Model("api_mail_logs aml")

		query = query.LeftJoin("mailstat_message_ids mi", "aml.message_id=mi.message_id")
		query = query.LeftJoin("mailstat_send_mails sm", "mi.postfix_message_id=sm.postfix_message_id")
		query = query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")
		query = query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")
		query = query.Where("aml.api_id", item.Id)

		// 统计各项数据
		query.Fields(
			"count(*) as sends",
			"coalesce(sum(case when sm.status='sent' and sm.dsn like '2.%' then 1 else 0 end), 0) as delivered",
			"count(distinct o.postfix_message_id) as opened",
			"count(distinct c.postfix_message_id) as clicked",
			"coalesce(sum(case when sm.status='bounced' then 1 else 0 end), 0) as bounced",
		)

		result, err := query.One()
		if err != nil {
			g.Log().Error(ctx, "统计API发送数据失败:", err)
			continue
		}

		// 填充统计数据
		item.SendCount = result["sends"].Int()
		item.SuccessCount = result["delivered"].Int()
		item.FailCount = result["bounced"].Int()

		if item.SendCount > 0 {
			// 送达率 = 成功送达数/总发送数
			item.DeliveryRate = public.Round(float64(item.SuccessCount)/float64(item.SendCount)*100, 2)
			// 退信率 = 退信数/总发送数
			item.BounceRate = public.Round(float64(item.FailCount)/float64(item.SendCount)*100, 2)
			// 打开率 = 打开数/总发送数
			item.OpenRate = public.Round(float64(result["opened"].Int())/float64(item.SendCount)*100, 2)
			// 点击率 = 点击数/总发送数
			item.ClickRate = public.Round(float64(result["clicked"].Int())/float64(item.SendCount)*100, 2)
		} else {
			item.DeliveryRate = 0
			item.BounceRate = 0
			item.OpenRate = 0
			item.ClickRate = 0
		}
	}

	res.Data.Total = total
	res.Data.List = list
	res.SetSuccess("获取API列表成功")
	return res, nil
}
