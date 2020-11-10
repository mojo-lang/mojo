package org.mojolang.mojo.parser;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.util.JsonFormat;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTree;
import org.mojolang.mojo.lang.SourceFile;;import java.io.IOException;

public class SourceFileParser {
    public static void main(String[] args) throws InvalidProtocolBufferException {
        //String mojoContent = "type Message { a:Int8\n b:String}\n interface Service { convert(from :From) -> Result }";
        //MojoLexer lexer = new MojoLexer(CharStreams.fromString(mojoContent));

        if (args.length == 0) {
            return;
        }

        try {
            MojoLexer lexer = new MojoLexer(CharStreams.fromFileName(args[0]));
            CommonTokenStream tokens = new CommonTokenStream(lexer);
            MojoParser parser = new MojoParser(tokens);

            ParseTree tree = parser.source_file();
            SourceFileVisitor visitor = new SourceFileVisitor();
            SourceFile sourceFile = (SourceFile) visitor.visit(tree);

            //System.out.println(sourceFile.toString());

            System.out.println(
                    JsonFormat.printer()
                            .preservingProtoFieldNames()
                            //.omittingInsignificantWhitespace()
                            .print(sourceFile)
            );
        } catch (IOException e) {
            return;
        }
    }
}
