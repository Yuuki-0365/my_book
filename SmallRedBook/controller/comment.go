package controller

import (
	"SmallRedBook/service"
	"SmallRedBook/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCommentToNote(c *gin.Context) {
	var addCommentToNoteService service.CommentService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addCommentToNoteService); err == nil {
		res := addCommentToNoteService.AddCommentToNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func DeleteCommentToNote(c *gin.Context) {
	var deleteCommentToNoteService service.CommentService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCommentToNoteService); err == nil {
		res := deleteCommentToNoteService.DeleteCommentToNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func AddCommentToComment(c *gin.Context) {
	var addCommentToCommentService service.CommentService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addCommentToCommentService); err == nil {
		res := addCommentToCommentService.AddCommentToComment(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func DeleteCommentToComment(c *gin.Context) {
	var deleteCommentToCommentService service.CommentService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCommentToCommentService); err == nil {
		res := deleteCommentToCommentService.DeleteCommentToComment(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func LikeComment(c *gin.Context) {
	var likeCommentService service.CommentService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&likeCommentService); err == nil {
		res := likeCommentService.LikeComment(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
