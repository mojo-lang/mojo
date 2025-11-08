package generator

import (
	"bytes"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/boot/generator/template"
	"io"
	"path"
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/templates"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type HandlerTemplate struct {
	Interface string
	Methods   string
	Extension string
}

type Options struct {
	Repository    string // the repository for the generated gokit service
	MixedInAPI    bool
	ApiRepository string
	Version       string
	VersionDate   string
	Output        string
	PreviousFiles map[string]io.Reader

	ExtensionData map[string]interface{}

	HandlerTemplate *HandlerTemplate
}

func (o *Options) SyncTo(ds *data.Service) {
	ds.CombinedAPI = o.MixedInAPI
	ds.Go.RepositoryPath = o.Repository
	ds.Go.ApiRepositoryPath = o.ApiRepository
}

func (o *Options) GenerateService(ds *data.Service) ([]*util.GeneratedFile, error) {
	o.SyncTo(ds)

	var codeGenFiles util.GeneratedFiles
	// Remove the suffix "-service" since it's added back in by templatePathToActual
	svcName := strcase.ToKebab(ds.Interface.BaredName)

	// Re-derive the actual path for this file based on the service output
	// path provided by the ncraft main.go
	for _, tmplPath := range template.ServiceNames() {
		actualPath := templatePathToActual(tmplPath, ds.Java.PackageName, svcName)
		if strings.Contains(actualPath, "ENTITY") {
			continue
		}

		file, err := generateTemplateFile(tmplPath, actualPath, ds, o.PreviousFiles[actualPath], template.Service)
		if err != nil {
			return nil, logs.NewErrorw("cannot render template", "error", err.Error())
		}

		codeGenFiles = append(codeGenFiles, &util.GeneratedFile{
			Name:        actualPath,
			Reader:      file,
			SkipIfExist: false,
		})
	}
	return codeGenFiles, nil
}

// generateTemplateFile
func generateTemplateFile(tmplPath string, actualPath string, ds *data.Service, prevFile io.Reader, getter templates.FileGetter) (io.Reader, error) {
	var genCode io.Reader
	var err error

	if genCode, err = applyTemplateFromPath(tmplPath, ds, getter); err != nil {
		return nil, errors.Wrapf(err, "cannot render template: %s", tmplPath)
	}

	codeBytes, err := io.ReadAll(genCode)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(codeBytes), nil
}

// templatePathToActual accepts a templateFilePath and the svcName of the
// service and returns what the relative file path of what should be written to
// disk
func templatePathToActual(tmplPath, pkgName string, svcName string) string {
	svcName = strcase.ToKebab(svcName)

	actual := strings.Replace(tmplPath, "NAME", svcName, -1)
	actual = strings.Replace(actual, "PACKAGE", strings.Replace(pkgName, ".", "/", -1), -1)
	actual = strings.TrimSuffix(actual, ".tmpl")

	if strings.HasSuffix(actual, ".java") {
		name := path.Base(actual)
		name = strings.TrimSuffix(name, ".java")
		actual = strings.Replace(actual, name+".java", strcase.ToCamel(name)+".java", -1)
	}
	return actual
}

// applyTemplateFromPath calls applyTemplate with the template at tmplPath
func applyTemplateFromPath(tmplPath string, ds *data.Service, getter templates.FileGetter) (io.Reader, error) {
	tmplBytes, err := getter(tmplPath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find template file: %v", tmplPath)
	}

	return util.ApplyTemplate(tmplPath, string(tmplBytes), ds, ds.FuncMap)
}
