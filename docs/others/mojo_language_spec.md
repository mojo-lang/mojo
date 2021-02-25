# Mojo Language Specification

[TOC]

## Introduction

Mojo...

- is an agile and dynamic language for the structure data process
- ​



## Syntax

This chapter covers the syntax of the Mojo programming language. The grammar of the language derives from the Java grammar, and allows certain simplifications.

### Comments

#### Single line comment

Single line comments start with `//` and can be found at any position in the line. The characters following `//`, till the end of the line, are considered part of the comment.

```mojo
// a standalone single line comment
print("hello"); // a comment till the end of the line
```

#### Multiline comment

A multiline comment starts with `/*` and can be found at any position in the line. The characters following `/*` will be considered part of the comment, including new line characters, up to the first `*/` closing the comment. Multiline comments can thus be put at the end of a statement, or even inside a statement.

```mojo
/* a standalone multiline comment
   spanning two lines */
print("hello"); /* a multiline comment starting
                   at the end of a statement */
print(1 /* one */ + 2 /* two */);
```

#### Single line doc comment

#### Multiline doc comment

```
///
/// A Struct description
///
struct Person {
    /// the name of the person
    name: String,
    
    ///
    ///
    ///
    role: String 
}
```

### Keywords

The following list represents all the keywords of the Mojo language:

Table 1. Keywords

| struct | enum  | fun   |         |
| ------ | ----- | ----- | ------- |
| null   | true  | false |         |
| for    | while | match | default |
| if     | else  | break | return  |
| and    | or    | not   | in      |

### Identifiers

#### Normal identifiers

Identifiers start with a letter,  or an underscore. They cannot start with a number.

A letter can be in the following ranges:

- 'a' to 'z' (lowercase ascii letter)
- 'A' to 'Z' (uppercase ascii letter)
- '\u00C0' to '\u00D6'
- '\u00D8' to '\u00F6'
- '\u00F8' to '\u00FF'
- '\u0100' to '\u_f_f_f_e'

Then following characters can contain letters and numbers.

Here are a few examples of valid identifiers (here, variable names):

```
def name
def item3
def with_underscore
def $dollar_start
```

But the following ones are invalid identifiers:

```
def 3tier
def a+b
def a#b
```

All keywords are also valid identifiers when following a dot:

```
foo.as
foo.assert
foo.break
foo.case
foo.catch
```

#### Unicode identifiers



### Strings

Text literals are represented in the form of chain of characters called strings. Groovy lets you instantiate `java.lang.String`objects, as well as GStrings (`groovy.lang.GString`) which are also called *interpolated strings* in other programming languages.

#### Single quoted string

Single quoted strings are a series of characters surrounded by single quotes:

```
'a single quoted string'
```

| **   | Single quoted strings are plain `java.lang.String` and don’t support interpolation. |
| ---- | ---------------------------------------- |
|      |                                          |

#### Double quoted string

Double quoted strings are a series of characters surrounded by double quotes:

```
"a double quoted string"
```

| **   | Double quoted strings are plain `java.lang.String` if there’s no interpolated expression, but are`groovy.lang.GString` instances if interpolation is present. |
| ---- | ---------------------------------------- |
|      |                                          |

| **   | To escape a double quote, you can use the backslash character: "A double quote: \"". |
| ---- | ---------------------------------------- |
|      |                                          |

#### Triple quoted string

Triple single quoted strings are a series of characters surrounded by triplets of single quotes:

```
'''a triple single quoted string'''
```

| **   | Triple single quoted strings are plain `java.lang.String` and don’t support interpolation. |
| ---- | ---------------------------------------- |
|      |                                          |

Triple single quoted strings are multiline. You can span the content of the string across line boundaries without the need to split the string in several pieces, without contatenation or newline escape characters:

```
def a_multiline_string = '''line one
line two
line three'''
```

If your code is indented, for example in the body of the method of a class, your string will contain the whitespace of the indentation. The Groovy Development Kit contains methods for stripping out the indentation with the `String#strip_indent()`method, and with the `String#strip_margin()` method that takes a delimiter character to identify the text to remove from the beginning of a string.

When creating a string as follows:

```
def starting_and_ending_with_a_newline = '''
line one
line two
line three
'''
```

You will notice that the resulting string contains a newline character as first character. It is possible to strip that character by escaping the newline with a backslash:

```
def stripped_first_newline = '''\
line one
line two
line three
'''

assert !stripped_first_newline.starts_with('\n')
```

#### Template string

#### Regex string (Slashy string)

Beyond the usual quoted strings, Groovy offers slashy strings, which use `/` as delimiters. Slashy strings are particularly useful for defining regular expressions and patterns, as there is no need to escape backslashes.

Example of a slashy string:

```
def foo_pattern = /.*foo.*/
assert foo_pattern == '.*foo.*'
```

Only forward slashes need to be escaped with a backslash:

```
def escape_slash = /The character \/ is a forward slash/
assert escape_slash == 'The character / is a forward slash'
```

Slashy strings are multiline:

```
def multiline_slashy = /one
    two
    three/

assert multiline_slashy.contains('\n')
```

Slashy strings can also be interpolated (ie. a GString):

```
def color = 'blue'
def interpolated_slashy = /a ${color} car/

assert interpolated_slashy == 'a blue car'
```

There are a few gotchas to be aware of.

An empty slashy string cannot be represented with a double forward slash, as it’s understood by the Groovy parser as a line comment. That’s why the following assert would actually not compile as it would look like a non-terminated statement:

```
assert '' == //
```

| **   | As slashy strings were mostly designed to make regexp easier so a few things that are errors in GStrings like `$()`will work with slashy strings. |
| ---- | ---------------------------------------- |
|      |                                          |

#### String concatenation

All the Groovy strings can be concatenated with the `+` operator:

```
assert 'ab' == 'a' + 'b'
```

#### Escaping special characters

You can escape single quotes with the the backslash character to avoid terminating the string literal:

```
'an escaped single quote: \' needs a backslash'
```

And you can escape the escape character itself with a double backslash:

```
'an escaped escape character: \\ needs a double backslash'
```

Some special characters also use the backslash as escape character:

| Escape sequence | Character                                |
| --------------- | ---------------------------------------- |
| '\t'            | tabulation                               |
| '\b'            | backspace                                |
| '\n'            | newline                                  |
| '\r'            | carriage return                          |
| '\f'            | formfeed                                 |
| '\\'            | backslash                                |
| '\''            | single quote (for single quoted and triple single quoted strings) |
| '\"'            | double quote (for double quoted and triple double quoted strings) |

#### Unicode escape sequence

For characters that are not present on your keyboard, you can use unicode escape sequences: a backslash, followed by 'u', then 4 hexadecimal digits.

For example, the Euro currency symbol can be represented with:

```
'The Euro currency symbol: \u20AC'
```

#### String summary table

| String name          | String syntax | Interpolated | Multiline | Escape character |
| -------------------- | ------------- | ------------ | --------- | ---------------- |
| Single quoted        | `'…​'`        | **           | **        | `\`              |
| Triple single quoted | `'''…​'''`    | **           | **        | `\`              |
| Double quoted        | `"…​"`        | **           | **        | `\`              |
| Triple double quoted | `"""…​"""`    | **           | **        | `\`              |
| Slashy               | `/…​/`        | **           | **        | `\`              |
| Dollar slashy        | `$/…​/$`      | **           | **        | `$`              |

#### Characters

Unlike Java, Groovy doesn’t have an explicit character literal. However, you can be explicit about making a Groovy string an actual character, by three different means:

```
char c1 = 'A' 
assert c1 instanceof Character

def c2 = 'B' as char 
assert c2 instanceof Character

def c3 = (char)'C' 
assert c3 instanceof Character
```

| **   | by being explicit when declaring a variable holding the character by specifying the `char` type |
| ---- | ---------------------------------------- |
| **   | by using type coercion with the `as` operator |
| **   | by using a cast to char operation        |

| **   | The first option *1* is interesting when the character is held in a variable, while the other two (*2* and *3*) are more interesting when a char value must be passed as argument of a method call. |
| ---- | ---------------------------------------- |
|      |                                          |

### Numbers

Mojo supports different kinds of integral literals and decimal literals.

#### Integral literals

The integral literal types are flowers:

- `Int8` / `UInt8`
- `Int16` / `UInt16`
- `Int32` / `UInt32`
- `Int64` / `UInt64`
- `Int` / `UInt`

You can create integral numbers of those types with the following declarations:

```
// primitive types
b : Int8 = 1;
char  c = 2
short s = 3
int   i = 4
long  l = 5
```

If you use optional typing by using the `def` keyword, the type of the integral number will vary: it’ll adapt to the capacity of the type that can hold that number.

For positive numbers:

```
a = 1;
assert(type_of(a) == Int);

// Int.max_value
b = 2147483647
assert b instanceof Integer

// Integer.MAX_VALUE + 1
c = 2147483648
assert c instanceof Long

// Long.MAX_VALUE
d = 9223372036854775807
assert d instanceof Long

```

As well as for negative numbers:

```
def na = -1
assert na instanceof Integer

// Integer.MIN_VALUE
def nb = -2147483648
assert nb instanceof Integer

// Integer.MIN_VALUE - 1
def nc = -2147483649
assert nc instanceof Long

// Long.MIN_VALUE
def nd = -9223372036854775808
assert nd instanceof Long

// Long.MIN_VALUE - 1
def ne = -9223372036854775809
assert ne instanceof BigInteger
```

#### Alternative non-base 10 representations

Numbers can also be represented in binary, octal, hexadecimal and decimal bases.

##### Binary literal

Binary numbers start with a 0b prefix:

```
int x_int = 0b10101111
assert x_int == 175

short x_short = 0b11001001
assert x_short == 201 as short

byte x_byte = 0b11
assert x_byte == 3 as byte

long x_long = 0b101101101101
assert x_long == 2925l
```

##### Octal literal

Octal numbers are specified in the typical format of `0` followed by octal digits.

```
int x_int = 077
assert x_int == 63

short x_short = 011
assert x_short == 9 as short

byte x_byte = 032
assert x_byte == 26 as byte

long x_long = 0246
assert x_long == 166l
```

##### Hexadecimal literal

Hexadecimal numbers are specified in the typical format of `0x`or `#` followed by hex digits.

```
int x_int = 0x77
assert x_int == 119

short x_short = 0xaa
assert x_short == 170 as short

byte x_byte = 0x3a
assert x_byte == 58 as byte

long x_long = 0xffff
assert x_long == 65535l

BigInteger x_big_integer = 0xaaaa
assert x_big_integer == 43690g
```

#### Real literals

#### Underscore in literals

When writing long literal numbers, it’s harder on the eye to figure out how some numbers are grouped together, for example with groups of thousands, of words, etc. By allowing you to place underscore in number literals, it’s easier to spot those groups:

```
long credit_card_number = 1234_5678_9012_3456L
long social_security_numbers = 999_99_9999L
double monetary_amount = 12_345_132.12
long hex_bytes = 0xFF_EC_DE_5E
long hex_words = 0xFFEC_DE5E
long max_long = 0x7fff_ffff_ffff_ffffL
long also_max_long = 9_223_372_036_854_775_807L
long bytes = 0b11010010_01101001_10010100_10010010
```

#### Number type suffixes

We can force a number (including binary, octals and hexadecimals) to have a specific type by giving a suffix (see table bellow), either uppercase or lowercase.

| Type    | Suffix     |
| ------- | ---------- |
| Long    | `L` or `l` |
| Integer | `I` or `i` |
| Float   | `F` or `f` |

Examples:

```
assert 42i == Int('42')
assert 42i == new Integer('42') // lowercase i more readable
assert 123L == new Long("123") // uppercase L more readable
assert 2147483648 == new Long('2147483648') // Long type used, value too large for an Integer
assert 123.45 == new BigDecimal('123.45') // default BigDecimal type used
assert 1.200065D == new Double('1.200065')
assert 1.234F == new Float('1.234')
assert 1.23E23D == new Double('1.23E23')
assert 0b1111L.class == Long // binary
assert 0xFFi.class == Integer // hexadecimal
assert 034G.class == BigInteger // octal
```

#### Math operations

Although [operators](http://docs.groovy-lang.org/latest/html/documentation/#_operators) are covered later on, it’s important to discuss the behavior of math operations and what their resulting types are.

##### The case of the division operator

##### The case of the power operator

### Array/Sequence

### Objects (Maps)

## Operators

This chapter covers the operators of the Mojo programming language.

### Arithmetic operators

Groovy supports the usual familiar arithmetic operators you find in mathematics and in other programming languages like Java. All the Java arithmetic operators are supported. Let’s go through them in the following examples.

#### Normal arithmetic operators

The following binary arithmetic operators are available in Groovy:

| Operator | Purpose        | Remarks                                  |
| -------- | -------------- | ---------------------------------------- |
| `+`      | addition       |                                          |
| `-`      | subtraction    |                                          |
| `*`      | multiplication |                                          |
| `/`      | division       | Use `intdiv()` for integer division, and see the section about [integer division](http://docs.groovy-lang.org/latest/html/documentation/#integer_division)for more information on the return type of the division. |
| `%`      | remainder      |                                          |
| `^`      | power          | See the section about [the power operation](http://docs.groovy-lang.org/latest/html/documentation/#power_operator) for more information on the return type of the operation. |

Here are a few examples of usage of those operators:

```
assert  1  + 2 == 3
assert  4  - 3 == 1
assert  3  * 5 == 15
assert  3  / 2 == 1.5
assert 10  % 3 == 1
assert  2 ^ 3 == 8
```

### Unary operators

The `+` and `-` operators are also available as unary operators:

```
assert +3 == 3
assert -4 == 0 - 4

assert -(-1) == 1  
```

| **   | Note the usage of parentheses to surround an expression to apply the unary minus to that surrounded expression. |
| ---- | ---------------------------------------- |
|      |                                          |

In terms of unary arithmetics operators, the `++` (increment) and `--` (decrement) operators are available, only in prefix notation:

```
def a = 2
def b = ++a * 3             

assert a == 3 && b == 6

def c = 3
def d = c-- * 2             

assert c == 2 && d == 6

def e = 1
def f = ++e + 3             

assert e == 2 && f == 5

def g = 4
def h = --g + 1             

assert g == 3 && h == 4
```

| **   | The postfix increment will increment `a` after the expression has been evaluated and assigned into `b` |
| ---- | ---------------------------------------- |
| **   | The postfix decrement will decrement `c` after the expression has been evaluated and assigned into `d` |
| **   | The prefix increment will increment `e` before the expression is evaluated and assigned into `f` |
| **   | The prefix decrement will decrement `g` before the expression is evaluated and assigned into `h` |

### Assignment arithmetic operators

The binary arithmetic operators we have seen above are also available in an assignment form:

- `+=`
- `-=`
- `*=`
- `/=`
- `%=`
- `^=`

Let’s see them in action:

```
def a = 4
a += 3

assert a == 7

def b = 5
b -= 3

assert b == 2

def c = 5
c *= 3

assert c == 15

def d = 10
d /= 2

assert d == 5

def e = 10
e %= 3

assert e == 1

def f = 3
f **= 2

assert f == 9
```



### Relational operators

Relational operators allow comparisons between objects, to know if two objects are the same or different, or if one is greater than, less than, or equal to the other.

The following operators are available:

| Operator | Purpose               |
| -------- | --------------------- |
| `==`     | equal                 |
| `!=`     | different             |
| `<`      | less than             |
| `<=`     | less than or equal    |
| `>`      | greater than          |
| `>=`     | greater than or equal |

Here are some examples of simple number comparisons using these operators:

```
assert 1 + 2 == 3
assert 3 != 4

assert -2 < 3
assert 2 <= 2
assert 3 <= 4

assert 5 > 1
assert 5 >= -2
```



### Logical operators

Groovy offers three logical operators for boolean expressions:

- `and`
- `or`
- `not` /  `!`

Let’s illustrate them with the following examples:

```
assert !false           
assert true and true     
assert true or false    
```

| **   | "not" false is true     |
| ---- | ----------------------- |
| **   | true "and" true is true |
| **   | true "or" false is true |

#### Precedence

The logical "not" has a higher priority than the logical "and".

```
assert (!false && false) == false   
```

| **   | Here, the assertion is true (as the expression in parentheses is false), because "not" has a higher precedence than "and", so it only applies to the first "false" term; otherwise, it would have applied to the result of the "and", turned it into true, and the assertion would have failed |
| ---- | ---------------------------------------- |
|      |                                          |

The logical "and" has a higher priority than the logical "or".

```
assert true || true && false        
```

| **   | Here, the assertion is true, because "and" has a higher precedence than "or", therefore the "or" is executed last and returns true, having one true argument; otherwise, the "and" would have executed last and returned false, having one false argument, and the assertion would have failed |
| ---- | ---------------------------------------- |
|      |                                          |



#### Short-circuiting

The logical `||` operator supports short-circuiting: if the left operand is true, it knows that the result will be true in any case, so it won’t evaluate the right operand. The right operand will be evaluated only if the left operand is false.

Likewise for the logical `&&` operator: if the left operand is false, it knows that the result will be false in any case, so it won’t evaluate the right operand. The right operand will be evaluated only if the left operand is true.

```
boolean check_if_called() {   
    called = true
}

called = false
true || check_if_called()
assert !called              

called = false
false || check_if_called()
assert called               

called = false
false && check_if_called()
assert !called              

called = false
true && check_if_called()
assert called               
```

| **   | We create a function that sets the `called` flag to true whenever it’s called |
| ---- | ---------------------------------------- |
| **   | In the first case, after resetting the called flag, we confirm that if the left operand to `\|\|` is true, the function is not called, as`\|\|` short-circuits the evaluation of the right operand |
| **   | In the second case, the left operand is false and so the function is called, as indicated by the fact our flag is now true |
| **   | Likewise for `&&`, we confirm that the function is not called with a false left operand |
| **   | But the function is called with a true left operand |

### Bitwise operators

### Conditional operators

#### Not operator

The "not" operator is represented with an exclamation mark (`!`) and inverts the result of the underlying boolean expression. In particular, it is possible to combine the `not` operator with the [Groovy truth](http://docs.groovy-lang.org/latest/html/documentation/#Groovy-Truth):

```
assert (!true)    == false                      
assert (!'foo')   == false                      
assert (!'')      == true                       
```

| **   | the negation of `true` is `false`        |
| ---- | ---------------------------------------- |
| **   | 'foo' is a non empty string, evaluating to `true`, so negation returns `false` |
| **   | '' is an empty string, evaluating to `false`, so negation returns `true` |

#### Ternary operator

The ternary operator is a shortcut expression that is equivalent to an if/else branch assigning some value to a variable.

Instead of:

```
if (string!=null && string.length()>0) {
    result = 'Found'
} else {
    result = 'Not found'
}
```

You can write:

```
result = (string!=null && string.length()>0) ? 'Found' : 'Not found'
```

The ternary operator is also compatible with the [Groovy truth](http://docs.groovy-lang.org/latest/html/documentation/#Groovy-Truth), so you can make it even simpler:

```
result = string ? 'Found' : 'Not found'
```

#### Elvis operator

The "Elvis operator" is a shortening of the ternary operator. One instance of where this is handy is for returning a 'sensible default' value if an expression resolves to `false`-ish (as in [Groovy truth](http://docs.groovy-lang.org/latest/html/documentation/#Groovy-Truth)). A simple example might look like this:

```
display_name = user.name ? user.name : 'Anonymous'   
display_name = user.name ?: 'Anonymous'              
```

| **   | with the ternary operator, you have to repeat the value you want to assign |
| ---- | ---------------------------------------- |
| **   | with the Elvis operator, the value, which is tested, is used if it is not `false`-ish |

Usage of the Elvis operator reduces the verbosity of your code and reduces the risks of errors in case of refactorings, by removing the need to duplicate the expression which is tested in both the condition and the positive return value.

### Object operators

##### Direct field access operator

Normally in Groovy, when you write code like this:

```
class User {
    public final String name                 
    User(String name) { this.name = name}
    String get_name() { "Name: $name" }       
}
def user = new User('Bob')
assert user.name == 'Name: Bob'              
```

| **   | public field `name`                      |
| ---- | ---------------------------------------- |
| **   | a getter for `name` that returns a custom string |
| **   | calls the getter                         |

The `user.name` call triggers a call to the property of the same name, that is to say, here, to the getter for `name`. If you want to retrieve the field instead of calling the getter, you can use the direct field access operator:

```
assert user.@name == 'Bob'                   
```

| **   | use of `.@` forces usage of the field instead of the getter |
| ---- | ---------------------------------------- |
|      |                                          |

#### Method pointer operator

The method pointer operator (`.&`) call be used to store a reference to a method in a variable, in order to call it later:

```
def str = 'example of method reference'            
def fun = str.&to_upper_case                         
def upper = fun()                                  
assert upper == str.to_upper_case()                  
```

| **   | the `str` variable contains a `String`   |
| ---- | ---------------------------------------- |
| **   | we store a reference to the `to_upper_case` method on the `str` instance inside a variable named `fun` |
| **   | `fun` can be called like a regular method |
| **   | we can check that the result is the same as if we had called it directly on `str` |

There are multiple advantages in using method pointers. First of all, the type of such a method pointer is a`groovy.lang.Closure`, so it can be used in any place a closure would be used. In particular, it is suitable to convert an existing method for the needs of the strategy pattern:

```
def transform(List elements, Closure action) {                    
    def result = []
    elements.each {
        result << action(it)
    }
    result
}
String describe(Person p) {                                       
    "$p.name is $p.age"
}
def action = this.&describe                                       
def list = [
    new Person(name: 'Bob',   age: 42),
    new Person(name: 'Julia', age: 35)]                           
assert transform(list, action) == ['Bob is 42', 'Julia is 35']    
```

| **   | the `transform` method takes each element of the list and calls the `action` closure on them, returning a new list |
| ---- | ---------------------------------------- |
| **   | we define a function that takes a `Person` and returns a `String` |
| **   | we create a method pointer on that function |
| **   | we create the list of elements we want to collect the descriptors |
| **   | the method pointer can be used where a `Closure` was expected |

Method pointers are bound by the receiver and a method name. Arguments are resolved at runtime, meaning that if you have multiple methods with the same name, the syntax is not different, only resolution of the appropriate method to be called will be done at runtime:

```
def do_something(String str) { str.to_upper_case() }    
def do_something(Integer x) { 2*x }                   
def reference = this.&do_something                    
assert reference('foo') == 'FOO'                     
assert reference(123)   == 246                       
```

| **   | define an overloaded `do_something` method accepting a `String` as an argument |
| ---- | ---------------------------------------- |
| **   | define an overloaded `do_something` method accepting an `Integer` as an argument |
| **   | create a single method pointer on `do_something`, without specifying argument types |
| **   | using the method pointer with a `String` calls the `String` version of `do_something` |
| **   | using the method pointer with an `Integer` calls the `Integer` version of `do_something` |

### Regular expression operators

#### Find operator

Alternatively to building a pattern, you can directly use the find operator `=~` to build a `java.util.regex.Matcher` instance:

```
def text = "some text to match"
def m = text =~ /match/                                           
assert m instanceof Matcher                                       
if (!m) {                                                         
    throw new RuntimeException("Oops, text not found!")
}
```

| **   | `=~` creates a matcher against the `text` variable, using the pattern on the right hand side |
| ---- | ---------------------------------------- |
| **   | the return type of `=~` is a `Matcher`   |
| **   | equivalent to calling `if (!m.find())`   |

Since a `Matcher` coerces to a `boolean` by calling its `find` method, the `=~` operator is consistent with the simple use of Perl’s`=~` operator, when it appears as a predicate (in `if`, `while`, etc.).

#### Match operator

The match operator (`==~`) is a slight variation of the find operator, that does not return a `Matcher` but a boolean and requires a strict match of the input string:

```
m = text ==~ /match/                                              
assert m instanceof Boolean                                       
if (m) {                                                          
    throw new RuntimeException("Should not reach that point!")
}
```

| **   | `==~` matches the subject with the regular expression, but match must be strict |
| ---- | ---------------------------------------- |
| **   | the return type of `==~` is therefore a `boolean` |
| **   | equivalent to calling `if (text ==~ /match/)` |



### Pipeline operator



### Other operators

#### Spread operator

#### Range operator

Groovy supports the concept of ranges and provides a notation (`..`) to create ranges of objects:

```
def range = 0..5                                    
assert (0..5).collect() == [0, 1, 2, 3, 4, 5]       
assert (0..<5).collect() == [0, 1, 2, 3, 4]         
assert (0..5) instanceof List                       
assert (0..5).size() == 6                           
```

| **   | a simple range of integers, stored into a local variable |
| ---- | ---------------------------------------- |
| **   | an `IntRange`, with inclusive bounds     |
| **   | an `IntRange`, with exclusive upper bound |
| **   | a `groovy.lang.Range` implements the `List` interface |
| **   | meaning that you can call the `size` method on it |

Ranges implementation is lightweight, meaning that only the lower and upper bounds are stored. You can create a range from any `Comparable` object that has `next()` and `previous()` methods to determine the next / previous item in the range. For example, you can create a range of characters this way:

```
assert ('a'..'d').collect() == ['a','b','c','d']
```

#### Subscript operator

The subscript operator is a short hand notation for `get_at` or `put_at`, depending on whether you find it on the left hand side or the right hand side of an assignment:

```
def list = [0,1,2,3,4]
assert list[2] == 2                         
list[2] = 4                                 
assert list[0..2] == [0,1,4]                
list[0..2] = [6,6,6]                        
assert list == [6,6,6,3,4]                  
```

| **   | `[2]` can be used instead of `get_at(2)`  |
| ---- | ---------------------------------------- |
| **   | if on left hand side of an assignment, will call `put_at` |
| **   | `get_at` also supports ranges             |
| **   | so does `put_at`                          |
| **   | the list is mutated                      |

The subscript operator, in combination with a custom implementation of `get_at`/`put_at` is a convenient way for destructuring objects:

```
class User {
    Long id
    String name
    def get_at(int i) {                                             
        switch (i) {
            case 0: return id
            case 1: return name
        }
        throw new IllegalArgumentException("No such element $i")
    }
    void put_at(int i, def value) {                                 
        switch (i) {
            case 0: id = value; return
            case 1: name = value; return
        }
        throw new IllegalArgumentException("No such element $i")
    }
}
def user = new User(id: 1, name: 'Alex')                           
assert user[0] == 1                                                
assert user[1] == 'Alex'                                           
user[1] = 'Bob'                                                    
assert user.name == 'Bob'                                          
```

| **   | the `User` class defines a custom `get_at` implementation |
| ---- | ---------------------------------------- |
| **   | the `User` class defines a custom `put_at` implementation |
| **   | create a sample user                     |
| **   | using the subscript operator with index 0 allows retrieving the user id |
| **   | using the subscript operator with index 1 allows retrieving the user name |
| **   | we can use the subscript operator to write to a property thanks to the delegation to `put_at` |
| **   | and check that it’s really the property `name` which was changed |

#### Membership operator

The membership operator (`in`) is equivalent to calling the `is_case` method. In the context of a `List`, it is equivalent to calling`contains`, like in the following example:

```
def list = ['Grace','Rob','Emmy']
assert ('Emmy' in list)                     
```

| **   | equivalent to calling `list.contains('Emmy')` or `list.is_case('Emmy')` |
| ---- | ---------------------------------------- |
|      |                                          |

#### Diamond operator

The diamond operator (`<>`) is a syntactic sugar only operator added to support compatibility with the operator of the same name in Java 7. It is used to indicate that generic types should be inferred from the declaration:

```
List<String> strings = new LinkedList<>()
```

In dynamic Groovy, this is totally unused. In statically type checked Groovy, it is also optional since the Groovy type checker performs type inference whether this operator is present or not.

#### Call operator

The call operator `()` is used to call a method named `call` implicitly. For any object which defines a `call` method, you can omit the `.call` part and use the call operator instead:

```
class MyCallable {
    int call(int x) {           
        2*x
    }
}

def mc = new MyCallable()
assert mc.call(2) == 4          
assert mc(2) == 4               
```

| **   | `MyCallable` defines a method named `call`. Note that it doesn’t need to implement `java.util.concurrent.Callable` |
| ---- | ---------------------------------------- |
| **   | we can call the method using the classic method call syntax |
| **   | or we can omit `.call` thanks to the call operator |

### Operator precedence

The table below lists all groovy operators in order of precedence.

| Level | Operator(s)                              | Name(s)                                  |
| ----- | ---------------------------------------- | ---------------------------------------- |
| 1     | `new`   `()`                             | object creation, explicit parentheses    |
|       | `()`   `{}`   `[]`                       | method call, closure, literal list/map   |
|       | `.`   `.&`   `.@`                        | member access, method closure, field/attribute access |
|       | `?.`   `*`   `*.`   `*:`                 | safe dereferencing, spread, spread-dot, spread-map |
|       | `~`   `!`   `(type)`                     | bitwise negate/pattern, not, typecast    |
|       | `[]`   `++`   `--`                       | list/map/array index, post inc/decrement |
| 2     | `**`                                     | power                                    |
| 3     | `++`   `--`   `+`   `-`                  | pre inc/decrement, unary plus, unary minus |
| 4     | `*`   `/`   `%`                          | multiply, div, remainder                 |
| 5     | `+`   `-`                                | addition, subtraction                    |
| 6     | `<<`   `>>`   `>>>`   `..`   `..<`       | left/right (unsigned) shift, inclusive/exclusive range |
| 7     | `<`   `<=`   `>`   `>=`   `in`   `instanceof`  `as` | less/greater than/or equal, in, instanceof, type coercion |
| 8     | `==`   `!=`   `<=>`                      | equals, not equals, compare to           |
|       | `=~`   `==~`                             | regex find, regex match                  |
| 9     | `&`                                      | binary/bitwise and                       |
| 10    | `^`                                      | binary/bitwise xor                       |
| 11    | `\|`                                     | binary/bitwise or                        |
| 12    | `&&`                                     | logical and                              |
| 13    | `\|\|`                                   | logical or                               |
| 14    | `? :`                                    | ternary conditional                      |
|       | `?:`                                     | elvis operator                           |
| 15    | `=`   `**=`   `*=`   `/=`   `%=`   `+=`   `-=`  `<<=`   `>>=`   `>>>=`   `&=`   `^=`   `\|=` | various assignments                      |

### Operator overloading

Groovy allows you to overload the various operators so that they can be used with your own classes. Consider this simple class:

```
struct Bucket {
    size : Int
}

func plus(lhs:Bucket, rhs:Bucket) -> Bucket {
	Bucket{lhs.size + rhs.size};  
}

plus = (lhs:Bucket, rhs:Bucket) -> Bucket{lhs.size + rhs.size};
```

| **   | `Bucket` implements a special method called `plus()` |
| ---- | ---------------------------------------- |
|      |                                          |

Just by implementing the `plus()` method, the `Bucket` class can now be used with the `+` operator like so:

```
def b1 = new Bucket(4)
def b2 = new Bucket(11)
assert (b1 + b2).size == 15                         
```

| **   | The two `Bucket` objects can be added together with the `+` operator |
| ---- | ---------------------------------------- |
|      |                                          |

All (non-comparator) Groovy operators have a corresponding method that you can implement in your own classes. The only requirements are that your method is public, has the correct name, and has the correct number of arguments. The argument types depend on what types you want to support on the right hand side of the operator. For example, you could support the statement

```
assert (b1 + 11).size == 15
```

by implementing the `plus()` method with this signature:

```
Bucket plus(int capacity) {
    return new Bucket(this.size + capacity)
}
```

Here is a complete list of the operators and their corresponding methods:

| Operator | Method        | Operator | Method       |
| -------- | ------------- | -------- | ------------ |
| `+`      | a.plus(b)     |          |              |
| `-`      | a.minus(b)    |          |              |
| `*`      | a.multiply(b) | `a in b` | b.is_case(a)  |
| `/`      | a.div(b)      |          |              |
| `%`      | a.mod(b)      |          |              |
| `^`      | a.power(b)    |          |              |
| `or`     | a.or(b)       | `++`     | a.next()     |
| `and`    | a.and(b)      | `--`     | a.previous() |
|          |               | `+a`     | a.positive() |
| `as`     | a.as_type(b)   | `-a`     | a.negative() |
|          |               |          |              |



## Program structure

This chapter covers the program structure of the Groovy programming language.

###  Package names

Package names play exactly the same role as in Java. They allows us to separate the code base without any conflicts. Groovy classes must specify their package before the class definition, else the default package is assumed.

Defining a package is very similar to Java:

```
// defining a package named com.yoursite
package com.yoursite
```

To refer to some class `Foo` in the `com.yoursite.com` package you will need to use the fully qualified name`com.yoursite.com.Foo`, or else you can use an `import` statement as we’ll see below.

### Imports

In order to refer to any class you need a qualified reference to its package. Groovy follows Java’s notion of allowing `import`statement to resolve class references.

For example, Groovy provides several builder classes, such as `MarkupBuilder`. `MarkupBuilder` is inside the package`groovy.xml` so in order to use this class, you need to `import` it as shown:

```
// importing the class MarkupBuilder
import groovy.xml.MarkupBuilder

// using the imported class to create an object
def xml = new MarkupBuilder()

assert xml != null
```

#### Default imports

Default imports are the imports that Groovy language provides by default. For example look at the following code:

```
new Date()
```

The same code in Java needs an import statement to `Date` class like this: import java.util.Date. Groovy by default imports these classes for you.

The below imports are added by groovy for you:

```
import java.lang.*
import java.util.*
import java.io.*
import java.net.*
import groovy.lang.*
import groovy.util.*
import java.math.BigInteger
import java.math.BigDecimal
```

This is done because the classes from these packages are most commonly used. By importing these boilerplate code is reduced.



#### Simple import

A simple import is an import statement where you fully define the class name along with the package. For example the import statement import groovy.xml.MarkupBuilder in the code below is a simple import which directly refers to a class inside a package.

```
// importing the class MarkupBuilder
import groovy.xml.MarkupBuilder

// using the imported class to create an object
def xml = new MarkupBuilder()

assert xml != null
```

#### Star import

Groovy, like Java, provides a special way to import all classes from a package using `*`, the so called star import. `MarkupBuilder`is a class which is in package `groovy.xml`, alongside another class called `StreamingMarkupBuilder`. In case you need to use both classes, you can do:

```
import groovy.xml.MarkupBuilder
import groovy.xml.StreamingMarkupBuilder

def markup_builder = new MarkupBuilder()

assert markup_builder != null

assert new StreamingMarkupBuilder() != null
```

That’s perfectly valid code. But with a `*` import, we can achieve the same effect with just one line. The star imports all the classes under package `groovy.xml`:

```
import groovy.xml.*

def markup_builder = new MarkupBuilder()

assert markup_builder != null

assert new StreamingMarkupBuilder() != null
```

One problem with `*` imports is that they can clutter your local namespace. But with the kinds of aliasing provided by Groovy, this can be solved easily.



#### Import aliasing

With type aliasing, we can refer to a fully qualified class name using a name of our choice. This can be done with the `as` keyword, as before.

For example we can import `java.sql.Date` as `SQLDate` and use it in the same file as `java.util.Date` without having to use the fully qualified name of either class:

```
import java.util.Date
import java.sql.Date as SQLDate

Date util_date = new Date(1000L)
SQLDate sql_date = new SQLDate(1000L)

assert util_date instanceof java.util.Date
assert sql_date instanceof java.sql.Date
```

### Script class

A [script](http://docs.groovy-lang.org/2.4.7/html/gapi/index.html?groovy/lang/Script.html) is always compiled into a class. The Groovy compiler will compile the class for you, with the body of the script copied into a `run` method. The previous example is therefore compiled as if it was the following:

Main.groovy

```
import org.codehaus.groovy.runtime.InvokerHelper
class Main extends Script {                     
    def run() {                                 
        println 'Groovy world!'                 
    }
    static void main(String[] args) {           
        InvokerHelper.run_script(Main, args)     
    }
}
```

| **   | The `Main` class extends the `groovy.lang.Script` class |
| ---- | ---------------------------------------- |
| **   | `groovy.lang.Script` requires a `run` method returning a value |
| **   | the script body goes into the `run` method |
| **   | the `main` method is automatically generated |
| **   | and delegates the execution of the script on the `run` method |

If the script is in a file, then the base name of the file is used to determine the name of the generated script class. In this example, if the name of the file is `Main.groovy`, then the script class is going to be `Main`.

### Methods

It is possible to define methods into a script, as illustrated here:

```
int fib(int n) {
    n < 2 ? 1 : fib(n-1) + fib(n-2)
}
assert fib(10)==89
```

You can also mix methods and code. The generated script class will carry all methods into the script class, and assemble all script bodies into the `run` method:

```
println 'Hello'                                 

int power(int n) { 2**n }                       

println "2^6==${power(6)}"                      
```

| **   | script begins                            |
| ---- | ---------------------------------------- |
| **   | a method is defined within the script body |
| **   | and script continues                     |

This code is internally converted into:

```
import org.codehaus.groovy.runtime.InvokerHelper
class Main extends Script {
    int power(int n) { 2** n}                   
    def run() {
        println 'Hello'                         
        println "2^6==${power(6)}"              
    }
    static void main(String[] args) {
        InvokerHelper.run_script(Main, args)
    }
}
```

| **   | the `power` method is copied as is into the generated script class |
| ---- | ---------------------------------------- |
| **   | first statement is copied into the `run` method |
| **   | second statement is copied into the `run` method |

| **   | Even if Groovy creates a class from your script, it is totally transparent for the user. In particular, scripts are compiled to bytecode, and line numbers are preserved. This implies that if an exception is thrown in a script, the stack trace will show line numbers corresponding to the original script, not the generated code that we have shown. |
| ---- | ---------------------------------------- |
|      |                                          |

##### Variables

Variables in a script do not require a type definition. This means that this script:

```
int x = 1
int y = 2
assert x+y == 3
```

will behave the same as:

```
x = 1
y = 2
assert x+y == 3
```

However there is a semantic difference between the two:

- if the variable is declared as in the first example, it is a *local variable*. It will be declared in the `run` method that the compiler will generate and will **not** be visible outside of the script main body. In particular, such a variable will **not** be visible in other methods of the script
- if the variable is undeclared, it goes into the [script binding](http://docs.groovy-lang.org/2.4.7/html/gapi/index.html?groovy/lang/Script.html#get_binding()). The binding is visible from the methods, and is especially important if you use a script to interact with an application and need to share data between the script and the application. Readers might refer to the [integration guide](http://docs.groovy-lang.org/latest/html/documentation/#_integrating_groovy_in_a_java_application) for more information.

| **   | If you want a variable to become a field of the class without going into the `Binding`, you can use the [@Field annotation](http://docs.groovy-lang.org/latest/html/documentation/#xform-Field). |
| ---- | ---------------------------------------- |
|      |                                          |



### 1.4. Object orientation

This chapter covers the object orientation of the Groovy programming language.

#### 1.4.1. Types

##### Primitive types

Groovy supports the same primitive types as those defined by the [Java Language Specification](http://docs.oracle.com/javase/specs/jls/se8/html/):

- integral types: `byte` (8 bit), `short` (16 bit), `int` (32 bit) and `long` (64 bit)
- floating-point types: `float` (32 bit) and `double` (64 bit)
- `boolean` type (exactly `true` or `false`)
- `char` type (16 bit, usable as a numeric type, representing an UTF-16 code)

While Groovy declares and stores primitive fields and variables as primitives, because it uses Objects for everything, it autowraps references to primitives. Just like Java, the wrappers it uses are

| Primitive type | Wrapper class |
| -------------- | ------------- |
| boolean        | Boolean       |
| char           | Character     |
| short          | Short         |
| int            | Integer       |
| long           | Long          |
| float          | Float         |
| double         | Double        |

Here’s an example using `int`

```
class Foo {
  static int i
}

assert Foo.class.get_declared_field('i').type == int.class
assert Foo.i != int.class && Foo.i.class == Integer.class
```

Now you may be concerned that this means every time you use a mathematical operator on a reference to a primitive that you’ll incur the cost of unboxing and reboxing the primitive. But this is not the case, as Groovy will compile your operators into their[method equivalents](http://docs.groovy-lang.org/latest/html/documentation/core-operators.html#_operator-overloading) and uses those instead. Additionally, Groovy will automatically unbox to a primitive when calling a Java method that takes a primitive parameter and automatically box primitive method return values from Java. However, be aware there are some [differences](http://docs.groovy-lang.org/latest/html/documentation/core-differences-java.html#_primitives_and_wrappers) from Java’s method resolution.

##### Constructors

Constructors are special methods used to initialize an object with a specific state. As in normal methods, it is possible for a class to declare more than one constructor. In Groovy there are two ways to invoke constructors: with positional parameters or named parameters. The former one is like we invoke Java constructors, while the second way allows one to specify the parameter names when invoking the constructor.

###### Positional argument constructor

To create an object by using positional argument constructors, the respective class needs to declare each of the constructors it allows being called. A side effect of this is that, once at least one constructor is declared, the class can only be instantiated by getting one of its constructors called. It is worth noting that, in this case, there is no way to create the class with named parameters.

There is three forms of using a declared constructor. The first one is the normal Java way, with the `new` keyword. The others rely on coercion of lists into the desired types. In this case, it is possible to coerce with the `as` keyword and by statically typing the variable.

```
class PersonConstructor {
    String name
    Integer age

    PersonConstructor(name, age) {          
        this.name = name
        this.age = age
    }
}

def person1 = new PersonConstructor('Marie', 1)  
def person2 = ['Marie', 2] as PersonConstructor  
PersonConstructor person3 = ['Marie', 3]         
```

| **   | Constructor declaration                  |
| ---- | ---------------------------------------- |
| **   | Constructor invocation, classic Java way |
| **   | Constructor usage, using coercion with `as` keyword |
| **   | Constructor usage, using coercion in assignment |

###### Named argument constructor

If no constructor is declared, it is possible to create objects by passing parameters in the form of a map (property/value pairs). This can be in handy in cases where one wants to allow several combinations of parameters. Otherwise, by using traditional positional parameters it would be necessary to declare all possible constructors.

```
class PersonWOConstructor {                                  
    String name
    Integer age
}

def person4 = new PersonWOConstructor()                      
def person5 = new PersonWOConstructor(name: 'Marie')         
def person6 = new PersonWOConstructor(age: 1)                
def person7 = new PersonWOConstructor(name: 'Marie', age: 2) 
```

| **   | No constructor declared                  |
| ---- | ---------------------------------------- |
| **   | No parameters given in the instantiation |
| **   | `name` parameter given in the instantiation |
| **   | `age` parameter given in the instantiation |
| **   | `name` and `age` parameters given in the instantiation |

It is important to highlight, however, that this approach gives more power to the constructor caller, while imposes a major responsibility to it. Thus, if a restriction is needed, one can just declare one or more constructors, and the instantiation by named parameters will no longer be available.

##### Methods

Groovy methods are quite similar to other languages. Some peculiarities will be shown in the next subsections.

###### Method definition

A method is defined with a return type or with the `def` keyword, to make the return type untyped. A method can also receive any number of arguments, which may not have their types explicitly declared. Java modifiers can be used normally, and if no visibility modifier is provided, the method is public.

Methods in Groovy always return some value. If no `return` statement is provided, the value evaluated in the last line executed will be returned. For instance, note that none of the following methods uses the `return` keyword.

```
def some_method() { 'method called' }                           
String another_method() { 'another method called' }             
def third_method(param1) { "$param1 passed" }                   
static String fourth_method(String param1) { "$param1 passed" } 
```

| **   | Method with no return type declared and no parameter |
| ---- | ---------------------------------------- |
| **   | Method with explicit return type and no parameter |
| **   | Method with a parameter with no type defined |
| **   | Static method with a String parameter    |

###### Named arguments

Like constructors, normal methods can also be called with named arguments. They need to receive the parameters as a map. In the method body, the values can be accessed as in normal maps (`map.key`).

```
def foo(Map args) { "${args.name}: ${args.age}" }
foo(name: 'Marie', age: 1)
```

###### Default arguments

Default arguments make parameters optional. If the argument is not supplied, the method assumes a default value.

```
def foo(String par1, Integer par2 = 1) { [name: par1, age: par2] }
assert foo('Marie').age == 1
```

Note that no mandatory parameter can be defined after a default parameter is present, only other default parameters.

###### Varargs

Groovy supports methods with a variable number of arguments. They are defined like this: `def foo(p1, …​, pn, T…​ args)`. Here `foo` supports `n` arguments by default, but also an unspecified number of further arguments exceeding `n`.

```
def foo(Object... args) { args.length }
assert foo() == 0
assert foo(1) == 1
assert foo(1, 2) == 2
```

This example defines a method `foo`, that can take any number of arguments, including no arguments at all. `args.length` will return the number of arguments given. Groovy allows `T[]` as a alternative notation to `T…​`. That means any method with an array as last parameter is seen by Groovy as a method that can take a variable number of arguments.

```
def foo(Object[] args) { args.length }
assert foo() == 0
assert foo(1) == 1
assert foo(1, 2) == 2
```

If a method with varargs is called with `null` as the vararg parameter, then the argument will be `null` and not an array of length one with `null` as the only element.

```
def foo(Object... args) { args }
assert foo(null) == null
```

If a varargs method is called with an array as an argument, then the argument will be that array instead of an array of length one containing the given array as the only element.

```
def foo(Object... args) { args }
Integer[] ints = [1, 2]
assert foo(ints) == [1, 2]
```

Another important point are varargs in combination with method overloading. In case of method overloading Groovy will select the most specific method. For example if a method `foo` takes a varargs argument of type `T` and another method `foo` also takes one argument of type `T`, the second method is preferred.

```
def foo(Object... args) { 1 }
def foo(Object x) { 2 }
assert foo() == 1
assert foo(1) == 2
assert foo(1, 2) == 1
```



###### Fields

A field is a member of a class or a trait which:

- a mandatory *access modifier* (`public`, `protected`, or `private`)
- one or more optional *modifiers* (`static`, `final`, `synchronized`)
- an optional *type*
- a mandatory *name*

```
class Data {
    private int id                                  
    protected String description                    
    public static final boolean DEBUG = false       
}
```

| **   | a `private` field named `id`, of type `int` |
| ---- | ---------------------------------------- |
| **   | a `protected` field named `description`, of type `String` |
| **   | a `public static final` field named *DEBUG* of type `boolean` |

A field may be initialized directly at declaration:

```
class Data {
    private String id = IDGenerator.next() 
    // ...
}
```

| **   | the private field `id` is initialized with `IDGenerator.next()` |
| ---- | ---------------------------------------- |
|      |                                          |

It is possible to omit the type declaration of a field. This is however considered a bad practice and in general it is a good idea to use strong typing for fields:

```
class BadPractice {
    private mapping                         
}
class GoodPractice {
    private Map<String,String> mapping      
}
```

| **   | the field `mapping` doesn’t declare a type |
| ---- | ---------------------------------------- |
| **   | the field `mapping` has a strong type    |

The difference between the two is important if you want to use optional type checking later. It is also important for documentation. However in some cases like scripting or if you want to rely on duck typing it may be interesting to omit the type.



##### Annotation

###### Annotation definition

##### Inheritance

##### Generics



## Function

This chapter covers Groovy Closures. A closure in Groovy is an open, anonymous, block of code that can take arguments, return a value and be assigned to a variable. A closure may reference variables declared in its surrounding scope. In opposition to the formal definition of a closure, `Closure` in the Groovy language can also contain free variables which are defined outside of its surrounding scope. While breaking the formal concept of a closure, it offers a variety of advantages which are described in this chapter.

### Syntax

##### Defining a closure

A closure definition follows this syntax:

```
{ [closure_parameters -> ] statements }
```

Where `[closure_parameters->]` is an optional comma-delimited list of parameters, and statements are 0 or more Groovy statements. The parameters look similar to a method parameter list, and these parameters may be typed or untyped.

When a parameter list is specified, the `->` character is required and serves to separate the arguments from the closure body. The*statements* portion consists of 0, 1, or many Groovy statements.

Some examples of valid closure definitions:

```
{ item++ }                                          

{ -> item++ }                                       

{ println it }                                      

{ it -> println it }                                

{ name -> println name }                            

{ String x, int y ->                                
    println "hey ${x} the value is ${y}"
}

{ reader ->                                         
    def line = reader.read_line()
    line.trim()
}
```

| **   | A closure referencing a variable named `item` |
| ---- | ---------------------------------------- |
| **   | It is possible to explicitly separate closure parameters from code by adding an arrow (`->`) |
| **   | A closure using an implicit parameter (`it`) |
| **   | An alternative version where `it` is an explicit parameter |
| **   | In that case it is often better to use an explicit name for the parameter |
| **   | A closure accepting two typed parameters |
| **   | A closure can contain multiple statements |



#### Parameters

##### Normal parameters

Parameters of closures follow the same principle as parameters of regular methods:

- an optional type
- a name
- an optional default value

Parameters are separated with commas:

```
def closure_with_one_arg = { str -> str.to_upper_case() }
assert closure_with_one_arg('groovy') == 'GROOVY'

def closure_with_one_arg_and_explicit_type = { String str -> str.to_upper_case() }
assert closure_with_one_arg_and_explicit_type('groovy') == 'GROOVY'

def closure_with_two_args = { a,b -> a+b }
assert closure_with_two_args(1,2) == 3

def closure_with_two_args_and_explicit_types = { int a, int b -> a+b }
assert closure_with_two_args_and_explicit_types(1,2) == 3

def closure_with_two_args_and_optional_types = { a, int b -> a+b }
assert closure_with_two_args_and_optional_types(1,2) == 3

def closure_with_two_arg_and_default_value = { int a, int b=2 -> a+b }
assert closure_with_two_arg_and_default_value(1) == 3
```



##### Implicit parameter

When a closure does not explicitly define a parameter list (using `->`), a closure **always** defines an implicit parameter, named `it`. This means that this code:

```
def greeting = { "Hello, $it!" }
assert greeting('Patrick') == 'Hello, Patrick!'
```

is stricly equivalent to this one:

```
def greeting = { it -> "Hello, $it!" }
assert greeting('Patrick') == 'Hello, Patrick!'
```

If you want to declare a closure which accepts no argument and must be restricted to calls without arguments, then you **must**declare it with an explicit empty argument list:

```
def magic_number = { -> 42 }

// this call will fail because the closure doesn't accept any argument
magic_number(11)
```

##### Varargs

It is possible for a closure to declare variable arguments like any other method. *Vargs* methods are methods that can accept a variable number of arguments if the last parameter is of variable length (or an array) like in the next examples:

```
def concat1 = { String... args -> args.join('') }           
assert concat1('abc','def') == 'abcdef'                     
def concat2 = { String[] args -> args.join('') }            
assert concat2('abc', 'def') == 'abcdef'

def multi_concat = { int n, String... args ->                
    args.join('')*n
}
assert multi_concat(2, 'abc','def') == 'abcdefabcdef'
```

| **   | A closure accepting a variable number of strings as first parameter |
| ---- | ---------------------------------------- |
| **   | It may be called using any number of arguments **without** having to explicitly wrap them into an array |
| **   | The same behavior is directly available if the *args* parameter is declared as an array |
| **   | As long as the **last** parameter is an array or an explicit vargs type |



## Semantics

This chapter covers the semantics of the Groovy programming language.

### Statements

#### Variable definition

Variables can be defined using either their type (like `String`) or by using the keyword `def`:

```
String x
def o
```

`def` is a replacement for a type name. In variable definitions it is used to indicate that you don’t care about the type. In variable definitions it is mandatory to either provide a type name explicitly or to use "def" in replacement. This is needed to make variable definitions detectable for the Groovy parser.

You can think of `def` as an alias of `Object` and you will understand it in an instant.

| **   | Variable definition types can be refined by using generics, like in `List names`. To learn more about the generics support, please read the [generics section](http://docs.groovy-lang.org/latest/html/documentation/#generics). |
| ---- | ---------------------------------------- |
|      |                                          |

#### Variable assignment

You can assign values to variables for later use. Try the following:

```
x = 1
println x

x = new java.util.Date()
println x

x = -3.1499392
println x

x = false
println x

x = "Hi"
println x
```

#### Control structures

##### Conditional structures

###### if / else

Groovy supports the usual if - else syntax from Java`def x = falsedef y = falseif ( !x ) {    x = true}assert x == trueif ( x ) {    x = false} else {    y = true}assert x == y`Groovy also supports the normal Java "nested" if then else if syntax:`if ( ... ) {    ...} else if (...) {    ...} else {    ...}`

###### match

The switch statement in Groovy is backwards compatible with Java code; so you can fall through cases sharing the same code for multiple matches.One difference though is that the Groovy switch statement can handle any kind of switch value and different kinds of matching can be performed.`def x = 1.23
def result = ""

switch ( x ) {
​    case "foo":
​        result = "found foo"
​        // lets fall through

​    case "bar":
​        result += "bar"

​    case [4, 5, 6, 'in_list']:
​        result = "list"
​        break

​    case 12..30:
​        result = "range"
​        break

​    case Integer:
​        result = "integer"
​        break

​    case Number:
​        result = "number"
​        break

​    case ~/fo*/: // to_string() representation of x matches the pattern?
​        result = "foo regex"
​        break

​    case { it < 0 }: // or { x < 0 }
​        result = "negative"
​        break

​    default:
​        result = "default"
}

assert result == "number"`Switch supports the following kinds of comparisons:Class case values match if the switch value is an instance of the class_regular expression case values match if the `to_string()` representation of the switch value matches the regex_collection case values match if the switch value is contained in the collection. This also includes ranges (since they are Lists)Closure case values match if the calling the closure returns a result which is true according to the [Groovy truth](http://docs.groovy-lang.org/latest/html/documentation/#Groovy-Truth)If none of the above are used then the case value matches if the case value equals the switch value**When using a closure case value, the default `it` parameter is actually the switch value (in our example, variable`x`).

##### Looping structures

###### Classic for loop

Groovy supports the standard Java / C for loop:`String message = ''for (int i = 0; i < 5; i++) {    message += 'Hi '}assert message == 'Hi Hi Hi Hi Hi '`

###### for in loop

The for loop in Groovy is much simpler and works with any kind of array, collection, Map, etc.`// iterate over a range
def x = 0
for ( i in 0..9 ) {
​    x += i
}
assert x == 45

// iterate over a list
x = 0
for ( i in [0, 1, 2, 3, 4] ) {
​    x += i
}
assert x == 10

// iterate over an array
def array = (0..4).to_array()
x = 0
for ( i in array ) {
​    x += i
}
assert x == 10

// iterate over a map
def map = ['abc':1, 'def':2, 'xyz':3]
x = 0
for ( e in map ) {
​    x += e.value
}
assert x == 6

// iterate over values in a map
x = 0
for ( v in map.values() ) {
​    x += v
}
assert x == 6

// iterate over the characters in a string
def text = "abc"
def list = []
for (c in text) {
​    list.add(c)
}
assert list == ["a", "b", "c"]`**Groovy also supports the Java colon variation with colons: `for (char c : text) {}`, where the type of the variable is mandatory.

##### while loop

Groovy supports the usual while {…} loops like Java:`def x = 0
def y = 5

while ( y-- > 0 ) {
​    x++
}

assert x == 5`

#### Power assertion

Unlike Java with which Groovy shares the `assert` keyword, the latter in Groovy behaves very differently. First of all, an assertion in Groovy is always executed, independently of the `-ea` flag of the JVM. It makes this a first class choice for unit tests. The notion of "power asserts" is directly related to how the Groovy `assert` behaves.

A power assertion is decomposed into 3 parts:

```
assert [left expression] == [right expression] : (optional message)
```

The result of the assertion is very different from what you would get in Java. If the assertion is true, then nothing happens. If the assertion is false, then it provides a visual representation of the value of each sub-expressions of the expression being asserted. For example:

```
assert 1+1 == 3
```

Will yield:

```
Caught: Assertion failed:

assert 1+1 == 3
        |  |
        2  false
```

Power asserts become very interesting when the expressions are more complex, like in the next example:

```
def x = 2
def y = 7
def z = 5
def calc = { a,b -> a*b+1 }
assert calc(x,y) == [x,z].sum()
```

Which will print the value for each sub-expression:

```
assert calc(x,y) == [x,z].sum()
       |    | |  |   | |  |
       15   2 7  |   2 5  7
                 false
```

In case you don’t want a pretty printed error message like above, you can fallback to a custom error message by changing the optional message part of the assertion, like in this example:

```
def x = 2
def y = 7
def z = 5
def calc = { a,b -> a*b+1 }
assert calc(x,y) == z*z : 'Incorrect computation result'
```

Which will print the following error message:

```
Incorrect computation result. Expression: (calc.call(x, y) == (z * z)). Values: z = 5, z = 
```



### Expressions

#### MPath expressions

`GPath` is a path expression language integrated into Groovy which allows parts of nested structured data to be identified. In this sense, it has similar aims and scope as XPath does for XML. GPath is often used in the context of processing XML, but it really applies to any object graph. Where XPath uses a filesystem-like path notation, a tree hierarchy with parts separated by a slash `/`, GPath **use a dot-object notation** to perform object navigation.

As an example, you can specify a path to an object or element of interest:

- `a.b.c` → for XML, yields all the `c` elements inside `b` inside `a`
- `a.b.c` → for POJOs, yields the `c` properties for all the `b` properties of `a` (sort of like `a.get_b().get_c()` in JavaBeans)

In both cases, the GPath expression can be viewed as a query on an object graph. For POJOs, the object graph is most often built by the program being written through object instantiation and composition; for XML processing, the object graph is the result of`parsing` the XML text, most often with classes like XmlParser or XmlSlurper. See <<../../../subprojects/groovy-xml/src/spec/doc/xml-userguide.adoc#Processing XML,Processing XML>> for more in-depth details on consuming XML in Groovy.

| **   | When querying the object graph generated from XmlParser or XmlSlurper, a GPath expression can refer to attributes defined on elements with the `@` notation:`a["@href"]` → map-like notation : the href attribute of all the a elements`a.'@href'` → property notation : an alternative way of expressing this`a.@href` → direct notation : yet another alternative way of expressing this |
| ---- | ---------------------------------------- |
|      |                                          |



#### Object navigation

Let’s see an example of a GPath expression on a simple *object graph*, the one obtained using java reflection. Suppose you are in a non-static method of a class having another method named `a_method_foo`

```
void a_method_foo() { println "This is a_method_foo." } 
```

the following GPath expression will get the name of that method:

```
assert ['a_method_foo'] == this.class.methods.name.grep(~/.*Foo/)
```

*More precisely*, the above GPath expression produces a list of String, each being the name of an existing method on `this` where that name ends with `Foo`.

Now, given the following methods also defined in that class:

```
void a_method_bar() { println "This is a_method_bar." } 
void another_foo_method() { println "This is another_foo_method." } 
void a_second_method_bar() { println "This is a_second_method_bar." } 
```

then the following GPath expression will get the names of **(1)** and **(3)**, but not **(2)** or **(0)**:

```
assert ['a_method_bar', 'a_second_method_bar'] as Set == this.class.methods.name.grep(~/.*Bar/) as Set
```

### Promotion and coercion

#### Number promotion

The rules of number promotion are specified in the section on [math operations](http://docs.groovy-lang.org/latest/html/documentation/#_math_operations).



#### Map to type coercion

Usually using a single closure to implement an interface or a class with multiple methods is not the way to go. As an alternative, Groovy allows you to coerce a map into an interface or a class. In that case, keys of the map are interpreted as method names, while the values are the method implementation. The following example illustrates the coercion of a map into an `Iterator`:

```
def map
map = [
  i: 10,
  has_next: { map.i > 0 },
  next: { map.i-- },
]
def iter = map as Iterator
```

Of course this is a rather contrived example, but illustrates the concept. You only need to implement those methods that are actually called, but if a method is called that doesn’t exist in the map a `MissingMethodException` or an`UnsupportedOperationException` is thrown, depending on the arguments passed to the call, as in the following example:

```
interface X {
    void f()
    void g(int n)
    void h(String s, int n)
}

x = [ f: {println "f called"} ] as X
x.f() // method exists
x.g() // MissingMethodException here
x.g(5) // UnsupportedOperationException here
```

The type of the exception depends on the call itself:

- `MissingMethodException` if the arguments of the call do not match those from the interface/class
- `UnsupportedOperationException` if the arguments of the call match one of the overloaded methods of the interface/class

##### String to enum coercion

Groovy allows transparent `String` (or `GString`) to enum values coercion. Imagine you define the following enum:

```
enum State {
    up,
    down
}
```

then you can assign a string to the enum without having to use an explicit `as` coercion:

```
State st = 'up'
assert st == State.up
```

It is also possible to use a `GString` as the value:

```
def val = "up"
State st = "${val}"
assert st == State.up
```

However, this would throw a runtime error (`IllegalArgumentException`):

```
State st = 'not an enum value'
```

Note that it is also possible to use implicit coercion in switch statements:

```
State switch_state(State st) {
    switch (st) {
        case 'up':
            return State.down // explicit constant
        case 'down':
            return 'up' // implicit coercion for return types
    }
}
```

in particular, see how the `case` use string constants. But if you call a method that uses an enum with a `String` argument, you still have to use an explicit `as` coercion:

```
assert switch_state('up' as State) == State.down
assert switch_state(State.down) == State.up
```



### Optionality



#### 1.6.5. The Groovy Truth

Groovy decides whether a expression is true or false by applying the rules given below.

##### Boolean expressions

True if the corresponding Boolean value is `true`.

```
assert true
assert !false
```

##### Collections and Arrays

Non-empty Collections and arrays are true.

```
assert [1, 2, 3]
assert ![]
```

##### Matchers

True if the Matcher has at least one match.

```
assert ('a' =~ /a/)
assert !('a' =~ /b/)
```

##### Iterators and Enumerations

Iterators and Enumerations with further elements are coerced to true.

```
assert [0].iterator()
assert ![].iterator()
Vector v = [0] as Vector
Enumeration enumeration = v.elements()
assert enumeration
enumeration.next_element()
assert !enumeration
```

##### Maps

Non-empty Maps are evaluated to true.

```
assert ['one' : 1]
assert ![:]
```

##### Strings

Non-empty Strings, GStrings and CharSequences are coerced to true.

```
assert 'a'
assert !''
def non_empty = 'a'
assert "$non_empty"
def empty = ''
assert !"$empty"
```

##### Numbers

Non-zero numbers are true.

```
assert 1
assert 3.5
assert !0
```

##### Object References

Non-null object references are coerced to true.

```
assert new Object()
assert !null
```

##### Customizing the truth with as_boolean() methods

In order to customize whether groovy evaluates your object to `true` or `false` implement the `as_boolean()` method:

```
class Color {
    String name

    boolean as_boolean(){
        name == 'green' ? true : false
    }
}
```

Groovy will call this method to coerce your object to a boolean value, e.g.:

```
assert new Color(name: 'green')
assert !new Color(name: 'red')
```



### Typing

##### Optional typing

Result typing



### Error process



