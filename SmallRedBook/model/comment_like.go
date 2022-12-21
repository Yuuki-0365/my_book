package model

type CommentLike struct {
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
	UserId    string `json:"user_id"`
	CommentId int64  `json:"comment_id"`
	IsLiked   bool   `json:"is_liked"`
}
