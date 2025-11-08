| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `code` | `string` | `ErrorCode` | N |  | The error code |
| `message` | `string` |  | N |  | A developer-facing error message, which should be in English. Anyuser-facing error message should be localized and sent in the[google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. |
| `details` | `Array<mojo.core.Any>` |  | N |  | A list of messages that carry the error details.  There is a common set ofmessage types for APIs to use. |
