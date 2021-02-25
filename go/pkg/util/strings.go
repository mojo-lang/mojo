package util

import "github.com/mojo-lang/core/go/pkg/mojo/core"

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
		strings.Values = UniqueStringSlice(strings.Values)
	}
}
