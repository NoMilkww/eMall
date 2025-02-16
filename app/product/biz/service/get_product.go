package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/feeeeling/eMall/app/product/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/product/biz/model"
	product "github.com/feeeeling/eMall/app/product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(2004001, "product id is empty")
	}
	prodcutQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := prodcutQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		},
	}, nil
}
