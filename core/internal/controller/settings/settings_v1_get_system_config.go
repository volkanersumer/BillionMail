package settings

import (
	"billionmail-core/api/settings/v1"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetSystemConfig(ctx context.Context, req *v1.GetSystemConfigReq) (res *v1.GetSystemConfigRes, err error) {
	res = &v1.GetSystemConfigRes{}

	envMap, err := loadEnvFile()
	if err != nil {
		res.SetError(gerror.New("failed to load environment variables"))
		return res, err
	}

	config := convertEnvToConfig(envMap)

	sslInfo, err := loadSSLInfo()
	if err == nil {
		config.SSL = *sslInfo
	}

	// TODO: Whitelist ip

	res.Data = config
	return res, nil
}
