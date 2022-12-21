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
	Like, count, err := likeNoteDao.LikeOrNot(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		like := &model.Like{
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
			UserId:   userId,
			NoteId:   uint(id),
			IsLiked:  true,
		}
		err = likeNoteDao.CreateLikeNote(like)
		if err != nil {
			return e.ThrowError(e.ErrorDataBase)
		}
		return tool.Response{
			Status: e.Success,
			Msg:    e.GetMsg(e.HasNotLiked),
		}
	}
	if count == 1 {
		if Like.IsLiked == true {
			err = likeNoteDao.UnLike(userId, int64(id))
			if err != nil {
				return e.ThrowError(e.ErrorDataBase)
			}
			return tool.Response{
				Status: e.Success,
				Msg:    e.GetMsg(e.HasLiked),
			}
		} else {
			err = likeNoteDao.Like(userId, int64(id))
			if err != nil {
				return e.ThrowError(e.ErrorDataBase)
			}
			return tool.Response{
				Status: e.Success,
				Msg:    e.GetMsg(e.HasNotLiked),
			}
		}
	}
	return e.ThrowError(e.Error)
}
