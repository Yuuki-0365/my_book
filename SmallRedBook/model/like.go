package model

type Like struct {
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	UserId   string `json:"user_id"`
	NoteId   uint   `json:"note_id"`
}
