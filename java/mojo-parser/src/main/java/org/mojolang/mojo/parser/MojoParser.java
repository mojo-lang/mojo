// Generated from /Users/frankee/Projects/mojo/mojo/antlr/Mojo.g4 by ANTLR 4.8
package org.mojolang.mojo.parser;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MojoParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		Type_name=32, Identifier=33, DOT=34, LCURLY=35, LPAREN=36, LBRACK=37, 
		RCURLY=38, RPAREN=39, RBRACK=40, COMMA=41, COLON=42, SEMI=43, LT=44, GT=45, 
		UNDERSCORE=46, BANG=47, QUESTION=48, AT=49, AND=50, SUB=51, EQUAL=52, 
		OR=53, DIV=54, ADD=55, MUL=56, MOD=57, CARET=58, TILDE=59, ELLIPSIS=60, 
		Operator_head_other=61, Operator_following_character=62, Implicit_parameter_name=63, 
		Binary_literal=64, Octal_literal=65, Decimal_literal=66, Pure_decimal_digits=67, 
		Hexadecimal_literal=68, Floating_point_literal=69, Static_string_literal=70, 
		Interpolated_string_literal=71, WS=72, Block_comment=73, Line_comment=74, 
		Line_document=75, Following_line_document=76;
	public static final int
		RULE_source_file = 0, RULE_statement = 1, RULE_loop_statement = 2, RULE_for_in_statement = 3, 
		RULE_while_statement = 4, RULE_condition_list = 5, RULE_condition = 6, 
		RULE_optional_binding_condition = 7, RULE_repeat_while_statement = 8, 
		RULE_branch_statement = 9, RULE_if_statement = 10, RULE_else_clause = 11, 
		RULE_match_statement = 12, RULE_match_cases = 13, RULE_match_case = 14, 
		RULE_control_transfer_statement = 15, RULE_break_statement = 16, RULE_continue_statement = 17, 
		RULE_return_statement = 18, RULE_generic_parameter_clause = 19, RULE_generic_parameter_list = 20, 
		RULE_generic_parameter = 21, RULE_generic_argument_clause = 22, RULE_generic_argument_list = 23, 
		RULE_generic_argument = 24, RULE_declaration = 25, RULE_code_block = 26, 
		RULE_package_declaration = 27, RULE_package_identifier = 28, RULE_import_declaration = 29, 
		RULE_import_path = 30, RULE_import_path_identifier = 31, RULE_import_identifier_as_clause = 32, 
		RULE_import_type_as_clause = 33, RULE_import_type_clause = 34, RULE_import_identifier_group_clause = 35, 
		RULE_import_identifier_group = 36, RULE_import_identifier = 37, RULE_import_type = 38, 
		RULE_constant_declaration = 39, RULE_pattern_initializer_list = 40, RULE_pattern_initializer = 41, 
		RULE_initializer = 42, RULE_typealias_declaration = 43, RULE_typealias_name = 44, 
		RULE_typealias_assignment = 45, RULE_function_declaration = 46, RULE_function_head = 47, 
		RULE_function_name = 48, RULE_function_signature = 49, RULE_function_result = 50, 
		RULE_function_body = 51, RULE_function_parameter_clause = 52, RULE_function_parameter_list = 53, 
		RULE_function_parameter = 54, RULE_enum_declaration = 55, RULE_enum_name = 56, 
		RULE_enum_members = 57, RULE_enum_member = 58, RULE_struct_declaration = 59, 
		RULE_struct_name = 60, RULE_struct_body = 61, RULE_struct_member = 62, 
		RULE_struct_member_declaration = 63, RULE_interface_declaration = 64, 
		RULE_interface_name = 65, RULE_interface_body = 66, RULE_interface_member = 67, 
		RULE_interface_method_declaration = 68, RULE_pattern = 69, RULE_wildcard_pattern = 70, 
		RULE_identifier_pattern = 71, RULE_tuple_pattern = 72, RULE_tuple_pattern_element_list = 73, 
		RULE_tuple_pattern_element = 74, RULE_optional_pattern = 75, RULE_expression_pattern = 76, 
		RULE_attribute = 77, RULE_attribute_name = 78, RULE_attribute_argument_clause = 79, 
		RULE_attributes = 80, RULE_expression = 81, RULE_expression_list = 82, 
		RULE_prefix_expression = 83, RULE_binary_expression = 84, RULE_binary_expressions = 85, 
		RULE_conditional_operator = 86, RULE_type_casting_operator = 87, RULE_primary_expression = 88, 
		RULE_literal_expression = 89, RULE_array_literal = 90, RULE_array_literal_items = 91, 
		RULE_array_literal_item = 92, RULE_object_literal = 93, RULE_object_literal_items = 94, 
		RULE_object_literal_item = 95, RULE_implicit_member_expression = 96, RULE_parenthesized_expression = 97, 
		RULE_tuple_expression = 98, RULE_tuple_element = 99, RULE_wildcard_expression = 100, 
		RULE_postfix_expression = 101, RULE_function_call_argument_clause = 102, 
		RULE_function_call_argument_list = 103, RULE_function_call_argument = 104, 
		RULE_argument_name_list = 105, RULE_argument_name = 106, RULE_type_ = 107, 
		RULE_type_annotation = 108, RULE_type_identifier = 109, RULE_type_identifier_clause = 110, 
		RULE_type_name = 111, RULE_tuple_type = 112, RULE_tuple_type_element_list = 113, 
		RULE_tuple_type_element = 114, RULE_function_type = 115, RULE_function_type_argument_clause = 116, 
		RULE_function_type_argument_list = 117, RULE_function_type_argument = 118, 
		RULE_argument_label = 119, RULE_array_type = 120, RULE_dictionary_type = 121, 
		RULE_type_inheritance_clause = 122, RULE_type_inheritance_list = 123, 
		RULE_declaration_identifier = 124, RULE_label_identifier = 125, RULE_path_identifier = 126, 
		RULE_identifier = 127, RULE_keyword_as_identifier_in_declarations = 128, 
		RULE_keyword_as_identifier_in_labels = 129, RULE_document = 130, RULE_following_document = 131, 
		RULE_assignment_operator = 132, RULE_negate_prefix_operator = 133, RULE_compilation_condition_AND = 134, 
		RULE_compilation_condition_OR = 135, RULE_compilation_condition_GE = 136, 
		RULE_arrow_operator = 137, RULE_range_operator = 138, RULE_same_type_equals = 139, 
		RULE_binary_operator = 140, RULE_prefix_operator = 141, RULE_postfix_operator = 142, 
		RULE_operator = 143, RULE_operator_character = 144, RULE_operator_head = 145, 
		RULE_dot_operator_head = 146, RULE_dot_operator_character = 147, RULE_literal = 148, 
		RULE_numeric_literal = 149, RULE_boolean_literal = 150, RULE_null_literal = 151, 
		RULE_integer_literal = 152, RULE_string_literal = 153;
	private static String[] makeRuleNames() {
		return new String[] {
			"source_file", "statement", "loop_statement", "for_in_statement", "while_statement", 
			"condition_list", "condition", "optional_binding_condition", "repeat_while_statement", 
			"branch_statement", "if_statement", "else_clause", "match_statement", 
			"match_cases", "match_case", "control_transfer_statement", "break_statement", 
			"continue_statement", "return_statement", "generic_parameter_clause", 
			"generic_parameter_list", "generic_parameter", "generic_argument_clause", 
			"generic_argument_list", "generic_argument", "declaration", "code_block", 
			"package_declaration", "package_identifier", "import_declaration", "import_path", 
			"import_path_identifier", "import_identifier_as_clause", "import_type_as_clause", 
			"import_type_clause", "import_identifier_group_clause", "import_identifier_group", 
			"import_identifier", "import_type", "constant_declaration", "pattern_initializer_list", 
			"pattern_initializer", "initializer", "typealias_declaration", "typealias_name", 
			"typealias_assignment", "function_declaration", "function_head", "function_name", 
			"function_signature", "function_result", "function_body", "function_parameter_clause", 
			"function_parameter_list", "function_parameter", "enum_declaration", 
			"enum_name", "enum_members", "enum_member", "struct_declaration", "struct_name", 
			"struct_body", "struct_member", "struct_member_declaration", "interface_declaration", 
			"interface_name", "interface_body", "interface_member", "interface_method_declaration", 
			"pattern", "wildcard_pattern", "identifier_pattern", "tuple_pattern", 
			"tuple_pattern_element_list", "tuple_pattern_element", "optional_pattern", 
			"expression_pattern", "attribute", "attribute_name", "attribute_argument_clause", 
			"attributes", "expression", "expression_list", "prefix_expression", "binary_expression", 
			"binary_expressions", "conditional_operator", "type_casting_operator", 
			"primary_expression", "literal_expression", "array_literal", "array_literal_items", 
			"array_literal_item", "object_literal", "object_literal_items", "object_literal_item", 
			"implicit_member_expression", "parenthesized_expression", "tuple_expression", 
			"tuple_element", "wildcard_expression", "postfix_expression", "function_call_argument_clause", 
			"function_call_argument_list", "function_call_argument", "argument_name_list", 
			"argument_name", "type_", "type_annotation", "type_identifier", "type_identifier_clause", 
			"type_name", "tuple_type", "tuple_type_element_list", "tuple_type_element", 
			"function_type", "function_type_argument_clause", "function_type_argument_list", 
			"function_type_argument", "argument_label", "array_type", "dictionary_type", 
			"type_inheritance_clause", "type_inheritance_list", "declaration_identifier", 
			"label_identifier", "path_identifier", "identifier", "keyword_as_identifier_in_declarations", 
			"keyword_as_identifier_in_labels", "document", "following_document", 
			"assignment_operator", "negate_prefix_operator", "compilation_condition_AND", 
			"compilation_condition_OR", "compilation_condition_GE", "arrow_operator", 
			"range_operator", "same_type_equals", "binary_operator", "prefix_operator", 
			"postfix_operator", "operator", "operator_character", "operator_head", 
			"dot_operator_head", "dot_operator_character", "literal", "numeric_literal", 
			"boolean_literal", "null_literal", "integer_literal", "string_literal"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'for'", "'in'", "'while'", "'let'", "'var'", "'repeat'", "'if'", 
			"'else'", "'match'", "'=>'", "'break'", "'continue'", "'return'", "'package'", 
			"'import'", "'as'", "'type'", "'func'", "'enum'", "'interface'", "'is'", 
			"'and'", "'attribute'", "'const'", "'false'", "'not'", "'null'", "'or'", 
			"'struct'", "'true'", "'xor'", null, null, "'.'", "'{'", "'('", "'['", 
			"'}'", "')'", "']'", "','", "':'", "';'", "'<'", "'>'", "'_'", "'!'", 
			"'?'", "'@'", "'&'", "'-'", "'='", "'|'", "'/'", "'+'", "'*'", "'%'", 
			"'^'", "'~'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "Type_name", "Identifier", 
			"DOT", "LCURLY", "LPAREN", "LBRACK", "RCURLY", "RPAREN", "RBRACK", "COMMA", 
			"COLON", "SEMI", "LT", "GT", "UNDERSCORE", "BANG", "QUESTION", "AT", 
			"AND", "SUB", "EQUAL", "OR", "DIV", "ADD", "MUL", "MOD", "CARET", "TILDE", 
			"ELLIPSIS", "Operator_head_other", "Operator_following_character", "Implicit_parameter_name", 
			"Binary_literal", "Octal_literal", "Decimal_literal", "Pure_decimal_digits", 
			"Hexadecimal_literal", "Floating_point_literal", "Static_string_literal", 
			"Interpolated_string_literal", "WS", "Block_comment", "Line_comment", 
			"Line_document", "Following_line_document"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Mojo.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public MojoParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class Source_fileContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(MojoParser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public Source_fileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_source_file; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterSource_file(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitSource_file(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitSource_file(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Source_fileContext source_file() throws RecognitionException {
		Source_fileContext _localctx = new Source_fileContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_source_file);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(309); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(308);
				statement();
				}
				}
				setState(311); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AT) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)) | (1L << (Line_document - 64)))) != 0) );
			setState(313);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public DeclarationContext declaration() {
			return getRuleContext(DeclarationContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(MojoParser.SEMI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		int _la;
		try {
			setState(323);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
				declaration();
				setState(317);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(316);
					match(SEMI);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(319);
				expression();
				setState(321);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(320);
					match(SEMI);
					}
				}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Loop_statementContext extends ParserRuleContext {
		public For_in_statementContext for_in_statement() {
			return getRuleContext(For_in_statementContext.class,0);
		}
		public While_statementContext while_statement() {
			return getRuleContext(While_statementContext.class,0);
		}
		public Repeat_while_statementContext repeat_while_statement() {
			return getRuleContext(Repeat_while_statementContext.class,0);
		}
		public Loop_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterLoop_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitLoop_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitLoop_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Loop_statementContext loop_statement() throws RecognitionException {
		Loop_statementContext _localctx = new Loop_statementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_loop_statement);
		try {
			setState(329);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EOF:
				enterOuterAlt(_localctx, 1);
				{
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 2);
				{
				setState(326);
				for_in_statement();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 3);
				{
				setState(327);
				while_statement();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(328);
				repeat_while_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class For_in_statementContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public For_in_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_in_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFor_in_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFor_in_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFor_in_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final For_in_statementContext for_in_statement() throws RecognitionException {
		For_in_statementContext _localctx = new For_in_statementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_for_in_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(331);
			match(T__0);
			setState(332);
			pattern(0);
			setState(333);
			match(T__1);
			setState(334);
			expression();
			setState(335);
			code_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class While_statementContext extends ParserRuleContext {
		public Condition_listContext condition_list() {
			return getRuleContext(Condition_listContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public While_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterWhile_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitWhile_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitWhile_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final While_statementContext while_statement() throws RecognitionException {
		While_statementContext _localctx = new While_statementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			match(T__2);
			setState(338);
			condition_list();
			setState(339);
			code_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Condition_listContext extends ParserRuleContext {
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Condition_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterCondition_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitCondition_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitCondition_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Condition_listContext condition_list() throws RecognitionException {
		Condition_listContext _localctx = new Condition_listContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_condition_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(341);
			condition();
			setState(346);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(342);
				match(COMMA);
				setState(343);
				condition();
				}
				}
				setState(348);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Optional_binding_conditionContext optional_binding_condition() {
			return getRuleContext(Optional_binding_conditionContext.class,0);
		}
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterCondition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitCondition(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitCondition(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_condition);
		try {
			setState(351);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(349);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(350);
				optional_binding_condition();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Optional_binding_conditionContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Optional_binding_conditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optional_binding_condition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterOptional_binding_condition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitOptional_binding_condition(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitOptional_binding_condition(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Optional_binding_conditionContext optional_binding_condition() throws RecognitionException {
		Optional_binding_conditionContext _localctx = new Optional_binding_conditionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_optional_binding_condition);
		try {
			setState(361);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(353);
				match(T__3);
				setState(354);
				pattern(0);
				setState(355);
				initializer();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 2);
				{
				setState(357);
				match(T__4);
				setState(358);
				pattern(0);
				setState(359);
				initializer();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Repeat_while_statementContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Repeat_while_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_repeat_while_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterRepeat_while_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitRepeat_while_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitRepeat_while_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Repeat_while_statementContext repeat_while_statement() throws RecognitionException {
		Repeat_while_statementContext _localctx = new Repeat_while_statementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_repeat_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(T__5);
			setState(364);
			code_block();
			setState(365);
			match(T__2);
			setState(366);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Branch_statementContext extends ParserRuleContext {
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public Match_statementContext match_statement() {
			return getRuleContext(Match_statementContext.class,0);
		}
		public Branch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_branch_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterBranch_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitBranch_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitBranch_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Branch_statementContext branch_statement() throws RecognitionException {
		Branch_statementContext _localctx = new Branch_statementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_branch_statement);
		try {
			setState(370);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__6:
				enterOuterAlt(_localctx, 1);
				{
				setState(368);
				if_statement();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 2);
				{
				setState(369);
				match_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class If_statementContext extends ParserRuleContext {
		public Condition_listContext condition_list() {
			return getRuleContext(Condition_listContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Else_clauseContext else_clause() {
			return getRuleContext(Else_clauseContext.class,0);
		}
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterIf_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitIf_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitIf_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(372);
			match(T__6);
			setState(373);
			condition_list();
			setState(374);
			code_block();
			setState(376);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(375);
				else_clause();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Else_clauseContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public Else_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterElse_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitElse_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitElse_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Else_clauseContext else_clause() throws RecognitionException {
		Else_clauseContext _localctx = new Else_clauseContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_else_clause);
		try {
			setState(382);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(378);
				match(T__7);
				setState(379);
				code_block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(380);
				match(T__7);
				setState(381);
				if_statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Match_statementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public Match_casesContext match_cases() {
			return getRuleContext(Match_casesContext.class,0);
		}
		public Match_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_match_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterMatch_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitMatch_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitMatch_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Match_statementContext match_statement() throws RecognitionException {
		Match_statementContext _localctx = new Match_statementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_match_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(384);
			match(T__8);
			setState(385);
			expression();
			setState(386);
			match(LCURLY);
			setState(388);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)))) != 0)) {
				{
				setState(387);
				match_cases();
				}
			}

			setState(390);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Match_casesContext extends ParserRuleContext {
		public Match_caseContext match_case() {
			return getRuleContext(Match_caseContext.class,0);
		}
		public Match_casesContext match_cases() {
			return getRuleContext(Match_casesContext.class,0);
		}
		public Match_casesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_match_cases; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterMatch_cases(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitMatch_cases(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitMatch_cases(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Match_casesContext match_cases() throws RecognitionException {
		Match_casesContext _localctx = new Match_casesContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_match_cases);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(392);
			match_case();
			setState(394);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)))) != 0)) {
				{
				setState(393);
				match_cases();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Match_caseContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Match_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_match_case; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterMatch_case(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitMatch_case(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitMatch_case(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Match_caseContext match_case() throws RecognitionException {
		Match_caseContext _localctx = new Match_caseContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_match_case);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(396);
			pattern(0);
			setState(397);
			match(T__9);
			setState(400);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(398);
				code_block();
				}
				break;
			case 2:
				{
				setState(399);
				expression();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Control_transfer_statementContext extends ParserRuleContext {
		public Break_statementContext break_statement() {
			return getRuleContext(Break_statementContext.class,0);
		}
		public Continue_statementContext continue_statement() {
			return getRuleContext(Continue_statementContext.class,0);
		}
		public Return_statementContext return_statement() {
			return getRuleContext(Return_statementContext.class,0);
		}
		public Control_transfer_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_transfer_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterControl_transfer_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitControl_transfer_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitControl_transfer_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Control_transfer_statementContext control_transfer_statement() throws RecognitionException {
		Control_transfer_statementContext _localctx = new Control_transfer_statementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_control_transfer_statement);
		try {
			setState(405);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__10:
				enterOuterAlt(_localctx, 1);
				{
				setState(402);
				break_statement();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 2);
				{
				setState(403);
				continue_statement();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 3);
				{
				setState(404);
				return_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Break_statementContext extends ParserRuleContext {
		public TerminalNode SEMI() { return getToken(MojoParser.SEMI, 0); }
		public Break_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterBreak_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitBreak_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitBreak_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Break_statementContext break_statement() throws RecognitionException {
		Break_statementContext _localctx = new Break_statementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_break_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(407);
			match(T__10);
			setState(409);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(408);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Continue_statementContext extends ParserRuleContext {
		public TerminalNode SEMI() { return getToken(MojoParser.SEMI, 0); }
		public Continue_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterContinue_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitContinue_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitContinue_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Continue_statementContext continue_statement() throws RecognitionException {
		Continue_statementContext _localctx = new Continue_statementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_continue_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(411);
			match(T__11);
			setState(413);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(412);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Return_statementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Return_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterReturn_statement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitReturn_statement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitReturn_statement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Return_statementContext return_statement() throws RecognitionException {
		Return_statementContext _localctx = new Return_statementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_return_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(415);
			match(T__12);
			setState(417);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)))) != 0)) {
				{
				setState(416);
				expression();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Generic_parameter_clauseContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(MojoParser.LT, 0); }
		public Generic_parameter_listContext generic_parameter_list() {
			return getRuleContext(Generic_parameter_listContext.class,0);
		}
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public Generic_parameter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_parameter_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterGeneric_parameter_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitGeneric_parameter_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitGeneric_parameter_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Generic_parameter_clauseContext generic_parameter_clause() throws RecognitionException {
		Generic_parameter_clauseContext _localctx = new Generic_parameter_clauseContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_generic_parameter_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(419);
			match(LT);
			setState(420);
			generic_parameter_list();
			setState(421);
			match(GT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Generic_parameter_listContext extends ParserRuleContext {
		public List<Generic_parameterContext> generic_parameter() {
			return getRuleContexts(Generic_parameterContext.class);
		}
		public Generic_parameterContext generic_parameter(int i) {
			return getRuleContext(Generic_parameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Generic_parameter_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_parameter_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterGeneric_parameter_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitGeneric_parameter_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitGeneric_parameter_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Generic_parameter_listContext generic_parameter_list() throws RecognitionException {
		Generic_parameter_listContext _localctx = new Generic_parameter_listContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_generic_parameter_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(423);
			generic_parameter();
			setState(428);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(424);
				match(COMMA);
				setState(425);
				generic_parameter();
				}
				}
				setState(430);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Generic_parameterContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Generic_parameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_parameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterGeneric_parameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitGeneric_parameter(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitGeneric_parameter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Generic_parameterContext generic_parameter() throws RecognitionException {
		Generic_parameterContext _localctx = new Generic_parameterContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_generic_parameter);
		try {
			setState(436);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(431);
				type_name();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(432);
				type_name();
				setState(433);
				match(COLON);
				setState(434);
				type_identifier();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Generic_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(MojoParser.LT, 0); }
		public Generic_argument_listContext generic_argument_list() {
			return getRuleContext(Generic_argument_listContext.class,0);
		}
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public Generic_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_argument_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterGeneric_argument_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitGeneric_argument_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitGeneric_argument_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Generic_argument_clauseContext generic_argument_clause() throws RecognitionException {
		Generic_argument_clauseContext _localctx = new Generic_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_generic_argument_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			match(LT);
			setState(439);
			generic_argument_list();
			setState(440);
			match(GT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Generic_argument_listContext extends ParserRuleContext {
		public List<Generic_argumentContext> generic_argument() {
			return getRuleContexts(Generic_argumentContext.class);
		}
		public Generic_argumentContext generic_argument(int i) {
			return getRuleContext(Generic_argumentContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Generic_argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_argument_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterGeneric_argument_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitGeneric_argument_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitGeneric_argument_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Generic_argument_listContext generic_argument_list() throws RecognitionException {
		Generic_argument_listContext _localctx = new Generic_argument_listContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_generic_argument_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(442);
			generic_argument();
			setState(447);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(443);
				match(COMMA);
				setState(444);
				generic_argument();
				}
				}
				setState(449);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Generic_argumentContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Generic_argumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generic_argument; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterGeneric_argument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitGeneric_argument(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitGeneric_argument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Generic_argumentContext generic_argument() throws RecognitionException {
		Generic_argumentContext _localctx = new Generic_argumentContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_generic_argument);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			type_(0);
			setState(452);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(451);
				attributes();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationContext extends ParserRuleContext {
		public Package_declarationContext package_declaration() {
			return getRuleContext(Package_declarationContext.class,0);
		}
		public Import_declarationContext import_declaration() {
			return getRuleContext(Import_declarationContext.class,0);
		}
		public Constant_declarationContext constant_declaration() {
			return getRuleContext(Constant_declarationContext.class,0);
		}
		public Typealias_declarationContext typealias_declaration() {
			return getRuleContext(Typealias_declarationContext.class,0);
		}
		public Function_declarationContext function_declaration() {
			return getRuleContext(Function_declarationContext.class,0);
		}
		public Enum_declarationContext enum_declaration() {
			return getRuleContext(Enum_declarationContext.class,0);
		}
		public Struct_declarationContext struct_declaration() {
			return getRuleContext(Struct_declarationContext.class,0);
		}
		public Interface_declarationContext interface_declaration() {
			return getRuleContext(Interface_declarationContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_declaration);
		try {
			setState(462);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(454);
				package_declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(455);
				import_declaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(456);
				constant_declaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(457);
				typealias_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(458);
				function_declaration();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(459);
				enum_declaration();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(460);
				struct_declaration();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(461);
				interface_declaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Code_blockContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public Code_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_code_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterCode_block(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitCode_block(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitCode_block(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Code_blockContext code_block() throws RecognitionException {
		Code_blockContext _localctx = new Code_blockContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_code_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(464);
			match(LCURLY);
			setState(466); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(465);
				statement();
				}
				}
				setState(468); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AT) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)) | (1L << (Line_document - 64)))) != 0) );
			setState(470);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Package_declarationContext extends ParserRuleContext {
		public Package_identifierContext package_identifier() {
			return getRuleContext(Package_identifierContext.class,0);
		}
		public Object_literalContext object_literal() {
			return getRuleContext(Object_literalContext.class,0);
		}
		public Package_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_package_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPackage_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPackage_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPackage_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Package_declarationContext package_declaration() throws RecognitionException {
		Package_declarationContext _localctx = new Package_declarationContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_package_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(472);
			match(T__13);
			setState(473);
			package_identifier();
			setState(474);
			object_literal();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Package_identifierContext extends ParserRuleContext {
		public List<TerminalNode> Identifier() { return getTokens(MojoParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(MojoParser.Identifier, i);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public Package_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_package_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPackage_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPackage_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPackage_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Package_identifierContext package_identifier() throws RecognitionException {
		Package_identifierContext _localctx = new Package_identifierContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_package_identifier);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(476);
			match(Identifier);
			setState(481);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(477);
					match(DOT);
					setState(478);
					match(Identifier);
					}
					} 
				}
				setState(483);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_declarationContext extends ParserRuleContext {
		public Import_pathContext import_path() {
			return getRuleContext(Import_pathContext.class,0);
		}
		public Import_identifier_as_clauseContext import_identifier_as_clause() {
			return getRuleContext(Import_identifier_as_clauseContext.class,0);
		}
		public Import_type_clauseContext import_type_clause() {
			return getRuleContext(Import_type_clauseContext.class,0);
		}
		public Import_identifier_group_clauseContext import_identifier_group_clause() {
			return getRuleContext(Import_identifier_group_clauseContext.class,0);
		}
		public Import_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_declarationContext import_declaration() throws RecognitionException {
		Import_declarationContext _localctx = new Import_declarationContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_import_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(484);
			match(T__14);
			setState(485);
			import_path();
			setState(489);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				setState(486);
				import_identifier_as_clause();
				}
				break;
			case 2:
				{
				setState(487);
				import_type_clause();
				}
				break;
			case 3:
				{
				setState(488);
				import_identifier_group_clause();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_pathContext extends ParserRuleContext {
		public List<Import_path_identifierContext> import_path_identifier() {
			return getRuleContexts(Import_path_identifierContext.class);
		}
		public Import_path_identifierContext import_path_identifier(int i) {
			return getRuleContext(Import_path_identifierContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public Import_pathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_path; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_path(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_path(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_path(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_pathContext import_path() throws RecognitionException {
		Import_pathContext _localctx = new Import_pathContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_import_path);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(491);
			import_path_identifier();
			setState(496);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(492);
					match(DOT);
					setState(493);
					import_path_identifier();
					}
					} 
				}
				setState(498);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_path_identifierContext extends ParserRuleContext {
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public Import_path_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_path_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_path_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_path_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_path_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_path_identifierContext import_path_identifier() throws RecognitionException {
		Import_path_identifierContext _localctx = new Import_path_identifierContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_import_path_identifier);
		try {
			setState(501);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case T__1:
			case T__2:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(499);
				declaration_identifier();
				}
				break;
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 2);
				{
				setState(500);
				operator();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_identifier_as_clauseContext extends ParserRuleContext {
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public Import_identifier_as_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_identifier_as_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_identifier_as_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_identifier_as_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_identifier_as_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_identifier_as_clauseContext import_identifier_as_clause() throws RecognitionException {
		Import_identifier_as_clauseContext _localctx = new Import_identifier_as_clauseContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_import_identifier_as_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(503);
			match(T__15);
			setState(504);
			declaration_identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_type_as_clauseContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Import_type_as_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_type_as_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_type_as_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_type_as_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_type_as_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_type_as_clauseContext import_type_as_clause() throws RecognitionException {
		Import_type_as_clauseContext _localctx = new Import_type_as_clauseContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_import_type_as_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			match(T__15);
			setState(507);
			type_name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_type_clauseContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Import_type_as_clauseContext import_type_as_clause() {
			return getRuleContext(Import_type_as_clauseContext.class,0);
		}
		public Import_type_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_type_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_type_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_type_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_type_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_type_clauseContext import_type_clause() throws RecognitionException {
		Import_type_clauseContext _localctx = new Import_type_clauseContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_import_type_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(509);
			match(DOT);
			setState(510);
			type_name();
			setState(512);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				{
				setState(511);
				import_type_as_clause();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_identifier_group_clauseContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public Import_identifier_groupContext import_identifier_group() {
			return getRuleContext(Import_identifier_groupContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public Import_identifier_group_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_identifier_group_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_identifier_group_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_identifier_group_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_identifier_group_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_identifier_group_clauseContext import_identifier_group_clause() throws RecognitionException {
		Import_identifier_group_clauseContext _localctx = new Import_identifier_group_clauseContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_import_identifier_group_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(514);
			match(DOT);
			setState(515);
			match(LCURLY);
			setState(516);
			import_identifier_group();
			setState(517);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_identifier_groupContext extends ParserRuleContext {
		public List<Import_identifierContext> import_identifier() {
			return getRuleContexts(Import_identifierContext.class);
		}
		public Import_identifierContext import_identifier(int i) {
			return getRuleContext(Import_identifierContext.class,i);
		}
		public List<Import_typeContext> import_type() {
			return getRuleContexts(Import_typeContext.class);
		}
		public Import_typeContext import_type(int i) {
			return getRuleContext(Import_typeContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Import_identifier_groupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_identifier_group; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_identifier_group(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_identifier_group(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_identifier_group(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_identifier_groupContext import_identifier_group() throws RecognitionException {
		Import_identifier_groupContext _localctx = new Import_identifier_groupContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_import_identifier_group);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(521);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case T__1:
			case T__2:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
			case Identifier:
				{
				setState(519);
				import_identifier();
				}
				break;
			case Type_name:
				{
				setState(520);
				import_type();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(530);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(523);
				match(COMMA);
				setState(526);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__0:
				case T__1:
				case T__2:
				case T__4:
				case T__5:
				case T__6:
				case T__7:
				case T__8:
				case T__10:
				case T__11:
				case T__12:
				case T__13:
				case T__14:
				case T__15:
				case T__16:
				case T__17:
				case T__18:
				case T__19:
				case T__20:
				case T__21:
				case T__22:
				case T__23:
				case T__24:
				case T__25:
				case T__26:
				case T__27:
				case T__28:
				case T__29:
				case T__30:
				case Identifier:
					{
					setState(524);
					import_identifier();
					}
					break;
				case Type_name:
					{
					setState(525);
					import_type();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				}
				setState(532);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_identifierContext extends ParserRuleContext {
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public Import_identifier_as_clauseContext import_identifier_as_clause() {
			return getRuleContext(Import_identifier_as_clauseContext.class,0);
		}
		public Import_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_identifierContext import_identifier() throws RecognitionException {
		Import_identifierContext _localctx = new Import_identifierContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_import_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(533);
			declaration_identifier();
			setState(535);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__15) {
				{
				setState(534);
				import_identifier_as_clause();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_typeContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Import_type_as_clauseContext import_type_as_clause() {
			return getRuleContext(Import_type_as_clauseContext.class,0);
		}
		public Import_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImport_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImport_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImport_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Import_typeContext import_type() throws RecognitionException {
		Import_typeContext _localctx = new Import_typeContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_import_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(537);
			type_name();
			setState(539);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__15) {
				{
				setState(538);
				import_type_as_clause();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Constant_declarationContext extends ParserRuleContext {
		public Pattern_initializer_listContext pattern_initializer_list() {
			return getRuleContext(Pattern_initializer_listContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Constant_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterConstant_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitConstant_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitConstant_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Constant_declarationContext constant_declaration() throws RecognitionException {
		Constant_declarationContext _localctx = new Constant_declarationContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_constant_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(542);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(541);
				attributes();
				}
			}

			setState(544);
			match(T__3);
			setState(545);
			pattern_initializer_list();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Pattern_initializer_listContext extends ParserRuleContext {
		public List<Pattern_initializerContext> pattern_initializer() {
			return getRuleContexts(Pattern_initializerContext.class);
		}
		public Pattern_initializerContext pattern_initializer(int i) {
			return getRuleContext(Pattern_initializerContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Pattern_initializer_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern_initializer_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPattern_initializer_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPattern_initializer_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPattern_initializer_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Pattern_initializer_listContext pattern_initializer_list() throws RecognitionException {
		Pattern_initializer_listContext _localctx = new Pattern_initializer_listContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_pattern_initializer_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(547);
			pattern_initializer();
			setState(552);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(548);
				match(COMMA);
				setState(549);
				pattern_initializer();
				}
				}
				setState(554);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Pattern_initializerContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Pattern_initializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern_initializer; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPattern_initializer(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPattern_initializer(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPattern_initializer(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Pattern_initializerContext pattern_initializer() throws RecognitionException {
		Pattern_initializerContext _localctx = new Pattern_initializerContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_pattern_initializer);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(555);
			pattern(0);
			setState(557);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				setState(556);
				initializer();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InitializerContext extends ParserRuleContext {
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public InitializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInitializer(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInitializer(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInitializer(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InitializerContext initializer() throws RecognitionException {
		InitializerContext _localctx = new InitializerContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_initializer);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(559);
			assignment_operator();
			setState(560);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Typealias_declarationContext extends ParserRuleContext {
		public Typealias_nameContext typealias_name() {
			return getRuleContext(Typealias_nameContext.class,0);
		}
		public Typealias_assignmentContext typealias_assignment() {
			return getRuleContext(Typealias_assignmentContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Typealias_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typealias_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTypealias_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTypealias_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTypealias_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Typealias_declarationContext typealias_declaration() throws RecognitionException {
		Typealias_declarationContext _localctx = new Typealias_declarationContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_typealias_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(563);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(562);
				attributes();
				}
			}

			setState(565);
			match(T__16);
			setState(566);
			typealias_name();
			setState(568);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(567);
				generic_parameter_clause();
				}
			}

			setState(570);
			typealias_assignment();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Typealias_nameContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Typealias_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typealias_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTypealias_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTypealias_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTypealias_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Typealias_nameContext typealias_name() throws RecognitionException {
		Typealias_nameContext _localctx = new Typealias_nameContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_typealias_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(572);
			type_name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Typealias_assignmentContext extends ParserRuleContext {
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Typealias_assignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typealias_assignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTypealias_assignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTypealias_assignment(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTypealias_assignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Typealias_assignmentContext typealias_assignment() throws RecognitionException {
		Typealias_assignmentContext _localctx = new Typealias_assignmentContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_typealias_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(574);
			assignment_operator();
			setState(575);
			type_(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_declarationContext extends ParserRuleContext {
		public Function_headContext function_head() {
			return getRuleContext(Function_headContext.class,0);
		}
		public Function_nameContext function_name() {
			return getRuleContext(Function_nameContext.class,0);
		}
		public Function_signatureContext function_signature() {
			return getRuleContext(Function_signatureContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Function_bodyContext function_body() {
			return getRuleContext(Function_bodyContext.class,0);
		}
		public Function_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_declarationContext function_declaration() throws RecognitionException {
		Function_declarationContext _localctx = new Function_declarationContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_function_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(577);
			function_head();
			setState(578);
			function_name();
			setState(580);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(579);
				generic_parameter_clause();
				}
			}

			setState(582);
			function_signature();
			setState(584);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				{
				setState(583);
				function_body();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_headContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Function_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_head; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_head(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_head(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_head(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_headContext function_head() throws RecognitionException {
		Function_headContext _localctx = new Function_headContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_function_head);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(587);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(586);
				attributes();
				}
			}

			setState(589);
			match(T__17);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_nameContext extends ParserRuleContext {
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public Function_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_nameContext function_name() throws RecognitionException {
		Function_nameContext _localctx = new Function_nameContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_function_name);
		try {
			setState(593);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case T__1:
			case T__2:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(591);
				declaration_identifier();
				}
				break;
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 2);
				{
				setState(592);
				operator();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_signatureContext extends ParserRuleContext {
		public Function_parameter_clauseContext function_parameter_clause() {
			return getRuleContext(Function_parameter_clauseContext.class,0);
		}
		public Function_resultContext function_result() {
			return getRuleContext(Function_resultContext.class,0);
		}
		public Function_signatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_signature; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_signature(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_signature(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_signature(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_signatureContext function_signature() throws RecognitionException {
		Function_signatureContext _localctx = new Function_signatureContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_function_signature);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(595);
			function_parameter_clause();
			setState(597);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				{
				setState(596);
				function_result();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_resultContext extends ParserRuleContext {
		public Arrow_operatorContext arrow_operator() {
			return getRuleContext(Arrow_operatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Function_resultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_result; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_result(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_result(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_result(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_resultContext function_result() throws RecognitionException {
		Function_resultContext _localctx = new Function_resultContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_function_result);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(599);
			arrow_operator();
			setState(600);
			type_(0);
			setState(602);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
			case 1:
				{
				setState(601);
				attributes();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_bodyContext extends ParserRuleContext {
		public Code_blockContext code_block() {
			return getRuleContext(Code_blockContext.class,0);
		}
		public Function_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_body; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_body(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_body(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_body(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_bodyContext function_body() throws RecognitionException {
		Function_bodyContext _localctx = new Function_bodyContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_function_body);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(604);
			code_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_parameter_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Function_parameter_listContext function_parameter_list() {
			return getRuleContext(Function_parameter_listContext.class,0);
		}
		public Function_parameter_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_parameter_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_parameter_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_parameter_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_parameter_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_parameter_clauseContext function_parameter_clause() throws RecognitionException {
		Function_parameter_clauseContext _localctx = new Function_parameter_clauseContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_function_parameter_clause);
		try {
			setState(612);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(606);
				match(LPAREN);
				setState(607);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(608);
				match(LPAREN);
				setState(609);
				function_parameter_list();
				setState(610);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_parameter_listContext extends ParserRuleContext {
		public List<Function_parameterContext> function_parameter() {
			return getRuleContexts(Function_parameterContext.class);
		}
		public Function_parameterContext function_parameter(int i) {
			return getRuleContext(Function_parameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Function_parameter_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_parameter_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_parameter_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_parameter_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_parameter_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_parameter_listContext function_parameter_list() throws RecognitionException {
		Function_parameter_listContext _localctx = new Function_parameter_listContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_function_parameter_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(614);
			function_parameter();
			setState(619);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(615);
				match(COMMA);
				setState(616);
				function_parameter();
				}
				}
				setState(621);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_parameterContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Function_parameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_parameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_parameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_parameter(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_parameter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_parameterContext function_parameter() throws RecognitionException {
		Function_parameterContext _localctx = new Function_parameterContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_function_parameter);
		int _la;
		try {
			setState(633);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(622);
				label_identifier();
				setState(623);
				type_annotation();
				setState(625);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==EQUAL) {
					{
					setState(624);
					initializer();
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(627);
				label_identifier();
				setState(628);
				type_(0);
				setState(629);
				match(ELLIPSIS);
				setState(631);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(630);
					attributes();
					}
				}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Enum_declarationContext extends ParserRuleContext {
		public Enum_nameContext enum_name() {
			return getRuleContext(Enum_nameContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Enum_membersContext enum_members() {
			return getRuleContext(Enum_membersContext.class,0);
		}
		public Enum_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterEnum_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitEnum_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitEnum_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Enum_declarationContext enum_declaration() throws RecognitionException {
		Enum_declarationContext _localctx = new Enum_declarationContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_enum_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(636);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Line_document) {
				{
				setState(635);
				document();
				}
			}

			setState(639);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(638);
				attributes();
				}
			}

			setState(641);
			match(T__18);
			setState(642);
			enum_name();
			setState(644);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(643);
				generic_parameter_clause();
				}
			}

			setState(647);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(646);
				type_inheritance_clause();
				}
			}

			setState(649);
			match(LCURLY);
			setState(651);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier || _la==Line_document) {
				{
				setState(650);
				enum_members();
				}
			}

			setState(653);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Enum_nameContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Enum_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterEnum_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitEnum_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitEnum_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Enum_nameContext enum_name() throws RecognitionException {
		Enum_nameContext _localctx = new Enum_nameContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_enum_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(655);
			type_name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Enum_membersContext extends ParserRuleContext {
		public List<Enum_memberContext> enum_member() {
			return getRuleContexts(Enum_memberContext.class);
		}
		public Enum_memberContext enum_member(int i) {
			return getRuleContext(Enum_memberContext.class,i);
		}
		public Enum_membersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_members; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterEnum_members(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitEnum_members(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitEnum_members(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Enum_membersContext enum_members() throws RecognitionException {
		Enum_membersContext _localctx = new Enum_membersContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_enum_members);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(658); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(657);
				enum_member();
				}
				}
				setState(660); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Identifier || _la==Line_document );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Enum_memberContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MojoParser.Identifier, 0); }
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Following_documentContext following_document() {
			return getRuleContext(Following_documentContext.class,0);
		}
		public Enum_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enum_member; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterEnum_member(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitEnum_member(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitEnum_member(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Enum_memberContext enum_member() throws RecognitionException {
		Enum_memberContext _localctx = new Enum_memberContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_enum_member);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(663);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Line_document) {
				{
				setState(662);
				document();
				}
			}

			setState(665);
			match(Identifier);
			setState(667);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(666);
				attributes();
				}
			}

			setState(670);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(669);
				initializer();
				}
			}

			setState(673);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Following_line_document) {
				{
				setState(672);
				following_document();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Struct_declarationContext extends ParserRuleContext {
		public Struct_nameContext struct_name() {
			return getRuleContext(Struct_nameContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Struct_bodyContext struct_body() {
			return getRuleContext(Struct_bodyContext.class,0);
		}
		public Struct_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterStruct_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitStruct_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitStruct_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Struct_declarationContext struct_declaration() throws RecognitionException {
		Struct_declarationContext _localctx = new Struct_declarationContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_struct_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(676);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Line_document) {
				{
				setState(675);
				document();
				}
			}

			setState(679);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(678);
				attributes();
				}
			}

			setState(681);
			match(T__16);
			setState(682);
			struct_name();
			setState(684);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,62,_ctx) ) {
			case 1:
				{
				setState(683);
				generic_parameter_clause();
				}
				break;
			}
			setState(687);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(686);
				type_inheritance_clause();
				}
			}

			setState(690);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,64,_ctx) ) {
			case 1:
				{
				setState(689);
				struct_body();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Struct_nameContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Struct_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterStruct_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitStruct_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitStruct_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Struct_nameContext struct_name() throws RecognitionException {
		Struct_nameContext _localctx = new Struct_nameContext(_ctx, getState());
		enterRule(_localctx, 120, RULE_struct_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(692);
			type_name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Struct_bodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<Struct_memberContext> struct_member() {
			return getRuleContexts(Struct_memberContext.class);
		}
		public Struct_memberContext struct_member(int i) {
			return getRuleContext(Struct_memberContext.class,i);
		}
		public Struct_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_body; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterStruct_body(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitStruct_body(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitStruct_body(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Struct_bodyContext struct_body() throws RecognitionException {
		Struct_bodyContext _localctx = new Struct_bodyContext(_ctx, getState());
		enterRule(_localctx, 122, RULE_struct_body);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(694);
			match(LCURLY);
			setState(698);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__16) | (1L << T__18) | (1L << Identifier) | (1L << AT))) != 0) || _la==Line_document) {
				{
				{
				setState(695);
				struct_member();
				}
				}
				setState(700);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(701);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Struct_memberContext extends ParserRuleContext {
		public Struct_member_declarationContext struct_member_declaration() {
			return getRuleContext(Struct_member_declarationContext.class,0);
		}
		public Struct_declarationContext struct_declaration() {
			return getRuleContext(Struct_declarationContext.class,0);
		}
		public Enum_declarationContext enum_declaration() {
			return getRuleContext(Enum_declarationContext.class,0);
		}
		public Constant_declarationContext constant_declaration() {
			return getRuleContext(Constant_declarationContext.class,0);
		}
		public Typealias_declarationContext typealias_declaration() {
			return getRuleContext(Typealias_declarationContext.class,0);
		}
		public Struct_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_member; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterStruct_member(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitStruct_member(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitStruct_member(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Struct_memberContext struct_member() throws RecognitionException {
		Struct_memberContext _localctx = new Struct_memberContext(_ctx, getState());
		enterRule(_localctx, 124, RULE_struct_member);
		try {
			setState(708);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,66,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(703);
				struct_member_declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(704);
				struct_declaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(705);
				enum_declaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(706);
				constant_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(707);
				typealias_declaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Struct_member_declarationContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MojoParser.Identifier, 0); }
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public Following_documentContext following_document() {
			return getRuleContext(Following_documentContext.class,0);
		}
		public Struct_member_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_member_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterStruct_member_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitStruct_member_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitStruct_member_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Struct_member_declarationContext struct_member_declaration() throws RecognitionException {
		Struct_member_declarationContext _localctx = new Struct_member_declarationContext(_ctx, getState());
		enterRule(_localctx, 126, RULE_struct_member_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(711);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Line_document) {
				{
				setState(710);
				document();
				}
			}

			setState(713);
			match(Identifier);
			setState(714);
			type_annotation();
			setState(716);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(715);
				initializer();
				}
			}

			setState(719);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Following_line_document) {
				{
				setState(718);
				following_document();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Interface_declarationContext extends ParserRuleContext {
		public Interface_nameContext interface_name() {
			return getRuleContext(Interface_nameContext.class,0);
		}
		public Interface_bodyContext interface_body() {
			return getRuleContext(Interface_bodyContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Type_inheritance_clauseContext type_inheritance_clause() {
			return getRuleContext(Type_inheritance_clauseContext.class,0);
		}
		public Interface_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interface_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInterface_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInterface_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInterface_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Interface_declarationContext interface_declaration() throws RecognitionException {
		Interface_declarationContext _localctx = new Interface_declarationContext(_ctx, getState());
		enterRule(_localctx, 128, RULE_interface_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(722);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Line_document) {
				{
				setState(721);
				document();
				}
			}

			setState(725);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(724);
				attributes();
				}
			}

			setState(727);
			match(T__19);
			setState(728);
			interface_name();
			setState(730);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(729);
				type_inheritance_clause();
				}
			}

			setState(732);
			interface_body();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Interface_nameContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Interface_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interface_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInterface_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInterface_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInterface_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Interface_nameContext interface_name() throws RecognitionException {
		Interface_nameContext _localctx = new Interface_nameContext(_ctx, getState());
		enterRule(_localctx, 130, RULE_interface_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(734);
			type_name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Interface_bodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<Interface_memberContext> interface_member() {
			return getRuleContexts(Interface_memberContext.class);
		}
		public Interface_memberContext interface_member(int i) {
			return getRuleContext(Interface_memberContext.class,i);
		}
		public Interface_bodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interface_body; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInterface_body(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInterface_body(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInterface_body(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Interface_bodyContext interface_body() throws RecognitionException {
		Interface_bodyContext _localctx = new Interface_bodyContext(_ctx, getState());
		enterRule(_localctx, 132, RULE_interface_body);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(736);
			match(LCURLY);
			setState(740);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LT) | (1L << GT) | (1L << BANG) | (1L << QUESTION) | (1L << AT) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || _la==Line_document) {
				{
				{
				setState(737);
				interface_member();
				}
				}
				setState(742);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(743);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Interface_memberContext extends ParserRuleContext {
		public Interface_method_declarationContext interface_method_declaration() {
			return getRuleContext(Interface_method_declarationContext.class,0);
		}
		public Typealias_declarationContext typealias_declaration() {
			return getRuleContext(Typealias_declarationContext.class,0);
		}
		public Interface_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interface_member; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInterface_member(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInterface_member(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInterface_member(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Interface_memberContext interface_member() throws RecognitionException {
		Interface_memberContext _localctx = new Interface_memberContext(_ctx, getState());
		enterRule(_localctx, 134, RULE_interface_member);
		try {
			setState(747);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,74,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(745);
				interface_method_declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(746);
				typealias_declaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Interface_method_declarationContext extends ParserRuleContext {
		public Function_nameContext function_name() {
			return getRuleContext(Function_nameContext.class,0);
		}
		public Function_signatureContext function_signature() {
			return getRuleContext(Function_signatureContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Generic_parameter_clauseContext generic_parameter_clause() {
			return getRuleContext(Generic_parameter_clauseContext.class,0);
		}
		public Interface_method_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interface_method_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInterface_method_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInterface_method_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInterface_method_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Interface_method_declarationContext interface_method_declaration() throws RecognitionException {
		Interface_method_declarationContext _localctx = new Interface_method_declarationContext(_ctx, getState());
		enterRule(_localctx, 136, RULE_interface_method_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(750);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Line_document) {
				{
				setState(749);
				document();
				}
			}

			setState(753);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(752);
				attributes();
				}
			}

			setState(755);
			function_name();
			setState(757);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(756);
				generic_parameter_clause();
				}
			}

			setState(759);
			function_signature();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PatternContext extends ParserRuleContext {
		public Wildcard_patternContext wildcard_pattern() {
			return getRuleContext(Wildcard_patternContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Identifier_patternContext identifier_pattern() {
			return getRuleContext(Identifier_patternContext.class,0);
		}
		public Tuple_patternContext tuple_pattern() {
			return getRuleContext(Tuple_patternContext.class,0);
		}
		public Optional_patternContext optional_pattern() {
			return getRuleContext(Optional_patternContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Expression_patternContext expression_pattern() {
			return getRuleContext(Expression_patternContext.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PatternContext pattern() throws RecognitionException {
		return pattern(0);
	}

	private PatternContext pattern(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PatternContext _localctx = new PatternContext(_ctx, _parentState);
		PatternContext _prevctx = _localctx;
		int _startState = 138;
		enterRecursionRule(_localctx, 138, RULE_pattern, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(778);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,81,_ctx) ) {
			case 1:
				{
				setState(762);
				wildcard_pattern();
				setState(764);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,78,_ctx) ) {
				case 1:
					{
					setState(763);
					type_annotation();
					}
					break;
				}
				}
				break;
			case 2:
				{
				setState(766);
				identifier_pattern();
				setState(768);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,79,_ctx) ) {
				case 1:
					{
					setState(767);
					type_annotation();
					}
					break;
				}
				}
				break;
			case 3:
				{
				setState(770);
				tuple_pattern();
				setState(772);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,80,_ctx) ) {
				case 1:
					{
					setState(771);
					type_annotation();
					}
					break;
				}
				}
				break;
			case 4:
				{
				setState(774);
				optional_pattern();
				}
				break;
			case 5:
				{
				setState(775);
				match(T__20);
				setState(776);
				type_(0);
				}
				break;
			case 6:
				{
				setState(777);
				expression_pattern();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(785);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,82,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PatternContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pattern);
					setState(780);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(781);
					match(T__15);
					setState(782);
					type_(0);
					}
					} 
				}
				setState(787);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,82,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Wildcard_patternContext extends ParserRuleContext {
		public TerminalNode UNDERSCORE() { return getToken(MojoParser.UNDERSCORE, 0); }
		public Wildcard_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_wildcard_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterWildcard_pattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitWildcard_pattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitWildcard_pattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Wildcard_patternContext wildcard_pattern() throws RecognitionException {
		Wildcard_patternContext _localctx = new Wildcard_patternContext(_ctx, getState());
		enterRule(_localctx, 140, RULE_wildcard_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(788);
			match(UNDERSCORE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Identifier_patternContext extends ParserRuleContext {
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public Identifier_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterIdentifier_pattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitIdentifier_pattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitIdentifier_pattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Identifier_patternContext identifier_pattern() throws RecognitionException {
		Identifier_patternContext _localctx = new Identifier_patternContext(_ctx, getState());
		enterRule(_localctx, 142, RULE_identifier_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(790);
			declaration_identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_patternContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Tuple_pattern_element_listContext tuple_pattern_element_list() {
			return getRuleContext(Tuple_pattern_element_listContext.class,0);
		}
		public Tuple_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_pattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_pattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_pattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_patternContext tuple_pattern() throws RecognitionException {
		Tuple_patternContext _localctx = new Tuple_patternContext(_ctx, getState());
		enterRule(_localctx, 144, RULE_tuple_pattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(792);
			match(LPAREN);
			setState(794);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)))) != 0)) {
				{
				setState(793);
				tuple_pattern_element_list();
				}
			}

			setState(796);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_pattern_element_listContext extends ParserRuleContext {
		public List<Tuple_pattern_elementContext> tuple_pattern_element() {
			return getRuleContexts(Tuple_pattern_elementContext.class);
		}
		public Tuple_pattern_elementContext tuple_pattern_element(int i) {
			return getRuleContext(Tuple_pattern_elementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Tuple_pattern_element_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_pattern_element_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_pattern_element_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_pattern_element_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_pattern_element_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_pattern_element_listContext tuple_pattern_element_list() throws RecognitionException {
		Tuple_pattern_element_listContext _localctx = new Tuple_pattern_element_listContext(_ctx, getState());
		enterRule(_localctx, 146, RULE_tuple_pattern_element_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(798);
			tuple_pattern_element();
			setState(803);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(799);
				match(COMMA);
				setState(800);
				tuple_pattern_element();
				}
				}
				setState(805);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_pattern_elementContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public Tuple_pattern_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_pattern_element; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_pattern_element(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_pattern_element(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_pattern_element(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_pattern_elementContext tuple_pattern_element() throws RecognitionException {
		Tuple_pattern_elementContext _localctx = new Tuple_pattern_elementContext(_ctx, getState());
		enterRule(_localctx, 148, RULE_tuple_pattern_element);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(806);
			pattern(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Optional_patternContext extends ParserRuleContext {
		public Identifier_patternContext identifier_pattern() {
			return getRuleContext(Identifier_patternContext.class,0);
		}
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public Optional_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optional_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterOptional_pattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitOptional_pattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitOptional_pattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Optional_patternContext optional_pattern() throws RecognitionException {
		Optional_patternContext _localctx = new Optional_patternContext(_ctx, getState());
		enterRule(_localctx, 150, RULE_optional_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(808);
			identifier_pattern();
			setState(809);
			match(QUESTION);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expression_patternContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Expression_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExpression_pattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExpression_pattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExpression_pattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Expression_patternContext expression_pattern() throws RecognitionException {
		Expression_patternContext _localctx = new Expression_patternContext(_ctx, getState());
		enterRule(_localctx, 152, RULE_expression_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(811);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AttributeContext extends ParserRuleContext {
		public TerminalNode AT() { return getToken(MojoParser.AT, 0); }
		public Attribute_nameContext attribute_name() {
			return getRuleContext(Attribute_nameContext.class,0);
		}
		public Attribute_argument_clauseContext attribute_argument_clause() {
			return getRuleContext(Attribute_argument_clauseContext.class,0);
		}
		public AttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterAttribute(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitAttribute(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitAttribute(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeContext attribute() throws RecognitionException {
		AttributeContext _localctx = new AttributeContext(_ctx, getState());
		enterRule(_localctx, 154, RULE_attribute);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(813);
			match(AT);
			setState(814);
			attribute_name();
			setState(816);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,85,_ctx) ) {
			case 1:
				{
				setState(815);
				attribute_argument_clause();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Attribute_nameContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MojoParser.Identifier, 0); }
		public TerminalNode Decimal_literal() { return getToken(MojoParser.Decimal_literal, 0); }
		public Attribute_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterAttribute_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitAttribute_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitAttribute_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Attribute_nameContext attribute_name() throws RecognitionException {
		Attribute_nameContext _localctx = new Attribute_nameContext(_ctx, getState());
		enterRule(_localctx, 156, RULE_attribute_name);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(818);
			_la = _input.LA(1);
			if ( !(_la==Identifier || _la==Decimal_literal) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Attribute_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public Expression_listContext expression_list() {
			return getRuleContext(Expression_listContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Attribute_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute_argument_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterAttribute_argument_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitAttribute_argument_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitAttribute_argument_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Attribute_argument_clauseContext attribute_argument_clause() throws RecognitionException {
		Attribute_argument_clauseContext _localctx = new Attribute_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 158, RULE_attribute_argument_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(820);
			match(LPAREN);
			setState(821);
			expression_list();
			setState(822);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AttributesContext extends ParserRuleContext {
		public List<AttributeContext> attribute() {
			return getRuleContexts(AttributeContext.class);
		}
		public AttributeContext attribute(int i) {
			return getRuleContext(AttributeContext.class,i);
		}
		public AttributesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributes; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterAttributes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitAttributes(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitAttributes(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributesContext attributes() throws RecognitionException {
		AttributesContext _localctx = new AttributesContext(_ctx, getState());
		enterRule(_localctx, 160, RULE_attributes);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(825); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(824);
					attribute();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(827); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,86,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public Prefix_expressionContext prefix_expression() {
			return getRuleContext(Prefix_expressionContext.class,0);
		}
		public Binary_expressionsContext binary_expressions() {
			return getRuleContext(Binary_expressionsContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExpression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 162, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(829);
			prefix_expression();
			setState(831);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,87,_ctx) ) {
			case 1:
				{
				setState(830);
				binary_expressions();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expression_listContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Expression_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExpression_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExpression_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExpression_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Expression_listContext expression_list() throws RecognitionException {
		Expression_listContext _localctx = new Expression_listContext(_ctx, getState());
		enterRule(_localctx, 164, RULE_expression_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(833);
			expression();
			setState(838);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(834);
				match(COMMA);
				setState(835);
				expression();
				}
				}
				setState(840);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Prefix_expressionContext extends ParserRuleContext {
		public Prefix_operatorContext prefix_operator() {
			return getRuleContext(Prefix_operatorContext.class,0);
		}
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public Prefix_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPrefix_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPrefix_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPrefix_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Prefix_expressionContext prefix_expression() throws RecognitionException {
		Prefix_expressionContext _localctx = new Prefix_expressionContext(_ctx, getState());
		enterRule(_localctx, 166, RULE_prefix_expression);
		try {
			setState(845);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,89,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(841);
				prefix_operator();
				setState(842);
				postfix_expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(844);
				postfix_expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Binary_expressionContext extends ParserRuleContext {
		public Binary_operatorContext binary_operator() {
			return getRuleContext(Binary_operatorContext.class,0);
		}
		public Prefix_expressionContext prefix_expression() {
			return getRuleContext(Prefix_expressionContext.class,0);
		}
		public Assignment_operatorContext assignment_operator() {
			return getRuleContext(Assignment_operatorContext.class,0);
		}
		public Conditional_operatorContext conditional_operator() {
			return getRuleContext(Conditional_operatorContext.class,0);
		}
		public Type_casting_operatorContext type_casting_operator() {
			return getRuleContext(Type_casting_operatorContext.class,0);
		}
		public Binary_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterBinary_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitBinary_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitBinary_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Binary_expressionContext binary_expression() throws RecognitionException {
		Binary_expressionContext _localctx = new Binary_expressionContext(_ctx, getState());
		enterRule(_localctx, 168, RULE_binary_expression);
		try {
			setState(857);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,90,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(847);
				binary_operator();
				setState(848);
				prefix_expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(850);
				assignment_operator();
				setState(851);
				prefix_expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(853);
				conditional_operator();
				setState(854);
				prefix_expression();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(856);
				type_casting_operator();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Binary_expressionsContext extends ParserRuleContext {
		public List<Binary_expressionContext> binary_expression() {
			return getRuleContexts(Binary_expressionContext.class);
		}
		public Binary_expressionContext binary_expression(int i) {
			return getRuleContext(Binary_expressionContext.class,i);
		}
		public Binary_expressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_expressions; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterBinary_expressions(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitBinary_expressions(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitBinary_expressions(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Binary_expressionsContext binary_expressions() throws RecognitionException {
		Binary_expressionsContext _localctx = new Binary_expressionsContext(_ctx, getState());
		enterRule(_localctx, 170, RULE_binary_expressions);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(860); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(859);
					binary_expression();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(862); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,91,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Conditional_operatorContext extends ParserRuleContext {
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Conditional_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterConditional_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitConditional_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitConditional_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Conditional_operatorContext conditional_operator() throws RecognitionException {
		Conditional_operatorContext _localctx = new Conditional_operatorContext(_ctx, getState());
		enterRule(_localctx, 172, RULE_conditional_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(864);
			match(QUESTION);
			setState(865);
			expression();
			setState(866);
			match(COLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_casting_operatorContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public Type_casting_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_casting_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_casting_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_casting_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_casting_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_casting_operatorContext type_casting_operator() throws RecognitionException {
		Type_casting_operatorContext _localctx = new Type_casting_operatorContext(_ctx, getState());
		enterRule(_localctx, 174, RULE_type_casting_operator);
		try {
			setState(872);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__20:
				enterOuterAlt(_localctx, 1);
				{
				setState(868);
				match(T__20);
				setState(869);
				type_(0);
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 2);
				{
				setState(870);
				match(T__15);
				setState(871);
				type_(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Primary_expressionContext extends ParserRuleContext {
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public Generic_argument_clauseContext generic_argument_clause() {
			return getRuleContext(Generic_argument_clauseContext.class,0);
		}
		public Literal_expressionContext literal_expression() {
			return getRuleContext(Literal_expressionContext.class,0);
		}
		public Parenthesized_expressionContext parenthesized_expression() {
			return getRuleContext(Parenthesized_expressionContext.class,0);
		}
		public Tuple_expressionContext tuple_expression() {
			return getRuleContext(Tuple_expressionContext.class,0);
		}
		public Implicit_member_expressionContext implicit_member_expression() {
			return getRuleContext(Implicit_member_expressionContext.class,0);
		}
		public Wildcard_expressionContext wildcard_expression() {
			return getRuleContext(Wildcard_expressionContext.class,0);
		}
		public Primary_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPrimary_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPrimary_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPrimary_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Primary_expressionContext primary_expression() throws RecognitionException {
		Primary_expressionContext _localctx = new Primary_expressionContext(_ctx, getState());
		enterRule(_localctx, 176, RULE_primary_expression);
		try {
			setState(883);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,94,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(874);
				declaration_identifier();
				setState(876);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,93,_ctx) ) {
				case 1:
					{
					setState(875);
					generic_argument_clause();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(878);
				literal_expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(879);
				parenthesized_expression();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(880);
				tuple_expression();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(881);
				implicit_member_expression();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(882);
				wildcard_expression();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Literal_expressionContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public Array_literalContext array_literal() {
			return getRuleContext(Array_literalContext.class,0);
		}
		public Object_literalContext object_literal() {
			return getRuleContext(Object_literalContext.class,0);
		}
		public Literal_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterLiteral_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitLiteral_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitLiteral_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Literal_expressionContext literal_expression() throws RecognitionException {
		Literal_expressionContext _localctx = new Literal_expressionContext(_ctx, getState());
		enterRule(_localctx, 178, RULE_literal_expression);
		try {
			setState(888);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__24:
			case T__26:
			case T__29:
			case SUB:
			case Binary_literal:
			case Octal_literal:
			case Decimal_literal:
			case Pure_decimal_digits:
			case Hexadecimal_literal:
			case Floating_point_literal:
			case Static_string_literal:
			case Interpolated_string_literal:
				enterOuterAlt(_localctx, 1);
				{
				setState(885);
				literal();
				}
				break;
			case LBRACK:
				enterOuterAlt(_localctx, 2);
				{
				setState(886);
				array_literal();
				}
				break;
			case LCURLY:
				enterOuterAlt(_localctx, 3);
				{
				setState(887);
				object_literal();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Array_literalContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public Array_literal_itemsContext array_literal_items() {
			return getRuleContext(Array_literal_itemsContext.class,0);
		}
		public Array_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArray_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArray_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArray_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Array_literalContext array_literal() throws RecognitionException {
		Array_literalContext _localctx = new Array_literalContext(_ctx, getState());
		enterRule(_localctx, 180, RULE_array_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(890);
			match(LBRACK);
			setState(892);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier) | (1L << DOT) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << LT) | (1L << GT) | (1L << UNDERSCORE) | (1L << BANG) | (1L << QUESTION) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE) | (1L << Operator_head_other))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)) | (1L << (Floating_point_literal - 64)) | (1L << (Static_string_literal - 64)) | (1L << (Interpolated_string_literal - 64)))) != 0)) {
				{
				setState(891);
				array_literal_items();
				}
			}

			setState(894);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Array_literal_itemsContext extends ParserRuleContext {
		public Array_literal_itemContext array_literal_item() {
			return getRuleContext(Array_literal_itemContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(MojoParser.COMMA, 0); }
		public Array_literal_itemsContext array_literal_items() {
			return getRuleContext(Array_literal_itemsContext.class,0);
		}
		public Array_literal_itemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_literal_items; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArray_literal_items(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArray_literal_items(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArray_literal_items(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Array_literal_itemsContext array_literal_items() throws RecognitionException {
		Array_literal_itemsContext _localctx = new Array_literal_itemsContext(_ctx, getState());
		enterRule(_localctx, 182, RULE_array_literal_items);
		int _la;
		try {
			setState(904);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,98,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(896);
				array_literal_item();
				setState(898);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(897);
					match(COMMA);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(900);
				array_literal_item();
				setState(901);
				match(COMMA);
				setState(902);
				array_literal_items();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Array_literal_itemContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Array_literal_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_literal_item; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArray_literal_item(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArray_literal_item(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArray_literal_item(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Array_literal_itemContext array_literal_item() throws RecognitionException {
		Array_literal_itemContext _localctx = new Array_literal_itemContext(_ctx, getState());
		enterRule(_localctx, 184, RULE_array_literal_item);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(906);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Object_literalContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public Object_literal_itemsContext object_literal_items() {
			return getRuleContext(Object_literal_itemsContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public Object_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterObject_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitObject_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitObject_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Object_literalContext object_literal() throws RecognitionException {
		Object_literalContext _localctx = new Object_literalContext(_ctx, getState());
		enterRule(_localctx, 186, RULE_object_literal);
		try {
			setState(914);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,99,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(908);
				match(LCURLY);
				setState(909);
				object_literal_items();
				setState(910);
				match(RCURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(912);
				match(LCURLY);
				setState(913);
				match(RCURLY);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Object_literal_itemsContext extends ParserRuleContext {
		public Object_literal_itemContext object_literal_item() {
			return getRuleContext(Object_literal_itemContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(MojoParser.COMMA, 0); }
		public Object_literal_itemsContext object_literal_items() {
			return getRuleContext(Object_literal_itemsContext.class,0);
		}
		public Object_literal_itemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object_literal_items; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterObject_literal_items(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitObject_literal_items(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitObject_literal_items(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Object_literal_itemsContext object_literal_items() throws RecognitionException {
		Object_literal_itemsContext _localctx = new Object_literal_itemsContext(_ctx, getState());
		enterRule(_localctx, 188, RULE_object_literal_items);
		int _la;
		try {
			setState(924);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,101,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(916);
				object_literal_item();
				setState(918);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(917);
					match(COMMA);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(920);
				object_literal_item();
				setState(921);
				match(COMMA);
				setState(922);
				object_literal_items();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Object_literal_itemContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Object_literal_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object_literal_item; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterObject_literal_item(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitObject_literal_item(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitObject_literal_item(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Object_literal_itemContext object_literal_item() throws RecognitionException {
		Object_literal_itemContext _localctx = new Object_literal_itemContext(_ctx, getState());
		enterRule(_localctx, 190, RULE_object_literal_item);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(926);
			expression();
			setState(927);
			match(COLON);
			setState(928);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Implicit_member_expressionContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Implicit_member_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implicit_member_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterImplicit_member_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitImplicit_member_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitImplicit_member_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Implicit_member_expressionContext implicit_member_expression() throws RecognitionException {
		Implicit_member_expressionContext _localctx = new Implicit_member_expressionContext(_ctx, getState());
		enterRule(_localctx, 192, RULE_implicit_member_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(930);
			match(DOT);
			setState(931);
			label_identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Parenthesized_expressionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Parenthesized_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parenthesized_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterParenthesized_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitParenthesized_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitParenthesized_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Parenthesized_expressionContext parenthesized_expression() throws RecognitionException {
		Parenthesized_expressionContext _localctx = new Parenthesized_expressionContext(_ctx, getState());
		enterRule(_localctx, 194, RULE_parenthesized_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(933);
			match(LPAREN);
			setState(934);
			expression();
			setState(935);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_expressionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public List<Tuple_elementContext> tuple_element() {
			return getRuleContexts(Tuple_elementContext.class);
		}
		public Tuple_elementContext tuple_element(int i) {
			return getRuleContext(Tuple_elementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Tuple_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_expressionContext tuple_expression() throws RecognitionException {
		Tuple_expressionContext _localctx = new Tuple_expressionContext(_ctx, getState());
		enterRule(_localctx, 196, RULE_tuple_expression);
		int _la;
		try {
			setState(949);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,103,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(937);
				match(LPAREN);
				setState(938);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(939);
				match(LPAREN);
				setState(940);
				tuple_element();
				setState(943); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(941);
					match(COMMA);
					setState(942);
					tuple_element();
					}
					}
					setState(945); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				setState(947);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_elementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Tuple_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_element; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_element(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_element(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_element(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_elementContext tuple_element() throws RecognitionException {
		Tuple_elementContext _localctx = new Tuple_elementContext(_ctx, getState());
		enterRule(_localctx, 198, RULE_tuple_element);
		try {
			setState(956);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,104,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(951);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(952);
				label_identifier();
				setState(953);
				match(COLON);
				setState(954);
				expression();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Wildcard_expressionContext extends ParserRuleContext {
		public TerminalNode UNDERSCORE() { return getToken(MojoParser.UNDERSCORE, 0); }
		public Wildcard_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_wildcard_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterWildcard_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitWildcard_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitWildcard_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Wildcard_expressionContext wildcard_expression() throws RecognitionException {
		Wildcard_expressionContext _localctx = new Wildcard_expressionContext(_ctx, getState());
		enterRule(_localctx, 200, RULE_wildcard_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(958);
			match(UNDERSCORE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Postfix_expressionContext extends ParserRuleContext {
		public Postfix_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfix_expression; }
	 
		public Postfix_expressionContext() { }
		public void copyFrom(Postfix_expressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Function_call_expressionContext extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public Function_call_argument_clauseContext function_call_argument_clause() {
			return getRuleContext(Function_call_argument_clauseContext.class,0);
		}
		public Function_call_expressionContext(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_call_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_call_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_call_expression(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class Subscript_expressionContext extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public Expression_listContext expression_list() {
			return getRuleContext(Expression_listContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public Subscript_expressionContext(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterSubscript_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitSubscript_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitSubscript_expression(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class Explicit_member_expression1Context extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public TerminalNode Pure_decimal_digits() { return getToken(MojoParser.Pure_decimal_digits, 0); }
		public Explicit_member_expression1Context(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExplicit_member_expression1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExplicit_member_expression1(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExplicit_member_expression1(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class Explicit_member_expression2Context extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public Generic_argument_clauseContext generic_argument_clause() {
			return getRuleContext(Generic_argument_clauseContext.class,0);
		}
		public Explicit_member_expression2Context(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExplicit_member_expression2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExplicit_member_expression2(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExplicit_member_expression2(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class Explicit_member_expression3Context extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Declaration_identifierContext declaration_identifier() {
			return getRuleContext(Declaration_identifierContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public Argument_name_listContext argument_name_list() {
			return getRuleContext(Argument_name_listContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Explicit_member_expression3Context(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExplicit_member_expression3(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExplicit_member_expression3(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExplicit_member_expression3(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class Explicit_member_expression4Context extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public Argument_name_listContext argument_name_list() {
			return getRuleContext(Argument_name_listContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Explicit_member_expression4Context(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterExplicit_member_expression4(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitExplicit_member_expression4(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitExplicit_member_expression4(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class Postfix_operationContext extends Postfix_expressionContext {
		public Postfix_expressionContext postfix_expression() {
			return getRuleContext(Postfix_expressionContext.class,0);
		}
		public Postfix_operatorContext postfix_operator() {
			return getRuleContext(Postfix_operatorContext.class,0);
		}
		public Postfix_operationContext(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPostfix_operation(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPostfix_operation(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPostfix_operation(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class PrimaryContext extends Postfix_expressionContext {
		public Primary_expressionContext primary_expression() {
			return getRuleContext(Primary_expressionContext.class,0);
		}
		public PrimaryContext(Postfix_expressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPrimary(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPrimary(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPrimary(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Postfix_expressionContext postfix_expression() throws RecognitionException {
		return postfix_expression(0);
	}

	private Postfix_expressionContext postfix_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Postfix_expressionContext _localctx = new Postfix_expressionContext(_ctx, _parentState);
		Postfix_expressionContext _prevctx = _localctx;
		int _startState = 202;
		enterRecursionRule(_localctx, 202, RULE_postfix_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new PrimaryContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(961);
			primary_expression();
			}
			_ctx.stop = _input.LT(-1);
			setState(995);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,107,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(993);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,106,_ctx) ) {
					case 1:
						{
						_localctx = new Postfix_operationContext(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(963);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(964);
						postfix_operator();
						}
						break;
					case 2:
						{
						_localctx = new Function_call_expressionContext(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(965);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(966);
						function_call_argument_clause();
						}
						break;
					case 3:
						{
						_localctx = new Explicit_member_expression1Context(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(967);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(968);
						match(DOT);
						setState(969);
						match(Pure_decimal_digits);
						}
						break;
					case 4:
						{
						_localctx = new Explicit_member_expression2Context(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(970);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(971);
						match(DOT);
						setState(972);
						declaration_identifier();
						setState(974);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,105,_ctx) ) {
						case 1:
							{
							setState(973);
							generic_argument_clause();
							}
							break;
						}
						}
						break;
					case 5:
						{
						_localctx = new Explicit_member_expression3Context(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(976);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(977);
						match(DOT);
						setState(978);
						declaration_identifier();
						setState(979);
						match(LPAREN);
						setState(980);
						argument_name_list();
						setState(981);
						match(RPAREN);
						}
						break;
					case 6:
						{
						_localctx = new Explicit_member_expression4Context(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(983);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(984);
						match(LPAREN);
						setState(985);
						argument_name_list();
						setState(986);
						match(RPAREN);
						}
						break;
					case 7:
						{
						_localctx = new Subscript_expressionContext(new Postfix_expressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_postfix_expression);
						setState(988);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(989);
						match(LBRACK);
						setState(990);
						expression_list();
						setState(991);
						match(RBRACK);
						}
						break;
					}
					} 
				}
				setState(997);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,107,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Function_call_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Function_call_argument_listContext function_call_argument_list() {
			return getRuleContext(Function_call_argument_listContext.class,0);
		}
		public Function_call_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call_argument_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_call_argument_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_call_argument_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_call_argument_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_call_argument_clauseContext function_call_argument_clause() throws RecognitionException {
		Function_call_argument_clauseContext _localctx = new Function_call_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 204, RULE_function_call_argument_clause);
		try {
			setState(1004);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,108,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(998);
				match(LPAREN);
				setState(999);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1000);
				match(LPAREN);
				setState(1001);
				function_call_argument_list();
				setState(1002);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_call_argument_listContext extends ParserRuleContext {
		public List<Function_call_argumentContext> function_call_argument() {
			return getRuleContexts(Function_call_argumentContext.class);
		}
		public Function_call_argumentContext function_call_argument(int i) {
			return getRuleContext(Function_call_argumentContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Function_call_argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call_argument_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_call_argument_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_call_argument_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_call_argument_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_call_argument_listContext function_call_argument_list() throws RecognitionException {
		Function_call_argument_listContext _localctx = new Function_call_argument_listContext(_ctx, getState());
		enterRule(_localctx, 206, RULE_function_call_argument_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1006);
			function_call_argument();
			setState(1011);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1007);
				match(COMMA);
				setState(1008);
				function_call_argument();
				}
				}
				setState(1013);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_call_argumentContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public Function_call_argumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call_argument; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_call_argument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_call_argument(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_call_argument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_call_argumentContext function_call_argument() throws RecognitionException {
		Function_call_argumentContext _localctx = new Function_call_argumentContext(_ctx, getState());
		enterRule(_localctx, 208, RULE_function_call_argument);
		try {
			setState(1024);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,110,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1014);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1015);
				label_identifier();
				setState(1016);
				match(COLON);
				setState(1017);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1019);
				operator();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1020);
				label_identifier();
				setState(1021);
				match(COLON);
				setState(1022);
				operator();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Argument_name_listContext extends ParserRuleContext {
		public List<Argument_nameContext> argument_name() {
			return getRuleContexts(Argument_nameContext.class);
		}
		public Argument_nameContext argument_name(int i) {
			return getRuleContext(Argument_nameContext.class,i);
		}
		public Argument_name_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_name_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArgument_name_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArgument_name_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArgument_name_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Argument_name_listContext argument_name_list() throws RecognitionException {
		Argument_name_listContext _localctx = new Argument_name_listContext(_ctx, getState());
		enterRule(_localctx, 210, RULE_argument_name_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1026);
			argument_name();
			setState(1030);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Identifier))) != 0)) {
				{
				{
				setState(1027);
				argument_name();
				}
				}
				setState(1032);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Argument_nameContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Argument_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArgument_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArgument_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArgument_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Argument_nameContext argument_name() throws RecognitionException {
		Argument_nameContext _localctx = new Argument_nameContext(_ctx, getState());
		enterRule(_localctx, 212, RULE_argument_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1033);
			label_identifier();
			setState(1034);
			match(COLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_Context extends ParserRuleContext {
		public Array_typeContext array_type() {
			return getRuleContext(Array_typeContext.class,0);
		}
		public Dictionary_typeContext dictionary_type() {
			return getRuleContext(Dictionary_typeContext.class,0);
		}
		public Function_typeContext function_type() {
			return getRuleContext(Function_typeContext.class,0);
		}
		public Type_identifierContext type_identifier() {
			return getRuleContext(Type_identifierContext.class,0);
		}
		public Tuple_typeContext tuple_type() {
			return getRuleContext(Tuple_typeContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public Type_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_Context type_() throws RecognitionException {
		return type_(0);
	}

	private Type_Context type_(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Type_Context _localctx = new Type_Context(_ctx, _parentState);
		Type_Context _prevctx = _localctx;
		int _startState = 214;
		enterRecursionRule(_localctx, 214, RULE_type_, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1042);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,112,_ctx) ) {
			case 1:
				{
				setState(1037);
				array_type();
				}
				break;
			case 2:
				{
				setState(1038);
				dictionary_type();
				}
				break;
			case 3:
				{
				setState(1039);
				function_type();
				}
				break;
			case 4:
				{
				setState(1040);
				type_identifier();
				}
				break;
			case 5:
				{
				setState(1041);
				tuple_type();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1048);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,113,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Type_Context(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_type_);
					setState(1044);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(1045);
					match(QUESTION);
					}
					} 
				}
				setState(1050);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,113,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Type_annotationContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Type_annotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_annotation; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_annotation(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_annotation(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_annotation(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_annotationContext type_annotation() throws RecognitionException {
		Type_annotationContext _localctx = new Type_annotationContext(_ctx, getState());
		enterRule(_localctx, 216, RULE_type_annotation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1051);
			match(COLON);
			setState(1052);
			type_(0);
			setState(1054);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,114,_ctx) ) {
			case 1:
				{
				setState(1053);
				attributes();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_identifierContext extends ParserRuleContext {
		public List<Type_identifier_clauseContext> type_identifier_clause() {
			return getRuleContexts(Type_identifier_clauseContext.class);
		}
		public Type_identifier_clauseContext type_identifier_clause(int i) {
			return getRuleContext(Type_identifier_clauseContext.class,i);
		}
		public Package_identifierContext package_identifier() {
			return getRuleContext(Package_identifierContext.class,0);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public Type_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_identifierContext type_identifier() throws RecognitionException {
		Type_identifierContext _localctx = new Type_identifierContext(_ctx, getState());
		enterRule(_localctx, 218, RULE_type_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1059);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(1056);
				package_identifier();
				setState(1057);
				match(DOT);
				}
			}

			setState(1061);
			type_identifier_clause();
			setState(1064);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,116,_ctx) ) {
			case 1:
				{
				setState(1062);
				match(DOT);
				setState(1063);
				type_identifier_clause();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_identifier_clauseContext extends ParserRuleContext {
		public Type_nameContext type_name() {
			return getRuleContext(Type_nameContext.class,0);
		}
		public Generic_argument_clauseContext generic_argument_clause() {
			return getRuleContext(Generic_argument_clauseContext.class,0);
		}
		public Type_identifier_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_identifier_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_identifier_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_identifier_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_identifier_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_identifier_clauseContext type_identifier_clause() throws RecognitionException {
		Type_identifier_clauseContext _localctx = new Type_identifier_clauseContext(_ctx, getState());
		enterRule(_localctx, 220, RULE_type_identifier_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1066);
			type_name();
			setState(1068);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,117,_ctx) ) {
			case 1:
				{
				setState(1067);
				generic_argument_clause();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_nameContext extends ParserRuleContext {
		public TerminalNode Type_name() { return getToken(MojoParser.Type_name, 0); }
		public Type_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_nameContext type_name() throws RecognitionException {
		Type_nameContext _localctx = new Type_nameContext(_ctx, getState());
		enterRule(_localctx, 222, RULE_type_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1070);
			match(Type_name);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_typeContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Tuple_type_element_listContext tuple_type_element_list() {
			return getRuleContext(Tuple_type_element_listContext.class,0);
		}
		public Tuple_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_typeContext tuple_type() throws RecognitionException {
		Tuple_typeContext _localctx = new Tuple_typeContext(_ctx, getState());
		enterRule(_localctx, 224, RULE_tuple_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1072);
			match(LPAREN);
			setState(1074);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Type_name) | (1L << Identifier) | (1L << LCURLY) | (1L << LPAREN) | (1L << LBRACK) | (1L << AT))) != 0)) {
				{
				setState(1073);
				tuple_type_element_list();
				}
			}

			setState(1076);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_type_element_listContext extends ParserRuleContext {
		public List<Tuple_type_elementContext> tuple_type_element() {
			return getRuleContexts(Tuple_type_elementContext.class);
		}
		public Tuple_type_elementContext tuple_type_element(int i) {
			return getRuleContext(Tuple_type_elementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Tuple_type_element_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type_element_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_type_element_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_type_element_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_type_element_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_type_element_listContext tuple_type_element_list() throws RecognitionException {
		Tuple_type_element_listContext _localctx = new Tuple_type_element_listContext(_ctx, getState());
		enterRule(_localctx, 226, RULE_tuple_type_element_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1078);
			tuple_type_element();
			setState(1083);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1079);
				match(COMMA);
				setState(1080);
				tuple_type_element();
				}
				}
				setState(1085);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tuple_type_elementContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Tuple_type_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type_element; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterTuple_type_element(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitTuple_type_element(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitTuple_type_element(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Tuple_type_elementContext tuple_type_element() throws RecognitionException {
		Tuple_type_elementContext _localctx = new Tuple_type_elementContext(_ctx, getState());
		enterRule(_localctx, 228, RULE_tuple_type_element);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1086);
			type_(0);
			setState(1088);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1087);
				attributes();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_typeContext extends ParserRuleContext {
		public Function_type_argument_clauseContext function_type_argument_clause() {
			return getRuleContext(Function_type_argument_clauseContext.class,0);
		}
		public Arrow_operatorContext arrow_operator() {
			return getRuleContext(Arrow_operatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Function_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_typeContext function_type() throws RecognitionException {
		Function_typeContext _localctx = new Function_typeContext(_ctx, getState());
		enterRule(_localctx, 230, RULE_function_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1091);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1090);
				attributes();
				}
			}

			setState(1093);
			function_type_argument_clause();
			setState(1094);
			arrow_operator();
			setState(1095);
			type_(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_type_argument_clauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public Function_type_argument_listContext function_type_argument_list() {
			return getRuleContext(Function_type_argument_listContext.class,0);
		}
		public Range_operatorContext range_operator() {
			return getRuleContext(Range_operatorContext.class,0);
		}
		public Function_type_argument_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type_argument_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_type_argument_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_type_argument_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_type_argument_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_type_argument_clauseContext function_type_argument_clause() throws RecognitionException {
		Function_type_argument_clauseContext _localctx = new Function_type_argument_clauseContext(_ctx, getState());
		enterRule(_localctx, 232, RULE_function_type_argument_clause);
		int _la;
		try {
			setState(1106);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,123,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1097);
				match(LPAREN);
				setState(1098);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1099);
				match(LPAREN);
				setState(1100);
				function_type_argument_list();
				setState(1102);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOT) {
					{
					setState(1101);
					range_operator();
					}
				}

				setState(1104);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_type_argument_listContext extends ParserRuleContext {
		public List<Function_type_argumentContext> function_type_argument() {
			return getRuleContexts(Function_type_argumentContext.class);
		}
		public Function_type_argumentContext function_type_argument(int i) {
			return getRuleContext(Function_type_argumentContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Function_type_argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type_argument_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_type_argument_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_type_argument_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_type_argument_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_type_argument_listContext function_type_argument_list() throws RecognitionException {
		Function_type_argument_listContext _localctx = new Function_type_argument_listContext(_ctx, getState());
		enterRule(_localctx, 234, RULE_function_type_argument_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1108);
			function_type_argument();
			setState(1113);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1109);
				match(COMMA);
				setState(1110);
				function_type_argument();
				}
				}
				setState(1115);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Function_type_argumentContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Argument_labelContext argument_label() {
			return getRuleContext(Argument_labelContext.class,0);
		}
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public Function_type_argumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type_argument; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFunction_type_argument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFunction_type_argument(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFunction_type_argument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_type_argumentContext function_type_argument() throws RecognitionException {
		Function_type_argumentContext _localctx = new Function_type_argumentContext(_ctx, getState());
		enterRule(_localctx, 236, RULE_function_type_argument);
		int _la;
		try {
			setState(1126);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,127,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1116);
				type_(0);
				setState(1118);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ELLIPSIS) {
					{
					setState(1117);
					match(ELLIPSIS);
					}
				}

				setState(1121);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1120);
					attributes();
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1123);
				argument_label();
				setState(1124);
				type_annotation();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Argument_labelContext extends ParserRuleContext {
		public Label_identifierContext label_identifier() {
			return getRuleContext(Label_identifierContext.class,0);
		}
		public Argument_labelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_label; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArgument_label(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArgument_label(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArgument_label(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Argument_labelContext argument_label() throws RecognitionException {
		Argument_labelContext _localctx = new Argument_labelContext(_ctx, getState());
		enterRule(_localctx, 238, RULE_argument_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1128);
			label_identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Array_typeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public Array_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArray_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArray_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArray_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Array_typeContext array_type() throws RecognitionException {
		Array_typeContext _localctx = new Array_typeContext(_ctx, getState());
		enterRule(_localctx, 240, RULE_array_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1130);
			match(LBRACK);
			setState(1131);
			type_(0);
			setState(1133);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1132);
				attributes();
				}
			}

			setState(1135);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dictionary_typeContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public List<Type_Context> type_() {
			return getRuleContexts(Type_Context.class);
		}
		public Type_Context type_(int i) {
			return getRuleContext(Type_Context.class,i);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public Dictionary_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictionary_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterDictionary_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitDictionary_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitDictionary_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Dictionary_typeContext dictionary_type() throws RecognitionException {
		Dictionary_typeContext _localctx = new Dictionary_typeContext(_ctx, getState());
		enterRule(_localctx, 242, RULE_dictionary_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1137);
			match(LCURLY);
			setState(1138);
			type_(0);
			setState(1139);
			match(COLON);
			setState(1140);
			type_(0);
			setState(1141);
			attributes();
			setState(1142);
			match(RCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_inheritance_clauseContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public Type_inheritance_listContext type_inheritance_list() {
			return getRuleContext(Type_inheritance_listContext.class,0);
		}
		public Type_inheritance_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_inheritance_clause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_inheritance_clause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_inheritance_clause(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_inheritance_clause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_inheritance_clauseContext type_inheritance_clause() throws RecognitionException {
		Type_inheritance_clauseContext _localctx = new Type_inheritance_clauseContext(_ctx, getState());
		enterRule(_localctx, 244, RULE_type_inheritance_clause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1144);
			match(COLON);
			setState(1145);
			type_inheritance_list();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_inheritance_listContext extends ParserRuleContext {
		public List<Type_identifierContext> type_identifier() {
			return getRuleContexts(Type_identifierContext.class);
		}
		public Type_identifierContext type_identifier(int i) {
			return getRuleContext(Type_identifierContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public Type_inheritance_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_inheritance_list; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterType_inheritance_list(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitType_inheritance_list(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitType_inheritance_list(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Type_inheritance_listContext type_inheritance_list() throws RecognitionException {
		Type_inheritance_listContext _localctx = new Type_inheritance_listContext(_ctx, getState());
		enterRule(_localctx, 246, RULE_type_inheritance_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1147);
			type_identifier();
			setState(1152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1148);
				match(COMMA);
				setState(1149);
				type_identifier();
				}
				}
				setState(1154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Declaration_identifierContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MojoParser.Identifier, 0); }
		public Keyword_as_identifier_in_declarationsContext keyword_as_identifier_in_declarations() {
			return getRuleContext(Keyword_as_identifier_in_declarationsContext.class,0);
		}
		public Declaration_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterDeclaration_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitDeclaration_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitDeclaration_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Declaration_identifierContext declaration_identifier() throws RecognitionException {
		Declaration_identifierContext _localctx = new Declaration_identifierContext(_ctx, getState());
		enterRule(_localctx, 248, RULE_declaration_identifier);
		try {
			setState(1157);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1155);
				match(Identifier);
				}
				break;
			case T__0:
			case T__1:
			case T__2:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
				enterOuterAlt(_localctx, 2);
				{
				setState(1156);
				keyword_as_identifier_in_declarations();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Label_identifierContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MojoParser.Identifier, 0); }
		public Keyword_as_identifier_in_labelsContext keyword_as_identifier_in_labels() {
			return getRuleContext(Keyword_as_identifier_in_labelsContext.class,0);
		}
		public Label_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterLabel_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitLabel_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitLabel_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Label_identifierContext label_identifier() throws RecognitionException {
		Label_identifierContext _localctx = new Label_identifierContext(_ctx, getState());
		enterRule(_localctx, 250, RULE_label_identifier);
		try {
			setState(1161);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1159);
				match(Identifier);
				}
				break;
			case T__0:
			case T__1:
			case T__2:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
				enterOuterAlt(_localctx, 2);
				{
				setState(1160);
				keyword_as_identifier_in_labels();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Path_identifierContext extends ParserRuleContext {
		public List<Declaration_identifierContext> declaration_identifier() {
			return getRuleContexts(Declaration_identifierContext.class);
		}
		public Declaration_identifierContext declaration_identifier(int i) {
			return getRuleContext(Declaration_identifierContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public Path_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_path_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPath_identifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPath_identifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPath_identifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Path_identifierContext path_identifier() throws RecognitionException {
		Path_identifierContext _localctx = new Path_identifierContext(_ctx, getState());
		enterRule(_localctx, 252, RULE_path_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1163);
			declaration_identifier();
			setState(1168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(1164);
				match(DOT);
				setState(1165);
				declaration_identifier();
				}
				}
				setState(1170);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MojoParser.Identifier, 0); }
		public TerminalNode Implicit_parameter_name() { return getToken(MojoParser.Implicit_parameter_name, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitIdentifier(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 254, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1171);
			_la = _input.LA(1);
			if ( !(_la==Identifier || _la==Implicit_parameter_name) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Keyword_as_identifier_in_declarationsContext extends ParserRuleContext {
		public Keyword_as_identifier_in_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword_as_identifier_in_declarations; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterKeyword_as_identifier_in_declarations(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitKeyword_as_identifier_in_declarations(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitKeyword_as_identifier_in_declarations(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Keyword_as_identifier_in_declarationsContext keyword_as_identifier_in_declarations() throws RecognitionException {
		Keyword_as_identifier_in_declarationsContext _localctx = new Keyword_as_identifier_in_declarationsContext(_ctx, getState());
		enterRule(_localctx, 256, RULE_keyword_as_identifier_in_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1173);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Keyword_as_identifier_in_labelsContext extends ParserRuleContext {
		public Keyword_as_identifier_in_labelsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword_as_identifier_in_labels; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterKeyword_as_identifier_in_labels(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitKeyword_as_identifier_in_labels(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitKeyword_as_identifier_in_labels(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Keyword_as_identifier_in_labelsContext keyword_as_identifier_in_labels() throws RecognitionException {
		Keyword_as_identifier_in_labelsContext _localctx = new Keyword_as_identifier_in_labelsContext(_ctx, getState());
		enterRule(_localctx, 258, RULE_keyword_as_identifier_in_labels);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1175);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DocumentContext extends ParserRuleContext {
		public List<TerminalNode> Line_document() { return getTokens(MojoParser.Line_document); }
		public TerminalNode Line_document(int i) {
			return getToken(MojoParser.Line_document, i);
		}
		public DocumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_document; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterDocument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitDocument(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitDocument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DocumentContext document() throws RecognitionException {
		DocumentContext _localctx = new DocumentContext(_ctx, getState());
		enterRule(_localctx, 260, RULE_document);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1178); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(1177);
				match(Line_document);
				}
				}
				setState(1180); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Line_document );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Following_documentContext extends ParserRuleContext {
		public TerminalNode Following_line_document() { return getToken(MojoParser.Following_line_document, 0); }
		public Following_documentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_following_document; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterFollowing_document(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitFollowing_document(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitFollowing_document(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Following_documentContext following_document() throws RecognitionException {
		Following_documentContext _localctx = new Following_documentContext(_ctx, getState());
		enterRule(_localctx, 262, RULE_following_document);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1182);
			match(Following_line_document);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Assignment_operatorContext extends ParserRuleContext {
		public TerminalNode EQUAL() { return getToken(MojoParser.EQUAL, 0); }
		public Assignment_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterAssignment_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitAssignment_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitAssignment_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Assignment_operatorContext assignment_operator() throws RecognitionException {
		Assignment_operatorContext _localctx = new Assignment_operatorContext(_ctx, getState());
		enterRule(_localctx, 264, RULE_assignment_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1184);
			match(EQUAL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Negate_prefix_operatorContext extends ParserRuleContext {
		public TerminalNode SUB() { return getToken(MojoParser.SUB, 0); }
		public Negate_prefix_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_negate_prefix_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterNegate_prefix_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitNegate_prefix_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitNegate_prefix_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Negate_prefix_operatorContext negate_prefix_operator() throws RecognitionException {
		Negate_prefix_operatorContext _localctx = new Negate_prefix_operatorContext(_ctx, getState());
		enterRule(_localctx, 266, RULE_negate_prefix_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1186);
			match(SUB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Compilation_condition_ANDContext extends ParserRuleContext {
		public List<TerminalNode> AND() { return getTokens(MojoParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(MojoParser.AND, i);
		}
		public Compilation_condition_ANDContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilation_condition_AND; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterCompilation_condition_AND(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitCompilation_condition_AND(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitCompilation_condition_AND(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Compilation_condition_ANDContext compilation_condition_AND() throws RecognitionException {
		Compilation_condition_ANDContext _localctx = new Compilation_condition_ANDContext(_ctx, getState());
		enterRule(_localctx, 268, RULE_compilation_condition_AND);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1188);
			match(AND);
			setState(1189);
			match(AND);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Compilation_condition_ORContext extends ParserRuleContext {
		public List<TerminalNode> OR() { return getTokens(MojoParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(MojoParser.OR, i);
		}
		public Compilation_condition_ORContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilation_condition_OR; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterCompilation_condition_OR(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitCompilation_condition_OR(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitCompilation_condition_OR(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Compilation_condition_ORContext compilation_condition_OR() throws RecognitionException {
		Compilation_condition_ORContext _localctx = new Compilation_condition_ORContext(_ctx, getState());
		enterRule(_localctx, 270, RULE_compilation_condition_OR);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1191);
			match(OR);
			setState(1192);
			match(OR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Compilation_condition_GEContext extends ParserRuleContext {
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public TerminalNode EQUAL() { return getToken(MojoParser.EQUAL, 0); }
		public Compilation_condition_GEContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilation_condition_GE; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterCompilation_condition_GE(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitCompilation_condition_GE(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitCompilation_condition_GE(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Compilation_condition_GEContext compilation_condition_GE() throws RecognitionException {
		Compilation_condition_GEContext _localctx = new Compilation_condition_GEContext(_ctx, getState());
		enterRule(_localctx, 272, RULE_compilation_condition_GE);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1194);
			match(GT);
			setState(1195);
			match(EQUAL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Arrow_operatorContext extends ParserRuleContext {
		public TerminalNode SUB() { return getToken(MojoParser.SUB, 0); }
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public Arrow_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrow_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterArrow_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitArrow_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitArrow_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Arrow_operatorContext arrow_operator() throws RecognitionException {
		Arrow_operatorContext _localctx = new Arrow_operatorContext(_ctx, getState());
		enterRule(_localctx, 274, RULE_arrow_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1197);
			match(SUB);
			setState(1198);
			match(GT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Range_operatorContext extends ParserRuleContext {
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public Range_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterRange_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitRange_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitRange_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Range_operatorContext range_operator() throws RecognitionException {
		Range_operatorContext _localctx = new Range_operatorContext(_ctx, getState());
		enterRule(_localctx, 276, RULE_range_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1200);
			match(DOT);
			setState(1201);
			match(DOT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Same_type_equalsContext extends ParserRuleContext {
		public List<TerminalNode> EQUAL() { return getTokens(MojoParser.EQUAL); }
		public TerminalNode EQUAL(int i) {
			return getToken(MojoParser.EQUAL, i);
		}
		public Same_type_equalsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_same_type_equals; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterSame_type_equals(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitSame_type_equals(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitSame_type_equals(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Same_type_equalsContext same_type_equals() throws RecognitionException {
		Same_type_equalsContext _localctx = new Same_type_equalsContext(_ctx, getState());
		enterRule(_localctx, 278, RULE_same_type_equals);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1203);
			match(EQUAL);
			setState(1204);
			match(EQUAL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Binary_operatorContext extends ParserRuleContext {
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public Binary_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterBinary_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitBinary_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitBinary_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Binary_operatorContext binary_operator() throws RecognitionException {
		Binary_operatorContext _localctx = new Binary_operatorContext(_ctx, getState());
		enterRule(_localctx, 280, RULE_binary_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1206);
			operator();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Prefix_operatorContext extends ParserRuleContext {
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public Prefix_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPrefix_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPrefix_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPrefix_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Prefix_operatorContext prefix_operator() throws RecognitionException {
		Prefix_operatorContext _localctx = new Prefix_operatorContext(_ctx, getState());
		enterRule(_localctx, 282, RULE_prefix_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1208);
			operator();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Postfix_operatorContext extends ParserRuleContext {
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public Postfix_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfix_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterPostfix_operator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitPostfix_operator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitPostfix_operator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Postfix_operatorContext postfix_operator() throws RecognitionException {
		Postfix_operatorContext _localctx = new Postfix_operatorContext(_ctx, getState());
		enterRule(_localctx, 284, RULE_postfix_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1210);
			operator();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OperatorContext extends ParserRuleContext {
		public Operator_headContext operator_head() {
			return getRuleContext(Operator_headContext.class,0);
		}
		public List<Operator_characterContext> operator_character() {
			return getRuleContexts(Operator_characterContext.class);
		}
		public Operator_characterContext operator_character(int i) {
			return getRuleContext(Operator_characterContext.class,i);
		}
		public Dot_operator_headContext dot_operator_head() {
			return getRuleContext(Dot_operator_headContext.class,0);
		}
		public List<Dot_operator_characterContext> dot_operator_character() {
			return getRuleContexts(Dot_operator_characterContext.class);
		}
		public Dot_operator_characterContext dot_operator_character(int i) {
			return getRuleContext(Dot_operator_characterContext.class,i);
		}
		public OperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterOperator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitOperator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 286, RULE_operator);
		try {
			int _alt;
			setState(1226);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 1);
				{
				setState(1212);
				operator_head();
				setState(1216);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,134,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1213);
						operator_character();
						}
						} 
					}
					setState(1218);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,134,_ctx);
				}
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(1219);
				dot_operator_head();
				setState(1223);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,135,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1220);
						dot_operator_character();
						}
						} 
					}
					setState(1225);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,135,_ctx);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Operator_characterContext extends ParserRuleContext {
		public Operator_headContext operator_head() {
			return getRuleContext(Operator_headContext.class,0);
		}
		public TerminalNode Operator_following_character() { return getToken(MojoParser.Operator_following_character, 0); }
		public Operator_characterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_character; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterOperator_character(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitOperator_character(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitOperator_character(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Operator_characterContext operator_character() throws RecognitionException {
		Operator_characterContext _localctx = new Operator_characterContext(_ctx, getState());
		enterRule(_localctx, 288, RULE_operator_character);
		try {
			setState(1230);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
				enterOuterAlt(_localctx, 1);
				{
				setState(1228);
				operator_head();
				}
				break;
			case Operator_following_character:
				enterOuterAlt(_localctx, 2);
				{
				setState(1229);
				match(Operator_following_character);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Operator_headContext extends ParserRuleContext {
		public TerminalNode DIV() { return getToken(MojoParser.DIV, 0); }
		public TerminalNode EQUAL() { return getToken(MojoParser.EQUAL, 0); }
		public TerminalNode SUB() { return getToken(MojoParser.SUB, 0); }
		public TerminalNode ADD() { return getToken(MojoParser.ADD, 0); }
		public TerminalNode BANG() { return getToken(MojoParser.BANG, 0); }
		public TerminalNode MUL() { return getToken(MojoParser.MUL, 0); }
		public TerminalNode MOD() { return getToken(MojoParser.MOD, 0); }
		public TerminalNode AND() { return getToken(MojoParser.AND, 0); }
		public TerminalNode OR() { return getToken(MojoParser.OR, 0); }
		public TerminalNode LT() { return getToken(MojoParser.LT, 0); }
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public TerminalNode CARET() { return getToken(MojoParser.CARET, 0); }
		public TerminalNode TILDE() { return getToken(MojoParser.TILDE, 0); }
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public TerminalNode Operator_head_other() { return getToken(MojoParser.Operator_head_other, 0); }
		public Operator_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_head; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterOperator_head(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitOperator_head(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitOperator_head(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Operator_headContext operator_head() throws RecognitionException {
		Operator_headContext _localctx = new Operator_headContext(_ctx, getState());
		enterRule(_localctx, 290, RULE_operator_head);
		int _la;
		try {
			setState(1234);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
				enterOuterAlt(_localctx, 1);
				{
				setState(1232);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << BANG) | (1L << QUESTION) | (1L << AND) | (1L << SUB) | (1L << EQUAL) | (1L << OR) | (1L << DIV) | (1L << ADD) | (1L << MUL) | (1L << MOD) | (1L << CARET) | (1L << TILDE))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case Operator_head_other:
				enterOuterAlt(_localctx, 2);
				{
				setState(1233);
				match(Operator_head_other);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dot_operator_headContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Dot_operator_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dot_operator_head; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterDot_operator_head(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitDot_operator_head(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitDot_operator_head(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Dot_operator_headContext dot_operator_head() throws RecognitionException {
		Dot_operator_headContext _localctx = new Dot_operator_headContext(_ctx, getState());
		enterRule(_localctx, 292, RULE_dot_operator_head);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1236);
			match(DOT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dot_operator_characterContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Operator_characterContext operator_character() {
			return getRuleContext(Operator_characterContext.class,0);
		}
		public Dot_operator_characterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dot_operator_character; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterDot_operator_character(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitDot_operator_character(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitDot_operator_character(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Dot_operator_characterContext dot_operator_character() throws RecognitionException {
		Dot_operator_characterContext _localctx = new Dot_operator_characterContext(_ctx, getState());
		enterRule(_localctx, 294, RULE_dot_operator_character);
		try {
			setState(1240);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(1238);
				match(DOT);
				}
				break;
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case SUB:
			case EQUAL:
			case OR:
			case DIV:
			case ADD:
			case MUL:
			case MOD:
			case CARET:
			case TILDE:
			case Operator_head_other:
			case Operator_following_character:
				enterOuterAlt(_localctx, 2);
				{
				setState(1239);
				operator_character();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LiteralContext extends ParserRuleContext {
		public Numeric_literalContext numeric_literal() {
			return getRuleContext(Numeric_literalContext.class,0);
		}
		public String_literalContext string_literal() {
			return getRuleContext(String_literalContext.class,0);
		}
		public Boolean_literalContext boolean_literal() {
			return getRuleContext(Boolean_literalContext.class,0);
		}
		public Null_literalContext null_literal() {
			return getRuleContext(Null_literalContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitLiteral(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 296, RULE_literal);
		try {
			setState(1246);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
			case Binary_literal:
			case Octal_literal:
			case Decimal_literal:
			case Pure_decimal_digits:
			case Hexadecimal_literal:
			case Floating_point_literal:
				enterOuterAlt(_localctx, 1);
				{
				setState(1242);
				numeric_literal();
				}
				break;
			case Static_string_literal:
			case Interpolated_string_literal:
				enterOuterAlt(_localctx, 2);
				{
				setState(1243);
				string_literal();
				}
				break;
			case T__24:
			case T__29:
				enterOuterAlt(_localctx, 3);
				{
				setState(1244);
				boolean_literal();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 4);
				{
				setState(1245);
				null_literal();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Numeric_literalContext extends ParserRuleContext {
		public Integer_literalContext integer_literal() {
			return getRuleContext(Integer_literalContext.class,0);
		}
		public Negate_prefix_operatorContext negate_prefix_operator() {
			return getRuleContext(Negate_prefix_operatorContext.class,0);
		}
		public TerminalNode Floating_point_literal() { return getToken(MojoParser.Floating_point_literal, 0); }
		public Numeric_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numeric_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterNumeric_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitNumeric_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitNumeric_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Numeric_literalContext numeric_literal() throws RecognitionException {
		Numeric_literalContext _localctx = new Numeric_literalContext(_ctx, getState());
		enterRule(_localctx, 298, RULE_numeric_literal);
		int _la;
		try {
			setState(1256);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,143,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1249);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SUB) {
					{
					setState(1248);
					negate_prefix_operator();
					}
				}

				setState(1251);
				integer_literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1253);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SUB) {
					{
					setState(1252);
					negate_prefix_operator();
					}
				}

				setState(1255);
				match(Floating_point_literal);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Boolean_literalContext extends ParserRuleContext {
		public Boolean_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolean_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterBoolean_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitBoolean_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitBoolean_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Boolean_literalContext boolean_literal() throws RecognitionException {
		Boolean_literalContext _localctx = new Boolean_literalContext(_ctx, getState());
		enterRule(_localctx, 300, RULE_boolean_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1258);
			_la = _input.LA(1);
			if ( !(_la==T__24 || _la==T__29) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Null_literalContext extends ParserRuleContext {
		public Null_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_null_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterNull_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitNull_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitNull_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Null_literalContext null_literal() throws RecognitionException {
		Null_literalContext _localctx = new Null_literalContext(_ctx, getState());
		enterRule(_localctx, 302, RULE_null_literal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1260);
			match(T__26);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Integer_literalContext extends ParserRuleContext {
		public TerminalNode Binary_literal() { return getToken(MojoParser.Binary_literal, 0); }
		public TerminalNode Octal_literal() { return getToken(MojoParser.Octal_literal, 0); }
		public TerminalNode Decimal_literal() { return getToken(MojoParser.Decimal_literal, 0); }
		public TerminalNode Pure_decimal_digits() { return getToken(MojoParser.Pure_decimal_digits, 0); }
		public TerminalNode Hexadecimal_literal() { return getToken(MojoParser.Hexadecimal_literal, 0); }
		public Integer_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterInteger_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitInteger_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitInteger_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Integer_literalContext integer_literal() throws RecognitionException {
		Integer_literalContext _localctx = new Integer_literalContext(_ctx, getState());
		enterRule(_localctx, 304, RULE_integer_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1262);
			_la = _input.LA(1);
			if ( !(((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (Binary_literal - 64)) | (1L << (Octal_literal - 64)) | (1L << (Decimal_literal - 64)) | (1L << (Pure_decimal_digits - 64)) | (1L << (Hexadecimal_literal - 64)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class String_literalContext extends ParserRuleContext {
		public TerminalNode Static_string_literal() { return getToken(MojoParser.Static_string_literal, 0); }
		public TerminalNode Interpolated_string_literal() { return getToken(MojoParser.Interpolated_string_literal, 0); }
		public String_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).enterString_literal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MojoListener ) ((MojoListener)listener).exitString_literal(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoVisitor ) return ((MojoVisitor<? extends T>)visitor).visitString_literal(this);
			else return visitor.visitChildren(this);
		}
	}

	public final String_literalContext string_literal() throws RecognitionException {
		String_literalContext _localctx = new String_literalContext(_ctx, getState());
		enterRule(_localctx, 306, RULE_string_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1264);
			_la = _input.LA(1);
			if ( !(_la==Static_string_literal || _la==Interpolated_string_literal) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 69:
			return pattern_sempred((PatternContext)_localctx, predIndex);
		case 101:
			return postfix_expression_sempred((Postfix_expressionContext)_localctx, predIndex);
		case 107:
			return type__sempred((Type_Context)_localctx, predIndex);
		}
		return true;
	}
	private boolean pattern_sempred(PatternContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean postfix_expression_sempred(Postfix_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 7);
		case 2:
			return precpred(_ctx, 6);
		case 3:
			return precpred(_ctx, 5);
		case 4:
			return precpred(_ctx, 4);
		case 5:
			return precpred(_ctx, 3);
		case 6:
			return precpred(_ctx, 2);
		case 7:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean type__sempred(Type_Context _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3N\u04f5\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\4U\tU\4V\tV\4W\tW\4X\tX\4Y\tY\4Z\tZ\4[\t[\4\\\t\\\4]\t]\4^\t^\4_\t_\4"+
		"`\t`\4a\ta\4b\tb\4c\tc\4d\td\4e\te\4f\tf\4g\tg\4h\th\4i\ti\4j\tj\4k\t"+
		"k\4l\tl\4m\tm\4n\tn\4o\to\4p\tp\4q\tq\4r\tr\4s\ts\4t\tt\4u\tu\4v\tv\4"+
		"w\tw\4x\tx\4y\ty\4z\tz\4{\t{\4|\t|\4}\t}\4~\t~\4\177\t\177\4\u0080\t\u0080"+
		"\4\u0081\t\u0081\4\u0082\t\u0082\4\u0083\t\u0083\4\u0084\t\u0084\4\u0085"+
		"\t\u0085\4\u0086\t\u0086\4\u0087\t\u0087\4\u0088\t\u0088\4\u0089\t\u0089"+
		"\4\u008a\t\u008a\4\u008b\t\u008b\4\u008c\t\u008c\4\u008d\t\u008d\4\u008e"+
		"\t\u008e\4\u008f\t\u008f\4\u0090\t\u0090\4\u0091\t\u0091\4\u0092\t\u0092"+
		"\4\u0093\t\u0093\4\u0094\t\u0094\4\u0095\t\u0095\4\u0096\t\u0096\4\u0097"+
		"\t\u0097\4\u0098\t\u0098\4\u0099\t\u0099\4\u009a\t\u009a\4\u009b\t\u009b"+
		"\3\2\6\2\u0138\n\2\r\2\16\2\u0139\3\2\3\2\3\3\3\3\5\3\u0140\n\3\3\3\3"+
		"\3\5\3\u0144\n\3\5\3\u0146\n\3\3\4\3\4\3\4\3\4\5\4\u014c\n\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\7\7\u015b\n\7\f\7\16\7\u015e"+
		"\13\7\3\b\3\b\5\b\u0162\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u016c"+
		"\n\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\5\13\u0175\n\13\3\f\3\f\3\f\3\f\5\f"+
		"\u017b\n\f\3\r\3\r\3\r\3\r\5\r\u0181\n\r\3\16\3\16\3\16\3\16\5\16\u0187"+
		"\n\16\3\16\3\16\3\17\3\17\5\17\u018d\n\17\3\20\3\20\3\20\3\20\5\20\u0193"+
		"\n\20\3\21\3\21\3\21\5\21\u0198\n\21\3\22\3\22\5\22\u019c\n\22\3\23\3"+
		"\23\5\23\u01a0\n\23\3\24\3\24\5\24\u01a4\n\24\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\7\26\u01ad\n\26\f\26\16\26\u01b0\13\26\3\27\3\27\3\27\3\27"+
		"\3\27\5\27\u01b7\n\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\7\31\u01c0\n"+
		"\31\f\31\16\31\u01c3\13\31\3\32\3\32\5\32\u01c7\n\32\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\3\33\5\33\u01d1\n\33\3\34\3\34\6\34\u01d5\n\34\r\34"+
		"\16\34\u01d6\3\34\3\34\3\35\3\35\3\35\3\35\3\36\3\36\3\36\7\36\u01e2\n"+
		"\36\f\36\16\36\u01e5\13\36\3\37\3\37\3\37\3\37\3\37\5\37\u01ec\n\37\3"+
		" \3 \3 \7 \u01f1\n \f \16 \u01f4\13 \3!\3!\5!\u01f8\n!\3\"\3\"\3\"\3#"+
		"\3#\3#\3$\3$\3$\5$\u0203\n$\3%\3%\3%\3%\3%\3&\3&\5&\u020c\n&\3&\3&\3&"+
		"\5&\u0211\n&\7&\u0213\n&\f&\16&\u0216\13&\3\'\3\'\5\'\u021a\n\'\3(\3("+
		"\5(\u021e\n(\3)\5)\u0221\n)\3)\3)\3)\3*\3*\3*\7*\u0229\n*\f*\16*\u022c"+
		"\13*\3+\3+\5+\u0230\n+\3,\3,\3,\3-\5-\u0236\n-\3-\3-\3-\5-\u023b\n-\3"+
		"-\3-\3.\3.\3/\3/\3/\3\60\3\60\3\60\5\60\u0247\n\60\3\60\3\60\5\60\u024b"+
		"\n\60\3\61\5\61\u024e\n\61\3\61\3\61\3\62\3\62\5\62\u0254\n\62\3\63\3"+
		"\63\5\63\u0258\n\63\3\64\3\64\3\64\5\64\u025d\n\64\3\65\3\65\3\66\3\66"+
		"\3\66\3\66\3\66\3\66\5\66\u0267\n\66\3\67\3\67\3\67\7\67\u026c\n\67\f"+
		"\67\16\67\u026f\13\67\38\38\38\58\u0274\n8\38\38\38\38\58\u027a\n8\58"+
		"\u027c\n8\39\59\u027f\n9\39\59\u0282\n9\39\39\39\59\u0287\n9\39\59\u028a"+
		"\n9\39\39\59\u028e\n9\39\39\3:\3:\3;\6;\u0295\n;\r;\16;\u0296\3<\5<\u029a"+
		"\n<\3<\3<\5<\u029e\n<\3<\5<\u02a1\n<\3<\5<\u02a4\n<\3=\5=\u02a7\n=\3="+
		"\5=\u02aa\n=\3=\3=\3=\5=\u02af\n=\3=\5=\u02b2\n=\3=\5=\u02b5\n=\3>\3>"+
		"\3?\3?\7?\u02bb\n?\f?\16?\u02be\13?\3?\3?\3@\3@\3@\3@\3@\5@\u02c7\n@\3"+
		"A\5A\u02ca\nA\3A\3A\3A\5A\u02cf\nA\3A\5A\u02d2\nA\3B\5B\u02d5\nB\3B\5"+
		"B\u02d8\nB\3B\3B\3B\5B\u02dd\nB\3B\3B\3C\3C\3D\3D\7D\u02e5\nD\fD\16D\u02e8"+
		"\13D\3D\3D\3E\3E\5E\u02ee\nE\3F\5F\u02f1\nF\3F\5F\u02f4\nF\3F\3F\5F\u02f8"+
		"\nF\3F\3F\3G\3G\3G\5G\u02ff\nG\3G\3G\5G\u0303\nG\3G\3G\5G\u0307\nG\3G"+
		"\3G\3G\3G\5G\u030d\nG\3G\3G\3G\7G\u0312\nG\fG\16G\u0315\13G\3H\3H\3I\3"+
		"I\3J\3J\5J\u031d\nJ\3J\3J\3K\3K\3K\7K\u0324\nK\fK\16K\u0327\13K\3L\3L"+
		"\3M\3M\3M\3N\3N\3O\3O\3O\5O\u0333\nO\3P\3P\3Q\3Q\3Q\3Q\3R\6R\u033c\nR"+
		"\rR\16R\u033d\3S\3S\5S\u0342\nS\3T\3T\3T\7T\u0347\nT\fT\16T\u034a\13T"+
		"\3U\3U\3U\3U\5U\u0350\nU\3V\3V\3V\3V\3V\3V\3V\3V\3V\3V\5V\u035c\nV\3W"+
		"\6W\u035f\nW\rW\16W\u0360\3X\3X\3X\3X\3Y\3Y\3Y\3Y\5Y\u036b\nY\3Z\3Z\5"+
		"Z\u036f\nZ\3Z\3Z\3Z\3Z\3Z\5Z\u0376\nZ\3[\3[\3[\5[\u037b\n[\3\\\3\\\5\\"+
		"\u037f\n\\\3\\\3\\\3]\3]\5]\u0385\n]\3]\3]\3]\3]\5]\u038b\n]\3^\3^\3_"+
		"\3_\3_\3_\3_\3_\5_\u0395\n_\3`\3`\5`\u0399\n`\3`\3`\3`\3`\5`\u039f\n`"+
		"\3a\3a\3a\3a\3b\3b\3b\3c\3c\3c\3c\3d\3d\3d\3d\3d\3d\6d\u03b2\nd\rd\16"+
		"d\u03b3\3d\3d\5d\u03b8\nd\3e\3e\3e\3e\3e\5e\u03bf\ne\3f\3f\3g\3g\3g\3"+
		"g\3g\3g\3g\3g\3g\3g\3g\3g\3g\3g\5g\u03d1\ng\3g\3g\3g\3g\3g\3g\3g\3g\3"+
		"g\3g\3g\3g\3g\3g\3g\3g\3g\7g\u03e4\ng\fg\16g\u03e7\13g\3h\3h\3h\3h\3h"+
		"\3h\5h\u03ef\nh\3i\3i\3i\7i\u03f4\ni\fi\16i\u03f7\13i\3j\3j\3j\3j\3j\3"+
		"j\3j\3j\3j\3j\5j\u0403\nj\3k\3k\7k\u0407\nk\fk\16k\u040a\13k\3l\3l\3l"+
		"\3m\3m\3m\3m\3m\3m\5m\u0415\nm\3m\3m\7m\u0419\nm\fm\16m\u041c\13m\3n\3"+
		"n\3n\5n\u0421\nn\3o\3o\3o\5o\u0426\no\3o\3o\3o\5o\u042b\no\3p\3p\5p\u042f"+
		"\np\3q\3q\3r\3r\5r\u0435\nr\3r\3r\3s\3s\3s\7s\u043c\ns\fs\16s\u043f\13"+
		"s\3t\3t\5t\u0443\nt\3u\5u\u0446\nu\3u\3u\3u\3u\3v\3v\3v\3v\3v\5v\u0451"+
		"\nv\3v\3v\5v\u0455\nv\3w\3w\3w\7w\u045a\nw\fw\16w\u045d\13w\3x\3x\5x\u0461"+
		"\nx\3x\5x\u0464\nx\3x\3x\3x\5x\u0469\nx\3y\3y\3z\3z\3z\5z\u0470\nz\3z"+
		"\3z\3{\3{\3{\3{\3{\3{\3{\3|\3|\3|\3}\3}\3}\7}\u0481\n}\f}\16}\u0484\13"+
		"}\3~\3~\5~\u0488\n~\3\177\3\177\5\177\u048c\n\177\3\u0080\3\u0080\3\u0080"+
		"\7\u0080\u0491\n\u0080\f\u0080\16\u0080\u0494\13\u0080\3\u0081\3\u0081"+
		"\3\u0082\3\u0082\3\u0083\3\u0083\3\u0084\6\u0084\u049d\n\u0084\r\u0084"+
		"\16\u0084\u049e\3\u0085\3\u0085\3\u0086\3\u0086\3\u0087\3\u0087\3\u0088"+
		"\3\u0088\3\u0088\3\u0089\3\u0089\3\u0089\3\u008a\3\u008a\3\u008a\3\u008b"+
		"\3\u008b\3\u008b\3\u008c\3\u008c\3\u008c\3\u008d\3\u008d\3\u008d\3\u008e"+
		"\3\u008e\3\u008f\3\u008f\3\u0090\3\u0090\3\u0091\3\u0091\7\u0091\u04c1"+
		"\n\u0091\f\u0091\16\u0091\u04c4\13\u0091\3\u0091\3\u0091\7\u0091\u04c8"+
		"\n\u0091\f\u0091\16\u0091\u04cb\13\u0091\5\u0091\u04cd\n\u0091\3\u0092"+
		"\3\u0092\5\u0092\u04d1\n\u0092\3\u0093\3\u0093\5\u0093\u04d5\n\u0093\3"+
		"\u0094\3\u0094\3\u0095\3\u0095\5\u0095\u04db\n\u0095\3\u0096\3\u0096\3"+
		"\u0096\3\u0096\5\u0096\u04e1\n\u0096\3\u0097\5\u0097\u04e4\n\u0097\3\u0097"+
		"\3\u0097\5\u0097\u04e8\n\u0097\3\u0097\5\u0097\u04eb\n\u0097\3\u0098\3"+
		"\u0098\3\u0099\3\u0099\3\u009a\3\u009a\3\u009b\3\u009b\3\u009b\2\5\u008c"+
		"\u00cc\u00d8\u009c\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60"+
		"\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084\u0086"+
		"\u0088\u008a\u008c\u008e\u0090\u0092\u0094\u0096\u0098\u009a\u009c\u009e"+
		"\u00a0\u00a2\u00a4\u00a6\u00a8\u00aa\u00ac\u00ae\u00b0\u00b2\u00b4\u00b6"+
		"\u00b8\u00ba\u00bc\u00be\u00c0\u00c2\u00c4\u00c6\u00c8\u00ca\u00cc\u00ce"+
		"\u00d0\u00d2\u00d4\u00d6\u00d8\u00da\u00dc\u00de\u00e0\u00e2\u00e4\u00e6"+
		"\u00e8\u00ea\u00ec\u00ee\u00f0\u00f2\u00f4\u00f6\u00f8\u00fa\u00fc\u00fe"+
		"\u0100\u0102\u0104\u0106\u0108\u010a\u010c\u010e\u0110\u0112\u0114\u0116"+
		"\u0118\u011a\u011c\u011e\u0120\u0122\u0124\u0126\u0128\u012a\u012c\u012e"+
		"\u0130\u0132\u0134\2\t\4\2##DD\4\2##AA\5\2\3\5\7\13\r!\5\2./\61\62\64"+
		"=\4\2\33\33  \3\2BF\3\2HI\2\u050f\2\u0137\3\2\2\2\4\u0145\3\2\2\2\6\u014b"+
		"\3\2\2\2\b\u014d\3\2\2\2\n\u0153\3\2\2\2\f\u0157\3\2\2\2\16\u0161\3\2"+
		"\2\2\20\u016b\3\2\2\2\22\u016d\3\2\2\2\24\u0174\3\2\2\2\26\u0176\3\2\2"+
		"\2\30\u0180\3\2\2\2\32\u0182\3\2\2\2\34\u018a\3\2\2\2\36\u018e\3\2\2\2"+
		" \u0197\3\2\2\2\"\u0199\3\2\2\2$\u019d\3\2\2\2&\u01a1\3\2\2\2(\u01a5\3"+
		"\2\2\2*\u01a9\3\2\2\2,\u01b6\3\2\2\2.\u01b8\3\2\2\2\60\u01bc\3\2\2\2\62"+
		"\u01c4\3\2\2\2\64\u01d0\3\2\2\2\66\u01d2\3\2\2\28\u01da\3\2\2\2:\u01de"+
		"\3\2\2\2<\u01e6\3\2\2\2>\u01ed\3\2\2\2@\u01f7\3\2\2\2B\u01f9\3\2\2\2D"+
		"\u01fc\3\2\2\2F\u01ff\3\2\2\2H\u0204\3\2\2\2J\u020b\3\2\2\2L\u0217\3\2"+
		"\2\2N\u021b\3\2\2\2P\u0220\3\2\2\2R\u0225\3\2\2\2T\u022d\3\2\2\2V\u0231"+
		"\3\2\2\2X\u0235\3\2\2\2Z\u023e\3\2\2\2\\\u0240\3\2\2\2^\u0243\3\2\2\2"+
		"`\u024d\3\2\2\2b\u0253\3\2\2\2d\u0255\3\2\2\2f\u0259\3\2\2\2h\u025e\3"+
		"\2\2\2j\u0266\3\2\2\2l\u0268\3\2\2\2n\u027b\3\2\2\2p\u027e\3\2\2\2r\u0291"+
		"\3\2\2\2t\u0294\3\2\2\2v\u0299\3\2\2\2x\u02a6\3\2\2\2z\u02b6\3\2\2\2|"+
		"\u02b8\3\2\2\2~\u02c6\3\2\2\2\u0080\u02c9\3\2\2\2\u0082\u02d4\3\2\2\2"+
		"\u0084\u02e0\3\2\2\2\u0086\u02e2\3\2\2\2\u0088\u02ed\3\2\2\2\u008a\u02f0"+
		"\3\2\2\2\u008c\u030c\3\2\2\2\u008e\u0316\3\2\2\2\u0090\u0318\3\2\2\2\u0092"+
		"\u031a\3\2\2\2\u0094\u0320\3\2\2\2\u0096\u0328\3\2\2\2\u0098\u032a\3\2"+
		"\2\2\u009a\u032d\3\2\2\2\u009c\u032f\3\2\2\2\u009e\u0334\3\2\2\2\u00a0"+
		"\u0336\3\2\2\2\u00a2\u033b\3\2\2\2\u00a4\u033f\3\2\2\2\u00a6\u0343\3\2"+
		"\2\2\u00a8\u034f\3\2\2\2\u00aa\u035b\3\2\2\2\u00ac\u035e\3\2\2\2\u00ae"+
		"\u0362\3\2\2\2\u00b0\u036a\3\2\2\2\u00b2\u0375\3\2\2\2\u00b4\u037a\3\2"+
		"\2\2\u00b6\u037c\3\2\2\2\u00b8\u038a\3\2\2\2\u00ba\u038c\3\2\2\2\u00bc"+
		"\u0394\3\2\2\2\u00be\u039e\3\2\2\2\u00c0\u03a0\3\2\2\2\u00c2\u03a4\3\2"+
		"\2\2\u00c4\u03a7\3\2\2\2\u00c6\u03b7\3\2\2\2\u00c8\u03be\3\2\2\2\u00ca"+
		"\u03c0\3\2\2\2\u00cc\u03c2\3\2\2\2\u00ce\u03ee\3\2\2\2\u00d0\u03f0\3\2"+
		"\2\2\u00d2\u0402\3\2\2\2\u00d4\u0404\3\2\2\2\u00d6\u040b\3\2\2\2\u00d8"+
		"\u0414\3\2\2\2\u00da\u041d\3\2\2\2\u00dc\u0425\3\2\2\2\u00de\u042c\3\2"+
		"\2\2\u00e0\u0430\3\2\2\2\u00e2\u0432\3\2\2\2\u00e4\u0438\3\2\2\2\u00e6"+
		"\u0440\3\2\2\2\u00e8\u0445\3\2\2\2\u00ea\u0454\3\2\2\2\u00ec\u0456\3\2"+
		"\2\2\u00ee\u0468\3\2\2\2\u00f0\u046a\3\2\2\2\u00f2\u046c\3\2\2\2\u00f4"+
		"\u0473\3\2\2\2\u00f6\u047a\3\2\2\2\u00f8\u047d\3\2\2\2\u00fa\u0487\3\2"+
		"\2\2\u00fc\u048b\3\2\2\2\u00fe\u048d\3\2\2\2\u0100\u0495\3\2\2\2\u0102"+
		"\u0497\3\2\2\2\u0104\u0499\3\2\2\2\u0106\u049c\3\2\2\2\u0108\u04a0\3\2"+
		"\2\2\u010a\u04a2\3\2\2\2\u010c\u04a4\3\2\2\2\u010e\u04a6\3\2\2\2\u0110"+
		"\u04a9\3\2\2\2\u0112\u04ac\3\2\2\2\u0114\u04af\3\2\2\2\u0116\u04b2\3\2"+
		"\2\2\u0118\u04b5\3\2\2\2\u011a\u04b8\3\2\2\2\u011c\u04ba\3\2\2\2\u011e"+
		"\u04bc\3\2\2\2\u0120\u04cc\3\2\2\2\u0122\u04d0\3\2\2\2\u0124\u04d4\3\2"+
		"\2\2\u0126\u04d6\3\2\2\2\u0128\u04da\3\2\2\2\u012a\u04e0\3\2\2\2\u012c"+
		"\u04ea\3\2\2\2\u012e\u04ec\3\2\2\2\u0130\u04ee\3\2\2\2\u0132\u04f0\3\2"+
		"\2\2\u0134\u04f2\3\2\2\2\u0136\u0138\5\4\3\2\u0137\u0136\3\2\2\2\u0138"+
		"\u0139\3\2\2\2\u0139\u0137\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u013b\3\2"+
		"\2\2\u013b\u013c\7\2\2\3\u013c\3\3\2\2\2\u013d\u013f\5\64\33\2\u013e\u0140"+
		"\7-\2\2\u013f\u013e\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u0146\3\2\2\2\u0141"+
		"\u0143\5\u00a4S\2\u0142\u0144\7-\2\2\u0143\u0142\3\2\2\2\u0143\u0144\3"+
		"\2\2\2\u0144\u0146\3\2\2\2\u0145\u013d\3\2\2\2\u0145\u0141\3\2\2\2\u0146"+
		"\5\3\2\2\2\u0147\u014c\3\2\2\2\u0148\u014c\5\b\5\2\u0149\u014c\5\n\6\2"+
		"\u014a\u014c\5\22\n\2\u014b\u0147\3\2\2\2\u014b\u0148\3\2\2\2\u014b\u0149"+
		"\3\2\2\2\u014b\u014a\3\2\2\2\u014c\7\3\2\2\2\u014d\u014e\7\3\2\2\u014e"+
		"\u014f\5\u008cG\2\u014f\u0150\7\4\2\2\u0150\u0151\5\u00a4S\2\u0151\u0152"+
		"\5\66\34\2\u0152\t\3\2\2\2\u0153\u0154\7\5\2\2\u0154\u0155\5\f\7\2\u0155"+
		"\u0156\5\66\34\2\u0156\13\3\2\2\2\u0157\u015c\5\16\b\2\u0158\u0159\7+"+
		"\2\2\u0159\u015b\5\16\b\2\u015a\u0158\3\2\2\2\u015b\u015e\3\2\2\2\u015c"+
		"\u015a\3\2\2\2\u015c\u015d\3\2\2\2\u015d\r\3\2\2\2\u015e\u015c\3\2\2\2"+
		"\u015f\u0162\5\u00a4S\2\u0160\u0162\5\20\t\2\u0161\u015f\3\2\2\2\u0161"+
		"\u0160\3\2\2\2\u0162\17\3\2\2\2\u0163\u0164\7\6\2\2\u0164\u0165\5\u008c"+
		"G\2\u0165\u0166\5V,\2\u0166\u016c\3\2\2\2\u0167\u0168\7\7\2\2\u0168\u0169"+
		"\5\u008cG\2\u0169\u016a\5V,\2\u016a\u016c\3\2\2\2\u016b\u0163\3\2\2\2"+
		"\u016b\u0167\3\2\2\2\u016c\21\3\2\2\2\u016d\u016e\7\b\2\2\u016e\u016f"+
		"\5\66\34\2\u016f\u0170\7\5\2\2\u0170\u0171\5\u00a4S\2\u0171\23\3\2\2\2"+
		"\u0172\u0175\5\26\f\2\u0173\u0175\5\32\16\2\u0174\u0172\3\2\2\2\u0174"+
		"\u0173\3\2\2\2\u0175\25\3\2\2\2\u0176\u0177\7\t\2\2\u0177\u0178\5\f\7"+
		"\2\u0178\u017a\5\66\34\2\u0179\u017b\5\30\r\2\u017a\u0179\3\2\2\2\u017a"+
		"\u017b\3\2\2\2\u017b\27\3\2\2\2\u017c\u017d\7\n\2\2\u017d\u0181\5\66\34"+
		"\2\u017e\u017f\7\n\2\2\u017f\u0181\5\26\f\2\u0180\u017c\3\2\2\2\u0180"+
		"\u017e\3\2\2\2\u0181\31\3\2\2\2\u0182\u0183\7\13\2\2\u0183\u0184\5\u00a4"+
		"S\2\u0184\u0186\7%\2\2\u0185\u0187\5\34\17\2\u0186\u0185\3\2\2\2\u0186"+
		"\u0187\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u0189\7(\2\2\u0189\33\3\2\2\2"+
		"\u018a\u018c\5\36\20\2\u018b\u018d\5\34\17\2\u018c\u018b\3\2\2\2\u018c"+
		"\u018d\3\2\2\2\u018d\35\3\2\2\2\u018e\u018f\5\u008cG\2\u018f\u0192\7\f"+
		"\2\2\u0190\u0193\5\66\34\2\u0191\u0193\5\u00a4S\2\u0192\u0190\3\2\2\2"+
		"\u0192\u0191\3\2\2\2\u0193\37\3\2\2\2\u0194\u0198\5\"\22\2\u0195\u0198"+
		"\5$\23\2\u0196\u0198\5&\24\2\u0197\u0194\3\2\2\2\u0197\u0195\3\2\2\2\u0197"+
		"\u0196\3\2\2\2\u0198!\3\2\2\2\u0199\u019b\7\r\2\2\u019a\u019c\7-\2\2\u019b"+
		"\u019a\3\2\2\2\u019b\u019c\3\2\2\2\u019c#\3\2\2\2\u019d\u019f\7\16\2\2"+
		"\u019e\u01a0\7-\2\2\u019f\u019e\3\2\2\2\u019f\u01a0\3\2\2\2\u01a0%\3\2"+
		"\2\2\u01a1\u01a3\7\17\2\2\u01a2\u01a4\5\u00a4S\2\u01a3\u01a2\3\2\2\2\u01a3"+
		"\u01a4\3\2\2\2\u01a4\'\3\2\2\2\u01a5\u01a6\7.\2\2\u01a6\u01a7\5*\26\2"+
		"\u01a7\u01a8\7/\2\2\u01a8)\3\2\2\2\u01a9\u01ae\5,\27\2\u01aa\u01ab\7+"+
		"\2\2\u01ab\u01ad\5,\27\2\u01ac\u01aa\3\2\2\2\u01ad\u01b0\3\2\2\2\u01ae"+
		"\u01ac\3\2\2\2\u01ae\u01af\3\2\2\2\u01af+\3\2\2\2\u01b0\u01ae\3\2\2\2"+
		"\u01b1\u01b7\5\u00e0q\2\u01b2\u01b3\5\u00e0q\2\u01b3\u01b4\7,\2\2\u01b4"+
		"\u01b5\5\u00dco\2\u01b5\u01b7\3\2\2\2\u01b6\u01b1\3\2\2\2\u01b6\u01b2"+
		"\3\2\2\2\u01b7-\3\2\2\2\u01b8\u01b9\7.\2\2\u01b9\u01ba\5\60\31\2\u01ba"+
		"\u01bb\7/\2\2\u01bb/\3\2\2\2\u01bc\u01c1\5\62\32\2\u01bd\u01be\7+\2\2"+
		"\u01be\u01c0\5\62\32\2\u01bf\u01bd\3\2\2\2\u01c0\u01c3\3\2\2\2\u01c1\u01bf"+
		"\3\2\2\2\u01c1\u01c2\3\2\2\2\u01c2\61\3\2\2\2\u01c3\u01c1\3\2\2\2\u01c4"+
		"\u01c6\5\u00d8m\2\u01c5\u01c7\5\u00a2R\2\u01c6\u01c5\3\2\2\2\u01c6\u01c7"+
		"\3\2\2\2\u01c7\63\3\2\2\2\u01c8\u01d1\58\35\2\u01c9\u01d1\5<\37\2\u01ca"+
		"\u01d1\5P)\2\u01cb\u01d1\5X-\2\u01cc\u01d1\5^\60\2\u01cd\u01d1\5p9\2\u01ce"+
		"\u01d1\5x=\2\u01cf\u01d1\5\u0082B\2\u01d0\u01c8\3\2\2\2\u01d0\u01c9\3"+
		"\2\2\2\u01d0\u01ca\3\2\2\2\u01d0\u01cb\3\2\2\2\u01d0\u01cc\3\2\2\2\u01d0"+
		"\u01cd\3\2\2\2\u01d0\u01ce\3\2\2\2\u01d0\u01cf\3\2\2\2\u01d1\65\3\2\2"+
		"\2\u01d2\u01d4\7%\2\2\u01d3\u01d5\5\4\3\2\u01d4\u01d3\3\2\2\2\u01d5\u01d6"+
		"\3\2\2\2\u01d6\u01d4\3\2\2\2\u01d6\u01d7\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8"+
		"\u01d9\7(\2\2\u01d9\67\3\2\2\2\u01da\u01db\7\20\2\2\u01db\u01dc\5:\36"+
		"\2\u01dc\u01dd\5\u00bc_\2\u01dd9\3\2\2\2\u01de\u01e3\7#\2\2\u01df\u01e0"+
		"\7$\2\2\u01e0\u01e2\7#\2\2\u01e1\u01df\3\2\2\2\u01e2\u01e5\3\2\2\2\u01e3"+
		"\u01e1\3\2\2\2\u01e3\u01e4\3\2\2\2\u01e4;\3\2\2\2\u01e5\u01e3\3\2\2\2"+
		"\u01e6\u01e7\7\21\2\2\u01e7\u01eb\5> \2\u01e8\u01ec\5B\"\2\u01e9\u01ec"+
		"\5F$\2\u01ea\u01ec\5H%\2\u01eb\u01e8\3\2\2\2\u01eb\u01e9\3\2\2\2\u01eb"+
		"\u01ea\3\2\2\2\u01eb\u01ec\3\2\2\2\u01ec=\3\2\2\2\u01ed\u01f2\5@!\2\u01ee"+
		"\u01ef\7$\2\2\u01ef\u01f1\5@!\2\u01f0\u01ee\3\2\2\2\u01f1\u01f4\3\2\2"+
		"\2\u01f2\u01f0\3\2\2\2\u01f2\u01f3\3\2\2\2\u01f3?\3\2\2\2\u01f4\u01f2"+
		"\3\2\2\2\u01f5\u01f8\5\u00fa~\2\u01f6\u01f8\5\u0120\u0091\2\u01f7\u01f5"+
		"\3\2\2\2\u01f7\u01f6\3\2\2\2\u01f8A\3\2\2\2\u01f9\u01fa\7\22\2\2\u01fa"+
		"\u01fb\5\u00fa~\2\u01fbC\3\2\2\2\u01fc\u01fd\7\22\2\2\u01fd\u01fe\5\u00e0"+
		"q\2\u01feE\3\2\2\2\u01ff\u0200\7$\2\2\u0200\u0202\5\u00e0q\2\u0201\u0203"+
		"\5D#\2\u0202\u0201\3\2\2\2\u0202\u0203\3\2\2\2\u0203G\3\2\2\2\u0204\u0205"+
		"\7$\2\2\u0205\u0206\7%\2\2\u0206\u0207\5J&\2\u0207\u0208\7(\2\2\u0208"+
		"I\3\2\2\2\u0209\u020c\5L\'\2\u020a\u020c\5N(\2\u020b\u0209\3\2\2\2\u020b"+
		"\u020a\3\2\2\2\u020c\u0214\3\2\2\2\u020d\u0210\7+\2\2\u020e\u0211\5L\'"+
		"\2\u020f\u0211\5N(\2\u0210\u020e\3\2\2\2\u0210\u020f\3\2\2\2\u0211\u0213"+
		"\3\2\2\2\u0212\u020d\3\2\2\2\u0213\u0216\3\2\2\2\u0214\u0212\3\2\2\2\u0214"+
		"\u0215\3\2\2\2\u0215K\3\2\2\2\u0216\u0214\3\2\2\2\u0217\u0219\5\u00fa"+
		"~\2\u0218\u021a\5B\"\2\u0219\u0218\3\2\2\2\u0219\u021a\3\2\2\2\u021aM"+
		"\3\2\2\2\u021b\u021d\5\u00e0q\2\u021c\u021e\5D#\2\u021d\u021c\3\2\2\2"+
		"\u021d\u021e\3\2\2\2\u021eO\3\2\2\2\u021f\u0221\5\u00a2R\2\u0220\u021f"+
		"\3\2\2\2\u0220\u0221\3\2\2\2\u0221\u0222\3\2\2\2\u0222\u0223\7\6\2\2\u0223"+
		"\u0224\5R*\2\u0224Q\3\2\2\2\u0225\u022a\5T+\2\u0226\u0227\7+\2\2\u0227"+
		"\u0229\5T+\2\u0228\u0226\3\2\2\2\u0229\u022c\3\2\2\2\u022a\u0228\3\2\2"+
		"\2\u022a\u022b\3\2\2\2\u022bS\3\2\2\2\u022c\u022a\3\2\2\2\u022d\u022f"+
		"\5\u008cG\2\u022e\u0230\5V,\2\u022f\u022e\3\2\2\2\u022f\u0230\3\2\2\2"+
		"\u0230U\3\2\2\2\u0231\u0232\5\u010a\u0086\2\u0232\u0233\5\u00a4S\2\u0233"+
		"W\3\2\2\2\u0234\u0236\5\u00a2R\2\u0235\u0234\3\2\2\2\u0235\u0236\3\2\2"+
		"\2\u0236\u0237\3\2\2\2\u0237\u0238\7\23\2\2\u0238\u023a\5Z.\2\u0239\u023b"+
		"\5(\25\2\u023a\u0239\3\2\2\2\u023a\u023b\3\2\2\2\u023b\u023c\3\2\2\2\u023c"+
		"\u023d\5\\/\2\u023dY\3\2\2\2\u023e\u023f\5\u00e0q\2\u023f[\3\2\2\2\u0240"+
		"\u0241\5\u010a\u0086\2\u0241\u0242\5\u00d8m\2\u0242]\3\2\2\2\u0243\u0244"+
		"\5`\61\2\u0244\u0246\5b\62\2\u0245\u0247\5(\25\2\u0246\u0245\3\2\2\2\u0246"+
		"\u0247\3\2\2\2\u0247\u0248\3\2\2\2\u0248\u024a\5d\63\2\u0249\u024b\5h"+
		"\65\2\u024a\u0249\3\2\2\2\u024a\u024b\3\2\2\2\u024b_\3\2\2\2\u024c\u024e"+
		"\5\u00a2R\2\u024d\u024c\3\2\2\2\u024d\u024e\3\2\2\2\u024e\u024f\3\2\2"+
		"\2\u024f\u0250\7\24\2\2\u0250a\3\2\2\2\u0251\u0254\5\u00fa~\2\u0252\u0254"+
		"\5\u0120\u0091\2\u0253\u0251\3\2\2\2\u0253\u0252\3\2\2\2\u0254c\3\2\2"+
		"\2\u0255\u0257\5j\66\2\u0256\u0258\5f\64\2\u0257\u0256\3\2\2\2\u0257\u0258"+
		"\3\2\2\2\u0258e\3\2\2\2\u0259\u025a\5\u0114\u008b\2\u025a\u025c\5\u00d8"+
		"m\2\u025b\u025d\5\u00a2R\2\u025c\u025b\3\2\2\2\u025c\u025d\3\2\2\2\u025d"+
		"g\3\2\2\2\u025e\u025f\5\66\34\2\u025fi\3\2\2\2\u0260\u0261\7&\2\2\u0261"+
		"\u0267\7)\2\2\u0262\u0263\7&\2\2\u0263\u0264\5l\67\2\u0264\u0265\7)\2"+
		"\2\u0265\u0267\3\2\2\2\u0266\u0260\3\2\2\2\u0266\u0262\3\2\2\2\u0267k"+
		"\3\2\2\2\u0268\u026d\5n8\2\u0269\u026a\7+\2\2\u026a\u026c\5n8\2\u026b"+
		"\u0269\3\2\2\2\u026c\u026f\3\2\2\2\u026d\u026b\3\2\2\2\u026d\u026e\3\2"+
		"\2\2\u026em\3\2\2\2\u026f\u026d\3\2\2\2\u0270\u0271\5\u00fc\177\2\u0271"+
		"\u0273\5\u00dan\2\u0272\u0274\5V,\2\u0273\u0272\3\2\2\2\u0273\u0274\3"+
		"\2\2\2\u0274\u027c\3\2\2\2\u0275\u0276\5\u00fc\177\2\u0276\u0277\5\u00d8"+
		"m\2\u0277\u0279\7>\2\2\u0278\u027a\5\u00a2R\2\u0279\u0278\3\2\2\2\u0279"+
		"\u027a\3\2\2\2\u027a\u027c\3\2\2\2\u027b\u0270\3\2\2\2\u027b\u0275\3\2"+
		"\2\2\u027co\3\2\2\2\u027d\u027f\5\u0106\u0084\2\u027e\u027d\3\2\2\2\u027e"+
		"\u027f\3\2\2\2\u027f\u0281\3\2\2\2\u0280\u0282\5\u00a2R\2\u0281\u0280"+
		"\3\2\2\2\u0281\u0282\3\2\2\2\u0282\u0283\3\2\2\2\u0283\u0284\7\25\2\2"+
		"\u0284\u0286\5r:\2\u0285\u0287\5(\25\2\u0286\u0285\3\2\2\2\u0286\u0287"+
		"\3\2\2\2\u0287\u0289\3\2\2\2\u0288\u028a\5\u00f6|\2\u0289\u0288\3\2\2"+
		"\2\u0289\u028a\3\2\2\2\u028a\u028b\3\2\2\2\u028b\u028d\7%\2\2\u028c\u028e"+
		"\5t;\2\u028d\u028c\3\2\2\2\u028d\u028e\3\2\2\2\u028e\u028f\3\2\2\2\u028f"+
		"\u0290\7(\2\2\u0290q\3\2\2\2\u0291\u0292\5\u00e0q\2\u0292s\3\2\2\2\u0293"+
		"\u0295\5v<\2\u0294\u0293\3\2\2\2\u0295\u0296\3\2\2\2\u0296\u0294\3\2\2"+
		"\2\u0296\u0297\3\2\2\2\u0297u\3\2\2\2\u0298\u029a\5\u0106\u0084\2\u0299"+
		"\u0298\3\2\2\2\u0299\u029a\3\2\2\2\u029a\u029b\3\2\2\2\u029b\u029d\7#"+
		"\2\2\u029c\u029e\5\u00a2R\2\u029d\u029c\3\2\2\2\u029d\u029e\3\2\2\2\u029e"+
		"\u02a0\3\2\2\2\u029f\u02a1\5V,\2\u02a0\u029f\3\2\2\2\u02a0\u02a1\3\2\2"+
		"\2\u02a1\u02a3\3\2\2\2\u02a2\u02a4\5\u0108\u0085\2\u02a3\u02a2\3\2\2\2"+
		"\u02a3\u02a4\3\2\2\2\u02a4w\3\2\2\2\u02a5\u02a7\5\u0106\u0084\2\u02a6"+
		"\u02a5\3\2\2\2\u02a6\u02a7\3\2\2\2\u02a7\u02a9\3\2\2\2\u02a8\u02aa\5\u00a2"+
		"R\2\u02a9\u02a8\3\2\2\2\u02a9\u02aa\3\2\2\2\u02aa\u02ab\3\2\2\2\u02ab"+
		"\u02ac\7\23\2\2\u02ac\u02ae\5z>\2\u02ad\u02af\5(\25\2\u02ae\u02ad\3\2"+
		"\2\2\u02ae\u02af\3\2\2\2\u02af\u02b1\3\2\2\2\u02b0\u02b2\5\u00f6|\2\u02b1"+
		"\u02b0\3\2\2\2\u02b1\u02b2\3\2\2\2\u02b2\u02b4\3\2\2\2\u02b3\u02b5\5|"+
		"?\2\u02b4\u02b3\3\2\2\2\u02b4\u02b5\3\2\2\2\u02b5y\3\2\2\2\u02b6\u02b7"+
		"\5\u00e0q\2\u02b7{\3\2\2\2\u02b8\u02bc\7%\2\2\u02b9\u02bb\5~@\2\u02ba"+
		"\u02b9\3\2\2\2\u02bb\u02be\3\2\2\2\u02bc\u02ba\3\2\2\2\u02bc\u02bd\3\2"+
		"\2\2\u02bd\u02bf\3\2\2\2\u02be\u02bc\3\2\2\2\u02bf\u02c0\7(\2\2\u02c0"+
		"}\3\2\2\2\u02c1\u02c7\5\u0080A\2\u02c2\u02c7\5x=\2\u02c3\u02c7\5p9\2\u02c4"+
		"\u02c7\5P)\2\u02c5\u02c7\5X-\2\u02c6\u02c1\3\2\2\2\u02c6\u02c2\3\2\2\2"+
		"\u02c6\u02c3\3\2\2\2\u02c6\u02c4\3\2\2\2\u02c6\u02c5\3\2\2\2\u02c7\177"+
		"\3\2\2\2\u02c8\u02ca\5\u0106\u0084\2\u02c9\u02c8\3\2\2\2\u02c9\u02ca\3"+
		"\2\2\2\u02ca\u02cb\3\2\2\2\u02cb\u02cc\7#\2\2\u02cc\u02ce\5\u00dan\2\u02cd"+
		"\u02cf\5V,\2\u02ce\u02cd\3\2\2\2\u02ce\u02cf\3\2\2\2\u02cf\u02d1\3\2\2"+
		"\2\u02d0\u02d2\5\u0108\u0085\2\u02d1\u02d0\3\2\2\2\u02d1\u02d2\3\2\2\2"+
		"\u02d2\u0081\3\2\2\2\u02d3\u02d5\5\u0106\u0084\2\u02d4\u02d3\3\2\2\2\u02d4"+
		"\u02d5\3\2\2\2\u02d5\u02d7\3\2\2\2\u02d6\u02d8\5\u00a2R\2\u02d7\u02d6"+
		"\3\2\2\2\u02d7\u02d8\3\2\2\2\u02d8\u02d9\3\2\2\2\u02d9\u02da\7\26\2\2"+
		"\u02da\u02dc\5\u0084C\2\u02db\u02dd\5\u00f6|\2\u02dc\u02db\3\2\2\2\u02dc"+
		"\u02dd\3\2\2\2\u02dd\u02de\3\2\2\2\u02de\u02df\5\u0086D\2\u02df\u0083"+
		"\3\2\2\2\u02e0\u02e1\5\u00e0q\2\u02e1\u0085\3\2\2\2\u02e2\u02e6\7%\2\2"+
		"\u02e3\u02e5\5\u0088E\2\u02e4\u02e3\3\2\2\2\u02e5\u02e8\3\2\2\2\u02e6"+
		"\u02e4\3\2\2\2\u02e6\u02e7\3\2\2\2\u02e7\u02e9\3\2\2\2\u02e8\u02e6\3\2"+
		"\2\2\u02e9\u02ea\7(\2\2\u02ea\u0087\3\2\2\2\u02eb\u02ee\5\u008aF\2\u02ec"+
		"\u02ee\5X-\2\u02ed\u02eb\3\2\2\2\u02ed\u02ec\3\2\2\2\u02ee\u0089\3\2\2"+
		"\2\u02ef\u02f1\5\u0106\u0084\2\u02f0\u02ef\3\2\2\2\u02f0\u02f1\3\2\2\2"+
		"\u02f1\u02f3\3\2\2\2\u02f2\u02f4\5\u00a2R\2\u02f3\u02f2\3\2\2\2\u02f3"+
		"\u02f4\3\2\2\2\u02f4\u02f5\3\2\2\2\u02f5\u02f7\5b\62\2\u02f6\u02f8\5("+
		"\25\2\u02f7\u02f6\3\2\2\2\u02f7\u02f8\3\2\2\2\u02f8\u02f9\3\2\2\2\u02f9"+
		"\u02fa\5d\63\2\u02fa\u008b\3\2\2\2\u02fb\u02fc\bG\1\2\u02fc\u02fe\5\u008e"+
		"H\2\u02fd\u02ff\5\u00dan\2\u02fe\u02fd\3\2\2\2\u02fe\u02ff\3\2\2\2\u02ff"+
		"\u030d\3\2\2\2\u0300\u0302\5\u0090I\2\u0301\u0303\5\u00dan\2\u0302\u0301"+
		"\3\2\2\2\u0302\u0303\3\2\2\2\u0303\u030d\3\2\2\2\u0304\u0306\5\u0092J"+
		"\2\u0305\u0307\5\u00dan\2\u0306\u0305\3\2\2\2\u0306\u0307\3\2\2\2\u0307"+
		"\u030d\3\2\2\2\u0308\u030d\5\u0098M\2\u0309\u030a\7\27\2\2\u030a\u030d"+
		"\5\u00d8m\2\u030b\u030d\5\u009aN\2\u030c\u02fb\3\2\2\2\u030c\u0300\3\2"+
		"\2\2\u030c\u0304\3\2\2\2\u030c\u0308\3\2\2\2\u030c\u0309\3\2\2\2\u030c"+
		"\u030b\3\2\2\2\u030d\u0313\3\2\2\2\u030e\u030f\f\4\2\2\u030f\u0310\7\22"+
		"\2\2\u0310\u0312\5\u00d8m\2\u0311\u030e\3\2\2\2\u0312\u0315\3\2\2\2\u0313"+
		"\u0311\3\2\2\2\u0313\u0314\3\2\2\2\u0314\u008d\3\2\2\2\u0315\u0313\3\2"+
		"\2\2\u0316\u0317\7\60\2\2\u0317\u008f\3\2\2\2\u0318\u0319\5\u00fa~\2\u0319"+
		"\u0091\3\2\2\2\u031a\u031c\7&\2\2\u031b\u031d\5\u0094K\2\u031c\u031b\3"+
		"\2\2\2\u031c\u031d\3\2\2\2\u031d\u031e\3\2\2\2\u031e\u031f\7)\2\2\u031f"+
		"\u0093\3\2\2\2\u0320\u0325\5\u0096L\2\u0321\u0322\7+\2\2\u0322\u0324\5"+
		"\u0096L\2\u0323\u0321\3\2\2\2\u0324\u0327\3\2\2\2\u0325\u0323\3\2\2\2"+
		"\u0325\u0326\3\2\2\2\u0326\u0095\3\2\2\2\u0327\u0325\3\2\2\2\u0328\u0329"+
		"\5\u008cG\2\u0329\u0097\3\2\2\2\u032a\u032b\5\u0090I\2\u032b\u032c\7\62"+
		"\2\2\u032c\u0099\3\2\2\2\u032d\u032e\5\u00a4S\2\u032e\u009b\3\2\2\2\u032f"+
		"\u0330\7\63\2\2\u0330\u0332\5\u009eP\2\u0331\u0333\5\u00a0Q\2\u0332\u0331"+
		"\3\2\2\2\u0332\u0333\3\2\2\2\u0333\u009d\3\2\2\2\u0334\u0335\t\2\2\2\u0335"+
		"\u009f\3\2\2\2\u0336\u0337\7&\2\2\u0337\u0338\5\u00a6T\2\u0338\u0339\7"+
		")\2\2\u0339\u00a1\3\2\2\2\u033a\u033c\5\u009cO\2\u033b\u033a\3\2\2\2\u033c"+
		"\u033d\3\2\2\2\u033d\u033b\3\2\2\2\u033d\u033e\3\2\2\2\u033e\u00a3\3\2"+
		"\2\2\u033f\u0341\5\u00a8U\2\u0340\u0342\5\u00acW\2\u0341\u0340\3\2\2\2"+
		"\u0341\u0342\3\2\2\2\u0342\u00a5\3\2\2\2\u0343\u0348\5\u00a4S\2\u0344"+
		"\u0345\7+\2\2\u0345\u0347\5\u00a4S\2\u0346\u0344\3\2\2\2\u0347\u034a\3"+
		"\2\2\2\u0348\u0346\3\2\2\2\u0348\u0349\3\2\2\2\u0349\u00a7\3\2\2\2\u034a"+
		"\u0348\3\2\2\2\u034b\u034c\5\u011c\u008f\2\u034c\u034d\5\u00ccg\2\u034d"+
		"\u0350\3\2\2\2\u034e\u0350\5\u00ccg\2\u034f\u034b\3\2\2\2\u034f\u034e"+
		"\3\2\2\2\u0350\u00a9\3\2\2\2\u0351\u0352\5\u011a\u008e\2\u0352\u0353\5"+
		"\u00a8U\2\u0353\u035c\3\2\2\2\u0354\u0355\5\u010a\u0086\2\u0355\u0356"+
		"\5\u00a8U\2\u0356\u035c\3\2\2\2\u0357\u0358\5\u00aeX\2\u0358\u0359\5\u00a8"+
		"U\2\u0359\u035c\3\2\2\2\u035a\u035c\5\u00b0Y\2\u035b\u0351\3\2\2\2\u035b"+
		"\u0354\3\2\2\2\u035b\u0357\3\2\2\2\u035b\u035a\3\2\2\2\u035c\u00ab\3\2"+
		"\2\2\u035d\u035f\5\u00aaV\2\u035e\u035d\3\2\2\2\u035f\u0360\3\2\2\2\u0360"+
		"\u035e\3\2\2\2\u0360\u0361\3\2\2\2\u0361\u00ad\3\2\2\2\u0362\u0363\7\62"+
		"\2\2\u0363\u0364\5\u00a4S\2\u0364\u0365\7,\2\2\u0365\u00af\3\2\2\2\u0366"+
		"\u0367\7\27\2\2\u0367\u036b\5\u00d8m\2\u0368\u0369\7\22\2\2\u0369\u036b"+
		"\5\u00d8m\2\u036a\u0366\3\2\2\2\u036a\u0368\3\2\2\2\u036b\u00b1\3\2\2"+
		"\2\u036c\u036e\5\u00fa~\2\u036d\u036f\5.\30\2\u036e\u036d\3\2\2\2\u036e"+
		"\u036f\3\2\2\2\u036f\u0376\3\2\2\2\u0370\u0376\5\u00b4[\2\u0371\u0376"+
		"\5\u00c4c\2\u0372\u0376\5\u00c6d\2\u0373\u0376\5\u00c2b\2\u0374\u0376"+
		"\5\u00caf\2\u0375\u036c\3\2\2\2\u0375\u0370\3\2\2\2\u0375\u0371\3\2\2"+
		"\2\u0375\u0372\3\2\2\2\u0375\u0373\3\2\2\2\u0375\u0374\3\2\2\2\u0376\u00b3"+
		"\3\2\2\2\u0377\u037b\5\u012a\u0096\2\u0378\u037b\5\u00b6\\\2\u0379\u037b"+
		"\5\u00bc_\2\u037a\u0377\3\2\2\2\u037a\u0378\3\2\2\2\u037a\u0379\3\2\2"+
		"\2\u037b\u00b5\3\2\2\2\u037c\u037e\7\'\2\2\u037d\u037f\5\u00b8]\2\u037e"+
		"\u037d\3\2\2\2\u037e\u037f\3\2\2\2\u037f\u0380\3\2\2\2\u0380\u0381\7*"+
		"\2\2\u0381\u00b7\3\2\2\2\u0382\u0384\5\u00ba^\2\u0383\u0385\7+\2\2\u0384"+
		"\u0383\3\2\2\2\u0384\u0385\3\2\2\2\u0385\u038b\3\2\2\2\u0386\u0387\5\u00ba"+
		"^\2\u0387\u0388\7+\2\2\u0388\u0389\5\u00b8]\2\u0389\u038b\3\2\2\2\u038a"+
		"\u0382\3\2\2\2\u038a\u0386\3\2\2\2\u038b\u00b9\3\2\2\2\u038c\u038d\5\u00a4"+
		"S\2\u038d\u00bb\3\2\2\2\u038e\u038f\7%\2\2\u038f\u0390\5\u00be`\2\u0390"+
		"\u0391\7(\2\2\u0391\u0395\3\2\2\2\u0392\u0393\7%\2\2\u0393\u0395\7(\2"+
		"\2\u0394\u038e\3\2\2\2\u0394\u0392\3\2\2\2\u0395\u00bd\3\2\2\2\u0396\u0398"+
		"\5\u00c0a\2\u0397\u0399\7+\2\2\u0398\u0397\3\2\2\2\u0398\u0399\3\2\2\2"+
		"\u0399\u039f\3\2\2\2\u039a\u039b\5\u00c0a\2\u039b\u039c\7+\2\2\u039c\u039d"+
		"\5\u00be`\2\u039d\u039f\3\2\2\2\u039e\u0396\3\2\2\2\u039e\u039a\3\2\2"+
		"\2\u039f\u00bf\3\2\2\2\u03a0\u03a1\5\u00a4S\2\u03a1\u03a2\7,\2\2\u03a2"+
		"\u03a3\5\u00a4S\2\u03a3\u00c1\3\2\2\2\u03a4\u03a5\7$\2\2\u03a5\u03a6\5"+
		"\u00fc\177\2\u03a6\u00c3\3\2\2\2\u03a7\u03a8\7&\2\2\u03a8\u03a9\5\u00a4"+
		"S\2\u03a9\u03aa\7)\2\2\u03aa\u00c5\3\2\2\2\u03ab\u03ac\7&\2\2\u03ac\u03b8"+
		"\7)\2\2\u03ad\u03ae\7&\2\2\u03ae\u03b1\5\u00c8e\2\u03af\u03b0\7+\2\2\u03b0"+
		"\u03b2\5\u00c8e\2\u03b1\u03af\3\2\2\2\u03b2\u03b3\3\2\2\2\u03b3\u03b1"+
		"\3\2\2\2\u03b3\u03b4\3\2\2\2\u03b4\u03b5\3\2\2\2\u03b5\u03b6\7)\2\2\u03b6"+
		"\u03b8\3\2\2\2\u03b7\u03ab\3\2\2\2\u03b7\u03ad\3\2\2\2\u03b8\u00c7\3\2"+
		"\2\2\u03b9\u03bf\5\u00a4S\2\u03ba\u03bb\5\u00fc\177\2\u03bb\u03bc\7,\2"+
		"\2\u03bc\u03bd\5\u00a4S\2\u03bd\u03bf\3\2\2\2\u03be\u03b9\3\2\2\2\u03be"+
		"\u03ba\3\2\2\2\u03bf\u00c9\3\2\2\2\u03c0\u03c1\7\60\2\2\u03c1\u00cb\3"+
		"\2\2\2\u03c2\u03c3\bg\1\2\u03c3\u03c4\5\u00b2Z\2\u03c4\u03e5\3\2\2\2\u03c5"+
		"\u03c6\f\t\2\2\u03c6\u03e4\5\u011e\u0090\2\u03c7\u03c8\f\b\2\2\u03c8\u03e4"+
		"\5\u00ceh\2\u03c9\u03ca\f\7\2\2\u03ca\u03cb\7$\2\2\u03cb\u03e4\7E\2\2"+
		"\u03cc\u03cd\f\6\2\2\u03cd\u03ce\7$\2\2\u03ce\u03d0\5\u00fa~\2\u03cf\u03d1"+
		"\5.\30\2\u03d0\u03cf\3\2\2\2\u03d0\u03d1\3\2\2\2\u03d1\u03e4\3\2\2\2\u03d2"+
		"\u03d3\f\5\2\2\u03d3\u03d4\7$\2\2\u03d4\u03d5\5\u00fa~\2\u03d5\u03d6\7"+
		"&\2\2\u03d6\u03d7\5\u00d4k\2\u03d7\u03d8\7)\2\2\u03d8\u03e4\3\2\2\2\u03d9"+
		"\u03da\f\4\2\2\u03da\u03db\7&\2\2\u03db\u03dc\5\u00d4k\2\u03dc\u03dd\7"+
		")\2\2\u03dd\u03e4\3\2\2\2\u03de\u03df\f\3\2\2\u03df\u03e0\7\'\2\2\u03e0"+
		"\u03e1\5\u00a6T\2\u03e1\u03e2\7*\2\2\u03e2\u03e4\3\2\2\2\u03e3\u03c5\3"+
		"\2\2\2\u03e3\u03c7\3\2\2\2\u03e3\u03c9\3\2\2\2\u03e3\u03cc\3\2\2\2\u03e3"+
		"\u03d2\3\2\2\2\u03e3\u03d9\3\2\2\2\u03e3\u03de\3\2\2\2\u03e4\u03e7\3\2"+
		"\2\2\u03e5\u03e3\3\2\2\2\u03e5\u03e6\3\2\2\2\u03e6\u00cd\3\2\2\2\u03e7"+
		"\u03e5\3\2\2\2\u03e8\u03e9\7&\2\2\u03e9\u03ef\7)\2\2\u03ea\u03eb\7&\2"+
		"\2\u03eb\u03ec\5\u00d0i\2\u03ec\u03ed\7)\2\2\u03ed\u03ef\3\2\2\2\u03ee"+
		"\u03e8\3\2\2\2\u03ee\u03ea\3\2\2\2\u03ef\u00cf\3\2\2\2\u03f0\u03f5\5\u00d2"+
		"j\2\u03f1\u03f2\7+\2\2\u03f2\u03f4\5\u00d2j\2\u03f3\u03f1\3\2\2\2\u03f4"+
		"\u03f7\3\2\2\2\u03f5\u03f3\3\2\2\2\u03f5\u03f6\3\2\2\2\u03f6\u00d1\3\2"+
		"\2\2\u03f7\u03f5\3\2\2\2\u03f8\u0403\5\u00a4S\2\u03f9\u03fa\5\u00fc\177"+
		"\2\u03fa\u03fb\7,\2\2\u03fb\u03fc\5\u00a4S\2\u03fc\u0403\3\2\2\2\u03fd"+
		"\u0403\5\u0120\u0091\2\u03fe\u03ff\5\u00fc\177\2\u03ff\u0400\7,\2\2\u0400"+
		"\u0401\5\u0120\u0091\2\u0401\u0403\3\2\2\2\u0402\u03f8\3\2\2\2\u0402\u03f9"+
		"\3\2\2\2\u0402\u03fd\3\2\2\2\u0402\u03fe\3\2\2\2\u0403\u00d3\3\2\2\2\u0404"+
		"\u0408\5\u00d6l\2\u0405\u0407\5\u00d6l\2\u0406\u0405\3\2\2\2\u0407\u040a"+
		"\3\2\2\2\u0408\u0406\3\2\2\2\u0408\u0409\3\2\2\2\u0409\u00d5\3\2\2\2\u040a"+
		"\u0408\3\2\2\2\u040b\u040c\5\u00fc\177\2\u040c\u040d\7,\2\2\u040d\u00d7"+
		"\3\2\2\2\u040e\u040f\bm\1\2\u040f\u0415\5\u00f2z\2\u0410\u0415\5\u00f4"+
		"{\2\u0411\u0415\5\u00e8u\2\u0412\u0415\5\u00dco\2\u0413\u0415\5\u00e2"+
		"r\2\u0414\u040e\3\2\2\2\u0414\u0410\3\2\2\2\u0414\u0411\3\2\2\2\u0414"+
		"\u0412\3\2\2\2\u0414\u0413\3\2\2\2\u0415\u041a\3\2\2\2\u0416\u0417\f\3"+
		"\2\2\u0417\u0419\7\62\2\2\u0418\u0416\3\2\2\2\u0419\u041c\3\2\2\2\u041a"+
		"\u0418\3\2\2\2\u041a\u041b\3\2\2\2\u041b\u00d9\3\2\2\2\u041c\u041a\3\2"+
		"\2\2\u041d\u041e\7,\2\2\u041e\u0420\5\u00d8m\2\u041f\u0421\5\u00a2R\2"+
		"\u0420\u041f\3\2\2\2\u0420\u0421\3\2\2\2\u0421\u00db\3\2\2\2\u0422\u0423"+
		"\5:\36\2\u0423\u0424\7$\2\2\u0424\u0426\3\2\2\2\u0425\u0422\3\2\2\2\u0425"+
		"\u0426\3\2\2\2\u0426\u0427\3\2\2\2\u0427\u042a\5\u00dep\2\u0428\u0429"+
		"\7$\2\2\u0429\u042b\5\u00dep\2\u042a\u0428\3\2\2\2\u042a\u042b\3\2\2\2"+
		"\u042b\u00dd\3\2\2\2\u042c\u042e\5\u00e0q\2\u042d\u042f\5.\30\2\u042e"+
		"\u042d\3\2\2\2\u042e\u042f\3\2\2\2\u042f\u00df\3\2\2\2\u0430\u0431\7\""+
		"\2\2\u0431\u00e1\3\2\2\2\u0432\u0434\7&\2\2\u0433\u0435\5\u00e4s\2\u0434"+
		"\u0433\3\2\2\2\u0434\u0435\3\2\2\2\u0435\u0436\3\2\2\2\u0436\u0437\7)"+
		"\2\2\u0437\u00e3\3\2\2\2\u0438\u043d\5\u00e6t\2\u0439\u043a\7+\2\2\u043a"+
		"\u043c\5\u00e6t\2\u043b\u0439\3\2\2\2\u043c\u043f\3\2\2\2\u043d\u043b"+
		"\3\2\2\2\u043d\u043e\3\2\2\2\u043e\u00e5\3\2\2\2\u043f\u043d\3\2\2\2\u0440"+
		"\u0442\5\u00d8m\2\u0441\u0443\5\u00a2R\2\u0442\u0441\3\2\2\2\u0442\u0443"+
		"\3\2\2\2\u0443\u00e7\3\2\2\2\u0444\u0446\5\u00a2R\2\u0445\u0444\3\2\2"+
		"\2\u0445\u0446\3\2\2\2\u0446\u0447\3\2\2\2\u0447\u0448\5\u00eav\2\u0448"+
		"\u0449\5\u0114\u008b\2\u0449\u044a\5\u00d8m\2\u044a\u00e9\3\2\2\2\u044b"+
		"\u044c\7&\2\2\u044c\u0455\7)\2\2\u044d\u044e\7&\2\2\u044e\u0450\5\u00ec"+
		"w\2\u044f\u0451\5\u0116\u008c\2\u0450\u044f\3\2\2\2\u0450\u0451\3\2\2"+
		"\2\u0451\u0452\3\2\2\2\u0452\u0453\7)\2\2\u0453\u0455\3\2\2\2\u0454\u044b"+
		"\3\2\2\2\u0454\u044d\3\2\2\2\u0455\u00eb\3\2\2\2\u0456\u045b\5\u00eex"+
		"\2\u0457\u0458\7+\2\2\u0458\u045a\5\u00eex\2\u0459\u0457\3\2\2\2\u045a"+
		"\u045d\3\2\2\2\u045b\u0459\3\2\2\2\u045b\u045c\3\2\2\2\u045c\u00ed\3\2"+
		"\2\2\u045d\u045b\3\2\2\2\u045e\u0460\5\u00d8m\2\u045f\u0461\7>\2\2\u0460"+
		"\u045f\3\2\2\2\u0460\u0461\3\2\2\2\u0461\u0463\3\2\2\2\u0462\u0464\5\u00a2"+
		"R\2\u0463\u0462\3\2\2\2\u0463\u0464\3\2\2\2\u0464\u0469\3\2\2\2\u0465"+
		"\u0466\5\u00f0y\2\u0466\u0467\5\u00dan\2\u0467\u0469\3\2\2\2\u0468\u045e"+
		"\3\2\2\2\u0468\u0465\3\2\2\2\u0469\u00ef\3\2\2\2\u046a\u046b\5\u00fc\177"+
		"\2\u046b\u00f1\3\2\2\2\u046c\u046d\7\'\2\2\u046d\u046f\5\u00d8m\2\u046e"+
		"\u0470\5\u00a2R\2\u046f\u046e\3\2\2\2\u046f\u0470\3\2\2\2\u0470\u0471"+
		"\3\2\2\2\u0471\u0472\7*\2\2\u0472\u00f3\3\2\2\2\u0473\u0474\7%\2\2\u0474"+
		"\u0475\5\u00d8m\2\u0475\u0476\7,\2\2\u0476\u0477\5\u00d8m\2\u0477\u0478"+
		"\5\u00a2R\2\u0478\u0479\7(\2\2\u0479\u00f5\3\2\2\2\u047a\u047b\7,\2\2"+
		"\u047b\u047c\5\u00f8}\2\u047c\u00f7\3\2\2\2\u047d\u0482\5\u00dco\2\u047e"+
		"\u047f\7+\2\2\u047f\u0481\5\u00dco\2\u0480\u047e\3\2\2\2\u0481\u0484\3"+
		"\2\2\2\u0482\u0480\3\2\2\2\u0482\u0483\3\2\2\2\u0483\u00f9\3\2\2\2\u0484"+
		"\u0482\3\2\2\2\u0485\u0488\7#\2\2\u0486\u0488\5\u0102\u0082\2\u0487\u0485"+
		"\3\2\2\2\u0487\u0486\3\2\2\2\u0488\u00fb\3\2\2\2\u0489\u048c\7#\2\2\u048a"+
		"\u048c\5\u0104\u0083\2\u048b\u0489\3\2\2\2\u048b\u048a\3\2\2\2\u048c\u00fd"+
		"\3\2\2\2\u048d\u0492\5\u00fa~\2\u048e\u048f\7$\2\2\u048f\u0491\5\u00fa"+
		"~\2\u0490\u048e\3\2\2\2\u0491\u0494\3\2\2\2\u0492\u0490\3\2\2\2\u0492"+
		"\u0493\3\2\2\2\u0493\u00ff\3\2\2\2\u0494\u0492\3\2\2\2\u0495\u0496\t\3"+
		"\2\2\u0496\u0101\3\2\2\2\u0497\u0498\t\4\2\2\u0498\u0103\3\2\2\2\u0499"+
		"\u049a\t\4\2\2\u049a\u0105\3\2\2\2\u049b\u049d\7M\2\2\u049c\u049b\3\2"+
		"\2\2\u049d\u049e\3\2\2\2\u049e\u049c\3\2\2\2\u049e\u049f\3\2\2\2\u049f"+
		"\u0107\3\2\2\2\u04a0\u04a1\7N\2\2\u04a1\u0109\3\2\2\2\u04a2\u04a3\7\66"+
		"\2\2\u04a3\u010b\3\2\2\2\u04a4\u04a5\7\65\2\2\u04a5\u010d\3\2\2\2\u04a6"+
		"\u04a7\7\64\2\2\u04a7\u04a8\7\64\2\2\u04a8\u010f\3\2\2\2\u04a9\u04aa\7"+
		"\67\2\2\u04aa\u04ab\7\67\2\2\u04ab\u0111\3\2\2\2\u04ac\u04ad\7/\2\2\u04ad"+
		"\u04ae\7\66\2\2\u04ae\u0113\3\2\2\2\u04af\u04b0\7\65\2\2\u04b0\u04b1\7"+
		"/\2\2\u04b1\u0115\3\2\2\2\u04b2\u04b3\7$\2\2\u04b3\u04b4\7$\2\2\u04b4"+
		"\u0117\3\2\2\2\u04b5\u04b6\7\66\2\2\u04b6\u04b7\7\66\2\2\u04b7\u0119\3"+
		"\2\2\2\u04b8\u04b9\5\u0120\u0091\2\u04b9\u011b\3\2\2\2\u04ba\u04bb\5\u0120"+
		"\u0091\2\u04bb\u011d\3\2\2\2\u04bc\u04bd\5\u0120\u0091\2\u04bd\u011f\3"+
		"\2\2\2\u04be\u04c2\5\u0124\u0093\2\u04bf\u04c1\5\u0122\u0092\2\u04c0\u04bf"+
		"\3\2\2\2\u04c1\u04c4\3\2\2\2\u04c2\u04c0\3\2\2\2\u04c2\u04c3\3\2\2\2\u04c3"+
		"\u04cd\3\2\2\2\u04c4\u04c2\3\2\2\2\u04c5\u04c9\5\u0126\u0094\2\u04c6\u04c8"+
		"\5\u0128\u0095\2\u04c7\u04c6\3\2\2\2\u04c8\u04cb\3\2\2\2\u04c9\u04c7\3"+
		"\2\2\2\u04c9\u04ca\3\2\2\2\u04ca\u04cd\3\2\2\2\u04cb\u04c9\3\2\2\2\u04cc"+
		"\u04be\3\2\2\2\u04cc\u04c5\3\2\2\2\u04cd\u0121\3\2\2\2\u04ce\u04d1\5\u0124"+
		"\u0093\2\u04cf\u04d1\7@\2\2\u04d0\u04ce\3\2\2\2\u04d0\u04cf\3\2\2\2\u04d1"+
		"\u0123\3\2\2\2\u04d2\u04d5\t\5\2\2\u04d3\u04d5\7?\2\2\u04d4\u04d2\3\2"+
		"\2\2\u04d4\u04d3\3\2\2\2\u04d5\u0125\3\2\2\2\u04d6\u04d7\7$\2\2\u04d7"+
		"\u0127\3\2\2\2\u04d8\u04db\7$\2\2\u04d9\u04db\5\u0122\u0092\2\u04da\u04d8"+
		"\3\2\2\2\u04da\u04d9\3\2\2\2\u04db\u0129\3\2\2\2\u04dc\u04e1\5\u012c\u0097"+
		"\2\u04dd\u04e1\5\u0134\u009b\2\u04de\u04e1\5\u012e\u0098\2\u04df\u04e1"+
		"\5\u0130\u0099\2\u04e0\u04dc\3\2\2\2\u04e0\u04dd\3\2\2\2\u04e0\u04de\3"+
		"\2\2\2\u04e0\u04df\3\2\2\2\u04e1\u012b\3\2\2\2\u04e2\u04e4\5\u010c\u0087"+
		"\2\u04e3\u04e2\3\2\2\2\u04e3\u04e4\3\2\2\2\u04e4\u04e5\3\2\2\2\u04e5\u04eb"+
		"\5\u0132\u009a\2\u04e6\u04e8\5\u010c\u0087\2\u04e7\u04e6\3\2\2\2\u04e7"+
		"\u04e8\3\2\2\2\u04e8\u04e9\3\2\2\2\u04e9\u04eb\7G\2\2\u04ea\u04e3\3\2"+
		"\2\2\u04ea\u04e7\3\2\2\2\u04eb\u012d\3\2\2\2\u04ec\u04ed\t\6\2\2\u04ed"+
		"\u012f\3\2\2\2\u04ee\u04ef\7\35\2\2\u04ef\u0131\3\2\2\2\u04f0\u04f1\t"+
		"\7\2\2\u04f1\u0133\3\2\2\2\u04f2\u04f3\t\b\2\2\u04f3\u0135\3\2\2\2\u0092"+
		"\u0139\u013f\u0143\u0145\u014b\u015c\u0161\u016b\u0174\u017a\u0180\u0186"+
		"\u018c\u0192\u0197\u019b\u019f\u01a3\u01ae\u01b6\u01c1\u01c6\u01d0\u01d6"+
		"\u01e3\u01eb\u01f2\u01f7\u0202\u020b\u0210\u0214\u0219\u021d\u0220\u022a"+
		"\u022f\u0235\u023a\u0246\u024a\u024d\u0253\u0257\u025c\u0266\u026d\u0273"+
		"\u0279\u027b\u027e\u0281\u0286\u0289\u028d\u0296\u0299\u029d\u02a0\u02a3"+
		"\u02a6\u02a9\u02ae\u02b1\u02b4\u02bc\u02c6\u02c9\u02ce\u02d1\u02d4\u02d7"+
		"\u02dc\u02e6\u02ed\u02f0\u02f3\u02f7\u02fe\u0302\u0306\u030c\u0313\u031c"+
		"\u0325\u0332\u033d\u0341\u0348\u034f\u035b\u0360\u036a\u036e\u0375\u037a"+
		"\u037e\u0384\u038a\u0394\u0398\u039e\u03b3\u03b7\u03be\u03d0\u03e3\u03e5"+
		"\u03ee\u03f5\u0402\u0408\u0414\u041a\u0420\u0425\u042a\u042e\u0434\u043d"+
		"\u0442\u0445\u0450\u0454\u045b\u0460\u0463\u0468\u046f\u0482\u0487\u048b"+
		"\u0492\u049e\u04c2\u04c9\u04cc\u04d0\u04d4\u04da\u04e0\u04e3\u04e7\u04ea";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}