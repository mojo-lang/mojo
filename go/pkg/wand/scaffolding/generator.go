package scaffolding

import (
	"errors"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"io/ioutil"
	path2 "path"
	"strings"
)

type Generator struct {
}

func (g *Generator) Generate(data *Data, output string) error {
	var files []*util.CodeGeneratedFile

	applyTemplate := func(name string, tmpl string, tmplData interface{}, filename string) error {
		if tmplData == nil {
			return errors.New("the input template data is nil")
		}

		str, err := ApplyTemplate(name, tmpl, tmplData, FuncMap)
		if err != nil {
			return err
		}
		files = append(files, &util.CodeGeneratedFile{
			Name:    path2.Join(output, data.Wand.Name, filename),
			Content: str,
		})
		return nil
	}

	if err := applyTemplate("wandFile", wandFile, data.Wand, "wand.mojo"); err != nil {
		return err
	}
	if err := applyTemplate("gitignoreFile", gitignoreFile, data.Wand, ".gitignore"); err != nil {
		return err
	}
	if err := applyTemplate("readmeFile", readmeFile, data.Wand, "README.md"); err != nil {
		return err
	}

	for _, f := range files {
		if len(f.Name) > 0 && len(f.Content) > 0 {
			name := f.Name
			path := path2.Dir(name)
			util.CreateDir(path)
			ioutil.WriteFile(name, []byte(f.Content), 0666)
		}
	}

	// mkdir mojo folders
	mojoPath := path2.Join(output, data.Wand.Name, "mojo", strings.ReplaceAll(data.Wand.FullName, ".", "/"))
	util.CreateDir(mojoPath)

	return nil
}
