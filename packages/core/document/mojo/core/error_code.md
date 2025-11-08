| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `code` | `integer` | `Int32` | Y |  |  |
| `name` | `string` |  | N |  | The name of the error code. This is a constant value that identifies theerror code. Error code name are unique within a particulardomain of errors. This should be at most 63 characters and match/[A-Z0-9_]+/. |
| `domain` | `string` |  | N |  | system, runtime, ... |
| `description` | `string` |  | N |  | a detail description for the code |
| `document` | `string` | `Url` | N |  | the api document url for this error code |
| `httpStatusCode` | `integer` | `Int32` | N |  | the http status code, which is error code mapping to |
