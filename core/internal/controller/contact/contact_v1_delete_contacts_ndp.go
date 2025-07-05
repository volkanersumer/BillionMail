package contact

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/contact/v1"
)

func (c *ControllerV1) DeleteContactsNDP(ctx context.Context, req *v1.DeleteContactsNDPReq) (res *v1.DeleteContactsNDPRes, err error) {
	res = &v1.DeleteContactsNDPRes{}

	if len(req.Ids) == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "The Ids field is required")))
		return
	}

	_, err = g.DB().Model("bm_contacts").WhereIn("id", req.Ids).Delete()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to delete contacts: {}", err.Error())))
		return
	}

	res.SetSuccess(public.LangCtx(ctx, "Contacts deleted successfully"))
	return
}
