package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type StructDeclarationVisitor struct {
	*BaseMojoParserVisitor

	StructDecl *lang.StructDecl
}

func NewStructDeclarationVisitor() *StructDeclarationVisitor {
	visitor := &StructDeclarationVisitor{}
	visitor.StructDecl = &lang.StructDecl{}
	visitor.StructDecl.Type = &lang.StructType{}
	return visitor
}

func (s *StructDeclarationVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	s.StructDecl.Document = GetDocument(ctx.Document())
	s.StructDecl.Attributes = GetAttributes(ctx.Attributes())

	structName := ctx.StructName()
	if structName != nil {
		s.StructDecl.Name = structName.Accept(s).(string)
	} else {
		// error happened
	}

	genericParameters := ctx.GenericParameterClause()
	if genericParameters != nil {
		visitor := NewGenericParametersVisitor()
		s.StructDecl.GenericParameters = genericParameters.Accept(visitor).([]*lang.GenericParameter)
	}

	inheritances := ctx.TypeInheritanceClause()
	if inheritances != nil {
		visitor := NewTypeInheritancesVisitor()
		s.StructDecl.Type.Inherits = inheritances.Accept(visitor).([]*lang.NominalType)
	}

	body := ctx.StructBody()
	if body != nil {
		body.Accept(s)
	}

	return s.StructDecl
}

func (s *StructDeclarationVisitor) VisitStructName(ctx *StructNameContext) interface{} {
	return ctx.GetText()
}

func (s *StructDeclarationVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	members := ctx.StructMembers()
	if members != nil {
		members.Accept(s)
	}
	return true
}

func (s *StructDeclarationVisitor) VisitStructMembers(ctx *StructMembersContext) interface{} {
	members := ctx.AllStructMember()
	for _, member := range members {
		member.Accept(s)
	}
	return true
}

func (s *StructDeclarationVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	this := s.StructDecl
	structCtx := ctx.StructDeclaration()
	if structCtx != nil {
		visitor := NewStructDeclarationVisitor()
		structDecl := structCtx.Accept(visitor).(*lang.StructDecl)
		if structCtx != nil {
			this.StructTypeDecls = append(this.StructTypeDecls, structDecl)
		}
	}

	enumCtx := ctx.EnumDeclaration()
	if enumCtx != nil {
		visitor := NewEnumDeclarationVisitor()
		enumDecl := enumCtx.Accept(visitor).(*lang.EnumDecl)
		if enumDecl != nil {
			this.EnumTypeDecls = append(this.EnumTypeDecls, enumDecl)
		}
	}

	memberCtx := ctx.StructMemberDeclaration()
	if memberCtx != nil {
		visitor := NewValueDeclarationVisitor()
		memberDecl := memberCtx.Accept(visitor).(*lang.ValueDecl)
		if memberDecl != nil {
			this.GetType().Fields = append(this.GetType().Fields, memberDecl)
		}
	}

	return true
}
