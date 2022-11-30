package service

import (
	"SmallRedBook/conf"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type UploadService struct {
}

func UploadAvatarToLocalStatic(file multipart.File, userId string, userName string) (filePath string, err error) {
	// ./static/imgs/avatar/user123/
	basePath := "." + conf.AvatarPath + "user" + userId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	// ./static/imgs/avatar/user123/yuuki.jpg
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	// 写入到该路径
	err = os.WriteFile(avatarPath, content, 0777)
	if err != nil {
		return "", err
	}
	// user123/yuuki.jpg
	return "user" + userId + "/" + userName + ".jpg", err
}

func UploadNoteFileToLocalStatic(file multipart.File, userId string, noteId string, count string) (filePath string, err error) {
	// ./static/imgs/note/user123/note123/
	basePath := "." + conf.NotePath + "user" + userId + "/" + "note" + noteId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	NoteFilePath := basePath + count + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(NoteFilePath, content, 0777)
	if err != nil {
		return "", err
	}
	return basePath, nil
}

func DirExistOrNot(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
