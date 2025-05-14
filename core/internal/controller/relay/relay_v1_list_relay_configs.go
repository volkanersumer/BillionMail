package relay

import (
	"billionmail-core/api/relay/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ListRelayConfigs(ctx context.Context, req *v1.ListRelayConfigsReq) (res *v1.ListRelayConfigsRes, err error) {
	res = &v1.ListRelayConfigsRes{}

	model := g.DB().Model("bm_relay").Safe()

	//if req.Active > 0 {
	//	model = model.Where("active", req.Active)
	//}
	// Search functionality (supports searching by remarks, relay server address, sender domain)
	if req.Keyword != "" {
		searchKey := "%" + req.Keyword + "%"
		model = model.WhereOr(
			"remark LIKE ?", searchKey,
		).WhereOr(
			"relay_host LIKE ?", searchKey,
		).WhereOr(
			"sender_domain LIKE ?", searchKey,
		)
	}

	count, err := model.Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to retrieve relay configuration count: {}", err.Error())))
		return res, nil
	}

	var list []*entity.BmRelay
	err = model.Page(req.Page, req.PageSize).Order("create_time DESC").Scan(&list)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to retrieve relay configuration list: {}", err.Error())))
		return res, nil
	}

	resultList := make([]*v1.BmRelayWithSPF, 0, len(list))
	for _, item := range list {
		// Generate SPF record suggestion
		spfRecord := GenerateSPFRecord(item.IP, item.Host, item.SenderDomain)

		bmRelay := toPostfixRelayResponse(ctx, item)

		relayWithSPF := &v1.BmRelayWithSPF{
			BmRelay: bmRelay,
			SPFRecord: v1.DNSRecord{
				Type:  "TXT",
				Host:  "@",
				Value: spfRecord,
			},
		}

		resultList = append(resultList, relayWithSPF)
	}

	res.Data.Total = count
	res.Data.List = resultList
	res.SetSuccess(public.LangCtx(ctx, "Relay configuration list retrieved successfully"))
	return res, nil
}

// toPostfixRelayResponse Converts database entity to API response structure
func toPostfixRelayResponse(ctx context.Context, relay *entity.BmRelay) *v1.BmRelay {
	password, _ := DecryptPassword(ctx, relay.AuthPassword)
	if password == "" {
		password = "********"
	}
	return &v1.BmRelay{
		Id:           relay.Id,
		Remark:       relay.Remark,
		SenderDomain: relay.SenderDomain,
		RelayHost:    relay.RelayHost,
		RelayPort:    relay.RelayPort,
		AuthUser:     relay.AuthUser,
		AuthPassword: password,
		IP:           relay.IP,
		Host:         relay.Host,
		Active:       relay.Active,
		CreateTime:   relay.CreateTime,
		UpdateTime:   relay.UpdateTime,
	}
}
