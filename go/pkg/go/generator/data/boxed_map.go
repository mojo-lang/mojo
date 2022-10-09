package data

type BoxedMap struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string
	FieldName     string
}

func (b *BoxedMap) GetPackageName() string {
	return b.PackageName
}

func (b *BoxedMap) GetFullName() string {
	return b.FullName
}
