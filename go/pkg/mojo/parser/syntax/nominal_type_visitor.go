package syntax

import (
	"fmt"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
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
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
			PackageName:   "mojo.core",
			Name:          "Array",
			GenericArguments: []*lang.NominalType{
				{
					StartPosition:    nominalType.StartPosition,
					EndPosition:      nominalType.EndPosition,
					Document:         nominalType.Document,
					PackageName:      nominalType.PackageName,
					Implicit:         nominalType.Implicit,
					Name:             nominalType.Name,
					GenericArguments: nominalType.GenericArguments,
					Attributes:       nominalType.Attributes,
					Enclosing:        nominalType.Enclosing,
				},
			},
		}
	}

	return nil
}

func (n *NominalTypeVisitor) VisitKeyAttributes(ctx *KeyAttributesContext) interface{} {
	return GetAttributes(ctx.Attributes())
}

func (n *NominalTypeVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	if len(ctx.AllType_()) != 2 {
		return nil
	}
	var keyAttributes []*lang.Attribute
	if ctx.KeyAttributes() != nil {
		if attributes, ok := ctx.KeyAttributes().Accept(n).([]*lang.Attribute); ok {
			keyAttributes = attributes
		}
	}

	if keyTypeCtx := ctx.Type_(0); keyTypeCtx != nil {
		if keyType, ok := keyTypeCtx.Accept(n).(*lang.NominalType); ok {
			if valueTypeCtx := ctx.Type_(1); valueTypeCtx != nil {
				if valueType, ok := valueTypeCtx.Accept(n).(*lang.NominalType); ok {
					return &lang.NominalType{
						StartPosition: GetPosition(ctx.GetStart()),
						EndPosition:   GetPosition(ctx.GetStop()),
						PackageName:   "mojo.core",
						Name:          "Map",
						GenericArguments: []*lang.NominalType{
							{
								StartPosition:    keyType.StartPosition,
								EndPosition:      keyType.EndPosition,
								Document:         keyType.Document,
								PackageName:      keyType.PackageName,
								Implicit:         keyType.Implicit,
								Name:             keyType.Name,
								GenericArguments: keyType.GenericArguments,
								Attributes:       keyAttributes,
								Enclosing:        keyType.Enclosing,
							},
							{
								StartPosition:    valueType.StartPosition,
								EndPosition:      valueType.EndPosition,
								Document:         valueType.Document,
								PackageName:      valueType.PackageName,
								Implicit:         valueType.Implicit,
								Name:             valueType.Name,
								GenericArguments: valueType.GenericArguments,
								Attributes:       GetAttributes(ctx.Attributes()),
								Enclosing:        valueType.Enclosing,
							},
						},
					}
				}
			}
		}
	}

	return nil
}

func (n *NominalTypeVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return nil
}

func (n *NominalTypeVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	if elementsCtx := ctx.TupleTypeElements(); elementsCtx != nil {
		return elementsCtx.Accept(n)
	}
	return nil
}

func (n *NominalTypeVisitor) VisitTupleTypeElements(ctx *TupleTypeElementsContext) interface{} {
	tupleType := &lang.NominalType{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
		PackageName:   "mojo.core",
		Name:          "Tuple",
	}

	elements := ctx.AllTupleTypeElement()
	for _, element := range elements {
		nominalType := element.Accept(n).(*lang.NominalType)
		tupleType.GenericArguments = append(tupleType.GenericArguments, nominalType)
	}

	// move to the semantic parser
	// argumentCount := len(tupleType.GenericArguments)
	// if argumentCount == 1 {
	//	label, _ := lang.GetStringAttribute(tupleType.GenericArguments[0].Attributes, "label")
	//	if len(label) == 0 { // only one element and has no label, just like paren expression for type
	//		nominalType := tupleType.GenericArguments[0]
	//		nominalType.Attributes = append(nominalType.Attributes, tupleType.Attributes...)
	//		return nominalType
	//	}
	// }
	return tupleType
}

func (n *NominalTypeVisitor) VisitTupleTypeElement(ctx *TupleTypeElementContext) interface{} {
	if typeCtx := ctx.Type_(); typeCtx != nil {
		nominalType := GetType(typeCtx)
		nominalType.Attributes = GetAttributes(ctx.Attributes())

		if identifier := ctx.DeclarationIdentifier(); identifier != nil {
			nominalType.Attributes =
				lang.SetStringAttribute(nominalType.Attributes, core.LabelAttributeName, identifier.GetText())
		}

		return nominalType
	}
	return nil
}

func (n *NominalTypeVisitor) VisitUnion(ctx *UnionContext) interface{} {
	unionType := &lang.NominalType{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
		PackageName:   "mojo.core",
		Name:          "Union",
	}

	typeCtx := ctx.AllBasicType()
	for _, t := range typeCtx {
		visitor := NewNominalTypeVisitor()
		if nominalType, ok := t.Accept(visitor).(*lang.NominalType); ok {
			unionType.GenericArguments = append(unionType.GenericArguments, nominalType)
		} else {
			// error
			return nil
		}
	}

	arguments := unionType.GenericArguments

	for i, attrCtx := range ctx.AllAttributes() {
		attrs := GetAttributes(attrCtx)
		if len(attrs) == 0 {
			continue
		}

		pos := lang.Attributes(attrs).GetEndPosition()
		for j := i + 1; j < len(arguments); j++ {
			if pos.Compare(arguments[j].StartPosition) < 0 {
				arguments[j-1].Attributes = attrs
				break
			}
		}
		lastArgument := arguments[len(arguments)-1]
		if pos.Compare(lastArgument.EndPosition) > 0 {
			lastArgument.Attributes = attrs
		}
	}
	for i, docCtx := range ctx.AllFollowingDocument() {
		doc := GetFollowingDocument(docCtx)
		if doc == nil {
			continue
		}

		pos := doc.GetEndPosition()
		for j := i + 1; j < len(arguments); j++ {
			if pos.Compare(arguments[j].StartPosition) < 0 {
				arguments[j-1].Document = doc
				break
			}
		}
		lastArgument := arguments[len(arguments)-1]
		if pos.Compare(lastArgument.EndPosition) > 0 {
			lastArgument.Document = doc
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
	identifier := GetTypeIdentifier(ctx.TypeIdentifier())
	if identifier != nil {
		return identifier
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
