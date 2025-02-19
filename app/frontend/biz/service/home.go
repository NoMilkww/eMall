package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/feeeeling/eMall/app/frontend/infra/rpc"
	rpcproduct "github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Home",
		"items": p.Products,
	}, nil
}
