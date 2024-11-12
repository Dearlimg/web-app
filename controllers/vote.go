package controllers

import (
	"web-app/logic"
	"web-app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandler(r *gin.Context) {
	//参数
	p := new(models.ParamVoted)
	if err := r.ShouldBindJSON(p); err != nil {
		zap.L().Error("VoteHandler fail :", zap.Error(err))
		ResponseError(r, CodeInvalidParam)
		return
	}
	//业务
	logic.PostVote()
	//返回
	ResponseSuccess(r, nil)
}
