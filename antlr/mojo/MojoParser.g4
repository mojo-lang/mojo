/*
 * [The "BSD license"]
 *  Copyright (c) 2014 Terence Parr
 *  All rights reserved.
 *
 *  Redistribution and use in source and binary forms, with or without
 *  modification, are permitted provided that the following conditions
 *  are met:
 *
 *  1. Redistributions of source code must retain the above copyright
 *     notice, this list of conditions and the following disclaimer.
 *  2. Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *  3. The name of the author may not be used to endorse or promote products
 *     derived from this software without specific prior written permission.
 *
 *  THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 *  IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
 *  OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 *  IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
 *  INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 *  NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *  DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *  THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 *  THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE
 *
 * Converted from Apple's doc, http://tinyurl.com/n8rkoue, to ANTLR's
 * meta-language.
 */
/*
 *
 */
parser grammar MojoParser;

options {
    tokenVocab=MojoLexer;
}

mojoFile : EOL* (statements)? EOL* EOF;

// Statements

// GRAMMAR OF A STATEMENT
statement
 : declaration
 | expression (ifModifier | whileModifier)?
 | loopStatement
 | branchStatement
 | controlTransferStatement
 | floatingStatement
 ;

ifModifier
    : KEYWORD_IF expression
    ;

whileModifier
    : KEYWORD_WHILE expression
    ;

floatingStatement
    : document
    ;

statements
    : statement (eos EOL* statement)* SEMI?
    ;

// GRAMMAR OF A LOOP STATEMENT

loopStatement
    : forInStatement
    | whileStatement
  //| repeatWhileStatement
    ;

// GRAMMAR OF A FOR_IN STATEMENT

forInStatement : KEYWORD_FOR pattern EOL* KEYWORD_IN expression EOL* codeBlock ;

// GRAMMAR OF A WHILE STATEMENT

whileStatement : KEYWORD_WHILE conditions EOL* codeBlock ;

conditions : condition (eov EOL* condition)* ;

condition
 : expression
 | optionalBindingCondition
 ;
 
optionalBindingCondition
 : KEYWORD_VAR pattern EOL* initializer
 ;

// GRAMMAR OF A REPEAT-WHILE STATEMENT

// repeatWhileStatement : KEYWORD_REPEAT codeBlock KEYWORD_WHILE expression ;

// GRAMMAR OF A BRANCH STATEMENT

branchStatement
 : ifStatement
 | matchStatement
 ;

// GRAMMAR OF AN IF STATEMENT

ifStatement
  : KEYWORD_IF conditions EOL* codeBlock EOL* elseClause?
  ;

elseClause
  : KEYWORD_ELSE EOL* codeBlock
  | KEYWORD_ELSE EOL* ifStatement
  ;

// GRAMMAR OF A MATCH STATEMENT

matchStatement : KEYWORD_MATCH expression EOL* LCURLY (EOL* matchCases)? EOL* RCURLY  ;
matchCases : matchCase (eos EOL* matchCase)* eos?;
matchCase : pattern ifModifier? EOL* RIGHT_RIGHT_ARROWS EOL* ( codeBlock | expression )  ;

// GRAMMAR OF A CONTROL TRANSFER STATEMENT

controlTransferStatement
 : breakStatement
 | continueStatement
 | returnStatement
 ;

// GRAMMAR OF A BREAK STATEMENT

breakStatement : KEYWORD_BREAK;

// GRAMMAR OF A CONTINUE STATEMENT

continueStatement : KEYWORD_CONTINUE;

// GRAMMAR OF A RETURN STATEMENT

returnStatement : KEYWORD_RETURN expression? ;

// Generic Parameters and Arguments

// GRAMMAR OF A GENERIC PARAMETER CLAUSE

genericParameterClause : LT EOL* genericParameters EOL* GT;
genericParameters : genericParameter (eovWithDocument EOL* genericParameter)* eovWithDocument?;
genericParameter
 : typeName
 | typeName ELLIPSIS
 | typeName typeAnnotation
 ;

// GRAMMAR OF A GENERIC ARGUMENT CLAUSE
genericArgumentClause : LT EOL* genericArguments EOL* GT;
genericArguments : genericArgument (eov EOL* genericArgument)*;
genericArgument : type_ attributes?;

// context-sensitive. Allow < as pre, post, or binary op
//lt : {_input.LT(1).getText().equals("<")}? operator ;
//gt : {_input.LT(1).getText().equals(">")}? operator ;

// GRAMMAR OF A DECLARATION
declaration : (document EOL)? (attributes EOL*)?
      ( packageDeclaration
      | importDeclaration
      | constantDeclaration
      | variableDeclaration
      | typeAliasDeclaration
      | functionDeclaration
      | enumDeclaration
      | structDeclaration
      | interfaceDeclaration
      | attributeDeclaration
      | attributeAliasDeclaration
      );

// GRAMMAR OF A CODE BLOCK
codeBlock : LCURLY (EOL* statements)? EOL* RCURLY ;

// GRAMMAR OF A PACKAGE DECLARATION
packageDeclaration : KEYWORD_PACKAGE packageIdentifier (EOL* objectLiteral)?;

packageIdentifier : packageName (DOT packageName)* ;
packageName : VALUE_IDENTIFIER;

// GRAMMAR OF AN IMPORT DECLARATION
importDeclaration
  : KEYWORD_IMPORT importPath (importAllClause | importValueAsClause | importTypeClause | importGroupClause)?
  ;
importPath : importPathIdentifier (DOT importPathIdentifier)*  ;
importPathIdentifier : declarationIdentifier;

importAllClause : DOT STAR;

importValueAsClause : KEYWORD_AS declarationIdentifier;

importTypeClause : DOT typeName importTypeAsClause?;
importTypeAsClause : KEYWORD_AS typeName;

importGroupClause : DOT LCURLY EOL* importGroup EOL* RCURLY;
importGroup : (importValue | importType) (eov EOL* (importValue | importType))* eov?;
importValue : declarationIdentifier importValueAsClause?;
importType : typeName importTypeAsClause?;

// GRAMMAR OF A CONSTANT DECLARATION

constantDeclaration
  : KEYWORD_CONST patternInitializers
  | KEYWORD_CONST LCURLY EOL* documentedPatternInitializer (eos EOL* documentedPatternInitializer)* eos? EOL* RCURLY
  ;

//FIXME will has error in golang parser if remove the patternInitializers
patternInitializers
  : patternInitializer
  ;

documentedPatternInitializer : (document EOL)? (attributes EOL)? patternInitializer;

/** rule is ambiguous. can match "var x = 1" with x as pattern
 *  OR with x as expression_pattern.
 *  ANTLR resolves in favor or first choice: pattern is x, 1 is initializer.
 */
patternInitializer : pattern initializer? ;
initializer : assignmentOperator EOL* expression  ;

variableDeclaration
  : KEYWORD_VAR patternInitializers
  | pattern COLON_EQUAL expression
  | KEYWORD_VAR LCURLY EOL* documentedPatternInitializer (eos EOL* documentedPatternInitializer)* eos? EOL* RCURLY
  ;

// GRAMMAR OF A TYPE ALIAS DECLARATION
typeAliasDeclaration
  : KEYWORD_TYPE typeAliasName genericParameterClause? EOL* typeAliasAssignment
  ;

typeAliasName : typeName ;
typeAliasAssignment : assignmentOperator EOL* type_ attributes? followingDocument?;

// GRAMMAR OF A FUNCTION DECLARATION
functionDeclaration : KEYWORD_FUNC functionName genericParameterClause? functionSignature (EOL* functionBody)? ;

functionName : declarationIdentifier | operator ;

functionSignature
 : functionParameterClause followingDocument? (EOL* functionResult)?
 ;
 
functionResult : arrowOperator (labelIdentifier COLON?)? type_ attributes? (EOL* followingDocument)? ;

functionBody : LCURLY followingDocument? (EOL* statements)? EOL* RCURLY ;

functionParameterClause
  : LPAREN RPAREN
  | LPAREN EOL* functionParameters EOL* RPAREN
  ;

functionParameters
  : functionParameter (eovWithDocument EOL* functionParameter)* eovWithDocument?
  ;

functionParameter
 : labelIdentifier typeAnnotation (EOL* initializer)?
 | labelIdentifier COLON? type_ ELLIPSIS attributes?
 ;

// GRAMMAR OF AN ENUMERATION DECLARATION

enumDeclaration
  : KEYWORD_ENUM enumName genericParameterClause? (EOL* typeInheritanceClause)? EOL* enumBody
  | KEYWORD_TYPE enumName assignmentOperator KEYWORD_ENUM genericParameterClause? (EOL* typeInheritanceClause)? EOL* enumBody
  ;

enumBody : LCURLY (followingDocument)? (EOL* enumMembers)? EOL* RCURLY ;

enumName: typeName;
enumMembers : enumMember (eosWithDocument EOL* enumMember)* eosWithDocument?;

enumMember
 : (document EOL)? (attributes EOL)? declarationIdentifier attributes? (EOL* initializer)?
 | floatingStatement
 ;

// GRAMMAR OF A STRUCTURE DECLARATION
structDeclaration
 : (KEYWORD_TYPE | KEYWORD_STRUCT) structName genericParameterClause? structType
 | KEYWORD_TYPE structName assignmentOperator KEYWORD_STRUCT genericParameterClause? structType
 ;

structName : typeName;
structType : (EOL* typeInheritanceClause)? (EOL* structBody)?;

structBody
   : LCURLY (followingDocument EOL)? (EOL* structMembers)? EOL* RCURLY
   ;

structMembers
  : structMember (eosWithDocument EOL* structMember)* eosWithDocument?
  //: structMember eosWithDocument?
  //| structMember eosWithDocument EOL* structMembers
  ;

structMember
 : (document EOL)? (attributes EOL)?
 ( structDeclaration
 | enumDeclaration
 | constantDeclaration
 | typeAliasDeclaration
 | structMemberDeclaration
 )
 | floatingStatement
 ;

structMemberDeclaration
  : declarationIdentifier typeAnnotation (EOL* initializer )?
  ;

// GRAMMAR OF A INTERFACE DECLARATION

interfaceDeclaration
  : KEYWORD_INTERFACE interfaceName genericParameterClause? interfaceType
  | KEYWORD_TYPE interfaceName assignmentOperator KEYWORD_INTERFACE genericParameterClause? interfaceType
  ;

interfaceName : typeName ;
interfaceType : (EOL* typeInheritanceClause)? EOL* interfaceBody;

interfaceBody : LCURLY followingDocument? (EOL* interfaceMembers)? EOL* RCURLY ;

interfaceMembers
  : interfaceMember (eosWithDocument EOL* interfaceMember)* eosWithDocument?
  ;

interfaceMember
 : (document EOL)? (attributes EOL)?
 ( typeAliasDeclaration
 | interfaceMethodDeclaration
 )
 | floatingStatement
 ;

// GRAMMAR OF A INTERFACE METHOD DECLARATION
interfaceMethodDeclaration : functionName genericParameterClause? EOL* functionSignature;

// GRAMMAR OF A ATTRIBUTE DECLARATION
attributeDeclaration
  : KEYWORD_ATTRIBUTE attributeName genericParameterClause? (structBody | typeAnnotation (EOL* initializer )? followingDocument?);

attributeAliasDeclaration
  : KEYWORD_ATTRIBUTE attributeName  genericParameterClause? EOL* attributeAliasAssignment;

attributeAliasAssignment : assignmentOperator EOL* (packageIdentifier DOT)? attributeName genericArgumentClause? followingDocument?;

// Patterns

// GRAMMAR OF A PATTERN
pattern
 : wildcardPattern typeAnnotation?
 | identifierPattern typeAnnotation?
 | tuplePattern typeAnnotation?
 | arrayPattern typeAnnotation?
 | type_ attributes?
 | enumValuePattern
 | optionalPattern
 | BANG? KEYWORD_IS type_
 | pattern KEYWORD_AS type_
 | expressionPattern
 ;

// GRAMMAR OF A WILDCARD PATTERN

wildcardPattern : UNDERSCORE  ;

// GRAMMAR OF AN IDENTIFIER PATTERN

identifierPattern : declarationIdentifier ;

// GRAMMAR OF A TUPLE PATTERN

tuplePattern : LPAREN tuplePatternElementList? RPAREN  ;
tuplePatternElementList
	:	tuplePatternElement (COMMA tuplePatternElement)*
	;
tuplePatternElement : pattern | ELLIPSIS ;

arrayPattern : LBRACK arrayPatternElements? RBRACK;
arrayPatternElements
    : arrayPatternElement (COMMA arrayPatternElement)*
    ;
arrayPatternElement: pattern | ELLIPSIS;

// GRAMMAR OF AN ENUMERATION CASE PATTERN

enumValuePattern : typeIdentifier? DOT declarationIdentifier tuplePattern? ;

// GRAMMAR OF AN OPTIONAL PATTERN
optionalPattern : identifierPattern QUESTION ;

// GRAMMAR OF AN EXPRESSION PATTERN
expressionPattern : expression  ;

// Attributes

// GRAMMAR OF AN ATTRIBUTE
attribute
 : '@' DECIMAL_LITERAL
 | '@' attributeIdentifier genericArgumentClause? attributeArgumentClause?
 ;

attributeIdentifier : ((packageIdentifier DOT)? attributeName);
attributeName : labelIdentifier;

attributeArgumentClause
    : LPAREN (EOL* attributeArguments)? EOL* RPAREN
    ;

attributeArgument : (labelIdentifier COLON)? expression;
attributeArguments : attributeArgument (eov EOL* attributeArgument)* eov?;

// GRAMMAR OF AN ATTRIBUTES
attributes : attribute (EOL? attribute)* ;

// Expressions

// GRAMMAR OF AN EXPRESSION
expression
    : prefixExpression binaryExpressions?
    ;

//expressions : expression (eov EOL* expression)* eov?;

// GRAMMAR OF A PREFIX EXPRESSION
prefixExpression
  : prefixOperator postfixExpression
  | postfixExpression
  ;

// GRAMMAR OF A BINARY EXPRESSION
binaryExpression
  : binaryOperator prefixExpression
  //| COMMA prefixExpression // comma-expression / tuple-expression
  //| assignmentOperator prefixExpression
  | conditionalOperator prefixExpression // conditional-expression
  | inOperator prefixExpression  // in-expression
  | ifOperator prefixExpression // if-expression
  //| typeCastingOperator // type-casting-expression
  | infixCallOperator prefixExpression // infix-call-expression
  ;

prefixCallOperator
    : labelIdentifier
    ;

infixCallOperator
    : VALUE_IDENTIFIER
  ;

binaryExpressions : binaryExpression+ ;

inOperator
    : BANG KEYWORD_IN | KEYWORD_IN
    ;

// GRAMMAR OF A CONDITIONAL OPERATOR
conditionalOperator : QUESTION expression COLON ;
ifOperator: KEYWORD_IF expression KEYWORD_ELSE;

// GRAMMAR OF A TYPE_CASTING OPERATOR
// is pattern
typeCastingOperator
  : (BANG KEYWORD_IS | KEYWORD_IS) (type_ | pattern)
  | KEYWORD_AS type_
  ;

// GRAMMAR OF A PRIMARY EXPRESSION
primaryExpression
 : literalExpression
 | declarationIdentifier genericArgumentClause?
 | typeIdentifier DOT declarationIdentifier genericArgumentClause?
 | closureExpression
 | tupleLiteralExpression
 | parenthesizedExpression
 | implicitMemberExpression
 | wildcardExpression
 | structConstructionExpression
 | ELLIPSIS
 //| key_path_expression
 ;

// GRAMMAR OF A LITERAL EXPRESSION
literalExpression
 : numericOperatorLiteral
 | stringOperatorLiteral
 | structLiteral
 | literal
 | arrayLiteral
 | mapLiteral
 | objectLiteral
 ;

numericOperatorLiteral : numericLiteral suffixLiteralOperator ;
stringOperatorLiteral
  : prefixLiteralOperator {p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? stringLiteral suffixLiteralOperator?
  | stringLiteral suffixLiteralOperator
  ;

suffixLiteralOperator
  : {p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? (TYPE_IDENTIFIER | VALUE_IDENTIFIER)
  ;

prefixLiteralOperator
  : VALUE_IDENTIFIER;

arrayLiteral : LBRACK (EOL* arrayLiteralItems)? EOL* RBRACK ;

arrayLiteralItems
  : arrayLiteralItem (eov EOL* arrayLiteralItem)* eov?
 ;
 
arrayLiteralItem : expression | ELLIPSIS ;

mapLiteral
   : LCURLY (EOL* mapLiteralItems)? EOL* RCURLY
   ;

mapLiteralItems
    : (mapLiteralItem (eov EOL* mapLiteralItem)* eov?)
    //| (mapLiteralIntegerItem (eov EOL* mapLiteralIntegerItem)* eov?)
    ;

mapLiteralItem
    : (stringLiteral | integerLiteral) COLON expression;

//mapLiteralStringItem
//    : stringLiteral COLON expression;
//
//mapLiteralIntegerItem
//    : integerLiteral COLON expression
//    ;

objectLiteral
 : LCURLY (EOL* objectLiteralItems)? EOL* RCURLY
 ;

objectLiteralItems
  : objectLiteralItem (eov EOL* objectLiteralItem)* eov?
  ;

objectLiteralItem
    : pathIdentifier (COLON expression)?
    ;

structLiteral
    : typeIdentifier objectLiteral
    ;

structConstructionExpression
    : typeIdentifier functionCallSuffix
    ;

//matchExpression : KEYWORD_MATCH expression EOL* LCURLY (EOL* matchExprCases)? EOL* RCURLY;
matchExprSuffix : KEYWORD_MATCH EOL* LCURLY (EOL* matchExprCases)? EOL* RCURLY;
matchExprCases : matchExprCase (eos EOL* matchExprCase)* eos?;
matchExprCase : pattern EOL* RIGHT_RIGHT_ARROWS EOL* expression;

// GRAMMAR OF A CLOSURE EXPRESSION
closureExpression
    : LCURLY EOL* statements EOL* RCURLY
    | LCURLY closureParameters EOL* RIGHT_ARROW (EOL* type_)? EOL* statements RCURLY
    ;

closureParameters
    : closureParameter (eov EOL* closureParameter)* eov?
    ;

closureParameter
    : functionParameter
    | labelIdentifier
    ;

// GRAMMAR OF A IMPLICIT MEMBER EXPRESSION

implicitMemberExpression : DOT labelIdentifier ; // let a: MyType = .default; static let `default` = MyType()

// GRAMMAR OF A PARENTHESIZED EXPRESSION

parenthesizedExpression
    : LPAREN EOL* (
        expression //|
        //expression GRAPH_RIGHT_PATH expression
      ) EOL* RPAREN ;
// GRAMMAR OF A TUPLE EXPRESSION

tupleLiteralExpression
 : LPAREN RPAREN
 | LPAREN tupleElement (COMMA tupleElement)+ RPAREN
 ;

tupleElement
 : expression
 | labelIdentifier COLON expression
 | ELLIPSIS
 ;

// GRAMMAR OF A WILDCARD EXPRESSION

wildcardExpression : UNDERSCORE ;

// GRAMMAR OF A POSTFIX EXPRESSION
postfixExpression: primaryExpression suffixExpression* postfixOperator?;

suffixExpression
    : functionCallSuffix
    | explicitMemberSuffix
    | subscriptSuffix
    | matchExprSuffix
    | typeCastingOperator
    ;

explicitMemberSuffix
    : DOT
		 ( DECIMAL_LITERAL
		| identifier (
			genericArgumentClause
			| LPAREN argumentNames RPAREN
		)?
	     )
	;

subscriptSuffix: LBRACK functionCallArguments RBRACK;

// GRAMMAR OF A FUNCTION CALL EXPRESSION
functionCallSuffix
    : functionCallArgumentClause? trailingClosures
    | functionCallArgumentClause
	;

functionCallArgumentClause
 : LPAREN EOL* RPAREN
 | LPAREN EOL* functionCallArguments EOL* RPAREN
 ;

functionCallArguments : functionCallArgument ( COMMA functionCallArgument )* ;

functionCallArgument
 : expression
 | labelIdentifier COLON expression
 //| operator
 //| labelIdentifier COLON operator
 ;

trailingClosures:
	closureExpression labeledTrailingClosures?;

labeledTrailingClosures: labeledTrailingClosure+;

labeledTrailingClosure: identifier COLON closureExpression;

// GRAMMAR OF AN EXPLICIT MEMBER EXPRESSION
argumentNames : argumentName (argumentName)* ;
argumentName : labelIdentifier COLON ;

// GRAMMAR OF A TYPE
type_
 : basicType
 | functionType
 | type_ BANG     // unit
 | type_ QUESTION // optional
 | type_ STAR     // reference
 | type_ PLUS     // public  RESERVED
 | type_ MINUS    // private RESERVED
 | type_ ELLIPSIS // type list
 ;

basicType
 : basicType attributes? (followingDocument EOL)? EOL* PIPE EOL* basicType attributes? (followingDocument EOL)? #Union
 | basicType attributes? (followingDocument EOL)? EOL* AND EOL* basicType attributes? (followingDocument EOL)?  #Intersection
 | primeType #Prime
 ;

primeType
 : arrayType
 | mapType
 | tupleType
 | typeIdentifier
 ;

// GRAMMAR OF A TYPE ANNOTATION
typeAnnotation : COLON? type_ attributes?;

// GRAMMAR OF A TYPE IDENTIFIER
typeIdentifier : (packageIdentifier DOT)? typeIdentifierClause (DOT typeIdentifierClause)* ;
typeIdentifierClause : typeName genericArgumentClause? ;

typeName : TYPE_IDENTIFIER ;

// GRAMMAR OF A TUPLE TYPE
tupleType : LPAREN (EOL* tupleTypeElements)? EOL* RPAREN ;
tupleTypeElements : tupleTypeElement (eovWithDocument EOL* tupleTypeElement)* eovWithDocument?;
tupleTypeElement : ( declarationIdentifier COLON? )? type_ attributes? ;

// GRAMMAR OF A UNION TYPE

// GRAMMAR OF A FUNCTION TYPE
functionType
 : functionParameterClause arrowOperator type_ attributes?
 ;

// GRAMMAR OF AN ARRAY TYPE

arrayType : LBRACK type_ attributes? RBRACK ;

// GRAMMAR OF A DICTIONARY TYPE

mapType : LCURLY type_ keyAttributes? COLON? type_  attributes? RCURLY ;

keyAttributes : attributes;

// GRAMMAR OF AN OPTIONAL TYPE

// The following sets of rules are mutually left-recursive [mojo_type, optional_type, implicitly_unwrapped_optional_type, metatype_type]
//optional_type : type_ QUESTION ;

// GRAMMAR OF A TYPE INHERITANCE CLAUSE

typeInheritanceClause : COLON EOL* typeInheritances ;
typeInheritances : typeInheritance (eovWithDocument EOL* typeInheritance)* eovWithDocument? ;
typeInheritance : basicType attributes? ;


// ---------- Lexical Structure -----------

// GRAMMAR OF AN IDENTIFIER

// identifier : Identifier | context_sensitive_keyword ;
//
// identifier is context sensitive

// var x = 1; funx x() {}; class x {}
declarationIdentifier
     : VALUE_IDENTIFIER
     | keywordAsIdentifierInDeclarations
     ;

// external, internal argument name
labelIdentifier
     : VALUE_IDENTIFIER
     | keywordAsIdentifierInLabels
     ;

pathIdentifier : declarationIdentifier ( DOT declarationIdentifier)*;

identifier : VALUE_IDENTIFIER | IMPLICIT_PARAMETER_NAME
 ;

// identifier_list : identifier (COMMA identifier)* ;
// 
// identifier is context sensitive
// See: closure_parameter_clause_identifier_list

// Keywords reserved in particular contexts: associativity, convenience, dynamic, didSet, final, get, infix, indirect, lazy, left, mutating, none, nonmutating, optional, override, postfix, precedence, prefix, Protocol, required, right, set, Type, unowned, weak, and willSet. Outside the context in which they appear in the grammar, they can be used as identifiers.
// context_sensitive_keyword :
//  'associativity' | 'convenience' | 'dynamic' | 'didSet'
//  | 'final' | 'get' | 'infix' | 'indirect'
//  | 'lazy' | 'left' | 'mutating' | 'none'
//  | 'nonmutating' | 'optional' | 'override' | 'postfix'
//  | 'precedence' | 'prefix' | 'Protocol' | 'required'
//  | 'right' | 'set' | 'Type' | 'unowned'
//  | 'weak' | 'willSet'
//
// ^- this does not work. "[10].index(of: 10)". "of" is a keyword. "mojo_type(of: self)"
 
 // Added by myself.
 // Tested all alphanumeric tokens in playground.
 // E.g. "let mutating = 1".
 // E.g. "func mutating() {}".
 //
 // In source code of swift there are multiple cases of error diag::keyword_cant_be_identifier.
 // Maybe it is not even a single error when keyword can't be identifier.
 //
keywordAsIdentifierInDeclarations
    : KEYWORD_AND
    | KEYWORD_AS
    | KEYWORD_ATTRIBUTE
    | KEYWORD_BREAK
    | KEYWORD_CONST
    | KEYWORD_CONTINUE
    | KEYWORD_ELSE
    | KEYWORD_ENUM
    | KEYWORD_FALSE
    | KEYWORD_FUNC
    //| KEYWORD_IF
    | KEYWORD_IMPORT
    | KEYWORD_IN
    | KEYWORD_INTERFACE
    | KEYWORD_IS
    | KEYWORD_MATCH
    | KEYWORD_NOT
    | KEYWORD_NULL
    | KEYWORD_OR
    | KEYWORD_PACKAGE
    | KEYWORD_STRUCT
    | KEYWORD_TRUE
    | KEYWORD_TYPE
    | KEYWORD_XOR
    ;

keywordAsIdentifierInLabels
    : KEYWORD_AND
    | KEYWORD_AS
    | KEYWORD_ATTRIBUTE
    | KEYWORD_BREAK
    | KEYWORD_CONST
    | KEYWORD_CONTINUE
    | KEYWORD_ELSE
    | KEYWORD_ENUM
    | KEYWORD_FALSE
    | KEYWORD_FOR
    | KEYWORD_FUNC
    //| KEYWORD_IF
    | KEYWORD_ELSE
    | KEYWORD_IMPORT
    | KEYWORD_IN
    | KEYWORD_INTERFACE
    | KEYWORD_IS
    | KEYWORD_MATCH
    | KEYWORD_NOT
    | KEYWORD_NULL
    | KEYWORD_OR
    | KEYWORD_PACKAGE
    | KEYWORD_RETURN
    | KEYWORD_STRUCT
    | KEYWORD_TRUE
    | KEYWORD_TYPE
    | KEYWORD_VAR
    | KEYWORD_WHILE
    | KEYWORD_XOR
    ;

// GRAMMAR A DOCUMENT_COMMENT

document : LINE_DOCUMENT (EOL LINE_DOCUMENT)*;

followingDocument : FOLLOWING_LINE_DOCUMENT (EOL FOLLOWING_LINE_DOCUMENT)*;

// GRAMMAR OF OPERATORS

/*
From doc on operators:
 The tokens =, ->, //, /*, *\/ [without the escape on \/], .,
 the prefix operators <, &, and ?, the infix
 operator ?, and the postfix operators >, !, and ? are reserved. These tokens
 can’t be overloaded, nor can they be used as custom operators.

 The whitespace around an operator is used to determine whether an operator
 is used as a prefix operator, a postfix operator, or a binary operator.

	* If an operator has whitespace around both sides or around neither
	  side, it is treated as a binary operator. As an example, the +
	  operator in a+b and a + b is treated as a binary operator.

	* If an operator has whitespace on the left side only, it is treated
	  as a prefix unary operator. As an example, the ++ operator in a ++b
	  is treated as a prefix unary operator.

	* If an operator has whitespace on the right side only, it is treated
	  as a postfix unary operator. As an example, the ++ operator in a++ b
	  is treated as a postfix unary operator.

	* If an operator has no whitespace on the left but is followed
	  immediately by a dot (.), it is treated as a postfix unary
	  operator. As an example, the ++ operator in a++.b is treated as a
	  postfix unary operator (a++ .b rather than a ++ .b).

 For the purposes of these rules, the characters (, [, and { before an operator,
 the characters ), ], and } after an operator, and the characters ,, ;, and :
 are also considered whitespace.

 There is one caveat to the rules above. If the ! or ? predefined operator has
 no whitespace on the left, it is treated as a postfix operator, regardless of
 whether it has whitespace on the right. To use the ? as the optional-chaining
 operator, it must not have whitespace on the left. To use it in the ternary
 conditional (? :) operator, it must have whitespace around both sides.

 In certain constructs, operators with a leading < or > may be split
 into two or more tokens. The remainder is treated the same way and may
 be split again. As a result, there is no need to use whitespace to
 disambiguate between the closing > characters in constructs like
 Map<String, Array<Int>>. In this example, the closing >
 characters are not treated as a single token that may then be
 misinterpreted as a bit shift >> operator.
*/

/* these following tokens are also a binaryOperator so much come first as special case */

assignmentOperator : EQUAL ;

/** Need to separate this out from prefixOperator as it's referenced in numericLiteral
 *  as specifically a negation prefix op.
 */
negatePrefixOperator : MINUS;

arrowOperator	      : RIGHT_ARROW;
rangeOperator	      : DOT_DOT ;
halfOpenRangeOperator : DOT_DOT_LT;

/**
 "If an operator has whitespace around both sides or around neither side,
 it is treated as a binary operator. As an example, the + operator in a+b
  and a + b is treated as a binary operator."
*/
binaryOperator
    : rangeOperator
    | halfOpenRangeOperator
    | operator
    | KEYWORD_AND
    | KEYWORD_OR
    ;

/**
 "If an operator has whitespace on the left side only, it is treated as a
 prefix unary operator. As an example, the ++ operator in a ++b is treated
 as a prefix unary operator."
*/
prefixOperator
    : operator
    | KEYWORD_NOT
    ;

/**
 "If an operator has whitespace on the right side only, it is treated as a
 postfix unary operator. As an example, the ++ operator in a++ b is treated
 as a postfix unary operator."

 "If an operator has no whitespace on the left but is followed immediately
 by a dot (.), it is treated as a postfix unary operator. As an example,
 the ++ operator in a++.b is treated as a postfix unary operator (a++ .b
 rather than a ++ .b)."
 */
postfixOperator: PLUS_PLUS | MINUS_MINUS | ELLIPSIS;

operator
  : operator_head     operator_characters?
  | dot_operator_head (dot_operator_character)+
  ;

operator_characters: (
		{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? operator_character
)+;

operator_character
  : operator_head
  | OPERATOR_FOLLOWING_CHARACTER
  ;

operator_head
  : ('/' | '=' | '-' | '+' | '!' | '*' | '%' | '&' | '|' | '<' | '>' | '^' | '~' | QUESTION) // wrapping in (..) makes it a fast set comparison
  | OPERATOR_HEAD_OTHER
  ;

dot_operator_head 		: DOT ;
dot_operator_character  : DOT | operator_character ;

// GRAMMAR OF A LITERAL
literal : numericLiteral | stringLiteral | boolLiteral | nullLiteral  ;

boolLiteral : KEYWORD_TRUE | KEYWORD_FALSE ;

nullLiteral : KEYWORD_NULL ;

numericLiteral
 : negatePrefixOperator? integerLiteral
 | negatePrefixOperator? FLOAT_LITERAL
 ;

// GRAMMAR OF AN INTEGER LITERAL
integerLiteral
 : BINARY_LITERAL
 | OCTAL_LITERAL
 | DECIMAL_LITERAL
 | PURE_DECIMAL_DIGITS
 | HEXADECIMAL_LITERAL
 ;

// GRAMMAR OF A STRING LITERAL
stringLiteral
  : STATIC_STRING_LITERAL
  | INTERPOLATED_STRING_LITERAL
  ;

eos : SEMI | EOL;
eov : COMMA | EOL;

eosWithDocument
  : SEMI (followingDocument EOL)?
  | followingDocument? EOL
  ;

eovWithDocument
  : COMMA (followingDocument EOL)?
  | followingDocument? EOL
  ;
