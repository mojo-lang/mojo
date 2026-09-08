| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `description` | `mojo.openapi.CachedDocument` |  | N |  | A brief description of the request body.This could contain examples of use. CommonMark syntax MAY be used for rich text representation. |
| `content` | `Map<string, mojo.openapi.MediaType>` |  | Y |  | The content of the request body.The key is a media type or media type range and the value describes it.For requests that match multiple keys, only the most specific key is applicable.e.g. text/plain overrides text/* |
| `required` | `boolean` |  | N |  | Determines if the request body is required in the request. |
