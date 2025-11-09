package testdata

import (
	_ "embed"
)

//go:embed addressbook.proto
var AddressBookProto string

//go:embed conformance.proto
var ConformanceProto string

//go:embed example.proto
var ExampleProto string

//go:embed helloworld.proto
var HelloWorldProto string

//go:embed helloworldreserved.proto
var HelloWorldReservedProto string
