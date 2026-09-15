package db_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/stretchr/testify/require"
	sqlite "modernc.org/sqlite"
)

// Import modernc explicitly to reproduce the combination introduced by the
// OpenAPI validator. Importing a second driver named "sqlite" panics at startup.
func TestNewSQLite(t *testing.T) {
	database := db.New(&db.Config{
		Driver: db.SqliteDriverName,
		Dsn:    filepath.Join(t.TempDir(), "test.db") + "?_pragma=busy_timeout(5000)",
	})
	require.NotNil(t, database)
	connection, err := database.DB.DB()
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, connection.Close()) })
	require.IsType(t, &sqlite.Driver{}, connection.Driver())

	var timeout int
	require.NoError(t, database.Raw("PRAGMA busy_timeout").Scan(&timeout).Error)
	require.Equal(t, 5000, timeout)

	type record struct {
		ID        uint `gorm:"primaryKey"`
		Name      string
		CreatedAt time.Time
	}
	require.NoError(t, database.AutoMigrate(&record{}))
	written := record{Name: "before"}
	require.NoError(t, database.Create(&written).Error)
	require.NotZero(t, written.ID)
	require.NoError(t, database.Model(&written).Update("name", "after").Error)
	var read record
	require.NoError(t, database.First(&read, written.ID).Error)
	require.Equal(t, "after", read.Name)
	require.True(t, read.CreatedAt.Equal(written.CreatedAt))
	require.NoError(t, database.Delete(&read).Error)
	var remaining int64
	require.NoError(t, database.Model(&record{}).Count(&remaining).Error)
	require.Zero(t, remaining)
}
