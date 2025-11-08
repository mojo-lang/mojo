
type FormDataContent {
    type Disposition {
        value: String @1 //< default is form-data
        name:  String @2
        file_name: String @3
    }

    disposition: Disposition @1
    type: MediaType @2
    transfer_encoding: String @3
    headers: Headers @4

    value: Value @5
}
