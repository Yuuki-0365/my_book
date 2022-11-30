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

type FavoriteService struct {
	NoteId string `json:"note_id" form:"note_id"`
}

func (service *FavoriteService) FavoriteNote(ctx context.Context, userId string) tool.Response {
	favoriteNoteDao := dao.NewFavoriteDao(ctx)
	id, _ := strconv.Atoi(service.NoteId)
	count, err := favoriteNoteDao.FavoriteOrNot(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count != 0 {
		return e.ThrowError(e.HasFavorited)
	}

	favorite := &model.Favorite{
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		UserId:   userId,
		NoteId:   uint(id),
	}
	err = favoriteNoteDao.FavoriteNote(favorite)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *FavoriteService) UnFavoriteNote(ctx context.Context, userId string) tool.Response {
	unFavoriteNoteDao := dao.NewFavoriteDao(ctx)
	id, _ := strconv.Atoi(service.NoteId)
	count, err := unFavoriteNoteDao.FavoriteOrNot(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		return e.ThrowError(e.HasNotFavorited)
	}

	err = unFavoriteNoteDao.UnFavoriteNote(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}
