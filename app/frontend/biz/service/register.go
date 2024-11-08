package service

import (
	"context"
	"fmt"

	auth "github.com/Cynthia/commence/app/frontend/hertz_gen/frontend/auth"
	"github.com/Cynthia/commence/app/frontend/infra/rpc"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *auth.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("req: ", req)
	userResp,err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
		ConfirmPassword: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return nil, nil
}