package service

import (
	"SmallRedBook/conf"
	"fmt"
	"os"
)

func GetAllFile(path string) (s []string, err error) {
	rd, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := conf.Host + conf.HttpPort + path[1:] + fi.Name()
			s = append(s, fullName)
		}
	}
	return
}
