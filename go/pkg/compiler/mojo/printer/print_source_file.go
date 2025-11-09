package printer

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/printer"
)

func (p *Printer) PrintSourceFile(ctx context.Context, file *lang.SourceFile) *Printer {
	if file == nil || p.GetError() != nil {
		return p
	}

	var packageDecl *lang.PackageDecl
	var importDecls []*lang.ImportDecl
	var statements []*lang.Statement
	for _, statement := range file.Statements {
		if decl := statement.GetDeclaration().GetPackageDecl(); decl != nil {
			packageDecl = decl
			continue
		}
		if decl := statement.GetDeclaration().GetImportDecl(); decl != nil {
			importDecls = append(importDecls, decl)
			continue
		}
		statements = append(statements, statement)
	}

	firstLineNoneBreaker := printer.NewFirstLineNoneBreaker(p.P)
	if packageDecl != nil {
		p.PrintPackageDecl(ctx, packageDecl)
	}

	if len(importDecls) > 0 {
		firstLineNoneBreaker.Break()
		for _, importDecl := range importDecls {
			p.PrintImportDecl(ctx, importDecl)
		}
	}

	if len(statements) > 0 {
		firstLineNoneBreaker.Break()
		for i, statement := range statements {
			if i > 0 {
				p.BreakLine()
			}
			p.PrintStatement(ctx, statement)
		}
	}

	return p
}
