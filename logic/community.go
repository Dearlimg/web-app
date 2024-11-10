package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
)

func GetCommunityList() ([]models.Community, error) {
	return mysql.GetCommunityList()
}
