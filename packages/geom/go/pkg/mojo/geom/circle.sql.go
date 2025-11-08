package geom

import (
	"database/sql/driver"
	fmt "fmt"
	"reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Brand
func (x *Circle) Value() (driver.Value, error) {
	if x.Radius == 0 && x.Center.IsEmpty() {
		return nil, nil
	}

	str := fmt.Sprintf("<(%f,%f),%f>", x.Center.Longitude, x.Center.Latitude, x.Radius)
	return str, nil
}

func (x *Circle) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}
	switch bs := src.(type) {
	case []byte:
		_, err := fmt.Sscanf(string(bs), "<(%f,%f),%f>", &x.Center.Longitude, &x.Center.Latitude, &x.Radius)
		return err
	case string:
		_, err := fmt.Sscanf(bs, "<(%f,%f),%f>", &x.Center.Longitude, &x.Center.Latitude, &x.Radius)
		return err
	default:
		return fmt.Errorf("Could not not Decode type %T -> %T", src, x)
	}
}
