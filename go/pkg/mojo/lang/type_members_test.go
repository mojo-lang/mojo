package lang

import (
	"errors"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestStructInheritedFieldsAndEarlyExit(t *testing.T) {
	id := &ValueDecl{Name: "id", Type: &NominalType{Name: "Int"}}
	name := &ValueDecl{Name: "name", Type: &NominalType{Name: "String"}}
	name.SetStringAttribute(core.AliasAttributeName, "displayName")
	base := &StructDecl{Name: "Base", Type: &StructType{Fields: []*ValueDecl{id}}}
	derived := &StructDecl{Name: "Derived", Type: &StructType{Inherits: []*NominalType{{Name: "Base", TypeDeclaration: NewStructTypeDeclaration(base)}}, Fields: []*ValueDecl{name}}}
	require.Equal(t, []*ValueDecl{id, name}, derived.GetAllFields())
	require.Equal(t, []*ValueDecl{id}, derived.GetInheritFields())
	require.Equal(t, []string{"id", "name"}, derived.GetAllFieldNames(FieldNamOptionDefault))
	require.Equal(t, []string{"id", "displayName"}, derived.GetAllFieldNames(FieldNamOptionUseAlias))
	require.Same(t, id, derived.GetField("id"))
	require.Same(t, name, derived.GetField("name"))
	require.Nil(t, derived.GetField("missing"))
	var seen []*ValueDecl
	require.NoError(t, derived.EachField(func(field *ValueDecl) error { seen = append(seen, field); return nil }))
	require.Equal(t, []*ValueDecl{id, name}, seen)
	stop := errors.New("stop")
	for _, stoppingField := range []*ValueDecl{id, name} {
		seen = nil
		err := derived.EachField(func(field *ValueDecl) error {
			seen = append(seen, field)
			if field == stoppingField {
				return stop
			}
			return nil
		})
		require.ErrorIs(t, err, stop)
		require.Same(t, stoppingField, seen[len(seen)-1])
	}
	require.Equal(t, 1, derived.GetFieldCount())
	require.False(t, derived.IsBodyEmpty())
	require.True(t, (&StructDecl{}).IsOpacity())
	require.True(t, (&StructDecl{}).IsBodyEmpty())
	boxed := &StructDecl{Type: &StructType{Inherits: []*NominalType{{TypeDeclaration: NewStructTypeDeclaration(&StructDecl{Name: "Int"})}}}}
	require.True(t, boxed.IsBoxed())
	boxed.Type.Fields = []*ValueDecl{id}
	require.False(t, boxed.IsBoxed())
	var absent *StructDecl
	require.Nil(t, absent.GetAllFields())
	require.Nil(t, absent.GetAllFieldNames(FieldNamOptionDefault))
	require.Nil(t, absent.GetInheritFields())
	require.Nil(t, absent.GetField("x"))
	require.NoError(t, absent.EachField(func(*ValueDecl) error { t.Fatal("nil struct must not invoke callback"); return nil }))
}

func TestMergeFieldsDeduplicatesIncomingNames(t *testing.T) {
	first := &ValueDecl{Name: "id", Type: &NominalType{Name: "Int"}}
	second := &ValueDecl{Name: "name", Type: &NominalType{Name: "String"}}
	conflict := &ValueDecl{Name: "id", Type: &NominalType{Name: "String"}}
	typ := &StructType{Fields: []*ValueDecl{first}}
	require.Same(t, typ, typ.MergeFields([]*ValueDecl{conflict, second, second}))
	require.Equal(t, []*ValueDecl{first, second}, typ.Fields)
}

func TestInterfaceMethodGroupsAndInheritedMethods(t *testing.T) {
	for _, method := range []string{GetMethod, ListMethod, CreateMethod, UpdateMethod, DeleteMethod, BatchGetMethod, BatchCreateMethod, BatchUpdateMethod, BatchDeleteMethod} {
		name := (&FunctionDecl{Name: method + "_users"}).parseStandardMethod()
		require.Equal(t, &MethodName{Method: method, Resource: "User"}, name)
	}
	for _, name := range []string{"get", "get_", "batch", "batch_get", "batch_get_", "batch_unknown_users", "custom"} {
		require.NotPanics(t, func() { require.Nil(t, (&FunctionDecl{Name: name}).parseStandardMethod(), name) })
	}
	get, list, custom := &FunctionDecl{Name: "get_user"}, &FunctionDecl{Name: "list_users"}, &FunctionDecl{Name: "activate_user"}
	typ := &InterfaceType{Methods: []*FunctionDecl{custom, get, list, {Name: "health"}}}
	groups := typ.GetMethodGroups()
	require.Len(t, groups, 1)
	require.Equal(t, []*FunctionDecl{get, list, custom}, groups["User"])
	for _, method := range groups["User"] {
		resource, err := method.GetStringAttribute(core.ResourceAttributeName)
		require.NoError(t, err)
		require.Equal(t, "User", resource)
		require.True(t, method.GetAttribute(core.MethodAttributeName).Implicit)
	}
	verb, err := custom.GetStringAttribute(core.MethodAttributeName)
	require.NoError(t, err)
	require.Equal(t, "activate", verb)
	base := &InterfaceDecl{Type: typ}
	derived := &InterfaceDecl{Type: &InterfaceType{Inherits: []*NominalType{{TypeDeclaration: NewTypeDeclarationFromDeclaration(NewInterfaceDeclaration(base))}}, Methods: []*FunctionDecl{{Name: "ping"}}}}
	require.Equal(t, typ.Methods, derived.GetInheritMethods())
	require.Equal(t, 1, derived.GetMethodCount())
	require.False(t, derived.IsBodyEmpty())
	require.Empty(t, derived.GetMethodGroups())
	require.True(t, (&InterfaceDecl{}).IsBodyEmpty())
}

func TestFunctionCopySeparatesOnlyRequestedFields(t *testing.T) {
	original := &FunctionDecl{Name: "get_user", Attributes: []*Attribute{NewStringAttribute("", "path", "/users")}, Signature: NewFunctionSignature(&NominalType{Name: "User"}, &ValueDecl{Name: "id"}), Body: &BlockStmt{Statements: []*Statement{NewExpressionStatement(NewStringLiteralExpressionFrom("body"))}}, Document: &Document{}}
	shallow := original.Copy()
	require.Same(t, original.Signature, shallow.Signature)
	require.Same(t, original.Body, shallow.Body)
	require.Same(t, original.Attributes[0], shallow.Attributes[0])
	deep := original.Copy("attributes", "signature", "body")
	require.True(t, proto.Equal(original, deep))
	require.Same(t, original.Document, deep.Document)
	deep.Attributes[0].SetString("changed")
	deep.Signature.Parameter.Decls[0].Name = "changed"
	deep.Body.Statements = nil
	path, err := original.GetStringAttribute("path")
	require.NoError(t, err)
	require.Equal(t, "/users", path)
	require.Equal(t, "id", original.Signature.Parameter.Decls[0].Name)
	require.Len(t, original.Body.Statements, 1)
}

func TestInterfaceMethodGroupsChooseTheMostSpecificResource(t *testing.T) {
	user := &FunctionDecl{Name: "get_user"}
	superUser := &FunctionDecl{Name: "get_super_user"}
	custom := &FunctionDecl{Name: "activate_super_users"}
	groups := (&InterfaceType{Methods: []*FunctionDecl{user, superUser, custom}}).GetMethodGroups()
	require.Equal(t, []*FunctionDecl{user}, groups["User"])
	require.Equal(t, []*FunctionDecl{superUser, custom}, groups["SuperUser"])
	verb, err := custom.GetStringAttribute(core.MethodAttributeName)
	require.NoError(t, err)
	require.Equal(t, "activate", verb)
	resource, err := custom.GetStringAttribute(core.ResourceAttributeName)
	require.NoError(t, err)
	require.Equal(t, "SuperUser", resource)
}
