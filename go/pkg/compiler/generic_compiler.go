package compiler

import (
	"errors"
	path2 "path"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

// compile the Generic Type to Nominal type except the Array, Map, Tuple, Union, Intersection
type GenericCompiler struct {
	Options  map[string]interface{}
	NewFiles []*lang.SourceFile
}

func NewGenericCompiler(options map[string]interface{}) *GenericCompiler {
	compiler := &GenericCompiler{
		Options: options,
	}

	if compiler.Options == nil {
		compiler.Options = make(map[string]interface{})
	}

	return compiler
}

func addIdentifiers(dest map[string]*lang.Identifier, identifiers []*lang.Identifier) {
	for _, identifier := range identifiers {
		dest[identifier.FullName] = identifier
	}
}

func (c *GenericCompiler) Compile(ctx *context.Context, pkg *lang.Package) error {
	c.NewFiles = nil
	for _, sourceFile := range pkg.SourceFiles {
		ctx.Open(sourceFile)
		identifiers := make(map[string]*lang.Identifier)
		for i, statement := range sourceFile.Statements {
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
					typeAliasDecl := decl.GetTypeAliasDecl()
					if typeAliasDecl != nil {
						ids, structDecl, err := c.CompileTypeAlias(ctx, typeAliasDecl)
						if err != nil {
							ctx.Close()
							return err
						}
						addIdentifiers(identifiers, ids)
						if structDecl != nil {
							sourceFile.Statements[i] = lang.NewStructDeclStatement(structDecl)
						}
					}
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
			if id.SourceFileName == sourceFile.FullName {
				continue
			}
			sourceFile.ResolvedIdentifiers = append(sourceFile.ResolvedIdentifiers, id)
		}

		ctx.Close()
	}

	for _, file := range c.NewFiles {
		pkg.SourceFiles = append(pkg.SourceFiles, file)
	}
	return nil
}

func (c *GenericCompiler) CompileStruct(ctx *context.Context, decl *lang.StructDecl) ([]*lang.Identifier, error) {
	// generic struct will not be compiled
	if len(decl.GenericParameters) > 0 {
		return decl.ResolvedIdentifiers, nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	identifiers := make(map[string]*lang.Identifier)
	for _, structDecl := range decl.StructDecls {
		/// TODO
		/// if the struct only have an inherit and the has the same name, should be replace the current struct
		ids, err := c.CompileStruct(ctx, structDecl)
		if err != nil {
			return nil, err
		}
		addIdentifiers(identifiers, ids)
	}

	typeAliasDecls := make([]*lang.TypeAliasDecl, 0, len(decl.TypeAliasDecls))
	for _, typeAliasDecl := range decl.TypeAliasDecls {
		ids, structDecl, err := c.CompileTypeAlias(ctx, typeAliasDecl)
		if err != nil {
			return nil, err
		}

		addIdentifiers(identifiers, ids)
		if structDecl != nil {
			structDecl.EnclosingType = &lang.NominalType{
				PackageName:     decl.PackageName,
				Name:            decl.Name,
				TypeDeclaration: lang.NewStructTypeDeclaration(decl),
			}
			decl.StructDecls = append(decl.StructDecls, structDecl)
		} else {
			typeAliasDecls = append(typeAliasDecls, typeAliasDecl)
		}
	}
	decl.TypeAliasDecls = typeAliasDecls

	if decl.Type != nil {
		identifierIndex := make(map[string]*lang.Identifier)
		for _, id := range decl.ResolvedIdentifiers {
			identifierIndex[id.FullName] = id
		}

		compileType := func(nominalType *lang.NominalType) (*lang.NominalType, error) {
			id, err := c.compileGenericNominalType(ctx, nominalType)
			if err != nil {
				return nil, err
			}

			if id != nil {
				identifiers[id.FullName] = id
				return &lang.NominalType{
					PackageName:     id.PackageName,
					Name:            id.Name,
					TypeDeclaration: lang.NewTypeDeclarationFromDeclaration(id.Declaration),
					Attributes:      nominalType.Attributes,
				}, nil
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
				if exceptionalTypes[inherit.GetFullName()] {
					for i, argument := range inherit.GenericArguments {
						id, err := c.compileGenericNominalType(ctx, argument)
						if err != nil {
							return nil, err
						}
						if id != nil {
							inherit.GenericArguments[i] = &lang.NominalType{
								PackageName:     id.PackageName,
								Name:            id.Name,
								TypeDeclaration: lang.NewTypeDeclarationFromDeclaration(id.Declaration),
								Attributes:      argument.Attributes,
							}
							identifiers[id.FullName] = id
						}
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
				if exceptionalTypes[field.Type.GetFullName()] {
					for i, argument := range field.Type.GenericArguments {
						id, err := c.compileGenericNominalType(ctx, argument)
						if err != nil {
							return nil, err
						}
						if id != nil {
							field.Type.GenericArguments[i] = &lang.NominalType{
								PackageName:     id.PackageName,
								Name:            id.Name,
								TypeDeclaration: lang.NewTypeDeclarationFromDeclaration(id.Declaration),
								Attributes:      argument.Attributes,
							}
							identifiers[id.FullName] = id
						}
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

func (c *GenericCompiler) CompileTypeAlias(ctx *context.Context, decl *lang.TypeAliasDecl) ([]*lang.Identifier, *lang.StructDecl, error) {
	// generic type alias will not be compiled
	if len(decl.GenericParameters) > 0 {
		return decl.ResolvedIdentifiers, nil, nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	decl.Type.Attributes = append(decl.Type.Attributes, decl.Attributes...)
	id, err := c.compileGenericNominalType(ctx, decl.Type)
	if err != nil || id == nil {
		return decl.ResolvedIdentifiers, nil, err
	}

	if id.Name == decl.Name {
		if structDecl := id.Declaration.GetStructDecl(); structDecl != nil {
			return structDecl.ResolvedIdentifiers, structDecl, nil
		} else if aliasDecl := id.Declaration.GetTypeAliasDecl(); aliasDecl != nil {
			decl.Type = aliasDecl.Type
			decl.ResolvedIdentifiers = aliasDecl.ResolvedIdentifiers
			return aliasDecl.ResolvedIdentifiers, nil, err
		}
	} else {
		decl.Type = &lang.NominalType{
			PackageName:     id.PackageName,
			Name:            id.Name,
			Attributes:      decl.Type.Attributes,
			TypeDeclaration: lang.NewTypeDeclarationFromDeclaration(id.Declaration),
		}

		return []*lang.Identifier{id}, nil, nil
	}

	return nil, nil, err
}

/// TODO support configure
var exceptionalTypes = map[string]bool{
	"mojo.core.Array":        true,
	"mojo.core.Map":          true,
	"mojo.core.Tuple":        true,
	"mojo.core.Union":        true,
	"mojo.core.Intersection": true,
}

func (c *GenericCompiler) compileGenericNominalType(ctx *context.Context, nominalType *lang.NominalType) (*lang.Identifier, error) {
	if len(nominalType.GenericArguments) == 0 {
		return nil, nil
	}

	if exceptionalTypes[nominalType.GetFullName()] {
		return nil, nil
	}

	// skip it, if all the arguments can't be instantiated
	for _, argument := range nominalType.GenericArguments {
		if argument.IsGenericParameter() {
			return nil, nil
		}
	}

	switch nominalType.TypeDeclaration.TypeDeclaration.(type) {
	case *lang.TypeDeclaration_TypeAliasDecl:
		return c.compileGenericTypeAlias(ctx, nominalType)
	case *lang.TypeDeclaration_StructDecl:
		return c.compileGenericStructType(ctx, nominalType)
	}

	return nil, nil
}

/// TODO support configure
// typename: typename template
// "Cached": "Cached{T}"
func newTypeName(t *lang.NominalType) string {
	name := t.Name
	if strings.HasSuffix(name, "ed") || strings.HasSuffix(name, "able") {
		for _, argument := range t.GenericArguments {
			name += argument.Name
		}
	} else {
		argumentName := ""
		for _, argument := range t.GenericArguments {
			argumentName += argument.Name
		}
		name = argumentName + name
	}
	return name
}

func (c *GenericCompiler) isCircleDependency(ctx *context.Context, nominalType *lang.NominalType) bool {
	values := ctx.OrderlyValues()
	for _, value := range values {
		scope := lang.GetScope(value)
		if scope != nil {
			for _, id := range scope.Identifiers {
				for _, argument := range nominalType.GenericArguments {
					if id.Name == argument.Name {
						return true
					}
				}
			}
		}
	}
	return false
}

func (c *GenericCompiler) instantiateTypeAlias(ctx *context.Context, nominalType *lang.NominalType) (*lang.NominalType, error) {
	aliasDecl := nominalType.TypeDeclaration.GetTypeAliasDecl()

	genericIndex := make(map[string]*lang.NominalType)
	for i, parameter := range aliasDecl.GenericParameters {
		genericIndex[parameter.Name] = nominalType.GenericArguments[i]
	}
	aliasType := &lang.NominalType{
		StartPosition:   aliasDecl.Type.StartPosition,
		EndPosition:     aliasDecl.Type.EndPosition,
		PackageName:     aliasDecl.Type.PackageName,
		Name:            aliasDecl.Type.Name,
		TypeDeclaration: aliasDecl.Type.TypeDeclaration,
		Attributes:      aliasDecl.Type.Attributes,
		EnclosingType:   aliasDecl.Type.EnclosingType,
	}
	for _, argument := range aliasDecl.Type.GenericArguments {
		generic := genericIndex[argument.Name]
		if generic != nil {
			for _, attribute := range argument.Attributes {
				generic.Attributes = append(generic.Attributes, attribute)
			}
			aliasType.GenericArguments = append(aliasType.GenericArguments, generic)
		} else {
			aliasType.GenericArguments = append(aliasType.GenericArguments, argument)
		}
	}

	for _, attribute := range nominalType.Attributes {
		aliasType.Attributes = append(aliasType.Attributes, attribute)
	}
	for _, attribute := range aliasDecl.Attributes {
		aliasType.Attributes = append(aliasType.Attributes, attribute)
	}

	// set or transmit the original type alias name to the end
	aliasName, err := lang.GetStringAttribute(nominalType.Attributes, lang.OriginalTypeAliasName)
	if err != nil {
		aliasName = nominalType.Name
	}
	aliasType.Attributes = lang.SetStringAttribute(aliasType.Attributes, lang.OriginalTypeAliasName, aliasName)

	if aliasType.IsTypeAlias() {
		return c.instantiateTypeAlias(ctx, aliasType)
	}

	return aliasType, nil
}

func (c *GenericCompiler) compileGenericTypeAlias(ctx *context.Context, nominalType *lang.NominalType) (*lang.Identifier, error) {
	aliasDecl := nominalType.TypeDeclaration.GetTypeAliasDecl()
	if aliasDecl == nil {
		return nil, nil
	}

	if len(nominalType.GenericArguments) != len(aliasDecl.GenericParameters) {
		return nil, errors.New("the generic arguments is NOT equals to generic parameters")
	}

	// new Type Name
	aliasType, err := c.instantiateTypeAlias(ctx, nominalType)
	if err != nil {
		return nil, err
	}

	// generate the new type name
	typeName := newTypeName(nominalType)
	ctxFile := ctx.GetSourceFile()

	// create the new file or just the current file
	var file *lang.SourceFile
	var scope *lang.Scope

	nameConflict := false
	for _, value := range ctx.OrderlyValues() {
		if scope = lang.GetScope(value); scope != nil {
			for name, i := range scope.Identifiers {
				if name == typeName {
					if i.Declaration.GetTypeAliasDecl() == nil { // only process the duplication with type alias
						///TODO logs error here
						return nil, nil
					}
					nameConflict = true
					file = ctxFile
					break
				}
			}
		}

		if nameConflict {
			break
		}

		if _, ok := value.(*lang.SourceFile); ok {
			break
		}
	}

	if file == nil {
		fileName := strcase.ToSnake(typeName) + ".mojo"
		fileFullName := path2.Join(path2.Dir(ctxFile.FullName), fileName)
		file = ctx.GetPackage().GetSourceFile(fileFullName)
		if file == nil {
			for _, f := range c.NewFiles {
				if f.FullName == fileFullName {
					if f.Scope != nil && f.Scope.Identifiers != nil && f.Scope.Identifiers[typeName] != nil {
						return f.Scope.Identifiers[typeName], nil
					}
					file = f
					break
				}
			}

			if file == nil {
				file = &lang.SourceFile{
					Name:        fileName,
					FullName:    fileFullName,
					PackageName: ctxFile.PackageName,
				}
				c.NewFiles = append(c.NewFiles, file)
			}
		}

		if file.Scope != nil {
			for name, i := range file.Scope.Identifiers {
				if name == typeName {
					if i.Declaration.GetTypeAliasDecl() == nil { // only process the duplication with type alias
						///TODO logs error here
						return nil, nil
					}
					nameConflict = true
					break
				}
			}
		}
	}

	compiledDecl := &lang.TypeAliasDecl{
		Document:       aliasDecl.Document,
		PackageName:    file.PackageName,
		SourceFileName: file.FullName,
		Implicit:       true,
		Name:           typeName,
		Type:           aliasType,
	}

	findIdentifier := func(t *lang.NominalType) *lang.Identifier {
		for _, id := range ctxFile.ResolvedIdentifiers {
			if id.FullName == t.GetFullName() {
				return id
			}
		}
		for name, id := range ctxFile.Scope.GetIdentifiers() {
			if name == t.Name {
				return id
			}
		}

		return nil
	}
	var resolvedIdentifiers []*lang.Identifier
	for _, argument := range nominalType.GenericArguments {
		if id := findIdentifier(argument); id != nil {
			resolvedIdentifiers = append(resolvedIdentifiers, id)
		}

		//resolvedIdentifiers = append(resolvedIdentifiers, &lang.Identifier{
		//	Package:     argument.Package,
		//	SourceFile:  argument.TypeDeclaration.GetSourceFileName(),
		//	Kind:        "",
		//	Name:        argument.Name,
		//	FullName:    argument.GetFullName(),
		//	Declaration: lang.NewDeclarationFromTypeDeclaration(argument.TypeDeclaration),
		//})
	}
	for _, id := range aliasDecl.ResolvedIdentifiers {
		if id.Kind == lang.Identifier_KIND_GENERIC_PARAMETER {
			continue
		}

		resolvedIdentifiers = append(resolvedIdentifiers, id)
	}
	compiledDecl.ResolvedIdentifiers = resolvedIdentifiers

	id := &lang.Identifier{
		PackageName:    compiledDecl.PackageName,
		Name:           compiledDecl.Name,
		FullName:       lang.GetFullName(compiledDecl.PackageName, lang.GetEnclosingNames(compiledDecl.EnclosingType), compiledDecl.Name),
		SourceFileName: file.FullName,
		Declaration:    lang.NewTypeAliasDeclaration(compiledDecl),
	}

	// update the identifiers in the current scope
	if nameConflict {
		// update the identifier in scope, may be in the file or the struct scope
		identifier := scope.Identifiers[aliasDecl.Name]
		if identifier == nil {
			identifier = scope.Identifiers[typeName]
		}

		if identifier != nil {
			identifier.Declaration = lang.NewTypeAliasDeclaration(compiledDecl)
			id = identifier
		}
	} else {
		statement := lang.NewTypeAliasDeclStatement(compiledDecl)
		file.Statements = append(file.Statements, statement)

		if file.Scope == nil || file.Scope.Identifiers == nil {
			file.Scope = &lang.Scope{
				Identifiers: map[string]*lang.Identifier{typeName: id},
			}
		} else {
			file.Scope.Identifiers[typeName] = id
		}
	}

	// update the resolved identifiers in file
	if file.ResolvedIdentifiers == nil {
		file.ResolvedIdentifiers = compiledDecl.ResolvedIdentifiers
	} else {
		file.ResolvedIdentifiers = append(file.ResolvedIdentifiers, compiledDecl.ResolvedIdentifiers...)
		file.ResolvedIdentifiers = lang.MergeDependencies(file.ResolvedIdentifiers)
	}

	return id, nil
}

func (c *GenericCompiler) compileGenericStructType(ctx *context.Context, nominalType *lang.NominalType) (*lang.Identifier, error) {
	structDecl := nominalType.TypeDeclaration.GetStructDecl()
	if structDecl == nil {
		return nil, nil
	}

	if len(structDecl.GenericParameters) != len(nominalType.GenericArguments) {
		return nil, errors.New("the count of generic argument not equal to parameter")
	}
	if structDecl.Type == nil {
		return nil, errors.New("opacity struct type declaration can't be compiled")
	}

	var compiledDecl *lang.StructDecl

	// generate the new type name
	typeName := newTypeName(nominalType)
	ctxFile := ctx.GetSourceFile()

	// create the new file or just the current file
	var file *lang.SourceFile
	fileName := strcase.ToSnake(typeName) + ".mojo"
	fileFullName := path2.Join(path2.Dir(ctxFile.FullName), fileName)
	file = ctx.GetPackage().GetSourceFile(fileFullName)
	if file == nil {
		for _, f := range c.NewFiles {
			if f.FullName == fileFullName {
				if f.Scope != nil && f.Scope.Identifiers != nil && f.Scope.Identifiers[typeName] != nil {
					return f.Scope.Identifiers[typeName], nil
				}
				file = f
				break
			}
		}

		if file == nil {
			file = &lang.SourceFile{
				Name:        fileName,
				FullName:    fileFullName,
				PackageName: ctxFile.PackageName,
			}
			c.NewFiles = append(c.NewFiles, file)
		}
	}

	nameConflict := false
	var scope *lang.Scope
	for i := 1; i < len(ctx.Values)-1; i++ {
		value := ctx.PreviousValue(i)
		switch t := value.(type) {
		case *lang.StructDecl:
			scope = t.Scope
			break
		case *lang.InterfaceDecl:
			scope = t.Scope
			break
		case *lang.SourceFile:
			scope = t.Scope
			break
		}
	}

	if scope != nil {
		for name, i := range scope.Identifiers {
			if name == typeName {
				if i.Declaration.GetTypeAliasDecl() == nil { // only process the duplication with type alias
					return nil, nil
				}
				nameConflict = true
				break
			}
		}
	}

	compiledDecl = &lang.StructDecl{
		Document:       structDecl.Document,
		PackageName:    file.PackageName,
		SourceFileName: file.FullName,
		Implicit:       true,
		Name:           typeName,
		Type:           &lang.StructType{},
	}

	replaces := make(map[string]*lang.NominalType)
	for i, parameter := range structDecl.GenericParameters {
		replaces[parameter.Name] = nominalType.GenericArguments[i]
	}

	for _, f := range structDecl.Type.Fields {
		field := &lang.ValueDecl{
			Document:     f.Document,
			Implicit:     true,
			Name:         f.Name,
			Attributes:   f.Attributes,
			Type:         nil,
			InitialValue: proto.Clone(f.InitialValue).(*lang.Expression),
		}
		replaced := replaces[f.Type.Name]
		if replaced != nil {
			field.Type = &lang.NominalType{
				Name:            replaced.Name,
				PackageName:     replaced.PackageName,
				TypeDeclaration: replaced.TypeDeclaration,
				Attributes:      f.Type.Attributes,
				EnclosingType:   replaced.EnclosingType,
			}
		} else {
			field.Type = &lang.NominalType{
				Name:            f.Type.Name,
				PackageName:     f.Type.PackageName,
				TypeDeclaration: f.Type.TypeDeclaration,
				Attributes:      f.Type.Attributes,
				EnclosingType:   f.Type.EnclosingType,
			}
			for i, argument := range f.Type.GenericArguments {
				replaced = replaces[argument.Name]
				if replaced != nil {
					field.Type.GenericArguments = append(field.Type.GenericArguments, &lang.NominalType{
						Name:            replaced.Name,
						PackageName:     replaced.PackageName,
						TypeDeclaration: replaced.TypeDeclaration,
						Attributes:      field.Type.GenericArguments[i].Attributes,
						EnclosingType:   replaced.EnclosingType,
					})
				} else {
					field.Type.GenericArguments = append(field.Type.GenericArguments, argument)
				}
			}
			/// TODO 循环编译内嵌的字段类型
		}

		compiledDecl.Type.Fields = append(compiledDecl.Type.Fields, field)
	}

	///TODO add inherits to compile

	var resolvedIdentifiers []*lang.Identifier
	for _, id := range structDecl.ResolvedIdentifiers {
		if id.Kind == lang.Identifier_KIND_GENERIC_PARAMETER {
			continue
		}
		resolvedIdentifiers = append(resolvedIdentifiers, id)
	}
	for _, argument := range nominalType.GenericArguments {
		resolvedIdentifiers = append(resolvedIdentifiers, &lang.Identifier{
			PackageName:        argument.PackageName,
			SourceFileName:     argument.TypeDeclaration.GetSourceFileName(),
			EnclosingTypeNames: nil,
			Name:               argument.Name,
			FullName:           argument.GetFullName(),
			Declaration:        lang.NewDeclarationFromTypeDeclaration(argument.TypeDeclaration),
		})
	}
	compiledDecl.ResolvedIdentifiers = lang.MergeDependencies(resolvedIdentifiers)

	id := &lang.Identifier{
		PackageName:    compiledDecl.PackageName,
		Name:           compiledDecl.Name,
		FullName:       lang.GetFullName(compiledDecl.PackageName, lang.GetEnclosingNames(compiledDecl.EnclosingType), compiledDecl.Name),
		SourceFileName: file.FullName,
		Declaration:    lang.NewStructDeclaration(compiledDecl),
	}

	// update the identifiers in the current scope
	if nameConflict {
		// update the identifier in scope, may be in the file or the struct scope
		identifier := scope.Identifiers[typeName]
		if identifier != nil {
			identifier.Declaration = lang.NewStructDeclaration(compiledDecl)
			id = identifier
		}
	} else {
		statement := lang.NewStructDeclStatement(compiledDecl)
		file.Statements = append(file.Statements, statement)

		if file.Scope == nil || file.Scope.Identifiers == nil {
			file.Scope = &lang.Scope{
				Identifiers: map[string]*lang.Identifier{typeName: id},
			}
		} else {
			file.Scope.Identifiers[typeName] = id
		}
	}

	// update the resolved identifiers in file
	if file.ResolvedIdentifiers == nil {
		file.ResolvedIdentifiers = compiledDecl.ResolvedIdentifiers
	} else {
		file.ResolvedIdentifiers = append(file.ResolvedIdentifiers, compiledDecl.ResolvedIdentifiers...)
		file.ResolvedIdentifiers = lang.MergeDependencies(file.ResolvedIdentifiers)
	}

	return id, nil
}
