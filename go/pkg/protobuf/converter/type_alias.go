package converter

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type TypeAlias struct {
}

func (s TypeAlias) Compile(ctx context.Context, decl *lang.TypeAliasDecl, desc *descriptor.Message) error {
	_ = desc
	decl.Type.Attributes = lang.SetStringAttribute(decl.Type.Attributes, lang.OriginalTypeAliasName, decl.Name)
	_, _, err := Nominal{}.Compile(ctx, decl.Type)
	return err
}
