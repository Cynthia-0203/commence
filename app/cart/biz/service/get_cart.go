package service

import (
	"context"

	"github.com/Cynthia/commence/app/cart/biz/dal/mysql"
	"github.com/Cynthia/commence/app/cart/model"
	cart "github.com/Cynthia/commence/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	items,err := model.GetCartByUserId(s.ctx,mysql.DB,req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002,err.Error())
	}
	var list []*cart.CartItem
	for _,item := range items {
		list = append(list, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Qty,
		})
	}
	return &cart.GetCartResp{Cart: list},nil
}
