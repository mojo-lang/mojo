package compiler

import (
	"bytes"
	"errors"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/grpc"
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

func ProtocGo(path string, pkg *lang.Package, files []*descriptor.FileDescriptor) (util.GeneratedFiles, error) {
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

	var outFiles util.GeneratedFiles

	/// gogoproto will clean the plugin after the first run `command.Generate(req)`
	/// make sure the service files will use the grpc plugin
	fileGroups := make([][]*descriptor.FileDescriptor, 2, 2)
	for _, file := range files {
		if file.HasService() {
			fileGroups[0] = append(fileGroups[0], file)
		} else {
			fileGroups[1] = append(fileGroups[1], file)
		}
	}

	for _, group := range fileGroups {
		if len(group) == 0 {
			continue
		}

		fileCmd := &exec.Cmd{
			Path: cmd.Path,
			Args: cmd.Args,
			Dir:  cmd.Dir,
		}

		for _, file := range group {
			fileCmd.Args = append(fileCmd.Args, *file.Name)
		}

		var stderr bytes.Buffer
		fileCmd.Stderr = &stderr
		out, err := fileCmd.Output()
		if err != nil {
			logs.Errorw("failed to run protoc cmd", "error", stderr.String(), "cmd", fileCmd.String())
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

		parameter := "plugins=grpc"
		req := &plugin.CodeGeneratorRequest{
			ProtoFile: fs,
			Parameter: &parameter,
		}
		for _, f := range fs {
			for _, file := range group {
				if *f.Name == *file.Name {
					req.FileToGenerate = append(req.FileToGenerate, *f.Name)
				}
			}
		}
		resp := command.Generate(req)
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		for _, f := range resp.File {
			if f.Name != nil && f.Content != nil {
				for _, file := range group {
					goPath := "pkg/" + strings.ReplaceAll(*file.Package, ".", "/")
					pos := strings.Index(*f.Name, goPath)
					if pos >= 0 {
						name := (*f.Name)[pos:]
						outFiles = append(outFiles, &util.GeneratedFile{
							Name:    name,
							Content: *f.Content,
						})
					}
				}
			}
		}
	}
	return outFiles, nil
}
