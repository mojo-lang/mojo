package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin/compiler"
)

var (
	pageSize = &lang.ValueDecl{
		Implicit: true,
		Name:     "page_size",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "Int32",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/numeric.mojo",
				Name:           "Int32",
			}),
			Attributes: []*lang.Attribute{{
				Name: "number",
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2000,
					}),
				},
			}},
		},
	}

	pageToken = &lang.ValueDecl{
		Implicit: true,
		Name:     "page_token",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "String",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/string.mojo",
				Name:           "String",
			}),
			Attributes: []*lang.Attribute{{
				Name: "number",
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2001,
					}),
				},
			}},
		},
	}

	skip = &lang.ValueDecl{
		Implicit: true,
		Name:     "skip",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "Int32",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/numeric.mojo",
				Name:           "Int32",
			}),
			Attributes: []*lang.Attribute{{
				Name: "number",
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2002,
					}),
				},
			}},
		},
	}

	totalCount = &lang.ValueDecl{
		Implicit: true,
		Name:     "total_count",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "Int32",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/numeric.mojo",
				Name:           "Int32",
			}),
			Attributes: []*lang.Attribute{{
				Name: "number",
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2000,
					}),
				},
			}},
		},
	}

	nextPageToken = &lang.ValueDecl{
		Implicit: true,
		Name:     "next_page_token",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "String",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/string.mojo",
				Name:           "String",
			}),
			Attributes: []*lang.Attribute{{
				Name: "number",
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2001,
					}),
				},
			}},
		},
	}
)

func PaginationRequestFields() []*lang.ValueDecl {
	return []*lang.ValueDecl{pageSize, pageToken, skip}
}

func PaginationResponseFields() []*lang.ValueDecl {
	return []*lang.ValueDecl{totalCount, nextPageToken}
}

type PaginationCompiler struct {
	compiler.DummyCompiler
}

func (c *PaginationCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	thisCtx := context.WithType(ctx, pkg)
	for _, sourceFile := range pkg.SourceFiles {
		fileCtx := context.WithType(thisCtx, sourceFile)
		for _, statement := range sourceFile.Statements {
			if decl := statement.GetDeclaration(); decl != nil {
				switch decl.Declaration.(type) {
				case *lang.Declaration_InterfaceDecl:
					if err := c.CompileInterface(fileCtx, decl.GetInterfaceDecl()); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (c *PaginationCompiler) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	//for _, s := range decl.StructDecls {
	//
	//}
	//
	//for _, field := range decl.Type.
	return nil
}
