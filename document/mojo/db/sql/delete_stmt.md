| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `endPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `implicit` | `boolean` |  | N |  | Implicit - Whether this statement is implicit. |
| `kind` | `integer` | `Int64` | N |  | Kind - The subclass of Stmt that this is. |
| `startPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `returning` | `mojo.db.sql.ReturningClause` |  | N |  |  |
| `tableName` | `mojo.db.sql.QualifiedTableName` |  | N |  |  |
| `where` | `mojo.db.sql.WhereClause` |  | N |  |  |
| `with` | `mojo.db.sql.WithClause` |  | N |  |  |
