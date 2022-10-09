package data

type FormatJSON struct {
	GoPackageName string

	PackageName string

	Name              string
	FullName          string
	EnclosingName     string
	EnclosingFullName string

	EncodingAsStruct bool
}

type FormatJSONs []*FormatJSON

func (j *FormatJSON) GetPackageName() string {
	return j.PackageName
}

func (j *FormatJSON) GetFullName() string {
	return j.FullName
}
