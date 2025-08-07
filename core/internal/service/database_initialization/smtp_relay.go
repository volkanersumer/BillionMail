package database_initialization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"time"
)

func init() {
	registerHandler(func() {
		sqlList := []string{

			`-- Relay configuration 
			CREATE TABLE IF NOT EXISTS bm_relay_config (
			id BIGSERIAL PRIMARY KEY,
			remark varchar(255),  
			rtype varchar(30),  -- Relay type: gmail, sendgrid, custom, aws, mailgun, local
			relay_host varchar(255) NOT NULL, -- Relay server address 
			relay_port varchar(10) NOT NULL, -- Relay server port 
			auth_user varchar(255) NOT NULL, -- SMTP auth username
			auth_password varchar(255) NOT NULL, -- SMTP auth password (encrypted)
			ip varchar(255), -- IP for reminding user to update SPF record (optional) 
			host varchar(255), --  Hostname for reminding user to update SPF record (optional)
			helo_name varchar(255), -- HELO hostname 
			skip_tls_verify smallint DEFAULT 0, -- kip TLS verification: 1-skip, 0-do not skip 
			smtp_name varchar(50), --  SMTP server name
			active smallint NOT NULL DEFAULT 1, 
			auth_method varchar(20), -- Authentication method: LOGIN, PLAIN, CRAM-MD5, NONE
			create_time int NOT NULL default 0,
			update_time int NOT NULL default 0
			) `,

			`-- Relay to domain mapping
			CREATE TABLE IF NOT EXISTS bm_relay_domain_mapping (
			id BIGSERIAL PRIMARY KEY,
			relay_id BIGINT NOT NULL, --  Reference to id in bm_relay_config  
			sender_domain varchar(255) NOT NULL, -- Sender domain, e.g. "example.com"
			create_time int NOT NULL default 0
			) `,

			`-- Domain to SMTP service name mapping table, used for Postfix sender_dependent_default_transport_maps
			CREATE TABLE IF NOT EXISTS bm_relay_domain_transport (
			id BIGSERIAL PRIMARY KEY,
			domain varchar(255) NOT NULL,  -- Sender domain, e.g. "@example.com"
			relay_type varchar(255) NOT NULL -- SMTP service name, e.g. "smtp_Relay_mail_xxxx_cn_gwt3_6_pvz_6"
			) `,

			`CREATE INDEX IF NOT EXISTS idx_relay_domain_transport_domain ON bm_relay_domain_transport(domain);`,
			`CREATE INDEX IF NOT EXISTS idx_relay_domain_transport_domain_relay ON bm_relay_domain_transport(domain,relay_type);`,
			`CREATE INDEX IF NOT EXISTS idx_relay_domain_mapping_domain ON bm_relay_domain_mapping(sender_domain);`,
			`CREATE INDEX IF NOT EXISTS idx_relay_domain_mapping_relay_id ON bm_relay_domain_mapping(relay_id);`,
		}

		for _, sql := range sqlList {
			_, err := g.DB().Exec(context.Background(), sql)
			if err != nil {
				g.Log().Error(context.Background(), "Failed to create mail server tables:", err)
				return
			}
		}

		migrateRelayData()

	})
}

func migrateRelayData() {
	ctx := context.Background()

	existsOld := checkTableExists(ctx, "bm_relay")
	if !existsOld {
		g.Log().Info(ctx, "Old table bm_relay does not exist, skip migration")
		return
	}

	existsNew := checkTableExists(ctx, "bm_relay_config")
	if !existsNew {
		g.Log().Error(ctx, "New table bm_relay_config does not exist, please create the table first")
		return
	}

	newCount, err := g.DB().Model("bm_relay_config").Count()
	if err != nil {
		g.Log().Error(ctx, "Failed to query record count of new table:", err)
		return
	}
	if newCount > 0 {
		g.Log().Warning(ctx, "bm_relay_config already has data (%d records), skip migration to prevent duplication", newCount)
		return
	}

	// 查询旧表数据（显式指定字段，避免多余字段干扰）
	var oldRelays []struct {
		Id            int64   `json:"id"`
		Remark        *string `json:"remark"`
		Rtype         *string `json:"rtype"`
		SenderDomain  *string `json:"sender_domain"`
		RelayHost     *string `json:"relay_host"`
		RelayPort     *string `json:"relay_port"`
		AuthUser      *string `json:"auth_user"`
		AuthPassword  *string `json:"auth_password"`
		Ip            *string `json:"ip"`
		Host          *string `json:"host"`
		HeloName      *string `json:"helo_name"`
		SkipTlsVerify *int    `json:"skip_tls_verify"`
		SmtpName      *string `json:"smtp_name"`
		Active        *int    `json:"active"`
		AuthMethod    *string `json:"auth_method"`
		CreateTime    int     `json:"create_time"`
		UpdateTime    int     `json:"update_time"`
	}
	err = g.DB().Model("bm_relay").
		Fields("id, remark, rtype, sender_domain, relay_host, relay_port, auth_user, auth_password, ip, host, helo_name, skip_tls_verify, smtp_name, active, auth_method, create_time, update_time").
		Scan(&oldRelays)
	if err != nil {
		g.Log().Error(ctx, "Failed to query old table  bm_relay:", err)
		return
	}

	if len(oldRelays) == 0 {
		g.Log().Info(ctx, "No data in old table bm_relay, no migration needed")
		return
	}

	g.Log().Infof(ctx, "Start migrating %d relay configurations", len(oldRelays))

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log().Error(ctx, "Failed to begin transaction:", err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			g.Log().Error(ctx, "Transaction rolled back")
		}
	}()

	for _, r := range oldRelays {
		// Build configData, skip nil fields
		configData := g.Map{}

		if r.Remark != nil {
			configData["remark"] = *r.Remark
		}
		if r.Rtype != nil {
			configData["rtype"] = *r.Rtype
		}
		if r.RelayHost == nil || *r.RelayHost == "" {
			g.Log().Errorf(ctx, "Skip record ID=%d: relay_host is empty", r.Id)
			continue
		}
		configData["relay_host"] = *r.RelayHost

		if r.RelayPort == nil || *r.RelayPort == "" {
			configData["relay_port"] = "25" // Default port
		} else {
			configData["relay_port"] = *r.RelayPort
		}

		if r.AuthUser != nil {
			configData["auth_user"] = *r.AuthUser
		}
		if r.AuthPassword != nil {
			configData["auth_password"] = *r.AuthPassword
		}
		if r.Ip != nil {
			configData["ip"] = *r.Ip
		}
		if r.Host != nil {
			configData["host"] = *r.Host
		}
		if r.HeloName != nil {
			configData["helo_name"] = *r.HeloName
		}
		if r.SkipTlsVerify != nil {
			configData["skip_tls_verify"] = *r.SkipTlsVerify
		} else {
			configData["skip_tls_verify"] = 0
		}
		if r.SmtpName != nil {
			configData["smtp_name"] = *r.SmtpName
		} else {
			// Generate unique name
			configData["smtp_name"] = "MigratedRelay_" + time.Now().Format("20060102_150405") + "_" + strings.ReplaceAll(*r.RelayHost, ".", "_")
		}
		if r.Active != nil {
			configData["active"] = *r.Active
		} else {
			configData["active"] = 1
		}
		if r.AuthMethod != nil {
			configData["auth_method"] = *r.AuthMethod
		}

		configData["create_time"] = r.CreateTime
		configData["update_time"] = r.UpdateTime

		// Insert into bm_relay_config
		result, err := tx.Model("bm_relay_config").Data(configData).Insert()
		if err != nil {
			g.Log().Errorf(ctx, "Failed to insert into bm_relay_config (ID=%d): %v", r.Id, err)
			err = err
			return
		}

		relayConfigId, err := result.LastInsertId()
		if err != nil {
			g.Log().Errorf(ctx, "Failed to get inserted ID (ID=%d): %v", r.Id, err)
			err = err
			return
		}

		// Insert domain mapping
		if r.SenderDomain != nil && *r.SenderDomain != "" {
			mappingData := g.Map{
				"relay_id":      relayConfigId,
				"sender_domain": *r.SenderDomain,
				"create_time":   r.CreateTime,
			}
			_, err = tx.Model("bm_relay_domain_mapping").Data(mappingData).Insert()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to insert into bm_relay_domain_mapping (ID=%d): %v", r.Id, err)
				err = err
				return
			}
		}
	}

	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, "Failed to commit transaction:", err)
		return
	}

	g.Log().Infof(ctx, "Data migration succeeded! Migrated %d relay configurations", len(oldRelays))
}

func checkTableExists(ctx context.Context, tableName string) bool {
	query := `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name = ?
	)`

	result, err := g.DB().GetOne(ctx, query, tableName)
	if err != nil {
		g.Log().Errorf(ctx, "Error checking if table %s exists: %v", tableName, err)
		return false
	}

	return result["exists"].Bool()
}
