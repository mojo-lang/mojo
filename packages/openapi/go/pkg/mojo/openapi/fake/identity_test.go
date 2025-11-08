package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentify(t *testing.T) {
	id := Identify()

	assert.NotEmpty(t, id)
	assert.Len(t, id, 18)
}
