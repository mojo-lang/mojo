package compiler

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

var (
	pageSize = &lang.ValueDecl{
		Implicit: true,
		Name:     "page_size",
		Type: &lang.NominalType{
			Package: "mojo.core",
			Name:    "Int32",
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
			Package: "mojo.core",
			Name:    "String",
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

func GeneratePaginationParameters() []*lang.ValueDecl {
	return []*lang.ValueDecl{pageSize, pageToken}
}
