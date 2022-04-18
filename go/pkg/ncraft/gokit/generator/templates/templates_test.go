package templates

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAssetNames(t *testing.T) {
    names := ServiceNames()
    assert.True(t, len(names) > 0)
}
