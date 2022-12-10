package mpm

import (
	"os"
	"os/exec"
	"path"
	"strings"

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

	Cache    map[string]*lang.Package
	MojoPkgs map[string]*lang.Package
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

func GetGitLatestCommit(dir string) *lang.Package_Requirement_Commit {
	cmd := exec.Command("git", "log", "-n", "1", `--format=format:"%H %cI"`)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		logs.Errorw("failed to run git cmd", "error", string(out), "cmd", cmd.String())
		return nil
	}
	logs.Debugw("finish to get git log", "workDir", dir, "cmd", cmd.String())
	segments := strings.Split(core.RemoveQuote(string(out), `"`), " ")
	if len(segments) == 2 {
		commit := &lang.Package_Requirement_Commit{
			Hash: segments[0],
		}
		t, err := core.ParseTimestamp(segments[1])
		if err != nil {
			return nil
		}
		commit.Date = t
		return commit
	}
	return nil
}

func (p *PackageCenter) Install(name string, requirement *lang.Package_Requirement) (string, error) {
	url := proto.Clone(requirement.Repository).(*core.Url)
	if len(url.Scheme) == 0 {
		url.Scheme = "https"
	}
	repoPath := p.getPkgPath(requirement)

	cmd := exec.Command("git", "clone", url.Format())
	cmd.Dir = path.Dir(repoPath)

	if err := core.CreateDir(cmd.Dir); err != nil {
		return "", err
	}

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

	commit := GetGitLatestCommit(repoPath)
	if commit != nil && requirement.Commit != nil && len(requirement.Commit.Hash) > 0 {
		if strings.HasPrefix(commit.Hash, requirement.Commit.Hash) {
			logs.Debugw("the mojo package are already latest", "package", name)
			return repoPath, nil
		}
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = repoPath

	logs.Debugw("begin to update mojo package", "package", name, "cmd", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		logs.Errorw("failed to run git cmd", "error", string(out), "cmd", cmd.String())
		return "", err
	}
	logs.Debugw("finish to update mojo package", "package", name, "cmd", cmd.String())

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
