// Package handlers discovers RPC implementations across a Go package and appends
// missing methods without replacing user implementations.
package handlers

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"path"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/handlers/templates"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/render"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

// ServerHandlerPath is the relative path to the server handler templates file
const ServerHandlerPath = "pkg/NAME-service/handlers/handlers.go.tmpl"

var (
	handlerInterface = templates.HandlerInterface
	handlerMethods   = templates.HandlerMethods
	handlerExtension = ""
)

func ResetHandlerInterface(str string) {
	handlerInterface = str
}

func ResetHandlerMethods(methods string) {
	handlerMethods = methods
}

func ResetHandlerExtension(extension string) {
	handlerExtension = extension
}

func GetHandlersTemplate() string {
	return templates.Handlers + handlerInterface + handlerMethods + handlerExtension
}

// New preserves the single-file API for callers without a package snapshot.
func New(svc *data.Interface, prev io.Reader) (render.Renderer, error) {
	files := map[string][]byte{}
	if prev != nil {
		source, err := io.ReadAll(prev)
		if err != nil {
			return nil, err
		}
		files["handlers.go"] = source
	}
	return NewPackage(svc, "handlers.go", files)
}

// NewPackage reads a snapshot of the immediate handlers package. It never
// changes sibling files and does not require generated imports to compile.
func NewPackage(svc *data.Interface, target string, files map[string][]byte) (render.Renderer, error) {
	h := &handler{service: svc, target: target, fileSet: token.NewFileSet(), files: map[string]*ast.File{}, sources: map[string][]byte{}}
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if path.Dir(name) != path.Dir(target) || !isSourceFile(name) {
			continue
		}
		file, err := parser.ParseFile(h.fileSet, name, files[name], parser.ParseComments)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse handler package file %s", name)
		}
		h.files[name] = file
		h.sources[name] = append([]byte(nil), files[name]...)
	}
	return h, nil
}

type handler struct {
	service *data.Interface
	target  string
	fileSet *token.FileSet
	files   map[string]*ast.File
	sources map[string][]byte
}

func isSourceFile(name string) bool {
	base := path.Base(name)
	return strings.HasSuffix(base, ".go") && !strings.HasSuffix(base, "_test.go") && !strings.HasPrefix(base, "_") && !strings.HasPrefix(base, ".")
}

// Render appends missing direct RPC methods, preserving existing declarations
// and bodies. Incompatible signatures must be resolved explicitly by the user.
func (h *handler) Render(alias string, service *data.Service) (io.Reader, error) {
	if alias != ServerHandlerPath {
		return nil, errors.Errorf("cannot render unknown file: %q", alias)
	}
	skeleton, err := applyServerTmpl(service)
	if err != nil {
		return nil, err
	}
	source, err := io.ReadAll(skeleton)
	if err != nil {
		return nil, err
	}
	expectedSet := token.NewFileSet()
	expected, err := parser.ParseFile(expectedSet, "generated-handlers.go", source, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	goName, ok := service.FuncMap["GoName"].(func(string) string)
	if !ok {
		return nil, errors.New("handler template requires GoName func(string) string")
	}
	receiver := strcase.ToLowerCamel(service.Interface.ServerName)
	methods := map[string]*ast.FuncDecl{}
	for _, decl := range expected.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && receiverName(fn) == receiver {
			methods[fn.Name.Name] = fn
		}
	}
	expectedMethods := map[string]*ast.FuncDecl{}
	for _, method := range service.Interface.Methods {
		name := goName(method.Name)
		fn := methods[name]
		if fn == nil {
			return nil, errors.Errorf("handler template does not define %s.%s", receiver, name)
		}
		expectedMethods[name] = fn
	}
	// Sort diagnostics as well as newly generated declarations.
	names := make([]string, 0, len(h.files))
	for name := range h.files {
		names = append(names, name)
	}
	sort.Strings(names)
	found := map[string]string{}
	declarations := map[string]bool{}
	for _, name := range names {
		file := h.files[name]
		if file.Name.Name != expected.Name.Name {
			continue
		}
		conditional := conditionalSource(name, file)
		if name == h.target && conditional {
			return nil, errors.Errorf("%s: handler output must not have build constraints", name)
		}
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv == nil {
				if !conditional {
					for _, name := range declaredNames(decl) {
						declarations[name] = true
					}
				}
				continue
			}
			if receiverName(fn) != receiver {
				continue
			}
			want := expectedMethods[fn.Name.Name]
			if want == nil {
				continue
			}
			position := h.fileSet.Position(fn.Pos()).String()
			// A common output cannot safely fill only some build configurations.
			if conditional {
				return nil, errors.Errorf("%s: RPC %s has build constraints; keep an unconditional handler entry point and move conditional logic into helpers", position, fn.Name.Name)
			}
			if previous := found[fn.Name.Name]; previous != "" {
				return nil, errors.Errorf("%s: duplicate RPC %s (also declared at %s)", position, fn.Name.Name, previous)
			}
			actualSignature := signature(fn.Type, file)
			wantedSignature := signature(want.Type, expected)
			if actualSignature != wantedSignature {
				return nil, errors.Errorf("%s: RPC %s signature mismatch: have %s; want %s; existing implementation was not modified", position, fn.Name.Name, actualSignature, wantedSignature)
			}
			found[fn.Name.Name] = position
		}
	}
	ex := *service
	iface := *service.Interface
	iface.Methods = nil
	ex.Interface = &iface
	for _, method := range service.Interface.Methods {
		name := goName(method.Name)
		if found[name] == "" {
			iface.Methods = append(iface.Methods, method)
		}
	}
	if current, exists := h.sources[h.target]; exists {
		if h.files[h.target].Name.Name != expected.Name.Name {
			return nil, errors.Errorf("%s: expected package %s", h.target, expected.Name.Name)
		}
		if len(iface.Methods) == 0 {
			return bytes.NewReader(current), nil
		}
		additions, err := applyServerMethsTmpl(&ex)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(additions)
		if err != nil {
			return nil, err
		}
		// Reuse existing import aliases and add only the imports referenced by new
		// declarations; gofmt alone cannot repair alias collisions.
		return appendMethods(current, body, h.files[h.target], expected, h.fileSet)
	}
	// A user may also move the server type or constructor to another file.
	// Generate the remaining skeleton without duplicating those declarations.
	generated, err := applyServerTmpl(&ex)
	if err != nil {
		return nil, err
	}
	contents, err := io.ReadAll(generated)
	if err != nil {
		return nil, err
	}
	set := token.NewFileSet()
	file, err := parser.ParseFile(set, h.target, contents, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	var kept []ast.Decl
	for _, decl := range file.Decls {
		duplicate := false
		for _, name := range declaredNames(decl) {
			if declarations[name] {
				duplicate = true
			}
		}
		if !duplicate {
			kept = append(kept, decl)
		}
	}
	file.Decls = kept
	removeUnusedImports(file)
	var out bytes.Buffer
	if err := printer.Fprint(&out, set, file); err != nil {
		return nil, err
	}
	return &out, nil
}

func receiverName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) != 1 {
		return ""
	}
	typ := fn.Recv.List[0].Type
	if pointer, ok := typ.(*ast.StarExpr); ok {
		typ = pointer.X
	}
	if name, ok := typ.(*ast.Ident); ok {
		return name.Name
	}
	return ""
}

func declaredNames(decl ast.Decl) []string {
	if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv == nil {
		return []string{fn.Name.Name}
	}
	var names []string
	if gen, ok := decl.(*ast.GenDecl); ok {
		for _, spec := range gen.Specs {
			switch spec := spec.(type) {
			case *ast.TypeSpec:
				names = append(names, spec.Name.Name)
			case *ast.ValueSpec:
				for _, name := range spec.Names {
					if name.Name != "_" {
						names = append(names, name.Name)
					}
				}
			}
		}
	}
	return names
}

func applyServerTmpl(service *data.Service) (io.Reader, error) {
	logs.Debug("Rendering handler for the first time")
	return util.ApplyTemplate("ServerTmpl", GetHandlersTemplate(), service, service.FuncMap)
}

func applyServerMethsTmpl(service *data.Service) (io.Reader, error) {
	return util.ApplyTemplate("ServerMethsTmpl", handlerMethods, service, service.FuncMap)
}
