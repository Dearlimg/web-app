package models

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreatedTime time.Time `json:"create_time" db:"create_time"`
}

type ApiPost struct {
	Username          string `json:"username" db:"username"`
	*Post             `json:"post" db:"post"`
	*CommunityDetails `json:"community" db:"community"`
}
