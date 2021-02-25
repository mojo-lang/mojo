package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

// nominal type alias process except the Generic Type
type TypeAliasCompiler struct {
	Options map[string]interface{}
}

func (c *TypeAliasCompiler) Compile(ctx *context.Context, pkg *lang.Package) error {
	for _, sourceFile := range pkg.SourceFiles {
		ctx.Open(sourceFile)
		identifiers := make(map[string]*lang.Identifier)
		for _, statement := range sourceFile.Statements {
			if decl := statement.GetDeclaration(); decl != nil {
				switch decl.Declaration.(type) {
				case *lang.Declaration_StructDecl:
					structDecl := decl.GetStructDecl()
					if structDecl != nil {
						ids, err := c.CompileStruct(ctx, structDecl)
						if err != nil {
							ctx.Close()
							return err
						}
						addIdentifiers(identifiers, ids)
					}
				case *lang.Declaration_TypeAliasDecl:
					addIdentifiers(identifiers, decl.GetTypeAliasDecl().ResolvedIdentifiers)
				case *lang.Declaration_EnumDecl:
					addIdentifiers(identifiers, decl.GetEnumDecl().ResolvedIdentifiers)
				case *lang.Declaration_InterfaceDecl:
					addIdentifiers(identifiers, decl.GetInterfaceDecl().ResolvedIdentifiers)
					//case *lang.Declaration_AttributeDecl:
					//	addIdentifiers(identifiers, decl.GetAttributeDecl().ResolvedIdentifiers)
					//case *lang.Declaration_ConstantDecl:
					//	addIdentifiers(identifiers, decl.GetConstantDecl().ResolvedIdentifiers)
					//case *lang.Declaration_VariableDecl:
					//	addIdentifiers(identifiers, decl.GetVariableDecl().ResolvedIdentifiers)
					//case *lang.Declaration_FunctionDecl:
					//	addIdentifiers(identifiers, decl.GetFunctionDecl().ResolvedIdentifiers)
				}
			}
		}

		sourceFile.ResolvedIdentifiers = nil
		for _, id := range identifiers {
			if id.SourceFile == sourceFile.FullName {
				continue
			}
			sourceFile.ResolvedIdentifiers = append(sourceFile.ResolvedIdentifiers, id)
		}

		ctx.Close()
	}
	return nil
}

func (c *TypeAliasCompiler) CompileStruct(ctx *context.Context, decl *lang.StructDecl) ([]*lang.Identifier, error) {
	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	identifiers := make(map[string]*lang.Identifier)
	for _, structDecl := range decl.StructDecls {
		ids, err := c.CompileStruct(ctx, structDecl)
		if err != nil {
			return nil, err
		}
		addIdentifiers(identifiers, ids)
	}

	for _, typeAliasDecl := range decl.TypeAliasDecls {
		addIdentifiers(identifiers, typeAliasDecl.ResolvedIdentifiers)
	}

	if decl.Type != nil {
		identifierIndex := make(map[string]*lang.Identifier)
		for _, id := range decl.ResolvedIdentifiers {
			identifierIndex[id.FullName] = id
		}

		compileType := func(nominalType *lang.NominalType) (*lang.NominalType, error) {
			ids, t, err := c.compileNominalType(ctx, nominalType)
			if err != nil {
				return nil, err
			} else if t != nil {
				for _, id := range ids {
					identifiers[id.FullName] = id
				}
				return t, nil
			} else {
				fullName := nominalType.GetFullName()
				identifier := identifierIndex[fullName]
				if identifier != nil {
					identifiers[fullName] = identifier
				}

				for _, argument := range nominalType.GenericArguments {
					fullName = argument.GetFullName()
					identifier = identifierIndex[fullName]
					if identifier != nil {
						identifiers[fullName] = identifier
					}
				}
				return nil, nil
			}
		}

		var inherits []*lang.NominalType
		for _, inherit := range decl.Type.Inherits {
			t, err := compileType(inherit)
			if err != nil {
				return nil, err
			}
			if t != nil {
				inherits = append(inherits, t)
			} else {
				for i, argument := range inherit.GenericArguments {
					t, err = compileType(argument)
					if err != nil {
						return nil, err
					}
					if t != nil {
						inherit.GenericArguments[i] = t
					}
				}
				inherits = append(inherits, inherit)
			}
		}
		decl.Type.Inherits = inherits

		for _, field := range decl.Type.Fields {
			t, err := compileType(field.Type)
			if err != nil {
				return nil, err
			}
			if t != nil {
				field.Type = t
			} else {
				for i, argument := range field.Type.GenericArguments {
					t, err = compileType(argument)
					if err != nil {
						return nil, err
					}
					if t != nil {
						field.Type.GenericArguments[i] = t
					}
				}
			}
		}
	}

	decl.ResolvedIdentifiers = nil
	for _, id := range identifiers {
		decl.ResolvedIdentifiers = append(decl.ResolvedIdentifiers, id)
	}

	return decl.ResolvedIdentifiers, nil
}

func (c *TypeAliasCompiler) compileNominalType(ctx *context.Context, nominalType *lang.NominalType) ([]*lang.Identifier, *lang.NominalType, error) {
	_ = ctx
	if nominalType == nil || nominalType.TypeDeclaration == nil {
		return nil, nil, nil
	}

	aliasDecl := nominalType.TypeDeclaration.GetTypeAliasDecl()
	if aliasDecl == nil {
		return nil, nil, nil
	}

	if len(nominalType.GenericArguments) > 0 {
		return nil, nil, nil
	}

	if len(aliasDecl.Type.GenericArguments) > 0 {
		return nil, nil, nil
	}

	aliasType := &lang.NominalType{
		StartPosition:   aliasDecl.Type.StartPosition,
		EndPosition:     aliasDecl.Type.EndPosition,
		Package:         aliasDecl.Type.Package,
		Name:            aliasDecl.Type.Name,
		TypeDeclaration: aliasDecl.Type.TypeDeclaration,
		Attributes:      aliasDecl.Type.Attributes,
		EnclosingType:   aliasDecl.Type.EnclosingType,
	}

	for _, attribute := range nominalType.Attributes {
		aliasType.Attributes = append(aliasType.Attributes, attribute)
	}

	// set or transmit the original type alias name to the end
	aliasName, err := lang.GetStringAttribute(nominalType.Attributes, lang.OriginalTypeAliasName)
	if err != nil {
		aliasName = nominalType.Name
	}
	aliasType.Attributes = lang.SetStringAttribute(aliasType.Attributes, lang.OriginalTypeAliasName, aliasName)
	aliasType.Attributes = append(aliasType.Attributes, aliasDecl.Attributes...)

	if aliasType.IsTypeAlias() {
		ids, nt, err := c.compileNominalType(ctx, aliasType)
		if err != nil || nt != nil {
			return ids, nt, err
		}
	}

	return aliasDecl.ResolvedIdentifiers, aliasType, nil
}
