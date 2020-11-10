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

parser grammar MojoParser;

options {
    tokenVocab=MojoLexer;
}

mojoFile : (EOL* statements)? EOL* EOF;

// Statements

// GRAMMAR OF A STATEMENT
statement
 : declaration
 | expression
 | loopStatement
 | branchStatement
 | controlTransferStatement
 ;

statements
    : statement (eos EOL* statement)* SEMI?
    ;

// GRAMMAR OF A LOOP STATEMENT

loopStatement :
 | forInStatement
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

//repeatWhileStatement : KEYWORD_REPEATE codeBlock KEYWORD_WHILE expression ;

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
matchCases : matchCase (eos EOL* matchCase)* eos;
matchCase : pattern EOL* RIGHT_RIGHT_ARROWS EOL* ( codeBlock | expression )  ;

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
 | typeName typeAnnotation
 ;

// GRAMMAR OF A GENERIC ARGUMENT CLAUSE

genericArgumentClause : LT EOL* genericArguments EOL* GT;
genericArguments : genericArgument (eov EOL* genericArgument)* eov;
genericArgument : type_ attributes?;

// context-sensitive. Allow < as pre, post, or binary op
//lt : {_input.LT(1).getText().equals("<")}? operator ;
//gt : {_input.LT(1).getText().equals(">")}? operator ;
// Declarations

// GRAMMAR OF A DECLARATION

declaration
 : packageDeclaration
 | importDeclaration
 | constantDeclaration
 | variableDeclaration
 | typeAliasDeclaration
 | functionDeclaration
 | enumDeclaration
 | structDeclaration
 | interfaceDeclaration
 ;

//declarations : declaration+ ;

// GRAMMAR OF A CODE BLOCK

codeBlock : LCURLY (EOL* statements)? EOL* RCURLY ;

// GRAMMAR OF A PACKAGE DECLARATION
packageDeclaration : document? (EOL* attributes)? EOL* KEYWORD_PACKAGE packageIdentifier (EOL* objectLiteral)?;

packageIdentifier : Identifier (DOT Identifier)* ;

// GRAMMAR OF AN IMPORT DECLARATION
importDeclaration
  : KEYWORD_IMPORT importPath (importAllClause | importValueAsClause | importTypeClause | importGroupClause)?
  ;
importPath : importPathIdentifier (DOT importPathIdentifier)*  ;
importPathIdentifier : declarationIdentifier;

importAllClause : DOT MUL;

importValueAsClause : KEYWORD_AS declarationIdentifier;

importTypeClause : DOT typeName importTypeAsClause?;
importTypeAsClause : KEYWORD_AS typeName;

importGroupClause : DOT LCURLY EOL* importGroup EOL* RCURLY;
importGroup : (importValue | importType) (eov EOL* (importValue | importType))* eov?;
importValue : declarationIdentifier importValueAsClause?;
importType : typeName importTypeAsClause?;

// GRAMMAR OF A CONSTANT DECLARATION

constantDeclaration
  : document? (EOL* attributes)? EOL* KEYWORD_CONST patternInitializers
  ;
patternInitializers
  : patternInitializer (COMMA patternInitializer)*
  | LCURLY EOL* patternInitializer (COMMA patternInitializer)* EOL* RCURLY
  ;

/** rule is ambiguous. can match "var x = 1" with x as pattern
 *  OR with x as expression_pattern.
 *  ANTLR resolves in favor or first choice: pattern is x, 1 is initializer.
 */
patternInitializer : pattern initializer? ;
initializer : assignmentOperator EOL* expression  ;

variableDeclaration
  : KEYWORD_VAR patternInitializers
  | identifierPattern COLON_EQUAL expression
  ;

// GRAMMAR OF A TYPE ALIAS DECLARATION
typeAliasDeclaration
  : document? (EOL* attributes)? EOL* KEYWORD_TYPE typealiasName genericParameterClause? EOL* typealiasAssignment
  ;

typealiasName : typeName ;
typealiasAssignment : assignmentOperator EOL* type_ ;

// GRAMMAR OF A FUNCTION DECLARATION
functionDeclaration : functionHead functionName genericParameterClause? functionSignature (EOL* functionBody)? ;

functionHead : document? (EOL* attributes)? EOL* KEYWORD_FUNC ;

functionName : declarationIdentifier | operator ;

functionSignature
 : functionParameterClause (EOL* functionResult)?
 ;
 
functionResult : arrowOperator (labelIdentifier COLON)? type_ attributes? (EOL* followingDocument)? ;

functionBody : LCURLY (eov* followingDocument)? (EOL* statements)? EOL* RCURLY ;

functionParameterClause
  : LPAREN RPAREN
  | LPAREN EOL* functionParameters EOL* RPAREN
  ;

functionParameters
  : functionParameter (eovWithDocument EOL* functionParameter)* eovWithDocument?
  ;

functionParameter
 : labelIdentifier typeAnnotation (EOL* initializer)?
 | labelIdentifier COLON type_ ELLIPSIS attributes?
 ;

// GRAMMAR OF AN ENUMERATION DECLARATION

enumDeclaration
  : document? (EOL* attributes)? EOL* KEYWORD_ENUM enumName genericParameterClause? (EOL* typeInheritanceClause)? EOL* enumBody
  ;

enumBody : LCURLY (eov* followingDocument)? (EOL* enumMembers)? EOL* RCURLY ;

enumName: typeName;
enumMembers : enumMember (eovWithDocument EOL* enumMember)* eovWithDocument?;

enumMember
 :  document? (EOL* attributes)? EOL* declarationIdentifier attributes? (EOL* initializer)?
 ;

// GRAMMAR OF A STRUCTURE DECLARATION
structDeclaration
 : document? (EOL* attributes)? EOL* KEYWORD_TYPE structName genericParameterClause? (EOL* typeInheritanceClause)? (EOL* structBody)?
 ;

structName : typeName;

structBody
   : LCURLY (eov* followingDocument)? (EOL* structMembers)? EOL* RCURLY
   ;

structMembers
  : structMember (eosWithDocument EOL* structMember)* eosWithDocument?
  //: structMember eosWithDocument?
  //| structMember eosWithDocument EOL* structMembers
  ;

structMember 
 : structMemberDeclaration
 | structDeclaration
 | enumDeclaration
 | constantDeclaration
 | typeAliasDeclaration
 ;

structMemberDeclaration
  : document? (EOL* attributes)? EOL* declarationIdentifier typeAnnotation (EOL* initializer )?
  ;

// GRAMMAR OF A INTERFACE DECLARATION

interfaceDeclaration
  : document? (EOL* attributes)? EOL* KEYWORD_INTERFACE interfaceName genericParameterClause? (EOL* typeInheritanceClause)? EOL* interfaceBody ;

interfaceName : typeName ;
interfaceBody : LCURLY (eov* followingDocument)? (EOL* interfaceMembers)? EOL* RCURLY ;

interfaceMembers
  : interfaceMember (eosWithDocument EOL* interfaceMember)* eosWithDocument?
  ;

interfaceMember
 : interfaceMethodDeclaration
 | typeAliasDeclaration
 ;

// GRAMMAR OF A INTERFACE METHOD DECLARATION

interfaceMethodDeclaration : document? (EOL* attributes)? EOL* functionName genericParameterClause? EOL* functionSignature;

// Patterns

// GRAMMAR OF A PATTERN

pattern
 : wildcard_pattern typeAnnotation?
 | identifierPattern typeAnnotation?
 | tuple_pattern typeAnnotation?
 //| enum_value_pattern
 | optional_pattern
 | KEYWORD_IS type_
 | pattern KEYWORD_AS type_
 | expression_pattern
 ;

// GRAMMAR OF A WILDCARD PATTERN

wildcard_pattern : UNDERSCORE  ;

// GRAMMAR OF AN IDENTIFIER PATTERN

identifierPattern : declarationIdentifier ;

// GRAMMAR OF A TUPLE PATTERN

tuple_pattern : LPAREN tuple_pattern_element_list? RPAREN  ;
tuple_pattern_element_list
	:	tuple_pattern_element (COMMA tuple_pattern_element)*
	;
tuple_pattern_element : pattern  ;

// GRAMMAR OF AN ENUMERATION CASE PATTERN

//enum_value_pattern : typeIdentifier? DOT enum_case_name tuple_pattern? ;

// GRAMMAR OF AN OPTIONAL PATTERN
optional_pattern : identifierPattern QUESTION ;

// GRAMMAR OF AN EXPRESSION PATTERN
expression_pattern : expression  ;

// Attributes

// GRAMMAR OF AN ATTRIBUTE

attribute
 : '@' attributeName attributeArgumentClause?
 ;

attributeName :  attributeNameIdentifier | DecimalLiteral;
attributeNameIdentifier : Identifier ( DOT Identifier)*;

attributeArgumentClause : LPAREN (EOL*  expressions)? EOL* RPAREN  ;

// GRAMMAR OF AN ATTRIBUTES
attributes : attribute (EOL* attribute)* ;

// Expressions

// GRAMMAR OF AN EXPRESSION
expression : prefixExpression binaryExpressions? ;

expressions : expression (eov EOL* expression)* eov?;

// GRAMMAR OF A PREFIX EXPRESSION

prefixExpression
  : prefixOperator postfixExpression
  | postfixExpression
  ;

// GRAMMAR OF A BINARY EXPRESSION

binaryExpression
  : binaryOperator prefixExpression
  | assignmentOperator prefixExpression
  | conditional_operator prefixExpression
  | type_casting_operator
  ;

binaryExpressions : binaryExpression+ ;

// GRAMMAR OF A CONDITIONAL OPERATOR

conditional_operator : QUESTION expression COLON ;

// GRAMMAR OF A TYPE_CASTING OPERATOR

type_casting_operator
  : KEYWORD_IS type_
  | KEYWORD_AS type_
  ;

// GRAMMAR OF A PRIMARY EXPRESSION

primaryExpression
 : declarationIdentifier genericArgumentClause?
 | literalExpression
 //| closure_expression
 | parenthesizedExpression
 | tupleExpression
 | implicitMemberExpression
 | wildcardExpression
 //| key_path_expression
 ;

// GRAMMAR OF A LITERAL EXPRESSION

literalExpression
 : literal
 | arrayLiteral
 | objectLiteral
 ;

arrayLiteral : LBRACK (EOL* arrayLiteralItems)? EOL* RBRACK ;

arrayLiteralItems
  : arrayLiteralItem (eov EOL* arrayLiteralItem)* eov?
// : arrayLiteralItem eov?
// | arrayLiteralItem eov EOL* arrayLiteralItems
 ;
 
arrayLiteralItem : expression ;

objectLiteral
 : LCURLY (EOL* objectLiteralItems)? EOL* RCURLY
 ;

objectLiteralItems
  : objectLiteralItem (eovWithDocument EOL* objectLiteralItem)* eovWithDocument?
  ;
 
objectLiteralItem : document? expression COLON expression;

// GRAMMAR OF A CLOSURE EXPRESSION

// : LCURLY closure_signature? statement* RCURLY ;

//closure_signature
// : capture_list? closure_parameter_clause functionResult? 'in'
// | capture_list 'in'
// ;

//closure_parameter_clause
// : LPAREN RPAREN
// | LPAREN closure_parameter_list RPAREN
// | closure_parameter_clause_identifier_list
// ;

// Renamed rule "identifier_list"
//closure_parameter_clause_identifier_list : declarationIdentifier (COMMA declarationIdentifier)* ;

//closure_parameter_list : closure_parameter (COMMA closure_parameter)* ;

//closure_parameter
// : closure_parameter_name typeAnnotation?
// | closure_parameter_name typeAnnotation range_operator
// ;

//closure_parameter_name : labelIdentifier ;

//capture_list : LBRACK capture_list_items RBRACK ;

//capture_list_items : capture_list_item (COMMA capture_list_item)* ;

//capture_list_item : capture_specifier? expression ;

//capture_specifier : 'weak' | 'unowned' | 'unowned(safe)' | 'unowned(unsafe)'  ;

// GRAMMAR OF A IMPLICIT MEMBER EXPRESSION

implicitMemberExpression : DOT labelIdentifier ; // let a: MyType = .default; static let `default` = MyType()

// GRAMMAR OF A PARENTHESIZED EXPRESSION

parenthesizedExpression : LPAREN EOL* expression EOL* RPAREN ;

// GRAMMAR OF A TUPLE EXPRESSION

tupleExpression
 : LPAREN RPAREN
 | LPAREN tupleElement (COMMA tupleElement)+ RPAREN
 ;

tupleElement
 : expression
 | labelIdentifier COLON expression
 ;

// GRAMMAR OF A WILDCARD EXPRESSION

wildcardExpression : UNDERSCORE ;

// GRAMMAR OF A POSTFIX EXPRESSION (inlined many rules from spec to avoid indirect left-recursion)

postfixExpression
 : primaryExpression                                                   # primary
 | postfixExpression postfixOperator                                   # postfixOperation
 | postfixExpression functionCallArgumentClause                        # functionCallExpression
 //| postfixExpression functionCallArgumentClause? trailing_closure    # function_call_expression_with_closure
 | postfixExpression DOT PureDecimalDigits                           # explicitMemberExpression1
 | postfixExpression DOT declarationIdentifier genericArgumentClause?          # explicitMemberExpression2
 | postfixExpression DOT declarationIdentifier LPAREN argumentNameList RPAREN            # explicitMemberExpression3
// This does't exist in the swift grammar, but this valid swift statement fails without it
// self.addTarget(self, action: #selector(nameOfAction(_:)))
 | postfixExpression LPAREN argumentNameList RPAREN                           # explicitMemberExpression4
 | postfixExpression LBRACK expressions RBRACK                          # subscriptExpression
// ! is a postfix operator already
// | postfixExpression '!'                                            # forced_value_expression
// ? is a postfix operator already
// | postfixExpression QUESTION                                            # optional_chaining_expression
 ;

// GRAMMAR OF A FUNCTION CALL EXPRESSION

// See the optimization in postfixExpression. It should be doing exactly this:
//
// function-call-expression → postfix-expression­ function-call-argument-clause­
// function-call-expression → postfix-expression­ function-call-argument-clause­?­ trailing-closure

functionCallArgumentClause
 : LPAREN RPAREN
 | LPAREN function_call_argument_list RPAREN
 ;

function_call_argument_list : function_call_argument ( COMMA function_call_argument )* ;

function_call_argument
 : expression
 | labelIdentifier COLON expression
 | operator
 | labelIdentifier COLON operator
 ;

//trailing_closure : closure_expression ;

// GRAMMAR OF AN EXPLICIT MEMBER EXPRESSION
argumentNameList : argument_name (argument_name)* ;
argument_name : labelIdentifier COLON ;

// GRAMMAR OF A TYPE
type_
 : basicType
 | functionType
 | type_ QUESTION
 ;

basicType
 : basicType attributes? EOL* OR EOL* basicType attributes? #Union
 | basicType attributes? EOL* AND EOL* basicType attributes? #Intersection
 | primeType #Prime
 ;

primeType
 : arrayType
 | dictionaryType
 | tupleType
 | typeIdentifier
 ;

// GRAMMAR OF A TYPE ANNOTATION
typeAnnotation : COLON type_ attributes?;

// GRAMMAR OF A TYPE IDENTIFIER
typeIdentifier : (packageIdentifier DOT)? typeIdentifierClause (DOT typeIdentifierClause)* ;
typeIdentifierClause : typeName genericArgumentClause? ;

typeName : TypeName ;

// GRAMMAR OF A TUPLE TYPE
tupleType : LPAREN (EOL* tupleTypeElements)? EOL* RPAREN ;
tupleTypeElements : tupleTypeElement (eovWithDocument EOL* tupleTypeElement)* eovWithDocument?;
tupleTypeElement : ( declarationIdentifier COLON )? type_ attributes? ;

// GRAMMAR OF A UNION TYPE

// GRAMMAR OF A FUNCTION TYPE
functionType
 : functionTypeArgumentClause arrowOperator type_
 ;
 
functionTypeArgumentClause
 : LPAREN RPAREN
 | LPAREN functionTypeArguments range_operator? RPAREN
 ;
 
functionTypeArguments : functionTypeArgument (eov EOL* functionTypeArgument)* eov?
 ;
 
functionTypeArgument
 : labelIdentifier COLON type_ ELLIPSIS attributes?
 | labelIdentifier typeAnnotation
 ;

// GRAMMAR OF AN ARRAY TYPE

arrayType : LBRACK type_ attributes? RBRACK ;

// GRAMMAR OF A DICTIONARY TYPE

dictionaryType : LCURLY type_ COLON type_  attributes? RCURLY ;

// GRAMMAR OF AN OPTIONAL TYPE

// The following sets of rules are mutually left-recursive [mojo_type, optional_type, implicitly_unwrapped_optional_type, metatype_type]
//optional_type : type_ QUESTION ;

// GRAMMAR OF A TYPE INHERITANCE CLAUSE

typeInheritanceClause : COLON EOL* typeInheritances ;
typeInheritances : typeInheritance (eovWithDocument EOL* typeInheritance)* eovWithDocument? ;
typeInheritance : typeIdentifier attributes? ;


// ---------- Lexical Structure -----------

// GRAMMAR OF AN IDENTIFIER

// identifier : Identifier | context_sensitive_keyword ;
//
// identifier is context sensitive

// var x = 1; funx x() {}; class x {}
declarationIdentifier
     : Identifier
     | keyword_as_identifier_in_declarations
     ;

// external, internal argument name
labelIdentifier
     : Identifier
     | keyword_as_identifier_in_labels
     ;

path_identifier : declarationIdentifier ( DOT declarationIdentifier)*;

identifier : Identifier | Implicit_parameter_name
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
keyword_as_identifier_in_declarations
: KEYWORD_AND
| KEYWORD_AS
| 'attribute'
| 'break'
| 'const'
| 'continue'
| 'else'
| 'enum'
| 'false'
| 'for'
| 'func'
| 'if'
| 'import'
| 'in'
| 'interface'
| KEYWORD_IS
| 'match'
| 'not'
| 'null'
| 'or'
| 'package'
| 'repeat'
| 'return'
| 'struct'
| 'true'
| 'type'
| 'var'
| 'while'
| 'xor'
;

// func x(Any: Any)
keyword_as_identifier_in_labels
: 'and'
| KEYWORD_AS
| 'attribute'
| 'break'
| 'const'
| 'continue'
| 'else'
| 'enum'
| 'false'
| 'for'
| 'func'
| 'if'
| 'import'
| 'in'
| 'interface'
| KEYWORD_IS
| 'match'
| 'not'
| 'null'
| 'or'
| 'package'
| 'repeat'
| 'return'
| 'struct'
| 'true'
| 'type'
| 'var'
| 'while'
| 'xor'
 ;

// GRAMMAR A DOCUMENT_COMMENT

document : LineDocument (EOL LineDocument)* EOL;

followingDocument : FollowingLineDocument (EOL FollowingLineDocument)*;

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
 Dictionary<String, Array<Int>>. In this example, the closing >
 characters are not treated as a single token that may then be
 misinterpreted as a bit shift >> operator.
*/

//operator : binaryOperator | prefixOperator | postfixOperator ;

/* these following tokens are also a binaryOperator so much come first as special case */

assignmentOperator : EQUAL ;

/** Need to separate this out from prefixOperator as it's referenced in numericLiteral
 *  as specifically a negation prefix op.
 */
negatePrefixOperator : SUB;

compilation_condition_AND : AND AND ;
compilation_condition_OR  : OR OR ;
compilation_condition_GE  : GT EQUAL ;
arrowOperator	:  RIGHT_ARROW;
range_operator	:  DOT_DOT ;
range_right_open_operator : DOT_DOT_LT;
same_type_equals:  EQUAL EQUAL ;

/**
 "If an operator has whitespace around both sides or around neither side,
 it is treated as a binary operator. As an example, the + operator in a+b
  and a + b is treated as a binary operator."
*/
binaryOperator :  operator ;

/**
 "If an operator has whitespace on the left side only, it is treated as a
 prefix unary operator. As an example, the ++ operator in a ++b is treated
 as a prefix unary operator."
*/
prefixOperator :  operator ;

/**
 "If an operator has whitespace on the right side only, it is treated as a
 postfix unary operator. As an example, the ++ operator in a++ b is treated
 as a postfix unary operator."

 "If an operator has no whitespace on the left but is followed immediately
 by a dot (.), it is treated as a postfix unary operator. As an example,
 the ++ operator in a++.b is treated as a postfix unary operator (a++ .b
 rather than a ++ .b)."
 */
postfixOperator :  operator ;

operator
  : operator_head     (operator_character)*
  | dot_operator_head (dot_operator_character)*
  ;

operator_character
  : operator_head
  | Operator_following_character
  ;

operator_head
  : ('/' | '=' | '-' | '+' | '!' | '*' | '%' | '&' | '|' | '<' | '>' | '^' | '~' | QUESTION) // wrapping in (..) makes it a fast set comparison
  | Operator_head_other
  ;

dot_operator_head 		: DOT ;
dot_operator_character  : DOT | operator_character ;

// GRAMMAR OF A LITERAL

literal : numericLiteral | stringLiteral | BoolLiteral | NullLiteral  ;

numericLiteral
 : negatePrefixOperator? integerLiteral
 | negatePrefixOperator? FloatLiteral
 ;

// GRAMMAR OF AN INTEGER LITERAL

integerLiteral
 : BinaryLiteral
 | OctalLiteral
 | DecimalLiteral
 | PureDecimalDigits
 | HexadecimalLiteral
 ;

// GRAMMAR OF A STRING LITERAL

stringLiteral
  : StaticStringLiteral
  | InterpolatedStringLiteral
  ;

eos : SEMI | EOL+;
eov : COMMA | EOL+;

eosWithDocument
  : (SEMI (EOL* followingDocument)?
  | (EOL* followingDocument)? EOL+)
  ;

eovWithDocument
  : (COMMA (EOL* followingDocument)?
  | (EOL* followingDocument)? EOL+)
  ;