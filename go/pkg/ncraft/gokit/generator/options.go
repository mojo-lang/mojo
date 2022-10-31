package generator

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	_go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/httptransport"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/model"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/templates"
	util2 "github.com/mojo-lang/mojo/go/pkg/util"
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

func (o Options) SyncTo(ds *data.Service) {
	ds.CombinedAPI = o.MixedInAPI
	ds.Go.RepositoryPath = o.Repository
	ds.Go.ApiRepositoryPath = o.ApiRepository
}

func (o *Options) GenerateClient(ds *data.Service) ([]*util2.GeneratedFile, error) {
	o.SyncTo(ds)
	return o.generateTemplatedFiles(ds, templates.ClientNames(), templates.Client)
}

func (o *Options) GenerateService(ds *data.Service) ([]*util2.GeneratedFile, error) {
	o.SyncTo(ds)
	return o.generateTemplatedFiles(ds, templates.ServiceNames(), templates.Service)
}

// GenerateTemplatedFiles generate the service or client files
func (o *Options) generateTemplatedFiles(ds *data.Service, tmplPaths []string, getter templates.FileGetter) ([]*util2.GeneratedFile, error) {
	if tmpl := o.HandlerTemplate; tmpl != nil {
		if len(tmpl.Interface) > 0 {
			handlers.ResetHandlerInterface(tmpl.Interface)
		}
		if len(tmpl.Methods) > 0 {
			handlers.ResetHandlerMethods(tmpl.Methods)
		}
		if len(tmpl.Extension) > 0 {
			handlers.ResetHandlerExtension(tmpl.Extension)
		}
	}

	var codeGenFiles util2.GeneratedFiles

	// Remove the suffix "-service" since it's added back in by templatePathToActual
	svcName := strcase.ToKebab(ds.Interface.BaredName)

	for _, tmplPath := range tmplPaths {
		if tmplPath == model.TemplatePath {
			m := model.Model{}
			if files, err := m.Generate(tmplPath, ds); err != nil {
				logs.Errorw("failed to generate model templates files", "err", err.Error())
			} else {
				if len(files) == 0 {
					continue
				}
				codeGenFiles = append(codeGenFiles, files...)
			}
		}

		// Re-derive the actual path for this file based on the service output
		// path provided by the ncraft main.go
		actualPath := templatePathToActual(tmplPath, ds.Go.PackageName, svcName)
		file, err := generateTemplateFile(tmplPath, actualPath, ds, o.PreviousFiles[actualPath], getter)
		if err != nil {
			return nil, logs.NewErrorw("cannot render templates", "error", err.Error())
		}

		codeGenFiles = append(codeGenFiles, &util2.GeneratedFile{
			Name:              actualPath,
			Reader:            file,
			SkipIfExist:       false,
			SkipNoneGenerated: false,
		})
	}

	return codeGenFiles, nil
}

// generateTemplateFile
func generateTemplateFile(tmplPath string, actualPath string, ds *data.Service, prevFile io.Reader, getter templates.FileGetter) (io.Reader, error) {
	var genCode io.Reader
	var err error

	switch tmplPath {
	case handlers.ServerHandlerPath:
		h, err := handlers.New(ds.Interface, prevFile)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse previous handler: %q", actualPath)
		}

		if genCode, err = h.Render(tmplPath, ds); err != nil {
			return nil, errors.Wrapf(err, "cannot render templates: %s", tmplPath)
		}
	case handlers.HookPath:
		hook := handlers.NewHook(prevFile)
		if genCode, err = hook.Render(tmplPath, ds); err != nil {
			return nil, errors.Wrapf(err, "cannot render templates: %s", tmplPath)
		}
	case handlers.MiddlewaresPath:
		m := handlers.NewMiddlewares()
		m.Load(prevFile)
		if genCode, err = m.Render(tmplPath, ds); err != nil {
			return nil, errors.Wrapf(err, "cannot render templates: %s", tmplPath)
		}
	case httptransport.ServerHttpTransportPath:
		transport, err := httptransport.NewServerHttpTransport(ds)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create the server http transport")
		}
		if genCode, err = transport.Render(tmplPath, ds); err != nil {
			return nil, errors.Wrapf(err, "cannot render templates: %s", tmplPath)
		}
	default:
		if genCode, err = applyTemplateFromPath(tmplPath, ds, getter); err != nil {
			return nil, errors.Wrapf(err, "cannot render templates: %s", tmplPath)
		}
	}

	codeBytes, err := ioutil.ReadAll(genCode)
	if err != nil {
		return nil, err
	}

	// ignore error as we want to write the code either way to inspect after
	// writing to disk
	formattedCode := codeBytes
	if strings.HasSuffix(actualPath, ".go") {
		formattedCode = _go.FormatCodeBytes(codeBytes)
	}

	return bytes.NewReader(formattedCode), nil
}

// templatePathToActual accepts a templateFilePath and the svcName of the
// service and returns what the relative file path of what should be written to
// disk
func templatePathToActual(tmplPath, pkgName string, svcName string) string {
	// Switch "NAME" in path with svcName.
	pkgName = strcase.ToKebab(pkgName)
	svcName = strcase.ToKebab(svcName)

	actual := strings.Replace(tmplPath, "PKGNAME", pkgName, -1)
	actual = strings.Replace(actual, "NAME", svcName, -1)
	return strings.TrimSuffix(actual, ".tmpl")
}

// applyTemplateFromPath calls applyTemplate with the templates at tmplPath
func applyTemplateFromPath(tmplPath string, ds *data.Service, getter templates.FileGetter) (io.Reader, error) {
	tmplBytes, err := getter(tmplPath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find templates file: %v", tmplPath)
	}

	return util2.ApplyTemplate(tmplPath, string(tmplBytes), ds, ds.FuncMap)
}
