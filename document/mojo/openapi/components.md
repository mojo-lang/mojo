| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `schemas` | `Map<string, mojo.openapi.Schema>` |  | N |  | An object to hold reusable Schema Objects. |
| `responses` | `Map<string, mojo.openapi.Response>` |  | N |  | An object to hold reusable Response Objects. |
| `parameters` | `Map<string, mojo.openapi.Parameter>` |  | N |  | An object to hold reusable Parameter Objects. |
| `examples` | `Map<string, mojo.openapi.Example>` |  | N |  | An object to hold reusable Example Objects. |
| `requestBodies` | `Map<string, mojo.openapi.RequestBody>` |  | N |  | An object to hold reusable Request Body Objects. |
| `headers` | `Map<string, mojo.openapi.Header>` |  | N |  | An object to hold reusable Header Objects. |
| `securitySchemes` | `Map<string, mojo.openapi.SecurityScheme>` |  | N |  | An object to hold reusable Security Scheme Objects. |
| `links` | `Map<string, mojo.openapi.Link>` |  | N |  | An object to hold reusable Link Objects. |
| `callbacks` | `Map<string, Map<string, mojo.openapi.PathItem>>` |  | N |  | An object to hold reusable Callback Objects. |
| `pathItems` | `Map<string, mojo.openapi.PathItem>` |  | N |  | An object to hold reusable Path Item Object. |
