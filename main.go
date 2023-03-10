// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"tiktok-gateway/internal/middleware"
	routers "tiktok-gateway/internal/routers"
)

func main() {
	middleware.InitJwt()
	h := server.Default()
	routers.Register(h)
	h.Spin()
}
