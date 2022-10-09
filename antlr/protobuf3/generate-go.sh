#!/bin/bash
antlr4 -Dlanguage=Go -visitor -no-listener -package syntax Protobuf3.g4 -o ../../go/pkg/protobuf3/parser/syntax
