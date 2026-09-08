# Mojo Java libraries

The eight core IDL packages share this Maven library. Generated messages,
gRPC stubs, and handwritten runtime helpers live in `src/main/java`.
Tests and their source protos live in `src/test`. The old
`packages/<module>/java` source trees have been migrated here.

## Regenerate

From the repository's `go` directory:

```sh
go run ./cmd/mojo bootstrap -t java
go run ./cmd/mojo build -t java ../packages/core
```

Use `bootstrap -t go,java` to update both language libraries. The default
bootstrap target remains Go. Java generation uses the Mojo compiler, then
Protobuf, then protoc; it emits protobuf messages and gRPC stubs. NCraft's
Spring/Feign service templates are separate from this IDL library build.

The Java default output is the repository's `java` directory. For an external
Mojo project it remains that project's `java` directory. An explicit
`build -t java -o <directory>` writes under `<directory>/src/main/java`.
Intermediate protos stay under the source package's `protobuf` directory.

Each core package records its output files in `.mojo-generated`. Keep these
manifests with the generated sources: they allow a single-package rebuild to
remove obsolete classes without removing another package's output. Handwritten
files are preserved; they must not have a generated-code header.

Core packages use the stable `org.mojolang.mojo.*` namespace, independently of
the organization names in their metadata. `mojo.core` also generates
`org.mojolang.mojo.MojoProtos` from the handwritten custom-options proto.

## ANTLR

Run from any directory:

```sh
/path/to/mojo/antlr/mojo/generate-java.sh
```

The script uses `antlr4`, falls back to `antlr`, or accepts an explicit tool jar:

```sh
ANTLR_JAR=/path/to/antlr-4.13.2-complete.jar /path/to/mojo/antlr/mojo/generate-java.sh
```

An optional output-directory argument supports isolated generation. The script
converts Go-specific predicates in temporary grammar copies and publishes the
four Java lexer/parser/visitor files only after ANTLR succeeds. It never edits
the checked-in grammars.

`java/mojo` is an older Java AST-adapter prototype. It is preserved for reference
and is not a source root of this Maven library. Its handwritten visitors use
obsolete grammar/AST APIs; the active ANTLR classes are under `java/src/main/java`.

## Build and test

Use a JDK 17 or newer supported by the configured Lombok version and Maven:

```sh
mvn -f java/pom.xml test
mvn -f java/pom.xml package
```

The checked-in generated sources use protoc 33.4 / protobuf-java 4.33.4,
ANTLR 4.13.2, and the gRPC Java plugin 1.63.0. The Maven build pins the matching
Protobuf and ANTLR runtimes, plus the runtime and test dependencies needed by
the migrated helpers. `protoc-gen-grpc-java` must be on PATH when regenerating
packages containing services; it is not required for message-only packages.

Lombok is explicitly configured as an annotation processor so newer JDKs still
generate the runtime helpers' accessors. The pinned 1.18.42 release supports
JDK 25; see the [Lombok changelog](https://projectlombok.org/changelog).

Regenerate the test fixtures from their protos with:

```sh
make -C packages/core/protobuf java-tests
```

Generator regressions and script checks:

```sh
cd go
go test ./pkg/compiler/cmd/commander ./pkg/compiler/cmd/build/java ./pkg/compiler/protobuf/converter
cd ..
python3 antlr/mojo/tests/test_generate_java.py
```
