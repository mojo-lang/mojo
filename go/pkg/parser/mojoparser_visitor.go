// Code generated from MojoParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MojoParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MojoParser.
type MojoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MojoParser#mojoFile.
	VisitMojoFile(ctx *MojoFileContext) interface{}

	// Visit a parse tree produced by MojoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MojoParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by MojoParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#forInStatement.
	VisitForInStatement(ctx *ForInStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#conditions.
	VisitConditions(ctx *ConditionsContext) interface{}

	// Visit a parse tree produced by MojoParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by MojoParser#optionalBindingCondition.
	VisitOptionalBindingCondition(ctx *OptionalBindingConditionContext) interface{}

	// Visit a parse tree produced by MojoParser#branchStatement.
	VisitBranchStatement(ctx *BranchStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#elseClause.
	VisitElseClause(ctx *ElseClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#matchStatement.
	VisitMatchStatement(ctx *MatchStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#matchCases.
	VisitMatchCases(ctx *MatchCasesContext) interface{}

	// Visit a parse tree produced by MojoParser#matchCase.
	VisitMatchCase(ctx *MatchCaseContext) interface{}

	// Visit a parse tree produced by MojoParser#controlTransferStatement.
	VisitControlTransferStatement(ctx *ControlTransferStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MojoParser#genericParameterClause.
	VisitGenericParameterClause(ctx *GenericParameterClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#genericParameters.
	VisitGenericParameters(ctx *GenericParametersContext) interface{}

	// Visit a parse tree produced by MojoParser#genericParameter.
	VisitGenericParameter(ctx *GenericParameterContext) interface{}

	// Visit a parse tree produced by MojoParser#genericArgumentClause.
	VisitGenericArgumentClause(ctx *GenericArgumentClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#genericArguments.
	VisitGenericArguments(ctx *GenericArgumentsContext) interface{}

	// Visit a parse tree produced by MojoParser#genericArgument.
	VisitGenericArgument(ctx *GenericArgumentContext) interface{}

	// Visit a parse tree produced by MojoParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#codeBlock.
	VisitCodeBlock(ctx *CodeBlockContext) interface{}

	// Visit a parse tree produced by MojoParser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#packageIdentifier.
	VisitPackageIdentifier(ctx *PackageIdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by MojoParser#importPathIdentifier.
	VisitImportPathIdentifier(ctx *ImportPathIdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#importAllClause.
	VisitImportAllClause(ctx *ImportAllClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#importValueAsClause.
	VisitImportValueAsClause(ctx *ImportValueAsClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#importTypeClause.
	VisitImportTypeClause(ctx *ImportTypeClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#importTypeAsClause.
	VisitImportTypeAsClause(ctx *ImportTypeAsClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#importGroupClause.
	VisitImportGroupClause(ctx *ImportGroupClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#importGroup.
	VisitImportGroup(ctx *ImportGroupContext) interface{}

	// Visit a parse tree produced by MojoParser#importValue.
	VisitImportValue(ctx *ImportValueContext) interface{}

	// Visit a parse tree produced by MojoParser#importType.
	VisitImportType(ctx *ImportTypeContext) interface{}

	// Visit a parse tree produced by MojoParser#constantDeclaration.
	VisitConstantDeclaration(ctx *ConstantDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#patternInitializers.
	VisitPatternInitializers(ctx *PatternInitializersContext) interface{}

	// Visit a parse tree produced by MojoParser#patternInitializer.
	VisitPatternInitializer(ctx *PatternInitializerContext) interface{}

	// Visit a parse tree produced by MojoParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by MojoParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#typeAliasDeclaration.
	VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#typealiasName.
	VisitTypealiasName(ctx *TypealiasNameContext) interface{}

	// Visit a parse tree produced by MojoParser#typealiasAssignment.
	VisitTypealiasAssignment(ctx *TypealiasAssignmentContext) interface{}

	// Visit a parse tree produced by MojoParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#functionHead.
	VisitFunctionHead(ctx *FunctionHeadContext) interface{}

	// Visit a parse tree produced by MojoParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by MojoParser#functionSignature.
	VisitFunctionSignature(ctx *FunctionSignatureContext) interface{}

	// Visit a parse tree produced by MojoParser#functionResult.
	VisitFunctionResult(ctx *FunctionResultContext) interface{}

	// Visit a parse tree produced by MojoParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by MojoParser#functionParameterClause.
	VisitFunctionParameterClause(ctx *FunctionParameterClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#functionParameters.
	VisitFunctionParameters(ctx *FunctionParametersContext) interface{}

	// Visit a parse tree produced by MojoParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by MojoParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by MojoParser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by MojoParser#enumMembers.
	VisitEnumMembers(ctx *EnumMembersContext) interface{}

	// Visit a parse tree produced by MojoParser#enumMember.
	VisitEnumMember(ctx *EnumMemberContext) interface{}

	// Visit a parse tree produced by MojoParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#structName.
	VisitStructName(ctx *StructNameContext) interface{}

	// Visit a parse tree produced by MojoParser#structBody.
	VisitStructBody(ctx *StructBodyContext) interface{}

	// Visit a parse tree produced by MojoParser#structMembers.
	VisitStructMembers(ctx *StructMembersContext) interface{}

	// Visit a parse tree produced by MojoParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by MojoParser#structMemberDeclaration.
	VisitStructMemberDeclaration(ctx *StructMemberDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#interfaceName.
	VisitInterfaceName(ctx *InterfaceNameContext) interface{}

	// Visit a parse tree produced by MojoParser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) interface{}

	// Visit a parse tree produced by MojoParser#interfaceMembers.
	VisitInterfaceMembers(ctx *InterfaceMembersContext) interface{}

	// Visit a parse tree produced by MojoParser#interfaceMember.
	VisitInterfaceMember(ctx *InterfaceMemberContext) interface{}

	// Visit a parse tree produced by MojoParser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by MojoParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by MojoParser#wildcard_pattern.
	VisitWildcard_pattern(ctx *Wildcard_patternContext) interface{}

	// Visit a parse tree produced by MojoParser#identifierPattern.
	VisitIdentifierPattern(ctx *IdentifierPatternContext) interface{}

	// Visit a parse tree produced by MojoParser#tuple_pattern.
	VisitTuple_pattern(ctx *Tuple_patternContext) interface{}

	// Visit a parse tree produced by MojoParser#tuple_pattern_element_list.
	VisitTuple_pattern_element_list(ctx *Tuple_pattern_element_listContext) interface{}

	// Visit a parse tree produced by MojoParser#tuple_pattern_element.
	VisitTuple_pattern_element(ctx *Tuple_pattern_elementContext) interface{}

	// Visit a parse tree produced by MojoParser#optional_pattern.
	VisitOptional_pattern(ctx *Optional_patternContext) interface{}

	// Visit a parse tree produced by MojoParser#expression_pattern.
	VisitExpression_pattern(ctx *Expression_patternContext) interface{}

	// Visit a parse tree produced by MojoParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by MojoParser#attributeName.
	VisitAttributeName(ctx *AttributeNameContext) interface{}

	// Visit a parse tree produced by MojoParser#attributeNameIdentifier.
	VisitAttributeNameIdentifier(ctx *AttributeNameIdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#attributeArgumentClause.
	VisitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#attributes.
	VisitAttributes(ctx *AttributesContext) interface{}

	// Visit a parse tree produced by MojoParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by MojoParser#prefixExpression.
	VisitPrefixExpression(ctx *PrefixExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#binaryExpressions.
	VisitBinaryExpressions(ctx *BinaryExpressionsContext) interface{}

	// Visit a parse tree produced by MojoParser#conditional_operator.
	VisitConditional_operator(ctx *Conditional_operatorContext) interface{}

	// Visit a parse tree produced by MojoParser#type_casting_operator.
	VisitType_casting_operator(ctx *Type_casting_operatorContext) interface{}

	// Visit a parse tree produced by MojoParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by MojoParser#arrayLiteralItems.
	VisitArrayLiteralItems(ctx *ArrayLiteralItemsContext) interface{}

	// Visit a parse tree produced by MojoParser#arrayLiteralItem.
	VisitArrayLiteralItem(ctx *ArrayLiteralItemContext) interface{}

	// Visit a parse tree produced by MojoParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by MojoParser#objectLiteralItems.
	VisitObjectLiteralItems(ctx *ObjectLiteralItemsContext) interface{}

	// Visit a parse tree produced by MojoParser#objectLiteralItem.
	VisitObjectLiteralItem(ctx *ObjectLiteralItemContext) interface{}

	// Visit a parse tree produced by MojoParser#implicitMemberExpression.
	VisitImplicitMemberExpression(ctx *ImplicitMemberExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#tupleElement.
	VisitTupleElement(ctx *TupleElementContext) interface{}

	// Visit a parse tree produced by MojoParser#wildcardExpression.
	VisitWildcardExpression(ctx *WildcardExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#explicitMemberExpression1.
	VisitExplicitMemberExpression1(ctx *ExplicitMemberExpression1Context) interface{}

	// Visit a parse tree produced by MojoParser#postfixOperation.
	VisitPostfixOperation(ctx *PostfixOperationContext) interface{}

	// Visit a parse tree produced by MojoParser#explicitMemberExpression4.
	VisitExplicitMemberExpression4(ctx *ExplicitMemberExpression4Context) interface{}

	// Visit a parse tree produced by MojoParser#subscriptExpression.
	VisitSubscriptExpression(ctx *SubscriptExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#explicitMemberExpression3.
	VisitExplicitMemberExpression3(ctx *ExplicitMemberExpression3Context) interface{}

	// Visit a parse tree produced by MojoParser#explicitMemberExpression2.
	VisitExplicitMemberExpression2(ctx *ExplicitMemberExpression2Context) interface{}

	// Visit a parse tree produced by MojoParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by MojoParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by MojoParser#functionCallArgumentClause.
	VisitFunctionCallArgumentClause(ctx *FunctionCallArgumentClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#function_call_argument_list.
	VisitFunction_call_argument_list(ctx *Function_call_argument_listContext) interface{}

	// Visit a parse tree produced by MojoParser#function_call_argument.
	VisitFunction_call_argument(ctx *Function_call_argumentContext) interface{}

	// Visit a parse tree produced by MojoParser#argumentNameList.
	VisitArgumentNameList(ctx *ArgumentNameListContext) interface{}

	// Visit a parse tree produced by MojoParser#argument_name.
	VisitArgument_name(ctx *Argument_nameContext) interface{}

	// Visit a parse tree produced by MojoParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by MojoParser#Intersection.
	VisitIntersection(ctx *IntersectionContext) interface{}

	// Visit a parse tree produced by MojoParser#Prime.
	VisitPrime(ctx *PrimeContext) interface{}

	// Visit a parse tree produced by MojoParser#Union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by MojoParser#primeType.
	VisitPrimeType(ctx *PrimeTypeContext) interface{}

	// Visit a parse tree produced by MojoParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by MojoParser#typeIdentifier.
	VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#typeIdentifierClause.
	VisitTypeIdentifierClause(ctx *TypeIdentifierClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by MojoParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by MojoParser#tupleTypeElements.
	VisitTupleTypeElements(ctx *TupleTypeElementsContext) interface{}

	// Visit a parse tree produced by MojoParser#tupleTypeElement.
	VisitTupleTypeElement(ctx *TupleTypeElementContext) interface{}

	// Visit a parse tree produced by MojoParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by MojoParser#functionTypeArgumentClause.
	VisitFunctionTypeArgumentClause(ctx *FunctionTypeArgumentClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#functionTypeArguments.
	VisitFunctionTypeArguments(ctx *FunctionTypeArgumentsContext) interface{}

	// Visit a parse tree produced by MojoParser#functionTypeArgument.
	VisitFunctionTypeArgument(ctx *FunctionTypeArgumentContext) interface{}

	// Visit a parse tree produced by MojoParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by MojoParser#dictionaryType.
	VisitDictionaryType(ctx *DictionaryTypeContext) interface{}

	// Visit a parse tree produced by MojoParser#typeInheritanceClause.
	VisitTypeInheritanceClause(ctx *TypeInheritanceClauseContext) interface{}

	// Visit a parse tree produced by MojoParser#typeInheritances.
	VisitTypeInheritances(ctx *TypeInheritancesContext) interface{}

	// Visit a parse tree produced by MojoParser#typeInheritance.
	VisitTypeInheritance(ctx *TypeInheritanceContext) interface{}

	// Visit a parse tree produced by MojoParser#declarationIdentifier.
	VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#labelIdentifier.
	VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#path_identifier.
	VisitPath_identifier(ctx *Path_identifierContext) interface{}

	// Visit a parse tree produced by MojoParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MojoParser#keyword_as_identifier_in_declarations.
	VisitKeyword_as_identifier_in_declarations(ctx *Keyword_as_identifier_in_declarationsContext) interface{}

	// Visit a parse tree produced by MojoParser#keyword_as_identifier_in_labels.
	VisitKeyword_as_identifier_in_labels(ctx *Keyword_as_identifier_in_labelsContext) interface{}

	// Visit a parse tree produced by MojoParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by MojoParser#followingDocument.
	VisitFollowingDocument(ctx *FollowingDocumentContext) interface{}

	// Visit a parse tree produced by MojoParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#negatePrefixOperator.
	VisitNegatePrefixOperator(ctx *NegatePrefixOperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#compilation_condition_AND.
	VisitCompilation_condition_AND(ctx *Compilation_condition_ANDContext) interface{}

	// Visit a parse tree produced by MojoParser#compilation_condition_OR.
	VisitCompilation_condition_OR(ctx *Compilation_condition_ORContext) interface{}

	// Visit a parse tree produced by MojoParser#compilation_condition_GE.
	VisitCompilation_condition_GE(ctx *Compilation_condition_GEContext) interface{}

	// Visit a parse tree produced by MojoParser#arrowOperator.
	VisitArrowOperator(ctx *ArrowOperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#range_operator.
	VisitRange_operator(ctx *Range_operatorContext) interface{}

	// Visit a parse tree produced by MojoParser#range_right_open_operator.
	VisitRange_right_open_operator(ctx *Range_right_open_operatorContext) interface{}

	// Visit a parse tree produced by MojoParser#same_type_equals.
	VisitSame_type_equals(ctx *Same_type_equalsContext) interface{}

	// Visit a parse tree produced by MojoParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#prefixOperator.
	VisitPrefixOperator(ctx *PrefixOperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#postfixOperator.
	VisitPostfixOperator(ctx *PostfixOperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by MojoParser#operator_character.
	VisitOperator_character(ctx *Operator_characterContext) interface{}

	// Visit a parse tree produced by MojoParser#operator_head.
	VisitOperator_head(ctx *Operator_headContext) interface{}

	// Visit a parse tree produced by MojoParser#dot_operator_head.
	VisitDot_operator_head(ctx *Dot_operator_headContext) interface{}

	// Visit a parse tree produced by MojoParser#dot_operator_character.
	VisitDot_operator_character(ctx *Dot_operator_characterContext) interface{}

	// Visit a parse tree produced by MojoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by MojoParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by MojoParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by MojoParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by MojoParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by MojoParser#eov.
	VisitEov(ctx *EovContext) interface{}

	// Visit a parse tree produced by MojoParser#eosWithDocument.
	VisitEosWithDocument(ctx *EosWithDocumentContext) interface{}

	// Visit a parse tree produced by MojoParser#eovWithDocument.
	VisitEovWithDocument(ctx *EovWithDocumentContext) interface{}
}
