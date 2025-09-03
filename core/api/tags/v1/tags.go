package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type TagInfo struct {
	Id          int    `json:"id"          dc:"Tag ID"`
	GroupId     int    `json:"group_id"    dc:"Group ID"`
	Name        string `json:"name"       dc:"Tag Name"`
	Color       string `json:"color"      dc:"Tag Color"`
	Description string `json:"description"       dc:"Tag description"`
	CreateTime  int    `json:"create_time" dc:"Create Time"`
}

// 获取标签列表
type TagListReq struct {
	g.Meta        `path:"/tags/list" method:"get" tags:"Tag" summary:"Get tag list"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" v:"required|min:1" dc:"Group ID"`
	Keyword       string `json:"keyword" dc:"Search keyword"`
}
type TagListRes struct {
	api_v1.StandardRes
	Data []*TagInfo `json:"data" dc:"Tag list"`
}

// 创建标签
type TagCreateReq struct {
	g.Meta        `path:"/tags/create" method:"post" tags:"Tag" summary:"Create a new tag"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" v:"required|min:1" dc:"Group ID"`
	Name          string `json:"name" v:"required|max-length:50" dc:"Tag Name"`
	Color         string `json:"color" v:"max-length:20" dc:"Tag Color"`
	Description   string `json:"description" v:"max-length:255" dc:"Tag description"`
}
type TagCreateRes struct {
	api_v1.StandardRes
}

// 更新标签
type TagUpdateReq struct {
	g.Meta        `path:"/tags/update" method:"put" tags:"Tag" summary:"Update a tag"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Id            int    `json:"id" v:"required|min:1" dc:"Tag ID"`
	Name          string `json:"name" v:"required|max-length:50" dc:"Tag Name"`
	Color         string `json:"color" v:"max-length:20" dc:"Tag Color"`
	Description   string `json:"description" v:"max-length:255" dc:"Tag description"`
}
type TagUpdateRes struct {
	api_v1.StandardRes
}

// 删除标签
type TagDeleteReq struct {
	g.Meta        `path:"/tags/delete" method:"delete" tags:"Tag" summary:"Delete a tag"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Id            int    `json:"id" v:"required|min:1" dc:"Tag ID"`
}
type TagDeleteRes struct {
	api_v1.StandardRes
}
