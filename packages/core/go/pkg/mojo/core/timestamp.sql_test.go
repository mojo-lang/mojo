package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TimestampItem struct {
	ID        uint
	Name      string
	Timestamp *Timestamp
}

func TestTimestamp_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := TimestampItem{Name: "mojo", Timestamp: Now()}
	DB.Migrator().DropTable(&TimestampItem{})
	DB.AutoMigrate(&TimestampItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&TimestampItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var timestamp Timestamp
	err = DB.Model(&TimestampItem{}).Select("timestamp").Where("name = ?", item.Name).Scan(&timestamp).Error
	assert.NoError(t, err)

	assert.Equal(t, item.Timestamp.Seconds, timestamp.Seconds)
	assert.Equal(t, item.Timestamp.Nanoseconds, timestamp.Nanoseconds)
}
