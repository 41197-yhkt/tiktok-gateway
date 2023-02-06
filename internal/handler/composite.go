package handler

import (
	"context"
	"pkg/errno"
	douyin "tiktok-gateway/internal/model"
	"tiktok-gateway/internal/rpc"
	"tiktok-gateway/kitex_gen/composite"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinFeedMethod .
// @router /douyin/feed [GET]
// TODO: fix feed idl
func DouyinFeedMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinFavoriteActionMethod .
// @router /douyin/favorite/action [POST]
func DouyinFavoriteActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// TODO: Tocken中拿到用户名
	errNo := rpc.FavoriteAction(context.Background(), &composite.BasicFavoriteActionRequest{
		VedioId:    req.VedioID,
		ActionType: req.ActionType,
	})

	if errNo != *errno.Success {
		SendResponse(c, errNo)
	}
	resp := new(douyin.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinFavoriteListMethod .
// @router /douyin/favorite/list [GET]
func DouyinFavoriteListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinFavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinCommentActionMethod .
// @router /douyin/comment/action [POST]
func DouyinCommentActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinCommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinCommentListMethod .
// @router /douyin/comment/list [GET]
func DouyinCommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}
