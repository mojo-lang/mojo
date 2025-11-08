package compiler

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/decompiler"
)

type Struct struct {
	*data.Data
}

var whiteList = map[string]bool{}

func (s *Struct) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}
	if whiteList[decl.Name] {
		return nil
	}

	pkg := context.Package(ctx)
	thisCtx := context.WithType(ctx, decl)

	if pkg.GetFullName() != decl.GetPackageName() {
		return nil
	}

	cmp := &DbJson{Data: s.Data}
	if err := cmp.CompileStruct(thisCtx, decl); err != nil {
		return err
	}

	formatComp := &FormatJson{Data: s.Data}
	if err := formatComp.CompileStruct(thisCtx, decl); err != nil {
		return err
	}

	if boxed, err := lang.GetBoolAttribute(decl.Attributes, core.BoxedAttributeName); boxed && err == nil {
		switch decl.Type.Fields[0].Type.GetFullName() {
		case core.ArrayTypeFullName:
			ba := &BoxedArray{Data: s.Data}
			if err = ba.CompileStruct(thisCtx, decl); err != nil {
				return err
			}
		case core.MapTypeFullName:
			bm := &BoxedMap{Data: s.Data}
			if err = bm.CompileStruct(thisCtx, decl); err != nil {
				return err
			}
		case core.UnionTypeFullName:
			bu := &BoxedUnion{Data: s.Data}
			if err = bu.CompileStruct(thisCtx, decl); err != nil {
				return err
			}
		}
	}

	if union, _ := lang.GetBoolAttribute(decl.Attributes, "union"); union {
		bu := &BoxedUnion{Data: s.Data}
		if err := bu.CompileStruct(ctx, decl); err != nil {
			return err
		}
		return nil
	}

	for _, d := range decl.StructDecls {
		if err := s.CompileStruct(thisCtx, d); err != nil {
			return errors.Wrapf(err, "failed to parse the inner struct decl %s in %s", d.Name, decl.Name)
		}
	}

	for _, e := range decl.EnumDecls {
		enum := &Enum{Data: s.Data}
		if err := enum.CompileEnum(thisCtx, e); err != nil {
			return errors.Wrapf(err, "failed to parse the inner enum decl %s in %s", e.Name, decl.Name)
		}
	}

	if decl.Type != nil {
		if decl.IsBoxed() {
			valueType := decl.Type.Inherits[0]
			valueName := "val"
			if fullName := valueType.GetFullName(); fullName == core.ArrayTypeFullName || fullName == core.MapTypeFullName {
				valueName = "vals"
			}
			valueDecl := &lang.ValueDecl{
				Implicit: true,
				Name:     valueName,
				Type:     valueType,
			}
			valueType.Attributes = lang.SetIntegerAttribute(valueType.Attributes, core.NumberAttributeName, 1)
			if err := s.compileStructFields(thisCtx, []*lang.ValueDecl{valueDecl}); err != nil {
				return err
			}

			boxedStruct := &lang.StructDecl{
				PackageName:    decl.PackageName,
				SourceFileName: decl.SourceFileName,
				Implicit:       true,
				Name:           decl.Name,
				Enclosing:      decl.Enclosing,
				Type: &lang.StructType{
					Fields: []*lang.ValueDecl{valueDecl},
				},
			}
			boxedStruct.Attributes = lang.SetBoolAttribute(boxedStruct.Attributes, core.BoxedAttributeName, true)
			return s.CompileStruct(thisCtx, boxedStruct)
		} else {
			if err := s.compileStructFields(thisCtx, decl.Type.Fields); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Struct) compileStructField(ctx context.Context, field *lang.ValueDecl) error {
	thisCtx := context.WithType(ctx, field)

	switch field.Type.Name {
	case core.ArrayTypeName:
		if len(field.Type.GenericArguments) == 1 {
			argument := field.Type.GenericArguments[0]
			_, err := s.compileNominalType(thisCtx, argument)
			if err != nil {
				return errors.Wrapf(err, "failed to compile the type %s", argument.Name)
			}
			return nil
		} else {
			return errors.Errorf("unexpect array type of %s", field.Type.Name)
		}
	case core.UnionTypeName:
		// /TODO GenericArguments == 1
		if len(field.Type.GenericArguments) > 0 {
			for _, argument := range field.Type.GenericArguments {
				_, err := s.compileNominalType(thisCtx, argument)
				if err != nil {
					return errors.Wrapf(err, "failed to compile the type %s", argument.Name)
				}
			}
		}
	case core.MapTypeName:
		if len(field.Type.GenericArguments) == 2 {
			for _, argument := range field.Type.GenericArguments {
				_, err := s.compileNominalType(thisCtx, argument)
				if err != nil {
					return errors.Wrapf(err, "failed to compile the type %s", argument.Name)
				}
			}
			return nil
		} else {
			return errors.Errorf("unexpect map type of %s", field.Type.Name)
		}
	default:
	}

	// alias, err := lang.GetStringAttribute(field.Type.Attributes, "alias")

	return nil
}

func (s *Struct) compileStructFields(ctx context.Context, fields []*lang.ValueDecl) error {
	for _, field := range fields {
		thisCtx := context.WithType(ctx, field)
		err := s.compileStructField(thisCtx, field)
		if err != nil {
			return err
		}
	}
	return nil
}

// 处理protobuf不支持的类型，将其box成struct类型
// auto_generate struct
func (s *Struct) compileNominalType(ctx context.Context, t *lang.NominalType) (string, error) { // typeName, error
	pkg := context.Package(ctx)
	if t.IsScalar() {
		return "", nil
	}

	setAttributes := func(s *lang.StructDecl) {
		s.Attributes = lang.SetBoolAttribute(s.Attributes, core.BoxedAttributeName, true)
		s.Attributes = lang.SetBoolAttribute(s.Attributes, "auto_generated", true)
		goPkgName := pkg.GoPackageName()
		if goPkgName != pkg.Name {
			s.Attributes = lang.SetStringAttribute(s.Attributes, "go_package_name", goPkgName)
		}
	}

	if t.Name == core.ArrayTypeName {
		if len(t.GenericArguments) != 1 {
			return "", errors.New("")
		}

		t.Attributes = lang.SetIntegerAttribute(t.Attributes, core.NumberAttributeName, 1)

		decl := &lang.StructDecl{
			Implicit: true,
		}
		decl.Type = &lang.StructType{
			Fields: []*lang.ValueDecl{{
				Name: "vals",
				Type: t,
			}},
		}
		decl.PackageName = pkg.GetFullName()
		decl.Name = transformer.Plural(strcase.ToCamel(t.GenericArguments[0].Name))

		setAttributes(decl)

		// TODO set to be boxed type
		err := s.CompileStruct(ctx, decl)
		if err != nil {
			return "", errors.Wrapf(err, "failed to compile the arry field in %s", decl.Name)
		}
		return decl.Name, nil
	} else if t.Name == core.MapTypeName {
		if len(t.GenericArguments) != 2 {
			return "", errors.New("")
		}

		valType := t.GenericArguments[1]
		if len(valType.GenericArguments) > 0 {
			name, err := s.compileNominalType(ctx, valType)
			if err != nil {
				return "", err
			}
			valType.Name = name
		}

		decl := &lang.StructDecl{}
		decl.PackageName = pkg.GetFullName()
		decl.Type = &lang.StructType{
			Fields: []*lang.ValueDecl{{
				Name: "vals",
				Type: t,
			}},
		}
		t.Attributes = lang.SetIntegerAttribute(t.Attributes, core.NumberAttributeName, 1)

		setAttributes(decl)

		err := s.CompileStruct(ctx, decl)
		if err != nil {
			return "", errors.Wrapf(err, "failed to compile the map field in %s", decl.Name)
		}
		return decl.Name, nil
	} else if t.Name == core.UnionTypeName {
		decl, _ := decompiler.CompileUnion(ctx, t)
		if decl != nil { // skip if directly union declaration
			for _, field := range decl.Type.Fields {
				if _, err := s.compileNominalType(ctx, field.Type); err != nil {
					return "", nil
				}
			}
			decl.PackageName = pkg.GetFullName()

			setAttributes(decl)
			decl.Attributes = lang.SetBoolAttribute(decl.Attributes, "union", true)

			err := s.CompileStruct(ctx, decl)
			if err != nil {
				return "", errors.Wrap(err, "failed to compile struct")
			}

			return decl.Name, nil
		}
	}

	return t.Name, nil
}
