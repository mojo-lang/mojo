package data

import (
	"text/template"
)

// Service is passed to templates as the executing struct; its fields
// and methods are used to modify the templates
type Service struct {
	Version     string
	VersionDate string

	// PackageName is the name of the package containing the service definition
	PackageName     string
	PackageFullName string

	CombinedAPI bool

	ImportedMessages []*Message
	ImportedEnums    []*Enum
	Entities         []*Message
	AllInterfaces    []*Interface

	// GRPC/Protobuf service, with all parameters and return values accessible
	Interface *Interface

	FuncMap template.FuncMap

	Go         *GoService
	Java       *JavaService
	Extensions map[string]interface{}
}

type GoService struct {
	PackageName string

	// the service repository path
	RepositoryPath string

	ApiRepositoryPath string

	// import path for .pb.go files containing service structs
	ApiImportPath string

	ImportedTypePaths []string
}

type JavaService struct {
	PackageName string
}

func (s *Service) HasImported() bool {
	return s != nil && (len(s.ImportedMessages) > 0 || len(s.ImportedEnums) > 0)
}
