package mysql

import (
	"web-app/models"

	"go.uber.org/zap"
)

func InsertPost(date *models.Post) (err error) {
	sqlstr := "insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlstr, date.ID, date.Title, date.Content, date.AuthorID, date.CommunityID)
	if err != nil {
		return
	}
	return
}

func GetPostByID(postID int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlstr := "select post_id,title,content,author_id,community_id,create_time from post where post_id = ?"
	err = db.Get(post, sqlstr, postID)
	if err != nil {
		zap.L().Error("GetPostByID", zap.Error(err))
		return
	}
	return
}
