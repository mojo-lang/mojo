package openapi

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestPathsCodec_Decode(t *testing.T) {
	paths := &Paths{Vals: map[string]*PathItem{
		"/foo/bar": {Get: &Operation{Summary: "foo.bar"}},
	}}

	str, err := jsoniter.MarshalToString(paths)
	assert.NoError(t, err)

	ps := &Paths{}
	err = jsoniter.UnmarshalFromString(str, ps)
	assert.NoError(t, err)
	assert.Equal(t, paths.Vals["/foo/bar"].Get.Summary, ps.Vals["/foo/bar"].Get.Summary)
}
