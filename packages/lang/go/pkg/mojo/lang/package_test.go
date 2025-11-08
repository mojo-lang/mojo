package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackage_IsVersionTag(t *testing.T) {
	pkg := &Package{Name: "v1"}
	assert.True(t, pkg.IsVersionTag())

	pkg = &Package{Name: "alpha"}
	assert.True(t, pkg.IsVersionTag())

	pkg = &Package{Name: "beta"}
	assert.True(t, pkg.IsVersionTag())
}
