package service

import (
	"context"
	"fmt"

	home "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/home"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any,error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	p,err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	fmt.Println("http-home resp:",p)
	if err != nil{
		return nil,err
	}

	return map[string]any{
		"Title":"Hot Sale",
		"Items":p.Products,
	},nil
}
