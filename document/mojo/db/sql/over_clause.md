| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `frame` | `mojo.db.sql.Window.Frame` |  | N |  |  |
| `name` | `string` |  | N |  |
| `orderBy` | `mojo.db.sql.OrderByClause` |  | N |  |  |
| `partition` | `mojo.db.sql.PartitionClause` |  | N |  |  |
| `endPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `implicit` | `boolean` |  | N |  | Implicit - Whether this statement is implicit. |
| `kind` | `integer` | `Int64` | N |  | Kind - The subclass of Stmt that this is. |
| `startPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
