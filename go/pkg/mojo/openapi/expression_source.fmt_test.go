package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestExpressionSourceParseBoundaries(t *testing.T) {
	for _, input := range []string{"header.Content-Type", "query.id", "path.id", "body", "body#/items/0", "body#/a~1b/~0key"} {
		t.Run(input, func(t *testing.T) {
			source, err := ParseExpressionSource(input)
			require.NoError(t, err)
			assert.Equal(t, input, source.Format())
		})
	}
	for _, input := range []string{"", "bod", "header", "query", "path", "head", "quer", "body.extra", "body/invalid", "unknown"} {
		source := &Expression_Source{Location: Expression_LOCATION_PATH, Name: "id"}
		assert.Error(t, source.Parse(input), input)
		assert.Equal(t, "path.id", source.Format(), input)
	}
	source, err := ParseExpressionSource("header.accept")
	require.NoError(t, err)
	require.NoError(t, source.Parse("body#/id"))
	assert.Empty(t, source.Name)
	assert.Equal(t, "body#/id", source.Format())
	require.NoError(t, source.Parse("query.name"))
	assert.Empty(t, source.Path)
	require.NoError(t, source.Parse("body"))
	assert.Equal(t, "body", source.Format())
	assert.Empty(t, (*Expression_Source)(nil).Format())
	assert.NoError(t, (*Expression_Source)(nil).Parse("body"))
}
