// Generated from MojoParser.g4 by ANTLR 4.13.2
package org.mojolang.mojo.parser.syntax;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class MojoParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		KEYWORD_AND=1, KEYWORD_AS=2, KEYWORD_ATTRIBUTE=3, KEYWORD_BREAK=4, KEYWORD_CONST=5, 
		KEYWORD_CONTINUE=6, KEYWORD_ELSE=7, KEYWORD_ENUM=8, KEYWORD_FALSE=9, KEYWORD_FOR=10, 
		KEYWORD_FUNC=11, KEYWORD_IF=12, KEYWORD_IMPORT=13, KEYWORD_IN=14, KEYWORD_INTERFACE=15, 
		KEYWORD_IS=16, KEYWORD_MATCH=17, KEYWORD_NOT=18, KEYWORD_NULL=19, KEYWORD_OR=20, 
		KEYWORD_PACKAGE=21, KEYWORD_RETURN=22, KEYWORD_STRUCT=23, KEYWORD_TRUE=24, 
		KEYWORD_TYPE=25, KEYWORD_VAR=26, KEYWORD_WHILE=27, KEYWORD_XOR=28, DOT=29, 
		LCURLY=30, LPAREN=31, LBRACK=32, RCURLY=33, RPAREN=34, RBRACK=35, COMMA=36, 
		COLON=37, SEMI=38, LT=39, GT=40, BANG=41, QUESTION=42, AT=43, AND=44, 
		MINUS=45, EQUAL=46, PIPE=47, SLASH=48, PLUS=49, STAR=50, PERCENT=51, CARET=52, 
		TILDE=53, DOLLAR=54, BACKTICK=55, UNDERSCORE=56, PLUS_PLUS=57, MINUS_MINUS=58, 
		COLON_EQUAL=59, RIGHT_RIGHT_ARROWS=60, RIGHT_ARROW=61, DOT_DOT=62, DOT_DOT_LT=63, 
		DOT_DOT_EQUAL=64, ELLIPSIS=65, GRAPH_RIGHT_PATH=66, GRAPH_LEFT_PATH=67, 
		GRAPH_PATH=68, GRAPH_CONSTRAINT_PATH_LEFT=69, GRAPH_CONSTRAINT_PATH_LEFT_ARROW=70, 
		GRAPH_CONSTRAINT_PATH_RIGHT=71, GRAPH_CONSTRAINT_PATH_RIGHT_ARROW=72, 
		TYPE_IDENTIFIER=73, VALUE_IDENTIFIER=74, OPERATOR_HEAD_OTHER=75, IMPLICIT_PARAMETER_NAME=76, 
		BINARY_LITERAL=77, OCTAL_LITERAL=78, DECIMAL_LITERAL=79, PURE_DECIMAL_DIGITS=80, 
		HEXADECIMAL_LITERAL=81, FLOAT_LITERAL=82, STATIC_STRING_LITERAL=83, INTERPOLATED_STRING_LITERAL=84, 
		WS=85, BLOCK_COMMENT=86, LINE_COMMENT=87, LINE_COMMENT_DISTINCT_DOCUMENT=88, 
		EOL=89, LINE_DOCUMENT=90, FOLLOWING_LINE_DOCUMENT=91, INNER_LINE_DOCUMENT=92, 
		OPERATOR_FOLLOWING_CHARACTER=93;
	public static final int
		RULE_mojoFile = 0, RULE_statement = 1, RULE_ifModifier = 2, RULE_whileModifier = 3, 
		RULE_floatingStatement = 4, RULE_statements = 5, RULE_loopStatement = 6, 
		RULE_forInStatement = 7, RULE_whileStatement = 8, RULE_conditions = 9, 
		RULE_condition = 10, RULE_optionalBindingCondition = 11, RULE_branchStatement = 12, 
		RULE_ifStatement = 13, RULE_elseClause = 14, RULE_matchStatement = 15, 
		RULE_matchCases = 16, RULE_matchCase = 17, RULE_controlTransferStatement = 18, 
		RULE_breakStatement = 19, RULE_continueStatement = 20, RULE_returnStatement = 21, 
		RULE_genericParameterClause = 22, RULE_genericParameters = 23, RULE_genericParameter = 24, 
		RULE_genericArgumentClause = 25, RULE_genericArguments = 26, RULE_genericArgument = 27, 
		RULE_declaration = 28, RULE_codeBlock = 29, RULE_packageDeclaration = 30, 
		RULE_packageIdentifier = 31, RULE_packageName = 32, RULE_importDeclaration = 33, 
		RULE_importPath = 34, RULE_importPathIdentifier = 35, RULE_importAllClause = 36, 
		RULE_importValueAsClause = 37, RULE_importTypeClause = 38, RULE_importTypeAsClause = 39, 
		RULE_importGroupClause = 40, RULE_importGroup = 41, RULE_importValue = 42, 
		RULE_importType = 43, RULE_constantDeclaration = 44, RULE_patternInitializers = 45, 
		RULE_documentedPatternInitializer = 46, RULE_patternInitializer = 47, 
		RULE_initializer = 48, RULE_variableDeclaration = 49, RULE_typeAliasDeclaration = 50, 
		RULE_typeAliasName = 51, RULE_typeAliasAssignment = 52, RULE_functionDeclaration = 53, 
		RULE_functionName = 54, RULE_functionSignature = 55, RULE_functionResult = 56, 
		RULE_functionBody = 57, RULE_functionParameterClause = 58, RULE_functionParameters = 59, 
		RULE_functionParameter = 60, RULE_enumDeclaration = 61, RULE_enumBody = 62, 
		RULE_enumName = 63, RULE_enumMembers = 64, RULE_enumMember = 65, RULE_structDeclaration = 66, 
		RULE_structName = 67, RULE_structType = 68, RULE_structBody = 69, RULE_structMembers = 70, 
		RULE_structMember = 71, RULE_structMemberDeclaration = 72, RULE_interfaceDeclaration = 73, 
		RULE_interfaceName = 74, RULE_interfaceType = 75, RULE_interfaceBody = 76, 
		RULE_interfaceMembers = 77, RULE_interfaceMember = 78, RULE_interfaceMethodDeclaration = 79, 
		RULE_attributeDeclaration = 80, RULE_attributeAliasDeclaration = 81, RULE_attributeAliasAssignment = 82, 
		RULE_pattern = 83, RULE_wildcardPattern = 84, RULE_identifierPattern = 85, 
		RULE_tuplePattern = 86, RULE_tuplePatternElementList = 87, RULE_tuplePatternElement = 88, 
		RULE_arrayPattern = 89, RULE_arrayPatternElements = 90, RULE_arrayPatternElement = 91, 
		RULE_enumValuePattern = 92, RULE_optionalPattern = 93, RULE_expressionPattern = 94, 
		RULE_attribute = 95, RULE_attributeIdentifier = 96, RULE_attributeName = 97, 
		RULE_attributeArgumentClause = 98, RULE_attributeArgument = 99, RULE_attributeArguments = 100, 
		RULE_attributes = 101, RULE_expression = 102, RULE_prefixExpression = 103, 
		RULE_binaryExpression = 104, RULE_prefixCallOperator = 105, RULE_infixCallOperator = 106, 
		RULE_binaryExpressions = 107, RULE_inOperator = 108, RULE_conditionalOperator = 109, 
		RULE_ifOperator = 110, RULE_typeCastingOperator = 111, RULE_primaryExpression = 112, 
		RULE_literalExpression = 113, RULE_numericOperatorLiteral = 114, RULE_stringOperatorLiteral = 115, 
		RULE_suffixLiteralOperator = 116, RULE_prefixLiteralOperator = 117, RULE_arrayLiteral = 118, 
		RULE_arrayLiteralItems = 119, RULE_arrayLiteralItem = 120, RULE_mapLiteral = 121, 
		RULE_mapLiteralItems = 122, RULE_mapLiteralItem = 123, RULE_objectLiteral = 124, 
		RULE_objectLiteralItems = 125, RULE_objectLiteralItem = 126, RULE_structLiteral = 127, 
		RULE_structConstructionExpression = 128, RULE_matchExprSuffix = 129, RULE_matchExprCases = 130, 
		RULE_matchExprCase = 131, RULE_closureExpression = 132, RULE_closureParameters = 133, 
		RULE_closureParameter = 134, RULE_implicitMemberExpression = 135, RULE_parenthesizedExpression = 136, 
		RULE_tupleLiteralExpression = 137, RULE_tupleElement = 138, RULE_wildcardExpression = 139, 
		RULE_postfixExpression = 140, RULE_suffixExpression = 141, RULE_explicitMemberSuffix = 142, 
		RULE_subscriptSuffix = 143, RULE_functionCallSuffix = 144, RULE_functionCallArgumentClause = 145, 
		RULE_functionCallArguments = 146, RULE_functionCallArgument = 147, RULE_trailingClosures = 148, 
		RULE_labeledTrailingClosures = 149, RULE_labeledTrailingClosure = 150, 
		RULE_argumentNames = 151, RULE_argumentName = 152, RULE_type_ = 153, RULE_basicType = 154, 
		RULE_primeType = 155, RULE_typeAnnotation = 156, RULE_typeIdentifier = 157, 
		RULE_typeIdentifierClause = 158, RULE_typeName = 159, RULE_tupleType = 160, 
		RULE_tupleTypeElements = 161, RULE_tupleTypeElement = 162, RULE_functionType = 163, 
		RULE_arrayType = 164, RULE_mapType = 165, RULE_keyAttributes = 166, RULE_typeInheritanceClause = 167, 
		RULE_typeInheritances = 168, RULE_typeInheritance = 169, RULE_declarationIdentifier = 170, 
		RULE_labelIdentifier = 171, RULE_pathIdentifier = 172, RULE_identifier = 173, 
		RULE_keywordAsIdentifierInDeclarations = 174, RULE_keywordAsIdentifierInLabels = 175, 
		RULE_document = 176, RULE_followingDocument = 177, RULE_assignmentOperator = 178, 
		RULE_negatePrefixOperator = 179, RULE_arrowOperator = 180, RULE_rangeOperator = 181, 
		RULE_halfOpenRangeOperator = 182, RULE_closeRangeOperator = 183, RULE_binaryOperator = 184, 
		RULE_prefixOperator = 185, RULE_postfixOperator = 186, RULE_operator = 187, 
		RULE_operator_characters = 188, RULE_operator_character = 189, RULE_operator_head = 190, 
		RULE_dot_operator_head = 191, RULE_dot_operator_character = 192, RULE_literal = 193, 
		RULE_boolLiteral = 194, RULE_nullLiteral = 195, RULE_numericLiteral = 196, 
		RULE_integerLiteral = 197, RULE_stringLiteral = 198, RULE_eos = 199, RULE_eov = 200, 
		RULE_eosWithDocument = 201, RULE_eovWithDocument = 202;
	private static String[] makeRuleNames() {
		return new String[] {
			"mojoFile", "statement", "ifModifier", "whileModifier", "floatingStatement", 
			"statements", "loopStatement", "forInStatement", "whileStatement", "conditions", 
			"condition", "optionalBindingCondition", "branchStatement", "ifStatement", 
			"elseClause", "matchStatement", "matchCases", "matchCase", "controlTransferStatement", 
			"breakStatement", "continueStatement", "returnStatement", "genericParameterClause", 
			"genericParameters", "genericParameter", "genericArgumentClause", "genericArguments", 
			"genericArgument", "declaration", "codeBlock", "packageDeclaration", 
			"packageIdentifier", "packageName", "importDeclaration", "importPath", 
			"importPathIdentifier", "importAllClause", "importValueAsClause", "importTypeClause", 
			"importTypeAsClause", "importGroupClause", "importGroup", "importValue", 
			"importType", "constantDeclaration", "patternInitializers", "documentedPatternInitializer", 
			"patternInitializer", "initializer", "variableDeclaration", "typeAliasDeclaration", 
			"typeAliasName", "typeAliasAssignment", "functionDeclaration", "functionName", 
			"functionSignature", "functionResult", "functionBody", "functionParameterClause", 
			"functionParameters", "functionParameter", "enumDeclaration", "enumBody", 
			"enumName", "enumMembers", "enumMember", "structDeclaration", "structName", 
			"structType", "structBody", "structMembers", "structMember", "structMemberDeclaration", 
			"interfaceDeclaration", "interfaceName", "interfaceType", "interfaceBody", 
			"interfaceMembers", "interfaceMember", "interfaceMethodDeclaration", 
			"attributeDeclaration", "attributeAliasDeclaration", "attributeAliasAssignment", 
			"pattern", "wildcardPattern", "identifierPattern", "tuplePattern", "tuplePatternElementList", 
			"tuplePatternElement", "arrayPattern", "arrayPatternElements", "arrayPatternElement", 
			"enumValuePattern", "optionalPattern", "expressionPattern", "attribute", 
			"attributeIdentifier", "attributeName", "attributeArgumentClause", "attributeArgument", 
			"attributeArguments", "attributes", "expression", "prefixExpression", 
			"binaryExpression", "prefixCallOperator", "infixCallOperator", "binaryExpressions", 
			"inOperator", "conditionalOperator", "ifOperator", "typeCastingOperator", 
			"primaryExpression", "literalExpression", "numericOperatorLiteral", "stringOperatorLiteral", 
			"suffixLiteralOperator", "prefixLiteralOperator", "arrayLiteral", "arrayLiteralItems", 
			"arrayLiteralItem", "mapLiteral", "mapLiteralItems", "mapLiteralItem", 
			"objectLiteral", "objectLiteralItems", "objectLiteralItem", "structLiteral", 
			"structConstructionExpression", "matchExprSuffix", "matchExprCases", 
			"matchExprCase", "closureExpression", "closureParameters", "closureParameter", 
			"implicitMemberExpression", "parenthesizedExpression", "tupleLiteralExpression", 
			"tupleElement", "wildcardExpression", "postfixExpression", "suffixExpression", 
			"explicitMemberSuffix", "subscriptSuffix", "functionCallSuffix", "functionCallArgumentClause", 
			"functionCallArguments", "functionCallArgument", "trailingClosures", 
			"labeledTrailingClosures", "labeledTrailingClosure", "argumentNames", 
			"argumentName", "type_", "basicType", "primeType", "typeAnnotation", 
			"typeIdentifier", "typeIdentifierClause", "typeName", "tupleType", "tupleTypeElements", 
			"tupleTypeElement", "functionType", "arrayType", "mapType", "keyAttributes", 
			"typeInheritanceClause", "typeInheritances", "typeInheritance", "declarationIdentifier", 
			"labelIdentifier", "pathIdentifier", "identifier", "keywordAsIdentifierInDeclarations", 
			"keywordAsIdentifierInLabels", "document", "followingDocument", "assignmentOperator", 
			"negatePrefixOperator", "arrowOperator", "rangeOperator", "halfOpenRangeOperator", 
			"closeRangeOperator", "binaryOperator", "prefixOperator", "postfixOperator", 
			"operator", "operator_characters", "operator_character", "operator_head", 
			"dot_operator_head", "dot_operator_character", "literal", "boolLiteral", 
			"nullLiteral", "numericLiteral", "integerLiteral", "stringLiteral", "eos", 
			"eov", "eosWithDocument", "eovWithDocument"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'and'", "'as'", "'attribute'", "'break'", "'const'", "'continue'", 
			"'else'", "'enum'", "'false'", "'for'", "'func'", "'if'", "'import'", 
			"'in'", "'interface'", "'is'", "'match'", "'not'", "'null'", "'or'", 
			"'package'", "'return'", "'struct'", "'true'", "'type'", "'var'", "'while'", 
			"'xor'", "'.'", "'{'", "'('", "'['", "'}'", "')'", "']'", "','", "':'", 
			"';'", "'<'", "'>'", "'!'", "'?'", "'@'", "'&'", "'-'", "'='", "'|'", 
			"'/'", "'+'", "'*'", "'%'", "'^'", "'~'", "'$'", "'`'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "KEYWORD_AND", "KEYWORD_AS", "KEYWORD_ATTRIBUTE", "KEYWORD_BREAK", 
			"KEYWORD_CONST", "KEYWORD_CONTINUE", "KEYWORD_ELSE", "KEYWORD_ENUM", 
			"KEYWORD_FALSE", "KEYWORD_FOR", "KEYWORD_FUNC", "KEYWORD_IF", "KEYWORD_IMPORT", 
			"KEYWORD_IN", "KEYWORD_INTERFACE", "KEYWORD_IS", "KEYWORD_MATCH", "KEYWORD_NOT", 
			"KEYWORD_NULL", "KEYWORD_OR", "KEYWORD_PACKAGE", "KEYWORD_RETURN", "KEYWORD_STRUCT", 
			"KEYWORD_TRUE", "KEYWORD_TYPE", "KEYWORD_VAR", "KEYWORD_WHILE", "KEYWORD_XOR", 
			"DOT", "LCURLY", "LPAREN", "LBRACK", "RCURLY", "RPAREN", "RBRACK", "COMMA", 
			"COLON", "SEMI", "LT", "GT", "BANG", "QUESTION", "AT", "AND", "MINUS", 
			"EQUAL", "PIPE", "SLASH", "PLUS", "STAR", "PERCENT", "CARET", "TILDE", 
			"DOLLAR", "BACKTICK", "UNDERSCORE", "PLUS_PLUS", "MINUS_MINUS", "COLON_EQUAL", 
			"RIGHT_RIGHT_ARROWS", "RIGHT_ARROW", "DOT_DOT", "DOT_DOT_LT", "DOT_DOT_EQUAL", 
			"ELLIPSIS", "GRAPH_RIGHT_PATH", "GRAPH_LEFT_PATH", "GRAPH_PATH", "GRAPH_CONSTRAINT_PATH_LEFT", 
			"GRAPH_CONSTRAINT_PATH_LEFT_ARROW", "GRAPH_CONSTRAINT_PATH_RIGHT", "GRAPH_CONSTRAINT_PATH_RIGHT_ARROW", 
			"TYPE_IDENTIFIER", "VALUE_IDENTIFIER", "OPERATOR_HEAD_OTHER", "IMPLICIT_PARAMETER_NAME", 
			"BINARY_LITERAL", "OCTAL_LITERAL", "DECIMAL_LITERAL", "PURE_DECIMAL_DIGITS", 
			"HEXADECIMAL_LITERAL", "FLOAT_LITERAL", "STATIC_STRING_LITERAL", "INTERPOLATED_STRING_LITERAL", 
			"WS", "BLOCK_COMMENT", "LINE_COMMENT", "LINE_COMMENT_DISTINCT_DOCUMENT", 
			"EOL", "LINE_DOCUMENT", "FOLLOWING_LINE_DOCUMENT", "INNER_LINE_DOCUMENT", 
			"OPERATOR_FOLLOWING_CHARACTER"
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
	public String getGrammarFileName() { return "MojoParser.g4"; }

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

	@SuppressWarnings("CheckReturnValue")
	public static class MojoFileContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(MojoParser.EOF, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public MojoFileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mojoFile; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMojoFile(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MojoFileContext mojoFile() throws RecognitionException {
		MojoFileContext _localctx = new MojoFileContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_mojoFile);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(409);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(406);
					match(EOL);
					}
					} 
				}
				setState(411);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(413);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 90071451381530622L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 34600705L) != 0)) {
				{
				setState(412);
				statements();
				}
			}

			setState(418);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(415);
				match(EOL);
				}
				}
				setState(420);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(421);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public DeclarationContext declaration() {
			return getRuleContext(DeclarationContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public IfModifierContext ifModifier() {
			return getRuleContext(IfModifierContext.class,0);
		}
		public WhileModifierContext whileModifier() {
			return getRuleContext(WhileModifierContext.class,0);
		}
		public LoopStatementContext loopStatement() {
			return getRuleContext(LoopStatementContext.class,0);
		}
		public BranchStatementContext branchStatement() {
			return getRuleContext(BranchStatementContext.class,0);
		}
		public ControlTransferStatementContext controlTransferStatement() {
			return getRuleContext(ControlTransferStatementContext.class,0);
		}
		public FloatingStatementContext floatingStatement() {
			return getRuleContext(FloatingStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(433);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(423);
				declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(424);
				expression();
				setState(427);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case KEYWORD_IF:
					{
					setState(425);
					ifModifier();
					}
					break;
				case KEYWORD_WHILE:
					{
					setState(426);
					whileModifier();
					}
					break;
				case EOF:
				case RCURLY:
				case SEMI:
				case EOL:
					break;
				default:
					break;
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(429);
				loopStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(430);
				branchStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(431);
				controlTransferStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(432);
				floatingStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfModifierContext extends ParserRuleContext {
		public TerminalNode KEYWORD_IF() { return getToken(MojoParser.KEYWORD_IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public IfModifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifModifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIfModifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IfModifierContext ifModifier() throws RecognitionException {
		IfModifierContext _localctx = new IfModifierContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ifModifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(435);
			match(KEYWORD_IF);
			setState(436);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WhileModifierContext extends ParserRuleContext {
		public TerminalNode KEYWORD_WHILE() { return getToken(MojoParser.KEYWORD_WHILE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public WhileModifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileModifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitWhileModifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final WhileModifierContext whileModifier() throws RecognitionException {
		WhileModifierContext _localctx = new WhileModifierContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_whileModifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			match(KEYWORD_WHILE);
			setState(439);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FloatingStatementContext extends ParserRuleContext {
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public FloatingStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatingStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFloatingStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FloatingStatementContext floatingStatement() throws RecognitionException {
		FloatingStatementContext _localctx = new FloatingStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_floatingStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(441);
			document();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public TerminalNode SEMI() { return getToken(MojoParser.SEMI, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStatements(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatementsContext statements() throws RecognitionException {
		StatementsContext _localctx = new StatementsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_statements);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			statement();
			setState(455);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(444);
					eos();
					setState(448);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(445);
						match(EOL);
						}
						}
						setState(450);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(451);
					statement();
					}
					} 
				}
				setState(457);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			setState(459);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(458);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LoopStatementContext extends ParserRuleContext {
		public ForInStatementContext forInStatement() {
			return getRuleContext(ForInStatementContext.class,0);
		}
		public WhileStatementContext whileStatement() {
			return getRuleContext(WhileStatementContext.class,0);
		}
		public LoopStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loopStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitLoopStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LoopStatementContext loopStatement() throws RecognitionException {
		LoopStatementContext _localctx = new LoopStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_loopStatement);
		try {
			setState(463);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_FOR:
				enterOuterAlt(_localctx, 1);
				{
				setState(461);
				forInStatement();
				}
				break;
			case KEYWORD_WHILE:
				enterOuterAlt(_localctx, 2);
				{
				setState(462);
				whileStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ForInStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_FOR() { return getToken(MojoParser.KEYWORD_FOR, 0); }
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode KEYWORD_IN() { return getToken(MojoParser.KEYWORD_IN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public CodeBlockContext codeBlock() {
			return getRuleContext(CodeBlockContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ForInStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forInStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitForInStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ForInStatementContext forInStatement() throws RecognitionException {
		ForInStatementContext _localctx = new ForInStatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_forInStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(465);
			match(KEYWORD_FOR);
			setState(466);
			pattern(0);
			setState(470);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(467);
				match(EOL);
				}
				}
				setState(472);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(473);
			match(KEYWORD_IN);
			setState(474);
			expression();
			setState(478);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(475);
				match(EOL);
				}
				}
				setState(480);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(481);
			codeBlock();
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

	@SuppressWarnings("CheckReturnValue")
	public static class WhileStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_WHILE() { return getToken(MojoParser.KEYWORD_WHILE, 0); }
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public CodeBlockContext codeBlock() {
			return getRuleContext(CodeBlockContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public WhileStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitWhileStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final WhileStatementContext whileStatement() throws RecognitionException {
		WhileStatementContext _localctx = new WhileStatementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_whileStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(483);
			match(KEYWORD_WHILE);
			setState(484);
			conditions();
			setState(488);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(485);
				match(EOL);
				}
				}
				setState(490);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(491);
			codeBlock();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionsContext extends ParserRuleContext {
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ConditionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditions; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitConditions(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConditionsContext conditions() throws RecognitionException {
		ConditionsContext _localctx = new ConditionsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_conditions);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(493);
			condition();
			setState(505);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(494);
					eov();
					setState(498);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(495);
						match(EOL);
						}
						}
						setState(500);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(501);
					condition();
					}
					} 
				}
				setState(507);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public OptionalBindingConditionContext optionalBindingCondition() {
			return getRuleContext(OptionalBindingConditionContext.class,0);
		}
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitCondition(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_condition);
		try {
			setState(510);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_CONTINUE:
			case KEYWORD_ELSE:
			case KEYWORD_ENUM:
			case KEYWORD_FALSE:
			case KEYWORD_FUNC:
			case KEYWORD_IMPORT:
			case KEYWORD_IN:
			case KEYWORD_INTERFACE:
			case KEYWORD_IS:
			case KEYWORD_MATCH:
			case KEYWORD_NOT:
			case KEYWORD_NULL:
			case KEYWORD_OR:
			case KEYWORD_PACKAGE:
			case KEYWORD_STRUCT:
			case KEYWORD_TRUE:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
			case DOT:
			case LCURLY:
			case LPAREN:
			case LBRACK:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case UNDERSCORE:
			case ELLIPSIS:
			case TYPE_IDENTIFIER:
			case VALUE_IDENTIFIER:
			case OPERATOR_HEAD_OTHER:
			case BINARY_LITERAL:
			case OCTAL_LITERAL:
			case DECIMAL_LITERAL:
			case PURE_DECIMAL_DIGITS:
			case HEXADECIMAL_LITERAL:
			case FLOAT_LITERAL:
			case STATIC_STRING_LITERAL:
			case INTERPOLATED_STRING_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(508);
				expression();
				}
				break;
			case KEYWORD_VAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(509);
				optionalBindingCondition();
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

	@SuppressWarnings("CheckReturnValue")
	public static class OptionalBindingConditionContext extends ParserRuleContext {
		public TerminalNode KEYWORD_VAR() { return getToken(MojoParser.KEYWORD_VAR, 0); }
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public OptionalBindingConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optionalBindingCondition; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitOptionalBindingCondition(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OptionalBindingConditionContext optionalBindingCondition() throws RecognitionException {
		OptionalBindingConditionContext _localctx = new OptionalBindingConditionContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_optionalBindingCondition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(512);
			match(KEYWORD_VAR);
			setState(513);
			pattern(0);
			setState(517);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(514);
				match(EOL);
				}
				}
				setState(519);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(520);
			initializer();
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

	@SuppressWarnings("CheckReturnValue")
	public static class BranchStatementContext extends ParserRuleContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public MatchStatementContext matchStatement() {
			return getRuleContext(MatchStatementContext.class,0);
		}
		public BranchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_branchStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitBranchStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BranchStatementContext branchStatement() throws RecognitionException {
		BranchStatementContext _localctx = new BranchStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_branchStatement);
		try {
			setState(524);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(522);
				ifStatement();
				}
				break;
			case KEYWORD_MATCH:
				enterOuterAlt(_localctx, 2);
				{
				setState(523);
				matchStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_IF() { return getToken(MojoParser.KEYWORD_IF, 0); }
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public CodeBlockContext codeBlock() {
			return getRuleContext(CodeBlockContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ElseClauseContext elseClause() {
			return getRuleContext(ElseClauseContext.class,0);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIfStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_ifStatement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			match(KEYWORD_IF);
			setState(527);
			conditions();
			setState(531);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(528);
				match(EOL);
				}
				}
				setState(533);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(534);
			codeBlock();
			setState(538);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(535);
					match(EOL);
					}
					} 
				}
				setState(540);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			}
			setState(542);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_ELSE) {
				{
				setState(541);
				elseClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ElseClauseContext extends ParserRuleContext {
		public TerminalNode KEYWORD_ELSE() { return getToken(MojoParser.KEYWORD_ELSE, 0); }
		public CodeBlockContext codeBlock() {
			return getRuleContext(CodeBlockContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public ElseClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitElseClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ElseClauseContext elseClause() throws RecognitionException {
		ElseClauseContext _localctx = new ElseClauseContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_elseClause);
		int _la;
		try {
			setState(560);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(544);
				match(KEYWORD_ELSE);
				setState(548);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(545);
					match(EOL);
					}
					}
					setState(550);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(551);
				codeBlock();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(552);
				match(KEYWORD_ELSE);
				setState(556);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(553);
					match(EOL);
					}
					}
					setState(558);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(559);
				ifStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatchStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_MATCH() { return getToken(MojoParser.KEYWORD_MATCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MatchCasesContext matchCases() {
			return getRuleContext(MatchCasesContext.class,0);
		}
		public MatchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMatchStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MatchStatementContext matchStatement() throws RecognitionException {
		MatchStatementContext _localctx = new MatchStatementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_matchStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(562);
			match(KEYWORD_MATCH);
			setState(563);
			expression();
			setState(567);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(564);
				match(EOL);
				}
				}
				setState(569);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(570);
			match(LCURLY);
			setState(578);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				setState(574);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(571);
					match(EOL);
					}
					}
					setState(576);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(577);
				matchCases();
				}
				break;
			}
			setState(583);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(580);
				match(EOL);
				}
				}
				setState(585);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(586);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatchCasesContext extends ParserRuleContext {
		public List<MatchCaseContext> matchCase() {
			return getRuleContexts(MatchCaseContext.class);
		}
		public MatchCaseContext matchCase(int i) {
			return getRuleContext(MatchCaseContext.class,i);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MatchCasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchCases; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMatchCases(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MatchCasesContext matchCases() throws RecognitionException {
		MatchCasesContext _localctx = new MatchCasesContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_matchCases);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(588);
			matchCase();
			setState(600);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(589);
					eos();
					setState(593);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(590);
						match(EOL);
						}
						}
						setState(595);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(596);
					matchCase();
					}
					} 
				}
				setState(602);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			}
			setState(604);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(603);
				eos();
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatchCaseContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode RIGHT_RIGHT_ARROWS() { return getToken(MojoParser.RIGHT_RIGHT_ARROWS, 0); }
		public CodeBlockContext codeBlock() {
			return getRuleContext(CodeBlockContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public IfModifierContext ifModifier() {
			return getRuleContext(IfModifierContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MatchCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchCase; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMatchCase(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MatchCaseContext matchCase() throws RecognitionException {
		MatchCaseContext _localctx = new MatchCaseContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_matchCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(606);
			pattern(0);
			setState(608);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_IF) {
				{
				setState(607);
				ifModifier();
				}
			}

			setState(613);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(610);
				match(EOL);
				}
				}
				setState(615);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(616);
			match(RIGHT_RIGHT_ARROWS);
			setState(620);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(617);
				match(EOL);
				}
				}
				setState(622);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(625);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(623);
				codeBlock();
				}
				break;
			case 2:
				{
				setState(624);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ControlTransferStatementContext extends ParserRuleContext {
		public BreakStatementContext breakStatement() {
			return getRuleContext(BreakStatementContext.class,0);
		}
		public ContinueStatementContext continueStatement() {
			return getRuleContext(ContinueStatementContext.class,0);
		}
		public ReturnStatementContext returnStatement() {
			return getRuleContext(ReturnStatementContext.class,0);
		}
		public ControlTransferStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_controlTransferStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitControlTransferStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ControlTransferStatementContext controlTransferStatement() throws RecognitionException {
		ControlTransferStatementContext _localctx = new ControlTransferStatementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_controlTransferStatement);
		try {
			setState(630);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_BREAK:
				enterOuterAlt(_localctx, 1);
				{
				setState(627);
				breakStatement();
				}
				break;
			case KEYWORD_CONTINUE:
				enterOuterAlt(_localctx, 2);
				{
				setState(628);
				continueStatement();
				}
				break;
			case KEYWORD_RETURN:
				enterOuterAlt(_localctx, 3);
				{
				setState(629);
				returnStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_BREAK() { return getToken(MojoParser.KEYWORD_BREAK, 0); }
		public BreakStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitBreakStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BreakStatementContext breakStatement() throws RecognitionException {
		BreakStatementContext _localctx = new BreakStatementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_breakStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(632);
			match(KEYWORD_BREAK);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_CONTINUE() { return getToken(MojoParser.KEYWORD_CONTINUE, 0); }
		public ContinueStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continueStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitContinueStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ContinueStatementContext continueStatement() throws RecognitionException {
		ContinueStatementContext _localctx = new ContinueStatementContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_continueStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(634);
			match(KEYWORD_CONTINUE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends ParserRuleContext {
		public TerminalNode KEYWORD_RETURN() { return getToken(MojoParser.KEYWORD_RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ReturnStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStatement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitReturnStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ReturnStatementContext returnStatement() throws RecognitionException {
		ReturnStatementContext _localctx = new ReturnStatementContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_returnStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(636);
			match(KEYWORD_RETURN);
			setState(638);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 90062655082982398L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1046273L) != 0)) {
				{
				setState(637);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GenericParameterClauseContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(MojoParser.LT, 0); }
		public GenericParametersContext genericParameters() {
			return getRuleContext(GenericParametersContext.class,0);
		}
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public GenericParameterClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericParameterClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitGenericParameterClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GenericParameterClauseContext genericParameterClause() throws RecognitionException {
		GenericParameterClauseContext _localctx = new GenericParameterClauseContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_genericParameterClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(640);
			match(LT);
			setState(644);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(641);
				match(EOL);
				}
				}
				setState(646);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(647);
			genericParameters();
			setState(651);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(648);
				match(EOL);
				}
				}
				setState(653);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(654);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GenericParametersContext extends ParserRuleContext {
		public List<GenericParameterContext> genericParameter() {
			return getRuleContexts(GenericParameterContext.class);
		}
		public GenericParameterContext genericParameter(int i) {
			return getRuleContext(GenericParameterContext.class,i);
		}
		public List<EovWithDocumentContext> eovWithDocument() {
			return getRuleContexts(EovWithDocumentContext.class);
		}
		public EovWithDocumentContext eovWithDocument(int i) {
			return getRuleContext(EovWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public GenericParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericParameters; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitGenericParameters(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GenericParametersContext genericParameters() throws RecognitionException {
		GenericParametersContext _localctx = new GenericParametersContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_genericParameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(656);
			genericParameter();
			setState(668);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(657);
					eovWithDocument();
					setState(661);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(658);
						match(EOL);
						}
						}
						setState(663);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(664);
					genericParameter();
					}
					} 
				}
				setState(670);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
			}
			setState(672);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				{
				setState(671);
				eovWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class GenericParameterContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public GenericParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericParameter; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitGenericParameter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GenericParameterContext genericParameter() throws RecognitionException {
		GenericParameterContext _localctx = new GenericParameterContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_genericParameter);
		try {
			setState(681);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(674);
				typeName();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(675);
				typeName();
				setState(676);
				match(ELLIPSIS);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(678);
				typeName();
				setState(679);
				typeAnnotation();
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

	@SuppressWarnings("CheckReturnValue")
	public static class GenericArgumentClauseContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(MojoParser.LT, 0); }
		public GenericArgumentsContext genericArguments() {
			return getRuleContext(GenericArgumentsContext.class,0);
		}
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public GenericArgumentClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericArgumentClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitGenericArgumentClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GenericArgumentClauseContext genericArgumentClause() throws RecognitionException {
		GenericArgumentClauseContext _localctx = new GenericArgumentClauseContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_genericArgumentClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(683);
			match(LT);
			setState(687);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(684);
				match(EOL);
				}
				}
				setState(689);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(690);
			genericArguments();
			setState(694);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(691);
				match(EOL);
				}
				}
				setState(696);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(697);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GenericArgumentsContext extends ParserRuleContext {
		public List<GenericArgumentContext> genericArgument() {
			return getRuleContexts(GenericArgumentContext.class);
		}
		public GenericArgumentContext genericArgument(int i) {
			return getRuleContext(GenericArgumentContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public GenericArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericArguments; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitGenericArguments(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GenericArgumentsContext genericArguments() throws RecognitionException {
		GenericArgumentsContext _localctx = new GenericArgumentsContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_genericArguments);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(699);
			genericArgument();
			setState(711);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(700);
					eov();
					setState(704);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(701);
						match(EOL);
						}
						}
						setState(706);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(707);
					genericArgument();
					}
					} 
				}
				setState(713);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GenericArgumentContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public GenericArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericArgument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitGenericArgument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GenericArgumentContext genericArgument() throws RecognitionException {
		GenericArgumentContext _localctx = new GenericArgumentContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_genericArgument);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(714);
			type_(0);
			setState(716);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(715);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationContext extends ParserRuleContext {
		public PackageDeclarationContext packageDeclaration() {
			return getRuleContext(PackageDeclarationContext.class,0);
		}
		public ImportDeclarationContext importDeclaration() {
			return getRuleContext(ImportDeclarationContext.class,0);
		}
		public ConstantDeclarationContext constantDeclaration() {
			return getRuleContext(ConstantDeclarationContext.class,0);
		}
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public TypeAliasDeclarationContext typeAliasDeclaration() {
			return getRuleContext(TypeAliasDeclarationContext.class,0);
		}
		public FunctionDeclarationContext functionDeclaration() {
			return getRuleContext(FunctionDeclarationContext.class,0);
		}
		public EnumDeclarationContext enumDeclaration() {
			return getRuleContext(EnumDeclarationContext.class,0);
		}
		public StructDeclarationContext structDeclaration() {
			return getRuleContext(StructDeclarationContext.class,0);
		}
		public InterfaceDeclarationContext interfaceDeclaration() {
			return getRuleContext(InterfaceDeclarationContext.class,0);
		}
		public AttributeDeclarationContext attributeDeclaration() {
			return getRuleContext(AttributeDeclarationContext.class,0);
		}
		public AttributeAliasDeclarationContext attributeAliasDeclaration() {
			return getRuleContext(AttributeAliasDeclarationContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(721);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LINE_DOCUMENT) {
				{
				setState(718);
				document();
				setState(719);
				match(EOL);
				}
			}

			setState(730);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(723);
				attributes();
				setState(727);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(724);
					match(EOL);
					}
					}
					setState(729);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(743);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,50,_ctx) ) {
			case 1:
				{
				setState(732);
				packageDeclaration();
				}
				break;
			case 2:
				{
				setState(733);
				importDeclaration();
				}
				break;
			case 3:
				{
				setState(734);
				constantDeclaration();
				}
				break;
			case 4:
				{
				setState(735);
				variableDeclaration();
				}
				break;
			case 5:
				{
				setState(736);
				typeAliasDeclaration();
				}
				break;
			case 6:
				{
				setState(737);
				functionDeclaration();
				}
				break;
			case 7:
				{
				setState(738);
				enumDeclaration();
				}
				break;
			case 8:
				{
				setState(739);
				structDeclaration();
				}
				break;
			case 9:
				{
				setState(740);
				interfaceDeclaration();
				}
				break;
			case 10:
				{
				setState(741);
				attributeDeclaration();
				}
				break;
			case 11:
				{
				setState(742);
				attributeAliasDeclaration();
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

	@SuppressWarnings("CheckReturnValue")
	public static class CodeBlockContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public CodeBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_codeBlock; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitCodeBlock(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CodeBlockContext codeBlock() throws RecognitionException {
		CodeBlockContext _localctx = new CodeBlockContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_codeBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(745);
			match(LCURLY);
			setState(753);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
			case 1:
				{
				setState(749);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(746);
					match(EOL);
					}
					}
					setState(751);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(752);
				statements();
				}
				break;
			}
			setState(758);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(755);
				match(EOL);
				}
				}
				setState(760);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(761);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PackageDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_PACKAGE() { return getToken(MojoParser.KEYWORD_PACKAGE, 0); }
		public PackageIdentifierContext packageIdentifier() {
			return getRuleContext(PackageIdentifierContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public PackageDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_packageDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPackageDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PackageDeclarationContext packageDeclaration() throws RecognitionException {
		PackageDeclarationContext _localctx = new PackageDeclarationContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_packageDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(763);
			match(KEYWORD_PACKAGE);
			setState(764);
			packageIdentifier();
			setState(772);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
			case 1:
				{
				setState(768);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(765);
					match(EOL);
					}
					}
					setState(770);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(771);
				objectLiteral();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PackageIdentifierContext extends ParserRuleContext {
		public List<PackageNameContext> packageName() {
			return getRuleContexts(PackageNameContext.class);
		}
		public PackageNameContext packageName(int i) {
			return getRuleContext(PackageNameContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public PackageIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_packageIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPackageIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PackageIdentifierContext packageIdentifier() throws RecognitionException {
		PackageIdentifierContext _localctx = new PackageIdentifierContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_packageIdentifier);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(774);
			packageName();
			setState(779);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,56,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(775);
					match(DOT);
					setState(776);
					packageName();
					}
					} 
				}
				setState(781);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,56,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PackageNameContext extends ParserRuleContext {
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public PackageNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_packageName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPackageName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PackageNameContext packageName() throws RecognitionException {
		PackageNameContext _localctx = new PackageNameContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_packageName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(782);
			match(VALUE_IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_IMPORT() { return getToken(MojoParser.KEYWORD_IMPORT, 0); }
		public ImportPathContext importPath() {
			return getRuleContext(ImportPathContext.class,0);
		}
		public ImportAllClauseContext importAllClause() {
			return getRuleContext(ImportAllClauseContext.class,0);
		}
		public ImportValueAsClauseContext importValueAsClause() {
			return getRuleContext(ImportValueAsClauseContext.class,0);
		}
		public ImportTypeClauseContext importTypeClause() {
			return getRuleContext(ImportTypeClauseContext.class,0);
		}
		public ImportGroupClauseContext importGroupClause() {
			return getRuleContext(ImportGroupClauseContext.class,0);
		}
		public ImportDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportDeclarationContext importDeclaration() throws RecognitionException {
		ImportDeclarationContext _localctx = new ImportDeclarationContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_importDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(784);
			match(KEYWORD_IMPORT);
			setState(785);
			importPath();
			setState(790);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,57,_ctx) ) {
			case 1:
				{
				setState(786);
				importAllClause();
				}
				break;
			case 2:
				{
				setState(787);
				importValueAsClause();
				}
				break;
			case 3:
				{
				setState(788);
				importTypeClause();
				}
				break;
			case 4:
				{
				setState(789);
				importGroupClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportPathContext extends ParserRuleContext {
		public List<ImportPathIdentifierContext> importPathIdentifier() {
			return getRuleContexts(ImportPathIdentifierContext.class);
		}
		public ImportPathIdentifierContext importPathIdentifier(int i) {
			return getRuleContext(ImportPathIdentifierContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public ImportPathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importPath; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportPath(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportPathContext importPath() throws RecognitionException {
		ImportPathContext _localctx = new ImportPathContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_importPath);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(792);
			importPathIdentifier();
			setState(797);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,58,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(793);
					match(DOT);
					setState(794);
					importPathIdentifier();
					}
					} 
				}
				setState(799);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,58,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportPathIdentifierContext extends ParserRuleContext {
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public ImportPathIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importPathIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportPathIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportPathIdentifierContext importPathIdentifier() throws RecognitionException {
		ImportPathIdentifierContext _localctx = new ImportPathIdentifierContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_importPathIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(800);
			declarationIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportAllClauseContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public TerminalNode STAR() { return getToken(MojoParser.STAR, 0); }
		public ImportAllClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importAllClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportAllClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportAllClauseContext importAllClause() throws RecognitionException {
		ImportAllClauseContext _localctx = new ImportAllClauseContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_importAllClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(802);
			match(DOT);
			setState(803);
			match(STAR);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportValueAsClauseContext extends ParserRuleContext {
		public TerminalNode KEYWORD_AS() { return getToken(MojoParser.KEYWORD_AS, 0); }
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public ImportValueAsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importValueAsClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportValueAsClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportValueAsClauseContext importValueAsClause() throws RecognitionException {
		ImportValueAsClauseContext _localctx = new ImportValueAsClauseContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_importValueAsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(805);
			match(KEYWORD_AS);
			setState(806);
			declarationIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportTypeClauseContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public ImportTypeAsClauseContext importTypeAsClause() {
			return getRuleContext(ImportTypeAsClauseContext.class,0);
		}
		public ImportTypeClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importTypeClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportTypeClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportTypeClauseContext importTypeClause() throws RecognitionException {
		ImportTypeClauseContext _localctx = new ImportTypeClauseContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_importTypeClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(808);
			match(DOT);
			setState(809);
			typeName();
			setState(811);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_AS) {
				{
				setState(810);
				importTypeAsClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportTypeAsClauseContext extends ParserRuleContext {
		public TerminalNode KEYWORD_AS() { return getToken(MojoParser.KEYWORD_AS, 0); }
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public ImportTypeAsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importTypeAsClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportTypeAsClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportTypeAsClauseContext importTypeAsClause() throws RecognitionException {
		ImportTypeAsClauseContext _localctx = new ImportTypeAsClauseContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_importTypeAsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(813);
			match(KEYWORD_AS);
			setState(814);
			typeName();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportGroupClauseContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public ImportGroupContext importGroup() {
			return getRuleContext(ImportGroupContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ImportGroupClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importGroupClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportGroupClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportGroupClauseContext importGroupClause() throws RecognitionException {
		ImportGroupClauseContext _localctx = new ImportGroupClauseContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_importGroupClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(816);
			match(DOT);
			setState(817);
			match(LCURLY);
			setState(821);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(818);
				match(EOL);
				}
				}
				setState(823);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(824);
			importGroup();
			setState(828);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(825);
				match(EOL);
				}
				}
				setState(830);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(831);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportGroupContext extends ParserRuleContext {
		public List<ImportValueContext> importValue() {
			return getRuleContexts(ImportValueContext.class);
		}
		public ImportValueContext importValue(int i) {
			return getRuleContext(ImportValueContext.class,i);
		}
		public List<ImportTypeContext> importType() {
			return getRuleContexts(ImportTypeContext.class);
		}
		public ImportTypeContext importType(int i) {
			return getRuleContext(ImportTypeContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ImportGroupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importGroup; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportGroup(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportGroupContext importGroup() throws RecognitionException {
		ImportGroupContext _localctx = new ImportGroupContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_importGroup);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(835);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_CONTINUE:
			case KEYWORD_ELSE:
			case KEYWORD_ENUM:
			case KEYWORD_FALSE:
			case KEYWORD_FUNC:
			case KEYWORD_IMPORT:
			case KEYWORD_IN:
			case KEYWORD_INTERFACE:
			case KEYWORD_IS:
			case KEYWORD_MATCH:
			case KEYWORD_NOT:
			case KEYWORD_NULL:
			case KEYWORD_OR:
			case KEYWORD_PACKAGE:
			case KEYWORD_STRUCT:
			case KEYWORD_TRUE:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
			case VALUE_IDENTIFIER:
				{
				setState(833);
				importValue();
				}
				break;
			case TYPE_IDENTIFIER:
				{
				setState(834);
				importType();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(850);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,65,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(837);
					eov();
					setState(841);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(838);
						match(EOL);
						}
						}
						setState(843);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(846);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case KEYWORD_AND:
					case KEYWORD_AS:
					case KEYWORD_ATTRIBUTE:
					case KEYWORD_BREAK:
					case KEYWORD_CONST:
					case KEYWORD_CONTINUE:
					case KEYWORD_ELSE:
					case KEYWORD_ENUM:
					case KEYWORD_FALSE:
					case KEYWORD_FUNC:
					case KEYWORD_IMPORT:
					case KEYWORD_IN:
					case KEYWORD_INTERFACE:
					case KEYWORD_IS:
					case KEYWORD_MATCH:
					case KEYWORD_NOT:
					case KEYWORD_NULL:
					case KEYWORD_OR:
					case KEYWORD_PACKAGE:
					case KEYWORD_STRUCT:
					case KEYWORD_TRUE:
					case KEYWORD_TYPE:
					case KEYWORD_XOR:
					case VALUE_IDENTIFIER:
						{
						setState(844);
						importValue();
						}
						break;
					case TYPE_IDENTIFIER:
						{
						setState(845);
						importType();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					} 
				}
				setState(852);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,65,_ctx);
			}
			setState(854);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,66,_ctx) ) {
			case 1:
				{
				setState(853);
				eov();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportValueContext extends ParserRuleContext {
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public ImportValueAsClauseContext importValueAsClause() {
			return getRuleContext(ImportValueAsClauseContext.class,0);
		}
		public ImportValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importValue; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportValue(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportValueContext importValue() throws RecognitionException {
		ImportValueContext _localctx = new ImportValueContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_importValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(856);
			declarationIdentifier();
			setState(858);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_AS) {
				{
				setState(857);
				importValueAsClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImportTypeContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public ImportTypeAsClauseContext importTypeAsClause() {
			return getRuleContext(ImportTypeAsClauseContext.class,0);
		}
		public ImportTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImportType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportTypeContext importType() throws RecognitionException {
		ImportTypeContext _localctx = new ImportTypeContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_importType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(860);
			typeName();
			setState(862);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_AS) {
				{
				setState(861);
				importTypeAsClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConstantDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_CONST() { return getToken(MojoParser.KEYWORD_CONST, 0); }
		public PatternInitializersContext patternInitializers() {
			return getRuleContext(PatternInitializersContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public List<DocumentedPatternInitializerContext> documentedPatternInitializer() {
			return getRuleContexts(DocumentedPatternInitializerContext.class);
		}
		public DocumentedPatternInitializerContext documentedPatternInitializer(int i) {
			return getRuleContext(DocumentedPatternInitializerContext.class,i);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public ConstantDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitConstantDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConstantDeclarationContext constantDeclaration() throws RecognitionException {
		ConstantDeclarationContext _localctx = new ConstantDeclarationContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_constantDeclaration);
		int _la;
		try {
			int _alt;
			setState(900);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,74,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(864);
				match(KEYWORD_CONST);
				setState(865);
				patternInitializers();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(866);
				match(KEYWORD_CONST);
				setState(867);
				match(LCURLY);
				setState(871);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(868);
					match(EOL);
					}
					}
					setState(873);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(874);
				documentedPatternInitializer();
				setState(886);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,71,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(875);
						eos();
						setState(879);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(876);
							match(EOL);
							}
							}
							setState(881);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(882);
						documentedPatternInitializer();
						}
						} 
					}
					setState(888);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,71,_ctx);
				}
				setState(890);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,72,_ctx) ) {
				case 1:
					{
					setState(889);
					eos();
					}
					break;
				}
				setState(895);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(892);
					match(EOL);
					}
					}
					setState(897);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(898);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PatternInitializersContext extends ParserRuleContext {
		public PatternInitializerContext patternInitializer() {
			return getRuleContext(PatternInitializerContext.class,0);
		}
		public PatternInitializersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_patternInitializers; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPatternInitializers(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PatternInitializersContext patternInitializers() throws RecognitionException {
		PatternInitializersContext _localctx = new PatternInitializersContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_patternInitializers);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(902);
			patternInitializer();
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

	@SuppressWarnings("CheckReturnValue")
	public static class DocumentedPatternInitializerContext extends ParserRuleContext {
		public PatternInitializerContext patternInitializer() {
			return getRuleContext(PatternInitializerContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public DocumentedPatternInitializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_documentedPatternInitializer; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitDocumentedPatternInitializer(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DocumentedPatternInitializerContext documentedPatternInitializer() throws RecognitionException {
		DocumentedPatternInitializerContext _localctx = new DocumentedPatternInitializerContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_documentedPatternInitializer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(907);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LINE_DOCUMENT) {
				{
				setState(904);
				document();
				setState(905);
				match(EOL);
				}
			}

			setState(912);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(909);
				attributes();
				setState(910);
				match(EOL);
				}
			}

			setState(914);
			patternInitializer();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PatternInitializerContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public PatternInitializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_patternInitializer; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPatternInitializer(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PatternInitializerContext patternInitializer() throws RecognitionException {
		PatternInitializerContext _localctx = new PatternInitializerContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_patternInitializer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(916);
			pattern(0);
			setState(918);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(917);
				initializer();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InitializerContext extends ParserRuleContext {
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public InitializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInitializer(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InitializerContext initializer() throws RecognitionException {
		InitializerContext _localctx = new InitializerContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_initializer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(920);
			assignmentOperator();
			setState(924);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(921);
				match(EOL);
				}
				}
				setState(926);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(927);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_VAR() { return getToken(MojoParser.KEYWORD_VAR, 0); }
		public PatternInitializersContext patternInitializers() {
			return getRuleContext(PatternInitializersContext.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode COLON_EQUAL() { return getToken(MojoParser.COLON_EQUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public List<DocumentedPatternInitializerContext> documentedPatternInitializer() {
			return getRuleContexts(DocumentedPatternInitializerContext.class);
		}
		public DocumentedPatternInitializerContext documentedPatternInitializer(int i) {
			return getRuleContext(DocumentedPatternInitializerContext.class,i);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public VariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitVariableDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final VariableDeclarationContext variableDeclaration() throws RecognitionException {
		VariableDeclarationContext _localctx = new VariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_variableDeclaration);
		int _la;
		try {
			int _alt;
			setState(969);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,84,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(929);
				match(KEYWORD_VAR);
				setState(930);
				patternInitializers();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(931);
				pattern(0);
				setState(932);
				match(COLON_EQUAL);
				setState(933);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(935);
				match(KEYWORD_VAR);
				setState(936);
				match(LCURLY);
				setState(940);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(937);
					match(EOL);
					}
					}
					setState(942);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(943);
				documentedPatternInitializer();
				setState(955);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,81,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(944);
						eos();
						setState(948);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(945);
							match(EOL);
							}
							}
							setState(950);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(951);
						documentedPatternInitializer();
						}
						} 
					}
					setState(957);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,81,_ctx);
				}
				setState(959);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,82,_ctx) ) {
				case 1:
					{
					setState(958);
					eos();
					}
					break;
				}
				setState(964);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(961);
					match(EOL);
					}
					}
					setState(966);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(967);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeAliasDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public TypeAliasNameContext typeAliasName() {
			return getRuleContext(TypeAliasNameContext.class,0);
		}
		public TypeAliasAssignmentContext typeAliasAssignment() {
			return getRuleContext(TypeAliasAssignmentContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public TypeAliasDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAliasDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeAliasDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeAliasDeclarationContext typeAliasDeclaration() throws RecognitionException {
		TypeAliasDeclarationContext _localctx = new TypeAliasDeclarationContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_typeAliasDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(971);
			match(KEYWORD_TYPE);
			setState(972);
			typeAliasName();
			setState(974);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(973);
				genericParameterClause();
				}
			}

			setState(979);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(976);
				match(EOL);
				}
				}
				setState(981);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(982);
			typeAliasAssignment();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeAliasNameContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public TypeAliasNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAliasName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeAliasName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeAliasNameContext typeAliasName() throws RecognitionException {
		TypeAliasNameContext _localctx = new TypeAliasNameContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_typeAliasName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(984);
			typeName();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeAliasAssignmentContext extends ParserRuleContext {
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public TypeAliasAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAliasAssignment; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeAliasAssignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeAliasAssignmentContext typeAliasAssignment() throws RecognitionException {
		TypeAliasAssignmentContext _localctx = new TypeAliasAssignmentContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_typeAliasAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(986);
			assignmentOperator();
			setState(990);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(987);
				match(EOL);
				}
				}
				setState(992);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(993);
			type_(0);
			setState(995);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(994);
				attributes();
				}
			}

			setState(998);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,89,_ctx) ) {
			case 1:
				{
				setState(997);
				followingDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_FUNC() { return getToken(MojoParser.KEYWORD_FUNC, 0); }
		public FunctionNameContext functionName() {
			return getRuleContext(FunctionNameContext.class,0);
		}
		public FunctionSignatureContext functionSignature() {
			return getRuleContext(FunctionSignatureContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionDeclarationContext functionDeclaration() throws RecognitionException {
		FunctionDeclarationContext _localctx = new FunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_functionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1000);
			match(KEYWORD_FUNC);
			setState(1001);
			functionName();
			setState(1003);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1002);
				genericParameterClause();
				}
			}

			setState(1005);
			functionSignature();
			setState(1013);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,92,_ctx) ) {
			case 1:
				{
				setState(1009);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1006);
					match(EOL);
					}
					}
					setState(1011);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1012);
				functionBody();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionNameContext extends ParserRuleContext {
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public FunctionNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionNameContext functionName() throws RecognitionException {
		FunctionNameContext _localctx = new FunctionNameContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_functionName);
		try {
			setState(1017);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_CONTINUE:
			case KEYWORD_ELSE:
			case KEYWORD_ENUM:
			case KEYWORD_FALSE:
			case KEYWORD_FUNC:
			case KEYWORD_IMPORT:
			case KEYWORD_IN:
			case KEYWORD_INTERFACE:
			case KEYWORD_IS:
			case KEYWORD_MATCH:
			case KEYWORD_NOT:
			case KEYWORD_NULL:
			case KEYWORD_OR:
			case KEYWORD_PACKAGE:
			case KEYWORD_STRUCT:
			case KEYWORD_TRUE:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(1015);
				declarationIdentifier();
				}
				break;
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case OPERATOR_HEAD_OTHER:
				enterOuterAlt(_localctx, 2);
				{
				setState(1016);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionSignatureContext extends ParserRuleContext {
		public FunctionParameterClauseContext functionParameterClause() {
			return getRuleContext(FunctionParameterClauseContext.class,0);
		}
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public FunctionResultContext functionResult() {
			return getRuleContext(FunctionResultContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionSignatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionSignature; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionSignature(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionSignatureContext functionSignature() throws RecognitionException {
		FunctionSignatureContext _localctx = new FunctionSignatureContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_functionSignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1019);
			functionParameterClause();
			setState(1021);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,94,_ctx) ) {
			case 1:
				{
				setState(1020);
				followingDocument();
				}
				break;
			}
			setState(1030);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,96,_ctx) ) {
			case 1:
				{
				setState(1026);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1023);
					match(EOL);
					}
					}
					setState(1028);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1029);
				functionResult();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionResultContext extends ParserRuleContext {
		public ArrowOperatorContext arrowOperator() {
			return getRuleContext(ArrowOperatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionResultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionResult; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionResult(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionResultContext functionResult() throws RecognitionException {
		FunctionResultContext _localctx = new FunctionResultContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_functionResult);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1032);
			arrowOperator();
			setState(1037);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,98,_ctx) ) {
			case 1:
				{
				setState(1033);
				labelIdentifier();
				setState(1035);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1034);
					match(COLON);
					}
				}

				}
				break;
			}
			setState(1039);
			type_(0);
			setState(1041);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(1040);
				attributes();
				}
			}

			setState(1050);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,101,_ctx) ) {
			case 1:
				{
				setState(1046);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1043);
					match(EOL);
					}
					}
					setState(1048);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1049);
				followingDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionBodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionBody; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionBody(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionBodyContext functionBody() throws RecognitionException {
		FunctionBodyContext _localctx = new FunctionBodyContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_functionBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1052);
			match(LCURLY);
			setState(1054);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1053);
				followingDocument();
				}
			}

			setState(1063);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,104,_ctx) ) {
			case 1:
				{
				setState(1059);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1056);
					match(EOL);
					}
					}
					setState(1061);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1062);
				statements();
				}
				break;
			}
			setState(1068);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1065);
				match(EOL);
				}
				}
				setState(1070);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1071);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionParameterClauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public FunctionParametersContext functionParameters() {
			return getRuleContext(FunctionParametersContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionParameterClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionParameterClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionParameterClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionParameterClauseContext functionParameterClause() throws RecognitionException {
		FunctionParameterClauseContext _localctx = new FunctionParameterClauseContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_functionParameterClause);
		int _la;
		try {
			setState(1091);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,108,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1073);
				match(LPAREN);
				setState(1074);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1075);
				match(LPAREN);
				setState(1079);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1076);
					match(EOL);
					}
					}
					setState(1081);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1082);
				functionParameters();
				setState(1086);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1083);
					match(EOL);
					}
					}
					setState(1088);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1089);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionParametersContext extends ParserRuleContext {
		public List<FunctionParameterContext> functionParameter() {
			return getRuleContexts(FunctionParameterContext.class);
		}
		public FunctionParameterContext functionParameter(int i) {
			return getRuleContext(FunctionParameterContext.class,i);
		}
		public List<EovWithDocumentContext> eovWithDocument() {
			return getRuleContexts(EovWithDocumentContext.class);
		}
		public EovWithDocumentContext eovWithDocument(int i) {
			return getRuleContext(EovWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionParameters; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionParameters(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionParametersContext functionParameters() throws RecognitionException {
		FunctionParametersContext _localctx = new FunctionParametersContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_functionParameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1093);
			functionParameter();
			setState(1105);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,110,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1094);
					eovWithDocument();
					setState(1098);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1095);
						match(EOL);
						}
						}
						setState(1100);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1101);
					functionParameter();
					}
					} 
				}
				setState(1107);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,110,_ctx);
			}
			setState(1109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,111,_ctx) ) {
			case 1:
				{
				setState(1108);
				eovWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionParameterContext extends ParserRuleContext {
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public FunctionParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionParameter; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionParameter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionParameterContext functionParameter() throws RecognitionException {
		FunctionParameterContext _localctx = new FunctionParameterContext(_ctx, getState());
		enterRule(_localctx, 120, RULE_functionParameter);
		int _la;
		try {
			setState(1131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,116,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1111);
				labelIdentifier();
				setState(1112);
				typeAnnotation();
				setState(1120);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,113,_ctx) ) {
				case 1:
					{
					setState(1116);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1113);
						match(EOL);
						}
						}
						setState(1118);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1119);
					initializer();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1122);
				labelIdentifier();
				setState(1124);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1123);
					match(COLON);
					}
				}

				setState(1126);
				type_(0);
				setState(1127);
				match(ELLIPSIS);
				setState(1129);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1128);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnumDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_ENUM() { return getToken(MojoParser.KEYWORD_ENUM, 0); }
		public EnumNameContext enumName() {
			return getRuleContext(EnumNameContext.class,0);
		}
		public EnumBodyContext enumBody() {
			return getRuleContext(EnumBodyContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public TypeInheritanceClauseContext typeInheritanceClause() {
			return getRuleContext(TypeInheritanceClauseContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public EnumDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEnumDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EnumDeclarationContext enumDeclaration() throws RecognitionException {
		EnumDeclarationContext _localctx = new EnumDeclarationContext(_ctx, getState());
		enterRule(_localctx, 122, RULE_enumDeclaration);
		int _la;
		try {
			setState(1179);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_ENUM:
				enterOuterAlt(_localctx, 1);
				{
				setState(1133);
				match(KEYWORD_ENUM);
				setState(1134);
				enumName();
				setState(1136);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1135);
					genericParameterClause();
					}
				}

				setState(1145);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,119,_ctx) ) {
				case 1:
					{
					setState(1141);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1138);
						match(EOL);
						}
						}
						setState(1143);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1144);
					typeInheritanceClause();
					}
					break;
				}
				setState(1150);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1147);
					match(EOL);
					}
					}
					setState(1152);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1153);
				enumBody();
				}
				break;
			case KEYWORD_TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(1155);
				match(KEYWORD_TYPE);
				setState(1156);
				enumName();
				setState(1157);
				assignmentOperator();
				setState(1158);
				match(KEYWORD_ENUM);
				setState(1160);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1159);
					genericParameterClause();
					}
				}

				setState(1169);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,123,_ctx) ) {
				case 1:
					{
					setState(1165);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1162);
						match(EOL);
						}
						}
						setState(1167);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1168);
					typeInheritanceClause();
					}
					break;
				}
				setState(1174);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1171);
					match(EOL);
					}
					}
					setState(1176);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1177);
				enumBody();
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnumBodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public EnumMembersContext enumMembers() {
			return getRuleContext(EnumMembersContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public EnumBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumBody; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEnumBody(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EnumBodyContext enumBody() throws RecognitionException {
		EnumBodyContext _localctx = new EnumBodyContext(_ctx, getState());
		enterRule(_localctx, 124, RULE_enumBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1181);
			match(LCURLY);
			setState(1183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1182);
				followingDocument();
				}
			}

			setState(1192);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,128,_ctx) ) {
			case 1:
				{
				setState(1188);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1185);
					match(EOL);
					}
					}
					setState(1190);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1191);
				enumMembers();
				}
				break;
			}
			setState(1197);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1194);
				match(EOL);
				}
				}
				setState(1199);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1200);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnumNameContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public EnumNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEnumName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EnumNameContext enumName() throws RecognitionException {
		EnumNameContext _localctx = new EnumNameContext(_ctx, getState());
		enterRule(_localctx, 126, RULE_enumName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1202);
			typeName();
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnumMembersContext extends ParserRuleContext {
		public List<EnumMemberContext> enumMember() {
			return getRuleContexts(EnumMemberContext.class);
		}
		public EnumMemberContext enumMember(int i) {
			return getRuleContext(EnumMemberContext.class,i);
		}
		public List<EosWithDocumentContext> eosWithDocument() {
			return getRuleContexts(EosWithDocumentContext.class);
		}
		public EosWithDocumentContext eosWithDocument(int i) {
			return getRuleContext(EosWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public EnumMembersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumMembers; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEnumMembers(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EnumMembersContext enumMembers() throws RecognitionException {
		EnumMembersContext _localctx = new EnumMembersContext(_ctx, getState());
		enterRule(_localctx, 128, RULE_enumMembers);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1204);
			enumMember();
			setState(1216);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,131,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1205);
					eosWithDocument();
					setState(1209);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1206);
						match(EOL);
						}
						}
						setState(1211);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1212);
					enumMember();
					}
					} 
				}
				setState(1218);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,131,_ctx);
			}
			setState(1220);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,132,_ctx) ) {
			case 1:
				{
				setState(1219);
				eosWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnumMemberContext extends ParserRuleContext {
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public List<AttributesContext> attributes() {
			return getRuleContexts(AttributesContext.class);
		}
		public AttributesContext attributes(int i) {
			return getRuleContext(AttributesContext.class,i);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public FloatingStatementContext floatingStatement() {
			return getRuleContext(FloatingStatementContext.class,0);
		}
		public EnumMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumMember; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEnumMember(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EnumMemberContext enumMember() throws RecognitionException {
		EnumMemberContext _localctx = new EnumMemberContext(_ctx, getState());
		enterRule(_localctx, 130, RULE_enumMember);
		int _la;
		try {
			setState(1246);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,138,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1225);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LINE_DOCUMENT) {
					{
					setState(1222);
					document();
					setState(1223);
					match(EOL);
					}
				}

				setState(1230);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1227);
					attributes();
					setState(1228);
					match(EOL);
					}
				}

				setState(1232);
				declarationIdentifier();
				setState(1234);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1233);
					attributes();
					}
				}

				setState(1243);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,137,_ctx) ) {
				case 1:
					{
					setState(1239);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1236);
						match(EOL);
						}
						}
						setState(1241);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1242);
					initializer();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1245);
				floatingStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclarationContext extends ParserRuleContext {
		public StructNameContext structName() {
			return getRuleContext(StructNameContext.class,0);
		}
		public StructTypeContext structType() {
			return getRuleContext(StructTypeContext.class,0);
		}
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public TerminalNode KEYWORD_STRUCT() { return getToken(MojoParser.KEYWORD_STRUCT, 0); }
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public StructDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructDeclarationContext structDeclaration() throws RecognitionException {
		StructDeclarationContext _localctx = new StructDeclarationContext(_ctx, getState());
		enterRule(_localctx, 132, RULE_structDeclaration);
		int _la;
		try {
			setState(1264);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,141,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1248);
				_la = _input.LA(1);
				if ( !(_la==KEYWORD_STRUCT || _la==KEYWORD_TYPE) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(1249);
				structName();
				setState(1251);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1250);
					genericParameterClause();
					}
				}

				setState(1253);
				structType();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1255);
				match(KEYWORD_TYPE);
				setState(1256);
				structName();
				setState(1257);
				assignmentOperator();
				setState(1258);
				match(KEYWORD_STRUCT);
				setState(1260);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1259);
					genericParameterClause();
					}
				}

				setState(1262);
				structType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructNameContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public StructNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructNameContext structName() throws RecognitionException {
		StructNameContext _localctx = new StructNameContext(_ctx, getState());
		enterRule(_localctx, 134, RULE_structName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1266);
			typeName();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructTypeContext extends ParserRuleContext {
		public TypeInheritanceClauseContext typeInheritanceClause() {
			return getRuleContext(TypeInheritanceClauseContext.class,0);
		}
		public StructBodyContext structBody() {
			return getRuleContext(StructBodyContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public StructTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructTypeContext structType() throws RecognitionException {
		StructTypeContext _localctx = new StructTypeContext(_ctx, getState());
		enterRule(_localctx, 136, RULE_structType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1275);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,143,_ctx) ) {
			case 1:
				{
				setState(1271);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1268);
					match(EOL);
					}
					}
					setState(1273);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1274);
				typeInheritanceClause();
				}
				break;
			}
			setState(1284);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,145,_ctx) ) {
			case 1:
				{
				setState(1280);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1277);
					match(EOL);
					}
					}
					setState(1282);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1283);
				structBody();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructBodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public StructMembersContext structMembers() {
			return getRuleContext(StructMembersContext.class,0);
		}
		public StructBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structBody; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructBody(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructBodyContext structBody() throws RecognitionException {
		StructBodyContext _localctx = new StructBodyContext(_ctx, getState());
		enterRule(_localctx, 138, RULE_structBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1286);
			match(LCURLY);
			setState(1290);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1287);
				followingDocument();
				setState(1288);
				match(EOL);
				}
			}

			setState(1299);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,148,_ctx) ) {
			case 1:
				{
				setState(1295);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1292);
					match(EOL);
					}
					}
					setState(1297);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1298);
				structMembers();
				}
				break;
			}
			setState(1304);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1301);
				match(EOL);
				}
				}
				setState(1306);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1307);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructMembersContext extends ParserRuleContext {
		public List<StructMemberContext> structMember() {
			return getRuleContexts(StructMemberContext.class);
		}
		public StructMemberContext structMember(int i) {
			return getRuleContext(StructMemberContext.class,i);
		}
		public List<EosWithDocumentContext> eosWithDocument() {
			return getRuleContexts(EosWithDocumentContext.class);
		}
		public EosWithDocumentContext eosWithDocument(int i) {
			return getRuleContext(EosWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public StructMembersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMembers; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructMembers(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructMembersContext structMembers() throws RecognitionException {
		StructMembersContext _localctx = new StructMembersContext(_ctx, getState());
		enterRule(_localctx, 140, RULE_structMembers);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1309);
			structMember();
			setState(1321);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,151,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1310);
					eosWithDocument();
					setState(1314);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1311);
						match(EOL);
						}
						}
						setState(1316);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1317);
					structMember();
					}
					} 
				}
				setState(1323);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,151,_ctx);
			}
			setState(1325);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,152,_ctx) ) {
			case 1:
				{
				setState(1324);
				eosWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructMemberContext extends ParserRuleContext {
		public StructDeclarationContext structDeclaration() {
			return getRuleContext(StructDeclarationContext.class,0);
		}
		public EnumDeclarationContext enumDeclaration() {
			return getRuleContext(EnumDeclarationContext.class,0);
		}
		public ConstantDeclarationContext constantDeclaration() {
			return getRuleContext(ConstantDeclarationContext.class,0);
		}
		public TypeAliasDeclarationContext typeAliasDeclaration() {
			return getRuleContext(TypeAliasDeclarationContext.class,0);
		}
		public StructMemberDeclarationContext structMemberDeclaration() {
			return getRuleContext(StructMemberDeclarationContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public FloatingStatementContext floatingStatement() {
			return getRuleContext(FloatingStatementContext.class,0);
		}
		public StructMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMember; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructMember(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructMemberContext structMember() throws RecognitionException {
		StructMemberContext _localctx = new StructMemberContext(_ctx, getState());
		enterRule(_localctx, 142, RULE_structMember);
		int _la;
		try {
			setState(1345);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,156,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1330);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LINE_DOCUMENT) {
					{
					setState(1327);
					document();
					setState(1328);
					match(EOL);
					}
				}

				setState(1335);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1332);
					attributes();
					setState(1333);
					match(EOL);
					}
				}

				setState(1342);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,155,_ctx) ) {
				case 1:
					{
					setState(1337);
					structDeclaration();
					}
					break;
				case 2:
					{
					setState(1338);
					enumDeclaration();
					}
					break;
				case 3:
					{
					setState(1339);
					constantDeclaration();
					}
					break;
				case 4:
					{
					setState(1340);
					typeAliasDeclaration();
					}
					break;
				case 5:
					{
					setState(1341);
					structMemberDeclaration();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1344);
				floatingStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructMemberDeclarationContext extends ParserRuleContext {
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public StructMemberDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMemberDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructMemberDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructMemberDeclarationContext structMemberDeclaration() throws RecognitionException {
		StructMemberDeclarationContext _localctx = new StructMemberDeclarationContext(_ctx, getState());
		enterRule(_localctx, 144, RULE_structMemberDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1347);
			declarationIdentifier();
			setState(1348);
			typeAnnotation();
			setState(1356);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,158,_ctx) ) {
			case 1:
				{
				setState(1352);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1349);
					match(EOL);
					}
					}
					setState(1354);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1355);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_INTERFACE() { return getToken(MojoParser.KEYWORD_INTERFACE, 0); }
		public InterfaceNameContext interfaceName() {
			return getRuleContext(InterfaceNameContext.class,0);
		}
		public InterfaceTypeContext interfaceType() {
			return getRuleContext(InterfaceTypeContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public InterfaceDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceDeclarationContext interfaceDeclaration() throws RecognitionException {
		InterfaceDeclarationContext _localctx = new InterfaceDeclarationContext(_ctx, getState());
		enterRule(_localctx, 146, RULE_interfaceDeclaration);
		int _la;
		try {
			setState(1374);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_INTERFACE:
				enterOuterAlt(_localctx, 1);
				{
				setState(1358);
				match(KEYWORD_INTERFACE);
				setState(1359);
				interfaceName();
				setState(1361);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1360);
					genericParameterClause();
					}
				}

				setState(1363);
				interfaceType();
				}
				break;
			case KEYWORD_TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(1365);
				match(KEYWORD_TYPE);
				setState(1366);
				interfaceName();
				setState(1367);
				assignmentOperator();
				setState(1368);
				match(KEYWORD_INTERFACE);
				setState(1370);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(1369);
					genericParameterClause();
					}
				}

				setState(1372);
				interfaceType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceNameContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public InterfaceNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceNameContext interfaceName() throws RecognitionException {
		InterfaceNameContext _localctx = new InterfaceNameContext(_ctx, getState());
		enterRule(_localctx, 148, RULE_interfaceName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1376);
			typeName();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceTypeContext extends ParserRuleContext {
		public InterfaceBodyContext interfaceBody() {
			return getRuleContext(InterfaceBodyContext.class,0);
		}
		public TypeInheritanceClauseContext typeInheritanceClause() {
			return getRuleContext(TypeInheritanceClauseContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public InterfaceTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceTypeContext interfaceType() throws RecognitionException {
		InterfaceTypeContext _localctx = new InterfaceTypeContext(_ctx, getState());
		enterRule(_localctx, 150, RULE_interfaceType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1385);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,163,_ctx) ) {
			case 1:
				{
				setState(1381);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1378);
					match(EOL);
					}
					}
					setState(1383);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1384);
				typeInheritanceClause();
				}
				break;
			}
			setState(1390);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1387);
				match(EOL);
				}
				}
				setState(1392);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1393);
			interfaceBody();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceBodyContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public InterfaceMembersContext interfaceMembers() {
			return getRuleContext(InterfaceMembersContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public InterfaceBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceBody; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceBody(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceBodyContext interfaceBody() throws RecognitionException {
		InterfaceBodyContext _localctx = new InterfaceBodyContext(_ctx, getState());
		enterRule(_localctx, 152, RULE_interfaceBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1395);
			match(LCURLY);
			setState(1397);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1396);
				followingDocument();
				}
			}

			setState(1406);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,167,_ctx) ) {
			case 1:
				{
				setState(1402);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1399);
					match(EOL);
					}
					}
					setState(1404);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1405);
				interfaceMembers();
				}
				break;
			}
			setState(1411);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1408);
				match(EOL);
				}
				}
				setState(1413);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1414);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceMembersContext extends ParserRuleContext {
		public List<InterfaceMemberContext> interfaceMember() {
			return getRuleContexts(InterfaceMemberContext.class);
		}
		public InterfaceMemberContext interfaceMember(int i) {
			return getRuleContext(InterfaceMemberContext.class,i);
		}
		public List<EosWithDocumentContext> eosWithDocument() {
			return getRuleContexts(EosWithDocumentContext.class);
		}
		public EosWithDocumentContext eosWithDocument(int i) {
			return getRuleContext(EosWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public InterfaceMembersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceMembers; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceMembers(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceMembersContext interfaceMembers() throws RecognitionException {
		InterfaceMembersContext _localctx = new InterfaceMembersContext(_ctx, getState());
		enterRule(_localctx, 154, RULE_interfaceMembers);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1416);
			interfaceMember();
			setState(1428);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,170,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1417);
					eosWithDocument();
					setState(1421);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1418);
						match(EOL);
						}
						}
						setState(1423);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1424);
					interfaceMember();
					}
					} 
				}
				setState(1430);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,170,_ctx);
			}
			setState(1432);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,171,_ctx) ) {
			case 1:
				{
				setState(1431);
				eosWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceMemberContext extends ParserRuleContext {
		public TypeAliasDeclarationContext typeAliasDeclaration() {
			return getRuleContext(TypeAliasDeclarationContext.class,0);
		}
		public InterfaceMethodDeclarationContext interfaceMethodDeclaration() {
			return getRuleContext(InterfaceMethodDeclarationContext.class,0);
		}
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public FloatingStatementContext floatingStatement() {
			return getRuleContext(FloatingStatementContext.class,0);
		}
		public InterfaceMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceMember; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceMember(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceMemberContext interfaceMember() throws RecognitionException {
		InterfaceMemberContext _localctx = new InterfaceMemberContext(_ctx, getState());
		enterRule(_localctx, 156, RULE_interfaceMember);
		int _la;
		try {
			setState(1449);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,175,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1437);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LINE_DOCUMENT) {
					{
					setState(1434);
					document();
					setState(1435);
					match(EOL);
					}
				}

				setState(1442);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1439);
					attributes();
					setState(1440);
					match(EOL);
					}
				}

				setState(1446);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,174,_ctx) ) {
				case 1:
					{
					setState(1444);
					typeAliasDeclaration();
					}
					break;
				case 2:
					{
					setState(1445);
					interfaceMethodDeclaration();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1448);
				floatingStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceMethodDeclarationContext extends ParserRuleContext {
		public FunctionNameContext functionName() {
			return getRuleContext(FunctionNameContext.class,0);
		}
		public FunctionSignatureContext functionSignature() {
			return getRuleContext(FunctionSignatureContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public InterfaceMethodDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceMethodDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInterfaceMethodDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InterfaceMethodDeclarationContext interfaceMethodDeclaration() throws RecognitionException {
		InterfaceMethodDeclarationContext _localctx = new InterfaceMethodDeclarationContext(_ctx, getState());
		enterRule(_localctx, 158, RULE_interfaceMethodDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1451);
			functionName();
			setState(1453);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1452);
				genericParameterClause();
				}
			}

			setState(1458);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1455);
				match(EOL);
				}
				}
				setState(1460);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1461);
			functionSignature();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_ATTRIBUTE() { return getToken(MojoParser.KEYWORD_ATTRIBUTE, 0); }
		public AttributeNameContext attributeName() {
			return getRuleContext(AttributeNameContext.class,0);
		}
		public StructBodyContext structBody() {
			return getRuleContext(StructBodyContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributeDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeDeclarationContext attributeDeclaration() throws RecognitionException {
		AttributeDeclarationContext _localctx = new AttributeDeclarationContext(_ctx, getState());
		enterRule(_localctx, 160, RULE_attributeDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1463);
			match(KEYWORD_ATTRIBUTE);
			setState(1464);
			attributeName();
			setState(1466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1465);
				genericParameterClause();
				}
			}

			setState(1482);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,182,_ctx) ) {
			case 1:
				{
				setState(1468);
				structBody();
				}
				break;
			case 2:
				{
				setState(1469);
				typeAnnotation();
				setState(1477);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,180,_ctx) ) {
				case 1:
					{
					setState(1473);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1470);
						match(EOL);
						}
						}
						setState(1475);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1476);
					initializer();
					}
					break;
				}
				setState(1480);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(1479);
					followingDocument();
					}
				}

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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeAliasDeclarationContext extends ParserRuleContext {
		public TerminalNode KEYWORD_ATTRIBUTE() { return getToken(MojoParser.KEYWORD_ATTRIBUTE, 0); }
		public AttributeNameContext attributeName() {
			return getRuleContext(AttributeNameContext.class,0);
		}
		public AttributeAliasAssignmentContext attributeAliasAssignment() {
			return getRuleContext(AttributeAliasAssignmentContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributeAliasDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeAliasDeclaration; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeAliasDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeAliasDeclarationContext attributeAliasDeclaration() throws RecognitionException {
		AttributeAliasDeclarationContext _localctx = new AttributeAliasDeclarationContext(_ctx, getState());
		enterRule(_localctx, 162, RULE_attributeAliasDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1484);
			match(KEYWORD_ATTRIBUTE);
			setState(1485);
			attributeName();
			setState(1487);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1486);
				genericParameterClause();
				}
			}

			setState(1492);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1489);
				match(EOL);
				}
				}
				setState(1494);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1495);
			attributeAliasAssignment();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeAliasAssignmentContext extends ParserRuleContext {
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public AttributeNameContext attributeName() {
			return getRuleContext(AttributeNameContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public PackageIdentifierContext packageIdentifier() {
			return getRuleContext(PackageIdentifierContext.class,0);
		}
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public GenericArgumentClauseContext genericArgumentClause() {
			return getRuleContext(GenericArgumentClauseContext.class,0);
		}
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public AttributeAliasAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeAliasAssignment; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeAliasAssignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeAliasAssignmentContext attributeAliasAssignment() throws RecognitionException {
		AttributeAliasAssignmentContext _localctx = new AttributeAliasAssignmentContext(_ctx, getState());
		enterRule(_localctx, 164, RULE_attributeAliasAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1497);
			assignmentOperator();
			setState(1501);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1498);
				match(EOL);
				}
				}
				setState(1503);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1507);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,186,_ctx) ) {
			case 1:
				{
				setState(1504);
				packageIdentifier();
				setState(1505);
				match(DOT);
				}
				break;
			}
			setState(1509);
			attributeName();
			setState(1511);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1510);
				genericArgumentClause();
				}
			}

			setState(1514);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1513);
				followingDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PatternContext extends ParserRuleContext {
		public WildcardPatternContext wildcardPattern() {
			return getRuleContext(WildcardPatternContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public IdentifierPatternContext identifierPattern() {
			return getRuleContext(IdentifierPatternContext.class,0);
		}
		public TuplePatternContext tuplePattern() {
			return getRuleContext(TuplePatternContext.class,0);
		}
		public ArrayPatternContext arrayPattern() {
			return getRuleContext(ArrayPatternContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public EnumValuePatternContext enumValuePattern() {
			return getRuleContext(EnumValuePatternContext.class,0);
		}
		public OptionalPatternContext optionalPattern() {
			return getRuleContext(OptionalPatternContext.class,0);
		}
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public TerminalNode BANG() { return getToken(MojoParser.BANG, 0); }
		public ExpressionPatternContext expressionPattern() {
			return getRuleContext(ExpressionPatternContext.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode KEYWORD_AS() { return getToken(MojoParser.KEYWORD_AS, 0); }
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPattern(this);
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
		int _startState = 166;
		enterRecursionRule(_localctx, 166, RULE_pattern, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1545);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,195,_ctx) ) {
			case 1:
				{
				setState(1517);
				wildcardPattern();
				setState(1519);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,189,_ctx) ) {
				case 1:
					{
					setState(1518);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 2:
				{
				setState(1521);
				identifierPattern();
				setState(1523);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,190,_ctx) ) {
				case 1:
					{
					setState(1522);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 3:
				{
				setState(1525);
				tuplePattern();
				setState(1527);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,191,_ctx) ) {
				case 1:
					{
					setState(1526);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 4:
				{
				setState(1529);
				arrayPattern();
				setState(1531);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,192,_ctx) ) {
				case 1:
					{
					setState(1530);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 5:
				{
				setState(1533);
				type_(0);
				setState(1535);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,193,_ctx) ) {
				case 1:
					{
					setState(1534);
					attributes();
					}
					break;
				}
				}
				break;
			case 6:
				{
				setState(1537);
				enumValuePattern();
				}
				break;
			case 7:
				{
				setState(1538);
				optionalPattern();
				}
				break;
			case 8:
				{
				setState(1540);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==BANG) {
					{
					setState(1539);
					match(BANG);
					}
				}

				setState(1542);
				match(KEYWORD_IS);
				setState(1543);
				type_(0);
				}
				break;
			case 9:
				{
				setState(1544);
				expressionPattern();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1552);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,196,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PatternContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pattern);
					setState(1547);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(1548);
					match(KEYWORD_AS);
					setState(1549);
					type_(0);
					}
					} 
				}
				setState(1554);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,196,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WildcardPatternContext extends ParserRuleContext {
		public TerminalNode UNDERSCORE() { return getToken(MojoParser.UNDERSCORE, 0); }
		public WildcardPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_wildcardPattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitWildcardPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final WildcardPatternContext wildcardPattern() throws RecognitionException {
		WildcardPatternContext _localctx = new WildcardPatternContext(_ctx, getState());
		enterRule(_localctx, 168, RULE_wildcardPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1555);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierPatternContext extends ParserRuleContext {
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public IdentifierPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierPattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIdentifierPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IdentifierPatternContext identifierPattern() throws RecognitionException {
		IdentifierPatternContext _localctx = new IdentifierPatternContext(_ctx, getState());
		enterRule(_localctx, 170, RULE_identifierPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1557);
			declarationIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TuplePatternContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public TuplePatternElementListContext tuplePatternElementList() {
			return getRuleContext(TuplePatternElementListContext.class,0);
		}
		public TuplePatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuplePattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTuplePattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TuplePatternContext tuplePattern() throws RecognitionException {
		TuplePatternContext _localctx = new TuplePatternContext(_ctx, getState());
		enterRule(_localctx, 172, RULE_tuplePattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1559);
			match(LPAREN);
			setState(1561);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 90062655082982398L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1046273L) != 0)) {
				{
				setState(1560);
				tuplePatternElementList();
				}
			}

			setState(1563);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TuplePatternElementListContext extends ParserRuleContext {
		public List<TuplePatternElementContext> tuplePatternElement() {
			return getRuleContexts(TuplePatternElementContext.class);
		}
		public TuplePatternElementContext tuplePatternElement(int i) {
			return getRuleContext(TuplePatternElementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public TuplePatternElementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuplePatternElementList; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTuplePatternElementList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TuplePatternElementListContext tuplePatternElementList() throws RecognitionException {
		TuplePatternElementListContext _localctx = new TuplePatternElementListContext(_ctx, getState());
		enterRule(_localctx, 174, RULE_tuplePatternElementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1565);
			tuplePatternElement();
			setState(1570);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1566);
				match(COMMA);
				setState(1567);
				tuplePatternElement();
				}
				}
				setState(1572);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TuplePatternElementContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public TuplePatternElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuplePatternElement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTuplePatternElement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TuplePatternElementContext tuplePatternElement() throws RecognitionException {
		TuplePatternElementContext _localctx = new TuplePatternElementContext(_ctx, getState());
		enterRule(_localctx, 176, RULE_tuplePatternElement);
		try {
			setState(1575);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,199,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1573);
				pattern(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1574);
				match(ELLIPSIS);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayPatternContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public ArrayPatternElementsContext arrayPatternElements() {
			return getRuleContext(ArrayPatternElementsContext.class,0);
		}
		public ArrayPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayPattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayPatternContext arrayPattern() throws RecognitionException {
		ArrayPatternContext _localctx = new ArrayPatternContext(_ctx, getState());
		enterRule(_localctx, 178, RULE_arrayPattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1577);
			match(LBRACK);
			setState(1579);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 90062655082982398L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1046273L) != 0)) {
				{
				setState(1578);
				arrayPatternElements();
				}
			}

			setState(1581);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayPatternElementsContext extends ParserRuleContext {
		public List<ArrayPatternElementContext> arrayPatternElement() {
			return getRuleContexts(ArrayPatternElementContext.class);
		}
		public ArrayPatternElementContext arrayPatternElement(int i) {
			return getRuleContext(ArrayPatternElementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public ArrayPatternElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayPatternElements; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayPatternElements(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayPatternElementsContext arrayPatternElements() throws RecognitionException {
		ArrayPatternElementsContext _localctx = new ArrayPatternElementsContext(_ctx, getState());
		enterRule(_localctx, 180, RULE_arrayPatternElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1583);
			arrayPatternElement();
			setState(1588);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1584);
				match(COMMA);
				setState(1585);
				arrayPatternElement();
				}
				}
				setState(1590);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayPatternElementContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public ArrayPatternElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayPatternElement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayPatternElement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayPatternElementContext arrayPatternElement() throws RecognitionException {
		ArrayPatternElementContext _localctx = new ArrayPatternElementContext(_ctx, getState());
		enterRule(_localctx, 182, RULE_arrayPatternElement);
		try {
			setState(1593);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,202,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1591);
				pattern(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1592);
				match(ELLIPSIS);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnumValuePatternContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public TypeIdentifierContext typeIdentifier() {
			return getRuleContext(TypeIdentifierContext.class,0);
		}
		public TuplePatternContext tuplePattern() {
			return getRuleContext(TuplePatternContext.class,0);
		}
		public EnumValuePatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumValuePattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEnumValuePattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EnumValuePatternContext enumValuePattern() throws RecognitionException {
		EnumValuePatternContext _localctx = new EnumValuePatternContext(_ctx, getState());
		enterRule(_localctx, 184, RULE_enumValuePattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1596);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TYPE_IDENTIFIER || _la==VALUE_IDENTIFIER) {
				{
				setState(1595);
				typeIdentifier();
				}
			}

			setState(1598);
			match(DOT);
			setState(1599);
			declarationIdentifier();
			setState(1601);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,204,_ctx) ) {
			case 1:
				{
				setState(1600);
				tuplePattern();
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

	@SuppressWarnings("CheckReturnValue")
	public static class OptionalPatternContext extends ParserRuleContext {
		public IdentifierPatternContext identifierPattern() {
			return getRuleContext(IdentifierPatternContext.class,0);
		}
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public OptionalPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optionalPattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitOptionalPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OptionalPatternContext optionalPattern() throws RecognitionException {
		OptionalPatternContext _localctx = new OptionalPatternContext(_ctx, getState());
		enterRule(_localctx, 186, RULE_optionalPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1603);
			identifierPattern();
			setState(1604);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionPatternContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionPattern; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitExpressionPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionPatternContext expressionPattern() throws RecognitionException {
		ExpressionPatternContext _localctx = new ExpressionPatternContext(_ctx, getState());
		enterRule(_localctx, 188, RULE_expressionPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1606);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeContext extends ParserRuleContext {
		public TerminalNode AT() { return getToken(MojoParser.AT, 0); }
		public TerminalNode DECIMAL_LITERAL() { return getToken(MojoParser.DECIMAL_LITERAL, 0); }
		public AttributeIdentifierContext attributeIdentifier() {
			return getRuleContext(AttributeIdentifierContext.class,0);
		}
		public GenericArgumentClauseContext genericArgumentClause() {
			return getRuleContext(GenericArgumentClauseContext.class,0);
		}
		public AttributeArgumentClauseContext attributeArgumentClause() {
			return getRuleContext(AttributeArgumentClauseContext.class,0);
		}
		public AttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttribute(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeContext attribute() throws RecognitionException {
		AttributeContext _localctx = new AttributeContext(_ctx, getState());
		enterRule(_localctx, 190, RULE_attribute);
		try {
			setState(1618);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,207,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1608);
				match(AT);
				setState(1609);
				match(DECIMAL_LITERAL);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1610);
				match(AT);
				setState(1611);
				attributeIdentifier();
				setState(1613);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,205,_ctx) ) {
				case 1:
					{
					setState(1612);
					genericArgumentClause();
					}
					break;
				}
				setState(1616);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,206,_ctx) ) {
				case 1:
					{
					setState(1615);
					attributeArgumentClause();
					}
					break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeIdentifierContext extends ParserRuleContext {
		public AttributeNameContext attributeName() {
			return getRuleContext(AttributeNameContext.class,0);
		}
		public PackageIdentifierContext packageIdentifier() {
			return getRuleContext(PackageIdentifierContext.class,0);
		}
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public AttributeIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeIdentifierContext attributeIdentifier() throws RecognitionException {
		AttributeIdentifierContext _localctx = new AttributeIdentifierContext(_ctx, getState());
		enterRule(_localctx, 192, RULE_attributeIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(1623);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,208,_ctx) ) {
			case 1:
				{
				setState(1620);
				packageIdentifier();
				setState(1621);
				match(DOT);
				}
				break;
			}
			setState(1625);
			attributeName();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeNameContext extends ParserRuleContext {
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public AttributeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeNameContext attributeName() throws RecognitionException {
		AttributeNameContext _localctx = new AttributeNameContext(_ctx, getState());
		enterRule(_localctx, 194, RULE_attributeName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1627);
			labelIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeArgumentClauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public AttributeArgumentsContext attributeArguments() {
			return getRuleContext(AttributeArgumentsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributeArgumentClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeArgumentClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeArgumentClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeArgumentClauseContext attributeArgumentClause() throws RecognitionException {
		AttributeArgumentClauseContext _localctx = new AttributeArgumentClauseContext(_ctx, getState());
		enterRule(_localctx, 196, RULE_attributeArgumentClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1629);
			match(LPAREN);
			setState(1637);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,210,_ctx) ) {
			case 1:
				{
				setState(1633);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1630);
					match(EOL);
					}
					}
					setState(1635);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1636);
				attributeArguments();
				}
				break;
			}
			setState(1642);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1639);
				match(EOL);
				}
				}
				setState(1644);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1645);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeArgumentContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public AttributeArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeArgument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeArgument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeArgumentContext attributeArgument() throws RecognitionException {
		AttributeArgumentContext _localctx = new AttributeArgumentContext(_ctx, getState());
		enterRule(_localctx, 198, RULE_attributeArgument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1650);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,212,_ctx) ) {
			case 1:
				{
				setState(1647);
				labelIdentifier();
				setState(1648);
				match(COLON);
				}
				break;
			}
			setState(1652);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeArgumentsContext extends ParserRuleContext {
		public List<AttributeArgumentContext> attributeArgument() {
			return getRuleContexts(AttributeArgumentContext.class);
		}
		public AttributeArgumentContext attributeArgument(int i) {
			return getRuleContext(AttributeArgumentContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributeArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeArguments; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributeArguments(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeArgumentsContext attributeArguments() throws RecognitionException {
		AttributeArgumentsContext _localctx = new AttributeArgumentsContext(_ctx, getState());
		enterRule(_localctx, 200, RULE_attributeArguments);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1654);
			attributeArgument();
			setState(1666);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,214,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1655);
					eov();
					setState(1659);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1656);
						match(EOL);
						}
						}
						setState(1661);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1662);
					attributeArgument();
					}
					} 
				}
				setState(1668);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,214,_ctx);
			}
			setState(1670);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,215,_ctx) ) {
			case 1:
				{
				setState(1669);
				eov();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributesContext extends ParserRuleContext {
		public List<AttributeContext> attribute() {
			return getRuleContexts(AttributeContext.class);
		}
		public AttributeContext attribute(int i) {
			return getRuleContext(AttributeContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public AttributesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributes; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAttributes(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributesContext attributes() throws RecognitionException {
		AttributesContext _localctx = new AttributesContext(_ctx, getState());
		enterRule(_localctx, 202, RULE_attributes);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1672);
			attribute();
			setState(1679);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,217,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1674);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==EOL) {
						{
						setState(1673);
						match(EOL);
						}
					}

					setState(1676);
					attribute();
					}
					} 
				}
				setState(1681);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,217,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public PrefixExpressionContext prefixExpression() {
			return getRuleContext(PrefixExpressionContext.class,0);
		}
		public BinaryExpressionsContext binaryExpressions() {
			return getRuleContext(BinaryExpressionsContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 204, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1682);
			prefixExpression();
			setState(1684);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,218,_ctx) ) {
			case 1:
				{
				setState(1683);
				binaryExpressions();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrefixExpressionContext extends ParserRuleContext {
		public PrefixOperatorContext prefixOperator() {
			return getRuleContext(PrefixOperatorContext.class,0);
		}
		public PostfixExpressionContext postfixExpression() {
			return getRuleContext(PostfixExpressionContext.class,0);
		}
		public PrefixExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefixExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrefixExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrefixExpressionContext prefixExpression() throws RecognitionException {
		PrefixExpressionContext _localctx = new PrefixExpressionContext(_ctx, getState());
		enterRule(_localctx, 206, RULE_prefixExpression);
		try {
			setState(1690);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,219,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1686);
				prefixOperator();
				setState(1687);
				postfixExpression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1689);
				postfixExpression();
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

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExpressionContext extends ParserRuleContext {
		public BinaryOperatorContext binaryOperator() {
			return getRuleContext(BinaryOperatorContext.class,0);
		}
		public PrefixExpressionContext prefixExpression() {
			return getRuleContext(PrefixExpressionContext.class,0);
		}
		public ConditionalOperatorContext conditionalOperator() {
			return getRuleContext(ConditionalOperatorContext.class,0);
		}
		public InOperatorContext inOperator() {
			return getRuleContext(InOperatorContext.class,0);
		}
		public IfOperatorContext ifOperator() {
			return getRuleContext(IfOperatorContext.class,0);
		}
		public InfixCallOperatorContext infixCallOperator() {
			return getRuleContext(InfixCallOperatorContext.class,0);
		}
		public BinaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitBinaryExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BinaryExpressionContext binaryExpression() throws RecognitionException {
		BinaryExpressionContext _localctx = new BinaryExpressionContext(_ctx, getState());
		enterRule(_localctx, 208, RULE_binaryExpression);
		try {
			setState(1707);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,220,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1692);
				binaryOperator();
				setState(1693);
				prefixExpression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1695);
				conditionalOperator();
				setState(1696);
				prefixExpression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1698);
				inOperator();
				setState(1699);
				prefixExpression();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1701);
				ifOperator();
				setState(1702);
				prefixExpression();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1704);
				infixCallOperator();
				setState(1705);
				prefixExpression();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrefixCallOperatorContext extends ParserRuleContext {
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public PrefixCallOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefixCallOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrefixCallOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrefixCallOperatorContext prefixCallOperator() throws RecognitionException {
		PrefixCallOperatorContext _localctx = new PrefixCallOperatorContext(_ctx, getState());
		enterRule(_localctx, 210, RULE_prefixCallOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1709);
			labelIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class InfixCallOperatorContext extends ParserRuleContext {
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public InfixCallOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_infixCallOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInfixCallOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InfixCallOperatorContext infixCallOperator() throws RecognitionException {
		InfixCallOperatorContext _localctx = new InfixCallOperatorContext(_ctx, getState());
		enterRule(_localctx, 212, RULE_infixCallOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1711);
			match(VALUE_IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExpressionsContext extends ParserRuleContext {
		public List<BinaryExpressionContext> binaryExpression() {
			return getRuleContexts(BinaryExpressionContext.class);
		}
		public BinaryExpressionContext binaryExpression(int i) {
			return getRuleContext(BinaryExpressionContext.class,i);
		}
		public BinaryExpressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryExpressions; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitBinaryExpressions(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BinaryExpressionsContext binaryExpressions() throws RecognitionException {
		BinaryExpressionsContext _localctx = new BinaryExpressionsContext(_ctx, getState());
		enterRule(_localctx, 214, RULE_binaryExpressions);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1714); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(1713);
					binaryExpression();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1716); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,221,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InOperatorContext extends ParserRuleContext {
		public TerminalNode BANG() { return getToken(MojoParser.BANG, 0); }
		public TerminalNode KEYWORD_IN() { return getToken(MojoParser.KEYWORD_IN, 0); }
		public InOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitInOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InOperatorContext inOperator() throws RecognitionException {
		InOperatorContext _localctx = new InOperatorContext(_ctx, getState());
		enterRule(_localctx, 216, RULE_inOperator);
		try {
			setState(1721);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BANG:
				enterOuterAlt(_localctx, 1);
				{
				setState(1718);
				match(BANG);
				setState(1719);
				match(KEYWORD_IN);
				}
				break;
			case KEYWORD_IN:
				enterOuterAlt(_localctx, 2);
				{
				setState(1720);
				match(KEYWORD_IN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionalOperatorContext extends ParserRuleContext {
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public ConditionalOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionalOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitConditionalOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConditionalOperatorContext conditionalOperator() throws RecognitionException {
		ConditionalOperatorContext _localctx = new ConditionalOperatorContext(_ctx, getState());
		enterRule(_localctx, 218, RULE_conditionalOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1723);
			match(QUESTION);
			setState(1724);
			expression();
			setState(1725);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfOperatorContext extends ParserRuleContext {
		public TerminalNode KEYWORD_IF() { return getToken(MojoParser.KEYWORD_IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode KEYWORD_ELSE() { return getToken(MojoParser.KEYWORD_ELSE, 0); }
		public IfOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIfOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IfOperatorContext ifOperator() throws RecognitionException {
		IfOperatorContext _localctx = new IfOperatorContext(_ctx, getState());
		enterRule(_localctx, 220, RULE_ifOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1727);
			match(KEYWORD_IF);
			setState(1728);
			expression();
			setState(1729);
			match(KEYWORD_ELSE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeCastingOperatorContext extends ParserRuleContext {
		public TerminalNode BANG() { return getToken(MojoParser.BANG, 0); }
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode KEYWORD_AS() { return getToken(MojoParser.KEYWORD_AS, 0); }
		public TypeCastingOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeCastingOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeCastingOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeCastingOperatorContext typeCastingOperator() throws RecognitionException {
		TypeCastingOperatorContext _localctx = new TypeCastingOperatorContext(_ctx, getState());
		enterRule(_localctx, 222, RULE_typeCastingOperator);
		try {
			setState(1742);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_IS:
			case BANG:
				enterOuterAlt(_localctx, 1);
				{
				setState(1734);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case BANG:
					{
					setState(1731);
					match(BANG);
					setState(1732);
					match(KEYWORD_IS);
					}
					break;
				case KEYWORD_IS:
					{
					setState(1733);
					match(KEYWORD_IS);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1738);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,224,_ctx) ) {
				case 1:
					{
					setState(1736);
					type_(0);
					}
					break;
				case 2:
					{
					setState(1737);
					pattern(0);
					}
					break;
				}
				}
				break;
			case KEYWORD_AS:
				enterOuterAlt(_localctx, 2);
				{
				setState(1740);
				match(KEYWORD_AS);
				setState(1741);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionContext extends ParserRuleContext {
		public LiteralExpressionContext literalExpression() {
			return getRuleContext(LiteralExpressionContext.class,0);
		}
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public GenericArgumentClauseContext genericArgumentClause() {
			return getRuleContext(GenericArgumentClauseContext.class,0);
		}
		public TypeIdentifierContext typeIdentifier() {
			return getRuleContext(TypeIdentifierContext.class,0);
		}
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public ClosureExpressionContext closureExpression() {
			return getRuleContext(ClosureExpressionContext.class,0);
		}
		public TupleLiteralExpressionContext tupleLiteralExpression() {
			return getRuleContext(TupleLiteralExpressionContext.class,0);
		}
		public ParenthesizedExpressionContext parenthesizedExpression() {
			return getRuleContext(ParenthesizedExpressionContext.class,0);
		}
		public ImplicitMemberExpressionContext implicitMemberExpression() {
			return getRuleContext(ImplicitMemberExpressionContext.class,0);
		}
		public WildcardExpressionContext wildcardExpression() {
			return getRuleContext(WildcardExpressionContext.class,0);
		}
		public StructConstructionExpressionContext structConstructionExpression() {
			return getRuleContext(StructConstructionExpressionContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public PrimaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrimaryExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrimaryExpressionContext primaryExpression() throws RecognitionException {
		PrimaryExpressionContext _localctx = new PrimaryExpressionContext(_ctx, getState());
		enterRule(_localctx, 224, RULE_primaryExpression);
		try {
			setState(1762);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,228,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1744);
				literalExpression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1745);
				declarationIdentifier();
				setState(1747);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,226,_ctx) ) {
				case 1:
					{
					setState(1746);
					genericArgumentClause();
					}
					break;
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1749);
				typeIdentifier();
				setState(1750);
				match(DOT);
				setState(1751);
				declarationIdentifier();
				setState(1753);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,227,_ctx) ) {
				case 1:
					{
					setState(1752);
					genericArgumentClause();
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1755);
				closureExpression();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1756);
				tupleLiteralExpression();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1757);
				parenthesizedExpression();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1758);
				implicitMemberExpression();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(1759);
				wildcardExpression();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(1760);
				structConstructionExpression();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(1761);
				match(ELLIPSIS);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralExpressionContext extends ParserRuleContext {
		public NumericOperatorLiteralContext numericOperatorLiteral() {
			return getRuleContext(NumericOperatorLiteralContext.class,0);
		}
		public StringOperatorLiteralContext stringOperatorLiteral() {
			return getRuleContext(StringOperatorLiteralContext.class,0);
		}
		public StructLiteralContext structLiteral() {
			return getRuleContext(StructLiteralContext.class,0);
		}
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public MapLiteralContext mapLiteral() {
			return getRuleContext(MapLiteralContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public LiteralExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitLiteralExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LiteralExpressionContext literalExpression() throws RecognitionException {
		LiteralExpressionContext _localctx = new LiteralExpressionContext(_ctx, getState());
		enterRule(_localctx, 226, RULE_literalExpression);
		try {
			setState(1771);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,229,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1764);
				numericOperatorLiteral();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1765);
				stringOperatorLiteral();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1766);
				structLiteral();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1767);
				literal();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1768);
				arrayLiteral();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1769);
				mapLiteral();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1770);
				objectLiteral();
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

	@SuppressWarnings("CheckReturnValue")
	public static class NumericOperatorLiteralContext extends ParserRuleContext {
		public NumericLiteralContext numericLiteral() {
			return getRuleContext(NumericLiteralContext.class,0);
		}
		public SuffixLiteralOperatorContext suffixLiteralOperator() {
			return getRuleContext(SuffixLiteralOperatorContext.class,0);
		}
		public NumericOperatorLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numericOperatorLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitNumericOperatorLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NumericOperatorLiteralContext numericOperatorLiteral() throws RecognitionException {
		NumericOperatorLiteralContext _localctx = new NumericOperatorLiteralContext(_ctx, getState());
		enterRule(_localctx, 228, RULE_numericOperatorLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1773);
			numericLiteral();
			setState(1774);
			suffixLiteralOperator();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StringOperatorLiteralContext extends ParserRuleContext {
		public PrefixLiteralOperatorContext prefixLiteralOperator() {
			return getRuleContext(PrefixLiteralOperatorContext.class,0);
		}
		public StringLiteralContext stringLiteral() {
			return getRuleContext(StringLiteralContext.class,0);
		}
		public SuffixLiteralOperatorContext suffixLiteralOperator() {
			return getRuleContext(SuffixLiteralOperatorContext.class,0);
		}
		public StringOperatorLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringOperatorLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStringOperatorLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StringOperatorLiteralContext stringOperatorLiteral() throws RecognitionException {
		StringOperatorLiteralContext _localctx = new StringOperatorLiteralContext(_ctx, getState());
		enterRule(_localctx, 230, RULE_stringOperatorLiteral);
		try {
			setState(1785);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(1776);
				prefixLiteralOperator();
				setState(1777);
				if (!(_input.index() > 0 && _input.get(_input.index()-1).getType() != WS)) throw new FailedPredicateException(this, "_input.index() > 0 && _input.get(_input.index()-1).getType() != WS");
				setState(1778);
				stringLiteral();
				setState(1780);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,230,_ctx) ) {
				case 1:
					{
					setState(1779);
					suffixLiteralOperator();
					}
					break;
				}
				}
				break;
			case STATIC_STRING_LITERAL:
			case INTERPOLATED_STRING_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(1782);
				stringLiteral();
				setState(1783);
				suffixLiteralOperator();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SuffixLiteralOperatorContext extends ParserRuleContext {
		public TerminalNode TYPE_IDENTIFIER() { return getToken(MojoParser.TYPE_IDENTIFIER, 0); }
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public SuffixLiteralOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_suffixLiteralOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitSuffixLiteralOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SuffixLiteralOperatorContext suffixLiteralOperator() throws RecognitionException {
		SuffixLiteralOperatorContext _localctx = new SuffixLiteralOperatorContext(_ctx, getState());
		enterRule(_localctx, 232, RULE_suffixLiteralOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1787);
			if (!(_input.index() > 0 && _input.get(_input.index()-1).getType() != WS)) throw new FailedPredicateException(this, "_input.index() > 0 && _input.get(_input.index()-1).getType() != WS");
			setState(1788);
			_la = _input.LA(1);
			if ( !(_la==TYPE_IDENTIFIER || _la==VALUE_IDENTIFIER) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrefixLiteralOperatorContext extends ParserRuleContext {
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public PrefixLiteralOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefixLiteralOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrefixLiteralOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrefixLiteralOperatorContext prefixLiteralOperator() throws RecognitionException {
		PrefixLiteralOperatorContext _localctx = new PrefixLiteralOperatorContext(_ctx, getState());
		enterRule(_localctx, 234, RULE_prefixLiteralOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1790);
			match(VALUE_IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayLiteralContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public ArrayLiteralItemsContext arrayLiteralItems() {
			return getRuleContext(ArrayLiteralItemsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ArrayLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayLiteralContext arrayLiteral() throws RecognitionException {
		ArrayLiteralContext _localctx = new ArrayLiteralContext(_ctx, getState());
		enterRule(_localctx, 236, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1792);
			match(LBRACK);
			setState(1800);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,233,_ctx) ) {
			case 1:
				{
				setState(1796);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1793);
					match(EOL);
					}
					}
					setState(1798);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1799);
				arrayLiteralItems();
				}
				break;
			}
			setState(1805);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1802);
				match(EOL);
				}
				}
				setState(1807);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1808);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayLiteralItemsContext extends ParserRuleContext {
		public List<ArrayLiteralItemContext> arrayLiteralItem() {
			return getRuleContexts(ArrayLiteralItemContext.class);
		}
		public ArrayLiteralItemContext arrayLiteralItem(int i) {
			return getRuleContext(ArrayLiteralItemContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ArrayLiteralItemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteralItems; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayLiteralItems(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayLiteralItemsContext arrayLiteralItems() throws RecognitionException {
		ArrayLiteralItemsContext _localctx = new ArrayLiteralItemsContext(_ctx, getState());
		enterRule(_localctx, 238, RULE_arrayLiteralItems);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1810);
			arrayLiteralItem();
			setState(1822);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,236,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1811);
					eov();
					setState(1815);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1812);
						match(EOL);
						}
						}
						setState(1817);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1818);
					arrayLiteralItem();
					}
					} 
				}
				setState(1824);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,236,_ctx);
			}
			setState(1826);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,237,_ctx) ) {
			case 1:
				{
				setState(1825);
				eov();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayLiteralItemContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public ArrayLiteralItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteralItem; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayLiteralItem(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayLiteralItemContext arrayLiteralItem() throws RecognitionException {
		ArrayLiteralItemContext _localctx = new ArrayLiteralItemContext(_ctx, getState());
		enterRule(_localctx, 240, RULE_arrayLiteralItem);
		try {
			setState(1830);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,238,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1828);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1829);
				match(ELLIPSIS);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MapLiteralContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public MapLiteralItemsContext mapLiteralItems() {
			return getRuleContext(MapLiteralItemsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MapLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMapLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MapLiteralContext mapLiteral() throws RecognitionException {
		MapLiteralContext _localctx = new MapLiteralContext(_ctx, getState());
		enterRule(_localctx, 242, RULE_mapLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1832);
			match(LCURLY);
			setState(1840);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,240,_ctx) ) {
			case 1:
				{
				setState(1836);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1833);
					match(EOL);
					}
					}
					setState(1838);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1839);
				mapLiteralItems();
				}
				break;
			}
			setState(1845);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1842);
				match(EOL);
				}
				}
				setState(1847);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1848);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MapLiteralItemsContext extends ParserRuleContext {
		public List<MapLiteralItemContext> mapLiteralItem() {
			return getRuleContexts(MapLiteralItemContext.class);
		}
		public MapLiteralItemContext mapLiteralItem(int i) {
			return getRuleContext(MapLiteralItemContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MapLiteralItemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapLiteralItems; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMapLiteralItems(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MapLiteralItemsContext mapLiteralItems() throws RecognitionException {
		MapLiteralItemsContext _localctx = new MapLiteralItemsContext(_ctx, getState());
		enterRule(_localctx, 244, RULE_mapLiteralItems);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(1850);
			mapLiteralItem();
			setState(1862);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,243,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1851);
					eov();
					setState(1855);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1852);
						match(EOL);
						}
						}
						setState(1857);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1858);
					mapLiteralItem();
					}
					} 
				}
				setState(1864);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,243,_ctx);
			}
			setState(1866);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,244,_ctx) ) {
			case 1:
				{
				setState(1865);
				eov();
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class MapLiteralItemContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public StringLiteralContext stringLiteral() {
			return getRuleContext(StringLiteralContext.class,0);
		}
		public IntegerLiteralContext integerLiteral() {
			return getRuleContext(IntegerLiteralContext.class,0);
		}
		public MapLiteralItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapLiteralItem; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMapLiteralItem(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MapLiteralItemContext mapLiteralItem() throws RecognitionException {
		MapLiteralItemContext _localctx = new MapLiteralItemContext(_ctx, getState());
		enterRule(_localctx, 246, RULE_mapLiteralItem);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1870);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STATIC_STRING_LITERAL:
			case INTERPOLATED_STRING_LITERAL:
				{
				setState(1868);
				stringLiteral();
				}
				break;
			case BINARY_LITERAL:
			case OCTAL_LITERAL:
			case DECIMAL_LITERAL:
			case PURE_DECIMAL_DIGITS:
			case HEXADECIMAL_LITERAL:
				{
				setState(1869);
				integerLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(1872);
			match(COLON);
			setState(1873);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ObjectLiteralContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public ObjectLiteralItemsContext objectLiteralItems() {
			return getRuleContext(ObjectLiteralItemsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ObjectLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitObjectLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjectLiteralContext objectLiteral() throws RecognitionException {
		ObjectLiteralContext _localctx = new ObjectLiteralContext(_ctx, getState());
		enterRule(_localctx, 248, RULE_objectLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1875);
			match(LCURLY);
			setState(1883);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,247,_ctx) ) {
			case 1:
				{
				setState(1879);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1876);
					match(EOL);
					}
					}
					setState(1881);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1882);
				objectLiteralItems();
				}
				break;
			}
			setState(1888);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1885);
				match(EOL);
				}
				}
				setState(1890);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1891);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ObjectLiteralItemsContext extends ParserRuleContext {
		public List<ObjectLiteralItemContext> objectLiteralItem() {
			return getRuleContexts(ObjectLiteralItemContext.class);
		}
		public ObjectLiteralItemContext objectLiteralItem(int i) {
			return getRuleContext(ObjectLiteralItemContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ObjectLiteralItemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectLiteralItems; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitObjectLiteralItems(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjectLiteralItemsContext objectLiteralItems() throws RecognitionException {
		ObjectLiteralItemsContext _localctx = new ObjectLiteralItemsContext(_ctx, getState());
		enterRule(_localctx, 250, RULE_objectLiteralItems);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1893);
			objectLiteralItem();
			setState(1905);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,250,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1894);
					eov();
					setState(1898);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1895);
						match(EOL);
						}
						}
						setState(1900);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1901);
					objectLiteralItem();
					}
					} 
				}
				setState(1907);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,250,_ctx);
			}
			setState(1909);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,251,_ctx) ) {
			case 1:
				{
				setState(1908);
				eov();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ObjectLiteralItemContext extends ParserRuleContext {
		public PathIdentifierContext pathIdentifier() {
			return getRuleContext(PathIdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ObjectLiteralItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectLiteralItem; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitObjectLiteralItem(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjectLiteralItemContext objectLiteralItem() throws RecognitionException {
		ObjectLiteralItemContext _localctx = new ObjectLiteralItemContext(_ctx, getState());
		enterRule(_localctx, 252, RULE_objectLiteralItem);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1911);
			pathIdentifier();
			setState(1914);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(1912);
				match(COLON);
				setState(1913);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructLiteralContext extends ParserRuleContext {
		public TypeIdentifierContext typeIdentifier() {
			return getRuleContext(TypeIdentifierContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public StructLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructLiteralContext structLiteral() throws RecognitionException {
		StructLiteralContext _localctx = new StructLiteralContext(_ctx, getState());
		enterRule(_localctx, 254, RULE_structLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1916);
			typeIdentifier();
			setState(1917);
			objectLiteral();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructConstructionExpressionContext extends ParserRuleContext {
		public TypeIdentifierContext typeIdentifier() {
			return getRuleContext(TypeIdentifierContext.class,0);
		}
		public FunctionCallSuffixContext functionCallSuffix() {
			return getRuleContext(FunctionCallSuffixContext.class,0);
		}
		public StructConstructionExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structConstructionExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStructConstructionExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructConstructionExpressionContext structConstructionExpression() throws RecognitionException {
		StructConstructionExpressionContext _localctx = new StructConstructionExpressionContext(_ctx, getState());
		enterRule(_localctx, 256, RULE_structConstructionExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1919);
			typeIdentifier();
			setState(1920);
			functionCallSuffix();
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatchExprSuffixContext extends ParserRuleContext {
		public TerminalNode KEYWORD_MATCH() { return getToken(MojoParser.KEYWORD_MATCH, 0); }
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MatchExprCasesContext matchExprCases() {
			return getRuleContext(MatchExprCasesContext.class,0);
		}
		public MatchExprSuffixContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchExprSuffix; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMatchExprSuffix(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MatchExprSuffixContext matchExprSuffix() throws RecognitionException {
		MatchExprSuffixContext _localctx = new MatchExprSuffixContext(_ctx, getState());
		enterRule(_localctx, 258, RULE_matchExprSuffix);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1922);
			match(KEYWORD_MATCH);
			setState(1926);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1923);
				match(EOL);
				}
				}
				setState(1928);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1929);
			match(LCURLY);
			setState(1937);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,255,_ctx) ) {
			case 1:
				{
				setState(1933);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1930);
					match(EOL);
					}
					}
					setState(1935);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1936);
				matchExprCases();
				}
				break;
			}
			setState(1942);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1939);
				match(EOL);
				}
				}
				setState(1944);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1945);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatchExprCasesContext extends ParserRuleContext {
		public List<MatchExprCaseContext> matchExprCase() {
			return getRuleContexts(MatchExprCaseContext.class);
		}
		public MatchExprCaseContext matchExprCase(int i) {
			return getRuleContext(MatchExprCaseContext.class,i);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MatchExprCasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchExprCases; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMatchExprCases(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MatchExprCasesContext matchExprCases() throws RecognitionException {
		MatchExprCasesContext _localctx = new MatchExprCasesContext(_ctx, getState());
		enterRule(_localctx, 260, RULE_matchExprCases);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1947);
			matchExprCase();
			setState(1959);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,258,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1948);
					eos();
					setState(1952);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1949);
						match(EOL);
						}
						}
						setState(1954);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1955);
					matchExprCase();
					}
					} 
				}
				setState(1961);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,258,_ctx);
			}
			setState(1963);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,259,_ctx) ) {
			case 1:
				{
				setState(1962);
				eos();
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatchExprCaseContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode RIGHT_RIGHT_ARROWS() { return getToken(MojoParser.RIGHT_RIGHT_ARROWS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public MatchExprCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchExprCase; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMatchExprCase(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MatchExprCaseContext matchExprCase() throws RecognitionException {
		MatchExprCaseContext _localctx = new MatchExprCaseContext(_ctx, getState());
		enterRule(_localctx, 262, RULE_matchExprCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1965);
			pattern(0);
			setState(1969);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1966);
				match(EOL);
				}
				}
				setState(1971);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1972);
			match(RIGHT_RIGHT_ARROWS);
			setState(1976);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1973);
				match(EOL);
				}
				}
				setState(1978);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1979);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ClosureExpressionContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ClosureParametersContext closureParameters() {
			return getRuleContext(ClosureParametersContext.class,0);
		}
		public TerminalNode RIGHT_ARROW() { return getToken(MojoParser.RIGHT_ARROW, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public ClosureExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closureExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitClosureExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ClosureExpressionContext closureExpression() throws RecognitionException {
		ClosureExpressionContext _localctx = new ClosureExpressionContext(_ctx, getState());
		enterRule(_localctx, 264, RULE_closureExpression);
		int _la;
		try {
			setState(2024);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,268,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1981);
				match(LCURLY);
				setState(1985);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1982);
					match(EOL);
					}
					}
					setState(1987);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1988);
				statements();
				setState(1992);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1989);
					match(EOL);
					}
					}
					setState(1994);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1995);
				match(RCURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1997);
				match(LCURLY);
				setState(1998);
				closureParameters();
				setState(2002);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1999);
					match(EOL);
					}
					}
					setState(2004);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2005);
				match(RIGHT_ARROW);
				setState(2013);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,266,_ctx) ) {
				case 1:
					{
					setState(2009);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(2006);
						match(EOL);
						}
						}
						setState(2011);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(2012);
					type_(0);
					}
					break;
				}
				setState(2018);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(2015);
					match(EOL);
					}
					}
					setState(2020);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2021);
				statements();
				setState(2022);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ClosureParametersContext extends ParserRuleContext {
		public List<ClosureParameterContext> closureParameter() {
			return getRuleContexts(ClosureParameterContext.class);
		}
		public ClosureParameterContext closureParameter(int i) {
			return getRuleContext(ClosureParameterContext.class,i);
		}
		public List<EovContext> eov() {
			return getRuleContexts(EovContext.class);
		}
		public EovContext eov(int i) {
			return getRuleContext(EovContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ClosureParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closureParameters; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitClosureParameters(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ClosureParametersContext closureParameters() throws RecognitionException {
		ClosureParametersContext _localctx = new ClosureParametersContext(_ctx, getState());
		enterRule(_localctx, 266, RULE_closureParameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2026);
			closureParameter();
			setState(2038);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,270,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2027);
					eov();
					setState(2031);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(2028);
						match(EOL);
						}
						}
						setState(2033);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(2034);
					closureParameter();
					}
					} 
				}
				setState(2040);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,270,_ctx);
			}
			setState(2042);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,271,_ctx) ) {
			case 1:
				{
				setState(2041);
				eov();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ClosureParameterContext extends ParserRuleContext {
		public FunctionParameterContext functionParameter() {
			return getRuleContext(FunctionParameterContext.class,0);
		}
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public ClosureParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closureParameter; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitClosureParameter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ClosureParameterContext closureParameter() throws RecognitionException {
		ClosureParameterContext _localctx = new ClosureParameterContext(_ctx, getState());
		enterRule(_localctx, 268, RULE_closureParameter);
		try {
			setState(2046);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,272,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2044);
				functionParameter();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2045);
				labelIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ImplicitMemberExpressionContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public ImplicitMemberExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implicitMemberExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitImplicitMemberExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImplicitMemberExpressionContext implicitMemberExpression() throws RecognitionException {
		ImplicitMemberExpressionContext _localctx = new ImplicitMemberExpressionContext(_ctx, getState());
		enterRule(_localctx, 270, RULE_implicitMemberExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2048);
			match(DOT);
			setState(2049);
			labelIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParenthesizedExpressionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public ParenthesizedExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parenthesizedExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitParenthesizedExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ParenthesizedExpressionContext parenthesizedExpression() throws RecognitionException {
		ParenthesizedExpressionContext _localctx = new ParenthesizedExpressionContext(_ctx, getState());
		enterRule(_localctx, 272, RULE_parenthesizedExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2051);
			match(LPAREN);
			setState(2055);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(2052);
				match(EOL);
				}
				}
				setState(2057);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			{
			setState(2058);
			expression();
			}
			setState(2062);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(2059);
				match(EOL);
				}
				}
				setState(2064);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(2065);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TupleLiteralExpressionContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public List<TupleElementContext> tupleElement() {
			return getRuleContexts(TupleElementContext.class);
		}
		public TupleElementContext tupleElement(int i) {
			return getRuleContext(TupleElementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public TupleLiteralExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleLiteralExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTupleLiteralExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TupleLiteralExpressionContext tupleLiteralExpression() throws RecognitionException {
		TupleLiteralExpressionContext _localctx = new TupleLiteralExpressionContext(_ctx, getState());
		enterRule(_localctx, 274, RULE_tupleLiteralExpression);
		int _la;
		try {
			setState(2079);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,276,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2067);
				match(LPAREN);
				setState(2068);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2069);
				match(LPAREN);
				setState(2070);
				tupleElement();
				setState(2073); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(2071);
					match(COMMA);
					setState(2072);
					tupleElement();
					}
					}
					setState(2075); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				setState(2077);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TupleElementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public TupleElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleElement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTupleElement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TupleElementContext tupleElement() throws RecognitionException {
		TupleElementContext _localctx = new TupleElementContext(_ctx, getState());
		enterRule(_localctx, 276, RULE_tupleElement);
		try {
			setState(2087);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,277,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2081);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2082);
				labelIdentifier();
				setState(2083);
				match(COLON);
				setState(2084);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(2086);
				match(ELLIPSIS);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WildcardExpressionContext extends ParserRuleContext {
		public TerminalNode UNDERSCORE() { return getToken(MojoParser.UNDERSCORE, 0); }
		public WildcardExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_wildcardExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitWildcardExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final WildcardExpressionContext wildcardExpression() throws RecognitionException {
		WildcardExpressionContext _localctx = new WildcardExpressionContext(_ctx, getState());
		enterRule(_localctx, 278, RULE_wildcardExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2089);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PostfixExpressionContext extends ParserRuleContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public List<SuffixExpressionContext> suffixExpression() {
			return getRuleContexts(SuffixExpressionContext.class);
		}
		public SuffixExpressionContext suffixExpression(int i) {
			return getRuleContext(SuffixExpressionContext.class,i);
		}
		public PostfixOperatorContext postfixOperator() {
			return getRuleContext(PostfixOperatorContext.class,0);
		}
		public PostfixExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfixExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPostfixExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PostfixExpressionContext postfixExpression() throws RecognitionException {
		PostfixExpressionContext _localctx = new PostfixExpressionContext(_ctx, getState());
		enterRule(_localctx, 280, RULE_postfixExpression);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2091);
			primaryExpression();
			setState(2095);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,278,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2092);
					suffixExpression();
					}
					} 
				}
				setState(2097);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,278,_ctx);
			}
			setState(2099);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,279,_ctx) ) {
			case 1:
				{
				setState(2098);
				postfixOperator();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SuffixExpressionContext extends ParserRuleContext {
		public FunctionCallSuffixContext functionCallSuffix() {
			return getRuleContext(FunctionCallSuffixContext.class,0);
		}
		public ExplicitMemberSuffixContext explicitMemberSuffix() {
			return getRuleContext(ExplicitMemberSuffixContext.class,0);
		}
		public SubscriptSuffixContext subscriptSuffix() {
			return getRuleContext(SubscriptSuffixContext.class,0);
		}
		public MatchExprSuffixContext matchExprSuffix() {
			return getRuleContext(MatchExprSuffixContext.class,0);
		}
		public TypeCastingOperatorContext typeCastingOperator() {
			return getRuleContext(TypeCastingOperatorContext.class,0);
		}
		public SuffixExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_suffixExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitSuffixExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SuffixExpressionContext suffixExpression() throws RecognitionException {
		SuffixExpressionContext _localctx = new SuffixExpressionContext(_ctx, getState());
		enterRule(_localctx, 282, RULE_suffixExpression);
		try {
			setState(2106);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LCURLY:
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(2101);
				functionCallSuffix();
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2102);
				explicitMemberSuffix();
				}
				break;
			case LBRACK:
				enterOuterAlt(_localctx, 3);
				{
				setState(2103);
				subscriptSuffix();
				}
				break;
			case KEYWORD_MATCH:
				enterOuterAlt(_localctx, 4);
				{
				setState(2104);
				matchExprSuffix();
				}
				break;
			case KEYWORD_AS:
			case KEYWORD_IS:
			case BANG:
				enterOuterAlt(_localctx, 5);
				{
				setState(2105);
				typeCastingOperator();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExplicitMemberSuffixContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public TerminalNode DECIMAL_LITERAL() { return getToken(MojoParser.DECIMAL_LITERAL, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public GenericArgumentClauseContext genericArgumentClause() {
			return getRuleContext(GenericArgumentClauseContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public ArgumentNamesContext argumentNames() {
			return getRuleContext(ArgumentNamesContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public ExplicitMemberSuffixContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_explicitMemberSuffix; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitExplicitMemberSuffix(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExplicitMemberSuffixContext explicitMemberSuffix() throws RecognitionException {
		ExplicitMemberSuffixContext _localctx = new ExplicitMemberSuffixContext(_ctx, getState());
		enterRule(_localctx, 284, RULE_explicitMemberSuffix);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2108);
			match(DOT);
			setState(2118);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECIMAL_LITERAL:
				{
				setState(2109);
				match(DECIMAL_LITERAL);
				}
				break;
			case VALUE_IDENTIFIER:
			case IMPLICIT_PARAMETER_NAME:
				{
				setState(2110);
				identifier();
				setState(2116);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,281,_ctx) ) {
				case 1:
					{
					setState(2111);
					genericArgumentClause();
					}
					break;
				case 2:
					{
					setState(2112);
					match(LPAREN);
					setState(2113);
					argumentNames();
					setState(2114);
					match(RPAREN);
					}
					break;
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SubscriptSuffixContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public FunctionCallArgumentsContext functionCallArguments() {
			return getRuleContext(FunctionCallArgumentsContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public SubscriptSuffixContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subscriptSuffix; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitSubscriptSuffix(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SubscriptSuffixContext subscriptSuffix() throws RecognitionException {
		SubscriptSuffixContext _localctx = new SubscriptSuffixContext(_ctx, getState());
		enterRule(_localctx, 286, RULE_subscriptSuffix);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2120);
			match(LBRACK);
			setState(2121);
			functionCallArguments();
			setState(2122);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallSuffixContext extends ParserRuleContext {
		public TrailingClosuresContext trailingClosures() {
			return getRuleContext(TrailingClosuresContext.class,0);
		}
		public FunctionCallArgumentClauseContext functionCallArgumentClause() {
			return getRuleContext(FunctionCallArgumentClauseContext.class,0);
		}
		public FunctionCallSuffixContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCallSuffix; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionCallSuffix(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionCallSuffixContext functionCallSuffix() throws RecognitionException {
		FunctionCallSuffixContext _localctx = new FunctionCallSuffixContext(_ctx, getState());
		enterRule(_localctx, 288, RULE_functionCallSuffix);
		int _la;
		try {
			setState(2129);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,284,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2125);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LPAREN) {
					{
					setState(2124);
					functionCallArgumentClause();
					}
				}

				setState(2127);
				trailingClosures();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2128);
				functionCallArgumentClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallArgumentClauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FunctionCallArgumentsContext functionCallArguments() {
			return getRuleContext(FunctionCallArgumentsContext.class,0);
		}
		public FunctionCallArgumentClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCallArgumentClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionCallArgumentClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionCallArgumentClauseContext functionCallArgumentClause() throws RecognitionException {
		FunctionCallArgumentClauseContext _localctx = new FunctionCallArgumentClauseContext(_ctx, getState());
		enterRule(_localctx, 290, RULE_functionCallArgumentClause);
		int _la;
		try {
			setState(2155);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,288,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2131);
				match(LPAREN);
				setState(2135);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(2132);
					match(EOL);
					}
					}
					setState(2137);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2138);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2139);
				match(LPAREN);
				setState(2143);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(2140);
					match(EOL);
					}
					}
					setState(2145);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2146);
				functionCallArguments();
				setState(2150);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(2147);
					match(EOL);
					}
					}
					setState(2152);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2153);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallArgumentsContext extends ParserRuleContext {
		public List<FunctionCallArgumentContext> functionCallArgument() {
			return getRuleContexts(FunctionCallArgumentContext.class);
		}
		public FunctionCallArgumentContext functionCallArgument(int i) {
			return getRuleContext(FunctionCallArgumentContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(MojoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MojoParser.COMMA, i);
		}
		public FunctionCallArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCallArguments; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionCallArguments(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionCallArgumentsContext functionCallArguments() throws RecognitionException {
		FunctionCallArgumentsContext _localctx = new FunctionCallArgumentsContext(_ctx, getState());
		enterRule(_localctx, 292, RULE_functionCallArguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2157);
			functionCallArgument();
			setState(2162);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(2158);
				match(COMMA);
				setState(2159);
				functionCallArgument();
				}
				}
				setState(2164);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallArgumentContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public FunctionCallArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCallArgument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionCallArgument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionCallArgumentContext functionCallArgument() throws RecognitionException {
		FunctionCallArgumentContext _localctx = new FunctionCallArgumentContext(_ctx, getState());
		enterRule(_localctx, 294, RULE_functionCallArgument);
		try {
			setState(2170);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,290,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2165);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2166);
				labelIdentifier();
				setState(2167);
				match(COLON);
				setState(2168);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TrailingClosuresContext extends ParserRuleContext {
		public ClosureExpressionContext closureExpression() {
			return getRuleContext(ClosureExpressionContext.class,0);
		}
		public LabeledTrailingClosuresContext labeledTrailingClosures() {
			return getRuleContext(LabeledTrailingClosuresContext.class,0);
		}
		public TrailingClosuresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_trailingClosures; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTrailingClosures(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TrailingClosuresContext trailingClosures() throws RecognitionException {
		TrailingClosuresContext _localctx = new TrailingClosuresContext(_ctx, getState());
		enterRule(_localctx, 296, RULE_trailingClosures);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2172);
			closureExpression();
			setState(2174);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,291,_ctx) ) {
			case 1:
				{
				setState(2173);
				labeledTrailingClosures();
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

	@SuppressWarnings("CheckReturnValue")
	public static class LabeledTrailingClosuresContext extends ParserRuleContext {
		public List<LabeledTrailingClosureContext> labeledTrailingClosure() {
			return getRuleContexts(LabeledTrailingClosureContext.class);
		}
		public LabeledTrailingClosureContext labeledTrailingClosure(int i) {
			return getRuleContext(LabeledTrailingClosureContext.class,i);
		}
		public LabeledTrailingClosuresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labeledTrailingClosures; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitLabeledTrailingClosures(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LabeledTrailingClosuresContext labeledTrailingClosures() throws RecognitionException {
		LabeledTrailingClosuresContext _localctx = new LabeledTrailingClosuresContext(_ctx, getState());
		enterRule(_localctx, 298, RULE_labeledTrailingClosures);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2177); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(2176);
					labeledTrailingClosure();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(2179); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,292,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LabeledTrailingClosureContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public ClosureExpressionContext closureExpression() {
			return getRuleContext(ClosureExpressionContext.class,0);
		}
		public LabeledTrailingClosureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labeledTrailingClosure; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitLabeledTrailingClosure(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LabeledTrailingClosureContext labeledTrailingClosure() throws RecognitionException {
		LabeledTrailingClosureContext _localctx = new LabeledTrailingClosureContext(_ctx, getState());
		enterRule(_localctx, 300, RULE_labeledTrailingClosure);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2181);
			identifier();
			setState(2182);
			match(COLON);
			setState(2183);
			closureExpression();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentNamesContext extends ParserRuleContext {
		public List<ArgumentNameContext> argumentName() {
			return getRuleContexts(ArgumentNameContext.class);
		}
		public ArgumentNameContext argumentName(int i) {
			return getRuleContext(ArgumentNameContext.class,i);
		}
		public ArgumentNamesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentNames; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArgumentNames(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArgumentNamesContext argumentNames() throws RecognitionException {
		ArgumentNamesContext _localctx = new ArgumentNamesContext(_ctx, getState());
		enterRule(_localctx, 302, RULE_argumentNames);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2185);
			argumentName();
			setState(2189);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 536866814L) != 0) || _la==VALUE_IDENTIFIER) {
				{
				{
				setState(2186);
				argumentName();
				}
				}
				setState(2191);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentNameContext extends ParserRuleContext {
		public LabelIdentifierContext labelIdentifier() {
			return getRuleContext(LabelIdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public ArgumentNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArgumentName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArgumentNameContext argumentName() throws RecognitionException {
		ArgumentNameContext _localctx = new ArgumentNameContext(_ctx, getState());
		enterRule(_localctx, 304, RULE_argumentName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2192);
			labelIdentifier();
			setState(2193);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Type_Context extends ParserRuleContext {
		public BasicTypeContext basicType() {
			return getRuleContext(BasicTypeContext.class,0);
		}
		public FunctionTypeContext functionType() {
			return getRuleContext(FunctionTypeContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode BANG() { return getToken(MojoParser.BANG, 0); }
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public TerminalNode STAR() { return getToken(MojoParser.STAR, 0); }
		public TerminalNode PLUS() { return getToken(MojoParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(MojoParser.MINUS, 0); }
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public Type_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitType_(this);
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
		int _startState = 306;
		enterRecursionRule(_localctx, 306, RULE_type_, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2198);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,294,_ctx) ) {
			case 1:
				{
				setState(2196);
				basicType(0);
				}
				break;
			case 2:
				{
				setState(2197);
				functionType();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(2214);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,296,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(2212);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,295,_ctx) ) {
					case 1:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2200);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(2201);
						match(BANG);
						}
						break;
					case 2:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2202);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(2203);
						match(QUESTION);
						}
						break;
					case 3:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2204);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(2205);
						match(STAR);
						}
						break;
					case 4:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2206);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(2207);
						match(PLUS);
						}
						break;
					case 5:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2208);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(2209);
						match(MINUS);
						}
						break;
					case 6:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(2210);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(2211);
						match(ELLIPSIS);
						}
						break;
					}
					} 
				}
				setState(2216);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,296,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BasicTypeContext extends ParserRuleContext {
		public BasicTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_basicType; }
	 
		public BasicTypeContext() { }
		public void copyFrom(BasicTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntersectionContext extends BasicTypeContext {
		public List<BasicTypeContext> basicType() {
			return getRuleContexts(BasicTypeContext.class);
		}
		public BasicTypeContext basicType(int i) {
			return getRuleContext(BasicTypeContext.class,i);
		}
		public TerminalNode AND() { return getToken(MojoParser.AND, 0); }
		public List<AttributesContext> attributes() {
			return getRuleContexts(AttributesContext.class);
		}
		public AttributesContext attributes(int i) {
			return getRuleContext(AttributesContext.class,i);
		}
		public List<FollowingDocumentContext> followingDocument() {
			return getRuleContexts(FollowingDocumentContext.class);
		}
		public FollowingDocumentContext followingDocument(int i) {
			return getRuleContext(FollowingDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public IntersectionContext(BasicTypeContext ctx) { copyFrom(ctx); }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIntersection(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimeContext extends BasicTypeContext {
		public PrimeTypeContext primeType() {
			return getRuleContext(PrimeTypeContext.class,0);
		}
		public PrimeContext(BasicTypeContext ctx) { copyFrom(ctx); }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrime(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnionContext extends BasicTypeContext {
		public List<BasicTypeContext> basicType() {
			return getRuleContexts(BasicTypeContext.class);
		}
		public BasicTypeContext basicType(int i) {
			return getRuleContext(BasicTypeContext.class,i);
		}
		public TerminalNode PIPE() { return getToken(MojoParser.PIPE, 0); }
		public List<AttributesContext> attributes() {
			return getRuleContexts(AttributesContext.class);
		}
		public AttributesContext attributes(int i) {
			return getRuleContext(AttributesContext.class,i);
		}
		public List<FollowingDocumentContext> followingDocument() {
			return getRuleContexts(FollowingDocumentContext.class);
		}
		public FollowingDocumentContext followingDocument(int i) {
			return getRuleContext(FollowingDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public UnionContext(BasicTypeContext ctx) { copyFrom(ctx); }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitUnion(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BasicTypeContext basicType() throws RecognitionException {
		return basicType(0);
	}

	private BasicTypeContext basicType(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		BasicTypeContext _localctx = new BasicTypeContext(_ctx, _parentState);
		BasicTypeContext _prevctx = _localctx;
		int _startState = 308;
		enterRecursionRule(_localctx, 308, RULE_basicType, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new PrimeContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(2218);
			primeType();
			}
			_ctx.stop = _input.LT(-1);
			setState(2284);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,310,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(2282);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,309,_ctx) ) {
					case 1:
						{
						_localctx = new UnionContext(new BasicTypeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_basicType);
						setState(2220);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(2222);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==AT) {
							{
							setState(2221);
							attributes();
							}
						}

						setState(2227);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==FOLLOWING_LINE_DOCUMENT) {
							{
							setState(2224);
							followingDocument();
							setState(2225);
							match(EOL);
							}
						}

						setState(2232);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(2229);
							match(EOL);
							}
							}
							setState(2234);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(2235);
						match(PIPE);
						setState(2239);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(2236);
							match(EOL);
							}
							}
							setState(2241);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(2242);
						basicType(0);
						setState(2244);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,301,_ctx) ) {
						case 1:
							{
							setState(2243);
							attributes();
							}
							break;
						}
						setState(2249);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,302,_ctx) ) {
						case 1:
							{
							setState(2246);
							followingDocument();
							setState(2247);
							match(EOL);
							}
							break;
						}
						}
						break;
					case 2:
						{
						_localctx = new IntersectionContext(new BasicTypeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_basicType);
						setState(2251);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(2253);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==AT) {
							{
							setState(2252);
							attributes();
							}
						}

						setState(2258);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==FOLLOWING_LINE_DOCUMENT) {
							{
							setState(2255);
							followingDocument();
							setState(2256);
							match(EOL);
							}
						}

						setState(2263);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(2260);
							match(EOL);
							}
							}
							setState(2265);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(2266);
						match(AND);
						setState(2270);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(2267);
							match(EOL);
							}
							}
							setState(2272);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(2273);
						basicType(0);
						setState(2275);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,307,_ctx) ) {
						case 1:
							{
							setState(2274);
							attributes();
							}
							break;
						}
						setState(2280);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,308,_ctx) ) {
						case 1:
							{
							setState(2277);
							followingDocument();
							setState(2278);
							match(EOL);
							}
							break;
						}
						}
						break;
					}
					} 
				}
				setState(2286);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,310,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrimeTypeContext extends ParserRuleContext {
		public ArrayTypeContext arrayType() {
			return getRuleContext(ArrayTypeContext.class,0);
		}
		public MapTypeContext mapType() {
			return getRuleContext(MapTypeContext.class,0);
		}
		public TupleTypeContext tupleType() {
			return getRuleContext(TupleTypeContext.class,0);
		}
		public TypeIdentifierContext typeIdentifier() {
			return getRuleContext(TypeIdentifierContext.class,0);
		}
		public PrimeTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primeType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrimeType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrimeTypeContext primeType() throws RecognitionException {
		PrimeTypeContext _localctx = new PrimeTypeContext(_ctx, getState());
		enterRule(_localctx, 310, RULE_primeType);
		try {
			setState(2291);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LBRACK:
				enterOuterAlt(_localctx, 1);
				{
				setState(2287);
				arrayType();
				}
				break;
			case LCURLY:
				enterOuterAlt(_localctx, 2);
				{
				setState(2288);
				mapType();
				}
				break;
			case LPAREN:
				enterOuterAlt(_localctx, 3);
				{
				setState(2289);
				tupleType();
				}
				break;
			case TYPE_IDENTIFIER:
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 4);
				{
				setState(2290);
				typeIdentifier();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeAnnotationContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public TypeAnnotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAnnotation; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeAnnotation(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeAnnotationContext typeAnnotation() throws RecognitionException {
		TypeAnnotationContext _localctx = new TypeAnnotationContext(_ctx, getState());
		enterRule(_localctx, 312, RULE_typeAnnotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2294);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(2293);
				match(COLON);
				}
			}

			setState(2296);
			type_(0);
			setState(2298);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,313,_ctx) ) {
			case 1:
				{
				setState(2297);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeIdentifierContext extends ParserRuleContext {
		public List<TypeIdentifierClauseContext> typeIdentifierClause() {
			return getRuleContexts(TypeIdentifierClauseContext.class);
		}
		public TypeIdentifierClauseContext typeIdentifierClause(int i) {
			return getRuleContext(TypeIdentifierClauseContext.class,i);
		}
		public PackageIdentifierContext packageIdentifier() {
			return getRuleContext(PackageIdentifierContext.class,0);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public TypeIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeIdentifierContext typeIdentifier() throws RecognitionException {
		TypeIdentifierContext _localctx = new TypeIdentifierContext(_ctx, getState());
		enterRule(_localctx, 314, RULE_typeIdentifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VALUE_IDENTIFIER) {
				{
				setState(2300);
				packageIdentifier();
				setState(2301);
				match(DOT);
				}
			}

			setState(2305);
			typeIdentifierClause();
			setState(2310);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,315,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2306);
					match(DOT);
					setState(2307);
					typeIdentifierClause();
					}
					} 
				}
				setState(2312);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,315,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeIdentifierClauseContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public GenericArgumentClauseContext genericArgumentClause() {
			return getRuleContext(GenericArgumentClauseContext.class,0);
		}
		public TypeIdentifierClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeIdentifierClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeIdentifierClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeIdentifierClauseContext typeIdentifierClause() throws RecognitionException {
		TypeIdentifierClauseContext _localctx = new TypeIdentifierClauseContext(_ctx, getState());
		enterRule(_localctx, 316, RULE_typeIdentifierClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2313);
			typeName();
			setState(2315);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,316,_ctx) ) {
			case 1:
				{
				setState(2314);
				genericArgumentClause();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeNameContext extends ParserRuleContext {
		public TerminalNode TYPE_IDENTIFIER() { return getToken(MojoParser.TYPE_IDENTIFIER, 0); }
		public TypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeName; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeNameContext typeName() throws RecognitionException {
		TypeNameContext _localctx = new TypeNameContext(_ctx, getState());
		enterRule(_localctx, 318, RULE_typeName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2317);
			match(TYPE_IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TupleTypeContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(MojoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(MojoParser.RPAREN, 0); }
		public TupleTypeElementsContext tupleTypeElements() {
			return getRuleContext(TupleTypeElementsContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public TupleTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTupleType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TupleTypeContext tupleType() throws RecognitionException {
		TupleTypeContext _localctx = new TupleTypeContext(_ctx, getState());
		enterRule(_localctx, 320, RULE_tupleType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2319);
			match(LPAREN);
			setState(2327);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,318,_ctx) ) {
			case 1:
				{
				setState(2323);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(2320);
					match(EOL);
					}
					}
					setState(2325);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2326);
				tupleTypeElements();
				}
				break;
			}
			setState(2332);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(2329);
				match(EOL);
				}
				}
				setState(2334);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(2335);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TupleTypeElementsContext extends ParserRuleContext {
		public List<TupleTypeElementContext> tupleTypeElement() {
			return getRuleContexts(TupleTypeElementContext.class);
		}
		public TupleTypeElementContext tupleTypeElement(int i) {
			return getRuleContext(TupleTypeElementContext.class,i);
		}
		public List<EovWithDocumentContext> eovWithDocument() {
			return getRuleContexts(EovWithDocumentContext.class);
		}
		public EovWithDocumentContext eovWithDocument(int i) {
			return getRuleContext(EovWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public TupleTypeElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleTypeElements; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTupleTypeElements(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TupleTypeElementsContext tupleTypeElements() throws RecognitionException {
		TupleTypeElementsContext _localctx = new TupleTypeElementsContext(_ctx, getState());
		enterRule(_localctx, 322, RULE_tupleTypeElements);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2337);
			tupleTypeElement();
			setState(2349);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,321,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2338);
					eovWithDocument();
					setState(2342);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(2339);
						match(EOL);
						}
						}
						setState(2344);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(2345);
					tupleTypeElement();
					}
					} 
				}
				setState(2351);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,321,_ctx);
			}
			setState(2353);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,322,_ctx) ) {
			case 1:
				{
				setState(2352);
				eovWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TupleTypeElementContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public DeclarationIdentifierContext declarationIdentifier() {
			return getRuleContext(DeclarationIdentifierContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public TupleTypeElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleTypeElement; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTupleTypeElement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TupleTypeElementContext tupleTypeElement() throws RecognitionException {
		TupleTypeElementContext _localctx = new TupleTypeElementContext(_ctx, getState());
		enterRule(_localctx, 324, RULE_tupleTypeElement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2359);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,324,_ctx) ) {
			case 1:
				{
				setState(2355);
				declarationIdentifier();
				setState(2357);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(2356);
					match(COLON);
					}
				}

				}
				break;
			}
			setState(2361);
			type_(0);
			setState(2363);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2362);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionTypeContext extends ParserRuleContext {
		public FunctionParameterClauseContext functionParameterClause() {
			return getRuleContext(FunctionParameterClauseContext.class,0);
		}
		public ArrowOperatorContext arrowOperator() {
			return getRuleContext(ArrowOperatorContext.class,0);
		}
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public FunctionTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFunctionType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FunctionTypeContext functionType() throws RecognitionException {
		FunctionTypeContext _localctx = new FunctionTypeContext(_ctx, getState());
		enterRule(_localctx, 326, RULE_functionType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2365);
			functionParameterClause();
			setState(2366);
			arrowOperator();
			setState(2367);
			type_(0);
			setState(2369);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,326,_ctx) ) {
			case 1:
				{
				setState(2368);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayTypeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(MojoParser.LBRACK, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode RBRACK() { return getToken(MojoParser.RBRACK, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public ArrayTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrayType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayTypeContext arrayType() throws RecognitionException {
		ArrayTypeContext _localctx = new ArrayTypeContext(_ctx, getState());
		enterRule(_localctx, 328, RULE_arrayType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2371);
			match(LBRACK);
			setState(2372);
			type_(0);
			setState(2374);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2373);
				attributes();
				}
			}

			setState(2376);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MapTypeContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public List<Type_Context> type_() {
			return getRuleContexts(Type_Context.class);
		}
		public Type_Context type_(int i) {
			return getRuleContext(Type_Context.class,i);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public KeyAttributesContext keyAttributes() {
			return getRuleContext(KeyAttributesContext.class,0);
		}
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public MapTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapType; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitMapType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MapTypeContext mapType() throws RecognitionException {
		MapTypeContext _localctx = new MapTypeContext(_ctx, getState());
		enterRule(_localctx, 330, RULE_mapType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2378);
			match(LCURLY);
			setState(2379);
			type_(0);
			setState(2381);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2380);
				keyAttributes();
				}
			}

			setState(2384);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(2383);
				match(COLON);
				}
			}

			setState(2386);
			type_(0);
			setState(2388);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2387);
				attributes();
				}
			}

			setState(2390);
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

	@SuppressWarnings("CheckReturnValue")
	public static class KeyAttributesContext extends ParserRuleContext {
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public KeyAttributesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyAttributes; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitKeyAttributes(this);
			else return visitor.visitChildren(this);
		}
	}

	public final KeyAttributesContext keyAttributes() throws RecognitionException {
		KeyAttributesContext _localctx = new KeyAttributesContext(_ctx, getState());
		enterRule(_localctx, 332, RULE_keyAttributes);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2392);
			attributes();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeInheritanceClauseContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public TypeInheritancesContext typeInheritances() {
			return getRuleContext(TypeInheritancesContext.class,0);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public TypeInheritanceClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeInheritanceClause; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeInheritanceClause(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeInheritanceClauseContext typeInheritanceClause() throws RecognitionException {
		TypeInheritanceClauseContext _localctx = new TypeInheritanceClauseContext(_ctx, getState());
		enterRule(_localctx, 334, RULE_typeInheritanceClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2394);
			match(COLON);
			setState(2398);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(2395);
				match(EOL);
				}
				}
				setState(2400);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(2401);
			typeInheritances();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeInheritancesContext extends ParserRuleContext {
		public List<TypeInheritanceContext> typeInheritance() {
			return getRuleContexts(TypeInheritanceContext.class);
		}
		public TypeInheritanceContext typeInheritance(int i) {
			return getRuleContext(TypeInheritanceContext.class,i);
		}
		public List<EovWithDocumentContext> eovWithDocument() {
			return getRuleContexts(EovWithDocumentContext.class);
		}
		public EovWithDocumentContext eovWithDocument(int i) {
			return getRuleContext(EovWithDocumentContext.class,i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public TypeInheritancesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeInheritances; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeInheritances(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeInheritancesContext typeInheritances() throws RecognitionException {
		TypeInheritancesContext _localctx = new TypeInheritancesContext(_ctx, getState());
		enterRule(_localctx, 336, RULE_typeInheritances);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2403);
			typeInheritance();
			setState(2415);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,333,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2404);
					eovWithDocument();
					setState(2408);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(2405);
						match(EOL);
						}
						}
						setState(2410);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(2411);
					typeInheritance();
					}
					} 
				}
				setState(2417);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,333,_ctx);
			}
			setState(2419);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,334,_ctx) ) {
			case 1:
				{
				setState(2418);
				eovWithDocument();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeInheritanceContext extends ParserRuleContext {
		public BasicTypeContext basicType() {
			return getRuleContext(BasicTypeContext.class,0);
		}
		public AttributesContext attributes() {
			return getRuleContext(AttributesContext.class,0);
		}
		public TypeInheritanceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeInheritance; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTypeInheritance(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeInheritanceContext typeInheritance() throws RecognitionException {
		TypeInheritanceContext _localctx = new TypeInheritanceContext(_ctx, getState());
		enterRule(_localctx, 338, RULE_typeInheritance);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2421);
			basicType(0);
			setState(2423);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2422);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationIdentifierContext extends ParserRuleContext {
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public KeywordAsIdentifierInDeclarationsContext keywordAsIdentifierInDeclarations() {
			return getRuleContext(KeywordAsIdentifierInDeclarationsContext.class,0);
		}
		public DeclarationIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarationIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitDeclarationIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DeclarationIdentifierContext declarationIdentifier() throws RecognitionException {
		DeclarationIdentifierContext _localctx = new DeclarationIdentifierContext(_ctx, getState());
		enterRule(_localctx, 340, RULE_declarationIdentifier);
		try {
			setState(2427);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2425);
				match(VALUE_IDENTIFIER);
				}
				break;
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_CONTINUE:
			case KEYWORD_ELSE:
			case KEYWORD_ENUM:
			case KEYWORD_FALSE:
			case KEYWORD_FUNC:
			case KEYWORD_IMPORT:
			case KEYWORD_IN:
			case KEYWORD_INTERFACE:
			case KEYWORD_IS:
			case KEYWORD_MATCH:
			case KEYWORD_NOT:
			case KEYWORD_NULL:
			case KEYWORD_OR:
			case KEYWORD_PACKAGE:
			case KEYWORD_STRUCT:
			case KEYWORD_TRUE:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(2426);
				keywordAsIdentifierInDeclarations();
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

	@SuppressWarnings("CheckReturnValue")
	public static class LabelIdentifierContext extends ParserRuleContext {
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public KeywordAsIdentifierInLabelsContext keywordAsIdentifierInLabels() {
			return getRuleContext(KeywordAsIdentifierInLabelsContext.class,0);
		}
		public LabelIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labelIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitLabelIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LabelIdentifierContext labelIdentifier() throws RecognitionException {
		LabelIdentifierContext _localctx = new LabelIdentifierContext(_ctx, getState());
		enterRule(_localctx, 342, RULE_labelIdentifier);
		try {
			setState(2431);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2429);
				match(VALUE_IDENTIFIER);
				}
				break;
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_CONTINUE:
			case KEYWORD_ELSE:
			case KEYWORD_ENUM:
			case KEYWORD_FALSE:
			case KEYWORD_FOR:
			case KEYWORD_FUNC:
			case KEYWORD_IMPORT:
			case KEYWORD_IN:
			case KEYWORD_INTERFACE:
			case KEYWORD_IS:
			case KEYWORD_MATCH:
			case KEYWORD_NOT:
			case KEYWORD_NULL:
			case KEYWORD_OR:
			case KEYWORD_PACKAGE:
			case KEYWORD_RETURN:
			case KEYWORD_STRUCT:
			case KEYWORD_TRUE:
			case KEYWORD_TYPE:
			case KEYWORD_VAR:
			case KEYWORD_WHILE:
			case KEYWORD_XOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(2430);
				keywordAsIdentifierInLabels();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PathIdentifierContext extends ParserRuleContext {
		public List<DeclarationIdentifierContext> declarationIdentifier() {
			return getRuleContexts(DeclarationIdentifierContext.class);
		}
		public DeclarationIdentifierContext declarationIdentifier(int i) {
			return getRuleContext(DeclarationIdentifierContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(MojoParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(MojoParser.DOT, i);
		}
		public PathIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pathIdentifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPathIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PathIdentifierContext pathIdentifier() throws RecognitionException {
		PathIdentifierContext _localctx = new PathIdentifierContext(_ctx, getState());
		enterRule(_localctx, 344, RULE_pathIdentifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2433);
			declarationIdentifier();
			setState(2438);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(2434);
				match(DOT);
				setState(2435);
				declarationIdentifier();
				}
				}
				setState(2440);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode VALUE_IDENTIFIER() { return getToken(MojoParser.VALUE_IDENTIFIER, 0); }
		public TerminalNode IMPLICIT_PARAMETER_NAME() { return getToken(MojoParser.IMPLICIT_PARAMETER_NAME, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIdentifier(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 346, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2441);
			_la = _input.LA(1);
			if ( !(_la==VALUE_IDENTIFIER || _la==IMPLICIT_PARAMETER_NAME) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordAsIdentifierInDeclarationsContext extends ParserRuleContext {
		public TerminalNode KEYWORD_AND() { return getToken(MojoParser.KEYWORD_AND, 0); }
		public TerminalNode KEYWORD_AS() { return getToken(MojoParser.KEYWORD_AS, 0); }
		public TerminalNode KEYWORD_ATTRIBUTE() { return getToken(MojoParser.KEYWORD_ATTRIBUTE, 0); }
		public TerminalNode KEYWORD_BREAK() { return getToken(MojoParser.KEYWORD_BREAK, 0); }
		public TerminalNode KEYWORD_CONST() { return getToken(MojoParser.KEYWORD_CONST, 0); }
		public TerminalNode KEYWORD_CONTINUE() { return getToken(MojoParser.KEYWORD_CONTINUE, 0); }
		public TerminalNode KEYWORD_ELSE() { return getToken(MojoParser.KEYWORD_ELSE, 0); }
		public TerminalNode KEYWORD_ENUM() { return getToken(MojoParser.KEYWORD_ENUM, 0); }
		public TerminalNode KEYWORD_FALSE() { return getToken(MojoParser.KEYWORD_FALSE, 0); }
		public TerminalNode KEYWORD_FUNC() { return getToken(MojoParser.KEYWORD_FUNC, 0); }
		public TerminalNode KEYWORD_IMPORT() { return getToken(MojoParser.KEYWORD_IMPORT, 0); }
		public TerminalNode KEYWORD_IN() { return getToken(MojoParser.KEYWORD_IN, 0); }
		public TerminalNode KEYWORD_INTERFACE() { return getToken(MojoParser.KEYWORD_INTERFACE, 0); }
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public TerminalNode KEYWORD_MATCH() { return getToken(MojoParser.KEYWORD_MATCH, 0); }
		public TerminalNode KEYWORD_NOT() { return getToken(MojoParser.KEYWORD_NOT, 0); }
		public TerminalNode KEYWORD_NULL() { return getToken(MojoParser.KEYWORD_NULL, 0); }
		public TerminalNode KEYWORD_OR() { return getToken(MojoParser.KEYWORD_OR, 0); }
		public TerminalNode KEYWORD_PACKAGE() { return getToken(MojoParser.KEYWORD_PACKAGE, 0); }
		public TerminalNode KEYWORD_STRUCT() { return getToken(MojoParser.KEYWORD_STRUCT, 0); }
		public TerminalNode KEYWORD_TRUE() { return getToken(MojoParser.KEYWORD_TRUE, 0); }
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public TerminalNode KEYWORD_XOR() { return getToken(MojoParser.KEYWORD_XOR, 0); }
		public KeywordAsIdentifierInDeclarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keywordAsIdentifierInDeclarations; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitKeywordAsIdentifierInDeclarations(this);
			else return visitor.visitChildren(this);
		}
	}

	public final KeywordAsIdentifierInDeclarationsContext keywordAsIdentifierInDeclarations() throws RecognitionException {
		KeywordAsIdentifierInDeclarationsContext _localctx = new KeywordAsIdentifierInDeclarationsContext(_ctx, getState());
		enterRule(_localctx, 348, RULE_keywordAsIdentifierInDeclarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2443);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 331344894L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordAsIdentifierInLabelsContext extends ParserRuleContext {
		public TerminalNode KEYWORD_AND() { return getToken(MojoParser.KEYWORD_AND, 0); }
		public TerminalNode KEYWORD_AS() { return getToken(MojoParser.KEYWORD_AS, 0); }
		public TerminalNode KEYWORD_ATTRIBUTE() { return getToken(MojoParser.KEYWORD_ATTRIBUTE, 0); }
		public TerminalNode KEYWORD_BREAK() { return getToken(MojoParser.KEYWORD_BREAK, 0); }
		public TerminalNode KEYWORD_CONST() { return getToken(MojoParser.KEYWORD_CONST, 0); }
		public TerminalNode KEYWORD_CONTINUE() { return getToken(MojoParser.KEYWORD_CONTINUE, 0); }
		public TerminalNode KEYWORD_ELSE() { return getToken(MojoParser.KEYWORD_ELSE, 0); }
		public TerminalNode KEYWORD_ENUM() { return getToken(MojoParser.KEYWORD_ENUM, 0); }
		public TerminalNode KEYWORD_FALSE() { return getToken(MojoParser.KEYWORD_FALSE, 0); }
		public TerminalNode KEYWORD_FOR() { return getToken(MojoParser.KEYWORD_FOR, 0); }
		public TerminalNode KEYWORD_FUNC() { return getToken(MojoParser.KEYWORD_FUNC, 0); }
		public TerminalNode KEYWORD_IMPORT() { return getToken(MojoParser.KEYWORD_IMPORT, 0); }
		public TerminalNode KEYWORD_IN() { return getToken(MojoParser.KEYWORD_IN, 0); }
		public TerminalNode KEYWORD_INTERFACE() { return getToken(MojoParser.KEYWORD_INTERFACE, 0); }
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public TerminalNode KEYWORD_MATCH() { return getToken(MojoParser.KEYWORD_MATCH, 0); }
		public TerminalNode KEYWORD_NOT() { return getToken(MojoParser.KEYWORD_NOT, 0); }
		public TerminalNode KEYWORD_NULL() { return getToken(MojoParser.KEYWORD_NULL, 0); }
		public TerminalNode KEYWORD_OR() { return getToken(MojoParser.KEYWORD_OR, 0); }
		public TerminalNode KEYWORD_PACKAGE() { return getToken(MojoParser.KEYWORD_PACKAGE, 0); }
		public TerminalNode KEYWORD_RETURN() { return getToken(MojoParser.KEYWORD_RETURN, 0); }
		public TerminalNode KEYWORD_STRUCT() { return getToken(MojoParser.KEYWORD_STRUCT, 0); }
		public TerminalNode KEYWORD_TRUE() { return getToken(MojoParser.KEYWORD_TRUE, 0); }
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public TerminalNode KEYWORD_VAR() { return getToken(MojoParser.KEYWORD_VAR, 0); }
		public TerminalNode KEYWORD_WHILE() { return getToken(MojoParser.KEYWORD_WHILE, 0); }
		public TerminalNode KEYWORD_XOR() { return getToken(MojoParser.KEYWORD_XOR, 0); }
		public KeywordAsIdentifierInLabelsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keywordAsIdentifierInLabels; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitKeywordAsIdentifierInLabels(this);
			else return visitor.visitChildren(this);
		}
	}

	public final KeywordAsIdentifierInLabelsContext keywordAsIdentifierInLabels() throws RecognitionException {
		KeywordAsIdentifierInLabelsContext _localctx = new KeywordAsIdentifierInLabelsContext(_ctx, getState());
		enterRule(_localctx, 350, RULE_keywordAsIdentifierInLabels);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2445);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 536866814L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class DocumentContext extends ParserRuleContext {
		public List<TerminalNode> LINE_DOCUMENT() { return getTokens(MojoParser.LINE_DOCUMENT); }
		public TerminalNode LINE_DOCUMENT(int i) {
			return getToken(MojoParser.LINE_DOCUMENT, i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public DocumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_document; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitDocument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DocumentContext document() throws RecognitionException {
		DocumentContext _localctx = new DocumentContext(_ctx, getState());
		enterRule(_localctx, 352, RULE_document);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2447);
			match(LINE_DOCUMENT);
			setState(2452);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,339,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2448);
					match(EOL);
					setState(2449);
					match(LINE_DOCUMENT);
					}
					} 
				}
				setState(2454);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,339,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FollowingDocumentContext extends ParserRuleContext {
		public List<TerminalNode> FOLLOWING_LINE_DOCUMENT() { return getTokens(MojoParser.FOLLOWING_LINE_DOCUMENT); }
		public TerminalNode FOLLOWING_LINE_DOCUMENT(int i) {
			return getToken(MojoParser.FOLLOWING_LINE_DOCUMENT, i);
		}
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
		public FollowingDocumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_followingDocument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFollowingDocument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FollowingDocumentContext followingDocument() throws RecognitionException {
		FollowingDocumentContext _localctx = new FollowingDocumentContext(_ctx, getState());
		enterRule(_localctx, 354, RULE_followingDocument);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2455);
			match(FOLLOWING_LINE_DOCUMENT);
			setState(2460);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,340,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2456);
					match(EOL);
					setState(2457);
					match(FOLLOWING_LINE_DOCUMENT);
					}
					} 
				}
				setState(2462);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,340,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentOperatorContext extends ParserRuleContext {
		public TerminalNode EQUAL() { return getToken(MojoParser.EQUAL, 0); }
		public AssignmentOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitAssignmentOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AssignmentOperatorContext assignmentOperator() throws RecognitionException {
		AssignmentOperatorContext _localctx = new AssignmentOperatorContext(_ctx, getState());
		enterRule(_localctx, 356, RULE_assignmentOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2463);
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

	@SuppressWarnings("CheckReturnValue")
	public static class NegatePrefixOperatorContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(MojoParser.MINUS, 0); }
		public NegatePrefixOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_negatePrefixOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitNegatePrefixOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NegatePrefixOperatorContext negatePrefixOperator() throws RecognitionException {
		NegatePrefixOperatorContext _localctx = new NegatePrefixOperatorContext(_ctx, getState());
		enterRule(_localctx, 358, RULE_negatePrefixOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2465);
			match(MINUS);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrowOperatorContext extends ParserRuleContext {
		public TerminalNode RIGHT_ARROW() { return getToken(MojoParser.RIGHT_ARROW, 0); }
		public ArrowOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrowOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitArrowOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrowOperatorContext arrowOperator() throws RecognitionException {
		ArrowOperatorContext _localctx = new ArrowOperatorContext(_ctx, getState());
		enterRule(_localctx, 360, RULE_arrowOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2467);
			match(RIGHT_ARROW);
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

	@SuppressWarnings("CheckReturnValue")
	public static class RangeOperatorContext extends ParserRuleContext {
		public TerminalNode DOT_DOT() { return getToken(MojoParser.DOT_DOT, 0); }
		public RangeOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitRangeOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final RangeOperatorContext rangeOperator() throws RecognitionException {
		RangeOperatorContext _localctx = new RangeOperatorContext(_ctx, getState());
		enterRule(_localctx, 362, RULE_rangeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2469);
			match(DOT_DOT);
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

	@SuppressWarnings("CheckReturnValue")
	public static class HalfOpenRangeOperatorContext extends ParserRuleContext {
		public TerminalNode DOT_DOT_LT() { return getToken(MojoParser.DOT_DOT_LT, 0); }
		public HalfOpenRangeOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_halfOpenRangeOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitHalfOpenRangeOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final HalfOpenRangeOperatorContext halfOpenRangeOperator() throws RecognitionException {
		HalfOpenRangeOperatorContext _localctx = new HalfOpenRangeOperatorContext(_ctx, getState());
		enterRule(_localctx, 364, RULE_halfOpenRangeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2471);
			match(DOT_DOT_LT);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CloseRangeOperatorContext extends ParserRuleContext {
		public TerminalNode DOT_DOT_EQUAL() { return getToken(MojoParser.DOT_DOT_EQUAL, 0); }
		public CloseRangeOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_closeRangeOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitCloseRangeOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CloseRangeOperatorContext closeRangeOperator() throws RecognitionException {
		CloseRangeOperatorContext _localctx = new CloseRangeOperatorContext(_ctx, getState());
		enterRule(_localctx, 366, RULE_closeRangeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2473);
			match(DOT_DOT_EQUAL);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryOperatorContext extends ParserRuleContext {
		public RangeOperatorContext rangeOperator() {
			return getRuleContext(RangeOperatorContext.class,0);
		}
		public HalfOpenRangeOperatorContext halfOpenRangeOperator() {
			return getRuleContext(HalfOpenRangeOperatorContext.class,0);
		}
		public CloseRangeOperatorContext closeRangeOperator() {
			return getRuleContext(CloseRangeOperatorContext.class,0);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public TerminalNode KEYWORD_AND() { return getToken(MojoParser.KEYWORD_AND, 0); }
		public TerminalNode KEYWORD_OR() { return getToken(MojoParser.KEYWORD_OR, 0); }
		public BinaryOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitBinaryOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BinaryOperatorContext binaryOperator() throws RecognitionException {
		BinaryOperatorContext _localctx = new BinaryOperatorContext(_ctx, getState());
		enterRule(_localctx, 368, RULE_binaryOperator);
		try {
			setState(2481);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT_DOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(2475);
				rangeOperator();
				}
				break;
			case DOT_DOT_LT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2476);
				halfOpenRangeOperator();
				}
				break;
			case DOT_DOT_EQUAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(2477);
				closeRangeOperator();
				}
				break;
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case OPERATOR_HEAD_OTHER:
				enterOuterAlt(_localctx, 4);
				{
				setState(2478);
				operator();
				}
				break;
			case KEYWORD_AND:
				enterOuterAlt(_localctx, 5);
				{
				setState(2479);
				match(KEYWORD_AND);
				}
				break;
			case KEYWORD_OR:
				enterOuterAlt(_localctx, 6);
				{
				setState(2480);
				match(KEYWORD_OR);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrefixOperatorContext extends ParserRuleContext {
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public TerminalNode KEYWORD_NOT() { return getToken(MojoParser.KEYWORD_NOT, 0); }
		public PrefixOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefixOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPrefixOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrefixOperatorContext prefixOperator() throws RecognitionException {
		PrefixOperatorContext _localctx = new PrefixOperatorContext(_ctx, getState());
		enterRule(_localctx, 370, RULE_prefixOperator);
		try {
			setState(2485);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT:
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case OPERATOR_HEAD_OTHER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2483);
				operator();
				}
				break;
			case KEYWORD_NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2484);
				match(KEYWORD_NOT);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PostfixOperatorContext extends ParserRuleContext {
		public TerminalNode PLUS_PLUS() { return getToken(MojoParser.PLUS_PLUS, 0); }
		public TerminalNode MINUS_MINUS() { return getToken(MojoParser.MINUS_MINUS, 0); }
		public TerminalNode ELLIPSIS() { return getToken(MojoParser.ELLIPSIS, 0); }
		public PostfixOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_postfixOperator; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitPostfixOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PostfixOperatorContext postfixOperator() throws RecognitionException {
		PostfixOperatorContext _localctx = new PostfixOperatorContext(_ctx, getState());
		enterRule(_localctx, 372, RULE_postfixOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2487);
			_la = _input.LA(1);
			if ( !(((((_la - 57)) & ~0x3f) == 0 && ((1L << (_la - 57)) & 259L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class OperatorContext extends ParserRuleContext {
		public Operator_headContext operator_head() {
			return getRuleContext(Operator_headContext.class,0);
		}
		public Operator_charactersContext operator_characters() {
			return getRuleContext(Operator_charactersContext.class,0);
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
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 374, RULE_operator);
		try {
			int _alt;
			setState(2499);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case OPERATOR_HEAD_OTHER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2489);
				operator_head();
				setState(2491);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,343,_ctx) ) {
				case 1:
					{
					setState(2490);
					operator_characters();
					}
					break;
				}
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2493);
				dot_operator_head();
				setState(2495); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(2494);
						dot_operator_character();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(2497); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,344,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	@SuppressWarnings("CheckReturnValue")
	public static class Operator_charactersContext extends ParserRuleContext {
		public List<Operator_characterContext> operator_character() {
			return getRuleContexts(Operator_characterContext.class);
		}
		public Operator_characterContext operator_character(int i) {
			return getRuleContext(Operator_characterContext.class,i);
		}
		public Operator_charactersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_characters; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitOperator_characters(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Operator_charactersContext operator_characters() throws RecognitionException {
		Operator_charactersContext _localctx = new Operator_charactersContext(_ctx, getState());
		enterRule(_localctx, 376, RULE_operator_characters);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2503); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(2501);
					if (!(_input.index() > 0 && _input.get(_input.index()-1).getType() != WS)) throw new FailedPredicateException(this, "_input.index() > 0 && _input.get(_input.index()-1).getType() != WS");
					setState(2502);
					operator_character();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(2505); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,346,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Operator_characterContext extends ParserRuleContext {
		public Operator_headContext operator_head() {
			return getRuleContext(Operator_headContext.class,0);
		}
		public TerminalNode OPERATOR_FOLLOWING_CHARACTER() { return getToken(MojoParser.OPERATOR_FOLLOWING_CHARACTER, 0); }
		public Operator_characterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_character; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitOperator_character(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Operator_characterContext operator_character() throws RecognitionException {
		Operator_characterContext _localctx = new Operator_characterContext(_ctx, getState());
		enterRule(_localctx, 378, RULE_operator_character);
		try {
			setState(2509);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case OPERATOR_HEAD_OTHER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2507);
				operator_head();
				}
				break;
			case OPERATOR_FOLLOWING_CHARACTER:
				enterOuterAlt(_localctx, 2);
				{
				setState(2508);
				match(OPERATOR_FOLLOWING_CHARACTER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Operator_headContext extends ParserRuleContext {
		public TerminalNode SLASH() { return getToken(MojoParser.SLASH, 0); }
		public TerminalNode EQUAL() { return getToken(MojoParser.EQUAL, 0); }
		public TerminalNode MINUS() { return getToken(MojoParser.MINUS, 0); }
		public TerminalNode PLUS() { return getToken(MojoParser.PLUS, 0); }
		public TerminalNode BANG() { return getToken(MojoParser.BANG, 0); }
		public TerminalNode STAR() { return getToken(MojoParser.STAR, 0); }
		public TerminalNode PERCENT() { return getToken(MojoParser.PERCENT, 0); }
		public TerminalNode AND() { return getToken(MojoParser.AND, 0); }
		public TerminalNode PIPE() { return getToken(MojoParser.PIPE, 0); }
		public TerminalNode LT() { return getToken(MojoParser.LT, 0); }
		public TerminalNode GT() { return getToken(MojoParser.GT, 0); }
		public TerminalNode CARET() { return getToken(MojoParser.CARET, 0); }
		public TerminalNode TILDE() { return getToken(MojoParser.TILDE, 0); }
		public TerminalNode QUESTION() { return getToken(MojoParser.QUESTION, 0); }
		public TerminalNode OPERATOR_HEAD_OTHER() { return getToken(MojoParser.OPERATOR_HEAD_OTHER, 0); }
		public Operator_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_head; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitOperator_head(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Operator_headContext operator_head() throws RecognitionException {
		Operator_headContext _localctx = new Operator_headContext(_ctx, getState());
		enterRule(_localctx, 380, RULE_operator_head);
		int _la;
		try {
			setState(2513);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
				enterOuterAlt(_localctx, 1);
				{
				setState(2511);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 18005052660645888L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case OPERATOR_HEAD_OTHER:
				enterOuterAlt(_localctx, 2);
				{
				setState(2512);
				match(OPERATOR_HEAD_OTHER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Dot_operator_headContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(MojoParser.DOT, 0); }
		public Dot_operator_headContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dot_operator_head; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitDot_operator_head(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Dot_operator_headContext dot_operator_head() throws RecognitionException {
		Dot_operator_headContext _localctx = new Dot_operator_headContext(_ctx, getState());
		enterRule(_localctx, 382, RULE_dot_operator_head);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2515);
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

	@SuppressWarnings("CheckReturnValue")
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
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitDot_operator_character(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Dot_operator_characterContext dot_operator_character() throws RecognitionException {
		Dot_operator_characterContext _localctx = new Dot_operator_characterContext(_ctx, getState());
		enterRule(_localctx, 384, RULE_dot_operator_character);
		try {
			setState(2519);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(2517);
				match(DOT);
				}
				break;
			case LT:
			case GT:
			case BANG:
			case QUESTION:
			case AND:
			case MINUS:
			case EQUAL:
			case PIPE:
			case SLASH:
			case PLUS:
			case STAR:
			case PERCENT:
			case CARET:
			case TILDE:
			case OPERATOR_HEAD_OTHER:
			case OPERATOR_FOLLOWING_CHARACTER:
				enterOuterAlt(_localctx, 2);
				{
				setState(2518);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public NumericLiteralContext numericLiteral() {
			return getRuleContext(NumericLiteralContext.class,0);
		}
		public StringLiteralContext stringLiteral() {
			return getRuleContext(StringLiteralContext.class,0);
		}
		public BoolLiteralContext boolLiteral() {
			return getRuleContext(BoolLiteralContext.class,0);
		}
		public NullLiteralContext nullLiteral() {
			return getRuleContext(NullLiteralContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 386, RULE_literal);
		try {
			setState(2525);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MINUS:
			case BINARY_LITERAL:
			case OCTAL_LITERAL:
			case DECIMAL_LITERAL:
			case PURE_DECIMAL_DIGITS:
			case HEXADECIMAL_LITERAL:
			case FLOAT_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(2521);
				numericLiteral();
				}
				break;
			case STATIC_STRING_LITERAL:
			case INTERPOLATED_STRING_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(2522);
				stringLiteral();
				}
				break;
			case KEYWORD_FALSE:
			case KEYWORD_TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(2523);
				boolLiteral();
				}
				break;
			case KEYWORD_NULL:
				enterOuterAlt(_localctx, 4);
				{
				setState(2524);
				nullLiteral();
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

	@SuppressWarnings("CheckReturnValue")
	public static class BoolLiteralContext extends ParserRuleContext {
		public TerminalNode KEYWORD_TRUE() { return getToken(MojoParser.KEYWORD_TRUE, 0); }
		public TerminalNode KEYWORD_FALSE() { return getToken(MojoParser.KEYWORD_FALSE, 0); }
		public BoolLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitBoolLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BoolLiteralContext boolLiteral() throws RecognitionException {
		BoolLiteralContext _localctx = new BoolLiteralContext(_ctx, getState());
		enterRule(_localctx, 388, RULE_boolLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2527);
			_la = _input.LA(1);
			if ( !(_la==KEYWORD_FALSE || _la==KEYWORD_TRUE) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class NullLiteralContext extends ParserRuleContext {
		public TerminalNode KEYWORD_NULL() { return getToken(MojoParser.KEYWORD_NULL, 0); }
		public NullLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nullLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitNullLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NullLiteralContext nullLiteral() throws RecognitionException {
		NullLiteralContext _localctx = new NullLiteralContext(_ctx, getState());
		enterRule(_localctx, 390, RULE_nullLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2529);
			match(KEYWORD_NULL);
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

	@SuppressWarnings("CheckReturnValue")
	public static class NumericLiteralContext extends ParserRuleContext {
		public IntegerLiteralContext integerLiteral() {
			return getRuleContext(IntegerLiteralContext.class,0);
		}
		public NegatePrefixOperatorContext negatePrefixOperator() {
			return getRuleContext(NegatePrefixOperatorContext.class,0);
		}
		public TerminalNode FLOAT_LITERAL() { return getToken(MojoParser.FLOAT_LITERAL, 0); }
		public NumericLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numericLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitNumericLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NumericLiteralContext numericLiteral() throws RecognitionException {
		NumericLiteralContext _localctx = new NumericLiteralContext(_ctx, getState());
		enterRule(_localctx, 392, RULE_numericLiteral);
		int _la;
		try {
			setState(2539);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,353,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2532);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS) {
					{
					setState(2531);
					negatePrefixOperator();
					}
				}

				setState(2534);
				integerLiteral();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2536);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS) {
					{
					setState(2535);
					negatePrefixOperator();
					}
				}

				setState(2538);
				match(FLOAT_LITERAL);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IntegerLiteralContext extends ParserRuleContext {
		public TerminalNode BINARY_LITERAL() { return getToken(MojoParser.BINARY_LITERAL, 0); }
		public TerminalNode OCTAL_LITERAL() { return getToken(MojoParser.OCTAL_LITERAL, 0); }
		public TerminalNode DECIMAL_LITERAL() { return getToken(MojoParser.DECIMAL_LITERAL, 0); }
		public TerminalNode PURE_DECIMAL_DIGITS() { return getToken(MojoParser.PURE_DECIMAL_DIGITS, 0); }
		public TerminalNode HEXADECIMAL_LITERAL() { return getToken(MojoParser.HEXADECIMAL_LITERAL, 0); }
		public IntegerLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integerLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitIntegerLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IntegerLiteralContext integerLiteral() throws RecognitionException {
		IntegerLiteralContext _localctx = new IntegerLiteralContext(_ctx, getState());
		enterRule(_localctx, 394, RULE_integerLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2541);
			_la = _input.LA(1);
			if ( !(((((_la - 77)) & ~0x3f) == 0 && ((1L << (_la - 77)) & 31L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class StringLiteralContext extends ParserRuleContext {
		public TerminalNode STATIC_STRING_LITERAL() { return getToken(MojoParser.STATIC_STRING_LITERAL, 0); }
		public TerminalNode INTERPOLATED_STRING_LITERAL() { return getToken(MojoParser.INTERPOLATED_STRING_LITERAL, 0); }
		public StringLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringLiteral; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitStringLiteral(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StringLiteralContext stringLiteral() throws RecognitionException {
		StringLiteralContext _localctx = new StringLiteralContext(_ctx, getState());
		enterRule(_localctx, 396, RULE_stringLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2543);
			_la = _input.LA(1);
			if ( !(_la==STATIC_STRING_LITERAL || _la==INTERPOLATED_STRING_LITERAL) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class EosContext extends ParserRuleContext {
		public TerminalNode SEMI() { return getToken(MojoParser.SEMI, 0); }
		public TerminalNode EOL() { return getToken(MojoParser.EOL, 0); }
		public EosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eos; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEos(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EosContext eos() throws RecognitionException {
		EosContext _localctx = new EosContext(_ctx, getState());
		enterRule(_localctx, 398, RULE_eos);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2545);
			_la = _input.LA(1);
			if ( !(_la==SEMI || _la==EOL) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class EovContext extends ParserRuleContext {
		public TerminalNode COMMA() { return getToken(MojoParser.COMMA, 0); }
		public TerminalNode EOL() { return getToken(MojoParser.EOL, 0); }
		public EovContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eov; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEov(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EovContext eov() throws RecognitionException {
		EovContext _localctx = new EovContext(_ctx, getState());
		enterRule(_localctx, 400, RULE_eov);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2547);
			_la = _input.LA(1);
			if ( !(_la==COMMA || _la==EOL) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class EosWithDocumentContext extends ParserRuleContext {
		public TerminalNode SEMI() { return getToken(MojoParser.SEMI, 0); }
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public TerminalNode EOL() { return getToken(MojoParser.EOL, 0); }
		public EosWithDocumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eosWithDocument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEosWithDocument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EosWithDocumentContext eosWithDocument() throws RecognitionException {
		EosWithDocumentContext _localctx = new EosWithDocumentContext(_ctx, getState());
		enterRule(_localctx, 402, RULE_eosWithDocument);
		int _la;
		try {
			setState(2559);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SEMI:
				enterOuterAlt(_localctx, 1);
				{
				setState(2549);
				match(SEMI);
				setState(2553);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(2550);
					followingDocument();
					setState(2551);
					match(EOL);
					}
				}

				}
				break;
			case EOL:
			case FOLLOWING_LINE_DOCUMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2556);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(2555);
					followingDocument();
					}
				}

				setState(2558);
				match(EOL);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EovWithDocumentContext extends ParserRuleContext {
		public TerminalNode COMMA() { return getToken(MojoParser.COMMA, 0); }
		public FollowingDocumentContext followingDocument() {
			return getRuleContext(FollowingDocumentContext.class,0);
		}
		public TerminalNode EOL() { return getToken(MojoParser.EOL, 0); }
		public EovWithDocumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eovWithDocument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitEovWithDocument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EovWithDocumentContext eovWithDocument() throws RecognitionException {
		EovWithDocumentContext _localctx = new EovWithDocumentContext(_ctx, getState());
		enterRule(_localctx, 404, RULE_eovWithDocument);
		int _la;
		try {
			setState(2571);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(2561);
				match(COMMA);
				setState(2565);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,357,_ctx) ) {
				case 1:
					{
					setState(2562);
					followingDocument();
					setState(2563);
					match(EOL);
					}
					break;
				}
				}
				break;
			case EOL:
			case FOLLOWING_LINE_DOCUMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2568);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(2567);
					followingDocument();
					}
				}

				setState(2570);
				match(EOL);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 83:
			return pattern_sempred((PatternContext)_localctx, predIndex);
		case 115:
			return stringOperatorLiteral_sempred((StringOperatorLiteralContext)_localctx, predIndex);
		case 116:
			return suffixLiteralOperator_sempred((SuffixLiteralOperatorContext)_localctx, predIndex);
		case 153:
			return type__sempred((Type_Context)_localctx, predIndex);
		case 154:
			return basicType_sempred((BasicTypeContext)_localctx, predIndex);
		case 188:
			return operator_characters_sempred((Operator_charactersContext)_localctx, predIndex);
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
	private boolean stringOperatorLiteral_sempred(StringOperatorLiteralContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return _input.index() > 0 && _input.get(_input.index()-1).getType() != WS;
		}
		return true;
	}
	private boolean suffixLiteralOperator_sempred(SuffixLiteralOperatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return _input.index() > 0 && _input.get(_input.index()-1).getType() != WS;
		}
		return true;
	}
	private boolean type__sempred(Type_Context _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 6);
		case 4:
			return precpred(_ctx, 5);
		case 5:
			return precpred(_ctx, 4);
		case 6:
			return precpred(_ctx, 3);
		case 7:
			return precpred(_ctx, 2);
		case 8:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean basicType_sempred(BasicTypeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 9:
			return precpred(_ctx, 3);
		case 10:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean operator_characters_sempred(Operator_charactersContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return _input.index() > 0 && _input.get(_input.index()-1).getType() != WS;
		}
		return true;
	}

	private static final String _serializedATNSegment0 =
		"\u0004\u0001]\u0a0e\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u00070\u00021\u00071\u0002"+
		"2\u00072\u00023\u00073\u00024\u00074\u00025\u00075\u00026\u00076\u0002"+
		"7\u00077\u00028\u00078\u00029\u00079\u0002:\u0007:\u0002;\u0007;\u0002"+
		"<\u0007<\u0002=\u0007=\u0002>\u0007>\u0002?\u0007?\u0002@\u0007@\u0002"+
		"A\u0007A\u0002B\u0007B\u0002C\u0007C\u0002D\u0007D\u0002E\u0007E\u0002"+
		"F\u0007F\u0002G\u0007G\u0002H\u0007H\u0002I\u0007I\u0002J\u0007J\u0002"+
		"K\u0007K\u0002L\u0007L\u0002M\u0007M\u0002N\u0007N\u0002O\u0007O\u0002"+
		"P\u0007P\u0002Q\u0007Q\u0002R\u0007R\u0002S\u0007S\u0002T\u0007T\u0002"+
		"U\u0007U\u0002V\u0007V\u0002W\u0007W\u0002X\u0007X\u0002Y\u0007Y\u0002"+
		"Z\u0007Z\u0002[\u0007[\u0002\\\u0007\\\u0002]\u0007]\u0002^\u0007^\u0002"+
		"_\u0007_\u0002`\u0007`\u0002a\u0007a\u0002b\u0007b\u0002c\u0007c\u0002"+
		"d\u0007d\u0002e\u0007e\u0002f\u0007f\u0002g\u0007g\u0002h\u0007h\u0002"+
		"i\u0007i\u0002j\u0007j\u0002k\u0007k\u0002l\u0007l\u0002m\u0007m\u0002"+
		"n\u0007n\u0002o\u0007o\u0002p\u0007p\u0002q\u0007q\u0002r\u0007r\u0002"+
		"s\u0007s\u0002t\u0007t\u0002u\u0007u\u0002v\u0007v\u0002w\u0007w\u0002"+
		"x\u0007x\u0002y\u0007y\u0002z\u0007z\u0002{\u0007{\u0002|\u0007|\u0002"+
		"}\u0007}\u0002~\u0007~\u0002\u007f\u0007\u007f\u0002\u0080\u0007\u0080"+
		"\u0002\u0081\u0007\u0081\u0002\u0082\u0007\u0082\u0002\u0083\u0007\u0083"+
		"\u0002\u0084\u0007\u0084\u0002\u0085\u0007\u0085\u0002\u0086\u0007\u0086"+
		"\u0002\u0087\u0007\u0087\u0002\u0088\u0007\u0088\u0002\u0089\u0007\u0089"+
		"\u0002\u008a\u0007\u008a\u0002\u008b\u0007\u008b\u0002\u008c\u0007\u008c"+
		"\u0002\u008d\u0007\u008d\u0002\u008e\u0007\u008e\u0002\u008f\u0007\u008f"+
		"\u0002\u0090\u0007\u0090\u0002\u0091\u0007\u0091\u0002\u0092\u0007\u0092"+
		"\u0002\u0093\u0007\u0093\u0002\u0094\u0007\u0094\u0002\u0095\u0007\u0095"+
		"\u0002\u0096\u0007\u0096\u0002\u0097\u0007\u0097\u0002\u0098\u0007\u0098"+
		"\u0002\u0099\u0007\u0099\u0002\u009a\u0007\u009a\u0002\u009b\u0007\u009b"+
		"\u0002\u009c\u0007\u009c\u0002\u009d\u0007\u009d\u0002\u009e\u0007\u009e"+
		"\u0002\u009f\u0007\u009f\u0002\u00a0\u0007\u00a0\u0002\u00a1\u0007\u00a1"+
		"\u0002\u00a2\u0007\u00a2\u0002\u00a3\u0007\u00a3\u0002\u00a4\u0007\u00a4"+
		"\u0002\u00a5\u0007\u00a5\u0002\u00a6\u0007\u00a6\u0002\u00a7\u0007\u00a7"+
		"\u0002\u00a8\u0007\u00a8\u0002\u00a9\u0007\u00a9\u0002\u00aa\u0007\u00aa"+
		"\u0002\u00ab\u0007\u00ab\u0002\u00ac\u0007\u00ac\u0002\u00ad\u0007\u00ad"+
		"\u0002\u00ae\u0007\u00ae\u0002\u00af\u0007\u00af\u0002\u00b0\u0007\u00b0"+
		"\u0002\u00b1\u0007\u00b1\u0002\u00b2\u0007\u00b2\u0002\u00b3\u0007\u00b3"+
		"\u0002\u00b4\u0007\u00b4\u0002\u00b5\u0007\u00b5\u0002\u00b6\u0007\u00b6"+
		"\u0002\u00b7\u0007\u00b7\u0002\u00b8\u0007\u00b8\u0002\u00b9\u0007\u00b9"+
		"\u0002\u00ba\u0007\u00ba\u0002\u00bb\u0007\u00bb\u0002\u00bc\u0007\u00bc"+
		"\u0002\u00bd\u0007\u00bd\u0002\u00be\u0007\u00be\u0002\u00bf\u0007\u00bf"+
		"\u0002\u00c0\u0007\u00c0\u0002\u00c1\u0007\u00c1\u0002\u00c2\u0007\u00c2"+
		"\u0002\u00c3\u0007\u00c3\u0002\u00c4\u0007\u00c4\u0002\u00c5\u0007\u00c5"+
		"\u0002\u00c6\u0007\u00c6\u0002\u00c7\u0007\u00c7\u0002\u00c8\u0007\u00c8"+
		"\u0002\u00c9\u0007\u00c9\u0002\u00ca\u0007\u00ca\u0001\u0000\u0005\u0000"+
		"\u0198\b\u0000\n\u0000\f\u0000\u019b\t\u0000\u0001\u0000\u0003\u0000\u019e"+
		"\b\u0000\u0001\u0000\u0005\u0000\u01a1\b\u0000\n\u0000\f\u0000\u01a4\t"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0003\u0001\u01ac\b\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0003\u0001\u01b2\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0005\u0005\u01bf\b\u0005\n\u0005\f\u0005\u01c2\t\u0005"+
		"\u0001\u0005\u0001\u0005\u0005\u0005\u01c6\b\u0005\n\u0005\f\u0005\u01c9"+
		"\t\u0005\u0001\u0005\u0003\u0005\u01cc\b\u0005\u0001\u0006\u0001\u0006"+
		"\u0003\u0006\u01d0\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007"+
		"\u01d5\b\u0007\n\u0007\f\u0007\u01d8\t\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0005\u0007\u01dd\b\u0007\n\u0007\f\u0007\u01e0\t\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\b\u0001\b\u0001\b\u0005\b\u01e7\b\b\n\b\f\b\u01ea\t"+
		"\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0005\t\u01f1\b\t\n\t\f\t\u01f4"+
		"\t\t\u0001\t\u0001\t\u0005\t\u01f8\b\t\n\t\f\t\u01fb\t\t\u0001\n\u0001"+
		"\n\u0003\n\u01ff\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u0204"+
		"\b\u000b\n\u000b\f\u000b\u0207\t\u000b\u0001\u000b\u0001\u000b\u0001\f"+
		"\u0001\f\u0003\f\u020d\b\f\u0001\r\u0001\r\u0001\r\u0005\r\u0212\b\r\n"+
		"\r\f\r\u0215\t\r\u0001\r\u0001\r\u0005\r\u0219\b\r\n\r\f\r\u021c\t\r\u0001"+
		"\r\u0003\r\u021f\b\r\u0001\u000e\u0001\u000e\u0005\u000e\u0223\b\u000e"+
		"\n\u000e\f\u000e\u0226\t\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005"+
		"\u000e\u022b\b\u000e\n\u000e\f\u000e\u022e\t\u000e\u0001\u000e\u0003\u000e"+
		"\u0231\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f\u0236\b"+
		"\u000f\n\u000f\f\u000f\u0239\t\u000f\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u023d\b\u000f\n\u000f\f\u000f\u0240\t\u000f\u0001\u000f\u0003\u000f\u0243"+
		"\b\u000f\u0001\u000f\u0005\u000f\u0246\b\u000f\n\u000f\f\u000f\u0249\t"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0005"+
		"\u0010\u0250\b\u0010\n\u0010\f\u0010\u0253\t\u0010\u0001\u0010\u0001\u0010"+
		"\u0005\u0010\u0257\b\u0010\n\u0010\f\u0010\u025a\t\u0010\u0001\u0010\u0003"+
		"\u0010\u025d\b\u0010\u0001\u0011\u0001\u0011\u0003\u0011\u0261\b\u0011"+
		"\u0001\u0011\u0005\u0011\u0264\b\u0011\n\u0011\f\u0011\u0267\t\u0011\u0001"+
		"\u0011\u0001\u0011\u0005\u0011\u026b\b\u0011\n\u0011\f\u0011\u026e\t\u0011"+
		"\u0001\u0011\u0001\u0011\u0003\u0011\u0272\b\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0003\u0012\u0277\b\u0012\u0001\u0013\u0001\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0015\u0001\u0015\u0003\u0015\u027f\b\u0015\u0001\u0016"+
		"\u0001\u0016\u0005\u0016\u0283\b\u0016\n\u0016\f\u0016\u0286\t\u0016\u0001"+
		"\u0016\u0001\u0016\u0005\u0016\u028a\b\u0016\n\u0016\f\u0016\u028d\t\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017"+
		"\u0294\b\u0017\n\u0017\f\u0017\u0297\t\u0017\u0001\u0017\u0001\u0017\u0005"+
		"\u0017\u029b\b\u0017\n\u0017\f\u0017\u029e\t\u0017\u0001\u0017\u0003\u0017"+
		"\u02a1\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0003\u0018\u02aa\b\u0018\u0001\u0019\u0001\u0019"+
		"\u0005\u0019\u02ae\b\u0019\n\u0019\f\u0019\u02b1\t\u0019\u0001\u0019\u0001"+
		"\u0019\u0005\u0019\u02b5\b\u0019\n\u0019\f\u0019\u02b8\t\u0019\u0001\u0019"+
		"\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0005\u001a\u02bf\b\u001a"+
		"\n\u001a\f\u001a\u02c2\t\u001a\u0001\u001a\u0001\u001a\u0005\u001a\u02c6"+
		"\b\u001a\n\u001a\f\u001a\u02c9\t\u001a\u0001\u001b\u0001\u001b\u0003\u001b"+
		"\u02cd\b\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u02d2\b"+
		"\u001c\u0001\u001c\u0001\u001c\u0005\u001c\u02d6\b\u001c\n\u001c\f\u001c"+
		"\u02d9\t\u001c\u0003\u001c\u02db\b\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u02e8\b\u001c\u0001\u001d\u0001"+
		"\u001d\u0005\u001d\u02ec\b\u001d\n\u001d\f\u001d\u02ef\t\u001d\u0001\u001d"+
		"\u0003\u001d\u02f2\b\u001d\u0001\u001d\u0005\u001d\u02f5\b\u001d\n\u001d"+
		"\f\u001d\u02f8\t\u001d\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0005\u001e\u02ff\b\u001e\n\u001e\f\u001e\u0302\t\u001e\u0001"+
		"\u001e\u0003\u001e\u0305\b\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0005"+
		"\u001f\u030a\b\u001f\n\u001f\f\u001f\u030d\t\u001f\u0001 \u0001 \u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0003!\u0317\b!\u0001\"\u0001\"\u0001"+
		"\"\u0005\"\u031c\b\"\n\"\f\"\u031f\t\"\u0001#\u0001#\u0001$\u0001$\u0001"+
		"$\u0001%\u0001%\u0001%\u0001&\u0001&\u0001&\u0003&\u032c\b&\u0001\'\u0001"+
		"\'\u0001\'\u0001(\u0001(\u0001(\u0005(\u0334\b(\n(\f(\u0337\t(\u0001("+
		"\u0001(\u0005(\u033b\b(\n(\f(\u033e\t(\u0001(\u0001(\u0001)\u0001)\u0003"+
		")\u0344\b)\u0001)\u0001)\u0005)\u0348\b)\n)\f)\u034b\t)\u0001)\u0001)"+
		"\u0003)\u034f\b)\u0005)\u0351\b)\n)\f)\u0354\t)\u0001)\u0003)\u0357\b"+
		")\u0001*\u0001*\u0003*\u035b\b*\u0001+\u0001+\u0003+\u035f\b+\u0001,\u0001"+
		",\u0001,\u0001,\u0001,\u0005,\u0366\b,\n,\f,\u0369\t,\u0001,\u0001,\u0001"+
		",\u0005,\u036e\b,\n,\f,\u0371\t,\u0001,\u0001,\u0005,\u0375\b,\n,\f,\u0378"+
		"\t,\u0001,\u0003,\u037b\b,\u0001,\u0005,\u037e\b,\n,\f,\u0381\t,\u0001"+
		",\u0001,\u0003,\u0385\b,\u0001-\u0001-\u0001.\u0001.\u0001.\u0003.\u038c"+
		"\b.\u0001.\u0001.\u0001.\u0003.\u0391\b.\u0001.\u0001.\u0001/\u0001/\u0003"+
		"/\u0397\b/\u00010\u00010\u00050\u039b\b0\n0\f0\u039e\t0\u00010\u00010"+
		"\u00011\u00011\u00011\u00011\u00011\u00011\u00011\u00011\u00011\u0005"+
		"1\u03ab\b1\n1\f1\u03ae\t1\u00011\u00011\u00011\u00051\u03b3\b1\n1\f1\u03b6"+
		"\t1\u00011\u00011\u00051\u03ba\b1\n1\f1\u03bd\t1\u00011\u00031\u03c0\b"+
		"1\u00011\u00051\u03c3\b1\n1\f1\u03c6\t1\u00011\u00011\u00031\u03ca\b1"+
		"\u00012\u00012\u00012\u00032\u03cf\b2\u00012\u00052\u03d2\b2\n2\f2\u03d5"+
		"\t2\u00012\u00012\u00013\u00013\u00014\u00014\u00054\u03dd\b4\n4\f4\u03e0"+
		"\t4\u00014\u00014\u00034\u03e4\b4\u00014\u00034\u03e7\b4\u00015\u0001"+
		"5\u00015\u00035\u03ec\b5\u00015\u00015\u00055\u03f0\b5\n5\f5\u03f3\t5"+
		"\u00015\u00035\u03f6\b5\u00016\u00016\u00036\u03fa\b6\u00017\u00017\u0003"+
		"7\u03fe\b7\u00017\u00057\u0401\b7\n7\f7\u0404\t7\u00017\u00037\u0407\b"+
		"7\u00018\u00018\u00018\u00038\u040c\b8\u00038\u040e\b8\u00018\u00018\u0003"+
		"8\u0412\b8\u00018\u00058\u0415\b8\n8\f8\u0418\t8\u00018\u00038\u041b\b"+
		"8\u00019\u00019\u00039\u041f\b9\u00019\u00059\u0422\b9\n9\f9\u0425\t9"+
		"\u00019\u00039\u0428\b9\u00019\u00059\u042b\b9\n9\f9\u042e\t9\u00019\u0001"+
		"9\u0001:\u0001:\u0001:\u0001:\u0005:\u0436\b:\n:\f:\u0439\t:\u0001:\u0001"+
		":\u0005:\u043d\b:\n:\f:\u0440\t:\u0001:\u0001:\u0003:\u0444\b:\u0001;"+
		"\u0001;\u0001;\u0005;\u0449\b;\n;\f;\u044c\t;\u0001;\u0001;\u0005;\u0450"+
		"\b;\n;\f;\u0453\t;\u0001;\u0003;\u0456\b;\u0001<\u0001<\u0001<\u0005<"+
		"\u045b\b<\n<\f<\u045e\t<\u0001<\u0003<\u0461\b<\u0001<\u0001<\u0003<\u0465"+
		"\b<\u0001<\u0001<\u0001<\u0003<\u046a\b<\u0003<\u046c\b<\u0001=\u0001"+
		"=\u0001=\u0003=\u0471\b=\u0001=\u0005=\u0474\b=\n=\f=\u0477\t=\u0001="+
		"\u0003=\u047a\b=\u0001=\u0005=\u047d\b=\n=\f=\u0480\t=\u0001=\u0001=\u0001"+
		"=\u0001=\u0001=\u0001=\u0001=\u0003=\u0489\b=\u0001=\u0005=\u048c\b=\n"+
		"=\f=\u048f\t=\u0001=\u0003=\u0492\b=\u0001=\u0005=\u0495\b=\n=\f=\u0498"+
		"\t=\u0001=\u0001=\u0003=\u049c\b=\u0001>\u0001>\u0003>\u04a0\b>\u0001"+
		">\u0005>\u04a3\b>\n>\f>\u04a6\t>\u0001>\u0003>\u04a9\b>\u0001>\u0005>"+
		"\u04ac\b>\n>\f>\u04af\t>\u0001>\u0001>\u0001?\u0001?\u0001@\u0001@\u0001"+
		"@\u0005@\u04b8\b@\n@\f@\u04bb\t@\u0001@\u0001@\u0005@\u04bf\b@\n@\f@\u04c2"+
		"\t@\u0001@\u0003@\u04c5\b@\u0001A\u0001A\u0001A\u0003A\u04ca\bA\u0001"+
		"A\u0001A\u0001A\u0003A\u04cf\bA\u0001A\u0001A\u0003A\u04d3\bA\u0001A\u0005"+
		"A\u04d6\bA\nA\fA\u04d9\tA\u0001A\u0003A\u04dc\bA\u0001A\u0003A\u04df\b"+
		"A\u0001B\u0001B\u0001B\u0003B\u04e4\bB\u0001B\u0001B\u0001B\u0001B\u0001"+
		"B\u0001B\u0001B\u0003B\u04ed\bB\u0001B\u0001B\u0003B\u04f1\bB\u0001C\u0001"+
		"C\u0001D\u0005D\u04f6\bD\nD\fD\u04f9\tD\u0001D\u0003D\u04fc\bD\u0001D"+
		"\u0005D\u04ff\bD\nD\fD\u0502\tD\u0001D\u0003D\u0505\bD\u0001E\u0001E\u0001"+
		"E\u0001E\u0003E\u050b\bE\u0001E\u0005E\u050e\bE\nE\fE\u0511\tE\u0001E"+
		"\u0003E\u0514\bE\u0001E\u0005E\u0517\bE\nE\fE\u051a\tE\u0001E\u0001E\u0001"+
		"F\u0001F\u0001F\u0005F\u0521\bF\nF\fF\u0524\tF\u0001F\u0001F\u0005F\u0528"+
		"\bF\nF\fF\u052b\tF\u0001F\u0003F\u052e\bF\u0001G\u0001G\u0001G\u0003G"+
		"\u0533\bG\u0001G\u0001G\u0001G\u0003G\u0538\bG\u0001G\u0001G\u0001G\u0001"+
		"G\u0001G\u0003G\u053f\bG\u0001G\u0003G\u0542\bG\u0001H\u0001H\u0001H\u0005"+
		"H\u0547\bH\nH\fH\u054a\tH\u0001H\u0003H\u054d\bH\u0001I\u0001I\u0001I"+
		"\u0003I\u0552\bI\u0001I\u0001I\u0001I\u0001I\u0001I\u0001I\u0001I\u0003"+
		"I\u055b\bI\u0001I\u0001I\u0003I\u055f\bI\u0001J\u0001J\u0001K\u0005K\u0564"+
		"\bK\nK\fK\u0567\tK\u0001K\u0003K\u056a\bK\u0001K\u0005K\u056d\bK\nK\f"+
		"K\u0570\tK\u0001K\u0001K\u0001L\u0001L\u0003L\u0576\bL\u0001L\u0005L\u0579"+
		"\bL\nL\fL\u057c\tL\u0001L\u0003L\u057f\bL\u0001L\u0005L\u0582\bL\nL\f"+
		"L\u0585\tL\u0001L\u0001L\u0001M\u0001M\u0001M\u0005M\u058c\bM\nM\fM\u058f"+
		"\tM\u0001M\u0001M\u0005M\u0593\bM\nM\fM\u0596\tM\u0001M\u0003M\u0599\b"+
		"M\u0001N\u0001N\u0001N\u0003N\u059e\bN\u0001N\u0001N\u0001N\u0003N\u05a3"+
		"\bN\u0001N\u0001N\u0003N\u05a7\bN\u0001N\u0003N\u05aa\bN\u0001O\u0001"+
		"O\u0003O\u05ae\bO\u0001O\u0005O\u05b1\bO\nO\fO\u05b4\tO\u0001O\u0001O"+
		"\u0001P\u0001P\u0001P\u0003P\u05bb\bP\u0001P\u0001P\u0001P\u0005P\u05c0"+
		"\bP\nP\fP\u05c3\tP\u0001P\u0003P\u05c6\bP\u0001P\u0003P\u05c9\bP\u0003"+
		"P\u05cb\bP\u0001Q\u0001Q\u0001Q\u0003Q\u05d0\bQ\u0001Q\u0005Q\u05d3\b"+
		"Q\nQ\fQ\u05d6\tQ\u0001Q\u0001Q\u0001R\u0001R\u0005R\u05dc\bR\nR\fR\u05df"+
		"\tR\u0001R\u0001R\u0001R\u0003R\u05e4\bR\u0001R\u0001R\u0003R\u05e8\b"+
		"R\u0001R\u0003R\u05eb\bR\u0001S\u0001S\u0001S\u0003S\u05f0\bS\u0001S\u0001"+
		"S\u0003S\u05f4\bS\u0001S\u0001S\u0003S\u05f8\bS\u0001S\u0001S\u0003S\u05fc"+
		"\bS\u0001S\u0001S\u0003S\u0600\bS\u0001S\u0001S\u0001S\u0003S\u0605\b"+
		"S\u0001S\u0001S\u0001S\u0003S\u060a\bS\u0001S\u0001S\u0001S\u0005S\u060f"+
		"\bS\nS\fS\u0612\tS\u0001T\u0001T\u0001U\u0001U\u0001V\u0001V\u0003V\u061a"+
		"\bV\u0001V\u0001V\u0001W\u0001W\u0001W\u0005W\u0621\bW\nW\fW\u0624\tW"+
		"\u0001X\u0001X\u0003X\u0628\bX\u0001Y\u0001Y\u0003Y\u062c\bY\u0001Y\u0001"+
		"Y\u0001Z\u0001Z\u0001Z\u0005Z\u0633\bZ\nZ\fZ\u0636\tZ\u0001[\u0001[\u0003"+
		"[\u063a\b[\u0001\\\u0003\\\u063d\b\\\u0001\\\u0001\\\u0001\\\u0003\\\u0642"+
		"\b\\\u0001]\u0001]\u0001]\u0001^\u0001^\u0001_\u0001_\u0001_\u0001_\u0001"+
		"_\u0003_\u064e\b_\u0001_\u0003_\u0651\b_\u0003_\u0653\b_\u0001`\u0001"+
		"`\u0001`\u0003`\u0658\b`\u0001`\u0001`\u0001a\u0001a\u0001b\u0001b\u0005"+
		"b\u0660\bb\nb\fb\u0663\tb\u0001b\u0003b\u0666\bb\u0001b\u0005b\u0669\b"+
		"b\nb\fb\u066c\tb\u0001b\u0001b\u0001c\u0001c\u0001c\u0003c\u0673\bc\u0001"+
		"c\u0001c\u0001d\u0001d\u0001d\u0005d\u067a\bd\nd\fd\u067d\td\u0001d\u0001"+
		"d\u0005d\u0681\bd\nd\fd\u0684\td\u0001d\u0003d\u0687\bd\u0001e\u0001e"+
		"\u0003e\u068b\be\u0001e\u0005e\u068e\be\ne\fe\u0691\te\u0001f\u0001f\u0003"+
		"f\u0695\bf\u0001g\u0001g\u0001g\u0001g\u0003g\u069b\bg\u0001h\u0001h\u0001"+
		"h\u0001h\u0001h\u0001h\u0001h\u0001h\u0001h\u0001h\u0001h\u0001h\u0001"+
		"h\u0001h\u0001h\u0003h\u06ac\bh\u0001i\u0001i\u0001j\u0001j\u0001k\u0004"+
		"k\u06b3\bk\u000bk\fk\u06b4\u0001l\u0001l\u0001l\u0003l\u06ba\bl\u0001"+
		"m\u0001m\u0001m\u0001m\u0001n\u0001n\u0001n\u0001n\u0001o\u0001o\u0001"+
		"o\u0003o\u06c7\bo\u0001o\u0001o\u0003o\u06cb\bo\u0001o\u0001o\u0003o\u06cf"+
		"\bo\u0001p\u0001p\u0001p\u0003p\u06d4\bp\u0001p\u0001p\u0001p\u0001p\u0003"+
		"p\u06da\bp\u0001p\u0001p\u0001p\u0001p\u0001p\u0001p\u0001p\u0003p\u06e3"+
		"\bp\u0001q\u0001q\u0001q\u0001q\u0001q\u0001q\u0001q\u0003q\u06ec\bq\u0001"+
		"r\u0001r\u0001r\u0001s\u0001s\u0001s\u0001s\u0003s\u06f5\bs\u0001s\u0001"+
		"s\u0001s\u0003s\u06fa\bs\u0001t\u0001t\u0001t\u0001u\u0001u\u0001v\u0001"+
		"v\u0005v\u0703\bv\nv\fv\u0706\tv\u0001v\u0003v\u0709\bv\u0001v\u0005v"+
		"\u070c\bv\nv\fv\u070f\tv\u0001v\u0001v\u0001w\u0001w\u0001w\u0005w\u0716"+
		"\bw\nw\fw\u0719\tw\u0001w\u0001w\u0005w\u071d\bw\nw\fw\u0720\tw\u0001"+
		"w\u0003w\u0723\bw\u0001x\u0001x\u0003x\u0727\bx\u0001y\u0001y\u0005y\u072b"+
		"\by\ny\fy\u072e\ty\u0001y\u0003y\u0731\by\u0001y\u0005y\u0734\by\ny\f"+
		"y\u0737\ty\u0001y\u0001y\u0001z\u0001z\u0001z\u0005z\u073e\bz\nz\fz\u0741"+
		"\tz\u0001z\u0001z\u0005z\u0745\bz\nz\fz\u0748\tz\u0001z\u0003z\u074b\b"+
		"z\u0001{\u0001{\u0003{\u074f\b{\u0001{\u0001{\u0001{\u0001|\u0001|\u0005"+
		"|\u0756\b|\n|\f|\u0759\t|\u0001|\u0003|\u075c\b|\u0001|\u0005|\u075f\b"+
		"|\n|\f|\u0762\t|\u0001|\u0001|\u0001}\u0001}\u0001}\u0005}\u0769\b}\n"+
		"}\f}\u076c\t}\u0001}\u0001}\u0005}\u0770\b}\n}\f}\u0773\t}\u0001}\u0003"+
		"}\u0776\b}\u0001~\u0001~\u0001~\u0003~\u077b\b~\u0001\u007f\u0001\u007f"+
		"\u0001\u007f\u0001\u0080\u0001\u0080\u0001\u0080\u0001\u0081\u0001\u0081"+
		"\u0005\u0081\u0785\b\u0081\n\u0081\f\u0081\u0788\t\u0081\u0001\u0081\u0001"+
		"\u0081\u0005\u0081\u078c\b\u0081\n\u0081\f\u0081\u078f\t\u0081\u0001\u0081"+
		"\u0003\u0081\u0792\b\u0081\u0001\u0081\u0005\u0081\u0795\b\u0081\n\u0081"+
		"\f\u0081\u0798\t\u0081\u0001\u0081\u0001\u0081\u0001\u0082\u0001\u0082"+
		"\u0001\u0082\u0005\u0082\u079f\b\u0082\n\u0082\f\u0082\u07a2\t\u0082\u0001"+
		"\u0082\u0001\u0082\u0005\u0082\u07a6\b\u0082\n\u0082\f\u0082\u07a9\t\u0082"+
		"\u0001\u0082\u0003\u0082\u07ac\b\u0082\u0001\u0083\u0001\u0083\u0005\u0083"+
		"\u07b0\b\u0083\n\u0083\f\u0083\u07b3\t\u0083\u0001\u0083\u0001\u0083\u0005"+
		"\u0083\u07b7\b\u0083\n\u0083\f\u0083\u07ba\t\u0083\u0001\u0083\u0001\u0083"+
		"\u0001\u0084\u0001\u0084\u0005\u0084\u07c0\b\u0084\n\u0084\f\u0084\u07c3"+
		"\t\u0084\u0001\u0084\u0001\u0084\u0005\u0084\u07c7\b\u0084\n\u0084\f\u0084"+
		"\u07ca\t\u0084\u0001\u0084\u0001\u0084\u0001\u0084\u0001\u0084\u0001\u0084"+
		"\u0005\u0084\u07d1\b\u0084\n\u0084\f\u0084\u07d4\t\u0084\u0001\u0084\u0001"+
		"\u0084\u0005\u0084\u07d8\b\u0084\n\u0084\f\u0084\u07db\t\u0084\u0001\u0084"+
		"\u0003\u0084\u07de\b\u0084\u0001\u0084\u0005\u0084\u07e1\b\u0084\n\u0084"+
		"\f\u0084\u07e4\t\u0084\u0001\u0084\u0001\u0084\u0001\u0084\u0003\u0084"+
		"\u07e9\b\u0084\u0001\u0085\u0001\u0085\u0001\u0085\u0005\u0085\u07ee\b"+
		"\u0085\n\u0085\f\u0085\u07f1\t\u0085\u0001\u0085\u0001\u0085\u0005\u0085"+
		"\u07f5\b\u0085\n\u0085\f\u0085\u07f8\t\u0085\u0001\u0085\u0003\u0085\u07fb"+
		"\b\u0085\u0001\u0086\u0001\u0086\u0003\u0086\u07ff\b\u0086\u0001\u0087"+
		"\u0001\u0087\u0001\u0087\u0001\u0088\u0001\u0088\u0005\u0088\u0806\b\u0088"+
		"\n\u0088\f\u0088\u0809\t\u0088\u0001\u0088\u0001\u0088\u0005\u0088\u080d"+
		"\b\u0088\n\u0088\f\u0088\u0810\t\u0088\u0001\u0088\u0001\u0088\u0001\u0089"+
		"\u0001\u0089\u0001\u0089\u0001\u0089\u0001\u0089\u0001\u0089\u0004\u0089"+
		"\u081a\b\u0089\u000b\u0089\f\u0089\u081b\u0001\u0089\u0001\u0089\u0003"+
		"\u0089\u0820\b\u0089\u0001\u008a\u0001\u008a\u0001\u008a\u0001\u008a\u0001"+
		"\u008a\u0001\u008a\u0003\u008a\u0828\b\u008a\u0001\u008b\u0001\u008b\u0001"+
		"\u008c\u0001\u008c\u0005\u008c\u082e\b\u008c\n\u008c\f\u008c\u0831\t\u008c"+
		"\u0001\u008c\u0003\u008c\u0834\b\u008c\u0001\u008d\u0001\u008d\u0001\u008d"+
		"\u0001\u008d\u0001\u008d\u0003\u008d\u083b\b\u008d\u0001\u008e\u0001\u008e"+
		"\u0001\u008e\u0001\u008e\u0001\u008e\u0001\u008e\u0001\u008e\u0001\u008e"+
		"\u0003\u008e\u0845\b\u008e\u0003\u008e\u0847\b\u008e\u0001\u008f\u0001"+
		"\u008f\u0001\u008f\u0001\u008f\u0001\u0090\u0003\u0090\u084e\b\u0090\u0001"+
		"\u0090\u0001\u0090\u0003\u0090\u0852\b\u0090\u0001\u0091\u0001\u0091\u0005"+
		"\u0091\u0856\b\u0091\n\u0091\f\u0091\u0859\t\u0091\u0001\u0091\u0001\u0091"+
		"\u0001\u0091\u0005\u0091\u085e\b\u0091\n\u0091\f\u0091\u0861\t\u0091\u0001"+
		"\u0091\u0001\u0091\u0005\u0091\u0865\b\u0091\n\u0091\f\u0091\u0868\t\u0091"+
		"\u0001\u0091\u0001\u0091\u0003\u0091\u086c\b\u0091\u0001\u0092\u0001\u0092"+
		"\u0001\u0092\u0005\u0092\u0871\b\u0092\n\u0092\f\u0092\u0874\t\u0092\u0001"+
		"\u0093\u0001\u0093\u0001\u0093\u0001\u0093\u0001\u0093\u0003\u0093\u087b"+
		"\b\u0093\u0001\u0094\u0001\u0094\u0003\u0094\u087f\b\u0094\u0001\u0095"+
		"\u0004\u0095\u0882\b\u0095\u000b\u0095\f\u0095\u0883\u0001\u0096\u0001"+
		"\u0096\u0001\u0096\u0001\u0096\u0001\u0097\u0001\u0097\u0005\u0097\u088c"+
		"\b\u0097\n\u0097\f\u0097\u088f\t\u0097\u0001\u0098\u0001\u0098\u0001\u0098"+
		"\u0001\u0099\u0001\u0099\u0001\u0099\u0003\u0099\u0897\b\u0099\u0001\u0099"+
		"\u0001\u0099\u0001\u0099\u0001\u0099\u0001\u0099\u0001\u0099\u0001\u0099"+
		"\u0001\u0099\u0001\u0099\u0001\u0099\u0001\u0099\u0001\u0099\u0005\u0099"+
		"\u08a5\b\u0099\n\u0099\f\u0099\u08a8\t\u0099\u0001\u009a\u0001\u009a\u0001"+
		"\u009a\u0001\u009a\u0001\u009a\u0003\u009a\u08af\b\u009a\u0001\u009a\u0001"+
		"\u009a\u0001\u009a\u0003\u009a\u08b4\b\u009a\u0001\u009a\u0005\u009a\u08b7"+
		"\b\u009a\n\u009a\f\u009a\u08ba\t\u009a\u0001\u009a\u0001\u009a\u0005\u009a"+
		"\u08be\b\u009a\n\u009a\f\u009a\u08c1\t\u009a\u0001\u009a\u0001\u009a\u0003"+
		"\u009a\u08c5\b\u009a\u0001\u009a\u0001\u009a\u0001\u009a\u0003\u009a\u08ca"+
		"\b\u009a\u0001\u009a\u0001\u009a\u0003\u009a\u08ce\b\u009a\u0001\u009a"+
		"\u0001\u009a\u0001\u009a\u0003\u009a\u08d3\b\u009a\u0001\u009a\u0005\u009a"+
		"\u08d6\b\u009a\n\u009a\f\u009a\u08d9\t\u009a\u0001\u009a\u0001\u009a\u0005"+
		"\u009a\u08dd\b\u009a\n\u009a\f\u009a\u08e0\t\u009a\u0001\u009a\u0001\u009a"+
		"\u0003\u009a\u08e4\b\u009a\u0001\u009a\u0001\u009a\u0001\u009a\u0003\u009a"+
		"\u08e9\b\u009a\u0005\u009a\u08eb\b\u009a\n\u009a\f\u009a\u08ee\t\u009a"+
		"\u0001\u009b\u0001\u009b\u0001\u009b\u0001\u009b\u0003\u009b\u08f4\b\u009b"+
		"\u0001\u009c\u0003\u009c\u08f7\b\u009c\u0001\u009c\u0001\u009c\u0003\u009c"+
		"\u08fb\b\u009c\u0001\u009d\u0001\u009d\u0001\u009d\u0003\u009d\u0900\b"+
		"\u009d\u0001\u009d\u0001\u009d\u0001\u009d\u0005\u009d\u0905\b\u009d\n"+
		"\u009d\f\u009d\u0908\t\u009d\u0001\u009e\u0001\u009e\u0003\u009e\u090c"+
		"\b\u009e\u0001\u009f\u0001\u009f\u0001\u00a0\u0001\u00a0\u0005\u00a0\u0912"+
		"\b\u00a0\n\u00a0\f\u00a0\u0915\t\u00a0\u0001\u00a0\u0003\u00a0\u0918\b"+
		"\u00a0\u0001\u00a0\u0005\u00a0\u091b\b\u00a0\n\u00a0\f\u00a0\u091e\t\u00a0"+
		"\u0001\u00a0\u0001\u00a0\u0001\u00a1\u0001\u00a1\u0001\u00a1\u0005\u00a1"+
		"\u0925\b\u00a1\n\u00a1\f\u00a1\u0928\t\u00a1\u0001\u00a1\u0001\u00a1\u0005"+
		"\u00a1\u092c\b\u00a1\n\u00a1\f\u00a1\u092f\t\u00a1\u0001\u00a1\u0003\u00a1"+
		"\u0932\b\u00a1\u0001\u00a2\u0001\u00a2\u0003\u00a2\u0936\b\u00a2\u0003"+
		"\u00a2\u0938\b\u00a2\u0001\u00a2\u0001\u00a2\u0003\u00a2\u093c\b\u00a2"+
		"\u0001\u00a3\u0001\u00a3\u0001\u00a3\u0001\u00a3\u0003\u00a3\u0942\b\u00a3"+
		"\u0001\u00a4\u0001\u00a4\u0001\u00a4\u0003\u00a4\u0947\b\u00a4\u0001\u00a4"+
		"\u0001\u00a4\u0001\u00a5\u0001\u00a5\u0001\u00a5\u0003\u00a5\u094e\b\u00a5"+
		"\u0001\u00a5\u0003\u00a5\u0951\b\u00a5\u0001\u00a5\u0001\u00a5\u0003\u00a5"+
		"\u0955\b\u00a5\u0001\u00a5\u0001\u00a5\u0001\u00a6\u0001\u00a6\u0001\u00a7"+
		"\u0001\u00a7\u0005\u00a7\u095d\b\u00a7\n\u00a7\f\u00a7\u0960\t\u00a7\u0001"+
		"\u00a7\u0001\u00a7\u0001\u00a8\u0001\u00a8\u0001\u00a8\u0005\u00a8\u0967"+
		"\b\u00a8\n\u00a8\f\u00a8\u096a\t\u00a8\u0001\u00a8\u0001\u00a8\u0005\u00a8"+
		"\u096e\b\u00a8\n\u00a8\f\u00a8\u0971\t\u00a8\u0001\u00a8\u0003\u00a8\u0974"+
		"\b\u00a8\u0001\u00a9\u0001\u00a9\u0003\u00a9\u0978\b\u00a9\u0001\u00aa"+
		"\u0001\u00aa\u0003\u00aa\u097c\b\u00aa\u0001\u00ab\u0001\u00ab\u0003\u00ab"+
		"\u0980\b\u00ab\u0001\u00ac\u0001\u00ac\u0001\u00ac\u0005\u00ac\u0985\b"+
		"\u00ac\n\u00ac\f\u00ac\u0988\t\u00ac\u0001\u00ad\u0001\u00ad\u0001\u00ae"+
		"\u0001\u00ae\u0001\u00af\u0001\u00af\u0001\u00b0\u0001\u00b0\u0001\u00b0"+
		"\u0005\u00b0\u0993\b\u00b0\n\u00b0\f\u00b0\u0996\t\u00b0\u0001\u00b1\u0001"+
		"\u00b1\u0001\u00b1\u0005\u00b1\u099b\b\u00b1\n\u00b1\f\u00b1\u099e\t\u00b1"+
		"\u0001\u00b2\u0001\u00b2\u0001\u00b3\u0001\u00b3\u0001\u00b4\u0001\u00b4"+
		"\u0001\u00b5\u0001\u00b5\u0001\u00b6\u0001\u00b6\u0001\u00b7\u0001\u00b7"+
		"\u0001\u00b8\u0001\u00b8\u0001\u00b8\u0001\u00b8\u0001\u00b8\u0001\u00b8"+
		"\u0003\u00b8\u09b2\b\u00b8\u0001\u00b9\u0001\u00b9\u0003\u00b9\u09b6\b"+
		"\u00b9\u0001\u00ba\u0001\u00ba\u0001\u00bb\u0001\u00bb\u0003\u00bb\u09bc"+
		"\b\u00bb\u0001\u00bb\u0001\u00bb\u0004\u00bb\u09c0\b\u00bb\u000b\u00bb"+
		"\f\u00bb\u09c1\u0003\u00bb\u09c4\b\u00bb\u0001\u00bc\u0001\u00bc\u0004"+
		"\u00bc\u09c8\b\u00bc\u000b\u00bc\f\u00bc\u09c9\u0001\u00bd\u0001\u00bd"+
		"\u0003\u00bd\u09ce\b\u00bd\u0001\u00be\u0001\u00be\u0003\u00be\u09d2\b"+
		"\u00be\u0001\u00bf\u0001\u00bf\u0001\u00c0\u0001\u00c0\u0003\u00c0\u09d8"+
		"\b\u00c0\u0001\u00c1\u0001\u00c1\u0001\u00c1\u0001\u00c1\u0003\u00c1\u09de"+
		"\b\u00c1\u0001\u00c2\u0001\u00c2\u0001\u00c3\u0001\u00c3\u0001\u00c4\u0003"+
		"\u00c4\u09e5\b\u00c4\u0001\u00c4\u0001\u00c4\u0003\u00c4\u09e9\b\u00c4"+
		"\u0001\u00c4\u0003\u00c4\u09ec\b\u00c4\u0001\u00c5\u0001\u00c5\u0001\u00c6"+
		"\u0001\u00c6\u0001\u00c7\u0001\u00c7\u0001\u00c8\u0001\u00c8\u0001\u00c9"+
		"\u0001\u00c9\u0001\u00c9\u0001\u00c9\u0003\u00c9\u09fa\b\u00c9\u0001\u00c9"+
		"\u0003\u00c9\u09fd\b\u00c9\u0001\u00c9\u0003\u00c9\u0a00\b\u00c9\u0001"+
		"\u00ca\u0001\u00ca\u0001\u00ca\u0001\u00ca\u0003\u00ca\u0a06\b\u00ca\u0001"+
		"\u00ca\u0003\u00ca\u0a09\b\u00ca\u0001\u00ca\u0003\u00ca\u0a0c\b\u00ca"+
		"\u0001\u00ca\u0000\u0003\u00a6\u0132\u0134\u00cb\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084\u0086\u0088"+
		"\u008a\u008c\u008e\u0090\u0092\u0094\u0096\u0098\u009a\u009c\u009e\u00a0"+
		"\u00a2\u00a4\u00a6\u00a8\u00aa\u00ac\u00ae\u00b0\u00b2\u00b4\u00b6\u00b8"+
		"\u00ba\u00bc\u00be\u00c0\u00c2\u00c4\u00c6\u00c8\u00ca\u00cc\u00ce\u00d0"+
		"\u00d2\u00d4\u00d6\u00d8\u00da\u00dc\u00de\u00e0\u00e2\u00e4\u00e6\u00e8"+
		"\u00ea\u00ec\u00ee\u00f0\u00f2\u00f4\u00f6\u00f8\u00fa\u00fc\u00fe\u0100"+
		"\u0102\u0104\u0106\u0108\u010a\u010c\u010e\u0110\u0112\u0114\u0116\u0118"+
		"\u011a\u011c\u011e\u0120\u0122\u0124\u0126\u0128\u012a\u012c\u012e\u0130"+
		"\u0132\u0134\u0136\u0138\u013a\u013c\u013e\u0140\u0142\u0144\u0146\u0148"+
		"\u014a\u014c\u014e\u0150\u0152\u0154\u0156\u0158\u015a\u015c\u015e\u0160"+
		"\u0162\u0164\u0166\u0168\u016a\u016c\u016e\u0170\u0172\u0174\u0176\u0178"+
		"\u017a\u017c\u017e\u0180\u0182\u0184\u0186\u0188\u018a\u018c\u018e\u0190"+
		"\u0192\u0194\u0000\f\u0002\u0000\u0017\u0017\u0019\u0019\u0001\u0000I"+
		"J\u0002\u0000JJLL\u0005\u0000\u0001\t\u000b\u000b\r\u0015\u0017\u0019"+
		"\u001c\u001c\u0002\u0000\u0001\u000b\r\u001c\u0002\u00009:AA\u0002\u0000"+
		"\'*,5\u0002\u0000\t\t\u0018\u0018\u0001\u0000MQ\u0001\u0000ST\u0002\u0000"+
		"&&YY\u0002\u0000$$YY\u0ae9\u0000\u0199\u0001\u0000\u0000\u0000\u0002\u01b1"+
		"\u0001\u0000\u0000\u0000\u0004\u01b3\u0001\u0000\u0000\u0000\u0006\u01b6"+
		"\u0001\u0000\u0000\u0000\b\u01b9\u0001\u0000\u0000\u0000\n\u01bb\u0001"+
		"\u0000\u0000\u0000\f\u01cf\u0001\u0000\u0000\u0000\u000e\u01d1\u0001\u0000"+
		"\u0000\u0000\u0010\u01e3\u0001\u0000\u0000\u0000\u0012\u01ed\u0001\u0000"+
		"\u0000\u0000\u0014\u01fe\u0001\u0000\u0000\u0000\u0016\u0200\u0001\u0000"+
		"\u0000\u0000\u0018\u020c\u0001\u0000\u0000\u0000\u001a\u020e\u0001\u0000"+
		"\u0000\u0000\u001c\u0230\u0001\u0000\u0000\u0000\u001e\u0232\u0001\u0000"+
		"\u0000\u0000 \u024c\u0001\u0000\u0000\u0000\"\u025e\u0001\u0000\u0000"+
		"\u0000$\u0276\u0001\u0000\u0000\u0000&\u0278\u0001\u0000\u0000\u0000("+
		"\u027a\u0001\u0000\u0000\u0000*\u027c\u0001\u0000\u0000\u0000,\u0280\u0001"+
		"\u0000\u0000\u0000.\u0290\u0001\u0000\u0000\u00000\u02a9\u0001\u0000\u0000"+
		"\u00002\u02ab\u0001\u0000\u0000\u00004\u02bb\u0001\u0000\u0000\u00006"+
		"\u02ca\u0001\u0000\u0000\u00008\u02d1\u0001\u0000\u0000\u0000:\u02e9\u0001"+
		"\u0000\u0000\u0000<\u02fb\u0001\u0000\u0000\u0000>\u0306\u0001\u0000\u0000"+
		"\u0000@\u030e\u0001\u0000\u0000\u0000B\u0310\u0001\u0000\u0000\u0000D"+
		"\u0318\u0001\u0000\u0000\u0000F\u0320\u0001\u0000\u0000\u0000H\u0322\u0001"+
		"\u0000\u0000\u0000J\u0325\u0001\u0000\u0000\u0000L\u0328\u0001\u0000\u0000"+
		"\u0000N\u032d\u0001\u0000\u0000\u0000P\u0330\u0001\u0000\u0000\u0000R"+
		"\u0343\u0001\u0000\u0000\u0000T\u0358\u0001\u0000\u0000\u0000V\u035c\u0001"+
		"\u0000\u0000\u0000X\u0384\u0001\u0000\u0000\u0000Z\u0386\u0001\u0000\u0000"+
		"\u0000\\\u038b\u0001\u0000\u0000\u0000^\u0394\u0001\u0000\u0000\u0000"+
		"`\u0398\u0001\u0000\u0000\u0000b\u03c9\u0001\u0000\u0000\u0000d\u03cb"+
		"\u0001\u0000\u0000\u0000f\u03d8\u0001\u0000\u0000\u0000h\u03da\u0001\u0000"+
		"\u0000\u0000j\u03e8\u0001\u0000\u0000\u0000l\u03f9\u0001\u0000\u0000\u0000"+
		"n\u03fb\u0001\u0000\u0000\u0000p\u0408\u0001\u0000\u0000\u0000r\u041c"+
		"\u0001\u0000\u0000\u0000t\u0443\u0001\u0000\u0000\u0000v\u0445\u0001\u0000"+
		"\u0000\u0000x\u046b\u0001\u0000\u0000\u0000z\u049b\u0001\u0000\u0000\u0000"+
		"|\u049d\u0001\u0000\u0000\u0000~\u04b2\u0001\u0000\u0000\u0000\u0080\u04b4"+
		"\u0001\u0000\u0000\u0000\u0082\u04de\u0001\u0000\u0000\u0000\u0084\u04f0"+
		"\u0001\u0000\u0000\u0000\u0086\u04f2\u0001\u0000\u0000\u0000\u0088\u04fb"+
		"\u0001\u0000\u0000\u0000\u008a\u0506\u0001\u0000\u0000\u0000\u008c\u051d"+
		"\u0001\u0000\u0000\u0000\u008e\u0541\u0001\u0000\u0000\u0000\u0090\u0543"+
		"\u0001\u0000\u0000\u0000\u0092\u055e\u0001\u0000\u0000\u0000\u0094\u0560"+
		"\u0001\u0000\u0000\u0000\u0096\u0569\u0001\u0000\u0000\u0000\u0098\u0573"+
		"\u0001\u0000\u0000\u0000\u009a\u0588\u0001\u0000\u0000\u0000\u009c\u05a9"+
		"\u0001\u0000\u0000\u0000\u009e\u05ab\u0001\u0000\u0000\u0000\u00a0\u05b7"+
		"\u0001\u0000\u0000\u0000\u00a2\u05cc\u0001\u0000\u0000\u0000\u00a4\u05d9"+
		"\u0001\u0000\u0000\u0000\u00a6\u0609\u0001\u0000\u0000\u0000\u00a8\u0613"+
		"\u0001\u0000\u0000\u0000\u00aa\u0615\u0001\u0000\u0000\u0000\u00ac\u0617"+
		"\u0001\u0000\u0000\u0000\u00ae\u061d\u0001\u0000\u0000\u0000\u00b0\u0627"+
		"\u0001\u0000\u0000\u0000\u00b2\u0629\u0001\u0000\u0000\u0000\u00b4\u062f"+
		"\u0001\u0000\u0000\u0000\u00b6\u0639\u0001\u0000\u0000\u0000\u00b8\u063c"+
		"\u0001\u0000\u0000\u0000\u00ba\u0643\u0001\u0000\u0000\u0000\u00bc\u0646"+
		"\u0001\u0000\u0000\u0000\u00be\u0652\u0001\u0000\u0000\u0000\u00c0\u0657"+
		"\u0001\u0000\u0000\u0000\u00c2\u065b\u0001\u0000\u0000\u0000\u00c4\u065d"+
		"\u0001\u0000\u0000\u0000\u00c6\u0672\u0001\u0000\u0000\u0000\u00c8\u0676"+
		"\u0001\u0000\u0000\u0000\u00ca\u0688\u0001\u0000\u0000\u0000\u00cc\u0692"+
		"\u0001\u0000\u0000\u0000\u00ce\u069a\u0001\u0000\u0000\u0000\u00d0\u06ab"+
		"\u0001\u0000\u0000\u0000\u00d2\u06ad\u0001\u0000\u0000\u0000\u00d4\u06af"+
		"\u0001\u0000\u0000\u0000\u00d6\u06b2\u0001\u0000\u0000\u0000\u00d8\u06b9"+
		"\u0001\u0000\u0000\u0000\u00da\u06bb\u0001\u0000\u0000\u0000\u00dc\u06bf"+
		"\u0001\u0000\u0000\u0000\u00de\u06ce\u0001\u0000\u0000\u0000\u00e0\u06e2"+
		"\u0001\u0000\u0000\u0000\u00e2\u06eb\u0001\u0000\u0000\u0000\u00e4\u06ed"+
		"\u0001\u0000\u0000\u0000\u00e6\u06f9\u0001\u0000\u0000\u0000\u00e8\u06fb"+
		"\u0001\u0000\u0000\u0000\u00ea\u06fe\u0001\u0000\u0000\u0000\u00ec\u0700"+
		"\u0001\u0000\u0000\u0000\u00ee\u0712\u0001\u0000\u0000\u0000\u00f0\u0726"+
		"\u0001\u0000\u0000\u0000\u00f2\u0728\u0001\u0000\u0000\u0000\u00f4\u073a"+
		"\u0001\u0000\u0000\u0000\u00f6\u074e\u0001\u0000\u0000\u0000\u00f8\u0753"+
		"\u0001\u0000\u0000\u0000\u00fa\u0765\u0001\u0000\u0000\u0000\u00fc\u0777"+
		"\u0001\u0000\u0000\u0000\u00fe\u077c\u0001\u0000\u0000\u0000\u0100\u077f"+
		"\u0001\u0000\u0000\u0000\u0102\u0782\u0001\u0000\u0000\u0000\u0104\u079b"+
		"\u0001\u0000\u0000\u0000\u0106\u07ad\u0001\u0000\u0000\u0000\u0108\u07e8"+
		"\u0001\u0000\u0000\u0000\u010a\u07ea\u0001\u0000\u0000\u0000\u010c\u07fe"+
		"\u0001\u0000\u0000\u0000\u010e\u0800\u0001\u0000\u0000\u0000\u0110\u0803"+
		"\u0001\u0000\u0000\u0000\u0112\u081f\u0001\u0000\u0000\u0000\u0114\u0827"+
		"\u0001\u0000\u0000\u0000\u0116\u0829\u0001\u0000\u0000\u0000\u0118\u082b"+
		"\u0001\u0000\u0000\u0000\u011a\u083a\u0001\u0000\u0000\u0000\u011c\u083c"+
		"\u0001\u0000\u0000\u0000\u011e\u0848\u0001\u0000\u0000\u0000\u0120\u0851"+
		"\u0001\u0000\u0000\u0000\u0122\u086b\u0001\u0000\u0000\u0000\u0124\u086d"+
		"\u0001\u0000\u0000\u0000\u0126\u087a\u0001\u0000\u0000\u0000\u0128\u087c"+
		"\u0001\u0000\u0000\u0000\u012a\u0881\u0001\u0000\u0000\u0000\u012c\u0885"+
		"\u0001\u0000\u0000\u0000\u012e\u0889\u0001\u0000\u0000\u0000\u0130\u0890"+
		"\u0001\u0000\u0000\u0000\u0132\u0896\u0001\u0000\u0000\u0000\u0134\u08a9"+
		"\u0001\u0000\u0000\u0000\u0136\u08f3\u0001\u0000\u0000\u0000\u0138\u08f6"+
		"\u0001\u0000\u0000\u0000\u013a\u08ff\u0001\u0000\u0000\u0000\u013c\u0909"+
		"\u0001\u0000\u0000\u0000\u013e\u090d\u0001\u0000\u0000\u0000\u0140\u090f"+
		"\u0001\u0000\u0000\u0000\u0142\u0921\u0001\u0000\u0000\u0000\u0144\u0937"+
		"\u0001\u0000\u0000\u0000\u0146\u093d\u0001\u0000\u0000\u0000\u0148\u0943"+
		"\u0001\u0000\u0000\u0000\u014a\u094a\u0001\u0000\u0000\u0000\u014c\u0958"+
		"\u0001\u0000\u0000\u0000\u014e\u095a\u0001\u0000\u0000\u0000\u0150\u0963"+
		"\u0001\u0000\u0000\u0000\u0152\u0975\u0001\u0000\u0000\u0000\u0154\u097b"+
		"\u0001\u0000\u0000\u0000\u0156\u097f\u0001\u0000\u0000\u0000\u0158\u0981"+
		"\u0001\u0000\u0000\u0000\u015a\u0989\u0001\u0000\u0000\u0000\u015c\u098b"+
		"\u0001\u0000\u0000\u0000\u015e\u098d\u0001\u0000\u0000\u0000\u0160\u098f"+
		"\u0001\u0000\u0000\u0000\u0162\u0997\u0001\u0000\u0000\u0000\u0164\u099f"+
		"\u0001\u0000\u0000\u0000\u0166\u09a1\u0001\u0000\u0000\u0000\u0168\u09a3"+
		"\u0001\u0000\u0000\u0000\u016a\u09a5\u0001\u0000\u0000\u0000\u016c\u09a7"+
		"\u0001\u0000\u0000\u0000\u016e\u09a9\u0001\u0000\u0000\u0000\u0170\u09b1"+
		"\u0001\u0000\u0000\u0000\u0172\u09b5\u0001\u0000\u0000\u0000\u0174\u09b7"+
		"\u0001\u0000\u0000\u0000\u0176\u09c3\u0001\u0000\u0000\u0000\u0178\u09c7"+
		"\u0001\u0000\u0000\u0000\u017a\u09cd\u0001\u0000\u0000\u0000\u017c\u09d1"+
		"\u0001\u0000\u0000\u0000\u017e\u09d3\u0001\u0000\u0000\u0000\u0180\u09d7"+
		"\u0001\u0000\u0000\u0000\u0182\u09dd\u0001\u0000\u0000\u0000\u0184\u09df"+
		"\u0001\u0000\u0000\u0000\u0186\u09e1\u0001\u0000\u0000\u0000\u0188\u09eb"+
		"\u0001\u0000\u0000\u0000\u018a\u09ed\u0001\u0000\u0000\u0000\u018c\u09ef"+
		"\u0001\u0000\u0000\u0000\u018e\u09f1\u0001\u0000\u0000\u0000\u0190\u09f3"+
		"\u0001\u0000\u0000\u0000\u0192\u09ff\u0001\u0000\u0000\u0000\u0194\u0a0b"+
		"\u0001\u0000\u0000\u0000\u0196\u0198\u0005Y\u0000\u0000\u0197\u0196\u0001"+
		"\u0000\u0000\u0000\u0198\u019b\u0001\u0000\u0000\u0000\u0199\u0197\u0001"+
		"\u0000\u0000\u0000\u0199\u019a\u0001\u0000\u0000\u0000\u019a\u019d\u0001"+
		"\u0000\u0000\u0000\u019b\u0199\u0001\u0000\u0000\u0000\u019c\u019e\u0003"+
		"\n\u0005\u0000\u019d\u019c\u0001\u0000\u0000\u0000\u019d\u019e\u0001\u0000"+
		"\u0000\u0000\u019e\u01a2\u0001\u0000\u0000\u0000\u019f\u01a1\u0005Y\u0000"+
		"\u0000\u01a0\u019f\u0001\u0000\u0000\u0000\u01a1\u01a4\u0001\u0000\u0000"+
		"\u0000\u01a2\u01a0\u0001\u0000\u0000\u0000\u01a2\u01a3\u0001\u0000\u0000"+
		"\u0000\u01a3\u01a5\u0001\u0000\u0000\u0000\u01a4\u01a2\u0001\u0000\u0000"+
		"\u0000\u01a5\u01a6\u0005\u0000\u0000\u0001\u01a6\u0001\u0001\u0000\u0000"+
		"\u0000\u01a7\u01b2\u00038\u001c\u0000\u01a8\u01ab\u0003\u00ccf\u0000\u01a9"+
		"\u01ac\u0003\u0004\u0002\u0000\u01aa\u01ac\u0003\u0006\u0003\u0000\u01ab"+
		"\u01a9\u0001\u0000\u0000\u0000\u01ab\u01aa\u0001\u0000\u0000\u0000\u01ab"+
		"\u01ac\u0001\u0000\u0000\u0000\u01ac\u01b2\u0001\u0000\u0000\u0000\u01ad"+
		"\u01b2\u0003\f\u0006\u0000\u01ae\u01b2\u0003\u0018\f\u0000\u01af\u01b2"+
		"\u0003$\u0012\u0000\u01b0\u01b2\u0003\b\u0004\u0000\u01b1\u01a7\u0001"+
		"\u0000\u0000\u0000\u01b1\u01a8\u0001\u0000\u0000\u0000\u01b1\u01ad\u0001"+
		"\u0000\u0000\u0000\u01b1\u01ae\u0001\u0000\u0000\u0000\u01b1\u01af\u0001"+
		"\u0000\u0000\u0000\u01b1\u01b0\u0001\u0000\u0000\u0000\u01b2\u0003\u0001"+
		"\u0000\u0000\u0000\u01b3\u01b4\u0005\f\u0000\u0000\u01b4\u01b5\u0003\u00cc"+
		"f\u0000\u01b5\u0005\u0001\u0000\u0000\u0000\u01b6\u01b7\u0005\u001b\u0000"+
		"\u0000\u01b7\u01b8\u0003\u00ccf\u0000\u01b8\u0007\u0001\u0000\u0000\u0000"+
		"\u01b9\u01ba\u0003\u0160\u00b0\u0000\u01ba\t\u0001\u0000\u0000\u0000\u01bb"+
		"\u01c7\u0003\u0002\u0001\u0000\u01bc\u01c0\u0003\u018e\u00c7\u0000\u01bd"+
		"\u01bf\u0005Y\u0000\u0000\u01be\u01bd\u0001\u0000\u0000\u0000\u01bf\u01c2"+
		"\u0001\u0000\u0000\u0000\u01c0\u01be\u0001\u0000\u0000\u0000\u01c0\u01c1"+
		"\u0001\u0000\u0000\u0000\u01c1\u01c3\u0001\u0000\u0000\u0000\u01c2\u01c0"+
		"\u0001\u0000\u0000\u0000\u01c3\u01c4\u0003\u0002\u0001\u0000\u01c4\u01c6"+
		"\u0001\u0000\u0000\u0000\u01c5\u01bc\u0001\u0000\u0000\u0000\u01c6\u01c9"+
		"\u0001\u0000\u0000\u0000\u01c7\u01c5\u0001\u0000\u0000\u0000\u01c7\u01c8"+
		"\u0001\u0000\u0000\u0000\u01c8\u01cb\u0001\u0000\u0000\u0000\u01c9\u01c7"+
		"\u0001\u0000\u0000\u0000\u01ca\u01cc\u0005&\u0000\u0000\u01cb\u01ca\u0001"+
		"\u0000\u0000\u0000\u01cb\u01cc\u0001\u0000\u0000\u0000\u01cc\u000b\u0001"+
		"\u0000\u0000\u0000\u01cd\u01d0\u0003\u000e\u0007\u0000\u01ce\u01d0\u0003"+
		"\u0010\b\u0000\u01cf\u01cd\u0001\u0000\u0000\u0000\u01cf\u01ce\u0001\u0000"+
		"\u0000\u0000\u01d0\r\u0001\u0000\u0000\u0000\u01d1\u01d2\u0005\n\u0000"+
		"\u0000\u01d2\u01d6\u0003\u00a6S\u0000\u01d3\u01d5\u0005Y\u0000\u0000\u01d4"+
		"\u01d3\u0001\u0000\u0000\u0000\u01d5\u01d8\u0001\u0000\u0000\u0000\u01d6"+
		"\u01d4\u0001\u0000\u0000\u0000\u01d6\u01d7\u0001\u0000\u0000\u0000\u01d7"+
		"\u01d9\u0001\u0000\u0000\u0000\u01d8\u01d6\u0001\u0000\u0000\u0000\u01d9"+
		"\u01da\u0005\u000e\u0000\u0000\u01da\u01de\u0003\u00ccf\u0000\u01db\u01dd"+
		"\u0005Y\u0000\u0000\u01dc\u01db\u0001\u0000\u0000\u0000\u01dd\u01e0\u0001"+
		"\u0000\u0000\u0000\u01de\u01dc\u0001\u0000\u0000\u0000\u01de\u01df\u0001"+
		"\u0000\u0000\u0000\u01df\u01e1\u0001\u0000\u0000\u0000\u01e0\u01de\u0001"+
		"\u0000\u0000\u0000\u01e1\u01e2\u0003:\u001d\u0000\u01e2\u000f\u0001\u0000"+
		"\u0000\u0000\u01e3\u01e4\u0005\u001b\u0000\u0000\u01e4\u01e8\u0003\u0012"+
		"\t\u0000\u01e5\u01e7\u0005Y\u0000\u0000\u01e6\u01e5\u0001\u0000\u0000"+
		"\u0000\u01e7\u01ea\u0001\u0000\u0000\u0000\u01e8\u01e6\u0001\u0000\u0000"+
		"\u0000\u01e8\u01e9\u0001\u0000\u0000\u0000\u01e9\u01eb\u0001\u0000\u0000"+
		"\u0000\u01ea\u01e8\u0001\u0000\u0000\u0000\u01eb\u01ec\u0003:\u001d\u0000"+
		"\u01ec\u0011\u0001\u0000\u0000\u0000\u01ed\u01f9\u0003\u0014\n\u0000\u01ee"+
		"\u01f2\u0003\u0190\u00c8\u0000\u01ef\u01f1\u0005Y\u0000\u0000\u01f0\u01ef"+
		"\u0001\u0000\u0000\u0000\u01f1\u01f4\u0001\u0000\u0000\u0000\u01f2\u01f0"+
		"\u0001\u0000\u0000\u0000\u01f2\u01f3\u0001\u0000\u0000\u0000\u01f3\u01f5"+
		"\u0001\u0000\u0000\u0000\u01f4\u01f2\u0001\u0000\u0000\u0000\u01f5\u01f6"+
		"\u0003\u0014\n\u0000\u01f6\u01f8\u0001\u0000\u0000\u0000\u01f7\u01ee\u0001"+
		"\u0000\u0000\u0000\u01f8\u01fb\u0001\u0000\u0000\u0000\u01f9\u01f7\u0001"+
		"\u0000\u0000\u0000\u01f9\u01fa\u0001\u0000\u0000\u0000\u01fa\u0013\u0001"+
		"\u0000\u0000\u0000\u01fb\u01f9\u0001\u0000\u0000\u0000\u01fc\u01ff\u0003"+
		"\u00ccf\u0000\u01fd\u01ff\u0003\u0016\u000b\u0000\u01fe\u01fc\u0001\u0000"+
		"\u0000\u0000\u01fe\u01fd\u0001\u0000\u0000\u0000\u01ff\u0015\u0001\u0000"+
		"\u0000\u0000\u0200\u0201\u0005\u001a\u0000\u0000\u0201\u0205\u0003\u00a6"+
		"S\u0000\u0202\u0204\u0005Y\u0000\u0000\u0203\u0202\u0001\u0000\u0000\u0000"+
		"\u0204\u0207\u0001\u0000\u0000\u0000\u0205\u0203\u0001\u0000\u0000\u0000"+
		"\u0205\u0206\u0001\u0000\u0000\u0000\u0206\u0208\u0001\u0000\u0000\u0000"+
		"\u0207\u0205\u0001\u0000\u0000\u0000\u0208\u0209\u0003`0\u0000\u0209\u0017"+
		"\u0001\u0000\u0000\u0000\u020a\u020d\u0003\u001a\r\u0000\u020b\u020d\u0003"+
		"\u001e\u000f\u0000\u020c\u020a\u0001\u0000\u0000\u0000\u020c\u020b\u0001"+
		"\u0000\u0000\u0000\u020d\u0019\u0001\u0000\u0000\u0000\u020e\u020f\u0005"+
		"\f\u0000\u0000\u020f\u0213\u0003\u0012\t\u0000\u0210\u0212\u0005Y\u0000"+
		"\u0000\u0211\u0210\u0001\u0000\u0000\u0000\u0212\u0215\u0001\u0000\u0000"+
		"\u0000\u0213\u0211\u0001\u0000\u0000\u0000\u0213\u0214\u0001\u0000\u0000"+
		"\u0000\u0214\u0216\u0001\u0000\u0000\u0000\u0215\u0213\u0001\u0000\u0000"+
		"\u0000\u0216\u021a\u0003:\u001d\u0000\u0217\u0219\u0005Y\u0000\u0000\u0218"+
		"\u0217\u0001\u0000\u0000\u0000\u0219\u021c\u0001\u0000\u0000\u0000\u021a"+
		"\u0218\u0001\u0000\u0000\u0000\u021a\u021b\u0001\u0000\u0000\u0000\u021b"+
		"\u021e\u0001\u0000\u0000\u0000\u021c\u021a\u0001\u0000\u0000\u0000\u021d"+
		"\u021f\u0003\u001c\u000e\u0000\u021e\u021d\u0001\u0000\u0000\u0000\u021e"+
		"\u021f\u0001\u0000\u0000\u0000\u021f\u001b\u0001\u0000\u0000\u0000\u0220"+
		"\u0224\u0005\u0007\u0000\u0000\u0221\u0223\u0005Y\u0000\u0000\u0222\u0221"+
		"\u0001\u0000\u0000\u0000\u0223\u0226\u0001\u0000\u0000\u0000\u0224\u0222"+
		"\u0001\u0000\u0000\u0000\u0224\u0225\u0001\u0000\u0000\u0000\u0225\u0227"+
		"\u0001\u0000\u0000\u0000\u0226\u0224\u0001\u0000\u0000\u0000\u0227\u0231"+
		"\u0003:\u001d\u0000\u0228\u022c\u0005\u0007\u0000\u0000\u0229\u022b\u0005"+
		"Y\u0000\u0000\u022a\u0229\u0001\u0000\u0000\u0000\u022b\u022e\u0001\u0000"+
		"\u0000\u0000\u022c\u022a\u0001\u0000\u0000\u0000\u022c\u022d\u0001\u0000"+
		"\u0000\u0000\u022d\u022f\u0001\u0000\u0000\u0000\u022e\u022c\u0001\u0000"+
		"\u0000\u0000\u022f\u0231\u0003\u001a\r\u0000\u0230\u0220\u0001\u0000\u0000"+
		"\u0000\u0230\u0228\u0001\u0000\u0000\u0000\u0231\u001d\u0001\u0000\u0000"+
		"\u0000\u0232\u0233\u0005\u0011\u0000\u0000\u0233\u0237\u0003\u00ccf\u0000"+
		"\u0234\u0236\u0005Y\u0000\u0000\u0235\u0234\u0001\u0000\u0000\u0000\u0236"+
		"\u0239\u0001\u0000\u0000\u0000\u0237\u0235\u0001\u0000\u0000\u0000\u0237"+
		"\u0238\u0001\u0000\u0000\u0000\u0238\u023a\u0001\u0000\u0000\u0000\u0239"+
		"\u0237\u0001\u0000\u0000\u0000\u023a\u0242\u0005\u001e\u0000\u0000\u023b"+
		"\u023d\u0005Y\u0000\u0000\u023c\u023b\u0001\u0000\u0000\u0000\u023d\u0240"+
		"\u0001\u0000\u0000\u0000\u023e\u023c\u0001\u0000\u0000\u0000\u023e\u023f"+
		"\u0001\u0000\u0000\u0000\u023f\u0241\u0001\u0000\u0000\u0000\u0240\u023e"+
		"\u0001\u0000\u0000\u0000\u0241\u0243\u0003 \u0010\u0000\u0242\u023e\u0001"+
		"\u0000\u0000\u0000\u0242\u0243\u0001\u0000\u0000\u0000\u0243\u0247\u0001"+
		"\u0000\u0000\u0000\u0244\u0246\u0005Y\u0000\u0000\u0245\u0244\u0001\u0000"+
		"\u0000\u0000\u0246\u0249\u0001\u0000\u0000\u0000\u0247\u0245\u0001\u0000"+
		"\u0000\u0000\u0247\u0248\u0001\u0000\u0000\u0000\u0248\u024a\u0001\u0000"+
		"\u0000\u0000\u0249\u0247\u0001\u0000\u0000\u0000\u024a\u024b\u0005!\u0000"+
		"\u0000\u024b\u001f\u0001\u0000\u0000\u0000\u024c\u0258\u0003\"\u0011\u0000"+
		"\u024d\u0251\u0003\u018e\u00c7\u0000\u024e\u0250\u0005Y\u0000\u0000\u024f"+
		"\u024e\u0001\u0000\u0000\u0000\u0250\u0253\u0001\u0000\u0000\u0000\u0251"+
		"\u024f\u0001\u0000\u0000\u0000\u0251\u0252\u0001\u0000\u0000\u0000\u0252"+
		"\u0254\u0001\u0000\u0000\u0000\u0253\u0251\u0001\u0000\u0000\u0000\u0254"+
		"\u0255\u0003\"\u0011\u0000\u0255\u0257\u0001\u0000\u0000\u0000\u0256\u024d"+
		"\u0001\u0000\u0000\u0000\u0257\u025a\u0001\u0000\u0000\u0000\u0258\u0256"+
		"\u0001\u0000\u0000\u0000\u0258\u0259\u0001\u0000\u0000\u0000\u0259\u025c"+
		"\u0001\u0000\u0000\u0000\u025a\u0258\u0001\u0000\u0000\u0000\u025b\u025d"+
		"\u0003\u018e\u00c7\u0000\u025c\u025b\u0001\u0000\u0000\u0000\u025c\u025d"+
		"\u0001\u0000\u0000\u0000\u025d!\u0001\u0000\u0000\u0000\u025e\u0260\u0003"+
		"\u00a6S\u0000\u025f\u0261\u0003\u0004\u0002\u0000\u0260\u025f\u0001\u0000"+
		"\u0000\u0000\u0260\u0261\u0001\u0000\u0000\u0000\u0261\u0265\u0001\u0000"+
		"\u0000\u0000\u0262\u0264\u0005Y\u0000\u0000\u0263\u0262\u0001\u0000\u0000"+
		"\u0000\u0264\u0267\u0001\u0000\u0000\u0000\u0265\u0263\u0001\u0000\u0000"+
		"\u0000\u0265\u0266\u0001\u0000\u0000\u0000\u0266\u0268\u0001\u0000\u0000"+
		"\u0000\u0267\u0265\u0001\u0000\u0000\u0000\u0268\u026c\u0005<\u0000\u0000"+
		"\u0269\u026b\u0005Y\u0000\u0000\u026a\u0269\u0001\u0000\u0000\u0000\u026b"+
		"\u026e\u0001\u0000\u0000\u0000\u026c\u026a\u0001\u0000\u0000\u0000\u026c"+
		"\u026d\u0001\u0000\u0000\u0000\u026d\u0271\u0001\u0000\u0000\u0000\u026e"+
		"\u026c\u0001\u0000\u0000\u0000\u026f\u0272\u0003:\u001d\u0000\u0270\u0272"+
		"\u0003\u00ccf\u0000\u0271\u026f\u0001\u0000\u0000\u0000\u0271\u0270\u0001"+
		"\u0000\u0000\u0000\u0272#\u0001\u0000\u0000\u0000\u0273\u0277\u0003&\u0013"+
		"\u0000\u0274\u0277\u0003(\u0014\u0000\u0275\u0277\u0003*\u0015\u0000\u0276"+
		"\u0273\u0001\u0000\u0000\u0000\u0276\u0274\u0001\u0000\u0000\u0000\u0276"+
		"\u0275\u0001\u0000\u0000\u0000\u0277%\u0001\u0000\u0000\u0000\u0278\u0279"+
		"\u0005\u0004\u0000\u0000\u0279\'\u0001\u0000\u0000\u0000\u027a\u027b\u0005"+
		"\u0006\u0000\u0000\u027b)\u0001\u0000\u0000\u0000\u027c\u027e\u0005\u0016"+
		"\u0000\u0000\u027d\u027f\u0003\u00ccf\u0000\u027e\u027d\u0001\u0000\u0000"+
		"\u0000\u027e\u027f\u0001\u0000\u0000\u0000\u027f+\u0001\u0000\u0000\u0000"+
		"\u0280\u0284\u0005\'\u0000\u0000\u0281\u0283\u0005Y\u0000\u0000\u0282"+
		"\u0281\u0001\u0000\u0000\u0000\u0283\u0286\u0001\u0000\u0000\u0000\u0284"+
		"\u0282\u0001\u0000\u0000\u0000\u0284\u0285\u0001\u0000\u0000\u0000\u0285"+
		"\u0287\u0001\u0000\u0000\u0000\u0286\u0284\u0001\u0000\u0000\u0000\u0287"+
		"\u028b\u0003.\u0017\u0000\u0288\u028a\u0005Y\u0000\u0000\u0289\u0288\u0001"+
		"\u0000\u0000\u0000\u028a\u028d\u0001\u0000\u0000\u0000\u028b\u0289\u0001"+
		"\u0000\u0000\u0000\u028b\u028c\u0001\u0000\u0000\u0000\u028c\u028e\u0001"+
		"\u0000\u0000\u0000\u028d\u028b\u0001\u0000\u0000\u0000\u028e\u028f\u0005"+
		"(\u0000\u0000\u028f-\u0001\u0000\u0000\u0000\u0290\u029c\u00030\u0018"+
		"\u0000\u0291\u0295\u0003\u0194\u00ca\u0000\u0292\u0294\u0005Y\u0000\u0000"+
		"\u0293\u0292\u0001\u0000\u0000\u0000\u0294\u0297\u0001\u0000\u0000\u0000"+
		"\u0295\u0293\u0001\u0000\u0000\u0000\u0295\u0296\u0001\u0000\u0000\u0000"+
		"\u0296\u0298\u0001\u0000\u0000\u0000\u0297\u0295\u0001\u0000\u0000\u0000"+
		"\u0298\u0299\u00030\u0018\u0000\u0299\u029b\u0001\u0000\u0000\u0000\u029a"+
		"\u0291\u0001\u0000\u0000\u0000\u029b\u029e\u0001\u0000\u0000\u0000\u029c"+
		"\u029a\u0001\u0000\u0000\u0000\u029c\u029d\u0001\u0000\u0000\u0000\u029d"+
		"\u02a0\u0001\u0000\u0000\u0000\u029e\u029c\u0001\u0000\u0000\u0000\u029f"+
		"\u02a1\u0003\u0194\u00ca\u0000\u02a0\u029f\u0001\u0000\u0000\u0000\u02a0"+
		"\u02a1\u0001\u0000\u0000\u0000\u02a1/\u0001\u0000\u0000\u0000\u02a2\u02aa"+
		"\u0003\u013e\u009f\u0000\u02a3\u02a4\u0003\u013e\u009f\u0000\u02a4\u02a5"+
		"\u0005A\u0000\u0000\u02a5\u02aa\u0001\u0000\u0000\u0000\u02a6\u02a7\u0003"+
		"\u013e\u009f\u0000\u02a7\u02a8\u0003\u0138\u009c\u0000\u02a8\u02aa\u0001"+
		"\u0000\u0000\u0000\u02a9\u02a2\u0001\u0000\u0000\u0000\u02a9\u02a3\u0001"+
		"\u0000\u0000\u0000\u02a9\u02a6\u0001\u0000\u0000\u0000\u02aa1\u0001\u0000"+
		"\u0000\u0000\u02ab\u02af\u0005\'\u0000\u0000\u02ac\u02ae\u0005Y\u0000"+
		"\u0000\u02ad\u02ac\u0001\u0000\u0000\u0000\u02ae\u02b1\u0001\u0000\u0000"+
		"\u0000\u02af\u02ad\u0001\u0000\u0000\u0000\u02af\u02b0\u0001\u0000\u0000"+
		"\u0000\u02b0\u02b2\u0001\u0000\u0000\u0000\u02b1\u02af\u0001\u0000\u0000"+
		"\u0000\u02b2\u02b6\u00034\u001a\u0000\u02b3\u02b5\u0005Y\u0000\u0000\u02b4"+
		"\u02b3\u0001\u0000\u0000\u0000\u02b5\u02b8\u0001\u0000\u0000\u0000\u02b6"+
		"\u02b4\u0001\u0000\u0000\u0000\u02b6\u02b7\u0001\u0000\u0000\u0000\u02b7"+
		"\u02b9\u0001\u0000\u0000\u0000\u02b8\u02b6\u0001\u0000\u0000\u0000\u02b9"+
		"\u02ba\u0005(\u0000\u0000\u02ba3\u0001\u0000\u0000\u0000\u02bb\u02c7\u0003"+
		"6\u001b\u0000\u02bc\u02c0\u0003\u0190\u00c8\u0000\u02bd\u02bf\u0005Y\u0000"+
		"\u0000\u02be\u02bd\u0001\u0000\u0000\u0000\u02bf\u02c2\u0001\u0000\u0000"+
		"\u0000\u02c0\u02be\u0001\u0000\u0000\u0000\u02c0\u02c1\u0001\u0000\u0000"+
		"\u0000\u02c1\u02c3\u0001\u0000\u0000\u0000\u02c2\u02c0\u0001\u0000\u0000"+
		"\u0000\u02c3\u02c4\u00036\u001b\u0000\u02c4\u02c6\u0001\u0000\u0000\u0000"+
		"\u02c5\u02bc\u0001\u0000\u0000\u0000\u02c6\u02c9\u0001\u0000\u0000\u0000"+
		"\u02c7\u02c5\u0001\u0000\u0000\u0000\u02c7\u02c8\u0001\u0000\u0000\u0000"+
		"\u02c85\u0001\u0000\u0000\u0000\u02c9\u02c7\u0001\u0000\u0000\u0000\u02ca"+
		"\u02cc\u0003\u0132\u0099\u0000\u02cb\u02cd\u0003\u00cae\u0000\u02cc\u02cb"+
		"\u0001\u0000\u0000\u0000\u02cc\u02cd\u0001\u0000\u0000\u0000\u02cd7\u0001"+
		"\u0000\u0000\u0000\u02ce\u02cf\u0003\u0160\u00b0\u0000\u02cf\u02d0\u0005"+
		"Y\u0000\u0000\u02d0\u02d2\u0001\u0000\u0000\u0000\u02d1\u02ce\u0001\u0000"+
		"\u0000\u0000\u02d1\u02d2\u0001\u0000\u0000\u0000\u02d2\u02da\u0001\u0000"+
		"\u0000\u0000\u02d3\u02d7\u0003\u00cae\u0000\u02d4\u02d6\u0005Y\u0000\u0000"+
		"\u02d5\u02d4\u0001\u0000\u0000\u0000\u02d6\u02d9\u0001\u0000\u0000\u0000"+
		"\u02d7\u02d5\u0001\u0000\u0000\u0000\u02d7\u02d8\u0001\u0000\u0000\u0000"+
		"\u02d8\u02db\u0001\u0000\u0000\u0000\u02d9\u02d7\u0001\u0000\u0000\u0000"+
		"\u02da\u02d3\u0001\u0000\u0000\u0000\u02da\u02db\u0001\u0000\u0000\u0000"+
		"\u02db\u02e7\u0001\u0000\u0000\u0000\u02dc\u02e8\u0003<\u001e\u0000\u02dd"+
		"\u02e8\u0003B!\u0000\u02de\u02e8\u0003X,\u0000\u02df\u02e8\u0003b1\u0000"+
		"\u02e0\u02e8\u0003d2\u0000\u02e1\u02e8\u0003j5\u0000\u02e2\u02e8\u0003"+
		"z=\u0000\u02e3\u02e8\u0003\u0084B\u0000\u02e4\u02e8\u0003\u0092I\u0000"+
		"\u02e5\u02e8\u0003\u00a0P\u0000\u02e6\u02e8\u0003\u00a2Q\u0000\u02e7\u02dc"+
		"\u0001\u0000\u0000\u0000\u02e7\u02dd\u0001\u0000\u0000\u0000\u02e7\u02de"+
		"\u0001\u0000\u0000\u0000\u02e7\u02df\u0001\u0000\u0000\u0000\u02e7\u02e0"+
		"\u0001\u0000\u0000\u0000\u02e7\u02e1\u0001\u0000\u0000\u0000\u02e7\u02e2"+
		"\u0001\u0000\u0000\u0000\u02e7\u02e3\u0001\u0000\u0000\u0000\u02e7\u02e4"+
		"\u0001\u0000\u0000\u0000\u02e7\u02e5\u0001\u0000\u0000\u0000\u02e7\u02e6"+
		"\u0001\u0000\u0000\u0000\u02e89\u0001\u0000\u0000\u0000\u02e9\u02f1\u0005"+
		"\u001e\u0000\u0000\u02ea\u02ec\u0005Y\u0000\u0000\u02eb\u02ea\u0001\u0000"+
		"\u0000\u0000\u02ec\u02ef\u0001\u0000\u0000\u0000\u02ed\u02eb\u0001\u0000"+
		"\u0000\u0000\u02ed\u02ee\u0001\u0000\u0000\u0000\u02ee\u02f0\u0001\u0000"+
		"\u0000\u0000\u02ef\u02ed\u0001\u0000\u0000\u0000\u02f0\u02f2\u0003\n\u0005"+
		"\u0000\u02f1\u02ed\u0001\u0000\u0000\u0000\u02f1\u02f2\u0001\u0000\u0000"+
		"\u0000\u02f2\u02f6\u0001\u0000\u0000\u0000\u02f3\u02f5\u0005Y\u0000\u0000"+
		"\u02f4\u02f3\u0001\u0000\u0000\u0000\u02f5\u02f8\u0001\u0000\u0000\u0000"+
		"\u02f6\u02f4\u0001\u0000\u0000\u0000\u02f6\u02f7\u0001\u0000\u0000\u0000"+
		"\u02f7\u02f9\u0001\u0000\u0000\u0000\u02f8\u02f6\u0001\u0000\u0000\u0000"+
		"\u02f9\u02fa\u0005!\u0000\u0000\u02fa;\u0001\u0000\u0000\u0000\u02fb\u02fc"+
		"\u0005\u0015\u0000\u0000\u02fc\u0304\u0003>\u001f\u0000\u02fd\u02ff\u0005"+
		"Y\u0000\u0000\u02fe\u02fd\u0001\u0000\u0000\u0000\u02ff\u0302\u0001\u0000"+
		"\u0000\u0000\u0300\u02fe\u0001\u0000\u0000\u0000\u0300\u0301\u0001\u0000"+
		"\u0000\u0000\u0301\u0303\u0001\u0000\u0000\u0000\u0302\u0300\u0001\u0000"+
		"\u0000\u0000\u0303\u0305\u0003\u00f8|\u0000\u0304\u0300\u0001\u0000\u0000"+
		"\u0000\u0304\u0305\u0001\u0000\u0000\u0000\u0305=\u0001\u0000\u0000\u0000"+
		"\u0306\u030b\u0003@ \u0000\u0307\u0308\u0005\u001d\u0000\u0000\u0308\u030a"+
		"\u0003@ \u0000\u0309\u0307\u0001\u0000\u0000\u0000\u030a\u030d\u0001\u0000"+
		"\u0000\u0000\u030b\u0309\u0001\u0000\u0000\u0000\u030b\u030c\u0001\u0000"+
		"\u0000\u0000\u030c?\u0001\u0000\u0000\u0000\u030d\u030b\u0001\u0000\u0000"+
		"\u0000\u030e\u030f\u0005J\u0000\u0000\u030fA\u0001\u0000\u0000\u0000\u0310"+
		"\u0311\u0005\r\u0000\u0000\u0311\u0316\u0003D\"\u0000\u0312\u0317\u0003"+
		"H$\u0000\u0313\u0317\u0003J%\u0000\u0314\u0317\u0003L&\u0000\u0315\u0317"+
		"\u0003P(\u0000\u0316\u0312\u0001\u0000\u0000\u0000\u0316\u0313\u0001\u0000"+
		"\u0000\u0000\u0316\u0314\u0001\u0000\u0000\u0000\u0316\u0315\u0001\u0000"+
		"\u0000\u0000\u0316\u0317\u0001\u0000\u0000\u0000\u0317C\u0001\u0000\u0000"+
		"\u0000\u0318\u031d\u0003F#\u0000\u0319\u031a\u0005\u001d\u0000\u0000\u031a"+
		"\u031c\u0003F#\u0000\u031b\u0319\u0001\u0000\u0000\u0000\u031c\u031f\u0001"+
		"\u0000\u0000\u0000\u031d\u031b\u0001\u0000\u0000\u0000\u031d\u031e\u0001"+
		"\u0000\u0000\u0000\u031eE\u0001\u0000\u0000\u0000\u031f\u031d\u0001\u0000"+
		"\u0000\u0000\u0320\u0321\u0003\u0154\u00aa\u0000\u0321G\u0001\u0000\u0000"+
		"\u0000\u0322\u0323\u0005\u001d\u0000\u0000\u0323\u0324\u00052\u0000\u0000"+
		"\u0324I\u0001\u0000\u0000\u0000\u0325\u0326\u0005\u0002\u0000\u0000\u0326"+
		"\u0327\u0003\u0154\u00aa\u0000\u0327K\u0001\u0000\u0000\u0000\u0328\u0329"+
		"\u0005\u001d\u0000\u0000\u0329\u032b\u0003\u013e\u009f\u0000\u032a\u032c"+
		"\u0003N\'\u0000\u032b\u032a\u0001\u0000\u0000\u0000\u032b\u032c\u0001"+
		"\u0000\u0000\u0000\u032cM\u0001\u0000\u0000\u0000\u032d\u032e\u0005\u0002"+
		"\u0000\u0000\u032e\u032f\u0003\u013e\u009f\u0000\u032fO\u0001\u0000\u0000"+
		"\u0000\u0330\u0331\u0005\u001d\u0000\u0000\u0331\u0335\u0005\u001e\u0000"+
		"\u0000\u0332\u0334\u0005Y\u0000\u0000\u0333\u0332\u0001\u0000\u0000\u0000"+
		"\u0334\u0337\u0001\u0000\u0000\u0000\u0335\u0333\u0001\u0000\u0000\u0000"+
		"\u0335\u0336\u0001\u0000\u0000\u0000\u0336\u0338\u0001\u0000\u0000\u0000"+
		"\u0337\u0335\u0001\u0000\u0000\u0000\u0338\u033c\u0003R)\u0000\u0339\u033b"+
		"\u0005Y\u0000\u0000\u033a\u0339\u0001\u0000\u0000\u0000\u033b\u033e\u0001"+
		"\u0000\u0000\u0000\u033c\u033a\u0001\u0000\u0000\u0000\u033c\u033d\u0001"+
		"\u0000\u0000\u0000\u033d\u033f\u0001\u0000\u0000\u0000\u033e\u033c\u0001"+
		"\u0000\u0000\u0000\u033f\u0340\u0005!\u0000\u0000\u0340Q\u0001\u0000\u0000"+
		"\u0000\u0341\u0344\u0003T*\u0000\u0342\u0344\u0003V+\u0000\u0343\u0341"+
		"\u0001\u0000\u0000\u0000\u0343\u0342\u0001\u0000\u0000\u0000\u0344\u0352"+
		"\u0001\u0000\u0000\u0000\u0345\u0349\u0003\u0190\u00c8\u0000\u0346\u0348"+
		"\u0005Y\u0000\u0000\u0347\u0346\u0001\u0000\u0000\u0000\u0348\u034b\u0001"+
		"\u0000\u0000\u0000\u0349\u0347\u0001\u0000\u0000\u0000\u0349\u034a\u0001"+
		"\u0000\u0000\u0000\u034a\u034e\u0001\u0000\u0000\u0000\u034b\u0349\u0001"+
		"\u0000\u0000\u0000\u034c\u034f\u0003T*\u0000\u034d\u034f\u0003V+\u0000"+
		"\u034e\u034c\u0001\u0000\u0000\u0000\u034e\u034d\u0001\u0000\u0000\u0000"+
		"\u034f\u0351\u0001\u0000\u0000\u0000\u0350\u0345\u0001\u0000\u0000\u0000"+
		"\u0351\u0354\u0001\u0000\u0000\u0000\u0352\u0350\u0001\u0000\u0000\u0000"+
		"\u0352\u0353\u0001\u0000\u0000\u0000\u0353\u0356\u0001\u0000\u0000\u0000"+
		"\u0354\u0352\u0001\u0000\u0000\u0000\u0355\u0357\u0003\u0190\u00c8\u0000"+
		"\u0356\u0355\u0001\u0000\u0000\u0000\u0356\u0357\u0001\u0000\u0000\u0000"+
		"\u0357S\u0001\u0000\u0000\u0000\u0358\u035a\u0003\u0154\u00aa\u0000\u0359"+
		"\u035b\u0003J%\u0000\u035a\u0359\u0001\u0000\u0000\u0000\u035a\u035b\u0001"+
		"\u0000\u0000\u0000\u035bU\u0001\u0000\u0000\u0000\u035c\u035e\u0003\u013e"+
		"\u009f\u0000\u035d\u035f\u0003N\'\u0000\u035e\u035d\u0001\u0000\u0000"+
		"\u0000\u035e\u035f\u0001\u0000\u0000\u0000\u035fW\u0001\u0000\u0000\u0000"+
		"\u0360\u0361\u0005\u0005\u0000\u0000\u0361\u0385\u0003Z-\u0000\u0362\u0363"+
		"\u0005\u0005\u0000\u0000\u0363\u0367\u0005\u001e\u0000\u0000\u0364\u0366"+
		"\u0005Y\u0000\u0000\u0365\u0364\u0001\u0000\u0000\u0000\u0366\u0369\u0001"+
		"\u0000\u0000\u0000\u0367\u0365\u0001\u0000\u0000\u0000\u0367\u0368\u0001"+
		"\u0000\u0000\u0000\u0368\u036a\u0001\u0000\u0000\u0000\u0369\u0367\u0001"+
		"\u0000\u0000\u0000\u036a\u0376\u0003\\.\u0000\u036b\u036f\u0003\u018e"+
		"\u00c7\u0000\u036c\u036e\u0005Y\u0000\u0000\u036d\u036c\u0001\u0000\u0000"+
		"\u0000\u036e\u0371\u0001\u0000\u0000\u0000\u036f\u036d\u0001\u0000\u0000"+
		"\u0000\u036f\u0370\u0001\u0000\u0000\u0000\u0370\u0372\u0001\u0000\u0000"+
		"\u0000\u0371\u036f\u0001\u0000\u0000\u0000\u0372\u0373\u0003\\.\u0000"+
		"\u0373\u0375\u0001\u0000\u0000\u0000\u0374\u036b\u0001\u0000\u0000\u0000"+
		"\u0375\u0378\u0001\u0000\u0000\u0000\u0376\u0374\u0001\u0000\u0000\u0000"+
		"\u0376\u0377\u0001\u0000\u0000\u0000\u0377\u037a\u0001\u0000\u0000\u0000"+
		"\u0378\u0376\u0001\u0000\u0000\u0000\u0379\u037b\u0003\u018e\u00c7\u0000"+
		"\u037a\u0379\u0001\u0000\u0000\u0000\u037a\u037b\u0001\u0000\u0000\u0000"+
		"\u037b\u037f\u0001\u0000\u0000\u0000\u037c\u037e\u0005Y\u0000\u0000\u037d"+
		"\u037c\u0001\u0000\u0000\u0000\u037e\u0381\u0001\u0000\u0000\u0000\u037f"+
		"\u037d\u0001\u0000\u0000\u0000\u037f\u0380\u0001\u0000\u0000\u0000\u0380"+
		"\u0382\u0001\u0000\u0000\u0000\u0381\u037f\u0001\u0000\u0000\u0000\u0382"+
		"\u0383\u0005!\u0000\u0000\u0383\u0385\u0001\u0000\u0000\u0000\u0384\u0360"+
		"\u0001\u0000\u0000\u0000\u0384\u0362\u0001\u0000\u0000\u0000\u0385Y\u0001"+
		"\u0000\u0000\u0000\u0386\u0387\u0003^/\u0000\u0387[\u0001\u0000\u0000"+
		"\u0000\u0388\u0389\u0003\u0160\u00b0\u0000\u0389\u038a\u0005Y\u0000\u0000"+
		"\u038a\u038c\u0001\u0000\u0000\u0000\u038b\u0388\u0001\u0000\u0000\u0000"+
		"\u038b\u038c\u0001\u0000\u0000\u0000\u038c\u0390\u0001\u0000\u0000\u0000"+
		"\u038d\u038e\u0003\u00cae\u0000\u038e\u038f\u0005Y\u0000\u0000\u038f\u0391"+
		"\u0001\u0000\u0000\u0000\u0390\u038d\u0001\u0000\u0000\u0000\u0390\u0391"+
		"\u0001\u0000\u0000\u0000\u0391\u0392\u0001\u0000\u0000\u0000\u0392\u0393"+
		"\u0003^/\u0000\u0393]\u0001\u0000\u0000\u0000\u0394\u0396\u0003\u00a6"+
		"S\u0000\u0395\u0397\u0003`0\u0000\u0396\u0395\u0001\u0000\u0000\u0000"+
		"\u0396\u0397\u0001\u0000\u0000\u0000\u0397_\u0001\u0000\u0000\u0000\u0398"+
		"\u039c\u0003\u0164\u00b2\u0000\u0399\u039b\u0005Y\u0000\u0000\u039a\u0399"+
		"\u0001\u0000\u0000\u0000\u039b\u039e\u0001\u0000\u0000\u0000\u039c\u039a"+
		"\u0001\u0000\u0000\u0000\u039c\u039d\u0001\u0000\u0000\u0000\u039d\u039f"+
		"\u0001\u0000\u0000\u0000\u039e\u039c\u0001\u0000\u0000\u0000\u039f\u03a0"+
		"\u0003\u00ccf\u0000\u03a0a\u0001\u0000\u0000\u0000\u03a1\u03a2\u0005\u001a"+
		"\u0000\u0000\u03a2\u03ca\u0003Z-\u0000\u03a3\u03a4\u0003\u00a6S\u0000"+
		"\u03a4\u03a5\u0005;\u0000\u0000\u03a5\u03a6\u0003\u00ccf\u0000\u03a6\u03ca"+
		"\u0001\u0000\u0000\u0000\u03a7\u03a8\u0005\u001a\u0000\u0000\u03a8\u03ac"+
		"\u0005\u001e\u0000\u0000\u03a9\u03ab\u0005Y\u0000\u0000\u03aa\u03a9\u0001"+
		"\u0000\u0000\u0000\u03ab\u03ae\u0001\u0000\u0000\u0000\u03ac\u03aa\u0001"+
		"\u0000\u0000\u0000\u03ac\u03ad\u0001\u0000\u0000\u0000\u03ad\u03af\u0001"+
		"\u0000\u0000\u0000\u03ae\u03ac\u0001\u0000\u0000\u0000\u03af\u03bb\u0003"+
		"\\.\u0000\u03b0\u03b4\u0003\u018e\u00c7\u0000\u03b1\u03b3\u0005Y\u0000"+
		"\u0000\u03b2\u03b1\u0001\u0000\u0000\u0000\u03b3\u03b6\u0001\u0000\u0000"+
		"\u0000\u03b4\u03b2\u0001\u0000\u0000\u0000\u03b4\u03b5\u0001\u0000\u0000"+
		"\u0000\u03b5\u03b7\u0001\u0000\u0000\u0000\u03b6\u03b4\u0001\u0000\u0000"+
		"\u0000\u03b7\u03b8\u0003\\.\u0000\u03b8\u03ba\u0001\u0000\u0000\u0000"+
		"\u03b9\u03b0\u0001\u0000\u0000\u0000\u03ba\u03bd\u0001\u0000\u0000\u0000"+
		"\u03bb\u03b9\u0001\u0000\u0000\u0000\u03bb\u03bc\u0001\u0000\u0000\u0000"+
		"\u03bc\u03bf\u0001\u0000\u0000\u0000\u03bd\u03bb\u0001\u0000\u0000\u0000"+
		"\u03be\u03c0\u0003\u018e\u00c7\u0000\u03bf\u03be\u0001\u0000\u0000\u0000"+
		"\u03bf\u03c0\u0001\u0000\u0000\u0000\u03c0\u03c4\u0001\u0000\u0000\u0000"+
		"\u03c1\u03c3\u0005Y\u0000\u0000\u03c2\u03c1\u0001\u0000\u0000\u0000\u03c3"+
		"\u03c6\u0001\u0000\u0000\u0000\u03c4\u03c2\u0001\u0000\u0000\u0000\u03c4"+
		"\u03c5\u0001\u0000\u0000\u0000\u03c5\u03c7\u0001\u0000\u0000\u0000\u03c6"+
		"\u03c4\u0001\u0000\u0000\u0000\u03c7\u03c8\u0005!\u0000\u0000\u03c8\u03ca"+
		"\u0001\u0000\u0000\u0000\u03c9\u03a1\u0001\u0000\u0000\u0000\u03c9\u03a3"+
		"\u0001\u0000\u0000\u0000\u03c9\u03a7\u0001\u0000\u0000\u0000\u03cac\u0001"+
		"\u0000\u0000\u0000\u03cb\u03cc\u0005\u0019\u0000\u0000\u03cc\u03ce\u0003"+
		"f3\u0000\u03cd\u03cf\u0003,\u0016\u0000\u03ce\u03cd\u0001\u0000\u0000"+
		"\u0000\u03ce\u03cf\u0001\u0000\u0000\u0000\u03cf\u03d3\u0001\u0000\u0000"+
		"\u0000\u03d0\u03d2\u0005Y\u0000\u0000\u03d1\u03d0\u0001\u0000\u0000\u0000"+
		"\u03d2\u03d5\u0001\u0000\u0000\u0000\u03d3\u03d1\u0001\u0000\u0000\u0000"+
		"\u03d3\u03d4\u0001\u0000\u0000\u0000\u03d4\u03d6\u0001\u0000\u0000\u0000"+
		"\u03d5\u03d3\u0001\u0000\u0000\u0000\u03d6\u03d7\u0003h4\u0000\u03d7e"+
		"\u0001\u0000\u0000\u0000\u03d8\u03d9\u0003\u013e\u009f\u0000\u03d9g\u0001"+
		"\u0000\u0000\u0000\u03da\u03de\u0003\u0164\u00b2\u0000\u03db\u03dd\u0005"+
		"Y\u0000\u0000\u03dc\u03db\u0001\u0000\u0000\u0000\u03dd\u03e0\u0001\u0000"+
		"\u0000\u0000\u03de\u03dc\u0001\u0000\u0000\u0000\u03de\u03df\u0001\u0000"+
		"\u0000\u0000\u03df\u03e1\u0001\u0000\u0000\u0000\u03e0\u03de\u0001\u0000"+
		"\u0000\u0000\u03e1\u03e3\u0003\u0132\u0099\u0000\u03e2\u03e4\u0003\u00ca"+
		"e\u0000\u03e3\u03e2\u0001\u0000\u0000\u0000\u03e3\u03e4\u0001\u0000\u0000"+
		"\u0000\u03e4\u03e6\u0001\u0000\u0000\u0000\u03e5\u03e7\u0003\u0162\u00b1"+
		"\u0000\u03e6\u03e5\u0001\u0000\u0000\u0000\u03e6\u03e7\u0001\u0000\u0000"+
		"\u0000\u03e7i\u0001\u0000\u0000\u0000\u03e8\u03e9\u0005\u000b\u0000\u0000"+
		"\u03e9\u03eb\u0003l6\u0000\u03ea\u03ec\u0003,\u0016\u0000\u03eb\u03ea"+
		"\u0001\u0000\u0000\u0000\u03eb\u03ec\u0001\u0000\u0000\u0000\u03ec\u03ed"+
		"\u0001\u0000\u0000\u0000\u03ed\u03f5\u0003n7\u0000\u03ee\u03f0\u0005Y"+
		"\u0000\u0000\u03ef\u03ee\u0001\u0000\u0000\u0000\u03f0\u03f3\u0001\u0000"+
		"\u0000\u0000\u03f1\u03ef\u0001\u0000\u0000\u0000\u03f1\u03f2\u0001\u0000"+
		"\u0000\u0000\u03f2\u03f4\u0001\u0000\u0000\u0000\u03f3\u03f1\u0001\u0000"+
		"\u0000\u0000\u03f4\u03f6\u0003r9\u0000\u03f5\u03f1\u0001\u0000\u0000\u0000"+
		"\u03f5\u03f6\u0001\u0000\u0000\u0000\u03f6k\u0001\u0000\u0000\u0000\u03f7"+
		"\u03fa\u0003\u0154\u00aa\u0000\u03f8\u03fa\u0003\u0176\u00bb\u0000\u03f9"+
		"\u03f7\u0001\u0000\u0000\u0000\u03f9\u03f8\u0001\u0000\u0000\u0000\u03fa"+
		"m\u0001\u0000\u0000\u0000\u03fb\u03fd\u0003t:\u0000\u03fc\u03fe\u0003"+
		"\u0162\u00b1\u0000\u03fd\u03fc\u0001\u0000\u0000\u0000\u03fd\u03fe\u0001"+
		"\u0000\u0000\u0000\u03fe\u0406\u0001\u0000\u0000\u0000\u03ff\u0401\u0005"+
		"Y\u0000\u0000\u0400\u03ff\u0001\u0000\u0000\u0000\u0401\u0404\u0001\u0000"+
		"\u0000\u0000\u0402\u0400\u0001\u0000\u0000\u0000\u0402\u0403\u0001\u0000"+
		"\u0000\u0000\u0403\u0405\u0001\u0000\u0000\u0000\u0404\u0402\u0001\u0000"+
		"\u0000\u0000\u0405\u0407\u0003p8\u0000\u0406\u0402\u0001\u0000\u0000\u0000"+
		"\u0406\u0407\u0001\u0000\u0000\u0000\u0407o\u0001\u0000\u0000\u0000\u0408"+
		"\u040d\u0003\u0168\u00b4\u0000\u0409\u040b\u0003\u0156\u00ab\u0000\u040a"+
		"\u040c\u0005%\u0000\u0000\u040b\u040a\u0001\u0000\u0000\u0000\u040b\u040c"+
		"\u0001\u0000\u0000\u0000\u040c\u040e\u0001\u0000\u0000\u0000\u040d\u0409"+
		"\u0001\u0000\u0000\u0000\u040d\u040e\u0001\u0000\u0000\u0000\u040e\u040f"+
		"\u0001\u0000\u0000\u0000\u040f\u0411\u0003\u0132\u0099\u0000\u0410\u0412"+
		"\u0003\u00cae\u0000\u0411\u0410\u0001\u0000\u0000\u0000\u0411\u0412\u0001"+
		"\u0000\u0000\u0000\u0412\u041a\u0001\u0000\u0000\u0000\u0413\u0415\u0005"+
		"Y\u0000\u0000\u0414\u0413\u0001\u0000\u0000\u0000\u0415\u0418\u0001\u0000"+
		"\u0000\u0000\u0416\u0414\u0001\u0000\u0000\u0000\u0416\u0417\u0001\u0000"+
		"\u0000\u0000\u0417\u0419\u0001\u0000\u0000\u0000\u0418\u0416\u0001\u0000"+
		"\u0000\u0000\u0419\u041b\u0003\u0162\u00b1\u0000\u041a\u0416\u0001\u0000"+
		"\u0000\u0000\u041a\u041b\u0001\u0000\u0000\u0000\u041bq\u0001\u0000\u0000"+
		"\u0000\u041c\u041e\u0005\u001e\u0000\u0000\u041d\u041f\u0003\u0162\u00b1"+
		"\u0000\u041e\u041d\u0001\u0000\u0000\u0000\u041e\u041f\u0001\u0000\u0000"+
		"\u0000\u041f\u0427\u0001\u0000\u0000\u0000\u0420\u0422\u0005Y\u0000\u0000"+
		"\u0421\u0420\u0001\u0000\u0000\u0000\u0422\u0425\u0001\u0000\u0000\u0000"+
		"\u0423\u0421\u0001\u0000\u0000\u0000\u0423\u0424\u0001\u0000\u0000\u0000"+
		"\u0424\u0426\u0001\u0000\u0000\u0000\u0425\u0423\u0001\u0000\u0000\u0000"+
		"\u0426\u0428\u0003\n\u0005\u0000\u0427\u0423\u0001\u0000\u0000\u0000\u0427"+
		"\u0428\u0001\u0000\u0000\u0000\u0428\u042c\u0001\u0000\u0000\u0000\u0429"+
		"\u042b\u0005Y\u0000\u0000\u042a\u0429\u0001\u0000\u0000\u0000\u042b\u042e"+
		"\u0001\u0000\u0000\u0000\u042c\u042a\u0001\u0000\u0000\u0000\u042c\u042d"+
		"\u0001\u0000\u0000\u0000\u042d\u042f\u0001\u0000\u0000\u0000\u042e\u042c"+
		"\u0001\u0000\u0000\u0000\u042f\u0430\u0005!\u0000\u0000\u0430s\u0001\u0000"+
		"\u0000\u0000\u0431\u0432\u0005\u001f\u0000\u0000\u0432\u0444\u0005\"\u0000"+
		"\u0000\u0433\u0437\u0005\u001f\u0000\u0000\u0434\u0436\u0005Y\u0000\u0000"+
		"\u0435\u0434\u0001\u0000\u0000\u0000\u0436\u0439\u0001\u0000\u0000\u0000"+
		"\u0437\u0435\u0001\u0000\u0000\u0000\u0437\u0438\u0001\u0000\u0000\u0000"+
		"\u0438\u043a\u0001\u0000\u0000\u0000\u0439\u0437\u0001\u0000\u0000\u0000"+
		"\u043a\u043e\u0003v;\u0000\u043b\u043d\u0005Y\u0000\u0000\u043c\u043b"+
		"\u0001\u0000\u0000\u0000\u043d\u0440\u0001\u0000\u0000\u0000\u043e\u043c"+
		"\u0001\u0000\u0000\u0000\u043e\u043f\u0001\u0000\u0000\u0000\u043f\u0441"+
		"\u0001\u0000\u0000\u0000\u0440\u043e\u0001\u0000\u0000\u0000\u0441\u0442"+
		"\u0005\"\u0000\u0000\u0442\u0444\u0001\u0000\u0000\u0000\u0443\u0431\u0001"+
		"\u0000\u0000\u0000\u0443\u0433\u0001\u0000\u0000\u0000\u0444u\u0001\u0000"+
		"\u0000\u0000\u0445\u0451\u0003x<\u0000\u0446\u044a\u0003\u0194\u00ca\u0000"+
		"\u0447\u0449\u0005Y\u0000\u0000\u0448\u0447\u0001\u0000\u0000\u0000\u0449"+
		"\u044c\u0001\u0000\u0000\u0000\u044a\u0448\u0001\u0000\u0000\u0000\u044a"+
		"\u044b\u0001\u0000\u0000\u0000\u044b\u044d\u0001\u0000\u0000\u0000\u044c"+
		"\u044a\u0001\u0000\u0000\u0000\u044d\u044e\u0003x<\u0000\u044e\u0450\u0001"+
		"\u0000\u0000\u0000\u044f\u0446\u0001\u0000\u0000\u0000\u0450\u0453\u0001"+
		"\u0000\u0000\u0000\u0451\u044f\u0001\u0000\u0000\u0000\u0451\u0452\u0001"+
		"\u0000\u0000\u0000\u0452\u0455\u0001\u0000\u0000\u0000\u0453\u0451\u0001"+
		"\u0000\u0000\u0000\u0454\u0456\u0003\u0194\u00ca\u0000\u0455\u0454\u0001"+
		"\u0000\u0000\u0000\u0455\u0456\u0001\u0000\u0000\u0000\u0456w\u0001\u0000"+
		"\u0000\u0000\u0457\u0458\u0003\u0156\u00ab\u0000\u0458\u0460\u0003\u0138"+
		"\u009c\u0000\u0459\u045b\u0005Y\u0000\u0000\u045a\u0459\u0001\u0000\u0000"+
		"\u0000\u045b\u045e\u0001\u0000\u0000\u0000\u045c\u045a\u0001\u0000\u0000"+
		"\u0000\u045c\u045d\u0001\u0000\u0000\u0000\u045d\u045f\u0001\u0000\u0000"+
		"\u0000\u045e\u045c\u0001\u0000\u0000\u0000\u045f\u0461\u0003`0\u0000\u0460"+
		"\u045c\u0001\u0000\u0000\u0000\u0460\u0461\u0001\u0000\u0000\u0000\u0461"+
		"\u046c\u0001\u0000\u0000\u0000\u0462\u0464\u0003\u0156\u00ab\u0000\u0463"+
		"\u0465\u0005%\u0000\u0000\u0464\u0463\u0001\u0000\u0000\u0000\u0464\u0465"+
		"\u0001\u0000\u0000\u0000\u0465\u0466\u0001\u0000\u0000\u0000\u0466\u0467"+
		"\u0003\u0132\u0099\u0000\u0467\u0469\u0005A\u0000\u0000\u0468\u046a\u0003"+
		"\u00cae\u0000\u0469\u0468\u0001\u0000\u0000\u0000\u0469\u046a\u0001\u0000"+
		"\u0000\u0000\u046a\u046c\u0001\u0000\u0000\u0000\u046b\u0457\u0001\u0000"+
		"\u0000\u0000\u046b\u0462\u0001\u0000\u0000\u0000\u046cy\u0001\u0000\u0000"+
		"\u0000\u046d\u046e\u0005\b\u0000\u0000\u046e\u0470\u0003~?\u0000\u046f"+
		"\u0471\u0003,\u0016\u0000\u0470\u046f\u0001\u0000\u0000\u0000\u0470\u0471"+
		"\u0001\u0000\u0000\u0000\u0471\u0479\u0001\u0000\u0000\u0000\u0472\u0474"+
		"\u0005Y\u0000\u0000\u0473\u0472\u0001\u0000\u0000\u0000\u0474\u0477\u0001"+
		"\u0000\u0000\u0000\u0475\u0473\u0001\u0000\u0000\u0000\u0475\u0476\u0001"+
		"\u0000\u0000\u0000\u0476\u0478\u0001\u0000\u0000\u0000\u0477\u0475\u0001"+
		"\u0000\u0000\u0000\u0478\u047a\u0003\u014e\u00a7\u0000\u0479\u0475\u0001"+
		"\u0000\u0000\u0000\u0479\u047a\u0001\u0000\u0000\u0000\u047a\u047e\u0001"+
		"\u0000\u0000\u0000\u047b\u047d\u0005Y\u0000\u0000\u047c\u047b\u0001\u0000"+
		"\u0000\u0000\u047d\u0480\u0001\u0000\u0000\u0000\u047e\u047c\u0001\u0000"+
		"\u0000\u0000\u047e\u047f\u0001\u0000\u0000\u0000\u047f\u0481\u0001\u0000"+
		"\u0000\u0000\u0480\u047e\u0001\u0000\u0000\u0000\u0481\u0482\u0003|>\u0000"+
		"\u0482\u049c\u0001\u0000\u0000\u0000\u0483\u0484\u0005\u0019\u0000\u0000"+
		"\u0484\u0485\u0003~?\u0000\u0485\u0486\u0003\u0164\u00b2\u0000\u0486\u0488"+
		"\u0005\b\u0000\u0000\u0487\u0489\u0003,\u0016\u0000\u0488\u0487\u0001"+
		"\u0000\u0000\u0000\u0488\u0489\u0001\u0000\u0000\u0000\u0489\u0491\u0001"+
		"\u0000\u0000\u0000\u048a\u048c\u0005Y\u0000\u0000\u048b\u048a\u0001\u0000"+
		"\u0000\u0000\u048c\u048f\u0001\u0000\u0000\u0000\u048d\u048b\u0001\u0000"+
		"\u0000\u0000\u048d\u048e\u0001\u0000\u0000\u0000\u048e\u0490\u0001\u0000"+
		"\u0000\u0000\u048f\u048d\u0001\u0000\u0000\u0000\u0490\u0492\u0003\u014e"+
		"\u00a7\u0000\u0491\u048d\u0001\u0000\u0000\u0000\u0491\u0492\u0001\u0000"+
		"\u0000\u0000\u0492\u0496\u0001\u0000\u0000\u0000\u0493\u0495\u0005Y\u0000"+
		"\u0000\u0494\u0493\u0001\u0000\u0000\u0000\u0495\u0498\u0001\u0000\u0000"+
		"\u0000\u0496\u0494\u0001\u0000\u0000\u0000\u0496\u0497\u0001\u0000\u0000"+
		"\u0000\u0497\u0499\u0001\u0000\u0000\u0000\u0498\u0496\u0001\u0000\u0000"+
		"\u0000\u0499\u049a\u0003|>\u0000\u049a\u049c\u0001\u0000\u0000\u0000\u049b"+
		"\u046d\u0001\u0000\u0000\u0000\u049b\u0483\u0001\u0000\u0000\u0000\u049c"+
		"{\u0001\u0000\u0000\u0000\u049d\u049f\u0005\u001e\u0000\u0000\u049e\u04a0"+
		"\u0003\u0162\u00b1\u0000\u049f\u049e\u0001\u0000\u0000\u0000\u049f\u04a0"+
		"\u0001\u0000\u0000\u0000\u04a0\u04a8\u0001\u0000\u0000\u0000\u04a1\u04a3"+
		"\u0005Y\u0000\u0000\u04a2\u04a1\u0001\u0000\u0000\u0000\u04a3\u04a6\u0001"+
		"\u0000\u0000\u0000\u04a4\u04a2\u0001\u0000\u0000\u0000\u04a4\u04a5\u0001"+
		"\u0000\u0000\u0000\u04a5\u04a7\u0001\u0000\u0000\u0000\u04a6\u04a4\u0001"+
		"\u0000\u0000\u0000\u04a7\u04a9\u0003\u0080@\u0000\u04a8\u04a4\u0001\u0000"+
		"\u0000\u0000\u04a8\u04a9\u0001\u0000\u0000\u0000\u04a9\u04ad\u0001\u0000"+
		"\u0000\u0000\u04aa\u04ac\u0005Y\u0000\u0000\u04ab\u04aa\u0001\u0000\u0000"+
		"\u0000\u04ac\u04af\u0001\u0000\u0000\u0000\u04ad\u04ab\u0001\u0000\u0000"+
		"\u0000\u04ad\u04ae\u0001\u0000\u0000\u0000\u04ae\u04b0\u0001\u0000\u0000"+
		"\u0000\u04af\u04ad\u0001\u0000\u0000\u0000\u04b0\u04b1\u0005!\u0000\u0000"+
		"\u04b1}\u0001\u0000\u0000\u0000\u04b2\u04b3\u0003\u013e\u009f\u0000\u04b3"+
		"\u007f\u0001\u0000\u0000\u0000\u04b4\u04c0\u0003\u0082A\u0000\u04b5\u04b9"+
		"\u0003\u0192\u00c9\u0000\u04b6\u04b8\u0005Y\u0000\u0000\u04b7\u04b6\u0001"+
		"\u0000\u0000\u0000\u04b8\u04bb\u0001\u0000\u0000\u0000\u04b9\u04b7\u0001"+
		"\u0000\u0000\u0000\u04b9\u04ba\u0001\u0000\u0000\u0000\u04ba\u04bc\u0001"+
		"\u0000\u0000\u0000\u04bb\u04b9\u0001\u0000\u0000\u0000\u04bc\u04bd\u0003"+
		"\u0082A\u0000\u04bd\u04bf\u0001\u0000\u0000\u0000\u04be\u04b5\u0001\u0000"+
		"\u0000\u0000\u04bf\u04c2\u0001\u0000\u0000\u0000\u04c0\u04be\u0001\u0000"+
		"\u0000\u0000\u04c0\u04c1\u0001\u0000\u0000\u0000\u04c1\u04c4\u0001\u0000"+
		"\u0000\u0000\u04c2\u04c0\u0001\u0000\u0000\u0000\u04c3\u04c5\u0003\u0192"+
		"\u00c9\u0000\u04c4\u04c3\u0001\u0000\u0000\u0000\u04c4\u04c5\u0001\u0000"+
		"\u0000\u0000\u04c5\u0081\u0001\u0000\u0000\u0000\u04c6\u04c7\u0003\u0160"+
		"\u00b0\u0000\u04c7\u04c8\u0005Y\u0000\u0000\u04c8\u04ca\u0001\u0000\u0000"+
		"\u0000\u04c9\u04c6\u0001\u0000\u0000\u0000\u04c9\u04ca\u0001\u0000\u0000"+
		"\u0000\u04ca\u04ce\u0001\u0000\u0000\u0000\u04cb\u04cc\u0003\u00cae\u0000"+
		"\u04cc\u04cd\u0005Y\u0000\u0000\u04cd\u04cf\u0001\u0000\u0000\u0000\u04ce"+
		"\u04cb\u0001\u0000\u0000\u0000\u04ce\u04cf\u0001\u0000\u0000\u0000\u04cf"+
		"\u04d0\u0001\u0000\u0000\u0000\u04d0\u04d2\u0003\u0154\u00aa\u0000\u04d1"+
		"\u04d3\u0003\u00cae\u0000\u04d2\u04d1\u0001\u0000\u0000\u0000\u04d2\u04d3"+
		"\u0001\u0000\u0000\u0000\u04d3\u04db\u0001\u0000\u0000\u0000\u04d4\u04d6"+
		"\u0005Y\u0000\u0000\u04d5\u04d4\u0001\u0000\u0000\u0000\u04d6\u04d9\u0001"+
		"\u0000\u0000\u0000\u04d7\u04d5\u0001\u0000\u0000\u0000\u04d7\u04d8\u0001"+
		"\u0000\u0000\u0000\u04d8\u04da\u0001\u0000\u0000\u0000\u04d9\u04d7\u0001"+
		"\u0000\u0000\u0000\u04da\u04dc\u0003`0\u0000\u04db\u04d7\u0001\u0000\u0000"+
		"\u0000\u04db\u04dc\u0001\u0000\u0000\u0000\u04dc\u04df\u0001\u0000\u0000"+
		"\u0000\u04dd\u04df\u0003\b\u0004\u0000\u04de\u04c9\u0001\u0000\u0000\u0000"+
		"\u04de\u04dd\u0001\u0000\u0000\u0000\u04df\u0083\u0001\u0000\u0000\u0000"+
		"\u04e0\u04e1\u0007\u0000\u0000\u0000\u04e1\u04e3\u0003\u0086C\u0000\u04e2"+
		"\u04e4\u0003,\u0016\u0000\u04e3\u04e2\u0001\u0000\u0000\u0000\u04e3\u04e4"+
		"\u0001\u0000\u0000\u0000\u04e4\u04e5\u0001\u0000\u0000\u0000\u04e5\u04e6"+
		"\u0003\u0088D\u0000\u04e6\u04f1\u0001\u0000\u0000\u0000\u04e7\u04e8\u0005"+
		"\u0019\u0000\u0000\u04e8\u04e9\u0003\u0086C\u0000\u04e9\u04ea\u0003\u0164"+
		"\u00b2\u0000\u04ea\u04ec\u0005\u0017\u0000\u0000\u04eb\u04ed\u0003,\u0016"+
		"\u0000\u04ec\u04eb\u0001\u0000\u0000\u0000\u04ec\u04ed\u0001\u0000\u0000"+
		"\u0000\u04ed\u04ee\u0001\u0000\u0000\u0000\u04ee\u04ef\u0003\u0088D\u0000"+
		"\u04ef\u04f1\u0001\u0000\u0000\u0000\u04f0\u04e0\u0001\u0000\u0000\u0000"+
		"\u04f0\u04e7\u0001\u0000\u0000\u0000\u04f1\u0085\u0001\u0000\u0000\u0000"+
		"\u04f2\u04f3\u0003\u013e\u009f\u0000\u04f3\u0087\u0001\u0000\u0000\u0000"+
		"\u04f4\u04f6\u0005Y\u0000\u0000\u04f5\u04f4\u0001\u0000\u0000\u0000\u04f6"+
		"\u04f9\u0001\u0000\u0000\u0000\u04f7\u04f5\u0001\u0000\u0000\u0000\u04f7"+
		"\u04f8\u0001\u0000\u0000\u0000\u04f8\u04fa\u0001\u0000\u0000\u0000\u04f9"+
		"\u04f7\u0001\u0000\u0000\u0000\u04fa\u04fc\u0003\u014e\u00a7\u0000\u04fb"+
		"\u04f7\u0001\u0000\u0000\u0000\u04fb\u04fc\u0001\u0000\u0000\u0000\u04fc"+
		"\u0504\u0001\u0000\u0000\u0000\u04fd\u04ff\u0005Y\u0000\u0000\u04fe\u04fd"+
		"\u0001\u0000\u0000\u0000\u04ff\u0502\u0001\u0000\u0000\u0000\u0500\u04fe"+
		"\u0001\u0000\u0000\u0000\u0500\u0501\u0001\u0000\u0000\u0000\u0501\u0503"+
		"\u0001\u0000\u0000\u0000\u0502\u0500\u0001\u0000\u0000\u0000\u0503\u0505"+
		"\u0003\u008aE\u0000\u0504\u0500\u0001\u0000\u0000\u0000\u0504\u0505\u0001"+
		"\u0000\u0000\u0000\u0505\u0089\u0001\u0000\u0000\u0000\u0506\u050a\u0005"+
		"\u001e\u0000\u0000\u0507\u0508\u0003\u0162\u00b1\u0000\u0508\u0509\u0005"+
		"Y\u0000\u0000\u0509\u050b\u0001\u0000\u0000\u0000\u050a\u0507\u0001\u0000"+
		"\u0000\u0000\u050a\u050b\u0001\u0000\u0000\u0000\u050b\u0513\u0001\u0000"+
		"\u0000\u0000\u050c\u050e\u0005Y\u0000\u0000\u050d\u050c\u0001\u0000\u0000"+
		"\u0000\u050e\u0511\u0001\u0000\u0000\u0000\u050f\u050d\u0001\u0000\u0000"+
		"\u0000\u050f\u0510\u0001\u0000\u0000\u0000\u0510\u0512\u0001\u0000\u0000"+
		"\u0000\u0511\u050f\u0001\u0000\u0000\u0000\u0512\u0514\u0003\u008cF\u0000"+
		"\u0513\u050f\u0001\u0000\u0000\u0000\u0513\u0514\u0001\u0000\u0000\u0000"+
		"\u0514\u0518\u0001\u0000\u0000\u0000\u0515\u0517\u0005Y\u0000\u0000\u0516"+
		"\u0515\u0001\u0000\u0000\u0000\u0517\u051a\u0001\u0000\u0000\u0000\u0518"+
		"\u0516\u0001\u0000\u0000\u0000\u0518\u0519\u0001\u0000\u0000\u0000\u0519"+
		"\u051b\u0001\u0000\u0000\u0000\u051a\u0518\u0001\u0000\u0000\u0000\u051b"+
		"\u051c\u0005!\u0000\u0000\u051c\u008b\u0001\u0000\u0000\u0000\u051d\u0529"+
		"\u0003\u008eG\u0000\u051e\u0522\u0003\u0192\u00c9\u0000\u051f\u0521\u0005"+
		"Y\u0000\u0000\u0520\u051f\u0001\u0000\u0000\u0000\u0521\u0524\u0001\u0000"+
		"\u0000\u0000\u0522\u0520\u0001\u0000\u0000\u0000\u0522\u0523\u0001\u0000"+
		"\u0000\u0000\u0523\u0525\u0001\u0000\u0000\u0000\u0524\u0522\u0001\u0000"+
		"\u0000\u0000\u0525\u0526\u0003\u008eG\u0000\u0526\u0528\u0001\u0000\u0000"+
		"\u0000\u0527\u051e\u0001\u0000\u0000\u0000\u0528\u052b\u0001\u0000\u0000"+
		"\u0000\u0529\u0527\u0001\u0000\u0000\u0000\u0529\u052a\u0001\u0000\u0000"+
		"\u0000\u052a\u052d\u0001\u0000\u0000\u0000\u052b\u0529\u0001\u0000\u0000"+
		"\u0000\u052c\u052e\u0003\u0192\u00c9\u0000\u052d\u052c\u0001\u0000\u0000"+
		"\u0000\u052d\u052e\u0001\u0000\u0000\u0000\u052e\u008d\u0001\u0000\u0000"+
		"\u0000\u052f\u0530\u0003\u0160\u00b0\u0000\u0530\u0531\u0005Y\u0000\u0000"+
		"\u0531\u0533\u0001\u0000\u0000\u0000\u0532\u052f\u0001\u0000\u0000\u0000"+
		"\u0532\u0533\u0001\u0000\u0000\u0000\u0533\u0537\u0001\u0000\u0000\u0000"+
		"\u0534\u0535\u0003\u00cae\u0000\u0535\u0536\u0005Y\u0000\u0000\u0536\u0538"+
		"\u0001\u0000\u0000\u0000\u0537\u0534\u0001\u0000\u0000\u0000\u0537\u0538"+
		"\u0001\u0000\u0000\u0000\u0538\u053e\u0001\u0000\u0000\u0000\u0539\u053f"+
		"\u0003\u0084B\u0000\u053a\u053f\u0003z=\u0000\u053b\u053f\u0003X,\u0000"+
		"\u053c\u053f\u0003d2\u0000\u053d\u053f\u0003\u0090H\u0000\u053e\u0539"+
		"\u0001\u0000\u0000\u0000\u053e\u053a\u0001\u0000\u0000\u0000\u053e\u053b"+
		"\u0001\u0000\u0000\u0000\u053e\u053c\u0001\u0000\u0000\u0000\u053e\u053d"+
		"\u0001\u0000\u0000\u0000\u053f\u0542\u0001\u0000\u0000\u0000\u0540\u0542"+
		"\u0003\b\u0004\u0000\u0541\u0532\u0001\u0000\u0000\u0000\u0541\u0540\u0001"+
		"\u0000\u0000\u0000\u0542\u008f\u0001\u0000\u0000\u0000\u0543\u0544\u0003"+
		"\u0154\u00aa\u0000\u0544\u054c\u0003\u0138\u009c\u0000\u0545\u0547\u0005"+
		"Y\u0000\u0000\u0546\u0545\u0001\u0000\u0000\u0000\u0547\u054a\u0001\u0000"+
		"\u0000\u0000\u0548\u0546\u0001\u0000\u0000\u0000\u0548\u0549\u0001\u0000"+
		"\u0000\u0000\u0549\u054b\u0001\u0000\u0000\u0000\u054a\u0548\u0001\u0000"+
		"\u0000\u0000\u054b\u054d\u0003`0\u0000\u054c\u0548\u0001\u0000\u0000\u0000"+
		"\u054c\u054d\u0001\u0000\u0000\u0000\u054d\u0091\u0001\u0000\u0000\u0000"+
		"\u054e\u054f\u0005\u000f\u0000\u0000\u054f\u0551\u0003\u0094J\u0000\u0550"+
		"\u0552\u0003,\u0016\u0000\u0551\u0550\u0001\u0000\u0000\u0000\u0551\u0552"+
		"\u0001\u0000\u0000\u0000\u0552\u0553\u0001\u0000\u0000\u0000\u0553\u0554"+
		"\u0003\u0096K\u0000\u0554\u055f\u0001\u0000\u0000\u0000\u0555\u0556\u0005"+
		"\u0019\u0000\u0000\u0556\u0557\u0003\u0094J\u0000\u0557\u0558\u0003\u0164"+
		"\u00b2\u0000\u0558\u055a\u0005\u000f\u0000\u0000\u0559\u055b\u0003,\u0016"+
		"\u0000\u055a\u0559\u0001\u0000\u0000\u0000\u055a\u055b\u0001\u0000\u0000"+
		"\u0000\u055b\u055c\u0001\u0000\u0000\u0000\u055c\u055d\u0003\u0096K\u0000"+
		"\u055d\u055f\u0001\u0000\u0000\u0000\u055e\u054e\u0001\u0000\u0000\u0000"+
		"\u055e\u0555\u0001\u0000\u0000\u0000\u055f\u0093\u0001\u0000\u0000\u0000"+
		"\u0560\u0561\u0003\u013e\u009f\u0000\u0561\u0095\u0001\u0000\u0000\u0000"+
		"\u0562\u0564\u0005Y\u0000\u0000\u0563\u0562\u0001\u0000\u0000\u0000\u0564"+
		"\u0567\u0001\u0000\u0000\u0000\u0565\u0563\u0001\u0000\u0000\u0000\u0565"+
		"\u0566\u0001\u0000\u0000\u0000\u0566\u0568\u0001\u0000\u0000\u0000\u0567"+
		"\u0565\u0001\u0000\u0000\u0000\u0568\u056a\u0003\u014e\u00a7\u0000\u0569"+
		"\u0565\u0001\u0000\u0000\u0000\u0569\u056a\u0001\u0000\u0000\u0000\u056a"+
		"\u056e\u0001\u0000\u0000\u0000\u056b\u056d\u0005Y\u0000\u0000\u056c\u056b"+
		"\u0001\u0000\u0000\u0000\u056d\u0570\u0001\u0000\u0000\u0000\u056e\u056c"+
		"\u0001\u0000\u0000\u0000\u056e\u056f\u0001\u0000\u0000\u0000\u056f\u0571"+
		"\u0001\u0000\u0000\u0000\u0570\u056e\u0001\u0000\u0000\u0000\u0571\u0572"+
		"\u0003\u0098L\u0000\u0572\u0097\u0001\u0000\u0000\u0000\u0573\u0575\u0005"+
		"\u001e\u0000\u0000\u0574\u0576\u0003\u0162\u00b1\u0000\u0575\u0574\u0001"+
		"\u0000\u0000\u0000\u0575\u0576\u0001\u0000\u0000\u0000\u0576\u057e\u0001"+
		"\u0000\u0000\u0000\u0577\u0579\u0005Y\u0000\u0000\u0578\u0577\u0001\u0000"+
		"\u0000\u0000\u0579\u057c\u0001\u0000\u0000\u0000\u057a\u0578\u0001\u0000"+
		"\u0000\u0000\u057a\u057b\u0001\u0000\u0000\u0000\u057b\u057d\u0001\u0000"+
		"\u0000\u0000\u057c\u057a\u0001\u0000\u0000\u0000\u057d\u057f\u0003\u009a"+
		"M\u0000\u057e\u057a\u0001\u0000\u0000\u0000\u057e\u057f\u0001\u0000\u0000"+
		"\u0000\u057f\u0583\u0001\u0000\u0000\u0000\u0580\u0582\u0005Y\u0000\u0000"+
		"\u0581\u0580\u0001\u0000\u0000\u0000\u0582\u0585\u0001\u0000\u0000\u0000"+
		"\u0583\u0581\u0001\u0000\u0000\u0000\u0583\u0584\u0001\u0000\u0000\u0000"+
		"\u0584\u0586\u0001\u0000\u0000\u0000\u0585\u0583\u0001\u0000\u0000\u0000"+
		"\u0586\u0587\u0005!\u0000\u0000\u0587\u0099\u0001\u0000\u0000\u0000\u0588"+
		"\u0594\u0003\u009cN\u0000\u0589\u058d\u0003\u0192\u00c9\u0000\u058a\u058c"+
		"\u0005Y\u0000\u0000\u058b\u058a\u0001\u0000\u0000\u0000\u058c\u058f\u0001"+
		"\u0000\u0000\u0000\u058d\u058b\u0001\u0000\u0000\u0000\u058d\u058e\u0001"+
		"\u0000\u0000\u0000\u058e\u0590\u0001\u0000\u0000\u0000\u058f\u058d\u0001"+
		"\u0000\u0000\u0000\u0590\u0591\u0003\u009cN\u0000\u0591\u0593\u0001\u0000"+
		"\u0000\u0000\u0592\u0589\u0001\u0000\u0000\u0000\u0593\u0596\u0001\u0000"+
		"\u0000\u0000\u0594\u0592\u0001\u0000\u0000\u0000\u0594\u0595\u0001\u0000"+
		"\u0000\u0000\u0595\u0598\u0001\u0000\u0000\u0000\u0596\u0594\u0001\u0000"+
		"\u0000\u0000\u0597\u0599\u0003\u0192\u00c9\u0000\u0598\u0597\u0001\u0000"+
		"\u0000\u0000\u0598\u0599\u0001\u0000\u0000\u0000\u0599\u009b\u0001\u0000"+
		"\u0000\u0000\u059a\u059b\u0003\u0160\u00b0\u0000\u059b\u059c\u0005Y\u0000"+
		"\u0000\u059c\u059e\u0001\u0000\u0000\u0000\u059d\u059a\u0001\u0000\u0000"+
		"\u0000\u059d\u059e\u0001\u0000\u0000\u0000\u059e\u05a2\u0001\u0000\u0000"+
		"\u0000\u059f\u05a0\u0003\u00cae\u0000\u05a0\u05a1\u0005Y\u0000\u0000\u05a1"+
		"\u05a3\u0001\u0000\u0000\u0000\u05a2\u059f\u0001\u0000\u0000\u0000\u05a2"+
		"\u05a3\u0001\u0000\u0000\u0000\u05a3\u05a6\u0001\u0000\u0000\u0000\u05a4"+
		"\u05a7\u0003d2\u0000\u05a5\u05a7\u0003\u009eO\u0000\u05a6\u05a4\u0001"+
		"\u0000\u0000\u0000\u05a6\u05a5\u0001\u0000\u0000\u0000\u05a7\u05aa\u0001"+
		"\u0000\u0000\u0000\u05a8\u05aa\u0003\b\u0004\u0000\u05a9\u059d\u0001\u0000"+
		"\u0000\u0000\u05a9\u05a8\u0001\u0000\u0000\u0000\u05aa\u009d\u0001\u0000"+
		"\u0000\u0000\u05ab\u05ad\u0003l6\u0000\u05ac\u05ae\u0003,\u0016\u0000"+
		"\u05ad\u05ac\u0001\u0000\u0000\u0000\u05ad\u05ae\u0001\u0000\u0000\u0000"+
		"\u05ae\u05b2\u0001\u0000\u0000\u0000\u05af\u05b1\u0005Y\u0000\u0000\u05b0"+
		"\u05af\u0001\u0000\u0000\u0000\u05b1\u05b4\u0001\u0000\u0000\u0000\u05b2"+
		"\u05b0\u0001\u0000\u0000\u0000\u05b2\u05b3\u0001\u0000\u0000\u0000\u05b3"+
		"\u05b5\u0001\u0000\u0000\u0000\u05b4\u05b2\u0001\u0000\u0000\u0000\u05b5"+
		"\u05b6\u0003n7\u0000\u05b6\u009f\u0001\u0000\u0000\u0000\u05b7\u05b8\u0005"+
		"\u0003\u0000\u0000\u05b8\u05ba\u0003\u00c2a\u0000\u05b9\u05bb\u0003,\u0016"+
		"\u0000\u05ba\u05b9\u0001\u0000\u0000\u0000\u05ba\u05bb\u0001\u0000\u0000"+
		"\u0000\u05bb\u05ca\u0001\u0000\u0000\u0000\u05bc\u05cb\u0003\u008aE\u0000"+
		"\u05bd\u05c5\u0003\u0138\u009c\u0000\u05be\u05c0\u0005Y\u0000\u0000\u05bf"+
		"\u05be\u0001\u0000\u0000\u0000\u05c0\u05c3\u0001\u0000\u0000\u0000\u05c1"+
		"\u05bf\u0001\u0000\u0000\u0000\u05c1\u05c2\u0001\u0000\u0000\u0000\u05c2"+
		"\u05c4\u0001\u0000\u0000\u0000\u05c3\u05c1\u0001\u0000\u0000\u0000\u05c4"+
		"\u05c6\u0003`0\u0000\u05c5\u05c1\u0001\u0000\u0000\u0000\u05c5\u05c6\u0001"+
		"\u0000\u0000\u0000\u05c6\u05c8\u0001\u0000\u0000\u0000\u05c7\u05c9\u0003"+
		"\u0162\u00b1\u0000\u05c8\u05c7\u0001\u0000\u0000\u0000\u05c8\u05c9\u0001"+
		"\u0000\u0000\u0000\u05c9\u05cb\u0001\u0000\u0000\u0000\u05ca\u05bc\u0001"+
		"\u0000\u0000\u0000\u05ca\u05bd\u0001\u0000\u0000\u0000\u05cb\u00a1\u0001"+
		"\u0000\u0000\u0000\u05cc\u05cd\u0005\u0003\u0000\u0000\u05cd\u05cf\u0003"+
		"\u00c2a\u0000\u05ce\u05d0\u0003,\u0016\u0000\u05cf\u05ce\u0001\u0000\u0000"+
		"\u0000\u05cf\u05d0\u0001\u0000\u0000\u0000\u05d0\u05d4\u0001\u0000\u0000"+
		"\u0000\u05d1\u05d3\u0005Y\u0000\u0000\u05d2\u05d1\u0001\u0000\u0000\u0000"+
		"\u05d3\u05d6\u0001\u0000\u0000\u0000\u05d4\u05d2\u0001\u0000\u0000\u0000"+
		"\u05d4\u05d5\u0001\u0000\u0000\u0000\u05d5\u05d7\u0001\u0000\u0000\u0000"+
		"\u05d6\u05d4\u0001\u0000\u0000\u0000\u05d7\u05d8\u0003\u00a4R\u0000\u05d8"+
		"\u00a3\u0001\u0000\u0000\u0000\u05d9\u05dd\u0003\u0164\u00b2\u0000\u05da"+
		"\u05dc\u0005Y\u0000\u0000\u05db\u05da\u0001\u0000\u0000\u0000\u05dc\u05df"+
		"\u0001\u0000\u0000\u0000\u05dd\u05db\u0001\u0000\u0000\u0000\u05dd\u05de"+
		"\u0001\u0000\u0000\u0000\u05de\u05e3\u0001\u0000\u0000\u0000\u05df\u05dd"+
		"\u0001\u0000\u0000\u0000\u05e0\u05e1\u0003>\u001f\u0000\u05e1\u05e2\u0005"+
		"\u001d\u0000\u0000\u05e2\u05e4\u0001\u0000\u0000\u0000\u05e3\u05e0\u0001"+
		"\u0000\u0000\u0000\u05e3\u05e4\u0001\u0000\u0000\u0000\u05e4\u05e5\u0001"+
		"\u0000\u0000\u0000\u05e5\u05e7\u0003\u00c2a\u0000\u05e6\u05e8\u00032\u0019"+
		"\u0000\u05e7\u05e6\u0001\u0000\u0000\u0000\u05e7\u05e8\u0001\u0000\u0000"+
		"\u0000\u05e8\u05ea\u0001\u0000\u0000\u0000\u05e9\u05eb\u0003\u0162\u00b1"+
		"\u0000\u05ea\u05e9\u0001\u0000\u0000\u0000\u05ea\u05eb\u0001\u0000\u0000"+
		"\u0000\u05eb\u00a5\u0001\u0000\u0000\u0000\u05ec\u05ed\u0006S\uffff\uffff"+
		"\u0000\u05ed\u05ef\u0003\u00a8T\u0000\u05ee\u05f0\u0003\u0138\u009c\u0000"+
		"\u05ef\u05ee\u0001\u0000\u0000\u0000\u05ef\u05f0\u0001\u0000\u0000\u0000"+
		"\u05f0\u060a\u0001\u0000\u0000\u0000\u05f1\u05f3\u0003\u00aaU\u0000\u05f2"+
		"\u05f4\u0003\u0138\u009c\u0000\u05f3\u05f2\u0001\u0000\u0000\u0000\u05f3"+
		"\u05f4\u0001\u0000\u0000\u0000\u05f4\u060a\u0001\u0000\u0000\u0000\u05f5"+
		"\u05f7\u0003\u00acV\u0000\u05f6\u05f8\u0003\u0138\u009c\u0000\u05f7\u05f6"+
		"\u0001\u0000\u0000\u0000\u05f7\u05f8\u0001\u0000\u0000\u0000\u05f8\u060a"+
		"\u0001\u0000\u0000\u0000\u05f9\u05fb\u0003\u00b2Y\u0000\u05fa\u05fc\u0003"+
		"\u0138\u009c\u0000\u05fb\u05fa\u0001\u0000\u0000\u0000\u05fb\u05fc\u0001"+
		"\u0000\u0000\u0000\u05fc\u060a\u0001\u0000\u0000\u0000\u05fd\u05ff\u0003"+
		"\u0132\u0099\u0000\u05fe\u0600\u0003\u00cae\u0000\u05ff\u05fe\u0001\u0000"+
		"\u0000\u0000\u05ff\u0600\u0001\u0000\u0000\u0000\u0600\u060a\u0001\u0000"+
		"\u0000\u0000\u0601\u060a\u0003\u00b8\\\u0000\u0602\u060a\u0003\u00ba]"+
		"\u0000\u0603\u0605\u0005)\u0000\u0000\u0604\u0603\u0001\u0000\u0000\u0000"+
		"\u0604\u0605\u0001\u0000\u0000\u0000\u0605\u0606\u0001\u0000\u0000\u0000"+
		"\u0606\u0607\u0005\u0010\u0000\u0000\u0607\u060a\u0003\u0132\u0099\u0000"+
		"\u0608\u060a\u0003\u00bc^\u0000\u0609\u05ec\u0001\u0000\u0000\u0000\u0609"+
		"\u05f1\u0001\u0000\u0000\u0000\u0609\u05f5\u0001\u0000\u0000\u0000\u0609"+
		"\u05f9\u0001\u0000\u0000\u0000\u0609\u05fd\u0001\u0000\u0000\u0000\u0609"+
		"\u0601\u0001\u0000\u0000\u0000\u0609\u0602\u0001\u0000\u0000\u0000\u0609"+
		"\u0604\u0001\u0000\u0000\u0000\u0609\u0608\u0001\u0000\u0000\u0000\u060a"+
		"\u0610\u0001\u0000\u0000\u0000\u060b\u060c\n\u0002\u0000\u0000\u060c\u060d"+
		"\u0005\u0002\u0000\u0000\u060d\u060f\u0003\u0132\u0099\u0000\u060e\u060b"+
		"\u0001\u0000\u0000\u0000\u060f\u0612\u0001\u0000\u0000\u0000\u0610\u060e"+
		"\u0001\u0000\u0000\u0000\u0610\u0611\u0001\u0000\u0000\u0000\u0611\u00a7"+
		"\u0001\u0000\u0000\u0000\u0612\u0610\u0001\u0000\u0000\u0000\u0613\u0614"+
		"\u00058\u0000\u0000\u0614\u00a9\u0001\u0000\u0000\u0000\u0615\u0616\u0003"+
		"\u0154\u00aa\u0000\u0616\u00ab\u0001\u0000\u0000\u0000\u0617\u0619\u0005"+
		"\u001f\u0000\u0000\u0618\u061a\u0003\u00aeW\u0000\u0619\u0618\u0001\u0000"+
		"\u0000\u0000\u0619\u061a\u0001\u0000\u0000\u0000\u061a\u061b\u0001\u0000"+
		"\u0000\u0000\u061b\u061c\u0005\"\u0000\u0000\u061c\u00ad\u0001\u0000\u0000"+
		"\u0000\u061d\u0622\u0003\u00b0X\u0000\u061e\u061f\u0005$\u0000\u0000\u061f"+
		"\u0621\u0003\u00b0X\u0000\u0620\u061e\u0001\u0000\u0000\u0000\u0621\u0624"+
		"\u0001\u0000\u0000\u0000\u0622\u0620\u0001\u0000\u0000\u0000\u0622\u0623"+
		"\u0001\u0000\u0000\u0000\u0623\u00af\u0001\u0000\u0000\u0000\u0624\u0622"+
		"\u0001\u0000\u0000\u0000\u0625\u0628\u0003\u00a6S\u0000\u0626\u0628\u0005"+
		"A\u0000\u0000\u0627\u0625\u0001\u0000\u0000\u0000\u0627\u0626\u0001\u0000"+
		"\u0000\u0000\u0628\u00b1\u0001\u0000\u0000\u0000\u0629\u062b\u0005 \u0000"+
		"\u0000\u062a\u062c\u0003\u00b4Z\u0000\u062b\u062a\u0001\u0000\u0000\u0000"+
		"\u062b\u062c\u0001\u0000\u0000\u0000\u062c\u062d\u0001\u0000\u0000\u0000"+
		"\u062d\u062e\u0005#\u0000\u0000\u062e\u00b3\u0001\u0000\u0000\u0000\u062f"+
		"\u0634\u0003\u00b6[\u0000\u0630\u0631\u0005$\u0000\u0000\u0631\u0633\u0003"+
		"\u00b6[\u0000\u0632\u0630\u0001\u0000\u0000\u0000\u0633\u0636\u0001\u0000"+
		"\u0000\u0000\u0634\u0632\u0001\u0000\u0000\u0000\u0634\u0635\u0001\u0000"+
		"\u0000\u0000\u0635\u00b5\u0001\u0000\u0000\u0000\u0636\u0634\u0001\u0000"+
		"\u0000\u0000\u0637\u063a\u0003\u00a6S\u0000\u0638\u063a\u0005A\u0000\u0000"+
		"\u0639\u0637\u0001\u0000\u0000\u0000\u0639\u0638\u0001\u0000\u0000\u0000"+
		"\u063a\u00b7\u0001\u0000\u0000\u0000\u063b\u063d\u0003\u013a\u009d\u0000"+
		"\u063c\u063b\u0001\u0000\u0000\u0000\u063c\u063d\u0001\u0000\u0000\u0000"+
		"\u063d\u063e\u0001\u0000\u0000\u0000\u063e\u063f\u0005\u001d\u0000\u0000"+
		"\u063f\u0641\u0003\u0154\u00aa\u0000\u0640\u0642\u0003\u00acV\u0000\u0641"+
		"\u0640\u0001\u0000\u0000\u0000\u0641\u0642\u0001\u0000\u0000\u0000\u0642"+
		"\u00b9\u0001\u0000\u0000\u0000\u0643\u0644\u0003\u00aaU\u0000\u0644\u0645"+
		"\u0005*\u0000\u0000\u0645\u00bb\u0001\u0000\u0000\u0000\u0646\u0647\u0003"+
		"\u00ccf\u0000\u0647\u00bd\u0001\u0000\u0000\u0000\u0648\u0649\u0005+\u0000"+
		"\u0000\u0649\u0653\u0005O\u0000\u0000\u064a\u064b\u0005+\u0000\u0000\u064b"+
		"\u064d\u0003\u00c0`\u0000\u064c\u064e\u00032\u0019\u0000\u064d\u064c\u0001"+
		"\u0000\u0000\u0000\u064d\u064e\u0001\u0000\u0000\u0000\u064e\u0650\u0001"+
		"\u0000\u0000\u0000\u064f\u0651\u0003\u00c4b\u0000\u0650\u064f\u0001\u0000"+
		"\u0000\u0000\u0650\u0651\u0001\u0000\u0000\u0000\u0651\u0653\u0001\u0000"+
		"\u0000\u0000\u0652\u0648\u0001\u0000\u0000\u0000\u0652\u064a\u0001\u0000"+
		"\u0000\u0000\u0653\u00bf\u0001\u0000\u0000\u0000\u0654\u0655\u0003>\u001f"+
		"\u0000\u0655\u0656\u0005\u001d\u0000\u0000\u0656\u0658\u0001\u0000\u0000"+
		"\u0000\u0657\u0654\u0001\u0000\u0000\u0000\u0657\u0658\u0001\u0000\u0000"+
		"\u0000\u0658\u0659\u0001\u0000\u0000\u0000\u0659\u065a\u0003\u00c2a\u0000"+
		"\u065a\u00c1\u0001\u0000\u0000\u0000\u065b\u065c\u0003\u0156\u00ab\u0000"+
		"\u065c\u00c3\u0001\u0000\u0000\u0000\u065d\u0665\u0005\u001f\u0000\u0000"+
		"\u065e\u0660\u0005Y\u0000\u0000\u065f\u065e\u0001\u0000\u0000\u0000\u0660"+
		"\u0663\u0001\u0000\u0000\u0000\u0661\u065f\u0001\u0000\u0000\u0000\u0661"+
		"\u0662\u0001\u0000\u0000\u0000\u0662\u0664\u0001\u0000\u0000\u0000\u0663"+
		"\u0661\u0001\u0000\u0000\u0000\u0664\u0666\u0003\u00c8d\u0000\u0665\u0661"+
		"\u0001\u0000\u0000\u0000\u0665\u0666\u0001\u0000\u0000\u0000\u0666\u066a"+
		"\u0001\u0000\u0000\u0000\u0667\u0669\u0005Y\u0000\u0000\u0668\u0667\u0001"+
		"\u0000\u0000\u0000\u0669\u066c\u0001\u0000\u0000\u0000\u066a\u0668\u0001"+
		"\u0000\u0000\u0000\u066a\u066b\u0001\u0000\u0000\u0000\u066b\u066d\u0001"+
		"\u0000\u0000\u0000\u066c\u066a\u0001\u0000\u0000\u0000\u066d\u066e\u0005"+
		"\"\u0000\u0000\u066e\u00c5\u0001\u0000\u0000\u0000\u066f\u0670\u0003\u0156"+
		"\u00ab\u0000\u0670\u0671\u0005%\u0000\u0000\u0671\u0673\u0001\u0000\u0000"+
		"\u0000\u0672\u066f\u0001\u0000\u0000\u0000\u0672\u0673\u0001\u0000\u0000"+
		"\u0000\u0673\u0674\u0001\u0000\u0000\u0000\u0674\u0675\u0003\u00ccf\u0000"+
		"\u0675\u00c7\u0001\u0000\u0000\u0000\u0676\u0682\u0003\u00c6c\u0000\u0677"+
		"\u067b\u0003\u0190\u00c8\u0000\u0678\u067a\u0005Y\u0000\u0000\u0679\u0678"+
		"\u0001\u0000\u0000\u0000\u067a\u067d\u0001\u0000\u0000\u0000\u067b\u0679"+
		"\u0001\u0000\u0000\u0000\u067b\u067c\u0001\u0000\u0000\u0000\u067c\u067e"+
		"\u0001\u0000\u0000\u0000\u067d\u067b\u0001\u0000\u0000\u0000\u067e\u067f"+
		"\u0003\u00c6c\u0000\u067f\u0681\u0001\u0000\u0000\u0000\u0680\u0677\u0001"+
		"\u0000\u0000\u0000\u0681\u0684\u0001\u0000\u0000\u0000\u0682\u0680\u0001"+
		"\u0000\u0000\u0000\u0682\u0683\u0001\u0000\u0000\u0000\u0683\u0686\u0001"+
		"\u0000\u0000\u0000\u0684\u0682\u0001\u0000\u0000\u0000\u0685\u0687\u0003"+
		"\u0190\u00c8\u0000\u0686\u0685\u0001\u0000\u0000\u0000\u0686\u0687\u0001"+
		"\u0000\u0000\u0000\u0687\u00c9\u0001\u0000\u0000\u0000\u0688\u068f\u0003"+
		"\u00be_\u0000\u0689\u068b\u0005Y\u0000\u0000\u068a\u0689\u0001\u0000\u0000"+
		"\u0000\u068a\u068b\u0001\u0000\u0000\u0000\u068b\u068c\u0001\u0000\u0000"+
		"\u0000\u068c\u068e\u0003\u00be_\u0000\u068d\u068a\u0001\u0000\u0000\u0000"+
		"\u068e\u0691\u0001\u0000\u0000\u0000\u068f\u068d\u0001\u0000\u0000\u0000"+
		"\u068f\u0690\u0001\u0000\u0000\u0000\u0690\u00cb\u0001\u0000\u0000\u0000"+
		"\u0691\u068f\u0001\u0000\u0000\u0000\u0692\u0694\u0003\u00ceg\u0000\u0693"+
		"\u0695\u0003\u00d6k\u0000\u0694\u0693\u0001\u0000\u0000\u0000\u0694\u0695"+
		"\u0001\u0000\u0000\u0000\u0695\u00cd\u0001\u0000\u0000\u0000\u0696\u0697"+
		"\u0003\u0172\u00b9\u0000\u0697\u0698\u0003\u0118\u008c\u0000\u0698\u069b"+
		"\u0001\u0000\u0000\u0000\u0699\u069b\u0003\u0118\u008c\u0000\u069a\u0696"+
		"\u0001\u0000\u0000\u0000\u069a\u0699\u0001\u0000\u0000\u0000\u069b\u00cf"+
		"\u0001\u0000\u0000\u0000\u069c\u069d\u0003\u0170\u00b8\u0000\u069d\u069e"+
		"\u0003\u00ceg\u0000\u069e\u06ac\u0001\u0000\u0000\u0000\u069f\u06a0\u0003"+
		"\u00dam\u0000\u06a0\u06a1\u0003\u00ceg\u0000\u06a1\u06ac\u0001\u0000\u0000"+
		"\u0000\u06a2\u06a3\u0003\u00d8l\u0000\u06a3\u06a4\u0003\u00ceg\u0000\u06a4"+
		"\u06ac\u0001\u0000\u0000\u0000\u06a5\u06a6\u0003\u00dcn\u0000\u06a6\u06a7"+
		"\u0003\u00ceg\u0000\u06a7\u06ac\u0001\u0000\u0000\u0000\u06a8\u06a9\u0003"+
		"\u00d4j\u0000\u06a9\u06aa\u0003\u00ceg\u0000\u06aa\u06ac\u0001\u0000\u0000"+
		"\u0000\u06ab\u069c\u0001\u0000\u0000\u0000\u06ab\u069f\u0001\u0000\u0000"+
		"\u0000\u06ab\u06a2\u0001\u0000\u0000\u0000\u06ab\u06a5\u0001\u0000\u0000"+
		"\u0000\u06ab\u06a8\u0001\u0000\u0000\u0000\u06ac\u00d1\u0001\u0000\u0000"+
		"\u0000\u06ad\u06ae\u0003\u0156\u00ab\u0000\u06ae\u00d3\u0001\u0000\u0000"+
		"\u0000\u06af\u06b0\u0005J\u0000\u0000\u06b0\u00d5\u0001\u0000\u0000\u0000"+
		"\u06b1\u06b3\u0003\u00d0h\u0000\u06b2\u06b1\u0001\u0000\u0000\u0000\u06b3"+
		"\u06b4\u0001\u0000\u0000\u0000\u06b4\u06b2\u0001\u0000\u0000\u0000\u06b4"+
		"\u06b5\u0001\u0000\u0000\u0000\u06b5\u00d7\u0001\u0000\u0000\u0000\u06b6"+
		"\u06b7\u0005)\u0000\u0000\u06b7\u06ba\u0005\u000e\u0000\u0000\u06b8\u06ba"+
		"\u0005\u000e\u0000\u0000\u06b9\u06b6\u0001\u0000\u0000\u0000\u06b9\u06b8"+
		"\u0001\u0000\u0000\u0000\u06ba\u00d9\u0001\u0000\u0000\u0000\u06bb\u06bc"+
		"\u0005*\u0000\u0000\u06bc\u06bd\u0003\u00ccf\u0000\u06bd\u06be\u0005%"+
		"\u0000\u0000\u06be\u00db\u0001\u0000\u0000\u0000\u06bf\u06c0\u0005\f\u0000"+
		"\u0000\u06c0\u06c1\u0003\u00ccf\u0000\u06c1\u06c2\u0005\u0007\u0000\u0000"+
		"\u06c2\u00dd\u0001\u0000\u0000\u0000\u06c3\u06c4\u0005)\u0000\u0000\u06c4"+
		"\u06c7\u0005\u0010\u0000\u0000\u06c5\u06c7\u0005\u0010\u0000\u0000\u06c6"+
		"\u06c3\u0001\u0000\u0000\u0000\u06c6\u06c5\u0001\u0000\u0000\u0000\u06c7"+
		"\u06ca\u0001\u0000\u0000\u0000\u06c8\u06cb\u0003\u0132\u0099\u0000\u06c9"+
		"\u06cb\u0003\u00a6S\u0000\u06ca\u06c8\u0001\u0000\u0000\u0000\u06ca\u06c9"+
		"\u0001\u0000\u0000\u0000\u06cb\u06cf\u0001\u0000\u0000\u0000\u06cc\u06cd"+
		"\u0005\u0002\u0000\u0000\u06cd\u06cf\u0003\u0132\u0099\u0000\u06ce\u06c6"+
		"\u0001\u0000\u0000\u0000\u06ce\u06cc\u0001\u0000\u0000\u0000\u06cf\u00df"+
		"\u0001\u0000\u0000\u0000\u06d0\u06e3\u0003\u00e2q\u0000\u06d1\u06d3\u0003"+
		"\u0154\u00aa\u0000\u06d2\u06d4\u00032\u0019\u0000\u06d3\u06d2\u0001\u0000"+
		"\u0000\u0000\u06d3\u06d4\u0001\u0000\u0000\u0000\u06d4\u06e3\u0001\u0000"+
		"\u0000\u0000\u06d5\u06d6\u0003\u013a\u009d\u0000\u06d6\u06d7\u0005\u001d"+
		"\u0000\u0000\u06d7\u06d9\u0003\u0154\u00aa\u0000\u06d8\u06da\u00032\u0019"+
		"\u0000\u06d9\u06d8\u0001\u0000\u0000\u0000\u06d9\u06da\u0001\u0000\u0000"+
		"\u0000\u06da\u06e3\u0001\u0000\u0000\u0000\u06db\u06e3\u0003\u0108\u0084"+
		"\u0000\u06dc\u06e3\u0003\u0112\u0089\u0000\u06dd\u06e3\u0003\u0110\u0088"+
		"\u0000\u06de\u06e3\u0003\u010e\u0087\u0000\u06df\u06e3\u0003\u0116\u008b"+
		"\u0000\u06e0\u06e3\u0003\u0100\u0080\u0000\u06e1\u06e3\u0005A\u0000\u0000"+
		"\u06e2\u06d0\u0001\u0000\u0000\u0000\u06e2\u06d1\u0001\u0000\u0000\u0000"+
		"\u06e2\u06d5\u0001\u0000\u0000\u0000\u06e2\u06db\u0001\u0000\u0000\u0000"+
		"\u06e2\u06dc\u0001\u0000\u0000\u0000\u06e2\u06dd\u0001\u0000\u0000\u0000"+
		"\u06e2\u06de\u0001\u0000\u0000\u0000\u06e2\u06df\u0001\u0000\u0000\u0000"+
		"\u06e2\u06e0\u0001\u0000\u0000\u0000\u06e2\u06e1\u0001\u0000\u0000\u0000"+
		"\u06e3\u00e1\u0001\u0000\u0000\u0000\u06e4\u06ec\u0003\u00e4r\u0000\u06e5"+
		"\u06ec\u0003\u00e6s\u0000\u06e6\u06ec\u0003\u00fe\u007f\u0000\u06e7\u06ec"+
		"\u0003\u0182\u00c1\u0000\u06e8\u06ec\u0003\u00ecv\u0000\u06e9\u06ec\u0003"+
		"\u00f2y\u0000\u06ea\u06ec\u0003\u00f8|\u0000\u06eb\u06e4\u0001\u0000\u0000"+
		"\u0000\u06eb\u06e5\u0001\u0000\u0000\u0000\u06eb\u06e6\u0001\u0000\u0000"+
		"\u0000\u06eb\u06e7\u0001\u0000\u0000\u0000\u06eb\u06e8\u0001\u0000\u0000"+
		"\u0000\u06eb\u06e9\u0001\u0000\u0000\u0000\u06eb\u06ea\u0001\u0000\u0000"+
		"\u0000\u06ec\u00e3\u0001\u0000\u0000\u0000\u06ed\u06ee\u0003\u0188\u00c4"+
		"\u0000\u06ee\u06ef\u0003\u00e8t\u0000\u06ef\u00e5\u0001\u0000\u0000\u0000"+
		"\u06f0\u06f1\u0003\u00eau\u0000\u06f1\u06f2\u0004s\u0001\u0000\u06f2\u06f4"+
		"\u0003\u018c\u00c6\u0000\u06f3\u06f5\u0003\u00e8t\u0000\u06f4\u06f3\u0001"+
		"\u0000\u0000\u0000\u06f4\u06f5\u0001\u0000\u0000\u0000\u06f5\u06fa\u0001"+
		"\u0000\u0000\u0000\u06f6\u06f7\u0003\u018c\u00c6\u0000\u06f7\u06f8\u0003"+
		"\u00e8t\u0000\u06f8\u06fa\u0001\u0000\u0000\u0000\u06f9\u06f0\u0001\u0000"+
		"\u0000\u0000\u06f9\u06f6\u0001\u0000\u0000\u0000\u06fa\u00e7\u0001\u0000"+
		"\u0000\u0000\u06fb\u06fc\u0004t\u0002\u0000\u06fc\u06fd\u0007\u0001\u0000"+
		"\u0000\u06fd\u00e9\u0001\u0000\u0000\u0000\u06fe\u06ff\u0005J\u0000\u0000"+
		"\u06ff\u00eb\u0001\u0000\u0000\u0000\u0700\u0708\u0005 \u0000\u0000\u0701"+
		"\u0703\u0005Y\u0000\u0000\u0702\u0701\u0001\u0000\u0000\u0000\u0703\u0706"+
		"\u0001\u0000\u0000\u0000\u0704\u0702\u0001\u0000\u0000\u0000\u0704\u0705"+
		"\u0001\u0000\u0000\u0000\u0705\u0707\u0001\u0000\u0000\u0000\u0706\u0704"+
		"\u0001\u0000\u0000\u0000\u0707\u0709\u0003\u00eew\u0000\u0708\u0704\u0001"+
		"\u0000\u0000\u0000\u0708\u0709\u0001\u0000\u0000\u0000\u0709\u070d\u0001"+
		"\u0000\u0000\u0000\u070a\u070c\u0005Y\u0000\u0000\u070b\u070a\u0001\u0000"+
		"\u0000\u0000\u070c\u070f\u0001\u0000\u0000\u0000\u070d\u070b\u0001\u0000"+
		"\u0000\u0000\u070d\u070e\u0001\u0000\u0000\u0000\u070e\u0710\u0001\u0000"+
		"\u0000\u0000\u070f\u070d\u0001\u0000\u0000\u0000\u0710\u0711\u0005#\u0000"+
		"\u0000\u0711\u00ed\u0001\u0000\u0000\u0000\u0712\u071e\u0003\u00f0x\u0000"+
		"\u0713\u0717\u0003\u0190\u00c8\u0000\u0714\u0716\u0005Y\u0000\u0000\u0715"+
		"\u0714\u0001\u0000\u0000\u0000\u0716\u0719\u0001\u0000\u0000\u0000\u0717"+
		"\u0715\u0001\u0000\u0000\u0000\u0717\u0718\u0001\u0000\u0000\u0000\u0718"+
		"\u071a\u0001\u0000\u0000\u0000\u0719\u0717\u0001\u0000\u0000\u0000\u071a"+
		"\u071b\u0003\u00f0x\u0000\u071b\u071d\u0001\u0000\u0000\u0000\u071c\u0713"+
		"\u0001\u0000\u0000\u0000\u071d\u0720\u0001\u0000\u0000\u0000\u071e\u071c"+
		"\u0001\u0000\u0000\u0000\u071e\u071f\u0001\u0000\u0000\u0000\u071f\u0722"+
		"\u0001\u0000\u0000\u0000\u0720\u071e\u0001\u0000\u0000\u0000\u0721\u0723"+
		"\u0003\u0190\u00c8\u0000\u0722\u0721\u0001\u0000\u0000\u0000\u0722\u0723"+
		"\u0001\u0000\u0000\u0000\u0723\u00ef\u0001\u0000\u0000\u0000\u0724\u0727"+
		"\u0003\u00ccf\u0000\u0725\u0727\u0005A\u0000\u0000\u0726\u0724\u0001\u0000"+
		"\u0000\u0000\u0726\u0725\u0001\u0000\u0000\u0000\u0727\u00f1\u0001\u0000"+
		"\u0000\u0000\u0728\u0730\u0005\u001e\u0000\u0000\u0729\u072b\u0005Y\u0000"+
		"\u0000\u072a\u0729\u0001\u0000\u0000\u0000\u072b\u072e\u0001\u0000\u0000"+
		"\u0000\u072c\u072a\u0001\u0000\u0000\u0000\u072c\u072d\u0001\u0000\u0000"+
		"\u0000\u072d\u072f\u0001\u0000\u0000\u0000\u072e\u072c\u0001\u0000\u0000"+
		"\u0000\u072f\u0731\u0003\u00f4z\u0000\u0730\u072c\u0001\u0000\u0000\u0000"+
		"\u0730\u0731\u0001\u0000\u0000\u0000\u0731\u0735\u0001\u0000\u0000\u0000"+
		"\u0732\u0734\u0005Y\u0000\u0000\u0733\u0732\u0001\u0000\u0000\u0000\u0734"+
		"\u0737\u0001\u0000\u0000\u0000\u0735\u0733\u0001\u0000\u0000\u0000\u0735"+
		"\u0736\u0001\u0000\u0000\u0000\u0736\u0738\u0001\u0000\u0000\u0000\u0737"+
		"\u0735\u0001\u0000\u0000\u0000\u0738\u0739\u0005!\u0000\u0000\u0739\u00f3"+
		"\u0001\u0000\u0000\u0000\u073a\u0746\u0003\u00f6{\u0000\u073b\u073f\u0003"+
		"\u0190\u00c8\u0000\u073c\u073e\u0005Y\u0000\u0000\u073d\u073c\u0001\u0000"+
		"\u0000\u0000\u073e\u0741\u0001\u0000\u0000\u0000\u073f\u073d\u0001\u0000"+
		"\u0000\u0000\u073f\u0740\u0001\u0000\u0000\u0000\u0740\u0742\u0001\u0000"+
		"\u0000\u0000\u0741\u073f\u0001\u0000\u0000\u0000\u0742\u0743\u0003\u00f6"+
		"{\u0000\u0743\u0745\u0001\u0000\u0000\u0000\u0744\u073b\u0001\u0000\u0000"+
		"\u0000\u0745\u0748\u0001\u0000\u0000\u0000\u0746\u0744\u0001\u0000\u0000"+
		"\u0000\u0746\u0747\u0001\u0000\u0000\u0000\u0747\u074a\u0001\u0000\u0000"+
		"\u0000\u0748\u0746\u0001\u0000\u0000\u0000\u0749\u074b\u0003\u0190\u00c8"+
		"\u0000\u074a\u0749\u0001\u0000\u0000\u0000\u074a\u074b\u0001\u0000\u0000"+
		"\u0000\u074b\u00f5\u0001\u0000\u0000\u0000\u074c\u074f\u0003\u018c\u00c6"+
		"\u0000\u074d\u074f\u0003\u018a\u00c5\u0000\u074e\u074c\u0001\u0000\u0000"+
		"\u0000\u074e\u074d\u0001\u0000\u0000\u0000\u074f\u0750\u0001\u0000\u0000"+
		"\u0000\u0750\u0751\u0005%\u0000\u0000\u0751\u0752\u0003\u00ccf\u0000\u0752"+
		"\u00f7\u0001\u0000\u0000\u0000\u0753\u075b\u0005\u001e\u0000\u0000\u0754"+
		"\u0756\u0005Y\u0000\u0000\u0755\u0754\u0001\u0000\u0000\u0000\u0756\u0759"+
		"\u0001\u0000\u0000\u0000\u0757\u0755\u0001\u0000\u0000\u0000\u0757\u0758"+
		"\u0001\u0000\u0000\u0000\u0758\u075a\u0001\u0000\u0000\u0000\u0759\u0757"+
		"\u0001\u0000\u0000\u0000\u075a\u075c\u0003\u00fa}\u0000\u075b\u0757\u0001"+
		"\u0000\u0000\u0000\u075b\u075c\u0001\u0000\u0000\u0000\u075c\u0760\u0001"+
		"\u0000\u0000\u0000\u075d\u075f\u0005Y\u0000\u0000\u075e\u075d\u0001\u0000"+
		"\u0000\u0000\u075f\u0762\u0001\u0000\u0000\u0000\u0760\u075e\u0001\u0000"+
		"\u0000\u0000\u0760\u0761\u0001\u0000\u0000\u0000\u0761\u0763\u0001\u0000"+
		"\u0000\u0000\u0762\u0760\u0001\u0000\u0000\u0000\u0763\u0764\u0005!\u0000"+
		"\u0000\u0764\u00f9\u0001\u0000\u0000\u0000\u0765\u0771\u0003\u00fc~\u0000"+
		"\u0766\u076a\u0003\u0190\u00c8\u0000\u0767\u0769\u0005Y\u0000\u0000\u0768"+
		"\u0767\u0001\u0000\u0000\u0000\u0769\u076c\u0001\u0000\u0000\u0000\u076a"+
		"\u0768\u0001\u0000\u0000\u0000\u076a\u076b\u0001\u0000\u0000\u0000\u076b"+
		"\u076d\u0001\u0000\u0000\u0000\u076c\u076a\u0001\u0000\u0000\u0000\u076d"+
		"\u076e\u0003\u00fc~\u0000\u076e\u0770\u0001\u0000\u0000\u0000\u076f\u0766"+
		"\u0001\u0000\u0000\u0000\u0770\u0773\u0001\u0000\u0000\u0000\u0771\u076f"+
		"\u0001\u0000\u0000\u0000\u0771\u0772\u0001\u0000\u0000\u0000\u0772\u0775"+
		"\u0001\u0000\u0000\u0000\u0773\u0771\u0001\u0000\u0000\u0000\u0774\u0776"+
		"\u0003\u0190\u00c8\u0000\u0775\u0774\u0001\u0000\u0000\u0000\u0775\u0776"+
		"\u0001\u0000\u0000\u0000\u0776\u00fb\u0001\u0000\u0000\u0000\u0777\u077a"+
		"\u0003\u0158\u00ac\u0000\u0778\u0779\u0005%\u0000\u0000\u0779\u077b\u0003"+
		"\u00ccf\u0000\u077a\u0778\u0001\u0000\u0000\u0000\u077a\u077b\u0001\u0000"+
		"\u0000\u0000\u077b\u00fd\u0001\u0000\u0000\u0000\u077c\u077d\u0003\u013a"+
		"\u009d\u0000\u077d\u077e\u0003\u00f8|\u0000\u077e\u00ff\u0001\u0000\u0000"+
		"\u0000\u077f\u0780\u0003\u013a\u009d\u0000\u0780\u0781\u0003\u0120\u0090"+
		"\u0000\u0781\u0101\u0001\u0000\u0000\u0000\u0782\u0786\u0005\u0011\u0000"+
		"\u0000\u0783\u0785\u0005Y\u0000\u0000\u0784\u0783\u0001\u0000\u0000\u0000"+
		"\u0785\u0788\u0001\u0000\u0000\u0000\u0786\u0784\u0001\u0000\u0000\u0000"+
		"\u0786\u0787\u0001\u0000\u0000\u0000\u0787\u0789\u0001\u0000\u0000\u0000"+
		"\u0788\u0786\u0001\u0000\u0000\u0000\u0789\u0791\u0005\u001e\u0000\u0000"+
		"\u078a\u078c\u0005Y\u0000\u0000\u078b\u078a\u0001\u0000\u0000\u0000\u078c"+
		"\u078f\u0001\u0000\u0000\u0000\u078d\u078b\u0001\u0000\u0000\u0000\u078d"+
		"\u078e\u0001\u0000\u0000\u0000\u078e\u0790\u0001\u0000\u0000\u0000\u078f"+
		"\u078d\u0001\u0000\u0000\u0000\u0790\u0792\u0003\u0104\u0082\u0000\u0791"+
		"\u078d\u0001\u0000\u0000\u0000\u0791\u0792\u0001\u0000\u0000\u0000\u0792"+
		"\u0796\u0001\u0000\u0000\u0000\u0793\u0795\u0005Y\u0000\u0000\u0794\u0793"+
		"\u0001\u0000\u0000\u0000\u0795\u0798\u0001\u0000\u0000\u0000\u0796\u0794"+
		"\u0001\u0000\u0000\u0000\u0796\u0797\u0001\u0000\u0000\u0000\u0797\u0799"+
		"\u0001\u0000\u0000\u0000\u0798\u0796\u0001\u0000\u0000\u0000\u0799\u079a"+
		"\u0005!\u0000\u0000\u079a\u0103\u0001\u0000\u0000\u0000\u079b\u07a7\u0003"+
		"\u0106\u0083\u0000\u079c\u07a0\u0003\u018e\u00c7\u0000\u079d\u079f\u0005"+
		"Y\u0000\u0000\u079e\u079d\u0001\u0000\u0000\u0000\u079f\u07a2\u0001\u0000"+
		"\u0000\u0000\u07a0\u079e\u0001\u0000\u0000\u0000\u07a0\u07a1\u0001\u0000"+
		"\u0000\u0000\u07a1\u07a3\u0001\u0000\u0000\u0000\u07a2\u07a0\u0001\u0000"+
		"\u0000\u0000\u07a3\u07a4\u0003\u0106\u0083\u0000\u07a4\u07a6\u0001\u0000"+
		"\u0000\u0000\u07a5\u079c\u0001\u0000\u0000\u0000\u07a6\u07a9\u0001\u0000"+
		"\u0000\u0000\u07a7\u07a5\u0001\u0000\u0000\u0000\u07a7\u07a8\u0001\u0000"+
		"\u0000\u0000\u07a8\u07ab\u0001\u0000\u0000\u0000\u07a9\u07a7\u0001\u0000"+
		"\u0000\u0000\u07aa\u07ac\u0003\u018e\u00c7\u0000\u07ab\u07aa\u0001\u0000"+
		"\u0000\u0000\u07ab\u07ac\u0001\u0000\u0000\u0000\u07ac\u0105\u0001\u0000"+
		"\u0000\u0000\u07ad\u07b1\u0003\u00a6S\u0000\u07ae\u07b0\u0005Y\u0000\u0000"+
		"\u07af\u07ae\u0001\u0000\u0000\u0000\u07b0\u07b3\u0001\u0000\u0000\u0000"+
		"\u07b1\u07af\u0001\u0000\u0000\u0000\u07b1\u07b2\u0001\u0000\u0000\u0000"+
		"\u07b2\u07b4\u0001\u0000\u0000\u0000\u07b3\u07b1\u0001\u0000\u0000\u0000"+
		"\u07b4\u07b8\u0005<\u0000\u0000\u07b5\u07b7\u0005Y\u0000\u0000\u07b6\u07b5"+
		"\u0001\u0000\u0000\u0000\u07b7\u07ba\u0001\u0000\u0000\u0000\u07b8\u07b6"+
		"\u0001\u0000\u0000\u0000\u07b8\u07b9\u0001\u0000\u0000\u0000\u07b9\u07bb"+
		"\u0001\u0000\u0000\u0000\u07ba\u07b8\u0001\u0000\u0000\u0000\u07bb\u07bc"+
		"\u0003\u00ccf\u0000\u07bc\u0107\u0001\u0000\u0000\u0000\u07bd\u07c1\u0005"+
		"\u001e\u0000\u0000\u07be\u07c0\u0005Y\u0000\u0000\u07bf\u07be\u0001\u0000"+
		"\u0000\u0000\u07c0\u07c3\u0001\u0000\u0000\u0000\u07c1\u07bf\u0001\u0000"+
		"\u0000\u0000\u07c1\u07c2\u0001\u0000\u0000\u0000\u07c2\u07c4\u0001\u0000"+
		"\u0000\u0000\u07c3\u07c1\u0001\u0000\u0000\u0000\u07c4\u07c8\u0003\n\u0005"+
		"\u0000\u07c5\u07c7\u0005Y\u0000\u0000\u07c6\u07c5\u0001\u0000\u0000\u0000"+
		"\u07c7\u07ca\u0001\u0000\u0000\u0000\u07c8\u07c6\u0001\u0000\u0000\u0000"+
		"\u07c8\u07c9\u0001\u0000\u0000\u0000\u07c9\u07cb\u0001\u0000\u0000\u0000"+
		"\u07ca\u07c8\u0001\u0000\u0000\u0000\u07cb\u07cc\u0005!\u0000\u0000\u07cc"+
		"\u07e9\u0001\u0000\u0000\u0000\u07cd\u07ce\u0005\u001e\u0000\u0000\u07ce"+
		"\u07d2\u0003\u010a\u0085\u0000\u07cf\u07d1\u0005Y\u0000\u0000\u07d0\u07cf"+
		"\u0001\u0000\u0000\u0000\u07d1\u07d4\u0001\u0000\u0000\u0000\u07d2\u07d0"+
		"\u0001\u0000\u0000\u0000\u07d2\u07d3\u0001\u0000\u0000\u0000\u07d3\u07d5"+
		"\u0001\u0000\u0000\u0000\u07d4\u07d2\u0001\u0000\u0000\u0000\u07d5\u07dd"+
		"\u0005=\u0000\u0000\u07d6\u07d8\u0005Y\u0000\u0000\u07d7\u07d6\u0001\u0000"+
		"\u0000\u0000\u07d8\u07db\u0001\u0000\u0000\u0000\u07d9\u07d7\u0001\u0000"+
		"\u0000\u0000\u07d9\u07da\u0001\u0000\u0000\u0000\u07da\u07dc\u0001\u0000"+
		"\u0000\u0000\u07db\u07d9\u0001\u0000\u0000\u0000\u07dc\u07de\u0003\u0132"+
		"\u0099\u0000\u07dd\u07d9\u0001\u0000\u0000\u0000\u07dd\u07de\u0001\u0000"+
		"\u0000\u0000\u07de\u07e2\u0001\u0000\u0000\u0000\u07df\u07e1\u0005Y\u0000"+
		"\u0000\u07e0\u07df\u0001\u0000\u0000\u0000\u07e1\u07e4\u0001\u0000\u0000"+
		"\u0000\u07e2\u07e0\u0001\u0000\u0000\u0000\u07e2\u07e3\u0001\u0000\u0000"+
		"\u0000\u07e3\u07e5\u0001\u0000\u0000\u0000\u07e4\u07e2\u0001\u0000\u0000"+
		"\u0000\u07e5\u07e6\u0003\n\u0005\u0000\u07e6\u07e7\u0005!\u0000\u0000"+
		"\u07e7\u07e9\u0001\u0000\u0000\u0000\u07e8\u07bd\u0001\u0000\u0000\u0000"+
		"\u07e8\u07cd\u0001\u0000\u0000\u0000\u07e9\u0109\u0001\u0000\u0000\u0000"+
		"\u07ea\u07f6\u0003\u010c\u0086\u0000\u07eb\u07ef\u0003\u0190\u00c8\u0000"+
		"\u07ec\u07ee\u0005Y\u0000\u0000\u07ed\u07ec\u0001\u0000\u0000\u0000\u07ee"+
		"\u07f1\u0001\u0000\u0000\u0000\u07ef\u07ed\u0001\u0000\u0000\u0000\u07ef"+
		"\u07f0\u0001\u0000\u0000\u0000\u07f0\u07f2\u0001\u0000\u0000\u0000\u07f1"+
		"\u07ef\u0001\u0000\u0000\u0000\u07f2\u07f3\u0003\u010c\u0086\u0000\u07f3"+
		"\u07f5\u0001\u0000\u0000\u0000\u07f4\u07eb\u0001\u0000\u0000\u0000\u07f5"+
		"\u07f8\u0001\u0000\u0000\u0000\u07f6\u07f4\u0001\u0000\u0000\u0000\u07f6"+
		"\u07f7\u0001\u0000\u0000\u0000\u07f7\u07fa\u0001\u0000\u0000\u0000\u07f8"+
		"\u07f6\u0001\u0000\u0000\u0000\u07f9\u07fb\u0003\u0190\u00c8\u0000\u07fa"+
		"\u07f9\u0001\u0000\u0000\u0000\u07fa\u07fb\u0001\u0000\u0000\u0000\u07fb"+
		"\u010b\u0001\u0000\u0000\u0000\u07fc\u07ff\u0003x<\u0000\u07fd\u07ff\u0003"+
		"\u0156\u00ab\u0000\u07fe\u07fc\u0001\u0000\u0000\u0000\u07fe\u07fd\u0001"+
		"\u0000\u0000\u0000\u07ff\u010d\u0001\u0000\u0000\u0000\u0800\u0801\u0005"+
		"\u001d\u0000\u0000\u0801\u0802\u0003\u0156\u00ab\u0000\u0802\u010f\u0001"+
		"\u0000\u0000\u0000\u0803\u0807\u0005\u001f\u0000\u0000\u0804\u0806\u0005"+
		"Y\u0000\u0000\u0805\u0804\u0001\u0000\u0000\u0000\u0806\u0809\u0001\u0000"+
		"\u0000\u0000\u0807\u0805\u0001\u0000\u0000\u0000\u0807\u0808\u0001\u0000"+
		"\u0000\u0000\u0808\u080a\u0001\u0000\u0000\u0000\u0809\u0807\u0001\u0000"+
		"\u0000\u0000\u080a\u080e\u0003\u00ccf\u0000\u080b\u080d\u0005Y\u0000\u0000"+
		"\u080c\u080b\u0001\u0000\u0000\u0000\u080d\u0810\u0001\u0000\u0000\u0000"+
		"\u080e\u080c\u0001\u0000\u0000\u0000\u080e\u080f\u0001\u0000\u0000\u0000"+
		"\u080f\u0811\u0001\u0000\u0000\u0000\u0810\u080e\u0001\u0000\u0000\u0000"+
		"\u0811\u0812\u0005\"\u0000\u0000\u0812\u0111\u0001\u0000\u0000\u0000\u0813"+
		"\u0814\u0005\u001f\u0000\u0000\u0814\u0820\u0005\"\u0000\u0000\u0815\u0816"+
		"\u0005\u001f\u0000\u0000\u0816\u0819\u0003\u0114\u008a\u0000\u0817\u0818"+
		"\u0005$\u0000\u0000\u0818\u081a\u0003\u0114\u008a\u0000\u0819\u0817\u0001"+
		"\u0000\u0000\u0000\u081a\u081b\u0001\u0000\u0000\u0000\u081b\u0819\u0001"+
		"\u0000\u0000\u0000\u081b\u081c\u0001\u0000\u0000\u0000\u081c\u081d\u0001"+
		"\u0000\u0000\u0000\u081d\u081e\u0005\"\u0000\u0000\u081e\u0820\u0001\u0000"+
		"\u0000\u0000\u081f\u0813\u0001\u0000\u0000\u0000\u081f\u0815\u0001\u0000"+
		"\u0000\u0000\u0820\u0113\u0001\u0000\u0000\u0000\u0821\u0828\u0003\u00cc"+
		"f\u0000\u0822\u0823\u0003\u0156\u00ab\u0000\u0823\u0824\u0005%\u0000\u0000"+
		"\u0824\u0825\u0003\u00ccf\u0000\u0825\u0828\u0001\u0000\u0000\u0000\u0826"+
		"\u0828\u0005A\u0000\u0000\u0827\u0821\u0001\u0000\u0000\u0000\u0827\u0822"+
		"\u0001\u0000\u0000\u0000\u0827\u0826\u0001\u0000\u0000\u0000\u0828\u0115"+
		"\u0001\u0000\u0000\u0000\u0829\u082a\u00058\u0000\u0000\u082a\u0117\u0001"+
		"\u0000\u0000\u0000\u082b\u082f\u0003\u00e0p\u0000\u082c\u082e\u0003\u011a"+
		"\u008d\u0000\u082d\u082c\u0001\u0000\u0000\u0000\u082e\u0831\u0001\u0000"+
		"\u0000\u0000\u082f\u082d\u0001\u0000\u0000\u0000\u082f\u0830\u0001\u0000"+
		"\u0000\u0000\u0830\u0833\u0001\u0000\u0000\u0000\u0831\u082f\u0001\u0000"+
		"\u0000\u0000\u0832\u0834\u0003\u0174\u00ba\u0000\u0833\u0832\u0001\u0000"+
		"\u0000\u0000\u0833\u0834\u0001\u0000\u0000\u0000\u0834\u0119\u0001\u0000"+
		"\u0000\u0000\u0835\u083b\u0003\u0120\u0090\u0000\u0836\u083b\u0003\u011c"+
		"\u008e\u0000\u0837\u083b\u0003\u011e\u008f\u0000\u0838\u083b\u0003\u0102"+
		"\u0081\u0000\u0839\u083b\u0003\u00deo\u0000\u083a\u0835\u0001\u0000\u0000"+
		"\u0000\u083a\u0836\u0001\u0000\u0000\u0000\u083a\u0837\u0001\u0000\u0000"+
		"\u0000\u083a\u0838\u0001\u0000\u0000\u0000\u083a\u0839\u0001\u0000\u0000"+
		"\u0000\u083b\u011b\u0001\u0000\u0000\u0000\u083c\u0846\u0005\u001d\u0000"+
		"\u0000\u083d\u0847\u0005O\u0000\u0000\u083e\u0844\u0003\u015a\u00ad\u0000"+
		"\u083f\u0845\u00032\u0019\u0000\u0840\u0841\u0005\u001f\u0000\u0000\u0841"+
		"\u0842\u0003\u012e\u0097\u0000\u0842\u0843\u0005\"\u0000\u0000\u0843\u0845"+
		"\u0001\u0000\u0000\u0000\u0844\u083f\u0001\u0000\u0000\u0000\u0844\u0840"+
		"\u0001\u0000\u0000\u0000\u0844\u0845\u0001\u0000\u0000\u0000\u0845\u0847"+
		"\u0001\u0000\u0000\u0000\u0846\u083d\u0001\u0000\u0000\u0000\u0846\u083e"+
		"\u0001\u0000\u0000\u0000\u0847\u011d\u0001\u0000\u0000\u0000\u0848\u0849"+
		"\u0005 \u0000\u0000\u0849\u084a\u0003\u0124\u0092\u0000\u084a\u084b\u0005"+
		"#\u0000\u0000\u084b\u011f\u0001\u0000\u0000\u0000\u084c\u084e\u0003\u0122"+
		"\u0091\u0000\u084d\u084c\u0001\u0000\u0000\u0000\u084d\u084e\u0001\u0000"+
		"\u0000\u0000\u084e\u084f\u0001\u0000\u0000\u0000\u084f\u0852\u0003\u0128"+
		"\u0094\u0000\u0850\u0852\u0003\u0122\u0091\u0000\u0851\u084d\u0001\u0000"+
		"\u0000\u0000\u0851\u0850\u0001\u0000\u0000\u0000\u0852\u0121\u0001\u0000"+
		"\u0000\u0000\u0853\u0857\u0005\u001f\u0000\u0000\u0854\u0856\u0005Y\u0000"+
		"\u0000\u0855\u0854\u0001\u0000\u0000\u0000\u0856\u0859\u0001\u0000\u0000"+
		"\u0000\u0857\u0855\u0001\u0000\u0000\u0000\u0857\u0858\u0001\u0000\u0000"+
		"\u0000\u0858\u085a\u0001\u0000\u0000\u0000\u0859\u0857\u0001\u0000\u0000"+
		"\u0000\u085a\u086c\u0005\"\u0000\u0000\u085b\u085f\u0005\u001f\u0000\u0000"+
		"\u085c\u085e\u0005Y\u0000\u0000\u085d\u085c\u0001\u0000\u0000\u0000\u085e"+
		"\u0861\u0001\u0000\u0000\u0000\u085f\u085d\u0001\u0000\u0000\u0000\u085f"+
		"\u0860\u0001\u0000\u0000\u0000\u0860\u0862\u0001\u0000\u0000\u0000\u0861"+
		"\u085f\u0001\u0000\u0000\u0000\u0862\u0866\u0003\u0124\u0092\u0000\u0863"+
		"\u0865\u0005Y\u0000\u0000\u0864\u0863\u0001\u0000\u0000\u0000\u0865\u0868"+
		"\u0001\u0000\u0000\u0000\u0866\u0864\u0001\u0000\u0000\u0000\u0866\u0867"+
		"\u0001\u0000\u0000\u0000\u0867\u0869\u0001\u0000\u0000\u0000\u0868\u0866"+
		"\u0001\u0000\u0000\u0000\u0869\u086a\u0005\"\u0000\u0000\u086a\u086c\u0001"+
		"\u0000\u0000\u0000\u086b\u0853\u0001\u0000\u0000\u0000\u086b\u085b\u0001"+
		"\u0000\u0000\u0000\u086c\u0123\u0001\u0000\u0000\u0000\u086d\u0872\u0003"+
		"\u0126\u0093\u0000\u086e\u086f\u0005$\u0000\u0000\u086f\u0871\u0003\u0126"+
		"\u0093\u0000\u0870\u086e\u0001\u0000\u0000\u0000\u0871\u0874\u0001\u0000"+
		"\u0000\u0000\u0872\u0870\u0001\u0000\u0000\u0000\u0872\u0873\u0001\u0000"+
		"\u0000\u0000\u0873\u0125\u0001\u0000\u0000\u0000\u0874\u0872\u0001\u0000"+
		"\u0000\u0000\u0875\u087b\u0003\u00ccf\u0000\u0876\u0877\u0003\u0156\u00ab"+
		"\u0000\u0877\u0878\u0005%\u0000\u0000\u0878\u0879\u0003\u00ccf\u0000\u0879"+
		"\u087b\u0001\u0000\u0000\u0000\u087a\u0875\u0001\u0000\u0000\u0000\u087a"+
		"\u0876\u0001\u0000\u0000\u0000\u087b\u0127\u0001\u0000\u0000\u0000\u087c"+
		"\u087e\u0003\u0108\u0084\u0000\u087d\u087f\u0003\u012a\u0095\u0000\u087e"+
		"\u087d\u0001\u0000\u0000\u0000\u087e\u087f\u0001\u0000\u0000\u0000\u087f"+
		"\u0129\u0001\u0000\u0000\u0000\u0880\u0882\u0003\u012c\u0096\u0000\u0881"+
		"\u0880\u0001\u0000\u0000\u0000\u0882\u0883\u0001\u0000\u0000\u0000\u0883"+
		"\u0881\u0001\u0000\u0000\u0000\u0883\u0884\u0001\u0000\u0000\u0000\u0884"+
		"\u012b\u0001\u0000\u0000\u0000\u0885\u0886\u0003\u015a\u00ad\u0000\u0886"+
		"\u0887\u0005%\u0000\u0000\u0887\u0888\u0003\u0108\u0084\u0000\u0888\u012d"+
		"\u0001\u0000\u0000\u0000\u0889\u088d\u0003\u0130\u0098\u0000\u088a\u088c"+
		"\u0003\u0130\u0098\u0000\u088b\u088a\u0001\u0000\u0000\u0000\u088c\u088f"+
		"\u0001\u0000\u0000\u0000\u088d\u088b\u0001\u0000\u0000\u0000\u088d\u088e"+
		"\u0001\u0000\u0000\u0000\u088e\u012f\u0001\u0000\u0000\u0000\u088f\u088d"+
		"\u0001\u0000\u0000\u0000\u0890\u0891\u0003\u0156\u00ab\u0000\u0891\u0892"+
		"\u0005%\u0000\u0000\u0892\u0131\u0001\u0000\u0000\u0000\u0893\u0894\u0006"+
		"\u0099\uffff\uffff\u0000\u0894\u0897\u0003\u0134\u009a\u0000\u0895\u0897"+
		"\u0003\u0146\u00a3\u0000\u0896\u0893\u0001\u0000\u0000\u0000\u0896\u0895"+
		"\u0001\u0000\u0000\u0000\u0897\u08a6\u0001\u0000\u0000\u0000\u0898\u0899"+
		"\n\u0006\u0000\u0000\u0899\u08a5\u0005)\u0000\u0000\u089a\u089b\n\u0005"+
		"\u0000\u0000\u089b\u08a5\u0005*\u0000\u0000\u089c\u089d\n\u0004\u0000"+
		"\u0000\u089d\u08a5\u00052\u0000\u0000\u089e\u089f\n\u0003\u0000\u0000"+
		"\u089f\u08a5\u00051\u0000\u0000\u08a0\u08a1\n\u0002\u0000\u0000\u08a1"+
		"\u08a5\u0005-\u0000\u0000\u08a2\u08a3\n\u0001\u0000\u0000\u08a3\u08a5"+
		"\u0005A\u0000\u0000\u08a4\u0898\u0001\u0000\u0000\u0000\u08a4\u089a\u0001"+
		"\u0000\u0000\u0000\u08a4\u089c\u0001\u0000\u0000\u0000\u08a4\u089e\u0001"+
		"\u0000\u0000\u0000\u08a4\u08a0\u0001\u0000\u0000\u0000\u08a4\u08a2\u0001"+
		"\u0000\u0000\u0000\u08a5\u08a8\u0001\u0000\u0000\u0000\u08a6\u08a4\u0001"+
		"\u0000\u0000\u0000\u08a6\u08a7\u0001\u0000\u0000\u0000\u08a7\u0133\u0001"+
		"\u0000\u0000\u0000\u08a8\u08a6\u0001\u0000\u0000\u0000\u08a9\u08aa\u0006"+
		"\u009a\uffff\uffff\u0000\u08aa\u08ab\u0003\u0136\u009b\u0000\u08ab\u08ec"+
		"\u0001\u0000\u0000\u0000\u08ac\u08ae\n\u0003\u0000\u0000\u08ad\u08af\u0003"+
		"\u00cae\u0000\u08ae\u08ad\u0001\u0000\u0000\u0000\u08ae\u08af\u0001\u0000"+
		"\u0000\u0000\u08af\u08b3\u0001\u0000\u0000\u0000\u08b0\u08b1\u0003\u0162"+
		"\u00b1\u0000\u08b1\u08b2\u0005Y\u0000\u0000\u08b2\u08b4\u0001\u0000\u0000"+
		"\u0000\u08b3\u08b0\u0001\u0000\u0000\u0000\u08b3\u08b4\u0001\u0000\u0000"+
		"\u0000\u08b4\u08b8\u0001\u0000\u0000\u0000\u08b5\u08b7\u0005Y\u0000\u0000"+
		"\u08b6\u08b5\u0001\u0000\u0000\u0000\u08b7\u08ba\u0001\u0000\u0000\u0000"+
		"\u08b8\u08b6\u0001\u0000\u0000\u0000\u08b8\u08b9\u0001\u0000\u0000\u0000"+
		"\u08b9\u08bb\u0001\u0000\u0000\u0000\u08ba\u08b8\u0001\u0000\u0000\u0000"+
		"\u08bb\u08bf\u0005/\u0000\u0000\u08bc\u08be\u0005Y\u0000\u0000\u08bd\u08bc"+
		"\u0001\u0000\u0000\u0000\u08be\u08c1\u0001\u0000\u0000\u0000\u08bf\u08bd"+
		"\u0001\u0000\u0000\u0000\u08bf\u08c0\u0001\u0000\u0000\u0000\u08c0\u08c2"+
		"\u0001\u0000\u0000\u0000\u08c1\u08bf\u0001\u0000\u0000\u0000\u08c2\u08c4"+
		"\u0003\u0134\u009a\u0000\u08c3\u08c5\u0003\u00cae\u0000\u08c4\u08c3\u0001"+
		"\u0000\u0000\u0000\u08c4\u08c5\u0001\u0000\u0000\u0000\u08c5\u08c9\u0001"+
		"\u0000\u0000\u0000\u08c6\u08c7\u0003\u0162\u00b1\u0000\u08c7\u08c8\u0005"+
		"Y\u0000\u0000\u08c8\u08ca\u0001\u0000\u0000\u0000\u08c9\u08c6\u0001\u0000"+
		"\u0000\u0000\u08c9\u08ca\u0001\u0000\u0000\u0000\u08ca\u08eb\u0001\u0000"+
		"\u0000\u0000\u08cb\u08cd\n\u0002\u0000\u0000\u08cc\u08ce\u0003\u00cae"+
		"\u0000\u08cd\u08cc\u0001\u0000\u0000\u0000\u08cd\u08ce\u0001\u0000\u0000"+
		"\u0000\u08ce\u08d2\u0001\u0000\u0000\u0000\u08cf\u08d0\u0003\u0162\u00b1"+
		"\u0000\u08d0\u08d1\u0005Y\u0000\u0000\u08d1\u08d3\u0001\u0000\u0000\u0000"+
		"\u08d2\u08cf\u0001\u0000\u0000\u0000\u08d2\u08d3\u0001\u0000\u0000\u0000"+
		"\u08d3\u08d7\u0001\u0000\u0000\u0000\u08d4\u08d6\u0005Y\u0000\u0000\u08d5"+
		"\u08d4\u0001\u0000\u0000\u0000\u08d6\u08d9\u0001\u0000\u0000\u0000\u08d7"+
		"\u08d5\u0001\u0000\u0000\u0000\u08d7\u08d8\u0001\u0000\u0000\u0000\u08d8"+
		"\u08da\u0001\u0000\u0000\u0000\u08d9\u08d7\u0001\u0000\u0000\u0000\u08da"+
		"\u08de\u0005,\u0000\u0000\u08db\u08dd\u0005Y\u0000\u0000\u08dc\u08db\u0001"+
		"\u0000\u0000\u0000\u08dd\u08e0\u0001\u0000\u0000\u0000\u08de\u08dc\u0001"+
		"\u0000\u0000\u0000\u08de\u08df\u0001\u0000\u0000\u0000\u08df\u08e1\u0001"+
		"\u0000\u0000\u0000\u08e0\u08de\u0001\u0000\u0000\u0000\u08e1\u08e3\u0003"+
		"\u0134\u009a\u0000\u08e2\u08e4\u0003\u00cae\u0000\u08e3\u08e2\u0001\u0000"+
		"\u0000\u0000\u08e3\u08e4\u0001\u0000\u0000\u0000\u08e4\u08e8\u0001\u0000"+
		"\u0000\u0000\u08e5\u08e6\u0003\u0162\u00b1\u0000\u08e6\u08e7\u0005Y\u0000"+
		"\u0000\u08e7\u08e9\u0001\u0000\u0000\u0000\u08e8\u08e5\u0001\u0000\u0000"+
		"\u0000\u08e8\u08e9\u0001\u0000\u0000\u0000\u08e9\u08eb\u0001\u0000\u0000"+
		"\u0000\u08ea\u08ac\u0001\u0000\u0000\u0000\u08ea\u08cb\u0001\u0000\u0000"+
		"\u0000\u08eb\u08ee\u0001\u0000\u0000\u0000\u08ec\u08ea\u0001\u0000\u0000"+
		"\u0000\u08ec\u08ed\u0001\u0000\u0000\u0000\u08ed\u0135\u0001\u0000\u0000"+
		"\u0000\u08ee\u08ec\u0001\u0000\u0000\u0000\u08ef\u08f4\u0003\u0148\u00a4"+
		"\u0000\u08f0\u08f4\u0003\u014a\u00a5\u0000\u08f1\u08f4\u0003\u0140\u00a0"+
		"\u0000\u08f2\u08f4\u0003\u013a\u009d\u0000\u08f3\u08ef\u0001\u0000\u0000"+
		"\u0000\u08f3\u08f0\u0001\u0000\u0000\u0000\u08f3\u08f1\u0001\u0000\u0000"+
		"\u0000\u08f3\u08f2\u0001\u0000\u0000\u0000\u08f4\u0137\u0001\u0000\u0000"+
		"\u0000\u08f5\u08f7\u0005%\u0000\u0000\u08f6\u08f5\u0001\u0000\u0000\u0000"+
		"\u08f6\u08f7\u0001\u0000\u0000\u0000\u08f7\u08f8\u0001\u0000\u0000\u0000"+
		"\u08f8\u08fa\u0003\u0132\u0099\u0000\u08f9\u08fb\u0003\u00cae\u0000\u08fa"+
		"\u08f9\u0001\u0000\u0000\u0000\u08fa\u08fb\u0001\u0000\u0000\u0000\u08fb"+
		"\u0139\u0001\u0000\u0000\u0000\u08fc\u08fd\u0003>\u001f\u0000\u08fd\u08fe"+
		"\u0005\u001d\u0000\u0000\u08fe\u0900\u0001\u0000\u0000\u0000\u08ff\u08fc"+
		"\u0001\u0000\u0000\u0000\u08ff\u0900\u0001\u0000\u0000\u0000\u0900\u0901"+
		"\u0001\u0000\u0000\u0000\u0901\u0906\u0003\u013c\u009e\u0000\u0902\u0903"+
		"\u0005\u001d\u0000\u0000\u0903\u0905\u0003\u013c\u009e\u0000\u0904\u0902"+
		"\u0001\u0000\u0000\u0000\u0905\u0908\u0001\u0000\u0000\u0000\u0906\u0904"+
		"\u0001\u0000\u0000\u0000\u0906\u0907\u0001\u0000\u0000\u0000\u0907\u013b"+
		"\u0001\u0000\u0000\u0000\u0908\u0906\u0001\u0000\u0000\u0000\u0909\u090b"+
		"\u0003\u013e\u009f\u0000\u090a\u090c\u00032\u0019\u0000\u090b\u090a\u0001"+
		"\u0000\u0000\u0000\u090b\u090c\u0001\u0000\u0000\u0000\u090c\u013d\u0001"+
		"\u0000\u0000\u0000\u090d\u090e\u0005I\u0000\u0000\u090e\u013f\u0001\u0000"+
		"\u0000\u0000\u090f\u0917\u0005\u001f\u0000\u0000\u0910\u0912\u0005Y\u0000"+
		"\u0000\u0911\u0910\u0001\u0000\u0000\u0000\u0912\u0915\u0001\u0000\u0000"+
		"\u0000\u0913\u0911\u0001\u0000\u0000\u0000\u0913\u0914\u0001\u0000\u0000"+
		"\u0000\u0914\u0916\u0001\u0000\u0000\u0000\u0915\u0913\u0001\u0000\u0000"+
		"\u0000\u0916\u0918\u0003\u0142\u00a1\u0000\u0917\u0913\u0001\u0000\u0000"+
		"\u0000\u0917\u0918\u0001\u0000\u0000\u0000\u0918\u091c\u0001\u0000\u0000"+
		"\u0000\u0919\u091b\u0005Y\u0000\u0000\u091a\u0919\u0001\u0000\u0000\u0000"+
		"\u091b\u091e\u0001\u0000\u0000\u0000\u091c\u091a\u0001\u0000\u0000\u0000"+
		"\u091c\u091d\u0001\u0000\u0000\u0000\u091d\u091f\u0001\u0000\u0000\u0000"+
		"\u091e\u091c\u0001\u0000\u0000\u0000\u091f\u0920\u0005\"\u0000\u0000\u0920"+
		"\u0141\u0001\u0000\u0000\u0000\u0921\u092d\u0003\u0144\u00a2\u0000\u0922"+
		"\u0926\u0003\u0194\u00ca\u0000\u0923\u0925\u0005Y\u0000\u0000\u0924\u0923"+
		"\u0001\u0000\u0000\u0000\u0925\u0928\u0001\u0000\u0000\u0000\u0926\u0924"+
		"\u0001\u0000\u0000\u0000\u0926\u0927\u0001\u0000\u0000\u0000\u0927\u0929"+
		"\u0001\u0000\u0000\u0000\u0928\u0926\u0001\u0000\u0000\u0000\u0929\u092a"+
		"\u0003\u0144\u00a2\u0000\u092a\u092c\u0001\u0000\u0000\u0000\u092b\u0922"+
		"\u0001\u0000\u0000\u0000\u092c\u092f\u0001\u0000\u0000\u0000\u092d\u092b"+
		"\u0001\u0000\u0000\u0000\u092d\u092e\u0001\u0000\u0000\u0000\u092e\u0931"+
		"\u0001\u0000\u0000\u0000\u092f\u092d\u0001\u0000\u0000\u0000\u0930\u0932"+
		"\u0003\u0194\u00ca\u0000\u0931\u0930\u0001\u0000\u0000\u0000\u0931\u0932"+
		"\u0001\u0000\u0000\u0000\u0932\u0143\u0001\u0000\u0000\u0000\u0933\u0935"+
		"\u0003\u0154\u00aa\u0000\u0934\u0936\u0005%\u0000\u0000\u0935\u0934\u0001"+
		"\u0000\u0000\u0000\u0935\u0936\u0001\u0000\u0000\u0000\u0936\u0938\u0001"+
		"\u0000\u0000\u0000\u0937\u0933\u0001\u0000\u0000\u0000\u0937\u0938\u0001"+
		"\u0000\u0000\u0000\u0938\u0939\u0001\u0000\u0000\u0000\u0939\u093b\u0003"+
		"\u0132\u0099\u0000\u093a\u093c\u0003\u00cae\u0000\u093b\u093a\u0001\u0000"+
		"\u0000\u0000\u093b\u093c\u0001\u0000\u0000\u0000\u093c\u0145\u0001\u0000"+
		"\u0000\u0000\u093d\u093e\u0003t:\u0000\u093e\u093f\u0003\u0168\u00b4\u0000"+
		"\u093f\u0941\u0003\u0132\u0099\u0000\u0940\u0942\u0003\u00cae\u0000\u0941"+
		"\u0940\u0001\u0000\u0000\u0000\u0941\u0942\u0001\u0000\u0000\u0000\u0942"+
		"\u0147\u0001\u0000\u0000\u0000\u0943\u0944\u0005 \u0000\u0000\u0944\u0946"+
		"\u0003\u0132\u0099\u0000\u0945\u0947\u0003\u00cae\u0000\u0946\u0945\u0001"+
		"\u0000\u0000\u0000\u0946\u0947\u0001\u0000\u0000\u0000\u0947\u0948\u0001"+
		"\u0000\u0000\u0000\u0948\u0949\u0005#\u0000\u0000\u0949\u0149\u0001\u0000"+
		"\u0000\u0000\u094a\u094b\u0005\u001e\u0000\u0000\u094b\u094d\u0003\u0132"+
		"\u0099\u0000\u094c\u094e\u0003\u014c\u00a6\u0000\u094d\u094c\u0001\u0000"+
		"\u0000\u0000\u094d\u094e\u0001\u0000\u0000\u0000\u094e\u0950\u0001\u0000"+
		"\u0000\u0000\u094f\u0951\u0005%\u0000\u0000\u0950\u094f\u0001\u0000\u0000"+
		"\u0000\u0950\u0951\u0001\u0000\u0000\u0000\u0951\u0952\u0001\u0000\u0000"+
		"\u0000\u0952\u0954\u0003\u0132\u0099\u0000\u0953\u0955\u0003\u00cae\u0000"+
		"\u0954\u0953\u0001\u0000\u0000\u0000\u0954\u0955\u0001\u0000\u0000\u0000"+
		"\u0955\u0956\u0001\u0000\u0000\u0000\u0956\u0957\u0005!\u0000\u0000\u0957"+
		"\u014b\u0001\u0000\u0000\u0000\u0958\u0959\u0003\u00cae\u0000\u0959\u014d"+
		"\u0001\u0000\u0000\u0000\u095a\u095e\u0005%\u0000\u0000\u095b\u095d\u0005"+
		"Y\u0000\u0000\u095c\u095b\u0001\u0000\u0000\u0000\u095d\u0960\u0001\u0000"+
		"\u0000\u0000\u095e\u095c\u0001\u0000\u0000\u0000\u095e\u095f\u0001\u0000"+
		"\u0000\u0000\u095f\u0961\u0001\u0000\u0000\u0000\u0960\u095e\u0001\u0000"+
		"\u0000\u0000\u0961\u0962\u0003\u0150\u00a8\u0000\u0962\u014f\u0001\u0000"+
		"\u0000\u0000\u0963\u096f\u0003\u0152\u00a9\u0000\u0964\u0968\u0003\u0194"+
		"\u00ca\u0000\u0965\u0967\u0005Y\u0000\u0000\u0966\u0965\u0001\u0000\u0000"+
		"\u0000\u0967\u096a\u0001\u0000\u0000\u0000\u0968\u0966\u0001\u0000\u0000"+
		"\u0000\u0968\u0969\u0001\u0000\u0000\u0000\u0969\u096b\u0001\u0000\u0000"+
		"\u0000\u096a\u0968\u0001\u0000\u0000\u0000\u096b\u096c\u0003\u0152\u00a9"+
		"\u0000\u096c\u096e\u0001\u0000\u0000\u0000\u096d\u0964\u0001\u0000\u0000"+
		"\u0000\u096e\u0971\u0001\u0000\u0000\u0000\u096f\u096d\u0001\u0000\u0000"+
		"\u0000\u096f\u0970\u0001\u0000\u0000\u0000\u0970\u0973\u0001\u0000\u0000"+
		"\u0000\u0971\u096f\u0001\u0000\u0000\u0000\u0972\u0974\u0003\u0194\u00ca"+
		"\u0000\u0973\u0972\u0001\u0000\u0000\u0000\u0973\u0974\u0001\u0000\u0000"+
		"\u0000\u0974\u0151\u0001\u0000\u0000\u0000\u0975\u0977\u0003\u0134\u009a"+
		"\u0000\u0976\u0978\u0003\u00cae\u0000\u0977\u0976\u0001\u0000\u0000\u0000"+
		"\u0977\u0978\u0001\u0000\u0000\u0000\u0978\u0153\u0001\u0000\u0000\u0000"+
		"\u0979\u097c\u0005J\u0000\u0000\u097a\u097c\u0003\u015c\u00ae\u0000\u097b"+
		"\u0979\u0001\u0000\u0000\u0000\u097b\u097a\u0001\u0000\u0000\u0000\u097c"+
		"\u0155\u0001\u0000\u0000\u0000\u097d\u0980\u0005J\u0000\u0000\u097e\u0980"+
		"\u0003\u015e\u00af\u0000\u097f\u097d\u0001\u0000\u0000\u0000\u097f\u097e"+
		"\u0001\u0000\u0000\u0000\u0980\u0157\u0001\u0000\u0000\u0000\u0981\u0986"+
		"\u0003\u0154\u00aa\u0000\u0982\u0983\u0005\u001d\u0000\u0000\u0983\u0985"+
		"\u0003\u0154\u00aa\u0000\u0984\u0982\u0001\u0000\u0000\u0000\u0985\u0988"+
		"\u0001\u0000\u0000\u0000\u0986\u0984\u0001\u0000\u0000\u0000\u0986\u0987"+
		"\u0001\u0000\u0000\u0000\u0987\u0159\u0001\u0000\u0000\u0000\u0988\u0986"+
		"\u0001\u0000\u0000\u0000\u0989\u098a\u0007\u0002\u0000\u0000\u098a\u015b"+
		"\u0001\u0000\u0000\u0000\u098b\u098c\u0007\u0003\u0000\u0000\u098c\u015d"+
		"\u0001\u0000\u0000\u0000\u098d\u098e\u0007\u0004\u0000\u0000\u098e\u015f"+
		"\u0001\u0000\u0000\u0000\u098f\u0994\u0005Z\u0000\u0000\u0990\u0991\u0005"+
		"Y\u0000\u0000\u0991\u0993\u0005Z\u0000\u0000\u0992\u0990\u0001\u0000\u0000"+
		"\u0000\u0993\u0996\u0001\u0000\u0000\u0000\u0994\u0992\u0001\u0000\u0000"+
		"\u0000\u0994\u0995\u0001\u0000\u0000\u0000\u0995\u0161\u0001\u0000\u0000"+
		"\u0000\u0996\u0994\u0001\u0000\u0000\u0000\u0997\u099c\u0005[";
	private static final String _serializedATNSegment1 =
		"\u0000\u0000\u0998\u0999\u0005Y\u0000\u0000\u0999\u099b\u0005[\u0000\u0000"+
		"\u099a\u0998\u0001\u0000\u0000\u0000\u099b\u099e\u0001\u0000\u0000\u0000"+
		"\u099c\u099a\u0001\u0000\u0000\u0000\u099c\u099d\u0001\u0000\u0000\u0000"+
		"\u099d\u0163\u0001\u0000\u0000\u0000\u099e\u099c\u0001\u0000\u0000\u0000"+
		"\u099f\u09a0\u0005.\u0000\u0000\u09a0\u0165\u0001\u0000\u0000\u0000\u09a1"+
		"\u09a2\u0005-\u0000\u0000\u09a2\u0167\u0001\u0000\u0000\u0000\u09a3\u09a4"+
		"\u0005=\u0000\u0000\u09a4\u0169\u0001\u0000\u0000\u0000\u09a5\u09a6\u0005"+
		">\u0000\u0000\u09a6\u016b\u0001\u0000\u0000\u0000\u09a7\u09a8\u0005?\u0000"+
		"\u0000\u09a8\u016d\u0001\u0000\u0000\u0000\u09a9\u09aa\u0005@\u0000\u0000"+
		"\u09aa\u016f\u0001\u0000\u0000\u0000\u09ab\u09b2\u0003\u016a\u00b5\u0000"+
		"\u09ac\u09b2\u0003\u016c\u00b6\u0000\u09ad\u09b2\u0003\u016e\u00b7\u0000"+
		"\u09ae\u09b2\u0003\u0176\u00bb\u0000\u09af\u09b2\u0005\u0001\u0000\u0000"+
		"\u09b0\u09b2\u0005\u0014\u0000\u0000\u09b1\u09ab\u0001\u0000\u0000\u0000"+
		"\u09b1\u09ac\u0001\u0000\u0000\u0000\u09b1\u09ad\u0001\u0000\u0000\u0000"+
		"\u09b1\u09ae\u0001\u0000\u0000\u0000\u09b1\u09af\u0001\u0000\u0000\u0000"+
		"\u09b1\u09b0\u0001\u0000\u0000\u0000\u09b2\u0171\u0001\u0000\u0000\u0000"+
		"\u09b3\u09b6\u0003\u0176\u00bb\u0000\u09b4\u09b6\u0005\u0012\u0000\u0000"+
		"\u09b5\u09b3\u0001\u0000\u0000\u0000\u09b5\u09b4\u0001\u0000\u0000\u0000"+
		"\u09b6\u0173\u0001\u0000\u0000\u0000\u09b7\u09b8\u0007\u0005\u0000\u0000"+
		"\u09b8\u0175\u0001\u0000\u0000\u0000\u09b9\u09bb\u0003\u017c\u00be\u0000"+
		"\u09ba\u09bc\u0003\u0178\u00bc\u0000\u09bb\u09ba\u0001\u0000\u0000\u0000"+
		"\u09bb\u09bc\u0001\u0000\u0000\u0000\u09bc\u09c4\u0001\u0000\u0000\u0000"+
		"\u09bd\u09bf\u0003\u017e\u00bf\u0000\u09be\u09c0\u0003\u0180\u00c0\u0000"+
		"\u09bf\u09be\u0001\u0000\u0000\u0000\u09c0\u09c1\u0001\u0000\u0000\u0000"+
		"\u09c1\u09bf\u0001\u0000\u0000\u0000\u09c1\u09c2\u0001\u0000\u0000\u0000"+
		"\u09c2\u09c4\u0001\u0000\u0000\u0000\u09c3\u09b9\u0001\u0000\u0000\u0000"+
		"\u09c3\u09bd\u0001\u0000\u0000\u0000\u09c4\u0177\u0001\u0000\u0000\u0000"+
		"\u09c5\u09c6\u0004\u00bc\u000b\u0000\u09c6\u09c8\u0003\u017a\u00bd\u0000"+
		"\u09c7\u09c5\u0001\u0000\u0000\u0000\u09c8\u09c9\u0001\u0000\u0000\u0000"+
		"\u09c9\u09c7\u0001\u0000\u0000\u0000\u09c9\u09ca\u0001\u0000\u0000\u0000"+
		"\u09ca\u0179\u0001\u0000\u0000\u0000\u09cb\u09ce\u0003\u017c\u00be\u0000"+
		"\u09cc\u09ce\u0005]\u0000\u0000\u09cd\u09cb\u0001\u0000\u0000\u0000\u09cd"+
		"\u09cc\u0001\u0000\u0000\u0000\u09ce\u017b\u0001\u0000\u0000\u0000\u09cf"+
		"\u09d2\u0007\u0006\u0000\u0000\u09d0\u09d2\u0005K\u0000\u0000\u09d1\u09cf"+
		"\u0001\u0000\u0000\u0000\u09d1\u09d0\u0001\u0000\u0000\u0000\u09d2\u017d"+
		"\u0001\u0000\u0000\u0000\u09d3\u09d4\u0005\u001d\u0000\u0000\u09d4\u017f"+
		"\u0001\u0000\u0000\u0000\u09d5\u09d8\u0005\u001d\u0000\u0000\u09d6\u09d8"+
		"\u0003\u017a\u00bd\u0000\u09d7\u09d5\u0001\u0000\u0000\u0000\u09d7\u09d6"+
		"\u0001\u0000\u0000\u0000\u09d8\u0181\u0001\u0000\u0000\u0000\u09d9\u09de"+
		"\u0003\u0188\u00c4\u0000\u09da\u09de\u0003\u018c\u00c6\u0000\u09db\u09de"+
		"\u0003\u0184\u00c2\u0000\u09dc\u09de\u0003\u0186\u00c3\u0000\u09dd\u09d9"+
		"\u0001\u0000\u0000\u0000\u09dd\u09da\u0001\u0000\u0000\u0000\u09dd\u09db"+
		"\u0001\u0000\u0000\u0000\u09dd\u09dc\u0001\u0000\u0000\u0000\u09de\u0183"+
		"\u0001\u0000\u0000\u0000\u09df\u09e0\u0007\u0007\u0000\u0000\u09e0\u0185"+
		"\u0001\u0000\u0000\u0000\u09e1\u09e2\u0005\u0013\u0000\u0000\u09e2\u0187"+
		"\u0001\u0000\u0000\u0000\u09e3\u09e5\u0003\u0166\u00b3\u0000\u09e4\u09e3"+
		"\u0001\u0000\u0000\u0000\u09e4\u09e5\u0001\u0000\u0000\u0000\u09e5\u09e6"+
		"\u0001\u0000\u0000\u0000\u09e6\u09ec\u0003\u018a\u00c5\u0000\u09e7\u09e9"+
		"\u0003\u0166\u00b3\u0000\u09e8\u09e7\u0001\u0000\u0000\u0000\u09e8\u09e9"+
		"\u0001\u0000\u0000\u0000\u09e9\u09ea\u0001\u0000\u0000\u0000\u09ea\u09ec"+
		"\u0005R\u0000\u0000\u09eb\u09e4\u0001\u0000\u0000\u0000\u09eb\u09e8\u0001"+
		"\u0000\u0000\u0000\u09ec\u0189\u0001\u0000\u0000\u0000\u09ed\u09ee\u0007"+
		"\b\u0000\u0000\u09ee\u018b\u0001\u0000\u0000\u0000\u09ef\u09f0\u0007\t"+
		"\u0000\u0000\u09f0\u018d\u0001\u0000\u0000\u0000\u09f1\u09f2\u0007\n\u0000"+
		"\u0000\u09f2\u018f\u0001\u0000\u0000\u0000\u09f3\u09f4\u0007\u000b\u0000"+
		"\u0000\u09f4\u0191\u0001\u0000\u0000\u0000\u09f5\u09f9\u0005&\u0000\u0000"+
		"\u09f6\u09f7\u0003\u0162\u00b1\u0000\u09f7\u09f8\u0005Y\u0000\u0000\u09f8"+
		"\u09fa\u0001\u0000\u0000\u0000\u09f9\u09f6\u0001\u0000\u0000\u0000\u09f9"+
		"\u09fa\u0001\u0000\u0000\u0000\u09fa\u0a00\u0001\u0000\u0000\u0000\u09fb"+
		"\u09fd\u0003\u0162\u00b1\u0000\u09fc\u09fb\u0001\u0000\u0000\u0000\u09fc"+
		"\u09fd\u0001\u0000\u0000\u0000\u09fd\u09fe\u0001\u0000\u0000\u0000\u09fe"+
		"\u0a00\u0005Y\u0000\u0000\u09ff\u09f5\u0001\u0000\u0000\u0000\u09ff\u09fc"+
		"\u0001\u0000\u0000\u0000\u0a00\u0193\u0001\u0000\u0000\u0000\u0a01\u0a05"+
		"\u0005$\u0000\u0000\u0a02\u0a03\u0003\u0162\u00b1\u0000\u0a03\u0a04\u0005"+
		"Y\u0000\u0000\u0a04\u0a06\u0001\u0000\u0000\u0000\u0a05\u0a02\u0001\u0000"+
		"\u0000\u0000\u0a05\u0a06\u0001\u0000\u0000\u0000\u0a06\u0a0c\u0001\u0000"+
		"\u0000\u0000\u0a07\u0a09\u0003\u0162\u00b1\u0000\u0a08\u0a07\u0001\u0000"+
		"\u0000\u0000\u0a08\u0a09\u0001\u0000\u0000\u0000\u0a09\u0a0a\u0001\u0000"+
		"\u0000\u0000\u0a0a\u0a0c\u0005Y\u0000\u0000\u0a0b\u0a01\u0001\u0000\u0000"+
		"\u0000\u0a0b\u0a08\u0001\u0000\u0000\u0000\u0a0c\u0195\u0001\u0000\u0000"+
		"\u0000\u0168\u0199\u019d\u01a2\u01ab\u01b1\u01c0\u01c7\u01cb\u01cf\u01d6"+
		"\u01de\u01e8\u01f2\u01f9\u01fe\u0205\u020c\u0213\u021a\u021e\u0224\u022c"+
		"\u0230\u0237\u023e\u0242\u0247\u0251\u0258\u025c\u0260\u0265\u026c\u0271"+
		"\u0276\u027e\u0284\u028b\u0295\u029c\u02a0\u02a9\u02af\u02b6\u02c0\u02c7"+
		"\u02cc\u02d1\u02d7\u02da\u02e7\u02ed\u02f1\u02f6\u0300\u0304\u030b\u0316"+
		"\u031d\u032b\u0335\u033c\u0343\u0349\u034e\u0352\u0356\u035a\u035e\u0367"+
		"\u036f\u0376\u037a\u037f\u0384\u038b\u0390\u0396\u039c\u03ac\u03b4\u03bb"+
		"\u03bf\u03c4\u03c9\u03ce\u03d3\u03de\u03e3\u03e6\u03eb\u03f1\u03f5\u03f9"+
		"\u03fd\u0402\u0406\u040b\u040d\u0411\u0416\u041a\u041e\u0423\u0427\u042c"+
		"\u0437\u043e\u0443\u044a\u0451\u0455\u045c\u0460\u0464\u0469\u046b\u0470"+
		"\u0475\u0479\u047e\u0488\u048d\u0491\u0496\u049b\u049f\u04a4\u04a8\u04ad"+
		"\u04b9\u04c0\u04c4\u04c9\u04ce\u04d2\u04d7\u04db\u04de\u04e3\u04ec\u04f0"+
		"\u04f7\u04fb\u0500\u0504\u050a\u050f\u0513\u0518\u0522\u0529\u052d\u0532"+
		"\u0537\u053e\u0541\u0548\u054c\u0551\u055a\u055e\u0565\u0569\u056e\u0575"+
		"\u057a\u057e\u0583\u058d\u0594\u0598\u059d\u05a2\u05a6\u05a9\u05ad\u05b2"+
		"\u05ba\u05c1\u05c5\u05c8\u05ca\u05cf\u05d4\u05dd\u05e3\u05e7\u05ea\u05ef"+
		"\u05f3\u05f7\u05fb\u05ff\u0604\u0609\u0610\u0619\u0622\u0627\u062b\u0634"+
		"\u0639\u063c\u0641\u064d\u0650\u0652\u0657\u0661\u0665\u066a\u0672\u067b"+
		"\u0682\u0686\u068a\u068f\u0694\u069a\u06ab\u06b4\u06b9\u06c6\u06ca\u06ce"+
		"\u06d3\u06d9\u06e2\u06eb\u06f4\u06f9\u0704\u0708\u070d\u0717\u071e\u0722"+
		"\u0726\u072c\u0730\u0735\u073f\u0746\u074a\u074e\u0757\u075b\u0760\u076a"+
		"\u0771\u0775\u077a\u0786\u078d\u0791\u0796\u07a0\u07a7\u07ab\u07b1\u07b8"+
		"\u07c1\u07c8\u07d2\u07d9\u07dd\u07e2\u07e8\u07ef\u07f6\u07fa\u07fe\u0807"+
		"\u080e\u081b\u081f\u0827\u082f\u0833\u083a\u0844\u0846\u084d\u0851\u0857"+
		"\u085f\u0866\u086b\u0872\u087a\u087e\u0883\u088d\u0896\u08a4\u08a6\u08ae"+
		"\u08b3\u08b8\u08bf\u08c4\u08c9\u08cd\u08d2\u08d7\u08de\u08e3\u08e8\u08ea"+
		"\u08ec\u08f3\u08f6\u08fa\u08ff\u0906\u090b\u0913\u0917\u091c\u0926\u092d"+
		"\u0931\u0935\u0937\u093b\u0941\u0946\u094d\u0950\u0954\u095e\u0968\u096f"+
		"\u0973\u0977\u097b\u097f\u0986\u0994\u099c\u09b1\u09b5\u09bb\u09c1\u09c3"+
		"\u09c9\u09cd\u09d1\u09d7\u09dd\u09e4\u09e8\u09eb\u09f9\u09fc\u09ff\u0a05"+
		"\u0a08\u0a0b";
	public static final String _serializedATN = Utils.join(
		new String[] {
			_serializedATNSegment0,
			_serializedATNSegment1
		},
		""
	);
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}