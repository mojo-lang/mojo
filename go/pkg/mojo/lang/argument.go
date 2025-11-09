package lang

import "errors"

func NewArgument(expr *Expression) *Argument {
	return &Argument{
		Value: expr,
	}
}

func NewBoolLiteralArgument(expr *BoolLiteralExpr) *Argument {
	return &Argument{
		Value: NewBoolLiteralExpression(expr),
	}
}

func NewIntegerLiteralArgument(expr *IntegerLiteralExpr) *Argument {
	return &Argument{
		Value: NewIntegerLiteralExpression(expr),
	}
}

func NewFloatLiteralArgument(expr *FloatLiteralExpr) *Argument {
	return &Argument{
		Value: NewFloatLiteralExpression(expr),
	}
}

func NewStringLiteralArgument(expr *StringLiteralExpr) *Argument {
	return &Argument{
		Value: NewStringLiteralExpression(expr),
	}
}

func NewStringArgument(val string) *Argument {
	return NewStringLiteralArgument(NewStringLiteralExpr(val))
}

func NewArrayLiteralArgument(expr *ArrayLiteralExpr) *Argument {
	return &Argument{
		Value: NewArrayLiteralExpression(expr),
	}
}

func NewMapLiteralArgument(expr *MapLiteralExpr) *Argument {
	return &Argument{
		Value: NewMapLiteralExpression(expr),
	}
}

func NewObjectLiteralArgument(expr *ObjectLiteralExpr) *Argument {
	return &Argument{
		Value: NewObjectLiteralExpression(expr),
	}
}

func (x *Argument) SetStartPosition(position *Position) *Argument {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
	return x
}

func (x *Argument) SetEndPosition(position *Position) *Argument {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
	return x
}

func (x *Argument) SetLabel(label string) *Argument {
	if x != nil {
		x.Label = label
	}
	return x
}

func (x *Argument) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil Argument")
}

func (x *Argument) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if len(x.Label) > 0 {
			terms = append(terms, &Term{
				StartPosition: x.StartPosition,
				EndPosition: &Position{
					Line:   x.StartPosition.Line,
					Column: x.StartPosition.Column + int64(len(x.Label)),
				},
				Type:  "Label",
				Value: x.Label,
			})
		}
		return append(terms, x.Value)
	}
	return nil
}

func (x *Argument) GetNullLiteralExpr() *NullLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetNullLiteralExpr()
	}
	return nil
}

func (x *Argument) GetIntegerLiteralExpr() *IntegerLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetIntegerLiteralExpr()
	}
	return nil
}

func (x *Argument) GetFloatLiteralExpr() *FloatLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetFloatLiteralExpr()
	}
	return nil
}

func (x *Argument) GetBoolLiteralExpr() *BoolLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetBoolLiteralExpr()
	}
	return nil
}

func (x *Argument) GetStringLiteralExpr() *StringLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetStringLiteralExpr()
	}
	return nil
}

func (x *Argument) GetObjectLiteralExpr() *ObjectLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetObjectLiteralExpr()
	}
	return nil
}

func (x *Argument) GetArrayLiteralExpr() *ArrayLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetArrayLiteralExpr()
	}
	return nil
}

func (x *Argument) GetMapLiteralExpr() *MapLiteralExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetMapLiteralExpr()
	}
	return nil
}

func (x *Argument) GetIdentifierExpr() *IdentifierExpr {
	if x != nil && x.Value != nil {
		return x.Value.GetIdentifierExpr()
	}
	return nil
}

func (x *Argument) GetBool() (bool, error) {
	return x.GetValue().EvalBoolLiteral()
}

func (x *Argument) GetInteger() (int64, error) {
	return x.GetValue().EvalIntegerLiteral()
}

func (x *Argument) GetFloat() (float64, error) {
	return x.GetValue().EvalFloatLiteral()
}

func (x *Argument) GetString() (string, error) {
	return x.GetValue().EvalStringLiteral()
}

func (x *Argument) GetArray() ([]*Expression, error) {
	return x.GetValue().EvalArrayLiteral()
}

func (x *Argument) GetIntegerArray() ([]int64, error) {
	return x.GetValue().EvalIntegerArrayLiteral()
}

func (x *Argument) GetFloatArray() ([]float64, error) {
	return x.GetValue().EvalFloatArrayLiteral()
}

func (x *Argument) GetStringArray() ([]string, error) {
	return x.GetValue().EvalStringArrayLiteral()
}

func (x *Argument) IterateStringMap(iterator func(key string, value *Expression) error) error {
	return x.GetValue().EvalStringMapLiteral(iterator)
}
