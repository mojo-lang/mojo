package generator

import (
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/printer"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

type Generator struct {
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateDescriptorsTo(files []*descriptor.File, out string) error {
	outs, err := g.GenerateDescriptors(files)
	if err != nil {
		return err
	}
	return g.writeGeneratedFiles(outs, out)
}

func (g *Generator) GenerateRequestTo(request *pluginpb.CodeGeneratorRequest, out string) error {
	response, err := g.GenerateRequest(request)
	if err != nil {
		return err
	}

	return g.writeGeneratedFiles(ConvertCodeGeneratorResponse(response), out)
}

func (g *Generator) GeneratePackageTo(pkg *lang.Package, out string) error {
	outs, err := g.GeneratePackage(pkg)
	if err != nil {
		return err
	}
	return g.writeGeneratedFiles(outs, out)
}

func (g *Generator) GenerateRequest(request *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	files := make([]*descriptor.File, 0, len(request.ProtoFile))

	genFileNames := make(map[string]bool)
	for _, n := range request.FileToGenerate {
		genFileNames[n] = true
	}

	for _, f := range request.ProtoFile {
		if _, ok := genFileNames[*f.Name]; ok {
			file := descriptor.NewFileFrom(f)
			files = append(files, file)
		}
	}

	out, err := g.GenerateDescriptors(files)
	if err != nil {
		return nil, err
	}

	return ConvertGeneratedFiles(out), nil
}

func (g *Generator) GenerateDescriptors(files []*descriptor.File) ([]*util.GeneratedFile, error) {
	var out []*util.GeneratedFile
	for _, file := range files {
		p := printer.New(nil)
		p.PrintDescriptorFile(context.Empty(), file)
		if p.Printer.Error != nil {
			return nil, p.Printer.Error
		}

		if content := p.Buffer.String(); len(content) > 0 {
			out = append(out, &util.GeneratedFile{
				Name:                file.GetName(),
				Content:             p.Buffer.String(),
				Reader:              nil,
				SkipIfExist:         false,
				SkipIfUserCodeMixed: false,
			})
		} else {
			logs.Infow("generate an empty file", "name", file.GetName())
		}
	}

	return out, nil
}

func (g *Generator) GeneratePackage(pkg *lang.Package) ([]*util.GeneratedFile, error) {
	var out []*util.GeneratedFile
	for _, file := range pkg.SourceFiles {
		p := printer.New(nil)
		p.PrintSourceFile(context.Empty(), file)
		if p.Printer.Error != nil {
			return nil, p.Printer.Error
		}

		if content := p.Buffer.String(); len(content) > 0 {
			out = append(out, &util.GeneratedFile{
				Name:                file.GetName(),
				Content:             p.Buffer.String(),
				Reader:              nil,
				SkipIfExist:         false,
				SkipIfUserCodeMixed: false,
			})
		} else {
			logs.Infow("generate an empty file", "name", file.GetName())
		}
	}
	return out, nil
}

func (g *Generator) writeGeneratedFiles(files []*util.GeneratedFile, out string) error {
	if len(files) > 0 {
		guard := &util.PathGuard{
			OnlyClearGenerated: true,
			Suffixes:           []string{".proto"},
		}

		for _, file := range files {
			if err := file.WriteTo(out, guard); err != nil {
				return err
			}
		}
	}
	return nil
}
