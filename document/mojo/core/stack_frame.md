| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `instruction` | `integer` | `Int64` | Y |  | The program counter location as an absolute virtual address.  For theinnermost called frame in a stack, this will be an exact program counteror instruction pointer value.  For all other frames, this will be withinthe instruction that caused execution to branch to a called function,but may not necessarily point to the exact beginning of that instruction. |
| `module` | `mojo.core.CodeModule` |  | N |  | The module in which the instruction resides. |
| `functionName` | `string` |  | N |  | The function name, may be omitted if debug symbols are not available. |
| `functionBase` | `integer` | `Int64` | N |  | The start address of the function, may be omitted if debug symbolsare not available. |
| `sourceFileName` | `string` |  | N |  | The source file name, may be omitted if debug symbols are not available. |
| `sourceLine` | `integer` | `Int32` | N |  | The (1-based) source line number, may be omitted if debug symbols arenot available. |
| `sourceLineBase` | `integer` | `Int64` | N |  | The start address of the source line, may be omitted if debug symbolsare not available. |
| `trust` | `string` |  | N |  | Amount of trust the stack walker has in the instruction pointerof this frame. |
