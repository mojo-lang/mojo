| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `code` | `integer` | `Int64` | N |  | The HTTP response status code, such as `200` and `404`. |
| `size` | `integer` | `Int64` | N |  | The HTTP response size in bytes. If unknown, it must be -1. |
| `headers` | `Map<string, string>` |  | N |  | The HTTP response headers. If multiple headers share the same key, theymust be merged according to HTTP spec. All header keys must belowercased, because HTTP header keys are case-insensitive. |
| `time` | `string` | `Timestamp` | N |  | The timestamp when the `destination` service sends the last byte ofthe response. |
| `backendLatency` | `string` | `Duration` | N |  | The length of time it takes the backend service to fully respond to arequest. Measured from when the destination service starts to send therequest to the backend until when the destination service receives thecomplete response from the backend. |
