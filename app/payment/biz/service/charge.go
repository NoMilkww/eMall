package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/feeeeling/eMall/app/payment/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/payment/biz/model"
	payment "github.com/feeeeling/eMall/app/payment/kitex_gen/payment"
	"github.com/google/uuid"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    string(req.CreditCard.CreditCardCvv),
		Month:  string(req.CreditCard.CreditCardExpirationMonth),
		Year:   string(req.CreditCard.CreditCardExpirationYear),
	}
	err = card.Validate()
	if err != nil {
		return nil, kerrors.NewBizStatusError(4004001, err.Error())
	}
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005001, err.Error())
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
		OrderId:       req.OrderId,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(4006001, err.Error())
	}
	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}, nil
}
