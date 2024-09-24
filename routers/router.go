package routers

import (
	"bluebell/api/middleware"
	"bluebell/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetupRouter 配置路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}
	// 创建路由实例 新建一个没有任何默认中间件的路由
	r := gin.New()
	//设置中间件
	r.Use(logger.GinLogger(),
		// Recovery 中间件会 recover掉项目可能出现的panic，并使用zap记录相关日志
		logger.GinRecovery(true),
		// 每两秒钟添加十个令牌  全局限流
		middleware.RateLimitMiddleware(2*time.Second, 40),
	)
	r.LoadHTMLFiles("templates/index.html") // 加载html
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusNotFound,
			"msg":  "404",
		})
	})
	return r
}
