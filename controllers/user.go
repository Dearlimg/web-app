package controllers

import (
	"errors"
	"web-app/dao/mysql"
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
			//c.JSON(200, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(Trans)))
		//c.JSON(200, gin.H{
		//	"msg": removeTopStruct(errs.Translate(Trans)),
		//})
		return
	}
	//业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp(p) failed", zap.Error(err))
		//c.JSON(200, gin.H{
		//	"msg": "",
		//})
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//参数校验绑定
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//c.JSON(200, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			//ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(Trans)))
		//c.JSON(200, gin.H{
		//	"msg": removeTopStruct(errs.Translate(Trans)),
		//})
		//ResponseErrorWithMsh(c, CodeInvalidParam)
		return
	}
	//业务处理
	token, err1 := logic.Login(&p)
	if err1 != nil {
		zap.L().Error("logic.Login(p) failed", zap.Error(err1))
		//c.JSON(200, gin.H{
		//	"msg": err1.Error(),
		//})
		if errors.Is(err1, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		//ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	//c.JSON(200, gin.H{
	//	"msg": "登录成功",
	//})
	ResponseSuccess(c, token)
}
