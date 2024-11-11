package service

import (
	"context"

	"github.com/Cynthia/commence/app/order/biz/dal/mysql"
	"github.com/Cynthia/commence/app/order/model"
	order "github.com/Cynthia/commence/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		return nil, kerrors.NewBizStatusError(500001,`order items is empty`)
	}
	
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId,err := uuid.NewUUID()

		o := model.Order{
			OrderId: orderId.String(),
			UserId:  req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
			o.Consignee.ZipCode = a.ZipCode
		}

		if err = tx.Create(&o).Error;err != nil{
			return err
		}
		var items []model.OrderItem
		for _,orderItem := range req.OrderItems {
			oi := model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId: orderItem.Item.ProductId,
				Quantity: orderItem.Item.Quantity,
				Cost: orderItem.Cost,
			}
			items = append(items, oi)
		}	

		if err := tx.Create(&items).Error;err != nil{
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})
	return
}
