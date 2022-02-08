package types

type Entity struct {
	GoPackageName string
	PackageName string

	Name string
	LowerCamelName string

	Import string

	Drivers []string
}
