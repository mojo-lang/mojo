package converter

import "strings"

func IsSystemFile(fileName string) bool {
	return strings.Contains(fileName, "mojo/core/numeric") ||
		strings.Contains(fileName, "mojo/core/bool") ||
		strings.Contains(fileName, "mojo/core/integer") ||
		strings.Contains(fileName, "mojo/core/float") ||
		strings.Contains(fileName, "mojo/core/string") ||
		strings.Contains(fileName, "mojo/core/array") ||
		strings.Contains(fileName, "mojo/core/map") ||
		strings.Contains(fileName, "mojo/core/union") ||
		strings.Contains(fileName, "mojo/core/bytes")
}

func GetProtoFile(fileName string) string {
	if strings.HasSuffix(fileName, ".mojo") {
		fileName = strings.TrimSuffix(fileName, "mojo") + "proto"
	} else {
		fileName = fileName + ".proto"
	}
	return fileName
}
