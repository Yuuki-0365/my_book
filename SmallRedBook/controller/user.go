package controller

import (
	"SmallRedBook/service"
	"SmallRedBook/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterCode(c *gin.Context) {
	var userRegisterCodeService service.UserService
	if err := c.ShouldBind(&userRegisterCodeService); err == nil {
		res := userRegisterCodeService.GetRegisterCode(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserLoginCode(c *gin.Context) {
	var userLoginCodeService service.UserService
	if err := c.ShouldBind(&userLoginCodeService); err == nil {
		res := userLoginCodeService.GetLoginCode(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserLoginById(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.LoginById(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserLoginByEmail(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.LoginByEmail(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserFollowersCount(c *gin.Context) {
	var userFollowersCountService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userFollowersCountService); err == nil {
		res := userFollowersCountService.GetFollowersCount(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserFansCount(c *gin.Context) {
	var userFansCountService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userFansCountService); err == nil {
		res := userFansCountService.GetFansCount(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserFollowers(c *gin.Context) {
	var userFollowersService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userFollowersService); err == nil {
		res := userFollowersService.GetFollowers(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

}

func UserFans(c *gin.Context) {
	var userFansService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userFansService); err == nil {
		res := userFansService.GetFans(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

}

func UserFollowed(c *gin.Context) {
	var userFollowedService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userFollowedService); err == nil {
		res := userFollowedService.UserFollowed(c.Request.Context(), claims.UserId, userFollowedService.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

}

func UserFollow(c *gin.Context) {
	var userFollowService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userFollowService); err == nil {
		res := userFollowService.UserFollow(c.Request.Context(), claims.UserId, userFollowService.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func FollowTogether(c *gin.Context) {
	var followTogetherService service.UserService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&followTogetherService); err == nil {
		res := followTogetherService.FollowTogether(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserUpdateCode(c *gin.Context) {
	var userUpdateCodeService service.UserService

	if err := c.ShouldBind(&userUpdateCodeService); err == nil {
		res := userUpdateCodeService.GetUpdateCode(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserUpdateEmailByEmail(c *gin.Context) {
	var userUpdateEmailByEmailService service.UserService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateEmailByEmailService); err == nil {
		res := userUpdateEmailByEmailService.UpdateEmailByEmail(c.Request.Context(), claims.UserId, c.PostForm("code_1"), c.PostForm("code_2"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserUpdateEmailByPassword(c *gin.Context) {
	var userUpdateEmailByPasswordService service.UserService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateEmailByPasswordService); err == nil {
		res := userUpdateEmailByPasswordService.UpdateEmailByPassword(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserUpdatePasswordByPassword(c *gin.Context) {
	var userUpdatePasswordByPasswordService service.UserService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdatePasswordByPasswordService); err == nil {
		res := userUpdatePasswordByPasswordService.UpdatePasswordByPassword(c.Request.Context(), claims.UserId, c.PostForm("new_password"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserUpdatePasswordByEmail(c *gin.Context) {
	var userUpdatePasswordService service.UserService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdatePasswordService); err == nil {
		res := userUpdatePasswordService.UpdatePasswordByEmail(c.Request.Context(), claims.UserId, c.PostForm("new_password"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UserUpdateInfo(c *gin.Context) {
	var userUpdateInfoService service.UserService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateInfoService); err == nil {
		file, _, _ := c.Request.FormFile("file")
		if file != nil {
			res := userUpdateInfoService.UpdateInfoIncludeAvatar(c.Request.Context(), claims.UserId, file)
			c.JSON(http.StatusOK, res)
		} else {
			res := userUpdateInfoService.UpdateInfo(c.Request.Context(), claims.UserId)
			c.JSON(http.StatusOK, res)
		}
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowUserInfoInUpdate(c *gin.Context) {
	var showUserInfoInUpdateService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showUserInfoInUpdateService); err == nil {
		res := showUserInfoInUpdateService.ShowUserInfoInUpdate(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowUserInfoAll(c *gin.Context) {
	var showUserInfoAllService service.UserService
	if err := c.ShouldBind(&showUserInfoAllService); err == nil {
		res := showUserInfoAllService.ShowUserInfoAll(c.Request.Context(), showUserInfoAllService.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowOwnUserInfoAll(c *gin.Context) {
	var showOwnUserInfoAllService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showOwnUserInfoAllService); err == nil {
		res := showOwnUserInfoAllService.ShowOwnUserInfoAll(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SearchUser(c *gin.Context) {
	var searchUserService service.UserService
	if err := c.ShouldBind(&searchUserService); err == nil {
		res := searchUserService.SearchUser(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func GetUserUpdateInfo(c *gin.Context) {
	var getUserUpdateInfoService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getUserUpdateInfoService); err == nil {
		res := getUserUpdateInfoService.GetUpdateInfo(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func DeleteUser(c *gin.Context) {
	var deleteUserService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteUserService); err == nil {
		res := deleteUserService.DeleteUser(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func AdminShowInfo(c *gin.Context) {
	var adminShowInfo service.UserService
	if err := c.ShouldBind(&adminShowInfo); err == nil {
		res := adminShowInfo.AdminShowInfo(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func AdminUpdateInfo(c *gin.Context) {
	var adminUpdateInfoService service.UserService

	if err := c.ShouldBind(&adminUpdateInfoService); err == nil {
		res := adminUpdateInfoService.UpdateInfo(c.Request.Context(), adminUpdateInfoService.UserId)
		fmt.Println(adminUpdateInfoService.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func AdminDeleteUser(c *gin.Context) {
	var adminDeleteInfoService service.UserService
	if err := c.ShouldBind(&adminDeleteInfoService); err == nil {
		res := adminDeleteInfoService.DeleteUser(c.Request.Context(), adminDeleteInfoService.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func AdminAddUser(c *gin.Context) {
	var adminAddUserService service.UserService
	if err := c.ShouldBind(&adminAddUserService); err == nil {
		res := adminAddUserService.AddUser(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func GetFollowUserNotes(c *gin.Context) {
	var getFollowUserNotesService service.UserService
	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getFollowUserNotesService); err == nil {
		res := getFollowUserNotesService.GetFollowUserNotes(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
