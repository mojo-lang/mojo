package core

import (
	"io"
	"sync/atomic"
)

// CountingReader is counter for io.Reader
type CountingReader struct {
	count int64
	io.Reader
}

// NewCountingReader function for create new CountingReader
func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader{
		Reader: r,
	}
}

func (r *CountingReader) Read(buf []byte) (int, error) {
	n, err := r.Reader.Read(buf)

	// Read() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Reader may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	if n >= 0 {
		atomic.AddInt64(&r.count, int64(n))
	}

	return n, err
}

// Count function return counted bytes
func (r *CountingReader) Count() int64 {
	return atomic.LoadInt64(&r.count)
}
