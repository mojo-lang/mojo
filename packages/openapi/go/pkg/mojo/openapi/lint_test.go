package openapi

import (
	_ "embed"
	"reflect"
	"testing"
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
