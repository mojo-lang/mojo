package boot

import (
	"github.com/mojo-lang/mojo/go/pkg/ncraft/boot/generator"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Options = generator.Options

func GenerateService(ds *data.Service, options Options) error {
	files, err := options.GenerateService(ds)
	if err != nil {
		return err
	}

	return generateFiles(files, options.Output)
}

func generateFiles(files []*util.GeneratedFile, output string) error {
	guard := &util.PathGuard{
		OnlyClearGenerated: true,
		Suffixes:           []string{".java"},
	}
	for _, file := range files {
		file.SkipIfUserCodeMixed = true
		if err := file.WriteTo(output, guard); err != nil {
			return err
		}
	}
	return nil
}
