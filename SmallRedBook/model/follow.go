package model

type Follow struct {
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	UserId         string `json:"user_id"`
	FollowedUserId string `json:"followed_user_id"`
	IsFollowed     bool   `json:"is_followed"`
}
