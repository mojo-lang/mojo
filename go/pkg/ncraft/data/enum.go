package data

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Enum struct {
	Decl        *lang.EnumDecl
	PackageName string // full package name
	Name        string

	Go         *GoEnum
	Extensions map[string]interface{}
}

type GoEnum struct {
	PackageName string
	ImportPath  string
}

func (e *Enum) GetGo() *GoEnum {
	if e != nil {
		return e.Go
	}
	return nil
}
