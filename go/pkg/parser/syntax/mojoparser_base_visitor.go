// Code generated from MojoParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package syntax // MojoParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMojoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMojoParserVisitor) VisitMojoFile(ctx *MojoFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFreeDocument(ctx *FreeDocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitForInStatement(ctx *ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitConditions(ctx *ConditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitOptionalBindingCondition(ctx *OptionalBindingConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBranchStatement(ctx *BranchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitElseClause(ctx *ElseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMatchStatement(ctx *MatchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMatchCases(ctx *MatchCasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMatchCase(ctx *MatchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitControlTransferStatement(ctx *ControlTransferStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitGenericParameterClause(ctx *GenericParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitGenericParameters(ctx *GenericParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitGenericParameter(ctx *GenericParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitGenericArgumentClause(ctx *GenericArgumentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitGenericArguments(ctx *GenericArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitGenericArgument(ctx *GenericArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitCodeBlock(ctx *CodeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPackageIdentifier(ctx *PackageIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPackageName(ctx *PackageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportPathIdentifier(ctx *ImportPathIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportAllClause(ctx *ImportAllClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportValueAsClause(ctx *ImportValueAsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportTypeClause(ctx *ImportTypeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportTypeAsClause(ctx *ImportTypeAsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportGroupClause(ctx *ImportGroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportGroup(ctx *ImportGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportValue(ctx *ImportValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImportType(ctx *ImportTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitConstantDeclaration(ctx *ConstantDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPatternInitializers(ctx *PatternInitializersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPatternInitializer(ctx *PatternInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeAliasName(ctx *TypeAliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeAliasAssignment(ctx *TypeAliasAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionHead(ctx *FunctionHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionSignature(ctx *FunctionSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionResult(ctx *FunctionResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionParameterClause(ctx *FunctionParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionParameters(ctx *FunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEnumName(ctx *EnumNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEnumMembers(ctx *EnumMembersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEnumMember(ctx *EnumMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructName(ctx *StructNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructMembers(ctx *StructMembersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructMemberDeclaration(ctx *StructMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInterfaceName(ctx *InterfaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInterfaceMembers(ctx *InterfaceMembersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInterfaceMember(ctx *InterfaceMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributeDeclaration(ctx *AttributeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitWildcard_pattern(ctx *Wildcard_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIdentifierPattern(ctx *IdentifierPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTuple_pattern(ctx *Tuple_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTuple_pattern_element_list(ctx *Tuple_pattern_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTuple_pattern_element(ctx *Tuple_pattern_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitOptional_pattern(ctx *Optional_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExpression_pattern(ctx *Expression_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributeIdentifier(ctx *AttributeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributeName(ctx *AttributeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributes(ctx *AttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBinaryExpressions(ctx *BinaryExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitConditional_operator(ctx *Conditional_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitType_casting_operator(ctx *Type_casting_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayLiteralItems(ctx *ArrayLiteralItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayLiteralItem(ctx *ArrayLiteralItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitObjectLiteralItems(ctx *ObjectLiteralItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitObjectLiteralItem(ctx *ObjectLiteralItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImplicitMemberExpression(ctx *ImplicitMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleElement(ctx *TupleElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitWildcardExpression(ctx *WildcardExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExplicitMemberExpression1(ctx *ExplicitMemberExpression1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPostfixOperation(ctx *PostfixOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExplicitMemberExpression4(ctx *ExplicitMemberExpression4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitSubscriptExpression(ctx *SubscriptExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExplicitMemberExpression3(ctx *ExplicitMemberExpression3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExplicitMemberExpression2(ctx *ExplicitMemberExpression2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionCallArgumentClause(ctx *FunctionCallArgumentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunction_call_argument_list(ctx *Function_call_argument_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunction_call_argument(ctx *Function_call_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArgumentNameList(ctx *ArgumentNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArgument_name(ctx *Argument_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIntersection(ctx *IntersectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrime(ctx *PrimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrimeType(ctx *PrimeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeIdentifierClause(ctx *TypeIdentifierClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleTypeElements(ctx *TupleTypeElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleTypeElement(ctx *TupleTypeElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionTypeArgumentClause(ctx *FunctionTypeArgumentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionTypeArguments(ctx *FunctionTypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionTypeArgument(ctx *FunctionTypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitDictionaryType(ctx *DictionaryTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeInheritanceClause(ctx *TypeInheritanceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeInheritances(ctx *TypeInheritancesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeInheritance(ctx *TypeInheritanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPath_identifier(ctx *Path_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitKeyword_as_identifier_in_declarations(ctx *Keyword_as_identifier_in_declarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitKeyword_as_identifier_in_labels(ctx *Keyword_as_identifier_in_labelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFollowingDocument(ctx *FollowingDocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitNegatePrefixOperator(ctx *NegatePrefixOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitCompilation_condition_AND(ctx *Compilation_condition_ANDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitCompilation_condition_OR(ctx *Compilation_condition_ORContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitCompilation_condition_GE(ctx *Compilation_condition_GEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrowOperator(ctx *ArrowOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitRange_operator(ctx *Range_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitHalf_open_range_operator(ctx *Half_open_range_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitSame_type_equals(ctx *Same_type_equalsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrefixOperator(ctx *PrefixOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPostfixOperator(ctx *PostfixOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitOperator_character(ctx *Operator_characterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitOperator_head(ctx *Operator_headContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitDot_operator_head(ctx *Dot_operator_headContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitDot_operator_character(ctx *Dot_operator_characterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEov(ctx *EovContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEosWithDocument(ctx *EosWithDocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEovWithDocument(ctx *EovWithDocumentContext) interface{} {
	return v.VisitChildren(ctx)
}
