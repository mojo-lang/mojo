package handlers

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"path"
	"sort"
	"strconv"
	"strings"
)

func importNames(file *ast.File) map[string]string {
	names := map[string]string{}
	for _, spec := range file.Imports {
		imported, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		name := path.Base(imported)
		if spec.Name != nil {
			name = spec.Name.Name
		}
		names[name] = imported
	}
	return names
}

// Compare types using import paths, ignoring parameter names and import aliases.
// This is deliberately syntax-only: the API may not have been generated yet.
func signature(fn *ast.FuncType, file *ast.File) string {
	imports := importNames(file)
	var typeString func(ast.Expr) string
	typeString = func(expr ast.Expr) string {
		switch expr := expr.(type) {
		case *ast.Ident:
			return expr.Name
		case *ast.StarExpr:
			return "*" + typeString(expr.X)
		case *ast.SelectorExpr:
			if pkg, ok := expr.X.(*ast.Ident); ok {
				if imported := imports[pkg.Name]; imported != "" {
					return strconv.Quote(imported) + "." + expr.Sel.Name
				}
			}
		case *ast.ParenExpr:
			return typeString(expr.X)
		}
		var out bytes.Buffer
		_ = printer.Fprint(&out, token.NewFileSet(), expr)
		return out.String()
	}
	fields := func(list *ast.FieldList) string {
		var types []string
		if list != nil {
			for _, field := range list.List {
				count := len(field.Names)
				if count == 0 {
					count = 1
				}
				for i := 0; i < count; i++ {
					types = append(types, typeString(field.Type))
				}
			}
		}
		return strings.Join(types, ", ")
	}
	return "func(" + fields(fn.Params) + ") (" + fields(fn.Results) + ")"
}

// Conditional RPC declarations require explicit handling rather than generating
// an unconditional method that could collide on another OS or with other tags.
func conditionalSource(name string, file *ast.File) bool {
	for _, group := range file.Comments {
		if group.Pos() > file.Package {
			break
		}
		for _, comment := range group.List {
			if strings.HasPrefix(comment.Text, "//go:build ") || strings.HasPrefix(comment.Text, "// +build ") {
				return true
			}
		}
	}
	parts := strings.Split(strings.TrimSuffix(path.Base(name), ".go"), "_")
	if len(parts) < 2 {
		return false
	}
	// Known GOOS/GOARCH suffixes as understood by the supported Go toolchains.
	targets := " aix android darwin dragonfly freebsd hurd illumos ios js linux nacl netbsd openbsd plan9 solaris wasip1 windows zos 386 amd64 amd64p32 arm armbe arm64 arm64be loong64 mips mipsle mips64 mips64le mips64p32 mips64p32le ppc ppc64 ppc64le riscv riscv64 s390 s390x sparc sparc64 wasm "
	return strings.Contains(targets, " "+parts[len(parts)-1]+" ")
}

func appendMethods(current, body []byte, file, expected *ast.File, set *token.FileSet) (io.Reader, error) {
	prefix := "package handlers\n"
	addedSet := token.NewFileSet()
	added, err := parser.ParseFile(addedSet, "additions.go", append([]byte(prefix), body...), parser.ParseComments)
	if err != nil {
		return nil, err
	}
	imports := importNames(file)
	available := map[string]string{}
	used := map[string]bool{}
	for alias, imported := range imports {
		available[imported] = alias
		used[alias] = true
	}
	for _, decl := range file.Decls {
		for _, name := range declaredNames(decl) {
			used[name] = true
		}
	}
	expectedImports := importNames(expected)
	needed := map[string]bool{}
	ast.Inspect(added, func(node ast.Node) bool {
		if sel, ok := node.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && expectedImports[id.Name] != "" {
				needed[id.Name] = true
			}
		}
		return true
	})
	aliases := make([]string, 0, len(needed))
	for alias := range needed {
		aliases = append(aliases, alias)
	}
	sort.Strings(aliases)
	rename := map[string]string{}
	var extra strings.Builder
	for _, alias := range aliases {
		imported := expectedImports[alias]
		chosen, exists := available[imported]
		if exists && (chosen == "." || chosen == "_") {
			return nil, fmt.Errorf("cannot append RPC using %s imported as %s; give the import an explicit usable alias", imported, chosen)
		}
		if !exists {
			chosen = alias
			for i := 2; used[chosen]; i++ {
				chosen = fmt.Sprintf("%s%d", alias, i)
			}
			extra.WriteString("\nimport " + chosen + " " + strconv.Quote(imported) + "\n")
			used[chosen] = true
			available[imported] = chosen
		}
		rename[alias] = chosen
	}
	type edit struct {
		start, end int
		text       string
	}
	var edits []edit
	ast.Inspect(added, func(node ast.Node) bool {
		if sel, ok := node.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				if to := rename[id.Name]; to != "" && to != id.Name {
					edits = append(edits, edit{addedSet.Position(id.Pos()).Offset - len(prefix), addedSet.Position(id.End()).Offset - len(prefix), to})
				}
			}
		}
		return true
	})
	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	body = append([]byte(nil), body...)
	for _, edit := range edits {
		body = append(append(append([]byte{}, body[:edit.start]...), []byte(edit.text)...), body[edit.end:]...)
	}
	offset := set.Position(file.Name.End()).Offset
	var out bytes.Buffer
	out.Write(current[:offset])
	out.WriteString(extra.String())
	out.Write(current[offset:])
	out.WriteString("\n")
	out.Write(body)
	return &out, nil
}

func removeUnusedImports(file *ast.File) {
	used := map[string]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		if sel, ok := node.(*ast.SelectorExpr); ok {
			if pkg, ok := sel.X.(*ast.Ident); ok {
				used[pkg.Name] = true
			}
		}
		return true
	})
	var declarations []ast.Decl
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT {
			declarations = append(declarations, decl)
			continue
		}
		var specs []ast.Spec
		for _, spec := range gen.Specs {
			imp := spec.(*ast.ImportSpec)
			imported, _ := strconv.Unquote(imp.Path.Value)
			name := path.Base(imported)
			if imp.Name != nil {
				name = imp.Name.Name
			}
			if name == "_" || name == "." || used[name] {
				specs = append(specs, spec)
			}
		}
		gen.Specs = specs
		if len(specs) > 0 {
			declarations = append(declarations, gen)
		}
	}
	file.Decls = declarations
}
