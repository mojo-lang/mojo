| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `filename` | `string` |  | N |  | filename, if any |
| `offset` | `integer` | `Int64` | N |  | offset, starting at 0 (byte count) |
| `line` | `integer` | `Int64` | N |  | line number, starting at 1 |
| `column` | `integer` | `Int64` | N |  | column number, starting at 1 (byte count) |
| `leadingComments` | `Array<mojo.lang.Comment>` |  | N |  | the comment before the AST node |
| `tailingComments` | `Array<mojo.lang.Comment>` |  | N |  | usually exist in the end position of the AST nodeand only one of the following situations:<br>```mojo
   a = 5 // following line comment 

```
 |
