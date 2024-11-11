package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
)

func Post(date *models.Post) (err error) {
	return mysql.InsertPost(date)
}
