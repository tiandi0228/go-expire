package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/app/controller"
	"hongcha/go-expire/internal/app/service"
	"hongcha/go-expire/internal/base/router/middleware"
	"net/http"
)

// InitWebRouter 初始化路由 唯一一个会阻塞的初始化动作，请放置在最后
func InitWebRouter(port string) {
	// 判断是否一release模式启动
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 开发阶段 使用gin的Recovery将日志格式化输出到控制台
	r.Use(gin.Recovery())
	// 生产阶段 使用zap输出到固定文件
	r.Use(RecoveryWithZap())
	// 解决前端跨域问题
	r.Use(CorsMiddleware())

	// 需要鉴权
	v1 := r.Group("/api/v1/expire")
	v1.Use(middleware.Authorize)
	controller.RouterRegisterAuth(v1)

	// 不需要鉴权
	v2 := r.Group("/api/v1/expire")
	controller.RouterRegisterUnAuth(v2)

	if err := r.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		panic(err)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Access-token")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{
				"code":    20000,
				"message": "Options Request!",
			})
		}

		c.Next()
	}
}

func CornInit() {
	go service.RunCorn()
}
