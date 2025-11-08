package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Duration
func (x *Checksum) Value() (driver.Value, error) {
	if x != nil {
		return x.Format(), nil
	}
	return nil, nil
}

func (x *Checksum) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch cs := src.(type) {
	case []byte:
		return x.Parse(string(cs))
	case string:
		return x.Parse(cs)
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, x)
	}
	return nil
}

func (x *Checksum) GormDataType() string {
	return "string"
}
