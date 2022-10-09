// Generated from /Users/frankee/Projects/mojo/mojo/antlr/Mojo.g4 by ANTLR 4.8
package org.mojolang.mojo.parser.syntax;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link MojoParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface MojoVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link MojoParser#source_file}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSource_file(MojoParser.Source_fileContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatement(MojoParser.StatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#loop_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLoop_statement(MojoParser.Loop_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#for_in_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFor_in_statement(MojoParser.For_in_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#while_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWhile_statement(MojoParser.While_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#condition_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCondition_list(MojoParser.Condition_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#condition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCondition(MojoParser.ConditionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#optional_binding_condition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOptional_binding_condition(MojoParser.Optional_binding_conditionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#repeat_while_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRepeat_while_statement(MojoParser.Repeat_while_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#branch_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBranch_statement(MojoParser.Branch_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#if_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIf_statement(MojoParser.If_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#else_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitElse_clause(MojoParser.Else_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#match_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatch_statement(MojoParser.Match_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#match_cases}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatch_cases(MojoParser.Match_casesContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#match_case}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMatch_case(MojoParser.Match_caseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#control_transfer_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitControl_transfer_statement(MojoParser.Control_transfer_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#break_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBreak_statement(MojoParser.Break_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#continue_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitContinue_statement(MojoParser.Continue_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#return_statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitReturn_statement(MojoParser.Return_statementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#generic_parameter_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGeneric_parameter_clause(MojoParser.Generic_parameter_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#generic_parameter_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGeneric_parameter_list(MojoParser.Generic_parameter_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#generic_parameter}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGeneric_parameter(MojoParser.Generic_parameterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#generic_argument_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGeneric_argument_clause(MojoParser.Generic_argument_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#generic_argument_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGeneric_argument_list(MojoParser.Generic_argument_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#generic_argument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGeneric_argument(MojoParser.Generic_argumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeclaration(MojoParser.DeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#code_block}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCode_block(MojoParser.Code_blockContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#package_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPackage_declaration(MojoParser.Package_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#package_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPackage_identifier(MojoParser.Package_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_declaration(MojoParser.Import_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_path}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_path(MojoParser.Import_pathContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_path_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_path_identifier(MojoParser.Import_path_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_identifier_as_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_identifier_as_clause(MojoParser.Import_identifier_as_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_type_as_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_type_as_clause(MojoParser.Import_type_as_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_type_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_type_clause(MojoParser.Import_type_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_identifier_group_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_identifier_group_clause(MojoParser.Import_identifier_group_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_identifier_group}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_identifier_group(MojoParser.Import_identifier_groupContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_identifier(MojoParser.Import_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#import_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImport_type(MojoParser.Import_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#constant_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConstant_declaration(MojoParser.Constant_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#pattern_initializer_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPattern_initializer_list(MojoParser.Pattern_initializer_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#pattern_initializer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPattern_initializer(MojoParser.Pattern_initializerContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#initializer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInitializer(MojoParser.InitializerContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typealias_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypealias_declaration(MojoParser.Typealias_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typealias_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypealias_name(MojoParser.Typealias_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#typealias_assignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypealias_assignment(MojoParser.Typealias_assignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_declaration(MojoParser.Function_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_head}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_head(MojoParser.Function_headContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_name(MojoParser.Function_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_signature}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_signature(MojoParser.Function_signatureContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_result}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_result(MojoParser.Function_resultContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_body}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_body(MojoParser.Function_bodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_parameter_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_parameter_clause(MojoParser.Function_parameter_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_parameter_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_parameter_list(MojoParser.Function_parameter_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_parameter}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_parameter(MojoParser.Function_parameterContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enum_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnum_declaration(MojoParser.Enum_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enum_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnum_name(MojoParser.Enum_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enum_members}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnum_members(MojoParser.Enum_membersContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#enum_member}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEnum_member(MojoParser.Enum_memberContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#struct_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStruct_declaration(MojoParser.Struct_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#struct_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStruct_name(MojoParser.Struct_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#struct_body}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStruct_body(MojoParser.Struct_bodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#struct_member}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStruct_member(MojoParser.Struct_memberContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#struct_member_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStruct_member_declaration(MojoParser.Struct_member_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interface_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterface_declaration(MojoParser.Interface_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interface_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterface_name(MojoParser.Interface_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interface_body}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterface_body(MojoParser.Interface_bodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interface_member}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterface_member(MojoParser.Interface_memberContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#interface_method_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterface_method_declaration(MojoParser.Interface_method_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPattern(MojoParser.PatternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#wildcard_pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWildcard_pattern(MojoParser.Wildcard_patternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#identifier_pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdentifier_pattern(MojoParser.Identifier_patternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_pattern(MojoParser.Tuple_patternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_pattern_element_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_pattern_element_list(MojoParser.Tuple_pattern_element_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_pattern_element}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_pattern_element(MojoParser.Tuple_pattern_elementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#optional_pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOptional_pattern(MojoParser.Optional_patternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#expression_pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression_pattern(MojoParser.Expression_patternContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attribute}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttribute(MojoParser.AttributeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attribute_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttribute_name(MojoParser.Attribute_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#attribute_argument_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttribute_argument_clause(MojoParser.Attribute_argument_clauseContext ctx);
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
	 * Visit a parse tree produced by {@link MojoParser#expression_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression_list(MojoParser.Expression_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#prefix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrefix_expression(MojoParser.Prefix_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#binary_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinary_expression(MojoParser.Binary_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#binary_expressions}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinary_expressions(MojoParser.Binary_expressionsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#conditional_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConditional_operator(MojoParser.Conditional_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_casting_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_casting_operator(MojoParser.Type_casting_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#primary_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimary_expression(MojoParser.Primary_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#literal_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLiteral_expression(MojoParser.Literal_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#array_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArray_literal(MojoParser.Array_literalContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#array_literal_items}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArray_literal_items(MojoParser.Array_literal_itemsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#array_literal_item}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArray_literal_item(MojoParser.Array_literal_itemContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#object_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObject_literal(MojoParser.Object_literalContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#object_literal_items}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObject_literal_items(MojoParser.Object_literal_itemsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#object_literal_item}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObject_literal_item(MojoParser.Object_literal_itemContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#implicit_member_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImplicit_member_expression(MojoParser.Implicit_member_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#parenthesized_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParenthesized_expression(MojoParser.Parenthesized_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_expression(MojoParser.Tuple_expressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_element}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_element(MojoParser.Tuple_elementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#wildcard_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWildcard_expression(MojoParser.Wildcard_expressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code function_call_expression}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_call_expression(MojoParser.Function_call_expressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code subscript_expression}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSubscript_expression(MojoParser.Subscript_expressionContext ctx);
	/**
	 * Visit a parse tree produced by the {@code explicit_member_expression1}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExplicit_member_expression1(MojoParser.Explicit_member_expression1Context ctx);
	/**
	 * Visit a parse tree produced by the {@code explicit_member_expression2}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExplicit_member_expression2(MojoParser.Explicit_member_expression2Context ctx);
	/**
	 * Visit a parse tree produced by the {@code explicit_member_expression3}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExplicit_member_expression3(MojoParser.Explicit_member_expression3Context ctx);
	/**
	 * Visit a parse tree produced by the {@code explicit_member_expression4}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExplicit_member_expression4(MojoParser.Explicit_member_expression4Context ctx);
	/**
	 * Visit a parse tree produced by the {@code postfix_operation}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPostfix_operation(MojoParser.Postfix_operationContext ctx);
	/**
	 * Visit a parse tree produced by the {@code primary}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimary(MojoParser.PrimaryContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_call_argument_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_call_argument_clause(MojoParser.Function_call_argument_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_call_argument_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_call_argument_list(MojoParser.Function_call_argument_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_call_argument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_call_argument(MojoParser.Function_call_argumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#argument_name_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgument_name_list(MojoParser.Argument_name_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#argument_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgument_name(MojoParser.Argument_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_(MojoParser.Type_Context ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_annotation}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_annotation(MojoParser.Type_annotationContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_identifier(MojoParser.Type_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_identifier_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_identifier_clause(MojoParser.Type_identifier_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_name(MojoParser.Type_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_type(MojoParser.Tuple_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_type_element_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_type_element_list(MojoParser.Tuple_type_element_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#tuple_type_element}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTuple_type_element(MojoParser.Tuple_type_elementContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_type(MojoParser.Function_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_type_argument_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_type_argument_clause(MojoParser.Function_type_argument_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_type_argument_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_type_argument_list(MojoParser.Function_type_argument_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#function_type_argument}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunction_type_argument(MojoParser.Function_type_argumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#argument_label}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgument_label(MojoParser.Argument_labelContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#array_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArray_type(MojoParser.Array_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#map_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMap_type(MojoParser.Map_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_inheritance_clause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_inheritance_clause(MojoParser.Type_inheritance_clauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#type_inheritance_list}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitType_inheritance_list(MojoParser.Type_inheritance_listContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#declaration_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeclaration_identifier(MojoParser.Declaration_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#label_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLabel_identifier(MojoParser.Label_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#path_identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPath_identifier(MojoParser.Path_identifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#identifier}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdentifier(MojoParser.IdentifierContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#keyword_as_identifier_in_declarations}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyword_as_identifier_in_declarations(MojoParser.Keyword_as_identifier_in_declarationsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#keyword_as_identifier_in_labels}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyword_as_identifier_in_labels(MojoParser.Keyword_as_identifier_in_labelsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#document}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDocument(MojoParser.DocumentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#following_document}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFollowing_document(MojoParser.Following_documentContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#assignment_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssignment_operator(MojoParser.Assignment_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#negate_prefix_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNegate_prefix_operator(MojoParser.Negate_prefix_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#compilation_condition_AND}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCompilation_condition_AND(MojoParser.Compilation_condition_ANDContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#compilation_condition_OR}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCompilation_condition_OR(MojoParser.Compilation_condition_ORContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#compilation_condition_GE}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCompilation_condition_GE(MojoParser.Compilation_condition_GEContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#arrow_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrow_operator(MojoParser.Arrow_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#range_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRange_operator(MojoParser.Range_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#same_type_equals}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSame_type_equals(MojoParser.Same_type_equalsContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#binary_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinary_operator(MojoParser.Binary_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#prefix_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrefix_operator(MojoParser.Prefix_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#postfix_operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPostfix_operator(MojoParser.Postfix_operatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperator(MojoParser.OperatorContext ctx);
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
	 * Visit a parse tree produced by {@link MojoParser#numeric_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumeric_literal(MojoParser.Numeric_literalContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#boolean_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBoolean_literal(MojoParser.Boolean_literalContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#null_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNull_literal(MojoParser.Null_literalContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#integer_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInteger_literal(MojoParser.Integer_literalContext ctx);
	/**
	 * Visit a parse tree produced by {@link MojoParser#string_literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitString_literal(MojoParser.String_literalContext ctx);
}