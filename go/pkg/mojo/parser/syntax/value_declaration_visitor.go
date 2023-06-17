package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ValueDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewValueDeclarationVisitor() *ValueDeclarationVisitor {
	return &ValueDeclarationVisitor{}
}

func (v *ValueDeclarationVisitor) VisitStructMemberDeclaration(ctx *StructMemberDeclarationContext) interface{} {
	if nameCtx := ctx.DeclarationIdentifier(); nameCtx != nil {
		if name, ok := nameCtx.Accept(v).(string); ok {
			decl := &lang.ValueDecl{
				StartPosition: GetPosition(ctx.GetStart()),
				EndPosition:   GetPosition(ctx.GetStop()),
				Name:          name,
				NamePosition:  GetPosition(nameCtx.GetStart()),
				Type:          GetTypeAnnotation(ctx.TypeAnnotation()),
			}

			if initializerCtx := ctx.Initializer(); initializerCtx != nil {
				if expr, ok := initializerCtx.Accept(v).(*lang.Expression); ok && expr != nil {
					decl.Initializer = &lang.Initializer{
						StartPosition: GetPosition(initializerCtx.GetStart()),
						EndPosition:   GetPosition(initializerCtx.GetStop()),
						Value:         expr,
					}
				}
			}
			return decl
		}
	}

	return nil
}

func (v *ValueDeclarationVisitor) VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{} {
	return ctx.GetText()
}

func (v *ValueDeclarationVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	initializersCtx := ctx.PatternInitializers()
	if initializersCtx != nil {
		if decls, ok := initializersCtx.Accept(v).([]*lang.ValueDecl); ok && len(decls) > 0 {
			decl := decls[0] // fixme
			return &lang.VariableDecl{
				StartPosition:  decl.StartPosition,
				EndPosition:    decl.EndPosition,
				Implicit:       decl.Implicit,
				Document:       decl.Document,
				PackageName:    decl.PackageName,
				SourceFileName: decl.SourceFileName,
				Name:           decl.Name,
				Attributes:     decl.Attributes,
				Group:          decl.Group,
				Type:           decl.Type,
				Initializer:    decl.Initializer,
			}
		}
	}
	return nil
}

func (v *ValueDeclarationVisitor) VisitPatternInitializers(ctx *PatternInitializersContext) interface{} {
	var decls []*lang.ValueDecl
	if decl, ok := ctx.PatternInitializer().Accept(v).(*lang.ValueDecl); ok && decl != nil {
		decls = append(decls, decl)
	}
	if len(decls) > 0 {
		return decls
	}
	return nil
}

func (v *ValueDeclarationVisitor) VisitPatternInitializer(ctx *PatternInitializerContext) interface{} {
	if patternCtx := ctx.Pattern(); patternCtx != nil {
		if pattern, ok := patternCtx.Accept(v).(*lang.ValueDecl); ok && pattern != nil {
			if initializerCtx := ctx.Initializer(); initializerCtx != nil {
				if expr, ok := initializerCtx.Accept(v).(*lang.Expression); ok && expr != nil {
					pattern.Initializer = &lang.Initializer{
						StartPosition: GetPosition(initializerCtx.GetStart()),
						EndPosition:   GetPosition(initializerCtx.GetStop()),
						Value:         expr,
					}
				}
			}
			return pattern
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
	if freeFloatingDocumentCtx := ctx.FloatingStatement(); freeFloatingDocumentCtx != nil {
		return GetFreeFloatingDocument(freeFloatingDocumentCtx)
	}

	decl := &lang.ValueDecl{}
	decl.Name = ctx.DeclarationIdentifier().GetText()
	decl.Document = GetDocument(ctx.Document())
	decl.StartPosition = GetPosition(ctx.GetStart())
	decl.EndPosition = GetPosition(ctx.GetStop())

	if initializerCtx := ctx.Initializer(); initializerCtx != nil {
		if expr, ok := initializerCtx.Accept(v).(*lang.Expression); ok && expr != nil {
			decl.Initializer = &lang.Initializer{
				StartPosition: GetPosition(initializerCtx.GetStart()),
				EndPosition:   GetPosition(initializerCtx.GetStop()),
				Value:         expr,
			}
		}
	}

	for _, attributesCtx := range ctx.AllAttributes() {
		decl.Attributes = append(decl.Attributes, GetAttributes(attributesCtx)...)
	}

	return decl
}

func (v *ValueDeclarationVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	decl := &lang.ValueDecl{}

	decl.Name = ctx.LabelIdentifier().GetText()
	decl.StartPosition = GetPosition(ctx.GetStart())
	decl.EndPosition = GetPosition(ctx.GetStop())

	if initializerCtx := ctx.Initializer(); initializerCtx != nil {
		if expr, ok := initializerCtx.Accept(v).(*lang.Expression); ok && expr != nil {
			decl.Initializer = &lang.Initializer{
				StartPosition: GetPosition(initializerCtx.GetStart()),
				EndPosition:   GetPosition(initializerCtx.GetStop()),
				Value:         expr,
			}
		}
	}

	decl.Type = GetTypeAnnotation(ctx.TypeAnnotation())
	return decl
}

func (v *ValueDeclarationVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	if ctx != nil {
		return GetExpression(ctx.Expression())
	}
	return nil
}
