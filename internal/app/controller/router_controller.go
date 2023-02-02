package controller

import "github.com/gin-gonic/gin"

func RouterRegisterAuth(r *gin.RouterGroup) {
	// 添加商品
	r.POST("/product/add", AddProduct)
	// 查询物品列表 分页
	r.GET("/product/page", GetProductListPage)
	// 查询所有分类
	r.GET("/category", GetCategoryAll)
	// 添加分类
	r.POST("/category/add", AddCategory)
	// 查询所有提醒
	r.GET("/remind", GetRemindAll)
}

// RouterRegisterUnAuth 不需要鉴权
func RouterRegisterUnAuth(r *gin.RouterGroup) {
	r.POST("/login", Login)
}
