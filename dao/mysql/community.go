package mysql

import (
	"web-app/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communitylist []models.Community, err error) {
	sqlstr := "select community_id,community_name from community "
	if err := db.Select(&communitylist, sqlstr); err != nil {
		zap.L().Warn("select community err:", zap.Error(err))
		return nil, err
	}
	return
}
