| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `baseAddress` | `integer` | `Int64` | N |  | The base address of this code module as it was loaded by the process. |
| `size` | `integer` | `Int64` | N |  | The size of the code module. |
| `codeFile` | `string` |  | N |  | The path or file name that the code module was loaded from. |
| `codeIdentifier` | `string` |  | N |  | An identifying string used to discriminate between multiple versions andbuilds of the same code module.  This may contain a uuid, timestamp,version number, or any combination of this or other information, in animplementation-defined format. |
| `debugFile` | `string` |  | N |  | The filename containing debugging information associated with the codemodule.  If debugging information is stored in a file separate from thecode module itself (as is the case when .pdb or .dSYM files are used),this will be different from code_file.  If debugging information isstored in the code module itself (possibly prior to stripping), thiswill be the same as code_file. |
| `debugIdentifier` | `string` |  | N |  | An identifying string similar to code_identifier, but identifies aspecific version and build of the associated debug file.  This may bethe same as code_identifier when the debug_file and code_file areidentical or when the same identifier is used to identify distinctdebug and code files. |
| `version` | `string` |  | N |  | A human-readable representation of the code module's version. |
