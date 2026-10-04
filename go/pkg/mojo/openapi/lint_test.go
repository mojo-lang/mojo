package openapi

import (
	_ "embed"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/api-docs-1.json
var openapiFile []byte

func TestLint(t *testing.T) {
	type args struct {
		spec []byte
	}
	tests := []struct {
		name      string
		args      args
		doNotWant *LintResult
	}{
		{
			name: "lint-1",
			args: args{
				spec: openapiFile,
			},
			doNotWant: &LintResult{Valid: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EasyLint(tt.args.spec); reflect.DeepEqual(got, tt.doNotWant) {
				t.Errorf("EasyLint() = %v", got)
			}
		})
	}
}

func TestEasyLintOperationResults(t *testing.T) {
	const valid = `{"swagger":"2.0","info":{"title":"Pets","version":"1"},"paths":{"/pets":{"get":{"parameters":[{"name":"id","in":"query","type":"string"}],"responses":{"200":{"description":"OK"}}}}}}`
	result := EasyLint([]byte(valid))
	require.NotNil(t, result)
	assert.True(t, result.Valid)
	require.Len(t, result.Operations, 1)
	assert.Equal(t, "/pets", result.Operations[0].Path)
	assert.Equal(t, "get", result.Operations[0].Method)
	assert.True(t, result.Operations[0].Valid)
	const duplicate = `{"swagger":"2.0","info":{"title":"Pets","version":"1"},"paths":{"/healthy":{"get":{"responses":{"200":{"description":"OK"}}}},"/pets":{"get":{"parameters":[{"name":"id","in":"query","type":"string"},{"name":"id","in":"query","type":"string"}],"responses":{"200":{"description":"OK"}}}}}}`
	result = EasyLint([]byte(duplicate))
	assert.False(t, result.Valid)
	require.Len(t, result.Operations, 2)
	assert.Equal(t, "/healthy", result.Operations[0].Path)
	assert.True(t, result.Operations[0].Valid)
	assert.Equal(t, "/pets", result.Operations[1].Path)
	assert.False(t, result.Operations[1].Valid)
	assert.NotEmpty(t, result.Operations[1].Description)
	refined, err := lintAndRefine([]byte(duplicate), jsoniter.Marshal, jsoniter.Unmarshal)
	require.NoError(t, err)
	assert.Equal(t, jsoniter.InvalidValue, jsoniter.Get(refined, "paths", "/pets", "get").ValueType())
	assert.Equal(t, jsoniter.ObjectValue, jsoniter.Get(refined, "paths", "/healthy", "get").ValueType())
	assert.False(t, EasyLint([]byte(`not: [valid`)).Valid)
}
