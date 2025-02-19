package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	frontendUtils "github.com/feeeeling/eMall/app/frontend/utils"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// TODO: 鉴权
		session := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, session.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			uri := []byte("/sign-in?next=" + c.FullPath())
			//klog.Infof("%s", uri)
			c.Redirect(302, uri)

			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
