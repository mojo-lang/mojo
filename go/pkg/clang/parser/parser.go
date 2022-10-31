package parser

import (
	"github.com/go-clang/clang-v13/clang"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type Parser struct {
}

func New(options core.Options) *Parser {
	return &Parser{}
}

func (p Parser) ParseFile(fileName string, cmdArgs []string) (*lang.SourceFile, error) {
	idx := clang.NewIndex(1, 0)
	defer idx.Dispose()

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
		case clang.Cursor_ClassDecl, clang.Cursor_EnumDecl, clang.Cursor_StructDecl, clang.Cursor_Namespace, clang.Cursor_UnionDecl:
			return clang.ChildVisit_Recurse
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

// func (p Parser) ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (*lang.SourceFile, error) {
//    if bytes, err := fs.ReadFile(fileSys, fileName); err != nil {
//        return nil, err
//    } else {
//    }
// }
