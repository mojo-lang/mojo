package strcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamel(t *testing.T) {
	RegisterIrregularCaseRule("IPV4", "Ipv4")
	RegisterIrregularCaseRule("B2B", "B2b")

	for _, c := range cases {
		assert.Equal(t, c[0], ToCamel(c[1]))
	}
}

func TestToLowerCamel(t *testing.T) {
	RegisterIrregularCaseRule("IPV4", "Ipv4")
	RegisterIrregularCaseRule("B2B", "B2b")

	for _, c := range cases {
		assert.Equal(t, c[1], ToLowerCamel(c[0]))
	}
}
