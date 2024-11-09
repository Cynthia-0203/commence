package service

import (
	"context"
	"fmt"

	category "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/category"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	p,err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err!=nil{
		return nil,err
	}
	fmt.Println("http-category resp:",p)
	return utils.H{
		"Title":"Category",
		"Items":p.Products,
	},nil
}
