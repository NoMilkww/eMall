package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/feeeeling/eMall/app/user/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/user/biz/model"
	user "github.com/feeeeling/eMall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService

type ValidateEmail struct {
	Email string `validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
}

func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Email == "" || req.Password == "" {
		a := fmt.Sprintf("%v", req)
		return nil, errors.New("email or password is empty " + a)
	}
	// check if email is valid
	validateEmail := ValidateEmail{Email: req.Email}
	if err = validator.Validate(validateEmail); err != nil {
		return nil, errors.New("email format error")
	}
	//if req.Password != req.PasswordConfirm {
	//	return nil, errors.New("password not match")
	//}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
