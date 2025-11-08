package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ValuesItem struct {
	ID     uint
	Name   string
	Values *Values
}

func TestValues_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	user := ValuesItem{Name: "mojo", Values: (&Values{}).AppendString("foo").AppendInt64(100).AppendString("bar")}
	DB.Migrator().DropTable(&ValuesItem{})
	DB.AutoMigrate(&ValuesItem{})
	DB.Save(&user)

	var count int64

	if DB.Model(&ValuesItem{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var values Values
	err = DB.Model(&ValuesItem{}).Select("values").Where("name = ?", user.Name).Scan(&values).Error
	assert.NoError(t, err)

	assert.Equal(t, 3, values.Len())
	assert.Equal(t, "foo", values.GetString(0))
	assert.Equal(t, int64(100), values.GetInt64(1))
}
