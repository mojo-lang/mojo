| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `service` | `string` |  | N |  | The API service name. It is a logical identifier for a networked API,such as "pubsub.googleapis.com". The naming syntax depends on theAPI management system being used for handling the request. |
| `operation` | `string` |  | N |  | The API operation name. For gRPC requests, it is the fully qualified APImethod name, such as "google.pubsub.v1.Publisher.Publish". For OpenAPIrequests, it is the `operationId`, such as "getPet". |
| `protocol` | `string` |  | N |  | The API protocol used for sending the request, such as "http", "https","grpc", or "internal". |
| `version` | `string` |  | N |  | The API version associated with the API operation above, such as "v1" or"v1alpha1". |
