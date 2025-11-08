| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | Replaces the name of the element/attribute used for the described schema property. When defined within items, it will affect the name of the individual XML elements within the list. When defined alongside type being array (outside the items), it will affect the wrapping element and only if wrapped is true. If wrapped is false, it will be ignored. |
| `namespace` | `string` | `Url` | N |  | The URI of the namespace definition. Value MUST be in the form of an absolute URI. |
| `prefix` | `string` |  | N |  | The prefix to be used for the name. |
| `attribute` | `boolean` |  | N |  | Declares whether the property definition translates to an attribute instead of an element. |
| `wrapped` | `boolean` |  | N |  | MAY be used only for an array definition. Signifies whether the array is wrapped (for example, `<books><book/><book/></books>`) or unwrapped (). Default value is false. The definition takes effect only when defined alongside type being array (outside the items). |
