package v1

import (
	"billionmail-core/internal/model/entity"
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type TagInfo struct {
	Id         int    `json:"id"          dc:"Tag ID"`
	GroupId    int    `json:"group_id"    dc:"Group ID"`
	Name       string `json:"name"       dc:"Tag Name"`
	CreateTime int    `json:"create_time" dc:"Create Time"`
}

// 获取标签列表
type TagListReq struct {
	g.Meta        `path:"/tags/list" method:"get" tags:"Tag" summary:"Get tag list"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" v:"required|min:1" dc:"Group ID"`
	Keyword       string `json:"keyword" dc:"Search keyword"`
	Page          int    `json:"page" dc:"Page" d:"1"`
	PageSize      int    `json:"page_size" dc:"Page size" d:"20"`
}
type TagListRes struct {
	api_v1.StandardRes
	Data struct {
		Total int                  `json:"total" dc:"Total Count"`
		List  []*entity.Tag `json:"list" dc:"Tag list"`
	} `json:"data"`
}

// 创建标签
type TagCreateReq struct {
	g.Meta        `path:"/tags/create" method:"post" tags:"Tag" summary:"Create a new tag"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" v:"required|min:1" dc:"Group ID"`
	Name          string `json:"name" v:"required|max-length:50" dc:"Tag Name"`
}
type TagCreateRes struct {
	api_v1.StandardRes
}

// 更新标签
type TagUpdateReq struct {
	g.Meta        `path:"/tags/update" method:"post" tags:"Tag" summary:"Update a tag"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Id            int    `json:"id" v:"required|min:1" dc:"Tag ID"`
	Name          string `json:"name" v:"required|max-length:50" dc:"Tag Name"`
}
type TagUpdateRes struct {
	api_v1.StandardRes
}

// 删除标签
type TagDeleteReq struct {
	g.Meta        `path:"/tags/delete" method:"post" tags:"Tag" summary:"Delete a tag"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Id            int    `json:"id" v:"required|min:1" dc:"Tag ID"`
}
type TagDeleteRes struct {
	api_v1.StandardRes
}


// 批量标记联系人  传 group_id  tag_id  data(联系人邮箱 换行分隔)  mark_include(标记包含1,标记不包含0)
type BatchTagContactsReq struct {
	g.Meta        `path:"/tags/batch_contacts" method:"post" tags:"Tag" summary:"Tag contacts in bulk"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" v:"required|min:1" dc:"Group ID"`
	TagId         int    `json:"tag_id" v:"required|min:1" dc:"Tag ID"`
	Data          string `json:"data" v:"required" dc:"Contact emails, separated by new lines"`
	MarkInclude   int    `json:"mark_include" v:"required|in:0,1" dc:"Mark include (1) or exclude (0)"`
}
type BatchTagContactsRes struct {
	api_v1.StandardRes
}

