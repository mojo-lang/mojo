package model

import (
	"bytes"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	_go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/model/templates"
	util2 "github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

// TemplatePath is the path to the entity gotemplate file.
const TemplatePath = "internal/model/ENTITY_model.go.tmpl"

type Model struct {
}

func (m Model) Generate(path string, service *data.Service) ([]*util2.GeneratedFile, error) {
	if path != TemplatePath {
		return nil, errors.Errorf("cannot render unknown file: %q", path)
	}

	var files []*util2.GeneratedFile

	for _, v := range service.Entities {
		reader, err := util2.ApplyTemplate("Model", templates.EntityModel, v, service.FuncMap)
		if err != nil {
			return nil, err
		}

		codeBytes, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		formattedCode := _go.FormatCodeBytes(codeBytes)
		files = append(files, &util2.GeneratedFile{
			Name:              strings.Replace(path, "ENTITY", v.Name, 1),
			Reader:            bytes.NewReader(formattedCode),
			SkipIfExist:       false,
			SkipNoneGenerated: false,
		})
	}

	return files, nil
}
