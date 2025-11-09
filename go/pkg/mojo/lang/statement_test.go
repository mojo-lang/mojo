package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatement_GetStartPosition(t *testing.T) {
	decl := &StructDecl{StartPosition: &Position{Line: 1, Column: 1}}
	pos := NewStructDeclStatement(decl).GetStartPosition()
	assert.Equal(t, int64(1), pos.Line)
	assert.Equal(t, int64(1), pos.Column)

	expr := &StringLiteralExpr{StartPosition: &Position{Line: 1, Column: 1}}
	pos = NewExpressionStatement(NewStringLiteralExpression(expr)).GetStartPosition()
	assert.Equal(t, int64(1), pos.Line)
	assert.Equal(t, int64(1), pos.Column)
}

func TestStatement_GetEndPosition(t *testing.T) {
	decl := &StructDecl{EndPosition: &Position{Line: 1, Column: 1}}
	pos := NewStructDeclStatement(decl).GetEndPosition()
	assert.Equal(t, int64(1), pos.Line)
	assert.Equal(t, int64(1), pos.Column)

	expr := &StringLiteralExpr{EndPosition: &Position{Line: 1, Column: 1}}
	pos = NewExpressionStatement(NewStringLiteralExpression(expr)).GetEndPosition()
	assert.Equal(t, int64(1), pos.Line)
	assert.Equal(t, int64(1), pos.Column)
}

func TestStatement_SetStartPosition(t *testing.T) {
	decl := &StructDecl{}
	NewStructDeclStatement(decl).SetStartPosition(&Position{
		Line:   1,
		Column: 1,
	})
	assert.Equal(t, int64(1), decl.StartPosition.Line)
	assert.Equal(t, int64(1), decl.StartPosition.Column)

	NewStructDeclStatement(decl).SetStartPosition(&Position{
		Line:   2,
		Column: 3,
	})
	assert.Equal(t, int64(2), decl.StartPosition.Line)
	assert.Equal(t, int64(3), decl.StartPosition.Column)

	NewStructDeclStatement(decl).SetStartPosition(&Position{
		LeadingComments: NewComments(&BlockComment{}),
	})
	assert.Equal(t, int64(2), decl.StartPosition.Line)
	assert.Equal(t, int64(3), decl.StartPosition.Column)
	assert.NotEmpty(t, decl.StartPosition.LeadingComments)
}

func TestStatement_SetEndPosition(t *testing.T) {
	decl := &StructDecl{}
	NewStructDeclStatement(decl).SetEndPosition(&Position{
		Line:   1,
		Column: 1,
	})
	assert.Equal(t, int64(1), decl.EndPosition.Line)
	assert.Equal(t, int64(1), decl.EndPosition.Column)

	NewStructDeclStatement(decl).SetEndPosition(&Position{
		Line:   2,
		Column: 3,
	})
	assert.Equal(t, int64(2), decl.EndPosition.Line)
	assert.Equal(t, int64(3), decl.EndPosition.Column)

	NewStructDeclStatement(decl).SetEndPosition(&Position{
		LeadingComments: NewComments(&BlockComment{}),
	})
	assert.Equal(t, int64(2), decl.EndPosition.Line)
	assert.Equal(t, int64(3), decl.EndPosition.Column)
	assert.NotEmpty(t, decl.EndPosition.LeadingComments)
}

func TestNewConstructorDeclStatement(t *testing.T) {
	stmt := NewConstructorDeclStatement(&ConstructorDecl{Name: "Foo"})
	assert.Equal(t, "Foo", stmt.GetDeclaration().GetConstructorDecl().GetName())
}

func TestNewDeclarationStatement(t *testing.T) {
	stmt := NewDeclarationStatement(NewStructDeclaration(&StructDecl{Name: "Foo"}))
	assert.Equal(t, "Foo", stmt.GetDeclaration().GetStructDecl().GetName())
}
