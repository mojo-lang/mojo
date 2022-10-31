package syntax

import (
	"fmt"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type TypeIdentifierVisitor struct {
	*BaseMojoParserVisitor
}

func GetTypeIdentifier(ctx ITypeIdentifierContext) *lang.NominalType {
	if ctx != nil {
		if nominal, ok := ctx.Accept(NewTypeIdentifierVisitor()).(*lang.NominalType); ok {
			return nominal
		}
	}
	return nil
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
		typ := identifier.Accept(t).(*lang.NominalType)
		typ.PackageName = packageIdentifier

		if nominalType == nil {
			nominalType = typ
		} else {
			typ.Enclosing = nominalType
			nominalType = typ
		}
	}

	return nominalType
}

func (t *TypeIdentifierVisitor) VisitTypeIdentifierClause(ctx *TypeIdentifierClauseContext) interface{} {
	nominalType := &lang.NominalType{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
	}

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
