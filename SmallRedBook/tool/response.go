package tool

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

type NoteInfo struct {
	NoteInfo    interface{} `json:"note_info"`
	CommentInfo interface{} `json:"comment_info"`
	Total       uint        `json:"total"`
}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserInfoAll struct {
	UserInfo          interface{} `json:"user_info"`
	OwnNoteInfo       interface{} `json:"own_note_info"`
	LikeNotesInfo     interface{} `json:"like_notes_info"`
	FavoriteNotesInfo interface{} `json:"favorite_notes_info"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Msg:    "oK",
		Data: DataList{
			Item:  item,
			Total: total,
		},
	}
}

func BuildNoteInfoResponse(noteInfo interface{}, commentInfo interface{}, total uint) Response {
	return Response{
		Status: 200,
		Msg:    "ok",
		Data: NoteInfo{
			NoteInfo:    noteInfo,
			CommentInfo: commentInfo,
			Total:       total,
		},
	}
}

func BuildUserInfoAll(userInfo, noteInfo, likeNoteInfo, favoriteNotesInfo interface{}) Response {
	return Response{
		Status: 200,
		Msg:    "ok",
		Data: UserInfoAll{
			UserInfo:          userInfo,
			OwnNoteInfo:       noteInfo,
			LikeNotesInfo:     likeNoteInfo,
			FavoriteNotesInfo: favoriteNotesInfo,
		},
	}
}
