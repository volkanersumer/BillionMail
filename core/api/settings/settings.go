// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package settings

import (
	"context"

	"billionmail-core/api/settings/v1"
)

type ISettingsV1 interface {
	GetVersion(ctx context.Context, req *v1.GetVersionReq) (res *v1.GetVersionRes, err error)
}
