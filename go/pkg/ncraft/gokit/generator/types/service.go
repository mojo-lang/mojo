package types

// Service is the top-level struct for the definition of a service.
type Service struct {
	// PkgName will be the package name of the go file(s) analyzed. So if a
	// Go file contained "package authz", then PkgName will be "authz". If
	// multiple Go files are analyzed, it will be the package name of the last
	// go file analyzed.
	PkgName     string
	FullPkgName string

	/// all the interface names in the same package
	AllInterfaceNames []string

	ApiImportPath string

	ImportPaths   []string
	ImportStructs []string
	ImportEnums   []string

	Messages []*Message
	Enums    []*Enum

	// Interface contains the sole service for this Service
	Interface *Interface
}
