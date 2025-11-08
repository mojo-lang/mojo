type Config {
    driver: String @1 @defualt("mysql") //< the selected sql driver
    dsn: String @2 //< DSN address, mysql://username:password@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s
    debug: Bool @3 //< debug mode, default is false
    max_idle_connections: Int32 @4 @default(10) //< max number of idle connections
    max_open_connections: Int32 @5 @default(100) //< max number of open connections

    connection_keep_alive: Duration @6 @default("300s") //< max number of seconds to keep alive
    slow_log_threshold: Duration @7 @default("500ms") //< max number of milliseconds to start log

    enable_metric: Bool @8
    enable_trace: Bool @9
    enable_detail_sql: Bool @10
    enable_log_access: Bool @11
    enable_log_access_request: Bool @12
    enable_log_access_response: Bool @13
    default_string_size UInt @14

    disable_auto_migrate: Bool @20 //< disable auto migration for the init the database
}
