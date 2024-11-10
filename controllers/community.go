package controllers

import (
	"fmt"
	"strconv"
	"web-app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//返回给前端数据
	date, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, date)
}

func CommunityDetailHandler(c *gin.Context) {
	isstr := c.Param("id")
	fmt.Println(isstr)
	id, err := strconv.ParseInt(isstr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err1 := logic.GetCommunityDetailList(id)
	if err1 != nil {
		zap.L().Error("logic.GetCommunityDetailList failed", zap.Error(err1))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
