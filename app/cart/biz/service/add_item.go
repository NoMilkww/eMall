package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/feeeeling/eMall/app/cart/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/cart/biz/model"
	"github.com/feeeeling/eMall/app/cart/rpc"
	cart "github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	p, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if p == nil || p.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}
	err = model.AddItem(s.ctx, mysql.DB, &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
