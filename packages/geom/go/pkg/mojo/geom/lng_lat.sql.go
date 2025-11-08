package geom

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

func (x *LngLat) Scan(value interface{}) (err error) {
	v := reflect.ValueOf(value)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}
	switch bs := value.(type) {
	case []byte:
		_, err := fmt.Sscanf(string(bs), "(%f,%f)", &x.Longitude, &x.Latitude)
		return err
	case string:
		_, err := fmt.Sscanf(bs, "(%f,%f)", &x.Longitude, &x.Latitude)
		return err
	default:
		return fmt.Errorf("failed to Decode type %T -> %T", value, x)
	}
}

func (x *LngLat) Value() (driver.Value, error) {
	if x.IsEmpty() {
		return nil, nil
	}

	str := fmt.Sprintf("(%f,%f)", x.Longitude, x.Latitude)
	return str, nil
}

// GormDataType gorm common data type
func (x *LngLat) GormDataType() string {
	return "Point"
}

// GormDBDataType gorm db data type
// func (m LngLat) GormDBDataType(db *gorm.DB, field *schema.Field) string {
//	 switch db.Dialector.Name() {
//	 case "sqlite":
//		return "Point"
//	 case "mysql":
//		return "Point"
//	 case "postgres":
//		return "Point"
//	 }
//	 return ""
// }
