package semantic

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestParser_TreePackages(t *testing.T) {
	packages := []*lang.Package{
		{FullName: "foo"},
		{FullName: "foo.bar"},
		{FullName: "mojo.bak"},
		{FullName: "mo.bar.car"},
		{FullName: "jo.bar.car.dir.far"},
	}

	root := treePackages(packages, nil)
	assert.Equal(t, true, root.IsGlobal())
	assert.Equal(t, "foo", root.Children[0].FullName)
	assert.Equal(t, "jo", root.Children[1].FullName)
	assert.Equal(t, "mo", root.Children[2].FullName)
	assert.Equal(t, "mojo", root.Children[3].FullName)

	assert.Equal(t, "foo.bar", root.Children[0].Children[0].FullName)
	assert.Equal(t, "jo.bar.car.dir.far", root.Children[1].Children[0].Children[0].Children[0].Children[0].FullName)
}

func TestParser_TreePackages_Single(t *testing.T) {
	packages := []*lang.Package{
		{FullName: "foo.bar"},
	}

	root := treePackages(packages, nil)
	assert.Equal(t, true, root.IsGlobal())
	assert.Equal(t, "foo", root.Children[0].FullName)
	assert.Equal(t, "foo.bar", root.Children[0].Children[0].FullName)
}
