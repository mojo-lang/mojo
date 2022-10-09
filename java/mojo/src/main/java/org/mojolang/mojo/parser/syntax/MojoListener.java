// Generated from /Users/frankee/Projects/mojo/mojo/antlr/Mojo.g4 by ANTLR 4.8
package org.mojolang.mojo.parser.syntax;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link MojoParser}.
 */
public interface MojoListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link MojoParser#source_file}.
	 * @param ctx the parse tree
	 */
	void enterSource_file(MojoParser.Source_fileContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#source_file}.
	 * @param ctx the parse tree
	 */
	void exitSource_file(MojoParser.Source_fileContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(MojoParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(MojoParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#loop_statement}.
	 * @param ctx the parse tree
	 */
	void enterLoop_statement(MojoParser.Loop_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#loop_statement}.
	 * @param ctx the parse tree
	 */
	void exitLoop_statement(MojoParser.Loop_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#for_in_statement}.
	 * @param ctx the parse tree
	 */
	void enterFor_in_statement(MojoParser.For_in_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#for_in_statement}.
	 * @param ctx the parse tree
	 */
	void exitFor_in_statement(MojoParser.For_in_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#while_statement}.
	 * @param ctx the parse tree
	 */
	void enterWhile_statement(MojoParser.While_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#while_statement}.
	 * @param ctx the parse tree
	 */
	void exitWhile_statement(MojoParser.While_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#condition_list}.
	 * @param ctx the parse tree
	 */
	void enterCondition_list(MojoParser.Condition_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#condition_list}.
	 * @param ctx the parse tree
	 */
	void exitCondition_list(MojoParser.Condition_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#condition}.
	 * @param ctx the parse tree
	 */
	void enterCondition(MojoParser.ConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#condition}.
	 * @param ctx the parse tree
	 */
	void exitCondition(MojoParser.ConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#optional_binding_condition}.
	 * @param ctx the parse tree
	 */
	void enterOptional_binding_condition(MojoParser.Optional_binding_conditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#optional_binding_condition}.
	 * @param ctx the parse tree
	 */
	void exitOptional_binding_condition(MojoParser.Optional_binding_conditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#repeat_while_statement}.
	 * @param ctx the parse tree
	 */
	void enterRepeat_while_statement(MojoParser.Repeat_while_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#repeat_while_statement}.
	 * @param ctx the parse tree
	 */
	void exitRepeat_while_statement(MojoParser.Repeat_while_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#branch_statement}.
	 * @param ctx the parse tree
	 */
	void enterBranch_statement(MojoParser.Branch_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#branch_statement}.
	 * @param ctx the parse tree
	 */
	void exitBranch_statement(MojoParser.Branch_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#if_statement}.
	 * @param ctx the parse tree
	 */
	void enterIf_statement(MojoParser.If_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#if_statement}.
	 * @param ctx the parse tree
	 */
	void exitIf_statement(MojoParser.If_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#else_clause}.
	 * @param ctx the parse tree
	 */
	void enterElse_clause(MojoParser.Else_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#else_clause}.
	 * @param ctx the parse tree
	 */
	void exitElse_clause(MojoParser.Else_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#match_statement}.
	 * @param ctx the parse tree
	 */
	void enterMatch_statement(MojoParser.Match_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#match_statement}.
	 * @param ctx the parse tree
	 */
	void exitMatch_statement(MojoParser.Match_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#match_cases}.
	 * @param ctx the parse tree
	 */
	void enterMatch_cases(MojoParser.Match_casesContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#match_cases}.
	 * @param ctx the parse tree
	 */
	void exitMatch_cases(MojoParser.Match_casesContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#match_case}.
	 * @param ctx the parse tree
	 */
	void enterMatch_case(MojoParser.Match_caseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#match_case}.
	 * @param ctx the parse tree
	 */
	void exitMatch_case(MojoParser.Match_caseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#control_transfer_statement}.
	 * @param ctx the parse tree
	 */
	void enterControl_transfer_statement(MojoParser.Control_transfer_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#control_transfer_statement}.
	 * @param ctx the parse tree
	 */
	void exitControl_transfer_statement(MojoParser.Control_transfer_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#break_statement}.
	 * @param ctx the parse tree
	 */
	void enterBreak_statement(MojoParser.Break_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#break_statement}.
	 * @param ctx the parse tree
	 */
	void exitBreak_statement(MojoParser.Break_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#continue_statement}.
	 * @param ctx the parse tree
	 */
	void enterContinue_statement(MojoParser.Continue_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#continue_statement}.
	 * @param ctx the parse tree
	 */
	void exitContinue_statement(MojoParser.Continue_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#return_statement}.
	 * @param ctx the parse tree
	 */
	void enterReturn_statement(MojoParser.Return_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#return_statement}.
	 * @param ctx the parse tree
	 */
	void exitReturn_statement(MojoParser.Return_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#generic_parameter_clause}.
	 * @param ctx the parse tree
	 */
	void enterGeneric_parameter_clause(MojoParser.Generic_parameter_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#generic_parameter_clause}.
	 * @param ctx the parse tree
	 */
	void exitGeneric_parameter_clause(MojoParser.Generic_parameter_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#generic_parameter_list}.
	 * @param ctx the parse tree
	 */
	void enterGeneric_parameter_list(MojoParser.Generic_parameter_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#generic_parameter_list}.
	 * @param ctx the parse tree
	 */
	void exitGeneric_parameter_list(MojoParser.Generic_parameter_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#generic_parameter}.
	 * @param ctx the parse tree
	 */
	void enterGeneric_parameter(MojoParser.Generic_parameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#generic_parameter}.
	 * @param ctx the parse tree
	 */
	void exitGeneric_parameter(MojoParser.Generic_parameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#generic_argument_clause}.
	 * @param ctx the parse tree
	 */
	void enterGeneric_argument_clause(MojoParser.Generic_argument_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#generic_argument_clause}.
	 * @param ctx the parse tree
	 */
	void exitGeneric_argument_clause(MojoParser.Generic_argument_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#generic_argument_list}.
	 * @param ctx the parse tree
	 */
	void enterGeneric_argument_list(MojoParser.Generic_argument_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#generic_argument_list}.
	 * @param ctx the parse tree
	 */
	void exitGeneric_argument_list(MojoParser.Generic_argument_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#generic_argument}.
	 * @param ctx the parse tree
	 */
	void enterGeneric_argument(MojoParser.Generic_argumentContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#generic_argument}.
	 * @param ctx the parse tree
	 */
	void exitGeneric_argument(MojoParser.Generic_argumentContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#declaration}.
	 * @param ctx the parse tree
	 */
	void enterDeclaration(MojoParser.DeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#declaration}.
	 * @param ctx the parse tree
	 */
	void exitDeclaration(MojoParser.DeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#code_block}.
	 * @param ctx the parse tree
	 */
	void enterCode_block(MojoParser.Code_blockContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#code_block}.
	 * @param ctx the parse tree
	 */
	void exitCode_block(MojoParser.Code_blockContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#package_declaration}.
	 * @param ctx the parse tree
	 */
	void enterPackage_declaration(MojoParser.Package_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#package_declaration}.
	 * @param ctx the parse tree
	 */
	void exitPackage_declaration(MojoParser.Package_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#package_identifier}.
	 * @param ctx the parse tree
	 */
	void enterPackage_identifier(MojoParser.Package_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#package_identifier}.
	 * @param ctx the parse tree
	 */
	void exitPackage_identifier(MojoParser.Package_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_declaration}.
	 * @param ctx the parse tree
	 */
	void enterImport_declaration(MojoParser.Import_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_declaration}.
	 * @param ctx the parse tree
	 */
	void exitImport_declaration(MojoParser.Import_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_path}.
	 * @param ctx the parse tree
	 */
	void enterImport_path(MojoParser.Import_pathContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_path}.
	 * @param ctx the parse tree
	 */
	void exitImport_path(MojoParser.Import_pathContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_path_identifier}.
	 * @param ctx the parse tree
	 */
	void enterImport_path_identifier(MojoParser.Import_path_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_path_identifier}.
	 * @param ctx the parse tree
	 */
	void exitImport_path_identifier(MojoParser.Import_path_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_identifier_as_clause}.
	 * @param ctx the parse tree
	 */
	void enterImport_identifier_as_clause(MojoParser.Import_identifier_as_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_identifier_as_clause}.
	 * @param ctx the parse tree
	 */
	void exitImport_identifier_as_clause(MojoParser.Import_identifier_as_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_type_as_clause}.
	 * @param ctx the parse tree
	 */
	void enterImport_type_as_clause(MojoParser.Import_type_as_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_type_as_clause}.
	 * @param ctx the parse tree
	 */
	void exitImport_type_as_clause(MojoParser.Import_type_as_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_type_clause}.
	 * @param ctx the parse tree
	 */
	void enterImport_type_clause(MojoParser.Import_type_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_type_clause}.
	 * @param ctx the parse tree
	 */
	void exitImport_type_clause(MojoParser.Import_type_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_identifier_group_clause}.
	 * @param ctx the parse tree
	 */
	void enterImport_identifier_group_clause(MojoParser.Import_identifier_group_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_identifier_group_clause}.
	 * @param ctx the parse tree
	 */
	void exitImport_identifier_group_clause(MojoParser.Import_identifier_group_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_identifier_group}.
	 * @param ctx the parse tree
	 */
	void enterImport_identifier_group(MojoParser.Import_identifier_groupContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_identifier_group}.
	 * @param ctx the parse tree
	 */
	void exitImport_identifier_group(MojoParser.Import_identifier_groupContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_identifier}.
	 * @param ctx the parse tree
	 */
	void enterImport_identifier(MojoParser.Import_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_identifier}.
	 * @param ctx the parse tree
	 */
	void exitImport_identifier(MojoParser.Import_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#import_type}.
	 * @param ctx the parse tree
	 */
	void enterImport_type(MojoParser.Import_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#import_type}.
	 * @param ctx the parse tree
	 */
	void exitImport_type(MojoParser.Import_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#constant_declaration}.
	 * @param ctx the parse tree
	 */
	void enterConstant_declaration(MojoParser.Constant_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#constant_declaration}.
	 * @param ctx the parse tree
	 */
	void exitConstant_declaration(MojoParser.Constant_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#pattern_initializer_list}.
	 * @param ctx the parse tree
	 */
	void enterPattern_initializer_list(MojoParser.Pattern_initializer_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#pattern_initializer_list}.
	 * @param ctx the parse tree
	 */
	void exitPattern_initializer_list(MojoParser.Pattern_initializer_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#pattern_initializer}.
	 * @param ctx the parse tree
	 */
	void enterPattern_initializer(MojoParser.Pattern_initializerContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#pattern_initializer}.
	 * @param ctx the parse tree
	 */
	void exitPattern_initializer(MojoParser.Pattern_initializerContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#initializer}.
	 * @param ctx the parse tree
	 */
	void enterInitializer(MojoParser.InitializerContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#initializer}.
	 * @param ctx the parse tree
	 */
	void exitInitializer(MojoParser.InitializerContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#typealias_declaration}.
	 * @param ctx the parse tree
	 */
	void enterTypealias_declaration(MojoParser.Typealias_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#typealias_declaration}.
	 * @param ctx the parse tree
	 */
	void exitTypealias_declaration(MojoParser.Typealias_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#typealias_name}.
	 * @param ctx the parse tree
	 */
	void enterTypealias_name(MojoParser.Typealias_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#typealias_name}.
	 * @param ctx the parse tree
	 */
	void exitTypealias_name(MojoParser.Typealias_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#typealias_assignment}.
	 * @param ctx the parse tree
	 */
	void enterTypealias_assignment(MojoParser.Typealias_assignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#typealias_assignment}.
	 * @param ctx the parse tree
	 */
	void exitTypealias_assignment(MojoParser.Typealias_assignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_declaration}.
	 * @param ctx the parse tree
	 */
	void enterFunction_declaration(MojoParser.Function_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_declaration}.
	 * @param ctx the parse tree
	 */
	void exitFunction_declaration(MojoParser.Function_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_head}.
	 * @param ctx the parse tree
	 */
	void enterFunction_head(MojoParser.Function_headContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_head}.
	 * @param ctx the parse tree
	 */
	void exitFunction_head(MojoParser.Function_headContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_name}.
	 * @param ctx the parse tree
	 */
	void enterFunction_name(MojoParser.Function_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_name}.
	 * @param ctx the parse tree
	 */
	void exitFunction_name(MojoParser.Function_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_signature}.
	 * @param ctx the parse tree
	 */
	void enterFunction_signature(MojoParser.Function_signatureContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_signature}.
	 * @param ctx the parse tree
	 */
	void exitFunction_signature(MojoParser.Function_signatureContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_result}.
	 * @param ctx the parse tree
	 */
	void enterFunction_result(MojoParser.Function_resultContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_result}.
	 * @param ctx the parse tree
	 */
	void exitFunction_result(MojoParser.Function_resultContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_body}.
	 * @param ctx the parse tree
	 */
	void enterFunction_body(MojoParser.Function_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_body}.
	 * @param ctx the parse tree
	 */
	void exitFunction_body(MojoParser.Function_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_parameter_clause}.
	 * @param ctx the parse tree
	 */
	void enterFunction_parameter_clause(MojoParser.Function_parameter_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_parameter_clause}.
	 * @param ctx the parse tree
	 */
	void exitFunction_parameter_clause(MojoParser.Function_parameter_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_parameter_list}.
	 * @param ctx the parse tree
	 */
	void enterFunction_parameter_list(MojoParser.Function_parameter_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_parameter_list}.
	 * @param ctx the parse tree
	 */
	void exitFunction_parameter_list(MojoParser.Function_parameter_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_parameter}.
	 * @param ctx the parse tree
	 */
	void enterFunction_parameter(MojoParser.Function_parameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_parameter}.
	 * @param ctx the parse tree
	 */
	void exitFunction_parameter(MojoParser.Function_parameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#enum_declaration}.
	 * @param ctx the parse tree
	 */
	void enterEnum_declaration(MojoParser.Enum_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#enum_declaration}.
	 * @param ctx the parse tree
	 */
	void exitEnum_declaration(MojoParser.Enum_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#enum_name}.
	 * @param ctx the parse tree
	 */
	void enterEnum_name(MojoParser.Enum_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#enum_name}.
	 * @param ctx the parse tree
	 */
	void exitEnum_name(MojoParser.Enum_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#enum_members}.
	 * @param ctx the parse tree
	 */
	void enterEnum_members(MojoParser.Enum_membersContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#enum_members}.
	 * @param ctx the parse tree
	 */
	void exitEnum_members(MojoParser.Enum_membersContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#enum_member}.
	 * @param ctx the parse tree
	 */
	void enterEnum_member(MojoParser.Enum_memberContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#enum_member}.
	 * @param ctx the parse tree
	 */
	void exitEnum_member(MojoParser.Enum_memberContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#struct_declaration}.
	 * @param ctx the parse tree
	 */
	void enterStruct_declaration(MojoParser.Struct_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#struct_declaration}.
	 * @param ctx the parse tree
	 */
	void exitStruct_declaration(MojoParser.Struct_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#struct_name}.
	 * @param ctx the parse tree
	 */
	void enterStruct_name(MojoParser.Struct_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#struct_name}.
	 * @param ctx the parse tree
	 */
	void exitStruct_name(MojoParser.Struct_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#struct_body}.
	 * @param ctx the parse tree
	 */
	void enterStruct_body(MojoParser.Struct_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#struct_body}.
	 * @param ctx the parse tree
	 */
	void exitStruct_body(MojoParser.Struct_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#struct_member}.
	 * @param ctx the parse tree
	 */
	void enterStruct_member(MojoParser.Struct_memberContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#struct_member}.
	 * @param ctx the parse tree
	 */
	void exitStruct_member(MojoParser.Struct_memberContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#struct_member_declaration}.
	 * @param ctx the parse tree
	 */
	void enterStruct_member_declaration(MojoParser.Struct_member_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#struct_member_declaration}.
	 * @param ctx the parse tree
	 */
	void exitStruct_member_declaration(MojoParser.Struct_member_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#interface_declaration}.
	 * @param ctx the parse tree
	 */
	void enterInterface_declaration(MojoParser.Interface_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#interface_declaration}.
	 * @param ctx the parse tree
	 */
	void exitInterface_declaration(MojoParser.Interface_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#interface_name}.
	 * @param ctx the parse tree
	 */
	void enterInterface_name(MojoParser.Interface_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#interface_name}.
	 * @param ctx the parse tree
	 */
	void exitInterface_name(MojoParser.Interface_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#interface_body}.
	 * @param ctx the parse tree
	 */
	void enterInterface_body(MojoParser.Interface_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#interface_body}.
	 * @param ctx the parse tree
	 */
	void exitInterface_body(MojoParser.Interface_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#interface_member}.
	 * @param ctx the parse tree
	 */
	void enterInterface_member(MojoParser.Interface_memberContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#interface_member}.
	 * @param ctx the parse tree
	 */
	void exitInterface_member(MojoParser.Interface_memberContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#interface_method_declaration}.
	 * @param ctx the parse tree
	 */
	void enterInterface_method_declaration(MojoParser.Interface_method_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#interface_method_declaration}.
	 * @param ctx the parse tree
	 */
	void exitInterface_method_declaration(MojoParser.Interface_method_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#pattern}.
	 * @param ctx the parse tree
	 */
	void enterPattern(MojoParser.PatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#pattern}.
	 * @param ctx the parse tree
	 */
	void exitPattern(MojoParser.PatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#wildcard_pattern}.
	 * @param ctx the parse tree
	 */
	void enterWildcard_pattern(MojoParser.Wildcard_patternContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#wildcard_pattern}.
	 * @param ctx the parse tree
	 */
	void exitWildcard_pattern(MojoParser.Wildcard_patternContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#identifier_pattern}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier_pattern(MojoParser.Identifier_patternContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#identifier_pattern}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier_pattern(MojoParser.Identifier_patternContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_pattern}.
	 * @param ctx the parse tree
	 */
	void enterTuple_pattern(MojoParser.Tuple_patternContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_pattern}.
	 * @param ctx the parse tree
	 */
	void exitTuple_pattern(MojoParser.Tuple_patternContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_pattern_element_list}.
	 * @param ctx the parse tree
	 */
	void enterTuple_pattern_element_list(MojoParser.Tuple_pattern_element_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_pattern_element_list}.
	 * @param ctx the parse tree
	 */
	void exitTuple_pattern_element_list(MojoParser.Tuple_pattern_element_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_pattern_element}.
	 * @param ctx the parse tree
	 */
	void enterTuple_pattern_element(MojoParser.Tuple_pattern_elementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_pattern_element}.
	 * @param ctx the parse tree
	 */
	void exitTuple_pattern_element(MojoParser.Tuple_pattern_elementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#optional_pattern}.
	 * @param ctx the parse tree
	 */
	void enterOptional_pattern(MojoParser.Optional_patternContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#optional_pattern}.
	 * @param ctx the parse tree
	 */
	void exitOptional_pattern(MojoParser.Optional_patternContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#expression_pattern}.
	 * @param ctx the parse tree
	 */
	void enterExpression_pattern(MojoParser.Expression_patternContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#expression_pattern}.
	 * @param ctx the parse tree
	 */
	void exitExpression_pattern(MojoParser.Expression_patternContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#attribute}.
	 * @param ctx the parse tree
	 */
	void enterAttribute(MojoParser.AttributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#attribute}.
	 * @param ctx the parse tree
	 */
	void exitAttribute(MojoParser.AttributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#attribute_name}.
	 * @param ctx the parse tree
	 */
	void enterAttribute_name(MojoParser.Attribute_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#attribute_name}.
	 * @param ctx the parse tree
	 */
	void exitAttribute_name(MojoParser.Attribute_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#attribute_argument_clause}.
	 * @param ctx the parse tree
	 */
	void enterAttribute_argument_clause(MojoParser.Attribute_argument_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#attribute_argument_clause}.
	 * @param ctx the parse tree
	 */
	void exitAttribute_argument_clause(MojoParser.Attribute_argument_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#attributes}.
	 * @param ctx the parse tree
	 */
	void enterAttributes(MojoParser.AttributesContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#attributes}.
	 * @param ctx the parse tree
	 */
	void exitAttributes(MojoParser.AttributesContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(MojoParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(MojoParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#expression_list}.
	 * @param ctx the parse tree
	 */
	void enterExpression_list(MojoParser.Expression_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#expression_list}.
	 * @param ctx the parse tree
	 */
	void exitExpression_list(MojoParser.Expression_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#prefix_expression}.
	 * @param ctx the parse tree
	 */
	void enterPrefix_expression(MojoParser.Prefix_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#prefix_expression}.
	 * @param ctx the parse tree
	 */
	void exitPrefix_expression(MojoParser.Prefix_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#binary_expression}.
	 * @param ctx the parse tree
	 */
	void enterBinary_expression(MojoParser.Binary_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#binary_expression}.
	 * @param ctx the parse tree
	 */
	void exitBinary_expression(MojoParser.Binary_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#binary_expressions}.
	 * @param ctx the parse tree
	 */
	void enterBinary_expressions(MojoParser.Binary_expressionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#binary_expressions}.
	 * @param ctx the parse tree
	 */
	void exitBinary_expressions(MojoParser.Binary_expressionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#conditional_operator}.
	 * @param ctx the parse tree
	 */
	void enterConditional_operator(MojoParser.Conditional_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#conditional_operator}.
	 * @param ctx the parse tree
	 */
	void exitConditional_operator(MojoParser.Conditional_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_casting_operator}.
	 * @param ctx the parse tree
	 */
	void enterType_casting_operator(MojoParser.Type_casting_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_casting_operator}.
	 * @param ctx the parse tree
	 */
	void exitType_casting_operator(MojoParser.Type_casting_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#primary_expression}.
	 * @param ctx the parse tree
	 */
	void enterPrimary_expression(MojoParser.Primary_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#primary_expression}.
	 * @param ctx the parse tree
	 */
	void exitPrimary_expression(MojoParser.Primary_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#literal_expression}.
	 * @param ctx the parse tree
	 */
	void enterLiteral_expression(MojoParser.Literal_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#literal_expression}.
	 * @param ctx the parse tree
	 */
	void exitLiteral_expression(MojoParser.Literal_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#array_literal}.
	 * @param ctx the parse tree
	 */
	void enterArray_literal(MojoParser.Array_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#array_literal}.
	 * @param ctx the parse tree
	 */
	void exitArray_literal(MojoParser.Array_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#array_literal_items}.
	 * @param ctx the parse tree
	 */
	void enterArray_literal_items(MojoParser.Array_literal_itemsContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#array_literal_items}.
	 * @param ctx the parse tree
	 */
	void exitArray_literal_items(MojoParser.Array_literal_itemsContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#array_literal_item}.
	 * @param ctx the parse tree
	 */
	void enterArray_literal_item(MojoParser.Array_literal_itemContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#array_literal_item}.
	 * @param ctx the parse tree
	 */
	void exitArray_literal_item(MojoParser.Array_literal_itemContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#object_literal}.
	 * @param ctx the parse tree
	 */
	void enterObject_literal(MojoParser.Object_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#object_literal}.
	 * @param ctx the parse tree
	 */
	void exitObject_literal(MojoParser.Object_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#object_literal_items}.
	 * @param ctx the parse tree
	 */
	void enterObject_literal_items(MojoParser.Object_literal_itemsContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#object_literal_items}.
	 * @param ctx the parse tree
	 */
	void exitObject_literal_items(MojoParser.Object_literal_itemsContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#object_literal_item}.
	 * @param ctx the parse tree
	 */
	void enterObject_literal_item(MojoParser.Object_literal_itemContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#object_literal_item}.
	 * @param ctx the parse tree
	 */
	void exitObject_literal_item(MojoParser.Object_literal_itemContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#implicit_member_expression}.
	 * @param ctx the parse tree
	 */
	void enterImplicit_member_expression(MojoParser.Implicit_member_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#implicit_member_expression}.
	 * @param ctx the parse tree
	 */
	void exitImplicit_member_expression(MojoParser.Implicit_member_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#parenthesized_expression}.
	 * @param ctx the parse tree
	 */
	void enterParenthesized_expression(MojoParser.Parenthesized_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#parenthesized_expression}.
	 * @param ctx the parse tree
	 */
	void exitParenthesized_expression(MojoParser.Parenthesized_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_expression}.
	 * @param ctx the parse tree
	 */
	void enterTuple_expression(MojoParser.Tuple_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_expression}.
	 * @param ctx the parse tree
	 */
	void exitTuple_expression(MojoParser.Tuple_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_element}.
	 * @param ctx the parse tree
	 */
	void enterTuple_element(MojoParser.Tuple_elementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_element}.
	 * @param ctx the parse tree
	 */
	void exitTuple_element(MojoParser.Tuple_elementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#wildcard_expression}.
	 * @param ctx the parse tree
	 */
	void enterWildcard_expression(MojoParser.Wildcard_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#wildcard_expression}.
	 * @param ctx the parse tree
	 */
	void exitWildcard_expression(MojoParser.Wildcard_expressionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code function_call_expression}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterFunction_call_expression(MojoParser.Function_call_expressionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code function_call_expression}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitFunction_call_expression(MojoParser.Function_call_expressionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code subscript_expression}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterSubscript_expression(MojoParser.Subscript_expressionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code subscript_expression}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitSubscript_expression(MojoParser.Subscript_expressionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code explicit_member_expression1}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterExplicit_member_expression1(MojoParser.Explicit_member_expression1Context ctx);
	/**
	 * Exit a parse tree produced by the {@code explicit_member_expression1}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitExplicit_member_expression1(MojoParser.Explicit_member_expression1Context ctx);
	/**
	 * Enter a parse tree produced by the {@code explicit_member_expression2}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterExplicit_member_expression2(MojoParser.Explicit_member_expression2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code explicit_member_expression2}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitExplicit_member_expression2(MojoParser.Explicit_member_expression2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code explicit_member_expression3}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterExplicit_member_expression3(MojoParser.Explicit_member_expression3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code explicit_member_expression3}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitExplicit_member_expression3(MojoParser.Explicit_member_expression3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code explicit_member_expression4}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterExplicit_member_expression4(MojoParser.Explicit_member_expression4Context ctx);
	/**
	 * Exit a parse tree produced by the {@code explicit_member_expression4}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitExplicit_member_expression4(MojoParser.Explicit_member_expression4Context ctx);
	/**
	 * Enter a parse tree produced by the {@code postfix_operation}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterPostfix_operation(MojoParser.Postfix_operationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code postfix_operation}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitPostfix_operation(MojoParser.Postfix_operationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primary}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void enterPrimary(MojoParser.PrimaryContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primary}
	 * labeled alternative in {@link MojoParser#postfix_expression}.
	 * @param ctx the parse tree
	 */
	void exitPrimary(MojoParser.PrimaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_call_argument_clause}.
	 * @param ctx the parse tree
	 */
	void enterFunction_call_argument_clause(MojoParser.Function_call_argument_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_call_argument_clause}.
	 * @param ctx the parse tree
	 */
	void exitFunction_call_argument_clause(MojoParser.Function_call_argument_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_call_argument_list}.
	 * @param ctx the parse tree
	 */
	void enterFunction_call_argument_list(MojoParser.Function_call_argument_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_call_argument_list}.
	 * @param ctx the parse tree
	 */
	void exitFunction_call_argument_list(MojoParser.Function_call_argument_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_call_argument}.
	 * @param ctx the parse tree
	 */
	void enterFunction_call_argument(MojoParser.Function_call_argumentContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_call_argument}.
	 * @param ctx the parse tree
	 */
	void exitFunction_call_argument(MojoParser.Function_call_argumentContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#argument_name_list}.
	 * @param ctx the parse tree
	 */
	void enterArgument_name_list(MojoParser.Argument_name_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#argument_name_list}.
	 * @param ctx the parse tree
	 */
	void exitArgument_name_list(MojoParser.Argument_name_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#argument_name}.
	 * @param ctx the parse tree
	 */
	void enterArgument_name(MojoParser.Argument_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#argument_name}.
	 * @param ctx the parse tree
	 */
	void exitArgument_name(MojoParser.Argument_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_}.
	 * @param ctx the parse tree
	 */
	void enterType_(MojoParser.Type_Context ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_}.
	 * @param ctx the parse tree
	 */
	void exitType_(MojoParser.Type_Context ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_annotation}.
	 * @param ctx the parse tree
	 */
	void enterType_annotation(MojoParser.Type_annotationContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_annotation}.
	 * @param ctx the parse tree
	 */
	void exitType_annotation(MojoParser.Type_annotationContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_identifier}.
	 * @param ctx the parse tree
	 */
	void enterType_identifier(MojoParser.Type_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_identifier}.
	 * @param ctx the parse tree
	 */
	void exitType_identifier(MojoParser.Type_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_identifier_clause}.
	 * @param ctx the parse tree
	 */
	void enterType_identifier_clause(MojoParser.Type_identifier_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_identifier_clause}.
	 * @param ctx the parse tree
	 */
	void exitType_identifier_clause(MojoParser.Type_identifier_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_name}.
	 * @param ctx the parse tree
	 */
	void enterType_name(MojoParser.Type_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_name}.
	 * @param ctx the parse tree
	 */
	void exitType_name(MojoParser.Type_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_type}.
	 * @param ctx the parse tree
	 */
	void enterTuple_type(MojoParser.Tuple_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_type}.
	 * @param ctx the parse tree
	 */
	void exitTuple_type(MojoParser.Tuple_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_type_element_list}.
	 * @param ctx the parse tree
	 */
	void enterTuple_type_element_list(MojoParser.Tuple_type_element_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_type_element_list}.
	 * @param ctx the parse tree
	 */
	void exitTuple_type_element_list(MojoParser.Tuple_type_element_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#tuple_type_element}.
	 * @param ctx the parse tree
	 */
	void enterTuple_type_element(MojoParser.Tuple_type_elementContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#tuple_type_element}.
	 * @param ctx the parse tree
	 */
	void exitTuple_type_element(MojoParser.Tuple_type_elementContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_type}.
	 * @param ctx the parse tree
	 */
	void enterFunction_type(MojoParser.Function_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_type}.
	 * @param ctx the parse tree
	 */
	void exitFunction_type(MojoParser.Function_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_type_argument_clause}.
	 * @param ctx the parse tree
	 */
	void enterFunction_type_argument_clause(MojoParser.Function_type_argument_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_type_argument_clause}.
	 * @param ctx the parse tree
	 */
	void exitFunction_type_argument_clause(MojoParser.Function_type_argument_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_type_argument_list}.
	 * @param ctx the parse tree
	 */
	void enterFunction_type_argument_list(MojoParser.Function_type_argument_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_type_argument_list}.
	 * @param ctx the parse tree
	 */
	void exitFunction_type_argument_list(MojoParser.Function_type_argument_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#function_type_argument}.
	 * @param ctx the parse tree
	 */
	void enterFunction_type_argument(MojoParser.Function_type_argumentContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#function_type_argument}.
	 * @param ctx the parse tree
	 */
	void exitFunction_type_argument(MojoParser.Function_type_argumentContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#argument_label}.
	 * @param ctx the parse tree
	 */
	void enterArgument_label(MojoParser.Argument_labelContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#argument_label}.
	 * @param ctx the parse tree
	 */
	void exitArgument_label(MojoParser.Argument_labelContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#array_type}.
	 * @param ctx the parse tree
	 */
	void enterArray_type(MojoParser.Array_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#array_type}.
	 * @param ctx the parse tree
	 */
	void exitArray_type(MojoParser.Array_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#map_type}.
	 * @param ctx the parse tree
	 */
	void enterMap_type(MojoParser.Map_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#map_type}.
	 * @param ctx the parse tree
	 */
	void exitMap_type(MojoParser.Map_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_inheritance_clause}.
	 * @param ctx the parse tree
	 */
	void enterType_inheritance_clause(MojoParser.Type_inheritance_clauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_inheritance_clause}.
	 * @param ctx the parse tree
	 */
	void exitType_inheritance_clause(MojoParser.Type_inheritance_clauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#type_inheritance_list}.
	 * @param ctx the parse tree
	 */
	void enterType_inheritance_list(MojoParser.Type_inheritance_listContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#type_inheritance_list}.
	 * @param ctx the parse tree
	 */
	void exitType_inheritance_list(MojoParser.Type_inheritance_listContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#declaration_identifier}.
	 * @param ctx the parse tree
	 */
	void enterDeclaration_identifier(MojoParser.Declaration_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#declaration_identifier}.
	 * @param ctx the parse tree
	 */
	void exitDeclaration_identifier(MojoParser.Declaration_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#label_identifier}.
	 * @param ctx the parse tree
	 */
	void enterLabel_identifier(MojoParser.Label_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#label_identifier}.
	 * @param ctx the parse tree
	 */
	void exitLabel_identifier(MojoParser.Label_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#path_identifier}.
	 * @param ctx the parse tree
	 */
	void enterPath_identifier(MojoParser.Path_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#path_identifier}.
	 * @param ctx the parse tree
	 */
	void exitPath_identifier(MojoParser.Path_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(MojoParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(MojoParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#keyword_as_identifier_in_declarations}.
	 * @param ctx the parse tree
	 */
	void enterKeyword_as_identifier_in_declarations(MojoParser.Keyword_as_identifier_in_declarationsContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#keyword_as_identifier_in_declarations}.
	 * @param ctx the parse tree
	 */
	void exitKeyword_as_identifier_in_declarations(MojoParser.Keyword_as_identifier_in_declarationsContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#keyword_as_identifier_in_labels}.
	 * @param ctx the parse tree
	 */
	void enterKeyword_as_identifier_in_labels(MojoParser.Keyword_as_identifier_in_labelsContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#keyword_as_identifier_in_labels}.
	 * @param ctx the parse tree
	 */
	void exitKeyword_as_identifier_in_labels(MojoParser.Keyword_as_identifier_in_labelsContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#document}.
	 * @param ctx the parse tree
	 */
	void enterDocument(MojoParser.DocumentContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#document}.
	 * @param ctx the parse tree
	 */
	void exitDocument(MojoParser.DocumentContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#following_document}.
	 * @param ctx the parse tree
	 */
	void enterFollowing_document(MojoParser.Following_documentContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#following_document}.
	 * @param ctx the parse tree
	 */
	void exitFollowing_document(MojoParser.Following_documentContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#assignment_operator}.
	 * @param ctx the parse tree
	 */
	void enterAssignment_operator(MojoParser.Assignment_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#assignment_operator}.
	 * @param ctx the parse tree
	 */
	void exitAssignment_operator(MojoParser.Assignment_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#negate_prefix_operator}.
	 * @param ctx the parse tree
	 */
	void enterNegate_prefix_operator(MojoParser.Negate_prefix_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#negate_prefix_operator}.
	 * @param ctx the parse tree
	 */
	void exitNegate_prefix_operator(MojoParser.Negate_prefix_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#compilation_condition_AND}.
	 * @param ctx the parse tree
	 */
	void enterCompilation_condition_AND(MojoParser.Compilation_condition_ANDContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#compilation_condition_AND}.
	 * @param ctx the parse tree
	 */
	void exitCompilation_condition_AND(MojoParser.Compilation_condition_ANDContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#compilation_condition_OR}.
	 * @param ctx the parse tree
	 */
	void enterCompilation_condition_OR(MojoParser.Compilation_condition_ORContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#compilation_condition_OR}.
	 * @param ctx the parse tree
	 */
	void exitCompilation_condition_OR(MojoParser.Compilation_condition_ORContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#compilation_condition_GE}.
	 * @param ctx the parse tree
	 */
	void enterCompilation_condition_GE(MojoParser.Compilation_condition_GEContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#compilation_condition_GE}.
	 * @param ctx the parse tree
	 */
	void exitCompilation_condition_GE(MojoParser.Compilation_condition_GEContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#arrow_operator}.
	 * @param ctx the parse tree
	 */
	void enterArrow_operator(MojoParser.Arrow_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#arrow_operator}.
	 * @param ctx the parse tree
	 */
	void exitArrow_operator(MojoParser.Arrow_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#range_operator}.
	 * @param ctx the parse tree
	 */
	void enterRange_operator(MojoParser.Range_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#range_operator}.
	 * @param ctx the parse tree
	 */
	void exitRange_operator(MojoParser.Range_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#same_type_equals}.
	 * @param ctx the parse tree
	 */
	void enterSame_type_equals(MojoParser.Same_type_equalsContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#same_type_equals}.
	 * @param ctx the parse tree
	 */
	void exitSame_type_equals(MojoParser.Same_type_equalsContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#binary_operator}.
	 * @param ctx the parse tree
	 */
	void enterBinary_operator(MojoParser.Binary_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#binary_operator}.
	 * @param ctx the parse tree
	 */
	void exitBinary_operator(MojoParser.Binary_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#prefix_operator}.
	 * @param ctx the parse tree
	 */
	void enterPrefix_operator(MojoParser.Prefix_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#prefix_operator}.
	 * @param ctx the parse tree
	 */
	void exitPrefix_operator(MojoParser.Prefix_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#postfix_operator}.
	 * @param ctx the parse tree
	 */
	void enterPostfix_operator(MojoParser.Postfix_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#postfix_operator}.
	 * @param ctx the parse tree
	 */
	void exitPostfix_operator(MojoParser.Postfix_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#operator}.
	 * @param ctx the parse tree
	 */
	void enterOperator(MojoParser.OperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#operator}.
	 * @param ctx the parse tree
	 */
	void exitOperator(MojoParser.OperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#operator_character}.
	 * @param ctx the parse tree
	 */
	void enterOperator_character(MojoParser.Operator_characterContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#operator_character}.
	 * @param ctx the parse tree
	 */
	void exitOperator_character(MojoParser.Operator_characterContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#operator_head}.
	 * @param ctx the parse tree
	 */
	void enterOperator_head(MojoParser.Operator_headContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#operator_head}.
	 * @param ctx the parse tree
	 */
	void exitOperator_head(MojoParser.Operator_headContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#dot_operator_head}.
	 * @param ctx the parse tree
	 */
	void enterDot_operator_head(MojoParser.Dot_operator_headContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#dot_operator_head}.
	 * @param ctx the parse tree
	 */
	void exitDot_operator_head(MojoParser.Dot_operator_headContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#dot_operator_character}.
	 * @param ctx the parse tree
	 */
	void enterDot_operator_character(MojoParser.Dot_operator_characterContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#dot_operator_character}.
	 * @param ctx the parse tree
	 */
	void exitDot_operator_character(MojoParser.Dot_operator_characterContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(MojoParser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(MojoParser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#numeric_literal}.
	 * @param ctx the parse tree
	 */
	void enterNumeric_literal(MojoParser.Numeric_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#numeric_literal}.
	 * @param ctx the parse tree
	 */
	void exitNumeric_literal(MojoParser.Numeric_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#boolean_literal}.
	 * @param ctx the parse tree
	 */
	void enterBoolean_literal(MojoParser.Boolean_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#boolean_literal}.
	 * @param ctx the parse tree
	 */
	void exitBoolean_literal(MojoParser.Boolean_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#null_literal}.
	 * @param ctx the parse tree
	 */
	void enterNull_literal(MojoParser.Null_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#null_literal}.
	 * @param ctx the parse tree
	 */
	void exitNull_literal(MojoParser.Null_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#integer_literal}.
	 * @param ctx the parse tree
	 */
	void enterInteger_literal(MojoParser.Integer_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#integer_literal}.
	 * @param ctx the parse tree
	 */
	void exitInteger_literal(MojoParser.Integer_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link MojoParser#string_literal}.
	 * @param ctx the parse tree
	 */
	void enterString_literal(MojoParser.String_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link MojoParser#string_literal}.
	 * @param ctx the parse tree
	 */
	void exitString_literal(MojoParser.String_literalContext ctx);
}