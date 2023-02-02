package controller

import (
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/app/service"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/httper"
	"hongcha/go-expire/internal/base/mistake"
	"hongcha/go-expire/internal/i18n"
	"net"
)

// Login 登录
// @Summary 登录
// @Description 登录
// @Tags 登录
// @User json
// @Router /login [post]
func Login(ctx *gin.Context) {
	ip := ctx.ClientIP()
	address := net.ParseIP(ip)

	if address == nil {
		httper.HandleResponse(ctx, mistake.NewReqErr(mistake.ErrDatabase, i18n.BaseReqBindError), nil)
		return
	}

	req := &val.GetLoginReq{}

	req.Ip = string(address)

	if httper.BindAndCheck(ctx, req) {
		return
	}

	resp, err := service.Login(req)
	httper.HandleResponse(ctx, err, resp)
}
