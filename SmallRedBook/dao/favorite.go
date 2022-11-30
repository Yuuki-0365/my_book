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
func (dao *FavoriteDao) FavoriteOrNot(userId string, noteId int64) (count int64, err error) {
	err = dao.DB.Table("favorite").
		Where("user_id=? and note_id=?", userId, noteId).
		Count(&count).Error
	return
}

func (dao *FavoriteDao) FavoriteNote(favorite *model.Favorite) (err error) {
	err = dao.DB.Table("favorite").
		Create(&favorite).Error
	return
}

func (dao *FavoriteDao) UnFavoriteNote(userId string, noteId int64) (err error) {
	err = dao.DB.Where("user_id=? and note_id=?", userId, noteId).
		Delete(&model.Favorite{}).
		Error
	return
}
