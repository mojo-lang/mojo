package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStringLike struct{}

func (testStringLike) ToString() string { return "string-like" }

type testFormatter struct{}

func (testFormatter) Format() string { return "formatted" }

type testStringAndFormatter struct {
	testStringLike
	testFormatter
}

func TestToString(t *testing.T) {
	for _, tt := range []struct {
		input interface{}
		want  string
	}{
		{true, "true"}, {false, "false"},
		{int(-12), "-12"}, {int8(-12), "-12"}, {int16(-12), "-12"}, {int32(-12), "-12"}, {int64(-12), "-12"},
		{uint(12), "12"}, {uint8(12), "12"}, {uint16(12), "12"}, {uint32(12), "12"}, {uint64(18446744073709551615), "18446744073709551615"},
		{float32(1.25), "1.25000000e+00"}, {float64(-1.25), "-1.25000000e+00"},
		{[]byte{0, 255, 16}, "00ff10"}, {"你好", "你好"}, {nil, ""}, {struct{}{}, ""},
		{testStringLike{}, "string-like"}, {testFormatter{}, "formatted"}, {testStringAndFormatter{}, "string-like"},
	} {
		t.Run(fmt.Sprintf("%T/%v", tt.input, tt.input), func(t *testing.T) {
			var got string
			require.NotPanics(t, func() { got = ToString(tt.input) })
			assert.Equal(t, tt.want, got)
		})
	}
}
