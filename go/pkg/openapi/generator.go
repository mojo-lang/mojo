package openapi

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"github.com/mojo-lang/yaml/go/pkg/yaml"
	"io/ioutil"
	"os"
	path2 "path"
	"strings"
)

// generate the following files:
// 1. openapi.yaml
// 2. struct.schema.json
type Generator struct {
	*OpenAPIs

	Package *lang.Package
	Files   map[string][]byte
}

func NewGenerator(pkg *lang.Package, apis map[string]*openapi.OpenAPI, components *openapi.Components) *Generator {
	return &Generator{
		OpenAPIs: &OpenAPIs{
			APIs:       apis,
			Components: components,
		},
		Package: pkg,
		Files: make(map[string][]byte),
	}
}

func (g *Generator) Generate(output string) error {
	if err := g.generateApi(); err != nil {
		return err
	}

	if err := g.generateSchema(); err != nil {
		return err
	}

	if err := g.writeFiles(output); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateApi() error {
	for name, api := range g.APIs {
		y, err := yaml.Marshal(api)
		if err != nil {
			return err
		}
		g.Files[name+".yaml"] = y
	}
	return nil
}

func toSchemaFileName(name string) string {
	return lang.TypeNameToFileName(name) + ".schema.json"
}

func (g *Generator) generateSchema() error {
	for name, schema := range g.Components.Schemas {
		if !strings.HasPrefix(name, g.Package.FullName) {
			continue
		}

		j, err := jsoniter.Marshal(schema)
		if err != nil {
			return err
		}
		g.Files[toSchemaFileName(name)] = j
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
