module github.com/mojo-lang/mojo/go

go 1.16

replace (
	github.com/mojo-lang/core/go => ../../core/go
	github.com/mojo-lang/document/go => ../../document/go
	github.com/mojo-lang/lang/go => ../../lang/go
	github.com/mojo-lang/openapi/go => ../../openapi/go
	github.com/mojo-lang/yaml/go => ../../yaml/go
)

require (
	github.com/antlr/antlr4 v0.0.0-20201010232522-9e64dfc6e99f
	github.com/gertd/go-pluralize v0.1.7
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.1
	github.com/iancoleman/strcase v0.1.3
	github.com/json-iterator/go v1.1.11
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mojo-lang/core/go v0.0.0-00010101000000-000000000000
	github.com/mojo-lang/document/go v0.0.0-00010101000000-000000000000
	github.com/mojo-lang/lang/go v0.0.0-00010101000000-000000000000
	github.com/mojo-lang/openapi/go v0.0.0-00010101000000-000000000000
	github.com/mojo-lang/yaml/go v0.0.0-00010101000000-000000000000
	github.com/otiai10/copy v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
)
