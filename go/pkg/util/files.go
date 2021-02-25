package util

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"os"
	"path/filepath"
	"strings"
)

func ClearFiles(path string, suffix string) error {
	var files []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f != nil && !f.IsDir() && strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		logs.Errorw("filepath.Walk() failed", "error", err.Error())
		return err
	}

	for _, file := range files {
		os.Remove(file)
	}

	return nil
}

//调用os.MkdirAll递归创建文件夹
func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}