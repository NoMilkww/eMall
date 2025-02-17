package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	common "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/common"
	"github.com/feeeeling/eMall/app/frontend/infra/rpc"
	frontendUtils "github.com/feeeeling/eMall/app/frontend/utils"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"
	"strconv"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	klog.Infof("cartResp = %+v, err = %+v", cartResp, err)
	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float64
	if cartResp.Cart == nil {
		return utils.H{
			"title": "Cart",
			"items": items,
			"total": total,
		}, nil
	}
	for _, item := range cartResp.Cart.Items {
		klog.Infof("item = %+v", item)
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(p.Price) * float64(item.Quantity)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
