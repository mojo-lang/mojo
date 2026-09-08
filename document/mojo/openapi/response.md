| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `description` | `string` |  | Y |  | A short description of the response. CommonMark syntax MAY be used for rich text representation. |
| `headers` | `Map<string, mojo.openapi.ReferenceableHeader>` |  | N |  | Maps a header name to its definition.RFC7230 states header names are case insensitive.If a response header is defined with the name "Content-Type", it SHALL be ignored. |
| `content` | `Map<string, mojo.openapi.MediaType>` |  | N |  | A map containing descriptions of potential response payloads.The key is a media type or media type range and the value describes it.For responses that match multiple keys, only the most specific key is applicable.e.g. text/plain overrides text/* |
| `links` | `Map<string, mojo.openapi.ReferenceableLink>` |  | N |  | A map of operations links that can be followed from the response.The key of the map is a short name for the link,following the naming constraints of the names for Component Objects. |
