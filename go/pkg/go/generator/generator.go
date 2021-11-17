package generator

import (
	path2 "path"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/go/compiler"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Generator struct {
	Files util.CodeGeneratedFiles
}

func packageToPath(pkg string) string {
	return path2.Join("pkg", strings.ReplaceAll(pkg, ".", "/"))
}

func (g *Generator) generateDecl(decl compiler.Decl, extTemplate string, fmtTemplate string, jsonTemplate string) error {
	if len(extTemplate) > 0 {
		str, err := ApplyTemplate("goExtFile", extTemplate, decl, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.CodeGeneratedFile{
			Name:    path2.Join(packageToPath(decl.GetPackageName()), strcase.ToSnake(decl.GetFullName())+".ext.go"),
			Content: FormatCode(str),
		})
	}

	if len(fmtTemplate) > 0 {
		str, err := ApplyTemplate("goFmtFile", fmtTemplate, decl, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.CodeGeneratedFile{
			Name:    path2.Join(packageToPath(decl.GetPackageName()), strcase.ToSnake(decl.GetFullName())+".fmt.go"),
			Content: FormatCode(str),
		})
	}

	if len(jsonTemplate) > 0 {
		str, err := ApplyTemplate("goJsonFile", jsonTemplate, decl, FuncMap)
		if err != nil {
			return err
		}
		g.Files = append(g.Files, &util.CodeGeneratedFile{
			Name:    path2.Join(packageToPath(decl.GetPackageName()), strcase.ToSnake(decl.GetFullName())+".json.go"),
			Content: FormatCode(str),
		})
	}
	return nil
}

func (g *Generator) Generate(data *compiler.Data) error {
	for _, decl := range data.BoxedArrays {
		g.generateDecl(decl, "", "", goArrayJsonFile)
	}

	for _, decl := range data.Enums {
		g.generateDecl(decl, "", goEnumFmtFile, goEnumJsonFile)
	}

	//for _, decl := range data.BoxedUnions {
	//	g.generateDecl(decl, "", "", goUnionJsonFile)
	//}

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
