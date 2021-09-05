package router

import (
	"cheese/controller"
	"cheese/logger"
	"cheese/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
	{
		//community
		v1.GET("/community", controller.GetCommunityHandler)
		v1.GET("/community/:id", controller.GetDetailHandler)
		//post
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/postorderlist", controller.GetOrderListHandler)
		v1.GET("/postlist", controller.GetPostListHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
