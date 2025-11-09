package testdata

import _ "embed"

//go:embed mine
var MineData []byte

//go:embed charset-mine
var CharsetMineData []byte

//go:embed nested-mime
var NestedMineData []byte
