package core

type ReferencedType interface {
	FieldMask() map[string]bool

	// Referenced return the referenced value
	Referenced() interface{}
}
