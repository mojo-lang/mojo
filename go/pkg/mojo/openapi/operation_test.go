package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testReference(t *testing.T, path string) *Reference {
	t.Helper()
	url, err := core.NewUrl("#" + path)
	require.NoError(t, err)
	return &Reference{Ref: url}
}

func TestOperationParameters(t *testing.T) {
	query := &Parameter{Name: "query", In: Parameter_LOCATION_QUERY}
	header := &Parameter{Name: "header", In: Parameter_LOCATION_HEADER}
	op := &Operation{Parameters: []*ReferenceableParameter{
		NewReferenceableParameter(query), nil, NewReferencedParameter(testReference(t, ParameterReferenceRoot+"header")), NewReferenceableParameter(header),
	}}
	assert.Same(t, query, op.GetParameter(0, nil))
	assert.Same(t, header, op.GetParameter(2, map[string]*Parameter{"header": header}))
	for _, index := range []int{-1, 1, 2, 4} {
		assert.Nil(t, op.GetParameter(index, nil))
	}
	assert.Nil(t, (*Operation)(nil).GetParameter(0, nil))
	assert.Equal(t, map[Parameter_Location][]*Parameter{Parameter_LOCATION_QUERY: {query}, Parameter_LOCATION_HEADER: {header}}, op.GetLocationParameters())
	assert.Empty(t, (*Operation)(nil).GetLocationParameters())
}

func TestOperationExamplesThroughReferences(t *testing.T) {
	parameter := &Parameter{Schema: exampleSchema()}
	requestMedia, responseMedia := &MediaType{Schema: exampleSchema()}, &MediaType{Schema: exampleSchema()}
	components := &Components{
		Parameters:    map[string]*Parameter{"id": parameter},
		RequestBodies: map[string]*RequestBody{"body": {Content: map[string]*MediaType{core.ApplicationJson: requestMedia, "empty": nil}}},
		Responses:     map[string]*Response{"ok": {Content: map[string]*MediaType{core.ApplicationJson: responseMedia, "empty": nil}}},
	}
	op := &Operation{
		Parameters:  []*ReferenceableParameter{nil, NewReferencedParameter(testReference(t, ParameterReferenceRoot+"id"))},
		RequestBody: NewReferencedRequestBody(testReference(t, RequestBodyReferenceRoot+"body")),
		Responses:   &Responses{Vals: map[string]*ReferenceableResponse{"200": NewReferencedResponse(testReference(t, ResponseReferenceRoot+"ok")), "default": nil}},
	}
	op.SupplementExample(components)
	for _, value := range []*core.Value{parameter.Example, requestMedia.Example, responseMedia.Example} {
		require.NotNil(t, value)
		assert.Equal(t, "generated", value.GetString())
	}
	explicit := core.NewStringValue("explicit")
	parameter.Example, requestMedia.Example, responseMedia.Example = explicit, explicit, explicit
	op.SupplementExample(components)
	assert.Same(t, explicit, parameter.Example)
	assert.Same(t, explicit, requestMedia.Example)
	assert.Same(t, explicit, responseMedia.Example)
	op.GenerateExample(components)
	for _, value := range []*core.Value{parameter.Example, requestMedia.Example, responseMedia.Example} {
		assert.Equal(t, "generated", value.GetString())
	}
	assert.NotPanics(t, func() {
		(*Operation)(nil).GenerateExample(nil)
		(*Operation)(nil).SupplementExample(nil)
		(&Operation{}).GenerateExample(nil)
		(&Operation{}).SupplementExample(nil)
	})
}
