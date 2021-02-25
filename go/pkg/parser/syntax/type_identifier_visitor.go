package syntax

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type TypeIdentifierVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeIdentifierVisitor() *TypeIdentifierVisitor {
	visitor := &TypeIdentifierVisitor{}
	return visitor
}

func (t *TypeIdentifierVisitor) VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{} {
	packageIdentifier := GetPackageIdentifier(ctx.PackageIdentifier())

	identifiers := ctx.AllTypeIdentifierClause()
	var nominalType *lang.NominalType

	for _, identifier := range identifiers {
		t := identifier.Accept(t).(*lang.NominalType)
		t.Package = packageIdentifier

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
		nominalType.Name = GetTypeName(nameCtx)
	}

	genericArgumentCtx := ctx.GenericArgumentClause()
	if genericArgumentCtx != nil {
		visitor := NewGenericArgumentsVisitor()
		if arguments, ok := genericArgumentCtx.Accept(visitor).([]*lang.NominalType); ok {
			nominalType.GenericArguments = arguments
		} else {
			fmt.Print("===> error")
		}
	}

	return nominalType
}
