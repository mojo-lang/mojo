package injection

import (
	"context"
	"go/ast"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
)

const (
	fileKey        = "@file"
	fileContentKey = "@fileContent"

	structTypeKey = "@structType"
	structNameKey = "@structName"
)

func GetFile(ctx context.Context) *ast.File {
	if file, ok := ctx.Value(fileKey).(*ast.File); ok {
		return file
	}
	return nil
}

func GetFileContent(ctx context.Context) []byte {
	if content, ok := ctx.Value(fileContentKey).([]byte); ok {
		return content
	}
	return nil
}

func GetStructType(ctx context.Context) *ast.StructType {
	if structType, ok := ctx.Value(structTypeKey).(*ast.StructType); ok {
		return structType
	}
	return nil
}

func GetStructName(ctx context.Context) string {
	if structName, ok := ctx.Value(structNameKey).(string); ok {
		return structName
	}
	return ""
}

func ToMojoStructName(name string) string {
	return strings.ReplaceAll(name, "_", ".")
}

func ToMojoFieldName(name string) string {
	return strcase.ToSnake(name)
}
