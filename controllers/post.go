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

func GetPostsHandler(r *gin.Context) {
	//数据
	page, size, err := GetPostParam(r)
	if err != nil {
		zap.L().Error("GetPostParam err", zap.Error(err))
		ResponseError(r, CodeInvalidParam)
		return
	}
	//业务处理
	date, err := logic.GetPosts(page, size)
	if err != nil {
		zap.L().Error("logic.GetPosts err", zap.Error(err))
		ResponseError(r, CodeServerBusy)
	}
	//返回
	ResponseSuccess(r, date)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
//func GetPostListHandler2(c *gin.Context) {
//	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
//	// 初始化结构体时指定初始参数
//	p := &models.ParamPostList{
//		Page:  1,
//		Size:  10,
//		Order: models.OrderTime,
//	}
//
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	data, err := logic.GetPostListNew(p)
//	// 获取数据
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//	// 返回响应
//}
