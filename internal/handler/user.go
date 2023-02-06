package handler

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"tiktok-gateway/internal/model"
	"tiktok-gateway/kitex_gen/user"
	"tiktok-gateway/kitex_gen/user/douyinservice"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinUserRegisterMethod .
// @router /douyin/user/register [POST]
func DouyinUserRegisterMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := douyinservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.DouyinUserRegisterMethod(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinUserLoginMethod .
// @router /douyin/user/login [POST]
func DouyinUserLoginMethod(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	var req user.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return nil, err
	}

	client, err := douyinservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.DouyinUserLoginMethod(ctx, &req)
	cancel()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DouyinUserMethod .
// @router /douyin/user [GET]
func DouyinUserMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := douyinservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.DouyinUserMethod(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
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
