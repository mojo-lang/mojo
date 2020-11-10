package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type TypeInheritancesVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeInheritancesVisitor() *TypeInheritancesVisitor {
	return &TypeInheritancesVisitor{}
}

func GetTypeInheritances(ctx ITypeInheritanceClauseContext) []*lang.NominalType {
	if ctx != nil {
		visitor := NewTypeInheritancesVisitor()
		return ctx.Accept(visitor).([]*lang.NominalType)
	}
	return []*lang.NominalType{}
}

func (t *TypeInheritancesVisitor) VisitTypeInheritanceClause(ctx *TypeInheritanceClauseContext) interface{} {
	inheritances := ctx.TypeInheritances()
	if inheritances != nil {
		return inheritances.Accept(t)
	}

	return []*lang.NominalType{}
}

func (t *TypeInheritancesVisitor) VisitTypeInheritances(ctx *TypeInheritancesContext) interface{} {
	var inheritances []*lang.NominalType
	identifierCtxes := ctx.AllTypeInheritance()
	for _, identifierCtx := range identifierCtxes {
		nominal := identifierCtx.Accept(t).(*lang.NominalType)
		inheritances = append(inheritances, nominal)
	}
	return inheritances
}

func (t *TypeInheritancesVisitor) VisitTypeInheritance(ctx *TypeInheritanceContext) interface{} {
	inheritance := &lang.NominalType{}
	identifierCtx := ctx.TypeIdentifier()
	if identifierCtx != nil {
		visitor := NewTypeIdentifierVisitor()
		inheritance = identifierCtx.Accept(visitor).(*lang.NominalType)
	}

	attributesCtx := ctx.Attributes()
	if attributesCtx != nil {
		visitor := NewAttributesVisitor()
		attributes := attributesCtx.Accept(visitor).([]*lang.Attribute)
		inheritance.Attributes = attributes
	}

	return inheritance
}