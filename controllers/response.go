package controllers

import "github.com/gin-gonic/gin"

type ResponseDate struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Date interface{} `json:"date"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(200, &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Date: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(200, &ResponseDate{
		Code: code,
		Msg:  msg,
		Date: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Date: data,
	})
}
