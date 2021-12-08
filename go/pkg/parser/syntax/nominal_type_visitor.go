package syntax

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

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
		if t, ok := ctx.Accept(visitor).(*lang.NominalType); ok {
			return t
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

func GetType(ctx IType_Context) *lang.NominalType {
	if ctx != nil {
		visitor := NewNominalTypeVisitor()
		if t, ok := ctx.Accept(visitor).(*lang.NominalType); ok {
			return t
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

func GetPrimeType(ctx IPrimeTypeContext) *lang.NominalType {
	if ctx != nil {
		visitor := NewNominalTypeVisitor()
		if t, ok := ctx.Accept(visitor).(*lang.NominalType); ok {
			return t
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

func GetBasicType(ctx IBasicTypeContext) *lang.NominalType {
	if ctx != nil {
		visitor := NewNominalTypeVisitor()
		if t, ok := ctx.Accept(visitor).(*lang.NominalType); ok {
			return t
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

func (n *NominalTypeVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	typeCtx := ctx.Type_()
	if typeCtx != nil {
		if nominalType, ok := typeCtx.Accept(n).(*lang.NominalType); ok {
			nominalType.Attributes = append(nominalType.Attributes, GetAttributes(ctx.Attributes())...)
			return nominalType
		}
	}
	return nil
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

	// optional annotation
	optional := ctx.QUESTION()
	if optional != nil {
		nominalType := GetType(ctx.Type_())
		if nominalType != nil {
			attribute := lang.NewImplicitBoolAttribute("mojo.core", "optional")
			nominalType.Attributes = append(nominalType.Attributes, attribute)
		}
		return nominalType
	}

	required := ctx.BANG()
	if required != nil {
		nominalType := GetType(ctx.Type_())
		if nominalType != nil {
			attribute := lang.NewImplicitBoolAttribute("mojo.core", "required")
			nominalType.Attributes = append(nominalType.Attributes, attribute)
		}
		return nominalType
	}

	return nil
}

func (n *NominalTypeVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	var nominalType *lang.NominalType

	typeCtx := ctx.Type_()
	if typeCtx != nil {
		if nt, ok := typeCtx.Accept(n).(*lang.NominalType); ok {
			nominalType = nt
		} else {
			fmt.Print("===> error")
		}
	}

	attributesCtx := ctx.Attributes()
	if attributesCtx != nil {
		visitor := NewAttributesVisitor()
		if attributes, ok := attributesCtx.Accept(visitor).([]*lang.Attribute); ok {
			nominalType.Attributes = attributes
		} else {
			fmt.Print("===> error")
		}
	}

	if nominalType != nil {
		return &lang.NominalType{
			Name: "Array",
			GenericArguments: []*lang.NominalType{
				{
					PackageName:      nominalType.PackageName,
					Name:             nominalType.Name,
					GenericArguments: nominalType.GenericArguments,
					Attributes:       nominalType.Attributes,
				},
			},
		}
	}

	return nil
}

func (n *NominalTypeVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	keyTypeCtx := ctx.Type_(0)
	if keyTypeCtx != nil {
		keyType := keyTypeCtx.Accept(n).(*lang.NominalType)

		valueTypeCtx := ctx.Type_(1)
		if valueTypeCtx != nil {
			valueType := valueTypeCtx.Accept(n).(*lang.NominalType)

			return &lang.NominalType{
				Name: "Map",
				GenericArguments: []*lang.NominalType{
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

func (n *NominalTypeVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	elementsCtx := ctx.TupleTypeElements()
	if elementsCtx != nil {
		return elementsCtx.Accept(n)
	}
	return nil
}

func (n *NominalTypeVisitor) VisitTupleTypeElements(ctx *TupleTypeElementsContext) interface{} {
	tupleType := &lang.NominalType{
		Name: "Tuple",
	}

	elements := ctx.AllTupleTypeElement()
	for _, element := range elements {
		nominalType := element.Accept(n).(*lang.NominalType)
		tupleType.GenericArguments = append(tupleType.GenericArguments, nominalType)
	}

	// move to the semantic parser
	//argumentCount := len(tupleType.GenericArguments)
	//if argumentCount == 1 {
	//	label, _ := lang.GetStringAttribute(tupleType.GenericArguments[0].Attributes, "label")
	//	if len(label) == 0 { // only one element and has no label, just like paren expression for type
	//		nominalType := tupleType.GenericArguments[0]
	//		nominalType.Attributes = append(nominalType.Attributes, tupleType.Attributes...)
	//		return nominalType
	//	}
	//}
	return tupleType
}

func (n *NominalTypeVisitor) VisitTupleTypeElement(ctx *TupleTypeElementContext) interface{} {
	typeCtx := ctx.Type_()
	if typeCtx != nil {
		nominalType := GetType(typeCtx)
		nominalType.Attributes = GetAttributes(ctx.Attributes())

		identifier := ctx.DeclarationIdentifier()
		if identifier != nil {
			nominalType.Attributes =
				lang.SetStringAttribute(nominalType.Attributes, "label", identifier.GetText())
		}

		return nominalType
	}
	return nil
}

func (n *NominalTypeVisitor) VisitUnion(ctx *UnionContext) interface{} {
	unionType := &lang.NominalType{
		Name: "Union",
	}

	typeCtx := ctx.AllBasicType()
	for i, t := range typeCtx {
		visitor := NewNominalTypeVisitor()
		if nominalType, ok := t.Accept(visitor).(*lang.NominalType); ok {
			nominalType.Attributes = GetAttributes(ctx.Attributes(i))
			unionType.GenericArguments = append(unionType.GenericArguments, nominalType)
		} else {
			// error
			return nil
		}
	}

	if len(unionType.GenericArguments) == 2 && unionType.GenericArguments[1].Name == "Union" {
		arguments := unionType.GenericArguments[1].GenericArguments
		unionType.GenericArguments = []*lang.NominalType{unionType.GenericArguments[0]}
		for _, t := range arguments {
			unionType.GenericArguments = append(unionType.GenericArguments, t)
		}
	}

	return unionType
}

func (n *NominalTypeVisitor) VisitIntersection(ctx *IntersectionContext) interface{} {
	return nil
}

func (n *NominalTypeVisitor) VisitPrime(ctx *PrimeContext) interface{} {
	primeCtx := ctx.PrimeType()
	if primeCtx != nil {
		return primeCtx.Accept(n)
	}

	return nil
}

func (n *NominalTypeVisitor) VisitPrimeType(ctx *PrimeTypeContext) interface{} {
	identifierCtx := ctx.TypeIdentifier()
	if identifierCtx != nil {
		visitor := NewTypeIdentifierVisitor()
		return identifierCtx.Accept(visitor)
	}

	arrayTypeCtx := ctx.ArrayType()
	if arrayTypeCtx != nil {
		return arrayTypeCtx.Accept(n)
	}

	dictTypeCtx := ctx.MapType()
	if dictTypeCtx != nil {
		return dictTypeCtx.Accept(n)
	}

	tupleCtx := ctx.TupleType()
	if tupleCtx != nil {
		return tupleCtx.Accept(n)
	}

	return nil
}
