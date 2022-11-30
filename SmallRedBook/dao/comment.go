package dao

import (
	"SmallRedBook/model"
	"context"
	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{NewDBClient(ctx)}
}

func NewCommentDaoByDb(db *gorm.DB) *CommentDao {
	return &CommentDao{db}
}

func (dao *CommentDao) CommentOrNot(userId string, noteId int64) (count int64, err error) {
	err = dao.DB.Table("comment").
		Where("user_id=? and note_id=?", userId, noteId).
		Count(&count).Error
	return
}

func (dao *CommentDao) AddComment(comment *model.Comment) (err error) {
	err = dao.DB.Table("comment").
		Create(&comment).Error
	return
}

func (dao *CommentDao) DeleteComment(userId string, noteId int64) (err error) {
	err = dao.DB.Where("user_id=? and note_id=?", userId, noteId).
		Delete(&model.Comment{}).
		Error
	return
}

func (dao *CommentDao) GetCommentInfo(noteId int64) (comments []map[string]interface{}, err error) {
	err = dao.DB.Table("comment").
		Select("user_name, content, avatar, user_id").
		Where("note_id=?", noteId).
		Order("create_at").
		Find(&comments).Error
	return
}
