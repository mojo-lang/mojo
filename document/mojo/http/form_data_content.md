| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `disposition` | `mojo.http.FormDataContent.Disposition` |  | N |  |  |
| `type` | `mojo.core.MediaType` |  | N |  | A media type (also known as a Multipurpose Internet Mail Extensions or MIME type) is a standardthat indicates the nature and format of a document, file, or assortment of bytes.It is defined and standardized in IETF's .<br>The simplest MIME type consists of a type and a subtype;these are each strings which, when concatenated with a slash (/) between them, comprise a MIME type.No whitespace is allowed in a MIME type: |
| `transferEncoding` | `string` |  | N |  |
| `headers` | `Map<string, Array<string>>` |  | N |  |  |
| `value` | `mojo.core.Value` |  | N |  |
