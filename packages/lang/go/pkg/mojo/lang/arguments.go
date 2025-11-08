package lang

import "errors"

type Arguments []*Argument

func NewArguments(exprs ...*Expression) Arguments {
	var arguments Arguments
	for _, expr := range exprs {
		arguments = append(arguments, NewArgument(expr))
	}
	return arguments
}

func (a Arguments) ToObjectLiteralExpr() (*ObjectLiteralExpr, error) {
	object := &ObjectLiteralExpr{}
	for _, argument := range a {
		if len(argument.Label) == 0 {
			return nil, errors.New("failed to convert to object because of lack of label in argument")
		}

		object.Fields = append(object.Fields, &ObjectLiteralExpr_Field{
			StartPosition: argument.StartPosition,
			EndPosition:   argument.EndPosition,
			Name:          argument.Label,
			Value:         argument.Value,
		})
	}
	object.Implicit = true
	return object, nil
}
