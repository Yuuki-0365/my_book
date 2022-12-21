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
	Favorite, count, err := favoriteNoteDao.FavoriteOrNot(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		favorite := &model.Favorite{
			CreateAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt:   time.Now().Format("2006-01-02 15:04:05"),
			UserId:     userId,
			NoteId:     uint(id),
			IsFavorite: true,
		}
		err = favoriteNoteDao.CreateFavoriteNote(favorite)
		if err != nil {
			return e.ThrowError(e.ErrorDataBase)
		}
		return tool.Response{
			Status: e.Success,
			Msg:    e.GetMsg(e.HasNotFavorited),
		}
	}
	if count == 1 {
		if Favorite.IsFavorite == true {
			err = favoriteNoteDao.UnFavorite(userId, int64(id))
			if err != nil {
				return e.ThrowError(e.ErrorDataBase)
			}
			return tool.Response{
				Status: e.Success,
				Msg:    e.GetMsg(e.HasFavorited),
			}
		} else {
			err = favoriteNoteDao.Favorite(userId, int64(id))
			if err != nil {
				return e.ThrowError(e.ErrorDataBase)
			}
			return tool.Response{
				Status: e.Success,
				Msg:    e.GetMsg(e.HasNotFavorited),
			}
		}
	}
	return e.ThrowError(e.Error)
}
