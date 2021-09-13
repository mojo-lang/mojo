package compiler

type Data struct {
	BoxedArrays       []*BoxedArray
	BoxedDictionaries []*BoxedDictionary
	BoxedStrings      []*BoxedString
	BoxedUnions       []*BoxedUnion
	Enums             []*Enum
	GoMod             *GoMod

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
