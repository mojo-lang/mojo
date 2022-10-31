package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type DeclaratorVisitor struct {
	*BaseCVisitor
}

func NewDeclaratorVisitor() *DeclaratorVisitor {
	visitor := &DeclaratorVisitor{}
	return visitor
}

func (v *DeclaratorVisitor) VisitInitDeclaratorList(ctx *InitDeclaratorListContext) interface{} {
	declarators := ctx.AllInitDeclarator()
	var decls []*lang.Declaration
	for _, declarator := range declarators {
		if d := declarator.Accept(v); d != nil {
			switch dt := d.(type) {
			case *lang.ValueDecl:
				decls = append(decls, lang.NewVariableDeclarationFromValueDecl(dt))
			case *lang.FunctionDecl:
				decls = append(decls, lang.NewFunctionDeclaration(dt))
			}
		}
	}
	return decls
}

func (v *DeclaratorVisitor) VisitInitDeclarator(ctx *InitDeclaratorContext) interface{} {
	if declarator := ctx.Declarator(); declarator != nil {
		return declarator.Accept(v)
	}
	return nil
}

func (v *DeclaratorVisitor) VisitDeclarator(ctx *DeclaratorContext) interface{} {
	if dd := ctx.DirectDeclarator(); dd != nil {
		return dd.Accept(v)
	}

	// if pointer := ctx.Pointer(); pointer != nil {
	//    if attributes, ok := pointer.Accept(v).([]*lang.Attribute); ok {
	//
	//    }
	// }
	return nil
}

func (v *DeclaratorVisitor) VisitDirectDeclarator(ctx *DirectDeclaratorContext) interface{} {
	if ds := ctx.DigitSequence(); ds != nil {
		return nil
	}
	if vc := ctx.VcSpecificModifer(); vc != nil {
		return nil
	}
	if identifier := ctx.Identifier(); identifier != nil {
		return &lang.ValueDecl{Name: identifier.GetText()}
	}
	if parameters := ctx.ParameterTypeList(); parameters != nil {
		if decl, ok := parameters.Accept(NewParameterVisitor()).([]*lang.ValueDecl); ok {
			if dd := ctx.DirectDeclarator().Accept(v); dd != nil {
				switch dt := dd.(type) {
				case *lang.ValueDecl:
					return &lang.FunctionDecl{
						Name: dt.Name,
						Signature: &lang.FunctionSignature{
							Parameter: lang.NewFunctionParameters(decl...),
						},
					}
				}
			}
		}
	}
	return nil
}

func (v *DeclaratorVisitor) VisitPointer(ctx *PointerContext) interface{} {
	cnt := len(ctx.AllStar())
	var attributes []*lang.Attribute
	for i := 0; i < cnt; i++ {
		if star := ctx.Star(i); star != nil {
			attributes = append(attributes, lang.NewBoolAttribute("c.type", "*"))
		}
		if caret := ctx.Caret(i); caret != nil {
			attributes = append(attributes, lang.NewBoolAttribute("c.type", "^"))
		}
		if qualifierList := ctx.TypeQualifierList(i); qualifierList != nil {
			if attrs, ok := qualifierList.Accept(v).([]*lang.Attribute); ok {
				attributes = append(attributes, attrs...)
			}
		}
	}
	return attributes
}

func (v *DeclaratorVisitor) VisitTypeQualifierList(ctx *TypeQualifierListContext) interface{} {
	qualifiers := ctx.AllTypeQualifier()
	var attributes []*lang.Attribute
	for _, qualifier := range qualifiers {
		if attribute, ok := qualifier.Accept(NewDeclarationVisitor()).(*lang.Attribute); ok {
			attributes = append(attributes, attribute)
		}
	}
	return attributes
}

func (v *DeclaratorVisitor) VisitStructDeclaratorList(ctx *StructDeclaratorListContext) interface{} {
	return nil
}

func (v *DeclaratorVisitor) VisitStructDeclarator(ctx *StructDeclaratorContext) interface{} {
	return nil
}
