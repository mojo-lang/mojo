package core

import (
	"bytes"
	"fmt"
	"strings"
)

const maxChecksumLength = 128 + 7
const checksumSeparator = ':'

func (x *Checksum) Format() string {
	if !x.IsEmpty() {
		buf := bytes.NewBuffer(make([]byte, 0, maxChecksumLength))

		buf.WriteString(x.Algorithm.Format())
		buf.WriteByte(checksumSeparator)
		buf.WriteString(x.Val)

		return buf.String()
	}

	return ""
}

func (x *Checksum) Parse(value string) error {
	if x != nil && len(value) > 0 {
		segments := strings.Split(value, string(checksumSeparator))
		if len(segments) == 2 {
			if err := x.Algorithm.Parse(segments[0]); err != nil {
				return fmt.Errorf("failed to parse the checksum alogorithm: %s, error: %w", segments[0], err)
			}

			x.Val = segments[1]
			if !x.IsValid() {
				return fmt.Errorf("invalid checksum string: %s", value)
			}
		} else {
			return fmt.Errorf("failed to parse the checksum string: %s", value)
		}
	}
	return nil
}
