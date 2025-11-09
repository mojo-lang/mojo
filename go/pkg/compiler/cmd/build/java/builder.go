package java

import (
	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/java/generator"
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"
	"github.com/otiai10/copy"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	gogen "github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/generator"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

type Builder struct {
	builder.Builder
	Output string
	Files  []*descriptor.File
}

func (b Builder) protocJava() error {
	if b.Package == nil {
		return errors.New("")
	}

	var tempPbDirs []string
	cmd := exec.Command("protoc", "-I.")
	for _, dep := range b.Package.ResolvedDependencies {
		wd := dep.GetExtraString("workingDir")
		p := dep.GetExtraString("path")
		if len(wd) == 0 && len(p) == 0 {
			pbDir, err := gogen.GenerateMojoPackageProtobuf(dep)
			if err != nil {
				logs.Errorw("failed to generate mojo package's protobuf files", "package", dep.FullName, "error", err)
				return err
			}
			if len(pbDir) > 0 {
				tempPbDirs = append(tempPbDirs, pbDir)
				cmd.Args = append(cmd.Args, "--proto_path="+pbDir)
			}
		} else {
			cmd.Args = append(cmd.Args, "--proto_path="+path.Join(wd, p, "protobuf"))
		}
	}

	cmd.Dir = path.Join(b.GetAbsolutePath(), "protobuf")

	//$ protoc --plugin=protoc-gen-grpc-java \
	//--grpc-java_out="$OUTPUT_FILE" --proto_path="$DIR_OF_PROTO_FILE" "$PROTO_FILE"

	// cmd.Args = append(cmd.Args, "--go_out=.")
	cmd.Args = append(cmd.Args, "--java_out=.")
	cmd.Args = append(cmd.Args, "--grpc-java_out=.")
	for _, file := range b.Files {
		if file.IsEmpty() {
			continue
		}

		fileCmd := &exec.Cmd{
			Path: cmd.Path,
			Args: cmd.Args,
			Dir:  cmd.Dir,
		}
		fileCmd.Args = append(fileCmd.Args, file.GetName())
		out, err := fileCmd.CombinedOutput()
		if err != nil {
			logs.Errorw("failed to run protoc cmd", "error", string(out), "cmd", cmd.String())
			return err
		}
	}
	defer func() {
		for _, d := range tempPbDirs {
			if err := os.RemoveAll(d); err != nil {
				logs.Warnw("failed to remove all the template protobuf files", "error", err)
			}
		}
	}()

	// move the generated files to destinations
	destDir := path.Join(b.GetAbsolutePath(), "java/src/main/java")
	if !core.IsExist(destDir) {
		if err := core.CreateDir(destDir); err != nil {
			return err
		}
	}

	if err := util.DeepClearGeneratedFiles(destDir, ".java"); err != nil {
		return err
	}

	for _, domain := range []string{"ai", "biz", "cn", "com", "edu", "gov", "net", "org", "info", "io", "tech"} {
		srcDir := path.Join(b.GetAbsolutePath(), "protobuf", domain)
		_, err := os.ReadDir(srcDir)
		if err == nil {
			if err = copy.Copy(srcDir, path.Join(destDir, domain)); err != nil {
				return err
			}
			if err = os.RemoveAll(srcDir); err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func (b Builder) build() error {
	logs.Infow("java begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

	cmp := generator.NewCompiler()
	options := make(core.Options)

	err := cmp.CompilePackage(context.WithOptions(context.Empty(), options), b.Package)
	if err != nil {
		logs.Errorw("failed to compile java", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	services := cmp.Services
	err = generator.GenerateService(services, b.Output)
	if err != nil {
		logs.Errorw("generate java failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return err
	}

	return nil
}

func (b Builder) Build() error {
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to build java.")
		return nil
	}
	logs.Infow("begin to build file.", "pwd", b.PWD, "path", b.Path)

	if len(b.Output) == 0 {
		b.Output = util.GetAbsolutePath(b.PWD, b.Path)
	}
	b.Output = path.Join(b.Output, "java")

	// protoc the protobuf files to java files
	if err := b.protocJava(); err != nil {
		return err
	}

	// other java files builder
	if err := b.build(); err != nil {
		return err
	}

	if err := generator.UpdateProtoJavaFiles(b.Output); err != nil {
		return err
	}

	return nil
}
