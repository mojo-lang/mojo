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
		if binary.Operator.Symbol == "?" && binary.RightHandArgument.GetConditionalExpr() != nil {
			conditional := binary.RightHandArgument.GetConditionalExpr()
			conditional.Condition = prefix
			return binary.RightHandArgument
		}
	}

	bins := []*lang.BinaryExpr{
		{Operator: &lang.Operator{}, RightHandArgument: prefix},
	}
	bins = append(bins, binaries...)
	binaries = bins

	for len(binaries) > 1 {
		highest := b.highestPrecedence(binaries)

		bins = []*lang.BinaryExpr{}
		for i, binary := range binaries {
			if binary.Operator.Precedence == highest {
				if i > 0 {
					last := bins[len(bins)-1]
					binary.LeftHandArgument = last.RightHandArgument
					last.RightHandArgument = lang.NewBinaryExpression(binary)
				} else {
					bins = append(bins, binary)
				}
			} else {
				bins = append(bins, binary)
			}
		}

		binaries = bins
	}

	return binaries[0].RightHandArgument
}

func (b BinaryExprParser) highestPrecedence(binaries []*lang.BinaryExpr) int32 {
	var highest int32
	for _, binary := range binaries {
		precedence := precedenceIndex[binary.Operator.Symbol]
		binary.Operator.Precedence = precedence

		if precedence > highest {
			highest = precedence
		}
	}
	return highest
}
