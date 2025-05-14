package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// BmRelay
type BmRelay struct {
	Id           int    `json:"id"`
	Remark       string `json:"remark"`        // Remark, e.g., "AWS SES Japan Region"
	SenderDomain string `json:"sender_domain"` // Sender domain, e.g., "example.com" (add "@" when writing to config)
	RelayHost    string `json:"relay_host"`    // Relay server address,
	RelayPort    string `json:"relay_port"`    // Relay server port, e.g., "587"
	AuthUser     string `json:"auth_user"`     // SMTP authentication username
	AuthPassword string `json:"auth_password"` // SMTP authentication password (consider encrypted storage)
	IP           string `json:"ip"`            // IP for reminding users to update SPF record (optional) +ip4:23.158.104.237
	Host         string `json:"host"`          // Host for reminding users to update SPF record (optional) include:lootk.cn
	Active       int    `json:"active"`        // Whether this relay configuration is active: 1-enabled, 0-disabled
	CreateTime   int    `json:"create_time"`   // Creation time
	UpdateTime   int    `json:"update_time"`   // Update time
}

type DNSRecord struct {
	Type  string `json:"type"`
	Host  string `json:"host"`
	Value string `json:"value"`
}
type BmRelayWithSPF struct {
	*BmRelay
	SPFRecord DNSRecord `json:"spf_record"`
}

// CreateRelayConfigReq Request to add a relay configuration
type CreateRelayConfigReq struct {
	g.Meta        `path:"/relay/add" method:"post" tags:"Relay Configs" summary:"Create a new relay configuration"`
	Authorization string `json:"authorization" in:"header" dc:"Authorization" v:"required"`
	Remark        string `json:"remark" v:"max-length:255" dc:"Remark, e.g., AWS SES Japan Region"`
	SenderDomain  string `json:"sender_domain" v:"required|max-length:255" dc:"Sender domain, e.g., example.com"`
	RelayHost     string `json:"relay_host" v:"required|max-length:255" dc:"Relay server address"`
	RelayPort     string `json:"relay_port" v:"required|max-length:10" dc:"Relay server port, e.g., 587"`
	AuthUser      string `json:"auth_user" v:"required|max-length:255" dc:"SMTP authentication username"`
	AuthPassword  string `json:"auth_password" v:"required|max-length:255" dc:"SMTP authentication password"`
	IP            string `json:"ip" v:"max-length:255" dc:"IP for reminding users to update SPF record, e.g., +ip4:23.158.104.237"`
	Host          string `json:"host" v:"max-length:255" dc:"Host for reminding users to update SPF record, e.g., include:lootk.cn"`
	Active        int    `json:"active" d:"1" v:"in:0,1" dc:"Whether enabled: 1-enabled, 0-disabled (default is 1)"`
}

type CreateRelayConfigRes struct {
	api_v1.StandardRes
	SPFRecord DNSRecord `json:"spf_record"` // Generated SPF record suggestion
}

// ListRelayConfigsReq Request to get a list of relay configurations
type ListRelayConfigsReq struct {
	g.Meta        `path:"/relay/list" method:"get" tags:"Relay Configs" summary:"Get the list of relay configurations"`
	Authorization string `json:"authorization" in:"header" dc:"Authorization" v:"required"`
	//Active        int    `json:"active,omitempty" v:"in:0,1" dc:"Filter by status: 1-enabled, 0-disabled"`
	Page     int    `json:"page" v:"required|min:1" dc:"Page number"`
	PageSize int    `json:"page_size" v:"required|min:1" dc:"Page size"`
	Keyword  string `json:"keyword" dc:"search"` // Search by remark, relay server address, or sender domain
}

type ListRelayConfigsRes struct {
	api_v1.StandardRes
	Data struct {
		Total int               `json:"total" dc:"Total records count"`
		List  []*BmRelayWithSPF `json:"list" dc:"Relay configuration list"`
	} `json:"data"`
}

// UpdateRelayConfigReq Request to update a relay configuration
type UpdateRelayConfigReq struct {
	g.Meta        `path:"/relay/edit" method:"post" tags:"Relay Configs" summary:"Update a relay configuration"`
	Authorization string `json:"authorization" in:"header" dc:"Authorization" v:"required"`
	ID            int    `json:"id" v:"required|min:1" dc:"The ID of the configuration to update"`
	Remark        string `json:"remark,omitempty" v:"max-length:255" dc:"Remark, e.g., AWS SES Japan Region"`
	SenderDomain  string `json:"sender_domain,omitempty" v:"max-length:255" dc:"Sender domain, e.g., example.com"`
	RelayHost     string `json:"relay_host,omitempty" v:"max-length:255" dc:"Relay server address"`
	RelayPort     string `json:"relay_port,omitempty" v:"max-length:10" dc:"Relay server port, e.g., 587"`
	AuthUser      string `json:"auth_user,omitempty" v:"max-length:255" dc:"SMTP authentication username"`
	AuthPassword  string `json:"auth_password,omitempty" v:"max-length:255" dc:"SMTP authentication password (no update if not provided)"`
	IP            string `json:"ip,omitempty" v:"max-length:255" dc:"IP for reminding users to update SPF record, e.g., +ip4:23.158.104.237"`
	Host          string `json:"host,omitempty" v:"max-length:255" dc:"Host for reminding users to update SPF record, e.g., include:lootk.cn"`
	Active        int    `json:"active,omitempty" v:"in:0,1" dc:"Whether enabled: 1-enabled, 0-disabled"` // Can be set to 0 as a toggle
}

type UpdateRelayConfigRes struct {
	api_v1.StandardRes
	SPFRecord DNSRecord `json:"spf_record"`
}

// DeleteRelayConfigReq Request to delete a relay configuration
type DeleteRelayConfigReq struct {
	g.Meta        `path:"/relay/delete" method:"post" tags:"Relay Configs" summary:"Delete a relay configuration"`
	Authorization string `json:"authorization" in:"header" dc:"Authorization" v:"required"`
	ID            int64  `json:"id" v:"required|min:1" dc:"The ID of the configuration to delete"`
}

type DeleteRelayConfigRes struct {
	api_v1.StandardRes
}
