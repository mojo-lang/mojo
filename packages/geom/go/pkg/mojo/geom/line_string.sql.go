package geom

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"strconv"
	"strings"
)

func parseLineString(str string) ([]*LngLat, error) {
	buffer := bytes.NewBufferString(str)
	var err error
	_, err = buffer.ReadString('[')

	var points []*LngLat
	for {
		var lng, lat string
		_, err = buffer.ReadString('(')
		if err != nil {
			break
		}
		lng, err = buffer.ReadString(',')
		if err != nil {
			return nil, err
		}
		lng = strings.TrimSuffix(strings.TrimSpace(lng), ",")
		lng = strings.TrimSpace(lng)

		lat, err = buffer.ReadString(')')
		if err != nil {
			return nil, err
		}
		lat = strings.TrimSuffix(strings.TrimSpace(lat), ")")
		lat = strings.TrimSpace(lat)

		var longitude, latitude float64
		longitude, err = strconv.ParseFloat(lng, 10)
		if err != nil {
			return nil, err
		}
		latitude, err = strconv.ParseFloat(lat, 10)
		if err != nil {
			return nil, err
		}
		points = append(points, &LngLat{Longitude: longitude, Latitude: latitude})

		_, err = buffer.ReadString(',')
		if err != nil {
			break
		}
	}

	_, err = buffer.ReadString(']')
	return points, nil
}

func (x *LineString) Scan(value interface{}) error {
	v := reflect.ValueOf(value)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch bs := value.(type) {
	case []byte:
		points, err := parseLineString(string(bs))
		if err == nil {
			x.Coordinates = append(x.Coordinates, points...)
		}
		return err
	case string:
		points, err := parseLineString(bs)
		if err == nil {
			x.Coordinates = append(x.Coordinates, points...)
		}
		return err
	default:
		return fmt.Errorf("failed to Decode type %T -> %T", value, x)
	}
}

func (x *LineString) Value() (driver.Value, error) {
	if len(x.Coordinates) == 0 {
		return nil, nil
	}

	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte('[')

	for i, point := range x.Coordinates {
		if i > 0 {
			buffer.WriteByte(',')
		}
		buffer.WriteString(fmt.Sprintf("(%f,%f)", point.Longitude, point.Latitude))
	}

	buffer.WriteByte(']')
	return buffer.String(), nil
}

// GormDataType gorm common data type
func (x *LineString) GormDataType() string {
	return "path"
}

// GormDBDataType gorm db data type
func (x *LineString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "LineString"
	case "mysql":
		return "LineString"
	case "postgres":
		return "path"
	case "clickhouse":
		return "Ring"
	}
	return ""
}
