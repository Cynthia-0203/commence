package service

import (
	"context"

	"github.com/Cynthia/commence/app/checkout/infra/rpc"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/cart"
	checkout "github.com/Cynthia/commence/rpc_gen/kitex_gen/checkout"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/order"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/payment"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult,err := rpc.CartClient.GetCart(s.ctx,&cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil,kerrors.NewGRPCBizStatusError(5005001,err.Error())
	}
	
	if cartResult == nil || cartResult.Cart == nil {
		return nil,kerrors.NewGRPCBizStatusError(5004001,"cart is empty")
	}

	var total float32
	var oi []*order.OrderItem
	for  _,cartItem := range cartResult.Cart {
		productResp,resultErr := rpc.ProductClient.GetProduct(s.ctx,&product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			return nil,resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(cartItem.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Cost: cost,
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity: cartItem.Quantity,
			},
		})
	}
	var orderId string

	orderResp,err := rpc.OrderClient.PlaceOrder(s.ctx,&order.PlaceOrderReq{
		UserId: req.UserId,
		Email: req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City: req.Address.City,
			State: req.Address.State,
			ZipCode: req.Address.ZipCode,
			Country: req.Address.Country,
		},
		OrderItems: oi,
	})
	if err != nil{
		return nil,kerrors.NewGRPCBizStatusError(5004002,err.Error())
	}
	if orderResp.Order != nil || orderResp.Order != nil{
		orderId = orderResp.Order.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId: req.UserId,
		OrderId: orderId,
		Amount: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber: req.CreditCard.CreditCardNumber,
			CreditCardCvv: req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear: req.CreditCard.CreditCardExpirationYear,
		},
	}

	_,err = rpc.CartClient.EmptyCart(s.ctx,&cart.EmptyCartReq{UserId: req.UserId})

	if err != nil {
		klog.Error(err)
	}

	paymentResult ,err := rpc.PaymentClient.Charge(s.ctx,payReq)
	if err != nil {
		return nil,err
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
