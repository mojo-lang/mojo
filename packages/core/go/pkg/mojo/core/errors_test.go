package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestNewAbortedError(t *testing.T) {
	err := NewAbortedError("test %s", "value")
	assert.Equal(t, Aborted.Code, err.Code.Code)
	assert.Equal(t, "test value", err.Message)
}

func TestNewAbortedError2(t *testing.T) {
	err := NewAbortedError("test value")
	assert.Equal(t, Aborted.Code, err.Code.Code)
	assert.Equal(t, "test value", err.Message)
}

func TestIsAbortedError(t *testing.T) {
	err := NewAbortedError("aborted")
	assert.True(t, IsAbortedError(err))
}

func TestIsCancelledError(t *testing.T) {
	assert.True(t, IsCancelledError(NewCancelledError("cancel")))
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("bad request")
	assert.Equal(t, BadRequest, err.Code)

	bytes, e := jsoniter.Marshal(err)
	assert.NoError(t, e)
	assert.NotEmpty(t, bytes)
	assert.Equal(t, []byte(`{"code":"400","message":"bad request"}`), bytes)
}

func TestErr_AddDetail(t *testing.T) {
	err := NewNotFoundError("not found")
	err.AddDetail(&RetryInfo{RetryDelay: NewDurationFromSeconds(100)})

	str, e := jsoniter.Marshal(err)
	assert.NoError(t, e)

	notFoundErr := &Error{}
	e = jsoniter.Unmarshal(str, notFoundErr)
	assert.NoError(t, e)
}

func TestErrorsIsError(t *testing.T) {
	err := NewNotFoundError("not found")
	assert.True(t, IsError(err))
}
