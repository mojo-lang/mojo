package parser

import (
	"os"
	"runtime"
	"strings"

	"github.com/go-clang/clang-v13/clang"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

const macOSDefaultIncludeDir = "/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include"

type Parser struct {
	Options core.Options
}

func New(options core.Options) *Parser {
	return &Parser{Options: options}
}

func (p Parser) initCmdArgs(cmdArgs []string) []string {
	parsePathEnv := func(p string, index map[string]bool) {
		paths := strings.Split(p, ":")
		for _, pth := range paths {
			index[pth] = true
		}
	}

	switch runtime.GOOS {
	case "linux":
		includes := make(map[string]bool)
		if path, ok := os.LookupEnv("C_INCLUDE_PATH"); ok {
			parsePathEnv(path, includes)
		}
		if path, ok := os.LookupEnv("CPLUS_INCLUDE_PATH"); ok {
			parsePathEnv(path, includes)
		}
		if len(includes) == 0 {
			includes["/usr/include"] = true
			includes["/usr/local/include"] = true
		}

		var args []string
		for _, arg := range cmdArgs {
			if strings.HasPrefix(arg, "-I") {
				paths := make(map[string]bool)
				parsePathEnv(strings.TrimSuffix(arg, "-I"), paths)
				for path := range paths {
					if _, ok := includes[path]; !ok {
						includes[path] = true
					}
				}
			} else {
				args = append(args, arg)
			}
		}

		var paths []string
		for path := range includes {
			paths = append(paths, path)
		}
		args = append(args, "-I"+strings.Join(paths, ":"))
		return args
	case "darwin":
		included := false
		for _, arg := range cmdArgs {
			if arg == "-I"+macOSDefaultIncludeDir {
				included = true
			}
		}
		if !included {
			cmdArgs = append(cmdArgs, "-I"+macOSDefaultIncludeDir)
		}
		return cmdArgs
	case "windows":
	}
	return cmdArgs
}

func (p Parser) ParseFile(fileName string, cmdArgs []string) (*lang.SourceFile, error) {
	idx := clang.NewIndex(1, 0)
	defer idx.Dispose()

	p.initCmdArgs(cmdArgs)
	tu := idx.ParseTranslationUnit(fileName, cmdArgs, nil, 0)
	defer tu.Dispose()

	diagnostics := tu.Diagnostics()
	for i, d := range diagnostics {
		if i < 8 {
			logs.Warnw("found diagnostics", "problem", d.Spelling())
		}
	}

	sourceFile := &lang.SourceFile{
		Name: fileName, // tu.File(fileName).Name()
	}

	// tu.FindIncludesInFile(tu.File(fileName))
	structs := make(map[string]*lang.StructDecl)
	enums := make(map[string]*lang.EnumDecl)

	cursor := tu.TranslationUnitCursor()
	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			logs.Warnw("cursor: <none>")
			return clang.ChildVisit_Continue
		}

		if !cursor.Location().IsFromMainFile() {
			return clang.ChildVisit_Continue
		}

		logs.Debugw("visit a cursor", "cursor", cursor.Spelling(), "kind", cursor.Kind().Spelling(), "USR", cursor.USR())

		switch cursor.Kind() {
		case clang.Cursor_ClassDecl:
			return clang.ChildVisit_Recurse
		case clang.Cursor_EnumDecl:
			typ := cursor.Type()
			decl := &lang.EnumDecl{
				Name: typ.Spelling(),
				Type: &lang.EnumType{},
			}
			enums[decl.Name] = decl
			sourceFile.Statements = append(sourceFile.Statements, lang.NewEnumDeclStatement(decl))
			return clang.ChildVisit_Recurse
		case clang.Cursor_EnumConstantDecl:
			name := cursor.Spelling()
			value := cursor.EnumConstantDeclValue()
			en := parent.Type().Spelling()
			if e, ok := enums[en]; ok {
				decl := &lang.ValueDecl{
					Name: name,
				}
				if value != 0 {
					decl.Initializer = &lang.Initializer{Value: lang.NewIntegerLiteralExpressionFrom(value)}
				}
				e.Type.Enumerators = append(e.Type.Enumerators, decl)
			}
			return clang.ChildVisit_Recurse
		case clang.Cursor_StructDecl:
			typ := cursor.Type()
			decl := &lang.StructDecl{
				Name: typ.Spelling(),
				Type: &lang.StructType{},
			}
			structs[decl.Name] = decl
			sourceFile.Statements = append(sourceFile.Statements, lang.NewStructDeclStatement(decl))
			return clang.ChildVisit_Recurse
		case clang.Cursor_FieldDecl:
			typ := cursor.Type()
			tName := typ.Spelling()
			fName := cursor.Spelling()
			sn := parent.Type().Spelling()
			if s, ok := structs[sn]; ok {
				decl := &lang.ValueDecl{
					Name:       fName,
					Attributes: nil,
					Type:       &lang.NominalType{Name: tName},
				}
				s.Type.Fields = append(s.Type.Fields, decl)
			}
			return clang.ChildVisit_Recurse
		case clang.Cursor_Namespace:
			return clang.ChildVisit_Recurse
		case clang.Cursor_UnionDecl:
			return clang.ChildVisit_Recurse
		case clang.Cursor_TypedefDecl:
			typ := cursor.Type()
			decl := &lang.TypeAliasDecl{
				Name: typ.Spelling(),
				Type: &lang.NominalType{
					Name: cursor.TypedefDeclUnderlyingType().Spelling(),
				},
			}
			sourceFile.Statements = append(sourceFile.Statements, lang.NewTypeAliasDeclStatement(decl))
		case clang.Cursor_InclusionDirective:
			// FIXME currently the include directive can't hit
		case clang.Cursor_FunctionDecl:
			decl := &lang.FunctionDecl{
				Name:      cursor.Spelling(),
				Signature: &lang.FunctionSignature{},
			}

			typ := cursor.Type()
			num := cursor.NumArguments()
			var args []*lang.ValueDecl
			for i := int32(0); i < num; i++ {
				cr := cursor.Argument(uint32(i))
				name := cr.Spelling()
				t := typ.ArgType(uint32(i)).Spelling()

				args = append(args, &lang.ValueDecl{
					Name: name,
					Type: &lang.NominalType{Name: t},
				})
			}
			if len(args) > 0 {
				decl.Signature.Parameter = lang.NewFunctionParameters(args...)
			}

			result := typ.ResultType().Spelling()
			decl.Signature.Result = lang.NewFunctionResult(&lang.NominalType{
				Name: result,
			})

			sourceFile.Statements = append(sourceFile.Statements, lang.NewFunctionDeclStatement(decl))
		}

		return clang.ChildVisit_Continue
	})

	if len(diagnostics) > 0 {
		logs.Warnw("There were problems while analyzing the given file")
	}

	return sourceFile, nil
}
