package generator

import (
	"errors"
	"fmt"
	compiler2 "github.com/mojo-lang/mojo/go/pkg/protobuf/generator/compiler"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Compiler struct {
	Descriptors *descriptor.Packages
}

func NewCompiler() *Compiler {
	c := &Compiler{
		Descriptors: descriptor.NewPackages(),
	}
	return c
}

func (c *Compiler) CompilePackages(packages map[string]*lang.Package) error {
	for _, pkg := range packages {
		err := c.compilePackage(context.Empty(), pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	thisCtx := context.WithType(ctx, pkg)

	for _, sourceFile := range pkg.SourceFiles {
		descriptor, err := c.compileFile(thisCtx, sourceFile)
		if err != nil {
			return err
		}
		if descriptor != nil {
			c.Descriptors.AddFile(descriptor)
		}
	}

	for _, child := range pkg.Children {
		if err := c.CompilePackage(thisCtx, child); err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) CompileFile(file *lang.SourceFile) (*descriptor.File, error) {
	return c.compileFile(context.Empty(), file)
}

func (c *Compiler) compilePackage(ctx context.Context, pkg *lang.Package) error {
	thisCtx := context.WithType(ctx, pkg)

	for _, sourceFile := range pkg.SourceFiles {
		descriptor, err := c.compileFile(thisCtx, sourceFile)
		if err != nil {
			return err
		}
		if descriptor != nil {
			c.Descriptors.AddFile(descriptor)
		}
	}
	return nil
}

func (c *Compiler) compileFile(ctx context.Context, file *lang.SourceFile) (*descriptor.File, error) {
	if file.IsGenericInstantiated() {
		return nil, nil
	}

	name := strings.TrimSuffix(file.FullName, ".mojo") + ".proto"

	fileDescriptor := descriptor.NewFile()
	fileDescriptor.SetProto3(true)
	fileDescriptor.SetName(name)

	// compile package decl
	thisCtx := context.WithDescriptor(context.WithType(ctx, file), fileDescriptor)
	if err := c.compilePackageDecl(thisCtx, file, fileDescriptor); err != nil {
		return nil, err
	}

	if err := c.compileFileOptions(thisCtx, file, fileDescriptor); err != nil {
		return nil, err
	}

	// compile statements
	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			if decl == nil {
				return nil, errors.New("declaration statement is nil")
			}

			switch decl.Declaration.(type) {
			case *lang.Declaration_TypeAliasDecl:
				if err := c.compileTypeAlias(thisCtx, decl.GetTypeAliasDecl(), descriptor.NewMessage(fileDescriptor)); err != nil {
					return nil, err
				}
			case *lang.Declaration_StructDecl:
				if err := c.compileStruct(thisCtx, decl.GetStructDecl(), descriptor.NewMessage(fileDescriptor)); err != nil {
					return nil, err
				}
			case *lang.Declaration_EnumDecl:
				if err := c.compileEnum(thisCtx, decl.GetEnumDecl(), descriptor.NewEnum(fileDescriptor)); err != nil {
					return nil, err
				}
			case *lang.Declaration_InterfaceDecl:
				if err := c.compileInterface(thisCtx, decl.GetInterfaceDecl(), descriptor.NewService(fileDescriptor)); err != nil {
					return nil, err
				}
			}
		case *lang.Statement_Expression:
		default:
		}
	}

	// compile imports
	if err := c.compileImport(file, fileDescriptor); err != nil {
		return nil, err
	}

	return fileDescriptor, nil
}

func (c *Compiler) compilePackageDecl(ctx context.Context, file *lang.SourceFile, descriptor *descriptor.File) error {
	descriptor.SetPackageName(file.PackageName)
	return nil
}

func (c *Compiler) compileFileOptions(ctx context.Context, file *lang.SourceFile, fileDescriptor *descriptor.File) error {
	pkg := context.Package(ctx)
	if pkg == nil {
		return errors.New("has no package found")
	}

	options := fileDescriptor.GetOptions()

	repository := pkg.Repository
	if repository != nil {
		goPackage := fmt.Sprintf("%s;%s", pkg.GoFullPackageName(), pkg.GoPackageName())
		options.GoPackage = &goPackage
	}

	javaMultipleFiles := true
	javaPackage := pkg.FullName

	organizationPrefix := getOrganizationPrefix(pkg)
	if len(organizationPrefix) > 0 {
		javaPackage = organizationPrefix + "." + javaPackage
	}

	javaClassName := strcase.ToCamel(strings.TrimSuffix(file.Name, ".mojo")) + "Proto"

	options.JavaMultipleFiles = &javaMultipleFiles
	options.JavaPackage = &javaPackage
	options.JavaOuterClassname = &javaClassName
	return nil
}

func (c *Compiler) compileImport(file *lang.SourceFile, descriptor *descriptor.File) error {
	for _, dependency := range file.ResolvedIdentifiers {
		if dependency.IsGenericInstantiated() {
			continue
		}

		fileName := compiler2.GetProtoFile(dependency.SourceFileName)

		if !compiler2.IsSystemFile(fileName) && fileName != descriptor.GetName() {
			descriptor.AppendDependency(fileName)
		}
	}

	descriptor.CleanDependency()
	return nil
}

func (c *Compiler) compileEnum(ctx context.Context, decl *lang.EnumDecl, descriptor *descriptor.Enum) error {
	if options, _ := lang.GetDisableGenerateAttribute(decl.Attributes); options.Including("protobuf", "") {
		return nil
	}

	return compiler2.Enum{}.Compile(ctx, decl, descriptor)
}

func (c *Compiler) compileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl, descriptor *descriptor.Message) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	if options, _ := lang.GetDisableGenerateAttribute(decl.Attributes); options.Including("protobuf", "") {
		return nil
	}

	return compiler2.TypeAlias{}.Compile(ctx, decl, descriptor)
}

func (c *Compiler) compileStruct(ctx context.Context, decl *lang.StructDecl, descriptor *descriptor.Message) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	if options, _ := lang.GetDisableGenerateAttribute(decl.Attributes); options.Including("protobuf", "") {
		return nil
	}

	return compiler2.Struct{}.Compile(ctx, decl, descriptor)
}

func (c *Compiler) compileInterface(ctx context.Context, decl *lang.InterfaceDecl, descriptor *descriptor.Service) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	if options, _ := lang.GetDisableGenerateAttribute(decl.Attributes); options.Including("protobuf", "") {
		return nil
	}

	return compiler2.Interface{}.Compile(ctx, decl, descriptor)
}

func getOrganizationPrefix(pkg *lang.Package) string {
	organization := pkg.GetOrganization()
	if len(organization) == 0 {
		if strings.HasPrefix(pkg.FullName, "mojo.") {
			organization = "mojolang.org"
		} else {
			logs.Warnw("you should add organization to package.mojo file, for build java files from protobuf", "package", pkg.FullName)
			return ""
		}
	}

	segments := strings.Split(organization, " ") // remove Co. Ltd in company Co. Ltd
	segments = strings.Split(segments[0], ".")   // separate company.com
	if len(segments) == 2 {
		return segments[1] + "." + segments[0]
	}
	return "com." + segments[0]
}
