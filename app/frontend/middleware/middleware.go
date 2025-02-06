package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func Register(h *server.Hertz) {
	// 注册中间件
	h.Use(GlobalAuth())
}
