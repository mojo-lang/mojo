package generator

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/compiler"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    path2 "path"
)

type Generator struct {
    Files util.GeneratedFiles
}

func packageToPath(pkg string) string {
    return path2.Join("pkg", lang.PackageNameToPath(pkg))
}

func (g *Generator) generateDecl(decl compiler.Decl, fileType string, template string) error {
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

func (g *Generator) Generate(data *compiler.Data) error {
    for _, decl := range data.BoxedArrays {
        g.generateDecl(decl, "json", goArrayJsonFile)
    }

    for _, decl := range data.Enums {
        g.generateDecl(decl, "fmt", goEnumFmtFile)
        g.generateDecl(decl, "json", goEnumJsonFile)
    }

    //for _, decl := range data.BoxedUnions {
    //	g.generateDecl(decl, "", "", goUnionJsonFile)
    //}

    for _, decl := range data.DbJSONs {
        g.generateDecl(decl, "sql", goDbJSONSqlFile)
    }

    for _, decl := range data.PaginationResults {
        g.generateDecl(decl, "json", goPaginationJsonFile)
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
