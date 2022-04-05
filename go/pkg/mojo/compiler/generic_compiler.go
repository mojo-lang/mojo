package compiler

import (
    "errors"
    "fmt"

    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
    "google.golang.org/protobuf/proto"
    path2 "path"
)

func init() {
    plugin.RegisterPlugin(NewGenericCompiler(nil))
}

const genericName = "compiler.generic"

// GenericCompiler compile the Generic Type to Nominal type except the Array, Map, Tuple, Union, Intersection
type GenericCompiler struct {
    plugin.BasicPlugin

    core.Options
    NewFiles []*lang.SourceFile
}

type IdentifierIndex map[string]*lang.Identifier

func NewGenericCompiler(options core.Options) *GenericCompiler {
    compiler := &GenericCompiler{
        Options: options,
        BasicPlugin: plugin.BasicPlugin{
            Name:          genericName,
            Group:         "compiler",
            GroupPriority: 10,
            Priority:      3,
            Creator: func(options core.Options) plugin.Plugin {
                return NewGenericCompiler(options)
            },
        },
    }
    if compiler.Options == nil {
        compiler.Options = make(core.Options)
    }

    return compiler
}

func addIdentifiers(dest map[string]*lang.Identifier, identifiers []*lang.Identifier) {
    for _, identifier := range identifiers {
        dest[identifier.FullName] = identifier
    }
}

func (c *GenericCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
    logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName)

    thisCtx := context.WithType(context.WithOptions(ctx, c.Options), pkg)

    c.NewFiles = nil
    for _, sourceFile := range pkg.SourceFiles {
        fileCtx := context.WithType(thisCtx, sourceFile)
        identifiers := make(map[string]*lang.Identifier)
        for i, statement := range sourceFile.Statements {
            if decl := statement.GetDeclaration(); decl != nil {
                switch decl.Declaration.(type) {
                case *lang.Declaration_StructDecl:
                    structDecl := decl.GetStructDecl()
                    if structDecl != nil {
                        ids, err := c.CompileStruct(fileCtx, structDecl)
                        if err != nil {
                            return err
                        }
                        addIdentifiers(identifiers, ids)
                    }
                case *lang.Declaration_TypeAliasDecl:
                    typeAliasDecl := decl.GetTypeAliasDecl()
                    if typeAliasDecl != nil {
                        ids, structDecl, err := c.CompileTypeAlias(fileCtx, typeAliasDecl)
                        if err != nil {
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
    }

    for _, file := range c.NewFiles {
        pkg.SourceFiles = append(pkg.SourceFiles, file)
    }

    for _, child := range pkg.Children {
        if err := c.CompilePackage(thisCtx, child); err != nil {
            return err
        }
    }
    return nil
}

func (c *GenericCompiler) CompileStruct(ctx context.Context, decl *lang.StructDecl) ([]*lang.Identifier, error) {
    // generic struct will not be compiled
    if len(decl.GenericParameters) > 0 {
        return decl.ResolvedIdentifiers, nil
    }

    thisCtx := context.WithType(ctx, decl)

    identifiers := make(map[string]*lang.Identifier)
    for _, structDecl := range decl.StructDecls {
        ids, err := c.CompileStruct(thisCtx, structDecl)
        if err != nil {
            return nil, err
        }
        addIdentifiers(identifiers, ids)
    }

    typeAliasDecls := make([]*lang.TypeAliasDecl, 0, len(decl.TypeAliasDecls))
    for _, typeAliasDecl := range decl.TypeAliasDecls {
        ids, structDecl, err := c.CompileTypeAlias(thisCtx, typeAliasDecl)
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
        for _, id := range decl.ResolvedIdentifiers {
            if structDecl := id.Declaration.GetStructDecl(); structDecl.IsGeneric() && structDecl.IsOpacity() {
                continue
            }
            identifiers[id.FullName] = id
        }

        var inherits []*lang.NominalType
        for _, inherit := range decl.Type.Inherits {
            t, err := c.compileNominalType(thisCtx, inherit, identifiers)
            if err != nil {
                return nil, err
            }
            if t != nil {
                inherits = append(inherits, t)
            } else {
                inherits = append(inherits, inherit)
            }
        }
        decl.Type.Inherits = inherits

        for _, field := range decl.Type.Fields {
            fieldCtx := context.WithType(thisCtx, field)
            t, err := c.compileNominalType(fieldCtx, field.Type, identifiers)
            if err != nil {
                return nil, err
            }
            if t != nil {
                field.Type = t
            }
        }
    }

    decl.ResolvedIdentifiers = nil
    for _, id := range identifiers {
        decl.ResolvedIdentifiers = append(decl.ResolvedIdentifiers, id)
    }

    return decl.ResolvedIdentifiers, nil
}

func (c *GenericCompiler) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) ([]*lang.Identifier, *lang.StructDecl, error) {
    // generic type alias will not be compiled
    if len(decl.GenericParameters) > 0 {
        return decl.ResolvedIdentifiers, nil, nil
    }
    thisCtx := context.WithType(ctx, decl)

    identifierIndex := make(IdentifierIndex)
    decl.Type.Attributes = append(decl.Type.Attributes, decl.Attributes...)
    decl.Type.Attributes = lang.SetStringAttribute(decl.Type.Attributes, lang.OriginalTypeAliasName, decl.Name)

    newType, err := c.compileNominalType(thisCtx, decl.Type, identifierIndex)
    if err != nil || newType == nil {
        return decl.ResolvedIdentifiers, nil, err
    }

    if newType.Name == decl.Name {
        if structDecl := newType.GetTypeDeclaration().GetStructDecl(); structDecl != nil {
            return structDecl.ResolvedIdentifiers, structDecl, nil
        }
        return decl.ResolvedIdentifiers, nil, nil
    } else {
        decl.Type = newType
        return []*lang.Identifier{newType.NewIdentifier()}, nil, nil
    }

    return nil, nil, err
}

func (c *GenericCompiler) compileNominalType(ctx context.Context, nominalType *lang.NominalType, index IdentifierIndex) (*lang.NominalType, error) {
    if len(nominalType.GenericArguments) == 0 {
        return nominalType, nil
    }

    // skip it, if all the arguments can't be instantiated
    if !nominalType.IsInstantiatedGeneric() {
        return nil, fmt.Errorf("the nominal type (%s) is not instantiated", nominalType.GetGenericFullName())
    }

    thisCtx := context.WithType(ctx, nominalType)
    for i, argument := range nominalType.GenericArguments {
        a, err := c.compileNominalType(thisCtx, argument, index)
        if err != nil {
            return nil, err
        }

        // add nominal to index

        nominalType.GenericArguments[i] = a
    }

    typeDecl := nominalType.GetTypeDeclaration()
    if decl := typeDecl.GetStructDecl(); decl != nil {
        if decl.IsOpacity() { // like Array<String>
            return nominalType, nil
        }
        if newType, err := c.compileStructType(thisCtx, nominalType, index); err != nil {
            return nil, err
        } else {
            newType.Attributes = nominalType.Attributes
            return newType, nil
        }
    }

    if decl := typeDecl.GetTypeAliasDecl(); decl != nil {
        if newType, err := c.compileTypeAlias(thisCtx, nominalType, index); err != nil {
            return nil, err
        } else {
            newType.Attributes = lang.MergeAttributes(newType.Attributes, nominalType.Attributes)
            return newType, nil
        }
    }

    return nominalType, nil
}

func (c *GenericCompiler) instantiateNominalType(ctx context.Context, nominalType *lang.NominalType, instantiates map[string]*lang.NominalType) *lang.NominalType {
    if instantiate, ok := instantiates[nominalType.Name]; ok {
        return &lang.NominalType{
            Name:            instantiate.Name,
            PackageName:     instantiate.PackageName,
            TypeDeclaration: instantiate.TypeDeclaration,
            Attributes:      nominalType.Attributes,
            EnclosingType:   instantiate.EnclosingType,
        }
    }

    newType := &lang.NominalType{
        Name:            nominalType.Name,
        PackageName:     nominalType.PackageName,
        TypeDeclaration: nominalType.TypeDeclaration,
        Attributes:      nominalType.Attributes,
        EnclosingType:   nominalType.EnclosingType,
    }

    for _, argument := range nominalType.GenericArguments {
        a := c.instantiateNominalType(ctx, argument, instantiates)
        newType.GenericArguments = append(newType.GenericArguments, a)
    }

    return newType
}

func (c *GenericCompiler) compileStructType(ctx context.Context, nominalType *lang.NominalType, index IdentifierIndex) (*lang.NominalType, error) {
    structDecl := nominalType.GetTypeDeclaration().GetStructDecl()
    if structDecl == nil {
        return nominalType, nil
    }
    if len(structDecl.GenericParameters) != len(nominalType.GenericArguments) {
        return nil, errors.New("the count of generic argument not equal to parameter")
    }
    if structDecl.Type == nil {
        return nil, errors.New("opacity struct type declaration can't be compiled")
    }

    // generate the new type name
    typeName := nominalType.GetGenericName()

    // create the new file or just the current file
    file, l, err := c.generateSourceFile(ctx, typeName, index)
    if err != nil {
        return nil, err
    }
    if l != nil {
        return l, nil
    }

    compiledDecl := &lang.StructDecl{
        Document:       structDecl.Document,
        Attributes:     nominalType.Attributes,
        PackageName:    file.PackageName,
        SourceFileName: file.FullName,
        Implicit:       true,
        Name:           typeName,
        Type:           &lang.StructType{},
    }

    instantiates := make(map[string]*lang.NominalType)
    for i, parameter := range structDecl.GenericParameters {
        instantiates[parameter.Name] = nominalType.GenericArguments[i]
    }

    // identifier index for the struct declaration
    identifierIndex := make(IdentifierIndex)

    // compile fields
    for _, f := range structDecl.Type.Fields {
        field := &lang.ValueDecl{
            StartPosition: f.StartPosition,
            EndPosition:   f.EndPosition,
            Document:      f.Document,
            Implicit:      true,
            Name:          f.Name,
            Attributes:    f.Attributes,
            Type:          nil,
            Initializer:   proto.Clone(f.Initializer).(*lang.Initializer),
        }

        var err error
        field.Type = c.instantiateNominalType(ctx, f.Type, instantiates)
        field.Type, err = c.compileNominalType(ctx, field.Type, identifierIndex)
        if err != nil {
            return nil, err
        }

        compiledDecl.Type.Fields = append(compiledDecl.Type.Fields, field)
    }

    // compile inherits
    for _, inherit := range structDecl.Type.Inherits {
        var err error
        newType := c.instantiateNominalType(ctx, inherit, instantiates)
        newType, err = c.compileNominalType(ctx, newType, identifierIndex)
        if err != nil {
            return nil, err
        }
        compiledDecl.Type.Inherits = append(compiledDecl.Type.Inherits, newType)
    }

    var resolvedIdentifiers []*lang.Identifier
    for _, id := range structDecl.ResolvedIdentifiers {
        if id.Kind == lang.Identifier_KIND_GENERIC_PARAMETER {
            continue
        }
        resolvedIdentifiers = append(resolvedIdentifiers, id)
    }

    ctxPkg := context.Package(ctx)
    for _, argument := range nominalType.GenericArguments {
        id := ctxPkg.GetResolvedIdentifier(argument.GetFullName())
        if id == nil {
            id = argument.NewIdentifier()
            logs.Warnw("failed to found the resolved identifier in the file", "compiledType", compiledDecl.GetFullName(), "argument", argument.GetFullName())
        }
        resolvedIdentifiers = append(resolvedIdentifiers, id)
    }
    for _, id := range identifierIndex {
        resolvedIdentifiers = append(resolvedIdentifiers, id)
    }

    compiledDecl.ResolvedIdentifiers = lang.MergeDependencies(resolvedIdentifiers)

    // update the identifiers in the current scope
    statement := lang.NewStructDeclStatement(compiledDecl)
    file.Statements = append(file.Statements, statement)

    id := c.updateStructIdentifier(compiledDecl, file, typeName, index)
    return id.ToNominalType(), nil
}

func (c *GenericCompiler) updateStructIdentifier(compiledDecl *lang.StructDecl, file *lang.SourceFile, typeName string, index IdentifierIndex) *lang.Identifier {
    id := &lang.Identifier{
        PackageName:    compiledDecl.PackageName,
        Name:           compiledDecl.Name,
        FullName:       lang.GetFullName(compiledDecl.PackageName, lang.GetEnclosingNames(compiledDecl.EnclosingType), compiledDecl.Name),
        SourceFileName: file.FullName,
        Kind:           lang.Identifier_KIND_STRUCT,
        Declaration:    lang.NewStructDeclaration(compiledDecl),
    }

    if file.Scope == nil || file.Scope.Identifiers == nil {
        file.Scope = &lang.Scope{
            Identifiers: map[string]*lang.Identifier{typeName: id},
        }
    } else {
        file.Scope.Identifiers[typeName] = id
    }

    // update the resolved identifiers in file
    if file.ResolvedIdentifiers == nil {
        file.ResolvedIdentifiers = compiledDecl.ResolvedIdentifiers
    } else {
        file.ResolvedIdentifiers = append(file.ResolvedIdentifiers, compiledDecl.ResolvedIdentifiers...)
        file.ResolvedIdentifiers = lang.MergeDependencies(file.ResolvedIdentifiers)
    }

    index[id.FullName] = id
    return id
}

func (c *GenericCompiler) generateSourceFile(ctx context.Context, typeName string, index IdentifierIndex) (*lang.SourceFile, *lang.NominalType, error) {
    var file *lang.SourceFile
    ctxFile := context.SourceFile(ctx)
    fileName := strcase.ToSnake(typeName) + ".mojo"
    fileFullName := path2.Join(path2.Dir(ctxFile.FullName), fileName)
    file = context.Package(ctx).GetSourceFile(fileFullName)
    if file == nil {
        for _, f := range c.NewFiles {
            if f.FullName == fileFullName {
                if f.Scope != nil && f.Scope.Identifiers != nil {
                    if id := f.Scope.Identifiers[typeName]; id != nil {
                        index[id.FullName] = id
                        return nil, id.ToNominalType(), nil
                    }
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
    return file, nil, nil
}

func (c *GenericCompiler) compileTypeAlias(ctx context.Context, nominalType *lang.NominalType, index IdentifierIndex) (*lang.NominalType, error) {
    aliasDecl := nominalType.GetTypeDeclaration().GetTypeAliasDecl()
    if aliasDecl == nil {
        return nil, fmt.Errorf("the nominal type (%s) do NOT have the resolved declaration", nominalType.GetFullName())
    }

    if len(nominalType.GenericArguments) != len(aliasDecl.GenericParameters) {
        return nil, errors.New("the generic arguments is NOT equals to generic parameters")
    }

    instantiates := make(map[string]*lang.NominalType)
    for i, parameter := range aliasDecl.GenericParameters {
        instantiates[parameter.Name] = nominalType.GenericArguments[i]
    }

    aliasType := c.instantiateNominalType(ctx, aliasDecl.Type, instantiates)
    for _, attribute := range nominalType.Attributes {
        aliasType.Attributes = append(aliasType.Attributes, attribute)
    }
    for _, attribute := range aliasDecl.Attributes {
        aliasType.Attributes = append(aliasType.Attributes, attribute)
    }

    // set or transmit the original type alias name to the end
    aliasName, err := lang.GetStringAttribute(nominalType.Attributes, lang.OriginalTypeAliasName)
    if err != nil {
        if nominalType.IsGeneric() {
            aliasName = nominalType.GetGenericName()
        } else {
            aliasName = nominalType.Name
        }
    }
    aliasType.Attributes = lang.SetStringAttribute(aliasType.Attributes, lang.OriginalTypeAliasName, aliasName)

    compiledType, err := c.compileNominalType(ctx, aliasType, index)
    if err != nil {
        return nil, err
    }

    if lang.IsGenericTypeName(aliasName) {
        // create the new file or just the current file
        file, l, err := c.generateSourceFile(ctx, aliasName, index)
        if err != nil {
            return nil, err
        }
        if l != nil {
            return l, nil
        }

        compiledDecl := &lang.TypeAliasDecl{
            Document:       nil,
            Attributes:     nil,
            PackageName:    file.PackageName,
            SourceFileName: file.FullName,
            Implicit:       true,
            Name:           aliasName,
            Type:           compiledType,
            Scope:          lang.NewScope(),
        }

        var resolvedIdentifiers []*lang.Identifier
        ctxPkg := context.Package(ctx)
        for _, argument := range compiledType.GenericArguments {
            id := ctxPkg.GetResolvedIdentifier(argument.GetFullName())
            if id == nil {
                id = argument.NewIdentifier()
                logs.Warnw("failed to found the resolved identifier in the file", "compiledType", compiledDecl.GetFullName(), "argument", argument.GetFullName())
            }
            resolvedIdentifiers = append(resolvedIdentifiers, id)
        }
        compiledDecl.ResolvedIdentifiers = lang.MergeDependencies(resolvedIdentifiers)

        statement := lang.NewTypeAliasDeclStatement(compiledDecl)
        file.Statements = append(file.Statements, statement)

        // update the identifiers in the current scope
        id := &lang.Identifier{
            PackageName:    compiledDecl.PackageName,
            Name:           compiledDecl.Name,
            FullName:       lang.GetFullName(compiledDecl.PackageName, lang.GetEnclosingNames(compiledDecl.EnclosingType), compiledDecl.Name),
            SourceFileName: file.FullName,
            Kind:           lang.Identifier_KIND_TYPE_ALIAS,
            Declaration:    lang.NewTypeAliasDeclaration(compiledDecl),
        }

        if file.Scope == nil || file.Scope.Identifiers == nil {
            file.Scope = &lang.Scope{
                Identifiers: map[string]*lang.Identifier{aliasName: id},
            }
        } else {
            file.Scope.Identifiers[aliasName] = id
        }

        // update the resolved identifiers in file
        if file.ResolvedIdentifiers == nil {
            file.ResolvedIdentifiers = compiledDecl.ResolvedIdentifiers
        } else {
            file.ResolvedIdentifiers = append(file.ResolvedIdentifiers, compiledDecl.ResolvedIdentifiers...)
            file.ResolvedIdentifiers = lang.MergeDependencies(file.ResolvedIdentifiers)
        }

        index[id.FullName] = id
        return id.ToNominalType(), nil
    }

    return compiledType, nil
}
