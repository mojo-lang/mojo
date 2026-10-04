package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	strs = []string{
		"$request.body#/user/uuid",
		"$response.body",
		"$url",
		"$response.header.Content-Type",
	}

	exprs = []*Expression{{
		Type:   Expression_TYPE_REQUEST,
		Source: &Expression_Source{Location: Expression_LOCATION_BODY, Path: "/user/uuid"},
	}, {
		Type:   Expression_TYPE_RESPONSE,
		Source: &Expression_Source{Location: Expression_LOCATION_BODY, Path: ""},
	}, {
		Type: Expression_TYPE_URL,
	}, {
		Type:   Expression_TYPE_RESPONSE,
		Source: &Expression_Source{Location: Expression_LOCATION_HEADER, Name: "Content-Type"},
	}}
)

func TestExpression_Format(t *testing.T) {
	for i, expr := range exprs {
		str := expr.Format()
		assert.Equal(t, strs[i], str, "expect to %s, but %s", strs[i], str)
	}
}

func TestExpressionParseBoundaries(t *testing.T) {
	for _, input := range []string{"$url", "$method", "$statusCode", "$request.header.accept", "$request.query.id", "$request.path.id", "$response.body#/items/0"} {
		t.Run(input, func(t *testing.T) {
			expression, err := ParseExpression(input)
			require.NoError(t, err)
			assert.Equal(t, input, expression.Format())
			assert.Equal(t, input, expression.ToString())
		})
	}
	for _, input := range []string{"", "$", "$ur", "$url.extra", "$methods", "$status", "$request", "$request.unknown", "$response", "$response.unknown", "invalid"} {
		t.Run("invalid "+input, func(t *testing.T) {
			expression := &Expression{Type: Expression_TYPE_METHOD}
			assert.Error(t, expression.Parse(input))
			assert.Equal(t, "$method", expression.Format(), "failed parsing must retain the old value")
		})
	}
	expression, err := ParseExpression("$request.path.id")
	require.NoError(t, err)
	require.NoError(t, expression.Parse("$url"))
	assert.Nil(t, expression.Source, "a simple expression must clear the previous source")
	assert.Equal(t, "$url", expression.Format())
	assert.Empty(t, (*Expression)(nil).Format())
	assert.NoError(t, (*Expression)(nil).Parse("$url"))
}

func TestExpression_Parse(t *testing.T) {
	for i, str := range strs {
		expr := &Expression{}
		err := expr.Parse(str)
		assert.NoError(t, err, "failed to parse the %s", str)

		assert.Equal(t, exprs[i].Type, expr.Type, "type expect to %s, but %s", exprs[i].Type, expr.Type)
		assert.Equal(t, exprs[i].Source.GetLocation(), expr.Source.GetLocation(), "source location expect to %s, but %s", exprs[i].Source.GetLocation(), expr.Source.GetLocation())
		assert.Equal(t, exprs[i].Source.GetName(), expr.Source.GetName(), "source name expect to %s, but %s", exprs[i].Source.GetName(), expr.Source.GetName())
		assert.Equal(t, exprs[i].Source.GetPath(), expr.Source.GetPath(), "source path expect to %s, but %s", exprs[i].Source.GetPath(), expr.Source.GetPath())
	}
}

func TestParseExpression(t *testing.T) {
	for i, str := range strs {
		expr, err := ParseExpression(str)
		assert.NoError(t, err, "failed to parse the %s", str)

		assert.Equal(t, exprs[i].Type, expr.Type, "type expect to %s, but %s", exprs[i].Type, expr.Type)
		assert.Equal(t, exprs[i].Source.GetLocation(), expr.Source.GetLocation(), "source location expect to %s, but %s", exprs[i].Source.GetLocation(), expr.Source.GetLocation())
		assert.Equal(t, exprs[i].Source.GetName(), expr.Source.GetName(), "source name expect to %s, but %s", exprs[i].Source.GetName(), expr.Source.GetName())
		assert.Equal(t, exprs[i].Source.GetPath(), expr.Source.GetPath(), "source path expect to %s, but %s", exprs[i].Source.GetPath(), expr.Source.GetPath())
	}
}
