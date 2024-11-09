package service

import (
	"context"

	"github.com/Cynthia/commence/app/product/biz/dal/mysql"
	"github.com/Cynthia/commence/app/product/model"
	product "github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx,mysql.DB)
	products,err := productQuery.SearchPeoduct(req.Query)
	if err!= nil{
		return nil,err
	}
	var results []*product.Product

	for _,v := range products{
		results = append(results,&product.Product{
			Id: uint32(v.ID),
			Picture: v.Picture,
			Price: v.Price,
			Description: v.Description,
			Name: v.Name,
		})
	}
	return &product.SearchProductsResp{Results: results},nil
}
