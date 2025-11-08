| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `condition` | `mojo.lang.Expression` |  | N |  | an expression whose value is used as a condition |
| `thenBranch` | `mojo.lang.Expression` |  | N |  | An expression which is evaluated if the condition evaluates to a truthy value(one which equals or can be converted to true). |
| `elseBranch` | `mojo.lang.Expression` |  | N |  | An expression which is executed if the condition is falsy(that is, has a value which can be converted to false). |
