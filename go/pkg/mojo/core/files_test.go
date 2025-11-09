package core

import (
	"os"
	"path"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkDir(t *testing.T) {
	dir := path.Join(os.TempDir(), "dir")

	err := MkDir(dir, 0777)
	assert.NoError(t, err)

	defer func() {
		os.Remove(dir)
	}()

	info, err := os.Stat(dir)
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0777), info.Mode()&os.ModePerm)
}

func TestCreateDir(t *testing.T) {
	dir := path.Join(os.TempDir(), "dir")

	err := CreateDir(dir)
	assert.NoError(t, err)

	defer func() {
		os.Remove(dir)
	}()

	mask := syscall.Umask(0)
	syscall.Umask(mask)

	info, err := os.Stat(dir)
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0777^mask), info.Mode()&os.ModePerm)
}

func TestIsExecutable(t *testing.T) {
	dir := path.Join(os.TempDir(), "dir")
	err := CreateDir(dir)
	assert.NoError(t, err)

	defer func() {
		os.RemoveAll(dir)
	}()

	data := `#!/bin/bash
echo "hello world"
`
	err = os.WriteFile(path.Join(dir, "hello.sh"), []byte(data), 0777)
	assert.NoError(t, err)

	info, err := os.Stat(path.Join(dir, "hello.sh"))
	assert.NoError(t, err)

	assert.True(t, IsExecutable(info.Mode()))
}
