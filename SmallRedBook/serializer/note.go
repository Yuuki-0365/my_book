package serializer

type NoteInfoLess struct {
	Title     string `json:"title"`
	Img       string `json:"img"`
	LikeCount int64  `json:"like_count"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
}
