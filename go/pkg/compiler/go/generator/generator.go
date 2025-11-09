package generator

import (
	"github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/generator"
	util2 "github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

type Generator struct {
	Data  *data.Data
	Files util2.GeneratedFiles
}

func NewGenerator(files util2.GeneratedFiles, data *data.Data) *Generator {
	return &Generator{
		Files: files,
		Data:  data,
	}
}

func (g *Generator) Generate(output string) error {
	gen := &generator.Generator{}
	err := gen.Generate(g.Data)
	if err != nil {
		return err
	}
	g.Files = append(g.Files, gen.Files...)

	guard := &util2.PathGuard{
		OnlyClearGenerated: true,
		Suffixes:           []string{".go"},
	}

	for _, f := range g.Files {
		f.SkipIfUserCodeMixed = true
		if err = f.WriteTo(output, guard); err != nil {
			return err
		}
	}
	return nil
}
