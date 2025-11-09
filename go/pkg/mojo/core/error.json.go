package core

func init() {
	RegisterJSONFieldEncoder("core.Error", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
	RegisterJSONFieldDecoder("core.Error", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
}
