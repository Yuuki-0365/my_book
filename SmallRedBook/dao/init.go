package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var _db *gorm.DB

func Database(conn string) {
	// 配置db
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn, // data source name
		DefaultStringSize:         256,  // string类型默认长度
		DisableDatetimePrecision:  true, // 禁止datetime精度，mysql5.6之前的不支持
		DontSupportRenameIndex:    true, // 重命名索引，要把索引先删除再重建，mysql5.7不支持
		DontSupportRenameColumn:   true, // 用change重命名列，mysql8之前的数据库不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return
	}

	// 连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)                  // 设置连接池
	sqlDB.SetMaxIdleConns(100)                 // 打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 最长连接时间
	_db = db

	migration()
}

// NewDBClient 新建DbClient
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
