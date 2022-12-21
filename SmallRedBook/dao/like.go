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

func (dao *LikeDao) LikeOrNot(userId string, noteId int64) (Like *model.Like, count int64, err error) {
	err = dao.DB.Table("like").
		Where("user_id=? and note_id=?", userId, noteId).
		Find(&Like).
		Count(&count).Error
	return
}

func (dao *LikeDao) CreateLikeNote(like *model.Like) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Table("like").
		Create(&like).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", like.NoteId).
		Update("like_count", gorm.Expr("like_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}

func (dao *LikeDao) UnLike(userId string, noteId int64) (err error) {
	tx := dao.DB
	err = tx.Model(&model.Like{}).
		Where("note_id = ? and user_id = ?", noteId, userId).
		Update("is_liked", false).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", noteId).
		Update("like_count", gorm.Expr("like_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}

func (dao *LikeDao) Like(userId string, noteId int64) (err error) {
	tx := dao.DB
	err = tx.Model(&model.Like{}).
		Where("note_id = ? and user_id = ?", noteId, userId).
		Update("is_liked", true).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", noteId).
		Update("like_count", gorm.Expr("like_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}
