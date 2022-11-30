package controller

import (
	"SmallRedBook/service"
	"SmallRedBook/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteNote(c *gin.Context) {
	var favoriteNoteService service.FavoriteService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&favoriteNoteService); err == nil {
		res := favoriteNoteService.FavoriteNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UnFavoriteNote(c *gin.Context) {
	var unFavoriteNoteService service.FavoriteService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&unFavoriteNoteService); err == nil {
		res := unFavoriteNoteService.UnFavoriteNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
