package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Declaration struct {
	Declaration isDeclarationDeclaration
}

func (d *Declaration) IsUnion() {
}

type isDeclarationDeclaration interface {
	isDeclarationDeclaration()
}

type DeclarationEnumDecl struct {
	EnumDecl *EnumDecl
}
type DeclarationStructDecl struct {
	StructDecl *StructDecl
}

func (*DeclarationEnumDecl) isDeclarationDeclaration()   {}
func (*DeclarationStructDecl) isDeclarationDeclaration() {}

type EnumDecl struct {
	Implicit       bool
	PackageName    string
	SourceFileName string
	Name           string
}

type StructDecl struct {
	Implicit       bool
	PackageName    string
	SourceFileName string
	Name           string
}

type Expression struct {
	Expression isExpressionExpression
}

func (x *Expression) IsUnion() {
}

type ExpressionNullLiteralExpr struct {
	NullLiteralExpr *NullLiteralExpr
}
type ExpressionIntegerLiteralExpr struct {
	IntegerLiteralExpr *IntegerLiteralExpr
}

type isExpressionExpression interface {
	isExpressionExpression()
}

func (*ExpressionNullLiteralExpr) isExpressionExpression()    {}
func (*ExpressionIntegerLiteralExpr) isExpressionExpression() {}

type NullLiteralExpr struct {
	Kind     int32
	Implicit bool
}

type IntegerLiteralExpr struct {
	Kind       int32
	Implicit   bool
	IsNegative bool
	Value      uint64
}

type Statement struct {
	Statement isStatementStatement
}

func (x *Statement) IsUnion() {
}

type isStatementStatement interface {
	isStatementStatement()
}

type StatementDeclaration struct {
	Declaration *Declaration
}
type StatementExpression struct {
	Expression *Expression
}

func (*StatementDeclaration) isStatementStatement() {}
func (*StatementExpression) isStatementStatement()  {}

func TestGetUnionPrimeType(t *testing.T) {
	null := &NullLiteralExpr{Implicit: true}
	union := &Statement{Statement: &StatementExpression{Expression: &Expression{Expression: &ExpressionNullLiteralExpr{null}}}}

	assert.Equal(t, null, GetUnionPrimeType(union))
}

func TestGetUnionPrimeType2(t *testing.T) {
	enum := &EnumDecl{Implicit: true}
	union := &Statement{Statement: &StatementDeclaration{Declaration: &Declaration{Declaration: &DeclarationEnumDecl{enum}}}}

	assert.Equal(t, enum, GetUnionPrimeType(union))
}

func TestGetUnionField(t *testing.T) {
	null := &NullLiteralExpr{Implicit: true}
	union := &Statement{Statement: &StatementExpression{Expression: &Expression{Expression: &ExpressionNullLiteralExpr{null}}}}
	assert.True(t, GetUnionField(union, "Implicit").Bool())
}

func TestSetUnionField(t *testing.T) {
	null := &NullLiteralExpr{Implicit: true}
	union := &Statement{Statement: &StatementExpression{Expression: &Expression{Expression: &ExpressionNullLiteralExpr{null}}}}
	assert.True(t, GetUnionField(union, "Implicit").Bool())

	SetUnionField(union, "Implicit", reflect.ValueOf(false))
	assert.False(t, GetUnionField(union, "Implicit").Bool())
}
