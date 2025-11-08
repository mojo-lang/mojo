package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Brand
func (x *Timestamp) Value() (driver.Value, error) {
	if x != nil {
		return x.ToTime(), nil
	}

	return nil, nil
}

func (x *Timestamp) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch bs := src.(type) {
	case []byte:
		return x.Parse(string(bs))
	case string:
		return x.Parse(bs)
	case time.Time:
		x.FromTime(bs)
	default:
		return fmt.Errorf("Could not not Decode type %T -> %T", src, x)
	}

	return nil
}

func (x *Timestamp) GormDataType() string {
	return "time"
}
