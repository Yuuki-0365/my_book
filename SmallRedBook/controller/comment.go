package controller

import (
	"SmallRedBook/service"
	"SmallRedBook/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddComment(c *gin.Context) {
	var addCommentService service.CommentService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addCommentService); err == nil {
		res := addCommentService.AddComment(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func DeleteComment(c *gin.Context) {
	var deleteCommentService service.CommentService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCommentService); err == nil {
		res := deleteCommentService.DeleteComment(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
