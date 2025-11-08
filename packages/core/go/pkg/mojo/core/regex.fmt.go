package core

func ParseRegex(fieldMask string) (*FieldMask, error) {
	mask := NewFieldMask(fieldMask)
	return mask, nil
}

func (x *Regex) Format() string {
	if x != nil {
		return x.Expression
	} else {
		return ""
	}
}

func (x *Regex) ToString() string {
	return x.Format()
}

func (x *Regex) Parse(expr string) error {
	if x != nil {
		// FIXME should fix regexp in java like: "^[\\u4E00-\\u9FA5A-Za-z0-9_.]{1,32}$"
		// if _, err := regexp.Compile(expr); err != nil {
		// 	return err
		// }

		x.Expression = expr
	}
	return nil
}
