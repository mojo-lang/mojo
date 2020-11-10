package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type ValueDeclarationVisitor struct {
	*BaseMojoParserVisitor

	ValueDecl *lang.ValueDecl
}

func NewValueDeclarationVisitor() *ValueDeclarationVisitor {
	visitor := &ValueDeclarationVisitor{}
	visitor.ValueDecl = &lang.ValueDecl{}
	return visitor
}

func (v *ValueDeclarationVisitor) VisitStructMemberDeclaration (ctx *StructMemberDeclarationContext) interface{} {
	v.ValueDecl.Name =  GetDeclarationIdentifier(ctx.DeclarationIdentifier())
	v.visitDocuments(ctx.Document(), nil)

	initializerCtx := ctx.Initializer()
	if initializerCtx != nil {
		initializerCtx.Accept(v)
	}

	v.ValueDecl.Type = GetTypeAnnotation(ctx.TypeAnnotation())

	return v.ValueDecl
}

func (v *ValueDeclarationVisitor) VisitEnumMember (ctx *EnumMemberContext) interface{} {
	v.ValueDecl.Name = ctx.DeclarationIdentifier().GetText()
	v.visitDocuments(ctx.Document(), nil)

	initializerCtx := ctx.Initializer()
	if initializerCtx != nil {
		initializerCtx.Accept(v)
	}

	//v.ValueDecl.Attributes = GetAttributes(ctx.Attributes())

	return v.ValueDecl
}

func (v *ValueDeclarationVisitor) VisitFunctionParameter (ctx *FunctionParameterContext) interface{} {
	v.ValueDecl.Name = ctx.LabelIdentifier().GetText()

	initializerCtx := ctx.Initializer()
	if initializerCtx != nil {
		initializerCtx.Accept(v)
	}

	v.ValueDecl.Type = GetTypeAnnotation(ctx.TypeAnnotation())

	return v.ValueDecl
}

func (v *ValueDeclarationVisitor) VisitInitializer (ctx * InitializerContext) interface{} {
	if ctx != nil {
		v.ValueDecl.DefaultValue = GetExpression(ctx.Expression())
	}
	return false
}

func (v *ValueDeclarationVisitor) visitDocuments(doc IDocumentContext, fdoc IFollowingDocumentContext) {
	if doc != nil {
		visitor := NewDocumentVisitor()
		v.ValueDecl.Document = doc.Accept(visitor).(*lang.Document)
	}

	if fdoc != nil {
		visitor := NewDocumentVisitor()
		document := fdoc.Accept(visitor).(*lang.Document)

		v.ValueDecl.Document.Lines = append(v.ValueDecl.Document.Lines, document.Lines...)
	}
}