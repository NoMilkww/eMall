package service

import (
	"context"
	user "github.com/feeeeling/eMall/app/user/kitex_gen/user"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	sessions.Default()
	return
}
