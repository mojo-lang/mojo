| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  | The unique ID for a request, which can be propagated to downstreamsystems. The ID should have low probability of collisionwithin a single day for a specific service. |
| `method` | `string` |  | N |  | The HTTP request method, such as `GET`, `POST`. |
| `headers` | `Map<string, string>` |  | N |  | The HTTP request headers. If multiple headers share the same key, theymust be merged according to the HTTP spec. All header keys must belowercased, because HTTP header keys are case-insensitive. |
| `path` | `string` |  | N |  | The HTTP URL path. |
| `host` | `string` |  | N |  | The HTTP request `Host` header value. |
| `scheme` | `string` |  | N |  | The HTTP URL scheme, such as `http` and `https`. |
| `query` | `string` |  | N |  | The HTTP URL query in the format of `name1=value1&name2=value2`, as itappears in the first line of the HTTP request. No decoding is performed. |
| `time` | `string` | `Timestamp` | N |  | The timestamp when the `destination` service receives the last byte ofthe request. |
| `size` | `integer` | `Int64` | N |  | The HTTP request size in bytes. If unknown, it must be -1. |
| `protocol` | `string` |  | N |  | The network protocol used with the request, such as "http/1.1","spdy/3", "h2", "h2c", "webrtc", "tcp", "udp", "quic". Seefor details. |
| `reason` | `string` |  | N |  | A special parameter for request reason. It is used by security systemsto associate auditing information with a request. |
| `auth` | `mojo.rpc.AttributeContext.Auth` |  | N |  | The request authentication. May be absent for unauthenticated requests.Derived from the HTTP request `Authorization` header or equivalent. |
