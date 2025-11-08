package db

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

type DurationTest struct {
	Id       string
	Duration *core.Duration
}

func TestDuration(t *testing.T) {
	testDurationWithDialector(t, sqlite.Open(filepath.Join(os.TempDir(), "gorm-duration-test.db")))
}

//postgres.Open("user=postgres password=postgres dbname=postgres port=15432 host=172.3.0.16 sslmode=disable TimeZone=Asia/Shanghai")

func testDurationWithDialector(t *testing.T, dialector gorm.Dialector) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return
	}

	db.Migrator().DropTable(&DurationTest{})
	err = db.AutoMigrate(&DurationTest{})
	if err != nil {
		return
	}

	duration := DurationTest{
		Id:       "1",
		Duration: core.NewDurationFromSeconds(1.000012),
	}
	db.Save(&duration)

	var count int64
	d := &DurationTest{}
	assert.NoError(t, db.First(d).Count(&count).Error)
	assert.Equal(t, int64(1), d.Duration.Seconds)
	assert.Equal(t, int32(12000), d.Duration.Nanoseconds)
}
