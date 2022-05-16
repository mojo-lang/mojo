package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type FormatJson struct {
    *data.Data
}

func (j *FormatJson) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
    if !decl.HasAttribute(core.FormatAttributeName) {
        return nil
    }

    pkg := context.Package(ctx)
    formatJSON := &data.FormatJSON{
        PackageName:       pkg.FullName,
        GoPackageName:     lang.GetGoPackageName(pkg.FullName),
        Name:              decl.Name,
        EnclosingName:     decl.EnclosingType.GetName(),
        EnclosingFullName: decl.EnclosingType.GetFullName(),
    }

    if decl.HasAttribute(core.EncodingAsStructAttributeName) {
        formatJSON.EncodingAsStruct = true
    }
    formatJSON.FullName = GetFullName(formatJSON.EnclosingName, formatJSON.Name)
    j.FormatJSONs = append(j.FormatJSONs, formatJSON)
    return nil
}

// CompileTypeAlias only process union type alias
func (j *FormatJson) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
    if !decl.HasAttribute(core.FormatAttributeName) { // core.FormatAttributeName) {
        return nil
    }
    if decl.Type.GetFullName() != core.UnionTypeFullName {
        return nil
    }

    pkg := context.Package(ctx)
    formatJSON := &data.FormatJSON{
        PackageName:       pkg.FullName,
        GoPackageName:     lang.GetGoPackageName(pkg.FullName),
        Name:              decl.Name,
        EnclosingName:     decl.EnclosingType.GetName(),
        EnclosingFullName: decl.EnclosingType.GetFullName(),
    }

    if decl.HasAttribute(core.EncodingAsStructAttributeName) {
        formatJSON.EncodingAsStruct = true
    }
    formatJSON.FullName = GetFullName(formatJSON.EnclosingName, formatJSON.Name)
    j.FormatJSONs = append(j.FormatJSONs, formatJSON)
    return nil
}
