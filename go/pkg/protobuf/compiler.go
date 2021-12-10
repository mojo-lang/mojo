package protobuf

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	langcompiler "github.com/mojo-lang/mojo/go/pkg/compiler"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/compiler"
	"github.com/mojo-lang/mojo/go/pkg/util"
	desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
	"sort"
	"strings"
)

var _ = compiler.ArrayPlugin{}
var _ = compiler.MapPlugin{}
var _ = compiler.TuplePlugin{}
var _ = compiler.UnionPlugin{}

type Compiler struct {
	Files   []*desc.FileDescriptor
}

func NewCompiler() *Compiler {
	c := &Compiler{}
	c.Files = make([]*desc.FileDescriptor, 0)
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

func (c *Compiler) CompilePackage(pkg *lang.Package) error {
	return c.compilePackage(context.Empty(), pkg)
}

func (c *Compiler) CompileFile(file *lang.SourceFile) error {
	_, err := c.compileFile(context.Empty(), file)
	return err
}

func (c *Compiler) compilePackage(ctx context.Context, pkg *lang.Package) error {
	thisCtx := context.WithType(ctx, pkg)

	for _, sourceFile := range pkg.SourceFiles {
		descriptor, err := c.compileFile(thisCtx, sourceFile)
		if err != nil {
			return err
		}
		if descriptor != nil {
			c.Files = append(c.Files, descriptor)
		}
	}
	return nil
}

func (c *Compiler) compileFile(ctx context.Context, file *lang.SourceFile) (*desc.FileDescriptor, error) {
	if file.IsGenericInstantiated() {
		return nil, nil
	}

	fileDescriptor := desc.NewFileDescriptor()
	thisCtx := context.WithDescriptor(context.WithType(ctx, file), fileDescriptor)
	fileDescriptor.Proto3 = true

	name := strings.TrimSuffix(file.FullName, ".mojo") + ".proto"
	fileDescriptor.Name = &name

	// compile package decl
	if err := c.compilePackageDecl(file, fileDescriptor); err != nil {
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
				err := c.compileTypeAlias(thisCtx, decl.GetTypeAliasDecl(), desc.NewMessageDescriptor(fileDescriptor))
				if err != nil {
					return nil, err
				}
			case *lang.Declaration_StructDecl:
				err := c.compileStruct(thisCtx, decl.GetStructDecl(), desc.NewMessageDescriptor(fileDescriptor))
				if err != nil {
					return nil, err
				}
			case *lang.Declaration_EnumDecl:
				err := c.compileEnum(thisCtx, decl.GetEnumDecl(), desc.NewEnumDescriptor(fileDescriptor))
				if err != nil {
					return nil, err
				}
			case *lang.Declaration_InterfaceDecl:
				err := c.compileInterface(thisCtx, decl.GetInterfaceDecl(), desc.NewServiceDescriptor(fileDescriptor))
				if err != nil {
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

func (c *Compiler) compilePackageDecl(file *lang.SourceFile, descriptor *desc.FileDescriptor) error {
	descriptor.Package = &file.PackageName
	return nil
}

func (c *Compiler) compileFileOptions(ctx context.Context, file *lang.SourceFile, fileDescriptor *desc.FileDescriptor) error {
	pkg := context.Package(ctx)
	if pkg == nil {
		return errors.New("has no package found")
	}

	fileDescriptor.Options = &descriptor.FileOptions{}

	repository := pkg.Repository
	if repository != nil {
		goPackage := fmt.Sprintf("%s;%s", pkg.GoFullPackageName(), pkg.GoPackageName())
		fileDescriptor.Options.GoPackage = &goPackage
	}

	javaMultipleFiles := true
	javaPackage := pkg.FullName

	organizationPrefix := getOrganizationPrefix(pkg)
	if len(organizationPrefix) > 0 {
		javaPackage = organizationPrefix + "." + javaPackage
	}

	javaClassName := strcase.ToCamel(strings.TrimSuffix(file.Name, ".mojo")) + "Proto"

	fileDescriptor.Options.JavaMultipleFiles = &javaMultipleFiles
	fileDescriptor.Options.JavaPackage = &javaPackage
	fileDescriptor.Options.JavaOuterClassname = &javaClassName
	return nil
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

func (c *Compiler) compileImport(file *lang.SourceFile, descriptor *desc.FileDescriptor) error {
	for _, dependency := range file.ResolvedIdentifiers {
		if dependency.IsGenericInstantiated() {
			continue
		}

		fileName := dependency.SourceFileName
		if strings.HasSuffix(fileName, ".mojo") {
			fileName = strings.TrimSuffix(fileName, "mojo") + "proto"
		} else {
			fileName = fileName + ".proto"
		}

		if !compiler.IsSystemFile(fileName) && fileName != *descriptor.Name {
			descriptor.Dependency = append(descriptor.Dependency, fileName)
		}
	}

	descriptor.Dependency = util.UniqueStringSlice(descriptor.Dependency)
	sort.Strings(descriptor.Dependency)
	return nil
}

func (c *Compiler) compileEnum(ctx context.Context, decl *lang.EnumDecl, descriptor *desc.EnumDescriptor) error {
	return compiler.CompileEnum(ctx, decl, descriptor)
}

func (c *Compiler) compileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl, descriptor *desc.MessageDescriptor) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	decl.Type.Attributes = lang.SetStringAttribute(decl.Type.Attributes, lang.OriginalTypeAliasName, decl.Name)
	_, _, err := compiler.CompileNominalType(ctx, decl.Type)
	return err
}

func (c *Compiler) compileStruct(ctx context.Context, decl *lang.StructDecl, descriptor *desc.MessageDescriptor) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	return compiler.CompileStruct(ctx, decl, descriptor)
}

func (c *Compiler) compileInterface(ctx context.Context, decl *lang.InterfaceDecl, serviceDescriptor *desc.ServiceDescriptor) error {
	thisCtx := context.WithDescriptor(context.WithType(ctx, decl), serviceDescriptor)
	serviceDescriptor.Name = &decl.Name

	//if i.Document != nil {
	//	for _, l := range i.Document.Lines {
	//		service.Document = append(service.Document, l.Line)
	//	}
	//}

	for _, method := range decl.Type.Methods {
		err := c.compileMethod(thisCtx, method, serviceDescriptor)
		if err != nil {
			return err
		}
	}

	file := context.FileDescriptor(ctx)
	if file != nil {
		file.Services = append(file.Services, serviceDescriptor)
	}
	return nil
}

func (c *Compiler) compileMethod(ctx context.Context, method *lang.FunctionDecl, serviceDescriptor *desc.ServiceDescriptor) error {
	m := &descriptor.MethodDescriptorProto{
		Name:       &method.Name,
		InputType:  nil,
		OutputType: nil,
		Options:    nil,
	}

	// generate the request message
	s := &lang.StructDecl{}
	s.Name = strcase.ToCamel(method.Name) + "Request"

	// add number attribute if there is no field has number attribute
	//for i, param := range method.Type.Parameters {
	//	param.Attributes
	//}

	s.Type = &lang.StructType{
		Fields: method.Signature.Parameters,
	}

	if pagination, _ := lang.GetBoolAttribute(method.Attributes, "pagination"); pagination {
		s.Type.Fields = append(s.Type.Fields, langcompiler.GeneratePaginationParameters()...)
	}

	file := context.FileDescriptor(ctx)
	c.compileStruct(ctx, s, desc.NewMessageDescriptor(file))

	m.InputType = &s.Name

	// FIXME move to the resolver of semantic parser
	var nullTypeName = "mojo.core.Null"
	result := method.Signature.Result
	if result == nil {
		m.OutputType = &nullTypeName

		if file != nil {
			file.Dependency = append(file.Dependency, "mojo/core/null.proto")
		}
	} else {
		if result.Name == "Tuple" {
			name := strcase.ToCamel(method.Name) + "Response"
			result.Attributes = lang.SetStringAttribute(result.Attributes, "alias", name)
		}

		_, name, err := compiler.CompileNominalType(ctx, result)
		if err != nil {
			return errors.New(
				fmt.Sprintf("failed to compile the type %s: %s", result.Name, err.Error()))
		}

		//TODO the type should be the struct
		m.OutputType = &name
	}

	// options

	serviceDescriptor.Method = append(serviceDescriptor.Method, m)
	return nil
}
