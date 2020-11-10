package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type TypeIdentifierVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeIdentifierVisitor() *TypeIdentifierVisitor {
	visitor := &TypeIdentifierVisitor{}
	return visitor
}

func (t *TypeIdentifierVisitor) VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{} {
	identifiers := ctx.AllTypeIdentifierClause()
	var nominalType *lang.NominalType

	for _, identifier := range identifiers {
		t := identifier.Accept(t).(*lang.NominalType)

		if nominalType == nil {
			nominalType = t
		} else {
			t.EnclosingType = nominalType
			nominalType = t
		}
	}

	return nominalType
}

func (t *TypeIdentifierVisitor) VisitTypeIdentifierClause(ctx *TypeIdentifierClauseContext) interface{} {
	nominalType := &lang.NominalType{}

	nameCtx := ctx.TypeName()
	if nameCtx != nil {
		visitor := NewTypeNameVisitor()
		nominalType.Name = nameCtx.Accept(visitor).(string)
	}

	genericArgumentCtx := ctx.GenericArgumentClause()
	if genericArgumentCtx != nil {
		visitor := NewGenericArgumentsVisitor()
		nominalType.GenericArguments = genericArgumentCtx.Accept(visitor).([]*lang.GenericArgument)
	}

	return nominalType
}
