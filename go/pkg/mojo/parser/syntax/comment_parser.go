package syntax

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type CommentParser struct {
}

func (c CommentParser) Parse(stream *antlr.CommonTokenStream) []*lang.Comment {
	var comments []*lang.Comment
	allTokens := stream.GetAllTokens()

	expectEOL := false
	var multiLineComment *lang.MultiLineComment
	closeMultiLineComment := func(token antlr.Token) {
		if multiLineComment != nil {
			endPosition := multiLineComment.Lines[len(multiLineComment.Lines)-1].EndPosition
			multiLineComment.EndPosition = &lang.Position{
				Line:   endPosition.Line,
				Column: endPosition.Column,
			}
			removeMultiLineCommentPrefix(multiLineComment)
			comments = append(comments, lang.NewMultiLineCommentComment(multiLineComment))
			multiLineComment = nil
		}
	}

	for i, token := range allTokens {
		tokenType := token.GetTokenType()
		if token.GetChannel() == antlr.TokenHiddenChannel {
			if isComment(tokenType) {
				if isLineComment(tokenType) {
					line := toLineComment(token)
					if i < len(allTokens)-1 {
						line.EndPosition = GetPosition(allTokens[i+1])
					}
					if !leadingWhiteSpaceInLine(token, i, allTokens) {
						line.Following = true
					}

					if multiLineComment == nil {
						multiLineComment = &lang.MultiLineComment{
							StartPosition: GetPosition(token),
						}
					}

					expectEOL = true
					multiLineComment.Lines = append(multiLineComment.Lines, line)
				} else {
					closeMultiLineComment(token)
					block := toBlockComment(token)
					block.HeadEmbedded = !leadingWhiteSpaceInLine(token, i, allTokens)
					block.TailEmbedded = !tailingWhiteSpaceInLine(token, i, allTokens)
					if i < len(allTokens)-1 {
						block.EndPosition = GetPosition(allTokens[i+1])
					}
					comments = append(comments, lang.NewBlockCommentComment(block))
				}
			}
		} else if tokenType == MojoParserEOL {
			if expectEOL {
				expectEOL = false
			} else {
				if multiLineComment != nil {
					closeMultiLineComment(token)
				}
			}
		} else {
			closeMultiLineComment(token)
		}
	}

	return comments
}

func leadingWhiteSpaceInLine(token antlr.Token, i int, allTokens []antlr.Token) bool {
	line := token.GetLine()
	preIndex := i - 1
	for preIndex >= 0 {
		preToken := allTokens[preIndex]
		if preToken.GetLine() == line {
			if preToken.GetTokenType() == MojoParserWS {
				preIndex--
				continue
			} else {
				return false
			}
		} else {
			return true
		}
	}
	return true // in first line
}

func tailingWhiteSpaceInLine(token antlr.Token, i int, allTokens []antlr.Token) bool {
	line := token.GetLine()
	i++
	for i < len(allTokens) {
		postToken := allTokens[i]
		if postToken.GetLine() == line {
			tokenType := postToken.GetTokenType()
			if tokenType == MojoParserWS || tokenType == MojoParserEOL || tokenType == MojoParserEOF {
				i++
				continue
			} else {
				return false
			}
		} else {
			return true
		}
	}
	return true // in first line
}

func toBlockComment(token antlr.Token) *lang.BlockComment {
	return &lang.BlockComment{
		StartPosition: GetPosition(token),
		Content:       token.GetText(),
	}
}

func toLineComment(token antlr.Token) *lang.LineComment {
	return &lang.LineComment{
		StartPosition: GetPosition(token),
		Content:       token.GetText(),
	}
}

func isComment(tokenType int) bool {
	return tokenType == MojoParserBLOCK_COMMENT || isLineComment(tokenType)
}

func isLineComment(tokenType int) bool {
	return tokenType == MojoParserLINE_COMMENT || tokenType == MojoParserLINE_COMMENT_DISTINCT_DOCUMENT
}

func removeMultiLineCommentPrefix(comment *lang.MultiLineComment) {
	if comment != nil && len(comment.Lines) > 0 {
		for _, line := range comment.Lines {
			line.Content = strings.TrimPrefix(line.Content, "//")
		}
	}
}
