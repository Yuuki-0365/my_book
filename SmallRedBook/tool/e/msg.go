package e

var MsgFlags = map[int]string{
	Success:                        "ok",
	Error:                          "fail",
	ErrorAuthCheckTokenTimeout:     "Token已过期，请重新登录",
	ErrorAuthCheckTokenFail:        "验证Token出错",
	ExistEmail:                     "邮箱已存在",
	ErrorDataBase:                  "数据库查询出错",
	ErrorSendEmail:                 "发送邮箱出错",
	PasswordIsShort:                "密码太短了，不足八位",
	ErrorRedis:                     "系统错误，请重新获取验证码",
	VcodeNotMatch:                  "验证码错误",
	LoginByIdError:                 "用户ID或密码错误，请重新输入",
	NotExistEmail:                  "邮箱未注册",
	ErrorAuthToken:                 "签发token失败",
	Follow:                         "关注了",
	NotFollow:                      "没有关注",
	FollowTogether:                 "已经互相关注了",
	NotFollowTogether:              "没有互相关注",
	ErrorOldPassword:               "原密码错误",
	OldVcodeNotMatch:               "旧邮箱验证码错误",
	NewVcodeNotMatch:               "新邮箱验证码错误",
	UploadAvatarToLocalStaticError: "头像保存到本地错误",
	MatchNewOldEmail:               "新旧邮箱相同，无须更改",
	MatchNewOldPassword:            "新旧密码相同，无须更改",
	ExistNewEmail:                  "新邮箱已经注册过",
	HasLiked:                       "取消点赞了",
	HasNotLiked:                    "点赞成功",
	HasFavorited:                   "取消收藏了",
	HasNotFavorited:                "收藏成功",
	HasComment:                     "已经评论过了",
	HasNotComment:                  "还没有评论过",
	HasFollowed:                    "取消关注了",
	HasNotFollowed:                 "关注成功",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
