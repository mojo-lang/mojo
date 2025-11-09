/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 * from (github.com/iancoleman/strcase)
 */
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strcase

import (
	"regexp"
	"strings"
	"unicode"
)

func init() {
	RegisterIrregularCaseRule("UInt", "Uint")
}

type caseRule struct {
	*regexp.Regexp
	Replacement string
}

var irregularSnakeRules []caseRule

func RegisterIrregularCaseRule(rule string, replacement string) {
	RegisterIrregularSnakeCaseRule(rule, replacement)
	if reg := regexp.MustCompile(`[a-zA-Z0-9]`); reg.Match([]byte(rule)) {
		RegisterIrregularCamelCaseRule(replacement, rule)
	}
}

func RegisterIrregularSnakeCaseRule(rule string, replacement string) {
	r, _ := regexp.Compile(rule)
	irregularSnakeRules = append(irregularSnakeRules, caseRule{
		Regexp:      r,
		Replacement: replacement,
	})
}

func ToSnake(name string) string {
	return ToDelimited(name, '_')
}

func ToSnakeWithIgnore(s string, ignore string) string {
	return ToScreamingDelimited(s, '_', ignore, false)
}

// ToScreamingSnake converts a string to SCREAMING_SNAKE_CASE
func ToScreamingSnake(s string) string {
	return ToScreamingDelimited(s, '_', "", true)
}

// ToKebab converts a string to kebab-case
func ToKebab(s string) string {
	return ToDelimited(s, '-')
}

func ToKebabWithIgnore(s string, ignore string) string {
	return ToScreamingDelimited(s, '-', ignore, false)
}

// ToScreamingKebab converts a string to SCREAMING-KEBAB-CASE
func ToScreamingKebab(s string) string {
	return ToScreamingDelimited(s, '-', "", true)
}

// ToDelimited converts a string to delimited.snake.case
// (in this case `delimiter = '.'`)
func ToDelimited(s string, delimiter uint8) string {
	return ToScreamingDelimited(s, delimiter, "", false)
}

func applySnakeRules(s string) string {
	for _, rule := range irregularSnakeRules {
		s = string(rule.Regexp.ReplaceAll([]byte(s), []byte(rule.Replacement)))
	}
	return s
}

func ToScreamingDelimited(s string, delimiter uint8, ignore string, screaming bool) string {
	s = applySnakeRules(strings.TrimSpace(s))

	n := strings.Builder{}
	n.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow && screaming {
			v += 'A'
			v -= 'a'
		} else if vIsCap && !screaming {
			v += 'a'
			v -= 'A'
		}

		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if i+1 < len(s) {
			next := s[i+1]
			vIsNum := v >= '0' && v <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			// add underscore if next letter case type is changed
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
				if (vIsCap && nextIsNum) || (vIsLow && nextIsNum) {
					n.WriteByte(v)
					continue
				}
				// treat the a2b as a whole word
				if vIsNum && v == '2' && (nextIsCap || nextIsLow) && i > 0 && unicode.IsLetter(rune(s[i-1])) {
					n.WriteByte(v)
					continue
				}

				prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(s[i-1]), ignore)
				if !prevIgnore {
					if vIsCap && nextIsLow {
						if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
							n.WriteByte(delimiter)
						}
					}
					n.WriteByte(v)
					if vIsLow || vIsNum || nextIsNum {
						n.WriteByte(delimiter)
					}
					continue
				}
			}
		}

		if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
			// replace space/underscore/hyphen/dot with delimiter
			n.WriteByte(delimiter)
		} else {
			n.WriteByte(v)
		}
	}

	return n.String()
}

// func sanitizeRule(rule string) *regexp.Regexp {
//	if isExpr(rule) {
//		return regexp.MustCompile(rule)
//	}
//
//	return regexp.MustCompile(`(?i)^` + rule + `$`)
// }
