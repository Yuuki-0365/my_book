package dao

import (
	"SmallRedBook/model"
	"context"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDb(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}
func (dao *FavoriteDao) FavoriteOrNot(userId string, noteId int64) (favorite *model.Favorite, count int64, err error) {
	err = dao.DB.Table("favorite").
		Where("user_id=? and note_id=?", userId, noteId).
		Find(&favorite).
		Count(&count).Error
	return
}

func (dao *FavoriteDao) CreateFavoriteNote(favorite *model.Favorite) (err error) {
	tx := dao.DB
	err = tx.Table("favorite").
		Create(&favorite).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", favorite.NoteId).
		Update("favorite_count", gorm.Expr("favorite_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}
func (dao *FavoriteDao) UnFavorite(userId string, noteId int64) (err error) {
	tx := dao.DB
	err = tx.Model(&model.Favorite{}).
		Where("note_id = ? and user_id = ?", noteId, userId).
		Update("is_favorite", false).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", noteId).
		Update("favorite_count", gorm.Expr("favorite_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}

func (dao *FavoriteDao) Favorite(userId string, noteId int64) (err error) {
	tx := dao.DB
	err = tx.Model(&model.Favorite{}).
		Where("note_id = ? and user_id = ?", noteId, userId).
		Update("is_favorite", true).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", noteId).
		Update("favorite_count", gorm.Expr("favorite_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}
