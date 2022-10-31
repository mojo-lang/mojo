package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

var precedenceIndex = map[string]int32{
	"|":   5,
	"=":   10,
	"+=":  10,
	"-=":  10,
	"**=": 10,
	"*=":  10,
	"/=":  10,
	"%=":  10,
	"<<=": 10,
	">>=": 10,
	"&=":  10,
	"^=":  10,
	"|=":  10,
	"||":  12,
	"&&":  13,
	"!=":  20,
	"==":  20,
	">=":  21,
	">":   21,
	"<=":  21,
	"<":   21,
	"<<":  30,
	">>":  30,
	"-":   40,
	"+":   40,
	"%":   45,
	"/":   45,
	"*":   45,
	"**":  47,
}

type BinaryExprParser struct {
}

func (b BinaryExprParser) Parse(prefix *lang.Expression, binaries []*lang.BinaryExpr) *lang.Expression {
	if len(binaries) == 0 || prefix == nil {
		return nil
	}

	if len(binaries) == 1 {
		binary := binaries[0]
		if binary.GetOperator().GetSymbol() == "?" && binary.RightArgument.GetConditionalExpr() != nil {
			conditional := binary.RightArgument.GetConditionalExpr()
			conditional.Condition = prefix
			return binary.RightArgument
		}
	}

	bins := []*lang.BinaryExpr{
		lang.NewBinaryExpr(&lang.Operator{}, nil, prefix),
	}
	bins = append(bins, binaries...)
	binaries = bins

	for len(binaries) > 1 {
		highest := b.highestPrecedence(binaries)

		bins = []*lang.BinaryExpr{}
		for i, binary := range binaries {
			if binary.GetOperator().GetPrecedence() == highest {
				if i > 0 {
					last := bins[len(bins)-1]
					binary.LeftArgument = last.RightArgument
					last.RightArgument = lang.NewBinaryExpression(binary)
				} else {
					bins = append(bins, binary)
				}
			} else {
				bins = append(bins, binary)
			}
		}

		binaries = bins
	}

	return binaries[0].RightArgument
}

func (b BinaryExprParser) highestPrecedence(binaries []*lang.BinaryExpr) int32 {
	var highest int32
	for _, binary := range binaries {
		precedence := precedenceIndex[binary.GetOperator().GetSymbol()]
		binary.GetOperator().Precedence = precedence

		if precedence > highest {
			highest = precedence
		}
	}
	return highest
}
