//go:build unix

package core

import (
	"os"
	"syscall"
)

// MkDir create dir with recursively, with the perm but no os mask
func MkDir(filePath string, perm os.FileMode) error {
	if !IsExist(filePath) {
		mask := syscall.Umask(0)
		defer syscall.Umask(mask)
		return os.MkdirAll(filePath, perm)
	}
	return nil
}
