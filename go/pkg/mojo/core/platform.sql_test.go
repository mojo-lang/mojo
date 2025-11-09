package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PlatformItem struct {
	ID       uint
	Name     string
	Platform *Platform
}

func TestPlatform_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := PlatformItem{Name: "mojo", Platform: NewPlatform(OS_OS_LINUX, Architecture_ARCHITECTURE_ARM64, "")}
	DB.Migrator().DropTable(&PlatformItem{})
	DB.AutoMigrate(&PlatformItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&PlatformItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var p Platform
	err = DB.Model(&PlatformItem{}).Select("platform").Where("name = ?", item.Name).Scan(&p).Error
	assert.NoError(t, err)

	assert.Equal(t, item.Platform.Os, p.Os)
	assert.Equal(t, item.Platform.Architecture, p.Architecture)
}
