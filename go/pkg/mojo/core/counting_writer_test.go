package core

import (
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type countingTestWriter func([]byte) (int, error)

func (f countingTestWriter) Write(p []byte) (int, error) { return f(p) }

func TestCountingWriter(t *testing.T) {
	var dst bytes.Buffer
	w := NewCountingWriter(&dst)
	assert.Zero(t, w.Count())
	for _, data := range []string{"hello", "", " world"} {
		n, err := w.Write([]byte(data))
		require.NoError(t, err)
		assert.Equal(t, len(data), n)
	}
	assert.Equal(t, "hello world", dst.String())
	assert.Equal(t, int64(11), w.Count())
}

func TestCountingWriterPartialAndInvalidWrites(t *testing.T) {
	for _, n := range []int{2, 0, -1} {
		w := NewCountingWriter(countingTestWriter(func([]byte) (int, error) { return n, io.ErrShortWrite }))
		for range 2 {
			got, err := w.Write([]byte("hello"))
			assert.Equal(t, n, got)
			assert.ErrorIs(t, err, io.ErrShortWrite)
		}
		assert.Equal(t, int64(max(n, 0)*2), w.Count())
	}
}

func TestCountingWriterConcurrentCount(t *testing.T) {
	w := NewCountingWriter(io.Discard)
	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 100 {
				_, _ = w.Write([]byte("abc"))
				_ = w.Count()
			}
		})
	}
	wg.Wait()
	assert.Equal(t, int64(2400), w.Count())
}
