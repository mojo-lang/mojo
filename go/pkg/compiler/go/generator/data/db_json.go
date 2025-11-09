package data

type DbJSON struct {
	GoPackageName string
	InDbPkg       bool

	PackageName string
	FullName    string

	Name               string
	UnderlyingTypeName string

	StructType bool
}

type DbJSONs []*DbJSON

func (j *DbJSON) GetPackageName() string {
	return j.PackageName
}

func (j *DbJSON) GetFullName() string {
	return j.FullName
}
