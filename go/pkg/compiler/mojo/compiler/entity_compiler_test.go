package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/semantic"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/semantic/circle"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/semantic/identifier"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
)

func TestEntityCompiler_CompilePackage(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-entity")

	// Fixme
	// graph.Builder{Builder: builder.Builder{Package: pkg}}.Build()
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}

func TestExplicitEntityImplicitID(t *testing.T) {
	pkg := &lang.Package{Name: "sample", FullName: "sample", Scope: lang.NewScope()}
	pkg.Scope.Declare(lang.NewStructDeclaration(&lang.StructDecl{Name: "String", PackageName: "mojo.core"}))
	field := &lang.ValueDecl{Name: "value", Type: &lang.NominalType{Name: "String", PackageName: "mojo.core"}}
	field.SetIntegerAttribute(core.NumberAttributeName, 536870911)
	decl := &lang.StructDecl{Name: "Manual", PackageName: "sample", Type: &lang.StructType{Fields: []*lang.ValueDecl{field}}}
	decl.SetBoolAttribute(core.EntityAttributeName, true)
	ctx := context.WithType(context.Empty(), pkg)
	c := NewEntityCompiler(nil)
	require.NoError(t, c.compileEntityNode(ctx, decl))
	node := pkg.GetEntityNode("sample.Manual")
	require.NotNil(t, node)
	require.Equal(t, "id", node.KeyField.Name)
	require.True(t, node.KeyField.Implicit)
	require.Equal(t, "mojo.core.String", node.KeyField.Type.GetFullName())
	number, err := node.KeyField.GetIntegerAttribute(core.NumberAttributeName)
	require.NoError(t, err)
	require.EqualValues(t, 536870910, number)
	require.NoError(t, c.compileEntityNode(ctx, decl))
	require.Len(t, decl.GetAllFields(), 2)
	require.Same(t, node.KeyField, pkg.GetEntityNode("sample.Manual").KeyField)
	require.False(t, decl.GetAttribute(core.EntityAttributeName).Implicit)
}
