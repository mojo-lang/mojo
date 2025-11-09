package data

import "github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"

type Service = data.Service

type ServicePackage struct {
	PackageName            string
	FullPackageName        string
	JavaPackageName        string
	JavaServicePackageName string // java service package name

	Services []*Service
}

type Data struct {
	Packages []*ServicePackage
}
