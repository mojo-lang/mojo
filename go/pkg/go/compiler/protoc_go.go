package compiler

import (
	"errors"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"os/exec"
	path2 "path"
	"strings"
)

func ProtocGo(path string, pkg *lang.Package, files []*descriptor.FileDescriptor) (util.CodeGeneratedFiles, error) {
	if pkg == nil {
		return nil, errors.New("")
	}

	cmd := exec.Command("protoc", "-I.")
	for _, dep := range pkg.ResolvedDependencies {
		p := dep.GetExtraString("path")
		cmd.Args = append(cmd.Args, "--proto_path="+path2.Join(p, "protobuf"))
	}

	cmd.Args = append(cmd.Args, "--include_source_info")
	cmd.Args = append(cmd.Args, "--include_imports")

	cmd.Dir = path2.Join(path, "protobuf")

	//cmd.Args = append(cmd.Args, "--go_out=.")
	//cmd.Args = append(cmd.Args, "--gogo_out=plugins=grpc:.")
	cmd.Args = append(cmd.Args, "--descriptor_set_out=/dev/stdout")

	var outFiles util.CodeGeneratedFiles
	for _, file := range files {
		fileCmd := &exec.Cmd{
			Path: cmd.Path,
			Args: cmd.Args,
			Dir:  cmd.Dir,
		}
		fileCmd.Args = append(fileCmd.Args, *file.Name)
		out, err := fileCmd.CombinedOutput()
		if err != nil {
			logs.Errorw("failed to run protoc cmd", "error", string(out), "cmd", cmd.String())
			return nil, err
		}

		fs, err := descriptor.UnmarshalFiles(out)
		if err != nil {
			return nil, err
		}

		//vanity.ForEachFile(fs, vanity.TurnOnMarshalerAll)
		//vanity.ForEachFile(fs, vanity.TurnOnSizerAll)
		//vanity.ForEachFile(fs, vanity.TurnOnUnmarshalerAll)
		//vanity.ForEachFile(fs, vanity.TurnOffGogoImport)
		vanity.ForEachFieldInFiles(fs, descriptor.JsonTagLowerCamelCase)

		req := &plugin.CodeGeneratorRequest{
			ProtoFile: fs,
		}
		for _, f := range fs {
			if *f.Name == *file.Name {
				req.FileToGenerate = append(req.FileToGenerate, *f.Name)
			}
		}
		resp := command.Generate(req)
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		for _, f := range resp.File {
			if f.Name != nil && f.Content != nil {
				goPath := "pkg/" + strings.ReplaceAll(*file.Package, ".", "/")
				pos := strings.Index(*f.Name, goPath)
				if pos >= 0 {
					name := (*f.Name)[pos:]
					outFiles = append(outFiles, &util.CodeGeneratedFile{
						Name:    name,
						Content: *f.Content,
					})
				}
			}
		}
	}
	return outFiles, nil
}
