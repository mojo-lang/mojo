package lang

import (
	"errors"
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

const (
	NullLiteralExprName          = "NullLiteralExpr"
	IntegerLiteralExprName       = "IntegerLiteralExpr"
	FloatLiteralExprName         = "FloatLiteralExpr"
	BoolLiteralExprName          = "BoolLiteralExpr"
	StringLiteralExprName        = "StringLiteralExpr"
	ObjectLiteralExprName        = "ObjectLiteralExpr"
	ArrayLiteralExprName         = "ArrayLiteralExpr"
	MapLiteralExprName           = "MapLiteralExpr"
	RangeLiteralExprName         = "RangeLiteralExpr"
	IdentifierExprName           = "IdentifierExpr"
	NumericSuffixLiteralExprName = "NumericSuffixLiteralExpr"
	StringPrefixLiteralExprName  = "StringPrefixLiteralExpr"
	StringSuffixLiteralExprName  = "StringSuffixLiteralExpr"
	StructLiteralExprName        = "StructLiteralExpr"
	ClosureExprName              = "ClosureExpr"
	ParenthesizedExprName        = "ParenthesizedExpr"
	ImplicitMemberExprName       = "ImplicitMemberExpr"
	WildcardExprName             = "WildcardExpr"
	StructConstructionExprName   = "StructConstructionExpr"
	TupleExprName                = "TupleExpr"
	PrefixUnaryExprName          = "PrefixUnaryExpr"
	PostfixUnaryExprName         = "PostfixUnaryExpr"
	FunctionCallExprName         = "FunctionCallExpr"
	ExplicitMemberExprName       = "ExplicitMemberExpr"
	SubscriptExprName            = "SubscriptExpr"
	BinaryExprName               = "BinaryExpr"
	ConditionalExprName          = "ConditionalExpr"
	TypeCastingExprName          = "TypeCastingExpr"
	AssignmentExprName           = "AssignmentExpr"
	ErrorExprName                = "ErrorExpr"
)

var exprInfos map[reflect.Type]StructJsonInfo

func init() {
	exprInfos = GetStructJsonInfos(jsoniter.Config{},
		NullLiteralExpr{},
		IntegerLiteralExpr{},
		FloatLiteralExpr{},
		BoolLiteralExpr{},
		StringLiteralExpr{},
		ObjectLiteralExpr{},
		ArrayLiteralExpr{},
		MapLiteralExpr{},
		RangeLiteralExpr{},
		IdentifierExpr{},
		NumericSuffixLiteralExpr{},
		StringPrefixLiteralExpr{},
		StringSuffixLiteralExpr{},
		StructLiteralExpr{},
		ClosureExpr{},
		ParenthesizedExpr{},
		ImplicitMemberExpr{},
		WildcardExpr{},
		StructConstructionExpr{},
		TupleExpr{},
		PrefixUnaryExpr{},
		PostfixUnaryExpr{},
		FunctionCallExpr{},
		ExplicitMemberExpr{},
		SubscriptExpr{},
		BinaryExpr{},
		ConditionalExpr{},
		TypeCastingExpr{},
		AssignmentExpr{},
		ErrorExpr{},
	)
}

func NewNullLiteralExpression(expr *NullLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_NullLiteralExpr{NullLiteralExpr: expr},
	}
}

func NewNullLiteralExpressionFrom() *Expression {
	return NewNullLiteralExpression(NewNullLiteralExpr())
}

func NewBoolLiteralExpression(expr *BoolLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_BoolLiteralExpr{BoolLiteralExpr: expr},
	}
}

func NewBoolLiteralExpressionFrom(value bool) *Expression {
	return NewBoolLiteralExpression(NewBoolLiteralExpr(value))
}

func NewIntegerLiteralExpression(expr *IntegerLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_IntegerLiteralExpr{IntegerLiteralExpr: expr},
	}
}

func NewIntegerLiteralExpressionFrom(value int64) *Expression {
	return NewIntegerLiteralExpression(NewIntegerLiteralExpr(value))
}

func NewFloatLiteralExpression(expr *FloatLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_FloatLiteralExpr{FloatLiteralExpr: expr},
	}
}

func NewFloatLiteralExpressionFrom(value float64) *Expression {
	return NewFloatLiteralExpression(NewFloatLiteralExpr(value))
}

func NewStringLiteralExpression(expr *StringLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringLiteralExpr{StringLiteralExpr: expr},
	}
}

func NewStringLiteralExpressionFrom(value string) *Expression {
	return NewStringLiteralExpression(NewStringLiteralExpr(value))
}

func NewArrayLiteralExpression(expr *ArrayLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ArrayLiteralExpr{ArrayLiteralExpr: expr},
	}
}

func NewArrayLiteralExpressionFrom(exprs ...*Expression) *Expression {
	return &Expression{
		Expression: &Expression_ArrayLiteralExpr{ArrayLiteralExpr: &ArrayLiteralExpr{Elements: exprs}},
	}
}

func NewMapLiteralExpression(expr *MapLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_MapLiteralExpr{MapLiteralExpr: expr},
	}
}

func NewObjectLiteralExpression(expr *ObjectLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ObjectLiteralExpr{ObjectLiteralExpr: expr},
	}
}

func NewStructLiteralExpression(expr *StructLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StructLiteralExpr{StructLiteralExpr: expr},
	}
}

func NewRangeLiteralExpression(expr *RangeLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_RangeLiteralExpr{RangeLiteralExpr: expr},
	}
}

func NewRangeLiteralExpressionFrom(value *core.IntRange) *Expression {
	return &Expression{
		Expression: &Expression_RangeLiteralExpr{RangeLiteralExpr: NewRangeLiteralExprFrom(value)},
	}
}

func NewIdentifierExpression(expr *IdentifierExpr) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{IdentifierExpr: expr},
	}
}

func NewIdentifierExpressionFrom(names ...string) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{IdentifierExpr: NewIdentifierExpr(names...)},
	}
}

func NewIdentifierExpressionFromType(nominal *NominalType) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{IdentifierExpr: NewIdentifierExprFromType(nominal)},
	}
}

func NewStringPrefixLiteralExpression(expr *StringPrefixLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringPrefixLiteralExpr{StringPrefixLiteralExpr: expr},
	}
}

func NewStringSuffixLiteralExpression(expr *StringSuffixLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringSuffixLiteralExpr{StringSuffixLiteralExpr: expr},
	}
}

func NewNumericSuffixLiteralExpression(expr *NumericSuffixLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_NumericSuffixLiteralExpr{NumericSuffixLiteralExpr: expr},
	}
}

func NewPrefixUnaryExpression(expr *PrefixUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_PrefixUnaryExpr{PrefixUnaryExpr: expr},
	}
}

func NewPostfixUnaryExpression(expr *PostfixUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_PostfixUnaryExpr{PostfixUnaryExpr: expr},
	}
}

func NewFunctionCallExpression(expr *FunctionCallExpr) *Expression {
	return &Expression{
		Expression: &Expression_FunctionCallExpr{FunctionCallExpr: expr},
	}
}

func NewExplicitMemberExpression(expr *ExplicitMemberExpr) *Expression {
	return &Expression{
		Expression: &Expression_ExplicitMemberExpr{ExplicitMemberExpr: expr},
	}
}

func NewSubscriptExpression(expr *SubscriptExpr) *Expression {
	return &Expression{
		Expression: &Expression_SubscriptExpr{SubscriptExpr: expr},
	}
}

func NewConditionalExpression(expr *ConditionalExpr) *Expression {
	return &Expression{
		Expression: &Expression_ConditionalExpr{ConditionalExpr: expr},
	}
}

func NewAssignmentExpression(expr *AssignmentExpr) *Expression {
	return &Expression{
		Expression: &Expression_AssignmentExpr{AssignmentExpr: expr},
	}
}

func NewBinaryExpression(expr *BinaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_BinaryExpr{BinaryExpr: expr},
	}
}

func NewWildcardExpression(expr *WildcardExpr) *Expression {
	return &Expression{
		Expression: &Expression_WildcardExpr{WildcardExpr: expr},
	}
}

func NewParenthesizedExpression(expr *ParenthesizedExpr) *Expression {
	return &Expression{
		Expression: &Expression_ParenthesizedExpr{ParenthesizedExpr: expr},
	}
}

func NewClosureExpression(expr *ClosureExpr) *Expression {
	return &Expression{
		Expression: &Expression_ClosureExpr{ClosureExpr: expr},
	}
}

func NewTupleExpression(expr *TupleExpr) *Expression {
	return &Expression{
		Expression: &Expression_TupleExpr{TupleExpr: expr},
	}
}

func (x *Expression) IsUnion() {
}

func (x *Expression) GetStartPosition() *Position {
	return GetUnionPosition(x, StartPositionFieldName)
}

func (x *Expression) GetEndPosition() *Position {
	return GetUnionPosition(x, EndPositionFieldName)
}

func (x *Expression) SetStartPosition(position *Position) {
	SetUnionPosition(x, StartPositionFieldName, position)
}

func (x *Expression) SetEndPosition(position *Position) {
	SetUnionPosition(x, EndPositionFieldName, position)
}

func (x *Expression) EvalBoolLiteral() (bool, error) {
	valueExpr := x.GetBoolLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return false, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (x *Expression) EvalIntegerLiteral() (int64, error) {
	valueExpr := x.GetIntegerLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalValue(), nil
	}
	return 0, errors.New("expression can NOT be evaluated to the Int64LiteralExpr")
}

func (x *Expression) EvalFloatLiteral() (float64, error) {
	valueExpr := x.GetFloatLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return 0, errors.New("expression can NOT be evaluated to the FloatLiteralExpr")
}

func (x *Expression) EvalStringLiteral() (string, error) {
	valueExpr := x.GetStringLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return "", errors.New("expression can NOT be evaluated to the StringLiteralExpr")
}

func (x *Expression) EvalRangeLiteral() (*core.IntRange, error) {
	valueExpr := x.GetRangeLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return nil, errors.New("expression can NOT be evaluated to the RangeLiteralExpr")
}

func (x *Expression) EvalArrayLiteral() ([]*Expression, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Elements, nil
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalIntegerArrayLiteral() ([]int64, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalIntegerArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalFloatArrayLiteral() ([]float64, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalFloatArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalStringArrayLiteral() ([]string, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalStringArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalStringMapLiteral(iterator func(key string, value *Expression) error) error {
	valueExpr := x.GetMapLiteralExpr()
	if valueExpr != nil {
		for _, entry := range valueExpr.Entries {
			err := iterator(entry.Key, entry.Value)
			if err != nil {
				if core.IsBreakError(err) {
					return nil
				}
				return err
			}
		}
		return nil
	}
	return errors.New("expression can NOT be evaluated to the MapLiteralExpr")
}

func (x *Expression) EvalIdentifier() (string, error) {
	valueExpr := x.GetIdentifierExpr()
	if valueExpr != nil {
		return valueExpr.Identifier.Name, nil
	}
	return "", errors.New("expression can NOT be evaluated to the IdentifierExpr")
}

func (x *Expression) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		value := reflect.ValueOf(x.Expression)
		value = reflect.Indirect(value).Field(0)
		if merger, ok := value.Interface().(CommentMerger); ok {
			return merger.MergeComment(comment)
		} else {
			// error
		}
	}

	return false, errors.New("nil Expression")
}
