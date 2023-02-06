package handler

import (
	"net/http"
	"pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	//"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err errno.ErrNo) {
	c.JSON(http.StatusOK, Response{
		Code:    int64(err.Code),
		Message: err.Msg,
	})
}
