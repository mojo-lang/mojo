package plugin

import (
	"errors"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

// lifecyclePlugin records the externally observable hook and compiler order.
type lifecyclePlugin struct {
	events             []string
	preErr, compileErr error
}

func (p *lifecyclePlugin) record(event string, err error) error {
	p.events = append(p.events, event)
	return err
}

func (p *lifecyclePlugin) PrePackage(context.Context, *lang.Package) error {
	return p.record("pre", p.preErr)
}
func (p *lifecyclePlugin) PostPackage(context.Context, *lang.Package) {
	p.events = append(p.events, "post")
}
func (p *lifecyclePlugin) CompilePackage(context.Context, *lang.Package) error {
	return p.record("compile", p.compileErr)
}

func (p *lifecyclePlugin) PreSourceFile(context.Context, *lang.SourceFile) error {
	return p.record("pre", p.preErr)
}
func (p *lifecyclePlugin) PostSourceFile(context.Context, *lang.SourceFile) {
	p.events = append(p.events, "post")
}
func (p *lifecyclePlugin) CompileSourceFile(context.Context, *lang.SourceFile) error {
	return p.record("compile", p.compileErr)
}

func (p *lifecyclePlugin) PreStruct(context.Context, *lang.StructDecl) error {
	return p.record("pre", p.preErr)
}
func (p *lifecyclePlugin) PostStruct(context.Context, *lang.StructDecl) {
	p.events = append(p.events, "post")
}
func (p *lifecyclePlugin) CompileStruct(context.Context, *lang.StructDecl) error {
	return p.record("compile", p.compileErr)
}

func (p *lifecyclePlugin) PreInterface(context.Context, *lang.InterfaceDecl) error {
	return p.record("pre", p.preErr)
}
func (p *lifecyclePlugin) PostInterface(context.Context, *lang.InterfaceDecl) {
	p.events = append(p.events, "post")
}
func (p *lifecyclePlugin) CompileInterface(context.Context, *lang.InterfaceDecl) error {
	return p.record("compile", p.compileErr)
}

func (p *lifecyclePlugin) PreEnum(context.Context, *lang.EnumDecl) error {
	return p.record("pre", p.preErr)
}
func (p *lifecyclePlugin) PostEnum(context.Context, *lang.EnumDecl) {
	p.events = append(p.events, "post")
}
func (p *lifecyclePlugin) CompileEnum(context.Context, *lang.EnumDecl) error {
	return p.record("compile", p.compileErr)
}

func (p *lifecyclePlugin) PreTypeAlias(context.Context, *lang.TypeAliasDecl) error {
	return p.record("pre", p.preErr)
}
func (p *lifecyclePlugin) PostTypeAlias(context.Context, *lang.TypeAliasDecl) {
	p.events = append(p.events, "post")
}
func (p *lifecyclePlugin) CompileTypeAlias(context.Context, *lang.TypeAliasDecl) error {
	return p.record("compile", p.compileErr)
}

func TestCompilerLifecycle(t *testing.T) {
	sentinel := errors.New("plugin failure")
	for _, dispatch := range []struct {
		name string
		call func(interface{}) error
	}{
		{"Package", func(p interface{}) error { return CompilePackage(p, context.Empty(), &lang.Package{}) }},
		{"SourceFile", func(p interface{}) error { return CompileSourceFile(p, context.Empty(), &lang.SourceFile{}) }},
		{"Struct", func(p interface{}) error { return CompileStruct(p, context.Empty(), &lang.StructDecl{}) }},
		{"Interface", func(p interface{}) error { return CompileInterface(p, context.Empty(), &lang.InterfaceDecl{}) }},
		{"Enum", func(p interface{}) error { return CompileEnum(p, context.Empty(), &lang.EnumDecl{}) }},
		{"TypeAlias", func(p interface{}) error { return CompileTypeAlias(p, context.Empty(), &lang.TypeAliasDecl{}) }},
	} {
		t.Run(dispatch.name, func(t *testing.T) {
			require.NoError(t, dispatch.call(nil))
			require.True(t, core.IsSkipError(dispatch.call(struct{}{})))
			for _, tt := range []struct {
				name                  string
				pre, compile, wantErr error
				events                []string
			}{
				{"success", nil, nil, nil, []string{"pre", "compile", "post"}},
				{"skip", core.NewSkipError(), nil, nil, []string{"pre"}},
				{"pre error", sentinel, nil, sentinel, []string{"pre"}},
				{"compile error", nil, sentinel, sentinel, []string{"pre", "compile", "post"}},
			} {
				t.Run(tt.name, func(t *testing.T) {
					p := &lifecyclePlugin{preErr: tt.pre, compileErr: tt.compile}
					require.ErrorIs(t, dispatch.call(p), tt.wantErr)
					require.Equal(t, tt.events, p.events)
				})
			}
		})
	}
	require.True(t, IsCompiler(&lifecyclePlugin{}))
	require.False(t, IsCompiler(struct{}{}))
	require.False(t, IsCompiler(nil))
}

type structCompilerFunc func(context.Context, *lang.StructDecl) error

func (f structCompilerFunc) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	return f(ctx, decl)
}

func TestCompilePackageTraversesSupportedDeclarations(t *testing.T) {
	file := &lang.SourceFile{Name: "root.mojo", Statements: []*lang.Statement{
		lang.NewEnumDeclStatement(&lang.EnumDecl{Name: "Status"}),
		lang.NewInterfaceDeclStatement(&lang.InterfaceDecl{Name: "Service"}),
		lang.NewTypeAliasDeclStatement(&lang.TypeAliasDecl{Name: "ID"}),
		lang.NewStructDeclStatement(&lang.StructDecl{Name: "First"}),
		lang.NewStructDeclStatement(&lang.StructDecl{Name: "Skipped"}),
		lang.NewStructDeclStatement(&lang.StructDecl{Name: "Last"}),
	}}
	childFile := &lang.SourceFile{Name: "child.mojo", Statements: []*lang.Statement{
		lang.NewStructDeclStatement(&lang.StructDecl{Name: "Child"}),
	}}
	child := &lang.Package{FullName: "root.child", SourceFiles: []*lang.SourceFile{childFile}}
	pkg := &lang.Package{FullName: "root", SourceFiles: []*lang.SourceFile{file}, Children: []*lang.Package{child}}
	var visits []string
	compiler := structCompilerFunc(func(ctx context.Context, decl *lang.StructDecl) error {
		visits = append(visits, decl.Name)
		if decl.Name == "Child" {
			require.Same(t, child, context.Package(ctx))
			require.Same(t, childFile, context.SourceFile(ctx))
		} else {
			require.Same(t, pkg, context.Package(ctx))
			require.Same(t, file, context.SourceFile(ctx))
		}
		if decl.Name == "Skipped" {
			return core.NewSkipError()
		}
		return nil
	})
	require.NoError(t, CompilePackage(compiler, context.Empty(), pkg))
	require.Equal(t, []string{"First", "Skipped", "Last", "Child"}, visits)
}

func TestCompilePackageStopsOnDeclarationError(t *testing.T) {
	sentinel := errors.New("invalid struct")
	pkg := &lang.Package{SourceFiles: []*lang.SourceFile{
		{Statements: []*lang.Statement{
			lang.NewStructDeclStatement(&lang.StructDecl{Name: "Bad"}),
			lang.NewStructDeclStatement(&lang.StructDecl{Name: "Unvisited"}),
		}},
	}, Children: []*lang.Package{{SourceFiles: []*lang.SourceFile{
		{Statements: []*lang.Statement{lang.NewStructDeclStatement(&lang.StructDecl{Name: "Child"})}},
	}}}}
	var visits []string
	err := CompilePackage(structCompilerFunc(func(ctx context.Context, decl *lang.StructDecl) error {
		visits = append(visits, decl.Name)
		return sentinel
	}), context.Empty(), pkg)
	require.ErrorIs(t, err, sentinel)
	require.Equal(t, []string{"Bad"}, visits)
}
