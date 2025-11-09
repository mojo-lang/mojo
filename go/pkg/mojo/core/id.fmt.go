package core

import "strconv"

func ParseId(id string) (*Id, error) {
	i := &Id{}
	err := i.Parse(id)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (x *Id) Format() string {
	switch x := x.Id.(type) {
	case *Id_Uint64Val:
		return strconv.FormatUint(x.Uint64Val, 10)
	case *Id_StringVal:
		return x.StringVal
	case *Id_UuidVal:
		return x.UuidVal.Format()
	default:
		return ""
	}
}

func (x *Id) Parse(value string) error {
	v, err := strconv.Atoi(value)
	if err == nil {
		x.SetInt(uint64(v))
	} else {
		uuid, err := ParseUuid(value)
		if err != nil {
			x.SetUuid(uuid)
		} else {
			x.SetString(value)
		}
	}

	return err
}
