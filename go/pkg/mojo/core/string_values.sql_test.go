package core

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

type StringValuesItem struct {
	ID     uint
	Name   string
	Values *StringValues
}

func TestStringValues_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := StringValuesItem{Name: "mojo", Values: NewStringValues("v1", "v2")}
	_ = DB.Migrator().DropTable(&StringValuesItem{})
	_ = DB.AutoMigrate(&StringValuesItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&StringValuesItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var values StringValues
	err = DB.Model(&StringValuesItem{}).Select("values").Where("name = ?", item.Name).Scan(&values).Error
	assert.NoError(t, err)

	assert.Equal(t, item.Values.Vals, values.Vals)
}

func TestStringValues_parseSql(t *testing.T) {
	vs := &StringValues{}
	err := vs.parseValues(`{"v1","v2"}`)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vs.Vals))
	assert.Equal(t, "v1", vs.Vals[0])
}
