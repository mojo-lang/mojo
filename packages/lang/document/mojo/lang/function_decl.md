| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  |  |
| `fullName` | `string` |  | N |  |  |
| `attributes` | `Array<mojo.lang.Attribute>` |  | N |  |  |
| `genericParameters` | `Array<mojo.lang.GenericParameter>` |  | N |  |  |
| `enclosing` | `mojo.lang.NominalType` |  | N |  |  |
| `namePosition` | `mojo.lang.Position` |  | N |  |  |
| `signature` | `mojo.lang.FunctionSignature` |  | Y |  |  |
| `body` | `mojo.lang.BlockStmt` |  | N |  |  |
| `receiver` | `mojo.lang.NominalType` |  | N |  |  |
| `scope` | `mojo.lang.Scope` |  | N |  |  |
| `interfaceDecl` | `mojo.lang.InterfaceDecl` |  | N |  |  |
