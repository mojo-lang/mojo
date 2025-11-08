// Package handlers manages the exported methods in the service handler code
// adding/removing exported methods to match the service definition.
package handlers

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers/templates"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/render"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

// NewInterface is an exported func that creates a new service
// it will not be defined in the service definition but is required
const ignoredFunc = "NewService"

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

// New returns a render.Renderer capable of updating server handlers.
// New should be passed the previous version of the server handler to parse.
func New(svc *data.Interface, prev io.Reader) (render.Renderer, error) {
	var h handler
	// logs.WithField("Interface Methods", len(svc.Methods)).Debug("Handler being created")
	h.methodMap = newMethodMap(svc.Methods)
	h.service = svc

	if prev == nil {
		return &h, nil
	}

	h.fileSet = token.NewFileSet()
	var err error
	if h.ast, err = parser.ParseFile(h.fileSet, "", prev, parser.ParseComments); err != nil {
		return nil, err
	}

	return &h, nil
}

// methodMap stores all defined service methods by name and is updated to
// remove service methods already in the handler file.
type methodMap map[string]*data.Method

func newMethodMap(meths []*data.Method) methodMap {
	mMap := make(methodMap, len(meths))
	for _, m := range meths {
		mMap[m.Name] = m
	}
	return mMap
}

type handler struct {
	fileSet   *token.FileSet
	service   *data.Interface
	methodMap methodMap
	ast       *ast.File
}

// Render returns a go code server handler that has functions for all
// InterfaceMethods in the service definition.
func (h *handler) Render(alias string, service *data.Service) (io.Reader, error) {
	if alias != ServerHandlerPath {
		return nil, errors.Errorf("cannot render unknown file: %q", alias)
	}
	if h.ast == nil {
		return applyServerTmpl(service)
	}

	// Remove exported methods not defined in service definition
	// and remove methods defined in the previous file from methodMap
	logs.Debugw("Before prune", "Interface Methods", len(h.methodMap))

	var prunedDecls []ast.Decl
	h.ast.Decls, prunedDecls = h.methodMap.pruneDecls(h.ast.Decls, strcase.ToLowerCamel(service.Interface.ServerName))
	logs.Debugw("After prune", "Interface Methods", len(h.methodMap))

	if len(prunedDecls) > 0 {
		var comments []*ast.CommentGroup
		for _, comment := range h.ast.Comments {
			if !commentOfDecls(h.fileSet, comment, prunedDecls) {
				comments = append(comments, comment)
			}
		}
		h.ast.Comments = comments
	}

	// If there are no methods to templates then exit early
	if len(h.methodMap) == 0 {
		return h.buffer()
	}

	// get the code out of the ast
	code, err := h.buffer()
	if err != nil {
		return nil, err
	}

	// create a new handlerData, and add all methods not defined in the previous file
	ex := &data.Service{
		Interface: &data.Interface{
			Name:       service.Interface.Name,
			ServerName: service.Interface.ServerName,
			Methods:    nil,
		},
		FuncMap: service.FuncMap,
	}
	for k, v := range h.methodMap {
		logs.Infow("Generating handler from rpc definition", "Method", k)
		ex.Interface.Methods = append(ex.Interface.Methods, v)
	}

	// render the server for all methods not already defined
	newCode, err := applyServerMethsTmpl(ex)
	if err != nil {
		return nil, err
	}

	if _, err = code.ReadFrom(newCode); err != nil {
		return nil, err
	}

	return code, nil
}

func (h *handler) buffer() (*bytes.Buffer, error) {
	code := bytes.NewBuffer(nil)
	err := printer.Fprint(code, h.fileSet, h.ast)

	if err != nil {
		return nil, err
	}

	return code, nil
}

// pruneDecls constructs a new []ast.Decls with the exported funcs in decls
// who's names are not keys in methodMap and/or does not have the function
// receiver svcName + "Interface" ("Handler func")  removed.
//
// When a "Handler func" is not removed from decls that funcs name is also
// deleted from methodMap, resulting in a methodMap only containing keys and
// values for functions defined in the service but not in the handler ast.
//
// In addition, pruneDecls will update un-removed "Handler func"s input
// parameters and output results to by the types described in methodMap's
// serviceMethod for that "Handler func".
func (m methodMap) pruneDecls(decls []ast.Decl, svcName string) ([]ast.Decl, []ast.Decl) {
	var newDecls []ast.Decl
	var prunedDecls []ast.Decl
	for _, d := range decls {
		switch x := d.(type) {
		case *ast.FuncDecl:
			name := x.Name.Name
			// Special case NewInterface and ignore unexported
			if name == ignoredFunc || !ast.IsExported(name) || x.Recv == nil {
				logs.Debugw("Ignoring", "Func", name)
				newDecls = append(newDecls, x)
				continue
			}
			if ok := isValidFunc(x, m, svcName); ok {
				indexName := strcase.ToSnake(name)
				updateParams(x, m[indexName])
				updateResults(x, m[indexName])
				newDecls = append(newDecls, x)
				delete(m, indexName)
			} else {
				prunedDecls = append(prunedDecls, x)
			}
		default:
			newDecls = append(newDecls, d)
		}
	}
	return newDecls, prunedDecls
}

// updateParams updates the second param of f to be `X`.(m.Request.Name).
// func ProtoMethod(ctx context.Context, *pb.Old) ...-> func ProtoMethod(ctx context.Context, *pb.(m.Request.Name))...
func updateParams(f *ast.FuncDecl, m *data.Method) {
	if f.Type.Params.NumFields() != 2 {
		logs.Warnw("Function params signature should be func NAME(ctx context.Context, in *pb.TYPE), cannot fix", "Function", f.Name.Name)
		return
	}
	updatePBFieldType(f.Type.Params.List[1].Type, m.Request.Name)
}

// updateResults updates the first result of f to be `X`.(m.Response.Name).
// func ProtoMethod(...) (*pb.Old, error) ->  func ProtoMethod(...) (*pb.(m.Response.Name), error)
func updateResults(f *ast.FuncDecl, m *data.Method) {
	if f.Type.Results.NumFields() != 2 {
		logs.Warnw("Function results signature should be (*pb.TYPE, error), cannot fix", "Function", f.Name.Name)
		return
	}
	updatePBFieldType(f.Type.Results.List[0].Type, m.Response.Name)
}

// updatePBFieldType updates t if in the form X.Sel/*X.Sel to X.newType/*X.newType.
func updatePBFieldType(t ast.Expr, newType string) {
	// *pb.TYPE -> pb.TYPE
	if ptr, _ := t.(*ast.StarExpr); ptr != nil {
		t = ptr.X
	}
	// pb.TYPE -> TYPE
	if sel, _ := t.(*ast.SelectorExpr); sel != nil {
		// pb.SOMETYPE -> pb.newType
		sel.Sel.Name = newType
	}
}

// isValidFunc returns false if f is exported and does not exist in m with
// receiver svcName + "Interface".
func isValidFunc(f *ast.FuncDecl, m methodMap, svcName string) bool {
	name := f.Name.String()
	if !ast.IsExported(name) {
		logs.Debugw("Unexported function; ignoring", "func", name)
		return true
	}

	v := m[strcase.ToSnake(name)]
	if v == nil {
		logs.Infow("Method does not exist in service definition as a rpc; removing", "Method", name)
		return false
	}

	rName := receiveTypeToString(f.Recv)
	if rName != svcName {
		logs.Infow("Func is exported with improper receiver; removing", "Func", name, "Receiver", rName)
		return false
	}

	logs.Debugw("Method already exists in service definition; ignoring", "Func", name)
	return true
}

// receiveTypeToString accepts an *ast.FuncDecl.Recv recv, and returns the
// string of the recv type.
//
//	func (s Foo) Test() {} -> "Foo"
func receiveTypeToString(recv *ast.FieldList) string {
	if recv == nil ||
		recv.List[0].Type == nil {
		logs.Debug("Function has no receiver")
		return ""
	}

	return exprString(recv.List[0].Type)
}

// exprString returns the string representation of
// ast.Expr for function receivers, parameters, and results.
func exprString(e ast.Expr) string {
	var prefix string
	// *Foo -> Foo
	if ptr, _ := e.(*ast.StarExpr); ptr != nil {
		prefix = "*"
		e = ptr.X
	}
	// *foo.Foo or foo.Foo
	if sel, _ := e.(*ast.SelectorExpr); sel != nil {
		// *foo.Foo -> foo.Foo
		if ptr, _ := e.(*ast.StarExpr); ptr != nil {
			prefix = "*"
			e = ptr.X
		}
		// foo.Foo
		if x, _ := sel.X.(*ast.Ident); x != nil {
			return prefix + x.Name + "." + sel.Sel.Name
		}
		return ""
	}

	// Foo
	if base, _ := e.(*ast.Ident); base != nil {
		return prefix + base.Name
	}

	return ""
}

func commentOfDecls(file *token.FileSet, comment *ast.CommentGroup, decls []ast.Decl) bool {
	for _, decl := range decls {
		if commentOfDecl(file, comment, decl) {
			return true
		}
	}
	return false
}

func commentOfDecl(file *token.FileSet, comment *ast.CommentGroup, decl ast.Decl) bool {
	lb := file.Position(comment.Pos()).Line
	le := file.Position(comment.End()).Line
	ldb := file.Position(decl.Pos()).Line
	lde := file.Position(decl.End()).Line

	if lb >= (ldb-1) && le < lde {
		return true
	}

	return false
}

func applyServerTmpl(service *data.Service) (io.Reader, error) {
	logs.Debug("Rendering handler for the first time")
	return util.ApplyTemplate("ServerTmpl", GetHandlersTemplate(), service, service.FuncMap)
}

func applyServerMethsTmpl(service *data.Service) (io.Reader, error) {
	return util.ApplyTemplate("ServerMethsTmpl", handlerMethods, service, service.FuncMap)
}
