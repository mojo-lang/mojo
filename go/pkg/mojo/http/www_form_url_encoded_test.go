package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWwwFormUrlencodedToJSON(t *testing.T) {
	json, err := WwwFormUrlencodedToJSON([]byte(`a=bb&foo_bar=baz&ids=12&ids=34`))
	assert.NoError(t, err)
	assert.NotEmpty(t, json)
}
