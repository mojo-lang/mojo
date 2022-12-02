package parser

import (
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

const structPrefix = "struct "
const cStructAttributionName = "c_struct"

type StructRenamer struct {
}

func NewStructRenamer(options core.Options) *StructRenamer {
	_ = options
	return &StructRenamer{}
}

func (r *StructRenamer) ParseSourceFile(ctx context.Context, file *lang.SourceFile) error {
	for _, statement := range file.Statements {
		if decl := statement.GetDeclaration(); decl != nil {
			var err error
			switch d := decl.Declaration.(type) {
			case *lang.Declaration_StructDecl:
				err = r.ParseStruct(ctx, d.StructDecl)
			case *lang.Declaration_TypeAliasDecl:
				err = r.ParseTypeAlias(ctx, d.TypeAliasDecl)
			case *lang.Declaration_FunctionDecl:
				err = r.ParseFunction(ctx, d.FunctionDecl)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *StructRenamer) ParseStruct(ctx context.Context, decl *lang.StructDecl) error {
	if strings.HasPrefix(decl.Name, structPrefix) {
		decl.Name = strings.TrimPrefix(decl.Name, structPrefix)
		decl.Attributes = append(decl.Attributes, lang.NewBoolAttribute("", cStructAttributionName))
	}

	if decl.Type != nil {
		for _, field := range decl.Type.Fields {
			if err := r.ParseNominalType(ctx, field.Type); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *StructRenamer) ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	return r.ParseNominalType(ctx, decl.Type)
}

func (r *StructRenamer) ParseFunction(ctx context.Context, decl *lang.FunctionDecl) error {
	if decl.Signature != nil {
		for _, param := range decl.Signature.Parameter.Decls {
			if err := r.ParseNominalType(ctx, param.Type); err != nil {
				return err
			}
		}
		return r.ParseNominalType(ctx, decl.Signature.Result.Type)
	}

	return nil
}

func (r *StructRenamer) ParseNominalType(ctx context.Context, typ *lang.NominalType) error {
	if typ != nil && strings.HasPrefix(typ.Name, structPrefix) {
		typ.Name = strings.TrimPrefix(typ.Name, structPrefix)
		typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cStructAttributionName))
	}
	return nil
}
