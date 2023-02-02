package controller

import (
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/app/service"
	"hongcha/go-expire/internal/base/httper"
)

// GetRemindAll 获取所有提醒
// @Summary 获取所有提醒
// @Description 获取所有提醒
// @Tags 获取所有提醒
// @Produce json
// @Router /remind [get]
func GetRemindAll(ctx *gin.Context) {
	remindAll, err := service.GetRemindAll()
	httper.HandleResponse(ctx, err, remindAll)
}
