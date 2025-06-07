package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) ApiTemplatesList(ctx context.Context, req *v1.ApiTemplatesListReq) (res *v1.ApiTemplatesListRes, err error) {
	res = &v1.ApiTemplatesListRes{}

	// 构建查询条件
	model := g.DB().Model("api_templates").Safe()

	// 添加api_name模糊搜索
	if req.Keyword != "" {
		model = model.WhereLike("api_name", "%"+req.Keyword+"%")
	}
	// active筛选
	if req.Active != -1 {
		model = model.Where("active", req.Active)
	}

	if req.StartTime > 0 && req.EndTime < 0 {
		req.EndTime = int(time.Now().Unix())
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
		query = query.Where("aml.api_id", item.Id)

		if req.StartTime > 0 {
			query.Where("sm.log_time_millis > ?", req.StartTime*1000-1)
		}

		if req.EndTime > 0 {
			query.Where("sm.log_time_millis < ?", req.EndTime*1000+1)
		}

		// 统计各项数据
		query.Fields(
			"count(*) as sends",
			"coalesce(sum(case when sm.status='sent' and sm.dsn like '2.%' then 1 else 0 end), 0) as delivered",
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

		// 统计打开、点击（直接用campaign_id）
		apiCampaignId := item.Id + 1000000000
		openedCount, _ := g.DB().Model("mailstat_opened").
			Where("campaign_id", apiCampaignId).
			WhereGTE("log_time_millis", req.StartTime*1000).
			WhereLTE("log_time_millis", req.EndTime*1000).
			Fields("count(distinct postfix_message_id) as opened").
			Value()
		clickedCount, _ := g.DB().Model("mailstat_clicked").
			Where("campaign_id", apiCampaignId).
			WhereGTE("log_time_millis", req.StartTime*1000).
			WhereLTE("log_time_millis", req.EndTime*1000).
			Fields("count(distinct postfix_message_id) as clicked").
			Value()

		if item.SendCount > 0 {
			item.DeliveryRate = public.Round(float64(item.SuccessCount)/float64(item.SendCount)*100, 2)
			item.BounceRate = public.Round(float64(item.FailCount)/float64(item.SendCount)*100, 2)
			item.OpenRate = public.Round(float64(openedCount.Int())/float64(item.SendCount)*100, 2)
			item.ClickRate = public.Round(float64(clickedCount.Int())/float64(item.SendCount)*100, 2)
		} else {
			item.DeliveryRate = 0
			item.BounceRate = 0
			item.OpenRate = 0
			item.ClickRate = 0
		}

		// 统计退订数量
		recipients := []string{}
		_, err = g.DB().Model("api_mail_logs").Where("api_id", item.Id).Fields("recipient").Array(&recipients)
		unsubscribeCount := 0
		if len(recipients) > 0 {
			unsubscribeCount, _ = g.DB().Model("bm_contacts").
				Where("email", recipients).
				Where("active", 0).
				WhereGTE("create_time", item.CreateTime).
				Count()
		}
		item.UnsubscribeCount = unsubscribeCount
	}

	// 汇总全局统计
	//var totalSend, totalDelivered, totalBounced, totalOpened, totalClicked, totalUnsub int
	//for _, item := range list {
	//	totalSend += item.SendCount
	//	totalDelivered += item.SuccessCount
	//	totalBounced += item.FailCount
	//	totalOpened += int(item.OpenRate * float64(item.SendCount) / 100)
	//	totalClicked += int(item.ClickRate * float64(item.SendCount) / 100)
	//	totalUnsub += item.UnsubscribeCount
	//}
	//var avgDeliveryRate, avgOpenRate, avgClickRate, avgBounceRate, avgUnsubRate float64
	//if totalSend > 0 {
	//	avgDeliveryRate = public.Round(float64(totalDelivered)/float64(totalSend)*100, 2)
	//	avgOpenRate = public.Round(float64(totalOpened)/float64(totalSend)*100, 2)
	//	avgClickRate = public.Round(float64(totalClicked)/float64(totalSend)*100, 2)
	//	avgBounceRate = public.Round(float64(totalBounced)/float64(totalSend)*100, 2)
	//	avgUnsubRate = public.Round(float64(totalUnsub)/float64(totalSend)*100, 2)
	//}
	//res.Data.Summary = map[string]interface{}{
	//	"total_send": totalSend,
	//	"avg_delivery_rate": avgDeliveryRate,
	//	"avg_open_rate": avgOpenRate,
	//	"avg_click_rate": avgClickRate,
	//	"avg_bounce_rate": avgBounceRate,
	//	"avg_unsub_rate": avgUnsubRate,
	//	"total_unsubscribe": totalUnsub,
	//}

	res.Data.Total = total
	res.Data.List = list
	res.SetSuccess("获取API列表成功")
	return res, nil
}
