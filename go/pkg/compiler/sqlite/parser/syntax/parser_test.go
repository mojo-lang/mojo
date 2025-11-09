package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseString(t *testing.T) {
	selectStmt := `SELECT CustomerName, City FROM Customers;`
	file, err := New(nil).ParseString(selectStmt)
	assert.NoError(t, err)
	assert.NotEmpty(t, file.Statements)
}
