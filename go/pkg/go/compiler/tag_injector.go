// Copyright (c) 2016, Diep Pham
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// [source](https://github.com/favadi/protoc-go-inject-tag)

package compiler

import (
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

var (
	rInject = regexp.MustCompile("`.+`$")
	rTags   = regexp.MustCompile(`[\w_]+:"[^"]+"`)
)

type TagInjector struct {
}

func (t TagInjector) InjectTag(fileName string, content []byte, xxxSkip []string, injections Injections) ([]byte, error) {
	areas, err := parseFile(fileName, content, xxxSkip, injections)
	if err != nil {
		return nil, err
	}

	return injectFile(content, areas), nil
}

type textArea struct {
	Start      int
	End        int
	CurrentTag string
	InjectTag  string
}

type TagItem struct {
	Key   string
	Value string
}

type TagItems []TagItem

type Injections map[string]TagItems

func parseFile(fileName string, content []byte, xxxSkip []string, injections Injections) (areas []textArea, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, content, parser.ParseComments)
	if err != nil {
		return
	}

	for _, decl := range f.Decls {
		// check if is generic declaration
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		var typeSpec *ast.TypeSpec
		for _, spec := range genDecl.Specs {
			if ts, tsOK := spec.(*ast.TypeSpec); tsOK {
				typeSpec = ts
				break
			}
		}

		// skip if can't get type spec
		if typeSpec == nil {
			continue
		}

		// not a struct, skip
		structDecl, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}

		var skipTags TagItems
		for _, skip := range xxxSkip {
			skipTags = append(skipTags, TagItem{
				Key:   skip,
				Value: "-",
			})
		}

		structName := typeSpec.Name.Name
		for _, field := range structDecl.Fields.List {
			if len(field.Names) > 0 {
				name := field.Names[0].Name
				currentTag := field.Tag.Value
				area := textArea{
					Start:      int(field.Pos()),
					End:        int(field.End()),
					CurrentTag: currentTag[1 : len(currentTag)-1],
				}

				if len(skipTags) > 0 && strings.HasPrefix(name, "XXX") {
					area.InjectTag = skipTags.format()
					areas = append(areas, area)
					//logs.Infow("inject field tag", "file", fileName, "struct", structName, "field", name, "injection", area.InjectTag)
					continue
				}

				key := structName + "." + name
				if injection, ok := injections[key]; ok {
					area.InjectTag = injection.format()
					areas = append(areas, area)
					//logs.Infow("inject field tag", "file", fileName, "struct", structName, "field", name, "injection", area.InjectTag)
				}
			}
		}
	}
	//logs.Infof("parsed file %s, number of fields to inject custom tags: %d", fileName, len(areas))
	return
}

func injectFile(input []byte, areas []textArea) []byte {
	// inject custom tags from tail of file first to preserve order
	for i := range areas {
		area := areas[len(areas)-i-1]
		//logs.Debugf("inject custom tag %q to expression %q", area.InjectTag, string(input[area.Start-1:area.End-1]))
		input = injectTag(input, area)
	}

	return input
}

func (t TagItems) format() string {
	var tags []string
	for _, item := range t {
		item.Value = util.RemoveQuote(item.Value)
		tags = append(tags, fmt.Sprintf(`%s:"%s"`, item.Key, item.Value))
	}
	return strings.Join(tags, " ")
}

func (t TagItems) override(items TagItems) TagItems {
	var overrided []TagItem
	for i := range t {
		var dup = -1
		for j := range items {
			if t[i].Key == items[j].Key {
				dup = j
				break
			}
		}
		if dup == -1 {
			overrided = append(overrided, t[i])
		} else {
			overrided = append(overrided, items[dup])
			items = append(items[:dup], items[dup+1:]...)
		}
	}
	return append(overrided, items...)
}

func newTagItems(tag string) TagItems {
	var items []TagItem
	splitted := rTags.FindAllString(tag, -1)

	for _, t := range splitted {
		sepPos := strings.Index(t, ":")
		items = append(items, TagItem{
			Key:   t[:sepPos],
			Value: t[sepPos+1:],
		})
	}
	return items
}

func injectTag(contents []byte, area textArea) (injected []byte) {
	expr := make([]byte, area.End-area.Start)
	copy(expr, contents[area.Start-1:area.End-1])
	cti := newTagItems(area.CurrentTag)
	iti := newTagItems(area.InjectTag)
	ti := cti.override(iti)
	expr = rInject.ReplaceAll(expr, []byte(fmt.Sprintf("`%s`", ti.format())))
	injected = append(injected, contents[:area.Start-1]...)
	injected = append(injected, expr...)
	injected = append(injected, contents[area.End-1:]...)
	return
}
