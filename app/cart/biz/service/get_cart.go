package service

import (
	"context"
	"github.com/feeeeling/eMall/app/cart/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/cart/biz/model"
	cart "github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	items, err := model.GetCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	respItems := make([]*cart.CartItem, 0, len(items))
	for _, item := range items {
		respItems = append(respItems, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  respItems,
		},
	}, nil
}
