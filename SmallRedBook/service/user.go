package service

import (
	"SmallRedBook/cache"
	"SmallRedBook/conf"
	"SmallRedBook/dao"
	"SmallRedBook/model"
	"SmallRedBook/serializer"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"SmallRedBook/tool/snowflake"
	"context"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strconv"
	"time"
)

type UserService struct {
	Password     string `json:"password" form:"password"`
	Email        string `json:"email" form:"email"`
	UserName     string `json:"user_name" form:"user_name"`
	Code         string `json:"code" form:"code"`
	UserId       string `json:"user_id" form:"user_id"`
	Key          string `json:"key" form:"key"`
	Avatar       string `json:"avatar" form:"avatar"`
	Introduction string `json:"introduction" form:"introduction"`
	Sex          string `json:"sex" form:"sex"`
}

func (service *UserService) GetRegisterCode(ctx context.Context) tool.Response {
	// 判断邮箱是否已经注册
	userDao := dao.NewUserDao(ctx)
	exist, err := userDao.ExistEmailOrNot(service.Email)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if exist {
		return e.ThrowError(e.ExistEmail)
	}

	// 邮箱格式
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err := noticeDao.GetNoticeById(1)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	Vcode := tool.GenerateVcode()
	conn := cache.RedisPool.Get()
	defer conn.Close()
	conn.Do("SET", cache.UserRegisterCode+service.Email, Vcode, "EX", 120)

	mailStr := notice.Text
	mailText := mailStr + Vcode
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "SmallRedBook")
	m.SetBody("text/html", mailText)

	// 发送
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		return e.ThrowError(e.ErrorSendEmail)
	}

	return e.ThrowSuccess()
}

func (service *UserService) GetLoginCode(ctx context.Context) tool.Response {
	// 判断邮箱是否已经注册
	userDao := dao.NewUserDao(ctx)
	exist, err := userDao.ExistEmailOrNot(service.Email)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if !exist {
		return e.ThrowError(e.NotExistEmail)
	}

	// 邮箱格式
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err := noticeDao.GetNoticeById(2)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	Vcode := tool.GenerateVcode()
	conn := cache.RedisPool.Get()
	defer conn.Close()
	conn.Do("SET", cache.UserLoginCode+service.Email, Vcode, "EX", 120)
	mailStr := notice.Text
	mailText := mailStr + Vcode
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "SmallRedBook")
	m.SetBody("text/html", mailText)

	// 发送
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		return e.ThrowError(e.ErrorSendEmail)
	}

	return e.ThrowSuccess()
}

func (service *UserService) Register(ctx context.Context, file multipart.File) tool.Response {

	// 判断密码长度
	if len(service.Password) < 8 {
		return e.ThrowError(e.PasswordIsShort)
	}

	// 验证码相关
	conn := cache.RedisPool.Get()
	defer conn.Close()
	Vcode, err := redis.String(conn.Do("GET", cache.UserRegisterCode+service.Email))
	if err != nil {
		return e.ThrowError(e.ErrorRedis)
	}

	if Vcode != service.Code {
		return e.ThrowError(e.VcodeNotMatch)
	}

	// 生成唯一ID
	snowflake.SetMachineId(12)
	userId := strconv.Itoa(int(snowflake.GetId()))
	path, err := UploadAvatarToLocalStatic(file, userId, service.UserName)
	if err != nil {
		return e.ThrowError(e.UploadAvatarToLocalStaticError)
	}

	// 数据初始化
	user := &model.User{
		UserId:       userId,
		UserName:     service.UserName,
		Email:        service.Email,
		Password:     tool.Encrypt.AesEncoding(service.Password),
		Avatar:       path,
		Introduction: "该用户太懒了，没有写自己的简介",
		Sex:          "男",
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	// 插入数据
	userDao := dao.NewUserDao(ctx)
	err = userDao.CreateUser(user)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	return tool.BuildListResponse(serializer.BuildUser(user), 1)
}

func (service *UserService) LoginById(ctx context.Context) tool.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	// 验证用户以及密码，并且查询用户
	user, err := userDao.GetUserByUserIdAndPassword(service.UserId, tool.Encrypt.AesEncoding(service.Password))
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if user == nil {
		return e.ThrowError(e.LoginByIdError)
	}

	// 签发token
	token, err := tool.GenerateToken(user.UserId, 0)
	if err != nil {
		return e.ThrowError(e.ErrorAuthToken)
	}

	return tool.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: tool.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}

func (service *UserService) LoginByEmail(ctx context.Context) tool.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	// 验证是否有email
	count, user, err := userDao.GetUserByEmail(service.Email)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		return e.ThrowError(e.NotExistEmail)
	}

	// 判断验证码
	conn := cache.RedisPool.Get()
	defer conn.Close()
	Vcode, err := redis.String(conn.Do("GET", cache.UserLoginCode+service.Email))
	if err != nil {
		return e.ThrowError(e.ErrorRedis)
	}
	if Vcode != service.Code {
		return e.ThrowError(e.VcodeNotMatch)
	}

	// 签发token
	token, err := tool.GenerateToken(user.UserId, 0)
	if err != nil {
		return e.ThrowError(e.ErrorAuthToken)
	}

	return tool.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: tool.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}

func (service *UserService) GetFollowersCount(ctx context.Context, userId string) tool.Response {
	code := e.Success
	getFollowersCountDao := dao.NewUserDao(ctx)

	count, err := getFollowersCountDao.GetFollowedCount(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	return tool.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   count,
	}
}

func (service *UserService) GetFansCount(ctx context.Context, userId string) tool.Response {
	code := e.Success

	getFansCountDao := dao.NewUserDao(ctx)
	count, err := getFansCountDao.GetFansCount(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	return tool.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   count,
	}
}

func (service *UserService) GetFollowers(ctx context.Context, userId string) tool.Response {
	getFollowersDao := dao.NewUserDao(ctx)
	count, users, err := getFollowersDao.GetFollowersById(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	return tool.BuildListResponse(serializer.BuildFollowers(users), uint(count))
}

func (service *UserService) GetFans(ctx context.Context, userId string) tool.Response {
	getFansDao := dao.NewUserDao(ctx)
	count, users, err := getFansDao.GetFansById(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	return tool.BuildListResponse(serializer.BuildFans(users), uint(count))
}

func (service *UserService) UserFollowed(ctx context.Context, userId string, followId string) tool.Response {
	code := e.Success

	userFollowedDao := dao.NewUserDao(ctx)
	count, err := userFollowedDao.UserFollowed(userId, followId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count != 0 {
		code = e.HasFollowed
		return tool.Response{
			Status: code,
			Msg:    userId + " " + e.GetMsg(code) + " " + followId,
		}
	} else {
		code = e.NotFollow
		return tool.Response{
			Status: code,
			Msg:    userId + " " + e.GetMsg(code) + " " + followId,
		}
	}
}

func (service *UserService) UserFollow(ctx context.Context, userId string, followId string) tool.Response {
	code := e.Success

	userFollowDao := dao.NewUserDao(ctx)
	// 首先判断是否已经有了
	count, err := userFollowDao.UserFollowed(userId, followId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	// 有了
	if count != 0 {
		code = e.HasFollowed
		return tool.Response{
			Status: code,
			Msg:    userId + " " + e.GetMsg(code) + " " + followId,
		}
	}

	// 没有
	err = userFollowDao.UserFollow(userId, followId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) UserUnFollow(ctx context.Context, userId string, unFollowId string) tool.Response {
	code := e.Success

	userFollowDao := dao.NewUserDao(ctx)

	count, err := userFollowDao.UserFollowed(userId, unFollowId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		code = e.NotFollow
		return tool.Response{
			Status: code,
			Msg:    userId + " " + e.GetMsg(code) + " " + unFollowId,
		}
	}

	// 取关首先要关注了
	err = userFollowDao.UserUnFollow(userId, unFollowId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	return e.ThrowSuccess()
}

func (service *UserService) UserFollowTogether(ctx context.Context, userId string, userFollowId string) tool.Response {
	userFollowTogetherDao := dao.NewUserDao(ctx)
	count, err := userFollowTogetherDao.UserFollowed(userId, userFollowId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		return e.ThrowError(e.NotFollow)
	}

	count, err = userFollowTogetherDao.UserFollowed(userFollowId, userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if count == 0 {
		return e.ThrowError(e.NotFollow)
	}

	return e.ThrowSuccess()
}

func (service *UserService) GetUpdateCode(ctx context.Context) tool.Response {
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err := noticeDao.GetNoticeById(3)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	Vcode := tool.GenerateVcode()
	conn := cache.RedisPool.Get()
	conn.Do("SET", cache.UserUpdateCode+service.Email, Vcode, "EX", 120)
	defer conn.Close()

	mailStr := notice.Text
	mailText := mailStr + Vcode
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "SmallRedBook")
	m.SetBody("text/html", mailText)

	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		return e.ThrowError(e.ErrorSendEmail)
	}

	return e.ThrowSuccess()
}

func (service *UserService) UpdateEmailByEmail(ctx context.Context, userId, oldCode, newCode string) tool.Response {
	updateEmailByEmailDao := dao.NewUserDao(ctx)
	oldEmail, err := updateEmailByEmailDao.GetEmailById(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if service.Email == oldEmail {
		return e.ThrowError(e.MatchNewOldEmail)
	}

	exist, err := updateEmailByEmailDao.ExistEmailOrNot(service.Email)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if exist {
		return e.ThrowError(e.ExistNewEmail)
	}

	conn := cache.RedisPool.Get()
	defer conn.Close()
	Vcode1, err := redis.String(conn.Do("GET", cache.UserUpdateCode+oldEmail))
	if Vcode1 != oldCode {
		return e.ThrowError(e.OldVcodeNotMatch)
	}
	if err != nil {
		return e.ThrowError(e.ErrorRedis)
	}

	Vcode2, err := redis.String(conn.Do("GET", cache.UserUpdateCode+service.Email))
	if Vcode2 != newCode {
		return e.ThrowError(e.NewVcodeNotMatch)
	}
	if err != nil {
		return e.ThrowError(e.ErrorRedis)
	}

	updateTime := time.Now().Format("2006-01-02 15:04:05")
	err = updateEmailByEmailDao.UpdateEmail(userId, service.Email, updateTime)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) UpdateEmailByPassword(ctx context.Context, userId string) tool.Response {
	oldPassword := service.Password
	newEmail := service.Email

	updateEmailByIdDao := dao.NewUserDao(ctx)
	oldEmail, err := updateEmailByIdDao.GetEmailById(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if newEmail == oldEmail {
		return e.ThrowError(e.MatchNewOldEmail)
	}

	oldPassword = tool.Encrypt.AesEncoding(oldPassword)
	_, err = updateEmailByIdDao.GetUserByUserIdAndPassword(userId, oldPassword)
	if err != nil {
		return e.ThrowError(e.ErrorOldPassword)
	}

	conn := cache.RedisPool.Get()
	defer conn.Close()
	Vcode, err := redis.String(conn.Do("GET", cache.UserUpdateCode+service.Email))
	if err != nil {
		return e.ThrowError(e.ErrorRedis)
	}

	if Vcode != service.Code {
		return e.ThrowError(e.VcodeNotMatch)
	}

	updateTime := time.Now().Format("2006-01-02 15:04:05")
	err = updateEmailByIdDao.UpdateEmail(userId, newEmail, updateTime)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) UpdatePasswordByPassword(ctx context.Context, userId string, newPassword string) tool.Response {
	if service.Password == newPassword {
		return e.ThrowError(e.MatchNewOldPassword)
	}

	// service.Password 旧密码
	password := tool.Encrypt.AesEncoding(service.Password)
	updatePasswordDao := dao.NewUserDao(ctx)
	// 判断用户id和密码是否对应的上
	_, err := updatePasswordDao.GetUserByUserIdAndPassword(userId, password)
	if err != nil {
		return e.ThrowError(e.ErrorOldPassword)
	}

	// 对应上了就更新
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	newPassword = tool.Encrypt.AesEncoding(newPassword)
	err = updatePasswordDao.UpdatePassword(userId, newPassword, updateTime)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) UpdatePasswordByEmail(ctx context.Context, userId string, newPassword string) tool.Response {
	updatePasswordDao := dao.NewUserDao(ctx)
	email, err := updatePasswordDao.GetEmailById(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	conn := cache.RedisPool.Get()
	defer conn.Close()
	Vcode, err := redis.String(conn.Do("GET", cache.UserUpdateCode+email))
	if err != nil {
		return e.ThrowError(e.ErrorRedis)
	}
	if Vcode != service.Code {
		return e.ThrowError(e.VcodeNotMatch)
	}

	updateTime := time.Now().Format("2006-01-02 15:04:05")
	newPassword = tool.Encrypt.AesEncoding(newPassword)
	err = updatePasswordDao.UpdatePassword(userId, newPassword, updateTime)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) UpdateInfo(ctx context.Context, userId string) tool.Response {
	// userName, avatar, introduction, sex
	updateInfoDao := dao.NewUserDao(ctx)
	user, err := updateInfoDao.GetUserInfoByUserId(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if service.UserName != "" {
		user.UserName = service.UserName
	}
	if service.Introduction != "" {
		user.Introduction = service.Introduction
	}
	if service.Sex != "" {
		user.Sex = service.Sex
	}
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = updateInfoDao.UpdateInfo(userId, user)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) UpdateInfoIncludeAvatar(ctx context.Context, userId string, file multipart.File) tool.Response {
	// userName, avatar, introduction, sex
	updateInfoDao := dao.NewUserDao(ctx)
	user, err := updateInfoDao.GetUserInfoByUserId(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	if service.UserName != "" {
		user.UserName = service.UserName
	}
	if service.Introduction != "" {
		user.Introduction = service.Introduction
	}
	if service.Sex != "" {
		user.Sex = service.Sex
	}
	path, err := UploadAvatarToLocalStatic(file, user.UserId, user.UserName)
	if err != nil {
		return e.ThrowError(e.UploadAvatarToLocalStaticError)
	}
	user.Avatar = path
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = updateInfoDao.UpdateInfo(userId, user)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return e.ThrowSuccess()
}

func (service *UserService) ShowUserInfoInUpdate(ctx context.Context, userId string) tool.Response {
	showUserInfoInUpdateDao := dao.NewUserDao(ctx)
	user, err := showUserInfoInUpdateDao.GetUserInfoInUpdate(userId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	return tool.BuildListResponse(serializer.BuildUserInfoInUpdate(user), 1)
}

func (service *UserService) ShowUserInfoAll(ctx context.Context) tool.Response {
	showUserInfoAllDao := dao.NewUserDao(ctx)
	userInfo, err := showUserInfoAllDao.ShowUserInfoAll(service.UserId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}
	userInfo["avatar"] = conf.Host + conf.HttpPort + conf.AvatarPath + userInfo["avatar"].(string)

	getNoteInfoLessDao := dao.NewNoteDaoByDb(showUserInfoAllDao.DB)
	noteInfo, err := getNoteInfoLessDao.GetNotesByUserId(service.UserId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	getLikeNoteInfoDao := dao.NewNoteDaoByDb(getNoteInfoLessDao.DB)
	likeNotesInfo, err := getLikeNoteInfoDao.GetLikeNotes(service.UserId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	getFavoriteNoteInfoDao := dao.NewNoteDaoByDb(getLikeNoteInfoDao.DB)
	favoriteNotesInfo, err := getFavoriteNoteInfoDao.GetFavoriteNotes(service.UserId)
	if err != nil {
		return e.ThrowError(e.ErrorDataBase)
	}

	for _, item := range noteInfo {
		files, _ := GetAllFile("." + item["file_path"].(string))
		var v1 interface{} = files
		item["file_path"] = v1
	}

	for _, item := range likeNotesInfo {
		files, _ := GetAllFile("." + item["file_path"].(string))
		var v1 interface{} = files
		item["file_path"] = v1
	}

	for _, item := range favoriteNotesInfo {
		files, _ := GetAllFile("." + item["file_path"].(string))
		var v1 interface{} = files
		item["file_path"] = v1
	}
	return tool.BuildUserInfoAll(userInfo, noteInfo, likeNotesInfo, favoriteNotesInfo)
}

func (service *UserService) SearchUser(ctx context.Context) tool.Response {
	searchUserDao := dao.NewUserDao(ctx)
	users, count, err := searchUserDao.SearchUser(service.UserName)
	if err != nil {
		e.ThrowError(e.ErrorDataBase)
	}
	return tool.BuildListResponse(serializer.BuildSearchUsers(users), uint(count))
}
