package routers

import (
	"api-server/handler/hello"
	"api-server/router/middleware"

	"github.com/1024casts/snake/handler"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 使用中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Logging())
	g.Use(middleware.RequestID())
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(handler.RouteNotFound)
	g.NoMethod(handler.RouteNotFound)

	// 静态资源，主要是图片
	g.Static("/static", "./static")

	// swagger api docs
	//g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router 性能分析路由
	// 默认关闭，开发环境下可以打开
	// 访问方式: HOST/debug/pprof
	// 通过 HOST/debug/pprof/profile 生成profile
	// 查看分析图 go tool pprof -http=:5000 profile
	// pprof.Register(g)

	// 认证相关路由
	g.GET("/v1/hello", hello.Hello)
	return g
}
