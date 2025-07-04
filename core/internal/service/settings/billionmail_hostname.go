package settings

import (
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/public"
	"os"
)

// 检查BILLIONMAIL_HOSTNAME的A记录是否与本机公网IP一致，写入或删除标记文件
func CheckHostname() {
	hostname := public.MustGetDockerEnv("BILLIONMAIL_HOSTNAME", "")
	flagFile := public.AbsPath("../core/data/billionmail_hostname.txt")

	if hostname == "" || hostname == "mail.example.com" {
		_ = os.Remove(flagFile)
		return
	}

	record, err := domains.GetARecord(hostname, true)
	if err != nil || !record.Valid {
		_ = os.Remove(flagFile)
		return
	}

	serverIP := record.Value
	if serverIP == "" {
		_ = os.Remove(flagFile)
		return
	}

	// 只要A记录有效且IP不为空就写入
	_ = os.WriteFile(flagFile, []byte(hostname), 0644)
}
