package dao

import (
	"SmallRedBook/model"
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDb(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistEmailOrNot(email string) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).
		Where("email=?", email).
		Count(&count).Error
	if err != nil {
		return true, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Table("user").
		Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) ExistUserIdOrNot(userId string, password string) (count int64, err error) {
	err = dao.DB.Table("user").
		Where("user_id=? and password=?", userId, password).
		Count(&count).Error
	return
}

func (dao *UserDao) GetUserByUserIdAndPassword(userId string, password string) (user *model.User, err error) {
	err = dao.DB.Table("user").
		Where("user_id=? and password=?", userId, password).
		First(&user).Error
	return
}

func (dao *UserDao) GetUserByEmail(email string) (count int64, user *model.User, err error) {
	err = dao.DB.Table("user").
		Where("email=?", email).
		Find(&user).
		Count(&count).Error
	return
}

func (dao *UserDao) GetEmailById(userId string) (email string, err error) {
	err = dao.DB.Table("user").
		Select("email").
		Where("user_id=?", userId).
		Find(&email).Error
	return
}

func (dao *UserDao) GetFollowedCount(userId string) (count int64, err error) {
	err = dao.DB.Table("user").
		Select("follow_count").
		Where("user_id=?", userId).
		Find(&count).Error
	return
}

func (dao *UserDao) GetFansCount(userId string) (count int64, err error) {
	err = dao.DB.Table("user").
		Select("fan_count").
		Where("user_id=?", userId).
		Find(&count).Error
	return
}

func (dao *UserDao) GetFollowersById(userId string) (count int64, user []*model.User, err error) {
	tmp := dao.DB.Table("follow").
		Select("followed_user_id").
		Where("user_id = ?", userId)
	err = dao.DB.Table("user").
		Select("*").
		Where("user_id in (?)", tmp).
		Find(&user).
		Count(&count).Error
	return
}

func (dao *UserDao) GetFansById(userId string) (count int64, user []*model.User, err error) {
	tmp := dao.DB.Table("fan").
		Select("follower_user_id").
		Where("user_id = ?", userId)

	err = dao.DB.Table("user").
		Select("*").
		Where("user_id in (?)", tmp).
		Find(&user).
		Count(&count).Error
	return
}

func (dao *UserDao) GetUserInfoByUserId(userId string) (user *model.User, err error) {
	err = dao.DB.Table("user").
		Where("user_id=?", userId).
		Find(&user).Error
	return
}

func (dao *UserDao) GetUserInfoInUpdate(userId string) (user *model.User, err error) {
	err = dao.DB.Table("user").
		Select("user_name, avatar, introduction, sex").
		Where("user_id=?", userId).
		First(&user).Error
	return
}

func (dao *UserDao) UserFollowed(userId string, followedId string) (count int64, err error) {
	err = dao.DB.Table("follow").
		Where("user_id = ? and followed_user_id = ?", userId, followedId).
		Count(&count).Error
	return
}

func (dao *UserDao) UserFollow(userId string, followId string) (err error) {
	tx := dao.DB
	tx.Begin()
	follow := model.Follow{
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		UserId:         userId,
		FollowedUserId: followId,
		Status:         1,
	}
	err = tx.Create(&follow).Error
	if err != nil {
		tx.Rollback()
		return
	}

	fan := model.Fan{
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		UserId:         followId,
		FollowerUserId: userId,
		Status:         1,
	}
	err = tx.Create(&fan).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Table("user").
		Where("user_id=?", userId).
		Update("follow_count", gorm.Expr("follow_count+?", 1)).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("user").
		Where("user_id=?", followId).
		Update("fan_count", gorm.Expr("fan_count+?", 1)).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (dao *UserDao) UserUnFollow(userId string, unFollowId string) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Where("user_id=?", userId).
		Delete(&model.Follow{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Where("user_id=?", unFollowId).
		Delete(&model.Fan{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("user").
		Where("user_id=?", userId).
		Update("follow_count", gorm.Expr("follow_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("user").
		Where("user_id=?", unFollowId).
		Update("fan_count", gorm.Expr("fan_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (dao *UserDao) UpdateEmail(userId string, email string, updateTime string) (err error) {
	err = dao.DB.Model(&model.User{}).
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"email": email, "updated_at": updateTime}).Error
	return
}

func (dao *UserDao) UpdatePassword(userId string, password string, updateTime string) (err error) {
	err = dao.DB.Model(&model.User{}).
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"password": password, "updated_at": updateTime}).Error
	return
}

func (dao *UserDao) UpdateInfo(userId string, user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).
		Where("user_id=?", userId).
		Save(&user).Error
	return
}

func (dao *UserDao) SearchUser(userName string) (users []*model.User, count int64, err error) {
	err = dao.DB.Table("user").
		Select("user_name, user_id, avatar, follow_count, fan_count").
		Where("user_name like ?", "%"+userName+"%").
		Find(&users).
		Count(&count).Error
	return
}

func (dao *UserDao) ShowUserInfoAll(userId string) (user map[string]interface{}, err error) {
	err = dao.DB.Table("user").
		Select("user_id, user_name, avatar, introduction, follow_count, fan_count, note_count").
		Where("user_id=?", userId).
		Find(&user).Error
	return
}
