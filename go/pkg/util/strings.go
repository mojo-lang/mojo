package util

import "github.com/mojo-lang/core/go/pkg/mojo/core"

func UniqueStringSlice(strings []string) []string {
	return core.NewStringValues(strings...).Unique().Vals
}
