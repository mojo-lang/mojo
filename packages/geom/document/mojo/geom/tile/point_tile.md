| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  | tile quad-key |
| `ids` | `Array<integer>` |  | N |  | for string ids, using keys or md5(string id) |
| `keys` | `Array<string>` |  | N |  | Dictionary encoding for keys |
| `properties` | `Map<string, mojo.geom.tile.VectorTile.Value>` |  | N |  | the properties for this tile |
| `rawValues` | `Array<integer>` |  | N |  | attribute index for which value is raw integer |
| `tags` | `Array<integer>` |  | N |  | Tags of this feature are encoded as repeated pairs ofintegers. |
| `type` | `string` |  | N |  | the point data typethe user defined point data type |
| `values` | `Array<mojo.geom.tile.VectorTile.Value>` |  | N |  | Dictionary encoding for values |
| `version` | `integer` | `Int32` | N |  | the version of the point tile specification |
| `xs` | `Array<integer>` |  | N |  | the coordinates for the point, using E7 of the latitude & longitude |
| `ys` | `Array<integer>` |  | N |  |
