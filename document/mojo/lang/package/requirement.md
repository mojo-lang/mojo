| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `version` | `mojo.lang.Package.Requirement.Version` |  | N |  | ^1.2.3~1.2.3<br> |
| `registry` | `string` |  | N |  |  |
| `path` | `string` |  | N |  |  |
| `repository` | `string` | `Url` | N |  |  |
| `branch` | `string` |  | N |  |  |
| `commit` | `mojo.lang.Package.Requirement.Commit` |  | N |  | will using the commit if the version is 0.0.0 |
