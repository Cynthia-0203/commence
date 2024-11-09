package service

import (
	"context"
	"fmt"

	product "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/product"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	rpcproduct "github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	p,err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	fmt.Println("http-getproduct resp:",p)
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Item":p.Product,
	},nil
}
