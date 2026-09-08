# Mojo

## Project layout

Package declarations are consolidated in `package.mojo`. Sources live in `mojo/`;
Go, Java, Markdown, OpenAPI, and Protobuf outputs share the root `go/`, `java/`,
`document/`, `openapi/`, and `protobuf/` directories. From `go/`, use
`go run ./cmd/mojo build -t api ..` for all declared packages, or
`go run ./cmd/mojo build -t go ../mojo/core` for one package.

## Regenerating the core libraries

From `go/`, run `go run ./cmd/mojo bootstrap`, then rebuild the CLI.
Use `go run ./cmd/mojo bootstrap -t java` to update the shared Java library,
or `-t go,java` to regenerate both languages. See [the Java build guide](java/README.md).
See [the bootstrap guide](docs/bootstrap.md) for the source layout, outputs, and single-package builds.


## What can do

1. Code generator
2. Database quey language for relational, document and graph database (inspired by the RQL of the Rethink_d_b)
   1. SQL
   2. Graph_q_l
   3. [Cypher Query Language](https://en.wikipedia.org/wiki/Cypher_Query_Language)
3. Script for data process
