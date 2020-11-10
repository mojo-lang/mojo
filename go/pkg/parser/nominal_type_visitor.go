package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type NominalTypeVisitor struct {
	*BaseMojoParserVisitor
}

func NewNominalTypeVisitor() *NominalTypeVisitor {
	visitor := &NominalTypeVisitor{}
	return visitor
}

func GetTypeAnnotation(ctx ITypeAnnotationContext) *lang.NominalType {
	if ctx != nil {
		visitor := NewNominalTypeVisitor()
		return ctx.Accept(visitor).(*lang.NominalType)
	}
	return nil
}

func GetType(ctx IType_Context) *lang.NominalType {
	if ctx != nil {
		visitor := NewNominalTypeVisitor()
		return ctx.Accept(visitor).(*lang.NominalType)
	}
	return nil
}

func (n *NominalTypeVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	var nominalType *lang.NominalType

	typeCtx := ctx.Type_()
	if typeCtx != nil {
		nominalType = typeCtx.Accept(n).(*lang.NominalType)
	}

	attributesCtx := ctx.Attributes()
	if attributesCtx != nil && nominalType != nil {
		visitor := NewAttributesVisitor()
		nominalType.Attributes = attributesCtx.Accept(visitor).([]*lang.Attribute)
	}

	return nominalType
}

func (n *NominalTypeVisitor) VisitType_(ctx *Type_Context) interface{} {
	basicCtx := ctx.BasicType()
	if basicCtx != nil {
		return basicCtx.Accept(n)
	}

	functionCtx := ctx.FunctionType()
	if functionCtx != nil {
		return functionCtx.Accept(n)
	}

	return nil
}

func (n *NominalTypeVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	var nominalType *lang.NominalType

	typeCtx := ctx.Type_()
	if typeCtx != nil {
		nominalType = typeCtx.Accept(n).(*lang.NominalType)
	}

	attributesCtx := ctx.Attributes()
	if attributesCtx != nil {
		visitor := NewAttributesVisitor()
		nominalType.Attributes = attributesCtx.Accept(visitor).([]*lang.Attribute)
	}

	return &lang.NominalType{
		Name: "Array",
		GenericArguments: []*lang.GenericArgument{
			{
				Name:             nominalType.Name,
				GenericArguments: nominalType.GenericArguments,
				Attributes:       nominalType.Attributes,
			},
		},
	}
}

func (n *NominalTypeVisitor) VisitDictionaryType(ctx *DictionaryTypeContext) interface{} {
	keyTypeCtx := ctx.Type_(0)
	if keyTypeCtx != nil {
		keyType := keyTypeCtx.Accept(n).(*lang.NominalType)

		valueTypeCtx := ctx.Type_(1)
		if valueTypeCtx != nil {
			valueType := valueTypeCtx.Accept(n).(*lang.NominalType)

			return &lang.NominalType{
				Name: "Dictionary",
				GenericArguments: []*lang.GenericArgument{
					{
						Name:             keyType.Name,
						GenericArguments: keyType.GenericArguments,
						Attributes:       keyType.Attributes,
					},
					{
						Name:             valueType.Name,
						GenericArguments: valueType.GenericArguments,
						Attributes:       valueType.Attributes,
					},
				},
			}
		}
	}

	return nil
}

func (n *NominalTypeVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return nil
}

func (n *NominalTypeVisitor) VisitTupleType(ctx *TupleTypeContext) interface{}  {
	return nil
}

func (n *NominalTypeVisitor) VisitTupleTypeElements(ctx *TupleTypeElementsContext) interface{} {
	return nil
}

func (n *NominalTypeVisitor) VisitUnion(ctx *UnionContext) interface{}  {
	return nil
}

func (n *NominalTypeVisitor) VisitIntersection(ctx *IntersectionContext) interface{}  {
	return nil
}

func (n *NominalTypeVisitor) VisitPrime(ctx *PrimeContext) interface{}  {
	primeCtx := ctx.PrimeType()
	if primeCtx != nil {
		return primeCtx.Accept(n)
	}

	return nil
}

func (n *NominalTypeVisitor) VisitPrimeType (ctx *PrimeTypeContext) interface{} {
	identifierCtx := ctx.TypeIdentifier()
	if identifierCtx != nil {
		visitor := NewTypeIdentifierVisitor()
		return identifierCtx.Accept(visitor)
	}

	arrayTypeCtx := ctx.ArrayType()
	if arrayTypeCtx != nil {
		return arrayTypeCtx.Accept(n)
	}

	dictTypeCtx := ctx.DictionaryType()
	if dictTypeCtx != nil {
		return dictTypeCtx.Accept(n)
	}

	return nil
}