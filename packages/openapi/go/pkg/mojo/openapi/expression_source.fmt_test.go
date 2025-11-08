package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sourceStrs = []string{
		"body#/user/uuid",
		"body",
		"header.Content-Type",
	}

	sources = []*Expression_Source{{
		Location: Expression_LOCATION_BODY,
		Path:     "/user/uuid",
	}, {
		Location: Expression_LOCATION_BODY,
		Path:     "",
	}, {
		Location: Expression_LOCATION_HEADER,
		Name:     "Content-Type",
	}}
)

func TestExpression_SourceFormat(t *testing.T) {
	for i, src := range sources {
		str := src.Format()
		assert.Equal(t, sourceStrs[i], str, "expect to %s, but %s", sourceStrs[i], str)
	}
}
