#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
repo_dir="$(cd -- "$script_dir/../.." && pwd)"
output_dir="${1:-$repo_dir/java/src/main/java/org/mojolang/mojo/parser/syntax}"
if (( $# > 1 )); then
    echo "Usage: $0 [output-directory]" >&2
    exit 2
fi
mkdir -p "$output_dir"
output_dir="$(cd -- "$output_dir" && pwd)"

if [[ -n "${ANTLR_JAR:-}" ]]; then
    antlr_command=(java -jar "$ANTLR_JAR")
elif command -v antlr4 >/dev/null 2>&1; then
    antlr_command=(antlr4)
elif command -v antlr >/dev/null 2>&1; then
    antlr_command=(antlr)
else
    echo "ANTLR 4 is required: install antlr4/antlr or set ANTLR_JAR." >&2
    exit 1
fi

work_dir="$(mktemp -d "${TMPDIR:-/tmp}/mojo-antlr-java.XXXXXX")"
trap 'rm -rf "$work_dir"' EXIT
cp "$script_dir/MojoLexer.g4" "$work_dir/MojoLexer.g4"
# Translate the Go-specific whitespace predicate in a temporary grammar.
# Leave the checked-in grammar unchanged, including when ANTLR fails.
sed 's/p\.GetTokenStream()\.Get(p\.GetTokenStream()\.Index()-1)\.GetTokenType() != MojoParserWS/_input.index() > 0 \&\& _input.get(_input.index()-1).getType() != WS/g' \
    "$script_dir/MojoParser.g4" > "$work_dir/MojoParser.g4"
(
    cd "$work_dir"
    "${antlr_command[@]}" -Dlanguage=Java -visitor -no-listener \
        -package org.mojolang.mojo.parser.syntax -Xexact-output-dir \
        -o "$work_dir/generated" MojoLexer.g4 MojoParser.g4
)
for name in MojoLexer MojoParser MojoParserBaseVisitor MojoParserVisitor; do
    cp "$work_dir/generated/$name.java" "$output_dir/$name.java"
done
