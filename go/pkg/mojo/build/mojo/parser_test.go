package mojo

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_TreePackages(t *testing.T) {
	packages := map[string]*lang.Package{
		"foo":                 {FullName: "foo"},
		"foo.bar":             {FullName: "foo.bar"},
		"foo.bak":             {FullName: "foo.bak"},
		"foo.bar.car":         {FullName: "foo.bar.car"},
		"foo.bar.car.dir.far": {FullName: "foo.bar.car.dir.far"},
	}

	parser := NewParser("")
	root, err := parser.TreePackages(packages)
	assert.NoError(t, err)
	assert.Equal(t, "foo", root.FullName)
	assert.Equal(t, "foo.bak", root.Children[0].FullName)
	assert.Equal(t, "foo.bar", root.Children[1].FullName)
	assert.Equal(t, "foo.bar.car", root.Children[1].Children[0].FullName)
	assert.Equal(t, "foo.bar.car.dir.far", root.Children[1].Children[0].Children[0].FullName)
}

func TestParser_TreePackages2(t *testing.T) {
	packages := map[string]*lang.Package{
		"foo":                 {FullName: "foo"},
		"foo.bar":             {FullName: "foo.bar"},
		"foo.bak":             {FullName: "foo.bak"},
		"foo.car.car":         {FullName: "foo.car.car"},
		"foo.car.car.dir.far": {FullName: "foo.car.car.dir.far"},
		"foo.car.car.eir.far": {FullName: "foo.bar.car.eir.far"},
	}

	parser := NewParser("")
	root, err := parser.TreePackages(packages)
	assert.NoError(t, err)
	assert.Equal(t, "foo", root.FullName)
	assert.Equal(t, "foo.bak", root.Children[0].FullName)
	assert.Equal(t, "foo.bar", root.Children[1].FullName)
	assert.Equal(t, "foo.car.car", root.Children[2].FullName)
	assert.Equal(t, "foo.car.car.dir.far", root.Children[2].Children[0].FullName)
}

func TestParser_TreePackagesError(t *testing.T) {
	packages := map[string]*lang.Package{
		"foo": {FullName: "foo"},
		"bar": {FullName: "bar"},
	}

	parser := NewParser("")
	_, err := parser.TreePackages(packages)
	assert.Error(t, err)
}
