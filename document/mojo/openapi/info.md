| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `title` | `string` |  | Y |  | The title of the application. |
| `description` | `mojo.openapi.CachedDocument` |  | N |  | A short description of the application.CommonMark syntax syntax MAY be used for rich text representation. |
| `termsOfService` | `string` |  | N |  | A URL to the Terms of Service for the API. |
| `contact` | `mojo.openapi.Contact` |  | N |  | The contact information for the exposed API. |
| `license` | `mojo.openapi.License` |  | N |  | The license information for the exposed API. |
| `version` | `string` | `Version` | Y |  | Provides the version of the application API (not to be confused with the specification version) |
