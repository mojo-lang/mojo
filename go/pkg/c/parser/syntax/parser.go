package syntax

import (
	"context"
	"io/fs"
	"path"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Parser struct {
	Options core.Options
}

func New(options core.Options) *Parser {
	return &Parser{Options: options}
}

func (p Parser) ParseString(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream("", input)
}

func (p Parser) ParseStream(fileName string, input antlr.CharStream) (*lang.SourceFile, error) {
	lexer := NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	errorListener := util.NewErrorListener(fileName, false)

	parser := NewCParser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.File()
	if sourceFile, ok := tree.Accept(NewFileVisitor()).(*lang.SourceFile); ok {
		if len(errorListener.Errors) == 0 {
			return sourceFile, nil
		} else {
			return nil, errorListener.Errors
		}
	}
	if len(errorListener.Errors) > 0 {
		return nil, logs.NewErrorw("failed to parse c file", "file", fileName, "error", errorListener.Errors.Error())
	}

	return nil, logs.NewErrorw("failed to parse mojo file", "file", fileName)
}

func (p Parser) ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (*lang.SourceFile, error) {
	_ = ctx
	if bytes, err := fs.ReadFile(fileSys, fileName); err != nil {
		return nil, err
	} else {
		if sourceFile, err := p.ParseStream(fileName, antlr.NewInputStream(string(bytes))); err != nil {
			return nil, err
		} else {
			sourceFile.Name = path.Base(fileName)
			sourceFile.FullName = fileName

			return sourceFile, nil
		}
	}
}

// func (p Parser) ParsePackage(ctx context.Context, pkg *lang.Package) (err error) {
//    if util.IsPackageProcessed(pkg, pluginName) {
//        return nil
//    }
//
//    pkgPath := pkg.GetExtraString("path")
//    if !strings.HasPrefix(pkg.FullName, "mojo.") {
//        pkgPath = path.Join(pkgPath, "mojo")
//    }
//
//    var files fs.FS
//    if fsCache := parser.ContextFsCache(ctx); fsCache != nil {
//        files = fsCache[pkg.FullName]
//    }
//    if files == nil {
//        workingDir := pkg.GetExtraString("workingDir")
//        if len(workingDir) == 0 {
//            workingDir = "."
//        }
//        files = os.DirFS(workingDir)
//    }
//
//    _, err = p.ParsePackagePath(parser.WithDeclaredPackage(ctx, pkg), pkgPath, files)
//    util.SetPackageProcessed(pkg, pluginName)
//    return
// }

// func (p Parser) ParsePackagePath(ctx context.Context, pkgPath string, fileSys fs.FS) (*lang.Package, error) {
//    currentPkg := parser.ContextDeclaredPackage(ctx)
//    currentPkgName := ""
//    if currentPkg != nil {
//        if util.IsPackageProcessed(currentPkg, pluginName) {
//            return currentPkg, nil
//        }
//        currentPkgName = currentPkg.FullName
//    } else {
//        currentPkgName = parser.ContextPackageName(ctx)
//        currentPkg = &lang.Package{
//            Name:     lang.GetPackageName(currentPkgName),
//            FullName: currentPkgName,
//        }
//    }
//
//    currentPath := path.Join(pkgPath, lang.PackageNameToPath(currentPkgName))
//    files, err := fs.ReadDir(fileSys, currentPath)
//    if err != nil {
//        logs.Errorw("failed to read source directory", "path", currentPath, "error", err.Error())
//        return nil, err
//    }
//
//    parentPkg := context.Package(ctx)
//    if parentPkg != nil {
//        parentPkg.Children = append(parentPkg.Children, currentPkg)
//
//        currentPkg.Authors = parentPkg.Authors
//        currentPkg.Repository = parentPkg.Repository
//        currentPkg.Version = parentPkg.Version
//    }
//
//    thisCtx := context.WithType(ctx, currentPkg)
//    for _, f := range files {
//        if f.IsDir() {
//            pkgName := ""
//            if len(currentPkgName) == 0 {
//                pkgName = f.Name()
//            } else {
//                pkgName = currentPkgName + "." + f.Name()
//            }
//
//            newCtx := parser.WithDeclaredPackage(context.WithType(thisCtx, currentPkg), &lang.Package{
//                Name:     lang.GetPackageName(pkgName),
//                FullName: pkgName,
//            })
//            if _, err = p.ParsePackagePath(newCtx, pkgPath, fileSys); err != nil {
//                return nil, err
//            }
//        } else {
//            if !strings.HasSuffix(f.Name(), ".mojo") {
//                continue
//            }
//
//            if sourceFile, err := p.ParseFile(thisCtx, path.Join(currentPath, f.Name()), fileSys); err != nil {
//                return nil, err
//            } else {
//                sourceFile.PackageName = currentPkgName
//                sourceFile.FullName = path.Join(lang.PackageNameToPath(currentPkgName), sourceFile.Name)
//                currentPkg.SourceFiles = append(currentPkg.SourceFiles, sourceFile)
//            }
//        }
//    }
//
//    util.SetPackageProcessed(currentPkg, pluginName)
//    return currentPkg, nil
// }
