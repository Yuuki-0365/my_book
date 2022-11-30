package conf

import (
	"SmallRedBook/cache"
	"SmallRedBook/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host       string
	AvatarPath string
	NotePath   string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("C:\\Users\\15314\\GolandProjects\\SmallRedBook\\conf\\conf.ini")
	if err != nil {
		panic(err)
	}
	LoadService(file)
	LoadDB(file)
	LoadEmail(file)
	LoadImgPath(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(path)
	cache.Init()
}

func LoadService(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadDB(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}
func LoadImgPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
	NotePath = file.Section("path").Key("NotePath").String()
}
