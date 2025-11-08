package lang

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

func TestNewPackageRequirementVersionOfCaret(t *testing.T) {
	// ^1.2.3  :=  >=1.2.3, <2.0.0
	// ^1.2    :=  >=1.2.0, <2.0.0
	// ^1      :=  >=1.0.0, <2.0.0
	// ^0.2.3  :=  >=0.2.3, <0.3.0
	// ^0.2    :=  >=0.2.0, <0.3.0
	// ^0.0.3  :=  >=0.0.3, <0.0.4
	// ^0.0    :=  >=0.0.0, <0.1.0
	// ^0      :=  >=0.0.0, <1.0.0

	test := func(t *testing.T, input string, min *core.Version, max *core.Version) {
		v := NewPackageRequirementVersion(input)
		assert.Equal(t, Package_Requirement_Version_TYPE_CARET, v.Type)
		assert.Equal(t, min.Major, v.Range.Start.Major)
		assert.Equal(t, min.Minor, v.Range.Start.Minor)
		assert.Equal(t, min.Patch, v.Range.Start.Patch)
		assert.Equal(t, max.Major, v.Range.End.Major)
		assert.Equal(t, max.Minor, v.Range.End.Minor)
		assert.Equal(t, max.Patch, v.Range.End.Patch)
		assert.False(t, v.Range.StartExcluded)
		assert.False(t, v.Range.EndIncluded)
	}

	test(t, "^1.2.3", &core.Version{Major: 1, Minor: 2, Patch: 3}, &core.Version{Major: 2, Minor: 0, Patch: 0})
	test(t, "^1.2", &core.Version{Major: 1, Minor: 2, Patch: 0}, &core.Version{Major: 2, Minor: 0, Patch: 0})
	test(t, "^1", &core.Version{Major: 1, Minor: 0, Patch: 0}, &core.Version{Major: 2, Minor: 0, Patch: 0})
	test(t, "^0.2.3", &core.Version{Major: 0, Minor: 2, Patch: 3}, &core.Version{Major: 0, Minor: 3, Patch: 0})
	test(t, "^0.2", &core.Version{Major: 0, Minor: 2, Patch: 0}, &core.Version{Major: 0, Minor: 3, Patch: 0})
	test(t, "^0.0.3", &core.Version{Major: 0, Minor: 0, Patch: 3}, &core.Version{Major: 0, Minor: 0, Patch: 4})
	test(t, "^0.0", &core.Version{Major: 0, Minor: 0, Patch: 0}, &core.Version{Major: 0, Minor: 1, Patch: 0})
	test(t, "^0", &core.Version{Major: 0, Minor: 0, Patch: 0}, &core.Version{Major: 1, Minor: 0, Patch: 0})
	test(t, "1.2.3", &core.Version{Major: 1, Minor: 2, Patch: 3}, &core.Version{Major: 2, Minor: 0, Patch: 0})
	test(t, "1.2", &core.Version{Major: 1, Minor: 2, Patch: 0}, &core.Version{Major: 2, Minor: 0, Patch: 0})
	test(t, "1", &core.Version{Major: 1, Minor: 0, Patch: 0}, &core.Version{Major: 2, Minor: 0, Patch: 0})
	test(t, "0.2.3", &core.Version{Major: 0, Minor: 2, Patch: 3}, &core.Version{Major: 0, Minor: 3, Patch: 0})
	test(t, "0.2", &core.Version{Major: 0, Minor: 2, Patch: 0}, &core.Version{Major: 0, Minor: 3, Patch: 0})
	test(t, "0.0.3", &core.Version{Major: 0, Minor: 0, Patch: 3}, &core.Version{Major: 0, Minor: 0, Patch: 4})
	test(t, "0.0", &core.Version{Major: 0, Minor: 0, Patch: 0}, &core.Version{Major: 0, Minor: 1, Patch: 0})
	test(t, "0", &core.Version{Major: 0, Minor: 0, Patch: 0}, &core.Version{Major: 1, Minor: 0, Patch: 0})
}
