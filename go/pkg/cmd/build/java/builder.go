package java

import (
    "errors"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
    "github.com/otiai10/copy"
    "io/ioutil"
    "os"
    "os/exec"
    path2 "path"
)

type Builder struct {
    builder.Builder
    Output string
    Files  []*descriptor.File
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
        wd := dep.GetExtraString("workingDir")
        path := dep.GetExtraString("path")
        cmd.Args = append(cmd.Args, "--proto_path="+path2.Join(wd, path, "protobuf"))
    }

    cmd.Dir = path2.Join(b.GetAbsolutePath(), "protobuf")

    //cmd.Args = append(cmd.Args, "--go_out=.")
    cmd.Args = append(cmd.Args, "--java_out=.")
    for _, file := range b.Files {
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

    // move the generated files to destinations
    destDir := path2.Join(b.GetAbsolutePath(), "java/src/main/java")
    if !core.IsExist(destDir) {
        core.CreateDir(destDir)
    }

    util.DeepClearFiles(destDir, ".pb.java")

    for _, domain := range []string{"ai", "biz", "cn", "com", "edu", "gov", "net", "org", "info", "io", "tech"} {
        srcDir := path2.Join(b.GetAbsolutePath(), "protobuf", domain)
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
