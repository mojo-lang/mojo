package http

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func ParseStatus(status string) (*Status, error) {
	v := &Status{}
	err := v.Parse(status)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (x *Status) Parse(status string) error {
	if x != nil && len(status) > 0 {
		codeStr := status
		if pos := strings.Index(status, " "); pos > 0 {
			codeStr = status[0:pos]
			x.Reason = strings.TrimSpace(status[pos+1:])
		}
		if code, err := strconv.Atoi(codeStr); err != nil {
			return fmt.Errorf("failed to parse status in code (%s) part, error: %w", codeStr, err)
		} else {
			x.Code = int32(code)
		}
	}
	return nil
}

func (x *Status) Format() string {
	if x != nil {
		buffer := bytes.Buffer{}
		buffer.WriteString(strconv.FormatInt(int64(x.Code), 10))

		if len(x.Reason) > 0 {
			buffer.WriteByte(' ')
			buffer.WriteString(x.Reason)
		}

		return buffer.String()
	}
	return ""
}
