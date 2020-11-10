package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/lang/go/pkg/lang"
)

func GetPosition(token antlr.Token) *lang.Position {
	position := &lang.Position{}
	position.Line = int32(token.GetLine())
	position.Column = int32(token.GetColumn())
	return position
}
