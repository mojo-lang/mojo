package converter

import (
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/decompiler"
)

type Interface struct {
}

func (i Interface) Compile(ctx context.Context, decl *lang.InterfaceDecl, descriptor *descriptor.Service) error {
	thisCtx := context.WithDescriptor(context.WithType(ctx, decl), descriptor)
	descriptor.SetName(decl.Name)

	// if i.Document != nil {
	//	for _, l := range i.Document.Lines {
	//		service.Document = append(service.Document, l.Line)
	//	}
	// }

	var methods []*lang.FunctionDecl
	// if ms := decl.GetInheritMethods(); len(ms) > 0 {
	//    methods = append(methods, ms...)
	// }
	if decl.Type != nil && len(decl.Type.Methods) > 0 {
		methods = append(methods, decl.Type.Methods...)
	}

	for _, method := range methods {
		err := i.compileMethod(thisCtx, method, descriptor)
		if err != nil {
			return err
		}
	}

	file := context.FileDescriptor(ctx)
	if file != nil {
		file.Services = append(file.Services, descriptor)
	}
	return nil
}

func (i Interface) compileMethod(ctx context.Context, method *lang.FunctionDecl, service *descriptor.Service) error {
	m := descriptor.NewMethod(service).SetName(method.Name)

	file := context.FileDescriptor(ctx)
	if req, resp, err := decompiler.CompileMethod(ctx, method); err != nil {
		return logs.NewErrorw("failed to compile method", "method", method.Name, "service", service.FullName, "error", err.Error())
	} else {
		var input, output *descriptor.Message
		var registerInput, registerOutput bool
		if req.Implicit && req.PackageName == file.GetPackageName() {
			input = descriptor.NewMessage(file)
			registerInput = true
		} else {
			file.AppendDependency(GetProtoFile(req.SourceFileName))
			input = descriptor.NewMessage(descriptor.NewFileWithName(req.SourceFileName, req.PackageName))
		}

		inputCtx := context.WithValues(ctx, "register_struct", registerInput)
		err = Struct{}.Compile(inputCtx, req, input)
		if err != nil {
			return logs.NewErrorw("failed to compile request", "request", req.Name, "method", method.Name, "service", service.FullName, "error", err.Error())
		}
		m.Input = input

		if resp.Implicit && resp.PackageName == file.GetPackageName() {
			output = descriptor.NewMessage(file)
			registerOutput = true
		} else {
			file.AppendDependency(GetProtoFile(resp.SourceFileName))
			output = descriptor.NewMessage(descriptor.NewFileWithName(resp.SourceFileName, resp.PackageName))
		}
		outputCtx := context.WithValues(ctx, "register_struct", registerOutput)
		err = Struct{}.Compile(outputCtx, resp, output)
		if err != nil {
			return logs.NewErrorw("failed to compile response", "response", resp.GetFullName(), "method", method.Name, "service", service.FullName, "error", err.Error())
		}
		m.Output = output
	}

	service.AppendMethod(m)
	return nil
}
