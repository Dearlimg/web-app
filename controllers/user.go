package controllers

import (
	"web-app/logic"
	"web-app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(200, gin.H{
				"msg": err.Error(),
			})
		}

		c.JSON(200, gin.H{
			"msg": removeTopStruct(errs.Translate(Trans)),
		})
		return
	}
	//业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp(p) failed", zap.Error(err))
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//返回响应
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
