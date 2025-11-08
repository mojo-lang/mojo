package core

import (
	"os"
)

// CreateDir create dir with recursively, with default os.ModePerm and os mask (022 in mac)
func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		return os.MkdirAll(filePath, os.ModePerm)
	}
	return nil
}

// IsExist check the file or directory exist, return true if exist
func IsExist(path string) bool {
	_, err := os.Stat(path) // os.Stat get file info
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsOwnerExecutable(mode os.FileMode) bool {
	return mode&0100 != 0
}

func IsGroupExecutable(mode os.FileMode) bool {
	return mode&0010 != 0
}

func IsOtherExecutable(mode os.FileMode) bool {
	return mode&0001 != 0
}

func IsAllExecutable(mode os.FileMode) bool {
	return mode&0111 == 0111
}

func IsExecutable(mode os.FileMode) bool {
	return mode&0111 != 0
}
