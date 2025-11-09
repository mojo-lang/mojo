package util

import "strings"

type ParseError struct {
	FileName string
	Line     int64
	Column   int64

	Message string
}

func NewParseError(fileName string, line int64, column int64, message string) *ParseError {
	return &ParseError{
		FileName: fileName,
		Line:     line,
		Column:   column,
		Message:  message,
	}
}

func (e ParseError) Error() string {
	return e.Message
}

type ParseErrors []*ParseError

func (e ParseErrors) Error() string {
	builder := strings.Builder{}
	for _, err := range e {
		builder.WriteString(err.Error())
		builder.WriteString("\n")
	}
	return builder.String()
}
