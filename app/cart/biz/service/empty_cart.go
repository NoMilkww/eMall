package service

import (
	"context"
	"github.com/feeeeling/eMall/app/cart/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/cart/biz/model"
	cart "github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	err = model.EmptyCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	return
}
