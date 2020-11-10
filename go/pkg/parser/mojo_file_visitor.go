package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/lang/go/pkg/lang"
)

type MojoFileVisitor struct {
	*BaseMojoParserVisitor

	SourceFile *lang.SourceFile
}

func NewMojoFileVisitor() *MojoFileVisitor {
	visitor := &MojoFileVisitor{}
	visitor.SourceFile = &lang.SourceFile{}
	return visitor
}

func (m *MojoFileVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *MojoFileContext:
		return val.Accept(m).(bool)
	default:
		return false
	}
}

func (m *MojoFileVisitor) VisitMojoFile(ctx *MojoFileContext) interface{} {
	statements := ctx.Statements()
	if statements != nil {
		visitor := NewStatementsVisitor()
		m.SourceFile.Statements = statements.Accept(visitor).([]*lang.Statement)
	}

	return true
}
