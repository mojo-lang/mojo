package template

import (
	"embed"
	_ "embed"
	gkt "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/templates"
)

//go:embed service-java/*
var services embed.FS

type FileGetter = gkt.FileGetter

func Service(path string) ([]byte, error) {
	return gkt.FileContent(services, path, "service-java/")
}

func ServiceNames() []string {
	return gkt.FileNames(services, "service-java/")
}
