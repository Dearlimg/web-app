package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type ParamVoted struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int64  `json:"direction,string" binding:"oneof=1 0 -1"`
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

const (
	OrderTime = "time"
)
