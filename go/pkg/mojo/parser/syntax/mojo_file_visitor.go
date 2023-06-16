package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type MojoFileVisitor struct {
	*BaseMojoParserVisitor
}

func NewMojoFileVisitor() *MojoFileVisitor {
	visitor := &MojoFileVisitor{}
	return visitor
}

func (m *MojoFileVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *MojoFileContext:
		return val.Accept(m)
	default:
		return nil
	}
}

func (m *MojoFileVisitor) VisitMojoFile(ctx *MojoFileContext) interface{} {
	if statementsCtx := ctx.Statements(); statementsCtx != nil {
		visitor := NewStatementsVisitor()
		sourceFile := &lang.SourceFile{}
		if s, ok := statementsCtx.Accept(visitor).([]*lang.Statement); ok {
			sourceFile.Statements = append(sourceFile.Statements, s...)
		}

		if visitor.FreeFloatingDocument != nil {
			sourceFile.TailingComments = lang.NewComments(visitor.FreeFloatingDocument)
		}

		if len(sourceFile.Statements) > 0 || len(sourceFile.TailingComments) > 0 {
			return sourceFile
		}

		// maybe something wrong
		return nil
	}

	// source file without the statements but may be having comments
	return &lang.SourceFile{}
}
