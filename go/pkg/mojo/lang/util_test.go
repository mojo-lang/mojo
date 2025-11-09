package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTypeName(t *testing.T) {
	assert.True(t, IsTypeName("Test"))
	assert.True(t, IsTypeName("pkg.Test"))
	assert.True(t, IsTypeName("pkg.pkg.Test.Test"))
	assert.False(t, IsTypeName("pkg.pkg.Test.Test."))
	assert.False(t, IsTypeName("pkg.pkg"))
}

func TestGetTypePackageName(t *testing.T) {
	assert.Equal(t, "pkg.pkg", GetTypePackageName("pkg.pkg.Test"))
	assert.Equal(t, "", GetTypePackageName("Test"))
	assert.Equal(t, "pkg", GetTypePackageName("pkg.Test.Test"))
	assert.Equal(t, "pkg", GetTypePackageName("pkg"))
}

func TestGetTypeTypeName(t *testing.T) {
	assert.Equal(t, "Test", GetTypeTypeName("pkg.pkg.Test"))
	assert.Equal(t, "Test", GetTypeTypeName("Test"))
	assert.Equal(t, "Test.Test", GetTypeTypeName("pkg.Test.Test"))
	assert.Equal(t, "", GetTypeTypeName("pkg"))
}
