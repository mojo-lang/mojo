| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | position of first character belonging to the Document |
| `endPosition` | `mojo.lang.Position` |  | N |  | position of first character immediately after the Document |
| `following` | `boolean` |  | N |  | all the lines are following after the decl, but not at the beginning of the line |
| `private` | `boolean` |  | N |  | private or inner document, which will not to publish to public |
| `lines` | `Array<mojo.lang.Document.Line>` |  | N |  | original line documents |
| `structured` | `mojo.document.Document` |  | N |  | structured document for all the lines |
