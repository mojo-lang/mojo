package syntax

import (
	"fmt"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type TypeInheritancesVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeInheritancesVisitor() *TypeInheritancesVisitor {
	return &TypeInheritancesVisitor{}
}

func GetTypeInheritances(ctx ITypeInheritanceClauseContext) []*lang.NominalType {
	if ctx != nil {
		visitor := NewTypeInheritancesVisitor()
		if inheritances, ok := ctx.Accept(visitor).([]*lang.NominalType); ok {
			return inheritances
		} else {
			fmt.Print("===> error")
		}
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
	inheritance := GetBasicType(ctx.BasicType())

	attributesCtx := ctx.Attributes()
	if attributesCtx != nil {
		visitor := NewAttributesVisitor()
		if attributes, ok := attributesCtx.Accept(visitor).([]*lang.Attribute); ok {
			inheritance.Attributes = attributes
		} else {
			fmt.Print("===> error")
		}
	}

	return inheritance
}
