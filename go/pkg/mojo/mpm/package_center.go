package mpm

import (
    "os"
    "os/exec"
    "path"

    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "google.golang.org/protobuf/proto"
)

var packageCenter *PackageCenter

func GetPackageCenter() *PackageCenter {
    center := &PackageCenter{
        Cache: make(map[string]*lang.Package),
    }

    center.MojoHome = os.Getenv("MOJO_HOME")
    if len(center.MojoHome) == 0 {
        home, _ := os.UserHomeDir()
        center.MojoHome = path.Join(home, "mojo")
    }
    center.MojoPkgRoot = path.Join(center.MojoHome, "pkg")

    return center
}

type PackageCenter struct {
    MojoHome    string
    MojoPkgRoot string

    Cache map[string]*lang.Package
}

func (p *PackageCenter) Get(name string, requirement *lang.Package_Requirement) (string, error) {
    if pkg, ok := p.Cache[name]; ok {
        return pkg.ExtraInfo.GetString("path"), nil
    }
    repoPath := p.getPkgPath(requirement)
    if core.IsExist(repoPath) {
        return p.Update(name, requirement)
    }

    return p.Install(name, requirement)
}

func (p *PackageCenter) Install(name string, requirement *lang.Package_Requirement) (string, error) {
    url := proto.Clone(requirement.Repository).(*core.Url)
    if len(url.Scheme) == 0 {
        url.Scheme = "https"
    }
    repoPath := p.getPkgPath(requirement)

    cmd := exec.Command("git", "clone", url.Format())
    cmd.Dir = path.Dir(repoPath)

    core.CreateDir(cmd.Dir)

    logs.Debugw("begin to install mojo package", "package", name, "cmd", cmd.String())
    out, err := cmd.CombinedOutput()
    if err != nil {
        logs.Errorw("failed to run git cmd", "error", string(out), "cmd", cmd.String())
        return "", err
    }
    logs.Debugw("finish to install mojo package", "package", name, "cmd", cmd.String())

    p.Cache[name] = &lang.Package{
        Name:       lang.GetPackageName(name),
        FullName:   name,
        Version:    nil,
        Repository: requirement.Repository,
        ExtraInfo:  pathObject(repoPath),
    }

    return repoPath, nil
}

func (p *PackageCenter) Update(name string, requirement *lang.Package_Requirement) (string, error) {
    repoPath := p.getPkgPath(requirement)

    //cmd := exec.Command("git", "pull")
    //cmd.Dir = repoPath
    //
    //logs.Debugw("begin to update mojo package", "package", name, "cmd", cmd.String())
    //out, err := cmd.CombinedOutput()
    //if err != nil {
    //	logs.Errorw("failed to run git cmd", "error", string(out), "cmd", cmd.String())
    //	return "", err
    //}
    //logs.Debugw("finish to update mojo package", "package", name, "cmd", cmd.String())

    p.Cache[name] = &lang.Package{
        Name:       lang.GetPackageName(name),
        FullName:   name,
        Repository: requirement.Repository,
        ExtraInfo:  pathObject(repoPath),
    }

    return repoPath, nil
}

func (p *PackageCenter) getPkgPath(requirement *lang.Package_Requirement) string {
    return path.Join(p.MojoPkgRoot, requirement.Repository.FormatWithoutSchema())
}

func pathObject(path string) *core.Object {
    object := &core.Object{}
    object.SetString("path", path)
    return object
}
