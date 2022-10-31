package injection

import (
	"context"
	"go/ast"

	"github.com/fatih/structtag"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
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

		if starExpr, ok := field.Type.(*ast.StarExpr); ok {
			if _, ok := starExpr.X.(*ast.SelectorExpr); ok {
				typeName = string(fileContent[starExpr.X.Pos()-1 : starExpr.X.End()-1])
				newTypeDefine += " struct {" + typeName + "}"

				newTypeName = "*" + newTypeName
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

		recvStructName := "*" + structName
		if i.getters[recvStructName] == nil {
			i.getters[recvStructName] = make(map[string]string)
		}
		i.getters[recvStructName]["Get"+fieldName] = newTypeName

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

		if newTypeName := i.getters[recvTypeName][function.Name.Name]; len(newTypeName) > 0 {
			areaAppender(&TextArea{
				Start:   function.Type.Results.Pos(),
				End:     function.Type.Results.End(),
				Content: newTypeName,
			})
		}
	}
}
