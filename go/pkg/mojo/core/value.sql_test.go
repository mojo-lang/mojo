package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User3 struct {
	ID    uint
	Name  string
	Value *Value
}

func TestValue_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	user := User3{Name: "mojo", Value: NewStringValue("foobar")}
	DB.Migrator().DropTable(&User3{})
	DB.AutoMigrate(&User3{})
	DB.Save(&user)

	var count int64

	if DB.Model(&User3{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var value Value
	err = DB.Model(&User3{}).Select("value").Where("name = ?", user.Name).Scan(&value).Error
	assert.NoError(t, err)

	assert.Equal(t, "foobar", value.GetString())
}
