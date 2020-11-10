package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.antlr.v4.runtime.tree.TerminalNode;
import org.mojolang.mojo.lang.Document;
import org.mojolang.mojo.lang.LineDocument;

import java.util.List;

public class DocumentVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    @Override
    public GeneratedMessageV3 visitDocument(MojoParser.DocumentContext ctx) {
        List<TerminalNode> lines = ctx.Line_document();
        Document.Builder builder = Document.newBuilder();

        for (TerminalNode line: lines) {
            LineDocument.Builder lineBuilder = LineDocument.newBuilder();
            lineBuilder.setLine(line.getText());
            lineBuilder.setStartPosition(Util.getPosition(line.getSymbol()));

            builder.addLines(lineBuilder);
        }

        return builder.build();
    }
}
