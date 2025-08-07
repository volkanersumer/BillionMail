package relay

import (
	"billionmail-core/api/relay/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/public"
	relay_service "billionmail-core/internal/service/relay"
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ListRelayConfigs(ctx context.Context, req *v1.ListRelayConfigsReq) (res *v1.ListRelayConfigsRes, err error) {
	res = &v1.ListRelayConfigsRes{}
	res.Data.List = make([]*v1.BmRelayWithSPF, 0)

	model := g.DB().Model("bm_relay_config").Safe()

	if req.Keyword != "" {
		searchKey := "%" + req.Keyword + "%"
		model = model.WhereLike("remark", searchKey).
			WhereOrLike("relay_host", searchKey).
			WhereOrLike("auth_user", searchKey).
			WhereOrLike("smtp_name", searchKey)
	}

	if req.Rtype != "" {
		model = model.Where("rtype", req.Rtype)
	}

	total, err := model.Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to retrieve relay configuration count: {}", err.Error())))
		return res, nil
	}

	var configs []*entity.BmRelayConfig
	err = model.Order("create_time DESC").Scan(&configs)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to retrieve relay configurations: {}", err.Error())))
		return res, nil
	}

	for _, config := range configs {

		relayWithSPF := &v1.BmRelayWithSPF{
			BmRelay: &v1.BmRelay{
				Id:            config.Id,
				Remark:        config.Remark,
				Rtype:         config.Rtype,
				RelayHost:     config.RelayHost,
				RelayPort:     config.RelayPort,
				AuthUser:      config.AuthUser,
				AuthPassword:  "", // 不返回密码
				Ip:            config.Ip,
				Host:          config.Host,
				Active:        config.Active,
				CreateTime:    config.CreateTime,
				UpdateTime:    config.UpdateTime,
				AuthMethod:    config.AuthMethod,
				HeloName:      config.HeloName,
				SmtpName:      config.SmtpName,
				SkipTlsVerify: config.SkipTlsVerify,
				SenderDomains: []string{},
				//RelayDomains:   []*v1.RelayDomainMapping{},
			},
			SmtpStatus: v1.SmtpStatus{
				Status: true,
				Msg:    "Connected",
			},
		}

		var mappings []*entity.BmRelayDomainMapping
		err = g.DB().Model("bm_relay_domain_mapping").
			Where("relay_id", config.Id).
			Scan(&mappings)

		if err != nil {
			g.Log().Warning(ctx, "Failed to query domain mappings for relay ID", config.Id, ":", err)
		} else {

			for _, mapping := range mappings {

				relayWithSPF.BmRelay.SenderDomains = append(relayWithSPF.BmRelay.SenderDomains, mapping.SenderDomain)
			}
		}

		var firstDomain string
		if len(relayWithSPF.BmRelay.SenderDomains) > 0 {
			firstDomain = relayWithSPF.BmRelay.SenderDomains[0]
			if strings.HasPrefix(firstDomain, "@") {
				firstDomain = firstDomain[1:]
			}

			record, _ := domains.GetSPFRecord(firstDomain, false)
			if record.Value != "" {
				relayWithSPF.SPFRecord = v1.DNSRecord{
					Type:  record.Type,
					Host:  record.Host,
					Value: record.Value,
				}
			} else {

				spfValue := relay_service.GenerateSPFRecord(config.Ip, config.Host, firstDomain)
				relayWithSPF.SPFRecord = v1.DNSRecord{
					Type:  "TXT",
					Host:  firstDomain,
					Value: spfValue,
				}
			}
		}

		res.Data.List = append(res.Data.List, relayWithSPF)
	}

	res.Data.Total = total
	res.SetSuccess("Success")
	return res, nil
}
