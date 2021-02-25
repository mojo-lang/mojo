package _go

import (
	"github.com/mojo-lang/mojo/go/pkg/util"
	"io/ioutil"
	path2 "path"
)

type Generator struct {
	Files util.CodeGeneratedFiles
}

func NewGenerator(files util.CodeGeneratedFiles) *Generator {
	return &Generator{Files: files}
}

func (g *Generator) Generate(output string) error {
	for _, f := range g.Files {
		if len(f.Name) > 0 && len(f.Content) > 0 {
			name := path2.Join(output, f.Name)
			path := path2.Dir(name)
			util.CreateDir(path)
			ioutil.WriteFile(name, []byte(f.Content), 0666)
		}
	}
	return nil
}
