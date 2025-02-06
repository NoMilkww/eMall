package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const (
	SessionUserId SessionUserIdKey = "user_id"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// TODO: 鉴权
		session := sessions.Default(c)
		context.WithValue(ctx, SessionUserId, session.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in？next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
