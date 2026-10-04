package core

import (
	"io"
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringFSReadAndStat(t *testing.T) {
	const content = "hello 世界"
	f, err := StringFS(content).Open("input.txt")
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })
	before := time.Now()
	info, err := f.Stat()
	require.NoError(t, err)
	assert.Equal(t, "input.txt", info.Name())
	assert.Equal(t, int64(len(content)), info.Size())
	assert.Equal(t, fs.ModePerm, info.Mode())
	assert.False(t, info.IsDir())
	assert.Nil(t, info.Sys())
	assert.False(t, info.ModTime().Before(before))
	buf := make([]byte, 2)
	n, err := f.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 2, n)
	info, err = f.Stat()
	require.NoError(t, err)
	assert.Equal(t, int64(len(content)), info.Size(), "file size must not shrink after reading")
	rest, err := io.ReadAll(f)
	require.NoError(t, err)
	assert.Equal(t, content, string(buf)+string(rest))
	n, err = f.Read(buf)
	assert.Zero(t, n)
	assert.ErrorIs(t, err, io.EOF)
}

func TestStringFSIndependentHandles(t *testing.T) {
	filesystem := StringFS("abc")
	for range 2 {
		data, err := fs.ReadFile(filesystem, "input.txt")
		require.NoError(t, err)
		assert.Equal(t, "abc", string(data))
	}
	data, err := fs.ReadFile(StringFS(""), "empty")
	require.NoError(t, err)
	assert.Empty(t, data)
}

func TestStringFSClosedHandle(t *testing.T) {
	f, err := StringFS("abc").Open("input.txt")
	require.NoError(t, err)
	require.NoError(t, f.Close())
	assert.ErrorIs(t, f.Close(), fs.ErrClosed)
	_, err = f.Stat()
	assert.ErrorIs(t, err, fs.ErrClosed)
	n, err := f.Read(make([]byte, 3))
	assert.Zero(t, n)
	assert.ErrorIs(t, err, fs.ErrClosed)
}
