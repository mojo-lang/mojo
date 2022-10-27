#!/bin/bash

sed -i '' -e 's/{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}/{_input.get(_input.index()-1).getType()!=WS}/' MojoParser.g4

antlr4 -Dlanguage=Java -visitor -no-listener -package  org.mojolang.mojo.parser.syntax MojoParser.g4 MojoLexer.g4 -o ../../java/mojo/src/main/java/org/mojolang/mojo/parser/syntax

sed -i '' -e 's/{_input.get(_input.index()-1).getType()!=WS}/{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}/' MojoParser.g4