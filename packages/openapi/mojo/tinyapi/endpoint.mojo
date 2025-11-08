type Endpoint {
    method: String @1
    path: String @2

    path_parameters: {String: Schema}  @10
    query_parameters: {String: Schema} @11
    header_parameters: {String: Schema} @12
    body: {String: Schema} @13

    response: {String: Schema} @20

    schemas: {String: Schema} @30
}
