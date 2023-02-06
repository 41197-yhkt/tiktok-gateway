// Code generated by hertz generator.

package routers

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"tiktok-gateway/internal/handler"
	"tiktok-gateway/internal/middleware"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	group := r.Group("/douyin")
	group.POST("/user/register", handler.DouyinUserRegisterMethod, middleware.JwtMiddleware.LoginHandler)
	group.POST("/user/login", middleware.JwtMiddleware.LoginHandler)
	auth := group.Group("/user", middleware.JwtMiddleware.MiddlewareFunc())
	auth.GET("/", handler.DouyinUserMethod)
	auth = group.Group("/relation", middleware.JwtMiddleware.MiddlewareFunc())
	auth.POST("/action", handler.DouyinUserMethod)
	auth.GET("/follow/list", handler.DouyinUserMethod)
	auth.GET("/follower/list", handler.DouyinUserMethod)
	auth.GET("/friend/list", handler.DouyinUserMethod)

	// feed不需jwt
	group.GET("/feed", handler.DouyinFeedMethod)

	// TODO:为了测试方便，还没上jwt
	// auth2 := group.Group("/publish",middleware.JwtMiddleware.MiddlewareFunc())
	auth3 := group.Group("/favourite",middleware.JwtMiddleware.MiddlewareFunc())
	// auth4 := group.Group("/comment",middleware.JwtMiddleware.MiddlewareFunc())
	
	//group.POST("/publish/action", handler.DouyinPublishActionMethod)
	//group.GET("/publish/list", handler.DouyinPublishListMethod)
	auth3.POST("/action", handler.DouyinFavoriteActionMethodTest)
	auth3.GET("/list", handler.DouyinFavoriteListMethod)
	group.POST("/comment/action", handler.DouyinCommentActionMethod)
	group.GET("/comment/list", handler.DouyinCommentListMethod)
	
}
