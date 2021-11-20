module github.com/mojo-lang/mojo/go

go 1.16

replace (
	github.com/mojo-lang/core/go => ../../core/go
	github.com/mojo-lang/lang/go => ../../lang/go
	github.com/mojo-lang/openapi/go => ../../openapi/go
)

require (
	github.com/antlr/antlr4 v0.0.0-20211106181442-e4c1a74c66bd
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211106181442-e4c1a74c66bd // indirect
	github.com/gertd/go-pluralize v0.1.7
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/iancoleman/strcase v0.2.0
	github.com/json-iterator/go v1.1.12
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mojo-lang/core/go v0.0.0-20211117011141-4ab29472fc29 // indirect
	github.com/mojo-lang/document/go v0.0.0-20211117011435-ec994bc588a8 // indirect
	github.com/mojo-lang/lang/go v0.0.0-20211117011903-247e727c9b83 // indirect
	github.com/mojo-lang/openapi/go v0.0.0-20211117013305-c1f3afbc00e0 // indirect
	github.com/mojo-lang/yaml v0.0.0-20210912132100-a6a60fbf6f20
	github.com/otiai10/copy v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
)
