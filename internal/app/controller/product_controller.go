package controller

import (
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/app/service"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/httper"
)

// AddProduct 添加商品
// @Summary 添加商品
// @Description 添加商品
// @Tags 添加商品
// @Produce json
// @Router /product/add [post]
func AddProduct(ctx *gin.Context) {
	req := &val.GetProductReq{}
	if httper.BindAndCheck(ctx, req) {
		return
	}
	resp, err := service.AddProduct(req)
	httper.HandleResponse(ctx, err, resp)
}

// GetProductListPage 查询物品列表 分页
// @Summary 查询物品列表 分页
// @Description 查询作者列表 分页
// @Tags 查询物品列表
// @Produce json
// @Param name query string false "物品名称 默认空"
// @Param page query int false "页码 默认为1"
// @Param page_size query int false "每页大小 默认为10"
// @Router /product/page [get]
func GetProductListPage(ctx *gin.Context) {
	req := &val.GetProductWithPageReq{}

	if httper.BindAndCheck(ctx, req) {
		return
	}

	resp, err := service.GetProductListPage(req)
	httper.HandleResponse(ctx, err, resp)
}
