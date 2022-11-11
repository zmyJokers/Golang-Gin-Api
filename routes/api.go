package routes

import (
	"github.com/gin-gonic/gin"
	"middleground/app/controller"
	"middleground/app/controller/log"
	"middleground/app/middleware"
	"net/http"
)

/*
	引入文件路径
	Action 控制器列表
*/
type Controller struct {
	LogController  log.Controller // 日志控制器
	TestController controller.TestController
}

/*
	func RouteApiInit 初始化路由
*/
func RouteApiInit() *gin.Engine {

	var controllerInit Controller

	r := gin.New()

	// panic 异常监控
	r.Use(middleground.Recovery())

	// 响应记录
	r.Use(middleground.Return())

	// 路由执行动作 --》 控制器/业务控制器/业务动作
	r.GET("/api/test", controllerInit.TestController.TestAction)
	r.POST("/api/registerLog", controllerInit.LogController.LogSendAction) // 日志推送
	r.GET("/api/test1", func(context *gin.Context) {
		context.String(200, "测试一下")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404请求页面不存在")
		panic("404请求页面不存在")
	})
	return r
}
