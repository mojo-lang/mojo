package core

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"regexp"
	"strings"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Brand
func (x *StringValues) Value() (driver.Value, error) {
	if x != nil {
		if len(x.Vals) == 0 {
			return "", nil
		}

		sb := strings.Builder{}
		sb.WriteString("{")
		for i, val := range x.Vals {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(`"`)
			sb.WriteString(val)
			sb.WriteString(`"`)
		}
		sb.WriteString("}")

		return sb.String(), nil
	}

	return nil, nil
}

var valueRegex = regexp.MustCompile(`"([^"']*)"`)

func (x *StringValues) parseValues(values string) error {
	matched := valueRegex.FindAllStringSubmatch(values, -1)
	var vs []string
	for _, v := range matched {
		if len(v) == 2 {
			vs = append(vs, v[1])
		}
	}
	if len(vs) == 0 && len(values) > 0 {
		values = strings.TrimPrefix(strings.TrimSuffix(strings.TrimSpace(values), "}"), "{")
		vs = strings.Split(values, ",")
		for i, _ := range vs {
			vs[i] = strings.TrimSpace(vs[i])
		}
	}

	x.Vals = vs
	return nil
}

func (x *StringValues) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch bs := src.(type) {
	case []byte:
		return x.parseValues(string(bs))
	case string:
		return x.parseValues(bs)
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, x)
	}
}

func (x *StringValues) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "postgres":
		return "text[]"
	}
	return "string"
}

func (x *StringValues) GormDataType() string {
	return "string"
}
