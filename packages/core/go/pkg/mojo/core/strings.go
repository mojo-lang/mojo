package core

import "strings"

const (
	DoubleQuote = `"`
	SingleQuote = `'`
)

func RemoveDoubleQuote(str string) string {
	return RemoveQuote(str, DoubleQuote)
}

func RemoveSingleQuote(str string) string {
	return RemoveQuote(str, SingleQuote)
}

func RemoveQuote(str string, quote string) string {
	if strings.HasPrefix(str, quote) && strings.HasSuffix(str, quote) {
		return strings.TrimSuffix(strings.TrimPrefix(str, quote), quote)
	}
	return str
}

func IsQuotedString(str string, quote string) bool {
	return strings.HasPrefix(str, quote) && strings.HasSuffix(str, quote)
}

func QuoteString(str string) string {
	if strings.HasPrefix(str, DoubleQuote) {
		if strings.HasSuffix(str, DoubleQuote) {
			return str
		} else {
			return `"\"` + str[1:] + DoubleQuote
		}
	} else {
		if strings.HasSuffix(str, DoubleQuote) {
			return DoubleQuote + str[:len(str)-1] + `\""`
		} else {
			return DoubleQuote + str + DoubleQuote
		}
	}
}
