package logic

import (
	"strconv"
	"web-app/dao/redis"
	"web-app/models"

	"go.uber.org/zap"
)

func PostVote(userID int64, p *models.ParamVoted) error {
	zap.L().Debug("logic.PostVote",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
