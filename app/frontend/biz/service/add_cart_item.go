package service

import (
	"context"

	fcart "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/cart"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *fcart.AddCartItemReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_,err = rpc.CartClient.AddItem(h.Context, &cart.AddItemReq{
		UserId: uint32(frontendutils.GetUserIdFromCtx(h.Context)),
		Item: &cart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.Count),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
