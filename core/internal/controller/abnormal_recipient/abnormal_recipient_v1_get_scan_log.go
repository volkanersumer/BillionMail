package abnormal_recipient

import (
	"billionmail-core/api/abnormal_recipient/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"os"
	"strings"
)

func (c *ControllerV1) GetScanLog(ctx context.Context, req *v1.GetScanLogReq) (res *v1.GetScanLogRes, err error) {
	var logDir = public.HostWorkDir + "/logs/core/check_email_valid.txt"

	res = &v1.GetScanLogRes{}

	content, err := os.ReadFile(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			res.Data = ""
			res.SetSuccess(public.LangCtx(ctx, "Success"))
			return
		}

		if strings.Contains(err.Error(), "no such file or directory") {
			res.Data = ""
			res.SetSuccess(public.LangCtx(ctx, "Success"))
			return
		}
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to read log file: : {}", err.Error())))
		return
	}

	res.Data = string(content)
	res.SetSuccess(public.LangCtx(ctx, "Success"))
	return
}
