package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AnyItem struct {
	ID   uint
	Name string
	Any  *Any
}

func TestAny_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	const val = "this is a string"
	item := AnyItem{Name: "mojo", Any: NewAny(val)}
	DB.Migrator().DropTable(&AnyItem{})
	DB.AutoMigrate(&AnyItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&AnyItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var any Any
	err = DB.Model(&AnyItem{}).Select("any").Where("name = ?", item.Name).Scan(&any).Error
	assert.NoError(t, err)

	assert.Equal(t, "String", any.GetType())
	assert.Equal(t, val, any.Get().(string))
}
