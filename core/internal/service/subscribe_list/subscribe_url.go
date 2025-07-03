package subscribe_list

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"path/filepath"
	"strings"
)

func ReplaceSubmitUrl(ctx context.Context) {

	hostwork := public.HostWorkDir
	filePath := filepath.Join(hostwork, "/core/public/html/subscribe_form.html")
	content := gfile.GetContents(filePath)

	if !strings.Contains(content, "{{ SubmitURL . }}") {
		return
	}

	var submitUrl string
	hostUrl := public.GethostUrl()
	submitUrl = hostUrl + "/api/subscribe/submit"

	newContent := strings.ReplaceAll(content, "{{ SubmitURL . }}", submitUrl)

	err := gfile.PutContents(filePath, newContent)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to update subscribe form URL: %v", err)

	}

}
