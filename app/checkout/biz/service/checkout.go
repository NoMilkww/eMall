package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/feeeeling/eMall/app/checkout/infra/rpc"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart"
	checkout "github.com/feeeeling/eMall/rpc_gen/kitex_gen/checkout"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/payment"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/product"
	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(4004001, err.Error())
	}
	if cartResp == nil || cartResp.Cart == nil || cartResp.Cart.Items == nil {
		return nil, kerrors.NewBizStatusError(4004002, "cart is empty")
	}
	var total float32
	for _, item := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, kerrors.NewBizStatusError(4004003, err.Error())
		}
		if productResp == nil || productResp.Product == nil {
			continue
		}
		total += productResp.Product.Price * float32(item.Quantity)
	}
	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()

	payReq := &payment.ChargeReq{
		Amount: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: orderId,
		UserId:  req.UserId,
	}
	payResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005002, err.Error())
	}
	klog.Info(payResp)

	return &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: payResp.TransactionId,
	}, nil
}
