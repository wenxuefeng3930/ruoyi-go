package system

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	"robvi/app/system/controller"
)

// 加载路由
func init() {
	// 服务监控
	g1 := router.New("/monitor/server", token.TokenMiddleware(), auth.Auth)
	serverController := controller.ServiceController{}
	g1.GET("/", "monitor:server:view", serverController.Server)

	//登陆日志
	g2 := router.New("/monitor/logininfor", token.TokenMiddleware(), auth.Auth)
	loginInforController := controller.LoginInforController{}
	g2.GET("/", "monitor:logininfor:view", loginInforController.List)
	g2.POST("/list", "monitor:logininfor:list", loginInforController.ListAjax)
	g2.POST("/export", "monitor:logininfor:export", loginInforController.Export)
	g2.POST("/clean", "monitor:logininfor:remove", loginInforController.Clean)
	g2.POST("/remove", "monitor:logininfor:remove", loginInforController.Remove)
	g2.POST("/unlock", "monitor:logininfor:unlock", loginInforController.Unlock)

	//操作日志
	g3 := router.New("/monitor/operlog", token.TokenMiddleware(), auth.Auth)
	operController := controller.OperlogController{}
	g3.GET("/", "monitor:operlog:view", operController.List)
	g3.POST("/list", "monitor:operlog:list", operController.ListAjax)
	g3.POST("/export", "monitor:operlog:export", operController.Export)
	g3.POST("/remove", "monitor:operlog:export", operController.Remove)
	g3.POST("/clean", "monitor:operlog:export", operController.Clean)
	g3.GET("/detail", "monitor:operlog:detail", operController.Detail)

	//在线用户
	g4 := router.New("/monitor/online", token.TokenMiddleware(), auth.Auth)
	onlineController := controller.OnlineController{}
	g4.GET("/", "monitor:online:view", onlineController.List)
	g4.POST("/list", "monitor:online:list", onlineController.ListAjax)
	g4.POST("/forceLogout", "monitor:online:forceLogout", onlineController.ForceLogout)
	g4.POST("/batchForceLogout", "monitor:online:batchForceLogout", onlineController.BatchForceLogout)
}
