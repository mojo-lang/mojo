#!/bin/bash
antlr4 -Dlanguage=Go -visitor -no-listener -package syntax CPP14Lexer.g4 CPP14Parser.g4 -o ../../go/pkg/cpp/parser/syntax
