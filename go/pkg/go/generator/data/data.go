package data

type Data struct {
	BoxedArrays    []*BoxedArray
	BoxedMaps      []*BoxedMap
	BoxedUnions    []*BoxedUnion
	Enums          []*Enum
	DbJSONs        []*DbJSON
	FormatJSONs    []*FormatJSON
	ArrayResponses []*ArrayResponse
	GoMod          *GoMod

	NameIndex map[string]bool
}

func NewData() *Data {
	return &Data{
		NameIndex: make(map[string]bool),
	}
}

func (d *Data) IsExist(name string) bool {
	_, ok := d.NameIndex[name]
	return ok
}
