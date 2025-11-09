package scaffolding

import (
	"errors"
	"fmt"
	"io/ioutil"
	path2 "path"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

func GetPackageDirName(data *Data) string {
	return strcase.ToKebab(data.Package.Name)
}

type Generator struct {
}

func (g *Generator) Generate(data *Data, output string) error {
	var files []*util.GeneratedFile
	pkgDirName := GetPackageDirName(data)

	if len(data.Package.Version) == 0 {
		data.Package.Version = "0.1.0"
	}

	if data.IsMojoPackage() {
		if len(data.Package.Repository) == 0 {
			data.Package.Repository = "https://github.com/mojo-lang/" + pkgDirName
		}
		if len(data.Package.Organization) == 0 {
			data.Package.Organization = "mojo-lang.org"
		}
	}

	applyTemplate := func(name string, tmpl string, tmplData interface{}, filename string) error {
		if tmplData == nil {
			return errors.New("the input templates data is nil")
		}

		str, err := ApplyTemplate(name, tmpl, tmplData, FuncMap)
		if err != nil {
			return err
		}
		files = append(files, &util.GeneratedFile{
			Name:                path2.Join(output, pkgDirName, filename),
			Content:             str,
			SkipIfUserCodeMixed: true,
		})
		return nil
	}

	if err := applyTemplate("packageFile", packageFile, data.Package, "package.mojo"); err != nil {
		return err
	}
	if err := applyTemplate("gitignoreFile", gitignoreFile, data.Package, ".gitignore"); err != nil {
		return err
	}
	if err := applyTemplate("readmeFile", readmeFile, data.Package, "README.md"); err != nil {
		return err
	}

	if len(files) > 0 {
		if core.IsExist(path2.Dir(files[0].Name)) {
			return fmt.Errorf("the package %s is already generated", data.Package.Name)
		}
	}

	for _, f := range files {
		if len(f.Name) > 0 && len(f.Content) > 0 {
			name := f.Name
			path := path2.Dir(name)
			if err := core.CreateDir(path); err != nil {
				return err
			}
			if err := ioutil.WriteFile(name, []byte(f.Content), 0666); err != nil {
				return err
			}
		}
	}

	// mkdir mojo folders
	mojoPath := lang.PackageNameToPath(data.Package.FullName)
	mojoFullPath := path2.Join(output, pkgDirName, "mojo", mojoPath)
	if data.IsMojoPackage() {
		mojoFullPath = path2.Join(output, pkgDirName, mojoPath)
	}
	return core.CreateDir(mojoFullPath)
}
