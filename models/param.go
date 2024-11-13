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
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int64 `json:"direction,string" binding:"required,oneof=1 0 -1"`
}
