package java

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"os/exec"
	path2 "path"
)

type Builder struct {
	builder.Builder
	Output string
	Files  []*descriptor.FileDescriptor
}

func (b Builder) protocJava() error {
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to build java.")
		return nil
	}

	if b.Package == nil {
		return errors.New("")
	}

	cmd := exec.Command("protoc", "-I.")
	for _, dep := range b.Package.ResolvedDependencies {
		path := dep.GetExtraString("path")
		cmd.Args = append(cmd.Args, "--proto_path="+path2.Join(path, "protobuf"))
	}

	cmd.Dir = path2.Join(b.PWD, b.Path, "protobuf")

	//cmd.Args = append(cmd.Args, "--go_out=.")
	cmd.Args = append(cmd.Args, "--java_out=.")
	for _, file := range b.Files {
		fileCmd := &exec.Cmd{
			Path: cmd.Path,
			Args: cmd.Args,
			Dir:  cmd.Dir,
		}
		fileCmd.Args = append(fileCmd.Args, *file.Name)
		out, err := fileCmd.CombinedOutput()
		if err != nil {
			logs.Errorw("failed to run protoc cmd", "error", string(out), "cmd", cmd.String())
			return err
		}
	}

	// move the generated files to destinations
	destDir := path2.Join(b.PWD, b.Path, "java/src/main/java")
	if !util.IsExist(destDir) {
		util.CreateDir(destDir)
	}

	util.ClearFiles(destDir, ".pb.java")

	for _, domain := range []string{"ai", "com", "net", "org", "io"} {
		srcDir := path2.Join(b.PWD, b.Path, "protobuf", domain)
		_, err := ioutil.ReadDir(srcDir)
		if err == nil {
			err = copy.Copy(srcDir, path2.Join(destDir, domain))
			if err != nil {
				return err
			}

			os.RemoveAll(srcDir)
			break
		}
	}

	return nil
}

func (b Builder) Build() error {
	// protoc the protobuf files to golang files
	if err := b.protocJava(); err != nil {
		return err
	}

	return nil
}
