///
type TinyAPI {
	title: String @1
    description: String @2
    server: Url @5

    endpoints: [Endpoint] @10
    schemas: {String: Schema} @15
}





