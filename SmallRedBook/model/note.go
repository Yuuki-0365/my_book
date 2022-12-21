package model

type Note struct {
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	NoteId        int64      `json:"note_id" gorm:"column:note_id;primary_key;AUTO_INCREMENT"`
	UserId        string     `json:"user_id"`
	Content       string     `json:"content"`
	Title         string     `json:"title"`
	FilePath      string     `json:"file_path"` // 放图片的路径
	Likes         []Like     `json:"likes" gorm:"foreignKey:note_id"`
	LikeCount     int64      `json:"like_count"`
	Favorites     []Favorite `json:"favorites" gorm:"foreignKey:note_id"`
	FavoriteCount int64      `json:"favorite_count"`
	CommentCount  int64      `json:"comment_count"`
}
