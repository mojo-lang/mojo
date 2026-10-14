package query

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFieldQueryUsesBindingsAndNegation(t *testing.T) {
	database, err := gorm.Open(sqlite.Open(filepath.Join(t.TempDir(), "query.db")), &gorm.Config{DryRun: true})
	require.NoError(t, err)
	connection, err := database.DB()
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, connection.Close()) })
	for _, tt := range []struct {
		name, operator string
		not            bool
		values         []interface{}
	}{
		{"equal", "name = ?", false, []interface{}{"O'Reilly"}},
		{"not equal", "name != ?", true, []interface{}{"Ada"}},
		{"like", "name LIKE ?", false, []interface{}{"Ad%"}},
		{"not like", "name NOT LIKE ?", true, []interface{}{"%da"}},
		{"in", "name IN (?,?)", false, []interface{}{"Ada", "Grace"}},
		{"not in", "name NOT IN (?,?)", true, []interface{}{"Ada", "Grace"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ApplyFieldQuery(database.Table("items"), &FieldQuery{Name: "name", Not: tt.not, Keywords: tt.values}, nil)
			require.NoError(t, err)
			statement := result.Find(&[]QueryTable{}).Statement
			require.Contains(t, statement.SQL.String(), tt.operator)
			require.Equal(t, tt.values, statement.Vars)
		})
	}
	unchanged, err := ApplyFieldQuery(database, &FieldQuery{}, nil)
	require.NoError(t, err)
	require.Same(t, database, unchanged)
	query := (&Query{}).AddFieldQuery("name", "Ada").AddFieldQuery("name", "Grace").AddFieldQuery("age", 20)
	require.Len(t, query.FieldQueries, 2)
	require.Equal(t, []interface{}{"Ada", "Grace"}, query.FieldQueries[0].Keywords)
}

func TestProjectFieldBoundsAndNames(t *testing.T) {
	projection := &FieldProjection{Uniques: []string{"city", "country"}, Functions: []string{"count", "sum"}}
	for _, tt := range []struct {
		index          int
		name, function string
	}{{-1, "", ""}, {0, "city", "count"}, {1, "country", "sum"}, {2, "", ""}} {
		field := &ProjectField{Projection: projection, Index: tt.index}
		require.Equal(t, tt.name, field.GetName())
		require.Equal(t, tt.function, field.GetFunction())
	}
	projection.Name = "explicit"
	require.Equal(t, "explicit", (&ProjectField{Projection: projection, Index: 9}).GetName())
	var empty *ProjectField
	require.Empty(t, empty.GetName())
	require.Empty(t, empty.GetFunction())
}
