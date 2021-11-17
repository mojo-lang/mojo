package compiler

type Decl interface {
	GetPackageName() string
	GetFullName() string
}
