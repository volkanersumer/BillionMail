package batch_mail

import (
	"billionmail-core/internal/service/public"
	"context"
	random "math/rand"

	"billionmail-core/api/batch_mail/v1"
)

func (c *ControllerV1) ApiOverviewStats(ctx context.Context, req *v1.ApiOverviewStatsReq) (res *v1.ApiOverviewStatsRes, err error) {
	res = &v1.ApiOverviewStatsRes{}

	var totalSend, totalDelivered, totalBounced, totalOpened, totalClicked int

	sends := 1200000 + random.Intn(300000)
	bouncedRateBase := 0.1 + random.Float64()*0.9
	bounced := int(float64(sends) * bouncedRateBase / 100)
	delivered := sends - bounced
	totalSend += sends
	totalDelivered += delivered
	totalBounced += bounced

	openedCount := int(float64(delivered) * (15 + random.Float64()*10) / 100)
	clickedCount := int(float64(openedCount) * (5 + random.Float64()*5) / 100)
	totalOpened += openedCount
	totalClicked += clickedCount

	var avgDeliveryRate, avgOpenRate, avgClickRate, avgBounceRate float64
	if totalSend > 0 {
		avgDeliveryRate = public.Round(float64(totalDelivered)/float64(totalSend)*100, 2)
		avgOpenRate = public.Round(float64(totalOpened)/float64(totalSend)*100, 2)
		avgClickRate = public.Round(float64(totalClicked)/float64(totalSend)*100, 2)
		avgBounceRate = public.Round(float64(totalBounced)/float64(totalSend)*100, 2)
		//avgUnsubRate = public.Round(float64(totalUnsub)/float64(totalSend)*100, 2)
	}

	res.Data = v1.ApiSummaryStats{
		TotalSend:       totalSend,
		AvgDeliveryRate: avgDeliveryRate,
		AvgOpenRate:     avgOpenRate,
		AvgClickRate:    avgClickRate,
		AvgBounceRate:   avgBounceRate,
	}
	res.SetSuccess(public.LangCtx(ctx, "Statistics successful"))
	return res, nil
}
