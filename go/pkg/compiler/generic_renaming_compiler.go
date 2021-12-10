package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"strings"
)

type LocationType int32

const (
	LocationInherit         LocationType = 0
	LocationGenericArgument LocationType = 1
	LocationField           LocationType = 2
	LocationTypeAlias       LocationType = 3
)

type Reference struct {
	Key        string
	Types      map[LocationType][]*lang.NominalType
	Ids        map[string]*lang.Identifier
	AliasFiles map[string]*lang.SourceFile
	SourceFile *lang.SourceFile
}

type GenericRenamingCompiler struct {
	References map[string]*Reference
}

func NewGenericRenamingCompiler() *GenericRenamingCompiler {
	return &GenericRenamingCompiler{
		References: make(map[string]*Reference),
	}
}

func (c *GenericRenamingCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	thisCtx := context.WithType(ctx, pkg)

	for _, sourceFile := range pkg.SourceFiles {
		if sourceFile.IsGenericInstantiated() {
			addRef := func(key string) {
				ref := &Reference{
					Key:        key,
					SourceFile: sourceFile,
					Types:      make(map[LocationType][]*lang.NominalType),
					Ids:        make(map[string]*lang.Identifier),
					AliasFiles: make(map[string]*lang.SourceFile),
				}
				c.References[ref.Key] = ref
			}
			structs := sourceFile.GetStructDecls()
			if len(structs) > 0 {
				addRef(structs[0].GetFullName())
			}

			aliases := sourceFile.GetTypeAliasDecls()
			if len(aliases) > 0 {
				addRef(aliases[0].GetFullName())
			}
		}
	}

	for _, sourceFile := range pkg.SourceFiles {
		if sourceFile.IsGenericInstantiated() {
			continue
		}
		fileCtx := context.WithType(thisCtx, sourceFile)
		for _, statement := range sourceFile.Statements {
			if decl := statement.GetDeclaration(); decl != nil {
				switch decl.Declaration.(type) {
				case *lang.Declaration_StructDecl:
					if structDecl := decl.GetStructDecl(); structDecl != nil {
						err := c.CompileStruct(fileCtx, structDecl)
						if err != nil {
							return err
						}
					}
				case *lang.Declaration_TypeAliasDecl:
					if typeAliasDecl := decl.GetTypeAliasDecl(); typeAliasDecl != nil {
						err := c.CompileTypeAlias(fileCtx, typeAliasDecl)
						if err != nil {
							return err
						}
					}
				case *lang.Declaration_EnumDecl:
				case *lang.Declaration_InterfaceDecl:
					//case *lang.Declaration_AttributeDecl:
					//case *lang.Declaration_ConstantDecl:
					//case *lang.Declaration_VariableDecl:
					//case *lang.Declaration_FunctionDecl:
				}
			}
		}
	}

	return c.Renaming()
}

func (c *GenericRenamingCompiler) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	thisCtx := context.WithType(ctx, decl)

	for _, d := range decl.StructDecls {
		if err := c.CompileStruct(thisCtx, d); err != nil {
			return err
		}
	}
	for _, d := range decl.TypeAliasDecls {
		if err := c.CompileTypeAlias(thisCtx, d); err != nil {
			return err
		}
	}

	if decl.Type == nil {
		return nil
	}

	for _, i := range decl.Type.Inherits {
		c.compileNominalType(thisCtx, i, LocationInherit)
	}

	for _, f := range decl.Type.Fields {
		fieldCtx := context.WithType(thisCtx, f)
		c.compileNominalType(fieldCtx, f.Type, LocationField)
	}

	for _, id := range decl.Scope.Identifiers {
		if r, ok := c.References[id.FullName]; ok {
			r.Ids[id.FullName] = id
		}
	}

	for _, id := range decl.ResolvedIdentifiers {
		if r, ok := c.References[id.FullName]; ok {
			r.Ids[id.FullName] = id
		}
	}

	return nil
}

func (c *GenericRenamingCompiler) compileNominalType(ctx context.Context, t *lang.NominalType, location LocationType) error {
	if r, ok := c.References[t.GetFullName()]; ok {
		r.Types[location] = append(r.Types[location], t)
	} else {
		for _, a := range t.GenericArguments {
			if r, ok = c.References[a.GetFullName()]; ok {
				r.Types[LocationGenericArgument] = append(r.Types[LocationGenericArgument], a)
			}
		}
	}
	return nil
}

func (c *GenericRenamingCompiler) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	thisCtx := context.WithType(ctx, decl)

	if err := c.compileNominalType(thisCtx, decl.Type, LocationTypeAlias); err != nil {
		return err
	}

	if r, ok := c.References[decl.Type.GetFullName()]; ok {
		r.AliasFiles[decl.GetFullName()] = context.SourceFile(thisCtx)
	}

	for _, id := range decl.Scope.Identifiers {
		if r, ok := c.References[id.FullName]; ok {
			r.Ids[id.FullName] = id
		}
	}

	for _, id := range decl.ResolvedIdentifiers {
		if r, ok := c.References[id.FullName]; ok {
			r.Ids[id.FullName] = id
		}
	}

	return nil
}

func (c *GenericRenamingCompiler) Renaming() error {
	for key, ref := range c.References {
		if !needRename(ref) {
			continue
		}

		if aliasTypes, ok := ref.Types[LocationTypeAlias]; ok {
			if len(aliasTypes) != len(ref.AliasFiles) {
				// something wrong
			}

			decls := ref.SourceFile.GetStructDecls()
			if len(decls) > 0 {
				var newDecl *lang.StructDecl
				for k, f := range ref.AliasFiles {
					for i, statement := range f.GetStatements() {
						if decl := statement.GetDeclaration().GetTypeAliasDecl(); decl.GetFullName() == k {
							newDecl = &lang.StructDecl{
								StartPosition:         decls[0].StartPosition,
								EndPosition:           nil,
								Document:              decls[0].Document,
								PackageName:           decls[0].PackageName,
								SourceFileName:        decl.SourceFileName,
								Implicit:              true,
								Name:                  decl.Name,
								Attributes:            decls[0].Attributes,
								EnclosingType:         decl.EnclosingType,
								Group:                 decls[0].Group,
								ResolvedIdentifiers:   decls[0].ResolvedIdentifiers,
								UnresolvedIdentifiers: decls[0].UnresolvedIdentifiers,
								Type:                  decls[0].Type,
								Scope:                 decls[0].Scope,
							}
							f.Statements[i] = lang.NewStructDeclStatement(newDecl)

							enclosing := decl.EnclosingType
							for enclosing != nil {
								declaration := enclosing.TypeDeclaration.GetStructDecl()
								for _, id := range declaration.GetScope().GetIdentifiers() {
									if id.Name == decl.Name {
										id.Declaration = f.Statements[i].GetDeclaration()
									}
								}
								enclosing = declaration.EnclosingType
							}

							for _, id := range f.Scope.Identifiers {
								if id.Name == decl.Name {
									id.Declaration = f.Statements[i].GetDeclaration()
								}
							}

							// processing the package scope
							break
						}
					}
				}

				if newDecl != nil {
					for _, types := range ref.Types {
						for _, nt := range types {
							nt.Name = newDecl.Name
							nt.EnclosingType = newDecl.EnclosingType
							nt.TypeDeclaration = lang.NewStructTypeDeclaration(newDecl)
						}
					}

					for _, id := range ref.Ids {
						id.Name = newDecl.Name
						id.FullName = newDecl.GetFullName()
						id.SourceFileName = newDecl.SourceFileName
					}
				}
			}
		} else {
			newName := transformer.GenericTypeNamer{}.NameFrom(key)
			newFileName := strcase.ToSnake(newName) + ".mojo"
			ref.SourceFile.FullName = strings.Replace(ref.SourceFile.FullName, ref.SourceFile.Name, newFileName, 1)
			ref.SourceFile.Name = newFileName

			structDecls := ref.SourceFile.GetStructDecls()
			if len(structDecls) > 0 {
				structDecls[0].Name = newName
				structDecls[0].SourceFileName = ref.SourceFile.FullName
			}

			aliasDecls := ref.SourceFile.GetTypeAliasDecls()
			if len(aliasDecls) > 0 {
				aliasDecls[0].Name = newName
				aliasDecls[0].SourceFileName = ref.SourceFile.FullName
			}

			for _, types := range ref.Types {
				for _, nt := range types {
					nt.Name = newName
				}
			}

			for _, id := range ref.Ids {
				id.FullName = strings.Replace(id.FullName, id.Name, newName, 1)
				id.Name = newName
				id.SourceFileName = ref.SourceFile.FullName
			}
		}
	}

	return nil
}

func needRename(ref *Reference) bool {
	for l := range ref.Types {
		if l != LocationInherit {
			return true
		}
	}
	return false
}
