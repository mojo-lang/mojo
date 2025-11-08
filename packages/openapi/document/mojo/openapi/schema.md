| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `discriminator` | `mojo.openapi.Discriminator` |  | N |  | Adds support for polymorphism. The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description. See Composition and Inheritance for more details. |
| `xml` | `mojo.openapi.Xml` |  | N |  | This MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to describe the XML representation of this property. |
| `externalDocs` | `mojo.openapi.ExternalDocument` |  | N |  | Additional external documentation for this schema. |
| `example` | `mojo.core.Value` |  | N |  | A free-form property to include an example of an instance for this schema. To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary. |
| `title` | `string` |  | N |  |
| `description` | `mojo.openapi.CachedDocument` |  | N |  | support GFM |
| `type` | `string` |  | N |  | Value MUST be a string. Multiple types via an array are not supported. |
| `format` | `string` |  | N |  | format is an open string-valued property,and can have any value. Formats such as "Email", "Uuid", and so on,MAY be used even though undefined by this specification.Types that are not accompanied by a format property follow the type definition in the JSON Schema.Tools that do not recognize a specific format MAY default back to the type alone, as if the format is not specified. |
| `multipleOf` | `number` | `Float64` | N |  | The value of "multipleOf" MUST be a number, strictly greater than 0.<br>A numeric instance is valid only if division by this keyword's value results in an integer.<br>The `multiple_of` will be less than 1, when the numeric has a fixed decimals.For example, `multiple_of` will be `0.01`, when the numeric need to hold two decimals, and any greater precision will be rejected. |
| `maximum` | `mojo.core.Value` |  | N |  |
| `exclusiveMaximum` | `mojo.core.Value` |  | N |  |
| `minimum` | `mojo.core.Value` |  | N |  |
| `exclusiveMinimum` | `mojo.core.Value` |  | N |  |
| `maxLength` | `integer` | `UInt64` | N |  |
| `minLength` | `integer` | `UInt64` | N |  | A string instance is valid against this keyword if its length is greater than, or equal to, the value of this keyword.The length of a string instance is defined as the number of its characters as defined by RFC 7159.Omitting this keyword has the same behavior as a value of 0. |
| `pattern` | `mojo.core.Regex` |  | N |  |  |
| `items` | `mojo.openapi.ReferenceableSchema` |  | N |  | Value MUST be an object and not an array. Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema. items MUST be present if the type is array. |
| `maxItems` | `integer` | `UInt64` | N |  |
| `minItems` | `integer` | `UInt64` | N |  |
| `uniqueItems` | `boolean` |  | N |  |
| `properties` | `Map<string, mojo.openapi.ReferenceableSchema>` |  | N |  |
| `additionalProperties` | `mojo.openapi.ReferenceableSchema` |  | N |  |
| `maxProperties` | `integer` | `UInt64` | N |  |
| `minProperties` | `integer` | `UInt64` | N |  |
| `required` | `Array<string>` |  | N |  |
| `enum` | `Array<mojo.core.Value>` |  | N |  |
| `allOf` | `Array<mojo.openapi.ReferenceableSchema>` |  | N |  | Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.validates the value against all the sub-schemas |
| `oneOf` | `Array<mojo.openapi.ReferenceableSchema>` |  | N |  | validates the value against exactly one of the sub-schemas |
| `anyOf` | `Array<mojo.openapi.ReferenceableSchema>` |  | N |  | validates the value against any (one or more) of the sub-schemas |
| `not` | `mojo.openapi.ReferenceableSchema` |  | N |  | make sure the value is not valid against the specified schema. |
| `default` | `mojo.core.Value` |  | N |  |
| `nullable` | `boolean` |  | N |  | Allows sending a null value for the defined schema. Default value is false. |
| `readOnly` | `boolean` |  | N |  | Relevant only for Schema "properties" definitions. Declares the property as "read only". This means that it MAY be sent as part of a response but SHOULD NOT be sent as part of the request. If the property is marked as readOnly being true and is in the required list, the required will take effect on the response only. A property MUST NOT be marked as both readOnly and writeOnly being true. Default value is false. |
| `writeOnly` | `boolean` |  | N |  | Relevant only for Schema "properties" definitions. Declares the property as "write only". Therefore, it MAY be sent as part of a request but SHOULD NOT be sent as part of the response. If the property is marked as writeOnly being true and is in the required list, the required will take effect on the request only. A property MUST NOT be marked as both readOnly and writeOnly being true. Default value is false. |
| `deprecated` | `boolean` |  | N |  | Specifies that a schema is deprecated and SHOULD be transitioned out of usage. Default value is false. |
