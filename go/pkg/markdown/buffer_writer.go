package markdown

import (
	"bytes"
	"strings"
)

type Writer interface {
	Bytes() []byte
	String() string
	Len() int
	Truncate(n int)
	Write(p []byte) (n int, err error)
	WriteByte(b byte) error
	WriteString(s string) (n int, err error)
}

type indentWriter struct {
	w           Writer
	prefix      []byte
	wroteIndent bool
}

// New creates a new indent writer that indents non-empty lines with indent number of tabs.
func NewIndentWriter(w Writer, indent int) Writer {
	return &indentWriter{
		w:      w,
		prefix: bytes.Repeat([]byte{'\t'}, indent),
	}
}

func NewIndentWriterWithPrefix(w Writer, indent int, prefix string) Writer {
	return &indentWriter{
		w:      w,
		prefix: []byte(strings.Repeat("\t", indent) + prefix),
	}
}

func (iw *indentWriter) Write(p []byte) (n int, err error) {
	for i, b := range p {
		if b == '\n' {
			iw.wroteIndent = false
		} else {
			if !iw.wroteIndent {
				_, err = iw.w.Write(iw.prefix)
				if err != nil {
					return n, err
				}
				iw.wroteIndent = true
			}
		}
		_, err = iw.w.Write(p[i : i+1])
		if err != nil {
			return n, err
		}
		n++
	}
	return len(p), nil
}

func (iw *indentWriter) WriteByte(b byte) error {
	_, err := iw.Write([]byte{b})
	return err
}

func (iw *indentWriter) WriteString(s string) (n int, err error) {
	return iw.Write([]byte(s))
}

func (iw *indentWriter) Bytes() []byte {
	return iw.w.Bytes()
}

func (iw *indentWriter) String() string {
	return iw.w.String()
}

func (iw *indentWriter) Len() int {
	return iw.w.Len()
}

func (iw *indentWriter) Truncate(n int) {
	iw.w.Truncate(n)
}
