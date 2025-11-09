package lang

import (
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("lang.Statement", &StatementCodec{})
	jsoniter.RegisterTypeEncoder("lang.Statement", &StatementCodec{})
}

type StatementCodec struct {
}

func (codec *StatementCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	stmt := (*Statement)(ptr)
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("@type").ToString()
		switch t {
		case PackageDeclName, ImportDeclName, EnumDeclName, StructDeclName, InterfaceDeclName, TypeAliasDeclName,
			ConstantDeclName, VariableDeclName, AttributeDeclName, AttributeAliasDeclName, FunctionDeclName,
			ConstructorDeclName, GenericParameterName:
			decl := &Declaration{}
			a.ToVal(decl)
			stmt.Statement = &Statement_Declaration{Declaration: decl}
		case NullLiteralExprName:
			expr := &Expression{}
			a.ToVal(expr)
			stmt.Statement = &Statement_Expression{Expression: expr}
		case ReturnStmtName:
			s := &ReturnStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_ReturnStmt{ReturnStmt: s}
		case BreakStmtName:
			s := &BreakStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_BreakStmt{BreakStmt: s}
		case ContinueStmtName:
			s := &ContinueStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_ContinueStmt{ContinueStmt: s}
		case MatchStmtName:
			s := &MatchStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_MatchStmt{MatchStmt: s}
		case IfStmtName:
			s := &IfStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_IfStmt{IfStmt: s}
		case ForStmtName:
			s := &ForStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_ForStmt{ForStmt: s}
		case WhileStmtName:
			s := &WhileStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_WhileStmt{WhileStmt: s}
		case RepeatStmtName:
			s := &RepeatStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_RepeatStmt{RepeatStmt: s}
		case BlockStmtName:
			s := &BlockStmt{}
			a.ToVal(s)
			stmt.Statement = &Statement_BlockStmt{BlockStmt: s}
		}
	}
}

func (codec *StatementCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stmt := (*Statement)(ptr)
	if stmt != nil && stmt.Statement != nil {
		if decl, ok := stmt.Statement.(*Statement_Declaration); ok {
			(&DeclarationCodec{}).Encode(unsafe.Pointer(decl.Declaration), stream)
		} else if expr, ok := stmt.Statement.(*Statement_Expression); ok {
			(&ExpressionCodec{}).Encode(unsafe.Pointer(expr.Expression), stream)
		} else {
			stream.WriteObjectStart()

			typ := ""
			var v reflect.Value
			switch decl := stmt.Statement.(type) {
			case *Statement_BlockStmt:
				v = reflect.ValueOf(decl.BlockStmt).Elem()
				typ = BlockStmtName
			case *Statement_BreakStmt:
				v = reflect.ValueOf(decl.BreakStmt).Elem()
				typ = BreakStmtName
			case *Statement_ContinueStmt:
				v = reflect.ValueOf(decl.ContinueStmt).Elem()
				typ = ContinueStmtName
			case *Statement_ForStmt:
				v = reflect.ValueOf(decl.ForStmt).Elem()
				typ = ForStmtName
			case *Statement_IfStmt:
				v = reflect.ValueOf(decl.IfStmt).Elem()
				typ = IfStmtName
			case *Statement_WhileStmt:
				v = reflect.ValueOf(decl.WhileStmt).Elem()
				typ = WhileStmtName
			case *Statement_RepeatStmt:
				v = reflect.ValueOf(decl.RepeatStmt).Elem()
				typ = RepeatStmtName
			case *Statement_MatchStmt:
				v = reflect.ValueOf(decl.MatchStmt).Elem()
				typ = MatchStmtName
			case *Statement_ReturnStmt:
				v = reflect.ValueOf(decl.ReturnStmt).Elem()
				typ = ReturnStmtName
			}

			if v.IsValid() && !v.IsZero() {
				stream.WriteObjectField("@type")
				stream.WriteVal(typ)

				t := declInfos[v.Type()]

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
}

func (codec *StatementCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Statement)(ptr)).Statement == nil
}
