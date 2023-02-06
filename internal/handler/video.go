package handler

// import (
// 	"context"
// 	"tiktok-gateway/internal/model"
// 	"github.com/cloudwego/hertz/pkg/app"
// 	"github.com/cloudwego/hertz/pkg/protocol/consts"
// )

// // DouyinPublishActionMethod .
// // @router /douyin/publish/action [POST]
// func DouyinPublishActionMethod(ctx context.Context, c *app.RequestContext) {
// 	var err error
// 	var req douyin.DouyinPublishActionRequest
// 	err = c.BindAndValidate(&req)
// 	if err != nil {
// 		c.String(consts.StatusBadRequest, err.Error())
// 		return
// 	}

// 	resp := new(douyin.DouyinPublishActionResponse)

// 	c.JSON(consts.StatusOK, resp)
// }

// // DouyinPublishListMethod .
// // @router /douyin/publish/list [GET]
// func DouyinPublishListMethod(ctx context.Context, c *app.RequestContext) {
// 	var err error
// 	var req douyin.DouyinPublishListRequest
// 	err = c.BindAndValidate(&req)
// 	if err != nil {
// 		c.String(consts.StatusBadRequest, err.Error())
// 		return
// 	}

// 	resp := new(douyin.DouyinPublishListResponse)

// 	c.JSON(consts.StatusOK, resp)
// }