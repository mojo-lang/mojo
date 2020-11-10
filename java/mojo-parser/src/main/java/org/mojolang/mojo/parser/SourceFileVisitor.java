package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.*;

import java.util.List;

public class SourceFileVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private SourceFile.Builder builder_  = SourceFile.newBuilder();

    @Override
    public GeneratedMessageV3 visitSource_file(MojoParser.Source_fileContext ctx) {
        List<MojoParser.StatementContext> statements = ctx.statement();
        for (MojoParser.StatementContext statementCtx : statements) {
            MojoParser.DeclarationContext declarationCtx = statementCtx.declaration();
            if (declarationCtx != null) {
                visitDeclaration(declarationCtx);
            }
        }

        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitDeclaration(MojoParser.DeclarationContext ctx) {
        Statement.Builder statementBuilder = Statement.newBuilder();

        MojoParser.Struct_declarationContext structCtx = ctx.struct_declaration();
        if (structCtx != null) {
            StructDeclarationVisitor visitor = new StructDeclarationVisitor();
            StructDecl decl = (StructDecl) visitor.visitStruct_declaration(structCtx);

            Declaration.Builder declBuilder = Declaration.newBuilder();
            declBuilder.setStructDecl(decl);
            statementBuilder.setDeclaration(declBuilder);

            builder_.addStatements(statementBuilder);
        }

        MojoParser.Interface_declarationContext interfaceCtx = ctx.interface_declaration();
        if (interfaceCtx != null) {
            InterfaceDeclarationVisitor visitor = new InterfaceDeclarationVisitor();
            InterfaceDecl decl = (InterfaceDecl) visitor.visitInterface_declaration(interfaceCtx);

            Declaration.Builder declBuilder = Declaration.newBuilder();
            declBuilder.setInterfaceDecl(decl);
            statementBuilder.setDeclaration(declBuilder);

            builder_.addStatements(statementBuilder);
        }

        return null;
    }
}
