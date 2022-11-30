package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/ini.v1"
	"strconv"
	"time"
)

// RedisClient Redis缓存客户端单例
var (
	RedisPool   *redis.Pool
	RedisDb     string
	RedisAddr   string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("C:\\Users\\15314\\GolandProjects\\SmallRedBook\\conf\\conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	loadRedisData(file)
	RedisPool = newRedisPool()
}

func loadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

// NewRedisPool redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     20,
		MaxActive:   100,
		IdleTimeout: 6 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			dbName, _ := strconv.Atoi(RedisDbName)
			return redis.Dial("tcp", RedisAddr, redis.DialDatabase(dbName))
		},
	}
}
