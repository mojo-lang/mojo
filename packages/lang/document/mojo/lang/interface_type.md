| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `endPosition` | `mojo.lang.Position` |  | N |  | Positions<br>Position describes an arbitrary source positionincluding the file, line, and column location.A Position is valid if the line number is > 0. |
| `methods` | `Array<mojo.lang.FunctionDecl>` |  | N |  |  |
| `inherits` | `Array<mojo.lang.NominalType>` |  | N |  |  |
| `inheritePosition` | `mojo.lang.Position` |  | N |  |  |
