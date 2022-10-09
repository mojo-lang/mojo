#!/bin/bash
antlr4 -Dlanguage=Go -visitor -no-listener -package syntax Protobuf2.g4 -o ../../go/pkg/protobuf2/parser/syntax
