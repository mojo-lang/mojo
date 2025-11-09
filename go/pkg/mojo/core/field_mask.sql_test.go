package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FieldMaskItem struct {
	ID   uint
	Name string
	FM   *FieldMask
}

func TestFieldMask_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := FieldMaskItem{Name: "mojo", FM: NewFieldMask("foo{bar, baz}")}
	DB.Migrator().DropTable(&FieldMaskItem{})
	DB.AutoMigrate(&FieldMaskItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&FieldMaskItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var fm FieldMask
	err = DB.Model(&FieldMaskItem{}).Select("fm").Where("name = ?", item.Name).Scan(&fm).Error
	assert.NoError(t, err)

	assert.Equal(t, 2, len(fm.Paths))
	assert.Equal(t, "foo.bar", fm.Paths[0])
	assert.Equal(t, "foo.baz", fm.Paths[1])
}
