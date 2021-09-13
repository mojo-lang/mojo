package generator

import (
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/go/compiler"
	"github.com/mojo-lang/mojo/go/pkg/util"
	path2 "path"
	"strings"
)

type Generator struct {
	Files util.CodeGeneratedFiles
}

func packageToPath(pkg string) string {
	return path2.Join("pkg", strings.ReplaceAll(pkg, ".", "/"))
}

func (g *Generator) Generate(data *compiler.Data) error {
	for _, decl := range data.BoxedArrays {
		//str, err := ApplyTemplate("goArrayExtFile", goArrayExtFile, decl, FuncMap)
		//if err != nil {
		//	return err
		//}
		//g.Files = append(g.Files, &util.CodeGeneratedFile{
		//	Name:    path2.Join(packageToPath(decl.PackageName),  strcase.ToSnake(decl.FullName) + ".ext.go"),
		//	Content: str,
		//})

		str, err := ApplyTemplate("goArrayJsonFile", goArrayJsonFile, decl, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.CodeGeneratedFile{
			Name:    path2.Join(packageToPath(decl.PackageName), strcase.ToSnake(decl.FullName)+".json.go"),
			Content: FormatCode(str),
		})
	}

	if data.GoMod != nil {
		str, err := ApplyTemplate("goModFile", goModFile, data.GoMod, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.CodeGeneratedFile{
			Name:        "go.mod",
			Content:     str,
			SkipIfExist: true,
		})
	}

	return nil
}
