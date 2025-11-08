package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediaType_IsSame(t *testing.T) {
	mt := &MediaType{
		Type:    "application",
		Subtype: "json",
	}

	assert.True(t, mt.IsSame(ApplicationJson))
}
