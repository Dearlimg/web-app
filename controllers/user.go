package controllers

import (
	"fmt"
	"web-app/logic"
	"web-app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) with invalid param", zap.Error(err))
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}

	fmt.Println(p)

	//业务处理
	logic.SignUp(c)
	//返回响应
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
