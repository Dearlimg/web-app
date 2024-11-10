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
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"code": "PAGE_NOT_FOUND", "message": "404"})
	})

	_ = r.Run()
	return r
}

//// JWTAuthMiddleware 基于JWT的认证中间件
//func JWTAuthMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
//		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
//		// 这里的具体实现方式要依据你的实际业务情况决定
//		authHeader := c.Request.Header.Get("Authorization")
//		if authHeader == "" {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2003,
//				"msg":  "请求头中auth为空",
//			})
//			c.Abort()
//			return
//		}
//		// 按空格分割
//		parts := strings.SplitN(authHeader, " ", 2)
//		if !(len(parts) == 2 && parts[0] == "Bearer") {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2004,
//				"msg":  "请求头中auth格式有误",
//			})
//			c.Abort()
//			return
//		}
//		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
//		mc, err := jwt.ParseToken(parts[1])
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2005,
//				"msg":  "无效的Token",
//			})
//			c.Abort()
//			return
//		}
//		// 将当前请求的username信息保存到请求的上下文c上
//		c.Set("userid", mc.UserID)
//		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
//	}
//}
