package lang

import (
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("lang.Declaration", &DeclarationCodec{})
	jsoniter.RegisterTypeEncoder("lang.Declaration", &DeclarationCodec{})
}

type DeclarationCodec struct {
}

func (codec *DeclarationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	declaration := (*Declaration)(ptr)
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("@type").ToString()
		switch t {
		case PackageDeclName:
			decl := &PackageDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_PackageDecl{PackageDecl: decl}
		case ImportDeclName:
			decl := &ImportDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_ImportDecl{ImportDecl: decl}
		case EnumDeclName:
			decl := &EnumDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_EnumDecl{EnumDecl: decl}
		case StructDeclName:
			decl := &StructDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_StructDecl{StructDecl: decl}
		case InterfaceDeclName:
			decl := &InterfaceDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_InterfaceDecl{InterfaceDecl: decl}
		case TypeAliasDeclName:
			decl := &TypeAliasDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_TypeAliasDecl{TypeAliasDecl: decl}
		case ConstantDeclName:
			decl := &ConstantDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_ConstantDecl{ConstantDecl: decl}
		case VariableDeclName:
			decl := &VariableDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_VariableDecl{VariableDecl: decl}
		case AttributeDeclName:
			decl := &AttributeDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_AttributeDecl{AttributeDecl: decl}
		case AttributeAliasDeclName:
			decl := &AttributeAliasDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_AttributeAliasDecl{AttributeAliasDecl: decl}
		case FunctionDeclName:
			decl := &FunctionDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_FunctionDecl{FunctionDecl: decl}
		case ConstructorDeclName:
			decl := &ConstructorDecl{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_ConstructorDecl{ConstructorDecl: decl}
		case GenericParameterName:
			decl := &GenericParameter{}
			a.ToVal(decl)
			declaration.Declaration = &Declaration_GenericParameter{GenericParameter: decl}
		}
	}
}

func (codec *DeclarationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	declaration := (*Declaration)(ptr)
	if declaration != nil && declaration.Declaration != nil {
		stream.WriteObjectStart()

		typ := ""
		var v reflect.Value
		switch decl := declaration.Declaration.(type) {
		case *Declaration_PackageDecl:
			v = reflect.ValueOf(decl.PackageDecl).Elem()
			typ = PackageDeclName
		case *Declaration_ImportDecl:
			v = reflect.ValueOf(decl.ImportDecl).Elem()
			typ = ImportDeclName
		case *Declaration_EnumDecl:
			v = reflect.ValueOf(decl.EnumDecl).Elem()
			typ = EnumDeclName
		case *Declaration_StructDecl:
			v = reflect.ValueOf(decl.StructDecl).Elem()
			typ = StructDeclName
		case *Declaration_InterfaceDecl:
			v = reflect.ValueOf(decl.InterfaceDecl).Elem()
			typ = InterfaceDeclName
		case *Declaration_TypeAliasDecl:
			v = reflect.ValueOf(decl.TypeAliasDecl).Elem()
			typ = TypeAliasDeclName
		case *Declaration_ConstantDecl:
			v = reflect.ValueOf(decl.ConstantDecl).Elem()
			typ = ConstantDeclName
		case *Declaration_VariableDecl:
			v = reflect.ValueOf(decl.VariableDecl).Elem()
			typ = VariableDeclName
		case *Declaration_AttributeDecl:
			v = reflect.ValueOf(decl.AttributeDecl).Elem()
			typ = AttributeDeclName
		case *Declaration_AttributeAliasDecl:
			v = reflect.ValueOf(decl.AttributeAliasDecl).Elem()
			typ = AttributeAliasDeclName
		case *Declaration_FunctionDecl:
			v = reflect.ValueOf(decl.FunctionDecl).Elem()
			typ = FunctionDeclName
		case *Declaration_ConstructorDecl:
			v = reflect.ValueOf(decl.ConstructorDecl).Elem()
			typ = ConstructorDeclName
		case *Declaration_GenericParameter:
			v = reflect.ValueOf(decl.GenericParameter).Elem()
			typ = GenericParameterName
		}

		if v.IsValid() && !v.IsZero() {
			stream.WriteObjectField("@type")
			stream.WriteVal(typ)

			t := declInfos[v.Type()]

			// iterator the fields
			for i := 0; i < v.NumField(); i++ {
				info := t.Fields[i]
				field := v.Field(i)

				if info.Ignored {
					continue
				}

				if info.OmitEmpty {
					if field.IsZero() {
						continue
					}

					k := field.Kind()
					switch k {
					case reflect.Map, reflect.Ptr, reflect.Interface, reflect.Struct, reflect.Slice:
						if info.Encoder.IsEmpty(unsafe.Pointer(field.Pointer())) {
							continue
						}
					}
				}

				stream.WriteMore()
				stream.WriteObjectField(t.Fields[i].ToName)
				stream.WriteVal(v.Field(i).Interface())
			}
		}

		stream.WriteObjectEnd()

		if stream.Error != nil && stream.Error != io.EOF {
			// stream.Error = fmt.Errorf("%v.%s", encoder.typ, stream.Error.Error())
		}
	}
}

func (codec *DeclarationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Declaration)(ptr)).Declaration == nil
}
