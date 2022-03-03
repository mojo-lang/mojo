package compiler

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic/circle"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	path2 "path"
)

func init() {
	plugin.RegisterPlugin(NewCircleCompiler(nil))
}

const circleName = "compiler.circle"

type CircleCompiler struct {
	plugin.BasicPlugin
}

func NewCircleCompiler(options core.Options) *CircleCompiler {
	return &CircleCompiler{
		BasicPlugin: plugin.BasicPlugin{
			Name:          circleName,
			Group:         "compiler",
			GroupPriority: 10,
			Priority:      9,
			Creator: func(options core.Options) plugin.Plugin {
				return NewCircleCompiler(options)
			},
		},
	}
}

func (c *CircleCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName)

	polluted, err := c.compile(ctx, pkg)
	if err != nil {
		return err
	}

	const MaxRetry = 10
	retry := 0
	for polluted && retry < MaxRetry {
		polluted, err = c.compile(ctx, pkg)
		if err != nil {
			return err
		}
		retry++
	}

	if polluted {
		return errors.New("Can't clean the circle dependency in the " + pkg.FullName)
	}

	return nil
}

func (c *CircleCompiler) reParse(ctx context.Context, pkg *lang.Package) error {
	_ = ctx

	for _, f := range pkg.SourceFiles {
		f.Attributes = lang.RemoveAttribute(f.Attributes, circle.DependencyCircle)
	}

	resolver := circle.NewResolver(nil)
	resolver.Force = true
	return resolver.ParsePackage(context.Empty(), pkg)
}

func (c *CircleCompiler) compile(ctx context.Context, pkg *lang.Package) (polluted bool, err error) {
	if err = c.reParse(ctx, pkg); err != nil {
		return false, err
	}

	circles := make(map[string][]*lang.SourceFile)
	var files []*lang.SourceFile

	for _, sourceFile := range pkg.SourceFiles {
		if file, err := lang.GetStringAttribute(sourceFile.Attributes, circle.DependencyCircle); err == nil {
			for _, identifier := range sourceFile.Scope.Identifiers {
				identifier.SourceFileName = file
			}
			sourceFile.Attributes = lang.RemoveAttribute(sourceFile.Attributes, circle.DependencyCircle)

			circles[file] = append(circles[file], sourceFile)
		} else {
			files = append(files, sourceFile)
		}
	}

	getExistFile := func(fileName string) *lang.SourceFile {
		for _, file := range files {
			if file.FullName == fileName {
				return file
			}
		}
		return nil
	}

	polluted = false
	for fileName, circleFiles := range circles {
		existFile := true
		sourceFile := getExistFile(fileName)
		if sourceFile == nil {
			existFile = false
			sourceFile = &lang.SourceFile{
				Name:        path2.Base(fileName),
				FullName:    fileName,
				PackageName: pkg.FullName,
				Scope: &lang.Scope{
					Identifiers: make(map[string]*lang.Identifier),
				},
			}
		} else {
			polluted = true
		}

		// update all the identifiers
		for _, file := range circleFiles {
			for k, v := range file.Scope.Identifiers {
				sourceFile.Scope.Identifiers[k] = v
			}
		}

		for _, file := range circleFiles {
			sourceFile.ResolvedIdentifiers = append(sourceFile.ResolvedIdentifiers, file.ResolvedIdentifiers...)
			sourceFile.UnresolvedIdentifiers = append(sourceFile.UnresolvedIdentifiers, file.UnresolvedIdentifiers...)
			sourceFile.Attributes = append(sourceFile.Attributes, file.Attributes...)
			sourceFile.Statements = append(sourceFile.Statements, file.Statements...)
		}

		var resolvedIdentifiers []*lang.Identifier
		for _, identifier := range sourceFile.ResolvedIdentifiers {
			if sourceFile.Scope.Identifiers[identifier.FullName] != nil {
				continue
			}
			resolvedIdentifiers = append(resolvedIdentifiers, identifier)
		}
		sourceFile.ResolvedIdentifiers = lang.MergeDependencies(resolvedIdentifiers)

		if !existFile {
			files = append(files, sourceFile)
		}
	}

	pkg.SourceFiles = files

	return polluted, nil
}
