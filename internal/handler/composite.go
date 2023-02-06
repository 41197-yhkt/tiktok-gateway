package handler

import (
	"context"
	"pkg/errno"
	douyin "tiktok-gateway/internal/model"
	"tiktok-gateway/internal/rpc"
	"tiktok-gateway/kitex_gen/composite"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
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
	claim := jwt.ExtractClaims(ctx, c)

	user_id, ok := claim["user_id"]

	if !ok {
		hlog.DefaultLogger().Info("user id not exist in jwt")
		SendResponse(c, *errno.UnauthorizedTokenError)
		return
	}

	uid, ok := user_id.(int64)

	if !ok {
		hlog.DefaultLogger().Info("user id assert fail")
		SendResponse(c, *errno.UnauthorizedTokenError)
		return
	}

	hlog.DefaultLogger().Info("user_id=", user_id)

	errNo := rpc.FavoriteAction(context.Background(), &composite.BasicFavoriteActionRequest{
		VedioId:    req.VedioID,
		ActionType: req.ActionType,
		UserId: uid,
	})

	if errNo != *errno.Success {
		SendResponse(c, errNo)
		return
	}

	SendResponse(c, *errno.Success)
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
