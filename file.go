package helper

import (
	"os"
)

// 判断文件或者目录是否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断是否为目录
func IsDir(dir string) bool {
	s, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断是否为文件
func IsFile(file string) bool {
	return !IsDir(file)
}
