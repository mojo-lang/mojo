| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `type` | `string` |  | N |  | The type of PreconditionFailure. We recommend using a service-specificenum type to define the supported precondition violation subjects. Forexample, "TOS" for "Terms of Service violation". |
| `subject` | `string` |  | N |  | The subject, relative to the type, that failed.For example, "google.com/cloud" relative to the "TOS" type would indicatewhich terms of service is being referenced. |
| `description` | `string` |  | N |  | A description of how the precondition failed. Developers can use thisdescription to understand how to fix the failure.<br>For example: "Terms of service not accepted". |
