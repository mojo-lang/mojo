package lang

import (
	"reflect"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
)

func (x *ObjectLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *ObjectLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *ObjectLiteralExpr) To(object interface{}) error {
	if x != nil && object != nil {
		ot := reflect.TypeOf(object)
		ov := reflect.ValueOf(object).Elem()
		for _, field := range x.Fields {
			fieldName := strcase.ToCamel(field.Name)
			if f, found := ot.FieldByName(fieldName); found {
				v := ov.FieldByIndex(f.Index)
				switch v.Type().Kind() {
				case reflect.String:
					v.SetString(field.Value.GetStringLiteralExpr().GetValue())
				case reflect.Struct:

				}
			}
		}
	}

	return nil
}

func (x *ObjectLiteralExpr) From(object interface{}) error {
	return nil
}
