package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/common"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	return
}
