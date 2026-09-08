| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `caption` | `Array<mojo.document.Inline>` |  | N |  | caption for table |
| `alignment` | `string` |  | N |  | column alignments (required) |
| `width` | `number` | `Float64` | N |  | relative column widths |
| `header` | `Array<Array<mojo.document.Block>>` |  | N |  | table header, each is column header (each a list of blocks) |
| `rows` | `Array<Array<Array<mojo.document.Block>>>` |  | N |  | rows, a list of row |
