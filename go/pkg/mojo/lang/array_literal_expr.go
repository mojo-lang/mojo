package lang

import "errors"

func (x *ArrayLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *ArrayLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *ArrayLiteralExpr) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil ArrayLiteralExpr")
}

func (x *ArrayLiteralExpr) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}
		for _, element := range x.Elements {
			terms = append(terms, element)
		}

		return terms
	}
	return nil
}

func (x *ArrayLiteralExpr) EvalIntegerArray() ([]int64, error) {
	if x != nil {
		var values []int64
		for _, elem := range x.Elements {
			value, err := elem.EvalIntegerLiteral()
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	}
	return nil, errors.New("ArrayLiteralExpr is nil")
}

func (x *ArrayLiteralExpr) EvalFloatArray() ([]float64, error) {
	if x != nil {
		var values []float64
		for _, elem := range x.Elements {
			value, err := elem.EvalFloatLiteral()
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	}
	return nil, errors.New("ArrayLiteralExpr is nil")
}

func (x *ArrayLiteralExpr) EvalStringArray() ([]string, error) {
	if x != nil {
		var values []string
		for _, elem := range x.Elements {
			value, err := elem.EvalStringLiteral()
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	}
	return nil, errors.New("ArrayLiteralExpr is nil")
}
