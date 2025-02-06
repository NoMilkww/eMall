package service

import (
	"context"

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
	var resp = make(map[string]any)
	var items []map[string]any
	resp["title"] = "Hot Sales"
	resp["items"] = items
	return resp, nil
}
