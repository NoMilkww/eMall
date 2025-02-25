package service

import (
	"context"
	//order "github.com/feeeeling/eMall/app/order/kitex_gen/order"
	//MarkOrderPaid    order "github.com/feeeeling/eMall/rpc_gen/kitex_gen/order"  //replace
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
//MarkOrderPaid    
// func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
// 	// Finish your business logic.

// 	return
// }
