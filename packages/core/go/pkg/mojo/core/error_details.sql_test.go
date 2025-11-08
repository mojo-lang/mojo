package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ErrorDetailsItem struct {
	ID      uint
	Name    string
	Details ErrorDetails
}

func TestErrorDetails_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := ErrorDetailsItem{Name: "mojo"}
	DB.Migrator().DropTable(&ErrorDetailsItem{})
	DB.AutoMigrate(&ErrorDetailsItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&ErrorDetailsItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var ei ErrorDetailsItem
	err = DB.Model(&ErrorDetailsItem{}).Where("name = ?", item.Name).Scan(&ei).Error
	assert.NoError(t, err)
	assert.Nil(t, ei.Details)
}
