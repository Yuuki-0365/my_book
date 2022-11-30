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

type LikeService struct {
	NoteId string `json:"note_id" form:"note_id"`
}

func (service *LikeService) LikeNote(ctx context.Context, userId string) tool.Response {
	likeNoteDao := dao.NewLikeDao(ctx)
	id, _ := strconv.Atoi(service.NoteId)
	count, err := likeNoteDao.LikeOrNot(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count != 0 {
		return e.ThrowError(e.HasLiked)
	}

	like := &model.Like{
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		UserId:   userId,
		NoteId:   uint(id),
	}
	err = likeNoteDao.LikeNote(like)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *LikeService) UnLikeNote(ctx context.Context, userId string) tool.Response {
	unLikeNoteDao := dao.NewLikeDao(ctx)
	id, _ := strconv.Atoi(service.NoteId)
	count, err := unLikeNoteDao.LikeOrNot(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		return e.ThrowError(e.HasNotLiked)
	}

	err = unLikeNoteDao.UnLikeNote(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}
