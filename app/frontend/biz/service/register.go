package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/auth"
	"github.com/feeeeling/eMall/app/frontend/infra/rpc"
	user "github.com/feeeeling/eMall/rpc_gen/kitex_gen/user"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	hlog.CtxInfof(h.Context, "req = %+v", req)
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()
	return
}
