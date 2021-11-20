package _go

import (
	"github.com/mojo-lang/mojo/go/pkg/go/compiler"
	"github.com/mojo-lang/mojo/go/pkg/go/generator"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Generator struct {
	Data  *compiler.Data
	Files util.GeneratedFiles
}

func NewGenerator(files util.GeneratedFiles, data *compiler.Data) *Generator {
	return &Generator{
		Files: files,
		Data:  data,
	}
}

func (g *Generator) Generate(output string) error {
	generator := &generator.Generator{}
	err := generator.Generate(g.Data)
	if err != nil {
		return err
	}
	g.Files = append(g.Files, generator.Files...)

	guard := &util.PathGuard{
		OnlyClearGenerated: true,
		Suffixes:           []string{".go"},
	}

	for _, f := range g.Files {
		f.SkipNoneGenerated = true
		if err = f.WriteTo(output, guard); err != nil {
			return err
		}
	}
	return nil
}
