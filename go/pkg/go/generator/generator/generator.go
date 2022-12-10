package generator

import (
	path2 "path"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	data2 "github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Generator struct {
	Files util.GeneratedFiles
}

func packageToPath(pkg string) string {
	return path2.Join("pkg", lang.PackageNameToPath(pkg))
}

func (g *Generator) generateDecl(decl data2.Decl, fileType string, template string) error {
	if len(template) > 0 {
		str, err := ApplyTemplate("goFile"+fileType, template, decl, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.GeneratedFile{
			Name:    path2.Join(packageToPath(decl.GetPackageName()), strcase.ToSnake(decl.GetFullName())+"."+fileType+".go"),
			Content: FormatCode(str),
		})
	}
	return nil
}

func (g *Generator) Generate(data *data2.Data) error {
	for _, decl := range data.BoxedArrays {
		err := g.generateDecl(decl, "json", goArrayJsonFile)
		if err != nil {
			return err
		}
	}

	for _, decl := range data.BoxedMaps {
		err := g.generateDecl(decl, "json", goMapJsonFile)
		if err != nil {
			return err
		}
	}

	for _, decl := range data.Enums {
		err := g.generateDecl(decl, "fmt", goEnumFmtFile)
		if err != nil {
			return err
		}

		err = g.generateDecl(decl, "json", goEnumJsonFile)
		if err != nil {
			return err
		}
	}

	for _, decl := range data.BoxedUnions {
		err := g.generateDecl(decl, "json", goUnionJsonFile)
		if err != nil {
			return err
		}
	}

	for _, decl := range data.DbJSONs {
		err := g.generateDecl(decl, "sql", goDbJSONSqlFile)
		if err != nil {
			return err
		}
	}

	for _, decl := range data.FormatJSONs {
		err := g.generateDecl(decl, "json", goFormatJSONFile)
		if err != nil {
			return err
		}
	}

	for _, decl := range data.ArrayResponses {
		err := g.generateDecl(decl, "json", goArrayResponseJsonFile)
		if err != nil {
			return err
		}
	}

	if data.GoMod != nil {
		str, err := ApplyTemplate("goModFile", goModFile, data.GoMod, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.GeneratedFile{
			Name:        "go.mod",
			Content:     str,
			SkipIfExist: true,
		})
	}

	return nil
}
