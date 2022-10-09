package generator

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	compiler2 "github.com/mojo-lang/mojo/go/pkg/go/generator/compiler"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	generator2 "github.com/mojo-lang/mojo/go/pkg/go/generator/generator"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Compiler struct {
	Path    string
	Package *lang.Package
	Files   []*descriptor.File

	Data *data.Data
}

func NewCompiler(path string, files []*descriptor.File) *Compiler {
	return &Compiler{
		Path:  path,
		Files: files,
		Data:  data.NewData(),
	}
}

func (c *Compiler) CompilePackage(pkg *lang.Package) (util.GeneratedFiles, error) {
	var files util.GeneratedFiles

	fs, err := generator2.ProtocGenGo(c.Path, pkg, c.Files)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		if len(f.Content) > 0 {
			f.Content = generator2.FormatCode(f.Content)
		}
		files = append(files, f)
	}

	// other compilers
	err = c.compilePackage(context.Empty(), pkg)
	if err != nil {
		return nil, err
	}

	for _, p := range pkg.Children {
		err = c.compilePackage(context.Empty(), p)
		if err != nil {
			return nil, err
		}
	}

	gm := &compiler2.GoMod{Data: c.Data}
	if err = gm.CompilePackage(context.Empty(), pkg); err != nil {
		return nil, err
	}

	return files, nil
}

func (c *Compiler) compilePackage(ctx context.Context, pkg *lang.Package) error {
	thisCtx := context.WithType(ctx, pkg)
	for _, file := range pkg.SourceFiles {
		fileCtx := context.WithType(thisCtx, file)

		// compile statements
		for _, statement := range file.Statements {
			switch statement.Statement.(type) {
			case *lang.Statement_Declaration:
				if decl := statement.GetDeclaration(); decl != nil {
					var err error
					switch decl.Declaration.(type) {
					case *lang.Declaration_TypeAliasDecl:
						s := compiler2.TypeAlias{Data: c.Data}
						err = s.CompileTypeAlias(fileCtx, decl.GetTypeAliasDecl())
					case *lang.Declaration_StructDecl:
						s := compiler2.Struct{Data: c.Data}
						err = s.CompileStruct(fileCtx, decl.GetStructDecl())
					case *lang.Declaration_EnumDecl:
						e := compiler2.Enum{Data: c.Data}
						err = e.CompileEnum(fileCtx, decl.GetEnumDecl())
					case *lang.Declaration_InterfaceDecl:
						i := compiler2.Interface{Data: c.Data}
						err = i.CompileInterface(fileCtx, decl.GetInterfaceDecl())
					}

					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
