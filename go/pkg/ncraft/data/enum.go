package data

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Enum struct {
    Decl        *lang.EnumDecl
    PackageName string // full package name
    Name        string

    Extensions map[string]interface{}

    Go GoEnum
}

type GoEnum struct {
    PackageName string
    ImportPath  string
}
