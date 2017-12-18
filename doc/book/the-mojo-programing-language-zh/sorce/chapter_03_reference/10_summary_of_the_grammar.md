## 语法总结

### 词法结构（Lexical Structure）

> ###### GRAMMAR OF A COMMENT
>
> ```abnf
> comment = line-comment / block-comment
>
> line-comment = "//"
> ```

> ###### GRAMMAR OF A DOCUMENT
>
> ```abnf
> document = following-document / up-document
>
> following-document = "///"
> up-document = "//<"
> ```

> ###### GRAMMAR OF AN IDENTIFIER
>
> ```abnf
> identifier = identifier-head [ identifier-characters ] / implicit-parameter-name
> identifier-list = identifier / identifier "," identifier-list
>
> identifier-head = %x61-7A / %xA1-%xFF
>                / %x0100-%x02FF / %x0370-%x167F / %x1681-%x180D / %x180F-%x1DBF / %x1E00-%x1FFF
>                / %x200B-%x200D / %x202A-%x20CF / %x2100-%x23FF / %x2460-%x2FFF / %x3001-%xD7FF
>                / %xF900-%xFD3D / %xFD40-%xFDCF / %xFDF0-%xFE1F / %xFE30-%xFE44 / %xFE47-%xFFFD
>                / %x10000-1FFFD / %x20000-2FFFD / %x30000-3FFFD
>                / %x40000-4FFFD / %x50000-5FFFD / %x60000-6FFFD
>                / %x70000-7FFFD / %x80000-8FFFD / %x90000-9FFFD
>                / %xA0000-AFFFD / %xB0000-BFFFD / %xC0000-CFFFD
>                / %xD0000-DFFFD / %xE1000-EFFFD
> identifier-character = identifier-head / "_" / DIGIT / %x41-5A
>                      / %x0300-%x036F / %x1DC0-%x1DFF / %x20D0-%x20FF / %xFE20-%xFE2F
> identifier-characters = identifier-character [ identifier-characters ]
>
> type-name = type-name-head [ type-name-characters ]
> type-name-head = %x41-5A
> type-name-character = ALPHA / DIGIT
> type-name-characters = type-name-character [ type-name-characters ]
>
> implicit-parameter-name = "$" [ decimal-digits ]
> ```

> ###### GRAMMAR OF AN LITERAL
>
> ```ABNF
> literal = numeric-literal / string-literal / boolean-literal / null-literal
>
> numeric-literal = [ "-" ] integer-literal / [ "-" ] floating-point-literal
> boolean-literal = "true" / "false"
> null-literal = "null"
> ```

> ###### GRAMMAR OF AN INTEGER LITERAL
>
> ```ABNF
> integer-literal = binary-literal / octal-literal / decimal-literal / hexadecimal-literal
>
> binary-literal = "0b" binary-digit [ binary-literal-characters ]
> binary-digit = "0" / "1"
> binary-literal-character = binary-digit / "_"
> binary-literal-characters = binary-literal-character [ binary-literal-characters ]
>
> octal-literal = "0o" octal-digit [ octal-literal-characters ]
> octal-digit = %x30-37
> octal-digit-character = octal-digit / "_"
> octal-digit-characters = octal-digit-character [ octal-digit-characters ]
>
> decimal-literal = decimal-digit [ decimal-literal-characters ]
> decimal-digit = DIGIT
> decimal-digits = decimal-digit [ decimal-digits ]
> decimal-literal-character = decimal-digit / "_"
> decimal-literal-characters = decimal-literal-character [ decimal-literal-characters ]
>
> hexadecimal-literal = "0x" HEXDIG [ exadecimal-literal-characters ]
> hexadecimal-literal-character = HEXDIG / "_"
> hexadecimal-literal-characters = hexadecimal-literal-character [ hexadecimal-literal-characters ]
> ```

> ###### GRAMMAR OF A FLOATING-POINT LITERAL
>
> ```ABNF
> floating-point-literal = decimal-literal [ decimal-fraction ] [ decimal-exponent ]
>
> decimal-fraction = "." decimal-literal
> decimal-exponent = floating-point-e [ sign ] decimal-literal
> floating-point-e = "e" / "E"
> sign = "+" / "-"
> ```

> ###### GRAMMAR OF A STRING LITERAL
>
> ```ABNF
> string-literal = static-string-literal / interpolated-string-literal
>
> static-string-literal = """ [ double-quoted-text ] """ / "'" [ single-quoted-text ] "'"
> double-quoted-text = double-quoted-text-item [ double-quoted-text ]
> double-quoted-text-item = escaped-character / %x0-%x09 / %xB-%x0C / %x0E-%x21 / %x23-%x5B /
>                           %x5D-%xD7FF / %xE000-%x10FFFF
> single-quoted-text = single-quoted-text-item [ single-quoted-text ]
> single-quoted-text-item = escaped-character / %x0-%x09 / %xB-%x0C / %x0E-%x26 / %x28-%x5B /
>                           %x5D-%xD7FF / %xE000-%x10FFFF
>
> interpolated-string-literal = """ [ interpolated-text ] """ / "'" [ interpolated-text ] "'"
> interpolated-text = interpolated-text-item [ interpolated-text ]
> interpolated-text-item = "\{" expression "}" / quoted-text-item
>
> escaped-character = "\0" / "\\" / "\t" / "\n" / "\r" / "\"" / "\'" / "\u" 4HEXDIG 
>                   / "\u{" unicode-scalar-digits "}"
> unicode-scalar-digits = 1*8(HEXDIG)
> ```

> ###### GRAMMAR OF OPERATORS
>
> ```
> operator = prefix-operator / infix-operator / postfix-operator
>
> prefix-operator = "-" / "!" / "not"
>
> binary-operator = arithmetic-operator /
>                   assignment-operator /
>                   relational-operator /
>                   bitwise_shift_operator /
>                   equality_operator /
>                   equiment-operator /
>                   logic-operator /
>                   bitwise-operator /
>                   pipeline-operator /
>                   range_operator /
>                   range_in_operator
>
> postfix-operator = ""
>
> arithmetic-operator = "*" / "/" / "%" / "^" / "+" / "-"
> assignment-operator = "=" / "+=" / "-=" / "*=" / "/=" / "^="
> relational-operator = "<" / "<=" / ">" / ">=" / "==" / "!="
>
> logic-operator = "&&" / "||" / "and" / "or"
> bitwise-logic-operator = "&&&" / "|||" / "xor"
> bitwise-shift-operator = ">>>" / "<<<"
> pipeline-operator = "|"
> ```

### 类型（Types）

> ###### GRAMMAR OF A TYPE
>
> ```
> type = array-type / map-type / tuple-type / function-type / type-identifier /
>        metatype-type
> ```

> ###### GRAMMAR OF A TYPE ANNOTATION
>
> ```
> type-annotation = ":"  type [ attributes ]
> ```

> ###### GRAMMAR OF A TYPE IDENTIFIER
>
> ```
> bare-type-identifier = type-name [ generic-argument-clause ] /
>                        type-name [ generic-argument-clause ] "." bare-type-identifier
>
> type-identifier = [ package-name "." ] bare-type-identifier
> ```

> ###### GRAMMAR OF A TUPLE TYPE
>
> ```
> tuple-type = "(" tuple-type-body ")"
>
> tuple-type-body = tuple-type-element-list
> tuple-type-element-list = tuple-type-element /
>                           tuple-type-element "," tuple-type-element-list
> tuple-type-element = type [ attributes ] / element-name type-annotation
> element-name = identifier
> ```

> ###### GRAMMAR OF A FUNCTION TYPE
>
> ```
> function-type = "(" function-type-parameters ")" "->" type / type "->" type
>
> function-type-parameters = type / type "," function-type-parameters
> ```

> ###### GRAMMAR OF AN ARRAY TYPE
>
> ```
> array-type = "[" type "]"
> ```

> ###### GRAMMAR OF A DICTIONARY TYPE
>
> ```
> dictionary-type = "[" type ":" type "]"
> ```

> ###### GRAMMAR OF A METATYPE TYPE
>
> ```
> metatype-type = type ".Type"
> ```

> ###### GRAMMAR OF A TYPE INHERITANCE CLAUSE
>
> ```
> type-inheritance-clause = ":" type-inheritance-list
> type-inheritance-list = type-identifier / type-identifier "," type-inheritance-list
> ```

### 表达式（Expressions）

> ###### GRAMMAR OF A CONDITIONAL OPERATOR
>
> ```
> conditional-operator = "?" expression ":"
> ```

> ###### GRAMMAR OF A PRIMARY EXPRESSION
>
> ```
> primary-expression = identifier [ generic-argument-clause ] /
>                      literal-expression /
>                      closure-expression /
>                      parenthesized-expression /
>                      wildcard-expression
> ```

> ###### GRAMMAR OF A LITERAL EXPRESSION
>
> ```
> literal-expression = literal / array-literal / object-literal
>
> array-literal = "[" [ array-literal-items ] "]"
> array-literal-items = array-literal-item [ "," ] / array-list-item "," array-list-items
> array-literal-item = expression
>
> object-literal = "{" [ object-literal-items ] "}"
> object-literal-items = object-literal-item [ "," ] / object-literal-item "," object-literal-items
> object-literal-item = (identifier / literal) ":" expression
> ```

> ###### GRAMMAR OF A CLOSURE EXPRESSION
>
> ```
> closure-expression = closure-signature "->" statement /
>                      closure-signature "->" type-identifier "{" statements "}"
>
> closure-signature = closure-parameter-clause "->" function-list
> closure-parameter-clause = "(" ")" /
>                            "(" closure-parameter-list ")" /
>                            identifier-list
> closure-parameter-list = closure-parameter /
>                          closure-parameter "," closure-parameter-list
> closure-parameter = closure-parameter-name [ type-annotation ] /
>                     closure-parameter-name "..."
> closure-parameter-name = identifier
> ```

> ###### GRAMMAR OF A PARENTHESIZED EXPRESSION
>
> ```
> parenthesized-expression = "(" [ expression-element-list ] ")"
> expression-element-list = expression-element / expression-element "," expression-element-list
> expression-element = expression / identifier ":" expression
> ```

> GRAMMAR OF A WILDCARD EXPRESSION
>
> ```
> wildcard-expression = "_"
> ```
>
> 

> GRAMMAR OF A POSTFIX EXPRESSION
>
> ```
> postfix-expression = primary-expression /
>                      postfix-expression postfix-operator /
>                      function-call-expression /
>                      explicit-member-expression /
>                      subscript-expression /
>                      forced-value-expression /
>                      optional-chaining-expression
> ```
>
> 

### 语句（Statements）

> GRAMMAR OF A STATEMENT
>
> ```
> statement = expression [ ";" ] /
>             declaration [ ";" ] /
>             loop-statement [ ";" ] /
>             branch-statement [ ";" ] /
>             control-transfer-statement [ ";" ]
>
> statements = statement [ statements ]
> ```

### 声明（Declarations）



### 特性（Attributes）

> GRAMMAR OF AN ATTRIBUTE
>
> ```abnf
> attribute = "@" decimal-digits / ( path-identifier [ attribute-argument ] )
>
> attribute-argument = "(" expression ")"
>
> group-attribute = "@{" [ group-attribute-element-list ] "}"
>
> group-attribute-element-list = group-attribute-element [ group-attribute-element-list ]
> group-attribute-element = path-identifier ":" expression
>
> attributes = ( group-attribute / attribute ) [ attributes ]
> ```
>

### 模式（Patterns）

### 泛型参数（Generic Parameters and Arguments）

