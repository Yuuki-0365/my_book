package service

import (
	"SmallRedBook/conf"
	"SmallRedBook/dao"
	"SmallRedBook/model"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"context"
	"time"
)

type CommentService struct {
	ParentId  int64  `json:"parent_id" form:"parent_id"`
	Content   string `json:"content" form:"content"`
	CommentId int64  `json:"comment_id" form:"comment_id"`
	NoteId    int64  `json:"note_id" form:"note_id"`
}

func (service *CommentService) AddCommentToNote(ctx context.Context, userId string) tool.Response {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserInfoByUserId(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	addCommentToNoteDao := dao.NewCommentDaoByDb(userDao.DB)
	count, err := addCommentToNoteDao.Count()
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	comment := &model.Comment{
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UserId:       userId,
		UserName:     user.UserName,
		Content:      service.Content,
		Avatar:       user.Avatar,
		LikeCount:    0,
		CommentID:    count + 1,
		ParentId:     service.ParentId,
		NoteId:       service.ParentId,
		Type:         0,
		CommentCount: 0,
	}
	err = addCommentToNoteDao.AddComment(comment)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	comment.Avatar = conf.Host + conf.HttpPort + conf.AvatarPath + comment.Avatar
	return tool.BuildListResponse(comment, 1)
}

func (service *CommentService) DeleteCommentToNote(ctx context.Context, userId string) tool.Response {
	deleteCommentToNoteDao := dao.NewCommentDao(ctx)
	err := deleteCommentToNoteDao.DeleteComment(userId, service.ParentId, service.ParentId, service.CommentId, 0)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *CommentService) AddCommentToComment(ctx context.Context, userId string) tool.Response {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserInfoByUserId(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	addCommentToNoteDao := dao.NewCommentDaoByDb(userDao.DB)
	count, err := addCommentToNoteDao.Count()
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	comment := &model.Comment{
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UserId:       userId,
		UserName:     user.UserName,
		Content:      service.Content,
		Avatar:       user.Avatar,
		LikeCount:    0,
		CommentID:    count + 1,
		ParentId:     service.ParentId,
		NoteId:       service.NoteId,
		Type:         1,
		CommentCount: 0,
	}
	err = addCommentToNoteDao.AddComment(comment)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	comment.Avatar = conf.Host + conf.HttpPort + conf.AvatarPath + comment.Avatar
	return tool.BuildListResponse(comment, 1)
}
func (service *CommentService) DeleteCommentToComment(ctx context.Context, userId string) tool.Response {
	deleteCommentToNoteDao := dao.NewCommentDao(ctx)
	err := deleteCommentToNoteDao.DeleteComment(userId, service.NoteId, service.ParentId, service.CommentId, 1)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *CommentService) LikeComment(ctx context.Context, userId string) tool.Response {
	likeCommentDao := dao.NewCommentDao(ctx)
	Like, count, err := likeCommentDao.LikeOrNot(userId, service.CommentId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		like := &model.CommentLike{
			CreateAt:  time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt:  time.Now().Format("2006-01-02 15:04:05"),
			UserId:    userId,
			CommentId: service.CommentId,
			IsLiked:   true,
		}
		err = likeCommentDao.CreateLikeComment(like)
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
			err = likeCommentDao.UnLike(userId, service.CommentId)
			if err != nil {
				return e.ThrowError(e.ErrorDataBase)
			}
			return tool.Response{
				Status: e.Success,
				Msg:    e.GetMsg(e.HasLiked),
			}
		} else {
			err = likeCommentDao.Like(userId, service.CommentId)
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
