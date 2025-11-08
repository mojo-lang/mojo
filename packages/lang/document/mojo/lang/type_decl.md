| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  |  |
| `attributes` | `Array<mojo.lang.Attribute>` |  | N |  |  |
| `genericParameters` | `Array<mojo.lang.GenericParameter>` |  | N |  |  |
| `enclosing` | `mojo.lang.NominalType` |  | N |  |  |
| `group` | `mojo.lang.GroupDecl` |  | N |  |  |
| `resolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | N |  |  |
| `unresolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | N |  | unresolved identifiers in this file |
| `namePosition` | `mojo.lang.Position` |  | N |  | the start position for name part |
