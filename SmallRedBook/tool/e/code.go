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
	NotFollowTogether
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
	HasFollowed
	HasNotFollowed
	HasComment
	HasNotComment
	NotFollow
	Follow
)
