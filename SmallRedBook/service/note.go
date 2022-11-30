package service

import (
	"SmallRedBook/conf"
	"SmallRedBook/dao"
	"SmallRedBook/model"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"context"
	"mime/multipart"
	"strconv"
	"sync"
	"time"
)

type NoteService struct {
	Content    string `json:"content" form:"content"`
	Title      string `json:"title" form:"title"`
	PageSize   int    `json:"pageSize" form:"page_size"`
	PageNumber int    `json:"page_number" form:"page_number"`
	NoteId     string `json:"note_id" form:"note_id"`
}

func (service *NoteService) PublishNote(ctx context.Context, userId string, files []*multipart.FileHeader) tool.Response {
	publishNoteDao := dao.NewNoteDao(ctx)

	count, err := publishNoteDao.Count()
	if err != nil {
		e.ThrowError(e.ErrorDataBase)
	}

	if err != nil {
		e.ThrowError(e.UploadNoteFileToLocalStaticError)
	}
	noteId := strconv.Itoa(int(count + 1))

	note := &model.Note{
		UserId:        userId,
		Title:         service.Title,
		Content:       service.Content,
		FilePath:      conf.NotePath + "user" + userId + "/" + "note" + noteId + "/",
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		LikeCount:     0,
		FavoriteCount: 0,
		CommentCount:  0,
		Id:            count + 1,
	}
	err = publishNoteDao.CreateNote(note)
	if err != nil {
		e.ThrowError(e.ErrorDataBase)
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		tmp, _ := file.Open()
		_, err = UploadNoteFileToLocalStatic(tmp, userId, noteId, num)
		if err != nil {
			e.ThrowError(e.UploadNoteFileToLocalStaticError)
		}
		wg.Done()
	}
	wg.Wait()
	return e.ThrowSuccess()
}

func (service *NoteService) GetNotesInfoLess(ctx context.Context) tool.Response {
	getNoteInfoLessDao := dao.NewNoteDao(ctx)
	notesInfoLess, err := getNoteInfoLessDao.GetNotesInfoLess(service.PageNumber, service.PageSize)

	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	for _, item := range notesInfoLess {
		var v1 interface{}
		v1 = conf.Host + conf.HttpPort + conf.AvatarPath + "user" + item["user_id"].(string) + "/" + item["user_name"].(string) + ".jpg"
		item["avatar"] = v1
		var v2 interface{} = conf.Host + conf.HttpPort + item["file_path"].(string) + "0.jpg"
		item["file_path"] = v2
	}
	return tool.Response{
		Status: e.Success,
		Msg:    e.GetMsg(e.Success),
		Data: tool.DataList{
			Item:  notesInfoLess,
			Total: uint(len(notesInfoLess)),
		},
	}
}

func (service *NoteService) DeleteNote(ctx context.Context, userId string) tool.Response {
	deleteNoteDao := dao.NewNoteDao(ctx)
	id, _ := strconv.Atoi(service.NoteId)

	err := deleteNoteDao.DeleteNote(userId, int64(id))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *NoteService) SearchNote(ctx context.Context) tool.Response {
	searchNoteDao := dao.NewNoteDao(ctx)
	notes, err := searchNoteDao.SearchNote(service.PageNumber, service.PageSize, service.Title)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	for _, item := range notes {
		var v1 interface{} = conf.Host + conf.HttpPort + conf.AvatarPath + item["avatar"].(string)
		item["avatar"] = v1
		var v2 interface{} = conf.Host + conf.HttpPort + item["file_path"].(string) + "0.jpg"
		item["file_path"] = v2
	}
	return tool.Response{
		Status: e.Success,
		Msg:    e.GetMsg(e.Success),
		Data: tool.DataList{
			Item:  notes,
			Total: uint(len(notes)),
		},
	}
}

func (service *NoteService) GetNotesInfoMore(ctx context.Context) tool.Response {
	getNotesInfoMoreDao := dao.NewNoteDao(ctx)
	noteId, _ := strconv.Atoi(service.NoteId)
	noteInfo, err := getNotesInfoMoreDao.GetNotesInfoMore(int64(noteId)) // get the note
	for _, item := range noteInfo {
		files, _ := GetAllFile("." + item["file_path"].(string))
		var v1 interface{} = files
		item["file_path"] = v1
	}
	getCommentInfoDao := dao.NewCommentDaoByDb(getNotesInfoMoreDao.DB)
	commentInfo, err := getCommentInfoDao.GetCommentInfo(int64(noteId))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	for _, item := range commentInfo {
		var v1 interface{}
		v1 = conf.Host + conf.HttpPort + conf.AvatarPath + "user" + item["user_id"].(string) + "/" + item["user_name"].(string) + ".jpg"
		item["avatar"] = v1
	}

	return tool.BuildNoteInfoResponse(noteInfo, commentInfo, uint(len(commentInfo)))
}
