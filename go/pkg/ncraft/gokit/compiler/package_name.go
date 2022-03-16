package compiler

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

var externalMessages map[string]string

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

func init() {
    externalMessages = make(map[string]string)
}

func RegisterMessagePackage(msg string, pkg string) {
    externalMessages[msg] = pkg
}

func GetPackageName(typename string) string {
    if pkg, ok := externalMessages[typename]; ok {
        return pkg
    } else if _, ok = goStdTypeNames[typename]; ok {
        return ""
    } else {
        return "pb"
    }
}

func GetGoPackage(pkg string) string {
    return lang.GetGoPackageName(pkg)
}
