package _go

import (
	"github.com/mojo-lang/mojo/go/pkg/go/compiler"
	"github.com/mojo-lang/mojo/go/pkg/go/generator"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"io/ioutil"
	path2 "path"
)

type Generator struct {
	Data  *compiler.Data
	Files util.CodeGeneratedFiles
}

func NewGenerator(files util.CodeGeneratedFiles, data *compiler.Data) *Generator {
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

	guard := &util.PathGuard{}
	for _, f := range g.Files {
		if len(f.Name) > 0 && len(f.Content) > 0 {
			name := path2.Join(output, f.Name)
			path := path2.Dir(name)

			if err = guard.Check(path); err != nil {
				return err
			}

			if util.IsExist(name) {
				if f.SkipIfExist {
					continue
				}

				if !util.IsGeneratedFile(name) {
					continue
				}
			}

			ioutil.WriteFile(name, []byte(f.Content), 0666)
		}
	}
	return nil
}
