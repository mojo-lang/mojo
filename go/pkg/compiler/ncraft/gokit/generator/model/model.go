package model

import (
	"fmt"
	"go/format"
	"io"
	"path"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/model/templates"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
)

const TemplatePath = "pkg/model/ENTITY_model.go.tmpl"

type Model struct{}

type entityModel struct {
	Name, ImportPath, KeyName, KeyType string
	Fields                             []modelField
}

type modelField struct {
	Name, Type string
	Repeated   bool
}

func (m Model) Generate(templatePath string, service *data.Service) ([]*util.GeneratedFile, error) {
	if templatePath != TemplatePath {
		return nil, fmt.Errorf("cannot render unknown file: %q", templatePath)
	}
	return m.GenerateEntities(service.Entities)
}

func (m Model) GenerateEntities(entities []*data.Message) ([]*util.GeneratedFile, error) {
	var files []*util.GeneratedFile
	seen := make(map[string]string)
	for _, entity := range entities {
		if entity == nil || entity.Decl == nil {
			return nil, fmt.Errorf("cannot generate model without an entity declaration")
		}
		fullName := entity.Decl.GetFullName()
		name := lang.GetTypeGoTypeName(fullName)
		fileName := path.Join("pkg/model", strcase.ToSnake(name)+"_model.go")
		if previous, ok := seen[fileName]; ok {
			if previous == fullName {
				continue
			}
			return nil, fmt.Errorf("model name collision: %s and %s both generate %s", previous, fullName, fileName)
		}
		seen[fileName] = fullName
		if entity.KeyField == nil {
			return nil, fmt.Errorf("entity %s has no primary key", fullName)
		}
		if entity.Go == nil || entity.Go.ImportPath == "" {
			return nil, fmt.Errorf("entity %s has no Go import path", fullName)
		}
		keyType := scalarKeyType(entity.KeyField.GetType())
		if keyType == "" {
			return nil, fmt.Errorf("entity %s: key %s must be a string or integer", fullName, entity.KeyField.Name)
		}
		model := entityModel{Name: name, ImportPath: entity.Go.ImportPath, KeyName: strcase.ToCamel(entity.KeyField.Name), KeyType: keyType}
		for _, field := range entity.Decl.GetAllFields() {
			typ := field.GetType()
			repeated := typ.IsArrayType()
			if repeated && len(typ.GenericArguments) > 0 {
				typ = typ.GenericArguments[0]
			}
			kind := "JSON"
			switch typ.GetFullName() {
			case "mojo.core.Bool":
				kind = "Bool"
			case "mojo.core.Float32", "mojo.core.Float64":
				kind = "Float"
			case "mojo.core.String":
				kind = "String"
			case "mojo.core.Bytes":
				kind = "Bytes"
			case "mojo.core.Timestamp", "mojo.core.Date", "mojo.core.DateTime":
				kind = "Datetime"
			default:
				if scalarKeyType(typ) != "" {
					kind = "Integer"
				}
			}
			model.Fields = append(model.Fields, modelField{Name: field.Name, Type: kind, Repeated: repeated})
		}
		file, err := render(fileName, templates.EntityModel, model)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	if len(files) > 0 {
		file, err := render("pkg/model/db.go", templates.DB, nil)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name < files[j].Name })
	return files, nil
}

func scalarKeyType(typ *lang.NominalType) string {
	switch typ.GetFullName() {
	case "mojo.core.String":
		return "string"
	case "mojo.core.Int8", "mojo.core.Int16", "mojo.core.Int32":
		return "int32"
	case "mojo.core.Int", "mojo.core.Int64":
		return "int64"
	case "mojo.core.UInt8", "mojo.core.UInt16", "mojo.core.UInt32":
		return "uint32"
	case "mojo.core.UInt", "mojo.core.UInt64":
		return "uint64"
	}
	return ""
}

func render(name, source string, value interface{}) (*util.GeneratedFile, error) {
	reader, err := util.ApplyTemplate(name, source, value, util.FuncMap)
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	content, err = format.Source([]byte(strings.TrimSpace(string(content)) + "\n"))
	if err != nil {
		return nil, fmt.Errorf("format model %s: %w", name, err)
	}
	return &util.GeneratedFile{Name: name, Content: string(content), SkipIfUserCodeMixed: true}, nil
}
