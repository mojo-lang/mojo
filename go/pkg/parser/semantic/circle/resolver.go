package circle

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/plugin"
	"strings"
)

func init() {
	plugin.AddPlugin("circle-resolver", 5, &Resolver{})
}

const DependencyCircle = "dependency_circle"

type Resolver struct {
}

func (r *Resolver) Parse(ctx *plugin.Context, pkg *lang.Package, options map[string]interface{}) error {
	ctx.Open(pkg)
	if len(options) > 0 {
		ctx.SetOptions(options)
	}

	defer func() {
		ctx.Close()
	}()

	for _, child := range pkg.Children {
		if err := r.Parse(ctx, child, nil); err != nil {
			return err
		}
	}

	if err := r.parsePackage(ctx, pkg); err != nil {
		return err
	}

	return nil
}

func removeDuplicated(values []string) []string {
	index := make(map[string]bool)
	for _, value := range values {
		index[value] = true
	}

	values = []string{}
	for value := range index {
		values = append(values, value)
	}
	return values
}

func (r *Resolver) parsePackage(ctx *plugin.Context, pkg *lang.Package) error {
	// generated all the file node
	_ = ctx
	files := make(map[string]*fileNode)
	for _, file := range pkg.SourceFiles {
		node := &fileNode{
			name: file.FullName,
		}

		for _, dependency := range file.ResolvedIdentifiers {
			if dependency.PackageName == pkg.FullName {
				node.dependencies = append(node.dependencies, dependency.SourceFileName)
			}
		}
		node.dependencies = removeDuplicated(node.dependencies)

		files[node.name] = node
	}

	circles := NewSearcher().Search(files)
	if len(circles) == 0 {
		return nil
	}

	circleIndex := make(map[string]bool)
	for _, circle := range circles {
		for _, name := range circle {
			circleIndex[name] = true
		}
	}

	for _, file := range pkg.SourceFiles {
		if circleIndex[file.FullName] {
			circleName := r.generateCircleName(ctx, pkg)
			file.Attributes = lang.SetStringAttribute(file.Attributes, DependencyCircle, circleName)
		}
	}

	return nil
}

func (r *Resolver) generateCircleName(ctx *plugin.Context, pkg *lang.Package) string {
	_ = ctx
	return strings.ReplaceAll(pkg.FullName, ".", "/") + "/" + pkg.Name + ".mojo"
}
