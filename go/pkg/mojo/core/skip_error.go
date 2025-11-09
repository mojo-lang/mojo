package core

import "errors"

// SkipError used in the loop callback
type SkipError struct {
}

func (e SkipError) Error() string {
	return "skip"
}

func NewSkipError() error {
	return &SkipError{}
}

func IsSkipError(err error) bool {
	return errors.Is(err, &SkipError{})
}
