package generator

import (
	"github.com/mojo-lang/mojo/go/pkg/protobuf/generator/generator"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Generator struct {
	*generator.Generator
}

func NewGenerator(files []*descriptor.File) *Generator {
	return &Generator{
		Generator: generator.New(files),
	}
}

func (g *Generator) Generate(output string) ([]*descriptor.File, error) {
	g.GenerateAllFiles().WriteAllFiles(output)
	return g.GetGeneratedFiles(), nil
}
