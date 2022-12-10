package data

import (
	"strconv"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
)

var DigitEnglish = map[rune]string{
	'0': "zero",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
}

// EnglishNumber takes an integer and returns the english words that represents
// that number, in base ten. Examples:
//     1  -> "One"
//     5  -> "Five"
//     10 -> "OneZero"
//     48 -> "FourEight"
func EnglishNumber(i int) string {
	n := strconv.Itoa(i)
	rv := ""
	for _, c := range n {
		if e, ok := DigitEnglish[c]; ok {
			rv += strcase.ToCamel(e)
		}
	}
	return rv
}
