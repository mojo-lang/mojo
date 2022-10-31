package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetNames(t *testing.T) {
	names := ServiceNames()
	assert.True(t, len(names) > 0)
}
