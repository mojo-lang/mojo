package identifier

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/plugin"
)

type Parser interface {
	plugin.Parser

	ParserSourceFile(ctx *plugin.Context, file *lang.SourceFile) error

	ParseStruct(ctx *plugin.Context, decl *lang.StructDecl) error

	ParseEnum(ctx *plugin.Context, decl *lang.EnumDecl) error

	ParseInterface(ctx *plugin.Context, decl *lang.InterfaceDecl) error

	ParseTypeAlias(ctx *plugin.Context, decl *lang.TypeAliasDecl) error

	ParseImport(ctx *plugin.Context, decl *lang.ImportDecl) error
}
