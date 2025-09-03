package tags

import (
	"billionmail-core/api/tags/v1"
	"context"
)

func (c *ControllerV1) TagList(ctx context.Context, req *v1.TagListReq) (res *v1.TagListRes, err error) {
	res = &v1.TagListRes{}

	//// 验证联系人组是否存在
	//var groupCount int
	//groupCount, err = g.DB().Model("bm_contact_groups").Where("id", req.GroupId).Count()
	//if err != nil {
	//	res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check group")))
	//	return
	//}
	//if groupCount == 0 {
	//	res.SetError(gerror.New(public.LangCtx(ctx, "Group not found")))
	//	return
	//}

	// 构建查询条件
	//model := g.DB().Model("bm_tags t").
	//	LeftJoin("bm_contact_tags ct", "t.id = ct.tag_id").
	//	Fields("t.*, COUNT(ct.contact_id) as contact_count").
	//	Where("t.group_id", req.GroupId).
	//	Group("t.id")
	//
	//if req.Keyword != "" {
	//	model = model.WhereLike("t.name", "%"+req.Keyword+"%")
	//}
	//
	//// 获取总数
	//total, err := g.DB().Model("bm_tags").
	//	Where("group_id", req.GroupId).
	//	WhereLike("name", "%"+req.Keyword+"%").
	//	Count()
	//if err != nil {
	//	res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get total count")))
	//	return
	//}
	//
	//// 分页查询
	//var tags []*v1.Tag
	//err = model.
	//	Order("t.create_time DESC").
	//	Page(req.Page, req.PageSize).
	//	Scan(&tags)
	//if err != nil {
	//	res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get tag list")))
	//	return
	//}
	//
	//res.Data.Total = total
	//res.Data.List = tags
	//res.SetSuccess(public.LangCtx(ctx, "Get tag list successfully"))
	return
}
