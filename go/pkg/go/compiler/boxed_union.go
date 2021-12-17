package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"strings"
)

type UnionField struct {
	JsonType string // bool, number, string, array, object
	Type     string
	Name     string

	Boxed       bool
	HasDiscriminatorField bool
}

type BoxedUnion struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string
	Discriminator string

	HasBoolField   bool
	HasNumberField bool
	HasStringField bool
	hasArrayField  bool
	hasObjectField bool

	BoolField    *UnionField
	NumberFields []*UnionField
	StringFields []*UnionField
	ArrayFields  []*UnionField
	ObjectFields []*UnionField
}

func (b *BoxedUnion) Compile(decl *lang.StructDecl) error {
	b.PackageName = decl.GetPackageName()
	b.GoPackageName = GetGoPackage(decl.GetPackageName())

	if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
		b.GoPackageName = pkg
	}

	if discriminator, _ := lang.GetStringAttribute(decl.Attributes, "discriminator"); len(discriminator) > 0 {
		b.Discriminator = discriminator
	}

	b.Name = decl.Name
	b.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), ".")
	b.FullName = GetFullName(b.EnclosingName, b.Name)

	for _, field := range decl.Type.Fields {
		//FIXME need to process boxed type, formatted type
		switch field.Type.GetFullName() {
		case core.BoolTypeName:
		case core.IntTypeName:
		case core.StringTypeName, core.BytesTypeName:
		case core.ArrayTypeName:
		default:
			if f, e := b.compileObjectField(field, decl); e != nil {
				return e
			} else {
				b.ObjectFields = append(b.ObjectFields, f)
				b.hasObjectField = true
			}
		}
	}

	return nil
}

func (b *BoxedUnion) compileObjectField(field *lang.ValueDecl, structDecl *lang.StructDecl) (*UnionField, error) {
	unionField := &UnionField{
		JsonType:              "object",
		Type:                  field.Type.Name,
		Name:                  field.Name,
		Boxed:                 false,
	}

	if len(b.Discriminator) > 0 && !strings.HasPrefix(b.Discriminator, "@") {
		names := field.GetType().GetTypeDeclaration().GetStructDecl().FieldNames(lang.FieldNamOptionDefault)
		for _, name := range names {
			if name == b.Discriminator {
				unionField.HasDiscriminatorField = true
			}
		}
	}

	return unionField, nil
}

func (b *BoxedUnion) GetPackageName() string {
	return b.PackageName
}

func (b *BoxedUnion) GetFullName() string {
	return b.FullName
}
