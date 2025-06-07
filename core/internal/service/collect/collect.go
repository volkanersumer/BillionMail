package collect

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func Collect(ctx context.Context) {
	g.Log().Debug(ctx, "Collecting mail sent and relay counts")

	apiBase := "https://www.aapanel.com/api/panel/submit_feature_invoked"

	// Get last post time
	var lastPostTime int64
	_ = public.OptionsMgrInstance.GetOption(ctx, "last_post_time", &lastPostTime)

	// Get current time
	curTime := time.Now().Unix()

	if lastPostTime == 0 || curTime-lastPostTime > 24*60*60 {
		lastPostTime = curTime
	}

	defer func() {
		// Update last post time
		_ = public.OptionsMgrInstance.SetOption(ctx, "last_post_time", curTime)
	}()

	var cnt int

	// Get sent total count
	cnt, _ = g.DB().Model("mailstate_send_mails").Where("log_time > ?", lastPostTime).Count()

	resp, err := g.Client().ContentJson().Post(ctx, apiBase, g.Map{
		"feature": "bm.sent",
		"cnt":     cnt,
	})

	if err != nil {
		g.Log().Warning(ctx, "Failed to post sent count: ", err)
	}

	g.Log().Debug(ctx, "sent count response: ", string(resp.ReadAll()))

	// Get relay count
	cnt, _ = g.DB().Model("bm_relay").Where("create_time > ?", lastPostTime).Count()

	resp, err = g.Client().ContentJson().Post(ctx, apiBase, g.Map{
		"feature": "bm.relays",
		"cnt":     cnt,
	})

	if err != nil {
		g.Log().Warning(ctx, "Failed to post relay count: ", err)
	}

	g.Log().Debug(ctx, "relay count response: ", string(resp.ReadAll()))
}
