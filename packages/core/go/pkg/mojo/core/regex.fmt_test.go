package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex_Parse(t *testing.T) {
	regex := &Regex{}
	err := regex.Parse(regexExpr)

	assert.NoError(t, err)
	assert.Equal(t, regexExpr, regex.Expression)
}
