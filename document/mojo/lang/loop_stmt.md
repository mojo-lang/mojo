| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `endPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `kind` | `integer` | `Int64` | N |  | Kind - The subclass of Stmt that this is. |
| `implicit` | `boolean` |  | N |  | Implicit - Whether this statement is implicit. |
| `body` | `mojo.lang.BlockStmt` |  | N |  | BlockStmt - A brace enclosed sequence of expressions, statements, or declarations, like{ var x = 10; print(x) }. |
