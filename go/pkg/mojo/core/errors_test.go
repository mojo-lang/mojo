package core

import (
	"errors"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTypedErrorsAndWrapping(t *testing.T) {
	cases := []struct {
		name   string
		create func(string, ...interface{}) error
		is     func(error) bool
		code   *ErrorCode
	}{
		{"BadRequest", func(s string, args ...interface{}) error { return NewBadRequestError(s, args...) }, IsBadRequestError, BadRequest},
		{"InvalidArgument", func(s string, args ...interface{}) error { return NewInvalidArgumentError(s, args...) }, IsInvalidArgumentError, InvalidArgument},
		{"MalformedSyntax", func(s string, args ...interface{}) error { return NewMalformedSyntaxError(s, args...) }, IsMalformedSyntaxError, MalformedSyntax},
		{"FailedPrecondition", func(s string, args ...interface{}) error { return NewFailedPreconditionError(s, args...) }, IsFailedPreconditionError, FailedPrecondition},
		{"OutOfRange", func(s string, args ...interface{}) error { return NewOutOfRangeError(s, args...) }, IsOutOfRangeError, OutOfRange},
		{"Unauthenticated", func(s string, args ...interface{}) error { return NewUnauthenticatedError(s, args...) }, IsUnauthenticatedError, Unauthenticated},
		{"PermissionDenied", func(s string, args ...interface{}) error { return NewPermissionDeniedError(s, args...) }, IsPermissionDeniedError, PermissionDenied},
		{"NotFound", func(s string, args ...interface{}) error { return NewNotFoundError(s, args...) }, IsNotFoundError, NotFound},
		{"AlreadyExists", func(s string, args ...interface{}) error { return NewAlreadyExistsError(s, args...) }, IsAlreadyExistsError, AlreadyExists},
		{"Aborted", func(s string, args ...interface{}) error { return NewAbortedError(s, args...) }, IsAbortedError, Aborted},
		{"ResourceExhausted", func(s string, args ...interface{}) error { return NewResourceExhaustedError(s, args...) }, IsResourceExhaustedError, ResourceExhausted},
		{"Cancelled", func(s string, args ...interface{}) error { return NewCancelledError(s, args...) }, IsCancelledError, Cancelled},
		{"UnknownError", func(s string, args ...interface{}) error { return NewUnknownError(s, args...) }, IsUnknownError, UnknownError},
		{"InternalError", func(s string, args ...interface{}) error { return NewInternalError(s, args...) }, IsInternalError, InternalError},
		{"DataLoss", func(s string, args ...interface{}) error { return NewDataLossError(s, args...) }, IsDataLossError, DataLoss},
		{"Unimplemented", func(s string, args ...interface{}) error { return NewUnimplementedError(s, args...) }, IsUnimplementedError, Unimplemented},
		{"Unavailable", func(s string, args ...interface{}) error { return NewUnavailableError(s, args...) }, IsUnavailableError, Unavailable},
		{"DeadlineExceeded", func(s string, args ...interface{}) error { return NewDeadlineExceededError(s, args...) }, IsDeadlineExceededError, DeadlineExceeded},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.create("failure %d", 7)
			assert.Equal(t, "failure 7", err.Error())
			wrapped := fmt.Errorf("outer: %w", err)
			assert.True(t, tt.is(wrapped))
			assert.True(t, errors.Is(wrapped, tt.create("different message")))
			assert.False(t, tt.is(nil))
			assert.False(t, tt.is(errors.New("plain")))
			for _, other := range cases {
				if other.name != tt.name {
					assert.False(t, other.is(wrapped), other.name)
				}
			}
			base := AsError(wrapped)
			require.NotNil(t, base)
			assert.Equal(t, tt.code, base.Code)
			assert.True(t, IsError(wrapped))
			assert.Same(t, base, errors.Unwrap(err))
			typed := err.(interface {
				ToError() *Error
				StatusCode() int
				GetCode() *ErrorCode
				GetMessage() string
			})
			assert.Same(t, base, typed.ToError())
			assert.Equal(t, base.StatusCode(), typed.StatusCode())
			assert.Equal(t, tt.code, typed.GetCode())
			assert.Equal(t, "failure 7", typed.GetMessage())
		})
	}
}

func TestBasicErrorSQLAndDetails(t *testing.T) {
	original := NewNotFoundError("missing %s", "item")
	detail := &RetryInfo{RetryDelay: Second(1)}
	original.AddDetail(detail)
	assert.Len(t, original.GetDetails(), 1)
	assert.True(t, original.ContainsDetailType(&RetryInfo{}))
	encoded, err := original.Value()
	require.NoError(t, err)
	restored := newBasicError(NotFound, "")
	require.NoError(t, restored.Scan(encoded))
	assert.Equal(t, original.Code, restored.Code)
	assert.Equal(t, original.Message, restored.Message)
	assert.Equal(t, "string", restored.GormDataType())
	assert.NotEmpty(t, restored.String())
	var absent *basicError
	encoded, err = absent.Value()
	assert.NoError(t, err)
	assert.Nil(t, encoded)
	assert.NoError(t, absent.Scan("ignored"))
	assert.NoError(t, restored.Scan(nil))
}
