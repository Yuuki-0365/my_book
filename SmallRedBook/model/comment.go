package model

type Comment struct {
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	NoteId       int64         `json:"note_id"`
	UserId       string        `json:"user_id"`
	UserName     string        `json:"user_name"`
	Content      string        `json:"content"`
	Avatar       string        `json:"avatar"`
	LikeCount    int64         `json:"like_count"`
	CommentID    int64         `json:"comment_id" gorm:"column:comment_id;primary_key;AUTO_INCREMENT"`
	CommentLikes []CommentLike `json:"comment_likes" gorm:"foreignKey:comment_id"`
	ParentId     int64         `json:"parent_id"` // 上一级id
	Type         int           `json:"type"`      // 评论的是note(0) or comment(1)
	CommentCount int64         `json:"comment_count"`
}
