package openapi

import "github.com/mojo-lang/mojo/go/pkg/mojo/core"

func ExpressionFromValue(value *core.Value) (*Expression, error) {
	expr := &Expression{}
	if err := expr.FromValue(value); err != nil {
		return nil, err
	}
	return expr, nil
}

func (x *Expression) FromValue(value *core.Value) error {
	if value != nil && value.GetKind() == core.ValueKind_VALUE_KIND_STRING {
		return x.Parse(value.GetString())
	}
	return nil
}
