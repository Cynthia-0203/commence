package checkout

import (
	"context"
	"fmt"

	"github.com/Cynthia/commence/app/frontend/biz/service"
	"github.com/Cynthia/commence/app/frontend/biz/utils"
	checkout "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CheckOut .
// @router /checkout [GET]
func CheckOut(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.Empty
	fmt.Println("checkout start bind")
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Println("checkout end bind")
	// resp := &checkout.Empty{}
	resp, err := service.NewCheckOutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "checkout", utils.WarpResponse(ctx, c, resp))
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, resp))
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &checkout.Empty{}
	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "result", utils.WarpResponse(ctx, c, resp))
}