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

package injection

import (
	"context"
	"go/ast"
	"strings"
	"unicode"

	"github.com/fatih/structtag"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

type TagChanger = func(ctx context.Context, field *ast.Field, tags *structtag.Tags)
type TagInjections map[string]structtag.Tags

type TagInjector struct {
	Injector

	tagChangers []TagChanger
}

func NewTagInjector(xxxSkip []string, injections TagInjections) *TagInjector {
	injector := &TagInjector{}

	injector.OnStructField = injector.onStructField

	if len(xxxSkip) > 0 {
		injector.tagChangers = append(injector.tagChangers, func(ctx context.Context, field *ast.Field, tags *structtag.Tags) {
			fieldName := field.Names[0].Name
			if len(fieldName) > 0 && (strings.HasPrefix(fieldName, "XXX") || unicode.IsLower(rune(fieldName[0]))) {
				for _, skip := range xxxSkip {
					err := tags.Set(&structtag.Tag{
						Key:  skip,
						Name: "-",
					})
					if err != nil {
						logs.Warnw("failed to set the tag", "tag", skip, "error", err)
						return
					}
				}
			}
		})
	}
	if len(injections) > 0 {
		injector.tagChangers = append(injector.tagChangers, func(ctx context.Context, field *ast.Field, tags *structtag.Tags) {
			structName := GetStructName(ctx)
			fieldName := field.Names[0].Name

			key := structName + "." + fieldName
			if injection, ok := injections[key]; ok {
				for _, tag := range injection.Tags() {
					err := tags.Set(tag)
					if err != nil {
						logs.Warnw("failed to set the tag", "tagName", tag.Name, "tagKey", tag.Key, "error", err)
						return
					}
				}
			}
		})
	}

	return injector
}

func (i *TagInjector) RegisterTagChanger(changer TagChanger) *TagInjector {
	i.tagChangers = append(i.tagChangers, changer)
	return i
}

func (i *TagInjector) onStructField(ctx context.Context, field *ast.Field, areaAppender func(area Area)) {
	if field.Tag == nil {
		return
	}

	tagValue := core.RemoveQuote(field.Tag.Value, "`")
	tags, _ := structtag.Parse(tagValue)
	if tags == nil {
		return
	}

	if len(i.tagChangers) > 0 {
		for _, changer := range i.tagChangers {
			changer(ctx, field, tags)
		}

		newTag := tags.String()
		if tagValue != newTag {
			areaAppender(&TextArea{
				Start:   field.Tag.Pos(),
				End:     field.Tag.End(),
				Content: "`" + newTag + "`",
			})
		}
	}
}
