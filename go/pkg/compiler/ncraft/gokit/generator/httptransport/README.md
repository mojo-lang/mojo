# HTTP path patterns

HTTP bindings accept Gorilla mux's `{parameter:pattern}` syntax:

```mojo
@http.get("/articles/{category}/{id:[0-9]{2,4}}")
get_article(category: String @1, id: Int32 @2) -> Article

@http.get("/files/{name:.+}")
@http.head("/files/{name:.+}")
get_file(name: String @1) -> BinaryFile
```

`{name}` continues to match one path segment. `{name:.+}` matches a non-empty
name that may include `/`. Storage handlers must still validate object names.
Patterns follow Gorilla mux's Go regular expression rules, including its
requirement to use non-capturing groups such as `(?:images|docs)`.

The compiler separates the parameter name from its pattern and recognizes
balanced braces in quantifiers such as `{2,4}`. A nested field name such as
`{id.level:[0-9]+}` becomes `{id_level:[0-9]+}` in the server route; the regex
itself is preserved.

- Generated server routes retain the patterns and quote them as Go strings.
- Request decoders bind only the parameter name, without the regex suffix.
- OpenAPI paths and HTTP client templates use plain `{parameter}` placeholders.
  Clients substitute URL-escaped values; they do not send pattern text in URLs.

Regenerate HTTP transports after changing route annotations. Application-specific
paths belong in the Mojo interface, not in generated Go files.
