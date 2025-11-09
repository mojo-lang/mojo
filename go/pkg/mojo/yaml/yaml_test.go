package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Position struct {
	Filename string `json:"filename,omitempty"`
	Offset   int64  `json:"offset,omitempty"`
	Line     int64  `json:"line,omitempty"`
	Column   int64  `json:"column,omitempty"`
}

type ContinueStmt struct {
	StartPosition *Position `json:"startPosition,omitempty"`
	EndPosition   *Position `json:"endPosition,omitempty"`
	Kind          int64     `json:"kind,omitempty"`
	Implicit      bool      `json:"implicit,omitempty"`
}

var continueStmt = &ContinueStmt{
	StartPosition: &Position{
		Line:   12,
		Column: 0,
	},
	EndPosition: &Position{
		Line:   21,
		Column: 100,
	},
	Kind:     16,
	Implicit: true,
}

const continueStmtYml = "endPosition:\n  column: 100\n  line: 21\nimplicit: true\nkind: 16\nstartPosition:\n  line: 12\n"

func TestMarshal(t *testing.T) {
	y, err := Marshal(continueStmt)
	assert.NoError(t, err)
	assert.Equal(t, continueStmtYml, string(y))
}

func TestUnmarshal(t *testing.T) {
	stmt := &ContinueStmt{}
	err := Unmarshal([]byte(continueStmtYml), stmt)
	assert.NoError(t, err)
	assert.Equal(t, continueStmt.Kind, stmt.Kind)
	assert.Equal(t, continueStmt.Implicit, stmt.Implicit)
	assert.Equal(t, continueStmt.StartPosition.Line, stmt.StartPosition.Line)
	assert.Equal(t, continueStmt.StartPosition.Column, stmt.StartPosition.Column)
	assert.Equal(t, continueStmt.EndPosition.Line, stmt.EndPosition.Line)
	assert.Equal(t, continueStmt.EndPosition.Column, stmt.EndPosition.Column)
}
