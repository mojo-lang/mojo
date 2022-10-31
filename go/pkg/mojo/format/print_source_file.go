package format

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintSourceFile(ctx context.Context, file *lang.SourceFile) {
	if file == nil || p == nil || p.Error != nil {
		return
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

	firstLineNoneBreaker := NewFirstLineNoneBreaker(p)
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
}
