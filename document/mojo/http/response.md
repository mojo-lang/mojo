| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `version` | `mojo.http.Version` |  | N |  | The protocol version for incoming server requests. |
| `status` | `mojo.http.Status` |  | N |  |  |
| `headers` | `Map<string, Array<string>>` |  | N |  | Headers contains the request header fields either receivedby the server or to be sent by the client. |
| `body` | `mojo.core.Value` |  | N |  | body is the request's body, which ban be raw bytes or JSON object |
