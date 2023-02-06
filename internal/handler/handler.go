package handler

import (
	"github.com/41197-yhkt/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-gateway/internal/model"
	//"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, err errno.ErrNo) {
	c.JSON(http.StatusOK, douyin.BaseResp{
		StatusCode: int32(err.Code),
		StatusMsg:  &err.Msg,
	})
}
