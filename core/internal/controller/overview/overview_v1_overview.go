package overview

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	random "math/rand"
	"time"

	"billionmail-core/api/overview/v1"
)

func (c *ControllerV1) Overview(ctx context.Context, req *v1.OverviewReq) (res *v1.OverviewRes, err error) {
	res = &v1.OverviewRes{}

	res.SetSuccess("Success")

	columnType := "hourly"

	if req.EndTime == 0 {
		req.EndTime = time.Now().Unix()
	}

	if req.EndTime < req.StartTime {
		err = fmt.Errorf("start time must be greater or equal to end time")
		return
	}

	// Simulate base values for million-level data
	totalSends := 1200000 + random.Intn(300000)
	// Control bounce rate between 0.1% and 1%
	bounceRateBase := 0.1 + random.Float64()*0.9 // Between 0.1% and 1%
	bounceCount := int(float64(totalSends) * bounceRateBase / 100)
	// Calculate other metrics
	deliveredCount := totalSends - bounceCount
	deliveryRate := float64(deliveredCount) / float64(totalSends) * 100
	failureRate := 100.0 - deliveryRate

	// Simulate click and open data
	openCount := int(float64(deliveredCount) * (15 + random.Float64()*10) / 100) // 15-25% open rate
	clickCount := int(float64(openCount) * (5 + random.Float64()*5) / 100)       // 5-10% click rate (based on opens)
	openRate := float64(openCount) / float64(totalSends) * 100
	clickRate := float64(clickCount) / float64(totalSends) * 100

	// fake chart map
	m := map[string]map[string][]map[string]interface{}{
		"daily": {
			"send_mail":   {},
			"bounce_rate": {},
			"open_rate":   {},
			"click_rate":  {},
		},
		"hourly": {
			"send_mail":   make([]map[string]interface{}, 24),
			"bounce_rate": make([]map[string]interface{}, 24),
			"open_rate":   make([]map[string]interface{}, 24),
			"click_rate":  make([]map[string]interface{}, 24),
		},
	}

	// Generate hourly data distribution
	hourlyDistribution := make([]float64, 24)
	totalDistribution := 0.0

	// Create a more realistic distribution: more emails sent between 9AM and 6PM
	for i := 0; i < 24; i++ {
		if i >= 9 && i <= 18 {
			// Higher distribution during working hours
			hourlyDistribution[i] = 3.0 + random.Float64()*2.0
		} else {
			// Lower during non-working hours
			hourlyDistribution[i] = 0.5 + random.Float64()*1.0
		}
		totalDistribution += hourlyDistribution[i]
	}

	// Normalize distribution and generate hourly data
	for i := 0; i < 24; i++ {
		hourlyWeight := hourlyDistribution[i] / totalDistribution
		hourlySends := int(float64(totalSends) * hourlyWeight)

		// Generate slightly different bounce_rate for each hour, but still within 0.1-1% range
		hourlyBounceRate := bounceRateBase * (0.8 + random.Float64()*0.4) // 80%-120% of base bounce rate
		hourlyBounces := int(float64(hourlySends) * hourlyBounceRate / 100)
		hourlyDelivered := hourlySends - hourlyBounces

		hourlyOpenRate := openRate * (0.9 + random.Float64()*0.2)   // 90%-110% of base open rate
		hourlyClickRate := clickRate * (0.9 + random.Float64()*0.2) // 90%-110% of base click rate

		m["hourly"]["send_mail"][i] = map[string]interface{}{
			"delivered": hourlyDelivered,
			"failed":    hourlyBounces,
			"sends":     hourlySends,
			"x":         i,
		}
		m["hourly"]["bounce_rate"][i] = map[string]interface{}{
			"bounce_rate": public.Round(hourlyBounceRate, 3),
			"x":           i,
		}
		m["hourly"]["open_rate"][i] = map[string]interface{}{
			"open_rate": public.Round(hourlyOpenRate, 2),
			"x":         i,
		}
		m["hourly"]["click_rate"][i] = map[string]interface{}{
			"click_rate": public.Round(hourlyClickRate, 2),
			"x":          i,
		}
	}

	if req.EndTime-req.StartTime > 86400 {
		columnType = "daily"
		dayUnixTimes := public.GenDayUnixTimestampByTimeSection(req.StartTime, req.EndTime)
		dayCount := len(dayUnixTimes)

		// Calculate average sends per day
		dailyAvgSends := totalSends / dayCount

		m["daily"]["send_mail"] = make([]map[string]interface{}, dayCount)
		m["daily"]["bounce_rate"] = make([]map[string]interface{}, dayCount)
		m["daily"]["open_rate"] = make([]map[string]interface{}, dayCount)
		m["daily"]["click_rate"] = make([]map[string]interface{}, dayCount)

		for i, dayUnixTime := range dayUnixTimes {
			// Generate slightly different data for each day
			dailyFactor := 0.85 + random.Float64()*0.3 // Daily factor between 85%-115%
			dailySends := int(float64(dailyAvgSends) * dailyFactor)

			// Control bounce rate between 0.1% and 1%
			dailyBounceRate := bounceRateBase * (0.9 + random.Float64()*0.2) // 90%-110% of base bounce rate
			dailyBounces := int(float64(dailySends) * dailyBounceRate / 100)
			dailyDelivered := dailySends - dailyBounces

			dailyDeliveryRate := float64(dailyDelivered) / float64(dailySends) * 100
			dailyFailureRate := 100.0 - dailyDeliveryRate
			dailyOpenRate := openRate * (0.95 + random.Float64()*0.1)   // 95%-105% of base open rate
			dailyClickRate := clickRate * (0.95 + random.Float64()*0.1) // 95%-105% of base click rate

			m["daily"]["send_mail"][i] = map[string]interface{}{
				"delivered":     dailyDelivered,
				"delivery_rate": public.Round(dailyDeliveryRate, 2),
				"failed":        dailyBounces,
				"failure_rate":  public.Round(dailyFailureRate, 2),
				"sends":         dailySends,
				"x":             dayUnixTime,
			}
			m["daily"]["bounce_rate"][i] = map[string]interface{}{
				"bounce_rate": public.Round(dailyBounceRate, 3),
				"x":           dayUnixTime,
			}
			m["daily"]["open_rate"][i] = map[string]interface{}{
				"open_rate": public.Round(dailyOpenRate, 2),
				"x":         dayUnixTime,
			}
			m["daily"]["click_rate"][i] = map[string]interface{}{
				"click_rate": public.Round(dailyClickRate, 2),
				"x":          dayUnixTime,
			}
		}
	}

	// Set dashboard data, showing million-level total and 0.1-1% bounce rate
	res.Data.Dashboard = map[string]interface{}{
		"bounce_rate":   public.Round(bounceRateBase, 3),
		"bounced":       bounceCount,
		"click_rate":    public.Round(clickRate, 2),
		"clicked":       clickCount,
		"delivered":     deliveredCount,
		"delivery_rate": public.Round(deliveryRate, 2),
		"open_rate":     public.Round(openRate, 2),
		"opened":        openCount,
		"sends":         totalSends,
		"failure_rate":  public.Round(failureRate, 2),
		"failed":        bounceCount,
		"delay_queue":   0, // Placeholder for delay queue
	}

	res.Data.SendMailChart = map[string]interface{}{
		"column_type": columnType,
		"dashboard": map[string]interface{}{
			"delivered":     deliveredCount,
			"delivery_rate": public.Round(deliveryRate, 2),
			"failed":        bounceCount,
			"failure_rate":  public.Round(failureRate, 2),
			"sends":         totalSends,
		},
		"data": m[columnType]["send_mail"],
	}

	res.Data.BounceRateChart = map[string]interface{}{
		"column_type": columnType,
		"data":        m[columnType]["bounce_rate"],
	}

	res.Data.OpenRateChart = map[string]interface{}{
		"column_type": columnType,
		"data":        m[columnType]["open_rate"],
	}

	res.Data.ClickRateChart = map[string]interface{}{
		"column_type": columnType,
		"data":        m[columnType]["click_rate"],
	}

	// Simulate data distribution for major mail providers
	mailProviderDistribution := map[string]float64{
		"google":  0.45, // Google accounts for 45%
		"other":   0.20, // Other mail providers account for 20%
		"outlook": 0.18, // Outlook accounts for 18%
		"yahoo":   0.12, // Yahoo accounts for 12%
		"local":   0.05, // Local mail services account for 5%
	}

	res.Data.MailProviders = make([]map[string]interface{}, 0, len(mailProviderDistribution))

	for provider, ratio := range mailProviderDistribution {
		providerSends := int(float64(totalSends) * ratio)
		providerBounceRate := bounceRateBase * (0.8 + random.Float64()*0.4) // Different bounce rates per provider
		providerBounces := int(float64(providerSends) * providerBounceRate / 100)
		providerDelivered := providerSends - providerBounces
		providerDeliveryRate := float64(providerDelivered) / float64(providerSends) * 100

		providerOpenRate := openRate * (0.9 + random.Float64()*0.2)
		providerOpened := int(float64(providerDelivered) * providerOpenRate / 100)

		providerClickRate := clickRate * (0.9 + random.Float64()*0.2)
		providerClicked := int(float64(providerOpened) * providerClickRate / openRate)

		res.Data.MailProviders = append(res.Data.MailProviders.([]map[string]interface{}), map[string]interface{}{
			"bounce_rate":   public.Round(providerBounceRate, 3),
			"bounced":       providerBounces,
			"click_rate":    public.Round(providerClickRate, 2),
			"clicked":       providerClicked,
			"delivered":     providerDelivered,
			"delivery_rate": public.Round(providerDeliveryRate, 2),
			"mail_provider": provider,
			"open_rate":     public.Round(providerOpenRate, 2),
			"opened":        providerOpened,
			"sends":         providerSends,
		})
	}

	return
}
