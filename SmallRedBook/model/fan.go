package model

type Fan struct {
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	ID             int64  `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId         string `json:"user_id"`
	FollowerUserId string `json:"follower_user_id"`
	Status         int64  `json:"status"`
}
