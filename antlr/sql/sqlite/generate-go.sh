#!/bin/bash
antlr4 -Dlanguage=Go -visitor -no-listener -package syntax SQLiteParser.g4 SQLiteLexer.g4 -o ../../../go/pkg/sqlite/parser/syntax
