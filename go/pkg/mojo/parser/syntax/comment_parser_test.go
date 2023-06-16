package syntax

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/stretchr/testify/assert"
)

func getComments(str string) []*lang.Comment {
	stream := antlr.NewCommonTokenStream(NewMojoLexer(antlr.NewInputStream(str)), 0)
	parser := NewMojoParser(stream)
	parser.BuildParseTrees = true
	parser.MojoFile()

	return CommentParser{}.Parse(stream)
}

func TestCommentParser_Parse(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
type Mailbox {
    // free floating comment
    
    address: String
    /* block comment
    *///comment4
// comment5

	following: Bool // following comment
}
`
	comments := getComments(typeDecl)
	assert.NotEmpty(t, comments)
}

func TestCommentParser_Parse2(t *testing.T) {
	const typeDecl = `
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


// comment1

type Mailbox {
    address String
}
`

	comments := getComments(typeDecl)
	assert.NotEmpty(t, comments)
	assert.Equal(t, 2, len(comments))
}
