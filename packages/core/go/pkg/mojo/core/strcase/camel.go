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

var irregularCamelRules []caseRule

func RegisterIrregularCamelCaseRule(rule string, replacement string) {
	r, _ := regexp.Compile(rule)
	irregularCamelRules = append(irregularCamelRules, caseRule{
		Regexp:      r,
		Replacement: replacement,
	})
}

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	return applyCamelRules(toCamelInitCase(s, true))
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	s = applySnakeRules(strings.TrimSpace(s))
	return toCamelInitCase(s, false)
}

func applyCamelRules(s string) string {
	for _, rule := range irregularCamelRules {
		s = string(rule.Regexp.ReplaceAll([]byte(s), []byte(rule.Replacement)))
	}
	return s
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		vIsNum := v >= '0' && v <= '9'

		nextIsCap := false
		nextIsLow := false
		if i+1 < len(s) {
			next := s[i+1]
			nextIsCap = next >= 'A' && next <= 'Z'
			nextIsLow = next >= 'a' && next <= 'z'
		}

		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum {
			// treat the a2b as a whole word
			if v == '2' && (nextIsCap || nextIsLow) && i > 0 && unicode.IsLetter(rune(s[i-1])) {
				n.WriteByte(v)
				capNext = false
				continue
			}
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}
