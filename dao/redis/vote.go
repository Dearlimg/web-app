package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 60 * 60
	ScorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("vote time expire")
)

func VoteForPost(userID, postID string, value float64) error {
	Client.ZScore(getRedisKey(KeyPostTimeZSet), postID)
	if float64(time.Now().Unix())-float64(value) < oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	ov := Client.ZScore(getRedisKey(KeyPostScoreZSet+postID), userID).Val()
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	pipeline := Client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*ScorePerVote, postID)

	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostScoreZSet+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostScoreZSet+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}

func CreatePost(postid int64) error {
	pipeline := Client.TxPipeline()

	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postid,
	}).Result()

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postid,
	})
	_, err := pipeline.Exec()

	return err
}
