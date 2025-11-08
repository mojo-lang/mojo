package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Duration
func (x *Duration) Value() (driver.Value, error) {
	if x != nil {
		return x.ToSeconds(), nil
	}
	return nil, nil
}

func (x *Duration) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch duration := src.(type) {
	case float64:
		x.FromSeconds(duration)
	case string:
		seconds, err := strconv.ParseFloat(duration, 64)
		if err != nil {
			return err
		}
		x.FromSeconds(seconds)
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, x)
	}
	return nil
}

func (x *Duration) GormDataType() string {
	return "float"
}
