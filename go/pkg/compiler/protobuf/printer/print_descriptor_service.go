package printer

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func (p *Printer) PrintDescriptorService(ctx context.Context, service *descriptor.Service) *Printer {
	_ = ctx

	p.PrintLine("service ", service.GetName(), " {")
	p.Indent()

	pkg := service.GetPackageName()
	for _, method := range service.Methods {
		input := method.GetInput()
		inputName := input.GetName()
		if input.GetPackageName() != pkg {
			inputName = input.GetFullName()
		}
		output := method.GetOutput()
		outputName := output.GetName()
		if output.GetPackageName() != pkg {
			outputName = output.GetFullName()
		}
		p.PrintLine("rpc ", method.GetName(), "(", inputName, ") returns (", outputName, ");")
	}

	p.Outdent()
	p.PrintLine("}")

	return p
}
