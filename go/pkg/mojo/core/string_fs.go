package core

import (
	"bytes"
	"io/fs"
	"time"
)

func StringFS(content string) fs.FS {
	return stringFs(content)
}

type stringFs string

func (f stringFs) Open(name string) (fs.File, error) {
	handle := &file{
		name:    name,
		content: bytes.NewBufferString(string(f)),
	}
	return handle, nil
}

type file struct {
	name    string
	content *bytes.Buffer
	closed  bool
}

func (f *file) Stat() (fs.FileInfo, error) {
	if f.closed {
		return nil, fs.ErrClosed
	}
	fi := fileInfo{
		name: f.name,
		size: int64(f.content.Len()),
	}
	return &fi, nil
}

func (f *file) Read(b []byte) (int, error) {
	if f.closed {
		return 0, fs.ErrClosed
	}
	return f.content.Read(b)
}

func (f *file) Close() error {
	if f.closed {
		return fs.ErrClosed
	}
	f.closed = true
	return nil
}

type fileInfo struct {
	name string
	size int64
}

func (fi *fileInfo) Name() string {
	return fi.name
}

func (fi *fileInfo) Size() int64 {
	return fi.size
}

func (fi *fileInfo) Mode() fs.FileMode {
	return fs.ModePerm
}

func (fi *fileInfo) ModTime() time.Time {
	return time.Now()
}

func (fi *fileInfo) IsDir() bool {
	return false
}

func (fi *fileInfo) Sys() interface{} {
	return nil
}
