package protobuf

import (
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/generator"
)

type Generator struct {
	*generator.Generator
}

func NewGenerator(files []*descriptor.FileDescriptor) *Generator {
	return &Generator{
		Generator: generator.New(files),
	}
}

func (g *Generator) Generate(output string) ([]*descriptor.FileDescriptor, error) {
	g.GenerateAllFiles().WriteAllFiles(output)
	return g.GetGeneratedFiles(), nil
}
