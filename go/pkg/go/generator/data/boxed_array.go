package data

type BoxedArray struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string

	FieldName string
}

func (b *BoxedArray) GetPackageName() string {
	return b.PackageName
}

func (b *BoxedArray) GetFullName() string {
	return b.FullName
}
