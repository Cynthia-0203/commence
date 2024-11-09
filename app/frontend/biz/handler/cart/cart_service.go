package cart

import (
	"context"
	"fmt"

	"github.com/Cynthia/commence/app/frontend/biz/service"
	"github.com/Cynthia/commence/app/frontend/biz/utils"
	cart "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.Empty
	err = c.BindAndValidate(&req)
	
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}

// AddCartItem .
// @router /cart[POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)
	fmt.Println("http-cart req:", req.ProductId)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewAddCartItemService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Println("http-cart resp:", resp)
	c.Redirect(consts.StatusFound, []byte("/cart"))
}
