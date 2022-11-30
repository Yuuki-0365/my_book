package router

import (
	"SmallRedBook/controller"
	"SmallRedBook/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api")
	{
		userNotAuthed := v1.Group("/user")
		{
			register := userNotAuthed.Group("/register")
			{
				register.POST("/code", controller.UserRegisterCode)
				register.POST("", controller.UserRegister)
			}
			login := userNotAuthed.Group("/login")
			{
				login.POST("/userId", controller.UserLoginById)
				login.POST("/email", controller.UserLoginByEmail)
				login.POST("/code", controller.UserLoginCode)
			}
			info := userNotAuthed.Group("/info")
			{
				info.GET("/all", controller.ShowUserInfoAll)
			}
		}
		search := v1.Group("/search")
		{
			search.GET("/user", controller.SearchUser)
			search.GET("/note", controller.SearchNote)
		}

		note := v1.Group("/note")
		{
			note.GET("/info/less", controller.GetNoteInfoLess)
			note.GET("/info/more", controller.GetNoteInfoMore)
		}

		authed := v1.Group("/authed")
		authed.Use(middleware.JWT())
		{
			user := authed.Group("/user")
			{
				count := user.Group("/count")
				{
					count.GET("/follow", controller.UserFollowersCount)
					count.GET("/fan", controller.UserFansCount)
				}
				info := user.Group("/info")
				{
					info.GET("/follower", controller.UserFollowers)
					info.GET("/fan", controller.UserFans)
					info.GET("/update", controller.ShowUserInfoInUpdate)

				}
				follow := user.Group("")
				{
					follow.GET("/followed", controller.UserFollowed)
					follow.POST("/follow", controller.UserFollow)
					follow.DELETE("/unfollow", controller.UserUnFollow)
					follow.GET("/follow/together", controller.FollowTogether)
				}
				update := user.Group("/update")
				{
					update.GET("/code", controller.UserUpdateCode)
					update.POST("/email/email", controller.UserUpdateEmailByEmail)
					update.POST("/email/password", controller.UserUpdateEmailByPassword)
					update.POST("/password/password", controller.UserUpdatePasswordByPassword)
					update.POST("/password/email", controller.UserUpdatePasswordByEmail)
					update.POST("/info", controller.UserUpdateInfo)
				}
			}
			note := authed.Group("/note")
			{
				note.POST("/publish", controller.PublishNote)
				note.POST("/like", controller.LikeNote)
				note.DELETE("/unlike", controller.UnLikeNote)
				note.POST("/favorite", controller.FavoriteNote)
				note.DELETE("/unfavorite", controller.UnFavoriteNote)
				note.DELETE("/delete", controller.DeleteNote)
				note.POST("/comment/add", controller.AddComment)
				note.DELETE("/comment/delete", controller.DeleteComment)
			}
			//comment := authed.Group("/comment")
			//
			//{
			//	comment.POST("/like", controller.LikeCommment)
			//	comment.POST("/unlike", controller.UnLikeComment)
			//}
		}
	}
	return r
}
