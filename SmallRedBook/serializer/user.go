package serializer

import (
	"SmallRedBook/conf"
	"SmallRedBook/model"
)

type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	CreateAt string `json:"create_at"`
}

type Follower struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	CreateAt     string `json:"create_at"`
}

type Fan struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	CreateAt     string `json:"create_at"`
}

type SearchUser struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Avatar      string `json:"avatar"`
	FollowCount int64  `json:"follow_count"`
	FanCount    int64  `json:"fan_count"`
}

func BuildUser(user *model.User) *User {
	return &User{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt,
	}
}

func BuildUserInfoInUpdate(user *model.User) map[string]interface{} {
	return map[string]interface{}{
		"UserName":     user.UserName,
		"Sex":          user.Sex,
		"Introduction": user.Introduction,
		"Avatar":       conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
	}
}

func BuildUsers(items []*model.User) (users []*User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return
}

func BuildFollower(user *model.User) *Follower {
	return &Follower{
		UserName:     user.UserName,
		Avatar:       conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		Introduction: user.Introduction,
		CreateAt:     user.CreatedAt,
		UserId:       user.UserId,
	}
}

func BuildFollowers(items []*model.User) (followers []*Follower) {
	for _, item := range items {
		follower := BuildFollower(item)
		followers = append(followers, follower)
	}
	return
}

func BuildFan(user *model.User) *Fan {
	return &Fan{
		UserName:     user.UserName,
		Avatar:       conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		Introduction: user.Introduction,
		CreateAt:     user.CreatedAt,
		UserId:       user.UserId,
	}
}

func BuildFans(items []*model.User) (fans []*Fan) {
	for _, item := range items {
		fan := BuildFan(item)
		fans = append(fans, fan)
	}
	return
}

func BuildSearchUser(user *model.User) (searchUser *SearchUser) {
	return &SearchUser{
		UserId:      user.UserId,
		UserName:    user.UserName,
		Avatar:      conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		FollowCount: user.FollowCount,
		FanCount:    user.FanCount,
	}
}

func BuildSearchUsers(items []*model.User) (searchUsers []*SearchUser) {
	for _, item := range items {
		searchUser := BuildSearchUser(item)
		searchUsers = append(searchUsers, searchUser)
	}
	return
}
