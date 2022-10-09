package gokit

import (
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator"
	util2 "github.com/mojo-lang/mojo/go/pkg/util"
)

type Options = generator.Options

func GenerateClient(ds *data.Service, options Options) error {
	files, err := options.GenerateClient(ds)
	if err != nil {
		return err
	}

	return generateFiles(files, options.Output)
}

func GenerateService(ds *data.Service, options Options) error {
	files, err := options.GenerateService(ds)
	if err != nil {
		return err
	}

	return generateFiles(files, options.Output)
}

func generateFiles(files []*util2.GeneratedFile, output string) error {
	guard := &util2.PathGuard{
		OnlyClearGenerated: true,
		Suffixes:           []string{".go", ".mod", ".md", ".sh", ".yaml"},
	}
	for _, file := range files {
		file.SkipNoneGenerated = true
		if err := file.WriteTo(output, guard); err != nil {
			return err
		}
	}
	return nil
}
