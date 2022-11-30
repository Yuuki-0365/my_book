package model

type User struct {
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	UserId       string   `json:"user_id" gorm:"primaryKey"`
	UserName     string   `json:"user_name"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Avatar       string   `json:"avatar"` // 默认头像
	Introduction string   `json:"introduction"`
	Sex          string   `json:"sex"`
	Followers    []Follow `json:"followers" gorm:"foreignKey:UserId"`
	FollowCount  int64    `json:"follow_count"`
	Fans         []Fan    `json:"fans" gorm:"foreignKey:UserId"`
	FanCount     int64    `json:"fan_count"`
	Notes        []Note   `json:"notes" gorm:"foreignKey:UserId"`
	NoteCount    int64    `json:"note_count"`
}
