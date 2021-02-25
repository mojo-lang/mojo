package document

import (
	"github.com/mojo-lang/document/go/pkg/markdown"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"io/ioutil"
	"os"
	path2 "path"
)

type Generator struct {
	Documents Documents

	Files map[string][]byte
}

func NewGenerator(documents Documents) *Generator {
	return &Generator{
		Documents: documents,
		Files:     make(map[string][]byte),
	}
}

func (g *Generator) Generate(output string) error {
	if err := g.generateMarkdown(); err != nil {
		return err
	}

	if err := g.writeFiles(output); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateMarkdown() error {
	md := markdown.New()
	for name, document := range g.Documents {
		m, err := md.RenderToString(document)
		if err != nil {
			return err
		}
		g.Files[name+".md"] = []byte(m)
	}
	return nil
}

func (g *Generator) cleanFiles() error {
	return nil
}

func (g *Generator) writeFiles(dir string) error {
	if util.IsExist(dir) {
		os.RemoveAll(dir)
	}

	for name, file := range g.Files {
		name = path2.Join(dir, name)
		path := path2.Dir(name)
		util.CreateDir(path)
		ioutil.WriteFile(name, file, 0666)
	}
	return nil
}
