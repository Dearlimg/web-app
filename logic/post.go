package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
)

func Post(date *models.Post) (err error) {
	return mysql.InsertPost(date)
}

func GetPostByID(ID int64) (post *models.Post, err error) {
	return mysql.GetPostByID(ID)
}
