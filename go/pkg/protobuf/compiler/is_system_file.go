package compiler

import "strings"

func IsSystemFile(fileName string) bool {
	return strings.Contains(fileName, "mojo/core/numeric") ||
		strings.Contains(fileName, "mojo/core/bool") ||
		strings.Contains(fileName, "mojo/core/string") ||
		strings.Contains(fileName, "mojo/core/array") ||
		strings.Contains(fileName, "mojo/core/dictionary") ||
		strings.Contains(fileName, "mojo/core/union") ||
		strings.Contains(fileName, "mojo/core/bytes")
}
