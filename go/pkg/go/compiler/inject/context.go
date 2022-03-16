package inject

import (
    "context"
    "go/ast"
)

func GetFile(ctx context.Context) *ast.File {
    if file, ok := ctx.Value("@file").(*ast.File); ok {
        return file
    }
    return nil
}

func GetFileContent(ctx context.Context) []byte {
    if content, ok := ctx.Value("@fileContent").([]byte); ok {
        return content
    }
    return nil
}

func GetStructType(ctx context.Context) *ast.StructType {
    if structType, ok := ctx.Value("@structType").(*ast.StructType); ok {
        return structType
    }
    return nil
}

func GetStructName(ctx context.Context) string {
    if structName, ok := ctx.Value("@structName").(string); ok {
        return structName
    }
    return ""
}
