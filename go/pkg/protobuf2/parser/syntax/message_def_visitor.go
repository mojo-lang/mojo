package syntax

import (
	"strconv"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type MessageDefVisitor struct {
	*BaseProtobuf2Visitor
}

func NewMessageDefVisitor() *MessageDefVisitor {
	visitor := &MessageDefVisitor{}
	return visitor
}

func (v *MessageDefVisitor) VisitMessageDef(ctx *MessageDefContext) interface{} {
	if name := ctx.MessageName(); name != nil {
		if body := ctx.MessageBody(); body != nil {
			if decl, ok := body.Accept(v).(*lang.StructDecl); ok {
				decl.Name = name.GetText()
				return decl
			}
		}
	}
	return nil
}

func (v *MessageDefVisitor) VisitMessageBody(ctx *MessageBodyContext) interface{} {
	elements := ctx.AllMessageElement()

	decl := &lang.StructDecl{Type: &lang.StructType{}}
	for _, element := range elements {
		ele := element.Accept(v)
		if valueDecl, ok := ele.(*lang.ValueDecl); ok {
			decl.Type.Fields = append(decl.Type.Fields, valueDecl)
		}
		if enum, ok := ele.(*lang.EnumDecl); ok {
			decl.EnumDecls = append(decl.EnumDecls, enum)
		}
		if message, ok := ele.(*lang.StructDecl); ok {
			decl.StructDecls = append(decl.StructDecls, message)
		}
		if attr, ok := ele.(*lang.Attribute); ok {
			decl.Attributes = append(decl.Attributes, attr)
		}
	}

	return decl
}

func (v *MessageDefVisitor) VisitMessageElement(ctx *MessageElementContext) interface{} {
	if field := ctx.Field(); field != nil {
		return field.Accept(v)
	}
	if mapField := ctx.MapField(); mapField != nil {
		return mapField.Accept(v)
	}
	if oneof := ctx.Oneof(); oneof != nil {
		return oneof.Accept(v)
	}
	if reserved := ctx.Reserved(); reserved != nil {
		return reserved.Accept(v)
	}
	if enum := ctx.EnumDef(); enum != nil {
		return enum.Accept(NewEnumDefVisitor())
	}
	if message := ctx.MessageDef(); message != nil {
		return message.Accept(NewMessageDefVisitor())
	}
	if extend := ctx.ExtendDef(); extend != nil {
		return extend.Accept(NewExtendDefVisitor())
	}
	if extensions := ctx.Extensions(); extensions != nil {
		return extensions.Accept(v)
	}
	if option := ctx.OptionStatement(); option != nil {
		return option.Accept(NewOptionStatementVisitor())
	}
	return nil
}

func (v *MessageDefVisitor) VisitField(ctx *FieldContext) interface{} {
	if typ := ctx.Type_(); typ != nil {
		decl := &lang.ValueDecl{
			Name: ctx.FieldName().GetText(),
			Type: &lang.NominalType{},
		}

		if t, ok := typ.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
			decl.Type = t
		}

		if label := ctx.FieldLabel().GetText(); len(label) > 0 {
			switch label {
			case "repeated":
				decl.Type = lang.NewArrayNominalType(decl.Type)
			case "optional":
			case "required":
				decl.Type.Attributes = append(decl.Type.Attributes, lang.NewBoolAttribute("protobuf", "required"))
			}
		}

		number := ctx.FieldNumber().GetText()
		value, _ := strconv.ParseInt(number, 10, 64)
		decl.Type.Attributes = append(decl.Type.Attributes, lang.NewIntegerAttribute("protobuf", "number", value))

		if options := ctx.FieldOptions(); options != nil {
			if attrs, ok := options.Accept(v).([]*lang.Attribute); ok {
				decl.Type.Attributes = append(decl.Type.Attributes, attrs...)
			}
		}
		return decl
	}
	return nil
}

func (v *MessageDefVisitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	allOptions := ctx.AllFieldOption()
	var attrs []*lang.Attribute
	for _, option := range allOptions {
		if attr, ok := option.Accept(v).(*lang.Attribute); ok {
			attrs = append(attrs, attr)
		}
	}
	return attrs
}

func (v *MessageDefVisitor) VisitFieldOption(ctx *FieldOptionContext) interface{} {
	if name := ctx.OptionName(); name != nil {
		if attribute, ok := name.Accept(NewOptionStatementVisitor()).(*lang.Attribute); ok {
			if constant := ctx.Constant(); constant != nil {
				if expr, ok := constant.Accept(NewConstantVisitor()).(*lang.Expression); ok {
					attribute.Arguments = append(attribute.Arguments, &lang.Argument{
						Value: expr,
					})

					return attribute
				}
			}
		}
	}
	return nil
}

func (v *MessageDefVisitor) VisitOneof(ctx *OneofContext) interface{} {
	oneof := &lang.ValueDecl{Name: ctx.OneofName().GetText()}

	allFields := ctx.AllOneofField()
	var types []*lang.NominalType
	for _, field := range allFields {
		if decl, ok := field.Accept(v).(*lang.ValueDecl); ok {
			decl.Type.Attributes = append(decl.Type.Attributes, lang.NewStringAttribute("", "label", decl.Name))
			types = append(types, decl.Type)
		}
	}
	oneof.Type = lang.NewUnionNominalType(types...)

	allOptions := ctx.AllOptionStatement()
	for _, option := range allOptions {
		if attr, ok := option.Accept(NewOptionStatementVisitor()).(*lang.Attribute); ok {
			oneof.Type.Attributes = append(oneof.Type.Attributes, attr)
		}
	}

	return oneof
}

func (v *MessageDefVisitor) VisitOneofField(ctx *OneofFieldContext) interface{} {
	if typ := ctx.Type_(); typ != nil {
		decl := &lang.ValueDecl{
			Name: ctx.FieldName().GetText(),
		}

		if t, ok := typ.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
			decl.Type = t
		}

		number := ctx.FieldNumber().GetText()
		value, _ := strconv.ParseInt(number, 10, 64)
		decl.Type.Attributes = append(decl.Type.Attributes, lang.NewIntegerAttribute("protobuf", "number", value))

		if options := ctx.FieldOptions(); options != nil {
			if attrs, ok := options.Accept(v).([]*lang.Attribute); ok {
				decl.Type.Attributes = append(decl.Type.Attributes, attrs...)
			}
		}
		return decl
	}
	return nil
}

func (v *MessageDefVisitor) VisitMapField(ctx *MapFieldContext) interface{} {
	if typ := ctx.Type_(); typ != nil {
		decl := &lang.ValueDecl{
			Name: ctx.MapName().GetText(),
			Type: &lang.NominalType{},
		}

		if keyType, ok := typ.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
			if valueType, ok := typ.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
				decl.Type = lang.NewMapNominalType(keyType, valueType)
			}
		}

		number := ctx.FieldNumber().GetText()
		value, _ := strconv.ParseInt(number, 10, 64)
		decl.Type.Attributes = append(decl.Type.Attributes, lang.NewIntegerAttribute("protobuf", "number", value))

		if options := ctx.FieldOptions(); options != nil {
			if attrs, ok := options.Accept(v).([]*lang.Attribute); ok {
				decl.Type.Attributes = append(decl.Type.Attributes, attrs...)
			}
		}
		return decl
	}
	return nil
}

func (v *MessageDefVisitor) VisitReserved(ctx *ReservedContext) interface{} {
	if ranges := ctx.Ranges(); ranges != nil {
		if exprs, ok := ranges.Accept(v).([]*lang.Expression); ok {
			return &lang.Attribute{
				PackageName: "protobuf",
				Name:        "reserved",
				Arguments:   lang.NewArguments(exprs...),
			}
		}
	}
	if fieldNames := ctx.ReservedFieldNames(); fieldNames != nil {
		if exprs, ok := fieldNames.Accept(v).([]*lang.Expression); ok {
			return &lang.Attribute{
				PackageName: "protobuf",
				Name:        "reserved",
				Arguments:   lang.NewArguments(exprs...),
			}
		}
	}
	return nil
}

func (v *MessageDefVisitor) VisitRanges(ctx *RangesContext) interface{} {
	var exprs []*lang.Expression
	ranges := ctx.AllRange_()
	for _, rng := range ranges {
		if expr, ok := rng.Accept(v).(*lang.Expression); ok {
			exprs = append(exprs, expr)
		}
	}
	return exprs
}

func (v *MessageDefVisitor) VisitRange_(ctx *Range_Context) interface{} {
	lits := ctx.AllIntLit()
	if len(lits) == 1 {
		if expr, ok := lits[0].Accept(NewConstantVisitor()).(*lang.IntegerLiteralExpr); ok {
			if ctx.MAX() != nil {
				return lang.NewRangeLiteralExpressionFrom(&core.IntRange{Start: expr.EvalValue()})
			} else {
				return lang.NewIntegerLiteralExpression(expr)
			}
		}
	} else if len(lits) > 1 {
		if start, ok := lits[0].Accept(NewConstantVisitor()).(*lang.IntegerLiteralExpr); ok {
			if end, ok := lits[1].Accept(NewConstantVisitor()).(*lang.IntegerLiteralExpr); ok {
				return lang.NewRangeLiteralExpressionFrom(&core.IntRange{
					Start:       start.EvalValue(),
					End:         end.EvalValue(),
					EndIncluded: true,
				})
			}
		}
	}
	return nil
}

func (v *MessageDefVisitor) VisitReservedFieldNames(ctx *ReservedFieldNamesContext) interface{} {
	var exprs []*lang.Expression
	lits := ctx.AllStrLit()
	for _, lit := range lits {
		if expr, ok := lit.Accept(NewConstantVisitor()).(*lang.StringLiteralExpr); ok {
			exprs = append(exprs, lang.NewStringLiteralExpression(expr))
		}
	}
	return exprs
}

func (v *MessageDefVisitor) VisitExtensions(ctx *ExtensionsContext) interface{} {
	if ranges := ctx.Ranges(); ranges != nil {
		if exprs, ok := ranges.Accept(v).([]*lang.Expression); ok {
			return &lang.Attribute{
				PackageName: "protobuf",
				Name:        "extensions",
				Arguments:   lang.NewArguments(exprs...),
			}
		}
	}

	return nil
}
