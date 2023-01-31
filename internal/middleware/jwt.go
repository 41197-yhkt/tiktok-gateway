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
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte("tiktok"), //密钥
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			var req user.DouyinUserResponse
			err = c.BindAndValidate(&req)
			c.JSON(http.StatusOK, utils.H{
				"status_code": code,
				"token":       token,
				"user_id ":    req.GetUser().Id,
				"status_msg":  "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			return handler.DouyinUserLoginMethod(ctx, c)
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
