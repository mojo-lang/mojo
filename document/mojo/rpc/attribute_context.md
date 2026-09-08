| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `origin` | `mojo.rpc.AttributeContext.Peer` |  | N |  | The origin of a network activity. In a multi hop network activity,the origin represents the sender of the first hop. For the first hop,the `source` and the `origin` must have the same content. |
| `source` | `mojo.rpc.AttributeContext.Peer` |  | N |  | The source of a network activity, such as starting a TCP connection.In a multi hop network activity, the source represents the sender of thelast hop. |
| `destination` | `mojo.rpc.AttributeContext.Peer` |  | N |  | The destination of a network activity, such as accepting a TCP connection.In a multi hop network activity, the destination represents the receiver ofthe last hop. |
| `request` | `mojo.rpc.AttributeContext.Request` |  | N |  | Represents a network request, such as an HTTP request. |
| `response` | `mojo.rpc.AttributeContext.Response` |  | N |  | Represents a network response, such as an HTTP response. |
| `resource` | `mojo.rpc.AttributeContext.Resource` |  | N |  | Represents a target resource that is involved with a network activity.If multiple resources are involved with an activity, this must be theprimary one. |
| `api` | `mojo.rpc.AttributeContext.Api` |  | N |  | Represents an API operation that is involved to a network activity. |
| `extensions` | `Array<mojo.core.Any>` |  | N |  | Supports extensions for advanced use cases, such as logs and metrics. |
