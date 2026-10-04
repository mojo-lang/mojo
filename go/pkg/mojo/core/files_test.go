package core

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMkDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "parent", "child")
	require.NoError(t, MkDir(dir, 0750))
	info, err := os.Stat(dir)
	require.NoError(t, err)
	assert.True(t, info.IsDir())
	if runtime.GOOS != "windows" {
		assert.Equal(t, os.FileMode(0750), info.Mode().Perm())
	}
	require.NoError(t, MkDir(dir, 0700))
	after, err := os.Stat(dir)
	require.NoError(t, err)
	assert.Equal(t, info.Mode(), after.Mode(), "existing directory permissions are preserved")
}

func TestCreateDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "parent", "child")
	assert.False(t, IsExist(dir))
	require.NoError(t, CreateDir(dir))
	assert.True(t, IsExist(dir))
	info, err := os.Stat(dir)
	require.NoError(t, err)
	assert.True(t, info.IsDir())
	require.NoError(t, CreateDir(dir))
}

func TestDirectoryCreationErrors(t *testing.T) {
	file := filepath.Join(t.TempDir(), "file")
	require.NoError(t, os.WriteFile(file, []byte("content"), 0600))
	assert.True(t, IsExist(file))
	child := filepath.Join(file, "child")
	assert.False(t, IsExist(child))
	assert.Error(t, CreateDir(child))
	assert.Error(t, MkDir(child, 0700))
	assert.Error(t, CreateDir(""))
	assert.Error(t, MkDir("", 0700))
}

func TestIsExecutable(t *testing.T) {
	for _, tt := range []struct {
		name                          string
		mode                          os.FileMode
		owner, group, other, all, any bool
	}{
		{"none", 0644, false, false, false, false, false},
		{"owner", 0744, true, false, false, false, true},
		{"group", 0654, false, true, false, false, true},
		{"other", 0645, false, false, true, false, true},
		{"owner and group", 0754, true, true, false, false, true},
		{"all", 0755, true, true, true, true, true},
		{"directory without execute", os.ModeDir | 0600, false, false, false, false, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.owner, IsOwnerExecutable(tt.mode))
			assert.Equal(t, tt.group, IsGroupExecutable(tt.mode))
			assert.Equal(t, tt.other, IsOtherExecutable(tt.mode))
			assert.Equal(t, tt.all, IsAllExecutable(tt.mode))
			assert.Equal(t, tt.any, IsExecutable(tt.mode))
		})
	}
}
