package openapi

import (
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

const (
	urlExprType        = "$url"
	methodExprType     = "$method"
	statusCodeExprType = "$statusCode"
	requestExprPrefix  = "$request."
	responseExprPrefix = "$response."
)

func (x *Expression) Format() string {
	if x != nil {
		builder := new(strings.Builder)
		builder.WriteString("$")
		builder.WriteString(x.Type.Format())
		if x.Source != nil {
			builder.WriteByte('.')
			builder.WriteString(x.Source.Format())
		}

		return builder.String()
	}
	return ""
}

func (x *Expression) ToString() string {
	return x.Format()
}

func (x *Expression) Parse(value string) error {
	if x != nil {
		if len(value) >= 4 {
			switch value[0:4] {
			case urlExprType:
				x.Type = Expression_TYPE_URL
			case "$met":
				if value == methodExprType {
					x.Type = Expression_TYPE_METHOD
				} else {
					return core.NewMalformedSyntaxError("invalid expression method type: %s", value)
				}
			case "$sta":
				if value == statusCodeExprType {
					x.Type = Expression_TYPE_STATUS_CODE
				} else {
					return core.NewMalformedSyntaxError("invalid expression statusCode type: %s", value)
				}
			case "$req":
				if strings.HasPrefix(value, requestExprPrefix) {
					source, err := ParseExpressionSource(value[len(requestExprPrefix):])
					if err != nil {
						return core.NewMalformedSyntaxError("invalid expression request source type: %s, error: %s", value, err.Error())
					}
					x.Type = Expression_TYPE_REQUEST
					x.Source = source
				} else {
					return core.NewMalformedSyntaxError("invalid expression request type: %s", value)
				}
			case "$res":
				if strings.HasPrefix(value, responseExprPrefix) {
					source, err := ParseExpressionSource(value[len(responseExprPrefix):])
					if err != nil {
						return core.NewMalformedSyntaxError("invalid expression response source type: %s, error: %s", value, err.Error())
					}
					x.Type = Expression_TYPE_RESPONSE
					x.Source = source
				} else {
					return core.NewMalformedSyntaxError("invalid expression response type: %s", value)
				}
			default:
				return core.NewMalformedSyntaxError("invalid expression: %s", value)
			}
			return nil
		}
		return core.NewMalformedSyntaxError("invalid expression: %s", value)
	}
	return nil
}

func ParseExpression(value string) (*Expression, error) {
	v := &Expression{}
	if err := v.Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
