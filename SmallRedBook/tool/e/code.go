package e

const (
	Success = iota
	Error
	ErrorAuthCheckTokenFail
	ErrorAuthCheckTokenTimeout
	ExistEmail
	ErrorDataBase
	ErrorSendEmail
	PasswordIsShort
	ErrorRedis
	VcodeNotExists
	VcodeNotMatch
	LoginByIdError
	NotExistEmail
	ErrorAuthToken
	NotFollow
	HasFollowed
	FollowTogether
	ErrorOldPassword
	OldVcodeNotMatch
	NewVcodeNotMatch
	UploadAvatarToLocalStaticError
	UploadNoteFileToLocalStaticError
	MatchNewOldEmail
	MatchNewOldPassword
	ExistNewEmail
	HasLiked
	HasNotLiked
	HasFavorited
	HasNotFavorited
	HasComment
	HasNotComment
)
