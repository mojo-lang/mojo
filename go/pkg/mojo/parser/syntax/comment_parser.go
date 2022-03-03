package syntax

import "github.com/antlr/antlr4/runtime/Go/antlr"

type CommentParser struct {
}

func (c CommentParser) Parse(stream *antlr.CommonTokenStream) {
	var tokens []antlr.Token
	for _, token := range stream.GetAllTokens() {
		if token.GetChannel() == antlr.TokenHiddenChannel && isComment(token.GetTokenType()) {
			token.GetTokenType()
			tokens = append(tokens, token)
		}
	}
}

func isComment(tokenType int) bool {
	return tokenType == MojoParserDELIMITED_COMMENT ||
		tokenType == MojoParserLINE_COMMENT ||
		tokenType == MojoParserLINE_COMMENT_FOR_DOCUMENT
}
