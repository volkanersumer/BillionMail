package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type GetVersionReq struct {
	g.Meta `path:"/settings/get_version" tags:"Version" method:"get" summary:"Get current version information" in:"query"`
}

type GetVersionRes struct {
	api_v1.StandardRes
	Data struct {
		Version       string `json:"version" dc:"Version"`
		LatestVersion string `json:"latest_version" dc:"Latest version"`
		ReleaseDate   string `json:"release_date" dc:"Release date"`
		CanUpdate     bool   `json:"can_update" dc:"CanUpdate"`
	} `json:"data" dc:"Data"`
}
