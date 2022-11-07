package context

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestLookupIdentifier(t *testing.T) {
	ctx := WithScope(Empty())
	Scope(ctx).Declare(lang.NewTypeAliasDeclaration(&lang.TypeAliasDecl{Name: "Test", PackageName: "testdata"}))

	ctx2 := WithScope(ctx)
	Scope(ctx2).Declare(lang.NewTypeAliasDeclaration(&lang.TypeAliasDecl{Name: "Test2", PackageName: "testdata"}))

	ctx3 := WithScope(ctx)
	Scope(ctx3).Declare(lang.NewTypeAliasDeclaration(&lang.TypeAliasDecl{Name: "Test2", PackageName: "test3"}))

	assert.NotNil(t, LookupIdentifier(ctx2, "Test"))
	assert.NotNil(t, LookupIdentifier(ctx2, "Test2"))

	assert.Nil(t, LookupIdentifier(ctx, "Test2"))

	id := LookupIdentifier(ctx3, "Test2")
	assert.NotNil(t, id)
	assert.Equal(t, "test3", id.PackageName)
}
