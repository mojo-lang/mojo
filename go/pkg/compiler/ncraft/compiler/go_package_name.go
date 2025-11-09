package compiler

import (
	"sync"
)

var goStdTypeNames = map[string]bool{
	"bool":    true,
	"int":     true,
	"int8":    true,
	"int16":   true,
	"int32":   true,
	"int64":   true,
	"uint":    true,
	"uint8":   true,
	"uint16":  true,
	"uint32":  true,
	"uint64":  true,
	"float32": true,
	"float64": true,
	"string":  true,
}

var importedMessages map[string]string
var importedMessagesInService map[string]map[string]bool
var importedMessagesOnce sync.Once

func RegisterMessageGoPackageName(service string, msg string, pkg string) (existed bool) {
	importedMessagesOnce.Do(func() {
		importedMessages = make(map[string]string)
		importedMessagesInService = make(map[string]map[string]bool)
	})

	var index map[string]bool
	if i, ok := importedMessagesInService[service]; !ok {
		index = make(map[string]bool)
		importedMessagesInService[service] = index
	} else {
		index = i
	}

	importedMessages[msg] = pkg

	_, existed = index[msg]
	index[msg] = true
	return
}

func GoPackageName(typename string) string {
	if pkg, ok := importedMessages[typename]; ok {
		return pkg
	} else if _, ok = goStdTypeNames[typename]; ok {
		return ""
	} else {
		return "pb"
	}
}
