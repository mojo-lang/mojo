package lang

import "github.com/mojo-lang/mojo/go/pkg/mojo/core"

func NewRangeLiteralExpr(start int64, end int64) *RangeLiteralExpr {
	return &RangeLiteralExpr{
		Value: &core.IntRange{
			Start: start,
			End:   end,
		},
	}
}

func NewRangeLiteralExprFrom(value *core.IntRange) *RangeLiteralExpr {
	return &RangeLiteralExpr{Value: value}
}

func (x *RangeLiteralExpr) EvalValue() *core.IntRange {
	if x != nil {
		return x.Value
	}
	return nil
}
