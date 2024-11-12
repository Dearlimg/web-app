package logic

import (
	"web-app/dao/mysql"
	"web-app/models"

	"go.uber.org/zap"
)

func Post(date *models.Post) (err error) {
	return mysql.InsertPost(date)
}

func GetPostByID(ID int64) (post *models.ApiPost, err error) {
	date, err1 := mysql.GetPostByID(ID)
	if err1 != nil {
		zap.L().Error("mysql.GetPostByID fail", zap.Error(err1))
		return
	}
	Username, err2 := mysql.GetUserByID(date.AuthorID)
	if err2 != nil {
		zap.L().Error("mysql.GetUserByID fail", zap.Error(err2))
		return
	}
	Community, err3 := mysql.GetCommunityDetailList(date.CommunityID)
	if err3 != nil {
		zap.L().Error("mysql.GetCommunityDetailList fail", zap.Error(err3))
		return
	}
	post = &models.ApiPost{
		Username:         Username.Username,
		CommunityDetails: Community,
		Post:             date,
	}
	return post, err
}

func GetPosts(page, size int64) ([]*models.ApiPost, error) {
	apiposts := make([]*models.ApiPost, 0)
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList fail", zap.Error(err))
		return nil, err
	}
	for _, v := range posts {
		Username, err := mysql.GetUserByID(v.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID fail", zap.Error(err))
			return nil, err
		}
		Community, err := mysql.GetCommunityDetailList(v.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailList fail", zap.Error(err))
			return nil, err
		}
		apipost := &models.ApiPost{
			Username:         Username.Username,
			CommunityDetails: Community,
			Post:             v,
		}
		apiposts = append(apiposts, apipost)
	}
	return apiposts, nil
}
