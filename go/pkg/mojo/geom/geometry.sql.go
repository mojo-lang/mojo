package geom

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"reflect"
)

var wktPrefixes = map[string]bool{
	"point": true,
	"lines": true,
	"polyg": true,
	"multi": true,
	"geome": true,
	"srid=": true,
	"POINT": true,
	"LINES": true,
	"POLYG": true,
	"MULTI": true,
	"GEOME": true,
	"SRID=": true,
}

func (x *Geometry) Scan(value interface{}) (err error) {
	v := reflect.ValueOf(value)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	var bytes []byte
	wkt := false
	switch bs := value.(type) {
	case []byte: // wkb, using st_asbinary
		bytes = bs
	case string:
		if len(bs) >= 5 && wktPrefixes[bs[:5]] { // wkt types, using st_astext()
			bytes = []byte(bs)
			wkt = true
		} else { // default hexed binary string
			bytes, err = hex.DecodeString(bs)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("failed to Decode type %T -> %T", value, x)
	}

	if wkt {
		return x.FromWKT(string(bytes))
	} else {
		return x.FromWKB(bytes)
	}
}

func (x *Geometry) Value() (driver.Value, error) {
	return x.ToWKT(), nil
}

// GormDataType gorm common data type
func (x *Geometry) GormDataType() string {
	return "geometry"
}

// GormDBDataType gorm db data type
// func (x *Geometry) GormDBDataType(db *gorm.DB, field *schema.Field) string {
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
