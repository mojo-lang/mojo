package db

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONQuerySQLiteExecution(t *testing.T) {
	database := New(&Config{Driver: SqliteDriverName, Dsn: filepath.Join(t.TempDir(), "json.db")})
	require.NotNil(t, database)
	connection, err := database.DB.DB()
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, connection.Close()) })
	require.NoError(t, database.Exec("CREATE TABLE documents (id INTEGER PRIMARY KEY, payload TEXT)").Error)
	require.NoError(t, database.Exec("INSERT INTO documents(id,payload) VALUES (?,?),(?,?)", 1, "{\"user\":{\"name\":\"O'Reilly\"},\"active\":true}", 2, "{\"active\":false}").Error)
	for _, tt := range []struct {
		name  string
		query *JSONQueryExpression
		want  int64
	}{
		{"nested key", JSONQuery("payload").HasKey("user", "name"), 1},
		{"missing", JSONQuery("payload").HasKey("missing"), 0},
		{"quoted value", JSONQuery("payload").Equals("O'Reilly", "user", "name"), 1},
		{"boolean", JSONQuery("payload").Equals(false, "active"), 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var count int64
			require.NoError(t, database.Table("documents").Where(tt.query).Count(&count).Error)
			require.Equal(t, tt.want, count)
		})
	}
}
