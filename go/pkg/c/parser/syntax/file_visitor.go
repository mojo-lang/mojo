package syntax

import (
	"regexp"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type FileVisitor struct {
	*BaseCVisitor
}

func NewFileVisitor() *FileVisitor {
	visitor := &FileVisitor{}
	return visitor
}

func (v *FileVisitor) VisitFile(ctx *FileContext) interface{} {
	sourceFile := &lang.SourceFile{}

	allIncludeDirective := ctx.AllIncludeDirective()
	for _, include := range allIncludeDirective {
		if decl, ok := include.Accept(v).(*lang.ImportDecl); ok {
			sourceFile.Imports = append(sourceFile.Imports, decl)
		}
	}

	if compilationUnit := ctx.CompilationUnit(); compilationUnit != nil {
		if file, ok := compilationUnit.Accept(v).(*lang.SourceFile); ok {
			sourceFile.Statements = file.Statements
		}
	}

	return sourceFile
}

func (v *FileVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	if context := ctx.TranslationUnit(); context != nil {
		return context.Accept(v)
	}
	return nil
}

func (v *FileVisitor) VisitTranslationUnit(ctx *TranslationUnitContext) interface{} {
	allDecls := ctx.AllExternalDeclaration()
	sourceFile := &lang.SourceFile{}
	for _, declCtx := range allDecls {
		if decl, ok := declCtx.Accept(v).(*lang.Declaration); ok {
			sourceFile.Statements = append(sourceFile.Statements, lang.NewDeclarationStatement(decl))
		}
	}
	return sourceFile
}

func (v *FileVisitor) VisitExternalDeclaration(ctx *ExternalDeclarationContext) interface{} {
	if functionDefinition := ctx.FunctionDefinition(); functionDefinition != nil {
		return functionDefinition.Accept(NewFunctionDefinitionVisitor())
	}
	if declaration := ctx.Declaration(); declaration != nil {
		return declaration.Accept(NewDeclarationVisitor())
	}
	return nil
}

var includeRegex = regexp.MustCompile(`["<][^"<>]+[">]`)

func (v *FileVisitor) VisitIncludeDirective(ctx *IncludeDirectiveContext) interface{} {
	path := ctx.GetText()
	path = includeRegex.FindString(path)

	if len(path) > 0 {
		decl := &lang.ImportDecl{}
		if strings.HasPrefix(path, "<") {
			path = strings.TrimSuffix(strings.TrimPrefix(path, "<"), ">")
			decl.ImportFileName = strings.TrimSpace(path)
			decl.Attributes = append(decl.Attributes, lang.NewBoolAttribute("c", "system_include"))
		} else if strings.HasPrefix(path, "\"") {
			decl.ImportFileName = core.RemoveQuote(path, "\"")
		}

		return decl
	}
	return &lang.ImportDecl{}
}
