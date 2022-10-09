// Package handlers manages the exported methods in the service handler code
// adding/removing exported methods to match the service definition.
package handlers

import (
	"bytes"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers/templates"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/render"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
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
	//logs.WithField("Interface Methods", len(svc.Methods)).Debug("Handler being created")
	h.mMap = newMethodMap(svc.Methods)
	h.service = svc

	if prev == nil {
		return &h, nil
	}

	h.fset = token.NewFileSet()
	var err error
	if h.ast, err = parser.ParseFile(h.fset, "", prev, parser.ParseComments); err != nil {
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
	fset    *token.FileSet
	service *data.Interface
	mMap    methodMap
	ast     *ast.File
}

type handlerData struct {
	InterfaceName string
	Methods       []*data.Method
}

// Render returns a go code server handler that has functions for all
// InterfaceMethods in the service definition.
func (h *handler) Render(alias string, data *data.Service) (io.Reader, error) {
	if alias != ServerHandlerPath {
		return nil, errors.Errorf("cannot render unknown file: %q", alias)
	}
	if h.ast == nil {
		return applyServerTmpl(data)
	}

	// Remove exported methods not defined in service definition
	// and remove methods defined in the previous file from methodMap
	//log.WithField("Interface Methods", len(h.mMap)).Debug("Before prune")
	// Lowercase the service name before pruning because the templates all
	// lowercase the service name when generating code to ensure Identifiers
	// incorporating the service name remain unexported.
	h.ast.Decls = h.mMap.pruneDecls(h.ast.Decls, strcase.ToLowerCamel(data.Interface.Name))
	//log.WithField("Interface Methods", len(h.mMap)).Debug("After prune")

	// create a new handlerData, and add all methods not defined in the previous file
	ex := handlerData{
		InterfaceName: data.Interface.Name,
	}

	// If there are no methods to templates then exit early
	if len(h.mMap) == 0 {
		return h.buffer()
	}

	for _, v := range h.mMap {
		//log.WithField("Method", k).
		//	Info("Generating handler from rpc definition")
		ex.Methods = append(ex.Methods, v)
	}

	// get the code out of the ast
	code, err := h.buffer()
	if err != nil {
		return nil, err
	}

	// render the server for all methods not already defined
	//FIXME 使用ex构造一个data
	newCode, err := applyServerMethsTmpl(data)

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
	err := printer.Fprint(code, h.fset, h.ast)

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
// In addition pruneDecls will update unremoved "Handler func"s input
// paramaters and output results to by the types described in methodMap's
// serviceMethod for that "Handler func".
func (m methodMap) pruneDecls(decls []ast.Decl, svcName string) []ast.Decl {
	var newDecls []ast.Decl
	for _, d := range decls {
		switch x := d.(type) {
		case *ast.FuncDecl:
			name := x.Name.Name
			// Special case NewInterface and ignore unexported
			if name == ignoredFunc || !ast.IsExported(name) {
				//log.WithField("Func", name).
				//	Debug("Ignoring")
				newDecls = append(newDecls, x)
				continue
			}
			if ok := isValidFunc(x, m, svcName); ok == true {
				indexName := strcase.ToSnake(name)
				updateParams(x, m[indexName])
				updateResults(x, m[indexName])
				newDecls = append(newDecls, x)
				delete(m, indexName)
			}
		default:
			newDecls = append(newDecls, d)
		}

	}
	return newDecls
}

// updateParams updates the second param of f to be `X`.(m.Request.Name).
// func ProtoMethod(ctx context.Context, *pb.Old) ...-> func ProtoMethod(ctx context.Context, *pb.(m.Request.Name))...
func updateParams(f *ast.FuncDecl, m *data.Method) {
	if f.Type.Params.NumFields() != 2 {
		//log.WithField("Function", f.Name.Name).
		//	Warn("Function params signature should be func NAME(ctx context.Context, in *pb.TYPE), cannot fix")
		return
	}
	updatePBFieldType(f.Type.Params.List[1].Type, m.Request.Name)
}

// updateResults updates the first result of f to be `X`.(m.Response.Name).
// func ProtoMethod(...) (*pb.Old, error) ->  func ProtoMethod(...) (*pb.(m.Response.Name), error)
func updateResults(f *ast.FuncDecl, m *data.Method) {
	if f.Type.Results.NumFields() != 2 {
		//log.WithField("Function", f.Name.Name).
		//	Warn("Function results signature should be (*pb.TYPE, error), cannot fix")
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
		//pb.SOMETYPE -> pb.newType
		sel.Sel.Name = newType
	}
}

// isVaidFunc returns false if f is exported and does no exist in m with
// reciever svcName + "Interface".
func isValidFunc(f *ast.FuncDecl, m methodMap, svcName string) bool {
	name := f.Name.String()
	if !ast.IsExported(name) {
		//log.WithField("Func", name).
		//	Debug("Unexported function; ignoring")
		return true
	}

	v := m[strcase.ToSnake(name)]
	if v == nil {
		//log.WithField("Method", name).
		//	Info("Method does not exist in service definition as an rpc; removing")
		return false
	}

	rName := recvTypeToString(f.Recv)
	if rName != svcName {
		//log.WithField("Func", name).WithField("Receiver", rName).
		//	Info("Func is exported with improper receiver; removing")
		return false
	}

	//log.WithField("Func", name).
	//	Debug("Method already exists in service definition; ignoring")

	return true
}

// recvTypeToString accepts an *ast.FuncDecl.Recv recv, and returns the
// string of the recv type.
//	func (s Foo) Test() {} -> "Foo"
func recvTypeToString(recv *ast.FieldList) string {
	if recv == nil ||
		recv.List[0].Type == nil {
		//log.Debug("Function has no reciever")
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

func applyServerTmpl(service *data.Service) (io.Reader, error) {
	//logs.Debug("Rendering handler for the first time")
	return util.ApplyTemplate("ServerTmpl", GetHandlersTemplate(), service, service.FuncMap)
}

func applyServerMethsTmpl(service *data.Service) (io.Reader, error) {
	return util.ApplyTemplate("ServerMethsTmpl", handlerMethods, service, service.FuncMap)
}
