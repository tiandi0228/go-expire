package controller

import (
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/app/service"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/httper"
)

// GetCategoryAll 查询所有分类
// @Summary 查询所有分类
// @Description 查询所有分类
// @Tags 查询所有分类
// @Produce json
// @Router /category [get]
func GetCategoryAll(ctx *gin.Context) {
	categoryAll, err := service.GetCategoryAll()
	httper.HandleResponse(ctx, err, categoryAll)
}

// AddCategory 添加分类
// @Summary 添加分类
// @Description 添加分类
// @Tags 添加分类
// @Produce json
// @Router /category/add [post]
func AddCategory(ctx *gin.Context) {
	req := &val.GetCategoryReq{}
	if httper.BindAndCheck(ctx, req) {
		return
	}
	resp, err := service.AddCategory(req)
	httper.HandleResponse(ctx, err, resp)
}
