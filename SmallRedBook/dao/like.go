package dao

import (
	"SmallRedBook/model"
	"context"
	"gorm.io/gorm"
)

type LikeDao struct {
	*gorm.DB
}

func NewLikeDao(ctx context.Context) *LikeDao {
	return &LikeDao{NewDBClient(ctx)}
}

func NewLikeDaoByDb(db *gorm.DB) *LikeDao {
	return &LikeDao{db}
}

func (dao *LikeDao) LikeOrNot(userId string, noteId int64) (count int64, err error) {
	err = dao.DB.Table("like").
		Where("user_id=? and note_id=?", userId, noteId).
		Count(&count).Error
	return
}

func (dao *LikeDao) LikeNote(like *model.Like) (err error) {
	err = dao.DB.Table("like").
		Create(&like).Error
	return
}

func (dao *LikeDao) UnLikeNote(userId string, noteId int64) (err error) {
	err = dao.DB.Where("user_id=? and note_id=?", userId, noteId).
		Delete(&model.Like{}).
		Error
	return
}
