package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
)

func GetCommunityList() ([]models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetailList(id int64) (date *models.CommunityDetails, err error) {
	return mysql.GetCommunityDetailList(id)
}
