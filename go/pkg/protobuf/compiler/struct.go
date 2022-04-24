package compiler

import (
    "errors"
    "fmt"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/db/go/pkg/mojo/db"
    "google.golang.org/protobuf/runtime/protoimpl"
    "strings"

    "github.com/mojo-lang/core/go/pkg/mojo"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Struct struct {
}

const inheritSourceFileKey = "inherit-source-file"

func (s Struct) Compile(ctx context.Context, decl *lang.StructDecl, structDescriptor *descriptor.Message) error {
    thisCtx := context.WithDescriptor(context.WithType(ctx, decl), structDescriptor)

    file := context.FileDescriptor(ctx)
    structDescriptor.SetName(decl.Name)

    {
        enumCtx := context.WithValues(thisCtx, "register_enum", false)
        for _, e := range decl.EnumDecls {
            enum := descriptor.NewEnum(file)
            err := Enum{}.Compile(enumCtx, e, enum)
            if err != nil {
                return errors.New(fmt.Sprintf("failed to parse the inner enum decl %s in %s: %s", e.Name, decl.Name, err.Error()))
            }
            structDescriptor.AppendInnerEnum(enum)
        }
    }

    {
        structCtx := context.WithValues(thisCtx, "register_struct", false)
        for _, d := range decl.StructDecls {
            msg := descriptor.NewMessage(file)
            err := s.Compile(structCtx, d, msg)
            if err != nil {
                return errors.New(fmt.Sprintf("failed to parse the inner struct decl %s in %s: %s", d.Name, decl.Name, err.Error()))
            }

            structDescriptor.AppendMessage(msg)
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
            if err := s.compileStructFields(thisCtx, []*lang.ValueDecl{valueDecl}, structDescriptor); err != nil {
                return err
            }
        } else {
            for _, inherit := range decl.Type.Inherits {
                //ctx.SetOption(inheritSourceFileKey, make(map[string]bool))
                if err := s.compileStructInherit(thisCtx, inherit, structDescriptor); err != nil {
                    //ctx.DeleteOption(inheritSourceFileKey)
                    return err
                }
                //ctx.DeleteOption(inheritSourceFileKey)
            }

            if err := s.compileStructFields(thisCtx, decl.Type.Fields, structDescriptor); err != nil {
                return err
            }
        }
    }

    switch structDescriptor.GetName() {
    case "BoolValues", "StringValues", "Int32Values", "UInt32Values", "Int64Values", "UInt64Values", "Float32Values", "Float64Values":
        if file != nil && !strings.HasPrefix(file.GetName(), "mojo/core/boxed") {
            file.AppendDependency("mojo/core/boxed.proto")
            return nil
        }
    }

    if register, ok := ctx.Value("register_struct").(bool); !ok || register {
        if msg := context.MessageDescriptor(ctx); msg != nil {
            if !msg.IsMessageExist(structDescriptor.GetName()) {
                msg.AppendMessage(structDescriptor)
            }
        } else if file != nil {
            if !file.IsMessageExist(structDescriptor.GetName()) {
                file.Messages = append(file.Messages, structDescriptor)
            }
        }
    }

    return nil
}

//TODO 不在相同的package下的inherit不需要进行Boxed类型的处理
func (s Struct) compileStructInherit(ctx context.Context, inherit *lang.NominalType, msgDescriptor *descriptor.Message) error {
    decl := inherit.TypeDeclaration.GetStructDecl()
    if decl == nil || decl.Type == nil {
        return nil
    }

    for _, i := range decl.Type.Inherits {
        if err := s.compileStructInherit(ctx, i, msgDescriptor); err != nil {
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
        if file != nil && !IsSystemFile(fileName) && fileName != file.GetName() /*&& !sourceFiles[fileName]*/ {
            file.AppendDependency(fileName)
        }
    }

    return s.compileStructFields(ctx, decl.Type.Fields, msgDescriptor)
}

func (s Struct) compileStructFields(ctx context.Context, fields []*lang.ValueDecl, msgDescriptor *descriptor.Message) error {
    file := context.SourceFile(ctx)

    for _, field := range fields {
        member := descriptor.NewField(msgDescriptor, field.Name)
        fieldCtx := context.WithDescriptor(context.WithType(ctx, field), member)

        switch field.Type.Name {
        case "Array":
            if len(field.Type.GenericArguments) > 0 {
                argument := field.Type.GenericArguments[0]
                member.SetRepeated()
                t, name, err := Nominal{}.Compile(fieldCtx, argument)
                if err != nil {
                    return errors.New(
                        fmt.Sprintf("failed to compile the type %s: %s", argument.Name, err.Error()))
                }
                member.SetType(t).SetTypeName(name)
            } else {
                return errors.New(fmt.Sprintf("unexpect array type of %s", field.Type.Name))
            }
        case "Union":
            ///TODO GenericArguments == 1
            // TODO 如果number标号在Union类型上，则需要将该Union转换成message；如果不是则使用oneof
            if len(field.Type.GenericArguments) > 0 {
                if _, err := field.GetIntegerAttribute(core.NumberAttributeName); err == nil {
                    t, name, err := Nominal{}.Compile(fieldCtx, field.Type)
                    if err != nil {
                        return errors.New(
                            fmt.Sprintf("failed to compile the type %s: %s", field.Type.Name, err.Error()))
                    }
                    member.SetType(t).SetTypeName(name)
                    break
                }

                msgDescriptor.AppendOneofWith(field.Name)
                oneof := msgDescriptor.GetLastOneof()
                for _, argument := range field.Type.GenericArguments {
                    var argumentCtx context.Context
                    if member == nil {
                        member = descriptor.NewField(msgDescriptor, "")
                        argumentCtx = context.WithDescriptor(fieldCtx, member)
                    } else {
                        argumentCtx = fieldCtx
                    }

                    t, name, err := Nominal{}.Compile(argumentCtx, argument)
                    if err != nil {
                        return errors.New(
                            fmt.Sprintf("failed to compile the type %s: %s", argument.Name, err.Error()))
                    }
                    member.SetType(t)
                    member.SetTypeName(name)
                    member.SetName(strcase.ToSnake(name))

                    number, err := lang.GetIntegerAttribute(argument.Attributes, core.NumberAttributeName)
                    if err != nil {
                        return logs.NewErrorw("has not set the number in field", "field", field.Name, "struct", msgDescriptor.GetName(), "file", file.GetFullName(), "line", field.GetStartPosition().GetLine())
                    }
                    if number <= 0 {
                        return logs.NewErrorw("number attribute value must be positive", "field", field.Name, "struct", msgDescriptor.GetName(), "file", file.GetFullName(), "line", field.GetStartPosition().GetLine())
                    }

                    member.SetNumber(int32(number))

                    oneof.AppendField(member)
                    msgDescriptor.AppendField(member)
                    member = nil
                }

                continue
            }
        default:
            t, name, err := Nominal{}.Compile(fieldCtx, field.Type)
            if err != nil {
                return errors.New(
                    fmt.Sprintf("failed to compile the type %s: %s", field.Type.Name, err.Error()))
            }

            member.SetType(t).SetTypeName(name)
        }

        number, err := lang.GetIntegerAttribute(field.Type.Attributes, core.NumberAttributeName)
        if err != nil {
            return logs.NewErrorw("has not set the number in field", "field", field.Name, "struct", msgDescriptor.GetName(), "file", file.GetFullName(), "line", field.GetStartPosition().GetLine())
        }
        if number <= 0 {
            return errors.New("number attribute value must be positive")
        }
        member.SetNumber(int32(number))

        setStringOption := func(attribute string, info *protoimpl.ExtensionInfo) {
            if field.HasAttribute(attribute) {
                option := ""
                option, _ = field.GetStringAttribute(attribute)
                member.SetStringOption(info, option)
                addOptionsDependency(fieldCtx)
            }
        }

        setStringOption(core.AliasAttributeName, mojo.E_Alias)
        setStringOption(core.KeyAttributeName, mojo.E_Key)
        setStringOption(core.ReferenceAttributeName, mojo.E_Reference)
        setStringOption(core.BackReferenceAttributeName, mojo.E_BackReference)

        if lang.HasAttribute(field.Type.Attributes, db.JSONAttributeName) {
            member.SetBoolOption(mojo.E_DbJson, true)
            addOptionsDependency(fieldCtx)
        }
        if lang.HasAttribute(field.Type.Attributes, db.IgnoreAttributeName) {
            member.SetBoolOption(mojo.E_DbIgnore, true)
            addOptionsDependency(fieldCtx)
        }
        if index, err := lang.GetStringAttribute(field.Type.Attributes, db.IndexAttributeName); err == nil {
            member.SetStringOption(mojo.E_DbIndex, index)
            addOptionsDependency(fieldCtx)
        }
        if explode, err := lang.GetStringAttribute(field.Type.Attributes, db.ExplodeAttributeName); err == nil {
            member.SetStringOption(mojo.E_DbExplode, explode)
            addOptionsDependency(fieldCtx)
        }

        msgDescriptor.AppendField(member)
    }
    return nil
}

const OptionsDependency = "mojo/mojo.proto"

func addOptionsDependency(ctx context.Context) {
    fileDescriptor := context.FileDescriptor(ctx)
    for _, dep := range fileDescriptor.GetDependencies() {
        if dep == OptionsDependency {
            return
        }
    }

    fileDescriptor.AppendDependency(OptionsDependency)
}
