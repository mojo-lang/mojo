package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type DeclarationVisitor struct {
	*BaseCVisitor
}

func NewDeclarationVisitor() *DeclarationVisitor {
	visitor := &DeclarationVisitor{}
	return visitor
}

func (v *DeclarationVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	if specifiers := ctx.DeclarationSpecifiers(); specifiers != nil {
		var nominalType *lang.NominalType
		if typ, ok := specifiers.Accept(v).(*lang.NominalType); ok {
			nominalType = typ
		}
		txt := ctx.GetText()
		print(txt)

		if declarators := ctx.InitDeclaratorList(); declarators != nil {
			if decls, ok := declarators.Accept(NewDeclaratorVisitor()).([]*lang.Declaration); ok {
				if len(decls) > 1 {
					for _, decl := range decls {
						if variable := decl.GetVariableDecl(); variable != nil {
							variable.Type = nominalType
						}
					}

					return &lang.GroupDecl{
						Type:         lang.Identifier_KIND_VARIABLE,
						Attributes:   []*lang.Attribute{lang.NewBoolAttribute("c.declaration", "inline")},
						Declarations: decls,
					}
				} else if len(decls) == 1 {
					decl := decls[0]
					if funcDecl := decl.GetFunctionDecl(); funcDecl != nil {
						if funcDecl.Signature == nil {
							funcDecl.Signature = lang.NewFunctionSignature(nominalType)
						} else {
							funcDecl.Signature.Result = lang.NewFunctionResult(nominalType)
						}
					}
					if varDecl := decl.GetVariableDecl(); varDecl != nil {
						varDecl.Type = nominalType
					}
					if constDecl := decl.GetConstantDecl(); constDecl != nil {
						constDecl.Type = nominalType
					}

					return decl
				}
			}
		}
	}
	return nil
}

func (v *DeclarationVisitor) VisitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) interface{} {
	specifiers := ctx.AllDeclarationSpecifier()
	typ := &lang.NominalType{}
	for _, specifier := range specifiers {
		s := specifier.Accept(v)
		switch st := s.(type) {
		case *lang.Attribute:
			typ.Attributes = append(typ.Attributes, st)
		case *lang.NominalType:
			typ.Attributes = append(typ.Attributes, st.Attributes...)
			typ.Name = st.Name
			typ.TypeDeclaration = st.TypeDeclaration
		}
	}
	return typ
}

func (v *DeclarationVisitor) VisitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) interface{} {
	specifiers := ctx.AllDeclarationSpecifier()
	typ := &lang.NominalType{}
	for _, specifier := range specifiers {
		s := specifier.Accept(v)
		switch st := s.(type) {
		case *lang.Attribute:
			typ.Attributes = append(typ.Attributes, st)
		case *lang.NominalType:
			typ.Attributes = append(typ.Attributes, st.Attributes...)
			typ.Name = st.Name
			typ.TypeDeclaration = st.TypeDeclaration
		}
	}
	return typ
}

func (v *DeclarationVisitor) VisitDeclarationSpecifier(ctx *DeclarationSpecifierContext) interface{} {
	if context := ctx.StorageClassSpecifier(); context != nil {
		return context.Accept(v)
	}
	if context := ctx.TypeSpecifier(); context != nil {
		return context.Accept(v)
	}
	return nil
}

func (v *DeclarationVisitor) VisitStorageClassSpecifier(ctx *StorageClassSpecifierContext) interface{} {
	specifier := ctx.GetText()
	if len(specifier) > 0 {
		return lang.NewBoolAttribute("c.storage_class", specifier)
	}
	return nil
}

func (v *DeclarationVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	if node := ctx.GetText(); len(node) > 0 {
		switch node {
		case "void", "char", "short", "int", "long", "float", "double", "signed", "unsigned", "_Bool", "_Complex", "__m128", "__m128d", "__m128i":
			return &lang.NominalType{Name: node}
		case "__extension__(__m128)":
			return &lang.NominalType{Name: "__m128", Attributes: []*lang.Attribute{lang.NewBoolAttribute("c.type", "__extension__")}}
		case "__extension__(__m128d)":
			return &lang.NominalType{Name: "__m128d", Attributes: []*lang.Attribute{lang.NewBoolAttribute("c.type", "__extension__")}}
		case "__extension__(__m128i)":
			return &lang.NominalType{Name: "__m128i", Attributes: []*lang.Attribute{lang.NewBoolAttribute("c.type", "__extension__")}}
		}
	}

	if atomic := ctx.AtomicTypeSpecifier(); atomic != nil {
		return atomic.Accept(v)
	}

	if typedefName := ctx.TypedefName(); typedefName != nil {
		return &lang.NominalType{Name: typedefName.GetText()}
	}

	return nil
}

func (v *DeclarationVisitor) VisitAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) interface{} {
	if typeName := ctx.TypeName(); typeName != nil {
		if typ, ok := typeName.Accept(NewTypeNameVisitor()).(*lang.NominalType); ok {
			typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("c.type", "_Atomic"))
			return typ
		}
	}
	return nil
}

func (v *DeclarationVisitor) VisitTypeQualifier(ctx *TypeQualifierContext) interface{} {
	return lang.NewBoolAttribute("c.type", ctx.GetText())
}

func (v *DeclarationVisitor) VisitFunctionSpecifier(ctx *FunctionSpecifierContext) interface{} {
	if node := ctx.GetText(); len(node) > 0 {
		switch node {
		case "inline", "_Noreturn", "__inline__", "__stdcall":
			return lang.NewBoolAttribute("c.function", node)
		}
	}

	if identifier := ctx.Identifier(); identifier != nil {
		return lang.NewStringAttribute("c.function", "__declspec", identifier.GetText())
	}
	return nil
}

func (v *DeclarationVisitor) VisitAlignmentSpecifier(ctx *AlignmentSpecifierContext) interface{} {
	if typeName := ctx.TypeName(); typeName != nil {
		if typ, ok := typeName.Accept(NewTypeNameVisitor()).(*lang.NominalType); ok {
			return lang.NewStringAttribute("c.alignment", "_Alignas", typ.Name)
		}
	}
	return nil
}

func (v *DeclarationVisitor) VisitEnumSpecifier(ctx *EnumSpecifierContext) interface{} {
	return nil
}

func (v *DeclarationVisitor) VisitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) interface{} {
	identifier := ctx.Identifier().GetText()

	var decls []*lang.Declaration
	if declarationList := ctx.StructDeclarationList(); declarationList != nil {
		if ds, ok := declarationList.Accept(v).([]*lang.Declaration); ok {
			decls = ds
		}
	}

	if structOrUnion := ctx.StructOrUnion(); structOrUnion != nil {
		su := structOrUnion.GetText()
		if su == "struct" {
			var fields []*lang.ValueDecl
			var groups []*lang.GroupDecl

			for _, decl := range decls {
				if d := decl.GetVariableDecl(); d != nil {
					fields = append(fields, d.ToValueDecl())
				} else if group := decl.GetGroupDecl(); group != nil {
					groups = append(groups, group)
				}
			}

			structDecl := &lang.StructDecl{
				Name: identifier,
				Type: &lang.StructType{
					Fields: fields,
					Groups: groups,
				},
			}

			return structDecl
		} else if su == "union" {

		}
	}

	return nil
}

// VisitStructDeclarationList
// return []*lang.Declaration // ValueDecl | GroupDecl
func (v *DeclarationVisitor) VisitStructDeclarationList(ctx *StructDeclarationListContext) interface{} {
	return nil
}

func (v *DeclarationVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return nil
}

func (v *DeclarationVisitor) VisitSpecifierQualifierList(ctx *SpecifierQualifierListContext) interface{} {
	return nil
}
