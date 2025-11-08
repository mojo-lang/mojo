| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `summary` | `string` |  | N |  | Short description for the example. |
| `description` | `string` |  | N |  | Long description for the example. |
| `value` | `mojo.core.Value` |  | N |  | Embedded literal example.The value field and externalValue field are mutually exclusive.To represent examples of media types that cannot naturally represented in JSON or YAML,use a string value to contain the example, escaping where necessary. |
| `externalValue` | `string` | `Url` | N |  | A URL that points to the literal example.This provides the capability to reference examples that cannot easily be included in JSON or YAML documents.The value field and externalValue field are mutually exclusive. |
