package model

import (
	"bytes"
	"io"
	"strings"

	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	_go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/model/templates"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

// TemplatePath is the path to the entity gotemplate file.
const TemplatePath = "internal/model/ENTITY_model.go.tmpl"

type Model struct {
}

func (m Model) Generate(path string, service *data.Service) ([]*util.GeneratedFile, error) {
	if path != TemplatePath {
		return nil, errors.Errorf("cannot render unknown file: %q", path)
	}

	var files []*util.GeneratedFile

	for _, v := range service.Entities {
		reader, err := util.ApplyTemplate("Model", templates.EntityModel, v, service.FuncMap)
		if err != nil {
			return nil, err
		}

		codeBytes, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		formattedCode := _go.FormatCodeBytes(codeBytes)
		files = append(files, &util.GeneratedFile{
			Name:                strings.Replace(path, "ENTITY", v.Name, 1),
			Reader:              bytes.NewReader(formattedCode),
			SkipIfExist:         false,
			SkipIfUserCodeMixed: false,
		})
	}

	return files, nil
}
