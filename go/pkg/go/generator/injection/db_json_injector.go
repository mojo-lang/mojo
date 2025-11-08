package injection

import (
	"context"
	"fmt"
	"github.com/fatih/structtag"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"go/ast"
)

type DbJSONInjector struct {
	Injector

	newTypeDecl string
	getters     map[string]map[string]string
}

func NewDbJSONInjector() *DbJSONInjector {
	i := &DbJSONInjector{
		getters: make(map[string]map[string]string),
	}

	i.OnStructField = i.onStructField
	i.PreStruct = i.preStruct
	i.PostStruct = i.postStruct
	i.OnFunction = i.onFunction
	return i
}

func (i *DbJSONInjector) preStruct(ctx context.Context, structType *ast.StructType, structName string) {
	i.newTypeDecl = ""
}

func (i *DbJSONInjector) onStructField(ctx context.Context, field *ast.Field, appender func(area Area)) {
	if field.Tag == nil {
		return
	}

	fieldName := field.Names[0].Name
	tags, _ := structtag.Parse(core.RemoveQuote(field.Tag.Value, "`"))
	if tags == nil {
		return
	}

	dbTag, _ := tags.Get("db")
	if dbTag != nil && HasTagOption(dbTag, "json") {
		structName := GetStructName(ctx)
		fileContent := GetFileContent(ctx)

		newTypeName := structName + fieldName
		newTypeDefine := "type " + newTypeName

		typeName := string(fileContent[field.Type.Pos()-1 : field.Type.End()-1])
		baseTypeName := typeName
		if starExpr, ok := field.Type.(*ast.StarExpr); ok {
			if selExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
				typeName = string(fileContent[starExpr.X.Pos()-1 : starExpr.X.End()-1])
				baseTypeName = selExpr.Sel.Name

				newTypeDefine += " struct { *" + typeName + "}"
				newTypeDefine += fmt.Sprintf("\nfunc (x *%s) Init() { x.%s = &%s{} }", newTypeName, baseTypeName, typeName)
			} else {
				return
			}
		} else if _, ok := field.Type.(*ast.ArrayType); ok {
			newTypeDefine += " " + typeName
		} else if _, ok := field.Type.(*ast.MapType); ok {
			newTypeDefine += " " + typeName
		} else {
			return
		}

		if typeName != baseTypeName {
			recvStructName := "*" + structName
			if i.getters[recvStructName] == nil {
				i.getters[recvStructName] = make(map[string]string)
			}
			i.getters[recvStructName]["Get"+fieldName] = fieldName + "." + baseTypeName
		}

		appender(&TextArea{
			Start:   field.Type.Pos(),
			End:     field.Type.End(),
			Content: newTypeName,
		})

		i.newTypeDecl += "\n" + newTypeDefine
	}
}

func (i *DbJSONInjector) postStruct(ctx context.Context, areaAppender func(area Area)) {
	if len(i.newTypeDecl) > 0 {
		structType := GetStructType(ctx)

		areaAppender(&TextArea{
			Start:   structType.End(),
			End:     structType.End(),
			Content: i.newTypeDecl,
		})
	}
}

func (i *DbJSONInjector) onFunction(ctx context.Context, function *ast.FuncDecl, areaAppender func(area Area)) {
	if function.Recv != nil && len(function.Recv.List) > 0 {
		fileContent := GetFileContent(ctx)
		recvType := function.Recv.List[0].Type
		recvTypeName := string(fileContent[recvType.Pos()-1 : recvType.End()-1])

		if len(i.getters[recvTypeName]) == 0 {
			return
		}

		if newIdent := i.getters[recvTypeName][function.Name.Name]; len(newIdent) > 0 {
			var target *ast.Ident
			if ifStmt, ok := function.Body.List[0].(*ast.IfStmt); ok {
				if returnStmt, ok := ifStmt.Body.List[0].(*ast.ReturnStmt); ok {
					if selectorExpr, ok := returnStmt.Results[0].(*ast.SelectorExpr); ok {
						target = selectorExpr.Sel
					}
				}
			}
			areaAppender(&TextArea{
				Start:   target.Pos(),
				End:     target.End(),
				Content: newIdent,
			})
		}
	}
}
