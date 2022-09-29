package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "strings"
)

type BoxedUnion struct {
    *data.Data
}

func (b *BoxedUnion) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {

    bu := &data.BoxedUnion{}
    bu.PackageName = decl.GetPackageName()
    bu.GoPackageName = lang.GetGoPackageName(decl.GetPackageName())

    if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
        bu.GoPackageName = pkg
    }

    if discriminator, _ := lang.GetStringAttribute(decl.Attributes, "discriminator"); len(discriminator) > 0 {
        bu.Discriminator = discriminator
    }

    bu.Name = decl.Name
    bu.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.Enclosing), ".")
    bu.FullName = GetFullName(bu.EnclosingName, bu.Name)

    for _, field := range decl.Type.Fields {
        //FIXME need to process boxed type, formatted type
        switch field.Type.GetFullName() {
        case core.BoolTypeFullName:
        case core.IntTypeFullName:
        case core.StringTypeFullName, core.BytesTypeFullName:
        case core.ArrayTypeFullName:
        default:
            if f, e := b.compileObjectField(field, bu); e != nil {
                return e
            } else {
                bu.ObjectFields = append(bu.ObjectFields, f)
                bu.HasObjectField = true
            }
        }
    }

    return nil
}

func (b *BoxedUnion) compileObjectField(field *lang.ValueDecl, bu *data.BoxedUnion) (*data.UnionField, error) {
    unionField := &data.UnionField{
        JsonType: "object",
        Type:     field.Type.Name,
        Name:     field.Name,
        Boxed:    false,
    }

    if len(bu.Discriminator) > 0 && !strings.HasPrefix(bu.Discriminator, "@") {
        names := field.GetType().GetTypeDeclaration().GetStructDecl().GetAllFieldNames(lang.FieldNamOptionDefault)
        for _, name := range names {
            if name == bu.Discriminator {
                unionField.HasDiscriminatorField = true
            }
        }
    }

    return unionField, nil
}
