package util

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"strings"
)

func UniqueStringSlice(strings []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strings {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueStrings(strings *core.Strings) {
	if strings != nil {
		strings.Vals = UniqueStringSlice(strings.Vals)
	}
}

func RemoveQuote(str string) string {
	if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
		return strings.TrimSuffix(strings.TrimPrefix(str, `"`), `"`)
	}
	return str
}
