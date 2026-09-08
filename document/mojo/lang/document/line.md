| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | position of first character belonging to the LineDocument |
| `endPosition` | `mojo.lang.Position` |  | N |  | position of first character immediately after the LineDocument |
| `following` | `boolean` |  | N |  | following after the decl, but not at the beginning of the line |
| `private` | `boolean` |  | N |  | private or inner document, which will not to publish to public |
| `content` | `string` |  | N |  | the content of the line comment, remove the prefix and the EOL |
