package model

import (
	"time"
)

type Notice struct {
	Id        int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text" gorm:"type:text"`
}
