package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.*;


public class ValueDeclarationVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private ValueDecl.Builder builder_ = ValueDecl.newBuilder();

    @Override
    public GeneratedMessageV3 visitStruct_member_declaration(MojoParser.Struct_member_declarationContext ctx) {
        builder_.setName(ctx.Identifier().getText());
        setDocument(ctx.document(), ctx.following_document());

        if (ctx.initializer() != null) {
            visitInitializer(ctx.initializer());
        }

        if (ctx.type_annotation() != null) {
            NominalTypeVisitor visitor = new NominalTypeVisitor();
            builder_.setType((NominalType) visitor.visitType_annotation(ctx.type_annotation()));
        }

        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitEnum_member(MojoParser.Enum_memberContext ctx) {
        builder_.setName(ctx.Identifier().getText());
        setDocument(ctx.document(), ctx.following_document());

        if (ctx.initializer() != null) {
            visitInitializer(ctx.initializer());
        }

        if (ctx.attributes() != null) {
            builder_.addAllAttributes(Util.visitAttributes(ctx.attributes()));
        }

        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitFunction_parameter(MojoParser.Function_parameterContext ctx) {
        builder_.setName(ctx.label_identifier().getText());

        if (ctx.initializer() != null) {
            visitInitializer(ctx.initializer());
        }

        if (ctx.type_annotation() != null) {
            NominalTypeVisitor visitor = new NominalTypeVisitor();
            builder_.setType((NominalType) visitor.visitType_annotation(ctx.type_annotation()));
        }

        return builder_.build();
    }

    private void setDocument(MojoParser.DocumentContext ctx, MojoParser.Following_documentContext followingCtx) {
        Document.Builder builder = Document.newBuilder();
        if (ctx != null ) {
            builder.mergeFrom(Util.visitDocument(ctx));
        }

        if (followingCtx != null) {
            LineDocument.Builder lineBuilder = LineDocument.newBuilder();
            lineBuilder.setLine(followingCtx.Following_line_document().getText());
            lineBuilder.setStartPosition(Util.getPosition(followingCtx.Following_line_document().getSymbol()));

            builder.addLines(lineBuilder);
        }

        if (builder.isInitialized()) {
            builder_.setDocument(builder);
        }
    }

    @Override
    public GeneratedMessageV3 visitInitializer(MojoParser.InitializerContext ctx) {
        if (ctx != null) {
            ExpressionVisitor visitor = new ExpressionVisitor();
            Expression expression = (Expression) visitor.visitExpression(ctx.expression());
            builder_.setDefaultValue(expression);
        }
        return null;
    }
}
