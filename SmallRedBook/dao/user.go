package dao

import (
	"SmallRedBook/model"
	"context"
	"gorm.io/gorm"
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
		Where("user_id = ? and is_followed = ?", userId, true)
	err = dao.DB.Table("user").
		Select("*").
		Where("user_id in (?)", tmp).
		Find(&user).
		Count(&count).Error
	return
}

func (dao *UserDao) GetFansById(userId string) (count int64, user []*model.User, err error) {
	tmp := dao.DB.Table("follow").
		Select("user_id").
		Where("followed_user_id = ? and is_followed = ?", userId, true)

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
func (dao *UserDao) UserFollowed(userId string, followedId string) (follow *model.Follow, count int64, err error) {
	err = dao.DB.Table("follow").
		Where("user_id = ? and followed_user_id = ?", userId, followedId).
		Find(&follow).
		Count(&count).Error
	return
}

func (dao *UserDao) CreateFollow(follow *model.Follow) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Model(&model.Follow{}).
		Create(&follow).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Model(&model.User{}).
		Where("user_id=?", follow.UserId).
		Update("follow_count", gorm.Expr("follow_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Model(&model.User{}).
		Where("user_id=?", follow.FollowedUserId).
		Update("fan_count", gorm.Expr("fan_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (dao *UserDao) Follow(userId string, followId string) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Model(&model.Follow{}).
		Where("user_id = ? and followed_user_id = ?", userId, followId).
		Update("is_followed", true).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Model(&model.User{}).
		Where("user_id=?", userId).
		Update("follow_count", gorm.Expr("follow_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Model(&model.User{}).
		Where("user_id=?", followId).
		Update("fan_count", gorm.Expr("fan_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (dao *UserDao) UnFollow(userId string, followId string) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Model(&model.Follow{}).
		Where("user_id = ? and followed_user_id = ?", userId, followId).
		Update("is_followed", false).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Model(&model.User{}).
		Where("user_id=?", userId).
		Update("follow_count", gorm.Expr("follow_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Model(&model.User{}).
		Where("user_id=?", followId).
		Update("fan_count", gorm.Expr("fan_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
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
		Select("user_id, user_name, sex, avatar, introduction, follow_count, fan_count, note_count").
		Where("user_id=?", userId).
		Find(&user).Error
	return
}

func (dao *UserDao) ShowOwnUserInfoAll(userId string) (user map[string]interface{}, err error) {
	err = dao.DB.Table("user").
		Select("*").
		Where("user_id=?", userId).
		Find(&user).Error
	return
}

func (dao *UserDao) GetUpdateInfo(userId string) (user map[string]interface{}, err error) {
	err = dao.DB.Table("user").
		Select("avatar, user_name, user_id, introduction, sex, email, password").
		Where("user_id=?", userId).
		Find(&user).Error
	return
}

func (dao *UserDao) DeleteUser(userId string) (err error) {
	err = dao.DB.Where("user_id=?", userId).
		Delete(&model.User{}).Error
	return
}

func (dao *UserDao) AdminShowInfo() (user []map[string]interface{}, err error) {
	err = dao.DB.Table("user").
		Select("created_at, updated_at, user_id, user_name, email, password, introduction, sex").
		Find(&user).Error
	return
}

func (dao *UserDao) AdminUpdateInfo(user *model.User, userId string) (err error) {
	err = dao.DB.Model(&model.User{}).
		Where("user_id=?", userId).
		Updates(user).Error
	return
}

func (dao *UserDao) AdminDeleteInfo(userId string) (err error) {
	err = dao.DB.Where("user_id=?", userId).
		Delete(&model.User{}).Error
	return
}

func (dao *UserDao) GetFollowUserNotes(userId string) (notes []map[string]interface{}, count int64, err error) {
	tmp := dao.DB.Table("follow").
		Select("followed_user_id").
		Where("user_id=?", userId)
	err = dao.DB.Table("note, user").
		Select("note.*, user.avatar, user.user_name").
		Where("note.user_id in(?) and note.user_id = user.user_id", tmp).
		Order("note.created_at desc").
		Find(&notes).
		Count(&count).Error
	return
}
