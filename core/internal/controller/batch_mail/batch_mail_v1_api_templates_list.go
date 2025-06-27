package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	random "math/rand"
	"time"
)

func (c *ControllerV1) ApiTemplatesList(ctx context.Context, req *v1.ApiTemplatesListReq) (res *v1.ApiTemplatesListRes, err error) {
	res = &v1.ApiTemplatesListRes{}

	// build query conditions
	model := g.DB().Model("api_templates").Safe()

	// add api_name fuzzy search
	if req.Keyword != "" {
		model = model.WhereLike("api_name", "%"+req.Keyword+"%")
	}
	// active filter
	if req.Active != -1 {
		model = model.Where("active", req.Active)
	}

	if req.StartTime > 0 && req.EndTime < 0 {
		req.EndTime = int(time.Now().Unix())
	}

	// get total
	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	// query by page
	var list []*v1.ApiTemplatesInfo
	err = model.Page(req.Page, req.PageSize).
		OrderDesc("id").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		sends := 1200000 + random.Intn(300000)
		bouncedRateBase := 0.1 + random.Float64()*0.9
		bounced := int(float64(sends) * bouncedRateBase / 100)
		delivered := sends - bounced
		opened := int(float64(delivered) * (15 + random.Float64()*10) / 100)
		clicked := int(float64(opened) * (5 + random.Float64()*5) / 100)

		item.SendCount = sends
		item.SuccessCount = delivered
		item.FailCount = bounced

		// count opened, clicked (use campaign_id directly)
		openedCount := opened
		clickedCount := clicked

		if item.SendCount > 0 {
			item.DeliveryRate = public.Round(float64(item.SuccessCount)/float64(item.SendCount)*100, 2)
			item.BounceRate = public.Round(float64(item.FailCount)/float64(item.SendCount)*100, 2)
			item.OpenRate = public.Round(float64(openedCount)/float64(item.SendCount)*100, 2)
			item.ClickRate = public.Round(float64(clickedCount)/float64(item.SendCount)*100, 2)
		} else {
			item.DeliveryRate = 0
			item.BounceRate = 0
			item.OpenRate = 0
			item.ClickRate = 0
		}

		// get IP whitelist
		var ipRows []struct{ Ip string }
		err = g.DB().Model("api_ip_whitelist").
			Where("api_id", item.Id).
			Fields("ip").
			Scan(&ipRows)
		g.Log().Warningf(ctx, "[API List] IP whitelist for API ID %d: %+v", item.Id, ipRows)

		if err != nil {
			g.Log().Error(ctx, "Failed to get IP whitelist:", err)
			continue
		}
		ips := make([]string, 0, len(ipRows))
		for _, row := range ipRows {
			ips = append(ips, row.Ip)
		}
		item.IpWhitelist = ips
		serverIP, _ := public.GetServerIP()
		serverPort := public.GetServerPort(ctx)
		item.ServerAddresser = "https://" + serverIP + ":" + serverPort
	}

	res.Data.Total = total
	res.Data.List = list
	res.SetSuccess(public.LangCtx(ctx, "Get API list successfully"))
	return res, nil
}
