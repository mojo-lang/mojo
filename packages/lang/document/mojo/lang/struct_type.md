| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | position of the character "{" belonging to the StructType |
| `endPosition` | `mojo.lang.Position` |  | N |  | position of the character "}" |
| `fields` | `Array<mojo.lang.ValueDecl>` |  | N |  | A Field represents a Field declaration list in a struct type,a method list in an interface type, or a parameter/result declarationin a signature. |
| `inherits` | `Array<mojo.lang.NominalType>` |  | N |  |  |
| `groups` | `Array<mojo.lang.GroupDecl>` |  | N |  |  |
| `inheritPosition` | `mojo.lang.Position` |  | N |  |  |
