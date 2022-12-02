package parser

import (
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

const cPointerAttributionName = "c_pointer"
const cDoublePointerAttributionName = "c_double_pointer"

type PointerRenamer struct {
}

func NewPointerRenamer(options core.Options) *PointerRenamer {
	_ = options
	return &PointerRenamer{}
}

func (r *PointerRenamer) ParseSourceFile(ctx context.Context, file *lang.SourceFile) error {
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

func (r *PointerRenamer) ParseStruct(ctx context.Context, decl *lang.StructDecl) error {
	if decl.Type != nil {
		for _, field := range decl.Type.Fields {
			if err := r.ParseNominalType(ctx, field.Type); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *PointerRenamer) ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	return r.ParseNominalType(ctx, decl.Type)
}

func (r *PointerRenamer) ParseFunction(ctx context.Context, decl *lang.FunctionDecl) error {
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

func (r *PointerRenamer) ParseNominalType(ctx context.Context, typ *lang.NominalType) error {
	if typ != nil {
		if strings.HasSuffix(typ.Name, "**") {
			typ.Name = strings.TrimSuffix(typ.Name, "**")
			typ.Name = strings.TrimSpace(typ.Name)
			typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cDoublePointerAttributionName))
		} else if strings.HasSuffix(typ.Name, "*") && (!strings.HasPrefix(typ.Name, "char *") || !strings.HasPrefix(typ.Name, "char*")) {
			typ.Name = strings.TrimSuffix(typ.Name, "*")
			typ.Name = strings.TrimSpace(typ.Name)
			typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cPointerAttributionName))
		}
	}
	return nil
}
