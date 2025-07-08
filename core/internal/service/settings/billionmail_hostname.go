package settings

import (
	v1 "billionmail-core/api/domains/v1"
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"strings"
)

// Check whether the A record of BILLIONMAIL_HOSTNAME is consistent with the local public network IP and write or delete the tag file
func CheckHostname() {
	ctx := context.Background()
	g.Log().Info(ctx, "billionmail_hostname  start 111111111")
	hostname := public.MustGetDockerEnv("BILLIONMAIL_HOSTNAME", "")
	flagFile := public.AbsPath("../core/data/billionmail_hostname.txt")

	// ----------------------------------------------------------------------
	if gfile.Exists(flagFile) {
		content := gfile.GetContents(flagFile)
		g.Log().Info(ctx, "flagFile content 111111111: ", content)
	} else {
		g.Log().Info(ctx, "flagFile not exists 111111111")
	}
	// ----------------------------------------------------------------------

	if hostname == "" || hostname == "mail.example.com" {
		_ = os.Remove(flagFile)
		g.Log().Info(ctx, "No hostname 111111111")
		return
	}

	serverIP, err := public.GetServerIP()

	if err != nil {
		g.Log().Info(ctx, "No serverIP 111111111")
		return
	}

	recordType := "A"
	if strings.Contains(serverIP, ":") {
		recordType = "AAAA"
	}

	record := v1.DNSRecord{
		Type:  recordType,
		Host:  hostname,
		Value: serverIP,
	}

	Valid := domains.ValidateARecord(record)
	g.Log().Info(ctx, "Valid   111111111", Valid)
	if Valid {
		_ = os.WriteFile(flagFile, []byte(hostname), 0644)
	}

}
