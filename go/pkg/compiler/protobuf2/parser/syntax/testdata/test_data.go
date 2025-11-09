package testdata

import (
	_ "embed"
)

//go:embed demo.proto
var DemoProto string

//go:embed extend.proto
var ExtendProto string
