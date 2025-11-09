package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseStyle_Format(t *testing.T) {
	var cs CaseStyle
	assert.Empty(t, cs.Format())
}
