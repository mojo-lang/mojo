package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Duration
func (x *FieldMask) Value() (driver.Value, error) {
	if x != nil {
		return x.Format(), nil
	}
	return nil, nil
}

func (x *FieldMask) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch fm := src.(type) {
	case []byte:
		return x.Parse(string(fm))
	case string:
		return x.Parse(fm)
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, x)
	}
	return nil
}

func (x *FieldMask) GormDataType() string {
	return "string"
}
