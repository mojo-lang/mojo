package commander

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestBootstrapJava(t *testing.T) {
	if _, err := exec.LookPath("protoc"); err != nil {
		t.Skip("integration test requires protoc")
	}
	root := t.TempDir()
	write := func(name, content string) {
		file := filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(file), 0755))
		require.NoError(t, os.WriteFile(file, []byte(content), 0644))
	}
	read := func(name string) string {
		content, err := os.ReadFile(filepath.Join(root, name))
		require.NoError(t, err)
		return string(content)
	}
	module := "module " + lang.MojoGoModule + "\n\ngo 1.24.7\n"
	write("go/go.mod", module)
	manual := "package org.mojolang.mojo.core;\nclass HandWritten {}\n"
	write("java/src/main/java/org/mojolang/mojo/core/HandWritten.java", manual)
	for _, name := range mpm.MojoPackageNames {
		deps, source := "", "type String\ntype Probe { value: String @1 }"
		if name != "core" {
			deps = "dependencies: { 'mojo.core': {repository: 'github.com/mojo-lang/mojo/packages/core'} }"
			source = "type Reference { probe: mojo.core.Probe @1 }"
		}
		write("packages/"+name+"/package.mojo", fmt.Sprintf("package mojo.%s {\nrepository: 'github.com/mojo-lang/mojo/packages/%s'\nauthors: [{author: 'test', organization: 'mojo-lang.org'}]\n%s\n}", name, name, deps))
		write("packages/"+name+"/mojo/"+name+"/probe.mojo", source)
	}
	require.NoError(t, Bootstrap(filepath.Join(root, "go"), "java"))
	referencePath := "java/src/main/java/org/mojolang/mojo/document/Reference.java"
	reference := read(referencePath)
	require.Contains(t, reference, "org.mojolang.mojo.core.Probe")
	require.Contains(t, read("java/src/main/java/org/mojolang/mojo/rpc/Reference.java"), "package org.mojolang.mojo.rpc;")
	require.Contains(t, read("java/src/main/java/org/mojolang/mojo/core/Probe.java"), "getValue()")
	require.NoDirExists(t, filepath.Join(root, "packages/core/java"))
	require.NoDirExists(t, filepath.Join(root, "go/pkg/mojo"))
	require.Equal(t, module, read("go/go.mod"))
	// A single-package regeneration removes only that package's obsolete
	// generated classes. Handwritten and other packages' files survive.
	write("packages/core/mojo/core/probe.mojo", "type String\ntype ChangedProbe { revision: String @1 }")
	b := Builder{Pwd: root, Path: "packages/core", Targets: "java"}
	require.NoError(t, b.Execute())
	require.NoFileExists(t, filepath.Join(root, "java/src/main/java/org/mojolang/mojo/core/Probe.java"))
	require.Contains(t, read("java/src/main/java/org/mojolang/mojo/core/ChangedProbe.java"), "getRevision()")
	require.Equal(t, reference, read(referencePath))
	require.Equal(t, manual, read("java/src/main/java/org/mojolang/mojo/core/HandWritten.java"))
	b = Builder{Pwd: root, Path: "packages/core", Targets: "java", Output: "custom-java"}
	require.NoError(t, b.Execute())
	require.Contains(t, read("custom-java/src/main/java/org/mojolang/mojo/core/ChangedProbe.java"), "getRevision()")
}

func TestBootstrapJavaUnified(t *testing.T) {
	if _, err := exec.LookPath("protoc"); err != nil {
		t.Skip("integration test requires protoc")
	}
	root := t.TempDir()
	write := func(name, content string) {
		file := filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(file), 0755))
		require.NoError(t, os.WriteFile(file, []byte(content), 0644))
	}
	read := func(name string) string {
		content, err := os.ReadFile(filepath.Join(root, name))
		require.NoError(t, err)
		return string(content)
	}
	module := "module " + lang.MojoGoModule + "\n\ngo 1.24.7\n"
	write("go/go.mod", module)
	manual := "package org.mojolang.mojo.core;\nclass HandWritten {}\n"
	write("java/src/main/java/org/mojolang/mojo/core/HandWritten.java", manual)
	manifest := ""
	for _, name := range mpm.MojoPackageNames {
		deps, source := "", "type String\ntype Probe { value: String @1 }"
		if name != "core" {
			deps = "dependencies: { 'mojo.core': {repository: 'github.com/mojo-lang/mojo/mojo/core'} }"
			source = "type Reference { probe: mojo.core.Probe @1 }"
		}
		manifest = fmt.Sprintf("package mojo.%s {\nrepository: 'github.com/mojo-lang/mojo'\nauthors: [{author: 'test', organization: 'mojo-lang.org'}]\n%s\n}", name, deps) + "\n" + manifest
		write("mojo/"+name+"/probe.mojo", source)
	}
	write("package.mojo", manifest)
	all := Builder{Pwd: root, Path: root, Targets: "java"}
	require.NoError(t, all.Execute())
	require.NoError(t, Bootstrap(filepath.Join(root, "go"), "java"))
	referencePath := "java/src/main/java/org/mojolang/mojo/document/Reference.java"
	reference := read(referencePath)
	require.Contains(t, reference, "org.mojolang.mojo.core.Probe")
	require.Contains(t, read("java/src/main/java/org/mojolang/mojo/rpc/Reference.java"), "package org.mojolang.mojo.rpc;")
	require.Contains(t, read("java/src/main/java/org/mojolang/mojo/core/Probe.java"), "getValue()")
	require.NoDirExists(t, filepath.Join(root, "mojo/core/java"))
	require.NoDirExists(t, filepath.Join(root, "go/pkg/mojo"))
	require.Equal(t, module, read("go/go.mod"))
	// A single-package regeneration removes only that package's obsolete
	// generated classes. Handwritten and other packages' files survive.
	write("mojo/core/probe.mojo", "type String\ntype ChangedProbe { revision: String @1 }")
	b := Builder{Pwd: root, Path: "mojo/core", Targets: "java"}
	require.NoError(t, b.Execute())
	require.NoFileExists(t, filepath.Join(root, "java/src/main/java/org/mojolang/mojo/core/Probe.java"))
	require.Contains(t, read("java/src/main/java/org/mojolang/mojo/core/ChangedProbe.java"), "getRevision()")
	require.Equal(t, reference, read(referencePath))
	require.Equal(t, manual, read("java/src/main/java/org/mojolang/mojo/core/HandWritten.java"))
	b = Builder{Pwd: root, Path: "mojo/core", Targets: "java", Output: "custom-java"}
	require.NoError(t, b.Execute())
	require.Contains(t, read("custom-java/src/main/java/org/mojolang/mojo/core/ChangedProbe.java"), "getRevision()")
}
