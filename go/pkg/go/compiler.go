package _go

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/go/compiler"
	protocompiler "github.com/mojo-lang/mojo/go/pkg/protobuf/compiler"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/pkg/errors"
)

type Compiler struct {
	Path    string
	Package *lang.Package
	Files   []*descriptor.FileDescriptor

	Data *compiler.Data
}

func NewCompiler(path string, pkg *lang.Package, files []*descriptor.FileDescriptor) *Compiler {
	return &Compiler{
		Path:    path,
		Package: pkg,
		Files:   files,
		Data:    compiler.NewData(),
	}
}

func (c *Compiler) Compile() (util.GeneratedFiles, error) {
	var files util.GeneratedFiles

	fs, err := compiler.ProtocGo(c.Path, c.Package, c.Files)
	if err != nil {
		return nil, err
	}

	files = append(files, fs...)

	// other compilers
	err = c.compilePackage(&context.Context{}, c.Package)
	if err != nil {
		return nil, err
	}

	for _, pkg := range c.Package.Children {
		err = c.compilePackage(&context.Context{}, pkg)
		if err != nil {
			return nil, err
		}
	}

	err = c.compileGoMod(&context.Context{}, c.Package)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (c *Compiler) compileGoMod(ctx *context.Context, pkg *lang.Package) error {
	_ = ctx
	c.Data.GoMod = &compiler.GoMod{}
	return c.Data.GoMod.Compile(pkg)
}

func (c *Compiler) compilePackage(ctx *context.Context, pkg *lang.Package) error {
	ctx.Open(pkg)
	defer func() {
		ctx.Close()
	}()

	cmp := &compiler.Compiler{Data: c.Data}
	err := cmp.Compile(pkg)
	if err != nil {
		return err
	}

	for _, sourceFile := range pkg.SourceFiles {
		err = c.compileFile(ctx, sourceFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) compileFile(ctx *context.Context, file *lang.SourceFile) error {
	ctx.Open(file)
	defer func() {
		ctx.Close()
	}()

	// compile statements
	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			if decl == nil {
				return errors.New("declaration statement is nil")
			}

			var err error
			switch decl.Declaration.(type) {
			case *lang.Declaration_StructDecl:
				err = c.compileStruct(ctx, decl.GetStructDecl())
			case *lang.Declaration_EnumDecl:
				err = c.compileEnum(ctx, decl.GetEnumDecl())
			case *lang.Declaration_InterfaceDecl:
				err = c.compileInterface(ctx, decl.GetInterfaceDecl())
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}

var whiteList = map[string]bool{
	"Strings":       true,
	"StringValues":  true,
	"IntegerValues": true,
	"DoubleValues":  true,
}

func (c *Compiler) compileStruct(ctx *context.Context, decl *lang.StructDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	if whiteList[decl.Name] {
		return nil
	}

	if boxed, err := lang.GetBoolAttribute(decl.Attributes, "boxed"); boxed && err == nil {
		cmp := &compiler.Compiler{Data: c.Data}
		return cmp.Compile(decl)
	}

	for _, s := range decl.StructDecls {
		err := c.compileStruct(ctx, s)
		if err != nil {
			return errors.Wrapf(err, "failed to parse the inner struct decl %s in %s", s.Name, decl.Name)
		}
	}
	for _, e := range decl.EnumDecls {
		err := c.compileEnum(ctx, e)
		if err != nil {
			return errors.Wrapf(err, "failed to parse the inner enum decl %s in %s", e.Name, decl.Name)
		}
	}

	if decl.Type != nil {
		if decl.IsBoxedType() {
			valueType := decl.Type.Inherits[0]
			valueName := "val"
			if fullName := valueType.GetFullName(); fullName == "mojo.core.Array" || fullName == "mojo.core.Map" {
				valueName = "vals"
			}
			valueDecl := &lang.ValueDecl{
				Implicit: true,
				Name:     valueName,
				Type:     valueType,
			}
			valueType.Attributes = lang.SetIntegerAttribute(valueType.Attributes, "number", 1)
			if err := c.compileStructFields(ctx, []*lang.ValueDecl{valueDecl}); err != nil {
				return err
			}

			boxedStruct := &lang.StructDecl{
				PackageName:    decl.PackageName,
				SourceFileName: decl.SourceFileName,
				Implicit:       true,
				Name:           decl.Name,
				EnclosingType:  decl.EnclosingType,
				Type: &lang.StructType{
					Fields: []*lang.ValueDecl{valueDecl},
				},
			}
			boxedStruct.Attributes = lang.SetBoolAttribute(boxedStruct.Attributes, "boxed", true)
			return c.compileStruct(ctx, boxedStruct)
		} else {
			if err := c.compileStructFields(ctx, decl.Type.Fields); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Compiler) compileEnum(ctx *context.Context, decl *lang.EnumDecl) error {
	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	cmp := &compiler.Compiler{Data: c.Data}
	return cmp.Compile(decl)
}

func (c *Compiler) compileStructField(ctx *context.Context, field *lang.ValueDecl) error {
	ctx.Open(field)
	defer func() {
		ctx.Close()
	}()

	switch field.Type.Name {
	case "Array":
		if len(field.Type.GenericArguments) > 0 {
			argument := field.Type.GenericArguments[0]
			_, err := c.compileNominalType(ctx, argument)
			if err != nil {
				return errors.Wrapf(err, "failed to compile the type %s", argument.Name)
			}
			return nil
		} else {
			return errors.Errorf("unexpect array type of %s", field.Type.Name)
		}
	case "Union":
		///TODO GenericArguments == 1
		// TODO 如果number标号在Union类型上，则需要将该Union转换成message；如果不是则使用oneof
		if len(field.Type.GenericArguments) > 0 {
			for i, argument := range field.Type.GenericArguments {
				_, err := c.compileNominalType(ctx, argument)
				if err != nil {
					return errors.Wrapf(err, "failed to compile the type %s: %s", argument.Name)
				}

				if i < len(field.Type.GenericArguments)-1 {
					ctx.Open(field)
				}
			}
		}
	case "Map":

	default:
	}

	//alias, err := lang.GetStringAttribute(field.Type.Attributes, "alias")

	return nil
}

func (c *Compiler) compileStructFields(ctx *context.Context, fields []*lang.ValueDecl) error {
	for _, field := range fields {
		err := c.compileStructField(ctx, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) compileInterface(ctx *context.Context, decl *lang.InterfaceDecl) error {
	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	for _, method := range decl.Type.Methods {
		result := method.Signature.Result
		if result != nil {
			_, err := c.compileNominalType(ctx, result)
			if err != nil {
				return errors.Errorf("failed to compile the type %s: %s", result.Name, err.Error())
			}
		}
	}

	return nil
}

// 处理protobuf不支持的类型，将其box成struct类型
// auto_generate struct
func (c *Compiler) compileNominalType(ctx *context.Context, t *lang.NominalType) (string, error) { // typeName, error
	pkg := ctx.GetPackage()
	if t.IsScalar() {
		return "", nil
	}

	setAttributes := func(s *lang.StructDecl) {
		s.Attributes = lang.SetBoolAttribute(s.Attributes, "boxed", true)
		s.Attributes = lang.SetBoolAttribute(s.Attributes, "auto_generated", true)
		goPkgName := pkg.GoPackageName()
		if goPkgName != pkg.Name {
			s.Attributes = lang.SetStringAttribute(s.Attributes, "go_package_name", goPkgName)
		}
	}

	if t.Name == "Array" {
		if len(t.GenericArguments) != 1 {
			return "", errors.New("")
		}

		t.Attributes = lang.SetIntegerAttribute(t.Attributes, "number", 1)

		s := &lang.StructDecl{
			Implicit: true,
		}
		s.Type = &lang.StructType{
			Fields: []*lang.ValueDecl{{
				Name: "vals",
				Type: t,
			}},
		}
		s.PackageName = pkg.GetFullName()
		s.Name = transformer.Plural(strcase.ToCamel(t.GenericArguments[0].Name))

		setAttributes(s)

		//TODO set to be boxed type
		err := c.compileStruct(ctx, s)
		if err != nil {
			return "", errors.Wrapf(err, "failed to compile the arry field in %s", s.Name)
		}
		return s.Name, nil
	} else if t.Name == "Map" {
		if len(t.GenericArguments) != 2 {
			return "", errors.New("")
		}

		valType := t.GenericArguments[1]
		if len(valType.GenericArguments) > 0 {
			name, err := c.compileNominalType(ctx, valType)
			if err != nil {
				return "", err
			}
			valType.Name = name
		}

		s := &lang.StructDecl{}
		s.PackageName = pkg.GetFullName()
		s.Type = &lang.StructType{
			Fields: []*lang.ValueDecl{{
				Name: "vals",
				Type: t,
			}},
		}
		t.Attributes = lang.SetIntegerAttribute(t.Attributes, "number", 1)

		setAttributes(s)

		err := c.compileStruct(ctx, s)
		if err != nil {
			return "", errors.Wrapf(err, "failed to compile the map field in %s", s.Name)
		}
		return s.Name, nil
	} else if t.Name == "Union" {
		s := protocompiler.ConstructBoxedUnion(t)
		if s != nil { // skip if directly union declaration
			for _, field := range s.Type.Fields {
				c.compileNominalType(ctx, field.Type)
			}
			s.PackageName = pkg.GetFullName()

			setAttributes(s)
			s.Attributes = lang.SetBoolAttribute(s.Attributes, "union", true)

			err := c.compileStruct(ctx, s)
			if err != nil {
				return "", errors.Wrap(err, "failed to compile struct")
			}

			return s.Name, nil
		}
	}

	return t.Name, nil
}
