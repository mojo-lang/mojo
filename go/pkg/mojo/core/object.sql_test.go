package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User1 struct {
	ID     uint
	Name   string
	Object *Object
}

func TestObject_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	user := User1{Name: "mojo", Object: NewObject().SetString("foo", "bar").SetInt64("baz", 100)}
	DB.Migrator().DropTable(&User1{})
	DB.AutoMigrate(&User1{})
	DB.Save(&user)

	var count int64

	if DB.Model(&User1{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var object Object
	err = DB.Model(&User1{}).Select("object").Where("name = ?", user.Name).Scan(&object).Error
	assert.NoError(t, err)

	assert.Equal(t, "bar", object.GetString("foo"))
	assert.Equal(t, int64(100), object.GetInt64("baz"))
}
