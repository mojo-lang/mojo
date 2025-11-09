package core

import "os"

// MkDir create dir with recursively, with the perm but no os mask
func MkDir(filePath string, perm os.FileMode) error {
	if !IsExist(filePath) {
		return os.MkdirAll(filePath, perm)
	}
	return nil
}
