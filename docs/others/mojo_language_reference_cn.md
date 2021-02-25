# Mojo Language Reference

[TOC]

## About the Language Reference

## Lexicals

The *lexicals* of Mojo describes what sequence of characters form valid tokens of the language. These valid tokens form the lowest-level building blocks of the language and are used to describe the rest of the language in subsequent chapters. A token consists of an identifier, keyword, punctuation, literal, or operator.

In most cases, tokens are generated from the characters of a Mojo source file by considering the longest possible substring from the input text, within the constraints of the grammar that are specified below. This behavior is referred to as *longest match* or *maximal munch*.

### Whitespace

Whitespace has two uses: to separate tokens in the source file and to help determine whether an operator is a prefix or postfix (see Operators), but is otherwise ignored. The following characters are considered whitespace: space (U+0020), line feed (U+000A), carriage return (U+000D), horizontal tab (U+0009), vertical tab (U+000B), form feed (U+000C) and null (U+0000).

### Comments

Comments are treated as whitespace by the compiler.

#### Single line comment

Single line comments begin with `//` and continue until a line feed (U+000A) or carriage return (U+000D).

#### Multiline comment

Multiline comments begin with `/*` and end with `*/`. Nesting multiline comments is allowed, but the comment markers must be balanced.

### Documents

#### Single line document

Comments can contain additional formatting and markup, as described in comments can contain additional formatting and markup, as described in Markup Formatting Reference.

### Identifies

Identifiers begin with an uppercase or lowercase letter A through Z, an underscore (_), a noncombining alphanumeric Unicode character in the Basic Multilingual Plane, or a character outside the Basic Multilingual Plane that isn’t in a Private Use Area. After the first character, digits and combining Unicode characters are also allowed.

To use a reserved word as an identifier, put a backtick (`) before and after it. For example, class is not a valid identifier, but `class` is valid. The backticks are not considered part of the identifier; `x` and x have the same meaning.

Inside a closure with no explicit parameter names, the parameters are implicitly named $0, $1, $2, and so on. These names are valid identifiers within the scope of the closure.

### Keywords and Punctuation

The following keywords are reserved and can’t be used as identifiers, unless they’re escaped with backticks, as described above in Identifiers. 

- Keywords used in declarations: `enum`, `func`, `import`,`package`, `const`,`var`,`attribute`, `type`.
- Keywords used in statements: `break`, `continue`, `else`, `for`, `if`,` in`, `repeat`, `return`, `switch`, and `while`.
- Keywords used in expressions and types: `as`, `false`, `is`, `null`and `true`.
- Keywords used in Operators: `and`, `or`, `xor`,`not`and `in`.
- Keywords used in patterns: `_`.

| Declarations | Statements | Expressions | Operators | Patterns | Types |
| ------------ | ---------- | ----------- | --------- | -------- | ----- |
| `type`       | `if`       | `null`      | `and`     | `_`      | `as`  |
| `enum`       | `else`     | `true`      | `or`      |          | `is`  |
| `func`       | `for`      | `false`     | `xor`     |          |       |
| `import`     | `in`       |             | `not`     |          |       |
| `const`      | `while`    |             | `in`      |          |       |
| `var`        | `repeat`   |             |           |          |       |
| `package`    | `switch`   |             |           |          |       |
|              | `continue` |             |           |          |       |
|              | `break`    |             |           |          |       |

The following tokens are reserved as punctuation and can’t be used as custom operators:

|      |        |       |       |
| ---- | ------ | ----- | ----- |
| `(`  | `+=`   | `,`   | `'`   |
| `)`  | `-=`   | `:`   | `"`   |
| `[`  | `*=`   | `;`   | `'''` |
| `]`  | `/=`   | `.`   | `\`   |
| `{`  | `%=`   | `@`   | `$`   |
| `}`  | `^=`   | `?`   |       |
|      | `~=`   | `->`  |       |
|      | `>>=`  | `..`  |       |
|      | `<<=`  | `...` |       |
|      | `>>>=` |       |       |
|      | `=`    |       |       |

### Literals

#### Integer Literals

#### Floating-Point Literals

#### String Literals

### Operators

| Arithmetic Operators | Relational Operators | Logical Operators | Bitwise Operators |
| :------------------: | :------------------: | :---------------: | :---------------: |
|         `+`          |         `<`          |       `&&`        |        `~`        |
|         `-`          |         `>`          |      `\|\|`       |       `<<`        |
|         `*`          |         `<=`         |        `!`        |       `>>`        |
|         `/`          |         `>=`         |                   |       `>>>`       |
|         `%`          |         `==`         |                   |      `xand`       |
|         `^`          |         `!=`         |                   |       `xor`       |
|         `++`         |                      |                   |                   |
|         `--`         |                      |                   |                   |

#### Arithmetic operators

##### Unary operators

| Operator | Name | Purpose     | Remarks |
| -------- | ---- | ----------- | ------- |
| `+`      | \    | addition    |         |
| `-`      | \    | subtraction |         |
| ++       |      |             |         |
| --       |      |             |         |

##### Binary Operators

| Operator | Name    | Purpose        | Remarks                                  |
| -------- | ------- | -------------- | ---------------------------------------- |
| `+`      | `add`   | addition       |                                          |
| `-`      | `sub`   | subtraction    |                                          |
| `*`      | `multi` | multiplication |                                          |
| `/`      | `div`   | division       | Use `intdiv()` for integer division, and see the section about [integer division](http://docs.groovy-lang.org/latest/html/documentation/#integer_division)for more information on the return type of the division. |
| `%`      | `mod`   | remainder      |                                          |
| `^`      | `pow`   | power          | See the section about [the power operation]() for more information on the return type of the operation. |

#### Assignment arithmetic operators

The binary arithmetic operators we have seen above are also available in an assignment form:

- `+=`
- `-=`
- `*=`
- `/=`
- `%=`
- `^=`

#### Relational operators

Relational operators allow comparisons between objects, to know if two objects are the same or different, or if one is greater than, less than, or equal to the other.

The following operators are available:

| Operator | Purpose            |                       |
| -------- | ------------------ | --------------------- |
| `==`     | equal              |                       |
| `!=`     | different          |                       |
| `<`      | less than          |                       |
| `<=`     | less than or equal |                       |
| `>`      |                    | greater than          |
| `>=`     |                    | greater than or equal |

#### Logical operators

Mojo offers three logical operators for boolean expressions:

- `&&`
- `||`
- `not` /  `!` `~`

##### Short-circuiting



#### Bitwise operators

- `&&&`
- `|||`
- `~~~`
- `xor`
- `<<`
- `>>`
- `>>>`

## Types

### Type Annotation

“A type annotation explicitly specifies the type of a variable or expression. Type annotations begin with a colon (:) and end with a type, as the following examples show:”

```mojo
var some_tuple: (Double, Double) = (3.14159, 2.71828)
func some_function(a: Int) {}
type Some_type {
  	a: Int32 @1
}
```

> ##### GRAMMAR OF A TYPE ANNOTATION
>
> ```ABNF
> type-annotation = ":" type [ attributes ]
> ```



### pipeline

`|`
`<|` Passes the result of the expression on the right side to the function on left side (backward pipe operator).

