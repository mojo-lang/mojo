| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | base name of the file |
| `suffix` | `string` |  | N |  | suffix of the file name |
| `size` | `integer` | `Int64` | N |  | length in bytes for regular files; system-dependent for others |
| `changeTime` | `string` | `Timestamp` | N |  | 文件的元数据发生变化。比如权限，所有者等 |
| `modifyTime` | `string` | `Timestamp` | N |  | 文件内容被修改的最后时间 |
