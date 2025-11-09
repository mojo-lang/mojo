package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Brand
func (x *Url) Value() (driver.Value, error) {
	if x != nil {
		return x.Format(), nil
	}
	return nil, nil
}

func (x *Url) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch bs := src.(type) {
	case []byte:
		x.Parse(string(bs))
	case string:
		x.Parse(bs)
	default:
		return fmt.Errorf("failed to Decode type %T -> %T", src, x)
	}

	return nil
}

func (x *Url) GormDataType() string {
	return "string"
}
