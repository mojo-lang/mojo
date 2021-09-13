package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

var externalMessages map[string]string

func init() {
	externalMessages = make(map[string]string)
}

func RegisterMessagePackage(msg string, pkg string) {
	externalMessages[msg] = pkg
}

func GetPackageName(typename string) string {
	if pkg, ok := externalMessages[typename]; ok {
		return pkg
	} else {
		return "pb"
	}
}

func GetGoPackage(pkg string) string {
	return lang.GetGoPackageName(pkg)
}
