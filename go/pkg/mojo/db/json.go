package db

import (
	"database/sql/driver"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
)

type JSONValuer struct{}

func (v JSONValuer) Value(value interface{}) (driver.Value, error) {
	if val := reflect.ValueOf(value); val.IsZero() {
		return nil, nil
	}

	return jsoniter.ConfigFastest.MarshalToString(value)
}

type JSONScanner struct{}

func (s JSONScanner) Scan(value interface{}, src interface{}) error {
	if src == nil {
		return nil
	}
	if val := reflect.ValueOf(value); !val.IsValid() || (val.CanAddr() && val.IsNil()) {
		return nil
	}

	switch v := src.(type) {
	case []byte:
		return jsoniter.ConfigFastest.Unmarshal(v, value)
	case string:
		return jsoniter.ConfigFastest.UnmarshalFromString(v, value)
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: %v", src)
	}
}

type JSONDbDataType struct{}

func (JSONDbDataType) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return core.GormDBJSONType(db)
}
