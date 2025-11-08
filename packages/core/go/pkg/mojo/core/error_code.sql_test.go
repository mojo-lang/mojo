package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ErrorCodeItem struct {
	ID   uint
	Name string
	EC   *ErrorCode
}

func TestErrorCode_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := ErrorCodeItem{Name: "mojo", EC: NewErrorCode(404)}
	DB.Migrator().DropTable(&ErrorCodeItem{})
	DB.AutoMigrate(&ErrorCodeItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&ErrorCodeItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var ec ErrorCode
	err = DB.Model(&ErrorCodeItem{}).Select("ec").Where("name = ?", item.Name).Scan(&ec).Error
	assert.NoError(t, err)

	assert.Equal(t, int32(404), ec.Code)
}
