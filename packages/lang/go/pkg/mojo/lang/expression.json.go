package lang

import (
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("lang.Expression", &ExpressionCodec{})
	jsoniter.RegisterTypeEncoder("lang.Expression", &ExpressionCodec{})
}

type ExpressionCodec struct {
}

func (codec *ExpressionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	expr := (*Expression)(ptr)
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("@type").ToString()
		switch t {
		case NullLiteralExprName:
			exp := &NullLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_NullLiteralExpr{NullLiteralExpr: exp}
		case IntegerLiteralExprName:
			exp := &IntegerLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_IntegerLiteralExpr{IntegerLiteralExpr: exp}
		case FloatLiteralExprName:
			exp := &FloatLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_FloatLiteralExpr{FloatLiteralExpr: exp}
		case BoolLiteralExprName:
			exp := &BoolLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_BoolLiteralExpr{BoolLiteralExpr: exp}
		case StringLiteralExprName:
			exp := &StringLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_StringLiteralExpr{StringLiteralExpr: exp}
		case ObjectLiteralExprName:
			exp := &ObjectLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ObjectLiteralExpr{ObjectLiteralExpr: exp}
		case ArrayLiteralExprName:
			exp := &ArrayLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ArrayLiteralExpr{ArrayLiteralExpr: exp}
		case MapLiteralExprName:
			exp := &MapLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_MapLiteralExpr{MapLiteralExpr: exp}
		case RangeLiteralExprName:
			exp := &RangeLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_RangeLiteralExpr{RangeLiteralExpr: exp}
		case IdentifierExprName:
			exp := &IdentifierExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_IdentifierExpr{IdentifierExpr: exp}
		case NumericSuffixLiteralExprName:
			exp := &NumericSuffixLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_NumericSuffixLiteralExpr{NumericSuffixLiteralExpr: exp}
		case StringPrefixLiteralExprName:
			exp := &StringPrefixLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_StringPrefixLiteralExpr{StringPrefixLiteralExpr: exp}
		case StringSuffixLiteralExprName:
			exp := &StringSuffixLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_StringSuffixLiteralExpr{StringSuffixLiteralExpr: exp}
		case StructLiteralExprName:
			exp := &StructLiteralExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_StructLiteralExpr{StructLiteralExpr: exp}
		case ClosureExprName:
			exp := &ClosureExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ClosureExpr{ClosureExpr: exp}
		case ParenthesizedExprName:
			exp := &ParenthesizedExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ParenthesizedExpr{ParenthesizedExpr: exp}
		case ImplicitMemberExprName:
			exp := &ImplicitMemberExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ImplicitMemberExpr{ImplicitMemberExpr: exp}
		case WildcardExprName:
			exp := &WildcardExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_WildcardExpr{WildcardExpr: exp}
		case StructConstructionExprName:
			exp := &StructConstructionExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_StructConstructionExpr{StructConstructionExpr: exp}
		case TupleExprName:
			exp := &TupleExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_TupleExpr{TupleExpr: exp}
		case PrefixUnaryExprName:
			exp := &PrefixUnaryExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_PrefixUnaryExpr{PrefixUnaryExpr: exp}
		case PostfixUnaryExprName:
			exp := &PostfixUnaryExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_PostfixUnaryExpr{PostfixUnaryExpr: exp}
		case FunctionCallExprName:
			exp := &FunctionCallExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_FunctionCallExpr{FunctionCallExpr: exp}
		case ExplicitMemberExprName:
			exp := &ExplicitMemberExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ExplicitMemberExpr{ExplicitMemberExpr: exp}
		case SubscriptExprName:
			exp := &SubscriptExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_SubscriptExpr{SubscriptExpr: exp}
		case BinaryExprName:
			exp := &BinaryExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_BinaryExpr{BinaryExpr: exp}
		case ConditionalExprName:
			exp := &ConditionalExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ConditionalExpr{ConditionalExpr: exp}
		case TypeCastingExprName:
			exp := &TypeCastingExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_TypeCastingExpr{TypeCastingExpr: exp}
		case AssignmentExprName:
			exp := &AssignmentExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_AssignmentExpr{AssignmentExpr: exp}
		case ErrorExprName:
			exp := &ErrorExpr{}
			a.ToVal(exp)
			expr.Expression = &Expression_ErrorExpr{ErrorExpr: exp}
		}
	}
}

func (codec *ExpressionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	expr := (*Expression)(ptr)
	if expr != nil && expr.Expression != nil {
		stream.WriteObjectStart()

		typ := ""
		var v reflect.Value
		switch exp := expr.Expression.(type) {
		case *Expression_NullLiteralExpr:
			v = reflect.ValueOf(exp.NullLiteralExpr).Elem()
			typ = NullLiteralExprName
		case *Expression_IntegerLiteralExpr:
			v = reflect.ValueOf(exp.IntegerLiteralExpr).Elem()
			typ = IntegerLiteralExprName
		case *Expression_FloatLiteralExpr:
			v = reflect.ValueOf(exp.FloatLiteralExpr).Elem()
			typ = FloatLiteralExprName
		case *Expression_StringLiteralExpr:
			v = reflect.ValueOf(exp.StringLiteralExpr).Elem()
			typ = StringLiteralExprName
		case *Expression_ObjectLiteralExpr:
			v = reflect.ValueOf(exp.ObjectLiteralExpr).Elem()
			typ = ObjectLiteralExprName
		case *Expression_ArrayLiteralExpr:
			v = reflect.ValueOf(exp.ArrayLiteralExpr).Elem()
			typ = ArrayLiteralExprName
		case *Expression_MapLiteralExpr:
			v = reflect.ValueOf(exp.MapLiteralExpr).Elem()
			typ = MapLiteralExprName
		case *Expression_RangeLiteralExpr:
			v = reflect.ValueOf(exp.RangeLiteralExpr).Elem()
			typ = RangeLiteralExprName
		case *Expression_IdentifierExpr:
			v = reflect.ValueOf(exp.IdentifierExpr).Elem()
			typ = IdentifierExprName
		case *Expression_NumericSuffixLiteralExpr:
			v = reflect.ValueOf(exp.NumericSuffixLiteralExpr).Elem()
			typ = NumericSuffixLiteralExprName
		case *Expression_StringPrefixLiteralExpr:
			v = reflect.ValueOf(exp.StringPrefixLiteralExpr).Elem()
			typ = StringPrefixLiteralExprName
		case *Expression_StringSuffixLiteralExpr:
			v = reflect.ValueOf(exp.StringSuffixLiteralExpr).Elem()
			typ = StringSuffixLiteralExprName
		case *Expression_StructLiteralExpr:
			v = reflect.ValueOf(exp.StructLiteralExpr).Elem()
			typ = StructLiteralExprName
		case *Expression_ClosureExpr:
			v = reflect.ValueOf(exp.ClosureExpr).Elem()
			typ = ClosureExprName
		case *Expression_ParenthesizedExpr:
			v = reflect.ValueOf(exp.ParenthesizedExpr).Elem()
			typ = ParenthesizedExprName
		case *Expression_ImplicitMemberExpr:
			v = reflect.ValueOf(exp.ImplicitMemberExpr).Elem()
			typ = ImplicitMemberExprName
		case *Expression_WildcardExpr:
			v = reflect.ValueOf(exp.WildcardExpr).Elem()
			typ = WildcardExprName
		case *Expression_StructConstructionExpr:
			v = reflect.ValueOf(exp.StructConstructionExpr).Elem()
			typ = StructConstructionExprName
		case *Expression_TupleExpr:
			v = reflect.ValueOf(exp.TupleExpr).Elem()
			typ = TupleExprName
		case *Expression_PrefixUnaryExpr:
			v = reflect.ValueOf(exp.PrefixUnaryExpr).Elem()
			typ = PrefixUnaryExprName
		case *Expression_PostfixUnaryExpr:
			v = reflect.ValueOf(exp.PostfixUnaryExpr).Elem()
			typ = PostfixUnaryExprName
		case *Expression_FunctionCallExpr:
			v = reflect.ValueOf(exp.FunctionCallExpr).Elem()
			typ = FunctionCallExprName
		case *Expression_ExplicitMemberExpr:
			v = reflect.ValueOf(exp.ExplicitMemberExpr).Elem()
			typ = ExplicitMemberExprName
		case *Expression_SubscriptExpr:
			v = reflect.ValueOf(exp.SubscriptExpr).Elem()
			typ = SubscriptExprName
		case *Expression_BinaryExpr:
			v = reflect.ValueOf(exp.BinaryExpr).Elem()
			typ = BinaryExprName
		case *Expression_ConditionalExpr:
			v = reflect.ValueOf(exp.ConditionalExpr).Elem()
			typ = ConditionalExprName
		case *Expression_TypeCastingExpr:
			v = reflect.ValueOf(exp.TypeCastingExpr).Elem()
			typ = TypeCastingExprName
		case *Expression_AssignmentExpr:
			v = reflect.ValueOf(exp.AssignmentExpr).Elem()
			typ = AssignmentExprName
		case *Expression_ErrorExpr:
			v = reflect.ValueOf(exp.ErrorExpr).Elem()
			typ = ErrorExprName
		}

		if v.IsValid() && !v.IsZero() {
			stream.WriteObjectField("@type")
			stream.WriteVal(typ)

			t := exprInfos[v.Type()]

			// iterator the fields
			for i := 0; i < v.NumField(); i++ {
				info := t.Fields[i]
				field := v.Field(i)

				if info.Ignored {
					continue
				}

				if info.OmitEmpty {
					if field.IsZero() {
						continue
					}

					k := field.Kind()
					switch k {
					case reflect.Map, reflect.Ptr, reflect.Interface, reflect.Struct, reflect.Slice:
						if info.Encoder.IsEmpty(unsafe.Pointer(field.Pointer())) {
							continue
						}
					}
				}

				stream.WriteMore()
				stream.WriteObjectField(t.Fields[i].ToName)
				stream.WriteVal(v.Field(i).Interface())
			}
		}

		stream.WriteObjectEnd()

		if stream.Error != nil && stream.Error != io.EOF {
			// stream.Error = fmt.Errorf("%v.%s", encoder.typ, stream.Error.Error())
		}
	}
}

func (codec *ExpressionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Expression)(ptr)).Expression == nil
}
