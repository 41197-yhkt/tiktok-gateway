package handler

import (
	"context"
	"tiktok-gateway/internal/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinFeedMethod .
// @router /relation/feed [GET]
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
// @router /relation/favorite/action [POST]
func DouyinFavoriteActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinFavoriteListMethod .
// @router /relation/favorite/list [GET]
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
// @router /relation/comment/action [POST]
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
// @router /relation/comment/list [GET]
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
