package converter

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"
	"github.com/stretchr/testify/require"
)

func scalarType(name string, number int64) *lang.NominalType {
	return &lang.NominalType{Name: name, PackageName: "mojo.core",
		TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{Name: name, PackageName: "mojo.core"}),
		Attributes:      lang.SetIntegerAttribute(nil, core.NumberAttributeName, number)}
}

func descriptorContext() (context.Context, *descriptor.File) {
	file := descriptor.NewFileWithName("sample/item.proto", "sample")
	ctx := context.WithType(context.Empty(), &lang.Package{FullName: "sample"})
	ctx = context.WithType(ctx, &lang.SourceFile{FullName: "sample/item.mojo", PackageName: "sample"})
	return context.WithDescriptor(ctx, file), file
}

func TestStructConvertsFieldsAndDeduplicatesOptionImports(t *testing.T) {
	ctx, file := descriptorContext()
	name := scalarType("String", 1)
	name.Attributes = lang.SetStringAttribute(name.Attributes, core.AliasAttributeName, "display_name")
	id := scalarType("Int64", 2)
	id.Attributes = lang.SetStringAttribute(id.Attributes, db.IndexAttributeName, "unique")
	array := &lang.NominalType{Name: "Array", PackageName: "mojo.core", GenericArguments: []*lang.NominalType{scalarType("String", 0)}, Attributes: lang.SetIntegerAttribute(nil, core.NumberAttributeName, 3)}
	decl := &lang.StructDecl{Name: "Item", Type: &lang.StructType{Fields: []*lang.ValueDecl{
		{Name: "name", Type: name}, {Name: "id", Type: id}, {Name: "tags", Type: array},
	}}}
	msg := descriptor.NewMessage(file)
	require.NoError(t, (Struct{}).Compile(ctx, decl, msg))
	require.Len(t, file.Messages, 1)
	require.Same(t, msg, file.Messages[0])
	require.Len(t, msg.Fields, 3)
	for i, want := range []struct {
		name, typ string
		repeated  bool
	}{
		{"name", "string", false}, {"id", "int64", false}, {"tags", "string", true},
	} {
		field := msg.Fields[i]
		require.Equal(t, want.name, field.GetName())
		require.Equal(t, int32(i+1), field.GetNumber())
		require.Equal(t, want.typ, field.GetTypeName())
		require.Equal(t, want.repeated, field.IsRepeated())
	}
	require.Equal(t, "display_name", msg.Fields[0].GetStringOption(mojo.E_Alias))
	require.Equal(t, "unique", msg.Fields[1].GetStringOption(mojo.E_DbIndex))
	require.Equal(t, []string{OptionsDependency}, file.GetDependencies())
}

func TestStructRejectsInvalidFieldNumbersAndArrayArguments(t *testing.T) {
	for _, tt := range []struct {
		name    string
		typ     *lang.NominalType
		message string
	}{
		{"missing number", &lang.NominalType{Name: "String"}, "has not set the number"},
		{"zero", scalarType("String", 0), "must be positive"},
		{"negative", scalarType("String", -1), "must be positive"},
		{"empty array", &lang.NominalType{Name: "Array"}, "unexpect array type"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx, file := descriptorContext()
			decl := &lang.StructDecl{Name: "Item", Type: &lang.StructType{Fields: []*lang.ValueDecl{{Name: "bad", Type: tt.typ}}}}
			err := (Struct{}).Compile(ctx, decl, descriptor.NewMessage(file))
			require.ErrorContains(t, err, tt.message)
			require.Empty(t, file.Messages)
		})
	}
}

func TestStructInheritsFieldsAndRegistersNestedTypesOnce(t *testing.T) {
	ctx, file := descriptorContext()
	base := &lang.StructDecl{Name: "Base", Type: &lang.StructType{Fields: []*lang.ValueDecl{{Name: "id", Type: scalarType("Int64", 1)}}},
		ResolvedIdentifiers: []*lang.Identifier{{SourceFileName: "external/value.mojo"}, {SourceFileName: "mojo/core/string.mojo"}}}
	decl := &lang.StructDecl{Name: "Item", Type: &lang.StructType{
		Inherits: []*lang.NominalType{{Name: "Base", TypeDeclaration: lang.NewStructTypeDeclaration(base)}},
		Fields:   []*lang.ValueDecl{{Name: "name", Type: scalarType("String", 2)}},
	}, StructDecls: []*lang.StructDecl{{Name: "Nested"}}, EnumDecls: []*lang.EnumDecl{{Name: "Status", Type: &lang.EnumType{Enumerators: []*lang.ValueDecl{{Name: "Unknown"}}}}}}
	msg := descriptor.NewMessage(file)
	require.NoError(t, (Struct{}).Compile(ctx, decl, msg))
	require.Len(t, file.Messages, 1)
	require.Len(t, msg.Messages, 1)
	require.Equal(t, "Nested", msg.Messages[0].GetName())
	require.Len(t, msg.Enums, 1)
	require.Equal(t, "Status", msg.Enums[0].GetName())
	require.Empty(t, file.Enums)
	require.Len(t, msg.Fields, 2)
	require.Equal(t, "id", msg.Fields[0].GetName())
	require.Equal(t, "name", msg.Fields[1].GetName())
	require.Equal(t, []string{"external/value.proto"}, file.GetDependencies())
}

func TestMapConvertsSingleNestedEntry(t *testing.T) {
	ctx, file := descriptorContext()
	parent := descriptor.NewMessage(file).SetName("Item")
	ctx = context.WithDescriptor(context.WithType(ctx, &lang.ValueDecl{Name: "labels"}), parent)
	typ := &lang.NominalType{Name: "Map", PackageName: "mojo.core", GenericArguments: []*lang.NominalType{scalarType("String", 0), scalarType("String", 0)}}
	kind, name, err := (&Map{}).ConvertTo(ctx, typ)
	require.NoError(t, err)
	require.Equal(t, "struct", kind)
	require.Equal(t, "LabelsEntry", name)
	require.Len(t, parent.Messages, 1)
	require.Len(t, parent.Proto.NestedType, 1)
	entry := parent.Messages[0]
	require.True(t, entry.IsMapEntry())
	require.Len(t, entry.Fields, 2)
	require.Equal(t, "key", entry.Fields[0].GetName())
	require.Equal(t, int32(1), entry.Fields[0].GetNumber())
	require.Equal(t, "value", entry.Fields[1].GetName())
	require.Equal(t, int32(2), entry.Fields[1].GetNumber())
	require.Empty(t, file.Messages)
}

func TestMapRejectsInvalidTypeAndArity(t *testing.T) {
	for _, typ := range []*lang.NominalType{{Name: "String"}, {Name: "Map", PackageName: "mojo.core"}, {Name: "Map", PackageName: "mojo.core", GenericArguments: []*lang.NominalType{{Name: "String"}}}} {
		ctx, file := descriptorContext()
		kind, name, err := (&Map{}).ConvertTo(ctx, typ)
		require.Error(t, err)
		require.Empty(t, kind)
		require.Empty(t, name)
		require.Empty(t, file.Messages)
	}
}
