package service

import (
	"context"
	"github.com/feeeeling/eMall/app/user/biz/dal/mysql"
	user "github.com/feeeeling/eMall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{Email: "123456@qq.com", Password: "123456", PasswordConfirm: "123456"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
