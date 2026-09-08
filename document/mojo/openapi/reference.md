| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `ref` | `string` | `Url` | Y |  | The reference Url string. |
| `summary` | `string` |  | N |  | A short summary which by default SHOULD override that of the referenced component. If the referenced object-type does not allow a summary field, then this field has no effect. |
| `description` | `mojo.openapi.CachedDocument` |  | N |  | A description which by default SHOULD override that of the referenced component. CommonMark syntax MAY be used for rich text representation. If the referenced object-type does not allow a description field, then this field has no effect. |
