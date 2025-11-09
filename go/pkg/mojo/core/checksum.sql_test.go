package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ChecksumItem struct {
	ID   uint
	Name string
	CS   *Checksum
}

func TestChecksum_Scan(t *testing.T) {
	DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		t.Errorf("failed to connect database")
	}

	item := ChecksumItem{Name: "mojo", CS: NewChecksum(Checksum_ALGORITHM_MD5, []byte("md5-value"))}
	DB.Migrator().DropTable(&ChecksumItem{})
	DB.AutoMigrate(&ChecksumItem{})
	DB.Save(&item)

	var count int64

	if DB.Model(&ChecksumItem{}).Where("name = ?", item.Name).Count(&count).Error != nil || count != 1 {
		t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	}

	var cs Checksum
	err = DB.Model(&ChecksumItem{}).Select("cs").Where("name = ?", item.Name).Scan(&cs).Error
	assert.NoError(t, err)

	assert.Equal(t, Checksum_ALGORITHM_MD5, cs.Algorithm)
	assert.Equal(t, item.CS.Val, cs.Val)
}
