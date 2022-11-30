package dao

import (
	"SmallRedBook/model"
	"fmt"
)

func migration() {
	err := _db.Set("gorm_table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Notice{},
			&model.Follow{},
			&model.Fan{},
			&model.Favorite{},
			&model.Like{},
			&model.Comment{},
			&model.Note{})
	if err != nil {
		fmt.Println("err=", err)
		return
	}
}
