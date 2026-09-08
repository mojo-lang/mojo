| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `condition` | `mojo.lang.Expression` |  | N |  |
| `executeAtLeastOnce` | `boolean` |  | N |  | A statement that is executed at least once and is re-executed each time the condition evaluates to true.equals to `repeat { statements } while( condition )` |
