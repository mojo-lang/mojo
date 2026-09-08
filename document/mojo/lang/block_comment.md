| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | position of first character belonging to the Document |
| `endPosition` | `mojo.lang.Position` |  | N |  | position of first character immediately after the Document |
| `content` | `string` |  | N |  | the content of the block comment |
| `headEmbedded` | `boolean` |  | N |  | the head /* is after some code in the same line |
| `tailEmbedded` | `boolean` |  | N |  | the tail */ is before some code in the same line |
