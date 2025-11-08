package core

import (
	"io"
	"sync/atomic"
)

// CountingWriter is counter for io.Writer
type CountingWriter struct {
	count int64
	io.Writer
}

func NewCountingWriter(w io.Writer) *CountingWriter {
	return &CountingWriter{
		Writer: w,
	}
}

func (w *CountingWriter) Write(buf []byte) (int, error) {
	n, err := w.Writer.Write(buf)
	if n >= 0 {
		atomic.AddInt64(&w.count, int64(n))
	}

	return n, err
}

func (w *CountingWriter) Count() int64 {
	return atomic.LoadInt64(&w.count)
}
