#!/bin/bash
antlr4 -Dlanguage=Go -visitor -no-listener -package syntax MojoParser.g4 MojoLexer.g4 -o ../../go/pkg/mojo/parser/syntax
