/*
* @Author: 滚~韬
* @Date:   2023/1/28 15:00
 */
 package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"tiktok-gateway/internal/handler"
	"tiktok-gateway/kitex_gen/user"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte("tiktok"), //密钥
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.DouyinUserLoginResponse); ok {
				return jwt.MapClaims{
					IdentityKey: v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			id, _ := c.Get("user_id")
			c.JSON(http.StatusOK, utils.H{
				"status_code": code,
				"token":       token,
				"user_id ":    id,
				"status_msg":  "success",
			})
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &user.User{
				Id: claims[IdentityKey].(int64),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			method, _ := handler.DouyinUserLoginMethod(ctx, c)
			c.Set("user_id", method.(*user.DouyinUserLoginResponse).UserId)
			return method, nil
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err information => %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, statusCode int, statusMsg string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": statusCode,
				"status_msg":  statusMsg,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
