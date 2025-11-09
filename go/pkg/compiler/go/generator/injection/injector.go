package injection

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
)

type Injector struct {
	Options map[string]interface{}

	PreStruct     func(ctx context.Context, structType *ast.StructType, structName string)
	PostStruct    func(ctx context.Context, areaAppender func(Area))
	OnStructField func(ctx context.Context, field *ast.Field, areaAppender func(Area))
	OnFunction    func(ctx context.Context, function *ast.FuncDecl, areaAppender func(Area))
	OnImport      func(ctx context.Context, importSpec *ast.ImportSpec)
}

func (i Injector) Inject(fileName string, content []byte) ([]byte, error) {
	areas, err := i.parseFile(fileName, content)
	if err != nil {
		return nil, err
	}
	return injectFile(content, areas), nil
}

func (i Injector) parseFile(fileName string, content []byte) (areas []Area, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, content, parser.ParseComments)
	if err != nil {
		return
	}

	areasAppender := func(area Area) {
		areas = append(areas, area)
	}

	fileCtx := context.WithValue(context.WithValue(context.Background(), fileKey, f), fileContentKey, content)
	for _, decl := range f.Decls {

		if funcDecl, ok := decl.(*ast.FuncDecl); ok && i.OnFunction != nil {
			i.OnFunction(fileCtx, funcDecl, areasAppender)
		}

		// check if is generic declaration
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		var importSpec *ast.ImportSpec
		var typeSpec *ast.TypeSpec
		for _, spec := range genDecl.Specs {
			if ts, tsOK := spec.(*ast.TypeSpec); tsOK {
				typeSpec = ts
			} else if is, isOk := spec.(*ast.ImportSpec); isOk {
				importSpec = is
			}
		}

		if importSpec != nil && i.OnImport != nil {
			i.OnImport(fileCtx, importSpec)
		}

		// skip if can't get type spec
		if typeSpec == nil {
			continue
		}

		// not a struct, skip
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		structName := typeSpec.Name.Name

		structCtx := context.WithValue(context.WithValue(fileCtx, structTypeKey, structType), structNameKey, structName)
		if i.PreStruct != nil {
			i.PreStruct(structCtx, structType, structName)
		}

		if i.OnStructField != nil {
			for _, field := range structType.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				i.OnStructField(structCtx, field, areasAppender)
			}
		}

		if i.PostStruct != nil {
			i.PostStruct(structCtx, areasAppender)
		}
	}

	// logs.Infof("parsed file %s, number of fields to injection custom tags: %d", fileName, len(areas))
	return
}

func injectFile(input []byte, areas []Area) []byte {
	Sort(areas)

	// injection custom tags from tail of file first to preserve order
	for j := range areas {
		area := areas[len(areas)-j-1]
		// logs.Debugf("injection custom tag %q to expression %q", area.InjectTag, string(input[area.Start-1:area.End-1]))
		input = injectArea(input, area)
	}

	return input
}

func injectArea(contents []byte, area Area) (injected []byte) {
	injected = append(injected, contents[:area.GetStart()-1]...)
	injected = append(injected, area.GetContent()...)
	injected = append(injected, contents[area.GetEnd()-1:]...)
	return
}
