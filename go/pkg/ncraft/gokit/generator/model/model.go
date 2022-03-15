package model

import (
    "bytes"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/model/templates"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/render"
    "github.com/pkg/errors"
    "io/ioutil"
    "strings"
)

// TemplatePath is the path to the entity gotemplate file.
const TemplatePath = "internal/model/ENTITY_model.go.tmpl"

type Model struct {
}

func (m Model) Generate(path string, data *render.Data) ([]*util.GeneratedFile, error) {
    if path != TemplatePath {
        return nil, errors.Errorf("cannot render unknown file: %q", path)
    }

    var files []*util.GeneratedFile

    for k, v := range data.Entities {
        reader, err := render.ApplyTemplate(templates.EntityModel, "Model", v, render.FuncMap)
        if err != nil {
            return nil, err
        }

        codeBytes, err := ioutil.ReadAll(reader)
        if err != nil {
            return nil, err
        }

        formattedCode := render.FormatGoCode(codeBytes)
        files = append(files, &util.GeneratedFile{
            Name:              strings.Replace(path, "ENTITY", k, 1),
            Reader:            bytes.NewReader(formattedCode),
            SkipIfExist:       false,
            SkipNoneGenerated: false,
        })
    }

    return files, nil
}
