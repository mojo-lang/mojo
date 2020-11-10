package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type PackageDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewPackageDeclarationVisitor() *PackageDeclarationVisitor {
	visitor := &PackageDeclarationVisitor{}
	return visitor
}

func (p *PackageDeclarationVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	packageDecl := &lang.PackageDecl{}

	packageDecl.Document = GetDocument(ctx.Document())
	//attributes := GetAttributes(ctx.Attributes())
	//metaInfo := GetObjectLiteral(ctx.ObjectLiteral())

	identifierCtx := ctx.PackageIdentifier()
	if identifierCtx != nil {
		identifiers := identifierCtx.Accept(p).([]string)
		packageDecl.Name = identifiers[len(identifiers)-1]
	}

	return packageDecl
}

func (p *PackageDeclarationVisitor) VisitPackageIdentifier(ctx *PackageIdentifierContext) interface{}  {
	identifierCtxes := ctx.AllIdentifier()

	var identifiers []string
	for _, identifier := range identifierCtxes {
		identifiers = append(identifiers, identifier.GetText())
	}

	return identifiers
}
