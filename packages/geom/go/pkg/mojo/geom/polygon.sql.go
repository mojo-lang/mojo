package geom

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func parsePolygon(str string) ([]*LngLat, error) {
	buffer := bytes.NewBufferString(str)
	var err error
	_, err = buffer.ReadString('(')

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

	_, err = buffer.ReadString(')')
	return points, nil
}

func (x *Polygon) Scan(value interface{}) error {
	v := reflect.ValueOf(value)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch bs := value.(type) {
	case []byte:
		points, err := parsePolygon(string(bs))
		if err == nil {
			x.LineStrings = append(x.LineStrings, NewLineString(points...))
		}
		return err
	case string:
		points, err := parsePolygon(bs)
		if err == nil {
			x.LineStrings = append(x.LineStrings, NewLineString(points...))
		}
		return err
	default:
		return fmt.Errorf("failed to Decode type %T -> %T", value, x)
	}
}

func (x *Polygon) Value() (driver.Value, error) {
	if len(x.LineStrings) == 0 {
		return nil, nil
	}

	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte('(')

	for i, point := range x.LineStrings[0].Coordinates {
		if i > 0 {
			buffer.WriteByte(',')
		}
		buffer.WriteString(fmt.Sprintf("(%f,%f)", point.Longitude, point.Latitude))
	}

	buffer.WriteByte(')')
	return buffer.String(), nil
}

// GormDataType gorm common data type
func (x *Polygon) GormDataType() string {
	return "Polygon"
}
