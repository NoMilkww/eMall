package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/feeeeling/eMall/app/frontend/infra/rpc"
	rpcproduct "github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	category "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/category"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "category",
		"items": p.Products,
	}, nil
}
