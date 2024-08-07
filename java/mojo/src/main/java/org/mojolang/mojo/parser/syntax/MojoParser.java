// Generated from java-escape by ANTLR 4.11.1
package org.mojolang.mojo.parser.syntax;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class MojoParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.11.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		KEYWORD_AND=1, KEYWORD_AS=2, KEYWORD_ATTRIBUTE=3, KEYWORD_BREAK=4, KEYWORD_CONST=5, 
		KEYWORD_CONTINUE=6, KEYWORD_ELSE=7, KEYWORD_ENUM=8, KEYWORD_FALSE=9, KEYWORD_FOR=10, 
		KEYWORD_FUNC=11, KEYWORD_IF=12, KEYWORD_IMPORT=13, KEYWORD_IN=14, KEYWORD_INTERFACE=15, 
		KEYWORD_IS=16, KEYWORD_MATCH=17, KEYWORD_NOT=18, KEYWORD_NULL=19, KEYWORD_OR=20, 
		KEYWORD_PACKAGE=21, KEYWORD_REPEATE=22, KEYWORD_RETURN=23, KEYWORD_STRUCT=24, 
		KEYWORD_TRUE=25, KEYWORD_TYPE=26, KEYWORD_VAR=27, KEYWORD_WHILE=28, KEYWORD_XOR=29, 
		DOT=30, LCURLY=31, LPAREN=32, LBRACK=33, RCURLY=34, RPAREN=35, RBRACK=36, 
		COMMA=37, COLON=38, SEMI=39, LT=40, GT=41, BANG=42, QUESTION=43, AT=44, 
		AND=45, MINUS=46, EQUAL=47, PIPE=48, SLASH=49, PLUS=50, STAR=51, PERCENT=52, 
		CARET=53, TILDE=54, DOLLER=55, BACKTICK=56, UNDERSCORE=57, PLUS_PLUS=58, 
		MINUS_MINUS=59, COLON_EQUAL=60, RIGHT_RIGHT_ARROWS=61, RIGHT_ARROW=62, 
		DOT_DOT=63, DOT_DOT_LT=64, ELLIPSIS=65, GRAPH_RIGHT_PATH=66, GRAPH_LEFT_PATH=67, 
		GRAPH_PATH=68, GRAPH_CONSTRAINT_PATH_LEFT=69, GRAPH_CONSTRAINT_PATH_LEFT_ARROW=70, 
		GRAPH_CONSTRAINT_PATH_RIGHT=71, GRAPH_CONSTRAINT_PATH_RIGHT_ARROW=72, 
		TYPE_IDENTIFIER=73, VALUE_IDENTIFIER=74, OPERATOR_HEAD_OTHER=75, IMPLICIT_PARAMETER_NAME=76, 
		BINARY_LITERAL=77, OCTAL_LITERAL=78, DECIMAL_LITERAL=79, PURE_DECIMAL_DIGITS=80, 
		HEXADECIMAL_LITERAL=81, FLOAT_LITERAL=82, STATIC_STRING_LITERAL=83, INTERPOLATED_STRING_LITERAL=84, 
		WS=85, BLOCK_COMMENT=86, LINE_COMMENT=87, LINE_COMMENT_DISTINCT_DOCUMENT=88, 
		EOL=89, LINE_DOCUMENT=90, FOLLOWING_LINE_DOCUMENT=91, INNER_LINE_DOCUMENT=92, 
		OPERATOR_FOLLOWING_CHARACTER=93;
	public static final int
		RULE_mojoFile = 0, RULE_statement = 1, RULE_freeFloatingDocument = 2, 
		RULE_statements = 3, RULE_loopStatement = 4, RULE_forInStatement = 5, 
		RULE_whileStatement = 6, RULE_conditions = 7, RULE_condition = 8, RULE_optionalBindingCondition = 9, 
		RULE_branchStatement = 10, RULE_ifStatement = 11, RULE_elseClause = 12, 
		RULE_matchStatement = 13, RULE_matchCases = 14, RULE_matchCase = 15, RULE_controlTransferStatement = 16, 
		RULE_breakStatement = 17, RULE_continueStatement = 18, RULE_returnStatement = 19, 
		RULE_genericParameterClause = 20, RULE_genericParameters = 21, RULE_genericParameter = 22, 
		RULE_genericArgumentClause = 23, RULE_genericArguments = 24, RULE_genericArgument = 25, 
		RULE_declaration = 26, RULE_codeBlock = 27, RULE_packageDeclaration = 28, 
		RULE_packageIdentifier = 29, RULE_packageName = 30, RULE_importDeclaration = 31, 
		RULE_importPath = 32, RULE_importPathIdentifier = 33, RULE_importAllClause = 34, 
		RULE_importValueAsClause = 35, RULE_importTypeClause = 36, RULE_importTypeAsClause = 37, 
		RULE_importGroupClause = 38, RULE_importGroup = 39, RULE_importValue = 40, 
		RULE_importType = 41, RULE_constantDeclaration = 42, RULE_patternInitializers = 43, 
		RULE_documentedPatternInitializer = 44, RULE_patternInitializer = 45, 
		RULE_initializer = 46, RULE_variableDeclaration = 47, RULE_typeAliasDeclaration = 48, 
		RULE_typeAliasName = 49, RULE_typeAliasAssignment = 50, RULE_functionDeclaration = 51, 
		RULE_functionName = 52, RULE_functionSignature = 53, RULE_functionResult = 54, 
		RULE_functionBody = 55, RULE_functionParameterClause = 56, RULE_functionParameters = 57, 
		RULE_functionParameter = 58, RULE_enumDeclaration = 59, RULE_enumBody = 60, 
		RULE_enumName = 61, RULE_enumMembers = 62, RULE_enumMember = 63, RULE_structDeclaration = 64, 
		RULE_structName = 65, RULE_structType = 66, RULE_structBody = 67, RULE_structMembers = 68, 
		RULE_structMember = 69, RULE_structMemberDeclaration = 70, RULE_interfaceDeclaration = 71, 
		RULE_interfaceName = 72, RULE_interfaceType = 73, RULE_interfaceBody = 74, 
		RULE_interfaceMembers = 75, RULE_interfaceMember = 76, RULE_interfaceMethodDeclaration = 77, 
		RULE_attributeDeclaration = 78, RULE_attributeAliasDeclaration = 79, RULE_attributeAliasAssignment = 80, 
		RULE_pattern = 81, RULE_wildcardPattern = 82, RULE_identifierPattern = 83, 
		RULE_tuplePattern = 84, RULE_tuplePatternElementList = 85, RULE_tuplePatternElement = 86, 
		RULE_optionalPattern = 87, RULE_expressionPattern = 88, RULE_attribute = 89, 
		RULE_attributeIdentifier = 90, RULE_attributeName = 91, RULE_attributeArgumentClause = 92, 
		RULE_attributeArgument = 93, RULE_attributeArguments = 94, RULE_attributes = 95, 
		RULE_expression = 96, RULE_expressions = 97, RULE_prefixExpression = 98, 
		RULE_binaryExpression = 99, RULE_binaryExpressions = 100, RULE_conditionalOperator = 101, 
		RULE_typeCastingOperator = 102, RULE_primaryExpression = 103, RULE_literalExpression = 104, 
		RULE_numericOperatorLiteral = 105, RULE_stringOperatorLiteral = 106, RULE_suffixLiteralOperator = 107, 
		RULE_prefixLiteralOperator = 108, RULE_arrayLiteral = 109, RULE_arrayLiteralItems = 110, 
		RULE_arrayLiteralItem = 111, RULE_mapLiteral = 112, RULE_mapLiteralItems = 113, 
		RULE_mapLiteralItem = 114, RULE_objectLiteral = 115, RULE_objectLiteralItems = 116, 
		RULE_objectLiteralItem = 117, RULE_structLiteral = 118, RULE_structConstructionExpression = 119, 
		RULE_closureExpression = 120, RULE_closureParameters = 121, RULE_closureParameter = 122, 
		RULE_implicitMemberExpression = 123, RULE_parenthesizedExpression = 124, 
		RULE_tupleExpression = 125, RULE_tupleElement = 126, RULE_wildcardExpression = 127, 
		RULE_postfixExpression = 128, RULE_suffixExpression = 129, RULE_explicitMemberSuffix = 130, 
		RULE_subscriptSuffix = 131, RULE_functionCallSuffix = 132, RULE_functionCallArgumentClause = 133, 
		RULE_functionCallArguments = 134, RULE_functionCallArgument = 135, RULE_trailingClosures = 136, 
		RULE_labeledTrailingClosures = 137, RULE_labeledTrailingClosure = 138, 
		RULE_argumentNames = 139, RULE_argumentName = 140, RULE_type_ = 141, RULE_basicType = 142, 
		RULE_primeType = 143, RULE_typeAnnotation = 144, RULE_typeIdentifier = 145, 
		RULE_typeIdentifierClause = 146, RULE_typeName = 147, RULE_tupleType = 148, 
		RULE_tupleTypeElements = 149, RULE_tupleTypeElement = 150, RULE_functionType = 151, 
		RULE_arrayType = 152, RULE_mapType = 153, RULE_keyAttributes = 154, RULE_typeInheritanceClause = 155, 
		RULE_typeInheritances = 156, RULE_typeInheritance = 157, RULE_declarationIdentifier = 158, 
		RULE_labelIdentifier = 159, RULE_pathIdentifier = 160, RULE_identifier = 161, 
		RULE_keywordAsIdentifierInDeclarations = 162, RULE_keywordAsIdentifierInLabels = 163, 
		RULE_document = 164, RULE_followingDocument = 165, RULE_assignmentOperator = 166, 
		RULE_negatePrefixOperator = 167, RULE_arrowOperator = 168, RULE_rangeOperator = 169, 
		RULE_halfOpenRangeOperator = 170, RULE_binaryOperator = 171, RULE_prefixOperator = 172, 
		RULE_postfixOperator = 173, RULE_operator = 174, RULE_operator_characters = 175, 
		RULE_operator_character = 176, RULE_operator_head = 177, RULE_dot_operator_head = 178, 
		RULE_dot_operator_character = 179, RULE_literal = 180, RULE_boolLiteral = 181, 
		RULE_nullLiteral = 182, RULE_numericLiteral = 183, RULE_integerLiteral = 184, 
		RULE_stringLiteral = 185, RULE_eos = 186, RULE_eov = 187, RULE_eosWithDocument = 188, 
		RULE_eovWithDocument = 189;
	private static String[] makeRuleNames() {
		return new String[] {
			"mojoFile", "statement", "freeFloatingDocument", "statements", "loopStatement", 
			"forInStatement", "whileStatement", "conditions", "condition", "optionalBindingCondition", 
			"branchStatement", "ifStatement", "elseClause", "matchStatement", "matchCases", 
			"matchCase", "controlTransferStatement", "breakStatement", "continueStatement", 
			"returnStatement", "genericParameterClause", "genericParameters", "genericParameter", 
			"genericArgumentClause", "genericArguments", "genericArgument", "declaration", 
			"codeBlock", "packageDeclaration", "packageIdentifier", "packageName", 
			"importDeclaration", "importPath", "importPathIdentifier", "importAllClause", 
			"importValueAsClause", "importTypeClause", "importTypeAsClause", "importGroupClause", 
			"importGroup", "importValue", "importType", "constantDeclaration", "patternInitializers", 
			"documentedPatternInitializer", "patternInitializer", "initializer", 
			"variableDeclaration", "typeAliasDeclaration", "typeAliasName", "typeAliasAssignment", 
			"functionDeclaration", "functionName", "functionSignature", "functionResult", 
			"functionBody", "functionParameterClause", "functionParameters", "functionParameter", 
			"enumDeclaration", "enumBody", "enumName", "enumMembers", "enumMember", 
			"structDeclaration", "structName", "structType", "structBody", "structMembers", 
			"structMember", "structMemberDeclaration", "interfaceDeclaration", "interfaceName", 
			"interfaceType", "interfaceBody", "interfaceMembers", "interfaceMember", 
			"interfaceMethodDeclaration", "attributeDeclaration", "attributeAliasDeclaration", 
			"attributeAliasAssignment", "pattern", "wildcardPattern", "identifierPattern", 
			"tuplePattern", "tuplePatternElementList", "tuplePatternElement", "optionalPattern", 
			"expressionPattern", "attribute", "attributeIdentifier", "attributeName", 
			"attributeArgumentClause", "attributeArgument", "attributeArguments", 
			"attributes", "expression", "expressions", "prefixExpression", "binaryExpression", 
			"binaryExpressions", "conditionalOperator", "typeCastingOperator", "primaryExpression", 
			"literalExpression", "numericOperatorLiteral", "stringOperatorLiteral", 
			"suffixLiteralOperator", "prefixLiteralOperator", "arrayLiteral", "arrayLiteralItems", 
			"arrayLiteralItem", "mapLiteral", "mapLiteralItems", "mapLiteralItem", 
			"objectLiteral", "objectLiteralItems", "objectLiteralItem", "structLiteral", 
			"structConstructionExpression", "closureExpression", "closureParameters", 
			"closureParameter", "implicitMemberExpression", "parenthesizedExpression", 
			"tupleExpression", "tupleElement", "wildcardExpression", "postfixExpression", 
			"suffixExpression", "explicitMemberSuffix", "subscriptSuffix", "functionCallSuffix", 
			"functionCallArgumentClause", "functionCallArguments", "functionCallArgument", 
			"trailingClosures", "labeledTrailingClosures", "labeledTrailingClosure", 
			"argumentNames", "argumentName", "type_", "basicType", "primeType", "typeAnnotation", 
			"typeIdentifier", "typeIdentifierClause", "typeName", "tupleType", "tupleTypeElements", 
			"tupleTypeElement", "functionType", "arrayType", "mapType", "keyAttributes", 
			"typeInheritanceClause", "typeInheritances", "typeInheritance", "declarationIdentifier", 
			"labelIdentifier", "pathIdentifier", "identifier", "keywordAsIdentifierInDeclarations", 
			"keywordAsIdentifierInLabels", "document", "followingDocument", "assignmentOperator", 
			"negatePrefixOperator", "arrowOperator", "rangeOperator", "halfOpenRangeOperator", 
			"binaryOperator", "prefixOperator", "postfixOperator", "operator", "operator_characters", 
			"operator_character", "operator_head", "dot_operator_head", "dot_operator_character", 
			"literal", "boolLiteral", "nullLiteral", "numericLiteral", "integerLiteral", 
			"stringLiteral", "eos", "eov", "eosWithDocument", "eovWithDocument"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'and'", "'as'", "'attribute'", "'break'", "'const'", "'continue'", 
			"'else'", "'enum'", "'false'", "'for'", "'func'", "'if'", "'import'", 
			"'in'", "'interface'", "'is'", "'match'", "'not'", "'null'", "'or'", 
			"'package'", "'repeat'", "'return'", "'struct'", "'true'", "'type'", 
			"'var'", "'while'", "'xor'", "'.'", "'{'", "'('", "'['", "'}'", "')'", 
			"']'", "','", "':'", "';'", "'<'", "'>'", "'!'", "'?'", "'@'", "'&'", 
			"'-'", "'='", "'|'", "'/'", "'+'", "'*'", "'%'", "'^'", "'~'", "'$'", 
			"'`'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "KEYWORD_AND", "KEYWORD_AS", "KEYWORD_ATTRIBUTE", "KEYWORD_BREAK", 
			"KEYWORD_CONST", "KEYWORD_CONTINUE", "KEYWORD_ELSE", "KEYWORD_ENUM", 
			"KEYWORD_FALSE", "KEYWORD_FOR", "KEYWORD_FUNC", "KEYWORD_IF", "KEYWORD_IMPORT", 
			"KEYWORD_IN", "KEYWORD_INTERFACE", "KEYWORD_IS", "KEYWORD_MATCH", "KEYWORD_NOT", 
			"KEYWORD_NULL", "KEYWORD_OR", "KEYWORD_PACKAGE", "KEYWORD_REPEATE", "KEYWORD_RETURN", 
			"KEYWORD_STRUCT", "KEYWORD_TRUE", "KEYWORD_TYPE", "KEYWORD_VAR", "KEYWORD_WHILE", 
			"KEYWORD_XOR", "DOT", "LCURLY", "LPAREN", "LBRACK", "RCURLY", "RPAREN", 
			"RBRACK", "COMMA", "COLON", "SEMI", "LT", "GT", "BANG", "QUESTION", "AT", 
			"AND", "MINUS", "EQUAL", "PIPE", "SLASH", "PLUS", "STAR", "PERCENT", 
			"CARET", "TILDE", "DOLLER", "BACKTICK", "UNDERSCORE", "PLUS_PLUS", "MINUS_MINUS", 
			"COLON_EQUAL", "RIGHT_RIGHT_ARROWS", "RIGHT_ARROW", "DOT_DOT", "DOT_DOT_LT", 
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
	public String getGrammarFileName() { return "java-escape"; }

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
			setState(383);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(380);
					match(EOL);
					}
					} 
				}
				setState(385);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(387);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((_la) & ~0x3f) == 0 && ((1L << _la) & 180142902763061118L) != 0 || (((_la - 73)) & ~0x3f) == 0 && ((1L << (_la - 73)) & 135159L) != 0) {
				{
				setState(386);
				statements();
				}
			}

			setState(392);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(389);
				match(EOL);
				}
				}
				setState(394);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(395);
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
		public LoopStatementContext loopStatement() {
			return getRuleContext(LoopStatementContext.class,0);
		}
		public BranchStatementContext branchStatement() {
			return getRuleContext(BranchStatementContext.class,0);
		}
		public ControlTransferStatementContext controlTransferStatement() {
			return getRuleContext(ControlTransferStatementContext.class,0);
		}
		public FreeFloatingDocumentContext freeFloatingDocument() {
			return getRuleContext(FreeFloatingDocumentContext.class,0);
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
			setState(403);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(397);
				declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(398);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(399);
				loopStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(400);
				branchStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(401);
				controlTransferStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(402);
				freeFloatingDocument();
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
	public static class FreeFloatingDocumentContext extends ParserRuleContext {
		public DocumentContext document() {
			return getRuleContext(DocumentContext.class,0);
		}
		public FreeFloatingDocumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_freeFloatingDocument; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitFreeFloatingDocument(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FreeFloatingDocumentContext freeFloatingDocument() throws RecognitionException {
		FreeFloatingDocumentContext _localctx = new FreeFloatingDocumentContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_freeFloatingDocument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
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
		enterRule(_localctx, 6, RULE_statements);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(407);
			statement();
			setState(419);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(408);
					eos();
					setState(412);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(409);
						match(EOL);
						}
						}
						setState(414);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(415);
					statement();
					}
					} 
				}
				setState(421);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			setState(423);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(422);
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
		enterRule(_localctx, 8, RULE_loopStatement);
		try {
			setState(427);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_FOR:
				enterOuterAlt(_localctx, 1);
				{
				setState(425);
				forInStatement();
				}
				break;
			case KEYWORD_WHILE:
				enterOuterAlt(_localctx, 2);
				{
				setState(426);
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
		enterRule(_localctx, 10, RULE_forInStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(429);
			match(KEYWORD_FOR);
			setState(430);
			pattern(0);
			setState(434);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(431);
				match(EOL);
				}
				}
				setState(436);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(437);
			match(KEYWORD_IN);
			setState(438);
			expression();
			setState(442);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(439);
				match(EOL);
				}
				}
				setState(444);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(445);
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
		enterRule(_localctx, 12, RULE_whileStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(447);
			match(KEYWORD_WHILE);
			setState(448);
			conditions();
			setState(452);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(449);
				match(EOL);
				}
				}
				setState(454);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(455);
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
		enterRule(_localctx, 14, RULE_conditions);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			condition();
			setState(469);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(458);
					eov();
					setState(462);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(459);
						match(EOL);
						}
						}
						setState(464);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(465);
					condition();
					}
					} 
				}
				setState(471);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
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
		enterRule(_localctx, 16, RULE_condition);
		try {
			setState(474);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
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
			case KEYWORD_REPEATE:
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
				setState(472);
				expression();
				}
				break;
			case KEYWORD_VAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(473);
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
		enterRule(_localctx, 18, RULE_optionalBindingCondition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(476);
			match(KEYWORD_VAR);
			setState(477);
			pattern(0);
			setState(481);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(478);
				match(EOL);
				}
				}
				setState(483);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(484);
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
		enterRule(_localctx, 20, RULE_branchStatement);
		try {
			setState(488);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(486);
				ifStatement();
				}
				break;
			case KEYWORD_MATCH:
				enterOuterAlt(_localctx, 2);
				{
				setState(487);
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
		enterRule(_localctx, 22, RULE_ifStatement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(490);
			match(KEYWORD_IF);
			setState(491);
			conditions();
			setState(495);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(492);
				match(EOL);
				}
				}
				setState(497);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(498);
			codeBlock();
			setState(502);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(499);
					match(EOL);
					}
					} 
				}
				setState(504);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			setState(506);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_ELSE) {
				{
				setState(505);
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
		enterRule(_localctx, 24, RULE_elseClause);
		int _la;
		try {
			setState(524);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(508);
				match(KEYWORD_ELSE);
				setState(512);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(509);
					match(EOL);
					}
					}
					setState(514);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(515);
				codeBlock();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(516);
				match(KEYWORD_ELSE);
				setState(520);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(517);
					match(EOL);
					}
					}
					setState(522);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(523);
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
		enterRule(_localctx, 26, RULE_matchStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			match(KEYWORD_MATCH);
			setState(527);
			expression();
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
			match(LCURLY);
			setState(542);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				{
				setState(538);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(535);
					match(EOL);
					}
					}
					setState(540);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(541);
				matchCases();
				}
				break;
			}
			setState(547);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(544);
				match(EOL);
				}
				}
				setState(549);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(550);
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
		enterRule(_localctx, 28, RULE_matchCases);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(552);
			matchCase();
			setState(564);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(553);
					eos();
					setState(557);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(554);
						match(EOL);
						}
						}
						setState(559);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(560);
					matchCase();
					}
					} 
				}
				setState(566);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			setState(567);
			eos();
			}
		}
		catch (RecognitionException re) {
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
		enterRule(_localctx, 30, RULE_matchCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(569);
			pattern(0);
			setState(573);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(570);
				match(EOL);
				}
				}
				setState(575);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(576);
			match(RIGHT_RIGHT_ARROWS);
			setState(580);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(577);
				match(EOL);
				}
				}
				setState(582);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(585);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				{
				setState(583);
				codeBlock();
				}
				break;
			case 2:
				{
				setState(584);
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
		enterRule(_localctx, 32, RULE_controlTransferStatement);
		try {
			setState(590);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_BREAK:
				enterOuterAlt(_localctx, 1);
				{
				setState(587);
				breakStatement();
				}
				break;
			case KEYWORD_CONTINUE:
				enterOuterAlt(_localctx, 2);
				{
				setState(588);
				continueStatement();
				}
				break;
			case KEYWORD_RETURN:
				enterOuterAlt(_localctx, 3);
				{
				setState(589);
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
		enterRule(_localctx, 34, RULE_breakStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(592);
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
		enterRule(_localctx, 36, RULE_continueStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(594);
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
		enterRule(_localctx, 38, RULE_returnStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(596);
			match(KEYWORD_RETURN);
			setState(598);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((_la) & ~0x3f) == 0 && ((1L << _la) & 180125310165969726L) != 0 || (((_la - 73)) & ~0x3f) == 0 && ((1L << (_la - 73)) & 4087L) != 0) {
				{
				setState(597);
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
		enterRule(_localctx, 40, RULE_genericParameterClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(600);
			match(LT);
			setState(604);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(601);
				match(EOL);
				}
				}
				setState(606);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(607);
			genericParameters();
			setState(611);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(608);
				match(EOL);
				}
				}
				setState(613);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(614);
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
		enterRule(_localctx, 42, RULE_genericParameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(616);
			genericParameter();
			setState(628);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(617);
					eovWithDocument();
					setState(621);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(618);
						match(EOL);
						}
						}
						setState(623);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(624);
					genericParameter();
					}
					} 
				}
				setState(630);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			}
			setState(632);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				{
				setState(631);
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
		enterRule(_localctx, 44, RULE_genericParameter);
		try {
			setState(641);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(634);
				typeName();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(635);
				typeName();
				setState(636);
				match(ELLIPSIS);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(638);
				typeName();
				setState(639);
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
		enterRule(_localctx, 46, RULE_genericArgumentClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(643);
			match(LT);
			setState(647);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(644);
				match(EOL);
				}
				}
				setState(649);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(650);
			genericArguments();
			setState(654);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(651);
				match(EOL);
				}
				}
				setState(656);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(657);
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
		enterRule(_localctx, 48, RULE_genericArguments);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(659);
			genericArgument();
			setState(671);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,42,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(660);
					eov();
					setState(664);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(661);
						match(EOL);
						}
						}
						setState(666);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(667);
					genericArgument();
					}
					} 
				}
				setState(673);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,42,_ctx);
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
		enterRule(_localctx, 50, RULE_genericArgument);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(674);
			type_(0);
			setState(676);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(675);
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
		enterRule(_localctx, 52, RULE_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(681);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LINE_DOCUMENT) {
				{
				setState(678);
				document();
				setState(679);
				match(EOL);
				}
			}

			setState(690);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(683);
				attributes();
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
				}
			}

			setState(703);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,47,_ctx) ) {
			case 1:
				{
				setState(692);
				packageDeclaration();
				}
				break;
			case 2:
				{
				setState(693);
				importDeclaration();
				}
				break;
			case 3:
				{
				setState(694);
				constantDeclaration();
				}
				break;
			case 4:
				{
				setState(695);
				variableDeclaration();
				}
				break;
			case 5:
				{
				setState(696);
				typeAliasDeclaration();
				}
				break;
			case 6:
				{
				setState(697);
				functionDeclaration();
				}
				break;
			case 7:
				{
				setState(698);
				enumDeclaration();
				}
				break;
			case 8:
				{
				setState(699);
				structDeclaration();
				}
				break;
			case 9:
				{
				setState(700);
				interfaceDeclaration();
				}
				break;
			case 10:
				{
				setState(701);
				attributeDeclaration();
				}
				break;
			case 11:
				{
				setState(702);
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
		enterRule(_localctx, 54, RULE_codeBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(705);
			match(LCURLY);
			setState(713);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				{
				setState(709);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(706);
					match(EOL);
					}
					}
					setState(711);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(712);
				statements();
				}
				break;
			}
			setState(718);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(715);
				match(EOL);
				}
				}
				setState(720);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(721);
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
		enterRule(_localctx, 56, RULE_packageDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(723);
			match(KEYWORD_PACKAGE);
			setState(724);
			packageIdentifier();
			setState(732);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
			case 1:
				{
				setState(728);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(725);
					match(EOL);
					}
					}
					setState(730);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(731);
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
		enterRule(_localctx, 58, RULE_packageIdentifier);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(734);
			packageName();
			setState(739);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,53,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(735);
					match(DOT);
					setState(736);
					packageName();
					}
					} 
				}
				setState(741);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,53,_ctx);
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
		enterRule(_localctx, 60, RULE_packageName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(742);
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
		enterRule(_localctx, 62, RULE_importDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(744);
			match(KEYWORD_IMPORT);
			setState(745);
			importPath();
			setState(750);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
			case 1:
				{
				setState(746);
				importAllClause();
				}
				break;
			case 2:
				{
				setState(747);
				importValueAsClause();
				}
				break;
			case 3:
				{
				setState(748);
				importTypeClause();
				}
				break;
			case 4:
				{
				setState(749);
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
		enterRule(_localctx, 64, RULE_importPath);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(752);
			importPathIdentifier();
			setState(757);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,55,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(753);
					match(DOT);
					setState(754);
					importPathIdentifier();
					}
					} 
				}
				setState(759);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,55,_ctx);
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
		enterRule(_localctx, 66, RULE_importPathIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(760);
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
		enterRule(_localctx, 68, RULE_importAllClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(762);
			match(DOT);
			setState(763);
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
		enterRule(_localctx, 70, RULE_importValueAsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(765);
			match(KEYWORD_AS);
			setState(766);
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
		enterRule(_localctx, 72, RULE_importTypeClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(768);
			match(DOT);
			setState(769);
			typeName();
			setState(771);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_AS) {
				{
				setState(770);
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
		enterRule(_localctx, 74, RULE_importTypeAsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(773);
			match(KEYWORD_AS);
			setState(774);
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
		enterRule(_localctx, 76, RULE_importGroupClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(776);
			match(DOT);
			setState(777);
			match(LCURLY);
			setState(781);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(778);
				match(EOL);
				}
				}
				setState(783);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(784);
			importGroup();
			setState(788);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(785);
				match(EOL);
				}
				}
				setState(790);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(791);
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
		enterRule(_localctx, 78, RULE_importGroup);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(795);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_ENUM:
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
			case KEYWORD_REPEATE:
			case KEYWORD_STRUCT:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
			case VALUE_IDENTIFIER:
				{
				setState(793);
				importValue();
				}
				break;
			case TYPE_IDENTIFIER:
				{
				setState(794);
				importType();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(810);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(797);
					eov();
					setState(801);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(798);
						match(EOL);
						}
						}
						setState(803);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(806);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case KEYWORD_AND:
					case KEYWORD_AS:
					case KEYWORD_ATTRIBUTE:
					case KEYWORD_BREAK:
					case KEYWORD_CONST:
					case KEYWORD_ENUM:
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
					case KEYWORD_REPEATE:
					case KEYWORD_STRUCT:
					case KEYWORD_TYPE:
					case KEYWORD_XOR:
					case VALUE_IDENTIFIER:
						{
						setState(804);
						importValue();
						}
						break;
					case TYPE_IDENTIFIER:
						{
						setState(805);
						importType();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					} 
				}
				setState(812);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
			}
			setState(814);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,63,_ctx) ) {
			case 1:
				{
				setState(813);
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
		enterRule(_localctx, 80, RULE_importValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(816);
			declarationIdentifier();
			setState(818);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_AS) {
				{
				setState(817);
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
		enterRule(_localctx, 82, RULE_importType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(820);
			typeName();
			setState(822);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KEYWORD_AS) {
				{
				setState(821);
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
		enterRule(_localctx, 84, RULE_constantDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(824);
			match(KEYWORD_CONST);
			setState(825);
			patternInitializers();
			}
		}
		catch (RecognitionException re) {
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
		public List<PatternInitializerContext> patternInitializer() {
			return getRuleContexts(PatternInitializerContext.class);
		}
		public PatternInitializerContext patternInitializer(int i) {
			return getRuleContext(PatternInitializerContext.class,i);
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
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public List<DocumentedPatternInitializerContext> documentedPatternInitializer() {
			return getRuleContexts(DocumentedPatternInitializerContext.class);
		}
		public DocumentedPatternInitializerContext documentedPatternInitializer(int i) {
			return getRuleContext(DocumentedPatternInitializerContext.class,i);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
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
		enterRule(_localctx, 86, RULE_patternInitializers);
		int _la;
		try {
			int _alt;
			setState(875);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,73,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(827);
				patternInitializer();
				setState(839);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,67,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(828);
						eov();
						setState(832);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(829);
							match(EOL);
							}
							}
							setState(834);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(835);
						patternInitializer();
						}
						} 
					}
					setState(841);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,67,_ctx);
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(842);
				match(LCURLY);
				setState(846);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(843);
					match(EOL);
					}
					}
					setState(848);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(849);
				documentedPatternInitializer();
				setState(861);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,70,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(850);
						eov();
						setState(854);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(851);
							match(EOL);
							}
							}
							setState(856);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(857);
						documentedPatternInitializer();
						}
						} 
					}
					setState(863);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,70,_ctx);
				}
				setState(865);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,71,_ctx) ) {
				case 1:
					{
					setState(864);
					eov();
					}
					break;
				}
				setState(870);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(867);
					match(EOL);
					}
					}
					setState(872);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(873);
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
		enterRule(_localctx, 88, RULE_documentedPatternInitializer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(880);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LINE_DOCUMENT) {
				{
				setState(877);
				document();
				setState(878);
				match(EOL);
				}
			}

			setState(885);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(882);
				attributes();
				setState(883);
				match(EOL);
				}
			}

			setState(887);
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
		enterRule(_localctx, 90, RULE_patternInitializer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(889);
			pattern(0);
			setState(891);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(890);
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
		enterRule(_localctx, 92, RULE_initializer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(893);
			assignmentOperator();
			setState(897);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(894);
				match(EOL);
				}
				}
				setState(899);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(900);
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
		public IdentifierPatternContext identifierPattern() {
			return getRuleContext(IdentifierPatternContext.class,0);
		}
		public TerminalNode COLON_EQUAL() { return getToken(MojoParser.COLON_EQUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		enterRule(_localctx, 94, RULE_variableDeclaration);
		try {
			setState(908);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(902);
				match(KEYWORD_VAR);
				setState(903);
				patternInitializers();
				}
				break;
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_ENUM:
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
			case KEYWORD_REPEATE:
			case KEYWORD_STRUCT:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 2);
				{
				setState(904);
				identifierPattern();
				setState(905);
				match(COLON_EQUAL);
				setState(906);
				expression();
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
		enterRule(_localctx, 96, RULE_typeAliasDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(910);
			match(KEYWORD_TYPE);
			setState(911);
			typeAliasName();
			setState(913);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(912);
				genericParameterClause();
				}
			}

			setState(918);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(915);
				match(EOL);
				}
				}
				setState(920);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(921);
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
		enterRule(_localctx, 98, RULE_typeAliasName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(923);
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
		enterRule(_localctx, 100, RULE_typeAliasAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(925);
			assignmentOperator();
			setState(929);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(926);
				match(EOL);
				}
				}
				setState(931);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(932);
			type_(0);
			setState(934);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(933);
				attributes();
				}
			}

			setState(937);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,83,_ctx) ) {
			case 1:
				{
				setState(936);
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
		enterRule(_localctx, 102, RULE_functionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(939);
			match(KEYWORD_FUNC);
			setState(940);
			functionName();
			setState(942);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(941);
				genericParameterClause();
				}
			}

			setState(944);
			functionSignature();
			setState(952);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,86,_ctx) ) {
			case 1:
				{
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
		enterRule(_localctx, 104, RULE_functionName);
		try {
			setState(956);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_ENUM:
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
			case KEYWORD_REPEATE:
			case KEYWORD_STRUCT:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(954);
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
				setState(955);
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
		enterRule(_localctx, 106, RULE_functionSignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(958);
			functionParameterClause();
			setState(960);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,88,_ctx) ) {
			case 1:
				{
				setState(959);
				followingDocument();
				}
				break;
			}
			setState(969);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,90,_ctx) ) {
			case 1:
				{
				setState(965);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(962);
					match(EOL);
					}
					}
					setState(967);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(968);
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
		enterRule(_localctx, 108, RULE_functionResult);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(971);
			arrowOperator();
			setState(976);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,92,_ctx) ) {
			case 1:
				{
				setState(972);
				labelIdentifier();
				setState(974);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(973);
					match(COLON);
					}
				}

				}
				break;
			}
			setState(978);
			type_(0);
			setState(980);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(979);
				attributes();
				}
			}

			setState(989);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,95,_ctx) ) {
			case 1:
				{
				setState(985);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(982);
					match(EOL);
					}
					}
					setState(987);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(988);
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
		enterRule(_localctx, 110, RULE_functionBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(991);
			match(LCURLY);
			setState(993);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(992);
				followingDocument();
				}
			}

			setState(1002);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,98,_ctx) ) {
			case 1:
				{
				setState(998);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(995);
					match(EOL);
					}
					}
					setState(1000);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1001);
				statements();
				}
				break;
			}
			setState(1007);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1004);
				match(EOL);
				}
				}
				setState(1009);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1010);
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
		enterRule(_localctx, 112, RULE_functionParameterClause);
		int _la;
		try {
			setState(1030);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,102,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1012);
				match(LPAREN);
				setState(1013);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1014);
				match(LPAREN);
				setState(1018);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1015);
					match(EOL);
					}
					}
					setState(1020);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1021);
				functionParameters();
				setState(1025);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1022);
					match(EOL);
					}
					}
					setState(1027);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1028);
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
		enterRule(_localctx, 114, RULE_functionParameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1032);
			functionParameter();
			setState(1044);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,104,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1033);
					eovWithDocument();
					setState(1037);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1034);
						match(EOL);
						}
						}
						setState(1039);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1040);
					functionParameter();
					}
					} 
				}
				setState(1046);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,104,_ctx);
			}
			setState(1048);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,105,_ctx) ) {
			case 1:
				{
				setState(1047);
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
		enterRule(_localctx, 116, RULE_functionParameter);
		int _la;
		try {
			setState(1070);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,110,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1050);
				labelIdentifier();
				setState(1051);
				typeAnnotation();
				setState(1059);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,107,_ctx) ) {
				case 1:
					{
					setState(1055);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1052);
						match(EOL);
						}
						}
						setState(1057);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1058);
					initializer();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1061);
				labelIdentifier();
				setState(1063);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1062);
					match(COLON);
					}
				}

				setState(1065);
				type_(0);
				setState(1066);
				match(ELLIPSIS);
				setState(1068);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1067);
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
		enterRule(_localctx, 118, RULE_enumDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1072);
			match(KEYWORD_ENUM);
			setState(1073);
			enumName();
			setState(1075);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1074);
				genericParameterClause();
				}
			}

			setState(1084);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,113,_ctx) ) {
			case 1:
				{
				setState(1080);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1077);
					match(EOL);
					}
					}
					setState(1082);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1083);
				typeInheritanceClause();
				}
				break;
			}
			setState(1089);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1086);
				match(EOL);
				}
				}
				setState(1091);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1092);
			enumBody();
			}
		}
		catch (RecognitionException re) {
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
		enterRule(_localctx, 120, RULE_enumBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1094);
			match(LCURLY);
			setState(1096);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1095);
				followingDocument();
				}
			}

			setState(1105);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,117,_ctx) ) {
			case 1:
				{
				setState(1101);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1098);
					match(EOL);
					}
					}
					setState(1103);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1104);
				enumMembers();
				}
				break;
			}
			setState(1110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1107);
				match(EOL);
				}
				}
				setState(1112);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1113);
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
		enterRule(_localctx, 122, RULE_enumName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1115);
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
		enterRule(_localctx, 124, RULE_enumMembers);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1117);
			enumMember();
			setState(1129);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,120,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1118);
					eovWithDocument();
					setState(1122);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1119);
						match(EOL);
						}
						}
						setState(1124);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1125);
					enumMember();
					}
					} 
				}
				setState(1131);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,120,_ctx);
			}
			setState(1133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,121,_ctx) ) {
			case 1:
				{
				setState(1132);
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
		public FreeFloatingDocumentContext freeFloatingDocument() {
			return getRuleContext(FreeFloatingDocumentContext.class,0);
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
		enterRule(_localctx, 126, RULE_enumMember);
		int _la;
		try {
			setState(1159);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,127,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1138);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LINE_DOCUMENT) {
					{
					setState(1135);
					document();
					setState(1136);
					match(EOL);
					}
				}

				setState(1143);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1140);
					attributes();
					setState(1141);
					match(EOL);
					}
				}

				setState(1145);
				declarationIdentifier();
				setState(1147);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1146);
					attributes();
					}
				}

				setState(1156);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,126,_ctx) ) {
				case 1:
					{
					setState(1152);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1149);
						match(EOL);
						}
						}
						setState(1154);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1155);
					initializer();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1158);
				freeFloatingDocument();
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
		public TerminalNode KEYWORD_TYPE() { return getToken(MojoParser.KEYWORD_TYPE, 0); }
		public StructNameContext structName() {
			return getRuleContext(StructNameContext.class,0);
		}
		public StructTypeContext structType() {
			return getRuleContext(StructTypeContext.class,0);
		}
		public GenericParameterClauseContext genericParameterClause() {
			return getRuleContext(GenericParameterClauseContext.class,0);
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
		enterRule(_localctx, 128, RULE_structDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1161);
			match(KEYWORD_TYPE);
			setState(1162);
			structName();
			setState(1164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1163);
				genericParameterClause();
				}
			}

			setState(1166);
			structType();
			}
		}
		catch (RecognitionException re) {
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
		enterRule(_localctx, 130, RULE_structName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1168);
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
		enterRule(_localctx, 132, RULE_structType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1177);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,130,_ctx) ) {
			case 1:
				{
				setState(1173);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1170);
					match(EOL);
					}
					}
					setState(1175);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1176);
				typeInheritanceClause();
				}
				break;
			}
			setState(1186);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,132,_ctx) ) {
			case 1:
				{
				setState(1182);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1179);
					match(EOL);
					}
					}
					setState(1184);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1185);
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
		enterRule(_localctx, 134, RULE_structBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1188);
			match(LCURLY);
			setState(1192);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1189);
				followingDocument();
				setState(1190);
				match(EOL);
				}
			}

			setState(1201);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,135,_ctx) ) {
			case 1:
				{
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
				structMembers();
				}
				break;
			}
			setState(1206);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1203);
				match(EOL);
				}
				}
				setState(1208);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1209);
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
		enterRule(_localctx, 136, RULE_structMembers);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1211);
			structMember();
			setState(1223);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,138,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1212);
					eosWithDocument();
					setState(1216);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1213);
						match(EOL);
						}
						}
						setState(1218);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1219);
					structMember();
					}
					} 
				}
				setState(1225);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,138,_ctx);
			}
			setState(1227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,139,_ctx) ) {
			case 1:
				{
				setState(1226);
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
		public FreeFloatingDocumentContext freeFloatingDocument() {
			return getRuleContext(FreeFloatingDocumentContext.class,0);
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
		enterRule(_localctx, 138, RULE_structMember);
		int _la;
		try {
			setState(1247);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,143,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1232);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LINE_DOCUMENT) {
					{
					setState(1229);
					document();
					setState(1230);
					match(EOL);
					}
				}

				setState(1237);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1234);
					attributes();
					setState(1235);
					match(EOL);
					}
				}

				setState(1244);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,142,_ctx) ) {
				case 1:
					{
					setState(1239);
					structDeclaration();
					}
					break;
				case 2:
					{
					setState(1240);
					enumDeclaration();
					}
					break;
				case 3:
					{
					setState(1241);
					constantDeclaration();
					}
					break;
				case 4:
					{
					setState(1242);
					typeAliasDeclaration();
					}
					break;
				case 5:
					{
					setState(1243);
					structMemberDeclaration();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1246);
				freeFloatingDocument();
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
		enterRule(_localctx, 140, RULE_structMemberDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1249);
			declarationIdentifier();
			setState(1250);
			typeAnnotation();
			setState(1258);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,145,_ctx) ) {
			case 1:
				{
				setState(1254);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1251);
					match(EOL);
					}
					}
					setState(1256);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1257);
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
		enterRule(_localctx, 142, RULE_interfaceDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1260);
			match(KEYWORD_INTERFACE);
			setState(1261);
			interfaceName();
			setState(1263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1262);
				genericParameterClause();
				}
			}

			setState(1265);
			interfaceType();
			}
		}
		catch (RecognitionException re) {
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
		enterRule(_localctx, 144, RULE_interfaceName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1267);
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
		enterRule(_localctx, 146, RULE_interfaceType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1276);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,148,_ctx) ) {
			case 1:
				{
				setState(1272);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1269);
					match(EOL);
					}
					}
					setState(1274);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1275);
				typeInheritanceClause();
				}
				break;
			}
			setState(1281);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1278);
				match(EOL);
				}
				}
				setState(1283);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1284);
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
		enterRule(_localctx, 148, RULE_interfaceBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1286);
			match(LCURLY);
			setState(1288);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1287);
				followingDocument();
				}
			}

			setState(1297);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,152,_ctx) ) {
			case 1:
				{
				setState(1293);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1290);
					match(EOL);
					}
					}
					setState(1295);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1296);
				interfaceMembers();
				}
				break;
			}
			setState(1302);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1299);
				match(EOL);
				}
				}
				setState(1304);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1305);
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
		enterRule(_localctx, 150, RULE_interfaceMembers);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1307);
			interfaceMember();
			setState(1319);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,155,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1308);
					eosWithDocument();
					setState(1312);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1309);
						match(EOL);
						}
						}
						setState(1314);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1315);
					interfaceMember();
					}
					} 
				}
				setState(1321);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,155,_ctx);
			}
			setState(1323);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,156,_ctx) ) {
			case 1:
				{
				setState(1322);
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
		public FreeFloatingDocumentContext freeFloatingDocument() {
			return getRuleContext(FreeFloatingDocumentContext.class,0);
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
		enterRule(_localctx, 152, RULE_interfaceMember);
		int _la;
		try {
			setState(1340);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,160,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1328);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LINE_DOCUMENT) {
					{
					setState(1325);
					document();
					setState(1326);
					match(EOL);
					}
				}

				setState(1333);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AT) {
					{
					setState(1330);
					attributes();
					setState(1331);
					match(EOL);
					}
				}

				setState(1337);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,159,_ctx) ) {
				case 1:
					{
					setState(1335);
					typeAliasDeclaration();
					}
					break;
				case 2:
					{
					setState(1336);
					interfaceMethodDeclaration();
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1339);
				freeFloatingDocument();
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
		enterRule(_localctx, 154, RULE_interfaceMethodDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1342);
			functionName();
			setState(1344);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1343);
				genericParameterClause();
				}
			}

			setState(1349);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1346);
				match(EOL);
				}
				}
				setState(1351);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1352);
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
		enterRule(_localctx, 156, RULE_attributeDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1354);
			match(KEYWORD_ATTRIBUTE);
			setState(1355);
			attributeName();
			setState(1357);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1356);
				genericParameterClause();
				}
			}

			setState(1373);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,167,_ctx) ) {
			case 1:
				{
				setState(1359);
				structBody();
				}
				break;
			case 2:
				{
				setState(1360);
				typeAnnotation();
				setState(1368);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,165,_ctx) ) {
				case 1:
					{
					setState(1364);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1361);
						match(EOL);
						}
						}
						setState(1366);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1367);
					initializer();
					}
					break;
				}
				setState(1371);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(1370);
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
		enterRule(_localctx, 158, RULE_attributeAliasDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1375);
			match(KEYWORD_ATTRIBUTE);
			setState(1376);
			attributeName();
			setState(1378);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1377);
				genericParameterClause();
				}
			}

			setState(1383);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1380);
				match(EOL);
				}
				}
				setState(1385);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1386);
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
		enterRule(_localctx, 160, RULE_attributeAliasAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1388);
			assignmentOperator();
			setState(1392);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1389);
				match(EOL);
				}
				}
				setState(1394);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1398);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,171,_ctx) ) {
			case 1:
				{
				setState(1395);
				packageIdentifier();
				setState(1396);
				match(DOT);
				}
				break;
			}
			setState(1400);
			attributeName();
			setState(1402);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(1401);
				genericArgumentClause();
				}
			}

			setState(1405);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FOLLOWING_LINE_DOCUMENT) {
				{
				setState(1404);
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
		public OptionalPatternContext optionalPattern() {
			return getRuleContext(OptionalPatternContext.class,0);
		}
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
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
		int _startState = 162;
		enterRecursionRule(_localctx, 162, RULE_pattern, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1424);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,177,_ctx) ) {
			case 1:
				{
				setState(1408);
				wildcardPattern();
				setState(1410);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,174,_ctx) ) {
				case 1:
					{
					setState(1409);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 2:
				{
				setState(1412);
				identifierPattern();
				setState(1414);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,175,_ctx) ) {
				case 1:
					{
					setState(1413);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 3:
				{
				setState(1416);
				tuplePattern();
				setState(1418);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,176,_ctx) ) {
				case 1:
					{
					setState(1417);
					typeAnnotation();
					}
					break;
				}
				}
				break;
			case 4:
				{
				setState(1420);
				optionalPattern();
				}
				break;
			case 5:
				{
				setState(1421);
				match(KEYWORD_IS);
				setState(1422);
				type_(0);
				}
				break;
			case 6:
				{
				setState(1423);
				expressionPattern();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1431);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,178,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PatternContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pattern);
					setState(1426);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(1427);
					match(KEYWORD_AS);
					setState(1428);
					type_(0);
					}
					} 
				}
				setState(1433);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,178,_ctx);
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
		enterRule(_localctx, 164, RULE_wildcardPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1434);
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
		enterRule(_localctx, 166, RULE_identifierPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1436);
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
		enterRule(_localctx, 168, RULE_tuplePattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1438);
			match(LPAREN);
			setState(1440);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((_la) & ~0x3f) == 0 && ((1L << _la) & 180125310165969726L) != 0 || (((_la - 73)) & ~0x3f) == 0 && ((1L << (_la - 73)) & 4087L) != 0) {
				{
				setState(1439);
				tuplePatternElementList();
				}
			}

			setState(1442);
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
		enterRule(_localctx, 170, RULE_tuplePatternElementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1444);
			tuplePatternElement();
			setState(1449);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1445);
				match(COMMA);
				setState(1446);
				tuplePatternElement();
				}
				}
				setState(1451);
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
		enterRule(_localctx, 172, RULE_tuplePatternElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1452);
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
		enterRule(_localctx, 174, RULE_optionalPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1454);
			identifierPattern();
			setState(1455);
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
		enterRule(_localctx, 176, RULE_expressionPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1457);
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
		enterRule(_localctx, 178, RULE_attribute);
		try {
			setState(1469);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,183,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1459);
				match(AT);
				setState(1460);
				match(DECIMAL_LITERAL);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1461);
				match(AT);
				setState(1462);
				attributeIdentifier();
				setState(1464);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,181,_ctx) ) {
				case 1:
					{
					setState(1463);
					genericArgumentClause();
					}
					break;
				}
				setState(1467);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,182,_ctx) ) {
				case 1:
					{
					setState(1466);
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
		enterRule(_localctx, 180, RULE_attributeIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(1474);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,184,_ctx) ) {
			case 1:
				{
				setState(1471);
				packageIdentifier();
				setState(1472);
				match(DOT);
				}
				break;
			}
			setState(1476);
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
		enterRule(_localctx, 182, RULE_attributeName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1478);
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
		enterRule(_localctx, 184, RULE_attributeArgumentClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1480);
			match(LPAREN);
			setState(1488);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,186,_ctx) ) {
			case 1:
				{
				setState(1484);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1481);
					match(EOL);
					}
					}
					setState(1486);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1487);
				attributeArguments();
				}
				break;
			}
			setState(1493);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1490);
				match(EOL);
				}
				}
				setState(1495);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1496);
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
		enterRule(_localctx, 186, RULE_attributeArgument);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1502);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,189,_ctx) ) {
			case 1:
				{
				setState(1498);
				labelIdentifier();
				setState(1500);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1499);
					match(COLON);
					}
				}

				}
				break;
			}
			setState(1504);
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
		enterRule(_localctx, 188, RULE_attributeArguments);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1506);
			attributeArgument();
			setState(1518);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,191,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1507);
					eov();
					setState(1511);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1508);
						match(EOL);
						}
						}
						setState(1513);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1514);
					attributeArgument();
					}
					} 
				}
				setState(1520);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,191,_ctx);
			}
			setState(1522);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,192,_ctx) ) {
			case 1:
				{
				setState(1521);
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
		enterRule(_localctx, 190, RULE_attributes);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1524);
			attribute();
			setState(1531);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,194,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1526);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==EOL) {
						{
						setState(1525);
						match(EOL);
						}
					}

					setState(1528);
					attribute();
					}
					} 
				}
				setState(1533);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,194,_ctx);
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
		enterRule(_localctx, 192, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1534);
			prefixExpression();
			setState(1536);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,195,_ctx) ) {
			case 1:
				{
				setState(1535);
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
	public static class ExpressionsContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
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
		public ExpressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressions; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitExpressions(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionsContext expressions() throws RecognitionException {
		ExpressionsContext _localctx = new ExpressionsContext(_ctx, getState());
		enterRule(_localctx, 194, RULE_expressions);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1538);
			expression();
			setState(1550);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,197,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1539);
					eov();
					setState(1543);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1540);
						match(EOL);
						}
						}
						setState(1545);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1546);
					expression();
					}
					} 
				}
				setState(1552);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,197,_ctx);
			}
			setState(1554);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA || _la==EOL) {
				{
				setState(1553);
				eov();
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
		enterRule(_localctx, 196, RULE_prefixExpression);
		try {
			setState(1560);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,199,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1556);
				prefixOperator();
				setState(1557);
				postfixExpression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1559);
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
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public ConditionalOperatorContext conditionalOperator() {
			return getRuleContext(ConditionalOperatorContext.class,0);
		}
		public TypeCastingOperatorContext typeCastingOperator() {
			return getRuleContext(TypeCastingOperatorContext.class,0);
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
		enterRule(_localctx, 198, RULE_binaryExpression);
		try {
			setState(1572);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,200,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1562);
				binaryOperator();
				setState(1563);
				prefixExpression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1565);
				assignmentOperator();
				setState(1566);
				prefixExpression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1568);
				conditionalOperator();
				setState(1569);
				prefixExpression();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1571);
				typeCastingOperator();
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
		enterRule(_localctx, 200, RULE_binaryExpressions);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1575); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(1574);
					binaryExpression();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1577); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,201,_ctx);
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
		enterRule(_localctx, 202, RULE_conditionalOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1579);
			match(QUESTION);
			setState(1580);
			expression();
			setState(1581);
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
	public static class TypeCastingOperatorContext extends ParserRuleContext {
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
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
		enterRule(_localctx, 204, RULE_typeCastingOperator);
		try {
			setState(1587);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_IS:
				enterOuterAlt(_localctx, 1);
				{
				setState(1583);
				match(KEYWORD_IS);
				setState(1584);
				type_(0);
				}
				break;
			case KEYWORD_AS:
				enterOuterAlt(_localctx, 2);
				{
				setState(1585);
				match(KEYWORD_AS);
				setState(1586);
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
		public ParenthesizedExpressionContext parenthesizedExpression() {
			return getRuleContext(ParenthesizedExpressionContext.class,0);
		}
		public TupleExpressionContext tupleExpression() {
			return getRuleContext(TupleExpressionContext.class,0);
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
		enterRule(_localctx, 206, RULE_primaryExpression);
		try {
			setState(1606);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,205,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1589);
				literalExpression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1590);
				declarationIdentifier();
				setState(1592);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,203,_ctx) ) {
				case 1:
					{
					setState(1591);
					genericArgumentClause();
					}
					break;
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1594);
				typeIdentifier();
				setState(1595);
				match(DOT);
				setState(1596);
				declarationIdentifier();
				setState(1598);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,204,_ctx) ) {
				case 1:
					{
					setState(1597);
					genericArgumentClause();
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1600);
				closureExpression();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1601);
				parenthesizedExpression();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1602);
				tupleExpression();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1603);
				implicitMemberExpression();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(1604);
				wildcardExpression();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(1605);
				structConstructionExpression();
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
		enterRule(_localctx, 208, RULE_literalExpression);
		try {
			setState(1615);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,206,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1608);
				numericOperatorLiteral();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1609);
				stringOperatorLiteral();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1610);
				structLiteral();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1611);
				literal();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(1612);
				arrayLiteral();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(1613);
				mapLiteral();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(1614);
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
		enterRule(_localctx, 210, RULE_numericOperatorLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1617);
			numericLiteral();
			setState(1618);
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
		enterRule(_localctx, 212, RULE_stringOperatorLiteral);
		try {
			setState(1629);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(1620);
				prefixLiteralOperator();
				setState(1621);
				if (!(_input.get(_input.index()-1).getType()!=WS)) throw new FailedPredicateException(this, "_input.get(_input.index()-1).getType()!=WS");
				setState(1622);
				stringLiteral();
				setState(1624);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,207,_ctx) ) {
				case 1:
					{
					setState(1623);
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
				setState(1626);
				stringLiteral();
				setState(1627);
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
		enterRule(_localctx, 214, RULE_suffixLiteralOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1631);
			if (!(_input.get(_input.index()-1).getType()!=WS)) throw new FailedPredicateException(this, "_input.get(_input.index()-1).getType()!=WS");
			setState(1632);
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
		enterRule(_localctx, 216, RULE_prefixLiteralOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1634);
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
		enterRule(_localctx, 218, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1636);
			match(LBRACK);
			setState(1644);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,210,_ctx) ) {
			case 1:
				{
				setState(1640);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1637);
					match(EOL);
					}
					}
					setState(1642);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1643);
				arrayLiteralItems();
				}
				break;
			}
			setState(1649);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1646);
				match(EOL);
				}
				}
				setState(1651);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1652);
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
		enterRule(_localctx, 220, RULE_arrayLiteralItems);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1654);
			arrayLiteralItem();
			setState(1666);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,213,_ctx);
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
					arrayLiteralItem();
					}
					} 
				}
				setState(1668);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,213,_ctx);
			}
			setState(1670);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,214,_ctx) ) {
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
	public static class ArrayLiteralItemContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
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
		enterRule(_localctx, 222, RULE_arrayLiteralItem);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1672);
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
		enterRule(_localctx, 224, RULE_mapLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1674);
			match(LCURLY);
			setState(1682);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,216,_ctx) ) {
			case 1:
				{
				setState(1678);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1675);
					match(EOL);
					}
					}
					setState(1680);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1681);
				mapLiteralItems();
				}
				break;
			}
			setState(1687);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1684);
				match(EOL);
				}
				}
				setState(1689);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1690);
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
		enterRule(_localctx, 226, RULE_mapLiteralItems);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(1692);
			mapLiteralItem();
			setState(1704);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,219,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1693);
					eov();
					setState(1697);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1694);
						match(EOL);
						}
						}
						setState(1699);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1700);
					mapLiteralItem();
					}
					} 
				}
				setState(1706);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,219,_ctx);
			}
			setState(1708);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,220,_ctx) ) {
			case 1:
				{
				setState(1707);
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
		enterRule(_localctx, 228, RULE_mapLiteralItem);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1712);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STATIC_STRING_LITERAL:
			case INTERPOLATED_STRING_LITERAL:
				{
				setState(1710);
				stringLiteral();
				}
				break;
			case BINARY_LITERAL:
			case OCTAL_LITERAL:
			case DECIMAL_LITERAL:
			case PURE_DECIMAL_DIGITS:
			case HEXADECIMAL_LITERAL:
				{
				setState(1711);
				integerLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(1714);
			match(COLON);
			setState(1715);
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
		enterRule(_localctx, 230, RULE_objectLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1717);
			match(LCURLY);
			setState(1725);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,223,_ctx) ) {
			case 1:
				{
				setState(1721);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1718);
					match(EOL);
					}
					}
					setState(1723);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1724);
				objectLiteralItems();
				}
				break;
			}
			setState(1730);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1727);
				match(EOL);
				}
				}
				setState(1732);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1733);
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
		enterRule(_localctx, 232, RULE_objectLiteralItems);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1735);
			objectLiteralItem();
			setState(1747);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,226,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1736);
					eov();
					setState(1740);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1737);
						match(EOL);
						}
						}
						setState(1742);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1743);
					objectLiteralItem();
					}
					} 
				}
				setState(1749);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,226,_ctx);
			}
			setState(1751);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,227,_ctx) ) {
			case 1:
				{
				setState(1750);
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
		enterRule(_localctx, 234, RULE_objectLiteralItem);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1753);
			pathIdentifier();
			setState(1756);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(1754);
				match(COLON);
				setState(1755);
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
		enterRule(_localctx, 236, RULE_structLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1758);
			typeIdentifier();
			setState(1759);
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
		enterRule(_localctx, 238, RULE_structConstructionExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1761);
			typeIdentifier();
			setState(1762);
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
	public static class ClosureExpressionContext extends ParserRuleContext {
		public TerminalNode LCURLY() { return getToken(MojoParser.LCURLY, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public ClosureParametersContext closureParameters() {
			return getRuleContext(ClosureParametersContext.class,0);
		}
		public TerminalNode RIGHT_ARROW() { return getToken(MojoParser.RIGHT_ARROW, 0); }
		public List<TerminalNode> EOL() { return getTokens(MojoParser.EOL); }
		public TerminalNode EOL(int i) {
			return getToken(MojoParser.EOL, i);
		}
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
		enterRule(_localctx, 240, RULE_closureExpression);
		int _la;
		try {
			setState(1795);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,233,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1764);
				match(LCURLY);
				setState(1765);
				statements();
				setState(1766);
				match(RCURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1768);
				match(LCURLY);
				setState(1769);
				closureParameters();
				setState(1773);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1770);
					match(EOL);
					}
					}
					setState(1775);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1776);
				match(RIGHT_ARROW);
				setState(1784);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,231,_ctx) ) {
				case 1:
					{
					setState(1780);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1777);
						match(EOL);
						}
						}
						setState(1782);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1783);
					type_(0);
					}
					break;
				}
				setState(1789);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(1786);
					match(EOL);
					}
					}
					setState(1791);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(1792);
				statements();
				setState(1793);
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
		enterRule(_localctx, 242, RULE_closureParameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1797);
			closureParameter();
			setState(1809);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,235,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1798);
					eov();
					setState(1802);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(1799);
						match(EOL);
						}
						}
						setState(1804);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(1805);
					closureParameter();
					}
					} 
				}
				setState(1811);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,235,_ctx);
			}
			setState(1813);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,236,_ctx) ) {
			case 1:
				{
				setState(1812);
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
		enterRule(_localctx, 244, RULE_closureParameter);
		try {
			setState(1817);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,237,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1815);
				functionParameter();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1816);
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
		enterRule(_localctx, 246, RULE_implicitMemberExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1819);
			match(DOT);
			setState(1820);
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
		enterRule(_localctx, 248, RULE_parenthesizedExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1822);
			match(LPAREN);
			setState(1826);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1823);
				match(EOL);
				}
				}
				setState(1828);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			{
			setState(1829);
			expression();
			}
			setState(1833);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(1830);
				match(EOL);
				}
				}
				setState(1835);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(1836);
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
	public static class TupleExpressionContext extends ParserRuleContext {
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
		public TupleExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleExpression; }
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MojoParserVisitor ) return ((MojoParserVisitor<? extends T>)visitor).visitTupleExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TupleExpressionContext tupleExpression() throws RecognitionException {
		TupleExpressionContext _localctx = new TupleExpressionContext(_ctx, getState());
		enterRule(_localctx, 250, RULE_tupleExpression);
		int _la;
		try {
			setState(1850);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,241,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1838);
				match(LPAREN);
				setState(1839);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1840);
				match(LPAREN);
				setState(1841);
				tupleElement();
				setState(1844); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(1842);
					match(COMMA);
					setState(1843);
					tupleElement();
					}
					}
					setState(1846); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				setState(1848);
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
		enterRule(_localctx, 252, RULE_tupleElement);
		int _la;
		try {
			setState(1859);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,243,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1852);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1853);
				labelIdentifier();
				setState(1855);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1854);
					match(COLON);
					}
				}

				setState(1857);
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
		enterRule(_localctx, 254, RULE_wildcardExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1861);
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
		enterRule(_localctx, 256, RULE_postfixExpression);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1863);
			primaryExpression();
			setState(1867);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,244,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1864);
					suffixExpression();
					}
					} 
				}
				setState(1869);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,244,_ctx);
			}
			setState(1871);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,245,_ctx) ) {
			case 1:
				{
				setState(1870);
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
		enterRule(_localctx, 258, RULE_suffixExpression);
		try {
			setState(1876);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LCURLY:
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(1873);
				functionCallSuffix();
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(1874);
				explicitMemberSuffix();
				}
				break;
			case LBRACK:
				enterOuterAlt(_localctx, 3);
				{
				setState(1875);
				subscriptSuffix();
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
		public TerminalNode PURE_DECIMAL_DIGITS() { return getToken(MojoParser.PURE_DECIMAL_DIGITS, 0); }
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
		enterRule(_localctx, 260, RULE_explicitMemberSuffix);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1878);
			match(DOT);
			setState(1888);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PURE_DECIMAL_DIGITS:
				{
				setState(1879);
				match(PURE_DECIMAL_DIGITS);
				}
				break;
			case VALUE_IDENTIFIER:
			case IMPLICIT_PARAMETER_NAME:
				{
				setState(1880);
				identifier();
				setState(1886);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,247,_ctx) ) {
				case 1:
					{
					setState(1881);
					genericArgumentClause();
					}
					break;
				case 2:
					{
					setState(1882);
					match(LPAREN);
					setState(1883);
					argumentNames();
					setState(1884);
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
		enterRule(_localctx, 262, RULE_subscriptSuffix);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1890);
			match(LBRACK);
			setState(1891);
			functionCallArguments();
			setState(1892);
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
		enterRule(_localctx, 264, RULE_functionCallSuffix);
		int _la;
		try {
			setState(1899);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,250,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1895);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LPAREN) {
					{
					setState(1894);
					functionCallArgumentClause();
					}
				}

				setState(1897);
				trailingClosures();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1898);
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
		enterRule(_localctx, 266, RULE_functionCallArgumentClause);
		try {
			setState(1907);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,251,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1901);
				match(LPAREN);
				setState(1902);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1903);
				match(LPAREN);
				setState(1904);
				functionCallArguments();
				setState(1905);
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
		enterRule(_localctx, 268, RULE_functionCallArguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1909);
			functionCallArgument();
			setState(1914);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1910);
				match(COMMA);
				setState(1911);
				functionCallArgument();
				}
				}
				setState(1916);
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
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
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
		enterRule(_localctx, 270, RULE_functionCallArgument);
		int _la;
		try {
			setState(1931);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,255,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1917);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1918);
				labelIdentifier();
				setState(1920);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1919);
					match(COLON);
					}
				}

				setState(1922);
				expression();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1924);
				operator();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1925);
				labelIdentifier();
				setState(1927);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(1926);
					match(COLON);
					}
				}

				setState(1929);
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
		enterRule(_localctx, 272, RULE_trailingClosures);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1933);
			closureExpression();
			setState(1935);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,256,_ctx) ) {
			case 1:
				{
				setState(1934);
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
		enterRule(_localctx, 274, RULE_labeledTrailingClosures);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1938); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(1937);
					labeledTrailingClosure();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1940); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,257,_ctx);
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
		enterRule(_localctx, 276, RULE_labeledTrailingClosure);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1942);
			identifier();
			setState(1943);
			match(COLON);
			setState(1944);
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
		enterRule(_localctx, 278, RULE_argumentNames);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1946);
			argumentName();
			setState(1950);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((_la) & ~0x3f) == 0 && ((1L << _la) & 1073741822L) != 0 || _la==VALUE_IDENTIFIER) {
				{
				{
				setState(1947);
				argumentName();
				}
				}
				setState(1952);
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
		enterRule(_localctx, 280, RULE_argumentName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1953);
			labelIdentifier();
			setState(1954);
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
		int _startState = 282;
		enterRecursionRule(_localctx, 282, RULE_type_, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1959);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,259,_ctx) ) {
			case 1:
				{
				setState(1957);
				basicType(0);
				}
				break;
			case 2:
				{
				setState(1958);
				functionType();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1969);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,261,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(1967);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,260,_ctx) ) {
					case 1:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(1961);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(1962);
						match(BANG);
						}
						break;
					case 2:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(1963);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(1964);
						match(QUESTION);
						}
						break;
					case 3:
						{
						_localctx = new Type_Context(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type_);
						setState(1965);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(1966);
						match(ELLIPSIS);
						}
						break;
					}
					} 
				}
				setState(1971);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,261,_ctx);
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
		int _startState = 284;
		enterRecursionRule(_localctx, 284, RULE_basicType, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new PrimeContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(1973);
			primeType();
			}
			_ctx.stop = _input.LT(-1);
			setState(2039);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,275,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(2037);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,274,_ctx) ) {
					case 1:
						{
						_localctx = new UnionContext(new BasicTypeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_basicType);
						setState(1975);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(1977);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==AT) {
							{
							setState(1976);
							attributes();
							}
						}

						setState(1982);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==FOLLOWING_LINE_DOCUMENT) {
							{
							setState(1979);
							followingDocument();
							setState(1980);
							match(EOL);
							}
						}

						setState(1987);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(1984);
							match(EOL);
							}
							}
							setState(1989);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(1990);
						match(PIPE);
						setState(1994);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(1991);
							match(EOL);
							}
							}
							setState(1996);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(1997);
						basicType(0);
						setState(1999);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,266,_ctx) ) {
						case 1:
							{
							setState(1998);
							attributes();
							}
							break;
						}
						setState(2004);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,267,_ctx) ) {
						case 1:
							{
							setState(2001);
							followingDocument();
							setState(2002);
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
						setState(2006);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(2008);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==AT) {
							{
							setState(2007);
							attributes();
							}
						}

						setState(2013);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==FOLLOWING_LINE_DOCUMENT) {
							{
							setState(2010);
							followingDocument();
							setState(2011);
							match(EOL);
							}
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
						match(AND);
						setState(2025);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==EOL) {
							{
							{
							setState(2022);
							match(EOL);
							}
							}
							setState(2027);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(2028);
						basicType(0);
						setState(2030);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,272,_ctx) ) {
						case 1:
							{
							setState(2029);
							attributes();
							}
							break;
						}
						setState(2035);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,273,_ctx) ) {
						case 1:
							{
							setState(2032);
							followingDocument();
							setState(2033);
							match(EOL);
							}
							break;
						}
						}
						break;
					}
					} 
				}
				setState(2041);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,275,_ctx);
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
		enterRule(_localctx, 286, RULE_primeType);
		try {
			setState(2046);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LBRACK:
				enterOuterAlt(_localctx, 1);
				{
				setState(2042);
				arrayType();
				}
				break;
			case LCURLY:
				enterOuterAlt(_localctx, 2);
				{
				setState(2043);
				mapType();
				}
				break;
			case LPAREN:
				enterOuterAlt(_localctx, 3);
				{
				setState(2044);
				tupleType();
				}
				break;
			case TYPE_IDENTIFIER:
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 4);
				{
				setState(2045);
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
		enterRule(_localctx, 288, RULE_typeAnnotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2049);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(2048);
				match(COLON);
				}
			}

			setState(2051);
			type_(0);
			setState(2053);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,278,_ctx) ) {
			case 1:
				{
				setState(2052);
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
		enterRule(_localctx, 290, RULE_typeIdentifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2058);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VALUE_IDENTIFIER) {
				{
				setState(2055);
				packageIdentifier();
				setState(2056);
				match(DOT);
				}
			}

			setState(2060);
			typeIdentifierClause();
			setState(2065);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,280,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2061);
					match(DOT);
					setState(2062);
					typeIdentifierClause();
					}
					} 
				}
				setState(2067);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,280,_ctx);
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
		enterRule(_localctx, 292, RULE_typeIdentifierClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2068);
			typeName();
			setState(2070);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,281,_ctx) ) {
			case 1:
				{
				setState(2069);
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
		enterRule(_localctx, 294, RULE_typeName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2072);
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
		enterRule(_localctx, 296, RULE_tupleType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2074);
			match(LPAREN);
			setState(2082);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,283,_ctx) ) {
			case 1:
				{
				setState(2078);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==EOL) {
					{
					{
					setState(2075);
					match(EOL);
					}
					}
					setState(2080);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(2081);
				tupleTypeElements();
				}
				break;
			}
			setState(2087);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(2084);
				match(EOL);
				}
				}
				setState(2089);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(2090);
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
		enterRule(_localctx, 298, RULE_tupleTypeElements);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2092);
			tupleTypeElement();
			setState(2104);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,286,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2093);
					eovWithDocument();
					setState(2097);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(2094);
						match(EOL);
						}
						}
						setState(2099);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(2100);
					tupleTypeElement();
					}
					} 
				}
				setState(2106);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,286,_ctx);
			}
			setState(2108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,287,_ctx) ) {
			case 1:
				{
				setState(2107);
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
		enterRule(_localctx, 300, RULE_tupleTypeElement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2114);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,289,_ctx) ) {
			case 1:
				{
				setState(2110);
				declarationIdentifier();
				setState(2112);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(2111);
					match(COLON);
					}
				}

				}
				break;
			}
			setState(2116);
			type_(0);
			setState(2118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2117);
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
		enterRule(_localctx, 302, RULE_functionType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2120);
			functionParameterClause();
			setState(2121);
			arrowOperator();
			setState(2122);
			type_(0);
			setState(2124);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,291,_ctx) ) {
			case 1:
				{
				setState(2123);
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
		enterRule(_localctx, 304, RULE_arrayType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2126);
			match(LBRACK);
			setState(2127);
			type_(0);
			setState(2129);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2128);
				attributes();
				}
			}

			setState(2131);
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
		public TerminalNode COLON() { return getToken(MojoParser.COLON, 0); }
		public TerminalNode RCURLY() { return getToken(MojoParser.RCURLY, 0); }
		public KeyAttributesContext keyAttributes() {
			return getRuleContext(KeyAttributesContext.class,0);
		}
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
		enterRule(_localctx, 306, RULE_mapType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2133);
			match(LCURLY);
			setState(2134);
			type_(0);
			setState(2136);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2135);
				keyAttributes();
				}
			}

			setState(2138);
			match(COLON);
			setState(2139);
			type_(0);
			setState(2141);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2140);
				attributes();
				}
			}

			setState(2143);
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
		enterRule(_localctx, 308, RULE_keyAttributes);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2145);
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
		enterRule(_localctx, 310, RULE_typeInheritanceClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2147);
			match(COLON);
			setState(2151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EOL) {
				{
				{
				setState(2148);
				match(EOL);
				}
				}
				setState(2153);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(2154);
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
		enterRule(_localctx, 312, RULE_typeInheritances);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2156);
			typeInheritance();
			setState(2168);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,297,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2157);
					eovWithDocument();
					setState(2161);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==EOL) {
						{
						{
						setState(2158);
						match(EOL);
						}
						}
						setState(2163);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(2164);
					typeInheritance();
					}
					} 
				}
				setState(2170);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,297,_ctx);
			}
			setState(2172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,298,_ctx) ) {
			case 1:
				{
				setState(2171);
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
		enterRule(_localctx, 314, RULE_typeInheritance);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2174);
			basicType(0);
			setState(2176);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(2175);
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
		enterRule(_localctx, 316, RULE_declarationIdentifier);
		try {
			setState(2180);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2178);
				match(VALUE_IDENTIFIER);
				}
				break;
			case KEYWORD_AND:
			case KEYWORD_AS:
			case KEYWORD_ATTRIBUTE:
			case KEYWORD_BREAK:
			case KEYWORD_CONST:
			case KEYWORD_ENUM:
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
			case KEYWORD_REPEATE:
			case KEYWORD_STRUCT:
			case KEYWORD_TYPE:
			case KEYWORD_XOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(2179);
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
		enterRule(_localctx, 318, RULE_labelIdentifier);
		try {
			setState(2184);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VALUE_IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2182);
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
			case KEYWORD_IF:
			case KEYWORD_IMPORT:
			case KEYWORD_IN:
			case KEYWORD_INTERFACE:
			case KEYWORD_IS:
			case KEYWORD_MATCH:
			case KEYWORD_NOT:
			case KEYWORD_NULL:
			case KEYWORD_OR:
			case KEYWORD_PACKAGE:
			case KEYWORD_REPEATE:
			case KEYWORD_RETURN:
			case KEYWORD_STRUCT:
			case KEYWORD_TRUE:
			case KEYWORD_TYPE:
			case KEYWORD_VAR:
			case KEYWORD_WHILE:
			case KEYWORD_XOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(2183);
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
		enterRule(_localctx, 320, RULE_pathIdentifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2186);
			declarationIdentifier();
			setState(2191);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(2187);
				match(DOT);
				setState(2188);
				declarationIdentifier();
				}
				}
				setState(2193);
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
		enterRule(_localctx, 322, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2194);
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
		public TerminalNode KEYWORD_ENUM() { return getToken(MojoParser.KEYWORD_ENUM, 0); }
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
		public TerminalNode KEYWORD_REPEATE() { return getToken(MojoParser.KEYWORD_REPEATE, 0); }
		public TerminalNode KEYWORD_STRUCT() { return getToken(MojoParser.KEYWORD_STRUCT, 0); }
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
		enterRule(_localctx, 324, RULE_keywordAsIdentifierInDeclarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2196);
			_la = _input.LA(1);
			if ( !(((_la) & ~0x3f) == 0 && ((1L << _la) & 629139774L) != 0) ) {
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
		public TerminalNode KEYWORD_IF() { return getToken(MojoParser.KEYWORD_IF, 0); }
		public TerminalNode KEYWORD_IMPORT() { return getToken(MojoParser.KEYWORD_IMPORT, 0); }
		public TerminalNode KEYWORD_IN() { return getToken(MojoParser.KEYWORD_IN, 0); }
		public TerminalNode KEYWORD_INTERFACE() { return getToken(MojoParser.KEYWORD_INTERFACE, 0); }
		public TerminalNode KEYWORD_IS() { return getToken(MojoParser.KEYWORD_IS, 0); }
		public TerminalNode KEYWORD_MATCH() { return getToken(MojoParser.KEYWORD_MATCH, 0); }
		public TerminalNode KEYWORD_NOT() { return getToken(MojoParser.KEYWORD_NOT, 0); }
		public TerminalNode KEYWORD_NULL() { return getToken(MojoParser.KEYWORD_NULL, 0); }
		public TerminalNode KEYWORD_OR() { return getToken(MojoParser.KEYWORD_OR, 0); }
		public TerminalNode KEYWORD_PACKAGE() { return getToken(MojoParser.KEYWORD_PACKAGE, 0); }
		public TerminalNode KEYWORD_REPEATE() { return getToken(MojoParser.KEYWORD_REPEATE, 0); }
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
		enterRule(_localctx, 326, RULE_keywordAsIdentifierInLabels);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2198);
			_la = _input.LA(1);
			if ( !(((_la) & ~0x3f) == 0 && ((1L << _la) & 1073741822L) != 0) ) {
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
		enterRule(_localctx, 328, RULE_document);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2200);
			match(LINE_DOCUMENT);
			setState(2205);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,303,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2201);
					match(EOL);
					setState(2202);
					match(LINE_DOCUMENT);
					}
					} 
				}
				setState(2207);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,303,_ctx);
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
		enterRule(_localctx, 330, RULE_followingDocument);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2208);
			match(FOLLOWING_LINE_DOCUMENT);
			setState(2213);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,304,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(2209);
					match(EOL);
					setState(2210);
					match(FOLLOWING_LINE_DOCUMENT);
					}
					} 
				}
				setState(2215);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,304,_ctx);
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
		enterRule(_localctx, 332, RULE_assignmentOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2216);
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
		enterRule(_localctx, 334, RULE_negatePrefixOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2218);
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
		enterRule(_localctx, 336, RULE_arrowOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2220);
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
		enterRule(_localctx, 338, RULE_rangeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2222);
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
		enterRule(_localctx, 340, RULE_halfOpenRangeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2224);
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
	public static class BinaryOperatorContext extends ParserRuleContext {
		public RangeOperatorContext rangeOperator() {
			return getRuleContext(RangeOperatorContext.class,0);
		}
		public HalfOpenRangeOperatorContext halfOpenRangeOperator() {
			return getRuleContext(HalfOpenRangeOperatorContext.class,0);
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
		enterRule(_localctx, 342, RULE_binaryOperator);
		try {
			setState(2231);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT_DOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(2226);
				rangeOperator();
				}
				break;
			case DOT_DOT_LT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2227);
				halfOpenRangeOperator();
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
				enterOuterAlt(_localctx, 3);
				{
				setState(2228);
				operator();
				}
				break;
			case KEYWORD_AND:
				enterOuterAlt(_localctx, 4);
				{
				setState(2229);
				match(KEYWORD_AND);
				}
				break;
			case KEYWORD_OR:
				enterOuterAlt(_localctx, 5);
				{
				setState(2230);
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
		enterRule(_localctx, 344, RULE_prefixOperator);
		try {
			setState(2235);
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
				setState(2233);
				operator();
				}
				break;
			case KEYWORD_NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2234);
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
		enterRule(_localctx, 346, RULE_postfixOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2237);
			_la = _input.LA(1);
			if ( !(_la==PLUS_PLUS || _la==MINUS_MINUS) ) {
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
		enterRule(_localctx, 348, RULE_operator);
		try {
			int _alt;
			setState(2250);
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
				setState(2239);
				operator_head();
				setState(2241);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,307,_ctx) ) {
				case 1:
					{
					setState(2240);
					operator_characters();
					}
					break;
				}
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2243);
				dot_operator_head();
				setState(2247);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,308,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(2244);
						dot_operator_character();
						}
						} 
					}
					setState(2249);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,308,_ctx);
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
		enterRule(_localctx, 350, RULE_operator_characters);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(2254); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(2252);
					if (!(_input.get(_input.index()-1).getType()!=WS)) throw new FailedPredicateException(this, "_input.get(_input.index()-1).getType()!=WS");
					setState(2253);
					operator_character();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(2256); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,310,_ctx);
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
		enterRule(_localctx, 352, RULE_operator_character);
		try {
			setState(2260);
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
				setState(2258);
				operator_head();
				}
				break;
			case OPERATOR_FOLLOWING_CHARACTER:
				enterOuterAlt(_localctx, 2);
				{
				setState(2259);
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
		enterRule(_localctx, 354, RULE_operator_head);
		int _la;
		try {
			setState(2264);
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
				setState(2262);
				_la = _input.LA(1);
				if ( !(((_la) & ~0x3f) == 0 && ((1L << _la) & 36010105321291776L) != 0) ) {
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
				setState(2263);
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
		enterRule(_localctx, 356, RULE_dot_operator_head);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2266);
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
		enterRule(_localctx, 358, RULE_dot_operator_character);
		try {
			setState(2270);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(2268);
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
				setState(2269);
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
		enterRule(_localctx, 360, RULE_literal);
		try {
			setState(2276);
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
				setState(2272);
				numericLiteral();
				}
				break;
			case STATIC_STRING_LITERAL:
			case INTERPOLATED_STRING_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(2273);
				stringLiteral();
				}
				break;
			case KEYWORD_FALSE:
			case KEYWORD_TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(2274);
				boolLiteral();
				}
				break;
			case KEYWORD_NULL:
				enterOuterAlt(_localctx, 4);
				{
				setState(2275);
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
		enterRule(_localctx, 362, RULE_boolLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2278);
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
		enterRule(_localctx, 364, RULE_nullLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2280);
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
		enterRule(_localctx, 366, RULE_numericLiteral);
		int _la;
		try {
			setState(2290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,317,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(2283);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS) {
					{
					setState(2282);
					negatePrefixOperator();
					}
				}

				setState(2285);
				integerLiteral();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(2287);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS) {
					{
					setState(2286);
					negatePrefixOperator();
					}
				}

				setState(2289);
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
		enterRule(_localctx, 368, RULE_integerLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2292);
			_la = _input.LA(1);
			if ( !((((_la - 77)) & ~0x3f) == 0 && ((1L << (_la - 77)) & 31L) != 0) ) {
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
		enterRule(_localctx, 370, RULE_stringLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2294);
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
		enterRule(_localctx, 372, RULE_eos);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2296);
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
		enterRule(_localctx, 374, RULE_eov);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(2298);
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
		enterRule(_localctx, 376, RULE_eosWithDocument);
		int _la;
		try {
			setState(2310);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SEMI:
				enterOuterAlt(_localctx, 1);
				{
				setState(2300);
				match(SEMI);
				setState(2304);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(2301);
					followingDocument();
					setState(2302);
					match(EOL);
					}
				}

				}
				break;
			case EOL:
			case FOLLOWING_LINE_DOCUMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(2307);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(2306);
					followingDocument();
					}
				}

				setState(2309);
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
		enterRule(_localctx, 378, RULE_eovWithDocument);
		int _la;
		try {
			setState(2322);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(2312);
				match(COMMA);
				setState(2316);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,321,_ctx) ) {
				case 1:
					{
					setState(2313);
					followingDocument();
					setState(2314);
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
				setState(2319);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==FOLLOWING_LINE_DOCUMENT) {
					{
					setState(2318);
					followingDocument();
					}
				}

				setState(2321);
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
		case 81:
			return pattern_sempred((PatternContext)_localctx, predIndex);
		case 106:
			return stringOperatorLiteral_sempred((StringOperatorLiteralContext)_localctx, predIndex);
		case 107:
			return suffixLiteralOperator_sempred((SuffixLiteralOperatorContext)_localctx, predIndex);
		case 141:
			return type__sempred((Type_Context)_localctx, predIndex);
		case 142:
			return basicType_sempred((BasicTypeContext)_localctx, predIndex);
		case 175:
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
			return _input.get(_input.index()-1).getType()!=WS;
		}
		return true;
	}
	private boolean suffixLiteralOperator_sempred(SuffixLiteralOperatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return _input.get(_input.index()-1).getType()!=WS;
		}
		return true;
	}
	private boolean type__sempred(Type_Context _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean basicType_sempred(BasicTypeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 3);
		case 7:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean operator_characters_sempred(Operator_charactersContext _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return _input.get(_input.index()-1).getType()!=WS;
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001]\u0915\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"\u0002\u00bd\u0007\u00bd\u0001\u0000\u0005\u0000\u017e\b\u0000\n\u0000"+
		"\f\u0000\u0181\t\u0000\u0001\u0000\u0003\u0000\u0184\b\u0000\u0001\u0000"+
		"\u0005\u0000\u0187\b\u0000\n\u0000\f\u0000\u018a\t\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0003\u0001\u0194\b\u0001\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0005\u0003\u019b\b\u0003\n\u0003\f\u0003\u019e\t\u0003"+
		"\u0001\u0003\u0001\u0003\u0005\u0003\u01a2\b\u0003\n\u0003\f\u0003\u01a5"+
		"\t\u0003\u0001\u0003\u0003\u0003\u01a8\b\u0003\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u01ac\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005"+
		"\u01b1\b\u0005\n\u0005\f\u0005\u01b4\t\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0005\u0005\u01b9\b\u0005\n\u0005\f\u0005\u01bc\t\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006\u01c3\b\u0006"+
		"\n\u0006\f\u0006\u01c6\t\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0005\u0007\u01cd\b\u0007\n\u0007\f\u0007\u01d0\t\u0007"+
		"\u0001\u0007\u0001\u0007\u0005\u0007\u01d4\b\u0007\n\u0007\f\u0007\u01d7"+
		"\t\u0007\u0001\b\u0001\b\u0003\b\u01db\b\b\u0001\t\u0001\t\u0001\t\u0005"+
		"\t\u01e0\b\t\n\t\f\t\u01e3\t\t\u0001\t\u0001\t\u0001\n\u0001\n\u0003\n"+
		"\u01e9\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u01ee\b\u000b"+
		"\n\u000b\f\u000b\u01f1\t\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u01f5"+
		"\b\u000b\n\u000b\f\u000b\u01f8\t\u000b\u0001\u000b\u0003\u000b\u01fb\b"+
		"\u000b\u0001\f\u0001\f\u0005\f\u01ff\b\f\n\f\f\f\u0202\t\f\u0001\f\u0001"+
		"\f\u0001\f\u0005\f\u0207\b\f\n\f\f\f\u020a\t\f\u0001\f\u0003\f\u020d\b"+
		"\f\u0001\r\u0001\r\u0001\r\u0005\r\u0212\b\r\n\r\f\r\u0215\t\r\u0001\r"+
		"\u0001\r\u0005\r\u0219\b\r\n\r\f\r\u021c\t\r\u0001\r\u0003\r\u021f\b\r"+
		"\u0001\r\u0005\r\u0222\b\r\n\r\f\r\u0225\t\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0005\u000e\u022c\b\u000e\n\u000e\f\u000e\u022f"+
		"\t\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u0233\b\u000e\n\u000e\f\u000e"+
		"\u0236\t\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u023c\b\u000f\n\u000f\f\u000f\u023f\t\u000f\u0001\u000f\u0001\u000f\u0005"+
		"\u000f\u0243\b\u000f\n\u000f\f\u000f\u0246\t\u000f\u0001\u000f\u0001\u000f"+
		"\u0003\u000f\u024a\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010"+
		"\u024f\b\u0010\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0013"+
		"\u0001\u0013\u0003\u0013\u0257\b\u0013\u0001\u0014\u0001\u0014\u0005\u0014"+
		"\u025b\b\u0014\n\u0014\f\u0014\u025e\t\u0014\u0001\u0014\u0001\u0014\u0005"+
		"\u0014\u0262\b\u0014\n\u0014\f\u0014\u0265\t\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u026c\b\u0015\n\u0015"+
		"\f\u0015\u026f\t\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u0273\b\u0015"+
		"\n\u0015\f\u0015\u0276\t\u0015\u0001\u0015\u0003\u0015\u0279\b\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u0282\b\u0016\u0001\u0017\u0001\u0017\u0005\u0017\u0286"+
		"\b\u0017\n\u0017\f\u0017\u0289\t\u0017\u0001\u0017\u0001\u0017\u0005\u0017"+
		"\u028d\b\u0017\n\u0017\f\u0017\u0290\t\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0005\u0018\u0297\b\u0018\n\u0018\f\u0018"+
		"\u029a\t\u0018\u0001\u0018\u0001\u0018\u0005\u0018\u029e\b\u0018\n\u0018"+
		"\f\u0018\u02a1\t\u0018\u0001\u0019\u0001\u0019\u0003\u0019\u02a5\b\u0019"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0003\u001a\u02aa\b\u001a\u0001\u001a"+
		"\u0001\u001a\u0005\u001a\u02ae\b\u001a\n\u001a\f\u001a\u02b1\t\u001a\u0003"+
		"\u001a\u02b3\b\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0003\u001a\u02c0\b\u001a\u0001\u001b\u0001\u001b\u0005\u001b\u02c4"+
		"\b\u001b\n\u001b\f\u001b\u02c7\t\u001b\u0001\u001b\u0003\u001b\u02ca\b"+
		"\u001b\u0001\u001b\u0005\u001b\u02cd\b\u001b\n\u001b\f\u001b\u02d0\t\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0005\u001c"+
		"\u02d7\b\u001c\n\u001c\f\u001c\u02da\t\u001c\u0001\u001c\u0003\u001c\u02dd"+
		"\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u02e2\b\u001d"+
		"\n\u001d\f\u001d\u02e5\t\u001d\u0001\u001e\u0001\u001e\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u02ef"+
		"\b\u001f\u0001 \u0001 \u0001 \u0005 \u02f4\b \n \f \u02f7\t \u0001!\u0001"+
		"!\u0001\"\u0001\"\u0001\"\u0001#\u0001#\u0001#\u0001$\u0001$\u0001$\u0003"+
		"$\u0304\b$\u0001%\u0001%\u0001%\u0001&\u0001&\u0001&\u0005&\u030c\b&\n"+
		"&\f&\u030f\t&\u0001&\u0001&\u0005&\u0313\b&\n&\f&\u0316\t&\u0001&\u0001"+
		"&\u0001\'\u0001\'\u0003\'\u031c\b\'\u0001\'\u0001\'\u0005\'\u0320\b\'"+
		"\n\'\f\'\u0323\t\'\u0001\'\u0001\'\u0003\'\u0327\b\'\u0005\'\u0329\b\'"+
		"\n\'\f\'\u032c\t\'\u0001\'\u0003\'\u032f\b\'\u0001(\u0001(\u0003(\u0333"+
		"\b(\u0001)\u0001)\u0003)\u0337\b)\u0001*\u0001*\u0001*\u0001+\u0001+\u0001"+
		"+\u0005+\u033f\b+\n+\f+\u0342\t+\u0001+\u0001+\u0005+\u0346\b+\n+\f+\u0349"+
		"\t+\u0001+\u0001+\u0005+\u034d\b+\n+\f+\u0350\t+\u0001+\u0001+\u0001+"+
		"\u0005+\u0355\b+\n+\f+\u0358\t+\u0001+\u0001+\u0005+\u035c\b+\n+\f+\u035f"+
		"\t+\u0001+\u0003+\u0362\b+\u0001+\u0005+\u0365\b+\n+\f+\u0368\t+\u0001"+
		"+\u0001+\u0003+\u036c\b+\u0001,\u0001,\u0001,\u0003,\u0371\b,\u0001,\u0001"+
		",\u0001,\u0003,\u0376\b,\u0001,\u0001,\u0001-\u0001-\u0003-\u037c\b-\u0001"+
		".\u0001.\u0005.\u0380\b.\n.\f.\u0383\t.\u0001.\u0001.\u0001/\u0001/\u0001"+
		"/\u0001/\u0001/\u0001/\u0003/\u038d\b/\u00010\u00010\u00010\u00030\u0392"+
		"\b0\u00010\u00050\u0395\b0\n0\f0\u0398\t0\u00010\u00010\u00011\u00011"+
		"\u00012\u00012\u00052\u03a0\b2\n2\f2\u03a3\t2\u00012\u00012\u00032\u03a7"+
		"\b2\u00012\u00032\u03aa\b2\u00013\u00013\u00013\u00033\u03af\b3\u0001"+
		"3\u00013\u00053\u03b3\b3\n3\f3\u03b6\t3\u00013\u00033\u03b9\b3\u00014"+
		"\u00014\u00034\u03bd\b4\u00015\u00015\u00035\u03c1\b5\u00015\u00055\u03c4"+
		"\b5\n5\f5\u03c7\t5\u00015\u00035\u03ca\b5\u00016\u00016\u00016\u00036"+
		"\u03cf\b6\u00036\u03d1\b6\u00016\u00016\u00036\u03d5\b6\u00016\u00056"+
		"\u03d8\b6\n6\f6\u03db\t6\u00016\u00036\u03de\b6\u00017\u00017\u00037\u03e2"+
		"\b7\u00017\u00057\u03e5\b7\n7\f7\u03e8\t7\u00017\u00037\u03eb\b7\u0001"+
		"7\u00057\u03ee\b7\n7\f7\u03f1\t7\u00017\u00017\u00018\u00018\u00018\u0001"+
		"8\u00058\u03f9\b8\n8\f8\u03fc\t8\u00018\u00018\u00058\u0400\b8\n8\f8\u0403"+
		"\t8\u00018\u00018\u00038\u0407\b8\u00019\u00019\u00019\u00059\u040c\b"+
		"9\n9\f9\u040f\t9\u00019\u00019\u00059\u0413\b9\n9\f9\u0416\t9\u00019\u0003"+
		"9\u0419\b9\u0001:\u0001:\u0001:\u0005:\u041e\b:\n:\f:\u0421\t:\u0001:"+
		"\u0003:\u0424\b:\u0001:\u0001:\u0003:\u0428\b:\u0001:\u0001:\u0001:\u0003"+
		":\u042d\b:\u0003:\u042f\b:\u0001;\u0001;\u0001;\u0003;\u0434\b;\u0001"+
		";\u0005;\u0437\b;\n;\f;\u043a\t;\u0001;\u0003;\u043d\b;\u0001;\u0005;"+
		"\u0440\b;\n;\f;\u0443\t;\u0001;\u0001;\u0001<\u0001<\u0003<\u0449\b<\u0001"+
		"<\u0005<\u044c\b<\n<\f<\u044f\t<\u0001<\u0003<\u0452\b<\u0001<\u0005<"+
		"\u0455\b<\n<\f<\u0458\t<\u0001<\u0001<\u0001=\u0001=\u0001>\u0001>\u0001"+
		">\u0005>\u0461\b>\n>\f>\u0464\t>\u0001>\u0001>\u0005>\u0468\b>\n>\f>\u046b"+
		"\t>\u0001>\u0003>\u046e\b>\u0001?\u0001?\u0001?\u0003?\u0473\b?\u0001"+
		"?\u0001?\u0001?\u0003?\u0478\b?\u0001?\u0001?\u0003?\u047c\b?\u0001?\u0005"+
		"?\u047f\b?\n?\f?\u0482\t?\u0001?\u0003?\u0485\b?\u0001?\u0003?\u0488\b"+
		"?\u0001@\u0001@\u0001@\u0003@\u048d\b@\u0001@\u0001@\u0001A\u0001A\u0001"+
		"B\u0005B\u0494\bB\nB\fB\u0497\tB\u0001B\u0003B\u049a\bB\u0001B\u0005B"+
		"\u049d\bB\nB\fB\u04a0\tB\u0001B\u0003B\u04a3\bB\u0001C\u0001C\u0001C\u0001"+
		"C\u0003C\u04a9\bC\u0001C\u0005C\u04ac\bC\nC\fC\u04af\tC\u0001C\u0003C"+
		"\u04b2\bC\u0001C\u0005C\u04b5\bC\nC\fC\u04b8\tC\u0001C\u0001C\u0001D\u0001"+
		"D\u0001D\u0005D\u04bf\bD\nD\fD\u04c2\tD\u0001D\u0001D\u0005D\u04c6\bD"+
		"\nD\fD\u04c9\tD\u0001D\u0003D\u04cc\bD\u0001E\u0001E\u0001E\u0003E\u04d1"+
		"\bE\u0001E\u0001E\u0001E\u0003E\u04d6\bE\u0001E\u0001E\u0001E\u0001E\u0001"+
		"E\u0003E\u04dd\bE\u0001E\u0003E\u04e0\bE\u0001F\u0001F\u0001F\u0005F\u04e5"+
		"\bF\nF\fF\u04e8\tF\u0001F\u0003F\u04eb\bF\u0001G\u0001G\u0001G\u0003G"+
		"\u04f0\bG\u0001G\u0001G\u0001H\u0001H\u0001I\u0005I\u04f7\bI\nI\fI\u04fa"+
		"\tI\u0001I\u0003I\u04fd\bI\u0001I\u0005I\u0500\bI\nI\fI\u0503\tI\u0001"+
		"I\u0001I\u0001J\u0001J\u0003J\u0509\bJ\u0001J\u0005J\u050c\bJ\nJ\fJ\u050f"+
		"\tJ\u0001J\u0003J\u0512\bJ\u0001J\u0005J\u0515\bJ\nJ\fJ\u0518\tJ\u0001"+
		"J\u0001J\u0001K\u0001K\u0001K\u0005K\u051f\bK\nK\fK\u0522\tK\u0001K\u0001"+
		"K\u0005K\u0526\bK\nK\fK\u0529\tK\u0001K\u0003K\u052c\bK\u0001L\u0001L"+
		"\u0001L\u0003L\u0531\bL\u0001L\u0001L\u0001L\u0003L\u0536\bL\u0001L\u0001"+
		"L\u0003L\u053a\bL\u0001L\u0003L\u053d\bL\u0001M\u0001M\u0003M\u0541\b"+
		"M\u0001M\u0005M\u0544\bM\nM\fM\u0547\tM\u0001M\u0001M\u0001N\u0001N\u0001"+
		"N\u0003N\u054e\bN\u0001N\u0001N\u0001N\u0005N\u0553\bN\nN\fN\u0556\tN"+
		"\u0001N\u0003N\u0559\bN\u0001N\u0003N\u055c\bN\u0003N\u055e\bN\u0001O"+
		"\u0001O\u0001O\u0003O\u0563\bO\u0001O\u0005O\u0566\bO\nO\fO\u0569\tO\u0001"+
		"O\u0001O\u0001P\u0001P\u0005P\u056f\bP\nP\fP\u0572\tP\u0001P\u0001P\u0001"+
		"P\u0003P\u0577\bP\u0001P\u0001P\u0003P\u057b\bP\u0001P\u0003P\u057e\b"+
		"P\u0001Q\u0001Q\u0001Q\u0003Q\u0583\bQ\u0001Q\u0001Q\u0003Q\u0587\bQ\u0001"+
		"Q\u0001Q\u0003Q\u058b\bQ\u0001Q\u0001Q\u0001Q\u0001Q\u0003Q\u0591\bQ\u0001"+
		"Q\u0001Q\u0001Q\u0005Q\u0596\bQ\nQ\fQ\u0599\tQ\u0001R\u0001R\u0001S\u0001"+
		"S\u0001T\u0001T\u0003T\u05a1\bT\u0001T\u0001T\u0001U\u0001U\u0001U\u0005"+
		"U\u05a8\bU\nU\fU\u05ab\tU\u0001V\u0001V\u0001W\u0001W\u0001W\u0001X\u0001"+
		"X\u0001Y\u0001Y\u0001Y\u0001Y\u0001Y\u0003Y\u05b9\bY\u0001Y\u0003Y\u05bc"+
		"\bY\u0003Y\u05be\bY\u0001Z\u0001Z\u0001Z\u0003Z\u05c3\bZ\u0001Z\u0001"+
		"Z\u0001[\u0001[\u0001\\\u0001\\\u0005\\\u05cb\b\\\n\\\f\\\u05ce\t\\\u0001"+
		"\\\u0003\\\u05d1\b\\\u0001\\\u0005\\\u05d4\b\\\n\\\f\\\u05d7\t\\\u0001"+
		"\\\u0001\\\u0001]\u0001]\u0003]\u05dd\b]\u0003]\u05df\b]\u0001]\u0001"+
		"]\u0001^\u0001^\u0001^\u0005^\u05e6\b^\n^\f^\u05e9\t^\u0001^\u0001^\u0005"+
		"^\u05ed\b^\n^\f^\u05f0\t^\u0001^\u0003^\u05f3\b^\u0001_\u0001_\u0003_"+
		"\u05f7\b_\u0001_\u0005_\u05fa\b_\n_\f_\u05fd\t_\u0001`\u0001`\u0003`\u0601"+
		"\b`\u0001a\u0001a\u0001a\u0005a\u0606\ba\na\fa\u0609\ta\u0001a\u0001a"+
		"\u0005a\u060d\ba\na\fa\u0610\ta\u0001a\u0003a\u0613\ba\u0001b\u0001b\u0001"+
		"b\u0001b\u0003b\u0619\bb\u0001c\u0001c\u0001c\u0001c\u0001c\u0001c\u0001"+
		"c\u0001c\u0001c\u0001c\u0003c\u0625\bc\u0001d\u0004d\u0628\bd\u000bd\f"+
		"d\u0629\u0001e\u0001e\u0001e\u0001e\u0001f\u0001f\u0001f\u0001f\u0003"+
		"f\u0634\bf\u0001g\u0001g\u0001g\u0003g\u0639\bg\u0001g\u0001g\u0001g\u0001"+
		"g\u0003g\u063f\bg\u0001g\u0001g\u0001g\u0001g\u0001g\u0001g\u0003g\u0647"+
		"\bg\u0001h\u0001h\u0001h\u0001h\u0001h\u0001h\u0001h\u0003h\u0650\bh\u0001"+
		"i\u0001i\u0001i\u0001j\u0001j\u0001j\u0001j\u0003j\u0659\bj\u0001j\u0001"+
		"j\u0001j\u0003j\u065e\bj\u0001k\u0001k\u0001k\u0001l\u0001l\u0001m\u0001"+
		"m\u0005m\u0667\bm\nm\fm\u066a\tm\u0001m\u0003m\u066d\bm\u0001m\u0005m"+
		"\u0670\bm\nm\fm\u0673\tm\u0001m\u0001m\u0001n\u0001n\u0001n\u0005n\u067a"+
		"\bn\nn\fn\u067d\tn\u0001n\u0001n\u0005n\u0681\bn\nn\fn\u0684\tn\u0001"+
		"n\u0003n\u0687\bn\u0001o\u0001o\u0001p\u0001p\u0005p\u068d\bp\np\fp\u0690"+
		"\tp\u0001p\u0003p\u0693\bp\u0001p\u0005p\u0696\bp\np\fp\u0699\tp\u0001"+
		"p\u0001p\u0001q\u0001q\u0001q\u0005q\u06a0\bq\nq\fq\u06a3\tq\u0001q\u0001"+
		"q\u0005q\u06a7\bq\nq\fq\u06aa\tq\u0001q\u0003q\u06ad\bq\u0001r\u0001r"+
		"\u0003r\u06b1\br\u0001r\u0001r\u0001r\u0001s\u0001s\u0005s\u06b8\bs\n"+
		"s\fs\u06bb\ts\u0001s\u0003s\u06be\bs\u0001s\u0005s\u06c1\bs\ns\fs\u06c4"+
		"\ts\u0001s\u0001s\u0001t\u0001t\u0001t\u0005t\u06cb\bt\nt\ft\u06ce\tt"+
		"\u0001t\u0001t\u0005t\u06d2\bt\nt\ft\u06d5\tt\u0001t\u0003t\u06d8\bt\u0001"+
		"u\u0001u\u0001u\u0003u\u06dd\bu\u0001v\u0001v\u0001v\u0001w\u0001w\u0001"+
		"w\u0001x\u0001x\u0001x\u0001x\u0001x\u0001x\u0001x\u0005x\u06ec\bx\nx"+
		"\fx\u06ef\tx\u0001x\u0001x\u0005x\u06f3\bx\nx\fx\u06f6\tx\u0001x\u0003"+
		"x\u06f9\bx\u0001x\u0005x\u06fc\bx\nx\fx\u06ff\tx\u0001x\u0001x\u0001x"+
		"\u0003x\u0704\bx\u0001y\u0001y\u0001y\u0005y\u0709\by\ny\fy\u070c\ty\u0001"+
		"y\u0001y\u0005y\u0710\by\ny\fy\u0713\ty\u0001y\u0003y\u0716\by\u0001z"+
		"\u0001z\u0003z\u071a\bz\u0001{\u0001{\u0001{\u0001|\u0001|\u0005|\u0721"+
		"\b|\n|\f|\u0724\t|\u0001|\u0001|\u0005|\u0728\b|\n|\f|\u072b\t|\u0001"+
		"|\u0001|\u0001}\u0001}\u0001}\u0001}\u0001}\u0001}\u0004}\u0735\b}\u000b"+
		"}\f}\u0736\u0001}\u0001}\u0003}\u073b\b}\u0001~\u0001~\u0001~\u0003~\u0740"+
		"\b~\u0001~\u0001~\u0003~\u0744\b~\u0001\u007f\u0001\u007f\u0001\u0080"+
		"\u0001\u0080\u0005\u0080\u074a\b\u0080\n\u0080\f\u0080\u074d\t\u0080\u0001"+
		"\u0080\u0003\u0080\u0750\b\u0080\u0001\u0081\u0001\u0081\u0001\u0081\u0003"+
		"\u0081\u0755\b\u0081\u0001\u0082\u0001\u0082\u0001\u0082\u0001\u0082\u0001"+
		"\u0082\u0001\u0082\u0001\u0082\u0001\u0082\u0003\u0082\u075f\b\u0082\u0003"+
		"\u0082\u0761\b\u0082\u0001\u0083\u0001\u0083\u0001\u0083\u0001\u0083\u0001"+
		"\u0084\u0003\u0084\u0768\b\u0084\u0001\u0084\u0001\u0084\u0003\u0084\u076c"+
		"\b\u0084\u0001\u0085\u0001\u0085\u0001\u0085\u0001\u0085\u0001\u0085\u0001"+
		"\u0085\u0003\u0085\u0774\b\u0085\u0001\u0086\u0001\u0086\u0001\u0086\u0005"+
		"\u0086\u0779\b\u0086\n\u0086\f\u0086\u077c\t\u0086\u0001\u0087\u0001\u0087"+
		"\u0001\u0087\u0003\u0087\u0781\b\u0087\u0001\u0087\u0001\u0087\u0001\u0087"+
		"\u0001\u0087\u0001\u0087\u0003\u0087\u0788\b\u0087\u0001\u0087\u0001\u0087"+
		"\u0003\u0087\u078c\b\u0087\u0001\u0088\u0001\u0088\u0003\u0088\u0790\b"+
		"\u0088\u0001\u0089\u0004\u0089\u0793\b\u0089\u000b\u0089\f\u0089\u0794"+
		"\u0001\u008a\u0001\u008a\u0001\u008a\u0001\u008a\u0001\u008b\u0001\u008b"+
		"\u0005\u008b\u079d\b\u008b\n\u008b\f\u008b\u07a0\t\u008b\u0001\u008c\u0001"+
		"\u008c\u0001\u008c\u0001\u008d\u0001\u008d\u0001\u008d\u0003\u008d\u07a8"+
		"\b\u008d\u0001\u008d\u0001\u008d\u0001\u008d\u0001\u008d\u0001\u008d\u0001"+
		"\u008d\u0005\u008d\u07b0\b\u008d\n\u008d\f\u008d\u07b3\t\u008d\u0001\u008e"+
		"\u0001\u008e\u0001\u008e\u0001\u008e\u0001\u008e\u0003\u008e\u07ba\b\u008e"+
		"\u0001\u008e\u0001\u008e\u0001\u008e\u0003\u008e\u07bf\b\u008e\u0001\u008e"+
		"\u0005\u008e\u07c2\b\u008e\n\u008e\f\u008e\u07c5\t\u008e\u0001\u008e\u0001"+
		"\u008e\u0005\u008e\u07c9\b\u008e\n\u008e\f\u008e\u07cc\t\u008e\u0001\u008e"+
		"\u0001\u008e\u0003\u008e\u07d0\b\u008e\u0001\u008e\u0001\u008e\u0001\u008e"+
		"\u0003\u008e\u07d5\b\u008e\u0001\u008e\u0001\u008e\u0003\u008e\u07d9\b"+
		"\u008e\u0001\u008e\u0001\u008e\u0001\u008e\u0003\u008e\u07de\b\u008e\u0001"+
		"\u008e\u0005\u008e\u07e1\b\u008e\n\u008e\f\u008e\u07e4\t\u008e\u0001\u008e"+
		"\u0001\u008e\u0005\u008e\u07e8\b\u008e\n\u008e\f\u008e\u07eb\t\u008e\u0001"+
		"\u008e\u0001\u008e\u0003\u008e\u07ef\b\u008e\u0001\u008e\u0001\u008e\u0001"+
		"\u008e\u0003\u008e\u07f4\b\u008e\u0005\u008e\u07f6\b\u008e\n\u008e\f\u008e"+
		"\u07f9\t\u008e\u0001\u008f\u0001\u008f\u0001\u008f\u0001\u008f\u0003\u008f"+
		"\u07ff\b\u008f\u0001\u0090\u0003\u0090\u0802\b\u0090\u0001\u0090\u0001"+
		"\u0090\u0003\u0090\u0806\b\u0090\u0001\u0091\u0001\u0091\u0001\u0091\u0003"+
		"\u0091\u080b\b\u0091\u0001\u0091\u0001\u0091\u0001\u0091\u0005\u0091\u0810"+
		"\b\u0091\n\u0091\f\u0091\u0813\t\u0091\u0001\u0092\u0001\u0092\u0003\u0092"+
		"\u0817\b\u0092\u0001\u0093\u0001\u0093\u0001\u0094\u0001\u0094\u0005\u0094"+
		"\u081d\b\u0094\n\u0094\f\u0094\u0820\t\u0094\u0001\u0094\u0003\u0094\u0823"+
		"\b\u0094\u0001\u0094\u0005\u0094\u0826\b\u0094\n\u0094\f\u0094\u0829\t"+
		"\u0094\u0001\u0094\u0001\u0094\u0001\u0095\u0001\u0095\u0001\u0095\u0005"+
		"\u0095\u0830\b\u0095\n\u0095\f\u0095\u0833\t\u0095\u0001\u0095\u0001\u0095"+
		"\u0005\u0095\u0837\b\u0095\n\u0095\f\u0095\u083a\t\u0095\u0001\u0095\u0003"+
		"\u0095\u083d\b\u0095\u0001\u0096\u0001\u0096\u0003\u0096\u0841\b\u0096"+
		"\u0003\u0096\u0843\b\u0096\u0001\u0096\u0001\u0096\u0003\u0096\u0847\b"+
		"\u0096\u0001\u0097\u0001\u0097\u0001\u0097\u0001\u0097\u0003\u0097\u084d"+
		"\b\u0097\u0001\u0098\u0001\u0098\u0001\u0098\u0003\u0098\u0852\b\u0098"+
		"\u0001\u0098\u0001\u0098\u0001\u0099\u0001\u0099\u0001\u0099\u0003\u0099"+
		"\u0859\b\u0099\u0001\u0099\u0001\u0099\u0001\u0099\u0003\u0099\u085e\b"+
		"\u0099\u0001\u0099\u0001\u0099\u0001\u009a\u0001\u009a\u0001\u009b\u0001"+
		"\u009b\u0005\u009b\u0866\b\u009b\n\u009b\f\u009b\u0869\t\u009b\u0001\u009b"+
		"\u0001\u009b\u0001\u009c\u0001\u009c\u0001\u009c\u0005\u009c\u0870\b\u009c"+
		"\n\u009c\f\u009c\u0873\t\u009c\u0001\u009c\u0001\u009c\u0005\u009c\u0877"+
		"\b\u009c\n\u009c\f\u009c\u087a\t\u009c\u0001\u009c\u0003\u009c\u087d\b"+
		"\u009c\u0001\u009d\u0001\u009d\u0003\u009d\u0881\b\u009d\u0001\u009e\u0001"+
		"\u009e\u0003\u009e\u0885\b\u009e\u0001\u009f\u0001\u009f\u0003\u009f\u0889"+
		"\b\u009f\u0001\u00a0\u0001\u00a0\u0001\u00a0\u0005\u00a0\u088e\b\u00a0"+
		"\n\u00a0\f\u00a0\u0891\t\u00a0\u0001\u00a1\u0001\u00a1\u0001\u00a2\u0001"+
		"\u00a2\u0001\u00a3\u0001\u00a3\u0001\u00a4\u0001\u00a4\u0001\u00a4\u0005"+
		"\u00a4\u089c\b\u00a4\n\u00a4\f\u00a4\u089f\t\u00a4\u0001\u00a5\u0001\u00a5"+
		"\u0001\u00a5\u0005\u00a5\u08a4\b\u00a5\n\u00a5\f\u00a5\u08a7\t\u00a5\u0001"+
		"\u00a6\u0001\u00a6\u0001\u00a7\u0001\u00a7\u0001\u00a8\u0001\u00a8\u0001"+
		"\u00a9\u0001\u00a9\u0001\u00aa\u0001\u00aa\u0001\u00ab\u0001\u00ab\u0001"+
		"\u00ab\u0001\u00ab\u0001\u00ab\u0003\u00ab\u08b8\b\u00ab\u0001\u00ac\u0001"+
		"\u00ac\u0003\u00ac\u08bc\b\u00ac\u0001\u00ad\u0001\u00ad\u0001\u00ae\u0001"+
		"\u00ae\u0003\u00ae\u08c2\b\u00ae\u0001\u00ae\u0001\u00ae\u0005\u00ae\u08c6"+
		"\b\u00ae\n\u00ae\f\u00ae\u08c9\t\u00ae\u0003\u00ae\u08cb\b\u00ae\u0001"+
		"\u00af\u0001\u00af\u0004\u00af\u08cf\b\u00af\u000b\u00af\f\u00af\u08d0"+
		"\u0001\u00b0\u0001\u00b0\u0003\u00b0\u08d5\b\u00b0\u0001\u00b1\u0001\u00b1"+
		"\u0003\u00b1\u08d9\b\u00b1\u0001\u00b2\u0001\u00b2\u0001\u00b3\u0001\u00b3"+
		"\u0003\u00b3\u08df\b\u00b3\u0001\u00b4\u0001\u00b4\u0001\u00b4\u0001\u00b4"+
		"\u0003\u00b4\u08e5\b\u00b4\u0001\u00b5\u0001\u00b5\u0001\u00b6\u0001\u00b6"+
		"\u0001\u00b7\u0003\u00b7\u08ec\b\u00b7\u0001\u00b7\u0001\u00b7\u0003\u00b7"+
		"\u08f0\b\u00b7\u0001\u00b7\u0003\u00b7\u08f3\b\u00b7\u0001\u00b8\u0001"+
		"\u00b8\u0001\u00b9\u0001\u00b9\u0001\u00ba\u0001\u00ba\u0001\u00bb\u0001"+
		"\u00bb\u0001\u00bc\u0001\u00bc\u0001\u00bc\u0001\u00bc\u0003\u00bc\u0901"+
		"\b\u00bc\u0001\u00bc\u0003\u00bc\u0904\b\u00bc\u0001\u00bc\u0003\u00bc"+
		"\u0907\b\u00bc\u0001\u00bd\u0001\u00bd\u0001\u00bd\u0001\u00bd\u0003\u00bd"+
		"\u090d\b\u00bd\u0001\u00bd\u0003\u00bd\u0910\b\u00bd\u0001\u00bd\u0003"+
		"\u00bd\u0913\b\u00bd\u0001\u00bd\u0000\u0003\u00a2\u011a\u011c\u00be\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.02468:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084"+
		"\u0086\u0088\u008a\u008c\u008e\u0090\u0092\u0094\u0096\u0098\u009a\u009c"+
		"\u009e\u00a0\u00a2\u00a4\u00a6\u00a8\u00aa\u00ac\u00ae\u00b0\u00b2\u00b4"+
		"\u00b6\u00b8\u00ba\u00bc\u00be\u00c0\u00c2\u00c4\u00c6\u00c8\u00ca\u00cc"+
		"\u00ce\u00d0\u00d2\u00d4\u00d6\u00d8\u00da\u00dc\u00de\u00e0\u00e2\u00e4"+
		"\u00e6\u00e8\u00ea\u00ec\u00ee\u00f0\u00f2\u00f4\u00f6\u00f8\u00fa\u00fc"+
		"\u00fe\u0100\u0102\u0104\u0106\u0108\u010a\u010c\u010e\u0110\u0112\u0114"+
		"\u0116\u0118\u011a\u011c\u011e\u0120\u0122\u0124\u0126\u0128\u012a\u012c"+
		"\u012e\u0130\u0132\u0134\u0136\u0138\u013a\u013c\u013e\u0140\u0142\u0144"+
		"\u0146\u0148\u014a\u014c\u014e\u0150\u0152\u0154\u0156\u0158\u015a\u015c"+
		"\u015e\u0160\u0162\u0164\u0166\u0168\u016a\u016c\u016e\u0170\u0172\u0174"+
		"\u0176\u0178\u017a\u0000\u000b\u0001\u0000IJ\u0002\u0000JJLL\u0007\u0000"+
		"\u0001\u0005\b\b\u000b\u000b\r\u0016\u0018\u0018\u001a\u001a\u001d\u001d"+
		"\u0001\u0000\u0001\u001d\u0001\u0000:;\u0002\u0000(+-6\u0002\u0000\t\t"+
		"\u0019\u0019\u0001\u0000MQ\u0001\u0000ST\u0002\u0000\'\'YY\u0002\u0000"+
		"%%YY\u09cd\u0000\u017f\u0001\u0000\u0000\u0000\u0002\u0193\u0001\u0000"+
		"\u0000\u0000\u0004\u0195\u0001\u0000\u0000\u0000\u0006\u0197\u0001\u0000"+
		"\u0000\u0000\b\u01ab\u0001\u0000\u0000\u0000\n\u01ad\u0001\u0000\u0000"+
		"\u0000\f\u01bf\u0001\u0000\u0000\u0000\u000e\u01c9\u0001\u0000\u0000\u0000"+
		"\u0010\u01da\u0001\u0000\u0000\u0000\u0012\u01dc\u0001\u0000\u0000\u0000"+
		"\u0014\u01e8\u0001\u0000\u0000\u0000\u0016\u01ea\u0001\u0000\u0000\u0000"+
		"\u0018\u020c\u0001\u0000\u0000\u0000\u001a\u020e\u0001\u0000\u0000\u0000"+
		"\u001c\u0228\u0001\u0000\u0000\u0000\u001e\u0239\u0001\u0000\u0000\u0000"+
		" \u024e\u0001\u0000\u0000\u0000\"\u0250\u0001\u0000\u0000\u0000$\u0252"+
		"\u0001\u0000\u0000\u0000&\u0254\u0001\u0000\u0000\u0000(\u0258\u0001\u0000"+
		"\u0000\u0000*\u0268\u0001\u0000\u0000\u0000,\u0281\u0001\u0000\u0000\u0000"+
		".\u0283\u0001\u0000\u0000\u00000\u0293\u0001\u0000\u0000\u00002\u02a2"+
		"\u0001\u0000\u0000\u00004\u02a9\u0001\u0000\u0000\u00006\u02c1\u0001\u0000"+
		"\u0000\u00008\u02d3\u0001\u0000\u0000\u0000:\u02de\u0001\u0000\u0000\u0000"+
		"<\u02e6\u0001\u0000\u0000\u0000>\u02e8\u0001\u0000\u0000\u0000@\u02f0"+
		"\u0001\u0000\u0000\u0000B\u02f8\u0001\u0000\u0000\u0000D\u02fa\u0001\u0000"+
		"\u0000\u0000F\u02fd\u0001\u0000\u0000\u0000H\u0300\u0001\u0000\u0000\u0000"+
		"J\u0305\u0001\u0000\u0000\u0000L\u0308\u0001\u0000\u0000\u0000N\u031b"+
		"\u0001\u0000\u0000\u0000P\u0330\u0001\u0000\u0000\u0000R\u0334\u0001\u0000"+
		"\u0000\u0000T\u0338\u0001\u0000\u0000\u0000V\u036b\u0001\u0000\u0000\u0000"+
		"X\u0370\u0001\u0000\u0000\u0000Z\u0379\u0001\u0000\u0000\u0000\\\u037d"+
		"\u0001\u0000\u0000\u0000^\u038c\u0001\u0000\u0000\u0000`\u038e\u0001\u0000"+
		"\u0000\u0000b\u039b\u0001\u0000\u0000\u0000d\u039d\u0001\u0000\u0000\u0000"+
		"f\u03ab\u0001\u0000\u0000\u0000h\u03bc\u0001\u0000\u0000\u0000j\u03be"+
		"\u0001\u0000\u0000\u0000l\u03cb\u0001\u0000\u0000\u0000n\u03df\u0001\u0000"+
		"\u0000\u0000p\u0406\u0001\u0000\u0000\u0000r\u0408\u0001\u0000\u0000\u0000"+
		"t\u042e\u0001\u0000\u0000\u0000v\u0430\u0001\u0000\u0000\u0000x\u0446"+
		"\u0001\u0000\u0000\u0000z\u045b\u0001\u0000\u0000\u0000|\u045d\u0001\u0000"+
		"\u0000\u0000~\u0487\u0001\u0000\u0000\u0000\u0080\u0489\u0001\u0000\u0000"+
		"\u0000\u0082\u0490\u0001\u0000\u0000\u0000\u0084\u0499\u0001\u0000\u0000"+
		"\u0000\u0086\u04a4\u0001\u0000\u0000\u0000\u0088\u04bb\u0001\u0000\u0000"+
		"\u0000\u008a\u04df\u0001\u0000\u0000\u0000\u008c\u04e1\u0001\u0000\u0000"+
		"\u0000\u008e\u04ec\u0001\u0000\u0000\u0000\u0090\u04f3\u0001\u0000\u0000"+
		"\u0000\u0092\u04fc\u0001\u0000\u0000\u0000\u0094\u0506\u0001\u0000\u0000"+
		"\u0000\u0096\u051b\u0001\u0000\u0000\u0000\u0098\u053c\u0001\u0000\u0000"+
		"\u0000\u009a\u053e\u0001\u0000\u0000\u0000\u009c\u054a\u0001\u0000\u0000"+
		"\u0000\u009e\u055f\u0001\u0000\u0000\u0000\u00a0\u056c\u0001\u0000\u0000"+
		"\u0000\u00a2\u0590\u0001\u0000\u0000\u0000\u00a4\u059a\u0001\u0000\u0000"+
		"\u0000\u00a6\u059c\u0001\u0000\u0000\u0000\u00a8\u059e\u0001\u0000\u0000"+
		"\u0000\u00aa\u05a4\u0001\u0000\u0000\u0000\u00ac\u05ac\u0001\u0000\u0000"+
		"\u0000\u00ae\u05ae\u0001\u0000\u0000\u0000\u00b0\u05b1\u0001\u0000\u0000"+
		"\u0000\u00b2\u05bd\u0001\u0000\u0000\u0000\u00b4\u05c2\u0001\u0000\u0000"+
		"\u0000\u00b6\u05c6\u0001\u0000\u0000\u0000\u00b8\u05c8\u0001\u0000\u0000"+
		"\u0000\u00ba\u05de\u0001\u0000\u0000\u0000\u00bc\u05e2\u0001\u0000\u0000"+
		"\u0000\u00be\u05f4\u0001\u0000\u0000\u0000\u00c0\u05fe\u0001\u0000\u0000"+
		"\u0000\u00c2\u0602\u0001\u0000\u0000\u0000\u00c4\u0618\u0001\u0000\u0000"+
		"\u0000\u00c6\u0624\u0001\u0000\u0000\u0000\u00c8\u0627\u0001\u0000\u0000"+
		"\u0000\u00ca\u062b\u0001\u0000\u0000\u0000\u00cc\u0633\u0001\u0000\u0000"+
		"\u0000\u00ce\u0646\u0001\u0000\u0000\u0000\u00d0\u064f\u0001\u0000\u0000"+
		"\u0000\u00d2\u0651\u0001\u0000\u0000\u0000\u00d4\u065d\u0001\u0000\u0000"+
		"\u0000\u00d6\u065f\u0001\u0000\u0000\u0000\u00d8\u0662\u0001\u0000\u0000"+
		"\u0000\u00da\u0664\u0001\u0000\u0000\u0000\u00dc\u0676\u0001\u0000\u0000"+
		"\u0000\u00de\u0688\u0001\u0000\u0000\u0000\u00e0\u068a\u0001\u0000\u0000"+
		"\u0000\u00e2\u069c\u0001\u0000\u0000\u0000\u00e4\u06b0\u0001\u0000\u0000"+
		"\u0000\u00e6\u06b5\u0001\u0000\u0000\u0000\u00e8\u06c7\u0001\u0000\u0000"+
		"\u0000\u00ea\u06d9\u0001\u0000\u0000\u0000\u00ec\u06de\u0001\u0000\u0000"+
		"\u0000\u00ee\u06e1\u0001\u0000\u0000\u0000\u00f0\u0703\u0001\u0000\u0000"+
		"\u0000\u00f2\u0705\u0001\u0000\u0000\u0000\u00f4\u0719\u0001\u0000\u0000"+
		"\u0000\u00f6\u071b\u0001\u0000\u0000\u0000\u00f8\u071e\u0001\u0000\u0000"+
		"\u0000\u00fa\u073a\u0001\u0000\u0000\u0000\u00fc\u0743\u0001\u0000\u0000"+
		"\u0000\u00fe\u0745\u0001\u0000\u0000\u0000\u0100\u0747\u0001\u0000\u0000"+
		"\u0000\u0102\u0754\u0001\u0000\u0000\u0000\u0104\u0756\u0001\u0000\u0000"+
		"\u0000\u0106\u0762\u0001\u0000\u0000\u0000\u0108\u076b\u0001\u0000\u0000"+
		"\u0000\u010a\u0773\u0001\u0000\u0000\u0000\u010c\u0775\u0001\u0000\u0000"+
		"\u0000\u010e\u078b\u0001\u0000\u0000\u0000\u0110\u078d\u0001\u0000\u0000"+
		"\u0000\u0112\u0792\u0001\u0000\u0000\u0000\u0114\u0796\u0001\u0000\u0000"+
		"\u0000\u0116\u079a\u0001\u0000\u0000\u0000\u0118\u07a1\u0001\u0000\u0000"+
		"\u0000\u011a\u07a7\u0001\u0000\u0000\u0000\u011c\u07b4\u0001\u0000\u0000"+
		"\u0000\u011e\u07fe\u0001\u0000\u0000\u0000\u0120\u0801\u0001\u0000\u0000"+
		"\u0000\u0122\u080a\u0001\u0000\u0000\u0000\u0124\u0814\u0001\u0000\u0000"+
		"\u0000\u0126\u0818\u0001\u0000\u0000\u0000\u0128\u081a\u0001\u0000\u0000"+
		"\u0000\u012a\u082c\u0001\u0000\u0000\u0000\u012c\u0842\u0001\u0000\u0000"+
		"\u0000\u012e\u0848\u0001\u0000\u0000\u0000\u0130\u084e\u0001\u0000\u0000"+
		"\u0000\u0132\u0855\u0001\u0000\u0000\u0000\u0134\u0861\u0001\u0000\u0000"+
		"\u0000\u0136\u0863\u0001\u0000\u0000\u0000\u0138\u086c\u0001\u0000\u0000"+
		"\u0000\u013a\u087e\u0001\u0000\u0000\u0000\u013c\u0884\u0001\u0000\u0000"+
		"\u0000\u013e\u0888\u0001\u0000\u0000\u0000\u0140\u088a\u0001\u0000\u0000"+
		"\u0000\u0142\u0892\u0001\u0000\u0000\u0000\u0144\u0894\u0001\u0000\u0000"+
		"\u0000\u0146\u0896\u0001\u0000\u0000\u0000\u0148\u0898\u0001\u0000\u0000"+
		"\u0000\u014a\u08a0\u0001\u0000\u0000\u0000\u014c\u08a8\u0001\u0000\u0000"+
		"\u0000\u014e\u08aa\u0001\u0000\u0000\u0000\u0150\u08ac\u0001\u0000\u0000"+
		"\u0000\u0152\u08ae\u0001\u0000\u0000\u0000\u0154\u08b0\u0001\u0000\u0000"+
		"\u0000\u0156\u08b7\u0001\u0000\u0000\u0000\u0158\u08bb\u0001\u0000\u0000"+
		"\u0000\u015a\u08bd\u0001\u0000\u0000\u0000\u015c\u08ca\u0001\u0000\u0000"+
		"\u0000\u015e\u08ce\u0001\u0000\u0000\u0000\u0160\u08d4\u0001\u0000\u0000"+
		"\u0000\u0162\u08d8\u0001\u0000\u0000\u0000\u0164\u08da\u0001\u0000\u0000"+
		"\u0000\u0166\u08de\u0001\u0000\u0000\u0000\u0168\u08e4\u0001\u0000\u0000"+
		"\u0000\u016a\u08e6\u0001\u0000\u0000\u0000\u016c\u08e8\u0001\u0000\u0000"+
		"\u0000\u016e\u08f2\u0001\u0000\u0000\u0000\u0170\u08f4\u0001\u0000\u0000"+
		"\u0000\u0172\u08f6\u0001\u0000\u0000\u0000\u0174\u08f8\u0001\u0000\u0000"+
		"\u0000\u0176\u08fa\u0001\u0000\u0000\u0000\u0178\u0906\u0001\u0000\u0000"+
		"\u0000\u017a\u0912\u0001\u0000\u0000\u0000\u017c\u017e\u0005Y\u0000\u0000"+
		"\u017d\u017c\u0001\u0000\u0000\u0000\u017e\u0181\u0001\u0000\u0000\u0000"+
		"\u017f\u017d\u0001\u0000\u0000\u0000\u017f\u0180\u0001\u0000\u0000\u0000"+
		"\u0180\u0183\u0001\u0000\u0000\u0000\u0181\u017f\u0001\u0000\u0000\u0000"+
		"\u0182\u0184\u0003\u0006\u0003\u0000\u0183\u0182\u0001\u0000\u0000\u0000"+
		"\u0183\u0184\u0001\u0000\u0000\u0000\u0184\u0188\u0001\u0000\u0000\u0000"+
		"\u0185\u0187\u0005Y\u0000\u0000\u0186\u0185\u0001\u0000\u0000\u0000\u0187"+
		"\u018a\u0001\u0000\u0000\u0000\u0188\u0186\u0001\u0000\u0000\u0000\u0188"+
		"\u0189\u0001\u0000\u0000\u0000\u0189\u018b\u0001\u0000\u0000\u0000\u018a"+
		"\u0188\u0001\u0000\u0000\u0000\u018b\u018c\u0005\u0000\u0000\u0001\u018c"+
		"\u0001\u0001\u0000\u0000\u0000\u018d\u0194\u00034\u001a\u0000\u018e\u0194"+
		"\u0003\u00c0`\u0000\u018f\u0194\u0003\b\u0004\u0000\u0190\u0194\u0003"+
		"\u0014\n\u0000\u0191\u0194\u0003 \u0010\u0000\u0192\u0194\u0003\u0004"+
		"\u0002\u0000\u0193\u018d\u0001\u0000\u0000\u0000\u0193\u018e\u0001\u0000"+
		"\u0000\u0000\u0193\u018f\u0001\u0000\u0000\u0000\u0193\u0190\u0001\u0000"+
		"\u0000\u0000\u0193\u0191\u0001\u0000\u0000\u0000\u0193\u0192\u0001\u0000"+
		"\u0000\u0000\u0194\u0003\u0001\u0000\u0000\u0000\u0195\u0196\u0003\u0148"+
		"\u00a4\u0000\u0196\u0005\u0001\u0000\u0000\u0000\u0197\u01a3\u0003\u0002"+
		"\u0001\u0000\u0198\u019c\u0003\u0174\u00ba\u0000\u0199\u019b\u0005Y\u0000"+
		"\u0000\u019a\u0199\u0001\u0000\u0000\u0000\u019b\u019e\u0001\u0000\u0000"+
		"\u0000\u019c\u019a\u0001\u0000\u0000\u0000\u019c\u019d\u0001\u0000\u0000"+
		"\u0000\u019d\u019f\u0001\u0000\u0000\u0000\u019e\u019c\u0001\u0000\u0000"+
		"\u0000\u019f\u01a0\u0003\u0002\u0001\u0000\u01a0\u01a2\u0001\u0000\u0000"+
		"\u0000\u01a1\u0198\u0001\u0000\u0000\u0000\u01a2\u01a5\u0001\u0000\u0000"+
		"\u0000\u01a3\u01a1\u0001\u0000\u0000\u0000\u01a3\u01a4\u0001\u0000\u0000"+
		"\u0000\u01a4\u01a7\u0001\u0000\u0000\u0000\u01a5\u01a3\u0001\u0000\u0000"+
		"\u0000\u01a6\u01a8\u0005\'\u0000\u0000\u01a7\u01a6\u0001\u0000\u0000\u0000"+
		"\u01a7\u01a8\u0001\u0000\u0000\u0000\u01a8\u0007\u0001\u0000\u0000\u0000"+
		"\u01a9\u01ac\u0003\n\u0005\u0000\u01aa\u01ac\u0003\f\u0006\u0000\u01ab"+
		"\u01a9\u0001\u0000\u0000\u0000\u01ab\u01aa\u0001\u0000\u0000\u0000\u01ac"+
		"\t\u0001\u0000\u0000\u0000\u01ad\u01ae\u0005\n\u0000\u0000\u01ae\u01b2"+
		"\u0003\u00a2Q\u0000\u01af\u01b1\u0005Y\u0000\u0000\u01b0\u01af\u0001\u0000"+
		"\u0000\u0000\u01b1\u01b4\u0001\u0000\u0000\u0000\u01b2\u01b0\u0001\u0000"+
		"\u0000\u0000\u01b2\u01b3\u0001\u0000\u0000\u0000\u01b3\u01b5\u0001\u0000"+
		"\u0000\u0000\u01b4\u01b2\u0001\u0000\u0000\u0000\u01b5\u01b6\u0005\u000e"+
		"\u0000\u0000\u01b6\u01ba\u0003\u00c0`\u0000\u01b7\u01b9\u0005Y\u0000\u0000"+
		"\u01b8\u01b7\u0001\u0000\u0000\u0000\u01b9\u01bc\u0001\u0000\u0000\u0000"+
		"\u01ba\u01b8\u0001\u0000\u0000\u0000\u01ba\u01bb\u0001\u0000\u0000\u0000"+
		"\u01bb\u01bd\u0001\u0000\u0000\u0000\u01bc\u01ba\u0001\u0000\u0000\u0000"+
		"\u01bd\u01be\u00036\u001b\u0000\u01be\u000b\u0001\u0000\u0000\u0000\u01bf"+
		"\u01c0\u0005\u001c\u0000\u0000\u01c0\u01c4\u0003\u000e\u0007\u0000\u01c1"+
		"\u01c3\u0005Y\u0000\u0000\u01c2\u01c1\u0001\u0000\u0000\u0000\u01c3\u01c6"+
		"\u0001\u0000\u0000\u0000\u01c4\u01c2\u0001\u0000\u0000\u0000\u01c4\u01c5"+
		"\u0001\u0000\u0000\u0000\u01c5\u01c7\u0001\u0000\u0000\u0000\u01c6\u01c4"+
		"\u0001\u0000\u0000\u0000\u01c7\u01c8\u00036\u001b\u0000\u01c8\r\u0001"+
		"\u0000\u0000\u0000\u01c9\u01d5\u0003\u0010\b\u0000\u01ca\u01ce\u0003\u0176"+
		"\u00bb\u0000\u01cb\u01cd\u0005Y\u0000\u0000\u01cc\u01cb\u0001\u0000\u0000"+
		"\u0000\u01cd\u01d0\u0001\u0000\u0000\u0000\u01ce\u01cc\u0001\u0000\u0000"+
		"\u0000\u01ce\u01cf\u0001\u0000\u0000\u0000\u01cf\u01d1\u0001\u0000\u0000"+
		"\u0000\u01d0\u01ce\u0001\u0000\u0000\u0000\u01d1\u01d2\u0003\u0010\b\u0000"+
		"\u01d2\u01d4\u0001\u0000\u0000\u0000\u01d3\u01ca\u0001\u0000\u0000\u0000"+
		"\u01d4\u01d7\u0001\u0000\u0000\u0000\u01d5\u01d3\u0001\u0000\u0000\u0000"+
		"\u01d5\u01d6\u0001\u0000\u0000\u0000\u01d6\u000f\u0001\u0000\u0000\u0000"+
		"\u01d7\u01d5\u0001\u0000\u0000\u0000\u01d8\u01db\u0003\u00c0`\u0000\u01d9"+
		"\u01db\u0003\u0012\t\u0000\u01da\u01d8\u0001\u0000\u0000\u0000\u01da\u01d9"+
		"\u0001\u0000\u0000\u0000\u01db\u0011\u0001\u0000\u0000\u0000\u01dc\u01dd"+
		"\u0005\u001b\u0000\u0000\u01dd\u01e1\u0003\u00a2Q\u0000\u01de\u01e0\u0005"+
		"Y\u0000\u0000\u01df\u01de\u0001\u0000\u0000\u0000\u01e0\u01e3\u0001\u0000"+
		"\u0000\u0000\u01e1\u01df\u0001\u0000\u0000\u0000\u01e1\u01e2\u0001\u0000"+
		"\u0000\u0000\u01e2\u01e4\u0001\u0000\u0000\u0000\u01e3\u01e1\u0001\u0000"+
		"\u0000\u0000\u01e4\u01e5\u0003\\.\u0000\u01e5\u0013\u0001\u0000\u0000"+
		"\u0000\u01e6\u01e9\u0003\u0016\u000b\u0000\u01e7\u01e9\u0003\u001a\r\u0000"+
		"\u01e8\u01e6\u0001\u0000\u0000\u0000\u01e8\u01e7\u0001\u0000\u0000\u0000"+
		"\u01e9\u0015\u0001\u0000\u0000\u0000\u01ea\u01eb\u0005\f\u0000\u0000\u01eb"+
		"\u01ef\u0003\u000e\u0007\u0000\u01ec\u01ee\u0005Y\u0000\u0000\u01ed\u01ec"+
		"\u0001\u0000\u0000\u0000\u01ee\u01f1\u0001\u0000\u0000\u0000\u01ef\u01ed"+
		"\u0001\u0000\u0000\u0000\u01ef\u01f0\u0001\u0000\u0000\u0000\u01f0\u01f2"+
		"\u0001\u0000\u0000\u0000\u01f1\u01ef\u0001\u0000\u0000\u0000\u01f2\u01f6"+
		"\u00036\u001b\u0000\u01f3\u01f5\u0005Y\u0000\u0000\u01f4\u01f3\u0001\u0000"+
		"\u0000\u0000\u01f5\u01f8\u0001\u0000\u0000\u0000\u01f6\u01f4\u0001\u0000"+
		"\u0000\u0000\u01f6\u01f7\u0001\u0000\u0000\u0000\u01f7\u01fa\u0001\u0000"+
		"\u0000\u0000\u01f8\u01f6\u0001\u0000\u0000\u0000\u01f9\u01fb\u0003\u0018"+
		"\f\u0000\u01fa\u01f9\u0001\u0000\u0000\u0000\u01fa\u01fb\u0001\u0000\u0000"+
		"\u0000\u01fb\u0017\u0001\u0000\u0000\u0000\u01fc\u0200\u0005\u0007\u0000"+
		"\u0000\u01fd\u01ff\u0005Y\u0000\u0000\u01fe\u01fd\u0001\u0000\u0000\u0000"+
		"\u01ff\u0202\u0001\u0000\u0000\u0000\u0200\u01fe\u0001\u0000\u0000\u0000"+
		"\u0200\u0201\u0001\u0000\u0000\u0000\u0201\u0203\u0001\u0000\u0000\u0000"+
		"\u0202\u0200\u0001\u0000\u0000\u0000\u0203\u020d\u00036\u001b\u0000\u0204"+
		"\u0208\u0005\u0007\u0000\u0000\u0205\u0207\u0005Y\u0000\u0000\u0206\u0205"+
		"\u0001\u0000\u0000\u0000\u0207\u020a\u0001\u0000\u0000\u0000\u0208\u0206"+
		"\u0001\u0000\u0000\u0000\u0208\u0209\u0001\u0000\u0000\u0000\u0209\u020b"+
		"\u0001\u0000\u0000\u0000\u020a\u0208\u0001\u0000\u0000\u0000\u020b\u020d"+
		"\u0003\u0016\u000b\u0000\u020c\u01fc\u0001\u0000\u0000\u0000\u020c\u0204"+
		"\u0001\u0000\u0000\u0000\u020d\u0019\u0001\u0000\u0000\u0000\u020e\u020f"+
		"\u0005\u0011\u0000\u0000\u020f\u0213\u0003\u00c0`\u0000\u0210\u0212\u0005"+
		"Y\u0000\u0000\u0211\u0210\u0001\u0000\u0000\u0000\u0212\u0215\u0001\u0000"+
		"\u0000\u0000\u0213\u0211\u0001\u0000\u0000\u0000\u0213\u0214\u0001\u0000"+
		"\u0000\u0000\u0214\u0216\u0001\u0000\u0000\u0000\u0215\u0213\u0001\u0000"+
		"\u0000\u0000\u0216\u021e\u0005\u001f\u0000\u0000\u0217\u0219\u0005Y\u0000"+
		"\u0000\u0218\u0217\u0001\u0000\u0000\u0000\u0219\u021c\u0001\u0000\u0000"+
		"\u0000\u021a\u0218\u0001\u0000\u0000\u0000\u021a\u021b\u0001\u0000\u0000"+
		"\u0000\u021b\u021d\u0001\u0000\u0000\u0000\u021c\u021a\u0001\u0000\u0000"+
		"\u0000\u021d\u021f\u0003\u001c\u000e\u0000\u021e\u021a\u0001\u0000\u0000"+
		"\u0000\u021e\u021f\u0001\u0000\u0000\u0000\u021f\u0223\u0001\u0000\u0000"+
		"\u0000\u0220\u0222\u0005Y\u0000\u0000\u0221\u0220\u0001\u0000\u0000\u0000"+
		"\u0222\u0225\u0001\u0000\u0000\u0000\u0223\u0221\u0001\u0000\u0000\u0000"+
		"\u0223\u0224\u0001\u0000\u0000\u0000\u0224\u0226\u0001\u0000\u0000\u0000"+
		"\u0225\u0223\u0001\u0000\u0000\u0000\u0226\u0227\u0005\"\u0000\u0000\u0227"+
		"\u001b\u0001\u0000\u0000\u0000\u0228\u0234\u0003\u001e\u000f\u0000\u0229"+
		"\u022d\u0003\u0174\u00ba\u0000\u022a\u022c\u0005Y\u0000\u0000\u022b\u022a"+
		"\u0001\u0000\u0000\u0000\u022c\u022f\u0001\u0000\u0000\u0000\u022d\u022b"+
		"\u0001\u0000\u0000\u0000\u022d\u022e\u0001\u0000\u0000\u0000\u022e\u0230"+
		"\u0001\u0000\u0000\u0000\u022f\u022d\u0001\u0000\u0000\u0000\u0230\u0231"+
		"\u0003\u001e\u000f\u0000\u0231\u0233\u0001\u0000\u0000\u0000\u0232\u0229"+
		"\u0001\u0000\u0000\u0000\u0233\u0236\u0001\u0000\u0000\u0000\u0234\u0232"+
		"\u0001\u0000\u0000\u0000\u0234\u0235\u0001\u0000\u0000\u0000\u0235\u0237"+
		"\u0001\u0000\u0000\u0000\u0236\u0234\u0001\u0000\u0000\u0000\u0237\u0238"+
		"\u0003\u0174\u00ba\u0000\u0238\u001d\u0001\u0000\u0000\u0000\u0239\u023d"+
		"\u0003\u00a2Q\u0000\u023a\u023c\u0005Y\u0000\u0000\u023b\u023a\u0001\u0000"+
		"\u0000\u0000\u023c\u023f\u0001\u0000\u0000\u0000\u023d\u023b\u0001\u0000"+
		"\u0000\u0000\u023d\u023e\u0001\u0000\u0000\u0000\u023e\u0240\u0001\u0000"+
		"\u0000\u0000\u023f\u023d\u0001\u0000\u0000\u0000\u0240\u0244\u0005=\u0000"+
		"\u0000\u0241\u0243\u0005Y\u0000\u0000\u0242\u0241\u0001\u0000\u0000\u0000"+
		"\u0243\u0246\u0001\u0000\u0000\u0000\u0244\u0242\u0001\u0000\u0000\u0000"+
		"\u0244\u0245\u0001\u0000\u0000\u0000\u0245\u0249\u0001\u0000\u0000\u0000"+
		"\u0246\u0244\u0001\u0000\u0000\u0000\u0247\u024a\u00036\u001b\u0000\u0248"+
		"\u024a\u0003\u00c0`\u0000\u0249\u0247\u0001\u0000\u0000\u0000\u0249\u0248"+
		"\u0001\u0000\u0000\u0000\u024a\u001f\u0001\u0000\u0000\u0000\u024b\u024f"+
		"\u0003\"\u0011\u0000\u024c\u024f\u0003$\u0012\u0000\u024d\u024f\u0003"+
		"&\u0013\u0000\u024e\u024b\u0001\u0000\u0000\u0000\u024e\u024c\u0001\u0000"+
		"\u0000\u0000\u024e\u024d\u0001\u0000\u0000\u0000\u024f!\u0001\u0000\u0000"+
		"\u0000\u0250\u0251\u0005\u0004\u0000\u0000\u0251#\u0001\u0000\u0000\u0000"+
		"\u0252\u0253\u0005\u0006\u0000\u0000\u0253%\u0001\u0000\u0000\u0000\u0254"+
		"\u0256\u0005\u0017\u0000\u0000\u0255\u0257\u0003\u00c0`\u0000\u0256\u0255"+
		"\u0001\u0000\u0000\u0000\u0256\u0257\u0001\u0000\u0000\u0000\u0257\'\u0001"+
		"\u0000\u0000\u0000\u0258\u025c\u0005(\u0000\u0000\u0259\u025b\u0005Y\u0000"+
		"\u0000\u025a\u0259\u0001\u0000\u0000\u0000\u025b\u025e\u0001\u0000\u0000"+
		"\u0000\u025c\u025a\u0001\u0000\u0000\u0000\u025c\u025d\u0001\u0000\u0000"+
		"\u0000\u025d\u025f\u0001\u0000\u0000\u0000\u025e\u025c\u0001\u0000\u0000"+
		"\u0000\u025f\u0263\u0003*\u0015\u0000\u0260\u0262\u0005Y\u0000\u0000\u0261"+
		"\u0260\u0001\u0000\u0000\u0000\u0262\u0265\u0001\u0000\u0000\u0000\u0263"+
		"\u0261\u0001\u0000\u0000\u0000\u0263\u0264\u0001\u0000\u0000\u0000\u0264"+
		"\u0266\u0001\u0000\u0000\u0000\u0265\u0263\u0001\u0000\u0000\u0000\u0266"+
		"\u0267\u0005)\u0000\u0000\u0267)\u0001\u0000\u0000\u0000\u0268\u0274\u0003"+
		",\u0016\u0000\u0269\u026d\u0003\u017a\u00bd\u0000\u026a\u026c\u0005Y\u0000"+
		"\u0000\u026b\u026a\u0001\u0000\u0000\u0000\u026c\u026f\u0001\u0000\u0000"+
		"\u0000\u026d\u026b\u0001\u0000\u0000\u0000\u026d\u026e\u0001\u0000\u0000"+
		"\u0000\u026e\u0270\u0001\u0000\u0000\u0000\u026f\u026d\u0001\u0000\u0000"+
		"\u0000\u0270\u0271\u0003,\u0016\u0000\u0271\u0273\u0001\u0000\u0000\u0000"+
		"\u0272\u0269\u0001\u0000\u0000\u0000\u0273\u0276\u0001\u0000\u0000\u0000"+
		"\u0274\u0272\u0001\u0000\u0000\u0000\u0274\u0275\u0001\u0000\u0000\u0000"+
		"\u0275\u0278\u0001\u0000\u0000\u0000\u0276\u0274\u0001\u0000\u0000\u0000"+
		"\u0277\u0279\u0003\u017a\u00bd\u0000\u0278\u0277\u0001\u0000\u0000\u0000"+
		"\u0278\u0279\u0001\u0000\u0000\u0000\u0279+\u0001\u0000\u0000\u0000\u027a"+
		"\u0282\u0003\u0126\u0093\u0000\u027b\u027c\u0003\u0126\u0093\u0000\u027c"+
		"\u027d\u0005A\u0000\u0000\u027d\u0282\u0001\u0000\u0000\u0000\u027e\u027f"+
		"\u0003\u0126\u0093\u0000\u027f\u0280\u0003\u0120\u0090\u0000\u0280\u0282"+
		"\u0001\u0000\u0000\u0000\u0281\u027a\u0001\u0000\u0000\u0000\u0281\u027b"+
		"\u0001\u0000\u0000\u0000\u0281\u027e\u0001\u0000\u0000\u0000\u0282-\u0001"+
		"\u0000\u0000\u0000\u0283\u0287\u0005(\u0000\u0000\u0284\u0286\u0005Y\u0000"+
		"\u0000\u0285\u0284\u0001\u0000\u0000\u0000\u0286\u0289\u0001\u0000\u0000"+
		"\u0000\u0287\u0285\u0001\u0000\u0000\u0000\u0287\u0288\u0001\u0000\u0000"+
		"\u0000\u0288\u028a\u0001\u0000\u0000\u0000\u0289\u0287\u0001\u0000\u0000"+
		"\u0000\u028a\u028e\u00030\u0018\u0000\u028b\u028d\u0005Y\u0000\u0000\u028c"+
		"\u028b\u0001\u0000\u0000\u0000\u028d\u0290\u0001\u0000\u0000\u0000\u028e"+
		"\u028c\u0001\u0000\u0000\u0000\u028e\u028f\u0001\u0000\u0000\u0000\u028f"+
		"\u0291\u0001\u0000\u0000\u0000\u0290\u028e\u0001\u0000\u0000\u0000\u0291"+
		"\u0292\u0005)\u0000\u0000\u0292/\u0001\u0000\u0000\u0000\u0293\u029f\u0003"+
		"2\u0019\u0000\u0294\u0298\u0003\u0176\u00bb\u0000\u0295\u0297\u0005Y\u0000"+
		"\u0000\u0296\u0295\u0001\u0000\u0000\u0000\u0297\u029a\u0001\u0000\u0000"+
		"\u0000\u0298\u0296\u0001\u0000\u0000\u0000\u0298\u0299\u0001\u0000\u0000"+
		"\u0000\u0299\u029b\u0001\u0000\u0000\u0000\u029a\u0298\u0001\u0000\u0000"+
		"\u0000\u029b\u029c\u00032\u0019\u0000\u029c\u029e\u0001\u0000\u0000\u0000"+
		"\u029d\u0294\u0001\u0000\u0000\u0000\u029e\u02a1\u0001\u0000\u0000\u0000"+
		"\u029f\u029d\u0001\u0000\u0000\u0000\u029f\u02a0\u0001\u0000\u0000\u0000"+
		"\u02a01\u0001\u0000\u0000\u0000\u02a1\u029f\u0001\u0000\u0000\u0000\u02a2"+
		"\u02a4\u0003\u011a\u008d\u0000\u02a3\u02a5\u0003\u00be_\u0000\u02a4\u02a3"+
		"\u0001\u0000\u0000\u0000\u02a4\u02a5\u0001\u0000\u0000\u0000\u02a53\u0001"+
		"\u0000\u0000\u0000\u02a6\u02a7\u0003\u0148\u00a4\u0000\u02a7\u02a8\u0005"+
		"Y\u0000\u0000\u02a8\u02aa\u0001\u0000\u0000\u0000\u02a9\u02a6\u0001\u0000"+
		"\u0000\u0000\u02a9\u02aa\u0001\u0000\u0000\u0000\u02aa\u02b2\u0001\u0000"+
		"\u0000\u0000\u02ab\u02af\u0003\u00be_\u0000\u02ac\u02ae\u0005Y\u0000\u0000"+
		"\u02ad\u02ac\u0001\u0000\u0000\u0000\u02ae\u02b1\u0001\u0000\u0000\u0000"+
		"\u02af\u02ad\u0001\u0000\u0000\u0000\u02af\u02b0\u0001\u0000\u0000\u0000"+
		"\u02b0\u02b3\u0001\u0000\u0000\u0000\u02b1\u02af\u0001\u0000\u0000\u0000"+
		"\u02b2\u02ab\u0001\u0000\u0000\u0000\u02b2\u02b3\u0001\u0000\u0000\u0000"+
		"\u02b3\u02bf\u0001\u0000\u0000\u0000\u02b4\u02c0\u00038\u001c\u0000\u02b5"+
		"\u02c0\u0003>\u001f\u0000\u02b6\u02c0\u0003T*\u0000\u02b7\u02c0\u0003"+
		"^/\u0000\u02b8\u02c0\u0003`0\u0000\u02b9\u02c0\u0003f3\u0000\u02ba\u02c0"+
		"\u0003v;\u0000\u02bb\u02c0\u0003\u0080@\u0000\u02bc\u02c0\u0003\u008e"+
		"G\u0000\u02bd\u02c0\u0003\u009cN\u0000\u02be\u02c0\u0003\u009eO\u0000"+
		"\u02bf\u02b4\u0001\u0000\u0000\u0000\u02bf\u02b5\u0001\u0000\u0000\u0000"+
		"\u02bf\u02b6\u0001\u0000\u0000\u0000\u02bf\u02b7\u0001\u0000\u0000\u0000"+
		"\u02bf\u02b8\u0001\u0000\u0000\u0000\u02bf\u02b9\u0001\u0000\u0000\u0000"+
		"\u02bf\u02ba\u0001\u0000\u0000\u0000\u02bf\u02bb\u0001\u0000\u0000\u0000"+
		"\u02bf\u02bc\u0001\u0000\u0000\u0000\u02bf\u02bd\u0001\u0000\u0000\u0000"+
		"\u02bf\u02be\u0001\u0000\u0000\u0000\u02c05\u0001\u0000\u0000\u0000\u02c1"+
		"\u02c9\u0005\u001f\u0000\u0000\u02c2\u02c4\u0005Y\u0000\u0000\u02c3\u02c2"+
		"\u0001\u0000\u0000\u0000\u02c4\u02c7\u0001\u0000\u0000\u0000\u02c5\u02c3"+
		"\u0001\u0000\u0000\u0000\u02c5\u02c6\u0001\u0000\u0000\u0000\u02c6\u02c8"+
		"\u0001\u0000\u0000\u0000\u02c7\u02c5\u0001\u0000\u0000\u0000\u02c8\u02ca"+
		"\u0003\u0006\u0003\u0000\u02c9\u02c5\u0001\u0000\u0000\u0000\u02c9\u02ca"+
		"\u0001\u0000\u0000\u0000\u02ca\u02ce\u0001\u0000\u0000\u0000\u02cb\u02cd"+
		"\u0005Y\u0000\u0000\u02cc\u02cb\u0001\u0000\u0000\u0000\u02cd\u02d0\u0001"+
		"\u0000\u0000\u0000\u02ce\u02cc\u0001\u0000\u0000\u0000\u02ce\u02cf\u0001"+
		"\u0000\u0000\u0000\u02cf\u02d1\u0001\u0000\u0000\u0000\u02d0\u02ce\u0001"+
		"\u0000\u0000\u0000\u02d1\u02d2\u0005\"\u0000\u0000\u02d27\u0001\u0000"+
		"\u0000\u0000\u02d3\u02d4\u0005\u0015\u0000\u0000\u02d4\u02dc\u0003:\u001d"+
		"\u0000\u02d5\u02d7\u0005Y\u0000\u0000\u02d6\u02d5\u0001\u0000\u0000\u0000"+
		"\u02d7\u02da\u0001\u0000\u0000\u0000\u02d8\u02d6\u0001\u0000\u0000\u0000"+
		"\u02d8\u02d9\u0001\u0000\u0000\u0000\u02d9\u02db\u0001\u0000\u0000\u0000"+
		"\u02da\u02d8\u0001\u0000\u0000\u0000\u02db\u02dd\u0003\u00e6s\u0000\u02dc"+
		"\u02d8\u0001\u0000\u0000\u0000\u02dc\u02dd\u0001\u0000\u0000\u0000\u02dd"+
		"9\u0001\u0000\u0000\u0000\u02de\u02e3\u0003<\u001e\u0000\u02df\u02e0\u0005"+
		"\u001e\u0000\u0000\u02e0\u02e2\u0003<\u001e\u0000\u02e1\u02df\u0001\u0000"+
		"\u0000\u0000\u02e2\u02e5\u0001\u0000\u0000\u0000\u02e3\u02e1\u0001\u0000"+
		"\u0000\u0000\u02e3\u02e4\u0001\u0000\u0000\u0000\u02e4;\u0001\u0000\u0000"+
		"\u0000\u02e5\u02e3\u0001\u0000\u0000\u0000\u02e6\u02e7\u0005J\u0000\u0000"+
		"\u02e7=\u0001\u0000\u0000\u0000\u02e8\u02e9\u0005\r\u0000\u0000\u02e9"+
		"\u02ee\u0003@ \u0000\u02ea\u02ef\u0003D\"\u0000\u02eb\u02ef\u0003F#\u0000"+
		"\u02ec\u02ef\u0003H$\u0000\u02ed\u02ef\u0003L&\u0000\u02ee\u02ea\u0001"+
		"\u0000\u0000\u0000\u02ee\u02eb\u0001\u0000\u0000\u0000\u02ee\u02ec\u0001"+
		"\u0000\u0000\u0000\u02ee\u02ed\u0001\u0000\u0000\u0000\u02ee\u02ef\u0001"+
		"\u0000\u0000\u0000\u02ef?\u0001\u0000\u0000\u0000\u02f0\u02f5\u0003B!"+
		"\u0000\u02f1\u02f2\u0005\u001e\u0000\u0000\u02f2\u02f4\u0003B!\u0000\u02f3"+
		"\u02f1\u0001\u0000\u0000\u0000\u02f4\u02f7\u0001\u0000\u0000\u0000\u02f5"+
		"\u02f3\u0001\u0000\u0000\u0000\u02f5\u02f6\u0001\u0000\u0000\u0000\u02f6"+
		"A\u0001\u0000\u0000\u0000\u02f7\u02f5\u0001\u0000\u0000\u0000\u02f8\u02f9"+
		"\u0003\u013c\u009e\u0000\u02f9C\u0001\u0000\u0000\u0000\u02fa\u02fb\u0005"+
		"\u001e\u0000\u0000\u02fb\u02fc\u00053\u0000\u0000\u02fcE\u0001\u0000\u0000"+
		"\u0000\u02fd\u02fe\u0005\u0002\u0000\u0000\u02fe\u02ff\u0003\u013c\u009e"+
		"\u0000\u02ffG\u0001\u0000\u0000\u0000\u0300\u0301\u0005\u001e\u0000\u0000"+
		"\u0301\u0303\u0003\u0126\u0093\u0000\u0302\u0304\u0003J%\u0000\u0303\u0302"+
		"\u0001\u0000\u0000\u0000\u0303\u0304\u0001\u0000\u0000\u0000\u0304I\u0001"+
		"\u0000\u0000\u0000\u0305\u0306\u0005\u0002\u0000\u0000\u0306\u0307\u0003"+
		"\u0126\u0093\u0000\u0307K\u0001\u0000\u0000\u0000\u0308\u0309\u0005\u001e"+
		"\u0000\u0000\u0309\u030d\u0005\u001f\u0000\u0000\u030a\u030c\u0005Y\u0000"+
		"\u0000\u030b\u030a\u0001\u0000\u0000\u0000\u030c\u030f\u0001\u0000\u0000"+
		"\u0000\u030d\u030b\u0001\u0000\u0000\u0000\u030d\u030e\u0001\u0000\u0000"+
		"\u0000\u030e\u0310\u0001\u0000\u0000\u0000\u030f\u030d\u0001\u0000\u0000"+
		"\u0000\u0310\u0314\u0003N\'\u0000\u0311\u0313\u0005Y\u0000\u0000\u0312"+
		"\u0311\u0001\u0000\u0000\u0000\u0313\u0316\u0001\u0000\u0000\u0000\u0314"+
		"\u0312\u0001\u0000\u0000\u0000\u0314\u0315\u0001\u0000\u0000\u0000\u0315"+
		"\u0317\u0001\u0000\u0000\u0000\u0316\u0314\u0001\u0000\u0000\u0000\u0317"+
		"\u0318\u0005\"\u0000\u0000\u0318M\u0001\u0000\u0000\u0000\u0319\u031c"+
		"\u0003P(\u0000\u031a\u031c\u0003R)\u0000\u031b\u0319\u0001\u0000\u0000"+
		"\u0000\u031b\u031a\u0001\u0000\u0000\u0000\u031c\u032a\u0001\u0000\u0000"+
		"\u0000\u031d\u0321\u0003\u0176\u00bb\u0000\u031e\u0320\u0005Y\u0000\u0000"+
		"\u031f\u031e\u0001\u0000\u0000\u0000\u0320\u0323\u0001\u0000\u0000\u0000"+
		"\u0321\u031f\u0001\u0000\u0000\u0000\u0321\u0322\u0001\u0000\u0000\u0000"+
		"\u0322\u0326\u0001\u0000\u0000\u0000\u0323\u0321\u0001\u0000\u0000\u0000"+
		"\u0324\u0327\u0003P(\u0000\u0325\u0327\u0003R)\u0000\u0326\u0324\u0001"+
		"\u0000\u0000\u0000\u0326\u0325\u0001\u0000\u0000\u0000\u0327\u0329\u0001"+
		"\u0000\u0000\u0000\u0328\u031d\u0001\u0000\u0000\u0000\u0329\u032c\u0001"+
		"\u0000\u0000\u0000\u032a\u0328\u0001\u0000\u0000\u0000\u032a\u032b\u0001"+
		"\u0000\u0000\u0000\u032b\u032e\u0001\u0000\u0000\u0000\u032c\u032a\u0001"+
		"\u0000\u0000\u0000\u032d\u032f\u0003\u0176\u00bb\u0000\u032e\u032d\u0001"+
		"\u0000\u0000\u0000\u032e\u032f\u0001\u0000\u0000\u0000\u032fO\u0001\u0000"+
		"\u0000\u0000\u0330\u0332\u0003\u013c\u009e\u0000\u0331\u0333\u0003F#\u0000"+
		"\u0332\u0331\u0001\u0000\u0000\u0000\u0332\u0333\u0001\u0000\u0000\u0000"+
		"\u0333Q\u0001\u0000\u0000\u0000\u0334\u0336\u0003\u0126\u0093\u0000\u0335"+
		"\u0337\u0003J%\u0000\u0336\u0335\u0001\u0000\u0000\u0000\u0336\u0337\u0001"+
		"\u0000\u0000\u0000\u0337S\u0001\u0000\u0000\u0000\u0338\u0339\u0005\u0005"+
		"\u0000\u0000\u0339\u033a\u0003V+\u0000\u033aU\u0001\u0000\u0000\u0000"+
		"\u033b\u0347\u0003Z-\u0000\u033c\u0340\u0003\u0176\u00bb\u0000\u033d\u033f"+
		"\u0005Y\u0000\u0000\u033e\u033d\u0001\u0000\u0000\u0000\u033f\u0342\u0001"+
		"\u0000\u0000\u0000\u0340\u033e\u0001\u0000\u0000\u0000\u0340\u0341\u0001"+
		"\u0000\u0000\u0000\u0341\u0343\u0001\u0000\u0000\u0000\u0342\u0340\u0001"+
		"\u0000\u0000\u0000\u0343\u0344\u0003Z-\u0000\u0344\u0346\u0001\u0000\u0000"+
		"\u0000\u0345\u033c\u0001\u0000\u0000\u0000\u0346\u0349\u0001\u0000\u0000"+
		"\u0000\u0347\u0345\u0001\u0000\u0000\u0000\u0347\u0348\u0001\u0000\u0000"+
		"\u0000\u0348\u036c\u0001\u0000\u0000\u0000\u0349\u0347\u0001\u0000\u0000"+
		"\u0000\u034a\u034e\u0005\u001f\u0000\u0000\u034b\u034d\u0005Y\u0000\u0000"+
		"\u034c\u034b\u0001\u0000\u0000\u0000\u034d\u0350\u0001\u0000\u0000\u0000"+
		"\u034e\u034c\u0001\u0000\u0000\u0000\u034e\u034f\u0001\u0000\u0000\u0000"+
		"\u034f\u0351\u0001\u0000\u0000\u0000\u0350\u034e\u0001\u0000\u0000\u0000"+
		"\u0351\u035d\u0003X,\u0000\u0352\u0356\u0003\u0176\u00bb\u0000\u0353\u0355"+
		"\u0005Y\u0000\u0000\u0354\u0353\u0001\u0000\u0000\u0000\u0355\u0358\u0001"+
		"\u0000\u0000\u0000\u0356\u0354\u0001\u0000\u0000\u0000\u0356\u0357\u0001"+
		"\u0000\u0000\u0000\u0357\u0359\u0001\u0000\u0000\u0000\u0358\u0356\u0001"+
		"\u0000\u0000\u0000\u0359\u035a\u0003X,\u0000\u035a\u035c\u0001\u0000\u0000"+
		"\u0000\u035b\u0352\u0001\u0000\u0000\u0000\u035c\u035f\u0001\u0000\u0000"+
		"\u0000\u035d\u035b\u0001\u0000\u0000\u0000\u035d\u035e\u0001\u0000\u0000"+
		"\u0000\u035e\u0361\u0001\u0000\u0000\u0000\u035f\u035d\u0001\u0000\u0000"+
		"\u0000\u0360\u0362\u0003\u0176\u00bb\u0000\u0361\u0360\u0001\u0000\u0000"+
		"\u0000\u0361\u0362\u0001\u0000\u0000\u0000\u0362\u0366\u0001\u0000\u0000"+
		"\u0000\u0363\u0365\u0005Y\u0000\u0000\u0364\u0363\u0001\u0000\u0000\u0000"+
		"\u0365\u0368\u0001\u0000\u0000\u0000\u0366\u0364\u0001\u0000\u0000\u0000"+
		"\u0366\u0367\u0001\u0000\u0000\u0000\u0367\u0369\u0001\u0000\u0000\u0000"+
		"\u0368\u0366\u0001\u0000\u0000\u0000\u0369\u036a\u0005\"\u0000\u0000\u036a"+
		"\u036c\u0001\u0000\u0000\u0000\u036b\u033b\u0001\u0000\u0000\u0000\u036b"+
		"\u034a\u0001\u0000\u0000\u0000\u036cW\u0001\u0000\u0000\u0000\u036d\u036e"+
		"\u0003\u0148\u00a4\u0000\u036e\u036f\u0005Y\u0000\u0000\u036f\u0371\u0001"+
		"\u0000\u0000\u0000\u0370\u036d\u0001\u0000\u0000\u0000\u0370\u0371\u0001"+
		"\u0000\u0000\u0000\u0371\u0375\u0001\u0000\u0000\u0000\u0372\u0373\u0003"+
		"\u00be_\u0000\u0373\u0374\u0005Y\u0000\u0000\u0374\u0376\u0001\u0000\u0000"+
		"\u0000\u0375\u0372\u0001\u0000\u0000\u0000\u0375\u0376\u0001\u0000\u0000"+
		"\u0000\u0376\u0377\u0001\u0000\u0000\u0000\u0377\u0378\u0003Z-\u0000\u0378"+
		"Y\u0001\u0000\u0000\u0000\u0379\u037b\u0003\u00a2Q\u0000\u037a\u037c\u0003"+
		"\\.\u0000\u037b\u037a\u0001\u0000\u0000\u0000\u037b\u037c\u0001\u0000"+
		"\u0000\u0000\u037c[\u0001\u0000\u0000\u0000\u037d\u0381\u0003\u014c\u00a6"+
		"\u0000\u037e\u0380\u0005Y\u0000\u0000\u037f\u037e\u0001\u0000\u0000\u0000"+
		"\u0380\u0383\u0001\u0000\u0000\u0000\u0381\u037f\u0001\u0000\u0000\u0000"+
		"\u0381\u0382\u0001\u0000\u0000\u0000\u0382\u0384\u0001\u0000\u0000\u0000"+
		"\u0383\u0381\u0001\u0000\u0000\u0000\u0384\u0385\u0003\u00c0`\u0000\u0385"+
		"]\u0001\u0000\u0000\u0000\u0386\u0387\u0005\u001b\u0000\u0000\u0387\u038d"+
		"\u0003V+\u0000\u0388\u0389\u0003\u00a6S\u0000\u0389\u038a\u0005<\u0000"+
		"\u0000\u038a\u038b\u0003\u00c0`\u0000\u038b\u038d\u0001\u0000\u0000\u0000"+
		"\u038c\u0386\u0001\u0000\u0000\u0000\u038c\u0388\u0001\u0000\u0000\u0000"+
		"\u038d_\u0001\u0000\u0000\u0000\u038e\u038f\u0005\u001a\u0000\u0000\u038f"+
		"\u0391\u0003b1\u0000\u0390\u0392\u0003(\u0014\u0000\u0391\u0390\u0001"+
		"\u0000\u0000\u0000\u0391\u0392\u0001\u0000\u0000\u0000\u0392\u0396\u0001"+
		"\u0000\u0000\u0000\u0393\u0395\u0005Y\u0000\u0000\u0394\u0393\u0001\u0000"+
		"\u0000\u0000\u0395\u0398\u0001\u0000\u0000\u0000\u0396\u0394\u0001\u0000"+
		"\u0000\u0000\u0396\u0397\u0001\u0000\u0000\u0000\u0397\u0399\u0001\u0000"+
		"\u0000\u0000\u0398\u0396\u0001\u0000\u0000\u0000\u0399\u039a\u0003d2\u0000"+
		"\u039aa\u0001\u0000\u0000\u0000\u039b\u039c\u0003\u0126\u0093\u0000\u039c"+
		"c\u0001\u0000\u0000\u0000\u039d\u03a1\u0003\u014c\u00a6\u0000\u039e\u03a0"+
		"\u0005Y\u0000\u0000\u039f\u039e\u0001\u0000\u0000\u0000\u03a0\u03a3\u0001"+
		"\u0000\u0000\u0000\u03a1\u039f\u0001\u0000\u0000\u0000\u03a1\u03a2\u0001"+
		"\u0000\u0000\u0000\u03a2\u03a4\u0001\u0000\u0000\u0000\u03a3\u03a1\u0001"+
		"\u0000\u0000\u0000\u03a4\u03a6\u0003\u011a\u008d\u0000\u03a5\u03a7\u0003"+
		"\u00be_\u0000\u03a6\u03a5\u0001\u0000\u0000\u0000\u03a6\u03a7\u0001\u0000"+
		"\u0000\u0000\u03a7\u03a9\u0001\u0000\u0000\u0000\u03a8\u03aa\u0003\u014a"+
		"\u00a5\u0000\u03a9\u03a8\u0001\u0000\u0000\u0000\u03a9\u03aa\u0001\u0000"+
		"\u0000\u0000\u03aae\u0001\u0000\u0000\u0000\u03ab\u03ac\u0005\u000b\u0000"+
		"\u0000\u03ac\u03ae\u0003h4\u0000\u03ad\u03af\u0003(\u0014\u0000\u03ae"+
		"\u03ad\u0001\u0000\u0000\u0000\u03ae\u03af\u0001\u0000\u0000\u0000\u03af"+
		"\u03b0\u0001\u0000\u0000\u0000\u03b0\u03b8\u0003j5\u0000\u03b1\u03b3\u0005"+
		"Y\u0000\u0000\u03b2\u03b1\u0001\u0000\u0000\u0000\u03b3\u03b6\u0001\u0000"+
		"\u0000\u0000\u03b4\u03b2\u0001\u0000\u0000\u0000\u03b4\u03b5\u0001\u0000"+
		"\u0000\u0000\u03b5\u03b7\u0001\u0000\u0000\u0000\u03b6\u03b4\u0001\u0000"+
		"\u0000\u0000\u03b7\u03b9\u0003n7\u0000\u03b8\u03b4\u0001\u0000\u0000\u0000"+
		"\u03b8\u03b9\u0001\u0000\u0000\u0000\u03b9g\u0001\u0000\u0000\u0000\u03ba"+
		"\u03bd\u0003\u013c\u009e\u0000\u03bb\u03bd\u0003\u015c\u00ae\u0000\u03bc"+
		"\u03ba\u0001\u0000\u0000\u0000\u03bc\u03bb\u0001\u0000\u0000\u0000\u03bd"+
		"i\u0001\u0000\u0000\u0000\u03be\u03c0\u0003p8\u0000\u03bf\u03c1\u0003"+
		"\u014a\u00a5\u0000\u03c0\u03bf\u0001\u0000\u0000\u0000\u03c0\u03c1\u0001"+
		"\u0000\u0000\u0000\u03c1\u03c9\u0001\u0000\u0000\u0000\u03c2\u03c4\u0005"+
		"Y\u0000\u0000\u03c3\u03c2\u0001\u0000\u0000\u0000\u03c4\u03c7\u0001\u0000"+
		"\u0000\u0000\u03c5\u03c3\u0001\u0000\u0000\u0000\u03c5\u03c6\u0001\u0000"+
		"\u0000\u0000\u03c6\u03c8\u0001\u0000\u0000\u0000\u03c7\u03c5\u0001\u0000"+
		"\u0000\u0000\u03c8\u03ca\u0003l6\u0000\u03c9\u03c5\u0001\u0000\u0000\u0000"+
		"\u03c9\u03ca\u0001\u0000\u0000\u0000\u03cak\u0001\u0000\u0000\u0000\u03cb"+
		"\u03d0\u0003\u0150\u00a8\u0000\u03cc\u03ce\u0003\u013e\u009f\u0000\u03cd"+
		"\u03cf\u0005&\u0000\u0000\u03ce\u03cd\u0001\u0000\u0000\u0000\u03ce\u03cf"+
		"\u0001\u0000\u0000\u0000\u03cf\u03d1\u0001\u0000\u0000\u0000\u03d0\u03cc"+
		"\u0001\u0000\u0000\u0000\u03d0\u03d1\u0001\u0000\u0000\u0000\u03d1\u03d2"+
		"\u0001\u0000\u0000\u0000\u03d2\u03d4\u0003\u011a\u008d\u0000\u03d3\u03d5"+
		"\u0003\u00be_\u0000\u03d4\u03d3\u0001\u0000\u0000\u0000\u03d4\u03d5\u0001"+
		"\u0000\u0000\u0000\u03d5\u03dd\u0001\u0000\u0000\u0000\u03d6\u03d8\u0005"+
		"Y\u0000\u0000\u03d7\u03d6\u0001\u0000\u0000\u0000\u03d8\u03db\u0001\u0000"+
		"\u0000\u0000\u03d9\u03d7\u0001\u0000\u0000\u0000\u03d9\u03da\u0001\u0000"+
		"\u0000\u0000\u03da\u03dc\u0001\u0000\u0000\u0000\u03db\u03d9\u0001\u0000"+
		"\u0000\u0000\u03dc\u03de\u0003\u014a\u00a5\u0000\u03dd\u03d9\u0001\u0000"+
		"\u0000\u0000\u03dd\u03de\u0001\u0000\u0000\u0000\u03dem\u0001\u0000\u0000"+
		"\u0000\u03df\u03e1\u0005\u001f\u0000\u0000\u03e0\u03e2\u0003\u014a\u00a5"+
		"\u0000\u03e1\u03e0\u0001\u0000\u0000\u0000\u03e1\u03e2\u0001\u0000\u0000"+
		"\u0000\u03e2\u03ea\u0001\u0000\u0000\u0000\u03e3\u03e5\u0005Y\u0000\u0000"+
		"\u03e4\u03e3\u0001\u0000\u0000\u0000\u03e5\u03e8\u0001\u0000\u0000\u0000"+
		"\u03e6\u03e4\u0001\u0000\u0000\u0000\u03e6\u03e7\u0001\u0000\u0000\u0000"+
		"\u03e7\u03e9\u0001\u0000\u0000\u0000\u03e8\u03e6\u0001\u0000\u0000\u0000"+
		"\u03e9\u03eb\u0003\u0006\u0003\u0000\u03ea\u03e6\u0001\u0000\u0000\u0000"+
		"\u03ea\u03eb\u0001\u0000\u0000\u0000\u03eb\u03ef\u0001\u0000\u0000\u0000"+
		"\u03ec\u03ee\u0005Y\u0000\u0000\u03ed\u03ec\u0001\u0000\u0000\u0000\u03ee"+
		"\u03f1\u0001\u0000\u0000\u0000\u03ef\u03ed\u0001\u0000\u0000\u0000\u03ef"+
		"\u03f0\u0001\u0000\u0000\u0000\u03f0\u03f2\u0001\u0000\u0000\u0000\u03f1"+
		"\u03ef\u0001\u0000\u0000\u0000\u03f2\u03f3\u0005\"\u0000\u0000\u03f3o"+
		"\u0001\u0000\u0000\u0000\u03f4\u03f5\u0005 \u0000\u0000\u03f5\u0407\u0005"+
		"#\u0000\u0000\u03f6\u03fa\u0005 \u0000\u0000\u03f7\u03f9\u0005Y\u0000"+
		"\u0000\u03f8\u03f7\u0001\u0000\u0000\u0000\u03f9\u03fc\u0001\u0000\u0000"+
		"\u0000\u03fa\u03f8\u0001\u0000\u0000\u0000\u03fa\u03fb\u0001\u0000\u0000"+
		"\u0000\u03fb\u03fd\u0001\u0000\u0000\u0000\u03fc\u03fa\u0001\u0000\u0000"+
		"\u0000\u03fd\u0401\u0003r9\u0000\u03fe\u0400\u0005Y\u0000\u0000\u03ff"+
		"\u03fe\u0001\u0000\u0000\u0000\u0400\u0403\u0001\u0000\u0000\u0000\u0401"+
		"\u03ff\u0001\u0000\u0000\u0000\u0401\u0402\u0001\u0000\u0000\u0000\u0402"+
		"\u0404\u0001\u0000\u0000\u0000\u0403\u0401\u0001\u0000\u0000\u0000\u0404"+
		"\u0405\u0005#\u0000\u0000\u0405\u0407\u0001\u0000\u0000\u0000\u0406\u03f4"+
		"\u0001\u0000\u0000\u0000\u0406\u03f6\u0001\u0000\u0000\u0000\u0407q\u0001"+
		"\u0000\u0000\u0000\u0408\u0414\u0003t:\u0000\u0409\u040d\u0003\u017a\u00bd"+
		"\u0000\u040a\u040c\u0005Y\u0000\u0000\u040b\u040a\u0001\u0000\u0000\u0000"+
		"\u040c\u040f\u0001\u0000\u0000\u0000\u040d\u040b\u0001\u0000\u0000\u0000"+
		"\u040d\u040e\u0001\u0000\u0000\u0000\u040e\u0410\u0001\u0000\u0000\u0000"+
		"\u040f\u040d\u0001\u0000\u0000\u0000\u0410\u0411\u0003t:\u0000\u0411\u0413"+
		"\u0001\u0000\u0000\u0000\u0412\u0409\u0001\u0000\u0000\u0000\u0413\u0416"+
		"\u0001\u0000\u0000\u0000\u0414\u0412\u0001\u0000\u0000\u0000\u0414\u0415"+
		"\u0001\u0000\u0000\u0000\u0415\u0418\u0001\u0000\u0000\u0000\u0416\u0414"+
		"\u0001\u0000\u0000\u0000\u0417\u0419\u0003\u017a\u00bd\u0000\u0418\u0417"+
		"\u0001\u0000\u0000\u0000\u0418\u0419\u0001\u0000\u0000\u0000\u0419s\u0001"+
		"\u0000\u0000\u0000\u041a\u041b\u0003\u013e\u009f\u0000\u041b\u0423\u0003"+
		"\u0120\u0090\u0000\u041c\u041e\u0005Y\u0000\u0000\u041d\u041c\u0001\u0000"+
		"\u0000\u0000\u041e\u0421\u0001\u0000\u0000\u0000\u041f\u041d\u0001\u0000"+
		"\u0000\u0000\u041f\u0420\u0001\u0000\u0000\u0000\u0420\u0422\u0001\u0000"+
		"\u0000\u0000\u0421\u041f\u0001\u0000\u0000\u0000\u0422\u0424\u0003\\."+
		"\u0000\u0423\u041f\u0001\u0000\u0000\u0000\u0423\u0424\u0001\u0000\u0000"+
		"\u0000\u0424\u042f\u0001\u0000\u0000\u0000\u0425\u0427\u0003\u013e\u009f"+
		"\u0000\u0426\u0428\u0005&\u0000\u0000\u0427\u0426\u0001\u0000\u0000\u0000"+
		"\u0427\u0428\u0001\u0000\u0000\u0000\u0428\u0429\u0001\u0000\u0000\u0000"+
		"\u0429\u042a\u0003\u011a\u008d\u0000\u042a\u042c\u0005A\u0000\u0000\u042b"+
		"\u042d\u0003\u00be_\u0000\u042c\u042b\u0001\u0000\u0000\u0000\u042c\u042d"+
		"\u0001\u0000\u0000\u0000\u042d\u042f\u0001\u0000\u0000\u0000\u042e\u041a"+
		"\u0001\u0000\u0000\u0000\u042e\u0425\u0001\u0000\u0000\u0000\u042fu\u0001"+
		"\u0000\u0000\u0000\u0430\u0431\u0005\b\u0000\u0000\u0431\u0433\u0003z"+
		"=\u0000\u0432\u0434\u0003(\u0014\u0000\u0433\u0432\u0001\u0000\u0000\u0000"+
		"\u0433\u0434\u0001\u0000\u0000\u0000\u0434\u043c\u0001\u0000\u0000\u0000"+
		"\u0435\u0437\u0005Y\u0000\u0000\u0436\u0435\u0001\u0000\u0000\u0000\u0437"+
		"\u043a\u0001\u0000\u0000\u0000\u0438\u0436\u0001\u0000\u0000\u0000\u0438"+
		"\u0439\u0001\u0000\u0000\u0000\u0439\u043b\u0001\u0000\u0000\u0000\u043a"+
		"\u0438\u0001\u0000\u0000\u0000\u043b\u043d\u0003\u0136\u009b\u0000\u043c"+
		"\u0438\u0001\u0000\u0000\u0000\u043c\u043d\u0001\u0000\u0000\u0000\u043d"+
		"\u0441\u0001\u0000\u0000\u0000\u043e\u0440\u0005Y\u0000\u0000\u043f\u043e"+
		"\u0001\u0000\u0000\u0000\u0440\u0443\u0001\u0000\u0000\u0000\u0441\u043f"+
		"\u0001\u0000\u0000\u0000\u0441\u0442\u0001\u0000\u0000\u0000\u0442\u0444"+
		"\u0001\u0000\u0000\u0000\u0443\u0441\u0001\u0000\u0000\u0000\u0444\u0445"+
		"\u0003x<\u0000\u0445w\u0001\u0000\u0000\u0000\u0446\u0448\u0005\u001f"+
		"\u0000\u0000\u0447\u0449\u0003\u014a\u00a5\u0000\u0448\u0447\u0001\u0000"+
		"\u0000\u0000\u0448\u0449\u0001\u0000\u0000\u0000\u0449\u0451\u0001\u0000"+
		"\u0000\u0000\u044a\u044c\u0005Y\u0000\u0000\u044b\u044a\u0001\u0000\u0000"+
		"\u0000\u044c\u044f\u0001\u0000\u0000\u0000\u044d\u044b\u0001\u0000\u0000"+
		"\u0000\u044d\u044e\u0001\u0000\u0000\u0000\u044e\u0450\u0001\u0000\u0000"+
		"\u0000\u044f\u044d\u0001\u0000\u0000\u0000\u0450\u0452\u0003|>\u0000\u0451"+
		"\u044d\u0001\u0000\u0000\u0000\u0451\u0452\u0001\u0000\u0000\u0000\u0452"+
		"\u0456\u0001\u0000\u0000\u0000\u0453\u0455\u0005Y\u0000\u0000\u0454\u0453"+
		"\u0001\u0000\u0000\u0000\u0455\u0458\u0001\u0000\u0000\u0000\u0456\u0454"+
		"\u0001\u0000\u0000\u0000\u0456\u0457\u0001\u0000\u0000\u0000\u0457\u0459"+
		"\u0001\u0000\u0000\u0000\u0458\u0456\u0001\u0000\u0000\u0000\u0459\u045a"+
		"\u0005\"\u0000\u0000\u045ay\u0001\u0000\u0000\u0000\u045b\u045c\u0003"+
		"\u0126\u0093\u0000\u045c{\u0001\u0000\u0000\u0000\u045d\u0469\u0003~?"+
		"\u0000\u045e\u0462\u0003\u017a\u00bd\u0000\u045f\u0461\u0005Y\u0000\u0000"+
		"\u0460\u045f\u0001\u0000\u0000\u0000\u0461\u0464\u0001\u0000\u0000\u0000"+
		"\u0462\u0460\u0001\u0000\u0000\u0000\u0462\u0463\u0001\u0000\u0000\u0000"+
		"\u0463\u0465\u0001\u0000\u0000\u0000\u0464\u0462\u0001\u0000\u0000\u0000"+
		"\u0465\u0466\u0003~?\u0000\u0466\u0468\u0001\u0000\u0000\u0000\u0467\u045e"+
		"\u0001\u0000\u0000\u0000\u0468\u046b\u0001\u0000\u0000\u0000\u0469\u0467"+
		"\u0001\u0000\u0000\u0000\u0469\u046a\u0001\u0000\u0000\u0000\u046a\u046d"+
		"\u0001\u0000\u0000\u0000\u046b\u0469\u0001\u0000\u0000\u0000\u046c\u046e"+
		"\u0003\u017a\u00bd\u0000\u046d\u046c\u0001\u0000\u0000\u0000\u046d\u046e"+
		"\u0001\u0000\u0000\u0000\u046e}\u0001\u0000\u0000\u0000\u046f\u0470\u0003"+
		"\u0148\u00a4\u0000\u0470\u0471\u0005Y\u0000\u0000\u0471\u0473\u0001\u0000"+
		"\u0000\u0000\u0472\u046f\u0001\u0000\u0000\u0000\u0472\u0473\u0001\u0000"+
		"\u0000\u0000\u0473\u0477\u0001\u0000\u0000\u0000\u0474\u0475\u0003\u00be"+
		"_\u0000\u0475\u0476\u0005Y\u0000\u0000\u0476\u0478\u0001\u0000\u0000\u0000"+
		"\u0477\u0474\u0001\u0000\u0000\u0000\u0477\u0478\u0001\u0000\u0000\u0000"+
		"\u0478\u0479\u0001\u0000\u0000\u0000\u0479\u047b\u0003\u013c\u009e\u0000"+
		"\u047a\u047c\u0003\u00be_\u0000\u047b\u047a\u0001\u0000\u0000\u0000\u047b"+
		"\u047c\u0001\u0000\u0000\u0000\u047c\u0484\u0001\u0000\u0000\u0000\u047d"+
		"\u047f\u0005Y\u0000\u0000\u047e\u047d\u0001\u0000\u0000\u0000\u047f\u0482"+
		"\u0001\u0000\u0000\u0000\u0480\u047e\u0001\u0000\u0000\u0000\u0480\u0481"+
		"\u0001\u0000\u0000\u0000\u0481\u0483\u0001\u0000\u0000\u0000\u0482\u0480"+
		"\u0001\u0000\u0000\u0000\u0483\u0485\u0003\\.\u0000\u0484\u0480\u0001"+
		"\u0000\u0000\u0000\u0484\u0485\u0001\u0000\u0000\u0000\u0485\u0488\u0001"+
		"\u0000\u0000\u0000\u0486\u0488\u0003\u0004\u0002\u0000\u0487\u0472\u0001"+
		"\u0000\u0000\u0000\u0487\u0486\u0001\u0000\u0000\u0000\u0488\u007f\u0001"+
		"\u0000\u0000\u0000\u0489\u048a\u0005\u001a\u0000\u0000\u048a\u048c\u0003"+
		"\u0082A\u0000\u048b\u048d\u0003(\u0014\u0000\u048c\u048b\u0001\u0000\u0000"+
		"\u0000\u048c\u048d\u0001\u0000\u0000\u0000\u048d\u048e\u0001\u0000\u0000"+
		"\u0000\u048e\u048f\u0003\u0084B\u0000\u048f\u0081\u0001\u0000\u0000\u0000"+
		"\u0490\u0491\u0003\u0126\u0093\u0000\u0491\u0083\u0001\u0000\u0000\u0000"+
		"\u0492\u0494\u0005Y\u0000\u0000\u0493\u0492\u0001\u0000\u0000\u0000\u0494"+
		"\u0497\u0001\u0000\u0000\u0000\u0495\u0493\u0001\u0000\u0000\u0000\u0495"+
		"\u0496\u0001\u0000\u0000\u0000\u0496\u0498\u0001\u0000\u0000\u0000\u0497"+
		"\u0495\u0001\u0000\u0000\u0000\u0498\u049a\u0003\u0136\u009b\u0000\u0499"+
		"\u0495\u0001\u0000\u0000\u0000\u0499\u049a\u0001\u0000\u0000\u0000\u049a"+
		"\u04a2\u0001\u0000\u0000\u0000\u049b\u049d\u0005Y\u0000\u0000\u049c\u049b"+
		"\u0001\u0000\u0000\u0000\u049d\u04a0\u0001\u0000\u0000\u0000\u049e\u049c"+
		"\u0001\u0000\u0000\u0000\u049e\u049f\u0001\u0000\u0000\u0000\u049f\u04a1"+
		"\u0001\u0000\u0000\u0000\u04a0\u049e\u0001\u0000\u0000\u0000\u04a1\u04a3"+
		"\u0003\u0086C\u0000\u04a2\u049e\u0001\u0000\u0000\u0000\u04a2\u04a3\u0001"+
		"\u0000\u0000\u0000\u04a3\u0085\u0001\u0000\u0000\u0000\u04a4\u04a8\u0005"+
		"\u001f\u0000\u0000\u04a5\u04a6\u0003\u014a\u00a5\u0000\u04a6\u04a7\u0005"+
		"Y\u0000\u0000\u04a7\u04a9\u0001\u0000\u0000\u0000\u04a8\u04a5\u0001\u0000"+
		"\u0000\u0000\u04a8\u04a9\u0001\u0000\u0000\u0000\u04a9\u04b1\u0001\u0000"+
		"\u0000\u0000\u04aa\u04ac\u0005Y\u0000\u0000\u04ab\u04aa\u0001\u0000\u0000"+
		"\u0000\u04ac\u04af\u0001\u0000\u0000\u0000\u04ad\u04ab\u0001\u0000\u0000"+
		"\u0000\u04ad\u04ae\u0001\u0000\u0000\u0000\u04ae\u04b0\u0001\u0000\u0000"+
		"\u0000\u04af\u04ad\u0001\u0000\u0000\u0000\u04b0\u04b2\u0003\u0088D\u0000"+
		"\u04b1\u04ad\u0001\u0000\u0000\u0000\u04b1\u04b2\u0001\u0000\u0000\u0000"+
		"\u04b2\u04b6\u0001\u0000\u0000\u0000\u04b3\u04b5\u0005Y\u0000\u0000\u04b4"+
		"\u04b3\u0001\u0000\u0000\u0000\u04b5\u04b8\u0001\u0000\u0000\u0000\u04b6"+
		"\u04b4\u0001\u0000\u0000\u0000\u04b6\u04b7\u0001\u0000\u0000\u0000\u04b7"+
		"\u04b9\u0001\u0000\u0000\u0000\u04b8\u04b6\u0001\u0000\u0000\u0000\u04b9"+
		"\u04ba\u0005\"\u0000\u0000\u04ba\u0087\u0001\u0000\u0000\u0000\u04bb\u04c7"+
		"\u0003\u008aE\u0000\u04bc\u04c0\u0003\u0178\u00bc\u0000\u04bd\u04bf\u0005"+
		"Y\u0000\u0000\u04be\u04bd\u0001\u0000\u0000\u0000\u04bf\u04c2\u0001\u0000"+
		"\u0000\u0000\u04c0\u04be\u0001\u0000\u0000\u0000\u04c0\u04c1\u0001\u0000"+
		"\u0000\u0000\u04c1\u04c3\u0001\u0000\u0000\u0000\u04c2\u04c0\u0001\u0000"+
		"\u0000\u0000\u04c3\u04c4\u0003\u008aE\u0000\u04c4\u04c6\u0001\u0000\u0000"+
		"\u0000\u04c5\u04bc\u0001\u0000\u0000\u0000\u04c6\u04c9\u0001\u0000\u0000"+
		"\u0000\u04c7\u04c5\u0001\u0000\u0000\u0000\u04c7\u04c8\u0001\u0000\u0000"+
		"\u0000\u04c8\u04cb\u0001\u0000\u0000\u0000\u04c9\u04c7\u0001\u0000\u0000"+
		"\u0000\u04ca\u04cc\u0003\u0178\u00bc\u0000\u04cb\u04ca\u0001\u0000\u0000"+
		"\u0000\u04cb\u04cc\u0001\u0000\u0000\u0000\u04cc\u0089\u0001\u0000\u0000"+
		"\u0000\u04cd\u04ce\u0003\u0148\u00a4\u0000\u04ce\u04cf\u0005Y\u0000\u0000"+
		"\u04cf\u04d1\u0001\u0000\u0000\u0000\u04d0\u04cd\u0001\u0000\u0000\u0000"+
		"\u04d0\u04d1\u0001\u0000\u0000\u0000\u04d1\u04d5\u0001\u0000\u0000\u0000"+
		"\u04d2\u04d3\u0003\u00be_\u0000\u04d3\u04d4\u0005Y\u0000\u0000\u04d4\u04d6"+
		"\u0001\u0000\u0000\u0000\u04d5\u04d2\u0001\u0000\u0000\u0000\u04d5\u04d6"+
		"\u0001\u0000\u0000\u0000\u04d6\u04dc\u0001\u0000\u0000\u0000\u04d7\u04dd"+
		"\u0003\u0080@\u0000\u04d8\u04dd\u0003v;\u0000\u04d9\u04dd\u0003T*\u0000"+
		"\u04da\u04dd\u0003`0\u0000\u04db\u04dd\u0003\u008cF\u0000\u04dc\u04d7"+
		"\u0001\u0000\u0000\u0000\u04dc\u04d8\u0001\u0000\u0000\u0000\u04dc\u04d9"+
		"\u0001\u0000\u0000\u0000\u04dc\u04da\u0001\u0000\u0000\u0000\u04dc\u04db"+
		"\u0001\u0000\u0000\u0000\u04dd\u04e0\u0001\u0000\u0000\u0000\u04de\u04e0"+
		"\u0003\u0004\u0002\u0000\u04df\u04d0\u0001\u0000\u0000\u0000\u04df\u04de"+
		"\u0001\u0000\u0000\u0000\u04e0\u008b\u0001\u0000\u0000\u0000\u04e1\u04e2"+
		"\u0003\u013c\u009e\u0000\u04e2\u04ea\u0003\u0120\u0090\u0000\u04e3\u04e5"+
		"\u0005Y\u0000\u0000\u04e4\u04e3\u0001\u0000\u0000\u0000\u04e5\u04e8\u0001"+
		"\u0000\u0000\u0000\u04e6\u04e4\u0001\u0000\u0000\u0000\u04e6\u04e7\u0001"+
		"\u0000\u0000\u0000\u04e7\u04e9\u0001\u0000\u0000\u0000\u04e8\u04e6\u0001"+
		"\u0000\u0000\u0000\u04e9\u04eb\u0003\\.\u0000\u04ea\u04e6\u0001\u0000"+
		"\u0000\u0000\u04ea\u04eb\u0001\u0000\u0000\u0000\u04eb\u008d\u0001\u0000"+
		"\u0000\u0000\u04ec\u04ed\u0005\u000f\u0000\u0000\u04ed\u04ef\u0003\u0090"+
		"H\u0000\u04ee\u04f0\u0003(\u0014\u0000\u04ef\u04ee\u0001\u0000\u0000\u0000"+
		"\u04ef\u04f0\u0001\u0000\u0000\u0000\u04f0\u04f1\u0001\u0000\u0000\u0000"+
		"\u04f1\u04f2\u0003\u0092I\u0000\u04f2\u008f\u0001\u0000\u0000\u0000\u04f3"+
		"\u04f4\u0003\u0126\u0093\u0000\u04f4\u0091\u0001\u0000\u0000\u0000\u04f5"+
		"\u04f7\u0005Y\u0000\u0000\u04f6\u04f5\u0001\u0000\u0000\u0000\u04f7\u04fa"+
		"\u0001\u0000\u0000\u0000\u04f8\u04f6\u0001\u0000\u0000\u0000\u04f8\u04f9"+
		"\u0001\u0000\u0000\u0000\u04f9\u04fb\u0001\u0000\u0000\u0000\u04fa\u04f8"+
		"\u0001\u0000\u0000\u0000\u04fb\u04fd\u0003\u0136\u009b\u0000\u04fc\u04f8"+
		"\u0001\u0000\u0000\u0000\u04fc\u04fd\u0001\u0000\u0000\u0000\u04fd\u0501"+
		"\u0001\u0000\u0000\u0000\u04fe\u0500\u0005Y\u0000\u0000\u04ff\u04fe\u0001"+
		"\u0000\u0000\u0000\u0500\u0503\u0001\u0000\u0000\u0000\u0501\u04ff\u0001"+
		"\u0000\u0000\u0000\u0501\u0502\u0001\u0000\u0000\u0000\u0502\u0504\u0001"+
		"\u0000\u0000\u0000\u0503\u0501\u0001\u0000\u0000\u0000\u0504\u0505\u0003"+
		"\u0094J\u0000\u0505\u0093\u0001\u0000\u0000\u0000\u0506\u0508\u0005\u001f"+
		"\u0000\u0000\u0507\u0509\u0003\u014a\u00a5\u0000\u0508\u0507\u0001\u0000"+
		"\u0000\u0000\u0508\u0509\u0001\u0000\u0000\u0000\u0509\u0511\u0001\u0000"+
		"\u0000\u0000\u050a\u050c\u0005Y\u0000\u0000\u050b\u050a\u0001\u0000\u0000"+
		"\u0000\u050c\u050f\u0001\u0000\u0000\u0000\u050d\u050b\u0001\u0000\u0000"+
		"\u0000\u050d\u050e\u0001\u0000\u0000\u0000\u050e\u0510\u0001\u0000\u0000"+
		"\u0000\u050f\u050d\u0001\u0000\u0000\u0000\u0510\u0512\u0003\u0096K\u0000"+
		"\u0511\u050d\u0001\u0000\u0000\u0000\u0511\u0512\u0001\u0000\u0000\u0000"+
		"\u0512\u0516\u0001\u0000\u0000\u0000\u0513\u0515\u0005Y\u0000\u0000\u0514"+
		"\u0513\u0001\u0000\u0000\u0000\u0515\u0518\u0001\u0000\u0000\u0000\u0516"+
		"\u0514\u0001\u0000\u0000\u0000\u0516\u0517\u0001\u0000\u0000\u0000\u0517"+
		"\u0519\u0001\u0000\u0000\u0000\u0518\u0516\u0001\u0000\u0000\u0000\u0519"+
		"\u051a\u0005\"\u0000\u0000\u051a\u0095\u0001\u0000\u0000\u0000\u051b\u0527"+
		"\u0003\u0098L\u0000\u051c\u0520\u0003\u0178\u00bc\u0000\u051d\u051f\u0005"+
		"Y\u0000\u0000\u051e\u051d\u0001\u0000\u0000\u0000\u051f\u0522\u0001\u0000"+
		"\u0000\u0000\u0520\u051e\u0001\u0000\u0000\u0000\u0520\u0521\u0001\u0000"+
		"\u0000\u0000\u0521\u0523\u0001\u0000\u0000\u0000\u0522\u0520\u0001\u0000"+
		"\u0000\u0000\u0523\u0524\u0003\u0098L\u0000\u0524\u0526\u0001\u0000\u0000"+
		"\u0000\u0525\u051c\u0001\u0000\u0000\u0000\u0526\u0529\u0001\u0000\u0000"+
		"\u0000\u0527\u0525\u0001\u0000\u0000\u0000\u0527\u0528\u0001\u0000\u0000"+
		"\u0000\u0528\u052b\u0001\u0000\u0000\u0000\u0529\u0527\u0001\u0000\u0000"+
		"\u0000\u052a\u052c\u0003\u0178\u00bc\u0000\u052b\u052a\u0001\u0000\u0000"+
		"\u0000\u052b\u052c\u0001\u0000\u0000\u0000\u052c\u0097\u0001\u0000\u0000"+
		"\u0000\u052d\u052e\u0003\u0148\u00a4\u0000\u052e\u052f\u0005Y\u0000\u0000"+
		"\u052f\u0531\u0001\u0000\u0000\u0000\u0530\u052d\u0001\u0000\u0000\u0000"+
		"\u0530\u0531\u0001\u0000\u0000\u0000\u0531\u0535\u0001\u0000\u0000\u0000"+
		"\u0532\u0533\u0003\u00be_\u0000\u0533\u0534\u0005Y\u0000\u0000\u0534\u0536"+
		"\u0001\u0000\u0000\u0000\u0535\u0532\u0001\u0000\u0000\u0000\u0535\u0536"+
		"\u0001\u0000\u0000\u0000\u0536\u0539\u0001\u0000\u0000\u0000\u0537\u053a"+
		"\u0003`0\u0000\u0538\u053a\u0003\u009aM\u0000\u0539\u0537\u0001\u0000"+
		"\u0000\u0000\u0539\u0538\u0001\u0000\u0000\u0000\u053a\u053d\u0001\u0000"+
		"\u0000\u0000\u053b\u053d\u0003\u0004\u0002\u0000\u053c\u0530\u0001\u0000"+
		"\u0000\u0000\u053c\u053b\u0001\u0000\u0000\u0000\u053d\u0099\u0001\u0000"+
		"\u0000\u0000\u053e\u0540\u0003h4\u0000\u053f\u0541\u0003(\u0014\u0000"+
		"\u0540\u053f\u0001\u0000\u0000\u0000\u0540\u0541\u0001\u0000\u0000\u0000"+
		"\u0541\u0545\u0001\u0000\u0000\u0000\u0542\u0544\u0005Y\u0000\u0000\u0543"+
		"\u0542\u0001\u0000\u0000\u0000\u0544\u0547\u0001\u0000\u0000\u0000\u0545"+
		"\u0543\u0001\u0000\u0000\u0000\u0545\u0546\u0001\u0000\u0000\u0000\u0546"+
		"\u0548\u0001\u0000\u0000\u0000\u0547\u0545\u0001\u0000\u0000\u0000\u0548"+
		"\u0549\u0003j5\u0000\u0549\u009b\u0001\u0000\u0000\u0000\u054a\u054b\u0005"+
		"\u0003\u0000\u0000\u054b\u054d\u0003\u00b6[\u0000\u054c\u054e\u0003(\u0014"+
		"\u0000\u054d\u054c\u0001\u0000\u0000\u0000\u054d\u054e\u0001\u0000\u0000"+
		"\u0000\u054e\u055d\u0001\u0000\u0000\u0000\u054f\u055e\u0003\u0086C\u0000"+
		"\u0550\u0558\u0003\u0120\u0090\u0000\u0551\u0553\u0005Y\u0000\u0000\u0552"+
		"\u0551\u0001\u0000\u0000\u0000\u0553\u0556\u0001\u0000\u0000\u0000\u0554"+
		"\u0552\u0001\u0000\u0000\u0000\u0554\u0555\u0001\u0000\u0000\u0000\u0555"+
		"\u0557\u0001\u0000\u0000\u0000\u0556\u0554\u0001\u0000\u0000\u0000\u0557"+
		"\u0559\u0003\\.\u0000\u0558\u0554\u0001\u0000\u0000\u0000\u0558\u0559"+
		"\u0001\u0000\u0000\u0000\u0559\u055b\u0001\u0000\u0000\u0000\u055a\u055c"+
		"\u0003\u014a\u00a5\u0000\u055b\u055a\u0001\u0000\u0000\u0000\u055b\u055c"+
		"\u0001\u0000\u0000\u0000\u055c\u055e\u0001\u0000\u0000\u0000\u055d\u054f"+
		"\u0001\u0000\u0000\u0000\u055d\u0550\u0001\u0000\u0000\u0000\u055e\u009d"+
		"\u0001\u0000\u0000\u0000\u055f\u0560\u0005\u0003\u0000\u0000\u0560\u0562"+
		"\u0003\u00b6[\u0000\u0561\u0563\u0003(\u0014\u0000\u0562\u0561\u0001\u0000"+
		"\u0000\u0000\u0562\u0563\u0001\u0000\u0000\u0000\u0563\u0567\u0001\u0000"+
		"\u0000\u0000\u0564\u0566\u0005Y\u0000\u0000\u0565\u0564\u0001\u0000\u0000"+
		"\u0000\u0566\u0569\u0001\u0000\u0000\u0000\u0567\u0565\u0001\u0000\u0000"+
		"\u0000\u0567\u0568\u0001\u0000\u0000\u0000\u0568\u056a\u0001\u0000\u0000"+
		"\u0000\u0569\u0567\u0001\u0000\u0000\u0000\u056a\u056b\u0003\u00a0P\u0000"+
		"\u056b\u009f\u0001\u0000\u0000\u0000\u056c\u0570\u0003\u014c\u00a6\u0000"+
		"\u056d\u056f\u0005Y\u0000\u0000\u056e\u056d\u0001\u0000\u0000\u0000\u056f"+
		"\u0572\u0001\u0000\u0000\u0000\u0570\u056e\u0001\u0000\u0000\u0000\u0570"+
		"\u0571\u0001\u0000\u0000\u0000\u0571\u0576\u0001\u0000\u0000\u0000\u0572"+
		"\u0570\u0001\u0000\u0000\u0000\u0573\u0574\u0003:\u001d\u0000\u0574\u0575"+
		"\u0005\u001e\u0000\u0000\u0575\u0577\u0001\u0000\u0000\u0000\u0576\u0573"+
		"\u0001\u0000\u0000\u0000\u0576\u0577\u0001\u0000\u0000\u0000\u0577\u0578"+
		"\u0001\u0000\u0000\u0000\u0578\u057a\u0003\u00b6[\u0000\u0579\u057b\u0003"+
		".\u0017\u0000\u057a\u0579\u0001\u0000\u0000\u0000\u057a\u057b\u0001\u0000"+
		"\u0000\u0000\u057b\u057d\u0001\u0000\u0000\u0000\u057c\u057e\u0003\u014a"+
		"\u00a5\u0000\u057d\u057c\u0001\u0000\u0000\u0000\u057d\u057e\u0001\u0000"+
		"\u0000\u0000\u057e\u00a1\u0001\u0000\u0000\u0000\u057f\u0580\u0006Q\uffff"+
		"\uffff\u0000\u0580\u0582\u0003\u00a4R\u0000\u0581\u0583\u0003\u0120\u0090"+
		"\u0000\u0582\u0581\u0001\u0000\u0000\u0000\u0582\u0583\u0001\u0000\u0000"+
		"\u0000\u0583\u0591\u0001\u0000\u0000\u0000\u0584\u0586\u0003\u00a6S\u0000"+
		"\u0585\u0587\u0003\u0120\u0090\u0000\u0586\u0585\u0001\u0000\u0000\u0000"+
		"\u0586\u0587\u0001\u0000\u0000\u0000\u0587\u0591\u0001\u0000\u0000\u0000"+
		"\u0588\u058a\u0003\u00a8T\u0000\u0589\u058b\u0003\u0120\u0090\u0000\u058a"+
		"\u0589\u0001\u0000\u0000\u0000\u058a\u058b\u0001\u0000\u0000\u0000\u058b"+
		"\u0591\u0001\u0000\u0000\u0000\u058c\u0591\u0003\u00aeW\u0000\u058d\u058e"+
		"\u0005\u0010\u0000\u0000\u058e\u0591\u0003\u011a\u008d\u0000\u058f\u0591"+
		"\u0003\u00b0X\u0000\u0590\u057f\u0001\u0000\u0000\u0000\u0590\u0584\u0001"+
		"\u0000\u0000\u0000\u0590\u0588\u0001\u0000\u0000\u0000\u0590\u058c\u0001"+
		"\u0000\u0000\u0000\u0590\u058d\u0001\u0000\u0000\u0000\u0590\u058f\u0001"+
		"\u0000\u0000\u0000\u0591\u0597\u0001\u0000\u0000\u0000\u0592\u0593\n\u0002"+
		"\u0000\u0000\u0593\u0594\u0005\u0002\u0000\u0000\u0594\u0596\u0003\u011a"+
		"\u008d\u0000\u0595\u0592\u0001\u0000\u0000\u0000\u0596\u0599\u0001\u0000"+
		"\u0000\u0000\u0597\u0595\u0001\u0000\u0000\u0000\u0597\u0598\u0001\u0000"+
		"\u0000\u0000\u0598\u00a3\u0001\u0000\u0000\u0000\u0599\u0597\u0001\u0000"+
		"\u0000\u0000\u059a\u059b\u00059\u0000\u0000\u059b\u00a5\u0001\u0000\u0000"+
		"\u0000\u059c\u059d\u0003\u013c\u009e\u0000\u059d\u00a7\u0001\u0000\u0000"+
		"\u0000\u059e\u05a0\u0005 \u0000\u0000\u059f\u05a1\u0003\u00aaU\u0000\u05a0"+
		"\u059f\u0001\u0000\u0000\u0000\u05a0\u05a1\u0001\u0000\u0000\u0000\u05a1"+
		"\u05a2\u0001\u0000\u0000\u0000\u05a2\u05a3\u0005#\u0000\u0000\u05a3\u00a9"+
		"\u0001\u0000\u0000\u0000\u05a4\u05a9\u0003\u00acV\u0000\u05a5\u05a6\u0005"+
		"%\u0000\u0000\u05a6\u05a8\u0003\u00acV\u0000\u05a7\u05a5\u0001\u0000\u0000"+
		"\u0000\u05a8\u05ab\u0001\u0000\u0000\u0000\u05a9\u05a7\u0001\u0000\u0000"+
		"\u0000\u05a9\u05aa\u0001\u0000\u0000\u0000\u05aa\u00ab\u0001\u0000\u0000"+
		"\u0000\u05ab\u05a9\u0001\u0000\u0000\u0000\u05ac\u05ad\u0003\u00a2Q\u0000"+
		"\u05ad\u00ad\u0001\u0000\u0000\u0000\u05ae\u05af\u0003\u00a6S\u0000\u05af"+
		"\u05b0\u0005+\u0000\u0000\u05b0\u00af\u0001\u0000\u0000\u0000\u05b1\u05b2"+
		"\u0003\u00c0`\u0000\u05b2\u00b1\u0001\u0000\u0000\u0000\u05b3\u05b4\u0005"+
		",\u0000\u0000\u05b4\u05be\u0005O\u0000\u0000\u05b5\u05b6\u0005,\u0000"+
		"\u0000\u05b6\u05b8\u0003\u00b4Z\u0000\u05b7\u05b9\u0003.\u0017\u0000\u05b8"+
		"\u05b7\u0001\u0000\u0000\u0000\u05b8\u05b9\u0001\u0000\u0000\u0000\u05b9"+
		"\u05bb\u0001\u0000\u0000\u0000\u05ba\u05bc\u0003\u00b8\\\u0000\u05bb\u05ba"+
		"\u0001\u0000\u0000\u0000\u05bb\u05bc\u0001\u0000\u0000\u0000\u05bc\u05be"+
		"\u0001\u0000\u0000\u0000\u05bd\u05b3\u0001\u0000\u0000\u0000\u05bd\u05b5"+
		"\u0001\u0000\u0000\u0000\u05be\u00b3\u0001\u0000\u0000\u0000\u05bf\u05c0"+
		"\u0003:\u001d\u0000\u05c0\u05c1\u0005\u001e\u0000\u0000\u05c1\u05c3\u0001"+
		"\u0000\u0000\u0000\u05c2\u05bf\u0001\u0000\u0000\u0000\u05c2\u05c3\u0001"+
		"\u0000\u0000\u0000\u05c3\u05c4\u0001\u0000\u0000\u0000\u05c4\u05c5\u0003"+
		"\u00b6[\u0000\u05c5\u00b5\u0001\u0000\u0000\u0000\u05c6\u05c7\u0003\u013e"+
		"\u009f\u0000\u05c7\u00b7\u0001\u0000\u0000\u0000\u05c8\u05d0\u0005 \u0000"+
		"\u0000\u05c9\u05cb\u0005Y\u0000\u0000\u05ca\u05c9\u0001\u0000\u0000\u0000"+
		"\u05cb\u05ce\u0001\u0000\u0000\u0000\u05cc\u05ca\u0001\u0000\u0000\u0000"+
		"\u05cc\u05cd\u0001\u0000\u0000\u0000\u05cd\u05cf\u0001\u0000\u0000\u0000"+
		"\u05ce\u05cc\u0001\u0000\u0000\u0000\u05cf\u05d1\u0003\u00bc^\u0000\u05d0"+
		"\u05cc\u0001\u0000\u0000\u0000\u05d0\u05d1\u0001\u0000\u0000\u0000\u05d1"+
		"\u05d5\u0001\u0000\u0000\u0000\u05d2\u05d4\u0005Y\u0000\u0000\u05d3\u05d2"+
		"\u0001\u0000\u0000\u0000\u05d4\u05d7\u0001\u0000\u0000\u0000\u05d5\u05d3"+
		"\u0001\u0000\u0000\u0000\u05d5\u05d6\u0001\u0000\u0000\u0000\u05d6\u05d8"+
		"\u0001\u0000\u0000\u0000\u05d7\u05d5\u0001\u0000\u0000\u0000\u05d8\u05d9"+
		"\u0005#\u0000\u0000\u05d9\u00b9\u0001\u0000\u0000\u0000\u05da\u05dc\u0003"+
		"\u013e\u009f\u0000\u05db\u05dd\u0005&\u0000\u0000\u05dc\u05db\u0001\u0000"+
		"\u0000\u0000\u05dc\u05dd\u0001\u0000\u0000\u0000\u05dd\u05df\u0001\u0000"+
		"\u0000\u0000\u05de\u05da\u0001\u0000\u0000\u0000\u05de\u05df\u0001\u0000"+
		"\u0000\u0000\u05df\u05e0\u0001\u0000\u0000\u0000\u05e0\u05e1\u0003\u00c0"+
		"`\u0000\u05e1\u00bb\u0001\u0000\u0000\u0000\u05e2\u05ee\u0003\u00ba]\u0000"+
		"\u05e3\u05e7\u0003\u0176\u00bb\u0000\u05e4\u05e6\u0005Y\u0000\u0000\u05e5"+
		"\u05e4\u0001\u0000\u0000\u0000\u05e6\u05e9\u0001\u0000\u0000\u0000\u05e7"+
		"\u05e5\u0001\u0000\u0000\u0000\u05e7\u05e8\u0001\u0000\u0000\u0000\u05e8"+
		"\u05ea\u0001\u0000\u0000\u0000\u05e9\u05e7\u0001\u0000\u0000\u0000\u05ea"+
		"\u05eb\u0003\u00ba]\u0000\u05eb\u05ed\u0001\u0000\u0000\u0000\u05ec\u05e3"+
		"\u0001\u0000\u0000\u0000\u05ed\u05f0\u0001\u0000\u0000\u0000\u05ee\u05ec"+
		"\u0001\u0000\u0000\u0000\u05ee\u05ef\u0001\u0000\u0000\u0000\u05ef\u05f2"+
		"\u0001\u0000\u0000\u0000\u05f0\u05ee\u0001\u0000\u0000\u0000\u05f1\u05f3"+
		"\u0003\u0176\u00bb\u0000\u05f2\u05f1\u0001\u0000\u0000\u0000\u05f2\u05f3"+
		"\u0001\u0000\u0000\u0000\u05f3\u00bd\u0001\u0000\u0000\u0000\u05f4\u05fb"+
		"\u0003\u00b2Y\u0000\u05f5\u05f7\u0005Y\u0000\u0000\u05f6\u05f5\u0001\u0000"+
		"\u0000\u0000\u05f6\u05f7\u0001\u0000\u0000\u0000\u05f7\u05f8\u0001\u0000"+
		"\u0000\u0000\u05f8\u05fa\u0003\u00b2Y\u0000\u05f9\u05f6\u0001\u0000\u0000"+
		"\u0000\u05fa\u05fd\u0001\u0000\u0000\u0000\u05fb\u05f9\u0001\u0000\u0000"+
		"\u0000\u05fb\u05fc\u0001\u0000\u0000\u0000\u05fc\u00bf\u0001\u0000\u0000"+
		"\u0000\u05fd\u05fb\u0001\u0000\u0000\u0000\u05fe\u0600\u0003\u00c4b\u0000"+
		"\u05ff\u0601\u0003\u00c8d\u0000\u0600\u05ff\u0001\u0000\u0000\u0000\u0600"+
		"\u0601\u0001\u0000\u0000\u0000\u0601\u00c1\u0001\u0000\u0000\u0000\u0602"+
		"\u060e\u0003\u00c0`\u0000\u0603\u0607\u0003\u0176\u00bb\u0000\u0604\u0606"+
		"\u0005Y\u0000\u0000\u0605\u0604\u0001\u0000\u0000\u0000\u0606\u0609\u0001"+
		"\u0000\u0000\u0000\u0607\u0605\u0001\u0000\u0000\u0000\u0607\u0608\u0001"+
		"\u0000\u0000\u0000\u0608\u060a\u0001\u0000\u0000\u0000\u0609\u0607\u0001"+
		"\u0000\u0000\u0000\u060a\u060b\u0003\u00c0`\u0000\u060b\u060d\u0001\u0000"+
		"\u0000\u0000\u060c\u0603\u0001\u0000\u0000\u0000\u060d\u0610\u0001\u0000"+
		"\u0000\u0000\u060e\u060c\u0001\u0000\u0000\u0000\u060e\u060f\u0001\u0000"+
		"\u0000\u0000\u060f\u0612\u0001\u0000\u0000\u0000\u0610\u060e\u0001\u0000"+
		"\u0000\u0000\u0611\u0613\u0003\u0176\u00bb\u0000\u0612\u0611\u0001\u0000"+
		"\u0000\u0000\u0612\u0613\u0001\u0000\u0000\u0000\u0613\u00c3\u0001\u0000"+
		"\u0000\u0000\u0614\u0615\u0003\u0158\u00ac\u0000\u0615\u0616\u0003\u0100"+
		"\u0080\u0000\u0616\u0619\u0001\u0000\u0000\u0000\u0617\u0619\u0003\u0100"+
		"\u0080\u0000\u0618\u0614\u0001\u0000\u0000\u0000\u0618\u0617\u0001\u0000"+
		"\u0000\u0000\u0619\u00c5\u0001\u0000\u0000\u0000\u061a\u061b\u0003\u0156"+
		"\u00ab\u0000\u061b\u061c\u0003\u00c4b\u0000\u061c\u0625\u0001\u0000\u0000"+
		"\u0000\u061d\u061e\u0003\u014c\u00a6\u0000\u061e\u061f\u0003\u00c4b\u0000"+
		"\u061f\u0625\u0001\u0000\u0000\u0000\u0620\u0621\u0003\u00cae\u0000\u0621"+
		"\u0622\u0003\u00c4b\u0000\u0622\u0625\u0001\u0000\u0000\u0000\u0623\u0625"+
		"\u0003\u00ccf\u0000\u0624\u061a\u0001\u0000\u0000\u0000\u0624\u061d\u0001"+
		"\u0000\u0000\u0000\u0624\u0620\u0001\u0000\u0000\u0000\u0624\u0623\u0001"+
		"\u0000\u0000\u0000\u0625\u00c7\u0001\u0000\u0000\u0000\u0626\u0628\u0003"+
		"\u00c6c\u0000\u0627\u0626\u0001\u0000\u0000\u0000\u0628\u0629\u0001\u0000"+
		"\u0000\u0000\u0629\u0627\u0001\u0000\u0000\u0000\u0629\u062a\u0001\u0000"+
		"\u0000\u0000\u062a\u00c9\u0001\u0000\u0000\u0000\u062b\u062c\u0005+\u0000"+
		"\u0000\u062c\u062d\u0003\u00c0`\u0000\u062d\u062e\u0005&\u0000\u0000\u062e"+
		"\u00cb\u0001\u0000\u0000\u0000\u062f\u0630\u0005\u0010\u0000\u0000\u0630"+
		"\u0634\u0003\u011a\u008d\u0000\u0631\u0632\u0005\u0002\u0000\u0000\u0632"+
		"\u0634\u0003\u011a\u008d\u0000\u0633\u062f\u0001\u0000\u0000\u0000\u0633"+
		"\u0631\u0001\u0000\u0000\u0000\u0634\u00cd\u0001\u0000\u0000\u0000\u0635"+
		"\u0647\u0003\u00d0h\u0000\u0636\u0638\u0003\u013c\u009e\u0000\u0637\u0639"+
		"\u0003.\u0017\u0000\u0638\u0637\u0001\u0000\u0000\u0000\u0638\u0639\u0001"+
		"\u0000\u0000\u0000\u0639\u0647\u0001\u0000\u0000\u0000\u063a\u063b\u0003"+
		"\u0122\u0091\u0000\u063b\u063c\u0005\u001e\u0000\u0000\u063c\u063e\u0003"+
		"\u013c\u009e\u0000\u063d\u063f\u0003.\u0017\u0000\u063e\u063d\u0001\u0000"+
		"\u0000\u0000\u063e\u063f\u0001\u0000\u0000\u0000\u063f\u0647\u0001\u0000"+
		"\u0000\u0000\u0640\u0647\u0003\u00f0x\u0000\u0641\u0647\u0003\u00f8|\u0000"+
		"\u0642\u0647\u0003\u00fa}\u0000\u0643\u0647\u0003\u00f6{\u0000\u0644\u0647"+
		"\u0003\u00fe\u007f\u0000\u0645\u0647\u0003\u00eew\u0000\u0646\u0635\u0001"+
		"\u0000\u0000\u0000\u0646\u0636\u0001\u0000\u0000\u0000\u0646\u063a\u0001"+
		"\u0000\u0000\u0000\u0646\u0640\u0001\u0000\u0000\u0000\u0646\u0641\u0001"+
		"\u0000\u0000\u0000\u0646\u0642\u0001\u0000\u0000\u0000\u0646\u0643\u0001"+
		"\u0000\u0000\u0000\u0646\u0644\u0001\u0000\u0000\u0000\u0646\u0645\u0001"+
		"\u0000\u0000\u0000\u0647\u00cf\u0001\u0000\u0000\u0000\u0648\u0650\u0003"+
		"\u00d2i\u0000\u0649\u0650\u0003\u00d4j\u0000\u064a\u0650\u0003\u00ecv"+
		"\u0000\u064b\u0650\u0003\u0168\u00b4\u0000\u064c\u0650\u0003\u00dam\u0000"+
		"\u064d\u0650\u0003\u00e0p\u0000\u064e\u0650\u0003\u00e6s\u0000\u064f\u0648"+
		"\u0001\u0000\u0000\u0000\u064f\u0649\u0001\u0000\u0000\u0000\u064f\u064a"+
		"\u0001\u0000\u0000\u0000\u064f\u064b\u0001\u0000\u0000\u0000\u064f\u064c"+
		"\u0001\u0000\u0000\u0000\u064f\u064d\u0001\u0000\u0000\u0000\u064f\u064e"+
		"\u0001\u0000\u0000\u0000\u0650\u00d1\u0001\u0000\u0000\u0000\u0651\u0652"+
		"\u0003\u016e\u00b7\u0000\u0652\u0653\u0003\u00d6k\u0000\u0653\u00d3\u0001"+
		"\u0000\u0000\u0000\u0654\u0655\u0003\u00d8l\u0000\u0655\u0656\u0004j\u0001"+
		"\u0000\u0656\u0658\u0003\u0172\u00b9\u0000\u0657\u0659\u0003\u00d6k\u0000"+
		"\u0658\u0657\u0001\u0000\u0000\u0000\u0658\u0659\u0001\u0000\u0000\u0000"+
		"\u0659\u065e\u0001\u0000\u0000\u0000\u065a\u065b\u0003\u0172\u00b9\u0000"+
		"\u065b\u065c\u0003\u00d6k\u0000\u065c\u065e\u0001\u0000\u0000\u0000\u065d"+
		"\u0654\u0001\u0000\u0000\u0000\u065d\u065a\u0001\u0000\u0000\u0000\u065e"+
		"\u00d5\u0001\u0000\u0000\u0000\u065f\u0660\u0004k\u0002\u0000\u0660\u0661"+
		"\u0007\u0000\u0000\u0000\u0661\u00d7\u0001\u0000\u0000\u0000\u0662\u0663"+
		"\u0005J\u0000\u0000\u0663\u00d9\u0001\u0000\u0000\u0000\u0664\u066c\u0005"+
		"!\u0000\u0000\u0665\u0667\u0005Y\u0000\u0000\u0666\u0665\u0001\u0000\u0000"+
		"\u0000\u0667\u066a\u0001\u0000\u0000\u0000\u0668\u0666\u0001\u0000\u0000"+
		"\u0000\u0668\u0669\u0001\u0000\u0000\u0000\u0669\u066b\u0001\u0000\u0000"+
		"\u0000\u066a\u0668\u0001\u0000\u0000\u0000\u066b\u066d\u0003\u00dcn\u0000"+
		"\u066c\u0668\u0001\u0000\u0000\u0000\u066c\u066d\u0001\u0000\u0000\u0000"+
		"\u066d\u0671\u0001\u0000\u0000\u0000\u066e\u0670\u0005Y\u0000\u0000\u066f"+
		"\u066e\u0001\u0000\u0000\u0000\u0670\u0673\u0001\u0000\u0000\u0000\u0671"+
		"\u066f\u0001\u0000\u0000\u0000\u0671\u0672\u0001\u0000\u0000\u0000\u0672"+
		"\u0674\u0001\u0000\u0000\u0000\u0673\u0671\u0001\u0000\u0000\u0000\u0674"+
		"\u0675\u0005$\u0000\u0000\u0675\u00db\u0001\u0000\u0000\u0000\u0676\u0682"+
		"\u0003\u00deo\u0000\u0677\u067b\u0003\u0176\u00bb\u0000\u0678\u067a\u0005"+
		"Y\u0000\u0000\u0679\u0678\u0001\u0000\u0000\u0000\u067a\u067d\u0001\u0000"+
		"\u0000\u0000\u067b\u0679\u0001\u0000\u0000\u0000\u067b\u067c\u0001\u0000"+
		"\u0000\u0000\u067c\u067e\u0001\u0000\u0000\u0000\u067d\u067b\u0001\u0000"+
		"\u0000\u0000\u067e\u067f\u0003\u00deo\u0000\u067f\u0681\u0001\u0000\u0000"+
		"\u0000\u0680\u0677\u0001\u0000\u0000\u0000\u0681\u0684\u0001\u0000\u0000"+
		"\u0000\u0682\u0680\u0001\u0000\u0000\u0000\u0682\u0683\u0001\u0000\u0000"+
		"\u0000\u0683\u0686\u0001\u0000\u0000\u0000\u0684\u0682\u0001\u0000\u0000"+
		"\u0000\u0685\u0687\u0003\u0176\u00bb\u0000\u0686\u0685\u0001\u0000\u0000"+
		"\u0000\u0686\u0687\u0001\u0000\u0000\u0000\u0687\u00dd\u0001\u0000\u0000"+
		"\u0000\u0688\u0689\u0003\u00c0`\u0000\u0689\u00df\u0001\u0000\u0000\u0000"+
		"\u068a\u0692\u0005\u001f\u0000\u0000\u068b\u068d\u0005Y\u0000\u0000\u068c"+
		"\u068b\u0001\u0000\u0000\u0000\u068d\u0690\u0001\u0000\u0000\u0000\u068e"+
		"\u068c\u0001\u0000\u0000\u0000\u068e\u068f\u0001\u0000\u0000\u0000\u068f"+
		"\u0691\u0001\u0000\u0000\u0000\u0690\u068e\u0001\u0000\u0000\u0000\u0691"+
		"\u0693\u0003\u00e2q\u0000\u0692\u068e\u0001\u0000\u0000\u0000\u0692\u0693"+
		"\u0001\u0000\u0000\u0000\u0693\u0697\u0001\u0000\u0000\u0000\u0694\u0696"+
		"\u0005Y\u0000\u0000\u0695\u0694\u0001\u0000\u0000\u0000\u0696\u0699\u0001"+
		"\u0000\u0000\u0000\u0697\u0695\u0001\u0000\u0000\u0000\u0697\u0698\u0001"+
		"\u0000\u0000\u0000\u0698\u069a\u0001\u0000\u0000\u0000\u0699\u0697\u0001"+
		"\u0000\u0000\u0000\u069a\u069b\u0005\"\u0000\u0000\u069b\u00e1\u0001\u0000"+
		"\u0000\u0000\u069c\u06a8\u0003\u00e4r\u0000\u069d\u06a1\u0003\u0176\u00bb"+
		"\u0000\u069e\u06a0\u0005Y\u0000\u0000\u069f\u069e\u0001\u0000\u0000\u0000"+
		"\u06a0\u06a3\u0001\u0000\u0000\u0000\u06a1\u069f\u0001\u0000\u0000\u0000"+
		"\u06a1\u06a2\u0001\u0000\u0000\u0000\u06a2\u06a4\u0001\u0000\u0000\u0000"+
		"\u06a3\u06a1\u0001\u0000\u0000\u0000\u06a4\u06a5\u0003\u00e4r\u0000\u06a5"+
		"\u06a7\u0001\u0000\u0000\u0000\u06a6\u069d\u0001\u0000\u0000\u0000\u06a7"+
		"\u06aa\u0001\u0000\u0000\u0000\u06a8\u06a6\u0001\u0000\u0000\u0000\u06a8"+
		"\u06a9\u0001\u0000\u0000\u0000\u06a9\u06ac\u0001\u0000\u0000\u0000\u06aa"+
		"\u06a8\u0001\u0000\u0000\u0000\u06ab\u06ad\u0003\u0176\u00bb\u0000\u06ac"+
		"\u06ab\u0001\u0000\u0000\u0000\u06ac\u06ad\u0001\u0000\u0000\u0000\u06ad"+
		"\u00e3\u0001\u0000\u0000\u0000\u06ae\u06b1\u0003\u0172\u00b9\u0000\u06af"+
		"\u06b1\u0003\u0170\u00b8\u0000\u06b0\u06ae\u0001\u0000\u0000\u0000\u06b0"+
		"\u06af\u0001\u0000\u0000\u0000\u06b1\u06b2\u0001\u0000\u0000\u0000\u06b2"+
		"\u06b3\u0005&\u0000\u0000\u06b3\u06b4\u0003\u00c0`\u0000\u06b4\u00e5\u0001"+
		"\u0000\u0000\u0000\u06b5\u06bd\u0005\u001f\u0000\u0000\u06b6\u06b8\u0005"+
		"Y\u0000\u0000\u06b7\u06b6\u0001\u0000\u0000\u0000\u06b8\u06bb\u0001\u0000"+
		"\u0000\u0000\u06b9\u06b7\u0001\u0000\u0000\u0000\u06b9\u06ba\u0001\u0000"+
		"\u0000\u0000\u06ba\u06bc\u0001\u0000\u0000\u0000\u06bb\u06b9\u0001\u0000"+
		"\u0000\u0000\u06bc\u06be\u0003\u00e8t\u0000\u06bd\u06b9\u0001\u0000\u0000"+
		"\u0000\u06bd\u06be\u0001\u0000\u0000\u0000\u06be\u06c2\u0001\u0000\u0000"+
		"\u0000\u06bf\u06c1\u0005Y\u0000\u0000\u06c0\u06bf\u0001\u0000\u0000\u0000"+
		"\u06c1\u06c4\u0001\u0000\u0000\u0000\u06c2\u06c0\u0001\u0000\u0000\u0000"+
		"\u06c2\u06c3\u0001\u0000\u0000\u0000\u06c3\u06c5\u0001\u0000\u0000\u0000"+
		"\u06c4\u06c2\u0001\u0000\u0000\u0000\u06c5\u06c6\u0005\"\u0000\u0000\u06c6"+
		"\u00e7\u0001\u0000\u0000\u0000\u06c7\u06d3\u0003\u00eau\u0000\u06c8\u06cc"+
		"\u0003\u0176\u00bb\u0000\u06c9\u06cb\u0005Y\u0000\u0000\u06ca\u06c9\u0001"+
		"\u0000\u0000\u0000\u06cb\u06ce\u0001\u0000\u0000\u0000\u06cc\u06ca\u0001"+
		"\u0000\u0000\u0000\u06cc\u06cd\u0001\u0000\u0000\u0000\u06cd\u06cf\u0001"+
		"\u0000\u0000\u0000\u06ce\u06cc\u0001\u0000\u0000\u0000\u06cf\u06d0\u0003"+
		"\u00eau\u0000\u06d0\u06d2\u0001\u0000\u0000\u0000\u06d1\u06c8\u0001\u0000"+
		"\u0000\u0000\u06d2\u06d5\u0001\u0000\u0000\u0000\u06d3\u06d1\u0001\u0000"+
		"\u0000\u0000\u06d3\u06d4\u0001\u0000\u0000\u0000\u06d4\u06d7\u0001\u0000"+
		"\u0000\u0000\u06d5\u06d3\u0001\u0000\u0000\u0000\u06d6\u06d8\u0003\u0176"+
		"\u00bb\u0000\u06d7\u06d6\u0001\u0000\u0000\u0000\u06d7\u06d8\u0001\u0000"+
		"\u0000\u0000\u06d8\u00e9\u0001\u0000\u0000\u0000\u06d9\u06dc\u0003\u0140"+
		"\u00a0\u0000\u06da\u06db\u0005&\u0000\u0000\u06db\u06dd\u0003\u00c0`\u0000"+
		"\u06dc\u06da\u0001\u0000\u0000\u0000\u06dc\u06dd\u0001\u0000\u0000\u0000"+
		"\u06dd\u00eb\u0001\u0000\u0000\u0000\u06de\u06df\u0003\u0122\u0091\u0000"+
		"\u06df\u06e0\u0003\u00e6s\u0000\u06e0\u00ed\u0001\u0000\u0000\u0000\u06e1"+
		"\u06e2\u0003\u0122\u0091\u0000\u06e2\u06e3\u0003\u0108\u0084\u0000\u06e3"+
		"\u00ef\u0001\u0000\u0000\u0000\u06e4\u06e5\u0005\u001f\u0000\u0000\u06e5"+
		"\u06e6\u0003\u0006\u0003\u0000\u06e6\u06e7\u0005\"\u0000\u0000\u06e7\u0704"+
		"\u0001\u0000\u0000\u0000\u06e8\u06e9\u0005\u001f\u0000\u0000\u06e9\u06ed"+
		"\u0003\u00f2y\u0000\u06ea\u06ec\u0005Y\u0000\u0000\u06eb\u06ea\u0001\u0000"+
		"\u0000\u0000\u06ec\u06ef\u0001\u0000\u0000\u0000\u06ed\u06eb\u0001\u0000"+
		"\u0000\u0000\u06ed\u06ee\u0001\u0000\u0000\u0000\u06ee\u06f0\u0001\u0000"+
		"\u0000\u0000\u06ef\u06ed\u0001\u0000\u0000\u0000\u06f0\u06f8\u0005>\u0000"+
		"\u0000\u06f1\u06f3\u0005Y\u0000\u0000\u06f2\u06f1\u0001\u0000\u0000\u0000"+
		"\u06f3\u06f6\u0001\u0000\u0000\u0000\u06f4\u06f2\u0001\u0000\u0000\u0000"+
		"\u06f4\u06f5\u0001\u0000\u0000\u0000\u06f5\u06f7\u0001\u0000\u0000\u0000"+
		"\u06f6\u06f4\u0001\u0000\u0000\u0000\u06f7\u06f9\u0003\u011a\u008d\u0000"+
		"\u06f8\u06f4\u0001\u0000\u0000\u0000\u06f8\u06f9\u0001\u0000\u0000\u0000"+
		"\u06f9\u06fd\u0001\u0000\u0000\u0000\u06fa\u06fc\u0005Y\u0000\u0000\u06fb"+
		"\u06fa\u0001\u0000\u0000\u0000\u06fc\u06ff\u0001\u0000\u0000\u0000\u06fd"+
		"\u06fb\u0001\u0000\u0000\u0000\u06fd\u06fe\u0001\u0000\u0000\u0000\u06fe"+
		"\u0700\u0001\u0000\u0000\u0000\u06ff\u06fd\u0001\u0000\u0000\u0000\u0700"+
		"\u0701\u0003\u0006\u0003\u0000\u0701\u0702\u0005\"\u0000\u0000\u0702\u0704"+
		"\u0001\u0000\u0000\u0000\u0703\u06e4\u0001\u0000\u0000\u0000\u0703\u06e8"+
		"\u0001\u0000\u0000\u0000\u0704\u00f1\u0001\u0000\u0000\u0000\u0705\u0711"+
		"\u0003\u00f4z\u0000\u0706\u070a\u0003\u0176\u00bb\u0000\u0707\u0709\u0005"+
		"Y\u0000\u0000\u0708\u0707\u0001\u0000\u0000\u0000\u0709\u070c\u0001\u0000"+
		"\u0000\u0000\u070a\u0708\u0001\u0000\u0000\u0000\u070a\u070b\u0001\u0000"+
		"\u0000\u0000\u070b\u070d\u0001\u0000\u0000\u0000\u070c\u070a\u0001\u0000"+
		"\u0000\u0000\u070d\u070e\u0003\u00f4z\u0000\u070e\u0710\u0001\u0000\u0000"+
		"\u0000\u070f\u0706\u0001\u0000\u0000\u0000\u0710\u0713\u0001\u0000\u0000"+
		"\u0000\u0711\u070f\u0001\u0000\u0000\u0000\u0711\u0712\u0001\u0000\u0000"+
		"\u0000\u0712\u0715\u0001\u0000\u0000\u0000\u0713\u0711\u0001\u0000\u0000"+
		"\u0000\u0714\u0716\u0003\u0176\u00bb\u0000\u0715\u0714\u0001\u0000\u0000"+
		"\u0000\u0715\u0716\u0001\u0000\u0000\u0000\u0716\u00f3\u0001\u0000\u0000"+
		"\u0000\u0717\u071a\u0003t:\u0000\u0718\u071a\u0003\u013e\u009f\u0000\u0719"+
		"\u0717\u0001\u0000\u0000\u0000\u0719\u0718\u0001\u0000\u0000\u0000\u071a"+
		"\u00f5\u0001\u0000\u0000\u0000\u071b\u071c\u0005\u001e\u0000\u0000\u071c"+
		"\u071d\u0003\u013e\u009f\u0000\u071d\u00f7\u0001\u0000\u0000\u0000\u071e"+
		"\u0722\u0005 \u0000\u0000\u071f\u0721\u0005Y\u0000\u0000\u0720\u071f\u0001"+
		"\u0000\u0000\u0000\u0721\u0724\u0001\u0000\u0000\u0000\u0722\u0720\u0001"+
		"\u0000\u0000\u0000\u0722\u0723\u0001\u0000\u0000\u0000\u0723\u0725\u0001"+
		"\u0000\u0000\u0000\u0724\u0722\u0001\u0000\u0000\u0000\u0725\u0729\u0003"+
		"\u00c0`\u0000\u0726\u0728\u0005Y\u0000\u0000\u0727\u0726\u0001\u0000\u0000"+
		"\u0000\u0728\u072b\u0001\u0000\u0000\u0000\u0729\u0727\u0001\u0000\u0000"+
		"\u0000\u0729\u072a\u0001\u0000\u0000\u0000\u072a\u072c\u0001\u0000\u0000"+
		"\u0000\u072b\u0729\u0001\u0000\u0000\u0000\u072c\u072d\u0005#\u0000\u0000"+
		"\u072d\u00f9\u0001\u0000\u0000\u0000\u072e\u072f\u0005 \u0000\u0000\u072f"+
		"\u073b\u0005#\u0000\u0000\u0730\u0731\u0005 \u0000\u0000\u0731\u0734\u0003"+
		"\u00fc~\u0000\u0732\u0733\u0005%\u0000\u0000\u0733\u0735\u0003\u00fc~"+
		"\u0000\u0734\u0732\u0001\u0000\u0000\u0000\u0735\u0736\u0001\u0000\u0000"+
		"\u0000\u0736\u0734\u0001\u0000\u0000\u0000\u0736\u0737\u0001\u0000\u0000"+
		"\u0000\u0737\u0738\u0001\u0000\u0000\u0000\u0738\u0739\u0005#\u0000\u0000"+
		"\u0739\u073b\u0001\u0000\u0000\u0000\u073a\u072e\u0001\u0000\u0000\u0000"+
		"\u073a\u0730\u0001\u0000\u0000\u0000\u073b\u00fb\u0001\u0000\u0000\u0000"+
		"\u073c\u0744\u0003\u00c0`\u0000\u073d\u073f\u0003\u013e\u009f\u0000\u073e"+
		"\u0740\u0005&\u0000\u0000\u073f\u073e\u0001\u0000\u0000\u0000\u073f\u0740"+
		"\u0001\u0000\u0000\u0000\u0740\u0741\u0001\u0000\u0000\u0000\u0741\u0742"+
		"\u0003\u00c0`\u0000\u0742\u0744\u0001\u0000\u0000\u0000\u0743\u073c\u0001"+
		"\u0000\u0000\u0000\u0743\u073d\u0001\u0000\u0000\u0000\u0744\u00fd\u0001"+
		"\u0000\u0000\u0000\u0745\u0746\u00059\u0000\u0000\u0746\u00ff\u0001\u0000"+
		"\u0000\u0000\u0747\u074b\u0003\u00ceg\u0000\u0748\u074a\u0003\u0102\u0081"+
		"\u0000\u0749\u0748\u0001\u0000\u0000\u0000\u074a\u074d\u0001\u0000\u0000"+
		"\u0000\u074b\u0749\u0001\u0000\u0000\u0000\u074b\u074c\u0001\u0000\u0000"+
		"\u0000\u074c\u074f\u0001\u0000\u0000\u0000\u074d\u074b\u0001\u0000\u0000"+
		"\u0000\u074e\u0750\u0003\u015a\u00ad\u0000\u074f\u074e\u0001\u0000\u0000"+
		"\u0000\u074f\u0750\u0001\u0000\u0000\u0000\u0750\u0101\u0001\u0000\u0000"+
		"\u0000\u0751\u0755\u0003\u0108\u0084\u0000\u0752\u0755\u0003\u0104\u0082"+
		"\u0000\u0753\u0755\u0003\u0106\u0083\u0000\u0754\u0751\u0001\u0000\u0000"+
		"\u0000\u0754\u0752\u0001\u0000\u0000\u0000\u0754\u0753\u0001\u0000\u0000"+
		"\u0000\u0755\u0103\u0001\u0000\u0000\u0000\u0756\u0760\u0005\u001e\u0000"+
		"\u0000\u0757\u0761\u0005P\u0000\u0000\u0758\u075e\u0003\u0142\u00a1\u0000"+
		"\u0759\u075f\u0003.\u0017\u0000\u075a\u075b\u0005 \u0000\u0000\u075b\u075c"+
		"\u0003\u0116\u008b\u0000\u075c\u075d\u0005#\u0000\u0000\u075d\u075f\u0001"+
		"\u0000\u0000\u0000\u075e\u0759\u0001\u0000\u0000\u0000\u075e\u075a\u0001"+
		"\u0000\u0000\u0000\u075e\u075f\u0001\u0000\u0000\u0000\u075f\u0761\u0001"+
		"\u0000\u0000\u0000\u0760\u0757\u0001\u0000\u0000\u0000\u0760\u0758\u0001"+
		"\u0000\u0000\u0000\u0761\u0105\u0001\u0000\u0000\u0000\u0762\u0763\u0005"+
		"!\u0000\u0000\u0763\u0764\u0003\u010c\u0086\u0000\u0764\u0765\u0005$\u0000"+
		"\u0000\u0765\u0107\u0001\u0000\u0000\u0000\u0766\u0768\u0003\u010a\u0085"+
		"\u0000\u0767\u0766\u0001\u0000\u0000\u0000\u0767\u0768\u0001\u0000\u0000"+
		"\u0000\u0768\u0769\u0001\u0000\u0000\u0000\u0769\u076c\u0003\u0110\u0088"+
		"\u0000\u076a\u076c\u0003\u010a\u0085\u0000\u076b\u0767\u0001\u0000\u0000"+
		"\u0000\u076b\u076a\u0001\u0000\u0000\u0000\u076c\u0109\u0001\u0000\u0000"+
		"\u0000\u076d\u076e\u0005 \u0000\u0000\u076e\u0774\u0005#\u0000\u0000\u076f"+
		"\u0770\u0005 \u0000\u0000\u0770\u0771\u0003\u010c\u0086\u0000\u0771\u0772"+
		"\u0005#\u0000\u0000\u0772\u0774\u0001\u0000\u0000\u0000\u0773\u076d\u0001"+
		"\u0000\u0000\u0000\u0773\u076f\u0001\u0000\u0000\u0000\u0774\u010b\u0001"+
		"\u0000\u0000\u0000\u0775\u077a\u0003\u010e\u0087\u0000\u0776\u0777\u0005"+
		"%\u0000\u0000\u0777\u0779\u0003\u010e\u0087\u0000\u0778\u0776\u0001\u0000"+
		"\u0000\u0000\u0779\u077c\u0001\u0000\u0000\u0000\u077a\u0778\u0001\u0000"+
		"\u0000\u0000\u077a\u077b\u0001\u0000\u0000\u0000\u077b\u010d\u0001\u0000"+
		"\u0000\u0000\u077c\u077a\u0001\u0000\u0000\u0000\u077d\u078c\u0003\u00c0"+
		"`\u0000\u077e\u0780\u0003\u013e\u009f\u0000\u077f\u0781\u0005&\u0000\u0000"+
		"\u0780\u077f\u0001\u0000\u0000\u0000\u0780\u0781\u0001\u0000\u0000\u0000"+
		"\u0781\u0782\u0001\u0000\u0000\u0000\u0782\u0783\u0003\u00c0`\u0000\u0783"+
		"\u078c\u0001\u0000\u0000\u0000\u0784\u078c\u0003\u015c\u00ae\u0000\u0785"+
		"\u0787\u0003\u013e\u009f\u0000\u0786\u0788\u0005&\u0000\u0000\u0787\u0786"+
		"\u0001\u0000\u0000\u0000\u0787\u0788\u0001\u0000\u0000\u0000\u0788\u0789"+
		"\u0001\u0000\u0000\u0000\u0789\u078a\u0003\u015c\u00ae\u0000\u078a\u078c"+
		"\u0001\u0000\u0000\u0000\u078b\u077d\u0001\u0000\u0000\u0000\u078b\u077e"+
		"\u0001\u0000\u0000\u0000\u078b\u0784\u0001\u0000\u0000\u0000\u078b\u0785"+
		"\u0001\u0000\u0000\u0000\u078c\u010f\u0001\u0000\u0000\u0000\u078d\u078f"+
		"\u0003\u00f0x\u0000\u078e\u0790\u0003\u0112\u0089\u0000\u078f\u078e\u0001"+
		"\u0000\u0000\u0000\u078f\u0790\u0001\u0000\u0000\u0000\u0790\u0111\u0001"+
		"\u0000\u0000\u0000\u0791\u0793\u0003\u0114\u008a\u0000\u0792\u0791\u0001"+
		"\u0000\u0000\u0000\u0793\u0794\u0001\u0000\u0000\u0000\u0794\u0792\u0001"+
		"\u0000\u0000\u0000\u0794\u0795\u0001\u0000\u0000\u0000\u0795\u0113\u0001"+
		"\u0000\u0000\u0000\u0796\u0797\u0003\u0142\u00a1\u0000\u0797\u0798\u0005"+
		"&\u0000\u0000\u0798\u0799\u0003\u00f0x\u0000\u0799\u0115\u0001\u0000\u0000"+
		"\u0000\u079a\u079e\u0003\u0118\u008c\u0000\u079b\u079d\u0003\u0118\u008c"+
		"\u0000\u079c\u079b\u0001\u0000\u0000\u0000\u079d\u07a0\u0001\u0000\u0000"+
		"\u0000\u079e\u079c\u0001\u0000\u0000\u0000\u079e\u079f\u0001\u0000\u0000"+
		"\u0000\u079f\u0117\u0001\u0000\u0000\u0000\u07a0\u079e\u0001\u0000\u0000"+
		"\u0000\u07a1\u07a2\u0003\u013e\u009f\u0000\u07a2\u07a3\u0005&\u0000\u0000"+
		"\u07a3\u0119\u0001\u0000\u0000\u0000\u07a4\u07a5\u0006\u008d\uffff\uffff"+
		"\u0000\u07a5\u07a8\u0003\u011c\u008e\u0000\u07a6\u07a8\u0003\u012e\u0097"+
		"\u0000\u07a7\u07a4\u0001\u0000\u0000\u0000\u07a7\u07a6\u0001\u0000\u0000"+
		"\u0000\u07a8\u07b1\u0001\u0000\u0000\u0000\u07a9\u07aa\n\u0003\u0000\u0000"+
		"\u07aa\u07b0\u0005*\u0000\u0000\u07ab\u07ac\n\u0002\u0000\u0000\u07ac"+
		"\u07b0\u0005+\u0000\u0000\u07ad\u07ae\n\u0001\u0000\u0000\u07ae\u07b0"+
		"\u0005A\u0000\u0000\u07af\u07a9\u0001\u0000\u0000\u0000\u07af\u07ab\u0001"+
		"\u0000\u0000\u0000\u07af\u07ad\u0001\u0000\u0000\u0000\u07b0\u07b3\u0001"+
		"\u0000\u0000\u0000\u07b1\u07af\u0001\u0000\u0000\u0000\u07b1\u07b2\u0001"+
		"\u0000\u0000\u0000\u07b2\u011b\u0001\u0000\u0000\u0000\u07b3\u07b1\u0001"+
		"\u0000\u0000\u0000\u07b4\u07b5\u0006\u008e\uffff\uffff\u0000\u07b5\u07b6"+
		"\u0003\u011e\u008f\u0000\u07b6\u07f7\u0001\u0000\u0000\u0000\u07b7\u07b9"+
		"\n\u0003\u0000\u0000\u07b8\u07ba\u0003\u00be_\u0000\u07b9\u07b8\u0001"+
		"\u0000\u0000\u0000\u07b9\u07ba\u0001\u0000\u0000\u0000\u07ba\u07be\u0001"+
		"\u0000\u0000\u0000\u07bb\u07bc\u0003\u014a\u00a5\u0000\u07bc\u07bd\u0005"+
		"Y\u0000\u0000\u07bd\u07bf\u0001\u0000\u0000\u0000\u07be\u07bb\u0001\u0000"+
		"\u0000\u0000\u07be\u07bf\u0001\u0000\u0000\u0000\u07bf\u07c3\u0001\u0000"+
		"\u0000\u0000\u07c0\u07c2\u0005Y\u0000\u0000\u07c1\u07c0\u0001\u0000\u0000"+
		"\u0000\u07c2\u07c5\u0001\u0000\u0000\u0000\u07c3\u07c1\u0001\u0000\u0000"+
		"\u0000\u07c3\u07c4\u0001\u0000\u0000\u0000\u07c4\u07c6\u0001\u0000\u0000"+
		"\u0000\u07c5\u07c3\u0001\u0000\u0000\u0000\u07c6\u07ca\u00050\u0000\u0000"+
		"\u07c7\u07c9\u0005Y\u0000\u0000\u07c8\u07c7\u0001\u0000\u0000\u0000\u07c9"+
		"\u07cc\u0001\u0000\u0000\u0000\u07ca\u07c8\u0001\u0000\u0000\u0000\u07ca"+
		"\u07cb\u0001\u0000\u0000\u0000\u07cb\u07cd\u0001\u0000\u0000\u0000\u07cc"+
		"\u07ca\u0001\u0000\u0000\u0000\u07cd\u07cf\u0003\u011c\u008e\u0000\u07ce"+
		"\u07d0\u0003\u00be_\u0000\u07cf\u07ce\u0001\u0000\u0000\u0000\u07cf\u07d0"+
		"\u0001\u0000\u0000\u0000\u07d0\u07d4\u0001\u0000\u0000\u0000\u07d1\u07d2"+
		"\u0003\u014a\u00a5\u0000\u07d2\u07d3\u0005Y\u0000\u0000\u07d3\u07d5\u0001"+
		"\u0000\u0000\u0000\u07d4\u07d1\u0001\u0000\u0000\u0000\u07d4\u07d5\u0001"+
		"\u0000\u0000\u0000\u07d5\u07f6\u0001\u0000\u0000\u0000\u07d6\u07d8\n\u0002"+
		"\u0000\u0000\u07d7\u07d9\u0003\u00be_\u0000\u07d8\u07d7\u0001\u0000\u0000"+
		"\u0000\u07d8\u07d9\u0001\u0000\u0000\u0000\u07d9\u07dd\u0001\u0000\u0000"+
		"\u0000\u07da\u07db\u0003\u014a\u00a5\u0000\u07db\u07dc\u0005Y\u0000\u0000"+
		"\u07dc\u07de\u0001\u0000\u0000\u0000\u07dd\u07da\u0001\u0000\u0000\u0000"+
		"\u07dd\u07de\u0001\u0000\u0000\u0000\u07de\u07e2\u0001\u0000\u0000\u0000"+
		"\u07df\u07e1\u0005Y\u0000\u0000\u07e0\u07df\u0001\u0000\u0000\u0000\u07e1"+
		"\u07e4\u0001\u0000\u0000\u0000\u07e2\u07e0\u0001\u0000\u0000\u0000\u07e2"+
		"\u07e3\u0001\u0000\u0000\u0000\u07e3\u07e5\u0001\u0000\u0000\u0000\u07e4"+
		"\u07e2\u0001\u0000\u0000\u0000\u07e5\u07e9\u0005-\u0000\u0000\u07e6\u07e8"+
		"\u0005Y\u0000\u0000\u07e7\u07e6\u0001\u0000\u0000\u0000\u07e8\u07eb\u0001"+
		"\u0000\u0000\u0000\u07e9\u07e7\u0001\u0000\u0000\u0000\u07e9\u07ea\u0001"+
		"\u0000\u0000\u0000\u07ea\u07ec\u0001\u0000\u0000\u0000\u07eb\u07e9\u0001"+
		"\u0000\u0000\u0000\u07ec\u07ee\u0003\u011c\u008e\u0000\u07ed\u07ef\u0003"+
		"\u00be_\u0000\u07ee\u07ed\u0001\u0000\u0000\u0000\u07ee\u07ef\u0001\u0000"+
		"\u0000\u0000\u07ef\u07f3\u0001\u0000\u0000\u0000\u07f0\u07f1\u0003\u014a"+
		"\u00a5\u0000\u07f1\u07f2\u0005Y\u0000\u0000\u07f2\u07f4\u0001\u0000\u0000"+
		"\u0000\u07f3\u07f0\u0001\u0000\u0000\u0000\u07f3\u07f4\u0001\u0000\u0000"+
		"\u0000\u07f4\u07f6\u0001\u0000\u0000\u0000\u07f5\u07b7\u0001\u0000\u0000"+
		"\u0000\u07f5\u07d6\u0001\u0000\u0000\u0000\u07f6\u07f9\u0001\u0000\u0000"+
		"\u0000\u07f7\u07f5\u0001\u0000\u0000\u0000\u07f7\u07f8\u0001\u0000\u0000"+
		"\u0000\u07f8\u011d\u0001\u0000\u0000\u0000\u07f9\u07f7\u0001\u0000\u0000"+
		"\u0000\u07fa\u07ff\u0003\u0130\u0098\u0000\u07fb\u07ff\u0003\u0132\u0099"+
		"\u0000\u07fc\u07ff\u0003\u0128\u0094\u0000\u07fd\u07ff\u0003\u0122\u0091"+
		"\u0000\u07fe\u07fa\u0001\u0000\u0000\u0000\u07fe\u07fb\u0001\u0000\u0000"+
		"\u0000\u07fe\u07fc\u0001\u0000\u0000\u0000\u07fe\u07fd\u0001\u0000\u0000"+
		"\u0000\u07ff\u011f\u0001\u0000\u0000\u0000\u0800\u0802\u0005&\u0000\u0000"+
		"\u0801\u0800\u0001\u0000\u0000\u0000\u0801\u0802\u0001\u0000\u0000\u0000"+
		"\u0802\u0803\u0001\u0000\u0000\u0000\u0803\u0805\u0003\u011a\u008d\u0000"+
		"\u0804\u0806\u0003\u00be_\u0000\u0805\u0804\u0001\u0000\u0000\u0000\u0805"+
		"\u0806\u0001\u0000\u0000\u0000\u0806\u0121\u0001\u0000\u0000\u0000\u0807"+
		"\u0808\u0003:\u001d\u0000\u0808\u0809\u0005\u001e\u0000\u0000\u0809\u080b"+
		"\u0001\u0000\u0000\u0000\u080a\u0807\u0001\u0000\u0000\u0000\u080a\u080b"+
		"\u0001\u0000\u0000\u0000\u080b\u080c\u0001\u0000\u0000\u0000\u080c\u0811"+
		"\u0003\u0124\u0092\u0000\u080d\u080e\u0005\u001e\u0000\u0000\u080e\u0810"+
		"\u0003\u0124\u0092\u0000\u080f\u080d\u0001\u0000\u0000\u0000\u0810\u0813"+
		"\u0001\u0000\u0000\u0000\u0811\u080f\u0001\u0000\u0000\u0000\u0811\u0812"+
		"\u0001\u0000\u0000\u0000\u0812\u0123\u0001\u0000\u0000\u0000\u0813\u0811"+
		"\u0001\u0000\u0000\u0000\u0814\u0816\u0003\u0126\u0093\u0000\u0815\u0817"+
		"\u0003.\u0017\u0000\u0816\u0815\u0001\u0000\u0000\u0000\u0816\u0817\u0001"+
		"\u0000\u0000\u0000\u0817\u0125\u0001\u0000\u0000\u0000\u0818\u0819\u0005"+
		"I\u0000\u0000\u0819\u0127\u0001\u0000\u0000\u0000\u081a\u0822\u0005 \u0000"+
		"\u0000\u081b\u081d\u0005Y\u0000\u0000\u081c\u081b\u0001\u0000\u0000\u0000"+
		"\u081d\u0820\u0001\u0000\u0000\u0000\u081e\u081c\u0001\u0000\u0000\u0000"+
		"\u081e\u081f\u0001\u0000\u0000\u0000\u081f\u0821\u0001\u0000\u0000\u0000"+
		"\u0820\u081e\u0001\u0000\u0000\u0000\u0821\u0823\u0003\u012a\u0095\u0000"+
		"\u0822\u081e\u0001\u0000\u0000\u0000\u0822\u0823\u0001\u0000\u0000\u0000"+
		"\u0823\u0827\u0001\u0000\u0000\u0000\u0824\u0826\u0005Y\u0000\u0000\u0825"+
		"\u0824\u0001\u0000\u0000\u0000\u0826\u0829\u0001\u0000\u0000\u0000\u0827"+
		"\u0825\u0001\u0000\u0000\u0000\u0827\u0828\u0001\u0000\u0000\u0000\u0828"+
		"\u082a\u0001\u0000\u0000\u0000\u0829\u0827\u0001\u0000\u0000\u0000\u082a"+
		"\u082b\u0005#\u0000\u0000\u082b\u0129\u0001\u0000\u0000\u0000\u082c\u0838"+
		"\u0003\u012c\u0096\u0000\u082d\u0831\u0003\u017a\u00bd\u0000\u082e\u0830"+
		"\u0005Y\u0000\u0000\u082f\u082e\u0001\u0000\u0000\u0000\u0830\u0833\u0001"+
		"\u0000\u0000\u0000\u0831\u082f\u0001\u0000\u0000\u0000\u0831\u0832\u0001"+
		"\u0000\u0000\u0000\u0832\u0834\u0001\u0000\u0000\u0000\u0833\u0831\u0001"+
		"\u0000\u0000\u0000\u0834\u0835\u0003\u012c\u0096\u0000\u0835\u0837\u0001"+
		"\u0000\u0000\u0000\u0836\u082d\u0001\u0000\u0000\u0000\u0837\u083a\u0001"+
		"\u0000\u0000\u0000\u0838\u0836\u0001\u0000\u0000\u0000\u0838\u0839\u0001"+
		"\u0000\u0000\u0000\u0839\u083c\u0001\u0000\u0000\u0000\u083a\u0838\u0001"+
		"\u0000\u0000\u0000\u083b\u083d\u0003\u017a\u00bd\u0000\u083c\u083b\u0001"+
		"\u0000\u0000\u0000\u083c\u083d\u0001\u0000\u0000\u0000\u083d\u012b\u0001"+
		"\u0000\u0000\u0000\u083e\u0840\u0003\u013c\u009e\u0000\u083f\u0841\u0005"+
		"&\u0000\u0000\u0840\u083f\u0001\u0000\u0000\u0000\u0840\u0841\u0001\u0000"+
		"\u0000\u0000\u0841\u0843\u0001\u0000\u0000\u0000\u0842\u083e\u0001\u0000"+
		"\u0000\u0000\u0842\u0843\u0001\u0000\u0000\u0000\u0843\u0844\u0001\u0000"+
		"\u0000\u0000\u0844\u0846\u0003\u011a\u008d\u0000\u0845\u0847\u0003\u00be"+
		"_\u0000\u0846\u0845\u0001\u0000\u0000\u0000\u0846\u0847\u0001\u0000\u0000"+
		"\u0000\u0847\u012d\u0001\u0000\u0000\u0000\u0848\u0849\u0003p8\u0000\u0849"+
		"\u084a\u0003\u0150\u00a8\u0000\u084a\u084c\u0003\u011a\u008d\u0000\u084b"+
		"\u084d\u0003\u00be_\u0000\u084c\u084b\u0001\u0000\u0000\u0000\u084c\u084d"+
		"\u0001\u0000\u0000\u0000\u084d\u012f\u0001\u0000\u0000\u0000\u084e\u084f"+
		"\u0005!\u0000\u0000\u084f\u0851\u0003\u011a\u008d\u0000\u0850\u0852\u0003"+
		"\u00be_\u0000\u0851\u0850\u0001\u0000\u0000\u0000\u0851\u0852\u0001\u0000"+
		"\u0000\u0000\u0852\u0853\u0001\u0000\u0000\u0000\u0853\u0854\u0005$\u0000"+
		"\u0000\u0854\u0131\u0001\u0000\u0000\u0000\u0855\u0856\u0005\u001f\u0000"+
		"\u0000\u0856\u0858\u0003\u011a\u008d\u0000\u0857\u0859\u0003\u0134\u009a"+
		"\u0000\u0858\u0857\u0001\u0000\u0000\u0000\u0858\u0859\u0001\u0000\u0000"+
		"\u0000\u0859\u085a\u0001\u0000\u0000\u0000\u085a\u085b\u0005&\u0000\u0000"+
		"\u085b\u085d\u0003\u011a\u008d\u0000\u085c\u085e\u0003\u00be_\u0000\u085d"+
		"\u085c\u0001\u0000\u0000\u0000\u085d\u085e\u0001\u0000\u0000\u0000\u085e"+
		"\u085f\u0001\u0000\u0000\u0000\u085f\u0860\u0005\"\u0000\u0000\u0860\u0133"+
		"\u0001\u0000\u0000\u0000\u0861\u0862\u0003\u00be_\u0000\u0862\u0135\u0001"+
		"\u0000\u0000\u0000\u0863\u0867\u0005&\u0000\u0000\u0864\u0866\u0005Y\u0000"+
		"\u0000\u0865\u0864\u0001\u0000\u0000\u0000\u0866\u0869\u0001\u0000\u0000"+
		"\u0000\u0867\u0865\u0001\u0000\u0000\u0000\u0867\u0868\u0001\u0000\u0000"+
		"\u0000\u0868\u086a\u0001\u0000\u0000\u0000\u0869\u0867\u0001\u0000\u0000"+
		"\u0000\u086a\u086b\u0003\u0138\u009c\u0000\u086b\u0137\u0001\u0000\u0000"+
		"\u0000\u086c\u0878\u0003\u013a\u009d\u0000\u086d\u0871\u0003\u017a\u00bd"+
		"\u0000\u086e\u0870\u0005Y\u0000\u0000\u086f\u086e\u0001\u0000\u0000\u0000"+
		"\u0870\u0873\u0001\u0000\u0000\u0000\u0871\u086f\u0001\u0000\u0000\u0000"+
		"\u0871\u0872\u0001\u0000\u0000\u0000\u0872\u0874\u0001\u0000\u0000\u0000"+
		"\u0873\u0871\u0001\u0000\u0000\u0000\u0874\u0875\u0003\u013a\u009d\u0000"+
		"\u0875\u0877\u0001\u0000\u0000\u0000\u0876\u086d\u0001\u0000\u0000\u0000"+
		"\u0877\u087a\u0001\u0000\u0000\u0000\u0878\u0876\u0001\u0000\u0000\u0000"+
		"\u0878\u0879\u0001\u0000\u0000\u0000\u0879\u087c\u0001\u0000\u0000\u0000"+
		"\u087a\u0878\u0001\u0000\u0000\u0000\u087b\u087d\u0003\u017a\u00bd\u0000"+
		"\u087c\u087b\u0001\u0000\u0000\u0000\u087c\u087d\u0001\u0000\u0000\u0000"+
		"\u087d\u0139\u0001\u0000\u0000\u0000\u087e\u0880\u0003\u011c\u008e\u0000"+
		"\u087f\u0881\u0003\u00be_\u0000\u0880\u087f\u0001\u0000\u0000\u0000\u0880"+
		"\u0881\u0001\u0000\u0000\u0000\u0881\u013b\u0001\u0000\u0000\u0000\u0882"+
		"\u0885\u0005J\u0000\u0000\u0883\u0885\u0003\u0144\u00a2\u0000\u0884\u0882"+
		"\u0001\u0000\u0000\u0000\u0884\u0883\u0001\u0000\u0000\u0000\u0885\u013d"+
		"\u0001\u0000\u0000\u0000\u0886\u0889\u0005J\u0000\u0000\u0887\u0889\u0003"+
		"\u0146\u00a3\u0000\u0888\u0886\u0001\u0000\u0000\u0000\u0888\u0887\u0001"+
		"\u0000\u0000\u0000\u0889\u013f\u0001\u0000\u0000\u0000\u088a\u088f\u0003"+
		"\u013c\u009e\u0000\u088b\u088c\u0005\u001e\u0000\u0000\u088c\u088e\u0003"+
		"\u013c\u009e\u0000\u088d\u088b\u0001\u0000\u0000\u0000\u088e\u0891\u0001"+
		"\u0000\u0000\u0000\u088f\u088d\u0001\u0000\u0000\u0000\u088f\u0890\u0001"+
		"\u0000\u0000\u0000\u0890\u0141\u0001\u0000\u0000\u0000\u0891\u088f\u0001"+
		"\u0000\u0000\u0000\u0892\u0893\u0007\u0001\u0000\u0000\u0893\u0143\u0001"+
		"\u0000\u0000\u0000\u0894\u0895\u0007\u0002\u0000\u0000\u0895\u0145\u0001"+
		"\u0000\u0000\u0000\u0896\u0897\u0007\u0003\u0000\u0000\u0897\u0147\u0001"+
		"\u0000\u0000\u0000\u0898\u089d\u0005Z\u0000\u0000\u0899\u089a\u0005Y\u0000"+
		"\u0000\u089a\u089c\u0005Z\u0000\u0000\u089b\u0899\u0001\u0000\u0000\u0000"+
		"\u089c\u089f\u0001\u0000\u0000\u0000\u089d\u089b\u0001\u0000\u0000\u0000"+
		"\u089d\u089e\u0001\u0000\u0000\u0000\u089e\u0149\u0001\u0000\u0000\u0000"+
		"\u089f\u089d\u0001\u0000\u0000\u0000\u08a0\u08a5\u0005[\u0000\u0000\u08a1"+
		"\u08a2\u0005Y\u0000\u0000\u08a2\u08a4\u0005[\u0000\u0000\u08a3\u08a1\u0001"+
		"\u0000\u0000\u0000\u08a4\u08a7\u0001\u0000\u0000\u0000\u08a5\u08a3\u0001"+
		"\u0000\u0000\u0000\u08a5\u08a6\u0001\u0000\u0000\u0000\u08a6\u014b\u0001"+
		"\u0000\u0000\u0000\u08a7\u08a5\u0001\u0000\u0000\u0000\u08a8\u08a9\u0005"+
		"/\u0000\u0000\u08a9\u014d\u0001\u0000\u0000\u0000\u08aa\u08ab\u0005.\u0000"+
		"\u0000\u08ab\u014f\u0001\u0000\u0000\u0000\u08ac\u08ad\u0005>\u0000\u0000"+
		"\u08ad\u0151\u0001\u0000\u0000\u0000\u08ae\u08af\u0005?\u0000\u0000\u08af"+
		"\u0153\u0001\u0000\u0000\u0000\u08b0\u08b1\u0005@\u0000\u0000\u08b1\u0155"+
		"\u0001\u0000\u0000\u0000\u08b2\u08b8\u0003\u0152\u00a9\u0000\u08b3\u08b8"+
		"\u0003\u0154\u00aa\u0000\u08b4\u08b8\u0003\u015c\u00ae\u0000\u08b5\u08b8"+
		"\u0005\u0001\u0000\u0000\u08b6\u08b8\u0005\u0014\u0000\u0000\u08b7\u08b2"+
		"\u0001\u0000\u0000\u0000\u08b7\u08b3\u0001\u0000\u0000\u0000\u08b7\u08b4"+
		"\u0001\u0000\u0000\u0000\u08b7\u08b5\u0001\u0000\u0000\u0000\u08b7\u08b6"+
		"\u0001\u0000\u0000\u0000\u08b8\u0157\u0001\u0000\u0000\u0000\u08b9\u08bc"+
		"\u0003\u015c\u00ae\u0000\u08ba\u08bc\u0005\u0012\u0000\u0000\u08bb\u08b9"+
		"\u0001\u0000\u0000\u0000\u08bb\u08ba\u0001\u0000\u0000\u0000\u08bc\u0159"+
		"\u0001\u0000\u0000\u0000\u08bd\u08be\u0007\u0004\u0000\u0000\u08be\u015b"+
		"\u0001\u0000\u0000\u0000\u08bf\u08c1\u0003\u0162\u00b1\u0000\u08c0\u08c2"+
		"\u0003\u015e\u00af\u0000\u08c1\u08c0\u0001\u0000\u0000\u0000\u08c1\u08c2"+
		"\u0001\u0000\u0000\u0000\u08c2\u08cb\u0001\u0000\u0000\u0000\u08c3\u08c7"+
		"\u0003\u0164\u00b2\u0000\u08c4\u08c6\u0003\u0166\u00b3\u0000\u08c5\u08c4"+
		"\u0001\u0000\u0000\u0000\u08c6\u08c9\u0001\u0000\u0000\u0000\u08c7\u08c5"+
		"\u0001\u0000\u0000\u0000\u08c7\u08c8\u0001\u0000\u0000\u0000\u08c8\u08cb"+
		"\u0001\u0000\u0000\u0000\u08c9\u08c7\u0001\u0000\u0000\u0000\u08ca\u08bf"+
		"\u0001\u0000\u0000\u0000\u08ca\u08c3\u0001\u0000\u0000\u0000\u08cb\u015d"+
		"\u0001\u0000\u0000\u0000\u08cc\u08cd\u0004\u00af\b\u0000\u08cd\u08cf\u0003"+
		"\u0160\u00b0\u0000\u08ce\u08cc\u0001\u0000\u0000\u0000\u08cf\u08d0\u0001"+
		"\u0000\u0000\u0000\u08d0\u08ce\u0001\u0000\u0000\u0000\u08d0\u08d1\u0001"+
		"\u0000\u0000\u0000\u08d1\u015f\u0001\u0000\u0000\u0000\u08d2\u08d5\u0003"+
		"\u0162\u00b1\u0000\u08d3\u08d5\u0005]\u0000\u0000\u08d4\u08d2\u0001\u0000"+
		"\u0000\u0000\u08d4\u08d3\u0001\u0000\u0000\u0000\u08d5\u0161\u0001\u0000"+
		"\u0000\u0000\u08d6\u08d9\u0007\u0005\u0000\u0000\u08d7\u08d9\u0005K\u0000"+
		"\u0000\u08d8\u08d6\u0001\u0000\u0000\u0000\u08d8\u08d7\u0001\u0000\u0000"+
		"\u0000\u08d9\u0163\u0001\u0000\u0000\u0000\u08da\u08db\u0005\u001e\u0000"+
		"\u0000\u08db\u0165\u0001\u0000\u0000\u0000\u08dc\u08df\u0005\u001e\u0000"+
		"\u0000\u08dd\u08df\u0003\u0160\u00b0\u0000\u08de\u08dc\u0001\u0000\u0000"+
		"\u0000\u08de\u08dd\u0001\u0000\u0000\u0000\u08df\u0167\u0001\u0000\u0000"+
		"\u0000\u08e0\u08e5\u0003\u016e\u00b7\u0000\u08e1\u08e5\u0003\u0172\u00b9"+
		"\u0000\u08e2\u08e5\u0003\u016a\u00b5\u0000\u08e3\u08e5\u0003\u016c\u00b6"+
		"\u0000\u08e4\u08e0\u0001\u0000\u0000\u0000\u08e4\u08e1\u0001\u0000\u0000"+
		"\u0000\u08e4\u08e2\u0001\u0000\u0000\u0000\u08e4\u08e3\u0001\u0000\u0000"+
		"\u0000\u08e5\u0169\u0001\u0000\u0000\u0000\u08e6\u08e7\u0007\u0006\u0000"+
		"\u0000\u08e7\u016b\u0001\u0000\u0000\u0000\u08e8\u08e9\u0005\u0013\u0000"+
		"\u0000\u08e9\u016d\u0001\u0000\u0000\u0000\u08ea\u08ec\u0003\u014e\u00a7"+
		"\u0000\u08eb\u08ea\u0001\u0000\u0000\u0000\u08eb\u08ec\u0001\u0000\u0000"+
		"\u0000\u08ec\u08ed\u0001\u0000\u0000\u0000\u08ed\u08f3\u0003\u0170\u00b8"+
		"\u0000\u08ee\u08f0\u0003\u014e\u00a7\u0000\u08ef\u08ee\u0001\u0000\u0000"+
		"\u0000\u08ef\u08f0\u0001\u0000\u0000\u0000\u08f0\u08f1\u0001\u0000\u0000"+
		"\u0000\u08f1\u08f3\u0005R\u0000\u0000\u08f2\u08eb\u0001\u0000\u0000\u0000"+
		"\u08f2\u08ef\u0001\u0000\u0000\u0000\u08f3\u016f\u0001\u0000\u0000\u0000"+
		"\u08f4\u08f5\u0007\u0007\u0000\u0000\u08f5\u0171\u0001\u0000\u0000\u0000"+
		"\u08f6\u08f7\u0007\b\u0000\u0000\u08f7\u0173\u0001\u0000\u0000\u0000\u08f8"+
		"\u08f9\u0007\t\u0000\u0000\u08f9\u0175\u0001\u0000\u0000\u0000\u08fa\u08fb"+
		"\u0007\n\u0000\u0000\u08fb\u0177\u0001\u0000\u0000\u0000\u08fc\u0900\u0005"+
		"\'\u0000\u0000\u08fd\u08fe\u0003\u014a\u00a5\u0000\u08fe\u08ff\u0005Y"+
		"\u0000\u0000\u08ff\u0901\u0001\u0000\u0000\u0000\u0900\u08fd\u0001\u0000"+
		"\u0000\u0000\u0900\u0901\u0001\u0000\u0000\u0000\u0901\u0907\u0001\u0000"+
		"\u0000\u0000\u0902\u0904\u0003\u014a\u00a5\u0000\u0903\u0902\u0001\u0000"+
		"\u0000\u0000\u0903\u0904\u0001\u0000\u0000\u0000\u0904\u0905\u0001\u0000"+
		"\u0000\u0000\u0905\u0907\u0005Y\u0000\u0000\u0906\u08fc\u0001\u0000\u0000"+
		"\u0000\u0906\u0903\u0001\u0000\u0000\u0000\u0907\u0179\u0001\u0000\u0000"+
		"\u0000\u0908\u090c\u0005%\u0000\u0000\u0909\u090a\u0003\u014a\u00a5\u0000"+
		"\u090a\u090b\u0005Y\u0000\u0000\u090b\u090d\u0001\u0000\u0000\u0000\u090c"+
		"\u0909\u0001\u0000\u0000\u0000\u090c\u090d\u0001\u0000\u0000\u0000\u090d"+
		"\u0913\u0001\u0000\u0000\u0000\u090e\u0910\u0003\u014a\u00a5\u0000\u090f"+
		"\u090e\u0001\u0000\u0000\u0000\u090f\u0910\u0001\u0000\u0000\u0000\u0910"+
		"\u0911\u0001\u0000\u0000\u0000\u0911\u0913\u0005Y\u0000\u0000\u0912\u0908"+
		"\u0001\u0000\u0000\u0000\u0912\u090f\u0001\u0000\u0000\u0000\u0913\u017b"+
		"\u0001\u0000\u0000\u0000\u0144\u017f\u0183\u0188\u0193\u019c\u01a3\u01a7"+
		"\u01ab\u01b2\u01ba\u01c4\u01ce\u01d5\u01da\u01e1\u01e8\u01ef\u01f6\u01fa"+
		"\u0200\u0208\u020c\u0213\u021a\u021e\u0223\u022d\u0234\u023d\u0244\u0249"+
		"\u024e\u0256\u025c\u0263\u026d\u0274\u0278\u0281\u0287\u028e\u0298\u029f"+
		"\u02a4\u02a9\u02af\u02b2\u02bf\u02c5\u02c9\u02ce\u02d8\u02dc\u02e3\u02ee"+
		"\u02f5\u0303\u030d\u0314\u031b\u0321\u0326\u032a\u032e\u0332\u0336\u0340"+
		"\u0347\u034e\u0356\u035d\u0361\u0366\u036b\u0370\u0375\u037b\u0381\u038c"+
		"\u0391\u0396\u03a1\u03a6\u03a9\u03ae\u03b4\u03b8\u03bc\u03c0\u03c5\u03c9"+
		"\u03ce\u03d0\u03d4\u03d9\u03dd\u03e1\u03e6\u03ea\u03ef\u03fa\u0401\u0406"+
		"\u040d\u0414\u0418\u041f\u0423\u0427\u042c\u042e\u0433\u0438\u043c\u0441"+
		"\u0448\u044d\u0451\u0456\u0462\u0469\u046d\u0472\u0477\u047b\u0480\u0484"+
		"\u0487\u048c\u0495\u0499\u049e\u04a2\u04a8\u04ad\u04b1\u04b6\u04c0\u04c7"+
		"\u04cb\u04d0\u04d5\u04dc\u04df\u04e6\u04ea\u04ef\u04f8\u04fc\u0501\u0508"+
		"\u050d\u0511\u0516\u0520\u0527\u052b\u0530\u0535\u0539\u053c\u0540\u0545"+
		"\u054d\u0554\u0558\u055b\u055d\u0562\u0567\u0570\u0576\u057a\u057d\u0582"+
		"\u0586\u058a\u0590\u0597\u05a0\u05a9\u05b8\u05bb\u05bd\u05c2\u05cc\u05d0"+
		"\u05d5\u05dc\u05de\u05e7\u05ee\u05f2\u05f6\u05fb\u0600\u0607\u060e\u0612"+
		"\u0618\u0624\u0629\u0633\u0638\u063e\u0646\u064f\u0658\u065d\u0668\u066c"+
		"\u0671\u067b\u0682\u0686\u068e\u0692\u0697\u06a1\u06a8\u06ac\u06b0\u06b9"+
		"\u06bd\u06c2\u06cc\u06d3\u06d7\u06dc\u06ed\u06f4\u06f8\u06fd\u0703\u070a"+
		"\u0711\u0715\u0719\u0722\u0729\u0736\u073a\u073f\u0743\u074b\u074f\u0754"+
		"\u075e\u0760\u0767\u076b\u0773\u077a\u0780\u0787\u078b\u078f\u0794\u079e"+
		"\u07a7\u07af\u07b1\u07b9\u07be\u07c3\u07ca\u07cf\u07d4\u07d8\u07dd\u07e2"+
		"\u07e9\u07ee\u07f3\u07f5\u07f7\u07fe\u0801\u0805\u080a\u0811\u0816\u081e"+
		"\u0822\u0827\u0831\u0838\u083c\u0840\u0842\u0846\u084c\u0851\u0858\u085d"+
		"\u0867\u0871\u0878\u087c\u0880\u0884\u0888\u088f\u089d\u08a5\u08b7\u08bb"+
		"\u08c1\u08c7\u08ca\u08d0\u08d4\u08d8\u08de\u08e4\u08eb\u08ef\u08f2\u0900"+
		"\u0903\u0906\u090c\u090f\u0912";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}