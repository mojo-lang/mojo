package data

type Dependency struct {
	Name    string
	Version string
	Path    string
}

type GoMod struct {
	Name         string
	Version      string // go
	Dependencies []*Dependency
}
