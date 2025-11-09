package plugin

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

const (
	parsePackageMethod    = "ParsePackage"
	parseSourceFileMethod = "ParseSourceFile"
	parseStructMethod     = "ParseStruct"
	parseInterfaceMethod  = "ParseInterface"
	parseEnumMethod       = "ParseEnum"
	parseTypeAliasMethod  = "ParseTypeAlias"
)

func ParseString(p interface{}, ctx context.Context, content string) (*lang.SourceFile, error) {
	if stringParser, ok := p.(StringParser); ok {
		return stringParser.ParseString(ctx, content)
	}
	return nil, errors.New(fmt.Sprintf("%v is not a StringParser", p))
}

func ParseFile(p interface{}, ctx context.Context, fileName string) (*lang.SourceFile, error) {
	if _, ok := p.(StringParser); ok {
		ctx = WithFilename(ctx, fileName)

		var err error
		var content []byte
		f := ContextFs(ctx)
		if f == nil {
			content, err = ioutil.ReadFile(fileName)
		} else {
			content, err = fs.ReadFile(f, fileName)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read the file %s, %w", fileName, err)
		}

		sf, err := ParseString(p, ctx, string(content))
		if err != nil {
			return nil, err
		}
		sf.Name = path.Base(fileName)
		sf.FullName = fileName
		return sf, nil
	}

	if fileParser, ok := p.(FileParser); ok {
		return fileParser.ParseFile(ctx, fileName)
	}
	return nil, errors.New("not a valid plugin type")
}

func ParsePackage(p interface{}, ctx context.Context, pkg *lang.Package) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, parsePackageMethod, parseSourceFileMethod, parseStructMethod,
		parseInterfaceMethod, parseEnumMethod, parseTypeAliasMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(PackagePreHook); ok {
		err := hk.PrePackage(ctx, pkg)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}

	defer func() {
		if hook, ok := p.(PackagePostHook); ok {
			hook.PostPackage(ctx, pkg)
		}
	}()

	if pkgParser, ok := p.(PackageParser); ok {
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

	if hk, ok := p.(SourceFilePreHook); ok {
		err := hk.PreSourceFile(ctx, file)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}
	defer func() {
		if hk, ok := p.(SourceFilePostHook); ok {
			hk.PostSourceFile(ctx, file)
		}
	}()

	if fileParser, ok := p.(SourceFileParser); ok {
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
