package data

type HTTPResponse struct {
	Body    *Field
	Headers map[string]string
}

func (r *HTTPResponse) GetBody() *Field {
	if r != nil {
		return r.Body
	}
	return nil
}
