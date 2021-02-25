/*
 * A Mojo grammar for ANTLR 4
 */

lexer grammar MojoLexer;

// Keywords
KEYWORD_AND       : 'and';
KEYWORD_AS        : 'as';
KEYWORD_ATTRIBUTE : 'attribute';
KEYWORD_BREAK     : 'break';
KEYWORD_CONST     : 'const';
KEYWORD_CONTINUE  : 'continue';
KEYWORD_ELSE      : 'else';
KEYWORD_ENUM      : 'enum';
KEYWORD_FALSE     : 'false';
KEYWORD_FOR       : 'for';
KEYWORD_FUNC      : 'func';
KEYWORD_IF        : 'if';
KEYWORD_IMPORT    : 'import';
KEYWORD_IN        : 'in';
KEYWORD_INTERFACE : 'interface';
KEYWORD_IS        : 'is';
KEYWORD_MATCH     : 'match';
KEYWORD_NOT       : 'not';
KEYWORD_NULL      : 'null';
KEYWORD_OR        : 'or';
KEYWORD_PACKAGE   : 'package';
KEYWORD_REPEATE   : 'repeat';
KEYWORD_RETURN    : 'return';
KEYWORD_STRUCT    : 'struct';       // reserved
KEYWORD_TRUE      : 'true';
KEYWORD_TYPE      : 'type';
KEYWORD_VAR       : 'var';
KEYWORD_WHILE     : 'while';
KEYWORD_XOR       : 'xor';


// Punctuation

DOT    	: '.' ;
LCURLY 	: '{' ;
LPAREN 	: '(' ;
LBRACK 	: '[' ;
RCURLY 	: '}' ;
RPAREN 	: ')' ;
RBRACK 	: ']' ;
COMMA  	: ',' ;
COLON  	: ':' ;
SEMI   	: ';' ;
LT 		: '<' ;
GT 		: '>' ;
UNDERSCORE : '_' ;
BANG 	: '!' ;
QUESTION: '?' ;
AT 		: '@' ;
AND 	: '&' ;
SUB 	: '-' ;
EQUAL 	: '=' ;
OR 		: '|' ;
DIV 	: '/' ;
ADD 	: '+' ;
MUL 	: '*' ;
MOD 	: '%' ;
CARET 	: '^' ;
TILDE 	: '~' ;
DOLLER  : '$' ;

COLON_EQUAL : ':' '=';

RIGHT_RIGHT_ARROWS : '=' '>';
RIGHT_ARROW : '-' '>';

DOT_DOT    : '.' '.';
DOT_DOT_LT : '.' '.' '<';

ELLIPSIS : '.' '.' '.';

TYPE_IDENTIFIER
 : [A-Z][a-zA-Z0-9]*
;

VALUE_IDENTIFIER
 : VALUE_IDENTIFIER_HEAD VALUE_IDENTIFIER_CHARACTERS?
;

fragment VALUE_IDENTIFIER_HEAD
 : [a-z]
 | '\u00A8' | '\u00AA' | '\u00AD' | '\u00AF' | [\u00B2-\u00B5] | [\u00B7-\u00BA]
 | [\u00BC-\u00BE] | [\u00C0-\u00D6] | [\u00D8-\u00F6] | [\u00F8-\u00FF]
 | [\u0100-\u02FF] | [\u0370-\u167F] | [\u1681-\u180D] | [\u180F-\u1DBF]
 | [\u1E00-\u1FFF]
 | [\u200B-\u200D] | [\u202A-\u202E] | [\u203F-\u2040] | '\u2054' | [\u2060-\u206F]
 | [\u2070-\u20CF] | [\u2100-\u218F] | [\u2460-\u24FF] | [\u2776-\u2793]
 | [\u2C00-\u2DFF] | [\u2E80-\u2FFF]
 | [\u3004-\u3007] | [\u3021-\u302F] | [\u3031-\u303F] | [\u3040-\uD7FF]
 | [\uF900-\uFD3D] | [\uFD40-\uFDCF] | [\uFDF0-\uFE1F] | [\uFE30-\uFE44]
 | [\uFE47-\uFFFD]
 ;

fragment VALUE_IDENTIFIER_CHARACTER : [0-9A-Z_]
 | [\u0300-\u036F] | [\u1DC0-\u1DFF] | [\u20D0-\u20FF] | [\uFE20-\uFE2F]
 | VALUE_IDENTIFIER_HEAD
 ;

fragment VALUE_IDENTIFIER_CHARACTERS : VALUE_IDENTIFIER_CHARACTER+ ;

OPERATOR_HEAD_OTHER // valid operator chars not used by Swift itself
  : [\u00A1-\u00A7]
  | [\u00A9\u00AB]
  | [\u00AC\u00AE]
  | [\u00B0-\u00B1\u00B6\u00BB\u00BF\u00D7\u00F7]
  | [\u2016-\u2017\u2020-\u2027]
  | [\u2030-\u203E]
  | [\u2041-\u2053]
  | [\u2055-\u205E]
  | [\u2190-\u23FF]
  | [\u2500-\u2775]
  | [\u2794-\u2BFF]
  | [\u2E00-\u2E7F]
  | [\u3001-\u3003]
  | [\u3008-\u3030]
  ;

OPERATOR_FOLLOWING_CHARACTER
  : [\u0300-\u036F]
  | [\u1DC0-\u1DFF]
  | [\u20D0-\u20FF]
  | [\uFE00-\uFE0F]
  | [\uFE20-\uFE2F]
  //| [\uE0100-\uE01EF]  ANTLR can't do >16bit char
  ;

IMPLICIT_PARAMETER_NAME : DOLLER PureDecimalDigits ;

BinaryLiteral : '0b' Binary_digit BinaryLiteral_characters? ;
fragment Binary_digit : [01] ;
fragment BinaryLiteral_character : Binary_digit | '_'  ;
fragment BinaryLiteral_characters : BinaryLiteral_character+ ;

OctalLiteral : '0o' Octal_digit OctalLiteral_characters? ;
fragment Octal_digit : [0-7] ;
fragment OctalLiteral_character : Octal_digit | '_'  ;
fragment OctalLiteral_characters : OctalLiteral_character+ ;

DecimalLiteral		: [0-9] [0-9_]* ;
PureDecimalDigits   : [0-9]+ ;
fragment Decimal_digit : [0-9] ;
fragment Decimal_literal_character : Decimal_digit | '_'  ;
fragment Decimal_literal_characters : Decimal_literal_character+ ;

HexadecimalLiteral : '0x' Hexadecimal_digit Hexadecimal_literal_characters? ;
fragment Hexadecimal_digit : [0-9a-fA-F] ;
fragment Hexadecimal_literal_character : Hexadecimal_digit | '_'  ;
fragment Hexadecimal_literal_characters : Hexadecimal_literal_character+ ;

// GRAMMAR OF A FLOATING_POINT LITERAL

FloatLiteral
 : DecimalLiteral Decimal_fraction? Decimal_exponent?
 | HexadecimalLiteral Hexadecimal_fraction? Hexadecimal_exponent
 ;

fragment Decimal_fraction : '.' DecimalLiteral ;
fragment Decimal_exponent : Floating_point_e Sign? DecimalLiteral ;
fragment Hexadecimal_fraction : '.' Hexadecimal_digit Hexadecimal_literal_characters? ;
fragment Hexadecimal_exponent : Floating_point_p Sign? DecimalLiteral ;
fragment Floating_point_e : [eE] ;
fragment Floating_point_p : [pP] ;
fragment Sign : [+\-] ;


StaticStringLiteral
    : '"' Double_quoted_text? '"'
    | '\'' Single_quoted_text? '\''
    ;

fragment Double_quoted_text : Double_quoted_text_item+ ;
fragment Single_quoted_text : Single_quoted_text_item+ ;

fragment Double_quoted_text_item
  : Escaped_character
  | ~["\n\r\\]
  ;

fragment Single_quoted_text_item
  : Escaped_character
  | ~['\n\r\\]
  ;

fragment
Escaped_character
  : '\\' [0\\tnr"']
  | '\\x' Hexadecimal_digit Hexadecimal_digit
  | '\\u' Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit
  | '\\u' '{' Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit '}'
  | '\\u' '{' Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit Hexadecimal_digit '}'
  ;

InterpolatedStringLiteral
  : '"' Double_interpolated_text_item* '"'
  | '\'' Single_interpolated_text_item* '\''
  ;

fragment
Double_interpolated_text_item
  : '\\{' (InterpolatedStringLiteral | Double_interpolated_text_item)+ '}' // nested strings allowed
  | Double_quoted_text_item
  ;

fragment
Single_interpolated_text_item
  : '\\{' (InterpolatedStringLiteral | Single_interpolated_text_item)+ '}' // nested strings allowed
  | Single_quoted_text_item
  ;

WS
  : [\u0020\u0009\u000B\u000C\u0000]
  -> skip ;

DelimitedComment
  : '/*' (DelimitedComment|.)*? '*/'
  -> channel(HIDDEN) ; // nesting comments allowed

LineComment
  : '//' (~[/<\u000A\u000D] ~[\u000A\u000D]*)?
  -> channel(HIDDEN) ;

LineCommentForDocument
  : '////' ~[\u000A\u000D]*
  -> channel(HIDDEN) ;

EOL   : '\u000A' | '\u000D' '\u000A'; // \r\n

LineDocument : '///' (~[/<\u000A\u000D] ~[\u000A\u000D]*)?;
FollowingLineDocument : '//<' ~[\u000A\u000D]*;
