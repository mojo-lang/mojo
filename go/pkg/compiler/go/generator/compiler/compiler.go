package compiler

import (
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
)

func GetGoPackage(pkg string) string {
	return lang.GetGoPackageName(pkg)
}

func GetFullName(enclosing string, name string) string {
	if len(enclosing) == 0 {
		return name
	}
	return strings.ReplaceAll(enclosing, ".", "_") + "_" + name
}
