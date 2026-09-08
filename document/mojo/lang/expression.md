| type | format | description |
|---|---|---|
| `mojo.lang.NullLiteralExpr` |  | The 'none' literal. |
| `mojo.lang.IntegerLiteralExpr` |  | \brief Integer literal with a '+' or '-' sign, like '+4' or '- 2'. |
| `mojo.lang.FloatLiteralExpr` |  | FloatLiteralExpr - Floating point literal, like '4.0'.  After semanticanalysis assigns types, this is guaranteed to only have aBuiltinFloatingPointType. |
| `mojo.lang.BoolLiteralExpr` |  | A Boolean literal ('true' or 'false') |
| `mojo.lang.StringLiteralExpr` |  | StringLiteralExpr - String literal, like '"foo"'. |
| `mojo.lang.ObjectLiteralExpr` |  |  |
| `mojo.lang.ArrayLiteralExpr` |  | An array literal expression [a, b, c]. |
| `mojo.lang.MapLiteralExpr` |  | \brief A map literal expression {"a" : x, "b" : y, "c" : z}. |
| `mojo.lang.RangeLiteralExpr` |  | RangeLiteralExpr - Range literal, like '1..4'. |
| `mojo.lang.IdentifierExpr` |  |  |
| `mojo.lang.NumericSuffixLiteralExpr` |  |  |
| `mojo.lang.StringPrefixLiteralExpr` |  |  |
| `mojo.lang.StringSuffixLiteralExpr` |  |  |
| `mojo.lang.StructLiteralExpr` |  |  |
| `mojo.lang.ClosureExpr` |  |  |
| `mojo.lang.ParenthesizedExpr` |  | A parenthesized expression like '(x+x)'.  Syntactically,this is just a TupleExpr with exactly one element that has no label.Semantically, however, it serves only as grouping parentheses anddoes not form an expression of tuple type (unless the sub-expressionhas tuple type, of course). |
| `mojo.lang.ImplicitMemberExpr` |  |  |
| `mojo.lang.WildcardExpr` |  |  |
| `mojo.lang.StructConstructionExpr` |  |  |
| `mojo.lang.TupleExpr` |  | TupleExpr - Parenthesized expressions like `(a: x+x)` and '(x, y, 4)'.  Alsoused to represent the operands to a binary operator.  Note thatexpressions like '(4)' are represented with a ParenExpr. |
| `mojo.lang.PrefixUnaryExpr` |  | PrefixUnaryExpr - Prefix unary expressions like '!y'. |
| `mojo.lang.PostfixUnaryExpr` |  | PostfixUnaryExpr - postfix unary expressions like 'y++'. |
| `mojo.lang.FunctionCallExpr` |  |  |
| `mojo.lang.ExplicitMemberExpr` |  | MemberRefExpr - This represents 'a.b' where we are referring to a memberof a type, such as a property or variable. |
| `mojo.lang.SubscriptExpr` |  | Subscripting expressions like a[i] that refer to an element within acontainer. |
| `mojo.lang.BinaryExpr` |  | BinaryExpr - Infix binary expressions like 'x+y'.  The argument is alwaysan implicit tuple expression of the type expected by the function. |
| `mojo.lang.ConditionalExpr` |  | The conditional expression 'x ? y : z'. |
| `mojo.lang.TypeCastingExpr` |  |  |
| `mojo.lang.AssignmentExpr` |  | AssignExpr - A value assignment, like "x = y". |
| `mojo.lang.ErrorExpr` |  | ErrorExpr - Represents a semantically erroneous subexpression in the AST,typically this will have an ErrorType. |
