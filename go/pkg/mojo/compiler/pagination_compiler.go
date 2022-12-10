package compiler

import (
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func init() {
	plugin.RegisterPlugin(NewPaginationCompiler(nil))
}

const paginationName = "compiler.pagination"

type PaginationCompiler struct {
	plugin.BasicPlugin

	Options core.Options
}

func NewPaginationCompiler(options core.Options) *PaginationCompiler {
	return &PaginationCompiler{
		BasicPlugin: plugin.BasicPlugin{
			Name:          paginationName,
			Group:         "compiler",
			GroupPriority: 10,
			Priority:      15,
			Creator: func(options core.Options) plugin.Plugin {
				return NewPaginationCompiler(options)
			},
		},
		Options: options,
	}
}

func (c *PaginationCompiler) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	pkg := context.Package(ctx)
	logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName, "interface", decl.Name)

	for _, method := range decl.GetType().GetMethods() {
		if !method.HasAttribute(core.PaginationAttributeName) {
			name := method.Name
			resultTypeName := method.GetSignature().GetResultType().GetFullName()
			if resultTypeName == core.ArrayTypeFullName &&
				(strings.HasPrefix(name, "list_") || !strings.HasPrefix(name, "batch_")) {
				method.SetBoolAttribute(core.PaginationAttributeName, true)
			}
		}
	}
	return nil
}

var (
	pageSize = &lang.ValueDecl{
		Implicit: true,
		Name:     "page_size",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "Int32",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/integer.mojo",
				Name:           "Int32",
			}),
			Attributes: []*lang.Attribute{{
				Name: core.NumberAttributeName,
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
				Name: core.NumberAttributeName,
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
				SourceFileName: "mojo/core/integer.mojo",
				Name:           "Int32",
			}),
			Attributes: []*lang.Attribute{{
				Name: core.NumberAttributeName,
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
				SourceFileName: "mojo/core/integer.mojo",
				Name:           "Int32",
			}),
			Attributes: []*lang.Attribute{{
				Name: core.NumberAttributeName,
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
				Name: core.NumberAttributeName,
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
