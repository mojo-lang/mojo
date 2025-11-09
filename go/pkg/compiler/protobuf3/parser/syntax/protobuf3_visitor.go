// Code generated from Protobuf3.g4 by ANTLR 4.13.0. DO NOT EDIT.

package syntax // Protobuf3
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Protobuf3Parser.
type Protobuf3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Protobuf3Parser#proto.
	VisitProto(ctx *ProtoContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#syntax.
	VisitSyntax(ctx *SyntaxContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#optionStatement.
	VisitOptionStatement(ctx *OptionStatementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#optionName.
	VisitOptionName(ctx *OptionNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldOption.
	VisitFieldOption(ctx *FieldOptionContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldNumber.
	VisitFieldNumber(ctx *FieldNumberContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#oneof.
	VisitOneof(ctx *OneofContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#oneofField.
	VisitOneofField(ctx *OneofFieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#mapField.
	VisitMapField(ctx *MapFieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#keyType.
	VisitKeyType(ctx *KeyTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by Protobuf3Parser#reserved.
	VisitReserved(ctx *ReservedContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#ranges.
	VisitRanges(ctx *RangesContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#range_.
	VisitRange_(ctx *Range_Context) interface{}

	// Visit a parse tree produced by Protobuf3Parser#reservedFieldNames.
	VisitReservedFieldNames(ctx *ReservedFieldNamesContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#topLevelDef.
	VisitTopLevelDef(ctx *TopLevelDefContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumDef.
	VisitEnumDef(ctx *EnumDefContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumElement.
	VisitEnumElement(ctx *EnumElementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumValueOptions.
	VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumValueOption.
	VisitEnumValueOption(ctx *EnumValueOptionContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageDef.
	VisitMessageDef(ctx *MessageDefContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageBody.
	VisitMessageBody(ctx *MessageBodyContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageElement.
	VisitMessageElement(ctx *MessageElementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#serviceDef.
	VisitServiceDef(ctx *ServiceDefContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#serviceElement.
	VisitServiceElement(ctx *ServiceElementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#rpc.
	VisitRpc(ctx *RpcContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#requestStream.
	VisitRequestStream(ctx *RequestStreamContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#responseStream.
	VisitResponseStream(ctx *ResponseStreamContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#blockLit.
	VisitBlockLit(ctx *BlockLitContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by Protobuf3Parser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fullIdent.
	VisitFullIdent(ctx *FullIdentContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageName.
	VisitMessageName(ctx *MessageNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#oneofName.
	VisitOneofName(ctx *OneofNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#mapName.
	VisitMapName(ctx *MapNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#rpcName.
	VisitRpcName(ctx *RpcNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageType.
	VisitMessageType(ctx *MessageTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumType.
	VisitEnumType(ctx *EnumTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#intLit.
	VisitIntLit(ctx *IntLitContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#strLit.
	VisitStrLit(ctx *StrLitContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#boolLit.
	VisitBoolLit(ctx *BoolLitContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#floatLit.
	VisitFloatLit(ctx *FloatLitContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#keywords.
	VisitKeywords(ctx *KeywordsContext) interface{}
}
