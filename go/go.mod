module github.com/mojo-lang/mojo/go

go 1.24.7

replace (
	github.com/mojo-lang/mojo/packages/core/go => ../packages/core/go
	github.com/mojo-lang/mojo/packages/db/go => ../packages/db/go
	github.com/mojo-lang/mojo/packages/document/go => ../packages/document/go
	github.com/mojo-lang/mojo/packages/geom/go => ../packages/geom/go
	github.com/mojo-lang/mojo/packages/http/go => ../packages/http/go
	github.com/mojo-lang/mojo/packages/lang/go => ../packages/lang/go
	github.com/mojo-lang/mojo/packages/openapi/go => ../packages/openapi/go
	github.com/mojo-lang/mojo/packages/protobuf/go => ../packages/protobuf/go
	github.com/mojo-lang/mojo/packages/rpc/go => ../packages/rpc/go
	github.com/mojo-lang/mojo/packages/yaml/go => ../packages/yaml/go
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.1
	github.com/edwin-luijten/go_mod_parser v0.0.0-20190307065647-27b9ee14b099
	github.com/fatih/structtag v1.2.0
	github.com/gertd/go-pluralize v0.2.1
	github.com/go-clang/clang-v10 v0.0.0-20211120055647-b59749ef6dbb
	github.com/goccy/go-graphviz v0.2.9
	github.com/json-iterator/go v1.1.12
	github.com/mojo-lang/mojo/packages/core/go v0.0.0
	github.com/mojo-lang/mojo/packages/db/go v0.0.0
	github.com/mojo-lang/mojo/packages/document/go v0.0.0
	github.com/mojo-lang/mojo/packages/geom/go v0.0.0
	github.com/mojo-lang/mojo/packages/http/go v0.0.0
	github.com/mojo-lang/mojo/packages/lang/go v0.0.0
	github.com/mojo-lang/mojo/packages/openapi/go v0.0.0
	github.com/mojo-lang/mojo/packages/protobuf/go v0.0.0
	github.com/mojo-lang/mojo/packages/rpc/go v0.0.0
	github.com/mojo-lang/mojo/packages/yaml/go v0.0.0
	github.com/otiai10/copy v1.14.1
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2
	github.com/stretchr/testify v1.11.1
	github.com/urfave/cli/v2 v2.27.7
	google.golang.org/protobuf v1.36.10
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/alecthomas/chroma/v2 v2.20.0 // indirect
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/brianvoe/gofakeit/v6 v6.28.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/daveshanley/vacuum v0.18.5 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/dop251/goja v0.0.0-20250630131328-58d95d85e994 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20250409162600-f7acab6894b0 // indirect
	github.com/flopp/go-findfont v0.1.0 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/getkin/kin-openapi v0.133.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/go-swiss/fonts v0.0.0-20230807175105-90067c2f5042 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/geo v0.0.0-20251014162054-f262919d8753 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20250501235452-c0086092b71a // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.6 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jellydator/ttlcache/v3 v3.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.32 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mmcloughlin/geohash v0.10.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oasdiff/yaml v0.0.0-20250309154309-f31be36b4037 // indirect
	github.com/oasdiff/yaml3 v0.0.0-20250309153720-d2182401db90 // indirect
	github.com/otiai10/mint v1.6.3 // indirect
	github.com/pb33f/doctor v0.0.39 // indirect
	github.com/pb33f/jsonpath v0.1.2 // indirect
	github.com/pb33f/libopenapi v0.28.0 // indirect
	github.com/pb33f/libopenapi-validator v0.6.3 // indirect
	github.com/pb33f/ordered-map/v2 v2.3.0 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/phpdave11/gofpdf v1.4.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smilextay/dm v0.0.0-20231121034257-b7bd60563f8b // indirect
	github.com/smilextay/gorm-dm v0.0.13 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/stephenafamo/goldmark-pdf v0.4.1 // indirect
	github.com/tetratelabs/wazero v1.8.1 // indirect
	github.com/woodsbury/decimal128 v1.3.0 // indirect
	github.com/xrash/smetrics v0.0.0-20250705151800-55b8f293f342 // indirect
	github.com/yuin/goldmark v1.7.13 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	go.yaml.in/yaml/v4 v4.0.0-rc.2 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/exp v0.0.0-20251009144603-d2f985daa21b // indirect
	golang.org/x/image v0.21.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.6.0 // indirect
	gorm.io/driver/postgres v1.6.0 // indirect
	gorm.io/driver/sqlite v1.6.0 // indirect
	gorm.io/gorm v1.31.0 // indirect
)
