package generator

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/java/generator/data"
	"path"

	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Generator struct {
	Files util.GeneratedFiles
}

func packageToPath(pkg string) string {
	return lang.PackageNameToPath(pkg)
}

func (g *Generator) generateWithPackage(sp *data.ServicePackage, fileName string, pathName string, template string) error {
	if len(template) > 0 {
		fileName = fileName + ".java"
		str, err := ApplyTemplate(fileName, template, sp, FuncMap)
		if err != nil {
			return err
		}

		g.Files = append(g.Files, &util.GeneratedFile{
			Name:    path.Join("src/main/java", packageToPath(sp.JavaPackageName), pathName, fileName),
			Content: str,
		})
	}
	return nil
}

func (g *Generator) generateWithService(sp *data.ServicePackage, service *data.Service, fileName string, pathName string, template string) error {
	if len(template) > 0 {
		fileName = strcase.ToCamel(service.Interface.Name+fileName) + ".java"
		str, err := ApplyTemplate(fileName, template, service, FuncMap)
		if err != nil {
			return err
		}

		g.Files = append(g.Files, &util.GeneratedFile{
			Name:    path.Join("src/main/java", packageToPath(sp.JavaServicePackageName), pathName, fileName),
			Content: str,
		})
	}
	return nil
}

func (g *Generator) Generate(d *data.Data) error {
	for _, sp := range d.Packages {
		err := g.generateWithPackage(sp, "ServiceNameConstants", "v1/constant/", javaServiceNameConstantsFile)
		if err != nil {
			return err
		}

		for _, service := range sp.Services {
			err = g.generateWithService(sp, service, "_http", "", javaServiceHttpFile)
			if err != nil {
				return err
			}

			err = g.generateWithService(sp, service, "_http_fallback_factory", "factory", javaServiceHttpFallbackFactoryFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
