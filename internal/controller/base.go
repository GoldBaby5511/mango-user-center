package controller

import (
	"mango-user-center/internal/middleware"
	"mango-user-center/pkg/response"

	"github.com/gin-gonic/gin"
)

var (
	Engine = NewEngine()
)

func init() {
}

func NewEngine() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	// engine.Use(middleware.Cors())
	engine.Use(
		middleware.ApiLog(),         // 日志记录
		middleware.CustomRecovery(), // 错误捕获
		middleware.Cors(),           // 跨域
		middleware.RateLimit(),      // 限流
	)

	// 404
	engine.NoRoute(func(c *gin.Context) {
		response.Echo(c, nil, response.Msg("URL does not exist."))
	})
	// ping
	engine.GET("ping", func(c *gin.Context) {
		response.Echo(c, "pong", nil)
	})

	return engine
}
