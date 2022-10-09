package data

type EnumItem struct {
	RawValue string
	Value    string
	Number   int
}

type Enum struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string
	WrapName      string
	CaseStyle     string

	DefaultItem int
	Items       []EnumItem
}

func (e *Enum) GetPackageName() string {
	return e.PackageName
}

func (e *Enum) GetFullName() string {
	return e.FullName
}
