module github.com/mojo-lang/mojo/packages/http/go

go 1.24.0

replace (
	github.com/mojo-lang/mojo/packages/core/go => ../../core/go
	github.com/mojo-lang/mojo/packages/document/go => ../../document/go
	github.com/mojo-lang/mojo/packages/lang/go => ../../lang/go
)

require (
	github.com/json-iterator/go v1.1.12
	github.com/mojo-lang/mojo/packages/core/go v0.0.0
	github.com/mojo-lang/mojo/packages/lang/go v0.0.0
	github.com/stretchr/testify v1.11.1
	google.golang.org/protobuf v1.36.10
)

require (
	github.com/alecthomas/chroma/v2 v2.10.0 // indirect
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/go-swiss/fonts v0.0.0-20230807175105-90067c2f5042 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jellydator/ttlcache/v3 v3.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mojo-lang/mojo/packages/document/go v0.0.0 // indirect
	github.com/phpdave11/gofpdf v1.4.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stephenafamo/goldmark-pdf v0.4.1 // indirect
	github.com/yuin/goldmark v1.7.13 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.31.0 // indirect
)
