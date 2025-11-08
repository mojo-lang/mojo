| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | the name of the file, usually it is the full path of the file |
| `isDir` | `boolean` |  | N |  | whether is the directory or not |
| `mode` | `string` |  | N |  | the value must be const to "dir" |
| `info` | `mojo.core.File.Info` |  | N |  |  |
| `files` | `Array<mojo.core.File>` |  | N |  | sub files in this file if is a directory |
