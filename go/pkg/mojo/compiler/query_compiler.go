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
	plugin.RegisterPlugin(NewQueryCompiler(nil))
}

const queryName = "compiler.query"

type QueryCompiler struct {
	plugin.BasicPlugin

	Options core.Options
}

func NewQueryCompiler(options core.Options) *QueryCompiler {
	return &QueryCompiler{
		BasicPlugin: plugin.BasicPlugin{
			Name:          queryName,
			Group:         "compiler",
			GroupPriority: 10,
			Priority:      16,
			Creator: func(options core.Options) plugin.Plugin {
				return NewQueryCompiler(options)
			},
		},
		Options: options,
	}
}

func (c *QueryCompiler) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	pkg := context.Package(ctx)
	logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName, "interface", decl.Name)

	for _, method := range decl.GetType().GetMethods() {
		if !method.HasAttribute(core.QueryAttributeName) {
			name := method.Name
			resultTypeName := method.GetSignature().GetResultType().GetFullName()
			if resultTypeName == core.ArrayTypeFullName &&
				(strings.HasPrefix(name, "list_") || !strings.HasPrefix(name, "batch_")) {
				method.SetBoolAttribute(core.QueryAttributeName, true)
			}
		}
	}
	return nil
}

var (
	filter = &lang.ValueDecl{
		Implicit: true,
		Name:     "filter",
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
						Value:    2011,
					}),
				},
			}},
		},
	}

	order = &lang.ValueDecl{
		Implicit: true,
		Name:     "order",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "Ordering",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/ordering.mojo",
				Name:           "Ordering",
			}),
			Attributes: []*lang.Attribute{{
				Name: core.NumberAttributeName,
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2012,
					}),
				},
			}},
		},
	}

	fieldMask = &lang.ValueDecl{
		Implicit: true,
		Name:     "field_mask",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "FieldMask",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/field_mask.mojo",
				Name:           "FieldMask",
			}),
			Attributes: []*lang.Attribute{{
				Name: core.NumberAttributeName,
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2013,
					}),
				},
			}},
		},
	}

	unique = &lang.ValueDecl{
		Implicit: true,
		Name:     "unique",
		Type: &lang.NominalType{
			PackageName: "mojo.core",
			Name:        "Bool",
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{
				PackageName:    "mojo.core",
				SourceFileName: "mojo/core/bool.mojo",
				Name:           "Bool",
			}),
			Attributes: []*lang.Attribute{{
				Name: core.NumberAttributeName,
				Arguments: []*lang.Argument{
					lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
						Kind:     0,
						Implicit: true,
						Value:    2014,
					}),
				},
			}},
		},
	}
)

func QueryRequestFields() []*lang.ValueDecl {
	return []*lang.ValueDecl{filter, order, fieldMask, unique}
}
