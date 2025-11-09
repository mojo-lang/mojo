package generator

import (
	"github.com/mojo-lang/mojo/go/pkg/compiler/java/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/java/generator/generator"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GenerateService(services []*data.Service, output string) error {
	d := toData(services)
	return generate(d, output)
}

func toData(services []*data.Service) *data.Data {
	ss := make(map[string][]*data.Service)
	for _, service := range services {
		ss[service.PackageFullName] = append(ss[service.PackageFullName], service)
	}

	d := &data.Data{}
	for _, s := range ss {
		pkg := &data.ServicePackage{
			Services: s,
		}

		if len(s) > 0 {
			pkg.PackageName = s[0].PackageName
			pkg.FullPackageName = s[0].PackageFullName
			pkg.JavaPackageName = s[0].Java.PackageName
			pkg.JavaServicePackageName = s[0].Java.ServicePackageName
			d.Packages = append(d.Packages, pkg)
		}
	}

	return d
}

func generate(d *data.Data, output string) error {
	gen := &generator.Generator{}
	err := gen.Generate(d)
	if err != nil {
		return err
	}

	var files util.GeneratedFiles
	files = append(files, gen.Files...)

	guard := &util.PathGuard{
		DisableClear: true,
		Suffixes:     []string{".java"},
	}

	for _, f := range files {
		if err = f.WriteTo(output, guard); err != nil {
			return err
		}
	}
	return nil
}

func UpdateProtoJavaFiles(dir string) error {
	var files []string
	err := filepath.Walk(dir, func(p string, f os.FileInfo, err error) error {
		if f != nil && !f.IsDir() && strings.HasSuffix(path.Base(f.Name()), "Proto.java") {
			files = append(files, p)
		}
		return nil
	})

	if err != nil {
		logs.Errorw("filepath.Walk() failed", "error", err.Error())
		return err
	}

	for _, file := range files {
		if !util.IsAllGeneratedFile(file) {
			continue
		}

		content, err := os.ReadFile(file)
		newContent := strings.ReplaceAll(string(content), "com.google.protobuf.MojoProtos", "org.mojolang.mojo.MojoProtos")

		if string(content) != newContent {
			if err = os.WriteFile(file, []byte(newContent), fs.ModePerm); err != nil {
				return err
			}
		}
	}

	return nil
}
