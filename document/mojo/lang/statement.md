| type | format | description |
|---|---|---|
| `mojo.lang.ReturnStmt` |  | ReturnStmt - A return statement.  The result is optional; "return" withoutan expression is semantically equivalent to "return ()".return 42 |
| `mojo.lang.BreakStmt` |  | BreakStmt - The "break" statement. |
| `mojo.lang.ContinueStmt` |  | ContinueStmt - The "continue" statement. |
| `mojo.lang.MatchStmt` |  | Match statement. |
| `mojo.lang.IfStmt` |  | IfStmt - if/then/else statement.  If no 'else' is specified, then theElseLoc location is not specified and the Else statement is null. Aftertype-checking, the condition is of type Builtin.Int1. |
| `mojo.lang.ForStmt` |  | ForStmt - foreach statement that iterates over the elements in acontainer. |
| `mojo.lang.WhileStmt` |  | WhileStmt - while statement. After type-checking, the condition is oftype Builtin.Int1. |
| `mojo.lang.RepeatStmt` |  | RepeatStmt - repeat/while statement. After type-checking, thecondition is of type Builtin.Int1. |
| `mojo.lang.BlockStmt` |  | BlockStmt - A brace enclosed sequence of expressions, statements, or declarations, like{ var x = 10; print(x) }. |
| `mojo.lang.Declaration` |  |  |
| `mojo.lang.Expression` |  |  |
