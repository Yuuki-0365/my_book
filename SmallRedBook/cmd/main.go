package main

import (
	"SmallRedBook/conf"
	"SmallRedBook/router"
)

func main() {
	// 初始化配置
	conf.Init()
	r := router.NewRouter()
	r.Run(conf.HttpPort)
}
