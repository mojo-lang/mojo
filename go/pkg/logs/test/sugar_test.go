package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func TestSugarFormatMessage(t *testing.T) {
	l := logs.New(nil)
	s := l.FormatMessage("", "foo", 100, "bar", "hello", "baz", true, "error", core.NewNotFoundError("error"))
	assert.NotEmpty(t, s)
	assert.Equal(t, "foo:100,bar:hello,baz:true,error:{\"code\":\"404\",\"message\":\"error\"}", s)
}
