package model

type Comment struct {
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
	NoteId    uint   `json:"note_id"`
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Content   string `json:"content"`
	Avatar    string `json:"avatar"`
	LikeCount int64  `json:"like_count"`
	ID        int64  `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
}
