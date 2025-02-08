package service

import (
	"context"
	"errors"
	"github.com/feeeeling/eMall/app/user/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/user/biz/model"
	user "github.com/feeeeling/eMall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
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
	// check if email is valid
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	validateEmail := ValidateEmail{Email: req.Email}
	if err = validator.Validate(validateEmail); err != nil {
		return nil, errors.New("email format error")
	}

	u, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, errors.New("email not found")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.Password)); err != nil {
		return nil, errors.New("password error")
	}

	return &user.LoginResp{UserId: int32(u.ID)}, nil
}
