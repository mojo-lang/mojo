package plugin

import (
    "errors"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin/hook"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin/parser"
    "io/fs"
)

const (
    parsePackageMethod    = "ParsePackage"
    parseSourceFileMethod = "ParseSourceFile"
    parseStructMethod     = "ParseStruct"
    parseInterfaceMethod  = "ParseInterface"
    parseEnumMethod       = "ParseEnum"
    parseTypeAliasMethod  = "ParseTypeAlias"
)

func ParseFile(p interface{}, fileName string, fileSys fs.FS) (*lang.SourceFile, error) {
    if fileParser, ok := p.(parser.FileParser); ok {
        fileParser.ParseFile(context.Empty(), fileName, fileSys)
    }
    return nil, errors.New("not a valid plugin type")
}

func ParseString(p interface{}, content string) (*lang.SourceFile, error) {
    return ParseFile(p, "", core.StringFS(content))
}

func ParsePackage(p interface{}, ctx context.Context, pkg *lang.Package) error {
    if p == nil {
        return nil
    }

    if !ContainsAnyMethod(p, parsePackageMethod, parseSourceFileMethod, parseStructMethod,
        parseInterfaceMethod, parseEnumMethod, parseTypeAliasMethod) {
        return core.NewSkipError()
    }

    if hk, ok := p.(hook.PackagePreHook); ok {
        err := hk.PrePackage(ctx, pkg)
        if core.IsSkipError(err) {
            return nil
        } else {
            return err
        }
    }

    defer func() {
        if hook, ok := p.(hook.PackagePostHook); ok {
            hook.PostPackage(ctx, pkg)
        }
    }()

    if pkgParser, ok := p.(parser.PackageParser); ok {
        return pkgParser.ParsePackage(ctx, pkg)
    }

    thisCtx := context.WithScopeType(ctx, pkg)
    for _, file := range pkg.SourceFiles {
        if err := ParseSourceFile(p, thisCtx, file); err != nil {
            return err
        }
    }

    for _, child := range pkg.Children {
        if err := ParsePackage(p, thisCtx, child); err != nil {
            return err
        }
    }
    return nil
}

func ParseSourceFile(p interface{}, ctx context.Context, file *lang.SourceFile) error {
    if p == nil {
        return nil
    }

    if !ContainsAnyMethod(p, parseSourceFileMethod, parseStructMethod,
        parseInterfaceMethod, parseEnumMethod, parseTypeAliasMethod) {
        return core.NewSkipError()
    }

    if hk, ok := p.(hook.SourceFilePreHook); ok {
        err := hk.PreSourceFile(ctx, file)
        if core.IsSkipError(err) {
            return nil
        } else {
            return err
        }
    }
    defer func() {
        if hk, ok := p.(hook.SourceFilePostHook); ok {
            hk.PostSourceFile(ctx, file)
        }
    }()

    if fileParser, ok := p.(parser.SourceFileParser); ok {
        return fileParser.ParseSourceFile(ctx, file)
    }

    thisCtx := context.WithScopeType(ctx, file)
    for _, statement := range file.Statements {
        switch statement.Statement.(type) {
        case *lang.Statement_Declaration:
            switch decl := statement.GetDeclaration().Declaration.(type) {
            case *lang.Declaration_StructDecl:
                if err := ParseStruct(p, thisCtx, decl.StructDecl); err != nil {
                    return err
                }
            }
        }
    }
    return nil
}

func ParseStruct(p interface{}, ctx context.Context, decl *lang.StructDecl) error {
    return nil
}
