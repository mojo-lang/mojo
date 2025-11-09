package query

type FieldType int

const (
	FieldTypeBool     FieldType = iota + 1
	FieldTypeInteger  FieldType = iota + 1
	FieldTypeFloat    FieldType = iota + 1
	FieldTypeString   FieldType = iota + 1
	FieldTypeBytes    FieldType = iota + 1
	FieldTypeDatetime FieldType = iota + 1
	FieldTypeJSON     FieldType = iota + 1
	FieldTypeGeometry FieldType = iota + 1
)

type Field struct {
	Type     FieldType
	Repeated bool
}
