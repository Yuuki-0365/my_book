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
				info.POST("/all", controller.ShowUserInfoAll)
			}
			admin := userNotAuthed.Group("/admin")
			{
				admin.POST("/info", controller.AdminShowInfo)
				admin.POST("/info/update", controller.AdminUpdateInfo)
				admin.POST("/info/delete", controller.AdminDeleteUser)
				admin.POST("/info/add", controller.AdminAddUser)
			}
		}
		search := v1.Group("/search")
		{
			search.POST("/user", controller.SearchUser)
			search.POST("/note", controller.SearchNote)
		}

		note := v1.Group("/note")
		{
			note.POST("/info/less", controller.GetNoteInfoLess)
			note.POST("/info/more", controller.GetNoteInfoMore)
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
					info.POST("/follower", controller.UserFollowers)
					info.POST("/fan", controller.UserFans)
					info.GET("/update", controller.ShowUserInfoInUpdate)
					info.POST("/all", controller.ShowOwnUserInfoAll)
				}
				follow := user.Group("")
				{
					follow.POST("/followed", controller.UserFollowed)
					follow.POST("/follow", controller.UserFollow)
					follow.POST("/follow/together", controller.FollowTogether)
				}
				update := user.Group("/update")
				{
					update.GET("/code", controller.UserUpdateCode)
					update.POST("/email/email", controller.UserUpdateEmailByEmail)
					update.POST("/email/password", controller.UserUpdateEmailByPassword)
					update.POST("/password/password", controller.UserUpdatePasswordByPassword)
					update.POST("/password/email", controller.UserUpdatePasswordByEmail)
					update.POST("/info", controller.UserUpdateInfo)
					update.POST("/info/all", controller.GetUserUpdateInfo)
				}
				user.POST("/delete", controller.DeleteUser)
				note := user.Group("/note")
				{
					note.POST("/follow", controller.GetFollowUserNotes)
				}
				comment := user.Group("/comment")
				{
					comment.POST("/add/note", controller.AddCommentToNote)
					comment.DELETE("/delete/note", controller.DeleteCommentToNote)
					comment.POST("/add/comment", controller.AddCommentToComment)
					comment.DELETE("/delete/comment", controller.DeleteCommentToComment)
					comment.POST("/like", controller.LikeComment)
				}
			}
			note := authed.Group("/note")
			{
				note.POST("/publish", controller.PublishNote)
				note.POST("/like", controller.LikeNote)
				note.POST("/favorite", controller.FavoriteNote)
				note.DELETE("/delete", controller.DeleteNote)
			}
		}
	}
	return r
}
