package printer

import (
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintDescriptorFile(ctx context.Context, file *descriptor.File) *Printer {
	for _, enum := range file.Enums {
		if enum.Parent == nil {
			p.PrintBlankLine()
			p.PrintDescriptorEnum(ctx, enum)
		}
	}

	for _, message := range file.Messages {
		if message.Parent == nil {
			if message.IsMapEntry() {
				continue
			}

			if isSystemMessage(message) && file.GetPackageName() == "mojo.core" {
				continue
			}

			p.PrintBlankLine()
			p.PrintDescriptorMessage(ctx, message)
		}
	}

	for _, service := range file.Services {
		p.PrintBlankLine()
		p.PrintDescriptorService(ctx, service)
	}

	// Run the plugins before the imports so we know which imports are necessary.
	// p.runPlugins(file)

	// Generate header and imports last, though they appear first in the output.
	rem := p.Buffer
	if rem.Len() == 0 {
		return p
	}

	p.Reset()

	p.PrintDescriptorFileHeader(ctx, file)

	if len(file.GetDependencies()) > 0 {
		p.PrintDescriptorFileImports(ctx, file)
	}

	if file.HasOptions() {
		p.PrintBlankLine()

		options := file.GetOptions()
		if file.HasGoOptions() {
			p.PrintLine(`option go_package = "`, options.GoPackage, `";`)
		}

		if file.HasJavaOptions() {
			if options.JavaMultipleFiles != nil {
				p.PrintLine("option java_multiple_files = ", options.JavaMultipleFiles, ";")
			}

			if options.JavaOuterClassname != nil {
				p.PrintLine(`option java_outer_classname = "`, options.JavaOuterClassname, `";`)
			}

			if options.JavaPackage != nil {
				p.PrintLine(`option java_package = "`, options.JavaPackage, `";`)
			}
		}
	}

	p.BreakLine()
	p.Buffer.Write(rem.Bytes())

	return p
}

// PrintDescriptorFileHeader print the header, including package definition
func (p *Printer) PrintDescriptorFileHeader(ctx context.Context, file *descriptor.File) {
	_ = ctx
	p.PrintLine("// Code generated by mojo. DO NOT EDIT.")
	if file.IsDeprecated() {
		p.PrintLine("//")
		p.PrintLine("// ", file.GetName(), " is a deprecated file.")
	}

	syntax := descriptor.Proto3Syntax
	if !file.IsProto3() {
		syntax = descriptor.Proto3Syntax
	}
	p.PrintBlankLine()
	p.PrintLine("syntax = \"", syntax, "\";")

	if len(file.GetPackageName()) > 0 {
		p.PrintBlankLine().PrintLine("package ", file.GetPackageName(), ";")
	} else {
		// error
	}
}

func (p *Printer) PrintDescriptorFileImports(ctx context.Context, file *descriptor.File) {
	_ = ctx
	dependencies := file.GetDependencies()
	if len(dependencies) > 0 {
		p.PrintBlankLine()
		for _, imp := range dependencies {
			p.PrintLine(`import "`, imp, `";`)
		}
	}
}
