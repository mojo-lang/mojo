package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func ParseErrorCode(code string) (*ErrorCode, error) {
	ec := &ErrorCode{}
	err := ec.Parse(code)
	if err != nil {
		return nil, err
	}
	return ec, nil
}

func (x *ErrorCode) Parse(code string) error {
	if x != nil && len(code) > 0 {
		segments := strings.Split(code, ".")
		value := ""
		if len(segments) > 1 {
			x.Domain = segments[0]
			value = segments[1]
		} else {
			value = segments[0]
		}

		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("failed to parse error code %w", err)
		}

		x.Code = int32(v)
	}
	return nil
}

func (x *ErrorCode) Format() string {
	if x != nil {
		buffer := bytes.Buffer{}

		if len(x.Domain) > 0 {
			buffer.WriteString(x.Domain)
			buffer.WriteByte('.')
		}
		buffer.WriteString(strconv.FormatInt(int64(x.Code), 10))
		return buffer.String()
	}
	return ""
}
