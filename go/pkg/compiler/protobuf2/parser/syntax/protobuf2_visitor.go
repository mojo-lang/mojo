// Code generated from Protobuf2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package syntax // Protobuf2
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Protobuf2Parser.
type Protobuf2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Protobuf2Parser#proto.
	VisitProto(ctx *ProtoContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#syntax.
	VisitSyntax(ctx *SyntaxContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#optionStatement.
	VisitOptionStatement(ctx *OptionStatementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#optionName.
	VisitOptionName(ctx *OptionNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#fieldLabel.
	VisitFieldLabel(ctx *FieldLabelContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#fieldOption.
	VisitFieldOption(ctx *FieldOptionContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#fieldNumber.
	VisitFieldNumber(ctx *FieldNumberContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#oneof.
	VisitOneof(ctx *OneofContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#oneofField.
	VisitOneofField(ctx *OneofFieldContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#mapField.
	VisitMapField(ctx *MapFieldContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#keyType.
	VisitKeyType(ctx *KeyTypeContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by Protobuf2Parser#extensions.
	VisitExtensions(ctx *ExtensionsContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#reserved.
	VisitReserved(ctx *ReservedContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#ranges.
	VisitRanges(ctx *RangesContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#range_.
	VisitRange_(ctx *Range_Context) interface{}

	// Visit a parse tree produced by Protobuf2Parser#reservedFieldNames.
	VisitReservedFieldNames(ctx *ReservedFieldNamesContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#topLevelDef.
	VisitTopLevelDef(ctx *TopLevelDefContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumDef.
	VisitEnumDef(ctx *EnumDefContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumElement.
	VisitEnumElement(ctx *EnumElementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumValueOptions.
	VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumValueOption.
	VisitEnumValueOption(ctx *EnumValueOptionContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#messageDef.
	VisitMessageDef(ctx *MessageDefContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#messageBody.
	VisitMessageBody(ctx *MessageBodyContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#messageElement.
	VisitMessageElement(ctx *MessageElementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#extendDef.
	VisitExtendDef(ctx *ExtendDefContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#extendElement.
	VisitExtendElement(ctx *ExtendElementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#serviceDef.
	VisitServiceDef(ctx *ServiceDefContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#serviceElement.
	VisitServiceElement(ctx *ServiceElementContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#rpc.
	VisitRpc(ctx *RpcContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#stream.
	VisitStream(ctx *StreamContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#blockLit.
	VisitBlockLit(ctx *BlockLitContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by Protobuf2Parser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#fullIdent.
	VisitFullIdent(ctx *FullIdentContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#messageName.
	VisitMessageName(ctx *MessageNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#groupName.
	VisitGroupName(ctx *GroupNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#oneofName.
	VisitOneofName(ctx *OneofNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#mapName.
	VisitMapName(ctx *MapNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#rpcName.
	VisitRpcName(ctx *RpcNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#streamName.
	VisitStreamName(ctx *StreamNameContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#messageType.
	VisitMessageType(ctx *MessageTypeContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#enumType.
	VisitEnumType(ctx *EnumTypeContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#intLit.
	VisitIntLit(ctx *IntLitContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#strLit.
	VisitStrLit(ctx *StrLitContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#boolLit.
	VisitBoolLit(ctx *BoolLitContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#floatLit.
	VisitFloatLit(ctx *FloatLitContext) interface{}

	// Visit a parse tree produced by Protobuf2Parser#keywords.
	VisitKeywords(ctx *KeywordsContext) interface{}
}
