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
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte("tiktok"), //密钥
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
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
