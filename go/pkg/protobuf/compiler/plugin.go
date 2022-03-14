package compiler

import (
    "errors"
    "fmt"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/db/go/pkg/mojo/db"
    "google.golang.org/protobuf/runtime/protoimpl"
    "strings"

    "github.com/golang/protobuf/protoc-gen-go/descriptor"
    "github.com/mojo-lang/core/go/pkg/mojo"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Plugin interface {
    Compile(ctx context.Context, t *lang.NominalType) (string, string, error)
}

var plugins = make(map[string][]Plugin)

// type alias transform
// type: Scalar, Enum, Struct
func CompileNominalType(ctx context.Context, t *lang.NominalType) (string, string, error) { // type, typeName, error
    pkg := context.Package(ctx)
    getName := func() string {
        if pkg != nil && pkg.FullName == t.PackageName {
            return t.Name
        } else {
            return t.GetFullName()
        }
    }

    if t.IsScalar() {
        return "Scalar", t.Name, nil
    }

    tp := "Struct"

    if decl := t.TypeDeclaration; decl != nil {
        if aliasDecl := decl.GetTypeAliasDecl(); aliasDecl != nil {
            if aliasDecl.Type != nil {
                //	// type_alias_original_name
                //	_, err := lang.GetStringAttribute(aliasDecl.Type.Attributes, lang.OriginalTypeAliasName)
                //	if err != nil { // not found
                //		aliasDecl.Type.Attributes = lang.SetStringAttribute(aliasDecl.Type.Attributes, lang.OriginalTypeAliasName, t.Name)
                //	}
                //
                //	return CompileNominalType(ctx, aliasDecl.Type)
            }
        }

        if decl.GetEnumDecl() != nil {
            tp = "Enum"
        }
    }

    ps := plugins[t.GetFullName()]
    for _, p := range ps {
        tp, typeName, err := p.Compile(ctx, t)
        if err == nil && len(typeName) > 0 {
            return tp, typeName, err
        }
    }

    return tp, getName(), nil
}

func CompileEnum(ctx context.Context, decl *lang.EnumDecl, enumDescriptor *desc.EnumDescriptor) error {
    thisCtx := context.WithDescriptor(context.WithType(ctx, decl), enumDescriptor)
    enumDescriptor.Name = &decl.Name

    for i, e := range decl.Type.Enumerators {
        name := decl.Name + "_" + e.Name
        name = strcase.ToScreamingSnake(name)
        value := &descriptor.EnumValueDescriptorProto{
            Name: &name,
            //Number:               nil,
            //Options:              nil,
        }

        number, err := lang.GetIntegerAttribute(e.Attributes, core.NumberAttributeName)
        if err == nil {
            if number < 0 {
                return errors.New("number attribute value must be positive")
            }
            n := int32(number)
            value.Number = &n
        } else {
            n := int32(i)
            value.Number = &n
        }

        enumDescriptor.Value = append(enumDescriptor.Value, value)
    }

    message := context.MessageDescriptor(thisCtx)
    file := context.FileDescriptor(thisCtx)
    if message == nil && file != nil {
        if register, ok := ctx.Value("register_enum").(bool); !ok || register {
            file.Enums = append(file.Enums, enumDescriptor)
        }
    }

    return nil
}

const inheritSourceFileKey = "inherit-source-file"

func CompileStruct(ctx context.Context, decl *lang.StructDecl, structDescriptor *desc.MessageDescriptor) error {
    thisCtx := context.WithDescriptor(context.WithType(ctx, decl), structDescriptor)

    file := context.FileDescriptor(ctx)
    structDescriptor.Name = &decl.Name

    {
        enumCtx := context.WithValues(thisCtx, "register_enum", false)
        for _, e := range decl.EnumDecls {
            enum := desc.NewEnumDescriptor(file)
            err := CompileEnum(enumCtx, e, enum)
            if err != nil {
                return errors.New(fmt.Sprintf("failed to parse the inner enum decl %s in %s: %s", e.Name, decl.Name, err.Error()))
            }
            structDescriptor.AddInnerEnum(enum)
        }
    }

    {
        structCtx := context.WithValues(thisCtx, "register_struct", false)
        for _, s := range decl.StructDecls {
            msg := desc.NewMessageDescriptor(file)
            err := CompileStruct(structCtx, s, msg)
            if err != nil {
                return errors.New(fmt.Sprintf("failed to parse the inner struct decl %s in %s: %s", s.Name, decl.Name, err.Error()))
            }

            structDescriptor.AddInnerMessage(msg)
        }
    }

    if decl.Type != nil {
        if decl.IsBoxed() {
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
            valueType.Attributes = lang.SetIntegerAttribute(valueType.Attributes, core.NumberAttributeName, 1)
            if err := compileStructFields(thisCtx, []*lang.ValueDecl{valueDecl}, structDescriptor); err != nil {
                return err
            }
        } else {
            for _, inherit := range decl.Type.Inherits {
                //ctx.SetOption(inheritSourceFileKey, make(map[string]bool))
                if err := compileStructInherit(thisCtx, inherit, structDescriptor); err != nil {
                    //ctx.DeleteOption(inheritSourceFileKey)
                    return err
                }
                //ctx.DeleteOption(inheritSourceFileKey)
            }

            if err := compileStructFields(thisCtx, decl.Type.Fields, structDescriptor); err != nil {
                return err
            }
        }
    }

    if *structDescriptor.Name == "Strings" || *structDescriptor.Name == "StringValues" || *structDescriptor.Name == "IntegerValues" || *structDescriptor.Name == "DoubleValues" {
        if file != nil && !strings.HasPrefix(*file.Name, "mojo/core/boxed") {
            file.Dependency = append(file.Dependency, "mojo/core/boxed.proto")
            return nil
        }
    }

    if register, ok := ctx.Value("register_struct").(bool); !ok || register {
        if msg := context.MessageDescriptor(ctx); msg != nil {
            if !msg.IsMessageExist(*structDescriptor.Name) {
                msg.AddInnerMessage(structDescriptor)
            }
        } else if file != nil {
            if !file.IsMessageExist(*structDescriptor.Name) {
                file.Messages = append(file.Messages, structDescriptor)
            }
        }
    }

    return nil
}

//TODO 不在相同的package下的inherit不需要进行Boxed类型的处理
func compileStructInherit(ctx context.Context, inherit *lang.NominalType, descriptor *desc.MessageDescriptor) error {
    decl := inherit.TypeDeclaration.GetStructDecl()
    if decl == nil || decl.Type == nil {
        return nil
    }

    for _, i := range decl.Type.Inherits {
        if err := compileStructInherit(ctx, i, descriptor); err != nil {
            return err
        }
    }

    unifyFileName := func(fileName string) string {
        if strings.HasSuffix(fileName, ".mojo") {
            return strings.TrimSuffix(fileName, "mojo") + "proto"
        } else {
            return fileName + ".proto"
        }
    }

    file := context.FileDescriptor(ctx)
    //FIXME remove the inherit dependency file to protobuf file imports
    //sourceFiles, ok := ctx.GetOption(inheritSourceFileKey).(map[string]bool)
    //if !ok {
    //	sourceFiles = make(map[string]bool)
    //}
    //sourceFiles[unifyFileName(decl.SourceFileName)] = true

    for _, dependency := range decl.ResolvedIdentifiers {
        fileName := unifyFileName(dependency.SourceFileName)
        if file != nil && !IsSystemFile(fileName) && fileName != *file.Name /*&& !sourceFiles[fileName]*/ {
            file.Dependency = append(file.Dependency, fileName)
        }
    }

    return compileStructFields(ctx, decl.Type.Fields, descriptor)
}

func compileStructFields(ctx context.Context, fields []*lang.ValueDecl, msgDescriptor *desc.MessageDescriptor) error {
    scope := lang.GetScope(context.TypeValue(ctx))
    file := context.SourceFile(ctx)

    for _, field := range fields {
        member := &descriptor.FieldDescriptorProto{}
        fieldCtx := context.WithDescriptor(context.WithType(ctx, field), member)
        member.Name = &field.Name

        if decl := field.Type.TypeDeclaration; decl != nil {
            if structDecl := decl.GetStructDecl(); structDecl != nil && structDecl.EnclosingType != nil {
                if scope == nil || scope.Identifiers == nil || scope.Identifiers[structDecl.Name] == nil {
                    field.Type.Name = lang.GetFullName("", lang.GetEnclosingNames(structDecl.EnclosingType), field.Type.Name)
                }
            } else if enumDecl := decl.GetEnumDecl(); enumDecl != nil && enumDecl.EnclosingType != nil {
                if scope == nil || scope.Identifiers == nil || scope.Identifiers[enumDecl.Name] == nil {
                    field.Type.Name = lang.GetFullName("", lang.GetEnclosingNames(enumDecl.EnclosingType), field.Type.Name)
                }
            }
        }

        switch field.Type.Name {
        case "Array":
            if len(field.Type.GenericArguments) > 0 {
                argument := field.Type.GenericArguments[0]
                repeated := descriptor.FieldDescriptorProto_LABEL_REPEATED
                member.Label = &repeated
                t, name, err := CompileNominalType(fieldCtx, argument)
                if err != nil {
                    return errors.New(
                        fmt.Sprintf("failed to compile the type %s: %s", argument.Name, err.Error()))
                }
                pType, pName := protoType(t, name)
                member.Type = &pType
                member.TypeName = &pName
            } else {
                return errors.New(fmt.Sprintf("unexpect array type of %s", field.Type.Name))
            }
        case "Union":
            ///TODO GenericArguments == 1
            // TODO 如果number标号在Union类型上，则需要将该Union转换成message；如果不是则使用oneof
            if len(field.Type.GenericArguments) > 0 {
                if _, err := field.GetIntegerAttribute(core.NumberAttributeName); err == nil {
                    t, name, err := CompileNominalType(fieldCtx, field.Type)
                    if err != nil {
                        return errors.New(
                            fmt.Sprintf("failed to compile the type %s: %s", field.Type.Name, err.Error()))
                    }

                    pType, pName := protoType(t, name)
                    member.Type = &pType
                    member.TypeName = &pName
                    break
                }

                msgDescriptor.OneofDecl = append(msgDescriptor.OneofDecl, &descriptor.OneofDescriptorProto{
                    Name: &field.Name,
                })

                index := int32(len(msgDescriptor.OneofDecl) - 1)
                for _, argument := range field.Type.GenericArguments {
                    var argumentCtx context.Context
                    if member == nil {
                        member = &descriptor.FieldDescriptorProto{}
                        argumentCtx = context.WithDescriptor(fieldCtx, member)
                    } else {
                        argumentCtx = fieldCtx
                    }

                    t, name, err := CompileNominalType(argumentCtx, argument)
                    if err != nil {
                        return errors.New(
                            fmt.Sprintf("failed to compile the type %s: %s", argument.Name, err.Error()))
                    }
                    pType, pName := protoType(t, name)
                    member.Type = &pType
                    member.TypeName = &pName
                    member.OneofIndex = &index

                    snakeName := strcase.ToSnake(pName)
                    member.Name = &snakeName

                    number, err := lang.GetIntegerAttribute(argument.Attributes, core.NumberAttributeName)
                    if err != nil {
                        return logs.NewErrorw("has not set the number in field", "field", field.Name, "struct", msgDescriptor.GetName(), "file", file.GetFullName(), "line", field.GetStartPosition().GetLine())
                    }
                    if number <= 0 {
                        return logs.NewErrorw("number attribute value must be positive", "field", field.Name, "struct", msgDescriptor.GetName(), "file", file.GetFullName(), "line", field.GetStartPosition().GetLine())
                    }
                    n := int32(number)
                    member.Number = &n

                    msgDescriptor.Field = append(msgDescriptor.Field, member)
                    member = nil
                }

                continue
            }
        default:
            t, name, err := CompileNominalType(fieldCtx, field.Type)
            if err != nil {
                return errors.New(
                    fmt.Sprintf("failed to compile the type %s: %s", field.Type.Name, err.Error()))
            }

            pType, pName := protoType(t, name)
            member.Type = &pType
            member.TypeName = &pName
        }

        number, err := lang.GetIntegerAttribute(field.Type.Attributes, core.NumberAttributeName)
        if err != nil {
            return logs.NewErrorw("has not set the number in field", "field", field.Name, "struct", msgDescriptor.GetName(), "file", file.GetFullName(), "line", field.GetStartPosition().GetLine())
        }
        if number <= 0 {
            return errors.New("number attribute value must be positive")
        }
        n := int32(number)
        member.Number = &n

        setStringOption := func(attribute string, info *protoimpl.ExtensionInfo) {
            if field.HasAttribute(attribute) {
                option := ""
                option, _ = field.GetStringAttribute(attribute)
                desc.SetStringFieldOption(info, option)(member)
                addOptionsDependency(fieldCtx)
            }
        }

        setStringOption(core.AliasAttributeName, mojo.E_Alias)
        setStringOption(core.KeyAttributeName, mojo.E_Key)
        setStringOption(core.ReferenceAttributeName, mojo.E_Reference)
        setStringOption(core.BackReferenceAttributeName, mojo.E_BackReference)

        if lang.HasAttribute(field.Type.Attributes, db.JSONAttributeName) {
            desc.SetBoolFieldOption(mojo.E_DbJson, true)(member)
            addOptionsDependency(fieldCtx)
        }
        if lang.HasAttribute(field.Type.Attributes, db.IgnoreAttributeName) {
            desc.SetBoolFieldOption(mojo.E_DbIgnore, true)(member)
            addOptionsDependency(fieldCtx)
        }
        if index, err := lang.GetStringAttribute(field.Type.Attributes, db.IndexAttributeName); err == nil {
            desc.SetStringFieldOption(mojo.E_DbIgnore, index)(member)
            addOptionsDependency(fieldCtx)
        }
        if explode, err := lang.GetStringAttribute(field.Type.Attributes, db.ExplodeAttributeName); err == nil {
            desc.SetStringFieldOption(mojo.E_DbExplode, explode)(member)
            addOptionsDependency(fieldCtx)
        }

        msgDescriptor.Field = append(msgDescriptor.Field, member)
    }
    return nil
}

const OptionsDependency = "mojo/mojo.proto"

func addOptionsDependency(ctx context.Context) {
    fileDescriptor := context.FileDescriptor(ctx)
    for _, dep := range fileDescriptor.Dependency {
        if dep == OptionsDependency {
            return
        }
    }

    fileDescriptor.Dependency = append(fileDescriptor.Dependency, OptionsDependency)
}

func protoType(t string, typeName string) (descriptor.FieldDescriptorProto_Type, string) {
    switch typeName {
    case "Double", "Float64":
        return descriptor.FieldDescriptorProto_TYPE_DOUBLE, ""
    case "Float", "Float32":
        return descriptor.FieldDescriptorProto_TYPE_FLOAT, ""
    case "Int64", "Int":
        return descriptor.FieldDescriptorProto_TYPE_INT64, ""
    case "UInt64", "UInt":
        return descriptor.FieldDescriptorProto_TYPE_UINT64, ""
    case "Int8", "Int16", "Int32":
        return descriptor.FieldDescriptorProto_TYPE_INT32, ""
    case "UInt8", "UInt16", "UInt32":
        return descriptor.FieldDescriptorProto_TYPE_INT32, ""
    case "Bool":
        return descriptor.FieldDescriptorProto_TYPE_BOOL, ""
    case "String":
        return descriptor.FieldDescriptorProto_TYPE_STRING, ""
    case "Bytes":
        return descriptor.FieldDescriptorProto_TYPE_BYTES, ""
    default:
        switch t {
        case "Enum":
            return descriptor.FieldDescriptorProto_TYPE_ENUM, typeName
        default:
            return descriptor.FieldDescriptorProto_TYPE_MESSAGE, typeName
        }
    }
}
