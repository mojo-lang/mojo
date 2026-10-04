package core

import (
	"errors"
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type countingTestReader func([]byte) (int, error)

func (f countingTestReader) Read(p []byte) (int, error) { return f(p) }

func TestCountingReader(t *testing.T) {
	r := NewCountingReader(strings.NewReader("hello"))
	assert.Zero(t, r.Count())
	data, err := io.ReadAll(r)
	require.NoError(t, err)
	assert.Equal(t, "hello", string(data))
	assert.Equal(t, int64(5), r.Count())
	n, err := r.Read(make([]byte, 2))
	assert.Zero(t, n)
	assert.ErrorIs(t, err, io.EOF)
	assert.Equal(t, int64(5), r.Count())
}

func TestCountingReaderPartialAndInvalidReads(t *testing.T) {
	failure := errors.New("read failed")
	for _, n := range []int{2, 0, -1} {
		r := NewCountingReader(countingTestReader(func([]byte) (int, error) { return n, failure }))
		for range 2 {
			got, err := r.Read(make([]byte, 4))
			assert.Equal(t, n, got)
			assert.ErrorIs(t, err, failure)
		}
		assert.Equal(t, int64(max(n, 0)*2), r.Count())
	}
}

func TestCountingReaderConcurrentCount(t *testing.T) {
	// The underlying reader is stateless; only the wrapper's counter is shared.
	r := NewCountingReader(countingTestReader(func(p []byte) (int, error) { return len(p), nil }))
	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 100 {
				_, _ = r.Read(make([]byte, 3))
				_ = r.Count()
			}
		})
	}
	wg.Wait()
	assert.Equal(t, int64(2400), r.Count())
}
