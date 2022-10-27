// Generated from java-escape by ANTLR 4.11.1
package org.mojolang.mojo.parser.syntax;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link MojoParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface MojoParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link MojoParser#mojoFile}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMojoFile(MojoParser.MojoFileContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatement(MojoParser.StatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#freeFloatingDocument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFreeFloatingDocument(MojoParser.FreeFloatingDocumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#statements}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatements(MojoParser.StatementsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#loopStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLoopStatement(MojoParser.LoopStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#forInStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitForInStatement(MojoParser.ForInStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#whileStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWhileStatement(MojoParser.WhileStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#conditions}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConditions(MojoParser.ConditionsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#condition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCondition(MojoParser.ConditionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#optionalBindingCondition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOptionalBindingCondition(MojoParser.OptionalBindingConditionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#branchStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBranchStatement(MojoParser.BranchStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIfStatement(MojoParser.IfStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#elseClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitElseClause(MojoParser.ElseClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#matchStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatchStatement(MojoParser.MatchStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#matchCases}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatchCases(MojoParser.MatchCasesContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#matchCase}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatchCase(MojoParser.MatchCaseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#controlTransferStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitControlTransferStatement(MojoParser.ControlTransferStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#breakStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBreakStatement(MojoParser.BreakStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#continueStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitContinueStatement(MojoParser.ContinueStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#returnStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitReturnStatement(MojoParser.ReturnStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#genericParameterClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGenericParameterClause(MojoParser.GenericParameterClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#genericParameters}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGenericParameters(MojoParser.GenericParametersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#genericParameter}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGenericParameter(MojoParser.GenericParameterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#genericArgumentClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGenericArgumentClause(MojoParser.GenericArgumentClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#genericArguments}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGenericArguments(MojoParser.GenericArgumentsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#genericArgument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGenericArgument(MojoParser.GenericArgumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeclaration(MojoParser.DeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#codeBlock}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCodeBlock(MojoParser.CodeBlockContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#packageDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPackageDeclaration(MojoParser.PackageDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#packageIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPackageIdentifier(MojoParser.PackageIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#packageName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPackageName(MojoParser.PackageNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportDeclaration(MojoParser.ImportDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importPath}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportPath(MojoParser.ImportPathContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importPathIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportPathIdentifier(MojoParser.ImportPathIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importAllClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportAllClause(MojoParser.ImportAllClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importValueAsClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportValueAsClause(MojoParser.ImportValueAsClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importTypeClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportTypeClause(MojoParser.ImportTypeClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importTypeAsClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportTypeAsClause(MojoParser.ImportTypeAsClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importGroupClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportGroupClause(MojoParser.ImportGroupClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importGroup}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportGroup(MojoParser.ImportGroupContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportValue(MojoParser.ImportValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#importType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportType(MojoParser.ImportTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#constantDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConstantDeclaration(MojoParser.ConstantDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#patternInitializers}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPatternInitializers(MojoParser.PatternInitializersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#documentedPatternInitializer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDocumentedPatternInitializer(MojoParser.DocumentedPatternInitializerContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#patternInitializer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPatternInitializer(MojoParser.PatternInitializerContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#initializer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInitializer(MojoParser.InitializerContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#variableDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVariableDeclaration(MojoParser.VariableDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeAliasDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeAliasDeclaration(MojoParser.TypeAliasDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeAliasName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeAliasName(MojoParser.TypeAliasNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeAliasAssignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeAliasAssignment(MojoParser.TypeAliasAssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionDeclaration(MojoParser.FunctionDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionName(MojoParser.FunctionNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionSignature}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionSignature(MojoParser.FunctionSignatureContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionResult}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionResult(MojoParser.FunctionResultContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionBody}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionBody(MojoParser.FunctionBodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionParameterClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionParameterClause(MojoParser.FunctionParameterClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionParameters}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionParameters(MojoParser.FunctionParametersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionParameter}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionParameter(MojoParser.FunctionParameterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enumDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnumDeclaration(MojoParser.EnumDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enumBody}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnumBody(MojoParser.EnumBodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enumName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnumName(MojoParser.EnumNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enumMembers}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnumMembers(MojoParser.EnumMembersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enumMember}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnumMember(MojoParser.EnumMemberContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructDeclaration(MojoParser.StructDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructName(MojoParser.StructNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructType(MojoParser.StructTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structBody}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructBody(MojoParser.StructBodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structMembers}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructMembers(MojoParser.StructMembersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structMember}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructMember(MojoParser.StructMemberContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structMemberDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructMemberDeclaration(MojoParser.StructMemberDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceDeclaration(MojoParser.InterfaceDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceName(MojoParser.InterfaceNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceType(MojoParser.InterfaceTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceBody}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceBody(MojoParser.InterfaceBodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceMembers}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceMembers(MojoParser.InterfaceMembersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceMember}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceMember(MojoParser.InterfaceMemberContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interfaceMethodDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceMethodDeclaration(MojoParser.InterfaceMethodDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeDeclaration(MojoParser.AttributeDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeAliasDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeAliasDeclaration(MojoParser.AttributeAliasDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeAliasAssignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeAliasAssignment(MojoParser.AttributeAliasAssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPattern(MojoParser.PatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#wildcardPattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWildcardPattern(MojoParser.WildcardPatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#identifierPattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdentifierPattern(MojoParser.IdentifierPatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuplePattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuplePattern(MojoParser.TuplePatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuplePatternElementList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuplePatternElementList(MojoParser.TuplePatternElementListContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuplePatternElement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuplePatternElement(MojoParser.TuplePatternElementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#optionalPattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOptionalPattern(MojoParser.OptionalPatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#expressionPattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionPattern(MojoParser.ExpressionPatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attribute}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttribute(MojoParser.AttributeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeIdentifier(MojoParser.AttributeIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeName(MojoParser.AttributeNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeArgumentClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeArgumentClause(MojoParser.AttributeArgumentClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeArgument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeArgument(MojoParser.AttributeArgumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributeArguments}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeArguments(MojoParser.AttributeArgumentsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attributes}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributes(MojoParser.AttributesContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression(MojoParser.ExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#expressions}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressions(MojoParser.ExpressionsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#prefixExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrefixExpression(MojoParser.PrefixExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#binaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryExpression(MojoParser.BinaryExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#binaryExpressions}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryExpressions(MojoParser.BinaryExpressionsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#conditionalOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConditionalOperator(MojoParser.ConditionalOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeCastingOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeCastingOperator(MojoParser.TypeCastingOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimaryExpression(MojoParser.PrimaryExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#literalExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLiteralExpression(MojoParser.LiteralExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#numericOperatorLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumericOperatorLiteral(MojoParser.NumericOperatorLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#stringOperatorLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStringOperatorLiteral(MojoParser.StringOperatorLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#suffixLiteralOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSuffixLiteralOperator(MojoParser.SuffixLiteralOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#prefixLiteralOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrefixLiteralOperator(MojoParser.PrefixLiteralOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#arrayLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrayLiteral(MojoParser.ArrayLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#arrayLiteralItems}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrayLiteralItems(MojoParser.ArrayLiteralItemsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#arrayLiteralItem}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrayLiteralItem(MojoParser.ArrayLiteralItemContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#mapLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMapLiteral(MojoParser.MapLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#mapLiteralItems}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMapLiteralItems(MojoParser.MapLiteralItemsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#mapLiteralItem}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMapLiteralItem(MojoParser.MapLiteralItemContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#objectLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObjectLiteral(MojoParser.ObjectLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#objectLiteralItems}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObjectLiteralItems(MojoParser.ObjectLiteralItemsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#objectLiteralItem}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObjectLiteralItem(MojoParser.ObjectLiteralItemContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructLiteral(MojoParser.StructLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#structConstructionExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructConstructionExpression(MojoParser.StructConstructionExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#closureExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitClosureExpression(MojoParser.ClosureExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#closureParameters}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitClosureParameters(MojoParser.ClosureParametersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#closureParameter}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitClosureParameter(MojoParser.ClosureParameterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#implicitMemberExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImplicitMemberExpression(MojoParser.ImplicitMemberExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#parenthesizedExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParenthesizedExpression(MojoParser.ParenthesizedExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tupleExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTupleExpression(MojoParser.TupleExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tupleElement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTupleElement(MojoParser.TupleElementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#wildcardExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWildcardExpression(MojoParser.WildcardExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#postfixExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPostfixExpression(MojoParser.PostfixExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#suffixExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSuffixExpression(MojoParser.SuffixExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#explicitMemberSuffix}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExplicitMemberSuffix(MojoParser.ExplicitMemberSuffixContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#subscriptSuffix}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSubscriptSuffix(MojoParser.SubscriptSuffixContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionCallSuffix}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCallSuffix(MojoParser.FunctionCallSuffixContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionCallArgumentClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCallArgumentClause(MojoParser.FunctionCallArgumentClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionCallArguments}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCallArguments(MojoParser.FunctionCallArgumentsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionCallArgument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCallArgument(MojoParser.FunctionCallArgumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#trailingClosures}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTrailingClosures(MojoParser.TrailingClosuresContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#labeledTrailingClosures}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLabeledTrailingClosures(MojoParser.LabeledTrailingClosuresContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#labeledTrailingClosure}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLabeledTrailingClosure(MojoParser.LabeledTrailingClosureContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#argumentNames}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgumentNames(MojoParser.ArgumentNamesContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#argumentName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgumentName(MojoParser.ArgumentNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_(MojoParser.Type_Context ctx);
	/**
	 * Visit a parse tree produced by the {@code Intersection}
	 * labeled alternative in {@link MojoParser#basicType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntersection(MojoParser.IntersectionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Prime}
	 * labeled alternative in {@link MojoParser#basicType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrime(MojoParser.PrimeContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Union}
	 * labeled alternative in {@link MojoParser#basicType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnion(MojoParser.UnionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#primeType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimeType(MojoParser.PrimeTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeAnnotation}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeAnnotation(MojoParser.TypeAnnotationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeIdentifier(MojoParser.TypeIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeIdentifierClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeIdentifierClause(MojoParser.TypeIdentifierClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeName(MojoParser.TypeNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tupleType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTupleType(MojoParser.TupleTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tupleTypeElements}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTupleTypeElements(MojoParser.TupleTypeElementsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tupleTypeElement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTupleTypeElement(MojoParser.TupleTypeElementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#functionType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionType(MojoParser.FunctionTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#arrayType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrayType(MojoParser.ArrayTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#mapType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMapType(MojoParser.MapTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#keyAttributes}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyAttributes(MojoParser.KeyAttributesContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeInheritanceClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeInheritanceClause(MojoParser.TypeInheritanceClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeInheritances}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeInheritances(MojoParser.TypeInheritancesContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typeInheritance}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeInheritance(MojoParser.TypeInheritanceContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#declarationIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeclarationIdentifier(MojoParser.DeclarationIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#labelIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLabelIdentifier(MojoParser.LabelIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#pathIdentifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPathIdentifier(MojoParser.PathIdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdentifier(MojoParser.IdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#keywordAsIdentifierInDeclarations}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeywordAsIdentifierInDeclarations(MojoParser.KeywordAsIdentifierInDeclarationsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#keywordAsIdentifierInLabels}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeywordAsIdentifierInLabels(MojoParser.KeywordAsIdentifierInLabelsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#document}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDocument(MojoParser.DocumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#followingDocument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFollowingDocument(MojoParser.FollowingDocumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#assignmentOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssignmentOperator(MojoParser.AssignmentOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#negatePrefixOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNegatePrefixOperator(MojoParser.NegatePrefixOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#arrowOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrowOperator(MojoParser.ArrowOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#rangeOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRangeOperator(MojoParser.RangeOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#halfOpenRangeOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitHalfOpenRangeOperator(MojoParser.HalfOpenRangeOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#binaryOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryOperator(MojoParser.BinaryOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#prefixOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrefixOperator(MojoParser.PrefixOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#postfixOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPostfixOperator(MojoParser.PostfixOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperator(MojoParser.OperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#operator_characters}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperator_characters(MojoParser.Operator_charactersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#operator_character}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperator_character(MojoParser.Operator_characterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#operator_head}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperator_head(MojoParser.Operator_headContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#dot_operator_head}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDot_operator_head(MojoParser.Dot_operator_headContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#dot_operator_character}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDot_operator_character(MojoParser.Dot_operator_characterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLiteral(MojoParser.LiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#boolLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBoolLiteral(MojoParser.BoolLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#nullLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNullLiteral(MojoParser.NullLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#numericLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumericLiteral(MojoParser.NumericLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#integerLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegerLiteral(MojoParser.IntegerLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#stringLiteral}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStringLiteral(MojoParser.StringLiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#eos}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEos(MojoParser.EosContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#eov}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEov(MojoParser.EovContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#eosWithDocument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEosWithDocument(MojoParser.EosWithDocumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#eovWithDocument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEovWithDocument(MojoParser.EovWithDocumentContext ctx);
}