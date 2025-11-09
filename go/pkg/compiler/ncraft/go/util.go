package _go

import (
	"go/format"
	"io"
	"io/ioutil"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

// DiffStrings returns the line differences of two strings. Useful for
// examining how generated code differs from expected code.
func DiffStrings(a, b string) string {
	t := difflib.UnifiedDiff{
		A:        difflib.SplitLines(a),
		B:        difflib.SplitLines(b),
		FromFile: "A",
		ToFile:   "B",
		Context:  5,
	}
	text, _ := difflib.GetUnifiedDiffString(t)
	return text
}

// DiffGoCode returns normalized versions of inA and inB using the go formatter
// so that differences in indentation or trailing spaces are ignored. A diff of
// inA and inB is also returned.
func DiffGoCode(inA, inB string) (outA, outB, diff string) {
	codeFormat := func(in string) string {
		// Trim starting and ending space so format starts indenting at 0 for
		// both strings
		out := strings.TrimSpace(in)

		// Format code, if we get an error we keep out the same,
		// otherwise we use the formmated version
		outBytes, err := format.Source([]byte(out))
		if err != nil {
			return "FAILED TO FORMAT\n" + out
		} else {
			return string(outBytes)
		}
	}
	outA = codeFormat(inA)
	outB = codeFormat(inB)
	diff = DiffStrings(
		outA,
		outB,
	)
	return
}

// TestFormat takes a string representing golang code and attempts to return a
// formatted copy of that code.
func TestFormat(code string) (string, error) {
	formatted, err := format.Source([]byte(code))

	if err != nil {
		return code, err
	}

	return string(formatted), nil
}

// FormatCodeBytes takes a string representing golang code and attempts to return a
// formatted copy of that code.  If formatting fails, a warning is logged and
// the original code is returned.
func FormatCodeBytes(code []byte) []byte {
	formatted, err := format.Source(code)
	if err != nil {
		// log.WithError(err).Warn("Code formatting error, generated service will not build, outputting unformatted code")
		// return code so at least we get something to examine
		return code
	}

	return formatted
}

func FormatCode(reader io.Reader) ([]byte, error) {
	if bytes, err := ioutil.ReadAll(reader); err != nil {
		return nil, err
	} else {
		formatted, err := format.Source(bytes)
		if err != nil {
			// log.WithError(err).Warn("Code formatting error, generated service will not build, outputting unformatted code")
			// return code so at least we get something to examine
			return bytes, nil
		}
		return formatted, nil
	}
}
