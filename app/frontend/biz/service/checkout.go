package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/feeeeling/eMall/app/frontend/infra/rpc"
	frontendutils "github.com/feeeeling/eMall/app/frontend/utils"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/feeeeling/eMall/app/frontend/hertz_gen/frontend/common"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return nil, err
	}
	if cartResp == nil || cartResp.Cart == nil || cartResp.Cart.Items == nil {
		return nil, err
	}
	var total float32
	for _, item := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp == nil || productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', -1, 64),
			"Picture": p.Picture,
			"Qty":     strconv.FormatUint(uint64(item.Quantity), 10),
		})
		total += p.Price * float32(item.Quantity)
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": total,
	}, nil
}
