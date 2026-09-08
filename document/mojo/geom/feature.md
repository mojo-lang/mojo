| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `type` | `string` |  | Y |  | the value must be const to "Feature" |
| `id` | `mojo.core.Id` |  | N |  |
| `geometry` | `mojo.geom.Geometry` |  | N |  |
| `bbox` | `mojo.geom.BoundingBox` |  | N |  |  |
| `crs` | `string` |  | N |  |
| `properties` | `Map<string, mojo.core.Value>` |  | N |  |
