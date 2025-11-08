| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `enum` | `Array<string>` |  | N |  | An enumeration of string values to be used if the substitution options are from a limited set. |
| `default` | `string` |  | Y |  | The default value to use for substitution, and to send, if an alternate value is ** supplied.Unlike the `Schema` Object's `default`, this value ** be provided by the consumer. |
| `description` | `mojo.openapi.CachedDocument` |  | N |  | An optional description for the server variable.CommonMark syntax MAY be used for rich text representation. |
