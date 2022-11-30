package service

import (
	"SmallRedBook/dao"
	"SmallRedBook/model"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"context"
	"strconv"
	"time"
)

type CommentService struct {
	NoteId  string `json:"note_id" form:"note_id"`
	Content string `json:"content" form:"content"`
}

func (service *CommentService) AddComment(ctx context.Context, userId string) tool.Response {
	noteId, _ := strconv.Atoi(service.NoteId)
	commentOrNotDao := dao.NewCommentDao(ctx)
	count, err := commentOrNotDao.CommentOrNot(userId, int64(noteId))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count != 0 {
		return e.ThrowError(e.HasComment)
	}

	getUserInfoDao := dao.NewUserDaoByDb(commentOrNotDao.DB)
	user, err := getUserInfoDao.GetUserInfoByUserId(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	comment := &model.Comment{
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		NoteId:   uint(noteId),
		UserId:   userId,
		UserName: user.UserName,
		Content:  service.Content,
		Avatar:   user.Avatar,
	}

	addCommentDao := dao.NewCommentDaoByDb(getUserInfoDao.DB)
	err = addCommentDao.AddComment(comment)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *CommentService) DeleteComment(ctx context.Context, userId string) tool.Response {
	noteId, _ := strconv.Atoi(service.NoteId)
	commentOrNotDao := dao.NewCommentDao(ctx)
	count, err := commentOrNotDao.CommentOrNot(userId, int64(noteId))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		return e.ThrowError(e.HasNotComment)
	}
	deleteCommentDao := dao.NewCommentDaoByDb(commentOrNotDao.DB)

	err = deleteCommentDao.DeleteComment(userId, int64(noteId))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}
