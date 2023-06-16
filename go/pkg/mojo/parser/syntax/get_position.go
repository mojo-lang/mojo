package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func GetPosition(token antlr.Token) *lang.Position {
	position := &lang.Position{}
	position.Line = int64(token.GetLine())
	position.Column = int64(token.GetColumn())
	return position
}
