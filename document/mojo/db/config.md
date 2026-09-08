| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `driver` | `string` |  | N |  | the selected sql driver |
| `dsn` | `string` |  | N |  | DSN address, mysql://username:password@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s |
| `debug` | `boolean` |  | N |  | debug mode, default is false |
| `maxIdleConnections` | `integer` | `Int32` | N |  | max number of idle connections |
| `maxOpenConnections` | `integer` | `Int32` | N |  | max number of open connections |
| `connectionKeepAlive` | `string` | `Duration` | N |  | max number of seconds to keep alive |
| `slowLogThreshold` | `string` | `Duration` | N |  | max number of milliseconds to start log |
| `enableMetric` | `boolean` |  | N |  |
| `enableTrace` | `boolean` |  | N |  |
| `enableDetailSql` | `boolean` |  | N |  |
| `enableLogAccess` | `boolean` |  | N |  |
| `enableLogAccessRequest` | `boolean` |  | N |  |
| `enableLogAccessResponse` | `boolean` |  | N |  |
| `defaultStringSize` | `integer` | `UInt64` | N |  |
| `disableAutoMigrate` | `boolean` |  | N |  | disable auto migration for the init the database |
