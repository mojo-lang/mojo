package parser

import (
	"regexp"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

const constKeyword = "const"
const constPrefix = "const "
const cConstAttributionName = "c_const"
const cPointerConstAttributeName = "c_pointer_const" // the pointer self is const
var constSuffix = regexp.MustCompile(`\* ?const$`)
var constInfix = regexp.MustCompile(`const ?\*`) // not support char * const *

type ConstRenamer struct {
}

func NewConstRenamer(options core.Options) *ConstRenamer {
	_ = options
	return &ConstRenamer{}
}

func (r *ConstRenamer) ParseSourceFile(ctx context.Context, file *lang.SourceFile) error {
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

func (r *ConstRenamer) ParseStruct(ctx context.Context, decl *lang.StructDecl) error {
	if decl != nil && decl.Type != nil {
		for _, field := range decl.Type.Fields {
			if err := r.ParseNominalType(ctx, field.Type); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ConstRenamer) ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	return r.ParseNominalType(ctx, decl.Type)
}

func (r *ConstRenamer) ParseFunction(ctx context.Context, decl *lang.FunctionDecl) error {
	if decl != nil && decl.Signature != nil {
		for _, param := range decl.Signature.Parameter.GetDecls() {
			if err := r.ParseNominalType(ctx, param.Type); err != nil {
				return err
			}
		}
		return r.ParseNominalType(ctx, decl.Signature.Result.GetType())
	}

	return nil
}

func (r *ConstRenamer) ParseNominalType(ctx context.Context, typ *lang.NominalType) error {
	if typ != nil {
		if strings.HasPrefix(typ.Name, constPrefix) {
			typ.Name = strings.TrimPrefix(typ.Name, constPrefix)
			typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cConstAttributionName))
			return r.ParseNominalType(ctx, typ)
		} else if str := constInfix.FindString(typ.Name); len(str) > 0 {
			segments := strings.Split(typ.Name, str)
			segments[0] = strings.TrimSpace(segments[0])
			if strings.HasSuffix(segments[0], "*") { // FIXME not support right now
				if constSuffix.MatchString(typ.Name) {
					typ.Name = strings.TrimSpace(strings.TrimSuffix(typ.Name, constKeyword))
					typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cPointerConstAttributeName))
				}
			} else {
				typ.Name = strings.Join(segments, " *")
				typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cConstAttributionName))

				return r.ParseNominalType(ctx, typ)
			}
		} else if constSuffix.MatchString(typ.Name) {
			typ.Name = strings.TrimSpace(strings.TrimSuffix(typ.Name, constKeyword))
			typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cPointerConstAttributeName))
		} else if strings.HasSuffix(typ.Name, constKeyword) {
			typ.Name = strings.TrimSpace(strings.TrimSuffix(typ.Name, constKeyword))
			typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("", cConstAttributionName))
		}
	}
	return nil
}
