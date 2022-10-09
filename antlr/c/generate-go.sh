#!/bin/bash
antlr4 -Dlanguage=Go -visitor -no-listener -package syntax C.g4 -o ../../go/pkg/c/parser/syntax
