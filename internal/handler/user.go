package handler

import (
	"context"
	"tiktok-gateway/internal/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinUserRegisterMethod .
// @router /douyin/user/register [POST]
func DouyinUserRegisterMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinUserLoginMethod .
// @router /douyin/user/login [POST]
func DouyinUserLoginMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinUserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinUserMethod .
// @router /douyin/user [GET]
func DouyinUserMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationActionMethod .
// @router /douyin/relation/action [POST]
func DouyinRelationActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinRelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFollowListMethod .
// @router /douyin/relation/follow/list [GET]
func DouyinRelationFollowListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinRelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFollowerListMethod .
// @router /douyin/relation/follower/list [GET]
func DouyinRelationFollowerListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinRelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFriendListMethod .
// @router /douyin/relation/friend/list [GET]
func DouyinRelationFriendListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinRelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}
