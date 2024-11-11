package service

import (
	"context"

	checkout "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/checkout"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	rpccheckout "github.com/Cynthia/commence/rpc_gen/kitex_gen/checkout"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	userId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CardNumber,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             req.Cvv,
		},
	})

	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
