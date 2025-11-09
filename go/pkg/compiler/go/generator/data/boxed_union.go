package data

type UnionField struct {
	JsonType string // bool, number, string, array, object
	Type     string
	Name     string

	Boxed                 bool
	HasDiscriminatorField bool
}

type BoxedUnion struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string
	Discriminator string

	HasBoolField   bool
	HasNumberField bool
	HasStringField bool
	HasArrayField  bool
	HasObjectField bool

	BoolField    *UnionField
	NumberFields []*UnionField
	StringFields []*UnionField
	ArrayFields  []*UnionField
	ObjectFields []*UnionField
}

func (b *BoxedUnion) GetPackageName() string {
	return b.PackageName
}

func (b *BoxedUnion) GetFullName() string {
	return b.FullName
}
