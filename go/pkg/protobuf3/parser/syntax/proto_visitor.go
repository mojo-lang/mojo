package syntax

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ProtoVisitor struct {
	*BaseProtobuf3Visitor
}

func NewProtoVisitor() *ProtoVisitor {
	visitor := &ProtoVisitor{}
	return visitor
}

func (v *ProtoVisitor) VisitProto(ctx *ProtoContext) interface{} {
	sourceFile := &lang.SourceFile{}

	if syntax := ctx.Syntax(); syntax != nil {
		if attr, ok := syntax.Accept(v).(*lang.Attribute); ok {
			sourceFile.Attributes = append(sourceFile.Attributes, attr)
		}
	}

	allPackages := ctx.AllPackageStatement()
	for _, pkg := range allPackages {
		if decl, ok := pkg.Accept(v).(*lang.PackageDecl); ok {
			sourceFile.Statements = append(sourceFile.Statements, lang.NewPackageDeclStatement(decl))
		}
	}

	allImports := ctx.AllImportStatement()
	for _, importCtx := range allImports {
		if decl, ok := importCtx.Accept(v).(*lang.ImportDecl); ok {
			sourceFile.Statements = append(sourceFile.Statements, lang.NewImportDeclStatement(decl))
		}
	}

	allOptions := ctx.AllOptionStatement()
	for _, option := range allOptions {
		if attr, ok := option.Accept(NewOptionStatementVisitor()).(*lang.Attribute); ok {
			sourceFile.Attributes = append(sourceFile.Attributes, attr)
		}
	}

	allDefs := ctx.AllTopLevelDef()
	for _, def := range allDefs {
		if decl, ok := def.Accept(v).(*lang.Declaration); ok {
			sourceFile.Statements = append(sourceFile.Statements, lang.NewDeclarationStatement(decl))
		}
	}

	return sourceFile
}

func (v *ProtoVisitor) VisitSyntax(ctx *SyntaxContext) interface{} {
	syntax := ctx.PROTO3_LIT_DOBULE()
	if syntax != nil {
		return lang.NewStringAttribute("protobuf", "syntax", core.RemoveDoubleQuote(syntax.GetText()))
	}

	syntax = ctx.PROTO3_LIT_SINGLE()
	if syntax != nil {
		return lang.NewStringAttribute("protobuf", "syntax", core.RemoveSingleQuote(syntax.GetText()))
	}

	return nil
}

func (v *ProtoVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	str := ctx.StrLit().GetText()
	decl := &lang.ImportDecl{
		ImportFileName: str,
	}
	if pub := ctx.PUBLIC(); pub != nil {
		decl.Filter = "public"
	}
	if weak := ctx.WEAK(); weak != nil {
		decl.Filter = "weak"
	}
	return decl
}

func (v *ProtoVisitor) VisitPackageStatement(ctx *PackageStatementContext) interface{} {
	pkg, name := lang.ParseIdentifierName(ctx.FullIdent().GetText())
	return &lang.PackageDecl{
		PackageName: pkg,
		Name:        name,
	}
}

func (v *ProtoVisitor) VisitTopLevelDef(ctx *TopLevelDefContext) interface{} {
	if enumDef := ctx.EnumDef(); enumDef != nil {
		if decl, ok := enumDef.Accept(NewEnumDefVisitor()).(*lang.EnumDecl); ok {
			return lang.NewEnumDeclaration(decl)
		}
	}
	if messageDef := ctx.MessageDef(); messageDef != nil {
		if decl, ok := messageDef.Accept(NewMessageDefVisitor()).(*lang.StructDecl); ok {
			return lang.NewStructDeclaration(decl)
		}
	}
	if serviceDef := ctx.ServiceDef(); serviceDef != nil {
		if decl, ok := serviceDef.Accept(NewServiceDefVisitor()).(*lang.InterfaceDecl); ok {
			return lang.NewInterfaceDeclaration(decl)
		}
	}
	return nil
}
