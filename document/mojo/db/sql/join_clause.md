| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `endPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `implicit` | `boolean` |  | N |  | Implicit - Whether this statement is implicit. |
| `kind` | `integer` | `Int64` | N |  | Kind - The subclass of Stmt that this is. |
| `startPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `constraint` | `mojo.db.sql.JoinClause.Constraint` |  | N |  |  |
| `left` | `mojo.db.sql.JoinClause` |  | N |  |  |
| `operator` | `string` |  | N |  |
| `right` | `mojo.db.sql.JoinClause` |  | N |  |  |
| `table` | `mojo.db.sql.Table` |  | N |  |
