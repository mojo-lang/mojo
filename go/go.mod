module github.com/mojo-lang/mojo/go

go 1.16

replace github.com/mojo-lang/mojo/go/pkg/mojo => ./pkg/mojo

require (
	github.com/fatih/structtag v1.2.0
	github.com/goccy/go-graphviz v0.0.9
	github.com/json-iterator/go v1.1.12
	github.com/mojo-lang/core/go v0.0.0-20220909081037-a971c08096fb
	github.com/mojo-lang/db/go v0.0.0-20220708073636-d85468d4ceca
	github.com/mojo-lang/document/go v0.0.0-20220504064505-9e47a78ac0ec
	github.com/mojo-lang/http/go v0.0.0-20220524115554-fc19f03dd3de
	github.com/mojo-lang/lang/go v0.0.0-20220929083528-e86973a7588b
	github.com/mojo-lang/mojo/go/pkg/mojo v0.0.0
	github.com/mojo-lang/openapi/go v0.0.0-20220504071403-44d2e722f238
	github.com/mojo-lang/protobuf/go v0.0.0-20220504071509-5bf0a646ba25
	github.com/mojo-lang/yaml/go v0.0.0-20220504064824-fa1822f5517c
	github.com/otiai10/copy v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/stretchr/testify v1.8.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/exp v0.0.0-20220921164117-439092de6870 // indirect
	google.golang.org/protobuf v1.28.1
)
