package batch_mail

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// TaskStatService 任务统计服务
type TaskStatService struct{}

func NewTaskStatService() *TaskStatService {
	return &TaskStatService{}
}

// buildBaseQuery 构建基础查询
func (s *TaskStatService) buildBaseQuery(taskId int64, domain string, startTime, endTime int64) *gdb.Model {
	// 基础查询
	query := g.DB().Model("mailstat_send_mails sm")

	// 关联发件人表
	query.LeftJoin("mailstat_senders s", "sm.postfix_message_id=s.postfix_message_id")
	query.Where("s.postfix_message_id is not null")

	// 关联消息ID和任务收件人信息
	query.LeftJoin("mailstat_message_ids mi", "sm.postfix_message_id=mi.postfix_message_id")
	query.LeftJoin("recipient_info ri", "mi.message_id=ri.message_id")

	// 过滤指定任务
	query.Where("ri.task_id = ?", taskId)

	// 域名过滤
	if domain != "" {
		query.Where("sm.recipient LIKE ?", "%@"+domain)
	}

	// 时间范围过滤
	if startTime > 0 {
		query.Where("sm.log_time_millis > ?", startTime*1000-1)
	}

	if endTime > 0 {
		query.Where("sm.log_time_millis < ?", endTime*1000+1)
	}

	return query
}

// GetTaskStatChart 获取任务统计图表
func (s *TaskStatService) GetTaskStatChart(taskId int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = s.filterAndPrepareTimeSection(startTime, endTime)

	return map[string]interface{}{
		"dashboard":         s.getTaskDashboard(taskId, domain, startTime, endTime),
		"mail_providers":    s.getTaskMailProviders(taskId, domain, startTime, endTime),
		"send_mail_chart":   s.getTaskSendMailChart(taskId, domain, startTime, endTime),
		"bounce_rate_chart": s.getTaskBounceRateChart(taskId, domain, startTime, endTime),
		"open_rate_chart":   s.getTaskOpenRateChart(taskId, domain, startTime, endTime),
		"click_rate_chart":  s.getTaskClickRateChart(taskId, domain, startTime, endTime),
	}
}

// filterAndPrepareTimeSection 处理时间区间
func (s *TaskStatService) filterAndPrepareTimeSection(startTime, endTime int64) (int64, int64) {
	if startTime > 0 && endTime < 0 {
		endTime = time.Now().Unix()
	}

	if startTime > endTime {
		panic(public.Lang("end_time must greater than start_time"))
	}

	return startTime, endTime
}

// prepareChartData 准备图表数据
func (s *TaskStatService) prepareChartData(startTime, endTime int64) (string, string) {
	columnType := "daily"
	secs := endTime - startTime
	xAxisField := "to_char(to_timestamp(sm.log_time_millis / 1000), 'YYYY-MM-DD') as x"

	if secs < 86400 {
		columnType = "hourly"
		xAxisField = "to_char(to_timestamp(sm.log_time_millis / 1000), 'HH24')::integer as x"
	}

	return columnType, xAxisField
}

// getTaskDashboard 获取仪表盘数据
func (s *TaskStatService) getTaskDashboard(taskId int64, domain string, startTime, endTime int64) map[string]interface{} {
	query := s.buildBaseQuery(taskId, domain, startTime, endTime)

	query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")
	query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")

	query.Fields(
		"count(*) as sends",
		"coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered",
		"count(distinct o.postfix_message_id) as opened",
		"count(distinct c.postfix_message_id) as clicked",
		"coalesce(sum(case when status='bounced' then 1 else 0 end), 0) as bounced",
	)

	result, err := query.One()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	stats := result.Map()
	sends := result["sends"].Int()

	if sends > 0 {
		stats["delivery_rate"] = public.Round(float64(result["delivered"].Int())/float64(sends)*100, 2)
		stats["bounce_rate"] = public.Round(float64(result["bounced"].Int())/float64(sends)*100, 2)
		stats["open_rate"] = public.Round(float64(result["opened"].Int())/float64(sends)*100, 2)
		stats["click_rate"] = public.Round(float64(result["clicked"].Int())/float64(sends)*100, 2)
	} else {
		stats["delivery_rate"] = 0.0
		stats["bounce_rate"] = 0.0
		stats["open_rate"] = 0.0
		stats["click_rate"] = 0.0
	}

	return stats
}

// fillChartData 填充图表数据
func (s *TaskStatService) fillChartData(data []map[string]interface{}, columnType string, startTime, endTime int64) []map[string]interface{} {
	switch columnType {
	case "daily":
		return s.fillChartDataDaily(data, startTime, endTime)
	case "hourly":
		return s.fillChartDataHourly(data)
	default:
		return data
	}
}

// fillChartDataHourly 填充小时数据
func (s *TaskStatService) fillChartDataHourly(data []map[string]interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 24)
	dataMap := make(map[int]map[string]interface{})

	for _, item := range data {
		hour := item["x"].(int)
		dataMap[hour] = item
	}

	for i := 0; i < 24; i++ {
		if item, exists := dataMap[i]; exists {
			result[i] = item
		} else {
			result[i] = map[string]interface{}{
				"x":         i,
				"sends":     0,
				"delivered": 0,
				"failed":    0,
			}
		}
	}

	return result
}

// fillChartDataDaily 填充每日数据
func (s *TaskStatService) fillChartDataDaily(data []map[string]interface{}, startTime, endTime int64) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	dataMap := make(map[string]map[string]interface{})

	for _, item := range data {
		date := item["x"].(string)
		dataMap[date] = item
	}

	for t := startTime; t <= endTime; t += 86400 {
		date := time.Unix(t, 0).Format("2006-01-02")
		if item, exists := dataMap[date]; exists {
			result = append(result, item)
		} else {
			result = append(result, map[string]interface{}{
				"x":         date,
				"sends":     0,
				"delivered": 0,
				"failed":    0,
			})
		}
	}

	return result
}

// getTaskBounceRateChart 获取退信率图表
func (s *TaskStatService) getTaskBounceRateChart(taskId int64, domain string, startTime, endTime int64) map[string]interface{} {
	query := s.buildBaseQuery(taskId, domain, startTime, endTime)
	columnType, xAxisField := s.prepareChartData(startTime, endTime)

	query.Fields(
		xAxisField,
		"case when count(*) > 0 then round(1.0 * coalesce(sum(case when status='bounced' then 1 else 0 end), 0) / count(*) * 100.0, 2) else 0.0 end as bounce_rate",
	)

	query.Group("x")
	query.Order("x")

	results, err := query.All()
	if err != nil {
		return nil
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        s.fillChartData(results.List(), columnType, startTime, endTime),
	}
}

// getTaskOpenRateChart 获取打开率图表
func (s *TaskStatService) getTaskOpenRateChart(taskId int64, domain string, startTime, endTime int64) map[string]interface{} {
	query := s.buildBaseQuery(taskId, domain, startTime, endTime)
	query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")

	columnType, xAxisField := s.prepareChartData(startTime, endTime)

	query.Fields(
		xAxisField,
		"case when count(*) > 0 then round(1.0 * count(distinct o.postfix_message_id) / count(*) * 100.0, 2) else 0.0 end as open_rate",
	)

	query.Group("x")
	query.Order("x")

	results, err := query.All()
	if err != nil {
		return nil
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        s.fillChartData(results.List(), columnType, startTime, endTime),
	}
}

// getTaskClickRateChart 获取点击率图表
func (s *TaskStatService) getTaskClickRateChart(taskId int64, domain string, startTime, endTime int64) map[string]interface{} {
	query := s.buildBaseQuery(taskId, domain, startTime, endTime)
	query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")

	columnType, xAxisField := s.prepareChartData(startTime, endTime)

	query.Fields(
		xAxisField,
		"case when count(*) > 0 then round(1.0 * count(distinct c.postfix_message_id) / count(*) * 100.0, 2) else 0.0 end as click_rate",
	)

	query.Group("x")
	query.Order("x")

	results, err := query.All()
	if err != nil {
		return nil
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        s.fillChartData(results.List(), columnType, startTime, endTime),
	}
}

// getTaskMailProviders 获取邮件服务商统计
func (s *TaskStatService) getTaskMailProviders(taskId int64, domain string, startTime, endTime int64) []map[string]interface{} {
	query := s.buildBaseQuery(taskId, domain, startTime, endTime)
	query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")
	query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")

	query.Fields(
		"sm.mail_provider",
		"count(*) as sends",
		"coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered",
		"count(distinct o.postfix_message_id) as opened",
		"count(distinct c.postfix_message_id) as clicked",
		"coalesce(sum(case when status='bounced' then 1 else 0 end), 0) as bounced",
	)

	query.Group("sm.mail_provider")

	results, err := query.All()
	if err != nil {
		return nil
	}

	providers := make([]map[string]interface{}, 0)
	for _, result := range results {
		provider := result.Map()
		sends := result["sends"].Int()
		if sends > 0 {
			provider["delivery_rate"] = public.Round(float64(result["delivered"].Int())/float64(sends)*100, 2)
			provider["bounce_rate"] = public.Round(float64(result["bounced"].Int())/float64(sends)*100, 2)
			provider["open_rate"] = public.Round(float64(result["opened"].Int())/float64(sends)*100, 2)
			provider["click_rate"] = public.Round(float64(result["clicked"].Int())/float64(sends)*100, 2)
		}
		providers = append(providers, provider)
	}

	return providers
}

// getTaskSendMailChart 获取发送邮件图表数据
func (s *TaskStatService) getTaskSendMailChart(taskId int64, domain string, startTime, endTime int64) map[string]interface{} {
	query := s.buildBaseQuery(taskId, domain, startTime, endTime)

	// 根据时间跨度决定按小时还是按天统计
	columnType, xAxisField := s.prepareChartData(startTime, endTime)

	query.Fields(
		xAxisField,
		"count(*) as sends",
		"coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered",
		"count(*) - coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as failed",
	)

	query.Group("x")
	query.Order("x")

	results, err := query.All()
	if err != nil {
		return nil
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        s.fillChartData(results.List(), columnType, startTime, endTime),
	}
}
