package controllers

import (
	"strconv"
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

func GetPostDetailHandler(r *gin.Context) {
	//参数校验
	idstr := r.Param("id")
	if len(idstr) == 0 {
		zap.L().Error("logic.GetPostDetail idstr is empty")
		ResponseError(r, CodeInvalidParam)
		return
	}
	id, _ := strconv.ParseInt(idstr, 10, 64)
	//事务处理
	date, err := logic.GetPostByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail err", zap.Error(err))
		ResponseError(r, CodeServerBusy)
	}
	//返回结果
	ResponseSuccess(r, date)
}
