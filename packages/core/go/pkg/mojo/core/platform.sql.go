package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Duration
func (x *Platform) Value() (driver.Value, error) {
	if x != nil {
		return x.Format(), nil
	}
	return nil, nil
}

func (x *Platform) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch p := src.(type) {
	case string:
		return x.Parse(p)
	case []byte:
		return x.Parse(string(p))
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, x)
	}
	return nil
}

func (x *Platform) GormDataType() string {
	return "string"
}
