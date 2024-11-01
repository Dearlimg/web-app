package routes

import (
	"net/http"
	"web-app/controllers"
	"web-app/logger"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controllers.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/fuck", func(c *gin.Context) {
		c.String(http.StatusOK, "fuck you!")
	})

	_ = r.Run()
	return r
}

//func Setup() interface{} {
//
//}
