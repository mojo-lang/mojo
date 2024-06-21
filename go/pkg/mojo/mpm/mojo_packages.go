package mpm

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/edwin-luijten/go_mod_parser"
	"github.com/edwin-luijten/go_mod_parser/module"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/geom/go/pkg/mojo/geom"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/rpc/go/pkg/mojo/rpc"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

// force including the rpc, geom package in go.mod for generating all the mojo packages
var _ = rpc.Message{}
var _ = geom.LngLat{}

//go:embed mojo/*
var packages embed.FS

var mojoPackages map[string]*lang.Package
var mojoPackagesOnce sync.Once

func GetMojoPackages() map[string]*lang.Package {
	mojoPackagesOnce.Do(func() {
		pkgs := make(map[string]*lang.Package)
		err := fs.WalkDir(packages, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".pb.binary") {
				return nil
			}

			f, err := packages.Open(path)
			if err != nil {
				return err
			}
			b, err := io.ReadAll(f)
			if err != nil {
				return err
			}

			pkg := &lang.Package{}
			if err = proto.Unmarshal(b, pkg); err != nil {
				return err
			}

			pkgs[pkg.FullName] = pkg
			return nil
		})

		if err != nil {
			logs.Errorw("failed to get embed mojo packages", "error", err)
			return
		}

		for _, pkg := range pkgs {
			if len(pkg.Dependencies) > 0 {
				if pkg.ResolvedDependencies == nil {
					pkg.ResolvedDependencies = make(map[string]*lang.Package)
				}
				for name := range pkg.Dependencies {
					pkg.ResolvedDependencies[name] = pkgs[name]
				}
			}
		}
		mojoPackages = pkgs
	})

	return mojoPackages
}

func GetMojoPackage(name string) *lang.Package {
	for n, pkg := range GetMojoPackages() {
		if n == name {
			return pkg
		}
	}
	return nil
}

func GetMojoPbFile(name string) *BinaryFile {
	var target *BinaryFile
	binaryFileName := "mojo/" + name + ".pb.binary"
	err := fs.WalkDir(packages, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if path == binaryFileName {
			f, e := packages.Open(path)
			if e != nil {
				return e
			}
			b, e := io.ReadAll(f)
			if e != nil {
				return e
			}

			file := &BinaryFile{Parts: make(map[string][]byte)}
			if e = file.Decode(b); e != nil {
				return e
			}

			target = file
			return core.NewBreakError()
		}
		return nil
	})
	if err != nil && !core.IsBreakError(err) {
		return nil
	}
	return target
}

// GenerateMojoPackages
// 1. parse the mojo go.mod
// 2. clone or update the git repo for all mojo packages
// 3. parse the core package using ony syntax parser
// 4. then other package one by one
func GenerateMojoPackages(projectPath string) error {
	pkgs := parseGoMod(projectPath)
	for _, name := range []string{"core", "document", "lang", "openapi", "http", "db", "rpc", "geom"} {
		if err := generatePackage(name, projectPath, pkgs); err != nil {
			return err
		}
	}
	return nil
}

type mojoPackage struct {
	Name       string
	Path       string
	Replaced   *mojoPackage
	RawVersion string
	Version    *core.Version
	Timestamp  string
	Hash       string
}

func parseGoMod(projectPath string) map[string]*mojoPackage {
	contents, _ := os.ReadFile(path.Join(projectPath, "go.mod"))
	if len(contents) == 0 {
		contents, _ = os.ReadFile(path.Join(projectPath, "go", "go.mod"))
	}

	p, err := go_mod_parser.Parse("go.mod", contents, nil)
	if err != nil {
		return nil
	}

	parse := func(v *module.Version) *mojoPackage {
		if !strings.HasPrefix(v.Path, "github.com/mojo-lang") {
			logs.Warnw("not mojo std package", "path", v.Path)
			return nil
		}

		pkg := &mojoPackage{
			Path:       strings.TrimSuffix(v.Path, "/go"),
			RawVersion: v.Version,
		}
		pkg.Name = path.Base(pkg.Path)

		version := v.Version
		segments := strings.Split(version, "-")
		if len(segments) > 0 {
			if err = pkg.Version.Parse(segments[0][1:]); err != nil {
				logs.Warnw("failed to parse version", "path", v.Path, "version", version, "error", err)
				return nil
			}
		}
		if len(segments) > 2 {
			pkg.Timestamp = segments[1]
			pkg.Hash = segments[2]
		}

		return pkg
	}

	pkgs := make(map[string]*mojoPackage)
	for _, required := range p.Require {
		pkg := parse(&required.Mod)
		if pkg == nil {
			continue
		}

		for _, replace := range p.Replace {
			if replace.Old.Path == pkg.Name {
				pkg.Replaced = parse(&replace.New)
			}
		}
		pkgs[pkg.Name] = pkg
	}
	return pkgs
}

func compileMojoPackage(pkg *mojoPackage) (*lang.Package, *BinaryFile, error) {
	url, err := core.NewUrl(pkg.Path)
	if err != nil {
		return nil, nil, err
	}
	requirement := &lang.Package_Requirement{
		Repository: url,
		Commit: &lang.Package_Requirement_Commit{
			Hash: pkg.Hash,
			Date: nil,
		},
	}
	if len(pkg.Timestamp) > 0 {
		ct, err := core.ParseTimestamp(pkg.Timestamp)
		if err != nil {
			return nil, nil, err
		}
		requirement.Commit.Date = ct
	}
	dir, err := GetPackageCenter().Get("mojo."+pkg.Name, requirement)
	if err != nil {
		return nil, nil, err
	}

	plugins := plugin.NewPlugins("mpm", "syntax")
	p, err := plugins.ParsePath(plugin.WithWorkingDir(context.Empty(), dir), dir)
	if err != nil {
		return nil, nil, err
	}

	var pbFile *BinaryFile
	pbDir := path.Join(dir, "protobuf")
	if core.IsExist(pbDir) {
		pbFile = &BinaryFile{Parts: make(map[string][]byte)}
		err = filepath.WalkDir(path.Join(dir, "protobuf"), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".proto") {
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			name, err := filepath.Rel(pbDir, path)
			if err != nil {
				return err
			}
			pbFile.Parts[name] = content
			return nil
		})
		if err != nil {
			return nil, nil, err
		}
	}

	return p, pbFile, nil
}

func shrinkPackage(pkg *lang.Package) {
	pkg.ExtraInfo.Delete("workingDir").Delete("path").Delete(pluginName)
	pkg.ResolvedDependencies = nil

	for _, p := range pkg.Children {
		shrinkPackage(p)
	}
}

func savePackage(projectPath string, pkg *lang.Package, pb *BinaryFile) error {
	if pkg != nil {
		shrinkPackage(pkg)

		bytes, err := proto.Marshal(pkg)
		if err != nil {
			return err
		}
		fileName := path.Join(projectPath, "pkg/mojo/mpm/mojo", pkg.Name+".binary")

		err = os.WriteFile(fileName, bytes, fs.ModePerm)
		if err != nil {
			return err
		}
		logs.Infow("save the pkg to binary file", "pkg", pkg.FullName, "size", len(bytes), "file", fileName)
	}
	if pb != nil {
		fileName := path.Join(projectPath, "pkg/mojo/mpm/mojo", pkg.Name+".pb.binary")
		bytes := pb.Encode()
		err := os.WriteFile(fileName, bytes, fs.ModePerm)
		if err != nil {
			return err
		}
		logs.Infow("save the pb files to binary file", "pkg", pkg.FullName, "size", len(bytes), "file", fileName)
	}

	return nil
}

func generatePackage(name string, projectPath string, pkgs map[string]*mojoPackage) error {
	if p, ok := pkgs[name]; ok {
		pkg, pb, err := compileMojoPackage(p)
		if err != nil {
			return err
		}
		if err = savePackage(projectPath, pkg, pb); err != nil {
			return err
		}
	}
	return nil
}

type BinaryFile struct {
	Parts map[string][]byte
}

func (f *BinaryFile) Encode() []byte {
	if f == nil || len(f.Parts) == 0 {
		return nil
	}

	var names []string
	for k := range f.Parts {
		names = append(names, k)
	}
	sort.Strings(names)

	var binary []byte
	for _, k := range names {
		var buf []byte
		v := f.Parts[k]
		buf = protowire.AppendTag(buf, 1, protowire.BytesType)
		buf = protowire.AppendString(buf, k)
		buf = protowire.AppendTag(buf, 2, protowire.BytesType)
		buf = protowire.AppendBytes(buf, v)

		binary = protowire.AppendTag(binary, 1, protowire.BytesType)
		binary = protowire.AppendBytes(binary, buf)
	}
	return binary
}

func consumeBytes(bytes []byte, number protowire.Number) (left []byte, value []byte, err error) {
	num, typ, length := protowire.ConsumeTag(bytes)
	if length < 0 {
		err = protowire.ParseError(length)
		return
	}
	if num != number || typ != protowire.BytesType {
		err = errors.New("invalid binary file format")
		return
	}
	if length == 0 {
		return
	}
	bytes = bytes[length:]

	value, length = protowire.ConsumeBytes(bytes)
	if length < 0 {
		err = protowire.ParseError(length)
		return
	}
	if length == 0 {
		return
	}
	left = bytes[length:]
	return
}

func DecodeBinaryFile(binary []byte) (*BinaryFile, error) {
	file := &BinaryFile{}
	if err := file.Decode(binary); err != nil {
		return nil, err
	}
	return file, nil
}

func (f *BinaryFile) Decode(binary []byte) error {
	if f != nil {
		if f.Parts == nil {
			f.Parts = make(map[string][]byte)
		}
		for len(binary) > 0 {
			left, value, err := consumeBytes(binary, 1)
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return nil
			}
			binary = left

			left, value, err = consumeBytes(value, 1)
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return errors.New("invalid key in the binary file")
			}
			key := string(value)

			left, value, err = consumeBytes(left, 2)
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return fmt.Errorf("invalid value of %s key in the binary file", key)
			}
			f.Parts[key] = value
		}
	}
	return nil
}
