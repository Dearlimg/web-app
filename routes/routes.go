package routes

import (
	"web-app/controllers"
	"web-app/logger"
	"web-app/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.PostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"code": "PAGE_NOT_FOUND", "message": "404"})
	})

	_ = r.Run()
	return r
}
