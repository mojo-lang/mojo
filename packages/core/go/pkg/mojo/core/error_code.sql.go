package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Duration
func (x *ErrorCode) Value() (driver.Value, error) {
	if x != nil {
		return x.Format(), nil
	}
	return nil, nil
}

func (x *ErrorCode) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch ec := src.(type) {
	case []byte:
		return x.Parse(string(ec))
	case string:
		return x.Parse(ec)
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, x)
	}
	return nil
}

func (x *ErrorCode) GormDataType() string {
	return "string"
}
