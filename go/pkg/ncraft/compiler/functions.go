package compiler

import (
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
)

func GoName(name string) string {
	if len(name) == 0 {
		return ""
	}

	segments := strings.Split(name, ".")
	if len(segments) == 1 {
		return strcase.ToCamel(segments[0])
	}

	var goNames []string
	for _, segment := range segments {
		goNames = append(goNames, strcase.ToCamel(segment))
	}
	return strings.Join(goNames, ".")
}

func HttpPathParameterName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}
