package service

import (
	"context"
	"github.com/feeeeling/eMall/app/product/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/product/biz/model"
	product "github.com/feeeeling/eMall/app/product/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	if err != nil {
		return nil, err
	}
	var result []*product.Product
	for _, p := range products {
		result = append(result, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		})
	}
	return &product.SearchProductsResp{
		Results: result,
	}, nil
}
