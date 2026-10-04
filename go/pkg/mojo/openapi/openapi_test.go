package openapi

import (
	_ "embed"
	"os"
	"path/filepath"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/api-docs_v2.json
var apiV2File []byte

func TestOpenAPI_GetOperation(t *testing.T) {
	o := New()
	o.Paths.Vals["/path/to/service"] = &PathItem{
		Get: &Operation{
			OperationId: "test_operation_id",
			Summary:     "test_operation_summary",
		},
		Post: &Operation{
			OperationId: "test_operation_id_2",
			Summary:     "test_operation_2_summary",
		}}

	p, m, api, e := o.GetOperation("test_operation_id")
	assert.NoError(t, e)
	assert.Equal(t, "/path/to/service", p)
	assert.Equal(t, "GET", m)
	assert.Equal(t, "test_operation_summary", api.Summary)
}

func TestOpenAPIEmptyAndLookupBoundaries(t *testing.T) {
	api := New()
	require.NotNil(t, api.Info)
	require.NotNil(t, api.Paths.Vals)
	require.NotNil(t, api.Components.Schemas)
	for _, candidate := range []*OpenAPI{nil, {}, {Paths: &Paths{Vals: map[string]*PathItem{"/nil": nil}}}} {
		assert.Nil(t, candidate.GetPath("/missing"))
		assert.NotPanics(t, func() {
			path, method, operation, err := candidate.GetOperation("missing")
			assert.True(t, core.IsNotFoundError(err))
			assert.Empty(t, path)
			assert.Empty(t, method)
			assert.Nil(t, operation)
			candidate.GenerateExample()
			candidate.SupplementExample()
		})
	}
	assert.False(t, (*OpenAPI)(nil).CheckVersion("3"))
	assert.False(t, (&OpenAPI{}).CheckVersion("3"))
	assert.False(t, api.CheckVersion("invalid"))
	item := &PathItem{}
	for _, method := range []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "HEAD", "PATCH", "TRACE"} {
		item.SetOperation(method, &Operation{OperationId: method})
	}
	api.Paths.Vals["/pets"] = item
	for _, method := range []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "HEAD", "PATCH", "TRACE"} {
		path, gotMethod, operation, err := api.GetOperation(method)
		require.NoError(t, err)
		assert.Equal(t, "/pets", path)
		assert.Equal(t, method, gotMethod)
		assert.Equal(t, method, operation.OperationId)
	}
	parameter := &Parameter{Schema: exampleSchema()}
	item.Get.Parameters = []*ReferenceableParameter{NewReferenceableParameter(parameter)}
	api.GenerateExample()
	assert.Equal(t, "generated", parameter.Example.GetString())
}

func TestOpenAPIParseFileFormatsAndErrors(t *testing.T) {
	const jsonDocument = `{"openapi":"3.0.3","info":{"title":"Pets","version":"1.0"},"paths":{}}`
	const yamlDocument = "openapi: 3.0.3\ninfo:\n  title: Pets\n  version: '1.0'\npaths: {}\n"
	for _, tt := range []struct{ name, content string }{{"api.json", jsonDocument}, {"api.yaml", yamlDocument}, {"api.yml", yamlDocument}} {
		filename := filepath.Join(t.TempDir(), tt.name)
		require.NoError(t, os.WriteFile(filename, []byte(tt.content), 0600))
		api, err := ParseFile(filename)
		require.NoError(t, err)
		assert.Equal(t, "Pets", api.Info.Title)
		assert.True(t, api.CheckVersion("3.0.3"))
	}
	for _, input := range []string{jsonDocument, yamlDocument} {
		api, err := Parse([]byte(input))
		require.NoError(t, err)
		assert.Equal(t, "Pets", api.Info.Title)
	}
	for _, input := range []string{`{"openapi":`, `{"openapi":"3.0.3"}`} {
		api, err := Parse([]byte(input))
		assert.Error(t, err)
		assert.Nil(t, api)
	}
	for _, name := range []string{"missing.json", "unsupported.txt"} {
		api, err := ParseFile(filepath.Join(t.TempDir(), name))
		assert.Error(t, err)
		assert.Nil(t, api)
	}
}

func TestOpenAPI_IsVersion(t *testing.T) {
	o := New()
	o.Openapi = &core.Version{Major: 3, Minor: 0, Patch: 3}

	assert.True(t, o.CheckVersion("3"))
	assert.True(t, o.CheckVersion("3.0"))
	assert.True(t, o.CheckVersion("3.0.3"))

	assert.False(t, o.CheckVersion("2"))
	assert.False(t, o.CheckVersion("3.1"))
	assert.False(t, o.CheckVersion("3.0.0"))
}

func TestOpenAPI_SupplementExample(t *testing.T) {
	o := New()
	o.Paths.Vals["/path/to/service"] = &PathItem{
		Get: &Operation{
			Summary: "test_operation_summary",
			Parameters: []*ReferenceableParameter{
				NewReferenceableParameter(&Parameter{
					Name:   "url",
					In:     Parameter_LOCATION_QUERY,
					Schema: NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING}),
				})},
		}}

	o.SupplementExample()
	assert.NotNil(t, o.GetPath("/path/to/service").GetGet().GetParameter(0, nil).GetExample())
}

func TestOpenAPI_SupplementExample2(t *testing.T) {
	api := &OpenAPI{}
	content, err := os.ReadFile("./testdata/supplement_example_test.json")
	if assert.NoError(t, err) {
		err = jsoniter.Unmarshal(content, api)
		if assert.NoError(t, err) {
			api.SupplementExample()
		}
	}

}

func TestOpenAPI_Parse(t *testing.T) {
	api, err := Parse(apiV2File)
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), api.Openapi.Major)
}

func TestOpenAPI_ParseFile(t *testing.T) {
	api, err := ParseFile("testdata/api-docs_v2.json")
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), api.Openapi.Major)
}
