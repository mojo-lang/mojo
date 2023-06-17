// Code generated from MojoParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package syntax // MojoParser
import "github.com/antlr4-go/antlr/v4"

type BaseMojoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMojoParserVisitor) VisitMojoFile(ctx *MojoFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIfModifier(ctx *IfModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitWhileModifier(ctx *WhileModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFloatingStatement(ctx *FloatingStatementContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitDocumentedPatternInitializer(ctx *DocumentedPatternInitializerContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitAttributeAliasDeclaration(ctx *AttributeAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributeAliasAssignment(ctx *AttributeAliasAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitWildcardPattern(ctx *WildcardPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIdentifierPattern(ctx *IdentifierPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTuplePattern(ctx *TuplePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTuplePatternElementList(ctx *TuplePatternElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTuplePatternElement(ctx *TuplePatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayPattern(ctx *ArrayPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayPatternElements(ctx *ArrayPatternElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArrayPatternElement(ctx *ArrayPatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitEnumValuePattern(ctx *EnumValuePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitOptionalPattern(ctx *OptionalPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExpressionPattern(ctx *ExpressionPatternContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitAttributeArgument(ctx *AttributeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributeArguments(ctx *AttributeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitAttributes(ctx *AttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrefixCallOperator(ctx *PrefixCallOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInfixCallOperator(ctx *InfixCallOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitBinaryExpressions(ctx *BinaryExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitInOperator(ctx *InOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitConditionalOperator(ctx *ConditionalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIfOperator(ctx *IfOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTypeCastingOperator(ctx *TypeCastingOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitNumericOperatorLiteral(ctx *NumericOperatorLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStringOperatorLiteral(ctx *StringOperatorLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitSuffixLiteralOperator(ctx *SuffixLiteralOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPrefixLiteralOperator(ctx *PrefixLiteralOperatorContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMapLiteralItems(ctx *MapLiteralItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMapLiteralItem(ctx *MapLiteralItemContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitStructConstructionExpression(ctx *StructConstructionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMatchExprSuffix(ctx *MatchExprSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMatchExprCases(ctx *MatchExprCasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMatchExprCase(ctx *MatchExprCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitClosureExpression(ctx *ClosureExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitClosureParameters(ctx *ClosureParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitClosureParameter(ctx *ClosureParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitImplicitMemberExpression(ctx *ImplicitMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleLiteralExpression(ctx *TupleLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTupleElement(ctx *TupleElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitWildcardExpression(ctx *WildcardExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitSuffixExpression(ctx *SuffixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitExplicitMemberSuffix(ctx *ExplicitMemberSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitSubscriptSuffix(ctx *SubscriptSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionCallSuffix(ctx *FunctionCallSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionCallArgumentClause(ctx *FunctionCallArgumentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionCallArguments(ctx *FunctionCallArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitFunctionCallArgument(ctx *FunctionCallArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitTrailingClosures(ctx *TrailingClosuresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLabeledTrailingClosures(ctx *LabeledTrailingClosuresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitLabeledTrailingClosure(ctx *LabeledTrailingClosureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArgumentNames(ctx *ArgumentNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitArgumentName(ctx *ArgumentNameContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitKeyAttributes(ctx *KeyAttributesContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitPathIdentifier(ctx *PathIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitKeywordAsIdentifierInDeclarations(ctx *KeywordAsIdentifierInDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitKeywordAsIdentifierInLabels(ctx *KeywordAsIdentifierInLabelsContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitArrowOperator(ctx *ArrowOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitRangeOperator(ctx *RangeOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMojoParserVisitor) VisitHalfOpenRangeOperator(ctx *HalfOpenRangeOperatorContext) interface{} {
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

func (v *BaseMojoParserVisitor) VisitOperator_characters(ctx *Operator_charactersContext) interface{} {
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
