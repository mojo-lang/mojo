package core

const FieldMaskTypeName = "FieldMask"
const FieldMaskTypeFullName = "mojo.core.FieldMask"

// NewFieldMask
// 支持如下的格式
// field1,field2{filed2-1, field2-2}, field3
func NewFieldMask(fields string) *FieldMask {
	mask := &FieldMask{}
	mask.Parse(fields)
	return mask
}

func (x *FieldMask) HasField(field string) bool {
	for _, f := range x.Paths {
		if f == field {
			return true
		}
	}
	return false
}
