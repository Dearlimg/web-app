package controllers

import (
	"web-app/logic"
	"web-app/models"
	"web-app/pkg/snowflake"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func PostHandler(r *gin.Context) {
	//处理数据
	date := new(models.Post)
	if err := r.ShouldBindJSON(date); err != nil {
		zap.L().Error("r.ShouldBindJSON(date)", zap.Error(err))
		ResponseError(r, CodeInvalidParam)
		return
	}
	date.AuthorID, _ = getCurrentUser(r)
	date.ID = snowflake.GenID()

	//业务处理
	if err := logic.Post(date); err != nil {
		ResponseError(r, CodeInvalidParam)
		return
	}
	//返回
	ResponseSuccess(r, CodeSuccess)
}
