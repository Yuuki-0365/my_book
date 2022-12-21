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

func (dao *CommentDao) Count() (count int64, err error) {
	err = dao.DB.Model(&model.Comment{}).
		Count(&count).Error
	if count == 0 {
		return
	} else {
		err = dao.DB.Table("comment").Select("MAX(comment_id)").Find(&count).Error
	}
	return
}

func (dao *CommentDao) GetCommentInfo(noteId int64) (comments []map[string]interface{}, err error) {
	err = dao.DB.Model(&model.Comment{}).
		Where("type = ? and parent_id = ?", 0, noteId).
		Find(&comments).Error
	return
}

func (dao *CommentDao) AddComment(comment *model.Comment) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Model(&model.Comment{}).Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("note").
		Where("note_id=?", comment.NoteId).
		Update("comment_count", gorm.Expr("comment_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if comment.Type == 1 {
		err = tx.Table("comment").
			Where("comment_id=?", comment.ParentId).
			Update("comment_count", gorm.Expr("comment_count+1")).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

func (dao *CommentDao) DeleteComment(userId string, noteId int64, parentId int64, commentId int64, Type int) (err error) {
	if Type == 0 {
		tx := dao.DB
		tx.Begin()
		err = tx.Where("user_id=? and parent_id=? and note_id = ? and comment_id=? and type=0", userId, parentId, parentId, commentId).
			Delete(&model.Comment{}).Error
		if err != nil {
			tx.Rollback()
			return
		}
		var count int64
		err = tx.Model(&model.Comment{}).
			Where("parent_id=? and type=1", commentId).
			Count(&count).Error
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Where("parent_id=? and type=1", commentId).
			Delete(&model.Comment{}).Error
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Table("note").
			Where("note_id=?", noteId).
			Update("comment_count", gorm.Expr("comment_count-?", count+1)).Error
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	} else {
		tx := dao.DB
		tx.Begin()
		err = tx.Where("user_id=? and parent_id=? and comment_id=? and type=1", userId, parentId, commentId).
			Delete(&model.Comment{}).Error
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Table("note").
			Where("note_id=?", noteId).
			Update("comment_count", gorm.Expr("comment_count-1")).Error
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Table("Comment").
			Where("comment_id=?", parentId).
			Update("comment_count", gorm.Expr("comment_count-1")).Error
		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}
	return
}

func (dao *CommentDao) LikeOrNot(userId string, commentId int64) (Like *model.CommentLike, count int64, err error) {
	err = dao.DB.Table("comment_like").
		Where("user_id=? and comment_id=?", userId, commentId).
		Find(&Like).
		Count(&count).Error
	return
}

func (dao *CommentDao) CreateLikeComment(like *model.CommentLike) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Table("comment_like").
		Create(&like).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("comment").
		Where("comment_id=?", like.CommentId).
		Update("like_count", gorm.Expr("like_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}

func (dao *CommentDao) UnLike(userId string, commentId int64) (err error) {
	tx := dao.DB
	err = tx.Model(&model.CommentLike{}).
		Where("comment_id = ? and user_id = ?", commentId, userId).
		Update("is_liked", false).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("comment").
		Where("comment_id=?", commentId).
		Update("like_count", gorm.Expr("like_count-1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}

func (dao *CommentDao) Like(userId string, commentId int64) (err error) {
	tx := dao.DB
	err = tx.Model(&model.CommentLike{}).
		Where("comment_id = ? and user_id = ?", commentId, userId).
		Update("is_liked", true).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("comment").
		Where("comment_id=?", commentId).
		Update("like_count", gorm.Expr("like_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return
}
