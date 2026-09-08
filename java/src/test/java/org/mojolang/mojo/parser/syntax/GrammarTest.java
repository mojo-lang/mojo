package org.mojolang.mojo.parser.syntax;

import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class GrammarTest {
    private MojoParser parse(String source) {
        MojoLexer lexer = new MojoLexer(CharStreams.fromString(source));
        MojoParser parser = new MojoParser(new CommonTokenStream(lexer));
        parser.removeErrorListeners();
        parser.mojoFile();
        return parser;
    }

    @Test
    void parsesCurrentIdlGrammar() {
        assertEquals(0, parse("package demo { version: '0.1.0' }\n"
                + "type Record { id: String @1\n value: Int32 @2 }\n").getNumberOfSyntaxErrors());
    }

    @Test
    void rejectsInvalidIdl() {
        assertTrue(parse("type Broken { value: }\n").getNumberOfSyntaxErrors() > 0);
    }

    @Test
    void loadsRegeneratedModelsAndRpcDescriptors() {
        assertEquals("mojo.lang.Package", org.mojolang.mojo.lang.Package.getDescriptor().getFullName());
        assertEquals("mojo.document.Document", org.mojolang.mojo.document.Document.getDescriptor().getFullName());
        assertEquals("mojo.rpc.longrunning.Operations",
                org.mojolang.mojo.rpc.longrunning.OperationsGrpc.getServiceDescriptor().getName());
    }
}
