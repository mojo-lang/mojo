package generator

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/packages/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/compiler"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/generator"
	"github.com/mojo-lang/mojo/go/pkg/util"
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

	fs, err := generator.ProtocGenGo(c.Path, pkg, c.Files)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		if len(f.Content) > 0 {
			f.Content = generator.FormatCode(f.Content)
		}
		files = append(files, f)
	}

	// other compilers
	err = c.compilePackage(context.Empty(), pkg)
	if err != nil {
		return nil, err
	}

	ctx := context.WithType(context.Empty(), pkg)
	for _, p := range pkg.Children {
		err = c.compilePackage(ctx, p)
		if err != nil {
			return nil, err
		}

		for _, cp := range p.Children {
			err = c.compilePackage(context.WithType(ctx, p), cp)
			if err != nil {
				return nil, err
			}
		}
	}

	gm := &compiler.GoMod{Data: c.Data}
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
						s := compiler.TypeAlias{Data: c.Data}
						err = s.CompileTypeAlias(fileCtx, decl.GetTypeAliasDecl())
					case *lang.Declaration_StructDecl:
						s := compiler.Struct{Data: c.Data}
						err = s.CompileStruct(fileCtx, decl.GetStructDecl())
					case *lang.Declaration_EnumDecl:
						e := compiler.Enum{Data: c.Data}
						err = e.CompileEnum(fileCtx, decl.GetEnumDecl())
					case *lang.Declaration_InterfaceDecl:
						i := compiler.Interface{Data: c.Data}
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
