| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `ip` | `string` |  | N |  | The IP address of the peer. |
| `port` | `integer` | `Int64` | N |  | The network port of the peer. |
| `labels` | `Map<string, string>` |  | N |  | The labels associated with the peer. |
| `principal` | `string` |  | N |  | The identity of this peer. Similar to `Request.auth.principal`, butrelative to the peer instead of the request. For example, theidenity associated with a load balancer that forwared the request. |
| `regionCode` | `string` |  | N |  | The CLDR country/region code associated with the above IP address.If the IP address is private, the `region_code` should reflect thephysical location where this peer is running. |
