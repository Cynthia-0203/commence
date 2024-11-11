package service

import (
	"context"
	"time"

	order "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/order"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	"github.com/Cynthia/commence/app/frontend/types"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	rpcorder "github.com/Cynthia/commence/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *order.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	listOrderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range listOrderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)

		for _,v := range v.OrderItems {
			total += v.Cost
			i := v.Item
			productResp,err := rpc.ProductClient.GetProduct(h.Context,&rpcproduct.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}

			p := productResp.Product

			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture: p.Picture,
				Cost: v.Cost,
				Quantity: i.Quantity,
			})
		}
		created := time.Unix(int64(v.CreatedAt),0)
		list = append(list, types.Order{
			OrderId: v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost: total,
			Items: items,
		})
	}	
	return utils.H{
		"Title":  "Order",
		"Orders": list,
	}, nil
	
}
