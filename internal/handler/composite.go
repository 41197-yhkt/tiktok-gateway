package handler

import (
	"context"
	"net/http"
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
	hlog.Info("in feed")
	var err error
	var req douyin.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid, _ := getUserIdFromJWT(ctx, c)

	errNo, videosRPC, nextTime := rpc.FeedMethod(ctx, &composite.BasicFeedRequest{
		UserId:      uid,
		LastestTime: req.LastestTime,
	})

	if errNo != *errno.Success {
		//TODO: 修复sendResponse
		return
	}

	// 烦人的类型转换
	var videosHTTP []*douyin.Vedio

	for _, v := range videosRPC {
		videoHttp := douyin.Vedio{
			ID: v.Id,
			Author: &douyin.User{
				FollowerCount: v.Author.FollowerCount,
				Name:          v.Author.Name,
				ID:            v.Author.Id,
				FollowCount:   v.Author.FollowCount,
				IsFollow:      v.Author.IsFollow,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
			IsFavorite:    v.IsFavorite,
		}
		videosHTTP = append(videosHTTP, &videoHttp)
	}

	resp := douyin.DouyinFeedResponse{
		VedioList: videosHTTP,
		NextTime:  nextTime,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
		},
	}

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
		UserId:     uid,
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

	errNo, videosRPC := rpc.FavoriteList(ctx, &composite.BasicFavoriteListRequest{
		UserId:  req.UserID,
		QueryId: req.UserID,
	})

	if errNo != *errno.Success {
		SendResponse(c, errNo)
		return
	}

	// 烦人的类型转换
	var videosHTTP []*douyin.Vedio

	for _, v := range videosRPC {
		videoHttp := douyin.Vedio{
			ID: v.Id,
			Author: &douyin.User{
				FollowerCount: v.Author.FollowerCount,
				Name:          v.Author.Name,
				ID:            v.Author.Id,
				FollowCount:   v.Author.FollowCount,
				IsFollow:      v.Author.IsFollow,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
			IsFavorite:    v.IsFavorite,
		}
		videosHTTP = append(videosHTTP, &videoHttp)
	}

	msg := "get success"
	resp := douyin.DouyinFavoriteListResponse{
		VedioList: videosHTTP,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
			StatusMsg:  &msg,
		},
	}

	c.JSON(http.StatusOK, resp)
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
	uid, errNo := getUserIdFromJWT(ctx, c)

	if err != *errno.Success {
		SendResponse(c, errNo)
		return
	}

	errNo = rpc.CommentAction(context.Background(), &composite.BasicCommentActionRequest{
		VedioId:     req.VedioID,
		UserId:      uid,
		ActionType:  req.ActionType,
		CommentId:   req.CommentID,
		CommentText: req.CommentText,
	})

	if err != *errno.Success {
		SendResponse(c, errNo)
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

	err, commentsRPC := rpc.CommentList(ctx, &composite.BasicCommentListRequest{
		VedioId: req.VedioID,
	})

	var commentsHTTP []*douyin.Comment
	// 类型转换
	for _, c := range commentsRPC {
		commentHTTP := douyin.Comment{
			ID: c.Id,
			User: &douyin.User{
				FollowerCount: c.User.FollowerCount,
				Name:          c.User.Name,
				ID:            c.User.Id,
				FollowCount:   c.User.FollowCount,
				IsFollow:      c.User.IsFollow,
			},
			Content:    c.Content,
			CreateDate: c.CreateDate,
		}
		commentsHTTP = append(commentsHTTP, &commentHTTP)
	}

	resp := douyin.DouyinCommentListResponse{
		CommentList: commentsHTTP,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
		},
	}

	c.JSON(consts.StatusOK, resp)
}

func getUserIdFromJWT(ctx context.Context, c *app.RequestContext) (int64, errno.ErrNo) {
	claim := jwt.ExtractClaims(ctx, c)

	user_id, ok := claim["user_id"]

	if !ok {
		hlog.DefaultLogger().Info("user id not exist in jwt")
		return 0, *errno.UnauthorizedTokenError
	}

	uid, ok := user_id.(int64)

	if !ok {
		hlog.DefaultLogger().Info("user id assert fail")
		return 0, *errno.UnauthorizedTokenError
	}

	return uid, *errno.Success
}
