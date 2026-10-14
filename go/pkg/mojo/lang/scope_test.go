package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScopeDeclaresTypes(t *testing.T) {
	enclosing := &NominalType{Name: "Outer"}
	for _, tt := range []struct {
		name string
		decl *Declaration
		kind Identifier_Kind
		full string
	}{
		{"struct", NewStructDeclaration(&StructDecl{Name: "Item", PackageName: "sample", Enclosing: enclosing}), Identifier_KIND_STRUCT, "sample.Outer.Item"},
		{"enum", NewEnumDeclaration(&EnumDecl{Name: "Status", PackageName: "sample", Enclosing: enclosing}), Identifier_KIND_ENUM, "sample.Outer.Status"},
		{"interface", NewInterfaceDeclaration(&InterfaceDecl{Name: "Service", PackageName: "sample"}), Identifier_KIND_INTERFACE, "sample.Service"},
		{"alias", NewTypeAliasDeclaration(&TypeAliasDecl{Name: "ID", PackageName: "sample"}), Identifier_KIND_TYPE_ALIAS, "sample.ID"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			scope := NewScope()
			id := scope.Declare(tt.decl)
			require.Equal(t, tt.kind, id.Kind)
			require.Equal(t, tt.full, id.FullName)
			require.Same(t, tt.decl, id.Declaration)
			require.Same(t, id, scope.GetIdentifier(id.Name))
			require.Nil(t, scope.GetIdentifier("Missing"))
		})
	}
	scope := NewScope()
	require.Nil(t, scope.Declare(nil))
	require.Nil(t, scope.Declare(NewImportDeclaration(&ImportDecl{})))
	param := &GenericParameter{Name: "T"}
	require.Equal(t, Identifier_KIND_GENERIC_PARAMETER, scope.Declare(NewGenericParameterDeclaration(param)).Kind)
}

func TestScopeContextShadowingAndClose(t *testing.T) {
	ctx := NewScopeContext()
	require.Nil(t, ctx.CurrentScope())
	require.Nil(t, ctx.GlobalScope())
	require.Nil(t, ctx.Lookup("Item"))
	require.NotPanics(t, ctx.Close)
	global, local := NewScope(), NewScope()
	require.Same(t, ctx, ctx.Open(global))
	original := ctx.Declare(NewStructDeclaration(&StructDecl{Name: "Item"}))
	ctx.Open(local)
	shadow := ctx.Declare(NewStructDeclaration(&StructDecl{Name: "Item"}))
	require.Same(t, global, ctx.GlobalScope())
	require.Same(t, local, ctx.CurrentScope())
	require.Same(t, shadow, ctx.Lookup("Item"))
	ctx.Close()
	require.Same(t, original, ctx.Lookup("Item"))
	ctx.Close()
	require.Nil(t, ctx.Lookup("Item"))
}
