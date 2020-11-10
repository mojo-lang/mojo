package protobuf

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/lang"
)

type Compiler struct {
	File *ProtoFile
}

func New() *Compiler {
	compiler := &Compiler{}
	compiler.File = &ProtoFile{}
	return compiler
}

func (c Compiler) CompileFile(file *lang.SourceFile) (*ProtoFile, error) {
	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			switch decl.Declaration.(type) {
			case *lang.Declaration_StructDecl:
				message , err := c.compileStruct(decl.GetStructDecl())
				if err != nil {
					return nil, err
				}
				c.File.Messages = append(c.File.Messages, message)
			case *lang.Declaration_EnumDecl:
				enum, err := c.compileEnum(decl.GetEnumDecl())
				if err != nil {
					return nil, err
				}
				c.File.Enums = append(c.File.Enums, enum)
			case *lang.Declaration_InterfaceDecl:
				service, err := c.compileInterface(decl.GetInterfaceDecl())
				if err != nil {
					return nil, err
				}
				c.File.Services = append(c.File.Services, service)
			case *lang.Declaration_PackageDecl:
			case *lang.Declaration_ImportDecl:
			}
		case *lang.Statement_Expression:
		default:
		}
	}
	return c.File, nil
}

func (c Compiler) compileEnum(decl *lang.EnumDecl) (*ProtoEnum, error) {
	return nil, nil
}

func (c Compiler) compileStruct(decl *lang.StructDecl) (*ProtoMessage, error) {
	message := &ProtoMessage{}
	message.Name = decl.Name

	for _, field := range decl.Type.Fields {
		member := &ProtoMessageField{}
		member.Name = field.Name

		switch field.Type.Name {
		case "Array":
			if len(field.Type.GenericArguments) > 0 {
				argument := field.Type.GenericArguments[0]
				member.Type = fmt.Sprint("repeated ", argument.Name)
			}
		case "Dictionary":
			if len(field.Type.GenericArguments) > 1 {
				key := field.Type.GenericArguments[0]
				value := field.Type.GenericArguments[1]
				member.Type = fmt.Sprintf("map<%s,%s>", key.Name, value.Name)
			}
		case "Union":
		default:
			member.Type = field.Type.Name
		}

		number, err := lang.GetIntegerAttribute(field.Type.Attributes, "number")
		if err != nil {
			return nil, errors.New("has not set the number in field")
		}
		if number <= 0 {
			return nil, errors.New("number attribute value must be positive")
		}
		member.Number = int32(number)

		message.Fields = append(message.Fields, member)
	}
	return message, nil
}

func (c Compiler) compileInterface(i *lang.InterfaceDecl) (*ProtoService, error) {
	service := &ProtoService{}
	if i == nil {
		return service, nil
	}

	service.Name = i.Name

	if i.Document != nil {
		for _, l := range i.Document.Lines {
			service.Document = append(service.Document, l.Line)
		}
	}

	for _, method := range i.Type.Methods {
		m, err := c.compileMethod(method)
		if err != nil {
			return nil, err
		}

		service.Methods = append(service.Methods, m)
	}

	return service, nil
}

func (c Compiler) compileMethod(method *lang.FuncDecl) (*ProtoMethod, error) {
	m := &ProtoMethod{
		Name: method.Name,
	}

	// generate the request message & response message
	s := &lang.StructDecl{}
	s.Name = method.Name + "Request"
	s.Type = &lang.StructType{
		Fields: method.Type.Parameters,
	}

	request, err := c.compileStruct(s)
	if err != nil {
		return nil, errors.New("")
	}

	c.File.Messages = append(c.File.Messages, request)
	m.InputType = s.Name

	if method.Type.ReturnType == nil {
		m.OutputType = "mojo.core.Null"
	} else {

	}

	return m, nil
}
