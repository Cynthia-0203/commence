package utils

import (
	"context"

	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendutils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId
	if userId > 0 {
		cartResp,err:=rpc.CartClient.GetCart(ctx,&cart.GetCartReq{
			UserId: uint32(userId),
		})
		if err != nil && cartResp != nil {
			content["cart_num"] = cartResp.Cart
		}
	}
	return content
}