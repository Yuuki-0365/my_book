package controller

import (
	"SmallRedBook/service"
	"SmallRedBook/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishNote(c *gin.Context) {
	// 多个
	form, _ := c.MultipartForm()
	files := form.File["file"]

	var publishNoteService service.NoteService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&publishNoteService); err == nil {
		res := publishNoteService.PublishNote(c.Request.Context(), claims.UserId, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func GetNoteInfoLess(c *gin.Context) {
	var getNoteInfoLessService service.NoteService
	if err := c.ShouldBind(&getNoteInfoLessService); err == nil {
		res := getNoteInfoLessService.GetNotesInfoLess(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func DeleteNote(c *gin.Context) {
	var deleteNoteService service.NoteService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteNoteService); err == nil {
		res := deleteNoteService.DeleteNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SearchNote(c *gin.Context) {
	var searchNote service.NoteService
	if err := c.ShouldBind(&searchNote); err == nil {
		res := searchNote.SearchNote(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetNoteInfoMore(c *gin.Context) {
	var getNoteInfoMore service.NoteService

	if err := c.ShouldBind(&getNoteInfoMore); err == nil {
		res := getNoteInfoMore.GetNotesInfoMore(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
