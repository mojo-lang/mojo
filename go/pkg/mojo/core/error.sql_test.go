package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ErrorItem struct {
	ID    uint
	Name  string
	Error *Error
}

func TestError_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	const val = "this is an error string"
	item := ErrorItem{Name: "mojo", Error: NewErrorFrom(404, val)}
	DB.Migrator().DropTable(&ErrorItem{})
	DB.AutoMigrate(&ErrorItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&ErrorItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var e Error
	err = DB.Model(&ErrorItem{}).Select("error").Where("name = ?", item.Name).Scan(&e).Error
	assert.NoError(t, err)

	assert.Equal(t, int32(404), e.Code.Code)
	assert.Equal(t, val, e.Message)
}

func TestError_Scan_2(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := ErrorItem{Name: "mojo"}
	DB.Migrator().DropTable(&ErrorItem{})
	DB.AutoMigrate(&ErrorItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&ErrorItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var ei ErrorItem
	err = DB.Model(&ErrorItem{}).Where("name = ?", item.Name).Scan(&ei).Error
	assert.NoError(t, err)
	assert.Equal(t, "mojo", ei.Name)
	assert.Nil(t, ei.Error)
}
