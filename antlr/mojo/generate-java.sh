#!/bin/bash

sed -i '' -e 's/\/\/{_input.get(_input.index()-1).getType()!=WS}? operator_character/{_input.get(_input.index()-1).getType()!=WS}? operator_character/' MojoParser.g4
sed -i '' -e 's/{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? operator_character/\/\/{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? operator_character/' MojoParser.g4

antlr4 -Dlanguage=Java -visitor -no-listener -package  org.mojolang.mojo.parser.syntax MojoParser.g4 MojoLexer.g4 -o ../../java/mojo/src/main/java/org/mojolang/mojo/parser/syntax

sed -i '' -e 's/{_input.get(_input.index()-1).getType()!=WS}? operator_character/\/\/{_input.get(_input.index()-1).getType()!=WS}? operator_character/' MojoParser.g4
sed -i '' -e 's/\/\/{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? operator_character/{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? operator_character/' MojoParser.g4
