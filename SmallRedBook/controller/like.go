package controller

import (
	"SmallRedBook/service"
	"SmallRedBook/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LikeNote(c *gin.Context) {
	var likeNoteService service.LikeService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&likeNoteService); err == nil {
		res := likeNoteService.LikeNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
