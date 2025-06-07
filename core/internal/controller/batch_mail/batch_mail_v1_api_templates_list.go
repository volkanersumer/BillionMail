package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
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

	// 获取每个API的发送统计
	for _, item := range list {
		// 统计发送数量
		count, err := g.DB().Model("api_mail_logs").
			Where("api_id", item.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "统计API发送数量失败:", err)
			continue
		}
		// 可以添加到ApiTemplates结构体中显示
		g.Log().Debug(ctx, "API:", item.ApiName, "发送数量:", count)

		item.SendCount = count
	}

	res.Data.Total = total
	res.Data.List = list
	res.SetSuccess("获取API列表成功")
	return res, nil
}
