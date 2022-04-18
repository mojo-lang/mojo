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

var externalMessages map[string]string
var externalMessagesOnce sync.Once

func RegisterMessageGoPackageName(msg string, pkg string) (existed bool) {
    externalMessagesOnce.Do(func() {
        externalMessages = make(map[string]string)
    })

    _, existed = externalMessages[msg]
    externalMessages[msg] = pkg
    return
}

func GoPackageName(typename string) string {
    if pkg, ok := externalMessages[typename]; ok {
        return pkg
    } else if _, ok = goStdTypeNames[typename]; ok {
        return ""
    } else {
        return "pb"
    }
}
