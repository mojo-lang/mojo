package generator

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
)

type Compiler struct {
	Services []*data.Service
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	services, err := compiler.CompilePackage(ctx, pkg)
	if err != nil {
		return err
	}

	c.Services = append(c.Services, services...)
	return nil
}

//
//type Compiler struct {
//	Path    string
//	Package *lang.Package
//	Files   []*descriptor.File
//
//	Data *data1.Data
//}
//
//func NewCompiler(path string, files []*descriptor.File) *Compiler {
//	return &Compiler{
//		Path:  path,
//		Files: files,
//		Data:  data1.NewData(),
//	}
//}
//
//func (c *Compiler) CompilePackage(pkg *lang.Package) (util.GeneratedFiles, error) {
//	var files util.GeneratedFiles
//
//	ctx := context.WithType(context.Empty(), pkg)
//	if pkg.DirectlyContainsService() {
//		spc := &compiler.ServicePackage{Data: c.Data}
//		if err := spc.CompilePackage(context.Empty(), pkg); err != nil {
//			return nil, err
//		}
//
//		sp := c.Data.Packages[len(c.Data.Packages)-1]
//		for _, p := range pkg.Children {
//			err := c.compilePackage(ctx, p, sp)
//			if err != nil {
//				return nil, err
//			}
//		}
//	} else {
//		for _, p := range pkg.Children {
//			if len(p.Children) > 0 {
//				spc := &compiler.ServicePackage{Data: c.Data}
//				if err := spc.CompilePackage(context.Empty(), p); err != nil {
//					return nil, err
//				}
//
//				sp := c.Data.Packages[len(c.Data.Packages)-1]
//				for _, cp := range p.Children {
//					err := c.compilePackage(context.WithType(ctx, p), cp, sp)
//					if err != nil {
//						return nil, err
//					}
//				}
//			}
//		}
//	}
//
//	d := &compiler.Data{Data: c.Data}
//	if err := d.CompilePackage(context.Empty(), pkg); err != nil {
//		return nil, err
//	}
//
//	return files, nil
//}
//
//func (c *Compiler) compilePackage(ctx context.Context, pkg *lang.Package, sp *data1.ServicePackage) error {
//	thisCtx := context.WithType(ctx, pkg)
//	for _, file := range pkg.SourceFiles {
//		fileCtx := context.WithType(thisCtx, file)
//
//		sp.JavaServicePackageName += "." + pkg.Name
//
//		// compile statements
//		for _, statement := range file.Statements {
//			switch statement.Statement.(type) {
//			case *lang.Statement_Declaration:
//				if decl := statement.GetDeclaration(); decl != nil {
//					var err error
//					switch decl.Declaration.(type) {
//					case *lang.Declaration_InterfaceDecl:
//						service := &httpd.Service{
//							PackageName:            sp.PackageName,
//							FullPackageName:        sp.FullPackageName,
//							JavaPackageName:        sp.JavaPackageName,
//							JavaServicePackageName: sp.JavaServicePackageName,
//							Methods:                nil,
//						}
//
//						ic := http.Interface{Service: service}
//						if err = ic.CompileInterface(fileCtx, decl.GetInterfaceDecl()); err != nil {
//							return err
//						}
//
//						sp.Services = append(sp.Services, service)
//					}
//
//					if err != nil {
//						return err
//					}
//				}
//			}
//		}
//	}
//
//	return nil
//}
