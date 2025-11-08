| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  |  |
| `fullName` | `string` |  | N |  |  |
| `attributes` | `Array<mojo.lang.Attribute>` |  | N |  |  |
| `genericParameters` | `Array<mojo.lang.GenericParameter>` |  | N |  |  |
| `group` | `mojo.lang.GroupDecl` |  | N |  |  |
| `resolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | N |  |  |
| `unresolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | N |  | unresolved identifiers in this file |
| `namePosition` | `mojo.lang.Position` |  | N |  |  |
| `attribute` | `mojo.lang.Attribute` |  | N |  |  |
| `scope` | `mojo.lang.Scope` |  | N |  | the 'generic-parameter' type identifier will be in here |
