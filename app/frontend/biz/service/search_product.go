package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/feeeeling/eMall/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/product"
	rpcproduct "github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.Q})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"items": p.Results,
		"q":     req.Q,
	}, nil
}
