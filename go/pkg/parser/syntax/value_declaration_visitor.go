package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ValueDeclarationVisitor struct {
	*BaseMojoParserVisitor

	ValueDecl *lang.ValueDecl
}

func NewValueDeclarationVisitor() *ValueDeclarationVisitor {
	visitor := &ValueDeclarationVisitor{}
	visitor.ValueDecl = &lang.ValueDecl{}
	return visitor
}

func (v *ValueDeclarationVisitor) VisitStructMemberDeclaration(ctx *StructMemberDeclarationContext) interface{} {
	v.ValueDecl.Name = GetDeclarationIdentifier(ctx.DeclarationIdentifier())

	initializerCtx := ctx.Initializer()
	if initializerCtx != nil {
		initializerCtx.Accept(v)
	}

	v.ValueDecl.Type = GetTypeAnnotation(ctx.TypeAnnotation())

	return v.ValueDecl
}

func (v *ValueDeclarationVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	initializersCtx := ctx.PatternInitializers()
	if initializersCtx != nil {
		return initializersCtx.Accept(v)
	}
	return nil
}

func (v *ValueDeclarationVisitor) VisitPatternInitializers(ctx *PatternInitializersContext) interface{} {
	initializerCtxes := ctx.AllPatternInitializer()
	var decls []interface{}
	for _, c := range initializerCtxes {
		decl := c.Accept(v)
		if decl != nil {
			decls = append(decls, decl)
		}
	}
	if len(decls) > 0 {
		return decls
	}
	return nil
}

func (v *ValueDeclarationVisitor) VisitPatternInitializer(ctx *PatternInitializerContext) interface{} {
	patternCtx := ctx.Pattern()
	if patternCtx != nil {
		if pattern, ok := patternCtx.Accept(v).(*lang.ValueDecl); ok {
			v.ValueDecl = pattern

			initializerCtx := ctx.Initializer()
			if initializerCtx != nil {
				initializerCtx.Accept(v)
			}
			return v.ValueDecl
		}
	}
	return nil
}

func (v *ValueDeclarationVisitor) VisitPattern(ctx *PatternContext) interface{} {
	identifierPatternCtx := ctx.IdentifierPattern()
	if identifierPatternCtx != nil {
		if name, ok := identifierPatternCtx.Accept(v).(string); ok {
			valueDecl := &lang.ValueDecl{}
			valueDecl.Name = name
			valueDecl.Type = GetTypeAnnotation(ctx.TypeAnnotation())
			return valueDecl
		}
	}

	return nil
}

func (v *ValueDeclarationVisitor) VisitIdentifierPattern(ctx *IdentifierPatternContext) interface{} {
	return ctx.GetText()
}

func (v *ValueDeclarationVisitor) VisitEnumMember(ctx *EnumMemberContext) interface{} {
	v.ValueDecl.Name = ctx.DeclarationIdentifier().GetText()
	v.visitDocuments(ctx.Document(), nil)

	initializerCtx := ctx.Initializer()
	if initializerCtx != nil {
		initializerCtx.Accept(v)
	}

	//v.ValueDecl.Attributes = GetAttributes(ctx.Attributes())

	return v.ValueDecl
}

func (v *ValueDeclarationVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	v.ValueDecl.Name = ctx.LabelIdentifier().GetText()

	initializerCtx := ctx.Initializer()
	if initializerCtx != nil {
		initializerCtx.Accept(v)
	}

	v.ValueDecl.Type = GetTypeAnnotation(ctx.TypeAnnotation())

	return v.ValueDecl
}

func (v *ValueDeclarationVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	if ctx != nil {
		v.ValueDecl.InitialValue = GetExpression(ctx.Expression())
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
		d := &lang.Document{}
		d.Lines = append(d.Lines, document.Lines...)
		v.ValueDecl.Document = d
	}
}
