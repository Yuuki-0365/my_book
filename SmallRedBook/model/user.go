package model

type User struct {
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	UserId        string        `json:"user_id" gorm:"primaryKey"`
	UserName      string        `json:"user_name"`
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	Avatar        string        `json:"avatar"` // 默认头像
	Introduction  string        `json:"introduction"`
	Sex           string        `json:"sex"`
	Notes         []Note        `json:"notes" gorm:"foreignKey:UserId"`
	NoteCount     int64         `json:"note_count"`
	Likes         []Like        `json:"likes" gorm:"foreignKey:UserId"`
	LikeCount     int64         `json:"like_count"`
	Favorites     []Favorite    `json:"favorites" gorm:"foreignKey:UserId"`
	FavoriteCount int64         `json:"favorite_count"`
	FollowCount   int64         `json:"follow_count"`
	FanCount      int64         `json:"fan_count"`
	CommentLikes  []CommentLike `json:"comment_likes" gorm:"foreignKey:UserId"`
}
