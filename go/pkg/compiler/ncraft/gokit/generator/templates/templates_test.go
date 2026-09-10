package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetNames(t *testing.T) {
	names := ServiceNames()
	assert.True(t, len(names) > 0)
}

func TestServiceTemplatesExcludeLegacyModels(t *testing.T) {
	for _, name := range ServiceNames() {
		assert.NotContains(t, name, "internal/model/")
		assert.NotContains(t, name, "ENTITY_")
	}
}
