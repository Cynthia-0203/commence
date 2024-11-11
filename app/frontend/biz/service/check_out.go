package service

import (
	"context"
	"strconv"

	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	checkout "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/checkout"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/cart"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckOutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutService(Context context.Context, RequestContext *app.RequestContext) *CheckOutService {
	return &CheckOutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	carts,err := rpc.CartClient.GetCart(h.Context,&cart.GetCartReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return nil,err
	}
	var total float32

	for  _,cartItem := range carts.Cart {
		productResp,err := rpc.ProductClient.GetProduct(h.Context,&product.GetProductReq{Id: cartItem.ProductId})
		if err != nil {
			return nil,err
		}
		if productResp.Product == nil {
			continue
		}

		p := productResp.Product

		items = append(items, map[string]string{
			"Name": p.Name,
			"Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Qty": strconv.Itoa(int(cartItem.Quantity)),
			"Picture": p.Picture,
		})

		cost := p.Price * float32(cartItem.Quantity)
		total += cost
	}
	
	return utils.H{
		"Title":"CheckOut",
		"Items": items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	},nil
}
