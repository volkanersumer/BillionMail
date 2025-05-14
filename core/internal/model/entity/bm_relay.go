package entity

// BmBcc
type BmRelay struct {
	Id           int    `json:"id"`
	Remark       string `json:"remark"`        // Remark, e.g., "AWS SES Japan Region"
	SenderDomain string `json:"sender_domain"` // Sender domain, e.g., "example.com" (add "@" when writing to config)
	RelayHost    string `json:"relay_host"`    // Relay server address
	RelayPort    string `json:"relay_port"`    // Relay server port, e.g., "587"
	AuthUser     string `json:"auth_user"`     // SMTP authentication username
	AuthPassword string `json:"auth_password"` // SMTP authentication password (consider encrypted storage)
	IP           string `json:"ip"`            // IP for reminding users to update SPF record (optional) +ip4:23.158.104.237
	Host         string `json:"host"`          // Host for reminding users to update SPF record (optional) include:lootk.cn
	Active       int    `json:"active"`        // Whether this relay configuration is active: 1-enabled, 0-disabled
	CreateTime   int    `json:"create_time"`   // Creation time
	UpdateTime   int    `json:"update_time"`   // Update time
}
