package dao

import (
	"SmallRedBook/model"
	"context"
	"gorm.io/gorm"
)

type NoteDao struct {
	*gorm.DB
}

func NewNoteDao(ctx context.Context) *NoteDao {
	return &NoteDao{NewDBClient(ctx)}
}

func NewNoteDaoByDb(db *gorm.DB) *NoteDao {
	return &NoteDao{db}
}

func (dao *NoteDao) Count() (count int64, err error) {
	err = dao.DB.Model(&model.Note{}).
		Count(&count).Error
	if count == 0 {
		return
	} else {
		err = dao.DB.Table("note").Select("MAX(note_id)").Find(&count).Error
	}
	return
}

func (dao *NoteDao) CreateNote(note *model.Note) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Model(&model.Note{}).Create(&note).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("user").
		Where("user_id=?", note.UserId).
		Update("note_count", gorm.Expr("note_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (dao *NoteDao) GetNotesInfoLess() (notesInfoLess []map[string]interface{}, err error) {
	err = dao.DB.Table("user, note").
		Select("note.note_id, user.user_name, user.avatar, user.user_id, note.title, note.like_count, note.file_path, note.comment_count").
		Where("user.user_id=note.user_id").
		Order("note.created_at desc").
		Find(&notesInfoLess).Error
	return
}

func (dao *NoteDao) DeleteNote(userId string, noteId int64) (err error) {
	tx := dao.DB
	tx.Begin()
	err = tx.Where("user_id=? and note_id=?", userId, noteId).
		Delete(&model.Note{}).
		Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("user").
		Where("user_id=?", userId).
		Update("note_count", gorm.Expr("note_count+1")).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (dao *NoteDao) SearchNote(pageNum int, pageSize int, title string) (notes []map[string]interface{}, err error) {
	err = dao.DB.Table("user, note").
		Select("user.user_name, user.avatar, note.title, note.like_count, note.file_path").
		Where("user.user_id=note.user_id and note.title like ?", "%"+title+"%").
		Offset((pageNum - 1) * (pageSize)).
		Limit(pageSize).
		Find(&notes).Error
	return
}

func (dao *NoteDao) GetNotesInfoMore(noteId int64) (noteInfo []map[string]interface{}, err error) {
	err = dao.DB.Table("note").
		Select("*").
		Where("note.note_id = ?", noteId).
		Find(&noteInfo).Error
	return
}

func (dao *NoteDao) GetNotesByUserId(userId string) (noteInfo []map[string]interface{}, err error) {
	err = dao.DB.Table("user, note").
		Select("user.user_name, user.avatar, note.title, note.like_count, note.file_path, note.comment_count, note.note_id").
		Where("user.user_id = ? and user.user_id=note.user_id", userId).
		Order("note.created_at desc").
		Find(&noteInfo).Error
	return
}

func (dao *NoteDao) GetLikeNotes(userId string) (noteInfo []map[string]interface{}, err error) {
	tmp := dao.DB.Table("like").
		Select("note_id").
		Where("user_id=?", userId)

	err = dao.DB.Table("user, note").
		Select("user.user_name, user.avatar, note.title, note.like_count, note.file_path, note.comment_count, note.note_id").
		Where("note.note_id in(?) and user.user_id = note.user_id", tmp).
		Order("note.created_at desc").
		Find(&noteInfo).Error
	return
}

func (dao *NoteDao) GetFavoriteNotes(userId string) (noteInfo []map[string]interface{}, err error) {
	tmp := dao.DB.Table("favorite").
		Select("note_id").
		Where("user_id=?", userId)

	err = dao.DB.Table("user, note").
		Select("user.user_name, user.avatar, note.title, note.like_count, note.file_path, note.comment_count, note.note_id").
		Where("note.note_id in(?) and user.user_id = note.user_id", tmp).
		Order("note.created_at desc").
		Find(&noteInfo).Error
	return
}
