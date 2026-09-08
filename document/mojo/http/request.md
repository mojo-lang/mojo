| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `method` | `string` |  | N |  | method specifies the HTTP method (GET, POST, PUT, etc.).For client requests, an empty string means GET. |
| `url` | `string` | `Url` | N |  | url specifies either the URI being requested (for serverrequests) or the URL to access (for client requests). |
| `version` | `mojo.http.Version` |  | N |  | The protocol version for incoming server requests. |
| `headers` | `Map<string, Array<string>>` |  | N |  | Headers contains the request header fields either receivedby the server or to be sent by the client. |
| `body` | `mojo.core.Value` |  | N |  | body is the request's body, which ban be raw bytes or JSON object |
