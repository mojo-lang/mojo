

type MultipartFormData {
	boundary: String @1
	parts:    [FormDataContent] @2
}
